package controller

import (
	"redis_gorm_fiber/domain"
	"strconv"

	"redis_gorm_fiber/model"

	"github.com/gofiber/fiber/v2"
)

type NovelController struct {
	novelUseCase domain.NovelUseCase
}

func NewNovelController(novelUseCase domain.NovelUseCase) *NovelController {
	return &NovelController{novelUseCase: novelUseCase}
}

// CreateNovel handles the HTTP POST request to create a new novel.
// It expects a JSON body with the novel details.
// It returns a JSON response with the status code and a message.
func (n *NovelController) CreateNovel(ctx *fiber.Ctx) error {
	// Parse the request body into a novelRequest struct
	var novelRequest model.Novel
	var response model.Response

	// If there is an error parsing the request body, return a 400 response with the error message
	if err := ctx.BodyParser(&novelRequest); err != nil {
		response = model.Response{
			StatusCode: 400,
			Message:    err.Error(),
		}
		return ctx.Status(400).JSON(response)
	}

	// Check if all the required fields are present in the request body
	if novelRequest.Author == "" || novelRequest.Name == "" || novelRequest.Description == "" {
		response = model.Response{
			StatusCode: 400,
			Message:    "All fields are required",
		}
		return ctx.Status(400).JSON(response)
	}

	// Call the novelUseCase's CreateNovel method to create the novel
	err := n.novelUseCase.CreateNovel(novelRequest)
	if err != nil {
		// If there is an error creating the novel, return a 500 response with the error message
		response = model.Response{
			StatusCode: 500,
			Message:    err.Error(),
		}
		return ctx.Status(500).JSON(response)
	}

	// If the novel is created successfully, return a 200 response with a success message
	response = model.Response{
		StatusCode: 200,
		Message:    "success",
	}
	return ctx.Status(200).JSON(response)
}

// GetNovelById handles the HTTP GET request to get a novel by its ID.
// It expects the ID as a parameter in the URL.
// It returns a JSON response with the status code, message, and data.
func (n *NovelController) GetNovelById(ctx *fiber.Ctx) error {
	// Extract the ID parameter from the URL
	id := ctx.Params("id")

	// Convert the ID parameter to an integer
	idInt, err := strconv.Atoi(id)
	if err != nil {
		// If there is an error parsing the ID, return a 400 response with an error message
		ctx.Status(400).JSON(fiber.Map{"message": "invalid id"})
		return err
	}

	// Call the novelUseCase's GetNovelById method to get the novel by its ID
	novel, err := n.novelUseCase.GetNovelById(idInt)
	if err != nil {
		// If there is an error getting the novel, return a 400 response with the error message
		return ctx.Status(400).JSON(fiber.Map{"message": err.Error()})
	}

	// Create a response struct based on the retrieved novel
	var res model.Response
	if novel.Name != "" {
		res = model.Response{
			StatusCode: 200,
			Message:    "success",
			Data:       novel,
		}
	} else {
		res = model.Response{
			StatusCode: 200,
			Message:    "data not found",
		}
	}

	// Return a 200 response with the response struct as JSON
	return ctx.Status(200).JSON(res)
}

func (n *NovelController) DeleteNovel(ctx *fiber.Ctx) error {

	id := ctx.Params("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"message": "invalid id"})
	}

	err = n.novelUseCase.DeleteNovel(idInt)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"message": err.Error()})
	}

	return ctx.Status(200).JSON(fiber.Map{"message": "success"})

}

func (n *NovelController) UpdateNovel(ctx *fiber.Ctx) error {

	id := ctx.Params("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"message": "invalid id"})
	}

	var novel model.Novel
	if err := ctx.BodyParser(&novel); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"message": err.Error()})
	}

	err = n.novelUseCase.UpdateNovel(idInt, novel)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"message": err.Error()})
	}

	return ctx.Status(200).JSON(fiber.Map{"message": "success"})
}
