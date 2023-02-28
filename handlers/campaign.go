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
	campaigns, err := h.service.GetCampaigns(user.Id)

	if err != nil {
		helper.SendErrorResponse(
			c,
			"Failed to get campaigns!",
			http.StatusBadRequest,
			"failed", err, nil)
		return
	}
	if len(campaigns) == 0 {
		helper.SendErrorResponse(
			c,
			"Campaigns not found!",
			http.StatusNoContent,
			"failed", nil, nil)
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

	formattedCampaign := campaign.FormatCampaignResponse(newCampaign)

	helper.SendResponse(c, "Campaign successfully created!", http.StatusOK, "success", formattedCampaign)
}
