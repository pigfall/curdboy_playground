package api_inquisitor

import (
	"fmt"
	"github.com/pigfall/tzzGoUtil/uuid"
	"math/rand"
	"time"
)

func randomStringValue() string {
	uid, err := uuid.GenUUID()
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("auto_%s", uid)
}

func randomTimeValue() time.Time {
	return time.Now()
}
func randomIntValue() int {
	return int(rand.Int31n(1000))
}

func randomFloat64Value() float64 {
	return float64(rand.Float32()) * float64(rand.Int31n(1000))
}

func randomBoolValue() bool {
	v := int(rand.Int31n(1000))
	return v > 500
}
