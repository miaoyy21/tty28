package hdo

import (
	"log"
	"math"
	"math/rand"
	"testing"
)

func TestDo(t *testing.T) {

	nums := []int{8, 5, 8, 3}

	for i := int64(0); i < math.MaxInt64; i++ {
		tbl := make([]int, len(nums)+3)
		for p := 0; p < len(nums)+3; p++ {
			tbl = append(tbl, rand.Intn(10))
		}

		for k := 0; k < len(tbl); k++ {
			if k+len(nums) > len(tbl) {
				continue
			}

			isMatch := true
			for x, n := range nums {
				if n != tbl[k+x] {
					isMatch = false
					break
				}
			}

			if isMatch && k+len(nums)+2 < len(tbl)-1 {
				log.Printf("\t相同的起始下标索引为【%d】，后3位【%d,%d,%d】\n", k, tbl[k+len(nums)], tbl[k+len(nums)+1], tbl[k+len(nums)+2])
			}
		}
	}
}
