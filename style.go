package main

import (
    "container/list"
    "github.com/gdamore/tcell"
)

type StyleSpan struct {
    gap, len int
    v tcell.Style
}

type StyleList struct {
    v *list.List
    end int
}

func NewStyleList() StyleList {
    return StyleList{ list.New(), 0 }
}

func (sl *StyleList) PushBack(idx, length int, style tcell.Style) {
    sl.v.PushBack(StyleSpan{ idx - sl.end, length, style })
    sl.end = idx + length
}

func (sl *StyleList) Iterate() StyleIterator {
    return StyleIterator{ 0, sl.v.Front() }
}

type StyleIterator struct {
    step int
    span *list.Element
}

func (it *StyleIterator) Step() tcell.Style {
    if it.span == nil {
        return tcell.StyleDefault
    }

    span := it.span.Value.(StyleSpan)
    it.step++

    if it.step <= span.gap {
        return tcell.StyleDefault
    }

    if it.step < span.gap + span.len {
        return span.v
    }

    it.step = 0
    it.span = it.span.Next()
    return span.v
}
