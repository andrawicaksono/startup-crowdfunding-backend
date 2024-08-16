package payment

import (
	"startup-crowdfunding-backend/user"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"github.com/spf13/viper"
)

type Service interface {
	GetPaymentURL(transaction Transaction, user user.User) (string, error)
}

type service struct {
	snapClient snap.Client
}

func NewService(viper *viper.Viper) *service {
	serverKey := viper.GetString("midtrans.server_key")
	var snapClient snap.Client
	snapClient.New(serverKey, midtrans.Sandbox)

	return &service{snapClient}
}

func (s *service) GetPaymentURL(transaction Transaction, user user.User) (string, error) {
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  transaction.Code,
			GrossAmt: int64(transaction.Amount),
		},

		CustomerDetail: &midtrans.CustomerDetails{
			FName: user.Name,
			Email: user.Email,
		},
	}

	snapResp, err := s.snapClient.CreateTransaction(req)
	if err != nil {
		return snapResp.RedirectURL, err
	}

	return snapResp.RedirectURL, nil
}
