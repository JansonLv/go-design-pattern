package main

// 水泥板
type cementSlab struct {
	cement string
}

// 水泥材料制成水泥板
func createCementSlab(cementName string) cementSlab {
	return cementSlab{
		cement: cementName,
	}
}
