package main

func (v *View) CursorX(n int) {
    ln := v.lines.CurrElem().Value.(Line)

    x := v.lines.cursor.x + n
    v.lines.cursor.x = Saturate(0, ln.Len() - 1, x)
}

func (v *View) CursorY(n int) {
    v.lines.Jump(n)
    v.CursorX(0)
}
