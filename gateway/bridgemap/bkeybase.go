// +build !nokeybase

package bridgemap

import (
	bkeybase "github.com/tzagim/matterbridge/bridge/keybase"
)

func init() {
	FullMap["keybase"] = bkeybase.New
}
