package transaction

import (
	"ourstartup/entities"
)

type CreateTransactionInput struct {
	CampaignSlug string `json:"campaign_slug" binding:"required"`
	Amount       int    `json:"amount" binding:"required"`
	IsSecret     bool   `json:"is_secret" binding:"required"`
	CallerId     string `json:"caller_id"`
	Campaign     entities.Campaign
	User         entities.User
}

type GetTransByCampaignId struct {
	Campaign   entities.Campaign
	User       entities.User
	IsAllTrans bool
}

type GetTransactionHistory struct {
	User entities.User
}

type GetTransactionByCode struct {
	Code string
}
