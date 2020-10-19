# time-interval [![godoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](github.com/go-follow/time-interval) [![license](https://img.shields.io/badge/license-MIT-green.svg?style=flat)](https://github.com/go-follow/time-interval/blob/master/LICENSE)

This package helps to work with time intervals. The functionality allows you to perform the following basic operations:
* Union - merging of time intervals
* Intersection - finding intersections of regular intervals
* Except - differentiation between fixed intervals
* Equal - comparison of time intervals
* IsIntersection - check for intersection of time intervals

## install
```
go get github.com/go-follow/time-interval

```
## usage
To start using the functionality, you need to initialize a time slot or time slots.
Attention: when initializing intervals, it is important to know that the end date of 
the interval must be greater than the start date of the interval, otherwise there will be panic.

An example of an example where panic: time start cannot be more time end
```go
package main
 
import (    
    "time"

    interval "github.com/go-follow/time-interval"
)

func main() {
    timeStart := time.Date(2020, 10, 17, 23, 0, 0, 0, time.UTC)
    timeEnd := time.Date(2020, 10, 17, 15, 0, 0, 0, time.UTC)
    interval.New(timeStart, timeEnd)
}
```

An example of using a package correctly
```go
package main
 
import (
    "fmt"
    "time"

    interval "github.com/go-follow/time-interval"
)

func main() {
    timeStart1 := time.Date(2020, 10, 17, 10, 0, 0, 0, time.UTC)
    timeEnd1 := time.Date(2020, 10, 17, 15, 0, 0, 0, time.UTC)
    timeStart2 := time.Date(2020, 10, 17, 14, 0, 0, 0, time.UTC)
    timeEnd2 := time.Date(2020, 10, 17, 19, 0, 0, 0, time.UTC)
    
    newInt := interval.New(timeStart1, timeEnd1)
    newInt2 := interval.New(timeStart2, timeEnd2)
    
    newIntMany := interval.NewMany(newInt, newInt2)
    result := newIntMany.Union()
    fmt.Println(result.String()) // 2020-10-17 10:00:00 +0000 UTC - 2020-10-17 19:00:00 +0000 UTC
}
```
## examples
* Except

For the Except operation, it is important where the subtraction comes from.
For SpanMany, before returning the final result, is sorted and merged
```go
package main
 
import (
    "fmt"
    "time"

    interval "github.com/go-follow/time-interval"
)

func main() {
    timeStart1 := time.Date(2020, 10, 18, 9, 0, 0, 0, time.UTC)
    timeEnd1 := time.Date(2020, 10, 18, 12, 0, 0, 0, time.UTC)
    timeStart2 := time.Date(2020, 10, 18, 11, 0, 0, 0, time.UTC)
    timeEnd2 := time.Date(2020, 10, 18, 15, 0, 0, 0, time.UTC)

    ti1 := interval.New(timeStart1, timeEnd1)
    ti2 := interval.New(timeStart2, timeEnd2)
    
    // ti1 \ ti2
    ti1ExceptTi2 := ti1.Except(ti2)
    fmt.Println(ti1ExceptTi2.String()) // [ 2020-10-18 10:00:00 +0000 UTC - 2020-10-18 11:00:00 +0000 UTC ]
    
    // ti2 \ ti1
    ti2ExceptTi1 := ti2.Except(ti1)
    fmt.Println(ti2ExceptTi1.String()) // [ 2020-10-18 12:00:00 +0000 UTC - 2020-10-18 15:00:00 +0000 UTC ]
    
    // Except for SpanMany
    intervalMany := interval.NewMany(
        interval.New(time.Date(2020, 10, 18, 7, 0, 0, 0, time.UTC), time.Date(2020, 10, 18, 10, 0, 0, 0, time.UTC)),
        interval.New(time.Date(2020, 10, 18, 14, 0, 0, 0, time.UTC), time.Date(2020, 10, 18, 15, 0, 0, 0, time.UTC)),
        interval.New(time.Date(2020, 10, 18, 12, 0, 0, 0, time.UTC), time.Date(2020, 10, 18, 16, 0, 0, 0, time.UTC)),
        interval.New(time.Date(2020, 10, 18, 12, 0, 0, 0, time.UTC), time.Date(2020, 10, 18, 14, 30, 0, 0, time.UTC)),
        interval.New(time.Date(2020, 10, 18, 14, 33, 0, 0, time.UTC), time.Date(2020, 10, 18, 18, 0, 0, 0, time.UTC)),
    )
    intervalInput := interval.New(time.Date(2020, 10, 18, 14, 0, 0, 0, time.UTC), time.Date(2020, 10, 18, 15, 0, 0, 0, time.UTC))
    exceptMany := intervalMany.Except(intervalInput)

    // [
    //    2020-10-18 07:00:00 +0000 UTC - 2020-10-18 10:00:00 +0000 UTC
    //    2020-10-18 12:00:00 +0000 UTC - 2020-10-18 14:00:00 +0000 UTC
    //    2020-10-18 15:00:00 +0000 UTC - 2020-10-18 18:00:00 +0000 UTC
    // ]

    fmt.Println(exceptMany.String())         				    			     
}
```
* Union

