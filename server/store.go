package server

import (
    os "os"
    // uuid "github.com/satori/go.uuid"
    redis "gopkg.in/redis.v4"
    msgpack "gopkg.in/vmihailenco/msgpack.v2"
    str "strconv"
    time "time"
)

var address = os.Getenv("REDIS_ADDR")

// var password = os.Getenv("REDIS_PASS")
var db, err = str.Atoi(os.Getenv("REDIS_DB"))

// Create a new redis client
var client = redis.NewClient(&redis.Options{
    Addr:     address,
    Password: "",
    DB:       db,
})

type StoreInterface interface {
    Put([]byte) (string, error)
    Next() ([]byte, error)
    // Lock(string) error
    // Unlock(string) error
}

type store struct{}
type Store store

// Construct the store
func NewStore() *Store {
    return new(Store)
}

// Put some data into the store
func (s *Store) Put(data []byte) error {
    encoded, err := msgpack.Marshal(data)

    if err != nil {
        return err
    }

    return client.RPush("waiting", encoded).Err()
}

// Get some data from the store
func (s *Store) Next() (out string, err error) {

    data, err := client.BLPop(1*time.Second, "waiting").Result()

    if err != nil {
        return
    }

    if data[1] == "" {
        return
    }

    msgpack.Unmarshal([]byte(data[1]), &out)

    return
}
