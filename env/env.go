package env

import (
	"os"
	"strconv"
)

func GetInfluxURL() string {
	return os.Getenv("INFLUX_URL")
}

func GetEmqxURL() string {
	return os.Getenv("EMQX_URL")
}

func GetEmqxAppID() string {
	return os.Getenv("EMQX_APP_ID")
}

func GetEmqxAppSecret() string {
	return os.Getenv("EMQX_APP_SECRET")
}

func GetAppPort() int {
	p, err := strconv.ParseInt(os.Getenv("APP_PORT"), 10, 32)
	if err != nil {
		return 8060
	}
	return int(p)
}
