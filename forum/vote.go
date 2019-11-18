package forum

import (
	eos "github.com/UncleAndy/cyberway-go"
)

// NewVote is an action representing a simple vote to be broadcast
// through the chain network.
func NewVote(permType, voter, contentAuthor, contentPermlink string, weight int) *eos.Action {
	a := &eos.Action{
		Account: ForumAN,
		Name:    ActN("upvote"),
		Authorization: []eos.PermissionLevel{
			{Actor: eos.AccountName(voter), Permission: eos.PermissionName(permType)},
		},
		ActionData: eos.NewActionData(Vote{
			Voter:        	eos.AccountName(voter),
			MessageId: 		eos.MssgId{
				Author:   eos.AccountName(contentAuthor),
				Permlink: contentPermlink,
			},
			Weight:			uint16(weight),
		}),
	}
	return a
}

// Vote represents the `eosio.forum::vote` action.
type Vote struct {
	Voter        eos.AccountName `json:"voter"`
	MessageId	 eos.MssgId		 `json:"message_id"`
	Weight		 uint16			 `json:"weight"`
}
