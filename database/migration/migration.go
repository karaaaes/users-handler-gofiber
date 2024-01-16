package migration

import (
	"fmt"
	"go-fiber/database"
	"go-fiber/model/entity"
)

func RunMigration() {
	// package_name.variable_db.orm_gorm
	err := database.DB.AutoMigrate(&entity.User{})
	if err != nil {
		fmt.Println("Migration failed")
	}

	fmt.Println("Migration completed")
}
