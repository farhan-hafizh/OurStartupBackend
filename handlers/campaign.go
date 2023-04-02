package handlers

import (
	"net/http"
	"ourstartup/helper"
	"ourstartup/services/campaign"
	"ourstartup/services/user"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	service     campaign.Service
	userService user.Service
}

func CreateCampaignHandler(service campaign.Service, userService user.Service) *campaignHandler {
	return &campaignHandler{service, userService}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	username := c.Query("username")
	user, err := h.userService.GetUserByUsername(username)
	if err != nil {
		helper.SendErrorResponse(
			c,
			"Failed to get campaign! No user found with that username!",
			http.StatusBadRequest,
			"failed", err, nil)
		return
	}
	// send user id neither it nil or not
	campaigns, err := h.service.GetCampaigns(user.Id)

	if err != nil {
		helper.SendErrorResponse(
			c,
			"Failed to get campaigns!",
			http.StatusBadRequest,
			"failed", err, nil)
		return
	}

	formattedCampaigns := campaign.FormatCampaignsResponse(campaigns)
	helper.SendResponse(c, "Successfully get campaigns!", http.StatusOK, "success", formattedCampaigns)

}

func (h *campaignHandler) CreateCampaign(c *gin.Context) {
	input := &campaign.CreateCampaignInput{}

	err := c.ShouldBindJSON(&input)

	if err != nil {
		helper.SendValidationErrorResponse(
			c,
			"Create campaign failed",
			http.StatusUnprocessableEntity,
			"failed",
			err)
		return
	}
	userData := c.MustGet("loggedInUser").(user.User)
	userId := userData.Id

	newCampaign, err := h.service.CreateCampaign(userId, *input)

	if err != nil {
		helper.SendErrorResponse(
			c,
			"Failed to create campaign!",
			http.StatusInternalServerError,
			"failed", err, nil)
		return
	}
	newCampaign.User.Name = userData.Name
	newCampaign.User.Username = userData.Username
	newCampaign.User.Occupation = userData.Occupation

	formattedCampaign := campaign.FormatCampaignResponse(newCampaign)

	helper.SendResponse(c, "Campaign successfully created!", http.StatusOK, "success", formattedCampaign)
}

func (h *campaignHandler) GetCampaignDetail(c *gin.Context) {
	var input campaign.GetCampaignSlugInput

	err := c.ShouldBindUri(&input)

	if err != nil {
		helper.SendValidationErrorResponse(
			c,
			"Get campaign detail failed!",
			http.StatusUnprocessableEntity,
			"failed",
			err)
		return
	}

	campaignData, err := h.service.GetCampaignBySlug(input.Slug)

	if err != nil {
		helper.SendErrorResponse(
			c,
			"Get campaign detail failed!",
			http.StatusInternalServerError,
			"failed", err, nil)
		return
	}
	formattedCampaign := campaign.FormatDetailCampaignResponse(campaignData)

	helper.SendResponse(c, "Successfully get campaign detail!", http.StatusOK, "success", formattedCampaign)
}
