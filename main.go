package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()
	ctx := context.Background()
	userID := 10
	val, err := fetchUserData(ctx, userID)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Result: ,", val)
	fmt.Println("took :", time.Since(start))

}

func fetchUserData(ctx context.Context, userID int) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200)
	defer cancel()
	val, err := fetchThirdPartyStuffWhichCanBeSlow()
	if err != nil {
		return 0, err
	}

	return val, nil
}

func fetchThirdPartyStuffWhichCanBeSlow() (int, error) {
	time.Sleep(time.Millisecond * 500)

	return 666, nil
}
