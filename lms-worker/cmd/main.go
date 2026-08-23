package main

func main() {
	app := NewApp()

	defer func() {
		app.entClient.Close()
	}()

	app.Start()
}
