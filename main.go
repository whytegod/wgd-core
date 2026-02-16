cat > main.go << 'EOF'
package main

import (
	"fmt"

	"github.com/whytegod/wgd-core/genesis"
	"github.com/whytegod/wgd-core/ledger"
)

func main() {
	config := genesis.DefaultGenesis()
	ledgerInstance := ledger.NewLedger(config.InitialSupply)

	fmt.Println("Protocol:", config.ProtocolName)
	fmt.Println("Version:", config.ProtocolVersion)
	fmt.Println("Total Supply:", ledgerInstance.TotalSupply())
	fmt.Println("Treasury Balance:", ledgerInstance.TreasuryBalance())
}
EOF