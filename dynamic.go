import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	nodes     = 5
	records   = 1000
	requests  = 200
)

type Request struct {
	key  int
	hops int
	time time.Duration
}

type Monitor struct {
	hopStats []int
}

func newMonitor() *Monitor {
	return &Monitor{hopStats: make([]int, nodes)}
}

func (m *Monitor) update(node int, hops int) {
	m.hopStats[node] += hops
}

func (m *Monitor) bestNode() int {
	min := m.hopStats[0]
	idx := 0
	for i := 1; i < len(m.hopStats); i++ {
		if m.hopStats[i] < min {
			min = m.hopStats[i]
			idx = i
		}
	}
	return idx
}

var database = make([]map[int]int, nodes)
var monitor = newMonitor()

func presentation(r Request) Request {
	time.Sleep(time.Millisecond * 1)
	return business(r)
}

func business(r Request) Request {
	time.Sleep(time.Millisecond * 2)
	return persistence(r)
}

func persistence(r Request) Request {
	node := monitor.bestNode()
	return router(r, node)
}

func router(r Request, node int) Request {
	hops := rand.Intn(2) + 1
	start := time.Now()
	time.Sleep(time.Millisecond * time.Duration(hops*2))
	_ = database[node][r.key]
	r.hops = hops
	r.time = time.Since(start)
	monitor.update(node, hops)
	return r
}

func worker(wg *sync.WaitGroup, ch chan Request) {
	for i := 0; i < requests; i++ {
		req := Request{key: rand.Intn(records)}
		res := presentation(req)
		ch <- res
	}
	wg.Done()
}

func main() {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < nodes; i++ {
		database[i] = make(map[int]int)
		for j := 0; j < records/nodes; j++ {
			database[i][j] = j
		}
	}

	var wg sync.WaitGroup
	results := make(chan Request, requests*4)

	start := time.Now()

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go worker(&wg, results)
	}

	wg.Wait()
	close(results)

	totalHops := 0
	totalTime := time.Duration(0)
	count := 0

	for r := range results {
		totalHops += r.hops
		totalTime += r.time
		count++
	}

	fmt.Println("Requests:", count)
	fmt.Println("Average Hops:", float64(totalHops)/float64(count))
	fmt.Println("Average Time(ms):", (totalTime/time.Duration(count)).Milliseconds())
	fmt.Println("Total Runtime(ms):", time.Since(start).Milliseconds())
}
