package repository

import (
	"errors"
	"sync"
	"tictactoe/internal/domain/model"
)

type InMemoryStore struct {
	store map[string]model.Game
	mu    sync.RWMutex
}

var (
	instance *InMemoryStore
	once     sync.Once
)

var (
	errorGameNotFound = errors.New("game not found")
)

func NewInMemoryStore() *InMemoryStore {
	once.Do(func() {
		instance = &InMemoryStore{
			store: make(map[string]model.Game),
		}
	})

	return instance
}

func (this *InMemoryStore) SaveGame(data model.Game) error {
	this.mu.Lock()
	defer this.mu.Unlock()

	this.store[data.GetID()] = data

	return nil
}

func (this *InMemoryStore) UpdateGame(data model.Game) error {
	this.mu.Lock()
	defer this.mu.Unlock()

	if _, exists := this.store[data.GetID()]; !exists {
		return errorGameNotFound
	}

	this.store[data.GetID()] = data
	return nil
}

func (this *InMemoryStore) GetGameByID(id string) (model.Game, error) {
	this.mu.RLock()
	defer this.mu.RUnlock()

	game, exists := this.store[id]
	if exists {
		return game, nil
	}
	return model.Game{}, errorGameNotFound
}
