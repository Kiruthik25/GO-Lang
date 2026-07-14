package students

import "errors"


type User struct{

	ID int
	Name string
	Age int

}

type UserRepo struct{
	users map[int]User
}

func NewUserRepo() *UserRepo{

	return &UserRepo{
		users : make(map[int]User),
	}
}


func (r *UserRepo) Create (user User) error{

		if _, exists := r.users[user.ID];
		 exists {
		return errors.New("user already exists")
	}
	r.users[user.ID] = user
	return nil
}


func (r *UserRepo) Get (id int) (User , bool){
		user , ok := r.users[id]

		return user , ok

}