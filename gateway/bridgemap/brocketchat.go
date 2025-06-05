// +build !norocketchat

package bridgemap

import (
	brocketchat "github.com/tzagim/matterbridge/bridge/rocketchat"
)

func init() {
	FullMap["rocketchat"] = brocketchat.New
}
