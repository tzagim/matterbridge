// +build !nowhatsapp
// +build !whatsappmulti

package bridgemap

import (
	bwhatsapp "github.com/tzagim/matterbridge/bridge/whatsapp"
)

func init() {
	FullMap["whatsapp"] = bwhatsapp.New
}
