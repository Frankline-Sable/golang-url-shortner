package storage

import "sync"

// Store represents an in-memory storage  for URL mappings
type Store struct {
	urls map[string]string // Key : Short URL, Value:Long URL
	mu   sync.RWMutex      // Mutex for concurrent access: 1 writer and multiple reader
}

// NewStore creates a new in memory store
func NewStore() *Store {
	return &Store{urls: make(map[string]string)}
}

// AddToStore saves the url mappings to the store
func (s *Store) AddToStore(shortUrl, longUrl string) {
	s.mu.Lock() //Lock for writing

	defer s.mu.Unlock()        // Unlock for writing
	s.urls[shortUrl] = longUrl // save the mapping
}

// GetFromStore reads the long url from store given the short one
func (s *Store) GetFromStore(shortUrl string) (string, bool) {
	s.mu.RLock()         // lock for reading
	defer s.mu.RUnlock() // unlocking for reading
	longUrl, exists := s.urls[shortUrl]

	return longUrl, exists // return the long ur; and its existing status
}

// RemoveFromStore  deletes from the store
func (s *Store) RemoveFromStore(shortUrl string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, exists := s.urls[shortUrl]
	if !exists {
		return false
	}
	delete(s.urls, shortUrl)
	return true
}
