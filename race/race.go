package race

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

type Race struct {
	time     int
	distance int
}

func GetRaces(input string) []Race {
	races := []Race{}

	lines := strings.Split(input, "\n")

	times := strings.Fields(lines[0])[1:]

	for _, t := range times {
		tt, _ := strconv.Atoi(t)
		r := Race{time: tt}
		races = append(races, r)
	}

	distances := strings.Fields(lines[1])[1:]

	for i, d := range distances {
		dd, _ := strconv.Atoi(d)
		races[i].distance = dd
	}

	return races
}

func GetRacesPartTwo(input string) []Race {
	races := []Race{}

	lines := strings.Split(input, "\n")

	timeSlice := strings.Fields(lines[0])[1:]
	timeStr := strings.Join(timeSlice, "")

	time, _ := strconv.Atoi(timeStr)

	distanceSlice := strings.Fields(lines[1])[1:]
	distanceStr := strings.Join(distanceSlice, "")

	distance, _ := strconv.Atoi(distanceStr)

	races = append(races, Race{
		time:     time,
		distance: distance,
	})

	return races
}

func (r *Race) CalculateWaysToWin() int {
	w := 0
	for i := 1; i < r.time; i++ {
		if i*(r.time-i) > r.distance {
			w++
		}
	}
	return w
}

func PartOne(races []Race) int {
	ch := make(chan int)
	var wg sync.WaitGroup

	fmt.Println("calculating ways to win")
	for _, r := range races {
		wg.Add(1)

		go func(r Race) {
			defer wg.Done()
			numW := r.CalculateWaysToWin()
			ch <- numW
		}(r)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	final := 1

	for s := range ch {
		final *= s
	}

	return final
}
