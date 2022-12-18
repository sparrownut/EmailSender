package utils

import (
	"fmt"
	"github.com/fatih/color"
)

func Printsuc(text string, args ...any) {
	c := color.New(color.FgHiGreen, color.Bold)
	_, err := c.Print("[+]")
	if err != nil {
		return
	}
	fmt.Printf(text+"\n", args...)

}
func Printerr(text string, args ...any) {
	c := color.New(color.FgHiRed)
	_, err := c.Print("[-]")
	if err != nil {
		return
	}
	fmt.Printf(text+"\n", args...)
}
func Printminfo(text string, args ...any) {
	c := color.New(color.FgYellow)
	_, err := c.Print("[>]")
	if err != nil {
		return
	}
	fmt.Printf(text+"\n", args...)
}
func Printhinfo(text string, args ...any) {
	c := color.New(color.FgHiYellow, color.Bold)
	_, err := c.Print("[!]")
	if err != nil {
		return
	}
	fmt.Printf(text+"\n", args...)
}
func Printcritical(text string, args ...any) {
	c := color.New(color.FgHiBlue, color.BgHiRed, color.Bold)
	_, err := c.Print("[-]")
	if err != nil {
		return
	}
	_, err = c.Printf(text+"\n", args...)
	if err != nil {
		return
	}
	fmt.Print()
}
