package main

import (
	"fmt"
	"log"

	"github.com/robfig/cron/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"  // or the Docker service name if running in another container
	port     = 5432         // default PostgreSQL port
	user     = "myuser"     // as defined in docker-compose.yml
	password = "mypassword" // as defined in docker-compose.yml
	dbname   = "mydatabase" // as defined in docker-compose.yml
)

func main() {
	// Configure your PostgreSQL database details here
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	// Migrate the schema
	db.AutoMigrate(&ExampleModel{})
	fmt.Println("Database migration completed!")

	c := cron.New(cron.WithSeconds())
	_, err = c.AddFunc("*/10 * * * * *", func() {
		go task(db)
	})

	if err != nil {
		l
		g
		flplt
		log.Fatal("Error scheduling a task:", err)
	}

	c.Start()

	// Block the main thread as the cron job runs in the background
	select {}
}

// func main() {
// 	c := cron.New(cron.WithSeconds())
// 	c.AddFunc("* * * * * *", func() { fmt.Println("Hi Fiat every second") })
// 	c.Start()

// 	select {}
// }
