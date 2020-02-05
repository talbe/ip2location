package rate

import (
	"testing"
	"time"
)

func sanity(t *testing.T){
	limiter := NewLimiter(1, 3)

	allowed := limiter.Allow()
	allowed = limiter.Allow() && allowed
	allowed = limiter.Allow() && allowed

	if ! allowed {
		t.Errorf("Not allowd after just 3 calls")
	}

	allowed = limiter.Allow() && allowed

	if allowed {
		t.Errorf("After 4, it is should not be allowed")
	}
}

func TestAllowedSanity(t *testing.T) {
	sanity(t)
}

func TestAllowedComplicated(t *testing.T) {

	sanity(t)

	// Sleep for 3 seconds and the sanity should work again
	time.Sleep(3000 * time.Millisecond)

	sanity(t)
}

func TestAllowedMultithreaded(t *testing.T) {
	limiter := NewLimiter(3, 1)

	go func(){
		limiter.Allow()
		time.Sleep(200 * time.Millisecond)
		limiter.Allow()
	}()

	time.Sleep(200 * time.Millisecond)
	limiter.Allow()

	// Sleep for 3 seconds and the sanity should work again
	time.Sleep(500 * time.Millisecond)

	if limiter.Allow() {
		t.Errorf("After 4, it is should not be allowed")
	}

}