package kvsbench_test

import (
	"os"
	"testing"

	"github.com/tauraamui/kvs/v2"
	"github.com/timshannon/badgerhold"
)

func BenchmarkBHLoad(b *testing.B) {
	defer func() {
		os.RemoveAll("./data")
	}()
	options := badgerhold.DefaultOptions
	options.Dir = "data"
	options.Logger = nil
	options.ValueDir = "data"

	store, err := badgerhold.Open(options)
	if err != nil {
		b.Fatal(err)
	}
	defer store.Close()

	store.Insert(kvs.RootOwner{}, &Balloon{Color: "WHITE", Size: 366})
	store.Insert(kvs.RootOwner{}, &Balloon{Color: "RED", Size: 695})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := []Balloon{}
		store.Find(&result, badgerhold.Where(badgerhold.Key).Eq(kvs.RootOwner{}))
	}
	b.StopTimer()
}

func BenchmarkBHLoad1000Records(b *testing.B) {
	defer func() {
		os.RemoveAll("./data")
	}()
	options := badgerhold.DefaultOptions
	options.Dir = "data"
	options.Logger = nil
	options.ValueDir = "data"

	store, err := badgerhold.Open(options)
	if err != nil {
		b.Fatal(err)
	}
	defer store.Close()

	for i := 0; i < 500; i++ {
		color := "RED"
		if i%2 == 0 {
			color = "BLUE"
		}
		store.Insert(kvs.RootOwner{}, &Balloon{Color: color, Size: i})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := []Balloon{}
		store.Find(&result, badgerhold.Where(badgerhold.Key).Eq(kvs.RootOwner{}))
	}
	b.StopTimer()
}
