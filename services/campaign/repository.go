package campaign

import "gorm.io/gorm"

type Repository interface {
	Save(campaign Campaign) (Campaign, error)
	FindBySlug(slug string) (Campaign, error)
	FindAll() ([]Campaign, error)
	FindByCreatorId(id int) ([]Campaign, error)
	Update(campaign Campaign) (Campaign, error)
	SaveImage(image CampaignImage) (CampaignImage, error)
	ChangeImageIsPrimary(campaignId int) (bool, error)
}

type repository struct {
	db *gorm.DB
}

func CreateRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// find campaign by id
func (r *repository) FindBySlug(slug string) (Campaign, error) {
	var campaign Campaign

	err := r.db.Where("slug = ?", slug).Preload("User").Preload("CampaignImages").Find(&campaign).Error

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

	err := r.db.Preload("User").Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error

	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (r *repository) FindByCreatorId(id int) ([]Campaign, error) {
	var campaigns []Campaign
	// pre load get campaign images befor getting the campaign with is primary = true and save it to CampaignImages
	err := r.db.Where("creator_id = ?", id).Preload("User").Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error

	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (r *repository) Update(campaign Campaign) (Campaign, error) {

	err := r.db.Save(&campaign).Error

	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

func (r *repository) SaveImage(image CampaignImage) (CampaignImage, error) {
	err := r.db.Create(&image).Error

	if err != nil {
		return image, err
	}
	return image, nil
}

func (r *repository) ChangeImageIsPrimary(campaignId int) (bool, error) {
	// UPDATE campaign_images SET is_primary = 0 WHERE id = campaignId AND is_primary=1
	err := r.db.Model(&CampaignImage{}).Where("campaign_id = ? AND is_primary = ?", campaignId, true).Update("is_primary", false).Error

	if err != nil {
		return false, err
	}

	return true, nil
}
