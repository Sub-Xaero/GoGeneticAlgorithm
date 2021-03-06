package ga

import (
	"errors"
	"fmt"
)

type Rule struct {
	Condition Bitstring
	Output    string
}

type RuleBase []Rule
type RulesMatchFunc func(Rule, Rule) (bool, error)
type EncodeRulesFunc func(RuleBase) (Bitstring, error)
type DecodeRulesFunc func(Bitstring, int, int) (RuleBase, error)

func (r Rule) String() string {
	return fmt.Sprintf("%v %v", r.Condition, r.Output)
}

// SetMutateFunc changes the mutate function to the function specified
func (genA *GeneticAlgorithm) SetRulesMatchFunc(f RulesMatchFunc) {
	genA.RulesMatch = f
}

var DefaultRulesMatchFunc RulesMatchFunc = func(rule1, rule2 Rule) (bool, error) {
	if len(rule1.Condition) != len(rule2.Condition) {
		return false, errors.New("conditions are not same length")
	}
	conditionMatches := true
	for i := range rule1.Condition {
		if string(rule1.Condition[i]) == "#" || string(rule2.Condition[i]) == "#" {
			continue
		} else if string(rule1.Condition[i]) != string(rule2.Condition[i]) {
			conditionMatches = false
			break
		}
	}
	outputMatches := rule1.Output == rule2.Output
	return outputMatches && conditionMatches, nil
}

// SetMutateFunc changes the mutate function to the function specified
func (genA *GeneticAlgorithm) SetDecodeRulesFunc(f DecodeRulesFunc) {
	genA.DecodeRules = f
}

var DefaultDecodeRulesFunc DecodeRulesFunc = func(sequence Bitstring, conditionLength, ruleLength int) (RuleBase, error) {
	NewRuleBase := make([]Rule, 0)
	for i := 0; i < len(sequence); i += ruleLength {
		condition := make(Bitstring, len(sequence[i:i+conditionLength]))
		copy(condition, sequence[i:i+conditionLength])
		rule := Rule{condition, sequence[i+conditionLength]}
		NewRuleBase = append(NewRuleBase, rule)
	}
	return NewRuleBase, nil
}

// SetMutateFunc changes the mutate function to the function specified
func (genA *GeneticAlgorithm) SetEncodeRulesFunc(f EncodeRulesFunc) {
	genA.EncodeRules = f
}

var DefaultEncodeRulesFunc = func(paramRuleBase RuleBase) (Bitstring, error) {
	var sequence Bitstring
	for _, rule := range paramRuleBase {
		for _, condition := range rule.Condition {
			sequence = append(sequence, string(condition))
		}
		for _, output := range rule.Output {
			sequence = append(sequence, string(output))
		}
	}
	return sequence, nil
}
