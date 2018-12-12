package main

import (
    "github.com/gdamore/tcell"
)

type Track struct {
    text []rune
    style StyleList
}

func NewTrack(w int) Track {
    return Track{ make([]rune, w), NewStyleList() }
}

func (t *Track) Len() int {
    return len(t.text)
}

func (t *Track) Fill(s string, style tcell.Style) {
    t.FillAt(s, 0, style)
}

func (t *Track) FillAt(s string, idx int, style tcell.Style) {
    i := 0
    runes := []rune(s)

    for idx + i < len(t.text) && i < len(s) {
        t.text[idx + i] = runes[i]
        i++
    }

    t.style.PushBack(idx, i, style)
    for idx + i < len(t.text) {
        t.text[idx + i] = ' '
        i++
    }
}

func (t *Track) Render(screen tcell.Screen, y int) {
    st := t.style.Iterate()
    for x := 0; x < t.Len(); x++ {
        screen.SetContent(x, y, t.text[x], nil, st.Step())
    }
}

func (t *Track) Clear() {
    for i := 0; i < len(t.text); i++ {
        t.text[i] = ' '
    }
}
