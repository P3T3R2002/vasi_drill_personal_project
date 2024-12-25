package main

import _ "github.com/lib/pq"

import (
	"os"
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
	"database/sql"
	"encoding/json"
	"github.com/joho/godotenv"
	"github.com/P3T3R2002/vasi_drill_personal_project/internal/database"
)

type apiConfig struct {
	db             *database.Queries
	POLKA_KEY      string
	fileserverHits atomic.Int32
}

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	fmt.Printf("Running on %v\n", port)
	const root = "."

	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Couldnt open Database!")
		return
	}
	dbQueries := database.New(db)

	apiCfg := apiConfig{
		db:             dbQueries,
		POLKA_KEY:      os.Getenv("POLKA_KEY"),
		fileserverHits: atomic.Int32{},
	}

	serveMux := http.NewServeMux()
	serveMux.Handle("/static", apiCfg.middlewareMetricsInc(http.StripPrefix("/static", http.FileServer(http.Dir(root)))))
	serveMux.HandleFunc("/", apiCfg.handleRunning)
	serveMux.HandleFunc("GET /admin/healthz", handleReadiness)
	serveMux.HandleFunc("GET /admin/metrics", apiCfg.handleMetrics)
	serveMux.HandleFunc("POST /admin/reset", apiCfg.handleReset)
	serveMux.HandleFunc("GET /admin/create_grid", apiCfg.handleCreateGrid)
	serveMux.HandleFunc("GET /admin/update_grid", apiCfg.handleUpdateGrid)

	serveMux.HandleFunc("GET /api/get_well_params", apiCfg.handleWellParams)
	serveMux.HandleFunc("POST /api/register_order", apiCfg.registerOrders)
	serveMux.HandleFunc("GET /api/get_order_codes", apiCfg.getOrderCodes)
	serveMux.HandleFunc("POST /api/delete_order", apiCfg.deleteOrder)
	serveMux.HandleFunc("GET /api/look_up_order", apiCfg.lookUpOrder)

	

	var server = &http.Server{
		Addr:    ":" + port,
		Handler: serveMux,
	}
	fmt.Println("Listening...")
	log.Fatal(server.ListenAndServe())
}

//------------------------------------------------------------

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits.Add(1)
		next.ServeHTTP(w, r)
	})
}

//**********

func respondJson(w http.ResponseWriter, r interface{}, statusCode int) {
	if r != nil {
		dat, err := json.Marshal(r)
		if err != nil {
			log.Printf("Error marshalling JSON: %s", err)
			w.WriteHeader(500)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		w.Write(dat)
	}
}

//**********

func respondError(w http.ResponseWriter, s error, statusCode int) {
	type returnVals struct {
		Error string `json:"error"`
	}

	respBody := returnVals{
		Error: s.Error(),
	}

	dat, err := json.Marshal(respBody)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(dat)
}

//------------------------------------------------------------

func (cfg *apiConfig) handleRunning(w http.ResponseWriter, r *http.Request) {
	respondJson(w, nil, 200)
	w.Write([]byte(fmt.Sprintf("<html>\n<body>\n<p> Hi Docker, I pushed a new version! </p>\n</body>\n</html> %v", cfg.fileserverHits.Load())))
}
