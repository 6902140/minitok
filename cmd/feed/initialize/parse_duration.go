package initialize

import (
	"minitok/cmd/feed/global"
	"time"
)

func ParseDuration() error {
	var err error
	global.ExpireDurationNullKey, err = time.ParseDuration(global.Configs.CacheExpire.NullKey)
	if err != nil {
		return err
	}
	global.ExpireDurationUserBaseInfo, err = time.ParseDuration(global.Configs.CacheExpire.UserBaseInfo)
	if err != nil {
		return err
	}
	global.ExpireDurationVideoBaseInfo, err = time.ParseDuration(global.Configs.CacheExpire.VideoBaseInfo)
	if err != nil {
		return err
	}
	return nil
}
