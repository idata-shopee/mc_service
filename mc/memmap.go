package mc

import (
	"errors"
	"strings"
	"sync"
)

type MemMap struct {
	memMap *sync.Map
}

func GetMemMap() *MemMap {
	var memMap sync.Map
	return &MemMap{&memMap}
}

func (mm *MemMap) Set(key string, value interface{}) error {
	var curMap = mm.memMap
	var subKeys = strings.Split(key, ".")
	// support json key: a.b.c
	for _, subKey := range subKeys[:len(subKeys)-1] {
		var subMap sync.Map
		nextI, _ := curMap.LoadOrStore(subKey, &subMap)
		if next, ok := nextI.(*sync.Map); !ok {
			return errors.New(`set to none-map object"`)
		} else {
			curMap = next
		}
	}

	curMap.Store(subKeys[len(subKeys)-1], value)
	return nil
}

func (mm *MemMap) Get(key string) (interface{}, error) {
	var curMap = mm.memMap

	var subKeys = strings.Split(key, ".")

	// support json key: a.b.c
	for _, subKey := range subKeys[:len(subKeys)-1] {
		if v, ok := curMap.Load(subKey); !ok {
			return nil, errors.New("key " + key + "is not exists")
		} else {
			if next, ok := v.(*sync.Map); !ok {
				return nil, errors.New(`get from none-map object"`)
			} else {
				curMap = next
			}
		}
	}

	if v, ok := curMap.Load(subKeys[len(subKeys)-1]); !ok {
		return nil, errors.New("key " + key + "is not exists")
	} else {
		return v, nil
	}
}
