package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"time"
)

var characters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func random(n int) string {
	buf := make([]rune, n)
	for i := range buf {
		buf[i] = characters[rand.Intn(len(characters))]
	}
	return string(buf)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	if len(os.Args) == 2 {
		switch os.Args[1] {
		case "help":
			fmt.Printf("matrix.exe [color]\n\tcolor = {black, red, green, yellow, blue, magenta, cyan, white}\nexample:\n\tmatrix.exe yellow")
			os.Exit(0)
		case "black":
			fmt.Print("\033[30;47m") // Black foreground and white background
		case "red":
			fmt.Print("\033[31;40m") // Red foreground and black background
		case "green":
			fmt.Print("\033[32;40m") // Green foreground and black background
		case "yellow":
			fmt.Print("\033[33;40m") // Yellow foreground and black background
		case "blue":
			fmt.Print("\033[34;40m") // Blue foreground and black background
		case "magenta":
			fmt.Print("\033[35;40m") // Magenta foreground and black background
		case "cyan":
			fmt.Print("\033[36;40m") // Cyan foreground and black background
		case "white":
			fmt.Print("\033[37;40m") // White foreground and black background
		default:
			fmt.Printf("matrix.exe [color]\n\tcolor = {black, red, green, yellow, blue, magenta, cyan, white}\nexample:\n\tmatrix.exe yellow")
			os.Exit(0)
		}
	} else {
		fmt.Print("\033[32;40m") // Green foreground and black background
	}

	fmt.Print("\033[2J") // Clear screen

	for i := 0; i < 1; i++ {
		go func() {
			for {
				fmt.Print(random(100))
			}
		}()
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	for {
		select {
		case <-interrupt:
			os.Exit(0)
			fmt.Print("\033[2J") // Clear screen
			return
		}
	}
}
