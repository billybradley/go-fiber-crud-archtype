package services

import (
	"github.com/gofiber/fiber/v2"
)

type Lister interface {
	GetAll() ([]interface{}, error)
}

type Getter interface {
	GetByID(id uint) (interface{}, error)
}

type Creator interface {
	Create(c *fiber.Ctx) error
}

type Updater interface {
	Update(c *fiber.Ctx) error
}

type Deleter interface {
	Delete(c *fiber.Ctx) error
}
