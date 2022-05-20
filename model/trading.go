package model

import (
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/chain/vm"
	"github.com/ipfs/go-cid"
	jsoniter "github.com/json-iterator/go"
)

var (
	// replace encoding/json
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

type Message struct {
	TipSet   *types.TipSet  `json:"tipSet"`
	MCid     cid.Cid        `json:"mCid"`
	Msg      *types.Message `json:"msg"`
	Ret      *vm.ApplyRet   `json:"ApplyRet"`
	Implicit bool           `json:"implicit"`
}

func (msg *Message) Defined() bool {
	return msg.MCid.Defined()
}

func (msg *Message) GetTipSet() *types.TipSet {
	return msg.TipSet
}

func (msg *Message) GetMCid() cid.Cid {
	return msg.MCid
}

func (msg *Message) GetMsg() *types.Message {
	return msg.Msg
}

func (msg *Message) Marshal() ([]byte, error) {
	return json.Marshal(msg)
}

func (msg *Message) GetRet() *vm.ApplyRet {
	return msg.Ret
}

func (msg *Message) GetImplicit() bool {
	return msg.Implicit
}

func UnmarshalMsg(b []byte) (msg *Message, err error) {
	err = json.Unmarshal(b, &msg)
	return
}
