package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var stringVar string
	var Rimskie, RimskFirst, RimskLast bool

	MassivRimskix := []string{
		"I",
		"II",
		"III",
		"IV",
		"V",
		"VI",
		"VII",
		"VIII",
		"IX",
		"X",
		"XI",
		"XII",
		"XIII",
		"XIV",
		"XV",
		"XVI",
		"XVII",
		"XVIII",
		"XIX",
		"XX",
		"XXI",
		"XXII",
		"XXIII",
		"XXIV",
		"XXV",
		"XXVI",
		"XXVII",
		"XXVIII",
		"XXIX",
		"XXX",
		"XXXI",
		"XXXII",
		"XXXIII",
		"XXXIV",
		"XXXV",
		"XXXVI",
		"XXXVII",
		"XXXVIII",
		"XXXIX",
		"XL",
		"XLI",
		"XLII",
		"XLIII",
		"XLIV",
		"XLV",
		"XLVI",
		"XLVII",
		"XLVIII",
		"XLIX",
		"L",
		"LI",
		"LII",
		"LIII",
		"LIV",
		"LV",
		"LVI",
		"LVII",
		"LVIII",
		"LIX",
		"LX",
		"LXI",
		"LXII",
		"LXIII",
		"LXIV",
		"LXV",
		"LXVI",
		"LXVII",
		"LXVIII",
		"LXIX",
		"LXX",
		"LXXI",
		"LXXII",
		"LXXIII",
		"LXXIV",
		"LXXV",
		"LXXVI",
		"LXXVII",
		"LXXVIII",
		"LXXIX",
		"LXXX",
		"LXXXI",
		"LXXXII",
		"LXXXIII",
		"LXXXIV",
		"LXXXV",
		"LXXXVI",
		"LXXXVII",
		"LXXXVIII",
		"LXXXIX",
		"XC",
		"XCI",
		"XCII",
		"XCIII",
		"XCIV",
		"XCV",
		"XCVI",
		"XCVII",
		"XCVIII",
		"XCIX",
		"C",
	}

	regAllOpers := regexp.MustCompile(`[+-/*]`)

	fmt.Println("Допустимые цифры: 0 1 2 3 5 6 7 8 9")
	fmt.Println("Допустимые римкие цифры: I II III IV V VI VII VIII IX X")
	fmt.Println("Допустимые римские символы: I V X")

	fmt.Println("Введите пример в одну строку")

	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		stringVar = scanner.Text()
	}

	stringVar = strings.ReplaceAll(stringVar, " ", "")

	if len(regAllOpers.FindAllString(stringVar, -1)) > 1 {
		fmt.Println("Формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
		os.Exit(4000)

	} else if len(regAllOpers.FindAllString(stringVar, -1)) == 1 {
		split := regAllOpers.FindStringSubmatch(stringVar)

		regVar := regAllOpers.Split(stringVar, -1)
		regVarRimskie := regAllOpers.Split(stringVar, -1)

		FirstVar, err := strconv.Atoi(regVar[0])
		LastVar, err := strconv.Atoi(regVar[1])

		for i, v := range MassivRimskix {
			if v == regVarRimskie[0] {
				RimskFirst = true
				FirstVar = i + 1
			}
		}

		for i, v := range MassivRimskix {
			if v == regVarRimskie[1] {
				RimskLast = true
				LastVar = i + 1
			}
		}

		if RimskFirst == true && RimskLast == true {
			Rimskie = true

		} else if RimskFirst == true && RimskLast == false {
			if _, err := strconv.Atoi(regVar[1]); err != nil {
				fmt.Printf("%q не число.\n", regVar[1])
				os.Exit(402)
			}

			fmt.Println("Используются одновременно разные системы счисления.")
			os.Exit(101)

		} else if RimskFirst == false && RimskLast == true {
			if _, err := strconv.Atoi(regVar[0]); err != nil {
				fmt.Printf("%q не число.\n", regVar[0])
				os.Exit(401)
			}
			fmt.Println("Используются одновременно разные системы счисления.")
			os.Exit(101)
		}

		if (len(regVar) > 2) || (regVar[0] == "") || (regVar[1] == "") {
			fmt.Println("Формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
			os.Exit(404)
		}

		if err != nil && Rimskie == false {
			panic("Строка не является математической операцией.")
		}

		if split[0] == "/" {
			if FirstVar <= 10 && LastVar <= 10 && FirstVar != 0 && LastVar != 0 {

				if FirstVar-LastVar < 1 && Rimskie == true {
					fmt.Println("В римской системе нет отрицательных чисел и нуля.")
					os.Exit(-1)
				} else {

					if Rimskie == true {
						Otvet := FirstVar / LastVar
						for i, v := range MassivRimskix {
							if i == Otvet-1 {
								fmt.Println(v)
								RimskFirst = true
								FirstVar = i + 1
							}
						}

					} else {
						fmt.Printf("%d", FirstVar/LastVar)
					}
				}
			} else {
				if FirstVar == 0 || LastVar == 0 {
					fmt.Println("Одно из чисел меньше 1")

				} else {
					fmt.Println("Одно из чисел больше 10")
				}
			}

		} else if split[0] == "*" {
			if FirstVar <= 10 && LastVar <= 10 && FirstVar != 0 && LastVar != 0 {
				if Rimskie == true {
					Otvet := FirstVar * LastVar

					for i, v := range MassivRimskix {
						if i == Otvet-1 {
							fmt.Println(v)
							RimskFirst = true
							FirstVar = i + 1
						}
					}

				} else {
					fmt.Printf("%d", FirstVar*LastVar)
				}

			} else {
				if FirstVar == 0 || LastVar == 0 {
					fmt.Println("Одно из чисел меньше 1")

				} else {
					fmt.Println("Одно из чисел больше 10")
				}
			}

		} else if split[0] == "+" {
			if FirstVar <= 10 && LastVar <= 10 && FirstVar != 0 && LastVar != 0 {
				if Rimskie == true {
					Otvet := FirstVar + LastVar
					// fmt.Println(Otvet)
					for i, v := range MassivRimskix {
						if i == Otvet-1 {
							fmt.Println(v)
							RimskFirst = true
							FirstVar = i + 1
						}
					}

				} else {
					fmt.Printf("%d", FirstVar+LastVar)
				}

			} else {
				if FirstVar == 0 || LastVar == 0 {
					fmt.Println("Одно из чисел меньше 1")

				} else {
					fmt.Println("Одно из чисел больше 10")
				}
			}

		} else if split[0] == "-" {
			if FirstVar <= 10 && LastVar <= 10 && FirstVar != 0 && LastVar != 0 {
				if FirstVar-LastVar < 1 && Rimskie == true {
					fmt.Println("В римской системе нет отрицательных чисел и нуля.")
					os.Exit(-1)

				} else {
					if Rimskie == true {
						Otvet := FirstVar - LastVar
						for i, v := range MassivRimskix {
							if i == Otvet-1 {
								fmt.Println(v)
								RimskFirst = true
								FirstVar = i + 1
							}
						}

					} else {
						fmt.Printf("%d", FirstVar-LastVar)
					}
				}

			} else {
				if FirstVar == 0 || LastVar == 0 {
					fmt.Println("Одно из чисел меньше 1")

				} else {
					fmt.Println("Одно из чисел больше 10")
				}
			}
		}
	} else if len(regAllOpers.FindAllString(stringVar, -1)) == 0 {
		fmt.Println("Строка не является математической операцией.")
		os.Exit(404)
	}
}
