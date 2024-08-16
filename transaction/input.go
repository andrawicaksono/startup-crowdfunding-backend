package transaction

type GetTransactionsByCampaignIDInput struct {
	CampaignID int `uri:"id" binding:"required"`
}
