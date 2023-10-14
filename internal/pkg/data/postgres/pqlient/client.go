package pqlient

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"sync"
	"sync/atomic"
	"time"

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

// Ping the database to verify DSN is valid and the
// server is accessible. If the ping fails exit the program with an error.
func (d *Client) ping(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := d.DB.PingContext(ctx); err != nil {
		log.Fatalf("ping failed with error: %v", err)
	}
}

func (d *Client) init() {
	// postgres: //jack:secret@pg.example.com:5432/mydb?sslmode=verify-ca&pool_max_conns=10
	/* dbConn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?%s",
		config.GetEnvValue("DB_USERNAME"),
		config.GetEnvValue("DB_PASS"),
		config.GetEnvValue("DB_HOST"),
		config.GetEnvValue("DB_PORT"),
		config.GetEnvValue("DB_NAME"),
		config.GetJsonValue("dbSslMode"),
	) */
	dbURI := fmt.Sprintf(
		"host=%s port=%s user=%s "+
			"password=%s dbname=%s sslmode=%s",
		config.GetJsonValue("dbHost"),
		config.GetJsonValue("dbPort"),
		config.GetJsonValue("dbUsername"),
		config.GetJsonValue("dbPass"),
		config.GetJsonValue("dbName"),
		config.GetJsonValue("dbSslMode"),
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
	// Ping the database to verify DSN is valid and the
	// server is accessible
	d.ping(context.Background())
	log.Println("Successfully connected!")
	/* db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(10) */
	stat := d.DB.Stats()
	log.Printf("DB.stats: idle=%d, inUse=%d,  maxOpen=%d", stat.Idle, stat.InUse, stat.MaxOpenConnections)
}
