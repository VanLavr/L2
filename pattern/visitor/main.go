package main

import "fmt"

type Transport interface {
	GetType() string
	Accept(Visitor)
}

type Bicycle struct{}

func (b *Bicycle) Accept(v Visitor) {
	v.VisitBicycle(b)
}

func (b *Bicycle) GetType() string {
	return "bicycle"
}

type Car struct{}

func (c *Car) Accept(v Visitor) {
	v.VisitCar(c)
}

func (c *Car) GetType() string {
	return "car"
}

type Tram struct{}

func (t *Tram) Accept(v Visitor) {
	v.VisitTram(t)
}

func (t *Tram) GetType() string {
	return "tram"
}

type Visitor interface {
	VisitBicycle(Transport)
	VisitCar(Transport)
	VisitTram(Transport)
}

type Passanger struct{}

func (p *Passanger) VisitBicycle(t Transport) {
	whatDoWeRiding := t.GetType()
	fmt.Println("getting huge quads on", whatDoWeRiding)
}

func (p *Passanger) VisitCar(t Transport) {
	whatDoWeRiding := t.GetType()
	fmt.Println("hit the traffic jam on", whatDoWeRiding)
}

func (p *Passanger) VisitTram(t Transport) {
	whatDoWeRiding := t.GetType()
	fmt.Println("pretending like you have a ticket on", whatDoWeRiding)
}

type Customer struct{}

func (c *Customer) VisitBicycle(t Transport) {
	bought := t.GetType()
	fmt.Println("Congrads, you bougt a", bought)
}

func (c *Customer) VisitCar(t Transport) {
	bought := t.GetType()
	fmt.Println("Congrads, you bougt a", bought)
}

func (c *Customer) VisitTram(t Transport) {
	bought := t.GetType()
	fmt.Println("Who the heck are gonna buy a", bought)
}

func main() {
	b := &Bicycle{}
	c := &Car{}
	t := &Tram{}

	Passanger := &Passanger{}
	b.Accept(Passanger)
	c.Accept(Passanger)
	t.Accept(Passanger)

	fmt.Println()

	Customer := &Customer{}
	b.Accept(Customer)
	c.Accept(Customer)
	t.Accept(Customer)
}
