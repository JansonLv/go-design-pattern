package main

import (
	"errors"
	"fmt"
)

//  一个请求过来，验证ip，token，验证权限
type handle interface {
	execute(*request) error
}

type request struct {
	ip    string
	token string
	auth  int
	resp  string
}

type ipHandle struct {
}

func (handle *ipHandle) execute(req *request) error {
	if req.ip == "" {
		return errors.New("token is err")
	}
	fmt.Println("ipHandle")
	return nil
}

type tokenHandle struct {
}

func (handle *tokenHandle) execute(req *request) error {
	if req.token == "" {
		return errors.New("token is err")
	}
	fmt.Println("tokenHandle")
	return nil
}

type authHandle struct{}

func (handle *authHandle) execute(req *request) error {
	if req.auth != 1 {
		return errors.New("permission")
	}
	fmt.Println("authHandle")
	return nil
}

type dealHandle struct{}

func (handle *dealHandle) execute(req *request) error {
	req.resp = "deal"
	fmt.Println("dealHandle")
	return nil
}

type handles struct {
	handles []handle
}

func (h *handles) execute(req *request) error {
	for _, handle := range h.handles {
		if err := handle.execute(req); err != nil {
			return err
		}
	}
	return nil
}

//
func main() {
	ipHandle := &ipHandle{}
	tokenHandle := &tokenHandle{}
	authHandle := &authHandle{}

	handles := &handles{[]handle{ipHandle, tokenHandle, authHandle, &dealHandle{}}}

	req := &request{
		ip:    "127.0.0.1",
		token: "token",
		auth:  1,
	}
	handles.execute(req)
	fmt.Printf("deaded: %s\n", req.resp)
}
