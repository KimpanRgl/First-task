package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
func main() {
	reader :=bufio.NewReader(os.Stdin)
	romeToArab := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":   10,
	}
	arabToRome := map[int]string{
		1:        "I",
		2:       "II",
		3:      "III",
		4:       "IV",
		5:        "V",
		6:       "VI",
		7:      "VII",
		8:     "VIII",
		9:       "IX",
		10:       "X",
		11:      "XI",
		12:     "XII",
		13:    "XIII",
		14:     "XIV",
		15:      "XV",
		16:     "XVI",
		17:    "XVII",
		18:   "XVIII",
		19:     "XIX",
		20:      "XX",
		21:     "XXI",
		24:    "XXIV",
		25:     "XXV",
		27:   "XXVII",
		28:  "XXVIII",
	 	30:     "XXX",
		32:   "XXXII",
		35:    "XXXV",
		36:   "XXXVI",
	 	40:      "XL",
		42:    "XLII",
		45:     "XLV",
		48:  "XLVIII",
		49:    "XLIX",
	 	50:       "L",
		54:     "LIV",
		56:     "LVI",
		60:      "LX",
		63:   "LXIII",
		64:    "LXIV",
		70:     "LXX",
		72:   "LXXII",
		80:    "LXXX",
		81:   "LXXXI",
		90:      "XC",
		100:      "C",
	}
	for {
		fmt.Println("Введите значение")
		text, _ := reader.ReadString('\n')
		array := strings.Fields(text)
		if len(array)==1 {
			switch array[0] {
			case "1","2","3","4","5","6","7","8","9","10","I","II","III","IV","V","VI","VII","VIII","IX","X":
				panic(fmt.Errorf("Выдача паники, так как строка не является математической операцией."))
			}
		}
		if len(array)==3 {
			switch array[0] {
			case "1","2","3","4","5","6","7","8","9","10":
				switch array[2] {
				case "1","2","3","4","5","6","7","8","9","10":
					a, _ := strconv.Atoi(array[0])             
					b, _ := strconv.Atoi(array[2])
					switch array[1] {
						case "+":
							fmt.Println(a+b)
						case "-":
							fmt.Println(a-b)
						case "*":
							fmt.Println(a*b)
						case "/":
							fmt.Println(a/b)
						default:
							panic(fmt.Errorf("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)."))
					}
				default:
					panic(fmt.Errorf("Выдача паники, так как используются одновременно разные системы счисления."))

				}
			case "I","II","III","IV","V","VI","VII","VIII","IX","X":
				switch array[2] {
				case "I","II","III","IV","V","VI","VII","VIII","IX","X":
					switch array[1] {
						case "+":
							fmt.Println(arabToRome[romeToArab[array[0]] + romeToArab[array[2]]])
						case "-":
							res := romeToArab[array[0]] - romeToArab[array[2]]
							if res == 0 {
								panic(fmt.Errorf("Выдача паники, так как в римской системе нет числа 0."))
							} else if res < 0 {
								panic(fmt.Errorf("Выдача паники, так как в римской системе нет отрицательных чисел."))
							} else {
								fmt.Println(arabToRome[res])	
							}
						case "*":
							fmt.Println(arabToRome[romeToArab[array[0]] * romeToArab[array[2]]])
						case "/":
							res := romeToArab[array[0]] / romeToArab[array[2]]
							if res == 0 {
								panic(fmt.Errorf("Выдача паники, так как в римской системе нет числа 0."))
							} else {
								fmt.Println(arabToRome[res])
							}
						default:
							panic(fmt.Errorf("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)."))
						}	
				default:
					panic(fmt.Errorf("Выдача паники, так как используются одновременно разные системы счисления."))
				}
			default:
				panic(fmt.Errorf("Выдача паники, так как были введены неподходящие числа"))
			}		
		} else {
			panic(fmt.Errorf("Выдача паники, так как формат математической операции не удовлетворяет заданию: два операнда и один оператор (+, -, /, *), калькулятор должен принимать на вход числа от 1 до 10."))
		}
	}
}