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
	JWT_secret     string
	POLKA_KEY      string
	fileserverHits atomic.Int32
}

func main() {
	fmt.Println("Running...")
	port := os.Getenv("PORT")
	const root = "."

	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Couldnt open Database!")
		return
	}
	dbQueries := database.New(db)

	apiCfg := apiConfig{
		db:             dbQueries,
		JWT_secret:     os.Getenv("JWT_secret"),
		POLKA_KEY:      os.Getenv("POLKA_KEY"),
		fileserverHits: atomic.Int32{},
	}

	serveMux := http.NewServeMux()
	serveMux.Handle("/app/", apiCfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(root)))))
	serveMux.HandleFunc("/", apiCfg.handleRunning)
	serveMux.HandleFunc("GET /admin/healthz", handleReadiness)
	serveMux.HandleFunc("GET /admin/metrics", apiCfg.handleMetrics)
	serveMux.HandleFunc("POST /admin/reset", apiCfg.handleReset)
	serveMux.HandleFunc("PUT /admin/wells_to_drill", handleToDrill)
	serveMux.HandleFunc("GET /admin/healthz", handleReadiness)

	serveMux.HandleFunc("GET /api/get_well_params", handleWellParams)

	serveMux.HandleFunc("POST /api/users", apiCfg.registerUsers)
	serveMux.HandleFunc("PUT /api/users", apiCfg.updateUsers)
	serveMux.HandleFunc("POST /api/login", apiCfg.loginUser)

	serveMux.HandleFunc("POST /api/refresh", apiCfg.handlerRefresh)
	serveMux.HandleFunc("POST /api/revoke", apiCfg.handlerRevoke)
	serveMux.HandleFunc("POST /api/polka/webhooks", apiCfg.handleWebhooks)

	var server = &http.Server{
		Addr:    ":" + port,
		Handler: serveMux,
	}
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

//**********

func respondError(w http.ResponseWriter, s string, statusCode int) {
	type returnVals struct {
		Error string `json:"error"`
	}

	respBody := returnVals{
		Error: s,
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
	w.Write([]byte(fmt.Sprintf("<html>\n<body>\n<p> Hi Docker, I pushed a new version! </p>\n</body>\n</html>", cfg.fileserverHits.Load())))
}
