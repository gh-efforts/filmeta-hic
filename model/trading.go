package model

import (
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
)

type Trading struct {
	TipSet *types.TipSet  `json:"tipSet"`
	MCid   cid.Cid        `json:"mCid"`
	Msg    *types.Message `json:"msg"`
	// todo add more fields
}

func (trading *Trading) IsEmpty() bool {
	return trading.MCid.String() == ""
}

func (trading *Trading) GetTipSet() *types.TipSet {
	return trading.TipSet
}

func (trading *Trading) GetMCid() cid.Cid {
	return trading.MCid
}

func (trading *Trading) GetMsg() *types.Message {
	return trading.Msg
}
