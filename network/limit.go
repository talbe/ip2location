package network

import (
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/IpLocation/configuration"
	"github.com/IpLocation/network/rate"
)

// Create a custom visitor struct which holds the rate limiter for each
// visitor and the last time that the visitor was seen.
type visitor struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

// Change the the map to hold values of the type visitor.
var visitors = make(map[string]*visitor)
var mu sync.Mutex

// Run a background goroutine to remove old entries from the visitors map.
func init() {
	go cleanupVisitors()
}

func getVisitor(ip string) (*rate.Limiter,error) {
	mu.Lock()
	defer mu.Unlock()

	v, exists := visitors[ip]
	if !exists {

		tokenIncreaseRate, err := configuration.ConfigInstance().TokenIncreaseRate()
		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		bucketSize, err := configuration.ConfigInstance().BucketSize()
		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		limiter := rate.NewLimiter(tokenIncreaseRate, bucketSize)
		visitors[ip] = &visitor{limiter, time.Now()}
		return limiter, nil
	}

	v.lastSeen = time.Now()
	return v.limiter, nil
}

// Every minute check the map for visitors that haven't been seen for
// more than X minutes and delete the entries.
func cleanupVisitors() error {
	visitorMinutesToLive, err := configuration.ConfigInstance().VisitorMinutesToLive()
	if err != nil {
		log.Fatal(err)
		return err
	}

	for {
		time.Sleep(time.Minute)

		mu.Lock()

		for ip, v := range visitors {
			if time.Now().Sub(v.lastSeen) > (time.Duration(visitorMinutesToLive) * time.Minute) {
				delete(visitors, ip)
			}
		}

		mu.Unlock()
	}
}

func limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			log.Println(err.Error())
			setError(&w, 500)
			return
		}

		limiter, err := getVisitor(ip)
		if err != nil {
			log.Println("Failed find the visitor")
			setError(&w, 500)
		}
		if limiter.Allow() == false {
			setError(&w, 429)
			return
		}

		next.ServeHTTP(w, r)
	})
}