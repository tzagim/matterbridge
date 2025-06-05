// +build !noxmpp

package bridgemap

import (
	bxmpp "github.com/tzagim/matterbridge/bridge/xmpp"
)

func init() {
	FullMap["xmpp"] = bxmpp.New
}
