package infrastructure

type MemoryStore struct {
	Store map[string]map[string]any
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		Store: make(map[string]map[string]any),
	}
}

func (s *MemoryStore) Get(id string) (map[string]any, bool) {
	value, found := s.Store[id]
	return value, found
}

func (s *MemoryStore) Put(id string, value map[string]any) {
	s.Store[id] = value
}

func (s *MemoryStore) Delete(id string) {
	delete(s.Store, id)
}

func (s *MemoryStore) HasKey(id string) bool {
	_, found := s.Store[id]
	return found
}
