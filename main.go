package main

func main() {
	storage, err := NewStorage("root:1234@tcp(127.0.0.1:3306)/main")
	if err != nil {
		panic(err)
	}
	_ = storage
}
