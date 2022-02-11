package timeinterval

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestEqualMany(t *testing.T) {
	testCases := []struct {
		name            string
		newIntervalMany SpanMany
		inputInterval   Span
		offset          time.Duration
		excepted        bool
	}{
		{
			name: "not_equal_slightly",
			newIntervalMany: NewMany(
				Span{
					time.Date(2020, 10, 12, 17, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 18, 0, 0, 12, time.UTC)},
				Span{
					time.Date(2020, 10, 12, 17, 0, 0, 1, time.UTC),
					time.Date(2020, 10, 12, 18, 0, 0, 11, time.UTC)},
				Span{
					time.Date(2020, 10, 12, 17, 0, 0, 1, time.UTC),
					time.Date(2020, 10, 12, 18, 0, 0, 12, time.UTC)},
			),
			inputInterval: Span{
				time.Date(2020, 10, 12, 17, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 12, 18, 0, 0, 11, time.UTC)},
			excepted: false,
		},
		{
			name: "not_equal_many",
			newIntervalMany: NewMany(
				Span{
					time.Date(2020, 10, 12, 12, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 14, 0, 0, 0, time.UTC)},
				Span{
					time.Date(2020, 10, 12, 15, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 16, 0, 0, 0, time.UTC)},
				Span{
					time.Date(2020, 10, 12, 21, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 22, 0, 0, 0, time.UTC)},
			),
			inputInterval: Span{
				time.Date(2020, 10, 12, 17, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 12, 18, 0, 0, 11, time.UTC)},
			excepted: false,
		},
		{
			name: "not_equal_offset",
			newIntervalMany: NewMany(
				Span{
					time.Date(2020, 10, 12, 9, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 10, 0, 0, 12, time.UTC)},
				Span{
					time.Date(2020, 10, 12, 19, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 20, 0, 0, 0, time.UTC)},
				Span{
					time.Date(2020, 10, 12, 16, 54, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 18, 0, 0, 11, time.UTC)},
			),
			offset: 5 * time.Minute,
			inputInterval: Span{
				time.Date(2020, 10, 12, 17, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 12, 18, 5, 0, 11, time.UTC)},
			excepted: false,
		},
		{
			name: "equal_offset",
			newIntervalMany: NewMany(
				Span{
					time.Date(2020, 10, 12, 9, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 10, 0, 0, 12, time.UTC)},
				Span{
					time.Date(2020, 10, 12, 19, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 20, 0, 0, 0, time.UTC)},
				Span{
					time.Date(2020, 10, 12, 16, 55, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 18, 0, 0, 11, time.UTC)},
			),
			offset: 5 * time.Minute,
			inputInterval: Span{
				time.Date(2020, 10, 12, 17, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 12, 18, 5, 0, 11, time.UTC)},
			excepted: true,
		},
		{
			name: "equal",
			newIntervalMany: NewMany(
				Span{
					time.Date(2020, 10, 12, 9, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 10, 0, 0, 12, time.UTC)},
				Span{
					time.Date(2020, 10, 12, 17, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 18, 0, 0, 11, time.UTC)},
				Span{
					time.Date(2020, 10, 12, 19, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 20, 0, 0, 0, time.UTC)},
				Span{
					time.Date(2020, 10, 12, 17, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 18, 0, 0, 11, time.UTC)},
			),
			inputInterval: Span{
				time.Date(2020, 10, 12, 17, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 12, 18, 0, 0, 11, time.UTC)},
			excepted: true,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result := tc.newIntervalMany.Equal(tc.inputInterval, tc.offset)
			assert.Equal(t, tc.excepted, result)
		})
	}
}

