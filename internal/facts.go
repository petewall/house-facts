package internal

import (
	"encoding/json"
	"fmt"
	"os"
)

type Facts struct {
	CircuitFacts []*CircuitFact `json:"circuitFacts,omitempty"`
	PaintColors  []*PaintColor  `json:"paintColors,omitempty"`
}

func LoadFacts(factsFile string) (*Facts, error) {
	factsFileBytes, err := os.ReadFile(factsFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read facts file: %w", err)
	}

	var facts *Facts
	err = json.Unmarshal(factsFileBytes, &facts)
	if err != nil {
		return nil, fmt.Errorf("failed to parse facts file: %w", err)
	}

	for _, circuitFact := range facts.CircuitFacts {
		circuitFact.CreateMetric()
	}
	for _, paintColor := range facts.PaintColors {
		paintColor.CreateMetric()
	}

	return facts, nil
}
