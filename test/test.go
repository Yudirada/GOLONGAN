package test

import (
	"fmt"
)

// Middle is the main struct
type Middle struct {
	Fxs int
}

// Add custom function to Fxs
func (M *Middle) Add(fx func()) {
	//Middle.Fxs = append(Middle.Fxs, fx)
	M.Fxs = 2
}

// Run all custom functions
func (*Middle) Run() {
	fmt.Println("asdasdaasddpoi")
}

// =============================
