// Package auth
package auth

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           string
	Email        string
	PasswordHash string
}

type MemoryStore struct {
	mu    sync.Mutex
	users map[string]*User
	seq   int
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		users: make(map[string]*User),
	}
}

func (s *MemoryStore) Register(ctx context.Context, email, password string) (*User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.users[email]; exists {
		return nil, errors.New("user already exists")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("hashing password: %w", err)
	}

	s.seq++
	user := &User{
		ID:           fmt.Sprintf("usr_%d", s.seq),
		Email:        email,
		PasswordHash: string(hash),
	}

	s.users[email] = user
	return user, nil
}

func (s *MemoryStore) Login(ctx context.Context, email, password string) (*User, error) {
	s.mu.Lock()
	user, exists := s.users[email]
	s.mu.Unlock()

	if !exists {
		return nil, errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

func (s *MemoryStore) FindByID(ctx context.Context, id string) (*User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, user := range s.users {
		if user.ID == id {
			return user, nil
		}
	}

	return nil, errors.New("user not found")
}
