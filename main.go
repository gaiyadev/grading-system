package main

import "fmt"

func main()  {
	grade(10)
	text("obed")
	fruitArrr()
}

func grade(score int)  {
	//score := 69

	if score >=70 && score <=100 {
		fmt.Println("Excellent, You got: A",score)
	}else if score >=60 && score<70 {
		fmt.Println("Congratulation, You got: B",score)
	}else if score >= 50 && score <60 {
		fmt.Println("Weldone, You got: C",score)
	}else if score>=40 && score <50 {
		fmt.Println("Good, You got: D",score)
	}else if score>=35 && score <40 {
		fmt.Println("Nice attempt, You got: E", score)
	}else if score>=0 && score <35 {
		fmt.Println("Bad, You got: F",score)
	}else {
		fmt.Println("Invalid input",score)
	}

}

func text(word string)  {
	fmt.Println(len(word))
}

func fruitArrr()  {
	fruit := []int{1,2,4,5,6,7,8}
	fmt.Println(len(fruit))
}