package goloco_repo

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"net/url"
)

type GormStore struct {
	Host     string
	Port     int32
	User     string
	Password string
	Dbname   string
}

func (store *GormStore) Connect() *gorm.DB {
	dsn := url.URL{
		User:     url.UserPassword(store.User, store.Password),
		Scheme:   "postgres",
		Host:     fmt.Sprintf("%s:%d", store.Host, store.Port),
		Path:     store.Dbname,
		RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
	}

	log.Printf("%v", dsn.String())

	db, err := gorm.Open("postgres", dsn.String())
	if err != nil {
		log.Printf("%v", err)
	}
	db.AutoMigrate(&LocationModel{})

	return db
}
