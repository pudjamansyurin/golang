package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	data := `[
		{"Name": "Puja", "Wins" : 10},
		{"Name": "Kusuma", "Wins": 30}
	]`

	t.Run("works with an empty file", func(t *testing.T) {
		db, cleanDb := createTempFile(t, "")
		defer cleanDb()

		_, err := NewFileSystemPlayerStore(db)

		assertNoError(t, err)
	})

	t.Run(`league from a reader`, func(t *testing.T) {
		db, cleanDb := createTempFile(t, data)
		defer cleanDb()
		store, _ := NewFileSystemPlayerStore(db)

		got := store.GetLeague()
		want := League{
			{"Kusuma", 30},
			{"Puja", 10},
		}

		assertLeague(t, got, want)

		// read again
		got = store.GetLeague()
		assertLeague(t, got, want)
	})

	t.Run(`get player score`, func(t *testing.T) {
		db, cleanDb := createTempFile(t, data)
		defer cleanDb()
		store, _ := NewFileSystemPlayerStore(db)

		got := store.GetPlayerScore("Kusuma")
		want := 30

		assertScoreEquals(t, got, want)
	})

	t.Run(`store wins for existing players`, func(t *testing.T) {
		db, cleanDb := createTempFile(t, data)
		defer cleanDb()
		store, _ := NewFileSystemPlayerStore(db)

		store.RecordWin("Kusuma")

		got := store.GetPlayerScore("Kusuma")
		want := 31

		assertScoreEquals(t, got, want)
	})

	t.Run(`store wins for new players`, func(t *testing.T) {
		db, cleanDb := createTempFile(t, data)
		defer cleanDb()
		store, _ := NewFileSystemPlayerStore(db)

		store.RecordWin("Erawan")

		got := store.GetPlayerScore("Erawan")
		want := 1

		assertScoreEquals(t, got, want)
	})

	t.Run(`league sorted`, func(t *testing.T) {
		db, cleanDb := createTempFile(t, data)
		defer cleanDb()

		store, err := NewFileSystemPlayerStore(db)
		assertNoError(t, err)

		got := store.GetLeague()
		want := League{
			{Name: "Kusuma", Wins: 30},
			{Name: "Puja", Wins: 10},
		}
		assertLeague(t, got, want)

		// read again
		got = store.GetLeague()
		assertLeague(t, got, want)
	})
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf(`didn't except an error but got one, %v`, err)
	}
}

func assertScoreEquals(t *testing.T, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf(`got %d, want %d`, got, want)
	}
}

func createTempFile(t testing.TB, initialData string) (*os.File, func()) {
	t.Helper()
	tmpfile, err := ioutil.TempFile("", "db")
	if err != nil {
		t.Fatalf("could not create temporary file, %v", err)
	}

	tmpfile.Write([]byte(initialData))

	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
}
