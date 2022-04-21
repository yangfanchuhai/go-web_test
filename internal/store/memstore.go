package store

import (
	"github.com/yangfanchuhai/go-web_test/store"
	"github.com/yangfanchuhai/go-web_test/store/factory"
	"sync"
)

func init()  {
	factory.Register("mem", &MemStore{
		bookMap: make(map[string]store.Book),
	})
}

type MemStore struct {
	mutex sync.Mutex
	bookMap map[string]store.Book
}

func (m *MemStore) Create(book *store.Book) error  {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.bookMap[book.Id] = *book
	return nil
}

func (m *MemStore) Update(book *store.Book) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.bookMap[book.Id] = *book
	return nil
}

func (m *MemStore) Get(id string) (store.Book, error) {
	return m.bookMap[id], nil
}

func (m *MemStore) GetAll() ([]store.Book, error) {
	r := make([]store.Book, len(m.bookMap))
	for _, b := range m.bookMap {
		r = append(r, b)
	}
	return r, nil
}

func (m *MemStore) Delete(id string) error {
	m.mutex.Lock()
	delete(m.bookMap, id)
	defer m.mutex.Unlock()
	return nil
}

