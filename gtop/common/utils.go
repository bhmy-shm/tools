package common

import (
	"strconv"
	"strings"
)

func AddressSplitIpPort(address string) (string, int) {

	sp := strings.Split(address, ":")
	if len(sp) < 2 {
		return "", 0
	}
	inter, _ := strconv.Atoi(sp[1])
	return sp[0], inter
}
