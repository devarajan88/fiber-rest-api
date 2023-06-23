package controller

import (
	db "fiber-rest-api/config"
	"fiber-rest-api/model"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"time"
)

func AdmitStudent(c *fiber.Ctx) error {
	var data map[string]string
	err := c.BodyParser(&data)
	fmt.Println(data)
	if err != nil {
		return c.Status(400).JSONP(
			fiber.Map{
				"success": false,
				"message": "Invalid data",
			})
	}
	if data["firstName"] == "" {
		return c.Status(400).JSONP(fiber.Map{
			"success": false,
			"message": "Student first name is required",
		})
	}
	if data["lastName"] == "" {
		return c.Status(400).JSONP(fiber.Map{
			"success": false,
			"message": "Student's last name is required",
		})
	}
	if data["department"] == "" {
		return c.Status(400).JSONP(fiber.Map{
			"success": false,
			"message": "Student's department is required",
		})
	}

	var student = model.Student{
		Id:         data["id"],
		FirstName:  data["firstName"],
		LastName:   data["lastName"],
		Department: data["department"],
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Time{},
	}
	fmt.Println(student)
	db.DB.Create(&student)
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Student is admitted successfully",
		"data":    student,
	})
}

func GetStudents(c *fiber.Ctx) error {
	var students []model.Student
	var count int64

	db.DB.Select("*").Find(&students).Count(&count)
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Students' records retrieved successfully",
		"data":    students,
	})
}

func GetStudentById(c *fiber.Ctx) error {

	studentId := c.Params("id")
	var student model.Student

	db.DB.Select("*").Where("id= ?", studentId).First(&student)

	studentData := make(map[string]interface{})
	studentData["id"] = student.Id
	studentData["firstName"] = student.FirstName
	studentData["lastName"] = student.LastName
	studentData["department"] = student.Department
	studentData["createdAt"] = student.CreatedAt
	studentData["updatedAt"] = student.UpdatedAt

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Student is retrieved for the given id successfully",
		"data":    studentData,
	})
}
