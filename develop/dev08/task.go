package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func handler(input string) error {
	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSpace(input)
	args := strings.Split(input, " ")
	switch args[0] {
	case "cd":
		if len(args) == 1 {
			log.Printf("Error! You must choose directory to move")
			return nil
		}
		return changeDir(args[1])
	case "pwd":
		return printWorkingDirectory()
	case "echo":
		return echo(args[1:])
	case "kill":
		if len(args) == 1 {
			log.Printf("Error! You must choose task to kill")
			return nil
		}
		return kill(args[1])
	case "ps":
		return ps()
	case "quit":
		os.Exit(0)
	default:
		fmt.Println(args[0])

		return errors.New("invalid command")
	}
	return nil
}

func changeDir(args string) error {
	err := os.Chdir(strings.Replace(args, "cd", "", 1))
	if err != nil {
		return err
	}
	return nil
}

func printWorkingDirectory() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", dir)
	return nil
}

func echo(args []string) error {
	var echoString string
	for i := range args {
		if i == 0 {
			echoString = args[i]
			continue
		}
		echoString += " " + args[i]
	}
	out, err := exec.Command("cmd", "/C", "echo", echoString).Output()
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", out)
	return nil
}

func kill(args string) error {
	out, err := exec.Command("cmd", "/C", "TASKKILL", "/PID", args).Output()

	if err != nil {
		return err
	}
	fmt.Printf("%s\n", out)

	return nil
}

func ps() error {
	out, err := exec.Command("cmd", "/C", "TASKLIST").Output()

	if err != nil {
		return err
	}
	fmt.Printf("%s\n", out)

	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		if err = handler(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
