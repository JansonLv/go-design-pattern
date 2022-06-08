package main

import "fmt"

// 定义产品接口
type iGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
	printDetails()
}

// 产品模型
type gun struct {
	name  string
	power int
}

func (g *gun) setName(name string) {
	g.name = name
}

func (g *gun) getName() string {
	return g.name
}

func (g *gun) setPower(power int) {
	g.power = power
}

func (g *gun) getPower() int {
	return g.power
}

// 产品描述
func (g *gun) printDetails() {
	fmt.Printf("Gun: %s", g.getName())
	fmt.Println()
	fmt.Printf("Power: %d", g.getPower())
	fmt.Println()
}

// ak47 产品
type ak47 struct {
	gun
}

func newAk47() iGun {
	return &ak47{
		gun: gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}

// musket产品
type musket struct {
	// 类似继承
	gun
}

func newMusket() iGun {
	return &musket{
		gun: gun{
			name:  "Musket gun",
			power: 1,
		},
	}
}

func getGun(gunType string) (iGun, error) {
	if gunType == "ak47" {
		return newAk47(), nil
	}
	if gunType == "musket" {
		return newMusket(), nil
	}
	return nil, fmt.Errorf("Wrong gun type passed")
}

func printDetails(g iGun) {
	g.printDetails()
}

func main() {
	ak47, _ := getGun("ak47")
	musket, _ := getGun("musket")

	ak47.printDetails()
	musket.printDetails()

	printDetails(newAk47())
	printDetails(newMusket())
}
