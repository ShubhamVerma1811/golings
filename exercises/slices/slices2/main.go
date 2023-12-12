// slices2
// Make me compile!

package main

import "fmt"

func main() {
	names := []string{"John", "Maria", "Carl", "Peter"}

	names = append(names, "He")
	lastTwoNames := names[len(names)-2:] // after figuring out the answer, try with other low/high bounds

	fmt.Println(lastTwoNames)
}
