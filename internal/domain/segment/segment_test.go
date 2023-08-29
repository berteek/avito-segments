package segment

import (
    "testing"

    "github.com/berteek/avito-segments/internal/domain/user"
)

func areEqual(u1, u2 []user.User) bool {
    if len(u1) != len(u2) {
        return false
    }
    for i := range u1 {
        if !u1[i].IsEqual(u2[i]) {
            return false
        }
    }
    return true
}

func TestAreEqual(t *testing.T) {
    users1 := []user.User{user.NewUser(1), user.NewUser(2), user.NewUser(3)}
    users2 := []user.User{user.NewUser(2), user.NewUser(2), user.NewUser(3)}

    if areEqual(users1, users2) {
        t.Fatalf("areEqual(users1, users2) is not working right. users1(=%v), users2(=%v) should not be equal!", users1, users2)
    }

    users1 = []user.User{user.NewUser(1), user.NewUser(2), user.NewUser(3)}
    users2 = []user.User{user.NewUser(1), user.NewUser(2)}

    if areEqual(users1, users2) {
        t.Fatalf("areEqual(users1, users2) is not working right. users1(=%v), users2(=%v) should not be equal!", users1, users2)
    }

    users1 = []user.User{user.NewUser(1), user.NewUser(2), user.NewUser(3)}
    users2 = []user.User{user.NewUser(1), user.NewUser(2), user.NewUser(3)}

    if !areEqual(users1, users2) {
        t.Fatalf("areEqual(users1, users2) is not working right. users1(=%v), users2(=%v) should be equal!", users1, users2)
    }
}

func TestNewSegment(t *testing.T) {
    id := 42
    name := "AVITO_TESTING"
    users := make([]user.User, 10)
    for i := range users {
        users[i] = user.NewUser(i)
    }

    s := NewSegment(id, name, users)

    returnedId := s.GetId()
    if id != returnedId {
        t.Fatalf("GetId() returns incorrect id. Got %d, expected %d", returnedId, id)
    }

    returnedName := s.GetName()
    if name != returnedName {
        t.Fatalf("GetName() returns incorrect name. Got %s, expected %s", returnedName, name)
    }

    returnedUsers := s.GetUsers()
    if !areEqual(users, returnedUsers) {
        t.Fatalf("GetUsers() returns incorrect users. Got %v, expected %v", returnedUsers, users)
    }
}
