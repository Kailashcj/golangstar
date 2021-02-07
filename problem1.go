package main

import (
	"fmt"
)

//*** since the solution was needed in tuple and go does not have tuple,
//i will create below struct to hold the string and int pairs. Note
//that slice and arrays can only hold one type
type storer struct {
	chr   string
	count int
}

func main() {
	arr := "aaaabbbcccaaaakkppppppppppqqqqt"
  
  //***previously i defined this slice of slice. this is also iteresting as to how you define that.
	//result := [][]string{}  //define an array of arrray. notice the single {}
  
  //result is a slice of struct storer, which can hold string & int
	result := []storer{}. //initiazed as well. did not use make as i do not know the length

	if len(arr) < 1 {
		fmt.Println("")
	}
	i, j := 0, 0
	for j < len(arr) {
		if arr[i] != arr[j] {
      //*** can negotiate for a print statement if you find it difficult
      //to output it to result(of course you have time limit to solve this problem)
			fmt.Println(string(arr[i]), j-i)
			//result = append(result, []string{string(arr[i]), strconv.Itoa(j - i)})
			result = append(result, storer{string(arr[i]), j - i})
			i = j
		}
		j++
	}
  
  //*** nice edge case handling when you have last element left in the string. 
  //when j becomes equal or greater than the string length, what to do?
  //print ith char and use regular j-1 for the count of chars***
	if j >= len(arr) { 
		fmt.Println(string(arr[i]), j-i)
		//result = append(result, []string{string(arr[i]), strconv.Itoa(j - i)})
		result = append(result, storer{string(arr[i]), j - i})
	}
	fmt.Println(result) //just printing result
}
