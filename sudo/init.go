package sudo

import eos "github.com/UncleAndy/cyberway-go"

func init() {
	eos.RegisterAction(AN("eosio.wrap"), ActN("exec"), Exec{})
}

var AN = eos.AN
var ActN = eos.ActN
