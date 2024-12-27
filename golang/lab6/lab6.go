package lab6

import "fmt"

type Television struct {
	brand   string
	channel int
	price   int
}

func (t Television) GetChannel() int {
	return t.channel
}

func (t *Television) SetChannel(channel int) {
	(*t).channel = channel
}

func (t Television) GetBrand() string {
	return t.brand
}

func (t *Television) SetBrand(brand string) {
	(*t).brand = brand
}

func (t Television) GetPrice() int {
	return t.price
}

func (t *Television) SetPrice(price int) {
	(*t).price = price
}

func (t Television) GetInfo() {
	fmt.Println("Телевизор", t.brand, "стоит", t.price, "и на нем включен канал", t.channel)
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
	var tv Television = Television{"Samsung", 4, 25000}
	var tvPointer *Television = &tv
	tvPointer.GetInfo()
	tvPointer.SetChannel(5)
	tvPointer.Display()

	tvPointer.SetChannel(12)
	tvPointer.Display()

	tvPointer.SetPrice(27000)
	tvPointer.SetBrand("Xiaomi")
	tvPointer.GetInfo()
}