func TestIsIntersectionMany(t *testing.T) {
	testCases := []struct {
		name            string
		newIntervalMany SpanMany
		inputInterval   Span
		excepted        bool
	}{
		{
			name: "not_intersection_slightly",
			newIntervalMany: NewMany(
				Span{
					time.Date(2020, 10, 12, 12, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 13, 0, 0, 12, time.UTC)},
				Span{
					time.Date(2020, 10, 12, 17, 0, 0, 1, time.UTC),
					time.Date(2020, 10, 12, 18, 0, 0, 11, time.UTC)},
				Span{
					time.Date(2020, 10, 12, 18, 0, 0, 11, time.UTC),
					time.Date(2020, 10, 12, 110, 0, 0, 0, time.UTC)},
			),
			inputInterval: Span{
				time.Date(2020, 10, 12, 17, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 12, 18, 0, 0, 11, time.UTC)},
			excepted: false,
		},
		{
			name: "not_intersection_many",
			newIntervalMany: NewMany(
				Span{
					time.Date(2020, 10, 12, 12, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 13, 0, 0, 12, time.UTC)},
				Span{
					time.Date(2020, 10, 12, 22, 0, 0, 1, time.UTC),
					time.Date(2020, 10, 12, 23, 0, 0, 11, time.UTC)},
				Span{
					time.Date(2020, 10, 12, 19, 0, 0, 11, time.UTC),
					time.Date(2020, 10, 12, 20, 0, 0, 0, time.UTC)},
			),
			inputInterval: Span{
				time.Date(2020, 10, 12, 17, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 12, 18, 0, 0, 11, time.UTC)},
			excepted: false,
		},
		{
			name: "equal",
			newIntervalMany: NewMany(
				Span{
					time.Date(2020, 10, 12, 9, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 10, 0, 0, 12, time.UTC)},
				Span{
					time.Date(2020, 10, 12, 17, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 18, 0, 0, 10, time.UTC)},
				Span{
					time.Date(2020, 10, 12, 19, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 20, 0, 0, 0, time.UTC)},
				Span{
					time.Date(2020, 10, 12, 17, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 18, 0, 0, 11, time.UTC)},
			),
			inputInterval: Span{
				time.Date(2020, 10, 12, 17, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 12, 18, 0, 0, 11, time.UTC)},
			excepted: true,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result := tc.newIntervalMany.Equal(tc.inputInterval)
			assert.Equal(t, tc.excepted, result)
		})
	}
}

func TestExceptionIfIntersection(t *testing.T) {
	testCases := []struct {
		name              string
		newIntervalMany   SpanMany
		inputIntervalMany SpanMany

		wantResult SpanMany
	}{
		{
			newIntervalMany: NewMany(
				Span{time.Date(2020, 10, 12, 7, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 8, 0, 0, 0, time.UTC)},
				Span{time.Date(2020, 10, 12, 7, 30, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 8, 30, 0, 0, time.UTC)},
				Span{time.Date(2020, 10, 12, 8, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 9, 0, 0, 0, time.UTC)},
				Span{time.Date(2020, 10, 12, 8, 30, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 9, 30, 0, 0, time.UTC)},
				Span{time.Date(2020, 10, 12, 9, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 10, 0, 0, 0, time.UTC)},
			),
			inputIntervalMany: NewMany(
				Span{time.Date(2020, 10, 12, 6, 30, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 7, 35, 0, 0, time.UTC)},
				Span{time.Date(2020, 10, 12, 9, 30, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 9, 31, 0, 0, time.UTC)},
				Span{time.Date(2020, 10, 12, 8, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 8, 0, 0, 1, time.UTC)},
			),

			wantResult: NewMany(
				Span{
					time.Date(2020, 10, 12, 8, 30, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 9, 30, 0, 0, time.UTC)},
			),
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			listSpan := tc.newIntervalMany.ExceptionIfIntersection(tc.inputIntervalMany)
			assert.Equal(t, tc.wantResult, listSpan)
		})
	}
}

