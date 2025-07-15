package utils

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUserIdFromContext(c *fiber.Ctx) (primitive.ObjectID, error) {
	contextUserID := "userId"
	userIdStr := c.Locals(contextUserID).(string)

	userId, err := primitive.ObjectIDFromHex(userIdStr)
	if err != nil {
		return primitive.NewObjectID(), errors.New("invalid User ID Format")
	}

	return userId, nil
}
