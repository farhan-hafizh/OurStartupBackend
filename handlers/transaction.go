package handlers

import (
	"ourstartup/services/transaction"
	"ourstartup/services/user"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	service     transaction.Service
	userService user.Service
}

func CreateTransactionHandler(service transaction.Service, userService user.Service) *transactionHandler {
	return &transactionHandler{service, userService}
}

func (h *transactionHandler) CreateTransaction(c *gin.Context) {

}

func (h *transactionHandler) GetTransHistory(c *gin.Context) {}
