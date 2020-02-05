package rate

import (
	"time"
	"sync/atomic"
)

type Limiter struct {
	bucketSize uint32
	numberOfTokens uint32
	tokenIncreaseRate uint32
}

func NewLimiter(tokenIncreaseRate uint32, bucketSize uint32) *Limiter {

	limiter := &Limiter{bucketSize: bucketSize, tokenIncreaseRate: tokenIncreaseRate, numberOfTokens: bucketSize}
    go limiter.addTokens()
	return limiter
}

func (this *Limiter) Allow() bool {
	actualNumberOfTokens := atomic.LoadUint32(&this.numberOfTokens)

	if actualNumberOfTokens == 0{
		return false
	}

	atomic.StoreUint32(&this.numberOfTokens, this.numberOfTokens - 1)
	return true
}

func (this *Limiter) addTokens() {

	for {
		time.Sleep(time.Second)

		increaseValue := this.tokenIncreaseRate
		for (this.numberOfTokens < this.bucketSize) && (increaseValue != 0) {
			atomic.AddUint32(&this.numberOfTokens, 1)
			increaseValue--
		}
	}
}