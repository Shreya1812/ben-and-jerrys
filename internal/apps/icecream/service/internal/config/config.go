package config

import "github.com/Shreya1812/ben-and-jerrys/internal/configs"

type IceCreamConfig struct {
	*configs.MongoDBConfig
	*configs.RedisConfig
	*configs.ElasticSearchConfig
}

func GetIceCreamConfig(c *configs.Config) *IceCreamConfig {
	return &IceCreamConfig{
		MongoDBConfig:       c.MongoDBConfig,
		RedisConfig:         c.RedisConfig,
		ElasticSearchConfig: c.ElasticSearchConfig,
	}
}
