package domain

import (
	eos "github.com/UncleAndy/cyberway-go"
)

// NewNewAccount returns a `newusername` action that lives on the
// `eosio.system` contract.
func NewNewUserName(creator, owner, name, permission string) *eos.Action {
	return &eos.Action{
		Account: eos.AN("cyber.domain"),
		Name:    eos.ActN("newusername"),
		Authorization: []eos.PermissionLevel{
			{Actor: eos.AccountName(creator), Permission: eos.PN(permission)},
		},
		ActionData: eos.NewActionData(NewUserName{
			Creator: eos.AccountName(creator),
			Owner:   eos.AccountName(owner),
			Name:    name,
		}),
	}
}

// NewUserName represents a `newusername` action on the `cyber.domain`
// contract. It is one of the rare ones to be hard-coded into the
// blockchain.
type NewUserName struct {
	Creator eos.AccountName 	`json:"creator"`
	Owner   eos.AccountName   	`json:"owner"`
	Name    string 				`json:"name"`
}
