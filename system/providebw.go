package system

import eos "github.com/UncleAndy/cyberway-go"

func NewProvideBW(account, provider eos.AccountName, usingAccount eos.AccountName, usingPermission eos.PermissionName) *eos.Action {
	a := &eos.Action{
		Account: AN("cyber"),
		Name:    ActN("providebw"),
		Authorization: []eos.PermissionLevel{
			{
				Actor:      usingAccount,
				Permission: usingPermission,
			},
		},
		ActionData: eos.NewActionData(ProvideBW{
			Account:    account,
			Provider:	provider,
		}),
	}

	return a
}

type ProvideBW struct {
	Provider    eos.AccountName    `json:"provider"`
	Account     eos.AccountName    `json:"account"`
}
