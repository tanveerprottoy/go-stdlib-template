package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"sync/atomic"

	_ "github.com/lib/pq"
	"github.com/tanveerprottoy/stdlib-go-template/pkg/config"
)

var (
	instance    *Client
	once        sync.Once
	mu          sync.Mutex
	initialized uint32
)

type Client struct {
	DB *sql.DB
}

func GetInstance() *Client {
	once.Do(func() {
		instance = new(Client)
		instance.init()
	})
	return instance
}

func GetInstanceMutex() *Client {
	if instance == nil {
		mu.Lock()
		defer mu.Unlock()
		if instance == nil {
			instance = new(Client)
			instance.init()
		}
	}
	return instance
}

func GetInstanceAtomic() *Client {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}
	mu.Lock()
	defer mu.Unlock()
	if initialized == 0 {
		instance = new(Client)
		instance.init()
		atomic.StoreUint32(&initialized, 1)
	}
	return instance
}

func (d *Client) init() {
	dbURI := fmt.Sprintf(
		"host=%s port=%s user=%s "+
			"password=%s dbname=%s",
		config.GetJsonValue("dbHost"),
		config.GetJsonValue("dbPort"),
		config.GetJsonValue("dbUsername"),
		config.GetJsonValue("dbPass"),
		config.GetJsonValue("dbName"),
	)
	dbRootCert := config.GetJsonValue("dbRootCert")
	if dbRootCert != nil {
		dbURI += fmt.Sprintf(" sslmode=require sslrootcert=%s sslcert=%s sslkey=%s",
			dbRootCert.(string), config.GetJsonValue("dbCert").(string), config.GetJsonValue("dbKey").(string))
	} else {
		dbURI += " sslmode=disable"
	}
	var err error
	d.DB, err = sql.Open("postgres", dbURI)
	if err != nil {
		panic(err)
	}
	// ping is necessary to create connection
	err = d.DB.Ping()
	if err != nil {
		panic(err)
	}
	log.Println("Successfully connected!")
}
