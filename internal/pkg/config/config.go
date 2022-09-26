package config

type Environment struct {
	DatabaseDNS string `env:"DATABASE_DNS,required=false"`
}
