package segment

import (
    "github.com/berteek/avito-segments/internal/domain/user"
)

type Segment struct {
    id    int
    name  string
    users []user.User
}

func NewSegment(id int, name string, users []user.User) Segment {
    return Segment{
        id:    id,
        name:  name,
        users: users,
    }
}

func (s Segment) GetId() int {
    return s.id
}

func (s Segment) GetName() string {
    return s.name
}

func (s Segment) GetUsers() []user.User {
    return s.users
}
