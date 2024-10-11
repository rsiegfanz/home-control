package resolvers

import "gorm.io/gorm"

type QueryResolver struct {
	DB *gorm.DB
}
