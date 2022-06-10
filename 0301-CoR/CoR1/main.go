package main

import (
	"errors"
	"fmt"
)

//  一个请求过来，验证ip，token，验证权限

type handle interface {
	execute(*request) error
	setNext(handle)
}

type request struct {
	ip    string
	token string
	auth  int
}

type ipHandle struct {
	next handle
}

func (handle *ipHandle) execute(req *request) error {
	if req.ip == "" {
		return errors.New("token is err")
	}
	fmt.Println("ipHandle")
	return handle.next.execute(req)
}

func (handle *ipHandle) setNext(next handle) {
	handle.next = next
}

type tokenHandle struct {
	next handle
}

func (handle *tokenHandle) execute(req *request) error {
	if req.token == "" {
		return errors.New("token is err")
	}
	fmt.Println("tokenHandle")
	return handle.next.execute(req)
}

func (handle *tokenHandle) setNext(next handle) {
	handle.next = next
}

type authHandle struct{}

func (handle *authHandle) execute(req *request) error {
	if req.auth != 1 {
		return errors.New("permission")
	}
	fmt.Println("authHandle")
	return nil
}

func (handle *authHandle) setNext(next handle) {
}

//
func main() {
	ipHandle := &ipHandle{}
	tokenHandle := &tokenHandle{}
	authHandle := &authHandle{}

	ipHandle.setNext(tokenHandle)
	tokenHandle.setNext(authHandle)

	ipHandle.execute(&request{
		ip:    "127.0.0.1",
		token: "token",
		auth:  1,
	})
}
