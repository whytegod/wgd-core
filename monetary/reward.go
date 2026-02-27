package monetary

import "github.com/wgd/wgd-core/constitutional"

func BaseReward(height uint64) uint64 {
    era := Era(height)

    if era >= constitutional.MaxEras {
        return 0
    }

    return constitutional.InitialReward >> era
}