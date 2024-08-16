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

type TransactionCampaignFormatter struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

type UserTransactionFormatter struct {
	ID        int                          `json:"id"`
	Amount    int                          `json:"amount"`
	Status    string                       `json:"status"`
	CreatedAt time.Time                    `json:"created_at"`
	Campaign  TransactionCampaignFormatter `json:"campaign"`
}

func FormatUserTransaction(transaction Transaction) UserTransactionFormatter {
	transactionCampaignFormatter := TransactionCampaignFormatter{
		Name:     transaction.Campaign.Name,
		ImageURL: transaction.Campaign.CampaignImages[0].FileName,
	}

	userTransactionFormatter := UserTransactionFormatter{
		ID:        transaction.ID,
		Amount:    transaction.Amount,
		Status:    transaction.Status,
		CreatedAt: transaction.CreatedAt,
		Campaign:  transactionCampaignFormatter,
	}

	return userTransactionFormatter
}

func FormatUserTransactions(transactions []Transaction) []UserTransactionFormatter {
	userTransactionsFormatter := []UserTransactionFormatter{}

	for _, transaction := range transactions {
		userTransactionFormatter := FormatUserTransaction(transaction)
		userTransactionsFormatter = append(userTransactionsFormatter, userTransactionFormatter)
	}

	return userTransactionsFormatter
}
