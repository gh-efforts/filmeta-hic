package random

import (
	"math/rand"
	"reflect"
	"strings"
	"time"
	"unsafe"
)

type Charset string

// Charsets
const (
	Uppercase    Charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Lowercase    Charset = "abcdefghijklmnopqrstuvwxyz"
	Alphabetic           = Uppercase + Lowercase
	Numeric      Charset = "0123456789"
	Alphanumeric         = Alphabetic + Numeric
	Symbols      Charset = "`" + `~!@#$%^&*()-_+={}[]|\;:"<>,./?`
	Hex                  = Numeric + "abcdef"
)

type Random struct {
}

var (
	global = New()
)

func New() *Random {
	rand.Seed(time.Now().UnixNano())
	return new(Random)
}

func (r *Random) String(length uint8, items ...string) string {
	charset := strings.Join(items, "")
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Int63()%int64(len(charset))]
	}
	return string(b)
}

func String(length uint8, charsets ...Charset) string {
	if len(charsets) == 0 {
		charsets = append(charsets, Alphanumeric)
	}
	var (
		items []string
	)
	for _, charset := range charsets {
		items = append(items, string(charset))
	}
	return global.String(length, items...)
}

// GetRandomNumber  get numbers for random
func GetRandomNumber(l int) string {
	str := "0123456789"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return BytesToString(result)
}

// BytesToString 0 copy from []byte to string
func BytesToString(b []byte) (s string) {
	_bptr := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	_sptr := (*reflect.StringHeader)(unsafe.Pointer(&s))
	_sptr.Data = _bptr.Data
	_sptr.Len = _bptr.Len
	return s
}
