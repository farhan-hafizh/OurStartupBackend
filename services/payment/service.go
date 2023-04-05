package payment

import (
	"ourstartup/config"
	"ourstartup/entities"
	"ourstartup/services/campaign"
	"ourstartup/services/transaction"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type service struct {
	config          config.Config
	transService    transaction.Service
	campaignService campaign.Service
}

type Service interface {
	GetRedirectUrl(transaction entities.Transaction) (string, error)
	ProcessPayment(input TransactionNotification) error
}

func CreateService(config config.Config, transService transaction.Service, campaignService campaign.Service) *service {
	return &service{config, transService, campaignService}
}

func (s *service) GetRedirectUrl(transaction entities.Transaction) (string, error) {
	var midclient snap.Client
	midclient.New(s.config.MidServer, midtrans.Sandbox)

	// Initiate Snap Request
	snapReq := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  transaction.Code,
			GrossAmt: int64(transaction.Amount),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: transaction.User.Name,
			Email: transaction.User.Email,
		},
	}

	snapRes, err := midclient.CreateTransaction(snapReq)

	if err != nil {
		return "", err
	}

	return snapRes.RedirectURL, nil
}

func (s *service) ProcessPayment(input TransactionNotification) error {

	inputTrans := transaction.GetTransactionByCode{
		Code: input.OrderId,
	}

	transaction, err := s.transService.GetTransByTransactionCode(inputTrans)

	if err != nil {
		return err
	}
	// MIDTRANS DOC
	if input.PaymentType == "credit_card" && input.TransactionStatus == "capture" && input.FraudStatus == "accept" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "settlement" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "deny" || input.TransactionStatus == "expire" || input.TransactionStatus == "cancel" {
		transaction.Status = "cancelled"
	}

	updatedTransaction, err := s.transService.UpdateTransaction(transaction)
	if err != nil {
		return err
	}

	if updatedTransaction.Status != "paid" {
		return nil
	}

	inputCampaign := campaign.GetCampaignInput{
		Id: updatedTransaction.CampaignId,
	}
	campaignData, err := s.campaignService.FindCampaignById(inputCampaign)
	if err != nil {
		return err
	}

	campaignData.BackerCount += 1
	campaignData.CurrentAmount += updatedTransaction.Amount

	campaignInput := campaign.GetCampaignInput{
		Id:              campaignData.Id,
		IsPaymentSucces: true,
	}
	campaignUpdate := campaign.CreateCampaignInput{
		BackerCount:   campaignData.BackerCount,
		CurrentAmount: campaignData.CurrentAmount,
	}

	_, err = s.campaignService.UpdateCampaign(campaignInput, campaignUpdate)
	return nil
}
