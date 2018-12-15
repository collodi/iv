package main

import (
    "container/list"
    "github.com/gdamore/tcell"
)

const (
    StyleNormal = 432346921437233152
    StyleGray = 432346818358018048
    StyleHighlight = 432346857012723712
    StyleYellow = 432347187725205504
    StyleBlue = 432346814063050752
    StyleRed = 432345568522534912
    StyleGreen = 432346835537887232
    StylePink = 432346809768083456
    StyleLime = 432345607177240576

    StyleBgDark = 432346921437233160
    StyleBgLight = 432345564227567932
)

type StyleSpan struct {
    gap, len int
    v tcell.Style
}

type StyleList struct {
    v *list.List
    end int
}

func NewStyleList() *StyleList {
    return &StyleList{ list.New(), 0 }
}

func (sl *StyleList) Clear() {
    sl.v.Init()
    sl.end = 0
}

func (sl *StyleList) ClearFrom(i int) {
    if sl.end <= i {
        return
    }

    e := 0
    for el := sl.v.Front(); el != nil; el = el.Next() {
        sp := el.Value.(StyleSpan)
        b := e + sp.gap
        e = b + sp.len

        if b >= i {
            sl.v.Remove(el)
        } else if e >= i {
            sp := el.Value.(StyleSpan)
            sp.len -= (e - i) + 1
        }
    }
}

func (sl *StyleList) PushBack(gap, size int, style tcell.Style) {
    sl.v.PushBack(StyleSpan{ gap, size, style })
    sl.end += gap + size
}

func (sl *StyleList) InsertAt(i, size int, style tcell.Style) {
    if i >= sl.end {
        sl.PushBack(i - sl.end, size, style)
    } else {
        arr := sl.ToArr(Max(i + size, sl.end))
        for k := 0; k < size; k++ {
            arr[i + k] = style
        }
        sl.FromArr(arr)
    }
}

func (sl *StyleList) ToArr(n int) []tcell.Style {
    it := sl.Iterate()
    arr := make([]tcell.Style, n)
    for i := 0; i < n; i++ {
        arr[i] = it()
    }
    return arr
}

func (sl *StyleList) FromArr(arr []tcell.Style) {
    sl.Clear()
    if len(arr) == 0 {
        return
    }

    gap := 0
    size := 0
    prev := arr[0]
    for _, v := range arr {
        if prev == v {
            size += 1
            continue
        }

        if prev == StyleNormal {
            gap = size
        } else {
            sl.PushBack(gap, size, prev)
            gap = 0
        }

        size = 1
        prev = v
    }

    if prev != StyleNormal {
        sl.PushBack(gap, size, prev)
    }
}

func (sl *StyleList) Iterate() func() tcell.Style {
    step := 0
    span := sl.v.Front()

    return func() tcell.Style {
        if span == nil {
            return StyleNormal
        }

        sp := span.Value.(StyleSpan)
        step++

        if step <= sp.gap {
            return StyleNormal
        }

        if step < sp.gap + sp.len {
            return sp.v
        }

        step = 0
        span = span.Next()
        return sp.v
    }
}
