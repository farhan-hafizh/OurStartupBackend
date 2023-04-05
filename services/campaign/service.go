package campaign

import (
	"errors"
	"fmt"
	"ourstartup/entities"
	"time"

	"github.com/gosimple/slug"
)

type Service interface {
	CreateCampaign(input CreateCampaignInput) (entities.Campaign, error)
	GetCampaigns(userId int) ([]entities.Campaign, error)
	GetCampaignBySlug(input GetCampaignInput) (entities.Campaign, error)
	UpdateCampaign(input GetCampaignInput, campaignData CreateCampaignInput) (entities.Campaign, error)
	SaveCampaignImage(input CreateCampaignImageInput, fileLocation string) (entities.CampaignImage, error)
	FindCampaignById(input GetCampaignInput) (entities.Campaign, error)
}

type service struct {
	repository Repository
}

func CreateService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindCampaignById(input GetCampaignInput) (entities.Campaign, error) {
	campaign, err := s.repository.FindById(input.Id)

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (s *service) GetCampaigns(userId int) ([]entities.Campaign, error) {
	if userId != 0 { // if user id exist get campaign created by user id
		campaign, err := s.repository.FindByCreatorId(userId)

		if err != nil {
			return campaign, err
		}

		return campaign, nil
	}
	// else get all campaign
	campaigns, err := s.repository.FindAll()

	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (s *service) CreateCampaign(input CreateCampaignInput) (entities.Campaign, error) {
	campaign := entities.Campaign{
		CreatorId:        input.User.Id,
		Name:             input.Name,
		ShortDescription: input.ShortDescription,
		Description:      input.Description,
		BackerCount:      0,
		GoalAmount:       input.GoalAmount,
		CurrentAmount:    0,
		Perks:            input.Perks,
		User:             input.User,
		Slug:             slug.Make(fmt.Sprintf("%s %d", input.Name, time.Now().Unix())),
	}

	newCampaign, err := s.repository.Save(campaign)

	if err != nil {
		return newCampaign, err
	}

	return newCampaign, nil
}

func (s *service) GetCampaignBySlug(input GetCampaignInput) (entities.Campaign, error) {
	campaign, err := s.repository.FindBySlug(input.Slug)

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (s *service) UpdateCampaign(input GetCampaignInput, campaignData CreateCampaignInput) (entities.Campaign, error) {

	var campaign entities.Campaign
	var err error

	if input.Id != 0 { // if there is the id use id if not use slug
		campaign, err = s.repository.FindById(input.Id)
	} else {
		campaign, err = s.repository.FindBySlug(input.Slug)
	}

	if err != nil {
		return campaign, err
	}
	// check if the current loggedin user is the campaign owner
	if campaign.User.Id != input.User.Id {
		return campaign, errors.New("Invalid campaign owner!")
	}

	if input.IsPaymentSucces { // from payment success
		campaign.BackerCount = campaignData.BackerCount
		campaign.CurrentAmount = campaignData.CurrentAmount
	} else { // update campaign info
		campaign.Name = campaignData.Name
		campaign.ShortDescription = campaignData.ShortDescription
		campaign.Description = campaignData.Description
		campaign.Perks = campaignData.Perks
		campaign.GoalAmount = campaignData.GoalAmount
	}

	updatedCampaign, err := s.repository.Update(campaign)

	if err != nil {
		return updatedCampaign, err
	}

	return updatedCampaign, nil
}

func (s *service) SaveCampaignImage(input CreateCampaignImageInput, fileLocation string) (entities.CampaignImage, error) {

	campaign, err := s.repository.FindBySlug(input.Slug)

	if err != nil {
		return entities.CampaignImage{}, err
	}

	if campaign.User.Id != input.User.Id {
		return entities.CampaignImage{}, errors.New("Invalid campaign owner!")
	}

	if input.IsPrimary {
		_, err := s.repository.ChangeImageIsPrimary(campaign.Id)

		if err != nil {
			return entities.CampaignImage{}, err
		}
	}

	image := entities.CampaignImage{
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
