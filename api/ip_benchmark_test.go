package ip

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/phayes/freeport"
	"github.com/wafer-bw/whatsmyip/spec"
	"google.golang.org/protobuf/proto"
)

var iterations = 1000
var maxConcurrent = 50
var client = http.Client{Timeout: time.Second * 2}

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}

func TestBenchmark(t *testing.T) {
	ports, err := freeport.GetFreePorts(2)
	if err != nil {
		panic(err)
	}
	os.Setenv("HTTP_PORT", fmt.Sprintf("%d", ports[0]))

	go func() {
		main()
	}()

	runTests("protobuf", protoTest)
	runTests("json\t", jsonTest)
}

func runTests(name string, test func(*sync.WaitGroup, *int32, *int32)) {
	bad, good := int32(0), int32(0)
	start := time.Now()
	wg := new(sync.WaitGroup)
	for i := 0; i < iterations; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			test(wg, &good, &bad)
		}()
		if i%maxConcurrent == 0 {
			wg.Wait()
		}
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("%s took %s, (ok: %d, !ok: %d)\n", name, elapsed, good, bad)
}

func protoTest(wg *sync.WaitGroup, good *int32, bad *int32) {
	headers := map[string]string{"Accept": "application/protobuf"}
	url := fmt.Sprintf("http://localhost:%s/", os.Getenv("HTTP_PORT"))
	body, err := getBody(url, headers)
	if err != nil {
		atomic.AddInt32(bad, 1)
		return
	}
	reply := &spec.IPReply{}
	err = proto.Unmarshal(body, reply)
	if err != nil {
		atomic.AddInt32(bad, 1)
		return
	}
	atomic.AddInt32(good, 1)
}

func jsonTest(wg *sync.WaitGroup, good *int32, bad *int32) {
	headers := map[string]string{"Accept": "application/json"}
	url := fmt.Sprintf("http://localhost:%s/", os.Getenv("HTTP_PORT"))
	body, err := getBody(url, headers)
	if err != nil {
		atomic.AddInt32(bad, 1)
		return
	}
	reply := &spec.IPReply{}
	err = json.Unmarshal(body, reply)
	if err != nil {
		atomic.AddInt32(bad, 1)
		return
	}
	atomic.AddInt32(good, 1)
}

func getBody(url string, headers map[string]string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	for key, val := range headers {
		req.Header.Set(key, val)
	}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}
