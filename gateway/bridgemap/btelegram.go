// +build !notelegram

package bridgemap

import (
	btelegram "github.com/tzagim/matterbridge/bridge/telegram"
)

func init() {
	FullMap["telegram"] = btelegram.New
}
