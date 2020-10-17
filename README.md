# time-interval [![license](https://img.shields.io/badge/license-MIT-green.svg?style=flat)](https://github.com/go-follow/time-interval/blob/master/LICENSE)

This package helps to work with time intervals. The functionality allows you to perform the following basic operations:
* Union - merging of time intervals
* Intersection - finding intersections of regular intervals
* Except - differentiation between fixed intervals
* Equal - comparison of time intervals

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
* Union
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


