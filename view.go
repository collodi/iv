package main

import (
    "os"
    "bufio"
)

type View struct {
    width, height int

    mode Mode
    file string
    lines Lines
    tracks []Track
    tcursor Cursor

    to chan Req
    from chan Res
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

    v := View{ w, h, ModeNormal, file, lines, tracks, Cursor{ 0, 0 }, make(chan Req), make(chan Res) }
    return &v, nil
}

func (v *View) Start() {
    v.UpdateTracks()
    v.from <- 0

    for _ = range v.to {
        v.UpdateTracks()
        v.from <- 0
    }
}

func (v *View) UpdateTracks() {
    v.lines.FillTracks(v.tracks[:v.height - 1])
    v.UpdateStatusBar()
}

func (v *View) UpdateStatusBar() {
    track := v.tracks[v.height - 1]

    track.FillAt(string(v.mode), 1)
    track.FillAt(v.file, 8)
}
