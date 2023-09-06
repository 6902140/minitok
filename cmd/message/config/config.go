package config

import "minitok/config"

type ServiceConfigs struct {
	JWT       config.JWT       `mapstructure:"jwt" yaml:"jwt"`
	ETCD      config.ETCD      `mapstructure:"etcd" yaml:"etcd"`
	RPCServer config.RPCServer `mapstructure:"rpc_server" yaml:"rpc_server"`
	MongoDB   config.MongoDB   `mapstructure:"mongodb" yaml:"mongodb"`
}
