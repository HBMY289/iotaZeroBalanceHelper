package reclaim

import (
	"encoding/json"
)

var Reclaims ReclaimAddresses

func init() {
	err := json.Unmarshal(reclaimJSON, &Reclaims)
	if err != nil {
		panic(err)
	}
}