// +build !nonctalk

package bridgemap

import (
	btalk "github.com/tzagim/matterbridge/bridge/nctalk"
)

func init() {
	FullMap["nctalk"] = btalk.New
}
