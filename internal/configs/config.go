package configs

type Config struct {
	*MongoDBConfig
	*RedisConfig
	*ElasticSearchConfig
	*JWTConfig
}

type MongoDBConfig struct {
	Host         string
	Port         string
	DatabaseName string
}

type JWTConfig struct {
	JWTSecret            string
	JwtExpirationMinutes int
}

type RedisConfig struct{}

type ElasticSearchConfig struct{}
