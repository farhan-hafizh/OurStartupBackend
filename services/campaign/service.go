package campaign

type Service interface {
	CreateCampaign(userId int, input CreateCampaignInput) (Campaign, error)
	GetCampaigns(userId int) ([]Campaign, error)
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
	campaign.Slug = "slug"

	newCampaign, err := s.repository.Save(campaign)

	if err != nil {
		return newCampaign, err
	}

	return newCampaign, nil
}
