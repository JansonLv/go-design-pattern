package main

import (
	"errors"
	"fmt"
)

type handleFunc func(*request) error
type MiddlewareFunc func(handleFunc) handleFunc
type request struct {
	ip    string
	token string
	auth  int
	resp  string
}

func ipExecuteFunc(hf handleFunc) handleFunc {
	return func(req *request) error {
		if req.token == "" {
			return errors.New("ip is err")
		}
		fmt.Println("ipHandle")
		return hf(req)
	}
}

func tokenExecuteFunc(hf handleFunc) handleFunc {
	return func(req *request) error {
		if req.token == "" {
			return errors.New("token is err")
		}
		fmt.Println("tokenHandle")
		return hf(req)
	}
}

func authExecuteFunc(hf handleFunc) handleFunc {
	return func(req *request) error {
		if req.auth != 1 {
			return errors.New("permission")
		}
		fmt.Println("authHandle")
		return hf(req)
	}
}

func dealExecute(req *request) error {
	req.resp = "xxxxxxxxxxxxxx"
	fmt.Println("dealHandle")
	return nil
}

func main() {
	f := dealExecute
	handles := []MiddlewareFunc{authExecuteFunc, tokenExecuteFunc, ipExecuteFunc}

	for _, handle := range handles {
		f = handle(f)
	}
	req := &request{
		ip:    "127.0.0.1",
		token: "token",
		auth:  1,
		resp:  "",
	}
	f(req)
	fmt.Println(req.resp)
}
