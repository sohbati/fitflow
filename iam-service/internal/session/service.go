package session

import (
	"time"

	"github.com/google/uuid"
)

type Service interface {
	CreateSession(userID uuid.UUID, tokenHash string, deviceInfo, ipAddress, userAgent string, expiresAt time.Time) (*Session, error)
	UpdateLastActivity(tokenHash string) error
	GetSessionByTokenHash(tokenHash string) (*Session, error)
	GetUserSessions(userID uuid.UUID) ([]*Session, error)
	GetActiveUserSessions(userID uuid.UUID) ([]*Session, error)
	DeleteSession(id uuid.UUID) error
	DeleteSessionByToken(tokenHash string) error
	DeleteExpiredSessions() error
	CountActiveUsers() (int64, error)
	CountActiveSessions() (int64, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateSession(userID uuid.UUID, tokenHash string, deviceInfo, ipAddress, userAgent string, expiresAt time.Time) (*Session, error) {
	session := &Session{
		UserID:         userID,
		TokenHash:      tokenHash,
		DeviceInfo:     deviceInfo,
		IPAddress:      ipAddress,
		UserAgent:      userAgent,
		ExpiresAt:      expiresAt,
		LastActivityAt: time.Now(),
	}

	if err := s.repo.Create(session); err != nil {
		return nil, err
	}

	return session, nil
}

func (s *service) UpdateLastActivity(tokenHash string) error {
	session, err := s.repo.GetByTokenHash(tokenHash)
	if err != nil {
		return err
	}

	session.LastActivityAt = time.Now()
	return s.repo.Update(session)
}

func (s *service) GetSessionByTokenHash(tokenHash string) (*Session, error) {
	return s.repo.GetByTokenHash(tokenHash)
}

func (s *service) GetUserSessions(userID uuid.UUID) ([]*Session, error) {
	return s.repo.GetByUserID(userID)
}

func (s *service) GetActiveUserSessions(userID uuid.UUID) ([]*Session, error) {
	return s.repo.GetActiveByUserID(userID)
}

func (s *service) DeleteSession(id uuid.UUID) error {
	return s.repo.Delete(id)
}

func (s *service) DeleteSessionByToken(tokenHash string) error {
	return s.repo.DeleteByTokenHash(tokenHash)
}

func (s *service) DeleteExpiredSessions() error {
	return s.repo.DeleteExpired()
}

func (s *service) CountActiveUsers() (int64, error) {
	return s.repo.CountActiveUsers()
}

func (s *service) CountActiveSessions() (int64, error) {
	return s.repo.CountActiveSessions()
}

