package cache

import (
	"encoding/json"
	"time"
	"url_shortener/models"
	"url_shortener/pkg/dbhelper"

	"github.com/go-redis/redis"
)

// In this application, we first initialize a Redis client and test
//  the connection to make sure it is working. Then we define
//   a function getLink that takes a key as input and first tries
//    to get the data from the cache using the getLinkFromCache
//    function. If that fails, it gets the Link from the Linkbase
//     using the getLinkFromDB function, and stores it in the cache
// 	 using the setLinkInCache function for next time.

// To use this caching function, you would simply call getLink
//  with the appropriate key. The first time you call it, it will
//  get the Link from the Linkbase and store it in the cache. The
//  next time you call it with the same key, it will retrieve the
//  Link from the cache instead of the Linkbase, which can significantly
//   speed up your application.

// Declaring global variable for creating instance of redis
var client *redis.Client

// Init function is a type of function in golang which execute atomaticly
func init() {
	// Initialize Redis client
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// Test the connection
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
}

// func New() *redis.Client {

// 	client := redis.NewClient(&redis.Options{
// 		Addr: "localhost:6379",
// 		// Password: "", // no password set
// 		// DB:       0,  // use default DB

// 	})
// 	// Test the connection
// 	_, err := client.Ping().Result()
// 	if err != nil {
// 		panic(err)
// 	}
// 	return client
// }

// If the data is not available in cache then it will get from database
func getLinkFromDB(code string) (models.Request, error) {
	// Simulate getting Link from the Linkbase
	link, err := dbhelper.GetLink(code)
	if err != nil {
		return models.Request{}, err

	}
	return link, nil
}

// This function return the data from inside cache only
func getLinkFromCache(key string) (models.Request, error) {
	// Get the Link from the cache
	var data models.Request
	link, err := client.Get(key).Result()
	if err != nil {
		return models.Request{}, err
	}
	err = json.Unmarshal([]byte(link), &data)
	if err != nil {
		return models.Request{}, err
	}
	return data, nil
}

func setLinkInCache(key string, value models.Request) error {
	// Set the Link in the cache
	err := client.Set(key, value.LongUrl, 10*time.Minute).Err()
	if err != nil {
		return err
	}
	return nil
}

func GetLink(key string) (models.Request, error) {
	// Try to get the Link from the cache first
	link, err := getLinkFromCache(key)
	if err != nil {
		// If there was an error, get the Link from the database
		link, _ = getLinkFromDB(key)
		// Then store the Link in the cache for next time
		err = setLinkInCache(key, link)
		if err != nil {
			return models.Request{}, err
		}
	}
	return link, nil
}

// Here is simple delete function which delete from cache
func DeleteCach(key string) error {

	err := client.Del(key)
	if err != nil {
		return err.Err()
	}

	return nil
}
