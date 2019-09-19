package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	timeFmt = "15:04:05"
)

func bestTime(now time.Time, later time.Time) time.Time {
	now = now.Local() // use local time to make things make sense
	nowh, nowm, nows := now.Clock()
	laterh, laterm, laters := later.Clock()
	add := false
	if nowh > laterh {
		add = true
	} else if (nowh == laterh) && (nowm > laterm) {
		add = true
	} else if (nowh == laterh) && (nowm == laterm) && (nows >= laters) {
		// >= in the case we're on the exact second; add a day because the alarm should have gone off by now otherwise!
		add = true
	}
	if add {
		now = now.AddDate(0, 0, 1)
	}
	return time.Date(now.Year(), now.Month(), now.Day(),
		laterh, laterm, laters, 0,
		now.Location())
}

func timeRand(duration int) int {
	rand.Seed(time.Now().UnixNano())
	max, min := 110, 91
	sum := 0
	for sum < (duration - 110) {
		sum += (rand.Intn(max-min) + min)
	}
	test := rand.Intn(max-min) + min
	if sum+test < duration {
		return sum + test
	}
	if duration-sum >= 91 && duration-sum <= 110 {
		return sum + (duration - sum)
	}
	return sum
}

func main() {
	fmt.Println("-----------------------------------------")
	fmt.Print("Please Enter Sleep Time: ")
	var x, y string
	fmt.Scan(&y)
	fmt.Print("Please Enter Alarm Time: ")
	fmt.Scan(&x)
	alaramTime, err := time.Parse(timeFmt, x)
	if err != nil {
		fmt.Print("Error: ", err)
		return
	}
	alaramTime2, err := time.Parse(timeFmt, y)
	if err != nil {
		fmt.Print("Error: ", err)
		return
	}
	now := time.Now()
	fmt.Println("Now Time: ", now.Format("15:04:05"))
	sTime := bestTime(now, alaramTime2)
	fmt.Println("Time For Start Sleeping : ", sTime.Format("15:04:05"))
	later := bestTime(sTime, alaramTime)
	fmt.Println("Alarm Time: ", later.Format("15:04:05"))
	duration := later.Sub(sTime)
	fmt.Printf("Duration Time: %.0f Minutes\n", duration.Minutes())
	Dut := int(duration.Minutes())
	addTime := timeRand(Dut)
	fmt.Println("Add Time: ", addTime)
	after := sTime.Add(time.Minute * time.Duration(addTime))
	fmt.Println("Awake Up Time: ", after.Format("15:04:05"))
	fmt.Println("-----------------------------------------")

}
