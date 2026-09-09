package user

import (
	"context"
	"errors"
	"sync"

	"github.com/google/uuid"
)

type MemoryRepository struct {
	mu    sync.Mutex
	users map[string]*User
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{users: make(map[string]*User)}
}

func (r *MemoryRepository) Create(ctx context.Context, u *User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	u.ID = uuid.NewString()
	r.users[u.ID] = u
	return nil
}

func (r *MemoryRepository) GetByID(ctx context.Context, id string) (*User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	u, ok := r.users[id]
	if !ok {
		return nil, errors.New("user not found")
	}
	return u, nil
}

func (r *MemoryRepository) List(ctx context.Context) ([]*User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	res := make([]*User, 0, len(r.users))
	for _, u := range r.users {
		res = append(res, u)
	}
	return res, nil
}

func (r *MemoryRepository) Update(ctx context.Context, u *User) (*User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.users[u.ID]; !ok {
		return nil, errors.New("user not found")
	}
	r.users[u.ID] = u
	return u, nil
}

func (r *MemoryRepository) Delete(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.users[id]; !ok {
		return errors.New("user not found")
	}
	delete(r.users, id)
	return nil
}
