package common

import "github.com/micro/go-micro/v2/config"

type MysqlConfig struct {
	Host     string `json:"host"`
	Port     int64  `json:"port"`
	User     string `json:"user"`
	Pwd      string `json:"pwd"`
	Database string `json:"database"`
}

// GetMysqlFromConsul 从配置中心获取MySQL的连接配置
func GetMysqlFromConsul(config config.Config, path ...string) *MysqlConfig {
	mysqlConfig := &MysqlConfig{}
	err := config.Get(path...).Scan(mysqlConfig)
	if err != nil {
		return nil
	}
	return mysqlConfig
}