for SpanMany union operation concatenates and sorts the original result
```go
package main
 
import (
    "fmt"
    "time"

    interval "github.com/go-follow/time-interval"
)

func main() {
    timeStart1 := time.Date(2020, 10, 17, 12, 0, 0, 0, time.UTC)
    timeEnd1 := time.Date(2020, 10, 17, 14, 0, 0, 0, time.UTC)
    timeStart2 := time.Date(2020, 10, 17, 22, 0, 0, 0, time.UTC)
    timeEnd2 := time.Date(2020, 10, 17, 23, 0, 0, 0, time.UTC)
    timeStart3 := time.Date(2020, 10, 17, 13, 0, 0, 0, time.UTC)
    timeEnd3 := time.Date(2020, 10, 17, 17, 0, 0, 0, time.UTC)
    timeStart4 := time.Date(2020, 10, 17, 7, 0, 0, 0, time.UTC)
    timeEnd4 :=  time.Date(2020, 10, 17, 10, 0, 0, 0, time.UTC)
    timeStart5 := time.Date(2020, 10, 17, 21, 0, 0, 0, time.UTC)
    timeEnd5 := time.Date(2020, 10, 17, 23, 0, 0, 0, time.UTC)
    timeStart6 := time.Date(2020, 10, 17, 11, 0, 0, 0, time.UTC)
    timeEnd6 := time.Date(2020, 10, 17, 15, 0, 0, 0, time.UTC)
    
    timeSpan := interval.New(timeStart1, timeEnd1)
    result := timeSpan.Union(interval.New(timeStart3, timeEnd3))

    // [
    //	2020-10-17 12:00:00 +0000 UTC - 2020-10-17 17:00:00 +0000 UTC
    // ]
    fmt.Println(result.String())

    timeSpanMany1 := interval.NewMany(
        timeSpan,
        interval.New(timeStart2, timeEnd2),
        interval.New(timeStart3, timeEnd3),
    )
    timeSpanMany2 := interval.NewMany()
    timeSpanMany2.Add(timeStart4, timeEnd4)
    timeSpanMany2.AddMany(interval.New(timeStart5, timeEnd5), interval.New(timeStart6, timeEnd6))

    resultMany := timeSpanMany1.Union(timeSpanMany2)

    // [
    //	2020-10-17 07:00:00 +0000 UTC - 2020-10-17 10:00:00 +0000 UTC
    //	2020-10-17 11:00:00 +0000 UTC - 2020-10-17 17:00:00 +0000 UTC
    //	2020-10-17 21:00:00 +0000 UTC - 2020-10-17 23:00:00 +0000 UTC
    // ]

    fmt.Println(resultMany.String())  
        				    			     
}
```
* Intersection
```go
package main
 
import (
    "fmt"
    "time"

    interval "github.com/go-follow/time-interval"
)

func main() {
    timeStart1 := time.Date(2020, 10, 17, 7, 0, 0, 0, time.UTC)
    timeEnd1 := time.Date(2020, 10, 17, 12, 0, 0, 0, time.UTC)
    timeStart2 := time.Date(2020, 10, 17, 10, 0, 0, 0, time.UTC)
    timeEnd2 := time.Date(2020, 10, 17, 15, 0, 0, 0, time.UTC)
    timeStart3 := time.Date(2020, 10, 17, 20, 0, 0, 0, time.UTC)
    timeEnd3 := time.Date(2020, 10, 17, 22, 0, 0, 0, time.UTC)
    timeStart4 := time.Date(2020, 10, 17, 18, 0, 0, 0, time.UTC)
    timeEnd4 :=	time.Date(2020, 10, 17, 23, 0, 0, 0, time.UTC)
    timeStart5 := time.Date(2020, 10, 17, 7, 0, 0, 0, time.UTC)
    timeEnd5 := time.Date(2020, 10, 17, 10, 0, 0, 0, time.UTC)

    newInt := interval.New(timeStart1, timeEnd1)
    newInt2 := interval.New(timeStart2, timeEnd2)

    result := newInt.Intersection(newInt2)
    fmt.Println(result.String()) // 2020-10-17 10:00:00 +0000 UTC - 2020-10-17 12:00:00 +0000 UTC

    timeStartInput := time.Date(2020, 10, 17, 10, 0, 0, 0, time.UTC)
    timeEndInput := time.Date(2020, 10, 17, 19, 0, 0, 0, time.UTC)
    intervalInput := interval.New(timeStartInput, timeEndInput)
    timeSpanMany := interval.NewMany(
        interval.New(timeStart1, timeEnd1),
        interval.New(timeStart2, timeEnd2),
        interval.New(timeStart3, timeEnd3),
        interval.New(timeStart4, timeEnd4),
        interval.New(timeStart5, timeEnd5),
    )
    resultMany := timeSpanMany.Intersection(intervalInput)
    // [
    //	2020-10-17 10:00:00 +0000 UTC - 2020-10-17 12:00:00 +0000 UTC
    //	2020-10-17 10:00:00 +0000 UTC - 2020-10-17 15:00:00 +0000 UTC
    //	2020-10-17 18:00:00 +0000 UTC - 2020-10-17 19:00:00 +0000 UTC
    // ]
    fmt.Println(resultMany.String())   
}
```
* Equal

