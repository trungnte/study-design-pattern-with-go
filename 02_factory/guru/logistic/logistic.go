package main

type Logistic struct {
	kind    string
	vehicle Transport
}

func (l *Logistic) planDelivery() {
	if l.kind == "sea" {
		l.vehicle = &SeaLogistic{}
	} else {
		l.vehicle = &RoadLogistic{}
	}
}
