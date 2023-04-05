package transaction

import (
	"ourstartup/services/campaign"
	"ourstartup/services/user"
)

type CreateTransactionInput struct {
	CampaignSlug string `json:"campaign_slug" binding:"required"`
	Amount       int    `json:"amount" binding:"required"`
	IsSecret     bool   `json:"is_secret" binding:"required"`
	CallerId     string `json:"caller_id"`
	Campaign     campaign.Campaign
	User         user.User
}

type GetTransByCampaignId struct {
	Campaign   campaign.Campaign
	User       user.User
	IsAllTrans bool
}

type GetTransactionHistory struct {
	User user.User
}
