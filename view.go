package main

import (
    "os"
    "bufio"
    "reflect"
)

type View struct {
    width, height int

    mode Mode
    file string
    lines Lines
    tracks []Track

    signal chan int
}

func NewView(file string, w, h int) (*View, error) {
    f, err := os.Open(file)
    if err != nil {
        return nil, err
    }
    defer f.Close()

    lines := NewLines()
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        lines.PushLine(scanner.Text())
    }

    tracks := make([]Track, h)
    for i := 0; i < h; i++ {
        tracks[i] = NewTrack(w)
    }

    v := View{ w, h, ModeNormal, file, lines, tracks, make(chan int) }
    return &v, nil
}

func (v *View) Update() {
    v.lines.FillTracks(v.tracks[:v.height - 1])
    v.UpdateStatusBar()

    v.signal <- 0
}

func (v *View) UpdateStatusBar() {
    track := &v.tracks[v.height - 1]

    track.style.InsertAt(0, 8, StyleBgDark)
    track.style.ClearFrom(9)

    track.FillAt(string(v.mode), 1)
    track.FillAt(v.file, 9)
}

func (v *View) TrackCursor() Cursor {
    width := v.width - v.lines.lnw

    x := v.lines.lnw + (v.lines.cursor.x % width)
    y := 0 + (v.lines.cursor.x / width)
    return Cursor{ x, y }
}

func (v *View) Cmd(name string, args ...interface{}) {
    argvals := make([]reflect.Value, len(args))
    for i, _ := range args {
        argvals[i] = reflect.ValueOf(args[i])
    }
    reflect.ValueOf(v).MethodByName(name).Call(argvals)
}
