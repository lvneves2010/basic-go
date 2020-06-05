package main

import "fmt"

func main() {
  var stationName string
  var nextTrainTime int8
  var isUptownTrain bool
  
  stationName = "Union Square"
  nextTrainTime = 12
  isUptownTrain = false
  
  fmt.Println("Current station:", stationName)
  fmt.Println("Next train:", nextTrainTime, "minutes")
  fmt.Println("Is uptown:", isUptownTrain)
  
  stationName = "Grand Central"
  nextTrainTime = 3
  isUptownTrain = true
  
  fmt.Println("")
  fmt.Println("Current station:", stationName)
  fmt.Println("Next train:", nextTrainTime, "minutes")
  fmt.Println("Is uptown:", isUptownTrain)

  // Create the constant earthsGravity
  // here and set it to 9.80665
  const earthsGravity = 9.80665
  
  // Here's where we print out the gravity:
  fmt.Println(earthsGravity)


  var numOfFlavors int
  // Assign a value for numOfFlavors below:
  numOfFlavors = 57
  
  fmt.Println(numOfFlavors)
  
  // Declare flavorScale below:
  var flavorScale float32 = 5.8

  fmt.Println(flavorScale)

  // Define a string variable
  // called favoriteSnack
  var favoriteSnack string 
  
  // Assign a value to
  // favoriteSnack
  favoriteSnack = "popcorn"
  
  // Print out the message
  // "My favorite snack is "
  // followed by the value in
  // favoriteSnack
  fmt.Println("My favorite snack is " + favoriteSnack)


  // Create three variables
  // emptyInt an int8
  var emptyInt int8
  // emptyFloat a float32
  var emptyFloat float32
  // and emptyString a string
  var emptyString string
  // Finally, print them all out
  fmt.Println(emptyInt, emptyFloat, emptyString)


  // Define daysOnVacation using := below:
  daysOnVacation := 7
  
  // Define hoursInDay using var and = below:
  var hoursInDay = 24
  
  fmt.Println("You have spent", daysOnVacation * hoursInDay, "hours on vacation.")

  // Define cupsOfCoffeeConsumed here
  var cupsOfCoffeeConsumed int
  // Give a value to cupsOfCoffeeConsumed
  cupsOfCoffeeConsumed = 4
  // Print out cupsOfCoffeeConsumed
  fmt.Println(cupsOfCoffeeConsumed)

  coolSneakers := 65.99
  niceNecklace := 45.50
  
  // Define taxCalculation here
  var taxCalculation float64
  // Add coolSneakers to taxCalculation
  taxCalculation += coolSneakers
  // Add niceNecklace to taxCalculation
  taxCalculation += niceNecklace
  // Compute the NYC sales tax
  // 8.875% of the purchase here:
  taxCalculation *= .08875
  // Uncomment this line for a receipt!
  //fmt.Println("Purchase of", coolSneakers + niceNecklace, "with 8.875% sales tax", taxCalculation, "equal to", coolSneakers + niceNecklace + taxCalculation)

  // Define magicNum and powerLevel below:
  var magicNum, powerLevel int32
  magicNum, powerLevel = 2048, 9001
  
  fmt.Println("magicNum is:", magicNum, "powerLevel is:", powerLevel)
  
  
  // Define amount and unit below:
  amount, unit := 10, "doll hairs"
  
  fmt.Println(amount, unit, ", that's expensive...")


	var congrats string
  congrats = "Congratulations"
  congrats += " Léo"
  congrats += "!!!"
  fmt.Println(congrats)
  
  var challenge string = "What else can you do?"
  fmt.Println(challenge)
  
  reminder := "Pratice is important!"
  fmt.Println(reminder)
}