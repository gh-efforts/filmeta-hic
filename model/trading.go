package model

import (
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	jsoniter "github.com/json-iterator/go"
)

var (
	// replace encoding/json
	json = jsoniter.ConfigCompatibleWithStandardLibrary
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

func (trading *Trading) Marshal() ([]byte, error) {
	return json.Marshal(trading)
}

func UnmarshalMsg(b []byte) (trading *Trading, err error) {
	err = json.Unmarshal(b, &trading)
	return
}
