package workers

// WG is waitgroup
type WG struct {
	main chan func()
	done chan bool
}

// New runs n number of workers
func New(n int) WG {

	wg := WG{
		main: make(chan func()),
		done: make(chan bool),
	}

	procDone := make(chan bool)

	for i := 0; i <= n; i++ {
		go func() {
			for f := range wg.main {
				f()
			}
			procDone <- true
		}()
	}

	go func() {
		for i := 0; i <= n; i++ {
			_ = <-procDone
		}
		wg.done <- true
	}()
	println(len(wg.main))
	return wg
}

// Add adds func f to main channel
func (wg WG) Add(f func()) {
	wg.main <- f
}

// Wait closes main channel and waits until all func are executed
func (wg WG) Wait() {
	close(wg.main)
	_ = <-wg.done
}