func TestExceptionIfNotEqual(t *testing.T) {
	testCases := []struct {
		name              string
		newIntervalMany   SpanMany
		inputIntervalMany SpanMany

		wantResult SpanMany
	}{
		{
			newIntervalMany: NewMany(
				Span{time.Date(2020, 10, 12, 8, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 9, 0, 0, 0, time.UTC)},
				Span{time.Date(2020, 10, 12, 8, 30, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 9, 30, 0, 0, time.UTC)},
				Span{time.Date(2020, 10, 12, 9, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 10, 0, 0, 0, time.UTC)},
				Span{time.Date(2020, 10, 12, 9, 30, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 10, 30, 0, 0, time.UTC)},
				Span{time.Date(2020, 10, 12, 10, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 11, 0, 0, 0, time.UTC)},
			),
			inputIntervalMany: NewMany(
				Span{time.Date(2020, 10, 12, 7, 30, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 8, 35, 0, 0, time.UTC)},
				Span{time.Date(2020, 10, 12, 8, 30, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 9, 29, 0, 0, time.UTC)},
				Span{time.Date(2020, 10, 12, 8, 30, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 9, 30, 0, 0, time.UTC)},
				Span{time.Date(2020, 10, 12, 10, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 11, 0, 0, 1, time.UTC)},
				Span{time.Date(2020, 10, 12, 10, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 11, 0, 0, 0, time.UTC)},
			),

			wantResult: NewMany(
				Span{
					time.Date(2020, 10, 12, 8, 30, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 9, 30, 0, 0, time.UTC)},
				Span{
					time.Date(2020, 10, 12, 10, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 11, 0, 0, 0, time.UTC)},
			),
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			listSpan := tc.newIntervalMany.ExceptionIfNotEqual(tc.inputIntervalMany)
			assert.Equal(t, tc.wantResult, listSpan)
		})
	}
}

func TestExceptionIfNotContains(t *testing.T) {
	testCases := []struct {
		name              string
		newIntervalMany   SpanMany
		inputIntervalMany SpanMany

		wantResult SpanMany
	}{
		{
			newIntervalMany: NewMany(
				Span{time.Date(2020, 11, 20, 8, 0, 0, 0, time.UTC),
					time.Date(2020, 11, 20, 9, 0, 0, 0, time.UTC)},
				Span{time.Date(2020, 11, 20, 8, 30, 0, 0, time.UTC),
					time.Date(2020, 11, 20, 9, 30, 0, 0, time.UTC)},
				Span{time.Date(2020, 11, 20, 9, 0, 0, 0, time.UTC),
					time.Date(2020, 11, 20, 10, 0, 0, 0, time.UTC)},
				Span{time.Date(2020, 11, 20, 9, 30, 0, 0, time.UTC),
					time.Date(2020, 11, 20, 10, 30, 0, 0, time.UTC)},
				Span{time.Date(2020, 11, 20, 10, 0, 0, 0, time.UTC),
					time.Date(2020, 11, 20, 11, 0, 0, 0, time.UTC)},
				Span{time.Date(2020, 11, 20, 1, 0, 0, 0, time.UTC),
					time.Date(2020, 11, 20, 1, 5, 0, 0, time.UTC)},
				Span{time.Date(2020, 11, 20, 1, 5, 0, 0, time.UTC),
					time.Date(2020, 11, 20, 1, 10, 0, 0, time.UTC)},
				Span{time.Date(2020, 11, 20, 1, 30, 0, 0, time.UTC),
					time.Date(2020, 11, 20, 1, 33, 0, 0, time.UTC)},
			),
			inputIntervalMany: NewMany(
				Span{time.Date(2020, 11, 20, 0, 0, 0, 0, time.UTC),
					time.Date(2020, 11, 20, 2, 0, 0, 0, time.UTC)},
				Span{time.Date(2020, 11, 20, 7, 30, 0, 0, time.UTC),
					time.Date(2020, 11, 20, 8, 35, 0, 0, time.UTC)},
				Span{time.Date(2020, 11, 20, 8, 30, 0, 0, time.UTC),
					time.Date(2020, 11, 20, 9, 29, 0, 0, time.UTC)},
				Span{time.Date(2020, 11, 20, 8, 30, 0, 0, time.UTC),
					time.Date(2020, 11, 20, 9, 30, 0, 0, time.UTC)},
				Span{time.Date(2020, 11, 20, 10, 0, 0, 0, time.UTC),
					time.Date(2020, 11, 20, 11, 0, 0, 1, time.UTC)},
				Span{time.Date(2020, 11, 20, 10, 0, 0, 0, time.UTC),
					time.Date(2020, 11, 20, 11, 0, 0, 0, time.UTC)},
			),

			wantResult: NewMany(
				Span{
					time.Date(2020, 11, 20, 8, 30, 0, 0, time.UTC),
					time.Date(2020, 11, 20, 9, 30, 0, 0, time.UTC)},
				Span{
					time.Date(2020, 11, 20, 10, 0, 0, 0, time.UTC),
					time.Date(2020, 11, 20, 11, 0, 0, 0, time.UTC)},
				Span{time.Date(2020, 11, 20, 1, 0, 0, 0, time.UTC),
					time.Date(2020, 11, 20, 1, 5, 0, 0, time.UTC)},
				Span{time.Date(2020, 11, 20, 1, 5, 0, 0, time.UTC),
					time.Date(2020, 11, 20, 1, 10, 0, 0, time.UTC)},
				Span{time.Date(2020, 11, 20, 1, 30, 0, 0, time.UTC),
					time.Date(2020, 11, 20, 1, 33, 0, 0, time.UTC)},
			),
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			listSpan := tc.newIntervalMany.ExceptionIfNotContains(tc.inputIntervalMany)
			assert.Equal(t, tc.wantResult, listSpan)
		})
	}
}

