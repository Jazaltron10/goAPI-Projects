package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jazaltron10/goAPI/chiAPI/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main() {

	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting GO API Service...")
	fmt.Println(`
  				  ______	  ______	   ______	   ______	   ___                               
  				 /\  ___\    /\  __ \     /\  __ \    /\  __ \    /\  \                              
  				 \ \ \___\   \ \ \/\ \    \ \  __ \   \ \  _\ \   \ \  \                                              
  				  \ \_____\   \ \_____\    \ \_\ \_\   \ \_\  /    \ \__\                                             
  				   \/_____/    \/_____/     \/_/\/_/    \/_/        \/__/                                              
   `)

	err := http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Error(err)
	}
}
