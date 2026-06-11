package main

import "time"

type City struct {
	name string
	long float32
	lat  float32
}

type Temp struct {
	city *City
	Date time.Time
	temp float32
}

