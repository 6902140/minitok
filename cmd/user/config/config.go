package config

import "minitok/config"

type ServiceConfigs struct {
	MySQL          config.MySQL          `mapstructure:"mysql" yaml:"mysql"`
	Redis          config.Redis          `mapstructure:"redis" yaml:"redis"`
	JWT            config.JWT            `mapstructure:"jwt" yaml:"jwt"`
	ETCD           config.ETCD           `mapstructure:"etcd" yaml:"etcd"`
	RPCServer      config.RPCServer      `mapstructure:"rpc_server" yaml:"rpc_server"`
	StaticResource config.StaticResource `mapstructure:"static_resource" yaml:"static_resource"`
	CacheExpire    config.CacheExpire    `mapstructure:"cache_expire" yaml:"cache_expire"`
	MongoDB        config.MongoDB        `mapstructure:"mongodb" yaml:"mongodb"`
}
