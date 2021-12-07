package config

//userConfig is a structure for main service configuration
type UserConfig struct {
	ServiceName string `yaml:"service_name"`
	HttpPort    string `yaml:"http_port"`
	BaseURL     string `yaml:"base_url"`
}
