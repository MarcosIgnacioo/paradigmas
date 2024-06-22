package employee

import (
	"errors"
	"fmt"
)

type Employee struct {
	Id       string
	Password string
}

func NewEmployee(id string, password string) (*Employee, error) {
	if len([]byte(id)) < 6 && len([]byte(password)) < 8 {
		return nil, errors.New(fmt.Sprintf("Invalid password (%v) or id (%v) length", password, id))
	}
	return &Employee{Id: id, Password: password}, nil
}

func (employee *Employee) String() string {
	return fmt.Sprintf("ID: %v\nPassword: %v", employee.Id, employee.Password)
}
