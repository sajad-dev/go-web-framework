package compile

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/fatih/color"
)

type Compiler struct {
	Addr string
}

func (co *Compiler) RunTest() {
	command := exec.Command("go", "test", "-v", co.Addr+"/test/check-system")
	var out, stderr strings.Builder
	command.Stdout = &out
	command.Stderr = &stderr
	err := command.Run()
	if err != nil {
		color.Red(stderr.String())
	}
	if strings.Contains(out.String(), "FAIL") {
		color.Yellow(out.String())
	} else {
		color.Cyan(out.String())
	}
}

func (co *Compiler) Compile() {

	for {
		if len(os.Args) > 1 {
			os.Exit(1)
		}
		co.Commands()
		time.Sleep(1 * time.Second)
	}
}

func (co *Compiler) build() {
	name := strings.Split(co.Addr, "/")
	command := exec.Command("go", "build","-o",co.Addr+"/build", co.Addr+"/cmd"+"/"+name[len(name)-1])
	var out, stderr strings.Builder
	command.Stdout = &out
	command.Stderr = &stderr

	err := command.Run()

	if err != nil {
		color.Red("Error In Compile: " + stderr.String())
	}

}

func (co *Compiler) Commands() {
	comand := exec.Command("git", "status", co.Addr)
	ou, err := comand.Output()
	if err != nil {
		return
	}

	output := strings.Split(string(ou), "Changes not staged for commit:")
	if len(output) > 1 && strings.Contains(output[1], "modified") {

		co.build()

		co.killProcess()
		comandGit := exec.Command("git", "add", ".", co.Addr)
		if err := comandGit.Run(); err != nil {
		}

		go co.Run()

	}
}

func (co *Compiler) Run() {
	currentTime := time.Now()

	fmt.Print("Compile in : ")
	color.Red(currentTime.Format("2006/01/02 15:04:05"))
	co.killProcess()
	comandRun := exec.Command("")

	filecompiledarr := strings.Split(co.Addr, "/")
	filecompiled := filecompiledarr[len(filecompiledarr)-1]
	if len(os.Args) > 1 {
		comandRun = exec.Command(co.Addr+"/build/"+filecompiled, os.Args[1:]...)

	} else {
		comandRun = exec.Command(co.Addr + "/build/" + filecompiled)

	}
	stdout, err := comandRun.StdoutPipe()
	if err != nil {
		color.Red(err.Error())
	}
	stderr, err := comandRun.StderrPipe()
	if err != nil {
		color.Red(err.Error())
	}

	err = comandRun.Start()
	if err != nil {
		co.build()
		color.Blue("Error: " + err.Error())

	}
	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			color.Green(scanner.Text())
		}
		scanner = bufio.NewScanner(stderr)
		for scanner.Scan() {
			color.Red(scanner.Text())
		}
	}()
}

func (co *Compiler) killProcess() {
	commandPr := exec.Command("lsof", "-t", "-i", ":8000")
	ou, err := commandPr.Output()
	if err != nil {
		return
	}
	pid := strings.TrimSpace(string(ou))
	pid = strings.Split(pid, "\n")[0]
	if pid != "" {
		commandKill := exec.Command("kill", pid)
		var out, stderr strings.Builder
		commandKill.Stdout = &out
		commandKill.Stderr = &stderr
		err := commandKill.Run()
		if err != nil {
			color.Blue(stderr.String() + " " + pid + "=> 2000")
		}
	}

	commandPr = exec.Command("lsof", "-t", "-i", ":3000")
	ou, err = commandPr.Output()
	if err != nil {
		return
	}
	pid = strings.TrimSpace(string(ou))
	pid = strings.Split(pid, "\n")[0]
	if pid != "" {
		commandKill := exec.Command("kill", pid)
		var out, stderr strings.Builder
		commandKill.Stdout = &out
		commandKill.Stderr = &stderr
		err := commandKill.Run()
		if err != nil {
			color.Blue(stderr.String() + " " + pid + "=> 3000")
		}
	}
}
