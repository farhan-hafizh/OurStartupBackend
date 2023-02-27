package handlers

import (
	"net/http"
	"ourstartup/helper"
	"ourstartup/services/campaign"
	"ourstartup/services/user"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	service campaign.Service
}

func CreateCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.service.GetCampaigns(userId)

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

	helper.SendResponse(c, "Successfully get campaigns!", http.StatusOK, "success", campaigns)

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

	campaign, err := h.service.CreateCampaign(userId, *input)

	if err != nil {
		helper.SendErrorResponse(
			c,
			"Failed to create campaign!",
			http.StatusInternalServerError,
			"failed", err, nil)
		return
	}

	helper.SendResponse(c, "Campaign successfully created!", http.StatusOK, "success", campaign)
}
