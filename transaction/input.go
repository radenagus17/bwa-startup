package transaction

import "bwastartup/user"

type GetTransactionsCampaignInput struct { 
	ID 				int 			`uri:"id" binding:"required"`
	User 			user.User
}

type CreateTransactionsInput struct {
	CampaignID 			int 				`json:"campaign_id" binding:"required"`
	Amount 					int 				`json:"amount" binding:"required"`
	User 						user.User
}

type TransactionNotificationInput struct {
	TransactionStatus string `json:"transaction_status"`
	OrderID string `json:"order_id"`
	PaymentType string `json:"payment_type"`
	FraudStatus string `json:"fraud_status"`
}