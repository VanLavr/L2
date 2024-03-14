package main

import "fmt"

type Command interface {
	Execute()
	Undo()
}

type TV struct{}

func (t *TV) TurnOn() {
	fmt.Println("tv is running")
}

func (t *TV) TurnOff() {
	fmt.Println("tv is sleeping")
}

type TvOnCommand struct {
	tv *TV
}

func NewTvOnCommand(tv *TV) *TvOnCommand {
	return &TvOnCommand{tv: tv}
}

func (tc *TvOnCommand) Execute() {
	tc.tv.TurnOn()
}

func (tc *TvOnCommand) Undo() {
	tc.tv.TurnOff()
}

type Pult struct {
	c Command
}

func (p *Pult) SetCommand(c Command) {
	p.c = c
}

func (p *Pult) PressButton() {
	p.c.Execute()
}

func (p *Pult) PressUndoButton() {
	p.c.Undo()
}

func main() {
	tv := &TV{}
	pult := &Pult{}

	pult.SetCommand(NewTvOnCommand(tv))

	pult.PressButton()
	pult.PressUndoButton()
}
