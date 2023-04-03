package campaign

import "ourstartup/services/user"

type CreateCampaignInput struct {
	Name             string `json:"name" binding:"required"`
	ShortDescription string `json:"short_description" binding:"required"`
	Description      string `json:"description" binding:"required"`
	GoalAmount       int    `json:"goal_amount" binding:"required"`
	Perks            string `json:"perks" binding:"required"`
}

type GetCampaignSlugInput struct {
	Slug string `uri:"slug" binding:"required`
	User user.User
}

type CreateCampaignImageInput struct {
	Slug      string `form:"slug" binding:"required"`
	IsPrimary bool   `form:"is_primary" binding:"required"`
	User      user.User
}
