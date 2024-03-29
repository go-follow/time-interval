package timeinterval

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	testCases := []struct {
		name     string
		start    time.Time
		end      time.Time
		wantErr  error
		wantSpan Span
	}{
		{
			name:     "start equal end",
			start:    time.Date(2022, 2, 12, 7, 30, 0, 0, time.UTC),
			end:      time.Date(2022, 2, 12, 7, 30, 0, 0, time.UTC),
			wantErr:  errors.New("time start cannot be more time end"),
			wantSpan: Span{},
		},
		{
			name:     "start more end",
			start:    time.Date(2022, 2, 12, 7, 30, 0, 1, time.UTC),
			end:      time.Date(2022, 2, 12, 7, 30, 0, 0, time.UTC),
			wantErr:  errors.New("time start cannot be more time end"),
			wantSpan: Span{},
		},
		{
			name:    "start less end",
			start:   time.Date(2022, 2, 12, 5, 30, 0, 0, time.UTC),
			end:     time.Date(2022, 2, 12, 7, 30, 0, 0, time.UTC),
			wantErr: nil,
			wantSpan: Span{
				start: time.Date(2022, 2, 12, 5, 30, 0, 0, time.UTC),
				end:   time.Date(2022, 2, 12, 7, 30, 0, 0, time.UTC),
			},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			interval, err := New(tc.start, tc.end)
			assert.Equal(t, tc.wantErr, err)
			assert.Equal(t, tc.wantSpan, interval)
		})
	}
}

func TestEqual(t *testing.T) {
	testCases := []struct {
		name          string
		newInterval   Span
		inputInterval Span
		offset        time.Duration
		excepted      bool
	}{
		{
			name: "not_equal_slightly",
			newInterval: Span{
				start: time.Date(2020, 10, 11, 17, 30, 0, 0, time.UTC),
				end:   time.Date(2020, 10, 11, 17, 30, 17, 12, time.UTC),
			},
			inputInterval: Span{
				start: time.Date(2020, 10, 11, 17, 30, 0, 0, time.UTC),
				end:   time.Date(2020, 10, 11, 17, 30, 17, 11, time.UTC),
			},
			excepted: false,
		},
		{
			name: "equal_slightly_with_offset",
			newInterval: Span{
				start: time.Date(2020, 10, 11, 15, 0, 5, 0, time.UTC),
				end:   time.Date(2020, 10, 11, 16, 0, 5, 0, time.UTC),
			},
			inputInterval: Span{
				start: time.Date(2020, 10, 11, 15, 0, 0, 0, time.UTC),
				end:   time.Date(2020, 10, 11, 16, 0, 0, 0, time.UTC),
			},
			offset:   time.Second * 5,
			excepted: true,
		},
		{
			name: "equal_offset_5_minute",
			newInterval: Span{
				start: time.Date(2020, 10, 11, 15, 5, 0, 0, time.UTC),
				end:   time.Date(2020, 10, 11, 16, 0, 0, 0, time.UTC),
			},
			inputInterval: Span{
				start: time.Date(2020, 10, 11, 15, 0, 0, 0, time.UTC),
				end:   time.Date(2020, 10, 11, 16, 5, 0, 0, time.UTC),
			},
			offset:   time.Minute * 5,
			excepted: true,
		},
		{
			name: "not_equal_offset_5_minute",
			newInterval: Span{
				start: time.Date(2020, 10, 11, 15, 30, 0, 0, time.UTC),
				end:   time.Date(2020, 10, 11, 16, 20, 0, 0, time.UTC),
			},
			inputInterval: Span{
				start: time.Date(2020, 10, 11, 16, 0, 0, 0, time.UTC),
				end:   time.Date(2020, 10, 11, 16, 20, 0, 0, time.UTC),
			},
			offset:   time.Minute * 5,
			excepted: false,
		},
		{
			name: "not_equal_many",
			newInterval: Span{
				start: time.Date(2020, 10, 11, 17, 30, 0, 0, time.UTC),
				end:   time.Date(2020, 10, 11, 18, 0, 0, 0, time.UTC),
			},
			inputInterval: Span{
				start: time.Date(2020, 10, 11, 22, 30, 0, 0, time.UTC),
				end:   time.Date(2020, 10, 11, 23, 30, 17, 11, time.UTC),
			},
			excepted: false,
		},
		{
			name: "equal_many_with_offset",
			newInterval: Span{
				start: time.Date(2020, 10, 11, 15, 0, 0, 0, time.UTC),
				end:   time.Date(2020, 10, 11, 21, 0, 0, 0, time.UTC),
			},
			inputInterval: Span{
				start: time.Date(2020, 10, 11, 14, 0, 0, 0, time.UTC),
				end:   time.Date(2020, 10, 11, 18, 0, 0, 0, time.UTC),
			},
			offset:   3 * time.Hour,
			excepted: true,
		},
		{
			name: "full_equal",
			newInterval: Span{
				start: time.Date(2020, 10, 11, 17, 30, 0, 0, time.UTC),
				end:   time.Date(2020, 10, 11, 17, 30, 17, 12, time.UTC),
			},
			inputInterval: Span{
				start: time.Date(2020, 10, 11, 17, 30, 0, 0, time.UTC),
				end:   time.Date(2020, 10, 11, 17, 30, 17, 12, time.UTC),
			},
			excepted: true,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result := tc.newInterval.Equal(tc.inputInterval, tc.offset)
			assert.Equal(t, tc.excepted, result)
		})
	}
}

