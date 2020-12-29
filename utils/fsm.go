package utils

import (
	"errors"
	"fmt"
)

const stateLength = 10

type Exchange [stateLength][stateLength]bool
type ExchangeUnit [2]int32
type ExchangeUnitList [100]ExchangeUnit

type StateFsm struct {
	exchange Exchange
}

func (sf *StateFsm) ConfigParse(list ExchangeUnitList) {
	// 解析状态机配置
	for i := 0; i < stateLength; i++ {
		for j := 0; j < stateLength; j++ {
			sf.exchange[i][j] = false
		}
	}
	for _, ex := range list {
		sf.exchange[ex[0]][ex[1]] = true
	}
}

func (sf *StateFsm) StateCheck(sfrom int32, sto int32) (bool, error) {
	// 状态变换检查
	if sf.exchange[sfrom][sto] {
		return sf.exchange[sfrom][sto], nil
	} else {
		return sf.exchange[sfrom][sto], errors.New(fmt.Sprintf("Exchange Error! from:%d to %d", sfrom, sto))
	}
}
