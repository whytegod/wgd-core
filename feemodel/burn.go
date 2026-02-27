package feemodel

import "github.com/wgd/wgd-core/constitutional"

func SplitFees(total uint64) (burn uint64, validators uint64) {
    burn = total * constitutional.BurnRatio / 100
    validators = total - burn
    return
}