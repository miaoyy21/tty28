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
		log.Printf("使用随机种子【%d】\n", i)
		rand.Seed(i)
		tbl := rand.Perm(math.MaxInt32)

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

			if isMatch && len(tbl)-k > 3 {
				log.Printf("\t随机种子【%d】，相同的起始下标索引为【%d】，后3位【%d,%d,%d】\n", i, k, tbl[k+len(nums)], tbl[k+len(nums)+1], tbl[k+len(nums)+2])
			}
		}
	}
}
