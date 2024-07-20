package systemcontract

import (
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/consensus/congress/vmcaller"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"math/big"
)

type SysDistribute struct {
	abi          abi.ABI
	contractAddr common.Address
}

func NewSysDistribute() *SysDistribute {
	return &SysDistribute{
		abi:          SysDistributeRateABI3,
		contractAddr: SysDistributeRateAddr,
	}
}

func (v *SysDistribute) GetBurnRate(statedb *state.StateDB, header *types.Header, chainContext core.ChainContext, config *params.ChainConfig) (int64, error) {
	method := "get_burn_rate"
	data, err := v.abi.Pack(method)
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
	rate, ok := ret[0].(*big.Int)
	if !ok {
		return 0, errors.New("invalid output")
	}
	return rate.Int64(), nil
}

func (v *SysDistribute) GetDistributeRate(statedb *state.StateDB, header *types.Header, chainContext core.ChainContext, config *params.ChainConfig) (int64, error) {
	method := "get_node_distribute_rate"
	data, err := v.abi.Pack(method)
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
	rate, ok := ret[0].(*big.Int)
	if !ok {
		return 0, errors.New("invalid output")
	}
	return rate.Int64(), nil
}

func (v *SysDistribute) GetAddress() common.Address {
	return v.contractAddr
}
