package transaction

import (
	"time"
)

type CampaignTransactionFormatter struct {
	ID 					int 			`json:"id"`
	Name 				string 		`json:"name"`
	Amount 			int 			`json:"amount"`
	CreatedAt 	time.Time `json:"created_at"`
}

func FormatCampaignTransaction(transaction Transaction) CampaignTransactionFormatter {
	formatter := CampaignTransactionFormatter{}
	formatter.ID = transaction.ID
	formatter.Name = transaction.User.Name
	formatter.Amount = transaction.Amount
	formatter.CreatedAt = transaction.CreatedAt

	return formatter
}

func FormatCampaignTransactions(transactions []Transaction) []CampaignTransactionFormatter{
	if len(transactions) == 0 {
		return []CampaignTransactionFormatter{}
	}

	var transactionsFormatter []CampaignTransactionFormatter

	for _, transaction := range transactions {
		formatter := FormatCampaignTransaction(transaction)
		transactionsFormatter = append(transactionsFormatter, formatter)
	}

	return transactionsFormatter
}

type UserTransactionFormatter struct {
	ID 					int 			`json:"id"`
	Amount 			int	 			`json:"amount"`
	Status 			string 		`json:"status"`
	CreatedAt 	time.Time `json:"created_at"`
	Campaign  CampaignFormatter `json:"campaign"`
}

type CampaignFormatter struct {
	Name					string				`json:"name"`
	ImageURL			string				`json:"image_url"`
}

func FormatUserTransaction(transactions Transaction) UserTransactionFormatter{
	formatter := UserTransactionFormatter{}
	formatter.ID = transactions.ID
	formatter.Amount = transactions.Amount
	formatter.Status = transactions.Status
	formatter.CreatedAt = transactions.CreatedAt

	campaignFormatter := CampaignFormatter{}
	campaignFormatter.Name = transactions.Campaign.Name
	campaignFormatter.ImageURL = ""

	if len(transactions.Campaign.CampaignImages) > 0 {
		campaignFormatter.ImageURL = transactions.Campaign.CampaignImages[0].FileName
	}

	formatter.Campaign = campaignFormatter

	return formatter
}

func FormatUserTransactions(transactions []Transaction) []UserTransactionFormatter{
	if len(transactions) == 0 {
		return []UserTransactionFormatter{}
	}

	var transactionsFormatter []UserTransactionFormatter

	for _, transaction := range transactions {
		formatter := FormatUserTransaction(transaction)
		transactionsFormatter = append(transactionsFormatter, formatter)
	}

	return transactionsFormatter
}