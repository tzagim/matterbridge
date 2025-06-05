// +build !nozulip

package bridgemap

import (
	bzulip "github.com/tzagim/matterbridge/bridge/zulip"
)

func init() {
	FullMap["zulip"] = bzulip.New
}
