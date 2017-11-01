package number

import (
	"testing"
	"fmt"
)

func TestSum(t *testing.T) {
	m1:=RandMatrix(5,2,10,50)
	m2:=RandMatrix(5,2,0,30)
	m3:=RandMatrix(5,2,0,30)
	fmt.Println(m1,m2,m3,Sum(m1,m2,m3))
}

func TestMulC(t *testing.T) {
	m1:=RandMatrix(5,2,10,50)
	fmt.Println(m1,MulC(m1,2))
}
func TestMul(t *testing.T) {
	m1:=RandMatrix(3,1,0,10)
	m2:=RandMatrix(1,3,0,10)
	fmt.Println(m1,m2,Mul(m1,m2))
}