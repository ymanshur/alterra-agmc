package seeder

import (
	"day-6/go-restful/database"
	"log"

	"gorm.io/gorm"
)

type seeder struct {
	DB    *gorm.DB
	model interface{}
}

func NewSeeder(db *gorm.DB, model interface{}) seeder {
	return seeder{
		DB:    db,
		model: model,
	}
}

func (s *seeder) Load() {
	if err := s.DB.Debug().Migrator().DropTable(&s.model); err != nil {
		log.Fatalf("Cannot drop table: %v", err)
	}

	database.Load(s.DB, &s.model)
}

func (s *seeder) Seed(instances ...interface{}) {
	if err := s.DB.Debug().Create(&instances).Error; err != nil {
		log.Fatalf("Cannot seed table: %v", err)
	}
	// for _, instance := range intances {
	// 	if err := s.DB.Debug().Model(&s.model).Create(&instance).Error; err != nil {
	// 		log.Fatalf("Cannot seed table: %v", err)
	// 	}
	// }
}
