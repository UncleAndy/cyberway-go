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
	Id       		cyberway.MssgId 			`json:"message_id"`
	ParentId 		cyberway.MssgId 			`json:"parent_id"`

	Beneficiaries 	[]cyberway.Beneficiary 		`json:"beneficiaries"`

	TokenProp     	uint16	     				`json:"tokenprop"`
	VestPayment   	bool     					`json:"vestpayment"`

	Header   		string 						`json:"headermssg"`
	Body     		string 						`json:"bodymssg"`
	Language 		string 						`json:"languagemssg"`

	Tags         	[]string 					`json:"tags"`
	JsonMetadata 	string   					`json:"jsonmetadata"`

	CuratorsPrcnt 	uint16      				`json:"curators_prcnt,ommitempty"`
	MaxPayout     	string     					`json:"max_payout,ommitempty"`
}
