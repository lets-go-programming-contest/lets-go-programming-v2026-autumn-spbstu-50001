package main

import (
	"errors"
	"fmt"

	"github.com/Alehandro-de-bah/task-2-2/internal/minheap"
	"github.com/Alehandro-de-bah/task-2-2/internal/myheap"
)

var (
	errDishesOutOfRange     = errors.New("number of dishes must be in range [1, 10000]")
	errRatingOutOfRange     = errors.New("rating of dish must be in range [-10000, 10000]")
	errPreferenceOutOfRange = errors.New("preference number must be in range [1, dishesNumber]")
)

const (
	minDishesNumber     = 1
	maxDishesNumber     = 10000
	minDishesRating     = -10000
	maxDishesRating     = 10000
	minPreferenceNumber = 1
)

func readDishesNumber() (int, error) {
	var dishesNumber int
	if _, err := fmt.Scan(&dishesNumber); err != nil {
		return 0, fmt.Errorf("invalid number of dishes, %w", err)
	}

	if dishesNumber < minDishesNumber || dishesNumber > maxDishesNumber {
		return 0, errDishesOutOfRange
	}

	return dishesNumber, nil
}

func readRatings(dishesNumber int) ([]int, error) {
	rating := make([]int, dishesNumber)

	for index := range rating {
		if _, err := fmt.Scan(&rating[index]); err != nil {
			return nil, fmt.Errorf("invalid rating of dish, %w", err)
		}

		if rating[index] < minDishesRating || rating[index] > maxDishesRating {
			return nil, errRatingOutOfRange
		}
	}

	return rating, nil
}

func readPreferenceNumber(dishesNumber int) (int, error) {
	var preferenceNumber int
	if _, err := fmt.Scan(&preferenceNumber); err != nil {
		return 0, fmt.Errorf("invalid preference number, %w", err)
	}

	if preferenceNumber < minPreferenceNumber || preferenceNumber > dishesNumber {
		return 0, errPreferenceOutOfRange
	}

	return preferenceNumber, nil
}

func findKthLargest(heap myheap.MyHeap, ratings []int, k int) (int, error) {
	for _, value := range ratings {
		heap.Push(value)

		if heap.Len() > k {
			if _, err := heap.Pop(); err != nil {
				return 0, fmt.Errorf("failed to pop from heap: %w", err)
			}
		}
	}

	result, err := heap.Peek()
	if err != nil {
		return 0, fmt.Errorf("failed to peek from heap: %w", err)
	}

	return result, nil
}

func main() {
	dishesNumber, err := readDishesNumber()
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())

		return
	}

	rating, err := readRatings(dishesNumber)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())

		return
	}

	preferenceNumber, err := readPreferenceNumber(dishesNumber)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())

		return
	}

	result, err := findKthLargest(minheap.NewMinHeap(), rating, preferenceNumber)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())

		return
	}

	fmt.Println(result)
}
