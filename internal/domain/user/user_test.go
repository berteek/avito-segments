package user

import (
    "testing"
)

func TestNewUser(t *testing.T) {
    id := 42
    u := NewUser(id)

    returnedId := u.GetId()
    if id != u.GetId() {
        t.Fatalf("GetId() returns incorrect id. Got %d, expected %d", returnedId, id)
    }
}

func TestIsEqual(t *testing.T) {
    id1 := 1
    id2 := 2
    u1 := NewUser(id1)
    u2 := NewUser(id2)

    if u1.IsEqual(u2) {
        t.Fatalf("Users shouldn't be equal! u1 is %+v, u2 is %+v", u1, u2)
    }

    id1 = 1
    id2 = 1
    u1 = NewUser(id1)
    u2 = NewUser(id2)

    if !u1.IsEqual(u2) {
        t.Fatalf("Users should be equal! u1 is %+v, u2 is %+v", u1, u2)
    }
}
