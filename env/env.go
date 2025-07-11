package env

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type SoapService struct {
	System   string
	Client   string
	Username string
	Password string

	APIEndpoint string
}

type CacheTimes struct {
	ShortCacheTime  time.Duration
	MediumCacheTime time.Duration
	LongCacheTime   time.Duration
}

type Credentials struct {
	DBUsername          string
	DBPassword          string
	AdminHashedPassword string
}

type Endpoints struct {
	CacheEndpoint string
	DBEndpoint    string
}

type Logs struct {
	FilePath string
}

type EmailCredentials struct {
	APIKey              string
	AgencyEmail         string
	InternalAgencyEmail string
}

type Vars struct {
	SOAPServices     []*SoapService
	Credentials      *Credentials
	CacheTimes       *CacheTimes
	Endpoints        *Endpoints
	Logs             *Logs
	EmailCredentials *EmailCredentials
}

func loadSoapServiceSoltropico() *SoapService {
	return &SoapService{
		System:      os.Getenv("SYSTEM_SOLTROPICO"),
		Client:      os.Getenv("CLIENT_SOLTROPICO"),
		Username:    os.Getenv("USERNAME_SOLTROPICO"),
		Password:    os.Getenv("PASSWORD_SOLTROPICO"),
		APIEndpoint: os.Getenv("API_ENDPOINT_SOLTROPICO"),
	}
}
func loadSoapServiceEgotravel() *SoapService {
	return &SoapService{
		System:      os.Getenv("SYSTEM_EGOTRAVEL"),
		Client:      os.Getenv("CLIENT_EGOTRAVEL"),
		Username:    os.Getenv("USERNAME_EGOTRAVEL"),
		Password:    os.Getenv("PASSWORD_EGOTRAVEL"),
		APIEndpoint: os.Getenv("API_ENDPOINT_EGOTRAVEL"),
	}
}
func loadSoapServiceViajarTours() *SoapService {
	return &SoapService{
		System:      os.Getenv("SYSTEM_VIAJARTOURS"),
		Client:      os.Getenv("CLIENT_VIAJARTOURS"),
		Username:    os.Getenv("USERNAME_VIAJARTOURS"),
		Password:    os.Getenv("PASSWORD_VIAJARTOURS"),
		APIEndpoint: os.Getenv("API_ENDPOINT_VIAJARTOURS"),
	}
}

func loadSoapServices() []*SoapService {
	var services []*SoapService
	// services = append(services, loadSoapServiceEgotravel()) // -- server is down
	services = append(services, loadSoapServiceSoltropico())
	services = append(services, loadSoapServiceViajarTours())
	return services
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
		ShortCacheTime:  *loadDuration("SHORT_CACHE_TIME"),
		MediumCacheTime: *loadDuration("MEDIUM_CACHE_TIME"),
		LongCacheTime:   *loadDuration("LONG_CACHE_TIME"),
	}
}

func loadCredentials() *Credentials {
	return &Credentials{
		DBUsername:          os.Getenv("DB_USERNAME"),
		DBPassword:          os.Getenv("DB_PASSWORD"),
		AdminHashedPassword: os.Getenv("ADMIN_HASHED_PASSWORD"),
	}
}

func loadEndpoints() *Endpoints {
	return &Endpoints{
		CacheEndpoint: os.Getenv("CACHE_ENDPOINT"),
		DBEndpoint:    os.Getenv("DB_ENDPOINT"),
	}
}

func loadLogs() *Logs {
	return &Logs{
		FilePath: os.Getenv("LOG_FILE_PATH"),
	}
}

func loadEmailCredentials() *EmailCredentials {
	return &EmailCredentials{
		APIKey:              os.Getenv("EMAIL_API_KEY"),
		AgencyEmail:         os.Getenv("AGENCY_EMAIL"),
		InternalAgencyEmail: os.Getenv("INTERNAL_AGENCY_EMAIL"),
	}
}

func LoadEnvFile() Vars {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return Vars{
		SOAPServices:     loadSoapServices(),
		CacheTimes:       loadCacheTimes(),
		Credentials:      loadCredentials(),
		Endpoints:        loadEndpoints(),
		Logs:             loadLogs(),
		EmailCredentials: loadEmailCredentials(),
	}
}