func TestIsIntersection(t *testing.T) {
	testCases := []struct {
		name          string
		newInterval   Span
		inputInterval Span
		offset        time.Duration
		excepted      bool
	}{
		{
			name: "invert_case",
			newInterval: Span{
				start: time.Date(2020, 10, 11, 11, 0, 0, 0, time.UTC),
				end:   time.Date(2020, 10, 11, 15, 0, 0, 0, time.UTC),
			},
			inputInterval: Span{
				start: time.Date(2020, 10, 11, 7, 0, 0, 0, time.UTC),
				end:   time.Date(2020, 10, 11, 12, 0, 0, 0, time.UTC),
			},
			excepted: true,
		},
		{
			name: "intersection_slightly",
			newInterval: Span{
				start: time.Date(2020, 10, 11, 17, 30, 0, 0, time.UTC),
				end:   time.Date(2020, 10, 11, 18, 22, 0, 0, time.UTC),
			},
			inputInterval: Span{
				start: time.Date(2020, 10, 11, 16, 0, 0, 0, time.UTC),
				end:   time.Date(2020, 10, 11, 17, 30, 0, 1, time.UTC),
			},
			excepted: true,
		},
		{
			name: "not_intersection_slightly_with_offset",
			newInterval: Span{
				start: time.Date(2020, 10, 11, 17, 30, 0, 0, time.UTC),
				end:   time.Date(2020, 10, 11, 18, 22, 0, 0, time.UTC),
			},
			inputInterval: Span{
				start: time.Date(2020, 10, 11, 16, 0, 0, 0, time.UTC),
				end:   time.Date(2020, 10, 11, 17, 30, 5, 0, time.UTC),
			},
			offset:   time.Second * 5,
			excepted: false,
		},
		{
			name: "intersection_slightly_with_offset",
			newInterval: Span{
				start: time.Date(2020, 10, 11, 17, 30, 0, 0, time.UTC),
				end:   time.Date(2020, 10, 11, 18, 22, 0, 0, time.UTC),
			},
			inputInterval: Span{
				start: time.Date(2020, 10, 11, 16, 0, 0, 0, time.UTC),
				end:   time.Date(2020, 10, 11, 17, 30, 6, 0, time.UTC),
			},
			offset:   time.Second * 5,
			excepted: true,
		},
		{
			name: "intersection_many",
			newInterval: Span{
				start: time.Date(2020, 10, 11, 17, 30, 0, 0, time.UTC),
				end:   time.Date(2020, 10, 11, 18, 22, 0, 0, time.UTC),
			},
			inputInterval: Span{
				start: time.Date(2020, 10, 11, 17, 10, 0, 0, time.UTC),
				end:   time.Date(2020, 10, 11, 18, 22, 0, 0, time.UTC),
			},
			excepted: true,
		},
		{
			name: "not_intersection_many",
			newInterval: Span{
				start: time.Date(2020, 10, 11, 17, 30, 0, 0, time.UTC),
				end:   time.Date(2020, 10, 11, 18, 22, 0, 0, time.UTC),
			},
			inputInterval: Span{
				start: time.Date(2020, 10, 11, 17, 10, 0, 0, time.UTC),
				end:   time.Date(2020, 10, 11, 18, 22, 0, 0, time.UTC),
			},
			offset:   time.Hour * 1,
			excepted: false,
		},
		{
			name: "not_intersection",
			newInterval: Span{
				start: time.Date(2020, 10, 11, 17, 30, 0, 0, time.UTC),
				end:   time.Date(2020, 10, 11, 18, 22, 0, 0, time.UTC),
			},
			inputInterval: Span{
				start: time.Date(2020, 10, 11, 18, 22, 0, 0, time.UTC),
				end:   time.Date(2020, 10, 11, 19, 0, 0, 0, time.UTC),
			},
			excepted: false,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result := tc.newInterval.IsIntersection(tc.inputInterval, tc.offset)
			assert.Equal(t, tc.excepted, result)
		})
	}
}

