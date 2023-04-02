package campaign

import (
	"ourstartup/services/user"
	"strings"
)

type CampaignResponse struct {
	Id               int                    `json:"id"`
	Name             string                 `json:"name"`
	Creator          user.UserBasicResponse `json:"creator"`
	ShortDescription string                 `json:"short_description"`
	ImageUrl         string                 `json:"image_url"`
	Slug             string                 `json:"slug"`
	GoalAmount       int                    `json:"goal_amount"`
	CurrentAmount    int                    `json:"current_amount"`
}

type CampaignDetailResponse struct {
	Name             string                       `json:"name"`
	ShortDescription string                       `json:"short_description"`
	ImageUrl         string                       `json:"image_url"`
	GoalAmount       int                          `json:"goal_amount"`
	CurrentAmount    int                          `json:"current_amount"`
	Description      string                       `json:"description"`
	Slug             string                       `json:"slug"`
	User             user.UserWithProfileResponse `json:"user"`
	Perks            []string                     `json:"perks"`
	Images           []CampaignImagesResponse     `json:"images"`
}

type CampaignImagesResponse struct {
	ImageUrl  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

func FormatDetailCampaignResponse(campaign Campaign) CampaignDetailResponse {
	var perksArray []string
	// split perks by koma
	perks := strings.Split(campaign.Perks, ",")
	//trim if perk has whitespace then append to perksArray
	for _, perk := range perks {
		perksArray = append(perksArray, strings.TrimSpace(perk))
	}

	// create primary image url
	imageUrl := ""
	if len(campaign.CampaignImages) > 0 {
		imageUrl = campaign.CampaignImages[0].FileName
	}

	// prevent null on empty array
	images := make([]CampaignImagesResponse, 0)
	for _, image := range campaign.CampaignImages {
		images = append(images, CampaignImagesResponse{
			ImageUrl:  image.FileName,
			IsPrimary: image.IsPrimary,
		})
	}

	return CampaignDetailResponse{
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		ImageUrl:         imageUrl, //primary image
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
		Description:      campaign.Description,
		Slug:             campaign.Slug,
		Images:           images,
		Perks:            perksArray,
		User:             user.FormatUserWithProfileResponse(campaign.User),
	}
}

// format campaign
func FormatCampaignResponse(campaign Campaign) CampaignResponse {

	imageUrl := ""

	if len(campaign.CampaignImages) > 0 {
		imageUrl = campaign.CampaignImages[0].FileName
	}

	return CampaignResponse{
		Id:               campaign.Id,
		Creator:          user.FormatUserBasicResponse(campaign.User),
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		GoalAmount:       campaign.GoalAmount,
		ImageUrl:         imageUrl,
		Slug:             campaign.Slug,
		CurrentAmount:    campaign.CurrentAmount,
	}
}

// format slice campaign
func FormatCampaignsResponse(campaigns []Campaign) []CampaignResponse {

	if len(campaigns) == 0 {
		return []CampaignResponse{}
	}
	var formattedCampaigns []CampaignResponse

	for _, campaign := range campaigns {
		formattedCampaign := FormatCampaignResponse(campaign)
		formattedCampaigns = append(formattedCampaigns, formattedCampaign)
	}

	return formattedCampaigns
}
