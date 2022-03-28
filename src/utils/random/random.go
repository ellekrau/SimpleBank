package random

import (
	"github.com/ellekrau/SimpleBank/utils/enum/currency"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvxz"

func init() {
	rand.Seed(time.Now().Unix())
}

// Int generates a random int64 between min and max
func Int(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// String generates a random string between min and max
func String(n int) string {

	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// Owner generates a random owner name
func Owner() string {
	return String(6)
}

// Amount generates a random amount value
func Amount() int64 {
	return Int(1, 10000)
}

// Currency generates a currency amount value
func Currency() string {
	currencies := []string{currency.Real, currency.USD}
	return currencies[rand.Intn(len(currencies))]
}
