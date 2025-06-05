// +build !noapi

package bridgemap

import (
	"github.com/tzagim/matterbridge/bridge/api"
)

func init() {
	FullMap["api"] = api.New
}
