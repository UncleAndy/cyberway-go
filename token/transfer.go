package token

import eos "github.com/UncleAndy/cyberway-go"

func NewTransfer(from, to string, quantity eos.Asset, memo string) *eos.Action {
	return &eos.Action{
		Account: AN("cyber.token"),
		Name:    ActN("transfer"),
		Authorization: []eos.PermissionLevel{
			{Actor: eos.AccountName(from), Permission: PN("active")},
		},
		ActionData: eos.NewActionData(Transfer{
			From:     eos.AccountName(from),
			To:       eos.AccountName(to),
			Quantity: quantity,
			Memo:     memo,
		}),
	}
}

// Transfer represents the `transfer` struct on `eosio.token` contract.
type Transfer struct {
	From     eos.AccountName `json:"from"`
	To       eos.AccountName `json:"to"`
	Quantity eos.Asset       `json:"quantity"`
	Memo     string          `json:"memo"`
}