func TestIntersection(t *testing.T) {
	testCases := []struct {
		name      string
		newSpan   Span
		inputSpan Span

		excepted Span
	}{
		{
			name: "input_contains_new",
			newSpan: Span{
				start: time.Date(2020, 10, 11, 15, 0, 0, 0, time.UTC),
				end:   time.Date(2020, 10, 11, 17, 0, 0, 0, time.UTC),
			},
			inputSpan: Span{
				start: time.Date(2020, 10, 11, 10, 0, 0, 0, time.UTC),
				end:   time.Date(2020, 10, 11, 19, 0, 0, 0, time.UTC),
			},
			excepted: Span{
				start: time.Date(2020, 10, 11, 15, 0, 0, 0, time.UTC),
				end:   time.Date(2020, 10, 11, 17, 0, 0, 0, time.UTC),
			},
		},
		{
			name: "new_contains_input",
			newSpan: Span{
				start: time.Date(2020, 10, 11, 11, 0, 0, 0, time.UTC),
				end:   time.Date(2020, 10, 11, 14, 0, 0, 0, time.UTC),
			},
			inputSpan: Span{
				start: time.Date(2020, 10, 11, 12, 0, 0, 0, time.UTC),
				end:   time.Date(2020, 10, 11, 13, 0, 0, 0, time.UTC),
			},
			excepted: Span{
				start: time.Date(2020, 10, 11, 12, 0, 0, 0, time.UTC),
				end:   time.Date(2020, 10, 11, 13, 0, 0, 0, time.UTC),
			},
		},
		{
			name: "not_intersection",
			newSpan: Span{
				start: time.Date(2020, 10, 11, 7, 0, 0, 0, time.UTC),
				end:   time.Date(2020, 10, 11, 12, 0, 0, 0, time.UTC),
			},
			inputSpan: Span{
				start: time.Date(2020, 10, 11, 12, 0, 0, 0, time.UTC),
				end:   time.Date(2020, 10, 11, 15, 0, 0, 0, time.UTC),
			},
			excepted: Span{},
		},
		{
			name: "not_intersection_many",
			newSpan: Span{
				start: time.Date(2020, 10, 11, 3, 0, 0, 0, time.UTC),
				end:   time.Date(2020, 10, 11, 7, 0, 0, 0, time.UTC),
			},
			inputSpan: Span{
				start: time.Date(2020, 10, 11, 22, 0, 0, 0, time.UTC),
				end:   time.Date(2020, 10, 11, 23, 0, 0, 0, time.UTC),
			},
			excepted: Span{},
		},
		{
			name: "intersection_new_left",
			newSpan: Span{
				start: time.Date(2020, 10, 11, 7, 0, 0, 0, time.UTC),
				end:   time.Date(2020, 10, 11, 12, 0, 0, 0, time.UTC),
			},
			inputSpan: Span{
				start: time.Date(2020, 10, 11, 10, 0, 0, 0, time.UTC),
				end:   time.Date(2020, 10, 11, 15, 0, 0, 0, time.UTC),
			},
			excepted: Span{
				start: time.Date(2020, 10, 11, 10, 0, 0, 0, time.UTC),
				end:   time.Date(2020, 10, 11, 12, 0, 0, 0, time.UTC),
			},
		},
		{
			name: "intersection_new_right",
			newSpan: Span{
				start: time.Date(2020, 10, 11, 11, 0, 0, 0, time.UTC),
				end:   time.Date(2020, 10, 11, 14, 0, 0, 0, time.UTC),
			},
			inputSpan: Span{
				start: time.Date(2020, 10, 11, 8, 0, 0, 0, time.UTC),
				end:   time.Date(2020, 10, 11, 12, 0, 0, 0, time.UTC),
			},
			excepted: Span{
				start: time.Date(2020, 10, 11, 11, 0, 0, 0, time.UTC),
				end:   time.Date(2020, 10, 11, 12, 0, 0, 0, time.UTC),
			},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result := tc.newSpan.Intersection(tc.inputSpan)
			fmt.Println(result.String())
			assert.Equal(t, tc.excepted, result)
		})
	}
}

