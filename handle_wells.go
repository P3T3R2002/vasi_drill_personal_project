package main

import (
	"fmt"
	"math"
	"errors"
	"net/http"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/P3T3R2002/vasi_drill_personal_project/internal/database"
)

func (cfg *apiConfig)handleWellParams(w http.ResponseWriter, r *http.Request) {
	const base_long int = 0
	const base_vert int = 0
	const pricePerMeter int = 20000
	
	type parameters struct{
		Num_long float64 `json:"num_long"`
		Num_vert float64 `json:"num_vert"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondError(w, err, 400)
		return
	} else {
		grid, err := cfg.getClosestGrid(r, params.Num_long, params.Num_vert)
		if err != nil {
			respondError(w, err, 400)
			return
		}
		expected_price := int(grid.ExpectedDepth)*pricePerMeter

		type returnVals struct {
			ID			uuid.UUID	`json:"id"`
			Num_long 	float64 	`json:"num_long"`
			Num_vert 	float64 	`json:"num_vert"`
			ExpectedDepth	int32 	`json:"expected_depth"`
			ExpectedPrice	int 	`json:"expected_price"`
		}

		respBody := returnVals{
			ID:			grid.ID,
			Num_long:	grid.NumLong,
			Num_vert:	grid.NumVert,
			ExpectedDepth:	grid.ExpectedDepth,
			ExpectedPrice:	expected_price,
		}
		respondJson(w, respBody, 200)
	}
}

func (cfg *apiConfig)handleToDrill(w http.ResponseWriter, r *http.Request) {

}

//****************//

func (cfg *apiConfig)getClosestGrid(r *http.Request, long, vert float64) (database.Grid, error) {
	grid_TR, err := cfg.db.GetClosestGrid_TR(r.Context(), database.GetClosestGrid_TRParams{
		NumLong: long, 
		NumVert: vert,
	})
	if err != nil {
		return database.Grid{}, err
	}
	dist_TR := getDistance(long, vert, grid_TR)

	grid_TL, err := cfg.db.GetClosestGrid_TL(r.Context(), database.GetClosestGrid_TLParams{
		NumLong: long, 
		NumVert: vert,
	})
	if err != nil {
		return database.Grid{}, err
	}
	dist_TL := getDistance(long, vert, grid_TL)

	grid_BR, err := cfg.db.GetClosestGrid_BR(r.Context(), database.GetClosestGrid_BRParams{
		NumLong: long, 
		NumVert: vert,
	})
	if err != nil {
		return database.Grid{}, err
	}
	dist_BR := getDistance(long, vert, grid_BR)

	grid_BL, err := cfg.db.GetClosestGrid_BL(r.Context(), database.GetClosestGrid_BLParams{
		NumLong: long, 
		NumVert: vert,
	})
	if err != nil {
		return database.Grid{}, err
	}
	dist_BL := getDistance(long, vert, grid_BL)

	switch math.Min(math.Min(dist_BL, dist_BR), math.Min(dist_TL, dist_TR)) {
	case dist_BL:
		return grid_BL, nil
	case dist_BR:
		return grid_BR, nil
	case dist_TL:
		return grid_TL, nil
	case dist_TR:
		return grid_TR, nil
	default:
		return database.Grid{}, errors.New("Problem at switch case")
	}
}

func getDistance(base_long, base_vert float64, grid database.Grid) float64{
	long := base_long-grid.NumLong
	vert := base_vert-grid.NumVert
	dist := math.Sqrt(long*long+vert*vert)
	fmt.Println(dist)
	return dist
}