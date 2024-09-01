package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DB is the global database instance
var DB *gorm.DB

func Initialize() error {
    var err error
    DB, err = gorm.Open(sqlite.Open("paste.db"), &gorm.Config{})
    if err != nil {
        return err
    }

    // Auto migrate the schema
    return DB.AutoMigrate(&Paste{})
}

// Blog represents the blog model
type Paste struct {
    ID      uint   `gorm:"primaryKey"`
    Content string `gorm:"type:text"`
    Slug    string `gorm:"unique;size:8"`
}

func CreatePaste(content, slug string) error {
    paste := Paste{Content: content, Slug: slug}
    result := DB.Create(&paste)
    return result.Error
}

func GetPaste(slug string) (*Paste, error) {
    var paste Paste
    result := DB.Where("slug = ?", slug).First(&paste)
    if result.Error != nil {
        return nil, result.Error
    }
    return &paste, nil
}


