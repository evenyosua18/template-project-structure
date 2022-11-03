package database

// ListDatabase add here and yml file if there is more connection to database
type ListDatabase struct {
	Example Database `yaml:"example"`
}

type Database struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Address  string `yaml:"address"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
	Adapter  string `yaml:"adapter"`
}
