package initialize

import (
	"github.com/google/uuid"
	"sync"
)

var Lookup sync.Map
var Match []string
var Size int

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
