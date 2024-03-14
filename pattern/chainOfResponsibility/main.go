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

// next is database logger
type DBLogger struct {
	level int // 0 - debug, 1 - warning, 2 - error
	next  Logger
}

func NewDBLogger(l int) Logger {
	return &DBLogger{level: l}
}

func (s *DBLogger) SetNext(nextLogger Logger) {
	s.next = nextLogger
}

func (s *DBLogger) Message(level int, msg string) {
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

func (s *DBLogger) writeMessage(msg string) {
	fmt.Printf("from db logger (saving this log in database): %d: %s\n", s.level, msg)
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

func (s *FileLogger) Message(level int, msg string) {
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

func (s *FileLogger) writeMessage(msg string) {
	fmt.Printf("from file logger (writing this log in file): %d: %s\n", s.level, msg)
}
