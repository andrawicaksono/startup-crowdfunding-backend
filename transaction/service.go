package transaction

import (
	"errors"
	"startup-crowdfunding-backend/campaign"
	"startup-crowdfunding-backend/helper"
)

type Service interface {
	GetTransactionsByCampaignID(input GetTransactionsByCampaignIDInput) ([]Transaction, error)
	GetTransactionsByUserID(userID int) ([]Transaction, error)
	CreateTransaction(input CreateTransactionInput) (Transaction, error)
}

type service struct {
	repository         Repository
	campaignRepository campaign.Repository
}

func NewService(repository Repository, campaignRepository campaign.Repository) *service {
	return &service{repository, campaignRepository}
}

func (s *service) GetTransactionsByCampaignID(input GetTransactionsByCampaignIDInput) ([]Transaction, error) {
	transactions := []Transaction{}

	campaign, err := s.campaignRepository.FindByID(input.CampaignID)
	if err != nil {
		return transactions, err
	}

	if input.CampaignID != campaign.ID {
		return transactions, errors.New("not an owner of the campaign")
	}

	transactions, err = s.repository.FindByCampaignID(input.CampaignID)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (s *service) GetTransactionsByUserID(userID int) ([]Transaction, error) {
	transactions, err := s.repository.FindByUserID(userID)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (s *service) CreateTransaction(input CreateTransactionInput) (Transaction, error) {
	transaction := Transaction{}
	campaign, err := s.campaignRepository.FindByID(input.CampaignID)
	if err != nil {
		return transaction, err
	}

	if campaign.ID == 0 {
		return transaction, errors.New("no campaign found on that ID")
	}

	transaction.UserID = input.User.ID
	transaction.CampaignID = campaign.ID
	transaction.Amount = input.Amount
	transaction.Status = "pending"
	transaction.Code = helper.GenerateTransactionCode(6)

	newTransaction, err := s.repository.Save(transaction)
	if err != nil {
		return transaction, err
	}

	return newTransaction, nil

}
