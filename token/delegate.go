package token

import eos "github.com/UncleAndy/cyberway-go"

func NewDelegateVesting(from, to string, val eos.Asset, interest uint16, permissionFrom, permissionTo string) *eos.Action {
	return &eos.Action{
		Account: AN("gls.vesting"),
		Name:    ActN("delegate"),
		Authorization: []eos.PermissionLevel{
			{Actor: AN(from), Permission: PN(permissionFrom)},
			{Actor: AN(to), Permission: PN(permissionTo)},
		},
		ActionData: eos.NewActionData(Delegate{
				From: AN(from),
				To: AN(to),
				Quantity: val,
				InterestRate: interest,
		}),
	}
}

type Delegate struct {
	From			eos.AccountName		`json:"from"`
	To				eos.AccountName		`json:"to"`
	Quantity		eos.Asset			`json:"quantity"`
	InterestRate	uint16			`json:"interest_rate"`
}
