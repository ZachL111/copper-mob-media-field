package policy

import "testing"

func TestFixtureDecisions(t *testing.T) {
	signal := Signal{Demand: 56, Capacity: 96, Latency: 18, Risk: 6, Weight: 13}
	if got := Score(signal); got != 188 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "accept" { t.Fatalf("decision = %s", got) }
	signal := Signal{Demand: 59, Capacity: 86, Latency: 13, Risk: 8, Weight: 5}
	if got := Score(signal); got != 150 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "review" { t.Fatalf("decision = %s", got) }
	signal := Signal{Demand: 100, Capacity: 77, Latency: 21, Risk: 11, Weight: 13}
	if got := Score(signal); got != 221 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "accept" { t.Fatalf("decision = %s", got) }
}