func TestUnion(t *testing.T) {
	testCases := []struct {
		name          string
		newInterval   Span
		inputInterval Span

		excepted SpanMany
	}{
		{
			name: "boundaries are equal",
			newInterval: Span{
				time.Date(2020, 10, 1, 7, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 1, 9, 0, 0, 0, time.UTC)},
			inputInterval: Span{
				time.Date(2020, 10, 1, 9, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 1, 15, 0, 0, 0, time.UTC)},
			excepted: NewMany(Span{
				time.Date(2020, 10, 1, 7, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 1, 15, 0, 0, 0, time.UTC)}),
		},
		{
			name: "new_absorbs_input",
			newInterval: Span{
				time.Date(2020, 10, 1, 7, 15, 0, 0, time.UTC),
				time.Date(2020, 10, 1, 17, 23, 11, 0, time.UTC)},
			inputInterval: Span{
				time.Date(2020, 10, 1, 10, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 1, 11, 0, 11, 0, time.UTC)},
			excepted: NewMany(Span{
				time.Date(2020, 10, 1, 7, 15, 0, 0, time.UTC),
				time.Date(2020, 10, 1, 17, 23, 11, 0, time.UTC)}),
		},
		{
			name: "input_absorbs_new",
			newInterval: Span{
				time.Date(2020, 10, 1, 12, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 1, 14, 0, 0, 0, time.UTC)},
			inputInterval: Span{
				time.Date(2020, 10, 1, 10, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 1, 15, 0, 0, 0, time.UTC)},
			excepted: NewMany(Span{
				time.Date(2020, 10, 1, 10, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 1, 15, 0, 0, 0, time.UTC)}),
		},
		{
			name: "new_left_input",
			newInterval: Span{
				time.Date(2020, 10, 1, 5, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 1, 12, 0, 0, 0, time.UTC)},
			inputInterval: Span{
				time.Date(2020, 10, 1, 10, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 1, 14, 0, 0, 0, time.UTC)},
			excepted: NewMany(Span{
				time.Date(2020, 10, 1, 5, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 1, 14, 0, 0, 0, time.UTC)}),
		},
		{
			name: "input_left_new",
			newInterval: Span{
				time.Date(2020, 10, 1, 21, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 1, 23, 0, 0, 0, time.UTC)},
			inputInterval: Span{
				time.Date(2020, 10, 1, 7, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 1, 22, 0, 0, 0, time.UTC)},
			excepted: NewMany(Span{
				time.Date(2020, 10, 1, 7, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 1, 23, 0, 0, 0, time.UTC)}),
		},
		{
			name: "new_next_input",
			newInterval: Span{
				time.Date(2020, 10, 1, 4, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 1, 7, 0, 0, 0, time.UTC)},
			inputInterval: Span{
				time.Date(2020, 10, 1, 15, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 1, 22, 0, 0, 0, time.UTC)},
			excepted: NewMany(
				Span{
					time.Date(2020, 10, 1, 4, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 1, 7, 0, 0, 0, time.UTC)},
				Span{
					time.Date(2020, 10, 1, 15, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 1, 22, 0, 0, 0, time.UTC)}),
		},
		{
			name: "input_next_new",
			newInterval: Span{
				time.Date(2020, 10, 1, 15, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 1, 22, 0, 0, 0, time.UTC)},
			inputInterval: Span{
				time.Date(2020, 10, 1, 4, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 1, 7, 0, 0, 0, time.UTC)},
			excepted: NewMany(
				Span{
					time.Date(2020, 10, 1, 15, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 1, 22, 0, 0, 0, time.UTC)},
				Span{
					time.Date(2020, 10, 1, 4, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 1, 7, 0, 0, 0, time.UTC)}),
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result := tc.newInterval.Union(tc.inputInterval)
			assert.Equal(t, tc.excepted, result)
		})
	}
}

