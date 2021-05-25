package poker

import (
	"encoding/json"
	"fmt"
	"io"
)

type League []Player

func NewLeague(rdr io.Reader) (League, error) {
	var league League
	err := json.NewDecoder(rdr).Decode(&league)
	if err != nil {
		err = fmt.Errorf(`problem parsing league, %v`, err)
	}
	return league, err
}

func (l League) Find(name string) *Player {
	for i, v := range l {
		if v.Name == name {
			return &l[i]
		}
	}
	return nil
}
