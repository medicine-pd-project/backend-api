package configs

type Server struct {
	Host string `envconfig:"HOST" required:"true"`
	Port int    `envconfig:"PORT" required:"true"`
}
