package transaction

import "startup-crowdfunding-backend/user"

type GetTransactionsByCampaignIDInput struct {
	CampaignID int `uri:"id" binding:"required"`
}

type CreateTransactionInput struct {
	CampaignID int `json:"campaign_id" binding:"required"`
	Amount     int `json:"amount" binding:"required"`
	User       user.User
}
