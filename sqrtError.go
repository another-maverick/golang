package main

import (
	"errors"
	"fmt"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error - %v %v %v", se.lat, se.long, se.err)

}

func main() {

	_, err := sqrt(-25.2)
	if err != nil {
		fmt.Println("here is the error received - ", err)
	}

}

func sqrt(f float64) (float64, error) {
	if f < 0 {

		return 0, sqrtError{"50.123 N", "34.231 S", errors.New("Need more coffee")}
	}
	return 42, nil
}
