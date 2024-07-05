package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	goredislib "github.com/redis/go-redis/v9"
)

type Config struct {
	DB    *sqlx.DB
	Redis *redsync.Redsync
}

func main() {
	// Init()

	router := gin.Default()
	// router.GET("/orders", GetOrders())
	c := Config{
		DB:    dbConnection(),
		Redis: redisConnection(),
	}

	router.GET("/orders", c.GetOrders)
	router.POST("/order", c.CreateOrder)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router.Handler(),
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
}

func dbConnection() *sqlx.DB {
	db, err := sqlx.Connect("postgres", "host=localhost port=5432 user=postgres password=password dbname=labs sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func redisConnection() *redsync.Redsync {
	client := goredislib.NewClient(&goredislib.Options{
		Addr: "localhost:6379",
	})
	pool := goredis.NewPool(client)

	return redsync.New(pool)
}

type Order struct {
	ID      int    `json:"id" db:"id"`
	OrderID string `json:"order_id" db:"order_id"`
}

// curl : curl --location 'http://localhost:8080/orders'
func (conf *Config) GetOrders(c *gin.Context) {
	orders := []Order{}
	err := conf.DB.Select(&orders, "SELECT id, order_id from orders")
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, orders)
}

// curl : curl --location --request POST 'http://localhost:8080/order'
func (conf *Config) CreateOrder(c *gin.Context) {
	redisMutexName := "generate-order-id"
	redisMutex := conf.Redis.NewMutex(redisMutexName)

	// Obtain a lock for our given mutex. After this is successful, no one else
	// can obtain the same lock (the same mutex name) until we unlock it.
	if err := redisMutex.Lock(); err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	order := Order{
		OrderID: generateOrderID(),
	}
	_, err := conf.DB.NamedExec("INSERT INTO orders (order_id) VALUES (:order_id)", order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	// Release the lock so other processes or threads can obtain a lock.
	if ok, err := redisMutex.Unlock(); !ok || err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, "Successfully created order")
}

func generateOrderID() string {
	return fmt.Sprintf("%s-%d", randomString(5), time.Now().UnixMilli())
}

func randomString(length int) string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}
