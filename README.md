# time-interval [![license](https://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://github.com/go-follow/time-interval/blob/master/LICENSE)

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
    fmt.Println(result.String())
}
```
