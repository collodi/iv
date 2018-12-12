package main

import (
    "github.com/gdamore/tcell"
)

type Track struct {
    text []rune
}

func NewTrack(w int) Track {
    return Track{ make([]rune, w) }
}

func (t *Track) Len() int {
    return len(t.text)
}

func (t *Track) Assign(s string, idx int) {
    i := 0
    runes := []rune(s)

    for i < len(t.text) && idx + i < len(s) {
        t.text[i] = runes[idx + i]
        i++
    }

    for i < len(t.text) {
        t.text[i] = ' '
        i++
    }
}

func (t *Track) Render(screen tcell.Screen, y int) {
    for x := 0; x < t.Len(); x++ {
        screen.SetContent(x, y, t.text[x], nil, tcell.StyleDefault)
    }
}

func (t *Track) Clear() {
    for i := 0; i < len(t.text); i++ {
        t.text[i] = ' '
    }
}
