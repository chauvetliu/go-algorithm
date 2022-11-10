package main

type AllArea struct {
	Province int64
	Area     int64
	City     int64
}

func main() {
	c := make(chan bool)

	go func() {

		for {
			select {
			case <-c:
				return

			default:

			}

		}
	}()

	c <- true
}
