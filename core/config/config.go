package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/bitrainforest/filmeta-hic/core/config/etcdx"

	"github.com/go-kratos/kratos/v2/config/file"

	"github.com/bitrainforest/filmeta-hic/core/assert"
	clientv3 "go.etcd.io/etcd/client/v3"

	"github.com/go-kratos/kratos/v2/config"
)

type (
	Schema   string
	SchemaFn func() (schema Schema, host string)
)

const (
	Etcd Schema = "etcd"
	File Schema = "file"
	// todo @remember support more config schema if needed
	// todo @remember add test for config schema
)

var (
	LoadConfigErr = errors.New("LoadConfigAndInitData err")
)

func LoadConfigAndInitData(prefix string, schemaFns ...SchemaFn) (config.Config, error) {
	var (
		conf config.Config
	)

Loop:
	for _, configSchema := range schemaFns {
		schema, host := configSchema()
		if host == "" || schema == "" {
			continue
		}
		switch schema {
		case Etcd:
			u, err := url.Parse(host)
			if err != nil {
				fmt.Println("Etcd err:", err)
				continue
			}
			prefix = "/" + strings.TrimLeft(prefix, "/")
			cli, err := clientv3.New(clientv3.Config{
				Endpoints: []string{u.Host},
				Username:  u.Query().Get("username"),
				Password:  u.Query().Get("password"),
			})
			if err != nil {
				fmt.Println("etcd clientv3.New err:", err)
				continue
			}
			source, err := etcdx.New(cli, etcdx.WithPrefix(true), etcdx.WithPath(prefix),
				etcdx.WithPrefix(true), etcdx.WithEnableFormatDotNotation(true))
			if err != nil {
				fmt.Println("etcdx.New err:", err)
				continue
			}
			conf = config.New(config.WithSource(source))
			if conf != nil {
				break Loop
			}
		case File:
			// if Schema is File,host means file path
			conf = config.New(config.WithSource(file.NewSource(host)))
			if conf != nil {
				break Loop
			}
		}
	}

	if conf == nil {
		assert.CheckErr(LoadConfigErr)
	}
	if err := conf.Load(); err != nil {
		assert.CheckErr(fmt.Errorf("load conf err:%v", err))
	}
	return conf, nil
}

func ScanConfValue(confSource config.Config, key string, item interface{}) error {
	value := confSource.Value(key)
	if val, err := value.String(); err == nil {
		return json.Unmarshal([]byte(val), item)
	}
	// if err !=nil, it's means that the val not string、int、[]byte or string() interface
	return value.Scan(item)
}
