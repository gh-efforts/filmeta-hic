package envx

import (
	"github.com/bitrainforest/filmeta-hic/core/assert"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/env"
)

// default env config

type FixedEnv struct {
	ConfigETCD string `json:"CONFIG_ETCD"` //   "etcd://127.0.0.1:2379"
	ConfigPath string `json:"CONFIG_PATH"`
	HttpAddr   string `json:"GAPI_ADDR"` //   default ":8080"
	GinMODE    string `json:"GIN_MODE"`  //  default "debug"
	// todo remember add more env config if needed
}

var (
	fixedEnv = &FixedEnv{HttpAddr: ":8088", GinMODE: "debug"}
)

func init() {
	MustSetup()
}

func GetEnvs() FixedEnv {
	if fixedEnv == nil {
		return FixedEnv{}
	}

	return *fixedEnv
}

func Setup() error {
	c := config.New(
		config.WithSource(env.NewSource()))

	if err := c.Load(); err != nil {
		return err
	}

	if err := c.Scan(fixedEnv); err != nil {
		return err
	}

	return nil
}

func MustSetup() {
	if err := Setup(); err != nil {
		assert.CheckErr(err)
	}
}
