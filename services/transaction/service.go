package transaction

import (
	"errors"
	"fmt"
	"time"
)

type Service interface {
	GetTransByCampaignId(input GetTransByCampaignId) ([]Transaction, error)
	CreateTransaction(input CreateTransactionInput) (Transaction, error)
	GetTransactionHistory(input GetTransactionHistory) ([]Transaction, error)
}

type service struct {
	repository Repository
}

func CreateService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetTransByCampaignId(input GetTransByCampaignId) ([]Transaction, error) {
	var transactions []Transaction
	var err error

	fmt.Println(input.IsAllTrans)

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

func (s *service) CreateTransaction(input CreateTransactionInput) (Transaction, error) {

	trans := Transaction{
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

	return newTransaction, nil
}

func (s *service) GetTransactionHistory(input GetTransactionHistory) ([]Transaction, error) {

	transactions, err := s.repository.GetByUserId(input.User.Id)

	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
