package model

import "github.com/google/uuid"

type AuthContext struct {
	UserID        uuid.UUID
	EmailVerified bool
}
