// +build !nomumble

package bridgemap

import (
	bmumble "github.com/tzagim/matterbridge/bridge/mumble"
)

func init() {
	FullMap["mumble"] = bmumble.New
}