func TestExcept(t *testing.T) {
	testCases := []struct {
		name      string
		newSpan   Span
		inputSpan Span

		excepted SpanMany
	}{
		{
			name: "not_intersection",
			newSpan: Span{
				time.Date(2020, 10, 14, 10, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 14, 12, 0, 0, 0, time.UTC),
			},
			inputSpan: Span{
				time.Date(2020, 10, 14, 12, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 14, 18, 0, 0, 0, time.UTC),
			},
			excepted: NewMany(
				Span{
					time.Date(2020, 10, 14, 10, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 14, 12, 0, 0, 0, time.UTC),
				},
			),
		},
		{
			name: "full_intersection",
			newSpan: Span{
				time.Date(2020, 10, 14, 12, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 14, 14, 0, 0, 0, time.UTC),
			},
			inputSpan: Span{
				time.Date(2020, 10, 14, 7, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 14, 15, 0, 0, 0, time.UTC),
			},
			excepted: NewMany(),
		},
		{
			name: "torn_result",
			newSpan: Span{
				time.Date(2020, 10, 14, 7, 15, 0, 0, time.UTC),
				time.Date(2020, 10, 14, 15, 22, 0, 0, time.UTC),
			},
			inputSpan: Span{
				time.Date(2020, 10, 14, 12, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 14, 14, 0, 0, 0, time.UTC),
			},
			excepted: NewMany(
				Span{
					time.Date(2020, 10, 14, 7, 15, 0, 0, time.UTC),
					time.Date(2020, 10, 14, 12, 0, 0, 0, time.UTC),
				},
				Span{
					time.Date(2020, 10, 14, 14, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 14, 15, 22, 0, 0, time.UTC),
				},
			),
		},
		{
			name: "right_takeover",
			newSpan: Span{
				time.Date(2020, 10, 14, 9, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 14, 12, 0, 0, 0, time.UTC),
			},
			inputSpan: Span{
				time.Date(2020, 10, 14, 11, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 14, 15, 0, 0, 0, time.UTC),
			},
			excepted: NewMany(
				Span{
					time.Date(2020, 10, 14, 9, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 14, 11, 0, 0, 0, time.UTC),
				},
			),
		},
		{
			name: "left_takeover",
			newSpan: Span{
				time.Date(2020, 10, 14, 10, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 14, 15, 0, 0, 0, time.UTC),
			},
			inputSpan: Span{
				time.Date(2020, 10, 14, 9, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 14, 14, 30, 0, 0, time.UTC),
			},
			excepted: NewMany(
				Span{
					time.Date(2020, 10, 14, 14, 30, 0, 0, time.UTC),
					time.Date(2020, 10, 14, 15, 0, 0, 0, time.UTC),
				},
			),
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result := tc.newSpan.Except(tc.inputSpan)
			assert.Equal(t, tc.excepted, result)
		})
	}
}

