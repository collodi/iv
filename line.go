package main

import (
)

type Line struct {
    num int
    text string
}

func (ln *Line) Len() int {
    return len(ln.text)
}

func (ln *Line) FillTracks(tracks []Track) int {
    idx := 0
    trknum := 0

    for trknum < len(tracks) {
        tracks[trknum].Assign(ln.text, idx)
        idx += tracks[trknum].Len()

        trknum++
        if idx >= ln.Len() {
            break
        }
    }
    return trknum
}
