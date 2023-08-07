package kvsbench_test

import (
	"io/fs"
	"os"
	"testing"

	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SQLBalloon struct {
	gorm.Model
	UUID uuid.UUID `gorm:"type:uuid"`
	Balloon
}

func BenchmarkSLLoad(b *testing.B) {
	defer func() {
		os.RemoveAll("./data")
	}()
	os.Mkdir("./data", fs.ModePerm)
	db, err := gorm.Open(sqlite.Open("./data/data.db"), &gorm.Config{})
	if err != nil {
		b.Fatal(err)
	}

	rootUUID := uuid.New()

	db.AutoMigrate(&SQLBalloon{})

	db.Create(&SQLBalloon{UUID: rootUUID, Balloon: Balloon{Color: "WHITE", Size: 366}})
	db.Create(&SQLBalloon{UUID: rootUUID, Balloon: Balloon{Color: "RED", Size: 695}})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		balloons := []SQLBalloon{}
		db.Where("uuid = ?", rootUUID).Find(&balloons)
	}
	b.StopTimer()
}
