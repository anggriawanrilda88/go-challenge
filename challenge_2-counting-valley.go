// Gary is an avid hiker. He tracks his hikes meticulously, paying close attention to small details like topography. During his last hike he took exactly n  steps. For every step he took, he noted if it was an uphill, U, or a downhill ,  D step. Gary's hikes start and end at sea level and each step up or down represents a 1 unit change in altitude. We define the following terms:

// A mountain is a sequence of consecutive steps above sea level, starting with a step up from sea level and ending with a step down to sea level.

// A valley is a sequence of consecutive steps below sea level, starting with a step down from sea level and ending with a step up to sea level.

// Given Gary's sequence of up and down steps during his last hike, find and print the number of valleys he walked through.

// For example, if Gary's path is , he first enters a valley  units deep. Then he climbs out an up onto a mountain  units high. Finally, he returns to sea level and ends his hike.

// Print a single integer that denotes the number of valleys Gary walked through during his hike.

package main

import (
	"fmt"
	"log"
	"math"
)

// Complete the countingValleys function below.
func countingValleys(n int32, s string) int32 {
	startAr := []int32{0}
	var start int32
	var valleyPoint int32 = 1
	for _, elem := range s {
		char := string(elem)
		checkInt := true
		checkInt = math.Signbit(float64(start))
		if char == "D" {
			start = start - 1
		} else {
			start = start + 1
		}

		if start == 0 {
			if checkInt == true {
				valleyPoint = valleyPoint + 1
			}
		}
		startAr = append(startAr, start)
	}
	fmt.Println(startAr)
	valley := valleyPoint - 1
	return valley
}

func main() {
	log.Print("Result: ", countingValleys(1000, "UDUUUDUDDD"))
	// log.Print("Result: ", countingValleys(1000, "UUUUDUUDDDDDUUUDDUUDUDDDUUDUDUDDUDUDDUUUDDDUDDUDDDDDUUUDDUUDDDDDDDDDDDUUDDDDDDUUUUDUDUDDDDUDUDUUUDUDDDDUUDDDDDDDUUUUUUUDDUUUUUDUUUDUDUDDDDUUDDUDDDDDUDUDUDDDDUUDUDUDDUUDDDDDUDDUUUUDDDUUUDUUUUDUUDUDUDUUDUUDDDUDDUDDDDDDDUUDDUDDDDDUDUDDDUUDUDUDDUUUUUUDDUDDDDUDUDDDUDDDUUUUDDDDDDDUUDDDDUDUDUDDUUDDUUDUUDUUDDDUUUDDDDUDDDUDUDUDUUUDUDUDUDUUUUDDDUUDUUUUDUUDDUUDUDUDDUDDDDUUDDDDUDDUUUDUUDUDDDUUUDUUDDDUDUDUUDUUUUUDDDUUUUUDDDUUDDUDDDUUUUUUUUDDDUUDUDUUDUUDUDUUDDDDDUUUUDUDDUUUDUDDDDUUDUUUDUUUDDDDUUDDUUDUUUDDUDDUDDUDDUUDDDUDUUUUUUUDDDUDUDDUUUUUDDUDUDDDDDUDDDUDUDUDDDDDDDDUUDDUDUUUUDUUUDUUDDUUUDDDUUDUUDDUDUDDUUUUUUUDDUDDUDUDDDUUDUUUDDUUDDDDUUUUUDDUUDDDUUUDDDUUDDDDDUDUUUDDUDUDDUDDUUDDDDDUUUUDUDUUUUUUDUUDUUUDDUUUUUUDDDDUDUDUUDUDDDUUDDDUUUUUDDDDDUDUDDDDDUUDDDUDUUUDDDUUUUDDUUUUUUUDUDDUUDDUDDDDDDDDDDDUDUDUUDDDDUDUUUDUDUUDUUUDDDDDUUDUDUUUDDUDUUUUDDUUDDUDDUDUUUDDUDUUDUDUUDDDDDDUUDUDDDUUUUDDDDUUUUUDDUUUDUUDUUDUUUUDDUDDUDDUDUUUDDUDDDDDUUDDUDUUDUUDUUDUDUUDDUUDUDUUUUDUUDDDDUUDDUUDUUUDDUUUDUDDUUUUUUUUUUUUUUUUUUUUUUUU"))
	// log.Print("Result: ", countingValleys(1000, "DDUUDDUDUUUD"))
}
