// +build !novk

package bridgemap

import (
	bvk "github.com/tzagim/matterbridge/bridge/vk"
)

func init() {
	FullMap["vk"] = bvk.New
}
