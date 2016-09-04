package messages

import "meshwalker.com/mws/pkg/types"

func ErrRethinkDb() *types.ErrMsg {
	return &types.ErrMsg{
		Status:	"error",
		Message: "Internal server error",
	}
}