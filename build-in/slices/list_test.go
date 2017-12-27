package main

import (
    "testing"
)

func TestLen(t *testing.T) {
    if New([]string{}).Len() != 0 {
        t.Error("Unexpected result")
    }

    if New([]string{"1", "2"}).Len() != 2 {
        t.Error("Unexpected result")
    }
}

func TestHas(t *testing.T) {
    if New([]string{}).Has(5) {
        t.Error("Unexpected result")
    }

    l := New([]string{"1", "2"})
    if !l.Has(0) || !l.Has(1) || l.Has(2) {
        t.Error("Unexpected result")
    }
}

func TestFind(t *testing.T) {
    l := New([]string{"1", "2", "3", "4", "5"})
    if ok, i := l.Find("3"); !ok || i != 2 {
        t.Error("Unexpected result")
    }

    if ok, i := l.Find("A"); ok || i > 0 {
        t.Error("Unexpected result")
    }
}

func TestCreate(t *testing.T) {
    l := New([]string{"1", "2", "3", "4", "5"})
    l.Create("6")
    if ok, i := l.Find("6"); !ok || i != 5 {
        t.Errorf("Unexpected result")
    }
}

func TestRead(t *testing.T) {
    l := New([]string{"1", "2", "3", "4", "5"})
    if s, err := l.Read(0); s != "1" || err != nil {
        t.Errorf("Unexpected result: %#v %#v", s, err)
    }

    if s, err := l.Read(4); s != "5" || err != nil {
        t.Errorf("Unexpected result: %#v %#v", s, err)
    }

    if s, err := l.Read(5); s != "" || err == nil {
        t.Errorf("Unexpected result: %#v %#v", s, err)
    }
}

func TestUpdate(t *testing.T) {
    l := New([]string{"2"})
    l.Update(0, "1")
    if s, err := l.Read(0); s != "1" && err != nil {
        t.Errorf("Unexpected result: %#v", s)
    }
}

func TestDelete(t *testing.T) {
    l := New([]string{"1", "2"})
    l.Delete(1)
    if l.Len() != 1 {
        t.Errorf("Unexpected result: %#v", l)
    }
    l.Delete(0)
    if l.Len() != 0 {
        t.Errorf("Unexpected result: %#v", l)
    }
    l.Delete(-1)
    if l.Len() != 0 {
        t.Errorf("Unexpected result: %#v", l)
    }
    l.Delete(1000)
    if l.Len() != 0 {
        t.Errorf("Unexpected result: %#v", l)
    }
}

func TestClone(t *testing.T) {
    l := New([]string{"1"})
    c := l.Clone()

    if &(*l)[0] == &(*c)[0] || (*l)[0] != (*c)[0] {
        t.Errorf("Unexpected result")
    }
}

func TestEach(t *testing.T) {
    var (
        gi int
        gs string
    )

    l := New([]string{"1"})
    l.Each(func(i int, s string) {
        gi = i
        gs = s
    })

    if gi != 0 || gs != "1" {
        t.Error("Unexpected result")
    }
}
