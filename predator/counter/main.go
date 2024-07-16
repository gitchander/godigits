package main

import (
	"log"
)

func main() {
	//checkError(testPredatorCountdown())
	checkError(drawPredatorCountdown())
	//checkError(drawPredatorCountdownOneSegment())
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
