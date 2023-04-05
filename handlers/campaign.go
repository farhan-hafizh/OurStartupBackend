package handlers

import (
	"fmt"
	"net/http"
	"ourstartup/helper"
	"ourstartup/services/campaign"
	"ourstartup/services/user"
	"time"

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
	input := campaign.CreateCampaignInput{}

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
	input.User = userData

	newCampaign, err := h.service.CreateCampaign(input)

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

	campaignData, err := h.service.GetCampaignBySlug(input)

	if err != nil {
		helper.SendErrorResponse(
			c,
			"Get campaign detail failed!",
			http.StatusInternalServerError,
			"failed", err, nil)
		return
	}
	var formattedCampaign interface{}
	if campaignData.Id != 0 {
		formattedCampaign = campaign.FormatDetailCampaignResponse(campaignData)
	}

	helper.SendResponse(c, "Successfully get campaign detail!", http.StatusOK, "success", formattedCampaign)
}

func (h *campaignHandler) UpdateCampaign(c *gin.Context) {
	var slugData campaign.GetCampaignSlugInput

	err := c.ShouldBindUri(&slugData)

	if err != nil {
		helper.SendValidationErrorResponse(
			c,
			"Update campaign failed!",
			http.StatusUnprocessableEntity,
			"failed",
			err)
		return
	}

	var inputCampaign campaign.CreateCampaignInput

	err = c.ShouldBindJSON(&inputCampaign)

	if err != nil {
		helper.SendValidationErrorResponse(
			c,
			"Update campaign failed!",
			http.StatusUnprocessableEntity,
			"failed",
			err)
		return
	}

	userData := c.MustGet("loggedInUser").(user.User)
	slugData.User = userData

	updatedCampaign, err := h.service.UpdateCampaign(slugData, inputCampaign)

	if err != nil {
		helper.SendErrorResponse(
			c,
			"Update campaign failed!",
			http.StatusInternalServerError,
			"failed", err, nil)
		return
	}

	updatedCampaign.User.Name = userData.Name
	updatedCampaign.User.Username = userData.Username
	updatedCampaign.User.Occupation = userData.Occupation

	formattedCampaign := campaign.FormatCampaignResponse(updatedCampaign)

	helper.SendResponse(c, "Campaign successfully updated!", http.StatusOK, "success", formattedCampaign)

}

func (h *campaignHandler) UploadCampaignImage(c *gin.Context) {
	var input campaign.CreateCampaignImageInput

	err := c.ShouldBind(&input)

	response := gin.H{"is_uploaded": false}

	if err != nil {
		helper.SendErrorResponse(
			c,
			"Upload campaign image failed!",
			http.StatusUnprocessableEntity,
			"failed",
			err, response)
		return
	}

	userData := c.MustGet("loggedInUser").(user.User)
	input.User = userData

	file, err := c.FormFile("file")

	if err != nil {
		helper.SendErrorResponse(
			c,
			"Upload campaign image failed!",
			http.StatusBadRequest,
			"failed",
			err, response)
		return
	}

	// create file path and filename
	path := fmt.Sprintf("images/campaign-%s-%d-%s", input.User.Username, time.Now().Unix(), file.Filename)

	// save uploaded file to filepath with filename
	err = c.SaveUploadedFile(file, path)

	if err != nil {
		helper.SendErrorResponse(
			c,
			"Upload campaign image failed!",
			http.StatusInternalServerError,
			"failed",
			err, response)
		return
	}

	_, err = h.service.SaveCampaignImage(input, path)

	if err != nil {
		helper.SendErrorResponse(
			c,
			"Update campaign failed!",
			http.StatusInternalServerError,
			"failed", err, response)
		return
	}
	response = gin.H{"is_uploaded": true}

	helper.SendResponse(c, "Campaign image successfully uploaded!", http.StatusOK, "success", response)

}
