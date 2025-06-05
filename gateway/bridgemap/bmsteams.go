// +build !nomsteams

package bridgemap

import (
	bmsteams "github.com/tzagim/matterbridge/bridge/msteams"
)

func init() {
	FullMap["msteams"] = bmsteams.New
}
