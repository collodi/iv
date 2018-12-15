package main

type Mode string
const (
    ModeNormal = "NORMAL"
    ModeInsert = "INSERT"
)

type Cursor struct {
    x, y int
}
