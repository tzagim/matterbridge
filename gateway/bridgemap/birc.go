// +build !noirc

package bridgemap

import (
	birc "github.com/tzagim/matterbridge/bridge/irc"
)

func init() {
	FullMap["irc"] = birc.New
}
