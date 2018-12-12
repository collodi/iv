package main

import (
    "github.com/gdamore/tcell"
)

type FrontEnd struct {
    screen tcell.Screen
    vw *View
}

func NewFrontEnd() (*FrontEnd, error) {
    screen, e := tcell.NewScreen()
    if e != nil {
        return nil, e
    }

    if e := screen.Init(); e != nil {
        return nil, e
    }

    return &FrontEnd{ screen, nil }, nil
}

func (fe *FrontEnd) SetView(v *View) {
    fe.vw = v
}

func (fe *FrontEnd) ListenTo(ch chan Res) {
    for _ = range ch {
        fe.UpdateScreen()
    }
}

func (fe *FrontEnd) UpdateScreen() {
    fe.screen.Clear()
    for i := 0; i < fe.vw.height; i++ {
        fe.vw.tracks[i].Render(fe.screen, i)
    }

    cursor := fe.vw.tcursor
    fe.screen.ShowCursor(cursor.x, cursor.y)

    fe.screen.Show()
}

func (fe *FrontEnd) Close() {
    fe.screen.Fini()
}
