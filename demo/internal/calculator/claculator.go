package calculator

import (
	"errors"
)

type calculator struct {}

func New() *calculator  {
	return &calculator{}
}

func (c *calculator) Add(a , b int) int{
	return  a+b
}

func (c *calculator) Sub(a , b int) int{
	return  a-b
}

func (c *calculator) Mul(a , b int) int{
	return  a*b
}

func (c *calculator) Div(a , b float32) (float32,error){

	if b==0 {
		return 0, errors.New("Divison by Zero")
	}

	return  a/b, nil
}