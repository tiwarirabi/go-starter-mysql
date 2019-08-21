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
	
	fmt.Printf("Starting API Server at http://localhost:%s \n", c.AppPort)
	fmt.Println("Press ctrl+C to stop the server.")

	srv := &http.Server{
		// Handler:      Handlers(),
		Addr:         port,
		ReadTimeout:  4 * time.Minute,
		WriteTimeout: 8 * time.Minute,
	}

	if serr := srv.ListenAndServe(); serr != nil {
		fmt.Printf("Error starting server %s", serr)
	}
}

//TODO: Create route handlers