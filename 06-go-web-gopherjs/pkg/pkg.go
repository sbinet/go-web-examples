package pkg

type Data struct {
	X float64
	Y float64
}

func Hello(s string) string {
	switch s {
	case "":
		return "Hello you!"
	default:
		return "Hello " + s + "!"
	}
}
