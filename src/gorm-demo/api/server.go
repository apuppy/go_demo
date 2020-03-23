package api

import (
	"fmt"
	"gorm-demo/api/route"
	"log"
	"net/http"
)

func listen(p int) {
	fmt.Printf("\n\nListening port %d...\n\n", p)
	port := fmt.Sprintf(":%d", p)
	r := route.NewRouter()
	log.Fatal(http.ListenAndServe(port, r))
}
