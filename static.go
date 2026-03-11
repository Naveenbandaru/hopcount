import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Request struct {
	id   int
	key  int
	time time.Duration
}

const (
	nodes   = 5
	records = 1000
)

var database = make([]map[int]int, nodes)

func presentation(r Request) Request {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(3)+1))
	return business(r)
}

func business(r Request) Request {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(4)+2))
	return persistence(r)
}

func persistence(r Request) Request {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(3)+1))
	return databaseLayer(r)
}

func databaseLayer(r Request) Request {
	start := time.Now()
	node := r.key % nodes
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(6)+3))
	_ = database[node][r.key]
	r.time = time.Since(start)
	return r
}

func worker(id int, wg *sync.WaitGroup, results chan time.Duration) {
	for i := 0; i < 50; i++ {
		req := Request{
			id:  id*100 + i,
			key: rand.Intn(records),
		}
		res := presentation(req)
		results <- res.time
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
	results := make(chan time.Duration, 500)

	start := time.Now()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(i, &wg, results)
	}

	wg.Wait()
	close(results)

	var total time.Duration
	count := 0

	for r := range results {
		total += r
		count++
	}

	fmt.Println("Requests:", count)
	fmt.Println("Average DB Access Time(ms):", (total/time.Duration(count)).Milliseconds())
	fmt.Println("Total Runtime(ms):", time.Since(start).Milliseconds())
}

