package elegaku

import "fmt"

// 仮置き
type Rank struct {
	Rank   int    `json:"rank"`
	GirlId string `json:"girl_id"`
}

func info() {
	fmt.Println("elegaku pkg")
}
