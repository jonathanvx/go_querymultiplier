package main

import (
    "fmt"
    "os"
    "math"
    "flag"
)

func main() {
flags := flag.NewFlagSet("",flag.ExitOnError)
querystring := flags.String("",os.Args[1],"The desired query to multiply")
rows := flags.Int("rows", 100, "The rows or chunk size.") 
minimum := flags.Int("min", 1, "The position you want the rows to start from.")
maximum := flags.Int("max", 100, "The maxiumum rows to change or the final id on the column you want.")
keytype := flags.String("keytype", "limit", "Enter 'limit' or a specific column name like 'id'")
flags.Parse(os.Args[2:])

looper := int(math.Ceil(float64(*maximum-*minimum+1)/float64(*rows)))

    for i := 0; i < looper; i++ {
        if ( *keytype == "LIMIT" ) || ( *keytype == "limit") {
                if ( i == looper-1) {
                        fmt.Printf("%s LIMIT %d,%d;\n", *querystring, (*rows*i)+*minimum-1, *maximum-((*rows*i)+*minimum)+1)
                } else {
                        fmt.Printf("%s LIMIT %d,%d;\n", *querystring, (*rows*i)+*minimum-1, *rows) 
                }
        } else {
                if ( i == looper-1) {
                        fmt.Printf("%s WHERE %s between %d and %d;\n", *querystring, *keytype, (*rows*i)+*minimum, *maximum)
                } else {
                        fmt.Printf("%s WHERE %s between %d and %d;\n", *querystring, *keytype, (*rows*i)+*minimum, *rows*(i+1)+*minimum-1)
                }
        }
        }
}
