package main
import (
	"fmt"
	"errors"
)

type div_err struct{
	dividend float64
	divisor float64
}
func (d *div_err) Error() string {
	return fmt.Sprintf("cannot divide %.2f by %.2f",d.dividend, d.divisor)
}
func divide (a, b float64) (float64,error){
	if(b==0){
		return 0,&div_err{dividend: a, divisor: b}

	
	}
	return a/b,nil
} 
func main() {
	output, err := divide(2,0)
	if(err!=nil) {
		fmt.Println(err)
		var ae *div_err
		// doubt
		if de, ok := err.(*div_err); ok {
			// If ok is true, then err is of type *DivideError
			fmt.Printf("Dividend: %.2f, Divisor: %.2f\n", de.dividend, de.divisor)
		}
		if errors.As(err, &ae) {
			fmt.Printf("Dividend1: %.2f, Divisor1: %.2f\n", ae.dividend, ae.divisor)

		}
		
	} else{
		fmt.Println(output)
	}
}