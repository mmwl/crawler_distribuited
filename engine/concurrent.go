package engine

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan interface{}
}

type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkerChan(chan Request)
}

func (c *ConcurrentEngine) Run(seeds ...Request) {

	in := make(chan Request)
	out := make(chan ParseResult)

	c.Scheduler.ConfigureMasterWorkerChan(in)

	for i := 0; i < c.WorkerCount; i++ {
		createWorker(in, out)
	}

	for _, v := range seeds {
		c.Scheduler.Submit(v)
	}

	itemCount := 1

	for {
		result := <-out
		itemCount++
		go func() {
			for _, v := range result.Items {
				//log.Printf("Got item %d : %v\n",itemCount, v)
				//save item
				go func() {
					c.ItemChan <- v
				}()
			}
		}()
		for _, v := range result.Requests {
			c.Scheduler.Submit(v)
		}
	}

}

func createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
