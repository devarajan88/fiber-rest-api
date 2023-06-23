package service

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func UploadFile(c *fiber.Ctx) error {

	file, err := c.FormFile("filename")

	if err != nil {
		fmt.Println("Error occurred.")
		return nil
	}

	fmt.Println("File name: ", file.Filename)
	fmt.Println("File size: ", file.Size)
	fmt.Println("File type: ", file.Header)
	content, _ := file.Open()
	//var csvRecords []model.CsvRecord
	//for _, csvRecord := range csvRecords {
	//	record, _ := strconv.Atoi()
	//}
	fmt.Println(content)
	return c.SaveFile(file, fmt.Sprintf("./%s", file.Filename))

}
