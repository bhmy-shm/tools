package errorx

import "encoding/json"

type (
	ErrCode uint64
	Status  struct {
		Code    uint64 `json:"Code,omitempty"`
		Reason  string `json:"reason,omitempty"`
		Message string `json:"message,omitempty"`
	}
)

const (
	errCodeBase = 0x00000000
)

const (
	ErrCodeSuccessful ErrCode = errCodeBase + iota
	ErrCodeSSHConnectionFailed
	ErrCodeSSHConnectionThreeLimit
)

var errMap map[ErrCode]*Status

func init() {
	errMap = map[ErrCode]*Status{
		ErrCodeSuccessful:              {Code: uint64(ErrCodeSuccessful), Message: "Ok"},
		ErrCodeSSHConnectionFailed:     {Code: uint64(ErrCodeSSHConnectionFailed), Message: "ssh connection failed"},
		ErrCodeSSHConnectionThreeLimit: {Code: uint64(ErrCodeSSHConnectionThreeLimit), Message: "ssh limit is reached 3 times"},
	}
}

func (e ErrCode) Error() string {
	if v, ok := errMap[e]; ok {
		b, _ := json.Marshal(v)
		return string(b)
	}
	return ""
}

func (e ErrCode) SetReason(r string) error {
	errMap[e].Reason = r
	return e
}
