package campaign

import (
	"errors"
	"fmt"
	"time"

	"github.com/gosimple/slug"
)

type Service interface {
	CreateCampaign(userId int, input CreateCampaignInput) (Campaign, error)
	GetCampaigns(userId int) ([]Campaign, error)
	GetCampaignBySlug(input GetCampaignSlugInput) (Campaign, error)
	UpdateCampaign(input GetCampaignSlugInput, campaignData CreateCampaignInput) (Campaign, error)
	SaveCampaignImage(input CreateCampaignImageInput, fileLocation string) (CampaignImage, error)
}

type service struct {
	repository Repository
}

func CreateService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetCampaigns(userId int) ([]Campaign, error) {
	if userId != 0 {
		campaign, err := s.repository.FindByCreatorId(userId)

		if err != nil {
			return campaign, err
		}

		return campaign, nil
	}

	campaigns, err := s.repository.FindAll()

	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (s *service) CreateCampaign(userId int, input CreateCampaignInput) (Campaign, error) {
	campaign := Campaign{}

	campaign.CreatorId = userId
	campaign.Name = input.Name
	campaign.ShortDescription = input.ShortDescription
	campaign.Description = input.Description
	campaign.BackerCount = 0
	campaign.GoalAmount = input.GoalAmount
	campaign.CurrentAmount = 0
	campaign.Perks = input.Perks
	campaign.Slug = slug.Make(fmt.Sprintf("%s %d", input.Name, time.Now().Unix()))
	newCampaign, err := s.repository.Save(campaign)

	if err != nil {
		return newCampaign, err
	}

	return newCampaign, nil
}

func (s *service) GetCampaignBySlug(input GetCampaignSlugInput) (Campaign, error) {
	campaign, err := s.repository.FindBySlug(input.Slug)

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (s *service) UpdateCampaign(input GetCampaignSlugInput, campaignData CreateCampaignInput) (Campaign, error) {
	campaign, err := s.repository.FindBySlug(input.Slug)

	if err != nil {
		return campaign, err
	}
	// check if the current loggedin user is the campaign owner
	if campaign.User.Id != input.User.Id {
		return campaign, errors.New("Invalid campaign owner!")
	}

	campaign.Name = campaignData.Name
	campaign.ShortDescription = campaignData.ShortDescription
	campaign.Description = campaignData.Description
	campaign.Perks = campaignData.Perks
	campaign.GoalAmount = campaignData.GoalAmount

	updatedCampaign, err := s.repository.Update(campaign)

	if err != nil {
		return updatedCampaign, err
	}

	return updatedCampaign, nil
}

func (s *service) SaveCampaignImage(input CreateCampaignImageInput, fileLocation string) (CampaignImage, error) {

	campaign, err := s.repository.FindBySlug(input.Slug)

	if err != nil {
		return CampaignImage{}, err
	}

	if campaign.User.Id != input.User.Id {
		return CampaignImage{}, errors.New("Invalid campaign owner!")
	}

	if input.IsPrimary {
		_, err := s.repository.ChangeImageIsPrimary(campaign.Id)

		if err != nil {
			return CampaignImage{}, err
		}
	}

	image := CampaignImage{
		CampaignId: campaign.Id,
		FileName:   fileLocation,
		IsPrimary:  input.IsPrimary,
	}

	updatedImage, err := s.repository.SaveImage(image)

	if err != nil {
		return updatedImage, err
	}

	return updatedImage, nil
}
