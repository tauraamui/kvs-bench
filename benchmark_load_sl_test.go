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

func BenchmarkSLLoad500Records(b *testing.B) {
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

	for i := 0; i < 500; i++ {
		color := "RED"
		if i%2 == 0 {
			color = "BLUE"
		}
		db.Create(&SQLBalloon{UUID: rootUUID, Balloon: Balloon{Color: color, Size: i}})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		balloons := []SQLBalloon{}
		db.Where("uuid = ?", rootUUID).Find(&balloons)
	}
	b.StopTimer()
}

func BenchmarkSLLoad100RecordsQueryColour(b *testing.B) {
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

	for i := 0; i < 100; i++ {
		color := "RED"
		if i%2 == 0 {
			color = "BLUE"
		}
		db.Create(&SQLBalloon{UUID: rootUUID, Balloon: Balloon{Color: color, Size: i}})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		balloons := []SQLBalloon{}
		db.Where("uuid = ? AND color = ?", rootUUID, "RED").Find(&balloons)
	}
	b.StopTimer()
}

func BenchmarkSLLoad500RecordsQuerySizeNoMatches(b *testing.B) {
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

	for i := 0; i < 500; i++ {
		color := "RED"
		if i%2 == 0 {
			color = "BLUE"
		}
		db.Create(&SQLBalloon{UUID: rootUUID, Balloon: Balloon{Color: color, Size: i}})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		balloons := []SQLBalloon{}
		db.Where("uuid = ? AND size = ?", rootUUID, i).Find(&balloons)
	}
	b.StopTimer()
}

func BenchmarkSLLoad500RecordsQueryColour(b *testing.B) {
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

	for i := 0; i < 100; i++ {
		color := "RED"
		if i%2 == 0 {
			color = "BLUE"
		}
		db.Create(&SQLBalloon{UUID: rootUUID, Balloon: Balloon{Color: color, Size: i}})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		balloons := []SQLBalloon{}
		db.Where("uuid = ? AND color = ?", rootUUID, "RED").Find(&balloons)
	}
	b.StopTimer()
}
