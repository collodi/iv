package main

import (
    "fmt"
    "github.com/gdamore/tcell"
)

func main() {
    fe, e := NewFrontEnd()
    if e != nil {
        fmt.Println("error creating a new frontend:", e)
        return
    }
    defer fe.Close()

    w, h := fe.screen.Size()
    vw, e := NewView("testfile", w, h)
    if e != nil {
        fmt.Println("error creating a new view:", e)
        return
    }

    fe.SetView(vw)
    go fe.ListenTo(vw.from)
    go vw.Start()

loop:
    for {
        ev := fe.screen.PollEvent()
        switch ev := ev.(type) {
            case *tcell.EventKey:
                if ev.Key() == tcell.KeyRune {
                    k := ev.Rune()
                    if k == 'q' {
                        break loop
                    }
                }
        }
    }
}
