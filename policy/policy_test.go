package policy

import "testing"

func TestFixtureDecisions(t *testing.T) {
	tests := []struct {
		name         string
		signal       Signal
		wantScore    int
		wantDecision string
	}{
		{name: "case_1", signal: Signal{Demand: 56, Capacity: 96, Latency: 18, Risk: 6, Weight: 13}, wantScore: 188, wantDecision: "accept"},
		{name: "case_2", signal: Signal{Demand: 59, Capacity: 86, Latency: 13, Risk: 8, Weight: 5}, wantScore: 150, wantDecision: "review"},
		{name: "case_3", signal: Signal{Demand: 100, Capacity: 77, Latency: 21, Risk: 11, Weight: 13}, wantScore: 221, wantDecision: "accept"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := Score(tc.signal); got != tc.wantScore {
				t.Fatalf("score = %d, want %d", got, tc.wantScore)
			}
			if got := Classify(tc.signal); got != tc.wantDecision {
				t.Fatalf("decision = %s, want %s", got, tc.wantDecision)
			}
		})
	}
}
