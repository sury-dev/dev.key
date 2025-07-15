package users

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthProvider string

const (
	AuthProviderLocal  AuthProvider = "local"
	AuthProviderGoogle AuthProvider = "google"
)

type User struct {
	ID              primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	FirstName       string             `json:"firstName" bson:"firstName"`
	LastName        string             `json:"lastName" bson:"lastName"`
	Email           string             `json:"email" bson:"email"`
	Password        *string            `json:"password,omitempty" bson:"password,omitempty"`
	AuthProvider    AuthProvider       `json:"authProvider" bson:"authProvider"`
	OAuthID         *string            `json:"oauthId,omitempty" bson:"oauthId,omitempty"`
	IsEmailVerified bool               `json:"isEmailVerified" bson:"isEmailVerified"`
	CreatedAt       time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt       time.Time          `json:"updatedAt" bson:"updatedAt"`
}
