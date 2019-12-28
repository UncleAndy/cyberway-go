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
			RamPayer: eos.AccountName(creator),
			Owner:    eos.AccountName(owner),
			Symbol:   symbol,
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
			RamPayer: eos.AccountName(creator),
			Owner:    eos.AccountName(owner),
			Symbol:   symbol,
		}),
	}
}

// Create represents the `create` struct on the `eosio.token` contract.
type Open struct {
	RamPayer	eos.AccountName		`json:"ram_payer"`
	Owner		eos.AccountName		`json:"owner"`
	Symbol		eos.Symbol			`json:"symbol"`
}
