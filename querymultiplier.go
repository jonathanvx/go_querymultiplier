ackage main

import (
    "fmt"
    "os"
    "strconv"
    "math"
)

func main() {
//var limit_rows, limit_total int
querystring := os.Args[1]
limit_rows, _ := strconv.Atoi(os.Args[2])
limit_total, _ := strconv.Atoi(os.Args[3])
looper := int(math.Ceil(float64(limit_total)/float64(limit_rows)))

    for i := 0; i < looper; i++ {
        if ( os.Args[4] == "LIMIT" ) || ( os.Args[4] == "limit") {
                if ( i == looper-1) {
                        fmt.Printf("%s LIMIT %d,%d;\n", querystring, limit_rows*i, limit_total-(limit_rows*i))
                } else {
                        fmt.Printf("%s LIMIT %d,%d;\n", querystring, limit_rows*i, limit_rows) 
                }
        } else {
                if ( i == looper-1) {
                        fmt.Printf("%s WHERE %s between %d and %d;\n", querystring, os.Args[4], (limit_rows*i)+1, limit_total)
                } else {
                        fmt.Printf("%s WHERE %s between %d and %d;\n", querystring, os.Args[4], (limit_rows*i)+1, limit_rows*(i+1))
                }
        }
        }
}
