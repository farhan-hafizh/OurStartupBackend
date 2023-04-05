package transaction

import (
	"ourstartup/services/campaign"
	"ourstartup/services/user"
	"time"
)

type Transaction struct {
	Id         int       `json:"id"`
	CallerId   string    `json:"caller_id"`
	CampaignId int       `json:"campaign_id"`
	UserId     int       `json:"user_id"`
	Amount     int       `json:"amount"`
	IsSecret   bool      `json:"is_secret"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	User       user.User
	Campaign   campaign.Campaign
}
