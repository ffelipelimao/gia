package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ffelipelimao/gia/internal/ai"
	"github.com/ffelipelimao/gia/internal/exec"
)

func main() {
	httpClient := &http.Client{
		Timeout: 15 * time.Second,
	}
	ai, err := ai.NewIA(httpClient)
	if err != nil {
		log.Fatal(err)
	}

	executor := exec.NewExecutor(ai)

	msg, err := executor.Start()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(msg)
}
