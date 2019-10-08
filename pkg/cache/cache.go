package cache

import (
	"sync"

	"github.com/andreylm/guru/pkg/errors"
	"github.com/andreylm/guru/pkg/models"
)

// Cache - cache
var Cache cache

func init() {
	Cache = cache{
		users:      make(map[uint64]*models.User),
		statistics: make(map[uint64]*models.Statistics),
	}
}

type cache struct {
	mu         sync.Mutex
	users      map[uint64]*models.User
	statistics map[uint64]*models.Statistics
}

func (c *cache) AddUser(user *models.User) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if _, ok := c.users[user.ID]; ok {
		return errors.UserExistError
	}
	c.users[user.ID] = user
	return nil
}

func (c *cache) GetUser(id uint64) (*models.User, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if user, ok := c.users[id]; ok {
		return user, nil
	}

	return nil, errors.UserNotExistError
}

func (c *cache) GetUserStatistics(id uint64) (*models.Statistics, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if statistics, ok := c.statistics[id]; ok {
		return statistics, nil
	}

	return nil, errors.UserStatisticsNotExistError
}
