package dataobjects

type User struct {
  ID string
  Name string
}

type Organization struct {
  User *User
}
