package storage

import(
    "fmt"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)


type Config struct {
    Host string
    Port string
    Password string
    User string
    Database string
    SSLMode *string
}{
    Host: "0.0.0.0",
    Port: "5432",
    Password: "test_password",
    User: "test_user",
    Database: "test_database",
}


func NewConnection(config Config) (*gorm.DB, error) {

    // Construct DSN (Data Source Name)
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.Host, config.User, config.Password, config.Database, config.Port, config.SSLMode,
	)
	fmt.Println("Connecting to DB...")
	fmt.Println("Connecting to DB with DSN:", dsn)

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }
    return db, nil
}
