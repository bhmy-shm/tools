package main

import (
    {{.importPackages}}
)

func main() {

	conf := gofkConf.New()

	gofks.Ignite("/v1").
		LoadWatch(conf).
		WireApply(
			wire.NewServiceWire(conf),
		).
		Mount({{.mount}}).
		Launch()
}