package kvsbench_test

import (
	"os"
	"testing"

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

	store.Insert(badgerhold.NextSequence(), &Balloon{Color: "WHITE", Size: 366})
	store.Insert(badgerhold.NextSequence(), &Balloon{Color: "RED", Size: 695})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := []Balloon{}
		store.Find(&result, &badgerhold.Query{})
	}
	b.StopTimer()
}

func BenchmarkBHLoad500Records(b *testing.B) {
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
		store.Insert(badgerhold.NextSequence(), &Balloon{Color: color, Size: i})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := []Balloon{}
		store.Find(&result, &badgerhold.Query{})
	}
	b.StopTimer()
}

func BenchmarkBHLoad100RecordsQueryColour(b *testing.B) {
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

	for i := 0; i < 100; i++ {
		color := "RED"
		if i%2 == 0 {
			color = "BLUE"
		}
		store.Insert(badgerhold.NextSequence(), &Balloon{Color: color, Size: i})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := []Balloon{}
		store.Find(&result, badgerhold.Where("Color").Eq("RED"))
	}
	b.StopTimer()
}

func BenchmarkBHLoad500RecordsQuerySizeNoMatches(b *testing.B) {
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
		store.Insert(badgerhold.NextSequence(), &Balloon{Color: color, Size: i})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := []Balloon{}
		store.Find(&result, badgerhold.Where("Size").Eq(988))
	}
	b.StopTimer()
}

func BenchmarkBHLoad500RecordsQueryColour(b *testing.B) {
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
		store.Insert(badgerhold.NextSequence(), &Balloon{Color: color, Size: i})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := []Balloon{}
		store.Find(&result, badgerhold.Where("Color").Eq("RED"))
	}
	b.StopTimer()
}
