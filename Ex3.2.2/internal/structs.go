package internal

type Config struct {
	Files []Data `yaml:"files"`
}

type Data struct {
	Filename string `yaml:"filename"`
	Substing string `yaml:"substring"`
}
