package campaign

import "ourstartup/entities"

type CreateCampaignInput struct {
	Name             string `json:"name" binding:"required"`
	ShortDescription string `json:"short_description" binding:"required"`
	Description      string `json:"description" binding:"required"`
	GoalAmount       int    `json:"goal_amount" binding:"required"`
	Perks            string `json:"perks" binding:"required"`
	User             entities.User
}

type GetCampaignSlugInput struct {
	Slug          string `uri:"slug" binding:"required"`
	CampaignOwner string `uri:"campaignOwner"` // for transaction
	User          entities.User
}

type CreateCampaignImageInput struct {
	Slug      string `form:"slug" binding:"required"`
	IsPrimary bool   `form:"is_primary"` // not required because by default is false
	User      entities.User
}
