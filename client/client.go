package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	req, err := http.NewRequest(http.MethodGet, "http://localhost:8090", nil)

	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	req = req.WithContext(ctx)
	defer cancel()
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Fatal(res.Status)
	}

	io.Copy(os.Stdout, res.Body)
}
