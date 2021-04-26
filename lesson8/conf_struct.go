package lesson8

type Conf struct {
	Host     string `env:"HOST"`
	Port     int    `env:"PORT"`
	Protocol string `env:"PROTOCOL"`
}