func TestIntersectionMany(t *testing.T) {
	testCases := []struct {
		name        string
		newSpanMany SpanMany
		inputSpan   Span

		excepted SpanMany
	}{
		{
			name: "intersection",
			newSpanMany: NewMany(
				Span{
					time.Date(2020, 10, 12, 15, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 17, 0, 0, 0, time.UTC)},
				Span{
					time.Date(2020, 10, 12, 7, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 12, 0, 0, 0, time.UTC)},
				Span{
					time.Date(2020, 10, 12, 20, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 22, 0, 0, 0, time.UTC)},
				Span{
					time.Date(2020, 10, 12, 18, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 23, 0, 0, 0, time.UTC)},
				Span{
					time.Date(2020, 10, 12, 7, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 10, 0, 0, 0, time.UTC)},
			),
			inputSpan: Span{
				time.Date(2020, 10, 12, 10, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 12, 19, 0, 0, 0, time.UTC),
			},
			excepted: NewMany(
				Span{
					time.Date(2020, 10, 12, 15, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 17, 0, 0, 0, time.UTC)},
				Span{
					time.Date(2020, 10, 12, 10, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 12, 0, 0, 0, time.UTC)},
				Span{
					time.Date(2020, 10, 12, 18, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 19, 0, 0, 0, time.UTC)},
			),
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result := tc.newSpanMany.Intersection(tc.inputSpan)
			assert.Equal(t, tc.excepted, result)
		})
	}
}

