package campaign

import "errors"

type Service interface {
	GetCampaigns(userID int) ([]Campaign, error)
	GetCampaign(input GetCampaignDetailInput) (Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetCampaigns(userID int) ([]Campaign, error) {
	var campaigns []Campaign
	var err error

	if userID != 0 {
		campaigns, err = s.repository.FindByUserID(userID)
	} else {
		campaigns, err = s.repository.FindAll()
	}

	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (s *service) GetCampaign(input GetCampaignDetailInput) (Campaign, error) {
	campaign, err := s.repository.FindByID(input.ID)
	if err != nil {
		return campaign, err
	}

	if campaign.ID == 0 {
		return campaign, errors.New("no campaign found on that ID")
	}

	return campaign, nil

}
