// handlers/generic_handler.go
package handlers

import (
	"CRUD-Fiber/services"
	"github.com/gofiber/fiber/v2"

	"strconv"
)

type GenericHandler struct {
	lister  services.Lister
	getter  services.Getter
	creator services.Creator
	updater services.Updater
	deleter services.Deleter
}

func NewGenericHandler(lister services.Lister, getter services.Getter, creator services.Creator, updater services.Updater, deleter services.Deleter) *GenericHandler {
	return &GenericHandler{
		lister:  lister,
		getter:  getter,
		creator: creator,
		updater: updater,
		deleter: deleter,
	}
}

func (h *GenericHandler) GetAll(c *fiber.Ctx) error {
	results, err := h.lister.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error getting models"})
	}
	return c.JSON(results)
}

func (h *GenericHandler) GetByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid model ID"})
	}

	result, err := h.getter.GetByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Model not found"})
	}

	return c.JSON(result)
}

func (h *GenericHandler) Create(c *fiber.Ctx) error {
	return h.creator.Create(c)
}

func (h *GenericHandler) Update(c *fiber.Ctx) error {
	return h.updater.Update(c)
}

func (h *GenericHandler) Delete(c *fiber.Ctx) error {
	return h.deleter.Delete(c)
}
