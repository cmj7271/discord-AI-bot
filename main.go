package main

import (
	"errors"
	"log"
)

func main() {
	if err := run(); err != nil {
		log.Printf("fail to start: %v\n", err)
	}
}

// TODO: context 추가해서 error 추가 구현
// TODO: PORT 번호는 환경변수로?
func run() error {
	return errors.New("something goes wrong")
}
