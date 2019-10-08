package cache

import (
	"sync"

	"github.com/andreylm/guru/pkg/errors"
	"github.com/andreylm/guru/pkg/models"
)

// Storage - cache
var Storage cache

// UserChangesType - type of changes for statistics
type UserChangesType uint

const (
	// UserChangesDeposit - deposit changes
	UserChangesDeposit UserChangesType = iota + 1
	// UserChangesBet - bet made
	UserChangesBet
	// UserChangesWin - win
	UserChangesWin
)

func init() {
	Storage = cache{
		users:         make(map[uint64]*models.User),
		modifiedUsers: make(map[uint64]*models.User),
		statistics:    make(map[uint64]*models.Statistics),
		transactions:  make(map[uint64]*models.Transaction),
	}
}

type cache struct {
	mu            sync.Mutex
	modifiedUsers map[uint64]*models.User
	users         map[uint64]*models.User
	statistics    map[uint64]*models.Statistics
	transactions  map[uint64]*models.Transaction
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

func (c *cache) GetUserStatistics(id uint64) *models.Statistics {
	c.mu.Lock()
	defer c.mu.Unlock()
	if _, ok := c.statistics[id]; !ok {
		c.statistics[id] = &models.Statistics{}
	}

	return c.statistics[id]
}

func (c *cache) AddModifiedUser(user *models.User) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.modifiedUsers[user.ID] = user
}

func (c *cache) GetModifiedUsers() map[uint64]*models.User {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.modifiedUsers
}

func (c *cache) RemoveModifiedUser(userID uint64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if _, ok := c.modifiedUsers[userID]; ok {
		delete(c.modifiedUsers, userID)
	}
}

func (c *cache) ClearModifiedUserCollection() {
	c.mu.Lock()
	defer c.mu.Unlock()
	for k := range c.modifiedUsers {
		delete(c.modifiedUsers, k)
	}
}

func (c *cache) UpdateUserStats(userID uint64, sum float64, changesType UserChangesType) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if _, ok := c.statistics[userID]; !ok {
		c.statistics[userID] = &models.Statistics{}
	}
	statistics := c.statistics[userID]
	switch changesType {
	case UserChangesDeposit:
		statistics.DepositCount++
		statistics.DepositSum += sum
	case UserChangesBet:
		statistics.BetCount++
		statistics.BetSum += sum
	case UserChangesWin:
		statistics.WinCount++
		statistics.WinSum += sum
	default:
		return errors.InvalidChangesTypeError
	}

	return nil
}