It is possible to pass an optional argument offset, which gives the possibility of a small error in the final result
```go
package main
 
import (
    "fmt"
    "time"

    interval "github.com/go-follow/time-interval"
)

func main() {
    timeStart1 := time.Date(2020, 10, 18, 15, 0, 0, 0, time.UTC)
    timeEnd1 := time.Date(2020, 10, 18, 21, 0, 0, 0, time.UTC)
    timeStart2 := time.Date(2020, 10, 18, 14, 50, 0, 0, time.UTC)
    timeEnd2 := time.Date(2020, 10, 18, 21, 10, 0, 0, time.UTC)
    ti1 := interval.New(timeStart1, timeEnd1)
    ti2 := interval.New(timeStart2, timeEnd2)
    ti2AddSecond := interval.New(timeStart2, timeEnd2.Add(1 * time.Second))
    // Equal without offset
    fmt.Println(ti1.Equal(ti2)) // false    
    // Equal with offset 10 minute
    fmt.Println(ti1.Equal(ti2, time.Minute * 10)) // true
    // Add 1 second to ti2
    fmt.Println(ti1.Equal(ti2AddSecond, time.Minute * 10)) // false        
    
    // Equal for SpanMany
    // If there is at least one match, return true
    intervalMany := interval.NewMany(
        interval.New(time.Date(2020, 10, 18, 9, 0, 0, 0, time.UTC), time.Date(2020, 10, 18, 10, 0, 0, 12, time.UTC)),
        interval.New(time.Date(2020, 10, 18, 19, 0, 0, 0, time.UTC), time.Date(2020, 10, 18, 20, 0, 0, 0, time.UTC)),
        interval.New(time.Date(2020, 10, 18, 16, 55, 0, 0, time.UTC), time.Date(2020, 10, 18, 18, 0, 0, 11, time.UTC)),
    )
    // Equal without offset
    intervalInput := interval.New(time.Date(2020, 10, 18, 17, 0, 0, 0, time.UTC), time.Date(2020, 10, 18, 18, 5, 0, 11, time.UTC))
    fmt.Println(intervalMany.Equal(intervalInput)) // false
    // Equal with offset
    fmt.Println(intervalMany.Equal(intervalInput, time.Minute * 5)) // true
    fmt.Println(intervalMany.Equal(intervalInput, time.Minute * 4)) // false      				    			     
}
```
* IsIntersection

It is possible to pass an optional argument offset, which gives the possibility of a small error in the final result
```go
package main
 
import (
    "fmt"
    "time"

    interval "github.com/go-follow/time-interval"
)

func main() {
    timeStart1 := time.Date(2020, 10, 18, 17, 30, 0, 0, time.UTC)
    timeEnd1 := time.Date(2020, 10, 18, 18, 22, 0, 0, time.UTC)
    timeStart2 := time.Date(2020, 10, 18, 16, 0, 0, 0, time.UTC)
    timeEnd2 := time.Date(2020, 10, 18, 17, 30, 7, 0, time.UTC)
    ti1 := interval.New(timeStart1, timeEnd1)
    ti2 := interval.New(timeStart2, timeEnd2)
    // IsIntersection without offset
    fmt.Println(ti1.IsIntersection(ti2)) // true
    // IsIntersection with offset 5 second
    fmt.Println(ti1.IsIntersection(ti2, time.Second * 5)) // true
    // IsIntersection with offset 10 second
    fmt.Println(ti1.IsIntersection(ti2, time.Second * 10)) // false        
    
    // IsIntersection for SpanMany
    // If there is at least one match, return true
    intervalMany := interval.NewMany(
        interval.New(time.Date(2020, 10, 18, 9, 0, 0, 0, time.UTC), time.Date(2020, 10, 18, 10, 0, 0, 0, time.UTC)),
        interval.New(time.Date(2020, 10, 18, 16, 0, 0, 0, time.UTC), time.Date(2020, 10, 18, 17, 5, 0, 0, time.UTC)),
        interval.New(time.Date(2020, 10, 18, 19, 0, 0, 0, time.UTC), time.Date(2020, 10, 18, 20, 0, 0, 0, time.UTC)),
    )
    // IsIntersection without offset
    intervalInput := interval.New(time.Date(2020, 10, 18, 17, 0, 0, 0, time.UTC), time.Date(2020, 10, 18, 18, 0, 0, 0, time.UTC))
    fmt.Println(intervalMany.IsIntersection(intervalInput)) // true
    // IsIntersection with offset
    fmt.Println(intervalMany.IsIntersection(intervalInput, time.Minute * 5)) // false
    fmt.Println(intervalMany.IsIntersection(intervalInput, time.Minute * 3)) // true    				    			     
}


