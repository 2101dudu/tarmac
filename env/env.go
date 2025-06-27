package env

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type CacheTimes struct {
	ShortCacheTime  *time.Duration
	MediumCacheTime *time.Duration
	LongCacheTime   *time.Duration
}

type Credentials struct {
	System     string
	Client     string
	Username   string
	Password   string
	DBUsername string
	DBPassword string
}

type Endpoints struct {
	APIEndpoint   string
	CacheEndpoint string
	DBEndpoint    string
}

type Logs struct {
	FilePath string
}

type Vars struct {
	Credentials *Credentials
	CacheTimes  *CacheTimes
	Endpoints   *Endpoints
	Logs        *Logs
}

func loadCacheTimes() *CacheTimes {
	loadDuration := func(envKey string) *time.Duration {
		if val, ok := os.LookupEnv(envKey); ok {
			if dur, err := time.ParseDuration(val); err == nil {
				return &dur
			} else {
				log.Fatal("Incorrect time syntax:", err)
			}
		}
		return nil
	}

	return &CacheTimes{
		ShortCacheTime:  loadDuration("SHORT_CACHE_TIME"),
		MediumCacheTime: loadDuration("MEDIUM_CACHE_TIME"),
		LongCacheTime:   loadDuration("LONG_CACHE_TIME"),
	}
}

func loadCredentials() *Credentials {
	return &Credentials{
		System:     os.Getenv("SYSTEM"),
		Client:     os.Getenv("CLIENT"),
		Username:   os.Getenv("USERNAME"),
		Password:   os.Getenv("PASSWORD"),
		DBUsername: os.Getenv("DB_USERNAME"),
		DBPassword: os.Getenv("DB_PASSWORD"),
	}
}

func loadEndpoints() *Endpoints {
	return &Endpoints{
		APIEndpoint:   os.Getenv("API_ENDPOINT"),
		CacheEndpoint: os.Getenv("CACHE_ENDPOINT"),
		DBEndpoint:    os.Getenv("DB_ENDPOINT"),
	}
}

func loadLogs() *Logs {
	return &Logs{
		FilePath: os.Getenv("LOG_FILE_PATH"),
	}
}

func LoadEnvFile() Vars {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return Vars{
		CacheTimes:  loadCacheTimes(),
		Credentials: loadCredentials(),
		Endpoints:   loadEndpoints(),
		Logs:        loadLogs(),
	}
}
