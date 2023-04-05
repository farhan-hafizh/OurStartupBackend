package payment

type TransactionNotification struct {
	TransactionStatus string `json:"transaction_status"`
	OrderId           string `json:"order_id"`
	PaymentType       string `json:"payment_type"`
	FraudStatus       string `json:"fraud_status"`
}
