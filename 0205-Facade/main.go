package main

import "fmt"

// 配置初始
type config struct{}

func (c *config) init() {
	fmt.Println("服务配置初始化")
}

// 服务发现
type discovery struct{}

func (d *discovery) discover() {
	fmt.Println("服务发现了")
}

// 服务注册
type register struct{}

func (*register) Registration() {
	fmt.Println("服务注册了")
}

// 启动服务
type app struct {
	config    *config
	register  *register
	discovery *discovery
}

func NewApp() *app {
	return &app{
		config:    &config{},
		discovery: &discovery{},
		register:  &register{},
	}
}

func (app *app) run() {
	fmt.Println("服务启动了")
}

func (app *app) Run() {
	app.config.init()
	app.discovery.discover()
	app.register.Registration()
	app.run()
}

func main() {
	app := NewApp()
	app.Run()
}
