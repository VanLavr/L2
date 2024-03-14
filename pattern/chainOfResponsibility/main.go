package main

import (
	"fmt"
)

type Logger interface {
	SetNext(Logger)
	Message(int, string)
	writeMessage(string)
}

func main() {
	std := NewStdoutLogger(0)
	file := NewFileLogger(1)
	db := NewDBLogger(2)

	std.SetNext(file)
	file.SetNext(db)

	std.Message(0, "this will be printed only in std, it is a DEBUG message")
	std.Message(1, "this will be printed in std and in file, it is a WARNING message")
	std.Message(2, "this will be printed in std, in file and wil be saved in database, it is an ERROR message")
}

// first class in chain - stdout logger
type StdoutLogger struct {
	level int // 0 - debug, 1 - warning, 2 - error
	next  Logger
}

func NewStdoutLogger(l int) Logger {
	return &StdoutLogger{level: l}
}

func (s *StdoutLogger) SetNext(nextLogger Logger) {
	s.next = nextLogger
}

func (s *StdoutLogger) Message(level int, msg string) {
	if level <= s.level {
		s.writeMessage(msg)
	} else {
		s.writeMessage(msg)
		if s.next != nil {
			s.next.Message(level, msg)
		} else {
			fmt.Println("ERR: no next logger was declared")
		}
	}
}

func (s *StdoutLogger) writeMessage(msg string) {
	fmt.Printf("from stdout logger: %d: %s\n", s.level, msg)
}

// next is database logger (has an extra method - saving log in database)
type DBLogger struct {
	level int // 0 - debug, 1 - warning, 2 - error
	next  Logger
}

func NewDBLogger(l int) Logger {
	return &DBLogger{level: l}
}

func (d *DBLogger) SetNext(nextLogger Logger) {
	d.next = nextLogger
}

func (d *DBLogger) Message(level int, msg string) {
	if level <= d.level {
		d.writeMessage(msg)
		d.saveInDataBase(msg)
	} else {
		d.writeMessage(msg)
		d.saveInDataBase(msg)
		if d.next != nil {
			d.next.Message(level, msg)
		} else {
			fmt.Println("ERR: no next logger was declared")
		}
	}
}

func (d *DBLogger) writeMessage(msg string) {
	fmt.Printf("from db logger: %d: %s\n", d.level, msg)
}

func (d *DBLogger) saveInDataBase(msg string) {
	fmt.Printf("saving this log in database (%s)\n", msg)
}

// next is file logger
type FileLogger struct {
	level int // 0 - debug, 1 - warning, 2 - error
	next  Logger
}

func NewFileLogger(l int) Logger {
	return &FileLogger{level: l}
}

func (s *FileLogger) SetNext(nextLogger Logger) {
	s.next = nextLogger
}

func (f *FileLogger) Message(level int, msg string) {
	if level <= f.level {
		f.writeMessage(msg)
		f.writeToFile(msg)
	} else {
		f.writeMessage(msg)
		f.writeToFile(msg)
		if f.next != nil {
			f.next.Message(level, msg)
		} else {
			fmt.Println("ERR: no next logger was declared")
		}
	}
}

func (f *FileLogger) writeMessage(msg string) {
	fmt.Printf("from file logger: %d: %s\n", f.level, msg)
}

func (f *FileLogger) writeToFile(msg string) {
	fmt.Printf("writing this log to a file (%s)\n", msg)
}
