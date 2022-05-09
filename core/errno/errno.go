package errno

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var _ Error = (*errno)(nil)

type (
	Error interface {
		GetCode() int
		GetHttpStatusCode() int
		GetErrMsg() string
		GetReason() string

		GetData() interface{}
		GetRawData() []byte

		// WithData return data if response is success
		WithData(data interface{}) Error

		FormatErrMsg(v ...interface{}) Error

		// WithReason return reason if response is failed
		WithReason(reason interface{}) Error

		// WithID set request id
		WithID(id string) Error

		// WithHttpStatusCode set http status code
		WithHttpStatusCode(httpStatusCode int) Error

		ToString() string

		ToBytes() []byte

		Error() string

		Render
	}

	errno struct {
		HttpStatusCode int         `json:"-"`
		Code           int         `json:"code"`             // code
		Msg            string      `json:"msg"`              // err msg
		Data           interface{} `json:"data"`             // data for success
		Reason         interface{} `json:"reason,omitempty"` // reason for err
		ID             string      `json:"id,omitempty"`     // global request id
		NowTime        int64       `json:"nowTime"`          // time.now
	}
)

func NewError(code int, errMsg string) Error {
	return errno{
		HttpStatusCode: http.StatusOK,
		Code:           code,
		Msg:            errMsg,
		Data:           make(map[string]interface{}),
		NowTime:        time.Now().Unix(),
	}
}

func (e errno) Error() string {
	return e.Msg
}

func (e errno) GetCode() int {
	return e.Code
}

func (e errno) GetHttpStatusCode() int {
	return e.HttpStatusCode
}

func (e errno) GetErrMsg() string {
	return e.Msg
}

func (e errno) GetReason() string {
	if v, ok := e.Reason.(string); ok {
		return v
	}

	return ""
}

func (e errno) GetData() interface{} {
	return e.Data
}

// GetRawData only support string/[]byte type data
func (e errno) GetRawData() []byte {
	if s, ok := e.Data.(string); ok {
		return []byte(s)
	}

	if s, ok := e.Data.([]byte); ok {
		return s
	}

	return []byte{}
}

func (e errno) FormatErrMsg(v ...interface{}) Error {
	e.Msg = fmt.Sprintf(e.Msg, v...)
	return e
}

func (e errno) WithData(data interface{}) Error {
	e.Data = data
	return e
}

func (e errno) WithID(id string) Error {
	e.ID = id
	return e
}

func (e errno) resetNowTime() Error {
	e.NowTime = time.Now().Unix()

	return e
}

// WithReason err reason
func (e errno) WithReason(reason interface{}) Error {
	// if reason is err, return err
	if v, ok := reason.(Error); ok {
		return v
	}

	if v, ok := reason.(error); ok {
		e.Reason = v.Error()

		return e
	}

	e.Reason = reason
	return e
}

// WithHttpStatusCode use net.http status code
func (e errno) WithHttpStatusCode(httpStatusCode int) Error {
	e.HttpStatusCode = httpStatusCode
	return e
}

func (e errno) ToString() string {
	return string(e.ToBytes())
}

func (e errno) ToBytes() []byte {
	if e.Data == nil {
		e.Data = make(map[string]interface{})
	}

	raw, err := json.Marshal(e)
	if err != nil {
		return nil
	}

	return raw
}
