package main

import (
    "fmt"
    "container/list"
)

type Line struct {
    num int
    text string
}

func (ln *Line) Len() int {
    return len(ln.text)
}

// lnw = line number width
func (ln *Line) FillTracks(lnw int, tracks []Track) int {
    idx := 0
    trknum := 0

    tracks[trknum].Fill(fmt.Sprintf("%*d", lnw - 1, ln.num))

    for trknum < len(tracks) {
        tracks[trknum].FillAt(ln.text[idx:], lnw + 1)
        idx += tracks[trknum].Len()

        trknum++
        if idx >= ln.Len() {
            break
        }
    }
    return trknum
}

type Lines struct {
    v *list.List
    lnw int
    cursor Cursor
}

func NewLines() Lines {
    return Lines{ list.New(), 3, Cursor{ 0, 0 } }
}

func (lns *Lines) PushLine(line string) {
    ln := Line{ lns.v.Len() + 1, line }
    lns.v.PushBack(ln)
    lns.lnw = len(string(ln.num)) + 2
}

func (lns *Lines) FillTracks(tracks []Track) {
    linenum := lns.cursor.y

    // skip to the line
    ln := lns.v.Front()
    for i := 0; i < linenum; i++ {
        ln = ln.Next()
    }

    trknum := 0
    for trknum < len(tracks) {
        line := ln.Value.(Line)
        trknum += line.FillTracks(lns.lnw, tracks[trknum:])

        ln = ln.Next()
        if ln == nil {
            break
        }
    }

    // tracks with no lines are filled with blanks
    for trknum < len(tracks) {
        tracks[trknum].Clear()
        trknum++
    }
}
