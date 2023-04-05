package campaign

import (
	"ourstartup/entities"

	"gorm.io/gorm"
)

type Repository interface {
	Save(campaign entities.Campaign) (entities.Campaign, error)
	FindBySlug(slug string) (entities.Campaign, error)
	FindAll() ([]entities.Campaign, error)
	FindByCreatorId(id int) ([]entities.Campaign, error)
	Update(campaign entities.Campaign) (entities.Campaign, error)
	SaveImage(image entities.CampaignImage) (entities.CampaignImage, error)
	ChangeImageIsPrimary(campaignId int) (bool, error)
}

type repository struct {
	db *gorm.DB
}

func CreateRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// find campaign by id
func (r *repository) FindBySlug(slug string) (entities.Campaign, error) {
	var campaign entities.Campaign

	err := r.db.Where("slug = ?", slug).Preload("User").Preload("CampaignImages").Find(&campaign).Error

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

// create campaign
func (r *repository) Save(campaign entities.Campaign) (entities.Campaign, error) {
	err := r.db.Create(&campaign).Error

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

// find all campaign
func (r *repository) FindAll() ([]entities.Campaign, error) {
	var campaigns []entities.Campaign

	err := r.db.Preload("User").Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error

	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (r *repository) FindByCreatorId(id int) ([]entities.Campaign, error) {
	var campaigns []entities.Campaign
	// pre load get campaign images befor getting the campaign with is primary = true and save it to CampaignImages
	err := r.db.Where("creator_id = ?", id).Preload("User").Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error

	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (r *repository) Update(campaign entities.Campaign) (entities.Campaign, error) {

	err := r.db.Save(&campaign).Error

	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

func (r *repository) SaveImage(image entities.CampaignImage) (entities.CampaignImage, error) {
	err := r.db.Create(&image).Error

	if err != nil {
		return image, err
	}
	return image, nil
}

func (r *repository) ChangeImageIsPrimary(campaignId int) (bool, error) {
	// UPDATE campaign_images SET is_primary = 0 WHERE id = campaignId AND is_primary=1
	err := r.db.Model(&entities.CampaignImage{}).Where("campaign_id = ? AND is_primary = ?", campaignId, true).Update("is_primary", false).Error

	if err != nil {
		return false, err
	}

	return true, nil
}
