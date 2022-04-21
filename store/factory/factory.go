package factory

import (
	"fmt"
	"github.com/yangfanchuhai/go-web_test/store"
	"sync"
)

var (
	providerMu sync.Mutex
	providerMap = make(map[string]store.Store)
)


func Register(name string, s store.Store)  {
	providerMu.Lock()
	defer providerMu.Unlock()

	if s == nil {
		panic("store:store is nil")
	}

	if _, dup := providerMap[name]; dup {
		panic("store:duplicated")
	}

	providerMap[name] = s
}

func New(name string) (store.Store, error)  {
	providerMu.Lock()
	s, ok := providerMap[name]
	if !ok {
		return nil, fmt.Errorf("store: %s is not exists", name)
	}
	return s, nil
}