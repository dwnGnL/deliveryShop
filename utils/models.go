package utils

type Config struct {
	Key     string `json:"key"`
	DBUri   string `json:"db"`
	Timeout int    `json:"timeout"`
	PortRun int    `json:"port"`
	Realm   string `json:"realm"`
}
