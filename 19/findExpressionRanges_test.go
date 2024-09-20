package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestLessThenAllInRange(t *testing.T) {
	success, failed := findConditionRanges(RangeRating{'a': {start: 1, end: 10}}, Condition{variable: 'a', operator: '<', value: 100})

	assert.Equal(t, success, RangeRating{'a': {start: 1, end: 10}})
	assert.Equal(t, failed, RangeRating{'a': {start: 0, end: 0}})
}

func TestLessThenAllOutOfRange(t *testing.T) {
	success, failed := findConditionRanges(RangeRating{'a': {start: 10, end: 20}}, Condition{variable: 'a', operator: '<', value: 5})

	assert.Equal(t, success, RangeRating{'a': {start: 0, end: 0}})
	assert.Equal(t, failed, RangeRating{'a': {start: 10, end: 20}})
}

func TestLessThenContainsRange(t *testing.T) {
	success, failed := findConditionRanges(RangeRating{'a': {start: 10, end: 40}}, Condition{variable: 'a', operator: '<', value: 20})

	assert.Equal(t, success, RangeRating{'a': {start: 10, end: 19}})
	assert.Equal(t, failed, RangeRating{'a': {start: 20, end: 40}})
}

func TestGreaterThenAllInRange(t *testing.T) {
	success, failed := findConditionRanges(RangeRating{'a': {start: 200, end: 300}}, Condition{variable: 'a', operator: '>', value: 100})

	assert.Equal(t, success, RangeRating{'a': {start: 200, end: 300}})
	assert.Equal(t, failed, RangeRating{'a': {start: 0, end: 0}})
}

func TestGreaterThenAllOutOfRange(t *testing.T) {
	success, failed := findConditionRanges(RangeRating{'a': {start: 100, end: 200}}, Condition{variable: 'a', operator: '>', value: 500})

	assert.Equal(t, success, RangeRating{'a': {start: 0, end: 0}})
	assert.Equal(t, failed, RangeRating{'a': {start: 100, end: 200}})
}

func TestGreaterThenContainsRange(t *testing.T) {
	success, failed := findConditionRanges(RangeRating{'a': {start: 10, end: 40}}, Condition{variable: 'a', operator: '>', value: 20})

	assert.Equal(t, success, RangeRating{'a': {start: 21, end: 40}})
	assert.Equal(t, failed, RangeRating{'a': {start: 10, end: 20}})
}
