package cache

import (
	"encoding/json"
	"time"
	"url_shortener/models"
	"url_shortener/pkg/dbhelper"

	"github.com/go-redis/redis"
)

func New() *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		// Password: "", // no password set
		// DB:       0,  // use default DB

	})
	// Test the connection
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	return client
}

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
func getLinkFromCache(key string, client *redis.Client) (models.Request, error) {
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

func setLinkInCache(key string, value models.Request, client *redis.Client) error {
	// Set the Link in the cache
	err := client.Set(key, value.LongUrl, 100*time.Minute).Err()
	if err != nil {
		return err
	}
	return nil
}

func GetLink(key string, client *redis.Client) (models.Request, error) {
	// Try to get the Link from the cache first
	// var request models.Request
	link, err := getLinkFromCache(key, client)

	if err != nil {
		// If there was an error, get the Link from the database
		link, _ = getLinkFromDB(key)
		// request.LongUrl = link.LongUrl
		// fmt.Println(link.LongUrl)
		// Then store the Link in the cache for next time
		err = setLinkInCache(key, link, client)
		if err != nil {
			return models.Request{}, err
		}
	}
	return link, nil
}

// Here is simple delete function which delete from cache
func DeleteCach(key string, client *redis.Client) error {

	err := client.Del(key)
	if err != nil {
		return err.Err()
	}

	return nil
}
