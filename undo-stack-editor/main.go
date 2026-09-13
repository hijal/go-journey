package main

import (
	"errors"
	"fmt"
	"slices"
)

type Editor struct {
	lines   []string
	history [][]string
}

func newEditor() *Editor {
	return &Editor{
		lines:   []string{},
		history: [][]string{},
	}
}

func (e *Editor) snapshot() {
	e.history = append(e.history, slices.Clone(e.lines))
}

func (e *Editor) Insert(line string) {
	e.snapshot()
	e.lines = append(e.lines, line)
}

func (e *Editor) DeleteLast() error {
	if len(e.lines) == 0 {
		return errors.New("nothing to delete")
	}
	e.snapshot()
	e.lines = e.lines[:len(e.lines)-1]
	return nil
}

func (e *Editor) Undo() error {
	if len(e.history) == 0 {
		return errors.New("nothing to undo")
	}
	e.lines = e.history[len(e.history)-1]
	e.history = e.history[:len(e.history)-1]
	return nil
}

func (e *Editor) View() []string {
	return slices.Clone(e.lines)
}

func main() {
	ed := newEditor()
	ed.Insert("line 1")
	ed.Insert("line 2")
	ed.Insert("line 3")

	fmt.Println("after inserts:", ed.View())

	if err := ed.DeleteLast(); err != nil {
		fmt.Println("delete error:", err)
	}
	fmt.Println("after delete:", ed.View())
	if err := ed.Undo(); err != nil {
		fmt.Println("undo error:", err)
	}
	fmt.Println("after undo:", ed.View())
}
