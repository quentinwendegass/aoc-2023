package main

import (
	"aoc2023/utils"
	"slices"
	"strconv"
	"strings"
)

type Range struct {
	start int
	end   int
}

type Condition struct {
	variable byte
	operator byte
	value    int
	outcome  string
}

type Workflow struct {
	conditions     []Condition
	defaultOutcome string
}

type Rating map[byte]int
type RangeRating map[byte]Range
type Workflows map[string]Workflow

func main() {
	utils.WithAOC(partOne, partTwo, utils.DefaultDataLoader)
}

func partOne(data []byte) int {
	ratings, workflows := convertInput(data)

	total := 0

	for _, rating := range ratings {
		if isRatingAccepted(rating, workflows) {
			for _, r := range rating {
				total += r
			}
		}
	}

	return total
}

func partOneWithRanges(data []byte) int {
	ratings, workflows := convertInput(data)

	total := 0

	for _, rating := range ratings {

		allPossibleRanges := findPossibleRanges("in", convertRatingToRange(rating), workflows)

		if len(allPossibleRanges) > 0 {
			for _, val := range rating {
				total += val
			}
		}
	}

	return total
}

func partTwo(data []byte) int {
	_, workflows := convertInput(data)
	ratingRanges := map[byte]Range{'x': {start: 1, end: 4000}, 'm': {start: 1, end: 4000}, 'a': {start: 1, end: 4000}, 's': {start: 1, end: 4000}}

	allPossibleRanges := findPossibleRanges("in", ratingRanges, workflows)

	allCombinations := 0
	for _, y := range allPossibleRanges {
		total := 1
		for _, v := range y {
			total *= v.end - v.start + 1
		}
		allCombinations += total
	}

	return allCombinations
}

func isRatingAccepted(rating Rating, workflows Workflows) bool {
	currentWorkflow := "in"

	for {
		if currentWorkflow == "R" {
			return false
		} else if currentWorkflow == "A" {
			return true
		}

		found := false
		for _, conditions := range workflows[currentWorkflow].conditions {
			isConditionTrue := checkCondition(rating, conditions)

			if isConditionTrue {
				currentWorkflow = conditions.outcome
				found = true
				break
			}
		}

		if !found {
			currentWorkflow = workflows[currentWorkflow].defaultOutcome
		}
	}
}

func checkCondition(rating Rating, condition Condition) bool {
	ratingValue := rating[condition.variable]
	return (condition.operator == '<' && ratingValue < condition.value) || (condition.operator == '>' && ratingValue > condition.value)
}

func findPossibleRanges(workflowKey string, rangeRating map[byte]Range, workflows Workflows) []RangeRating {
	if workflowKey == "A" {
		return []RangeRating{rangeRating}
	} else if workflowKey == "R" {
		return []RangeRating{}
	}

	workflow := workflows[workflowKey]

	successRangeRating := rangeRating
	failedRangeRating := rangeRating

	allPossibleRatingRanges := make([]RangeRating, 0)
	allPossibleRatingRangeKeys := make([]string, 0)

	for _, condition := range workflow.conditions {
		successRangeRating, failedRangeRating = findConditionRanges(failedRangeRating, condition)
		allPossibleRatingRanges = append(allPossibleRatingRanges, successRangeRating)
		allPossibleRatingRangeKeys = append(allPossibleRatingRangeKeys, condition.outcome)
	}

	allPossibleRatingRanges = append(allPossibleRatingRanges, failedRangeRating)
	allPossibleRatingRangeKeys = append(allPossibleRatingRangeKeys, workflow.defaultOutcome)

	resultRatingRanges := make([]RangeRating, 0)
	for i, possibleRatingRange := range allPossibleRatingRanges {
		possibleRatingRange := findPossibleRanges(allPossibleRatingRangeKeys[i], possibleRatingRange, workflows)
		if len(possibleRatingRange) > 0 && !containsEmptyRange(possibleRatingRange) {
			resultRatingRanges = append(resultRatingRanges, possibleRatingRange...)
		}
	}

	return resultRatingRanges
}

