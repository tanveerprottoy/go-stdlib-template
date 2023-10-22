package sqlxext

import (
	"fmt"
	"log"
	"sync"
	"sync/atomic"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/tanveerprottoy/stdlib-go-template/internal/pkg/config"
)

var (
	instance    *Client
	once        sync.Once
	mu          sync.Mutex
	initialized uint32
)

type Client struct {
	DB *sqlx.DB
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

func (c *Client) init() {
	// connection properties.
	info := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.GetEnvValue("DB_HOST"), config.GetEnvValue("DB_PORT"), config.GetEnvValue("DB_USER"), config.GetEnvValue("DB_PASS"), config.GetEnvValue("DB_NAME"))
	var err error
	c.DB, err = sqlx.Open("pgx", info)
	if err != nil {
		panic(err)
	}
	// ping is necessary to create connection
	err = c.DB.Ping()
	if err != nil {
		panic(err)
	}
	log.Println("Successfully connected!")
	// create table if not exists
	_, err = c.DB.Exec("CREATE TABLE IF NOT EXISTS users (id uuid PRIMARY KEY DEFAULT gen_random_uuid(), name VARCHAR NOT NULL, role VARCHAR NOT NULL, created_at BIGINT, updated_at BIGINT)")
	if err != nil {
		panic(err)
	}
	_, err = c.DB.Exec("CREATE TABLE IF NOT EXISTS contents (id uuid PRIMARY KEY DEFAULT gen_random_uuid(), name VARCHAR NOT NULL, created_at BIGINT, updated_at BIGINT)")
	if err != nil {
		panic(err)
	}
}
