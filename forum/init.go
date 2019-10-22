package forum

import eos "github.com/UncleAndy/cyberway-go"

func init() {
	eos.RegisterAction(ForumAN, ActN("createmssg"), MessageAction{})
	eos.RegisterAction(ForumAN, ActN("updatemssg"), MessageAction{})
	eos.RegisterAction(ForumAN, ActN("clnproposal"), CleanProposal{})
	eos.RegisterAction(ForumAN, ActN("expire"), Expire{})
	eos.RegisterAction(ForumAN, ActN("propose"), Propose{})
	eos.RegisterAction(ForumAN, ActN("status"), Status{})
	eos.RegisterAction(ForumAN, ActN("unpost"), UnPost{})
	eos.RegisterAction(ForumAN, ActN("unvote"), UnVote{})
	eos.RegisterAction(ForumAN, ActN("vote"), Vote{})
}

var AN = eos.AN
var PN = eos.PN
var ActN = eos.ActN

var ForumAN = AN("gls.publish")
