package token

import eos "github.com/UncleAndy/cyberway-go"

func NewTokenOpen(creator, owner string, symbol eos.Symbol, permission string) *eos.Action {
	return &eos.Action{
		Account: AN("cyber.token"),
		Name:    ActN("open"),
		Authorization: []eos.PermissionLevel{
			{Actor: AN(creator), Permission: PN(permission)},
		},
		ActionData: eos.NewActionData(Open{
			Owner:    eos.AccountName(owner),
			Symbol:   symbol,
			RamPayer: eos.AccountName(creator),
		}),
	}
}

func NewVestingOpen(creator, owner string, symbol eos.Symbol, permission string) *eos.Action {
	return &eos.Action{
		Account: AN("gls.vesting"),
		Name:    ActN("open"),
		Authorization: []eos.PermissionLevel{
			{Actor: AN(creator), Permission: PN(permission)},
		},
		ActionData: eos.NewActionData(Open{
			Owner:    eos.AccountName(owner),
			Symbol:   symbol,
			RamPayer: eos.AccountName(creator),
		}),
	}
}

// Create represents the `create` struct on the `eosio.token` contract.
type Open struct {
	Owner		eos.AccountName		`json:"owner"`
	Symbol		eos.Symbol			`json:"symbol"`
	RamPayer	eos.AccountName		`json:"ram_payer"`
}
