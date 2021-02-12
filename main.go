package main

import (
	"errors"
	"log"
	"time"
	"math/rand"
	"fmt"
	"context"
)

const (
	minLatency = 10
	maxLatency = 5000
	timeout = 3000
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Millisecond)
	defer cancel()
	result, err := searchWithTimeout(ctx, "Tokyo", "Taipei")
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(result)

}

func searchWithTimeout(ctx context.Context, from, to string) ([]string, error) {
	res := make(chan []string)

	go func(){
		res <- slowSearch(from, to)
		close(res)
	}()

	select {
	case result := <-res:
		log.Println("OK")
		return result, nil
	case <-ctx.Done():
		err := errors.New("Timeout")
		return nil, err
	}
	
}

// slowSearch is an simulation of an API which search for the flight of EVA and China airline.
func slowSearch(from, to string) []string {
	rand.Seed(time.Now().Unix())
	latency := rand.Intn(maxLatency-minLatency) + minLatency
	log.Println("Latency: ", latency)
	time.Sleep(time.Duration(latency)*time.Millisecond)
	return []string{from+" --> "+to+"- EVA airline 11am\n", from+" --> "+to+"- China airline 6pm\n"}
}