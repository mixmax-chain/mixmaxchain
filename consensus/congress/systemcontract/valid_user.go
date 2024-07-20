package systemcontract

import (
	//	"bytes"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/congress/vmcaller"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"math"
	"math/big"
	//	"strings"
	//	"sort"
)

type ValidatorUser struct {
	abi          abi.ABI
	contractAddr common.Address
}

func NewValidatorUser() *ValidatorUser {
	return &ValidatorUser{
		abi:          Valid_user_abi2,
		contractAddr: ValidatorsUserAddr,
	}
}

func (v *ValidatorUser) checkValidUser(val common.Address, statedb *state.StateDB, header *types.Header, chainContext core.ChainContext, config *params.ChainConfig) (uint, error) {
	method := "checkValidUser"
	data, err := v.abi.Pack(method, val)
	if err != nil {
		log.Error("Can't pack data for GetValidatorInfo", "error", err)
		return 0, err
	}
	msg := vmcaller.NewLegacyMessage(header.Coinbase, &v.contractAddr, 0, new(big.Int), math.MaxUint64, new(big.Int), data, false)

	// use parent
	result, err := vmcaller.ExecuteMsg(msg, statedb, header, chainContext, config)
	if err != nil {
		return 0, err
	}
	// unpack data
	ret, err := v.abi.Unpack(method, result)
	if err != nil {
		return 0, err
	}
	feeAddr, ok := ret[0].(uint)
	if !ok {
		return 0, errors.New("invalid output")
	}
	return feeAddr, nil
}
