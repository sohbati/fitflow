package session

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	Create(session *Session) error
	Update(session *Session) error
	GetByTokenHash(tokenHash string) (*Session, error)
	GetByUserID(userID uuid.UUID) ([]*Session, error)
	GetActiveByUserID(userID uuid.UUID) ([]*Session, error)
	Delete(id uuid.UUID) error
	DeleteByTokenHash(tokenHash string) error
	DeleteExpired() error
	CountActiveUsers() (int64, error)
	CountActiveSessions() (int64, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(session *Session) error {
	return r.db.Create(session).Error
}

func (r *repository) Update(session *Session) error {
	return r.db.Save(session).Error
}

func (r *repository) GetByTokenHash(tokenHash string) (*Session, error) {
	var session Session
	err := r.db.Where("token_hash = ?", tokenHash).First(&session).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("session not found")
		}
		return nil, err
	}
	return &session, nil
}

func (r *repository) GetByUserID(userID uuid.UUID) ([]*Session, error) {
	var sessions []*Session
	err := r.db.Where("user_id = ?", userID).Find(&sessions).Error
	if err != nil {
		return nil, err
	}
	return sessions, nil
}

func (r *repository) GetActiveByUserID(userID uuid.UUID) ([]*Session, error) {
	var sessions []*Session
	now := time.Now()
	err := r.db.Where("user_id = ? AND expires_at > ?", userID, now).Find(&sessions).Error
	if err != nil {
		return nil, err
	}
	return sessions, nil
}

func (r *repository) Delete(id uuid.UUID) error {
	return r.db.Delete(&Session{}, id).Error
}

func (r *repository) DeleteByTokenHash(tokenHash string) error {
	return r.db.Where("token_hash = ?", tokenHash).Delete(&Session{}).Error
}

func (r *repository) DeleteExpired() error {
	now := time.Now()
	return r.db.Where("expires_at < ?", now).Delete(&Session{}).Error
}

func (r *repository) CountActiveUsers() (int64, error) {
	var count int64
	now := time.Now()
	err := r.db.Model(&Session{}).
		Where("expires_at > ?", now).
		Distinct("user_id").
		Count(&count).Error
	return count, err
}

func (r *repository) CountActiveSessions() (int64, error) {
	var count int64
	now := time.Now()
	err := r.db.Model(&Session{}).
		Where("expires_at > ?", now).
		Count(&count).Error
	return count, err
}

