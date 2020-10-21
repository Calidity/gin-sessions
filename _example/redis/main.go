package main

import (
	"github.com/Calidity/gin-sessions"
	sredis "github.com/Calidity/gin-sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func main() {
	r := gin.Default()
	store, _ := sredis.NewRedisStore(redis.NewClient(&redis.Options{Addr: "localhost:6379"}), []byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/incr", func(c *gin.Context) {
		session := sessions.Default(c)
		var count int
		v := session.Get("count")
		if v == nil {
			count = 0
		} else {
			count = v.(int)
			count++
		}
		session.Set("count", count)
		_ = session.Save()
		c.JSON(200, gin.H{"count": count})
	})
	r.Run(":8000")
}
