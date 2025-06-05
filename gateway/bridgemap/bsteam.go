// +build !nosteam

package bridgemap

import (
	bsteam "github.com/tzagim/matterbridge/bridge/steam"
)

func init() {
	FullMap["steam"] = bsteam.New
}
