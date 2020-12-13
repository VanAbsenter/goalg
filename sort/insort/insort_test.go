// Testing of insertion sort
package insort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestIntInsort(t *testing.T) {
	slice := make([]int, 100, 100)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	fmt.Println("===Unsorted int slice===\n", slice)
	IntInsort(slice)
	fmt.Println("===Sorted int slice===\n", slice)
}

func TestFloat64Insort(t *testing.T) {
	slice := make([]float64, 100, 100)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		slice[i] = (rand.Float64() * 9.2) - rand.Float64()*3.3
	}
	fmt.Println("===Unsorted float64 slice===\n", slice)
	Float64Insort(slice)
	fmt.Println("===Sorted float64 slice===\n", slice)
}
