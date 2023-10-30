package main

import (
	"fmt"
	"sync"
)

var mutex = &sync.Mutex{}

type single struct {
}

var instance *single

func getInstance() *single {
	if instance == nil {
		mutex.Lock()
		defer mutex.Unlock()
		if instance == nil {
			fmt.Println("creation")
			instance = &single{}
		} else {
			fmt.Println("already exist.")
		}
	} else {
		fmt.Println("already exist")
	}

	return instance
}
