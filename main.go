/*
    Created By FurherBino AKA FurherAcker AKA MrHive AKA MrH1V3 AKA AckerPosiden AKA dev17V
    My Old Github: https://github.com/dev17V
    My New Github: https://github.com/dev17V2

    My First Official Golang Code Practice Wish Me Luck.

    Date Of Practice: 08/17/2023
*/

package main

import (
    "time" // for importing the local/date/current/sec/hrs etc module
    "fmt" //  // Nickname: Formating Strings Module
    //"os/exec" //os for importing the Operating System Module | exec for excutable system commands
    "os"
)

func main() {

    //array of integers with a fixed size of 5
    var myNumbers [5]int

    myNumbers[0] = 1 //all data starts from 0
    myNumbers[1] = 2
    myNumbers[2] = 3
    myNumbers[3] = 4
    myNumbers[4] = 5

    var friendsNumbers [5]int //sorry my folks old lol

    friendsNumbers[0] = 5 //all data starts from 0
    friendsNumbers[1] = 4
    friendsNumbers[2] = 3
    friendsNumbers[3] = 2
    friendsNumbers[4] = 1



    //////////////////////////////////////////
    /*
        side note Print is normal on one line/ Printf is for formating prints 
        putting(or/calling) varibles inside the string / 
        Println auto moves to the next line after your string ends
    */

    //My Variables UwU
    var name string
    var myAge int

    // String Input
    fmt.Print("Nickname: ")
    fmt.Scanln(&name)

    // int Input
    fmt.Print("Age: ")
    fmt.Scanln(&myAge)

    //%s == string // %d == data // 

    // Output
    fmt.Printf("Welcome To The Hive, %s!\nWe heard you just turned %d years old!\nAnd according to doxbin you won the lottery with the lucky numbers [%d]\n", name, myAge, myNumbers)

    friendName := os.Args[0] //all data starts from 0
    friendAge := os.Args[1]

    err := UnrelatedError

    err_one := fmt.Errorf("To Many Args. | Usage ./main friendName friendAge | Example: ./main Skid 24") //Phrase Our Own Error Message

    err_two := fmt.Errorf("To Many Args. | Usage ./main friendName friendAge | Example: ./main Skid 24") //Phrase Our Own Error Message

    //if statement
    if (os.Args[1] == " ") {
        fmt.Println("[%s] is not a valid argument", os.Args[1])
    } else if len(os.Args) < 2 { //If Greater Than Two Arguments /// side note when comparing os.Args use "len" for the length if len(os.Args) etc..///
        fmt.Println(err_one) //print on the next avalible line our custom err_one()
    } else if len(os.Args) > 2 { //If Less Than Two Arguments
        fmt.Println(err_two) //print on the next avalible line our custom err_two()
    } else if err != nil {
        fmt.Println("Error:", err)
    } else {
        //cmd := exec.Command("clear") //"ls", "-l" or one command
        //// Run the command and capture its output
        fmt.Println("YOUR A FUCKING L33T HaCkEr BrO.")
        time.Sleep(2 * time.Second) //FUCKING AWESOME VERY MUCH LIKE V-Lang UwU
    }
    //Output
    fmt.Println("Your Friends Name: %s\nYour Friends Age: %d\nYour Friends Lottery Numbers: %d", friendName, friendAge, friendsNumbers)
}

func UnrelatedError() {
    //return fmt.Errorf("Something went wrong, But only because your a fucking skid!") //lol HIVEWASHERE
}
