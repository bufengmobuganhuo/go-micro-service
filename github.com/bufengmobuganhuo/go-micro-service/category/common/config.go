package common

import (
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-plugins/config/source/consul/v2"
	"strconv"
)

// GetConsulConfig 设置配置中心
func GetConsulConfig(host string, port int64, prefix string) (config.Config, error) {
	consulSource := consul.NewSource(
		// 设置配置中心地址
		consul.WithAddress(host+":"+strconv.FormatInt(port, 10)),
		// 配置前缀，不设置的话默认为: /micro/config
		consul.WithPrefix(prefix),
		// 设置是否移除前缀
		consul.StripPrefix(true),
	)
	config, err := config.NewConfig()
	if err != nil {
		return config, err
	}
	// 加载配置
	err = config.Load(consulSource)
	return config, err
}
