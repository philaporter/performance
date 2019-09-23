package initialize

import (
	"github.com/google/uuid"
	"sync"
)

// Lookup is the sync.Map containing data that looks like it would be better suited in a []string.
var Lookup sync.Map

// Match is the []string list that says which uuids should be shared in the response
var Match []string

// Size is just a cheat code for knowing the size of the sync.Map that makes it easier to make "simple" loops
var Size int

// Init sets the sync.Map and []string that we'll use to bog down the performance
func Init() {
	Size = 10000
	Lookup = sync.Map{}
	for i := 0; i < Size; i++ {
		uuid := uuid.New().String()
		Lookup.Store(i, uuid)
		if i%7 == 0 || 1%13 == 0 {
			Match = append(Match, uuid)
		}
	}
}
