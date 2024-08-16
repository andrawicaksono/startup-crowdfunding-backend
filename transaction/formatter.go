package transaction

import "time"

type CampaignTransactionFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

func FormatCampaignTransaction(transaction Transaction) CampaignTransactionFormatter {
	transactionFormatter := CampaignTransactionFormatter{
		ID:        transaction.ID,
		Name:      transaction.User.Name,
		Amount:    transaction.Amount,
		CreatedAt: transaction.CreatedAt,
	}

	return transactionFormatter
}

func FormatCampaignTransactions(transactions []Transaction) []CampaignTransactionFormatter {
	transactionsFormatter := []CampaignTransactionFormatter{}

	for _, transaction := range transactions {
		transactionFormatter := FormatCampaignTransaction(transaction)

		transactionsFormatter = append(transactionsFormatter, transactionFormatter)
	}

	return transactionsFormatter
}
