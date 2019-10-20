package forum

import (
	eos "github.com/UncleAndy/cyberway-go"
)

// NewMessage is an action representing a simple message to be posted
// through the chain network.
func NewMessage(msg *MessageAction) *eos.Action {
	a := &eos.Action{
		Account: ForumAN,
		Name:    ActN("post"),
		Authorization: []eos.PermissionLevel{
			{Actor: poster, Permission: eos.PermissionName("active")},
		},
		ActionData: eos.NewActionData(msg),
	}
	return a
}

// Post represents the `eosio.forum::post` action.
type MessageAction struct {
	Id       *MessageIdType `json:"message_id"`
	ParentId *MessageIdType `json:"parent_id"`

	Language string `json:"languagemssg"`
	Header   string `json:"headermssg"`
	Body     string `json:"bodymssg"`

	Tags         []string `json:"tags"`
	JsonMetadata string   `json:"jsonmetadata"`

	TokenProp     int      `json:"tokenprop"`
	MaxPayout     *int     `json:"max_payout"`
	Beneficiaries []eos.AccountName `json:"beneficiaries"`
	CuratorsPrcnt int      `json:"curators_prcnt"`
	VestPayment   bool     `json:"vestpayment"`
}

type MessageIdType struct {
	Author		eos.AccountName	`json:"author"`
	Permlink	string			`json:"permlink"`
}
