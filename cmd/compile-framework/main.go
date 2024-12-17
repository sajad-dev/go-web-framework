package main

import (
	"os/exec"
	"strings"

	"github.com/fatih/color"
	"github.com/sajad-dev/go-web-framework/internal/compile"
)

func main() {
	ou := exec.Command("pwd")
	output, _ := ou.Output()

	Serve(string(output))
}

func Serve(address string) {

	addr := address

	com := compile.Compiler{Addr: strings.TrimSpace(addr)}

	color.Blue("Adress Local : http://127.0.0.1:8000")

	com.Run()
	com.RunTest()

	com.Compile()

}