/*
reads numbers from its command-line arguments or from 
the standard input if there are no arguments, 
and converts each number into units like temperature in 
Celsius and Fahrenheit, length in feet and meters, 
weight in pounds and kilograms, and the like.
*/

package main

import (
	"bufio"
	"fmt"
	"golang/chapterTwo/tempconv"
	"golang/chapterTwo/massconv"
	"golang/chapterTwo/lenconv"
	"os"
	"strconv"
)

func throwError(msg string, err error) {
	fmt.Fprintf(os.Stderr, msg, err, "\n")
	os.Exit(1)
}

func readFromStdin() (float64, string) {
	input := bufio.NewScanner(os.Stdin)
	fmt.Println("please input number for conversion")
	var number float64
	var unitType string
	var err error
	for i := 0; input.Scan(); {
		if i == 0 { 
			if number, err = strconv.ParseFloat(input.Text(), 64); err != nil {
				throwError("error parsing float:", err)
			}
			fmt.Println("please input unit type ('t' for temperature," +
				"'m' for mass, 'l' for length)")
		}
		if i == 1 {
			unitType = input.Text()
			break
		}
		i++
	}
	return number, unitType
}

func tempConvert(t float64) {
	f := tempconv.Fahrenheit(t)
	c := tempconv.Celsius(t)
	k := tempconv.Kelvin(t)
	fmt.Printf("%s = %s, %s = %s\n%s = %s, %s = %s\n",
		f, tempconv.FToC(f), c, tempconv.CToF(c), 
		k, tempconv.KToC(k), c, tempconv.CToK(c))
}

func massConvert(m float64) {
	kg := massconv.Kilogram(m)
	p := massconv.Pound(m)
	fmt.Printf("%s = %s, %s = %s\n",
		kg, massconv.KgToLb(kg), p, massconv.LbToKg(p))
}
func lenConvert(l float64) {
	km := lenconv.Kilometer(l)
	m := lenconv.Mile(l)
	fmt.Printf("%s = %s, %s = %s\n",
		km, lenconv.KmToM(km), m, lenconv.MToKm(m))
}

func main() {
	var number float64
	var unitType string
	var err error
    if len(os.Args) != 3 {
		number, unitType = readFromStdin()
	} else {
		number, err = strconv.ParseFloat(os.Args[1], 64) 
		if err != nil { throwError("error parsing float:", err) }
		unitType = os.Args[2]
	}

	switch {
	case unitType == "t":
		tempConvert(number)
	case unitType == "m":
		massConvert(number)
	case unitType == "l":
		lenConvert(number)
	default:
		fmt.Printf("unit type input = '%s'. Please input 't' for temperature," +
				"'m' for mass, or 'l' for length", unitType)
	}
}
