// +build whatsappmulti

package bridgemap

import (
	bwhatsapp "github.com/tzagim/matterbridge/bridge/whatsappmulti"
)

func init() {
	FullMap["whatsapp"] = bwhatsapp.New
}
