package payment

import (
	"ourstartup/config"
	"ourstartup/entities"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type service struct {
	config config.Config
}

type Service interface {
	GetRedirectUrl(transaction entities.Transaction) (string, error)
}

func CreateService(config config.Config) *service {
	return &service{config}
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
