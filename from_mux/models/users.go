package models

import(
    "time"
    "gorm.io/gorm"
)

type Users struct {
    ID        uint          `gorm:"primary_key;auto_increment;true"`
    Name      *string       `json:"name"`
    Email     *string       `json:"email"`
    Password  *string       `json:"password"`
    CreatedAt time.Time     `json:"created_at;default:CURRENT_TIMESTAMP"`
    UpdatedAt time.Time     `json:"updated_at;default:CURRENT_TIMESTAMP"`
}


func MigrateUsers(db *gorm.DB) error {
    err := db.AutoMigrate(&Users{})
    return err
}
