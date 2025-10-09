package main

import "fmt"

func main() {
	var s = []string{"1", "2", "3"}
	fmt.Printf("Slice address inside main() before modifySlice():\t%p\n", &s)
	modifySlice(s)
	fmt.Printf("Slice address inside main() after modifySlice():\t%p\n", &s)
	fmt.Printf("Slice inside main():\t%v\n", s)
}

func modifySlice(i []string) {
	fmt.Printf("Slice address inside modifySlice() before append():\t%p\n", &i)
	fmt.Printf("Before modify:\tlen=%d cap=%d addr=%p\n", len(i), cap(i), &i[0])
	i[0] = "3"
	fmt.Printf("Before append:\tlen=%d cap=%d addr=%p\n", len(i), cap(i), &i[0])
	i = append(i, "4")
	fmt.Printf("After append:\tlen=%d cap=%d addr=%p\n", len(i), cap(i), &i[0])
	i[1] = "5"
	i = append(i, "6")
	fmt.Printf("Slice inside modifySlice():\t%v\n", i)
	fmt.Printf("Slice address inside modifySlice() after append():\t%p\n", &i)
}
