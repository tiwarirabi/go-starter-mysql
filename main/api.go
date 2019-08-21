package api

import (
	"fmt"
	"net/http"
	"time"


	config "github.com/tiwarirabi/go-starter-mysql/config"
)

func Start() {
	
	c, err := config.ReadConfig()
	if err != nil {
		fmt.Println("Error Starting API Server.")
		return
	}

	config.SetupDatabase()

	port := fmt.Sprintf(":%s", c.AppPort)
	
	fmt.Println("Starting API Server.")
	fmt.Println("Press ctrl+E to stop the server.")

	srv := &http.Server{
		// Handler:      Handlers(),
		Addr:         port,
		ReadTimeout:  4 * time.Minute,
		WriteTimeout: 8 * time.Minute,
	}

	if serr := srv.ListenAndServe(); serr != nil {
		fmt.Println("Error starting server")
	}
}


//Handlers phil router with all the handlers
// func Handlers() *pweb.PhilRouter {
// 	ctx := context.Background()
// 	return controller.SetupRoutes(ctx)
// }