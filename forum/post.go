package forum

import (
	eos "github.com/UncleAndy/cyberway-go"
)

// NewMessage is an action representing a simple message to be posted
// through the chain network.
func CreateMessage(msg *CreateMssg) *eos.Action {
	a := &eos.Action{
		Account: ForumAN,
		Name:    ActN("createmssg"),
		Authorization: []eos.PermissionLevel{
			{Actor: msg.Id.Author, Permission: eos.PermissionName("posting")},
		},
		ActionData: eos.NewActionData(msg),
	}
	return a
}

// Post represents the `eosio.forum::post` action.
type Message struct {
	Id       		MssgId 			`json:"message_id"`
	ParentId 		MssgId 			`json:"parent_id"`

	Language 		string 			`json:"languagemssg"`
	Header   		string 			`json:"headermssg"`
	Body     		string 			`json:"bodymssg"`

	Tags         	[]string 		`json:"tags"`
	JsonMetadata 	string   		`json:"jsonmetadata"`

	TokenProp     	PercentT     	`json:"tokenprop"`
	MaxPayout     	Asset     		`json:"max_payout,ommitempty"`
	Beneficiaries 	[]Beneficiary 	`json:"beneficiaries"`
	CuratorsPrcnt 	uint16      	`json:"curators_prcnt,ommitempty"`
	VestPayment   	bool     		`json:"vestpayment"`
}

type CreateMssg struct {
	Message
}

type Asset string

type PercentT uint16

type Beneficiary struct {
	Account		eos.Name	`json:"account"`
	Weight		PercentT	`json:"weight"`
}

type MssgId struct {
	Author		eos.AccountName	`json:"author"`
	Permlink	string			`json:"permlink"`
}
