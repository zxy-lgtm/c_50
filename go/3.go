package main

func main() {

	i := 0
	for i < 5 {
		if i < 3 {
			n := 5 - i*2
			k := (5 - n) / 2
			for ; k > 0; k-- {
				print(" ")
			}
			for ; n > 0; n-- {
				print("*")
			}
		} else {
			n := i*2 - 3
			k := (5 - n) / 2
			for ; k > 0; k-- {
				print(" ")
			}
			for ; n > 0; n-- {
				print("*")
			}
		}
		print("\n")
		i++
	}
}
