// Compare copy() and append()
package mylib

import "fmt"

func Linked(s1 []int) {
	s2 := s1  // point to the same memory with s1
	s3 := s1[:]  // point to the same memory with s1
	
	s1[3] = 99

	fmt.Println("Linked():")
	fmt.Println("s1: ", s1, fmt.Sprintf("addr:%p, len:%d, cap:%d", s1, len(s1), cap(s1)))
	fmt.Println("s2: ", s2, fmt.Sprintf("addr:%p, len:%d, cap:%d", s2, len(s2), cap(s2)))
	fmt.Println("s3: ", s3, fmt.Sprintf("addr:%p, len:%d, cap:%d", s3, len(s3), cap(s3)))
}

func Append(s1 []int) {
	s2 := s1
	s3 := s1[:len(s1)-1]

	s1 = append(s1, -10)
	s1[3] = -99

	s3 = append(s3, -20)
	s3[0] = 0


	fmt.Println("NoLinked():")
	fmt.Println("s1: ", s1, fmt.Sprintf("addr:%p, len:%d, cap:%d", s1, len(s1), cap(s1)))
	fmt.Println("s2: ", s2, fmt.Sprintf("addr:%p, len:%d, cap:%d", s2, len(s2), cap(s2)))
	fmt.Println("s3: ", s3, fmt.Sprintf("addr:%p, len:%d, cap:%d", s3, len(s3), cap(s3)))
}

func Copy(s1 []int) {
	s2 := make([]int, len(s1))
	cpoied := copy(s2, s1)

	s1[1]=-1

	fmt.Println("cpoied: ", cpoied)
	fmt.Println("s1: ", s1, fmt.Sprintf("addr:%p, len:%d, cap:%d", s1, len(s1), cap(s1)))
	fmt.Println("s2: ", s2, fmt.Sprintf("addr:%p, len:%d, cap:%d", s2, len(s2), cap(s2)))
}