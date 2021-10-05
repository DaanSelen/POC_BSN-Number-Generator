package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var (
	validBSN int
)

func main() {
	generateBSN() //Initiate the generation function.
	fmt.Println("The BSN Generated is: ", validBSN)
	fmt.Scanln() //This is here so the program doesn't immediatly close after running.
}

func generateBSN() {
	buildingBSN := ""                //Creating an empty string to build the BSN number within.
	rand.Seed(time.Now().UnixNano()) //Creating a random seed to base the random numbers on, in this case the time.
	for i := 0; i < 9; i++ {
		buildingBSN += strconv.Itoa(rand.Intn(10)) //Generate random numbers by default from 1 to 10
	}
	trialBSN, err := strconv.Atoi(buildingBSN) //Convert string to int or leave an error behind telling the user what's wrong.
	if err != nil {
		log.Fatal("Kon de string niet converteren naar een int", err)
	}
	validateBSN(buildingBSN, trialBSN) //Initiate the validation function with the generated string of numbers.
}

func validateBSN(buildingBSN string, trialBSN int) {
	splitBSN := strings.Split(buildingBSN, "")
	a := 9
	number01 := 0
	for _, x := range splitBSN {
		dummyvar, err := strconv.Atoi(x)
		if err != nil {
			log.Fatal("Converteren mislukt, functie ValidateBSN ", err)
		}
		if a != 1 {
			number01 += (dummyvar * a)
			a--
		} else if a == 1 {
			number01 += (dummyvar * -1)
		}
	}
	if number01%11 == 0 {
		fmt.Println("Valid BSN generated!")
		validBSN = trialBSN //Since validation succeeded variable changes to validBSN
	} else {
		fmt.Println("FAILED BSN VALIDATION, retrying!")
		generateBSN()
	}
}
