package controller

import (
	"fmt"
	"net/http"
	"url_shortener/models"
	"url_shortener/pkg/cache"
	"url_shortener/pkg/dbhelper"

	"github.com/gin-gonic/gin"
)

// This file control the whole application from this file we are controlling
// what is our next process and what fille to imported and what function has to called
// This file is entry to structure of application.

// The job of this function to get short link from cache memory or database
// Get the short link
func GetShortLink(c *gin.Context) {

	code := c.Param("code")
	// id, _ := strconv.Atoi(code)
	// client := cache.New()
	request, err := cache.GetLink(code)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.Redirect(http.StatusPermanentRedirect, request.LongUrl)
	// c.JSONP(http.StatusOK, gin.H{
	// 	"code":    200,
	// 	"message": "success",
	// 	"data": map[string]string{
	// 		// "short_url": request.ShortUrl,
	// 		"long_url": request.LongUrl,
	// 	},
	// })

}

// This function is responsible for generating a short whenever user call this function
// Createurl ... Create url
func GenerateShortLink(c *gin.Context) {

	var request models.Request
	request.ShortUrl = dbhelper.RandomBase62String()
	c.ShouldBind(&request)
	fmt.Println(request.LongUrl)
	fmt.Println(request.ShortUrl)

	err := dbhelper.AddLink(request)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSONP(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data": map[string]string{
			"short_url": request.ShortUrl,
			// "long_url":  request.LongUrl,
		},
	})

}

func Delete(c *gin.Context) {
	key := c.Param("code")
	// var response models.Response
	err := cache.DeleteCach(key)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"msg":  "success",
	})
}
