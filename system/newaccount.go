package system

import (
	eos "github.com/UncleAndy/cyberway-go"
	"github.com/UncleAndy/cyberway-go/ecc"
)

// NewNewAccount returns a `newaccount` action that lives on the
// `eosio.system` contract.
func NewNewAccount(creator, newAccount eos.AccountName, ownerKey, activeKey ecc.PublicKey, permission string) *eos.Action {
	return &eos.Action{
		Account: AN("cyber"),
		Name:    ActN("newaccount"),
		Authorization: []eos.PermissionLevel{
			{Actor: creator, Permission: PN(permission)},
		},
		ActionData: eos.NewActionData(NewAccount{
			Creator: creator,
			Name:    newAccount,
			Owner: eos.Authority{
				Threshold: 1,
				Keys: []eos.KeyWeight{
					{
						PublicKey: ownerKey,
						Weight:    1,
					},
				},
				Accounts: []eos.PermissionLevelWeight{},
			},
			Active: eos.Authority{
				Threshold: 1,
				Keys: []eos.KeyWeight{
					{
						PublicKey: activeKey,
						Weight:    1,
					},
				},
				Accounts: []eos.PermissionLevelWeight{},
			},
		}),
	}
}

// NewDelegatedNewAccount returns a `newaccount` action that lives on the
// `eosio.system` contract. It is filled with an authority structure that
// delegates full control of the new account to an already existing account.
func NewDelegatedNewAccount(creator, newAccount eos.AccountName, delegatedTo eos.AccountName) *eos.Action {
	return &eos.Action{
		Account: AN("cyber"),
		Name:    ActN("newaccount"),
		Authorization: []eos.PermissionLevel{
			{Actor: creator, Permission: PN("active")},
		},
		ActionData: eos.NewActionData(NewAccount{
			Creator: creator,
			Name:    newAccount,
			Owner: eos.Authority{
				Threshold: 1,
				Keys:      []eos.KeyWeight{},
				Accounts: []eos.PermissionLevelWeight{
					eos.PermissionLevelWeight{
						Permission: eos.PermissionLevel{
							Actor:      delegatedTo,
							Permission: PN("active"),
						},
						Weight: 1,
					},
				},
			},
			Active: eos.Authority{
				Threshold: 1,
				Keys:      []eos.KeyWeight{},
				Accounts: []eos.PermissionLevelWeight{
					eos.PermissionLevelWeight{
						Permission: eos.PermissionLevel{
							Actor:      delegatedTo,
							Permission: PN("active"),
						},
						Weight: 1,
					},
				},
			},
		}),
	}
}

// NewCustomNewAccount returns a `newaccount` action that lives on the
// `eosio.system` contract. You can specify your own `owner` and
// `active` permissions.
func NewCustomNewAccount(creator, newAccount eos.AccountName, owner, active, posting eos.Authority) *eos.Action {
	return &eos.Action{
		Account: AN("cyber"),
		Name:    ActN("newaccount"),
		Authorization: []eos.PermissionLevel{
			{Actor: creator, Permission: PN("active")},
		},
		ActionData: eos.NewActionData(NewAccount{
			Creator: creator,
			Name:    newAccount,
			Owner:   owner,
			Active:  active,
		}),
	}
}

// NewAccount represents a `newaccount` action on the `eosio.system`
// contract. It is one of the rare ones to be hard-coded into the
// blockchain.
type NewAccount struct {
	Creator eos.AccountName `json:"creator"`
	Name    eos.AccountName `json:"name"`
	Owner   eos.Authority   `json:"owner"`
	Active  eos.Authority   `json:"active"`
}
