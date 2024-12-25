package main

import(
	"os"
	"fmt"
	"errors"
	"net/http"
)

func handleReadiness(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type:", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))
}

//------------------------------------------------------------

func (cfg *apiConfig)handleMetrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type:", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("<html>\n<body>\n<h1>Welcome, Chirpy Admin</h1>\n<p>Chirpy has been visited %d times!</p>\n</body>\n</html>", cfg.fileserverHits.Load())))
}

//------------------------------------------------------------

func (cfg *apiConfig)handleReset(w http.ResponseWriter, r *http.Request) {
	if os.Getenv("PLATFORM") != "dev_admin_peter" {
		respondError(w, errors.New("Unauthorized"), 403)
		return
	}
	cfg.db.DeleteOrders(r.Context())
	cfg.fileserverHits.Store(0)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hits reset to 0"))
}