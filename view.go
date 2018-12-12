package main

import (
    "os"
    "bufio"
    "container/list"
)

type View struct {
    width, height int

    mode Mode
    lines *list.List
    lcursor Cursor
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

    lines := list.New()
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        ln := Line{ lines.Len(), scanner.Text() }
        lines.PushBack(ln)
    }

    tracks := make([]Track, h)
    for i := 0; i < h; i++ {
        tracks[i] = NewTrack(w)
    }

    v := View{ w, h, ModeNormal, lines, Cursor{ 0, 0 }, tracks, Cursor{ 0, 0 }, make(chan Req), make(chan Res) }
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
    linenum := v.lcursor.y

    // skip to linenum
    ln := v.lines.Front()
    for i := 0; i < linenum; i++ {
        ln = ln.Next()
    }

    trknum := 0
    for trknum < v.height {
        line := ln.Value.(Line)
        trknum += line.FillTracks(v.tracks[trknum:])

        ln = ln.Next()
        if ln == nil {
            break
        }
    }

    for trknum < v.height {
        v.tracks[trknum].Clear()
        trknum++
    }
}
