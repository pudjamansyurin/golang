package poker

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
)

func FileSystemPlayerStoreFromFile(path string) (*FileSystemPlayerStore, func(), error) {
	db, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf(`problem opening %s, %v`, path, err)
	}

	closeFunc := func() {
		db.Close()
	}

	store, err := NewFileSystemPlayerStore(db)
	if err != nil {
		return nil, nil, fmt.Errorf(`problem creating file system store, %v`, err)
	}
	return store, closeFunc, nil
}

func NewFileSystemPlayerStore(file *os.File) (*FileSystemPlayerStore, error) {
	err := initialisePlayerDBFile(file)
	if err != nil {
		return nil, fmt.Errorf(`problem initialising player db file %v`, err)
	}

	league, err := NewLeague(file)

	if err != nil {
		return nil, fmt.Errorf(`problem parsing json file %s, %v`, file.Name(), err)
	}

	return &FileSystemPlayerStore{
		db:     json.NewEncoder(&Tape{file}),
		league: league,
	}, nil
}

type FileSystemPlayerStore struct {
	db     *json.Encoder
	league League
}

func (f *FileSystemPlayerStore) GetLeague() League {
	sort.Slice(f.league, func(i, j int) bool {
		return f.league[i].Wins > f.league[j].Wins
	})
	return f.league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	var wins int
	player := f.league.Find(name)

	if player != nil {
		wins = player.Wins
	}
	return wins
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	player := f.league.Find(name)

	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{name, 1})
	}

	// fmt.Println(name, "Wins")
	f.db.Encode(f.league)
}

func initialisePlayerDBFile(file *os.File) error {
	file.Seek(0, 0)

	info, err := file.Stat()
	if err != nil {
		return fmt.Errorf(`problem getting file information of %v, %v`, file.Name(), err)
	}

	if info.Size() == 0 {
		file.Write([]byte(`[]`))
		file.Seek(0, 0)
	}
	return nil
}