func findConditionRanges(rangeRating RangeRating, condition Condition) (success, failed RangeRating) {
	curRange := rangeRating[condition.variable]
	rangeRatingSuccess := make(RangeRating)
	rangeRatingFailed := make(RangeRating)

	for k, v := range rangeRating {
		rangeRatingSuccess[k] = v
		rangeRatingFailed[k] = v
	}

	if condition.operator == '<' {
		if curRange.end < condition.value {
			rangeRatingFailed[condition.variable] = Range{start: 0, end: 0}
		} else if curRange.start > condition.value {
			rangeRatingSuccess[condition.variable] = Range{start: 0, end: 0}
		} else {
			rangeRatingSuccess[condition.variable] = Range{start: curRange.start, end: condition.value - 1}
			rangeRatingFailed[condition.variable] = Range{start: condition.value, end: curRange.end}
		}
	}

	if condition.operator == '>' {
		if curRange.start > condition.value {
			rangeRatingFailed[condition.variable] = Range{start: 0, end: 0}
		} else if curRange.end < condition.value {
			rangeRatingSuccess[condition.variable] = Range{start: 0, end: 0}
		} else {
			rangeRatingSuccess[condition.variable] = Range{start: condition.value + 1, end: curRange.end}
			rangeRatingFailed[condition.variable] = Range{start: curRange.start, end: condition.value}
		}
	}

	for k := range rangeRating {
		if rangeRatingSuccess[k].start > rangeRatingSuccess[k].end {
			rangeRatingSuccess[k] = Range{start: 0, end: 0}
		}
		if rangeRatingFailed[k].start > rangeRatingFailed[k].end {
			rangeRatingFailed[k] = Range{start: 0, end: 0}
		}
	}

	return rangeRatingSuccess, rangeRatingFailed
}

func containsEmptyRange(rangeMaps []RangeRating) bool {
	emptyRange := Range{}
	return slices.ContainsFunc(rangeMaps, func(r RangeRating) bool {
		for _, v := range r {
			if v == emptyRange {
				return true
			}
		}
		return false
	})
}

func convertRatingToRange(rating Rating) RangeRating {
	rangeRating := make(RangeRating)

	for k, v := range rating {
		rangeRating[k] = Range{start: v, end: v}
	}

	return rangeRating
}

func convertInput(data []byte) ([]Rating, Workflows) {
	splitInput := strings.Split(string(data), "\n\n")
	workflowLines := strings.Split(splitInput[0], "\n")

	workflows := make(Workflows)

	for _, line := range workflowLines {
		lineSplit := strings.Split(line, "{")
		workflowKey := lineSplit[0]
		workflowFull := lineSplit[1][0 : len(lineSplit[1])-1]
		rawWorkflows := strings.Split(workflowFull, ",")

		conditions := make([]Condition, len(rawWorkflows)-1)

		for i := 0; i < len(rawWorkflows)-1; i++ {
			rawWorkflowSplit := strings.Split(rawWorkflows[i], ":")
			value, err := strconv.Atoi(rawWorkflowSplit[0][2:len(rawWorkflowSplit[0])])
			if err != nil {
				panic(err)
			}

			conditions[i] = Condition{variable: rawWorkflowSplit[0][0], operator: rawWorkflowSplit[0][1], value: value, outcome: rawWorkflowSplit[1]}
		}

		workflows[workflowKey] = Workflow{conditions: conditions, defaultOutcome: rawWorkflows[len(rawWorkflows)-1]}
	}

	ratingLines := strings.Split(splitInput[1], "\n")

	ratings := make([]Rating, len(ratingLines))
	for i, ratingLine := range ratingLines {
		rawRatings := strings.Split(ratingLine[1:len(ratingLine)-1], ",")
		ratingMap := make(Rating)

		for _, rawRating := range rawRatings {
			value, err := strconv.Atoi(rawRating[2:])
			if err != nil {
				panic(err)
			}

			ratingMap[rawRating[0]] = value
		}

		ratings[i] = ratingMap
	}

	return ratings, workflows
}
