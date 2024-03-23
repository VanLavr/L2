package main

import (
	"log"

	"github.com/VanLavr/L2/develop/11/internal/scheduler"
	schedulerrepo "github.com/VanLavr/L2/develop/11/internal/scheduler_repo"
	schedulertransport "github.com/VanLavr/L2/develop/11/internal/scheduler_transport"
)

func main() {
	repo := schedulerrepo.New()
	sched := scheduler.New(repo)
	srv := schedulertransport.New(":8080", sched)

	srv.RegisterRoutes()

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
