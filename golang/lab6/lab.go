package lab6

import "fmt"

type Television struct {
	channel int
}

func (t *Television) GetChannel() int {
	return t.channel
}

func (t *Television) SetChannel(channel int) {
	(*t).channel = channel
}

func (t *Television) Display() {
	channel := t.GetChannel()
	fmt.Println("╔════════════════╗")
	fmt.Println("║                ║")
	fmt.Println("║   Канал", channel, "     ║")
	fmt.Println("║                ║")
	fmt.Println("╚════════════════╝")
}

func RunLab6() {
	var tv Television = Television{4}
	var tvPointer *Television = &tv
	tvPointer.SetChannel(5)
	tvPointer.Display()

	tvPointer.SetChannel(12)
	tvPointer.Display()
}
