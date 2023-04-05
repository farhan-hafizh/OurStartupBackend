package handlers

import (
	"net/http"
	"ourstartup/entities"
	"ourstartup/helper"
	"ourstartup/services/campaign"
	"ourstartup/services/transaction"
	"ourstartup/services/user"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	service         transaction.Service
	userService     user.Service
	campaignService campaign.Service
}

func CreateTransactionHandler(service transaction.Service, userService user.Service, campaignService campaign.Service) *transactionHandler {
	return &transactionHandler{service, userService, campaignService}
}

func (h *transactionHandler) CreateTransaction(c *gin.Context) {

	var input transaction.CreateTransactionInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		helper.SendValidationErrorResponse(
			c,
			"Create transaction failed!",
			http.StatusUnprocessableEntity,
			"failed",
			err)
		return
	}
	userData := c.MustGet("loggedInUser").(entities.User)
	input.User = userData

	campaignSlug := campaign.GetCampaignSlugInput{
		Slug: input.CampaignSlug,
	}

	campaign, err := h.campaignService.GetCampaignBySlug(campaignSlug)

	if err != nil {
		helper.SendErrorResponse(
			c,
			"Create transaction failed! Failed to get campaign!",
			http.StatusBadRequest,
			"failed",
			err, nil)
		return
	}

	input.Campaign = campaign

	trans, err := h.service.CreateTransaction(input)

	if err != nil {
		helper.SendErrorResponse(
			c,
			"Create transaction failed!",
			http.StatusInternalServerError,
			"failed",
			err, nil)
		return
	}

	formattedTrans := transaction.FormatNewTransactionResponse(trans)

	helper.SendResponse(c, "Successfully create transaction!", http.StatusOK, "success", formattedTrans)
}

func (h *transactionHandler) GetTransHistoryByCampaign(c *gin.Context) {

	var input campaign.GetCampaignSlugInput

	err := c.ShouldBindUri(&input)

	if err != nil {
		helper.SendValidationErrorResponse(
			c,
			"Get transaction history failed!",
			http.StatusUnprocessableEntity,
			"failed",
			err)
		return
	}

	campaign, err := h.campaignService.GetCampaignBySlug(input)

	if err != nil {
		helper.SendErrorResponse(
			c,
			"Get transaction history failed!",
			http.StatusBadRequest,
			"failed",
			err, nil)
		return
	}

	userData := c.MustGet("loggedInUser").(entities.User)

	isAll := false
	// if uri campaign owner not null
	if input.CampaignOwner != "" {
		if input.CampaignOwner == "campaign-owner" {
			isAll = true
		} else { // if uri not exactly "campaign-owner"
			helper.SendErrorResponse(
				c,
				"Get transaction history failed!",
				http.StatusBadRequest,
				"failed",
				err, nil)
			return
		}
	}

	transInput := transaction.GetTransByCampaignId{
		Campaign:   campaign,
		IsAllTrans: isAll,
		User:       userData,
	}
	transactions, err := h.service.GetTransByCampaignId(transInput)

	if err != nil {
		helper.SendErrorResponse(
			c,
			"Get transaction history failed!",
			http.StatusInternalServerError,
			"failed",
			err, nil)
		return
	}

	formattedTransactions := transaction.FormatTransactionsResponse(transactions)

	helper.SendResponse(c, "Successfully get campaign detail!", http.StatusOK, "success", formattedTransactions)
}

func (h *transactionHandler) GetTransactionHistory(c *gin.Context) {
	userData := c.MustGet("loggedInUser").(entities.User)

	input := transaction.GetTransactionHistory{
		User: userData,
	}

	transactions, err := h.service.GetTransactionHistory(input)

	if err != nil {
		helper.SendErrorResponse(
			c,
			"Get transaction history failed!",
			http.StatusInternalServerError,
			"failed",
			err, nil)
		return
	}

	formattedTransactions := transaction.FormatTransHistoryResponse(transactions)
	helper.SendResponse(c, "Successfully get campaign detail!", http.StatusOK, "success", formattedTransactions)
}
