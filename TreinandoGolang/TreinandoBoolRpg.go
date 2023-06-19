package main

import "fmt"

func CanFastAttack(awake bool) bool {
	if awake == true {
		return false
	} else {
		return true
	}
}

func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {

	return knightIsAwake || archerIsAwake || prisonerIsAwake //they are already bool, so if any of those variables turn into "false", the hole return will be false
}

func CanSignalPrisoner(prisonerIsAwake, archerIsAwake bool) bool {
	if prisonerIsAwake == true && archerIsAwake == false {
		return true
	} else {
		return false
	}
}

func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
	if petDogIsPresent == true && archerIsAwake == false {
		return true
	} else if petDogIsPresent == false {
		if prisonerIsAwake == true && archerIsAwake == false && (knightIsAwake == false || archerIsAwake == false) {
			return true
		}
	}

	return false
}

func main() {
	var knightIsAwake = false
	var archerIsAwake = true
	var prisonerIsAwake = false
	var petDogIsPresent = false
	fmt.Println(CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent))
	// Output: false
}
