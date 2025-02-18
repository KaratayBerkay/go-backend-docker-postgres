package main

import (
    "log"
    "fmt"
    "time"
    "gorm.io/gorm"
    "github.com/joho/godotenv"
    "github.com/gofiber/fiber/v2"
    "github.com/KaratayBerkay/go-backend-docker-postgres/from_fiber/storage"
    "github.com/KaratayBerkay/go-backend-docker-postgres/from_fiber/models"
)


type Repository struct {
    DB *gorm.DB
}


type User struct {
    Email string `json:"email"`
    Password string `json:"password"`
    Name string `json:"name"`
}


func (r *Repository) CreateUser(c *fiber.Ctx) error {
    user := new(User)
    if err := c.BodyParser(user); err != nil {
        return c.Status(400).SendString(err.Error())
    }
    r.DB.Create(&user)
    return c.JSON(user)
}


func (r *Repository) GetUsers(c *fiber.Ctx) error {
    var users []User
    r.DB.Find(&users)
    return c.JSON(users)
}


func (r *Repository) SetupRoutes(app *fiber.App) {
    api := app.Group("/api")
    api.Post("/users", r.CreateUser)
    api.Get("/users", r.GetUsers)
}


func main() {

    err := godotenv.Load(".env")
    if err != nil {
        fmt.Println("Error loading .env file")
    }

	config := storage.Config{
		Host:     "postgres_db_for_go_app",
		Port:     "5432",
		User:     "test_user",
		Password: "test_password",
		Database: "test_database",
		SSLMode:  "disable", // Change as needed
	}

    time.Sleep(20 * time.Second)    // Wait for postgres to start

	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

    models.MigrateUsers(db)
    if err != nil {
        fmt.Println("Error migrating users")
    }

	fmt.Println("Database connection successful!", db)
    if err != nil {
        fmt.Println("Error connecting to database")
    }

    r := Repository{DB: db}
    app := fiber.New()
    r.SetupRoutes(app)
    fmt.Println("Api up and running on port 8000")
    app.Listen(":8000")
}
