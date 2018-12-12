package main

const (
    ModeNormal = "NORMAL"
    ModeInsert = "INSERT"
)

type Mode string
type Command int

type Cursor struct {
    x, y int
}

type Req struct {
    cmd Command
}

type Res int
