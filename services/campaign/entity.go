package campaign

import "time"

type Campaign struct {
	Id               int             `json:"id"`
	CreatorId        int             `json:"creator_id"`
	Name             string          `json:"name"`
	ShortDescription string          `json:"short_description"`
	Description      string          `json:"description"`
	BackerCount      int             `json:"backer_count"`
	GoalAmount       int             `json:"goal_amount"`
	CurrentAmount    int             `json:"current_amount"`
	Perks            string          `json:"perks"`
	Slug             string          `json:"slug"`
	CreatedAt        time.Time       `json:"createdAt"`
	UpdatedAt        time.Time       `json:"updatedAt"`
	CampaignImages   []CampaignImage `json:"campaign_image"`
}

type CampaignImage struct {
	ID         int       `json:"id"`
	CampaignId int       `json:"campaign_id"`
	FileName   string    `json:"file_name"`
	IsPrimary  bool      `json:"is_primary"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
