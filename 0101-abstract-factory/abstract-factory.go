package main

import "fmt"

// 抽象工厂接口
type iSportsFactory interface {
	makeShoe() iShoe
	makeShirt() iShirt
}

// 定义了两个厂家，一个是耐克一个是阿迪，且两个厂家都可以生产shirt和shoe
func getSportsFactory(brand string) (iSportsFactory, error) {
	if brand == "adidas" {
		return &adidas{}, nil
	}

	if brand == "nike" {
		return &nike{}, nil
	}

	return nil, fmt.Errorf("Wrong brand type passed")
}
