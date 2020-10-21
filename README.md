# sessions

[![GoDoc](https://godoc.org/github.com/gin-contrib/sessions?status.svg)](https://godoc.org/github.com/gin-contrib/sessions)
[![Go Report Card](https://goreportcard.com/badge/github.com/Calidity/gin-sessions)](https://goreportcard.com/report/github.com/Calidity/gin-sessions)

Gin middleware for session management with multi-backend support:

- [cookie-based](#cookie-based)
- [Redis](#redis) using [go-redis/redis/v8](https://github.com/go-redis/redis)

This Redis client allows for using an existing client with support for Redis Sentinel and cluster.

Forked from https://github.com/gin-contrib/sessions

## Usage

### Start using it

Download and install it:

```bash
$ go get github.com/Calidity/gin-sessions
```

Import it in your code:

```go
import "github.com/Calidity/gin-sessions"
```

## Basic Examples

### single session

```go
package main

import (
	"github.com/Calidity/gin-sessions"
	"github.com/Calidity/gin-sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/hello", func(c *gin.Context) {
		session := sessions.Default(c)

		if session.Get("hello") != "world" {
			session.Set("hello", "world")
			session.Save()
		}

		c.JSON(200, gin.H{"hello": session.Get("hello")})
	})
	r.Run(":8000")
}
```

### multiple sessions

```go
package main

import (
	"github.com/Calidity/gin-sessions"
	"github.com/Calidity/gin-sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	sessionNames := []string{"a", "b"}
	r.Use(sessions.SessionsMany(sessionNames, store))

	r.GET("/hello", func(c *gin.Context) {
		sessionA := sessions.DefaultMany(c, "a")
		sessionB := sessions.DefaultMany(c, "b")

		if sessionA.Get("hello") != "world!" {
			sessionA.Set("hello", "world!")
			sessionA.Save()
		}

		if sessionB.Get("hello") != "world?" {
			sessionB.Set("hello", "world?")
			sessionB.Save()
		}

		c.JSON(200, gin.H{
			"a": sessionA.Get("hello"),
			"b": sessionB.Get("hello"),
		})
	})
	r.Run(":8000")
}
```

## Backend Examples

### cookie-based

[embedmd]:# (_example/cookie/main.go go)
```go
package main

import (
	"github.com/Calidity/gin-sessions"
	"github.com/Calidity/gin-sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
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
		session.Save()
		c.JSON(200, gin.H{"count": count})
	})
	r.Run(":8000")
}
```

### Redis

[embedmd]:# (_example/redis/main.go go)
```go
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
		session.Save()
		c.JSON(200, gin.H{"count": count})
	})
	r.Run(":8000")
}
```
