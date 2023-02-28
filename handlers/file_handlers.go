package handlers

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

func GetPfp(c *fiber.Ctx) error {
	userName := c.Params("userName")
	fmt.Println(userName)

	_, err := os.Stat("assets/pfp/" + userName + ".png")
	if err != nil {
		return c.SendFile("assets/pfp/flop.png")

	}
	return c.SendFile("assets/pfp/" + userName + ".png")
}
