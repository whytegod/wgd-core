package monetary

import "github.com/wgd/wgd-core/constitutional"

func Era(height uint64) uint64 {
    return height / constitutional.HalvingInterval
}