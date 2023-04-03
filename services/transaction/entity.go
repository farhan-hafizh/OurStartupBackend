package transaction

import (
	"os/user"
	"ourstartup/services/campaign"
	"time"
)

type Transaction struct {
	Id         int               `json:"id"`
	CampaignId int               `json:"campaign_id"`
	Campaign   campaign.Campaign `gorm:"foreignKey:CampaignId"`
	UserId     int               `json:"user_id"`
	User       user.User         `gorm:"foreignKey:UserId"` //refer CreatorId as user's foreign key
	Amount     int               `json:"amount"`
	Status     string            `json:"status"`
	CreatedAt  time.Time         `json:"createdAt"`
	UpdatedAt  time.Time         `json:"updatedAt"`
}
