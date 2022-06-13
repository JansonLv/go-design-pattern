package main

import "fmt"

type user struct {
	name string
	age  int
}

// 具体迭代器
type userIterator struct {
	index int
	users []*user
}

func (u *userIterator) hasNext() bool {
	return u.index < len(u.users)
}
func (u *userIterator) getNext() *user {
	if u.hasNext() {
		user := u.users[u.index]
		u.index++
		return user
	}
	return nil
}

type userCollection struct {
	users []*user
}

func (u *userCollection) createIterator() iterator {
	return &userIterator{
		users: u.users,
	}
}

type iterator interface {
	hasNext() bool
	getNext() *user
}

func main() {
	user1 := &user{
		name: "a",
		age:  30,
	}
	user2 := &user{
		name: "b",
		age:  20,
	}

	userCollection := &userCollection{
		users: []*user{user1, user2},
	}

	iterator := userCollection.createIterator()

	for iterator.hasNext() {
		user := iterator.getNext()
		fmt.Printf("User is %+v\n", user)
	}
}
