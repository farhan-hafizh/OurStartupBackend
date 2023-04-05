package transaction

import (
	"ourstartup/entities"
	"time"
)

type BasicTransactionResponse struct {
	Id        string    `json:"id"` // code
	Amount    int       `json:"amount"`
	IsSecret  bool      `json:"is_secret"`
	CreatedAt time.Time `json:"created_at"`
}

type NewTransactionResponse struct {
	BasicTransactionResponse
	PaymentUrl string `json:"payment_url"`
	Status     string `json:"status"`
}

type TransactionResponse struct { // general transaction access
	BasicTransactionResponse
	Name string `json:"name"`
}

type TransactionHistoryResponse struct { // trans owner access
	BasicTransactionResponse                             //embed
	Status                   string                      `json:"status"`
	Campaign                 TransactionCampaignResponse `json:"campaign"`
}

type TransactionCampaignResponse struct {
	Name     string `json:"name"`
	Slug     string `json:"slug"`
	ImageUrl string `json:"image_url"`
}

func FormatBasicTransactionResponse(trans entities.Transaction) BasicTransactionResponse {
	return BasicTransactionResponse{
		Id:        trans.Code,
		Amount:    trans.Amount,
		IsSecret:  trans.IsSecret,
		CreatedAt: trans.CreatedAt,
	}
}

func FormatTransactionResponse(transaction entities.Transaction) TransactionResponse {
	return TransactionResponse{
		BasicTransactionResponse: FormatBasicTransactionResponse(transaction),
		Name:                     transaction.User.Name,
	}
}

func FormatTransactionsResponse(transactions []entities.Transaction) []TransactionResponse {

	if len(transactions) == 0 {
		return []TransactionResponse{}
	}

	var response []TransactionResponse

	for _, trans := range transactions {
		transResponse := FormatTransactionResponse(trans)
		response = append(response, transResponse)
	}

	return response
}

// single trans
func FormatTranHistoryResponse(transaction entities.Transaction) TransactionHistoryResponse {

	imageUrl := ""

	if len(transaction.Campaign.CampaignImages) > 0 {
		imageUrl = transaction.Campaign.CampaignImages[0].FileName
	}

	return TransactionHistoryResponse{
		BasicTransactionResponse: FormatBasicTransactionResponse(transaction),
		Status:                   transaction.Status,
		Campaign: TransactionCampaignResponse{
			Name:     transaction.Campaign.Name,
			Slug:     transaction.Campaign.Slug,
			ImageUrl: imageUrl,
		},
	}
}

// plural
func FormatTransHistoryResponse(transactions []entities.Transaction) []TransactionHistoryResponse {

	if len(transactions) == 0 {
		return []TransactionHistoryResponse{}
	}

	var response []TransactionHistoryResponse

	for _, data := range transactions {
		transResponse := FormatTranHistoryResponse(data)
		response = append(response, transResponse)
	}

	return response
}

func FormatNewTransactionResponse(transaction entities.Transaction) NewTransactionResponse {
	return NewTransactionResponse{
		BasicTransactionResponse: FormatBasicTransactionResponse(transaction),
		PaymentUrl:               transaction.PaymentUrl,
		Status:                   transaction.Status,
	}
}
