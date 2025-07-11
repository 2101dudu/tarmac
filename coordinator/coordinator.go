package coordinator

import (
	"tarmac/cache"
	"tarmac/db"
	"tarmac/mail"
	"tarmac/wsdl"
)

type Client struct {
	WSDLService  *wsdl.ServiceList
	CacheService *cache.Service
	DBService    *db.Service
	MailService  *mail.Service
}

type Service struct {
	wsdlService  *wsdl.ServiceList
	cacheService *cache.Service
	dbService    *db.Service
	mailService  *mail.Service
}

func (c *Client) NewCoordinatorService() *Service {
	return &Service{
		wsdlService:  c.WSDLService,
		cacheService: c.CacheService,
		dbService:    c.DBService,
		mailService:  c.MailService,
	}
}

func getData[T any](s *Service, cacheKey, collectionName, id string, fetchFunc func() (T, error)) (T, bool, bool, error) {
	// defer logger.Log.TrackTime()()
	var empty T

	// Check cache first
	cachedData := cache.CheckCacheHit[T](s.cacheService, cacheKey)
	if cachedData != nil {
		return *cachedData, false, false, nil
	}

	// Check DB next
	dbData, outdated := db.CheckDBHit[T](s.dbService, collectionName, id)
	if dbData != nil {
		return *dbData, outdated, true, nil
	}

	// Not in cache or DB, fetch from API
	newData, err := fetchFunc()
	if err != nil {
		return empty, true, true, err
	}

	return newData, true, true, nil
}
