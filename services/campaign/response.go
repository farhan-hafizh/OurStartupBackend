package campaign

import "ourstartup/services/user"

type CampaignResponse struct {
	Id               int                    `json:"id"`
	Name             string                 `json:"name"`
	Creator          user.UserBasicResponse `json:"creator"`
	ShortDescription string                 `json:"short_description"`
	ImageUrl         string                 `json:"image_url"`
	GoalAmount       int                    `json:"goal_amount"`
	CurrentAmount    int                    `json:"current_amount"`
}

func FormatCampaignResponse(campaign Campaign) CampaignResponse {
	imageUrl := ""

	if len(campaign.CampaignImages) > 0 {
		imageUrl = campaign.CampaignImages[0].FileName
	}

	return CampaignResponse{
		Id:               campaign.Id,
		Creator:          user.FormatUserBasicResponse(campaign.Users),
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		GoalAmount:       campaign.GoalAmount,
		ImageUrl:         imageUrl,
		CurrentAmount:    campaign.CurrentAmount,
	}
}

// format slice campaign
func FormatCampaignsResponse(campaigns []Campaign) []CampaignResponse {
	var formattedCampaigns []CampaignResponse

	for _, campaign := range campaigns {
		formattedCampaign := FormatCampaignResponse(campaign)
		formattedCampaigns = append(formattedCampaigns, formattedCampaign)
	}

	return formattedCampaigns
}
