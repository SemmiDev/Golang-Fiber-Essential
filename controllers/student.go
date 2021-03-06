package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/xid"
	"strconv"
)

// Student : student model
type Student struct {
	ID			int		`json:"id"`
	Identifier  string  `json:"identifier"`
	Name 		string  `json:"name"`
}

var students = []*Student{
	{
		ID: 1,
		Identifier: "2003113948",
		Name: "sammidev",
	},
	{
		ID: 2,
		Identifier: "2003113949",
		Name: "sam",
	},
}

// GetStudents : get all students
func GetStudents(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"students": students,
		},
	})
}

// CreateStudent : Create a student
func CreateStudent(c *fiber.Ctx) error {
	type Request struct {
		Name string `json:"name"`
	}

	// generated identifier automatically
	nim := xid.New()
	identifier := nim.String()

	var body Request
	err := c.BodyParser(&body)

	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	// create a student variable
	student := &Student{
		ID:         len(students) + 1,
		Identifier: identifier,
		Name: body.Name,
	}

	// append in students
	students = append(students, student)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"student": student,
		},
	})
}

// GetStudent : get a single student
// PARAM: id
func GetStudent(c *fiber.Ctx) error {
	// get parameter value
	paramID := c.Params("id")

	// convert parameter value string to int
	id, err := strconv.Atoi(paramID)

	// if error in parsing string to int
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse Id",
			"error":   err,
		})
	}

	// find student and return
	for _, student := range students {
		if student.ID == id {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"success": true,
				"data": fiber.Map{
					"student": student,
				},
			})
		}
	}

	// if student not available
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"success": false,
		"message": "Student not found",
	})
}

// UpdateStudent : Update a student
// PARAM: id
func UpdateStudent(c *fiber.Ctx) error {
	// find parameter
	paramID := c.Params("id")

	// convert parameter string to int
	id, err := strconv.Atoi(paramID)

	// if parameter cannot parse
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse id",
			"error":   err,
		})
	}

	// request structure
	type Request struct {
		Name *string   `json:"name"`
	}

	var body Request
	err = c.BodyParser(&body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	var student *Student

	for _, t := range students {
		if t.ID == id {
			student = t
			break
		}
	}

	if student.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Not found",
		})
	}

	if body.Name != nil {
		student.Name = *body.Name
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"student": student,
		},
	})
}

// DeleteStudent : Delete a student
// PARAM: id
func DeleteStudent(c *fiber.Ctx) error {
	// get param
	paramID := c.Params("id")

	// convert param string to int
	id, err := strconv.Atoi(paramID)

	// if parameter cannot parse
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse id",
			"error":   err,
		})
	}

	// find and delete student
	for i, student := range students {
		if student.ID == id {

			students = append(students[:i], students[i+1:]...)

			return c.SendStatus(fiber.StatusNoContent)
		}
	}

	// if student not found
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"success": false,
		"message": "Student not found",
	})
}
