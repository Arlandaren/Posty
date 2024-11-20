package kafka

type Config struct {
	Brokers []string
}

func NewConfig(brokers []string) *Config {
	return &Config{
		Brokers: brokers,
	}
}
