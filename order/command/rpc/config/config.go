package config

import (
	rpc_config "github.com/yakaa/grpcx/config"
)

type Config struct {
	RpcServerConfig *rpc_config.ServiceConf
	Mode            string `json:"mode"`
	Mysql           struct {
		DataSource string
		Table      struct {
			Order string
		}
	}
	Redis struct {
		DataSource string
		Auth       string
	}
	RabbitMq struct {
		DataSource  string
		VirtualHost string
		QueueName   string
	}
}
