package prose

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sort"
)

type Prosefile struct {
	Filetype string  `json:"filetype"`
	Version  string  `json:"version"`
	Entries  []Entry `json:"entries"`
}

func remove(entries []Entry, s int) []Entry {
	return append(entries[:s], entries[s+1:]...)
}

// TODO: return an error if need be instead of exiting
func RemoveEntryById(entries []Entry, id string) []Entry {
	removeIndex := -1
	for i := range entries {
		if entries[i].Id == id {
			removeIndex = i
			break
		}
	}

	if removeIndex == -1 {
		fmt.Println("Entry with id", id, "not found.")
		os.Exit(1)
	}

	return remove(entries, removeIndex)
}

func (p *Prosefile) Format() error {
	var seenIds map[string]bool = make(map[string]bool)

	// Format posts and ensure no duplicate IDs
	for index, entry := range p.Entries {
		p.Entries[index].Format()

		_, ok := seenIds[entry.Id]
		if ok {
			return errors.New("Invariant: found duplicate ID" + entry.Id)
		}

		seenIds[entry.Id] = true
	}

	// Sort posts
	sort.SliceStable(p.Entries, func(i, j int) bool {
		return p.Entries[i].Time.Unix() < p.Entries[j].Time.Unix()
	})

	return nil
}

func (p *Prosefile) Serialize(pretty bool) string {
	var outstr []byte
	if pretty {
		outstr, _ = json.MarshalIndent(p, "", "  ")
	} else {
		outstr, _ = json.Marshal(p)
	}

	return string(outstr)
}
