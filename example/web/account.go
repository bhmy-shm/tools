package main

import (
	"github.com/bhmy-shm/gofks"
	gofkConf "github.com/bhmy-shm/gofks/core/config"
	"web/controls/account"
	"web/wire"
)

func main() {

	conf := gofkConf.New()

	gofks.Ignite("/v1").
		LoadWatch(conf).
		WireApply(
			wire.NewServiceWire(conf),
		).
		Mount(account.NewAccountController()).
		Launch()
}
