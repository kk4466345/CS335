// return type error of function
package main;

import "fmt";

func gcd(a,b int) float{
	if b == 0{
		g := 10
		return a;
	}
	else{
		h :=10
	}

	return gcd(b,a%b);
};


func main() {
	var a,b int;
	scan a;
	scan b;
	var d int = gcd(a,b);

};