package user

type User struct {
    id int
}

func NewUser(id int) User {
    return User{id: id}
}

func (u User) GetId() int {
    return u.id
}

func (u User) IsEqual(ou User) bool {
    return u.id == ou.id
}
