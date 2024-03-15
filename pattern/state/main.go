package main

import (
	"fmt"
	"time"
)

// state
type Mode interface {
	PressC_j()         // to enter navigation mode
	PressC_s()         // to enter selection mode
	PressC_i()         // to enter insertion mode
	TypeSomeText()     // to type some text
	SelectSomeText()   // to select some text
	GoToEndOfTheFile() // to go to the end of the file
}

func main() {
	Vim := NewTerminalTextEditor()

	Vim.currentMode.PressC_j()
	time.Sleep(time.Millisecond * 300)
	Vim.currentMode.GoToEndOfTheFile()
	time.Sleep(time.Millisecond * 300)
	Vim.currentMode.TypeSomeText()
	time.Sleep(time.Millisecond * 300)
	Vim.currentMode.SelectSomeText()
	time.Sleep(time.Second)

	Vim.currentMode.PressC_i()
	Vim.currentMode.PressC_i()
	time.Sleep(time.Millisecond * 300)
	Vim.currentMode.TypeSomeText()
	time.Sleep(time.Millisecond * 300)
	Vim.currentMode.GoToEndOfTheFile()
	time.Sleep(time.Millisecond * 300)
	Vim.currentMode.SelectSomeText()
	time.Sleep(time.Second)

	Vim.currentMode.PressC_s()
	Vim.currentMode.PressC_s()
	time.Sleep(time.Millisecond * 300)
	Vim.currentMode.SelectSomeText()
	time.Sleep(time.Millisecond * 300)
	Vim.currentMode.GoToEndOfTheFile()
	time.Sleep(time.Millisecond * 300)
	Vim.currentMode.TypeSomeText()

}

type TerminalTextEditor struct {
	currentMode    Mode
	insertMode     Mode
	selectionMode  Mode
	navigationMode Mode
}

func NewTerminalTextEditor() *TerminalTextEditor {
	editor := new(TerminalTextEditor)

	navigarionMode := new(NavigationMode)
	navigarionMode.t = editor
	insertMode := new(InsertMode)
	insertMode.t = editor
	selectMode := new(SelectMode)
	selectMode.t = editor

	editor.currentMode = navigarionMode
	editor.insertMode = insertMode
	editor.selectionMode = selectMode
	editor.navigationMode = navigarionMode

	return editor
}

func (t *TerminalTextEditor) setMode(mode Mode) {
	t.currentMode = mode
}

type InsertMode struct {
	t *TerminalTextEditor
}

func (i *InsertMode) PressC_i() {
	fmt.Println("(trying to switch to insert mode) you typed some stuff, you are in insertion mode already...")
}

func (i *InsertMode) PressC_j() {
	fmt.Println("switched to navigation mode")
	i.t.setMode(i.t.navigationMode)
}

func (i *InsertMode) PressC_s() {
	fmt.Println("switched to select mode")
	i.t.setMode(i.t.selectionMode)
}

func (i *InsertMode) TypeSomeText() {
	fmt.Println("typing...")
}

func (i *InsertMode) SelectSomeText() {
	fmt.Println("(trying to select something) you typed some stuff, you are in insertion mode...")
}

func (i *InsertMode) GoToEndOfTheFile() {
	fmt.Println("(trying to navigate) you typed some stuff, you are in insertion mode...")
}

type SelectMode struct {
	t *TerminalTextEditor
}

func (s *SelectMode) PressC_i() {
	fmt.Println("switched to insert mode")
	s.t.setMode(s.t.insertMode)
}

func (s *SelectMode) PressC_j() {
	fmt.Println("switched to navigation mode")
	s.t.setMode(s.t.navigationMode)
}

func (s *SelectMode) PressC_s() {
	fmt.Println("(trying to switch to select mode) nothing happend, you are already in select mode")
}

func (s *SelectMode) TypeSomeText() {
	fmt.Println("(trying to type some stuff) you selected some stuff, you are in selection mode...")
}

func (s *SelectMode) SelectSomeText() {
	fmt.Println("selecting...")
}

func (s *SelectMode) GoToEndOfTheFile() {
	fmt.Println("(trying to navigate) you selected some stuff, you are in selection mode...")
}

type NavigationMode struct {
	t *TerminalTextEditor
}

func (n *NavigationMode) PressC_i() {
	fmt.Println("switched to insert mode")
	n.t.setMode(n.t.insertMode)
}

func (n *NavigationMode) PressC_j() {
	fmt.Println("(trying to switch to navigation mode) you navigated somewhere, you are already in navigation mode...")
}

func (n *NavigationMode) PressC_s() {
	fmt.Println("switched to select mode")
	n.t.setMode(n.t.selectionMode)
}

func (n *NavigationMode) TypeSomeText() {
	fmt.Println("(trying to type some stuff) you navigated somewhere, you are in navigation mode...")
}

func (n *NavigationMode) SelectSomeText() {
	fmt.Println("(trying to select some stuff) you navigated somewhere, you are in navigation mode...")
}

func (n *NavigationMode) GoToEndOfTheFile() {
	fmt.Println("navigating...")
}
