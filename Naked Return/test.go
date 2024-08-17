package main

import "fmt"

func HdrCol(hdr []string) (IdxtoCol map[int]string) {
	IdxtoCol = make(map[int]string)

	for idx, item := range hdr {
		switch item {
		case "employee":
			IdxtoCol[idx] = item
		case "empid":
			IdxtoCol[idx] = item
		case "hourly rate":
			IdxtoCol[idx] = item
		}
	}

	return IdxtoCol
}

func Naked_HdrCol(hdr []string) (IdxtoCol map[int]string) {
	IdxtoCol = make(map[int]string)

	for idx, item := range hdr {
		switch item {
		case "address":
			IdxtoCol[idx] = "Get Address"
		case "employee":
			IdxtoCol[idx] = "Get Employee"
		case "hours worked":
			IdxtoCol[idx] = "Get Hours Worked"
		}
	}

	return
}

func main() {
	hdr := []string{"empid", "employee", "address", "hours worked",
		"hourly rate", "manager"}
	Idx2Col := HdrCol(hdr)
	Naked_Idx2Col := Naked_HdrCol(hdr)

	fmt.Println(Idx2Col)
	fmt.Println()
	fmt.Println(Naked_Idx2Col)
}
