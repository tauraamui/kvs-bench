package kvsbench_test

import (
	"os"
	"testing"

	"github.com/dgraph-io/badger/v3"
	"github.com/tauraamui/kvs/v2"
	"github.com/tauraamui/kvs/v2/storage"
)

type Balloon struct {
	ID    uint32 `mdb:"ignore"`
	Color string
	Size  int
}

func (b Balloon) TableName() string { return "balloons" }

func BenchmarkKVSLoad(b *testing.B) {
	defer func() {
		os.RemoveAll("./data")
	}()
	conn, err := badger.Open(badger.DefaultOptions("").WithLogger(nil).WithDir("data").WithValueDir("data"))
	db, err := kvs.NewKVDB(conn)

	if err != nil {
		b.Fatal(err)
	}
	defer db.Close()

	store := storage.New(db)
	defer store.Close()

	store.Save(kvs.RootOwner{}, &Balloon{Color: "WHITE", Size: 366})
	store.Save(kvs.RootOwner{}, &Balloon{Color: "RED", Size: 695})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		storage.LoadAll[Balloon](store, kvs.RootOwner{})
	}
	b.StopTimer()
}

func BenchmarkKVSLoad500Records(b *testing.B) {
	defer func() {
		os.RemoveAll("./data")
	}()
	conn, err := badger.Open(badger.DefaultOptions("").WithLogger(nil).WithDir("data").WithValueDir("data"))
	db, err := kvs.NewKVDB(conn)

	if err != nil {
		b.Fatal(err)
	}
	defer db.Close()

	store := storage.New(db)
	defer store.Close()

	for i := 0; i < 500; i++ {
		color := "RED"
		if i%2 == 0 {
			color = "BLUE"
		}
		store.Save(kvs.RootOwner{}, &Balloon{Color: color, Size: i})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		storage.LoadAll[Balloon](store, kvs.RootOwner{})
	}
	b.StopTimer()
}

func BenchmarkKVSLoad1000Records(b *testing.B) {
	defer func() {
		os.RemoveAll("./data")
	}()
	conn, err := badger.Open(badger.DefaultOptions("").WithLogger(nil).WithDir("data").WithValueDir("data"))
	db, err := kvs.NewKVDB(conn)

	if err != nil {
		b.Fatal(err)
	}
	defer db.Close()

	store := storage.New(db)
	defer store.Close()

	for i := 0; i < 500; i++ {
		color := "RED"
		if i%2 == 0 {
			color = "BLUE"
		}
		store.Save(kvs.RootOwner{}, &Balloon{Color: color, Size: i})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		storage.LoadAll[Balloon](store, kvs.RootOwner{})
	}
	b.StopTimer()
}
