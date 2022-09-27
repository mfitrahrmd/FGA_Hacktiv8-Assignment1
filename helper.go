package main

import (
	"errors"
	"log"
	"os"
	"strconv"
)

func GetArgs() ([]int, error) {
	if len(os.Args) <= 1 {
		return nil, errors.New("kamu harus memberi setidaknya 1 CLI parameter, parameter diberikan : 0")
	}

	var args []int

	for _, arg := range os.Args[1:] {
		a, err := strconv.Atoi(arg)
		if err != nil {
			log.Fatalln("CLI parameter harus bertipe angka")
		}
		args = append(args, a)
	}

	return args, nil
}

func RepeatChar(char string, times int) string {
	c := char
	for i := 0; i < times; i++ {
		c += char
	}
	return c
}
