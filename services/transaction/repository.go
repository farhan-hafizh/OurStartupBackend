package transaction

import (
	"ourstartup/entities"

	"gorm.io/gorm"
)

type Repository interface {
	GetAllByCampaignId(campaignId int) ([]entities.Transaction, error)
	GetAllNotSecretByCampaignId(campaignId int) ([]entities.Transaction, error)
	Save(trans entities.Transaction) (entities.Transaction, error)
	GetByUserId(userId int) ([]entities.Transaction, error)
}

type repository struct {
	db *gorm.DB
}

func CreateRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// for admin or campaign owner
func (r *repository) GetAllByCampaignId(campaignId int) ([]entities.Transaction, error) {

	var transactions []entities.Transaction
	err := r.db.Where("campaign_id = ?", campaignId).Preload("User").Order("id desc").Find(&transactions).Error

	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (r *repository) GetAllNotSecretByCampaignId(campaignId int) ([]entities.Transaction, error) {

	var transactions []entities.Transaction
	err := r.db.Where("campaign_id = ? AND is_secret = ?", campaignId, false).Preload("User").Order("id desc").Find(&transactions).Error

	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (r *repository) Save(trans entities.Transaction) (entities.Transaction, error) {

	err := r.db.Create(&trans).Error

	if err != nil {
		return trans, err
	}

	return trans, nil
}

// get all logged in user transaction
func (r *repository) GetByUserId(userId int) ([]entities.Transaction, error) {
	var transactions []entities.Transaction
	// join campaign images by table campaign
	err := r.db.Where("user_id = ?", userId).Preload("Campaign.CampaignImages", "campaign_images.is_primary = 1").Order("id desc").Find(&transactions).Error

	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
