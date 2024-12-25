package main

import (
	"os"
    "fmt"
	"errors"
	"net/http"
	"github.com/google/uuid"
    //"github.com/xuri/excelize/v2"
	"github.com/P3T3R2002/vasi_drill_personal_project/internal/database"
)

func (cfg *apiConfig)handleCreateGrid(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating cells\n")
	if os.Getenv("PLATFORM") != "dev_admin_peter" {
		respondError(w, errors.New("Unauthorized"), 403)
		return
	}

	err := cfg.db.DeleteGrid(r.Context())
	if err != nil {
		respondError(w, err, 500)
		return
	}

	for i := range(100) {
		for j := range(100) {
			_, err = cfg.db.CreateCell(r.Context(), database.CreateCellParams{
				ID:			uuid.New(),
				NumLong:	float64(i),
				NumVert:	float64(j),
				ExpectedDepth:	0,
			})
			if err != nil {
				respondError(w, err, 500)
				return
			}
		}
	}
	respondJson(w, nil, 200)
}

func (cfg *apiConfig)handleUpdateGrid(w http.ResponseWriter, r *http.Request) {
	if os.Getenv("PLATFORM") != "dev_admin_peter" {
		respondError(w, errors.New("Unauthorized"), 403)
		return
	}

	for i := range(100) {
		for j := range(100) {
			_, err := cfg.db.UpdateCell(r.Context(), database.UpdateCellParams{
				NumLong:	float64(i),
				NumVert:	float64(j),
				ExpectedDepth:	int32(i+j),
			})
			if err != nil {
				respondError(w, err, 500)
				return
			}
		}
	}
	respondJson(w, nil, 200)

/*	f, err := excelize.OpenFile("water_level.xlsx")
	if err != nil {
		respondError(w, "Could not open xlsx file", 500)
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			respondError(w, "Could not close xlsx file", 500)
		}
	}()

/*	// Get all the rows in a sheet
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Iterate through each row
	for i, row := range rows {
		// Skip header if needed
		if i == 0 {
			continue
		}
		
		// row is a slice of strings, each element is a cell
		for j, cellValue := range row {
			// Do something with cellValue
			
		}
	}*/

}

//********

func getGridCoordinates(long int, vert int) (int, int) {
	return 0, 0
}