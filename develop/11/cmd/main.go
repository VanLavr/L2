package main

import (
	"fmt"
	"log"

	"github.com/VanLavr/L2/develop/11/internal/pkg/middlewares"
	"github.com/VanLavr/L2/develop/11/internal/scheduler"
	schedulerrepo "github.com/VanLavr/L2/develop/11/internal/scheduler_repo"
	schedulertransport "github.com/VanLavr/L2/develop/11/internal/scheduler_transport"
)

func main() {
	repo := schedulerrepo.New()
	sched := scheduler.New(repo)
	validator := middlewares.NewValidator("2006-01-02")
	srv := schedulertransport.New(":8080", sched, validator)

	srv.RegisterRoutes()

	fmt.Println("listening on :8080")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
