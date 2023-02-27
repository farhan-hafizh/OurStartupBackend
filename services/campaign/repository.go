package campaign

import "gorm.io/gorm"

type Repository interface {
	Save(campaign Campaign) (Campaign, error)
	FindById(id int) (Campaign, error)
	FindAll() ([]Campaign, error)
	FindByCreatorId(id int) ([]Campaign, error)
	// FindCampaignImage(id int, isPrimary bool) ([]CampaignImage, error)
}

type repository struct {
	db *gorm.DB
}

func CreateRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// find campaign by id
func (r *repository) FindById(id int) (Campaign, error) {
	var campaign Campaign

	err := r.db.Where("id = ?", id).Find(&campaign).Error

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

// create campaign
func (r *repository) Save(campaign Campaign) (Campaign, error) {
	err := r.db.Create(&campaign).Error

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

// find all campaign
func (r *repository) FindAll() ([]Campaign, error) {
	var campaigns []Campaign

	err := r.db.Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error

	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (r *repository) FindByCreatorId(id int) ([]Campaign, error) {
	var campaigns []Campaign
	// pre load get campaign images befor getting the campaign with is primary = true and save it to CampaignImages
	err := r.db.Where("creator_id = ?", id).Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error

	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

// func (r *repository) FindCampaignImage(id int)
