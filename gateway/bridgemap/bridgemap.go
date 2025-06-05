package bridgemap

import (
	"github.com/tzagim/matterbridge/bridge"
)

var (
	FullMap           = map[string]bridge.Factory{}
	UserTypingSupport = map[string]struct{}{}
)
