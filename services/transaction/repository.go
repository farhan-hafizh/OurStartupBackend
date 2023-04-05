package transaction

import "gorm.io/gorm"

type Repository interface {
	GetAllByCampaignId(campaignId int) ([]Transaction, error)
	GetAllNotSecretByCampaignId(campaignId int) ([]Transaction, error)
	Save(trans Transaction) (Transaction, error)
	GetByUserId(userId int) ([]Transaction, error)
}

type repository struct {
	db *gorm.DB
}

func CreateRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// for admin or campaign owner
func (r *repository) GetAllByCampaignId(campaignId int) ([]Transaction, error) {

	var transactions []Transaction
	err := r.db.Where("campaign_id = ?", campaignId).Preload("User").Order("id desc").Find(&transactions).Error

	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (r *repository) GetAllNotSecretByCampaignId(campaignId int) ([]Transaction, error) {

	var transactions []Transaction
	err := r.db.Where("campaign_id = ? AND is_secret = ?", campaignId, false).Preload("User").Order("id desc").Find(&transactions).Error

	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (r *repository) Save(trans Transaction) (Transaction, error) {

	err := r.db.Create(&trans).Error

	if err != nil {
		return trans, err
	}

	return trans, nil
}

// get all logged in user transaction
func (r *repository) GetByUserId(userId int) ([]Transaction, error) {
	var transactions []Transaction
	// join campaign images by table campaign
	err := r.db.Where("user_id = ?", userId).Preload("Campaign.CampaignImages", "campaign_images.is_primary = 1").Order("id desc").Find(&transactions).Error

	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
