package main

import (
	"fmt"
	"github.com/hpifu/go-kit/hflag"
	"math/rand"
	"strconv"
	"time"
)

func test() {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(2)
	fmt.Println(i)
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			test()
		}
	}()

	ticker := time.NewTicker(1 * time.Second)
	go func(ticker *time.Ticker) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println(r)
				test()
			}
		}()
		defer ticker.Stop()
		for range ticker.C {
			rand.Seed(time.Now().UnixNano())
			i := rand.Intn(2)
			fmt.Println(i)
		}
	}(ticker)
}

func lts(nums []int) {
	bis := [10]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	for i, num1 := range nums {
		for j, _ := range bis {
			if nums[j] < num1 {
				bis[i] = max(bis[i], bis[j]+1)
			}
		}
	}
}
func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

var m map[string]int = make(map[string]int)

func MaxString(s1, s2 string, i, j int) int {
	if i == len(s1) {
		return 0
	}
	if j == len(s2) {
		return 0
	}
	if s1[i] == s2[j] {
		key := strconv.Itoa(i) + ":" + strconv.Itoa(j)
		if _, ok := m[key]; !ok {
			m[key] = 1 + MaxString(s1, s2, i+1, j+1)
		}
		return m[key]
	} else {
		key1 := strconv.Itoa(i+1) + ":" + strconv.Itoa(j)
		if _, ok := m[key1]; !ok {
			m[key1] = MaxString(s1, s2, i+1, j)
		}
		key2 := strconv.Itoa(i) + ":" + strconv.Itoa(j+1)
		if _, ok := m[key2]; !ok {
			m[key2] = MaxString(s1, s2, i, j+1)
		}
		return max(m[key1], m[key2])
	}
}

func sub(s1, s2 string, i, j int) int {
	if i == len(s1) && j < len(s2) {
		return len(s2) - j - 1
	} else if j == len(s2) && i < len(s1) {
		return len(s1) - i - 1
	} else if j == len(s2) && i == len(s1) {
		return 0
	}
	if s1[i] == s2[j] {
		return sub(s1, s2, i+1, j+1)
	} else {
		return min(sub(s1, s2, i+1, j)+1, sub(s1, s2, i+1, j+1)+1, sub(s1, s2, i, j+1)+1)
	}

}

func min(i, j, k int) int {
	if j > k && i > k {
		return k
	} else if i > j && k > j {
		return j
	} else {
		return i
	}
}

func main() {
	hflag.AddFlag("project", "create project", hflag.Shorthand("p"), hflag.Type("string"))
	hflag.AddFlag("module", "create module", hflag.Shorthand("m"))
	hflag.AddFlag("o", "int slice flag", hflag.Type("[]int"), hflag.DefaultValue("1,2,3"))
	hflag.AddFlag("ip", "ip flag", hflag.Type("ip"))
	hflag.AddFlag("time", "time flag", hflag.Type("time"), hflag.DefaultValue("2019-11-27"))
	hflag.AddPosFlag("pos", "pos flag")
	if err := hflag.Parse(); err != nil {
		panic(err)
	}

	fmt.Println("int =>", hflag.GetInt("i"))
	fmt.Println("str =>", hflag.GetString("s"))
	fmt.Println("int-slice =>", hflag.GetIntSlice("int-slice"))
	fmt.Println("ip =>", hflag.GetIP("ip"))
	fmt.Println("time =>", hflag.GetTime("time"))
	fmt.Println("pos =>", hflag.GetString("pos"))

}
