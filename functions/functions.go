package functions

import (
	"log"
)

// it starts from capital letter so it is exposed to other packages
func Sum() {
	log.Println("Function Sum Invoked")
	subtract()
}

// it is not exposed to other packages ,it can be used only inside this package

func subtract() {
	log.Println("Function Subtract Invoked")

}

func WithArg(name string) {
	log.Println("Function invoked", name)
}

func WithArgsAndReturn(name string) string {
	log.Println("Function invoked", name)
	//c := strings.ToUpper(name)
	return name
}

// Function can return multiple values as well

func WithArgsAndMulReturn(name string) (string, int) {
	log.Println("Function invoked", name)
	//c := strings.ToUpper(name)
	return name, 90
}

// Function can return multiple values as well

func WithMulArgsAndMulReturn(name string, age int) (string, int) {
	log.Println("Function invoked", name)
	//c := strings.ToUpper(name)
	return name, age
}
