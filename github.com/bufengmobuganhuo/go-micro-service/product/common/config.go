package common

import (
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-plugins/config/source/consul/v2"
	"strconv"
)

func GetConsulConfig(host string, port int64, prefix string) (config.Config, error) {
	consulSource := consul.NewSource(
		consul.WithAddress(host+":"+strconv.FormatInt(port, 10)),
		consul.WithPrefix(prefix),
		consul.StripPrefix(true),
	)
	cfg, err := config.NewConfig()
	if err != nil {
		return nil, err
	}
	err = cfg.Load(consulSource)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
