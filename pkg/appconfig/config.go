package appconfig

type Config struct {
	Shift int `yaml:"shift"`
	Tags  map[uint32]string
}
