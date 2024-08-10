package main

func main() {
	if len(os.Args) > 4 {
		return
	}
	a := toNumber(os.Args[1])
	b := toNumber(os.Args[3])
	switch op := os.Args[2]; op {
		case "+":
			
		case "-":
		case "/":
		case "*":
		case "%":
		default:
			return
	}
}
