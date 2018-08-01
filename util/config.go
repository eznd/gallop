package util

type Config struct {
	RootPath string `env:"GALLOP_ROOT" default:"/go/src/eznd/gallop" required:"true"`
	Port     string `env:"GALLOP_PORT" default:"8080" required:"true"`
}
