package storage

type inMemoryStorage struct {
	urls map[string]string
}

func NewStorage() Storage {
	return Storage(&inMemoryStorage{
		urls: make(map[string]string),
	})
}

func (s *inMemoryStorage) Save(k, v string) {
	s.urls[k] = v
}

func (s *inMemoryStorage) Get(k string) (string, bool) {
	v, ok := s.urls[k]
	return v, ok
}
