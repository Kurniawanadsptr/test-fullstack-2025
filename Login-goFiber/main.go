package main

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

var ctx = context.Background()

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserData struct {
	Nama     string `json:"nama"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	r := redis.NewClient(&redis.Options{
		Addr: "localhost:3690",
	})
	app := fiber.New()
	app.Post("/login", func(c *fiber.Ctx) error {
		var body LoginRequest
		if err := c.BodyParser(&body); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid",
			})
		}
		k := fmt.Sprintf("login_%s", body.Username)
		raw, err := r.Get(ctx, k).Result()
		if err == redis.Nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "User not found",
			})
		} else if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "error",
			})
		}
		h := sha1.New()
		h.Write([]byte(body.Password))
		encoded := hex.EncodeToString(h.Sum(nil))

		if encoded != user.Password {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid akun",
			})
		}
		return c.JSON(fiber.Map{
			"message":  "Login berhasil",
			"realname": user.Nama,
			"email":    user.Email,
		})
	})

	log.Fatal(app.Listen(":3009"))
}
