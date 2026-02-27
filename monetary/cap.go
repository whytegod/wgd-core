package monetary

import "github.com/wgd/wgd-core/constitutional"

func EnforceHardCap(totalMinted, reward uint64) uint64 {
    if totalMinted >= constitutional.HardCap {
        return 0
    }

    if totalMinted+reward > constitutional.HardCap {
        return constitutional.HardCap - totalMinted
    }

    return reward
}