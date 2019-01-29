package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gkkkb/pokedex"
	"github.com/gkkkb/pokedex/pkg/api"
	"github.com/gkkkb/pokedex/pkg/api/response"
	"github.com/gkkkb/pokedex/pkg/log"
	"github.com/gkkkb/pokedex/route"

	"github.com/gkkkb/piston/metric"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

func main() {
	instance := pokedex.GetInstance() 
	
	router := httprouter.New()
	router.HandlerFunc("GET", "/metrics", metric.Handler)
	router.GET("/healthz", func(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
		resp := response.ResponseBody{
			Data:    instance.DB.Stats(),
			Message: "OK",
			Meta: response.MetaInfo{
				HTTPStatus: http.StatusOK},
		}
		response.Write(w, resp, http.StatusOK)
	})

	apis := route.Route()

	api.StartAPIs(router, apis)

	co := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PATCH", "DELETE", "PUT", "HEAD", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		MaxAge:         86400,
	})

	go instance.Loop()

	log.DevLog(fmt.Sprintf("Listening to port %s", os.Getenv("PORT")))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), honeybadger.Handler(co.Handler(router))))
}
