package main

import (
	b "bufio"
	f "fmt"
	o "os"
	s "strings"
)

func main() {
	replacer := s.NewReplacer("\r\n", "", "\r", "", "\n", "", "%0a", "", "%0d", "")
	reader := b.NewReader(o.Stdin)
	f.Println("Findian")
	f.Println("======================")
	f.Print("Enter text: ")
	t, _ := reader.ReadString('\n')
	t = replacer.Replace(t)
	t = s.ToUpper(t)
	if s.HasPrefix(t, "I") && s.Contains(t, "A") && s.HasSuffix(t, "N") {
		f.Println("Found!")
	} else {
		f.Println("Not Found!")
	}
}
