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

type NotifyMessage interface {
	Get() (b []byte, err error)
}

type (
	Message struct {
		TipSet   *types.TipSet  `json:"TipSet"`
		MCid     cid.Cid        `json:"MCid"`
		Msg      *types.Message `json:"Msg"`
		Ret      *vm.ApplyRet   `json:"ApplyRet"`
		Implicit bool           `json:"Implicit"`
	}

	OneMessage struct {
		Msg       types.Message `json:"Msg"`
		Implicit  bool          `json:"Implicit"`
		IsSubCall bool          `json:"IsSubCall"`
		TipSet    *types.TipSet `json:"TipSet"`
	}
)

func (msg *Message) Get() (b []byte, err error) {
	return json.Marshal(msg)
}

func (msg *Message) IsImplicit() bool {
	return msg.Implicit
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

//func (msg *Message) Marshal() ([]byte, error) {
//	return json.Marshal(msg)
//}

func (msg *Message) GetRet() *vm.ApplyRet {
	return msg.Ret
}

func (msg *Message) GetImplicit() bool {
	return msg.Implicit
}

func (one *OneMessage) Get() (b []byte, err error) {
	return json.Marshal(one)
}

func UnmarshalMsg(b []byte) (msg *Message, err error) {
	err = json.Unmarshal(b, &msg)
	return
}
