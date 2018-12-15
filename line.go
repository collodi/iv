package main

import (
    "fmt"
    "container/list"
)

type Line struct {
    num int
    text string
    style *StyleList
}

func (ln *Line) Len() int {
    return len(ln.text)
}

// lnw = line number width
func (ln *Line) FillTracks(lnw int, tracks []Track) int {
    idx := 0
    trk := 0

    linenum := fmt.Sprintf("%*d", lnw - 1, ln.num)
    tracks[trk].Fill(linenum)
    tracks[trk].style.InsertAt(0, lnw, StyleGray)

    for trk < len(tracks) {
        tracks[trk].FillAt(ln.text[idx:], lnw)
        idx += tracks[trk].Len()

        trk++
        if idx >= ln.Len() {
            break
        }
    }
    return trk
}

type Lines struct {
    v *list.List
    lnw int
    cursor Cursor
}

func NewLines() Lines {
    return Lines{ list.New(), 0, Cursor{ 0, 0 } }
}

func (lns *Lines) PushLine(line string) {
    ln := Line{ lns.v.Len() + 1, line, NewStyleList() }
    lns.v.PushBack(ln)
    lns.lnw = len(string(ln.num)) + 2
}

func (lns *Lines) CurrElem() *list.Element {
    ln := lns.v.Front()
    for i := 0; i < lns.cursor.y; i++ {
        ln = ln.Next()
    }
    return ln
}

func (lns *Lines) Jump(n int) {
    y := lns.cursor.y + n
    lns.cursor.y = Saturate(0, lns.v.Len() - 1, y)
}

func (lns *Lines) FillTracks(tracks []Track) {
    trknum := 0
    ln := lns.CurrElem()
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