func TestUnionMany(t *testing.T) {
	testCases := []struct {
		name              string
		newIntervalMany   SpanMany
		inputIntervalMany SpanMany

		excepted SpanMany
	}{
		{
			name:              "empty",
			newIntervalMany:   NewMany(),
			inputIntervalMany: NewMany(),

			excepted: NewMany(),
		},
		{
			name:            "new_empty",
			newIntervalMany: NewMany(),
			inputIntervalMany: NewMany(
				Span{
					time.Date(2020, 10, 17, 7, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 17, 8, 0, 0, 0, time.UTC)},
			),
			excepted: NewMany(
				Span{
					time.Date(2020, 10, 17, 7, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 17, 8, 0, 0, 0, time.UTC)},
			),
		},
		{
			name: "input_empty",
			newIntervalMany: NewMany(
				Span{
					time.Date(2020, 10, 17, 7, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 17, 8, 0, 0, 0, time.UTC)},
			),
			inputIntervalMany: NewMany(),
			excepted: NewMany(
				Span{
					time.Date(2020, 10, 17, 7, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 17, 8, 0, 0, 0, time.UTC)},
			),
		},
		{
			name: "many_test",
			newIntervalMany: NewMany(
				Span{time.Date(2020, 10, 17, 12, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 17, 14, 0, 0, 0, time.UTC)},
				Span{time.Date(2020, 10, 17, 22, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 17, 23, 0, 0, 0, time.UTC)},
				Span{time.Date(2020, 10, 17, 13, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 17, 17, 0, 0, 0, time.UTC)},
			),
			inputIntervalMany: NewMany(
				Span{time.Date(2020, 10, 17, 7, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 17, 10, 0, 0, 0, time.UTC)},
				Span{time.Date(2020, 10, 17, 21, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 17, 23, 0, 0, 0, time.UTC)},
				Span{time.Date(2020, 10, 17, 11, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 17, 15, 0, 0, 0, time.UTC)}),
			excepted: NewMany(
				Span{time.Date(2020, 10, 17, 7, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 17, 10, 0, 0, 0, time.UTC)},
				Span{time.Date(2020, 10, 17, 11, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 17, 17, 0, 0, 0, time.UTC)},
				Span{time.Date(2020, 10, 17, 21, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 17, 23, 0, 0, 0, time.UTC)},
			),
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result := tc.newIntervalMany.Union(tc.inputIntervalMany)
			assert.Equal(t, tc.excepted, result)
		})
	}
}

func TestExceptMany(t *testing.T) {
	testCases := []struct {
		name        string
		newSpanMany SpanMany
		inputSpan   Span

		excepted SpanMany
	}{
		{
			name: "many_result",
			newSpanMany: NewMany(
				Span{
					time.Date(2020, 10, 17, 7, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 17, 10, 0, 0, 0, time.UTC),
				},
				Span{
					time.Date(2020, 10, 17, 14, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 17, 15, 0, 0, 0, time.UTC),
				},
				Span{
					time.Date(2020, 10, 17, 12, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 17, 16, 0, 0, 0, time.UTC),
				},
				Span{
					time.Date(2020, 10, 17, 12, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 17, 14, 30, 0, 0, time.UTC),
				},
				Span{
					time.Date(2020, 10, 17, 14, 33, 0, 0, time.UTC),
					time.Date(2020, 10, 17, 18, 0, 0, 0, time.UTC),
				},
			),
			inputSpan: Span{
				time.Date(2020, 10, 17, 14, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 17, 15, 0, 0, 0, time.UTC),
			},
			excepted: NewMany(
				Span{
					time.Date(2020, 10, 17, 7, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 17, 10, 0, 0, 0, time.UTC),
				},
				Span{
					time.Date(2020, 10, 17, 12, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 17, 14, 0, 0, 0, time.UTC),
				},
				Span{
					time.Date(2020, 10, 17, 15, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 17, 18, 0, 0, 0, time.UTC),
				},
			),
		},
		{
			name:        "empty_new_span",
			newSpanMany: NewMany(),
			inputSpan: Span{
				time.Date(2020, 10, 17, 14, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 17, 15, 0, 0, 0, time.UTC),
			},
			excepted: NewMany(),
		},
		{
			name: "empty_input_span",
			newSpanMany: NewMany(
				Span{
					time.Date(2020, 10, 17, 22, 33, 0, 0, time.UTC),
					time.Date(2020, 10, 17, 23, 7, 0, 0, time.UTC),
				},
			),
			inputSpan: Span{},

			excepted: NewMany(
				Span{
					time.Date(2020, 10, 17, 22, 33, 0, 0, time.UTC),
					time.Date(2020, 10, 17, 23, 7, 0, 0, time.UTC),
				},
			),
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result := tc.newSpanMany.Except(tc.inputSpan)
			assert.Equal(t, tc.excepted, result)
		})
	}
}
