package forum

import (
	eos "github.com/UncleAndy/cyberway-go"
)

// NewMessage is an action representing a simple message to be posted
// through the chain network.
func CreateMessage(permType string, msg *Message) *eos.Action {
	a := &eos.Action{
		Account: ForumAN,
		Name:    ActN("createmssg"),
		Authorization: []eos.PermissionLevel{
			{Actor: msg.Id.Author, Permission: eos.PermissionName(permType)},
		},
		ActionData: eos.NewActionData(msg),
	}
	return a
}

func UpdateMessage(permType string, msg *Message) *eos.Action {
	a := &eos.Action{
		Account: ForumAN,
		Name:    ActN("updatemssg"),
		Authorization: []eos.PermissionLevel{
			{Actor: msg.Id.Author, Permission: eos.PermissionName(permType)},
		},
		ActionData: eos.NewActionData(msg),
	}
	return a
}

// Post represents the `eosio.forum::post` action.
type Message struct {
	Id       		eos.MssgId 					`json:"message_id"`
	ParentId 		eos.MssgId 					`json:"parent_id"`

	Beneficiaries 	[]eos.Beneficiary 			`json:"beneficiaries"`

	TokenProp     	uint16	     				`json:"tokenprop"`
	VestPayment   	bool     					`json:"vestpayment"`

	Header   		string 						`json:"headermssg"`
	Body     		string 						`json:"bodymssg"`
	Language 		string 						`json:"languagemssg"`

	Tags         	[]string 					`json:"tags"`
	JsonMetadata 	string   					`json:"jsonmetadata"`

	CuratorsPrcnt 	uint16      				`json:"curators_prcnt,ommitempty"`
	MaxPayout     	string 						`json:"max_payout,ommitempty"`
}
