 
package api

import (
	"log"
	"net/http"
)

func StartServer() {
	http.HandleFunc("/consensus", handleConsensus)
	http.HandleFunc("/verify", handleVerify)
	log.Println("API server running...")
	http.ListenAndServe(":8080", nil)
}

func handleConsensus(w http.ResponseWriter, r *http.Request) {
	// Handle consensus requests
}

func handleVerify(w http.ResponseWriter, r *http.Request) {
	// Handle proof verification requests
}
