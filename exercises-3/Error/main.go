package main

import (
	"fmt"
	"time"
)

type CoustomError struct {
	Timestamp time.Time
}

func (e CoustomError) Error() string {
	return fmt.Sprintf("[%s] Custom error occurred", e.Timestamp.Format(time.TimeOnly))
}

func main(){
	err := CoustomError{Timestamp: time.Now()}

		fmt.Println(err)
}