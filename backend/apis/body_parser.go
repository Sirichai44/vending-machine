package apis

import (
	"fmt"
	"log"
	"reflect"

	"github.com/gofiber/fiber/v2"
)

func HandleResponse[RES any](handle func(c *fiber.Ctx) (RES, error)) fiber.Handler {
	return func(c *fiber.Ctx) error {
		res, err := handle(c)
		if err != nil {
			return err
		}

		t := func(v any) any {
			ref := reflect.ValueOf(v)
			if (ref.Kind() == reflect.Pointer || ref.Kind() == reflect.Slice) && (v == nil || ref.IsNil()) {
				return []string{}
			} else {
				return v
			}
		}(res)

		return c.JSON(t)
	}
}

func HandleBodyParser[REQ, RES any](handle func(REQ, *fiber.Ctx) (RES, error)) fiber.Handler {
	return func(c *fiber.Ctx) error {
		const msg = "HandleBodyParser"
		var request REQ

		// parse request
		if err := c.BodyParser(&request); err != nil {
			log.Printf("%s-BodyParser: %v", msg, err)
			return fiber.ErrBadRequest
		}

		// validate
		if err := Validation(request); len(err) > 0 {
			for _, err := range err {
				log.Printf("%s-Validation: %s", msg, fmt.Sprintf("%+v", err))
			}
			return fiber.ErrBadRequest
		}

		// Todo implement function handler.
		res, err := handle(request, c)
		if err != nil {
			return err
		}

		return c.JSON(res)
	}
}
