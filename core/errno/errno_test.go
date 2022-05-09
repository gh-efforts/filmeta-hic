package errno

import (
	"github.com/stretchr/testify/assert"

	"reflect"
	"testing"
)

func TestNoDataError(t *testing.T) {
	errMsg := "Parameter error"
	badCode := -1
	err := NewError(badCode, errMsg)
	assert.Equal(t, err.GetErrMsg(), errMsg)
	assert.Equal(t, err.GetCode(), badCode)

	data := reflect.ValueOf(err.GetData())
	assert.Equal(t, data.Kind(), reflect.Map)
	assert.Equal(t, len(data.MapKeys()), 0)
}

func TestDataError(t *testing.T) {

	type Data struct {
		Name string
	}
	data := Data{Name: "remember"}
	err := NewError(0, "成功").WithData(data)
	assert.Equal(t, err.GetCode(), 0)
	assert.Equal(t, err.GetData().(Data).Name, "remember")
}
