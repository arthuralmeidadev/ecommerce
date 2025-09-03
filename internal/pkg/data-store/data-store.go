package datadataStore

import (
	"database/sql"
	"ecommerce/pkg/deuterium"
	"fmt"
	"os"
	"sync"
)

var (
	store *dataStore
	once  sync.Once
)

type dataStore struct {
	logger *deuterium.Logger
	db     *sql.DB
}

func formatDSN() string {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUsr := os.Getenv("DB_USER")
	dbPw := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	return fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUsr, dbPw, dbName)
}

func instantiate() (*dataStore, error) {
	driver := os.Getenv("DB_DRIVER")
	dsn := formatDSN()
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, fmt.Errorf("Error on database open: %v", err)
	}

	return &dataStore{
		db: db,
		logger: &deuterium.Logger{
			Context: "Database",
		},
	}, nil
}

func GetDataStore() *dataStore {
	var err error
	once.Do(func() {
		store, err = instantiate()
		if err != nil {
			deuterium.GetLogger().Fatal(fmt.Sprintf("%v", err))
		}
	})

	return store
}

func (s *dataStore) checkConn() error {
	if s.db == nil {
		msg := "Missing database connection"
		s.logger.Error(msg)
		return fmt.Errorf("%s", msg)
	}

	return nil
}

func (s *dataStore) queryError(q string, err error) {
	s.logger.Error(fmt.Sprintf("Failed to execute %s: %v", q, err))
}

func (s *dataStore) Close() error {
	s.checkConn()
	return s.db.Close()
}