func TestIsContains(t *testing.T) {
	testCases := []struct {
		name      string
		newSpan   Span
		inputSpan Span
		offset    time.Duration
		excepted  bool
	}{
		{
			name: "contains",
			newSpan: Span{
				time.Date(2020, 11, 20, 7, 0, 0, 0, time.UTC),
				time.Date(2020, 11, 20, 9, 0, 0, 0, time.UTC),
			},
			inputSpan: Span{
				time.Date(2020, 11, 20, 8, 0, 0, 0, time.UTC),
				time.Date(2020, 11, 20, 8, 30, 0, 0, time.UTC),
			},
			excepted: true,
		},
		{
			name: "equal",
			newSpan: Span{
				time.Date(2020, 11, 20, 7, 0, 0, 0, time.UTC),
				time.Date(2020, 11, 20, 9, 0, 0, 0, time.UTC),
			},
			inputSpan: Span{
				time.Date(2020, 11, 20, 7, 0, 0, 0, time.UTC),
				time.Date(2020, 11, 20, 9, 0, 0, 0, time.UTC),
			},
			excepted: true,
		},
		{
			name: "reverse_contains",
			newSpan: Span{
				time.Date(2020, 11, 20, 8, 0, 0, 0, time.UTC),
				time.Date(2020, 11, 20, 8, 30, 0, 0, time.UTC),
			},
			inputSpan: Span{
				time.Date(2020, 11, 20, 7, 0, 0, 0, time.UTC),
				time.Date(2020, 11, 20, 9, 0, 0, 0, time.UTC),
			},
			offset:   time.Hour,
			excepted: true,
		},
		{
			name: "not_contains_left",
			newSpan: Span{
				time.Date(2020, 11, 20, 7, 0, 0, 0, time.UTC),
				time.Date(2020, 11, 20, 9, 0, 0, 0, time.UTC),
			},
			inputSpan: Span{
				time.Date(2020, 11, 20, 6, 59, 0, 0, time.UTC),
				time.Date(2020, 11, 20, 9, 0, 0, 0, time.UTC),
			},
			excepted: false,
		},
		{
			name: "contains_left_offset",
			newSpan: Span{
				time.Date(2020, 11, 20, 7, 0, 0, 0, time.UTC),
				time.Date(2020, 11, 20, 9, 0, 0, 0, time.UTC),
			},
			inputSpan: Span{
				time.Date(2020, 11, 20, 6, 59, 0, 0, time.UTC),
				time.Date(2020, 11, 20, 9, 0, 0, 0, time.UTC),
			},
			offset:   time.Minute,
			excepted: true,
		},
		{
			name: "not_contains_right",
			newSpan: Span{
				time.Date(2020, 11, 20, 7, 0, 0, 0, time.UTC),
				time.Date(2020, 11, 20, 9, 0, 0, 0, time.UTC),
			},
			inputSpan: Span{
				time.Date(2020, 11, 20, 8, 0, 0, 0, time.UTC),
				time.Date(2020, 11, 20, 9, 0, 0, 1, time.UTC),
			},
			excepted: false,
		},
		{
			name: "contains_right_offset",
			newSpan: Span{
				time.Date(2020, 11, 20, 7, 0, 0, 0, time.UTC),
				time.Date(2020, 11, 20, 9, 0, 0, 0, time.UTC),
			},
			inputSpan: Span{
				time.Date(2020, 11, 20, 8, 0, 0, 0, time.UTC),
				time.Date(2020, 11, 20, 9, 0, 0, 1, time.UTC),
			},
			offset:   time.Second,
			excepted: true,
		},
		{
			name: "many_not_contains",
			newSpan: Span{
				time.Date(2020, 11, 20, 7, 0, 0, 0, time.UTC),
				time.Date(2020, 11, 20, 9, 0, 0, 0, time.UTC),
			},
			inputSpan: Span{
				time.Date(2020, 11, 20, 1, 0, 0, 0, time.UTC),
				time.Date(2020, 11, 20, 5, 0, 0, 0, time.UTC),
			},
			excepted: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.newSpan.IsContains(tc.inputSpan, tc.offset)
			assert.Equal(t, tc.excepted, result)
		})
	}
}
