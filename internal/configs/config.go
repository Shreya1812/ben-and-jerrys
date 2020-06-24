package configs

type Config struct {
	*MongoDBConfig
	*RedisConfig
	*ElasticSearchConfig
}

type MongoDBConfig struct {
	Host         string
	Port         string
	DatabaseName string
}

type RedisConfig struct{}

type ElasticSearchConfig struct{}
