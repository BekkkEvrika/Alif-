package base

type Setting struct {
	AppName    string `json:"appName"`
	DBUser     string `json:"db-user"`
	DBPassword string `json:"db-password"`
	DBHost     string `json:"db-host"`
	DBPort     int    `json:"db-port"`
	DBName     string `json:"db-name"`
}
