package service

import (
	"encoding/csv"
	"fiber-rest-api/model"
	"fmt"
	"gorm.io/gorm"
	"io"
	"log"
	"net/http"
	"time"
)

var DB *gorm.DB

func CsvProcessor() {
	http.HandleFunc("/upload", uploadCSV)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func uploadCSV(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Parse the multipart form
	err := r.ParseMultipartForm(10 << 20) // 10MB
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get the file from the request
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Parse the CSV file
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var rowData = model.CsvRecord{
			Company:          record[0],
			Person:           record[1],
			Name:             record[2],
			DeviceType:       record[3],
			MacAddress:       record[4],
			Registered:       record[5],
			Status:           record[6],
			UUIDCreationDate: record[7],
			DownloadDate:     record[8],
			HotDesking:       record[9],
			HotDeskingID:     record[10],
			HotDeskingPhone:  record[11],
			Location:         record[12],
			Group:            record[13],
			Comment:          record[14],
			Firmware:         record[15],
		}

		if rowData.Registered == "Ja" {
			createDate, _ := time.Parse("2006.01.02 15:04:00", rowData.UUIDCreationDate)
			fmt.Println(createDate, rowData.UUIDCreationDate)
			curr := time.Now().Add(-24 * time.Hour).Unix()
			fmt.Println(curr)
			//if createDate.After(curr) {
			//	db.DB.Create(&rowData)
			//}
		}
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "CSV file uploaded and processed successfully")
}
