package main

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	Id       string "json:id"
	Name     string "json:name"
	Age      int32  "json:age"
	Email    string "json:email"
	Password string "json:password"
}

var users []User

func main() {
	app := fiber.New()

	userGroup := app.Group("/users")

	userGroup.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusOK).JSON(users)
	})

	userGroup.Get("/:userId", func(ctx *fiber.Ctx) error {
		userId := ctx.Params("userId")

		if userId == "" {
			return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
				"message": "User ID is required",
			})
		}

		for _, user := range users {
			if user.Id == userId {
				return ctx.Status(http.StatusOK).JSON(user)
			}
		}

		return ctx.Status(http.StatusOK).JSON(fiber.Map{
			"message": "User not found",
		})
	})

	userGroup.Post("/", func(ctx *fiber.Ctx) error {

		var user User

		if err := ctx.BodyParser(&user); err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid request",
			})
		}

		users = append(users, user)

		return ctx.Status(http.StatusOK).JSON(fiber.Map{
			"message": "User created successfully",
		})
	})

	userGroup.Put("/:userId", func(ctx *fiber.Ctx) error {
		userId := ctx.Params("userId")

		for i, user := range users {
			if user.Id == userId {
				var u User

				if err := ctx.BodyParser(&u); err != nil {
					return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
						"message": "Invalid request",
					})
				}
				users[i] = u
			}
		}

		return ctx.Status(http.StatusOK).JSON(fiber.Map{
			"message": "User updated successfully",
		})
	})

	userGroup.Delete("/:userId", func(ctx *fiber.Ctx) error {
		userId := ctx.Params("userId")

		for i, user := range users {
			if user.Id == userId {
				users = append(users[:i], users[i+1:]...)
			}
		}

		return ctx.Status(http.StatusOK).JSON(fiber.Map{
			"message": "User deleted successfully",
		})
	})

	app.Listen(":3000")
}
