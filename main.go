package main

// func init() {
// 	opts := &mysql.MySQLContainerOptions{
// 		RootPassword: "1234",
// 		Database:     "main",
// 	}
// 	container := mysql.NewMySQLContainer(opts)
// 	container.Run()
// }

func main() {
	storage, err := NewStorage("root:1234@tcp(127.0.0.1:3306)/main")
	if err != nil {
		panic(err)
	}

	as := &apiserver{storage: storage}
	s := as.SetupServer()

	go s.Listen(":4000")
	storage.WatchBookings()
}
