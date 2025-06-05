// +build !nomattermost

package bridgemap

import (
	bmattermost "github.com/tzagim/matterbridge/bridge/mattermost"
)

func init() {
	FullMap["mattermost"] = bmattermost.New
}
