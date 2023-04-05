package entities

import (
	"time"
)

type Transaction struct {
	Id         int       `json:"id"`
	Code       string    `json:"code"`
	CampaignId int       `json:"campaign_id"`
	UserId     int       `json:"user_id"`
	Amount     int       `json:"amount"`
	IsSecret   bool      `json:"is_secret"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	User       User
	Campaign   Campaign
}
