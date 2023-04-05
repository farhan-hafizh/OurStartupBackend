package handlers

import (
	"net/http"
	"ourstartup/entities"
	"ourstartup/helper"
	"ourstartup/services/campaign"
	"ourstartup/services/payment"
	"ourstartup/services/transaction"
	"ourstartup/services/user"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	service         transaction.Service
	userService     user.Service
	campaignService campaign.Service
	paymentService  payment.Service
}

func CreateTransactionHandler(service transaction.Service, userService user.Service, campaignService campaign.Service, paymentService payment.Service) *transactionHandler {
	return &transactionHandler{service, userService, campaignService, paymentService}
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

	campaignSlug := campaign.GetCampaignInput{
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

	paymentUrl, err := h.paymentService.GetRedirectUrl(trans)

	if err != nil {
		helper.SendErrorResponse(
			c,
			"Create transaction failed!",
			http.StatusInternalServerError,
			"failed",
			err, nil)
		return
	}

	trans.PaymentUrl = paymentUrl

	updatedTrans, err := h.service.UpdateTransaction(trans)

	formattedTrans := transaction.FormatNewTransactionResponse(updatedTrans)

	helper.SendResponse(c, "Successfully create transaction!", http.StatusOK, "success", formattedTrans)
}

func (h *transactionHandler) GetTransHistoryByCampaign(c *gin.Context) {

	var input campaign.GetCampaignInput

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

func (h *transactionHandler) GetNotification(c *gin.Context) {
	var input payment.TransactionNotification

	err := c.ShouldBindJSON(&input)

	if err != nil {
		helper.SendValidationErrorResponse(
			c,
			"Failed to process notification!",
			http.StatusUnprocessableEntity,
			"failed",
			err)
		return
	}

	err = h.paymentService.ProcessPayment(input)

	if err != nil {
		helper.SendErrorResponse(
			c,
			"Failed to process notification!",
			http.StatusInternalServerError,
			"failed",
			err, nil)
		return
	}
	helper.SendResponse(c, "Successfully get campaign detail!", http.StatusOK, "success", input)
}
