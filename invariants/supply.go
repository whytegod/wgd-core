package invariants

import "github.com/wgd/wgd-core/constitutional"

func HardCapInvariant(totalMinted uint64) bool {
    return totalMinted <= constitutional.HardCap
}