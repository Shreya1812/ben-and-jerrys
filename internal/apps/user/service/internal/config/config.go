package config

import "github.com/Shreya1812/ben-and-jerrys/internal/configs"

type UserConfig struct {
	*configs.MongoDBConfig
	*configs.RedisConfig
	*configs.ElasticSearchConfig
}

func GetUserConfig(c *configs.Config) *UserConfig {
	return &UserConfig{
		MongoDBConfig:       c.MongoDBConfig,
		RedisConfig:         c.RedisConfig,
		ElasticSearchConfig: c.ElasticSearchConfig,
	}
}
