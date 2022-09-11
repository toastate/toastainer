package nodes

import (
	"github.com/toastate/toastcloud/internal/config"
)

func Init() error {
	if !config.NodeDiscovery {
		return nil
	}

	err := initAddrs()
	if err != nil {
		return err
	}

	startDNSNodeLookupRoutine()

	return nil
}
