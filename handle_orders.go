package main

import (
	"fmt"
	"time"
	"net/http"
	"math/rand"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/P3T3R2002/vasi_drill_personal_project/internal/database"
)

func (cfg *apiConfig) registerOrders(w http.ResponseWriter, r *http.Request) {
	type wellParam struct {
		CellID 	uuid.UUID 	`json:"cell_id"`
		NumLong	float64 `json:"num_long"`
		NumVert	float64 `json:"num_vert"`
		Expected_price	int32	`json:"expected_price"`
	}
	type orderParam struct {
		Wells 		[]wellParam `json:"wells"`
		Name        string `json:"name"`
		PhoneNum   	string `json:"phone_num"`
		Email     	string `json:"email"`
	}
	decoder := json.NewDecoder(r.Body)
	params := orderParam{}
	err := decoder.Decode(&params)
	if err != nil {
		respondError(w, err, 400)
		return
	} else {
		var full_price int32 = 0
		for _, well := range params.Wells{
			full_price += well.Expected_price
		} 

		order, err := cfg.db.CreateOrder(r.Context(), database.CreateOrderParams{		
			ID:			uuid.New(),
			Name:		params.Name,             
			PhoneNum:	params.PhoneNum,
			Email:		params.Email,
			LookUpCode:	cfg.generate_code(r),
			NumberOfWells: 	int32(len(params.Wells)),
			PredictedFullPrice: full_price,
		})
		if err != nil {
			respondError(w, err, 500)
			fmt.Println(err)
			return
		}

		for i, well := range params.Wells {
			_, err := cfg.db.CreateWell(r.Context(), database.CreateWellParams{		
				ID:			uuid.New(),
				GpsLong:	well.NumLong,
				GpsVert:	well.NumVert,
				Price:		well.Expected_price,
				CellID:		well.CellID,
				OrderID:	order.ID,
			})
			if err != nil {
				respondError(w, err, 500)
				fmt.Printf("1 %d\n", i)
				return
			}
		}

		_, err = cfg.db.CreateCode(r.Context(), database.CreateCodeParams{	
			Code:	order.LookUpCode,
			OrderID:	order.ID,
		})
		if err != nil {
			respondError(w, err, 500)
			fmt.Println(err)
			return
		}

		type returnVals struct {
			ID          uuid.UUID `json:"id"`
			Name 	 	string 	`json:"name"`
			Email    	string 	`json:"email"`
			LookUpCode	string 	`json:"look_up_code"`
		}
	
		respBody := returnVals{
			ID:         order.ID,
			Email:      order.Email,
			Name:		order.Name,
			LookUpCode: order.LookUpCode,
		}
	
		respondJson(w, respBody, 201)
	}
}

//------------------------------------------------------------

func (cfg *apiConfig) getOrderCodes(w http.ResponseWriter, r *http.Request) {
	codes, err := cfg.db.GetCodes(r.Context())
	if err != nil {
		respondError(w, err, 500)
		return
	}
	
	var orderCodes []string

	for _, code := range codes {
		orderCodes = append(orderCodes, code)
	}

	type codesParam struct {
		Codes 		[]string `json:"codes"`
	}

	respCodes := codesParam{
		Codes:	orderCodes,
	}
	respondJson(w, respCodes, 200)
}

//------------------------------------------------------------

func (cfg *apiConfig) deleteOrder(w http.ResponseWriter, r *http.Request) {
	type Params struct {
		LookUpCode	string	`json:"look_up_code"`
	}
	decoder := json.NewDecoder(r.Body)
	params := Params{}
	err := decoder.Decode(&params)
	if err != nil {
		respondError(w, err, 400)
		return
	} else {
		err = cfg.db.DeleteOrder(r.Context(), params.LookUpCode)
		if err != nil {
			respondError(w, err, 400)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Deleted"))
	}
}

//------------------------------------------------------------

func (cfg *apiConfig) lookUpOrder(w http.ResponseWriter, r *http.Request) {
	type Params struct {
		LookUpCode	string	`json:"look_up_code"`
	}
	decoder := json.NewDecoder(r.Body)
	params := Params{}
	err := decoder.Decode(&params)
	if err != nil {
		respondError(w, err, 400)
		return
	} else {
		order, err := cfg.db.GetOrderDetails(r.Context(), params.LookUpCode)
		if err != nil {
			respondError(w, err, 500)
			return
		}

		wells, err := cfg.db.GetWellDetails(r.Context(), order.ID)
		if err != nil {
			respondError(w, err, 500)
			return
		}

		type wellParam struct {
			NumLong	float64 `json:"num_long"`
			NumVert	float64 `json:"num_vert"`
			Expected_depth	int32	`json:"expected_depth"`
			Expected_price	int32	`json:"expected_price"`
		}
		var respWell []wellParam

		for _, well := range wells {
			respWell = append(respWell, wellParam{
				NumLong: 		well.GpsLong,
				NumVert: 		well.GpsVert,
				Expected_depth:	well.ExpectedDepth,
				Expected_price: well.Price,
			})
		}

		type orderParam struct {
			Wells 		[]wellParam `json:"wells"`
			Name        string `json:"name"`
			Email     	string `json:"email"`
		}

		respOrder := orderParam{
			Wells:	respWell,
			Name:	order.Name,
			Email:	order.Email,
		}

		respondJson(w, respOrder, 200)
	}
}

//**********

func (cfg *apiConfig) generate_code(r *http.Request) string {
	newCode := make([]byte, 8)
	i := false
	for i == false {
		fmt.Println("code")
		rand.Seed(time.Now().UnixNano())
		const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	
		for i := range newCode {
			newCode[i] = charset[rand.Intn(len(charset))]
		}
		fmt.Println(newCode, string(newCode))

		_, err := cfg.db.GetCode(r.Context(), string(newCode))
		if err != nil {
			i = true
		} else {
			fmt.Println("duplicate code")
		}
	}	
	return string(newCode)
}