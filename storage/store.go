package storage

import "sync"

// Store represents an in-memory storage  for URL mappings
type Store struct {
	urls map[string]string // Key : Short URL, Value:Long URL
	mu   sync.RWMutex      // Mutex for concurrent access: 1 writer and multiple reader
}

// NewStore creates a new in memory store and returns a pointer to the struct rather than copy
func NewStore() *Store {
	return &Store{urls: make(map[string]string)}
}

// AddToStore saves the url mappings to the store
func (s *Store) AddToStore(shortURL, longURL string) {
	s.mu.Lock() //Lock for writing

	defer s.mu.Unlock()        // Unlock for writing
	s.urls[shortURL] = longURL // save the mapping
}

// GetFromStore reads the long url from store given the short one
func (s *Store) GetFromStore(shortURL string) (string, bool) {
	s.mu.RLock()         // lock for reading
	defer s.mu.RUnlock() // unlocking for reading
	longURL, exists := s.urls[shortURL]

	return longURL, exists // return the long ur; and its existing status
}

// RemoveFromStore  deletes from the store
func (s *Store) RemoveFromStore(shortURL string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, exists := s.urls[shortURL]
	if !exists {
		return false
	}
	delete(s.urls, shortURL)
	return true
}

// GetCount To know how many shortened URLs are stored
func (s *Store) GetCount() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.urls)
}
