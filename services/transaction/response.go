package transaction

import (
	"time"
)

type BasicTransactionResponse struct {
	Id        string    `json:"id"`
	Amount    int       `json:"amount"`
	IsSecret  bool      `json:"is_secret"`
	CreatedAt time.Time `json:"created_at"`
}

type TransactionResponse struct {
	BasicTransactionResponse
	Name string `json:"name"`
}

type TransactionHistoryResponse struct {
	BasicTransactionResponse //embed
	Status                   string
	Campaign                 TransactionCampaignResponse `json:"campaign"`
}

type TransactionCampaignResponse struct {
	Name     string `json:"name"`
	Slug     string `json:"slug"`
	ImageUrl string `json:"image_url"`
}

func FormatBasicTransactionResponse(trans Transaction) BasicTransactionResponse {
	return BasicTransactionResponse{
		Id:        trans.CallerId,
		Amount:    trans.Amount,
		IsSecret:  trans.IsSecret,
		CreatedAt: trans.CreatedAt,
	}
}

func FormatTransactionResponse(transaction Transaction) TransactionResponse {
	return TransactionResponse{
		BasicTransactionResponse: FormatBasicTransactionResponse(transaction),
		Name:                     transaction.User.Name,
	}
}

func FormatTransactionsResponse(transactions []Transaction) []TransactionResponse {

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

func FormatTranHistoryResponse(transaction Transaction) TransactionHistoryResponse {

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

func FormatTransHistoryResponse(transactions []Transaction) []TransactionHistoryResponse {

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
