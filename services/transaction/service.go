package transaction

import (
	"errors"
	"fmt"
	"ourstartup/entities"
	"ourstartup/helper/payment"
	"time"
)

type Service interface {
	GetTransByCampaignId(input GetTransByCampaignId) ([]entities.Transaction, error)
	CreateTransaction(input CreateTransactionInput) (entities.Transaction, error)
	GetTransactionHistory(input GetTransactionHistory) ([]entities.Transaction, error)
}

type service struct {
	repository     Repository
	paymentService payment.Service
}

func CreateService(repository Repository, paymentService payment.Service) *service {
	return &service{repository, paymentService}
}

func (s *service) GetTransByCampaignId(input GetTransByCampaignId) ([]entities.Transaction, error) {
	var transactions []entities.Transaction
	var err error

	if input.IsAllTrans {
		// if not the owner
		if input.Campaign.User.Id != input.User.Id {
			return transactions, errors.New("Invalid campaign owner!")
		}

		transactions, err = s.repository.GetAllByCampaignId(input.Campaign.Id)
	} else {
		transactions, err = s.repository.GetAllNotSecretByCampaignId(input.Campaign.Id)
	}

	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (s *service) CreateTransaction(input CreateTransactionInput) (entities.Transaction, error) {

	trans := entities.Transaction{
		CampaignId: input.Campaign.Id,
		UserId:     input.User.Id,
		Amount:     input.Amount,
		IsSecret:   input.IsSecret,
		Status:     "pending",
		User:       input.User,
		Code:       fmt.Sprintf("Transaction-%d%d%d", input.User.Id, input.Campaign.Id, time.Now().Unix()),
	}

	newTransaction, err := s.repository.Save(trans)

	if err != nil {
		return newTransaction, err
	}
	paymentUrl, err := s.paymentService.GetRedirectUrl(newTransaction)

	newTransaction.PaymentUrl = paymentUrl

	if err != nil {
		return newTransaction, err
	}

	newTransaction, err = s.repository.Update(newTransaction)

	if err != nil {
		return newTransaction, err
	}

	return newTransaction, nil
}

func (s *service) GetTransactionHistory(input GetTransactionHistory) ([]entities.Transaction, error) {

	transactions, err := s.repository.GetByUserId(input.User.Id)

	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
