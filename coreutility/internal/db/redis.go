package db

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-redis/redis"
	"github.com/nilsyadv/ShopBillBuddy/common/pkg/config"
	wraperror "github.com/nilsyadv/ShopBillBuddy/common/pkg/error"
)

type RedisDB struct {
	rds *redis.Client
}

// NewRedisObject creates and returns a new instance of the RedisDB struct, which encapsulates a connection to a Redis database.
// It takes a configuration interface and uses it to set up the Redis client with the specified address, port, and password.
// The created RedisDB instance is configured with the obtained Redis client.
func NewRedisObject(conf config.InterfaceConfig) *RedisDB {
	client := redis.NewClient(&redis.Options{
		Addr:     conf.GetString("redis.addr") + ":" + conf.GetString("redis.port"), // Redis server address
		Password: conf.GetString("redis.pass"),                                      // No password
		DB:       0,                                                                 // Default DB
	})
	return &RedisDB{
		rds: client,
	}
}

// Get retrieves the value associated with a given key from the Redis database.
// It takes a key as a string and returns the value as a string and an error wrapped with additional information if encountered.
// If the key is not found, it returns an error with an HTTP 400 status.
// If the operation is successful, it returns the value and nil.
func (db *RedisDB) Get(key string) (string, *wraperror.WrappedError) {
	value, err := db.rds.Get(key).Result()
	if err != nil {
		if err == redis.Nil {
			werr := fmt.Sprintf("key %s not exist", key)
			return "", wraperror.Wrap(errors.New(werr), werr, "error", http.StatusBadRequest)
		}
		return "", wraperror.Wrap(err, "encounter error duering get", "error", http.StatusBadRequest)
	}

	return value, nil
}

// Set adds a new key-value pair to the Redis database or updates the value if the key already exists.
// It takes a key as a string and a value as an interface.
// It returns an error wrapped with additional information if encountered during the set operation.
// If the operation is successful, it returns nil.
func (db *RedisDB) Set(key string, value interface{}) *wraperror.WrappedError {
	err := db.rds.Set(key, value, 0).Err()
	if err != nil {
		return wraperror.Wrap(err, "encounter error duering set", "error", http.StatusInternalServerError)
	}
	return nil
}

// Update modifies the value associated with a given key in the Redis database.
// It takes a key as a string and a value as an interface.
// It returns an error wrapped with additional information if encountered during the update.
// If the operation is successful, it returns nil.
func (db *RedisDB) Update(key string, value interface{}) *wraperror.WrappedError {
	err := db.rds.Set(key, value, 0).Err()
	if err != nil {
		return wraperror.Wrap(err, "encounter error duering update", "error", http.StatusInternalServerError)
	}
	return nil
}

// Delete removes a key-value pair from the Redis database.
// It takes a key as a string and returns an error wrapped with additional information if encountered.
// If the operation is successful, it returns nil.
func (db *RedisDB) Delete(key string, value interface{}) *wraperror.WrappedError {
	err := db.rds.Del(key).Err()
	if err != nil {
		return wraperror.Wrap(err, "encounter error duering delete", "error", http.StatusInternalServerError)
	}
	return nil
}
