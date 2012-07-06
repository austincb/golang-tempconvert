// Converts temperatures from Fahrenheit -> Celsius
// Or Celsius -> Fahrenheit
// Fahrenheit -> Celsius by default, otherwise use "c" flag for Celsius -> Fahrenheit
package main
     
    import (
            "fmt"
            "flag"
    )
     
    var Celsius = flag.Bool("c", false, "convert from Celsius to Fahrenheit") // set a boolean flag for celsius conversions
     
    // fahrenheit -> celsius conversion
    func convertf(sum float32) (y float32) {
            x := sum - 32
            y = x * 0.555555556
            return y
    }
    
    // celsius -> fahrenheit conversion
    func convertc(sum float32) (y float32) { 
            x := sum * 1.8
            y = x + 32
            return y
    }
     
    // check for the c flag
    func main(){
            flag.Parse()
            if flag.NArg() == 1 {
                    // flag set, convert to Fahrenheit
                    var l float32
                    fmt.Printf("Enter the degrees in Celsius: ")
                    fmt.Scanf("%g", &l)
                    m := convertc(l)
                    fmt.Printf("%g\n", m)
            }
            if !*Celsius {
	            // no flag set, convert to Celsius
                    var l float32
                    fmt.Printf("Enter the degrees in Fahrenheit: ")
                    fmt.Scanf("%g°\n", &l)
                    m := convertf(l)
                    fmt.Printf("%g\n", m)
            }
    }
