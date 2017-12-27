package main

import (
    "errors"
)

// Queue is a slice which implement CRUD
type List []string

var (
    ErrNotExist = errors.New("Index does not exist")
)

func New(sl []string) *List {
    l := List(sl)
    return &l
}

func (l *List) Len() int {
    return len(*l)
}

func (l *List) Has(i int) bool {
    return l.Len() > i && i >= 0
}

func (l *List) Find(s string) (bool, int) {
    for i, e := range *l {
        if e == s {
            return true, i
        }
    }
    return false, -1
}

func (l *List) Create(s string) {
    p := *l
    p = append(p, s)
    *l = p
}

func (l *List) Read(i int) (string, error) {
    if l.Has(i) {
        return (*l)[i], nil
    }
    return "", ErrNotExist
}

func (l *List) Update(i int, s string) error {
    if !l.Has(i) {
        return ErrNotExist
    }

    (*l)[i] = s
    return nil
}

func (l *List) Delete(i int) {
    if l.Has(i) {
        p := *l
        p = append(p[:i], p[i+1:]...)
        *l = p
    }
}

func (l *List) Clone() *List {
    c := make(List, l.Len())
    copy(c, *l)
    return &c
}

func (l *List) Each(f func (i int, s string)) {
    for i, s := range *l {
        f(i, s)
    }
}
