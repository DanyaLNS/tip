package main

import (
	"fmt"
	"math"
)

// 1. Сумма цифр числа
func sumOfDigits(n int) int {
	sum := 0
	for n != 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

// 2. Преобразование температуры
func celsiusToFahrenheit(celsius float64) float64 {
	return celsius*9/5 + 32
}

// 3. Удвоение каждого элемента массива
func doubleArray(arr []int) []int {
	result := make([]int, len(arr))
	for i, v := range arr {
		result[i] = v * 2
	}
	return result
}

// 4. Объединение строк
func joinStrings(words []string) string {
	result := ""
	for i, word := range words {
		result += word
		if i < len(words)-1 {
			result += " "
		}
	}
	return result
}

// 5. Расчет расстояния между двумя точками
func distance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}

// 2.1 Проверка на четность/нечетность
func isEven(n int) string {
	if n%2 == 0 {
		return "Четное"
	}
	return "Нечетное"
}

// 2.2 Проверка високосного года
func isLeapYear(year int) string {
	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		return "Високосный"
	}
	return "Невисокосный"
}

// 2.3 Определение наибольшего из трех чисел
func maxOfThree(a, b, c int) int {
	if a >= b && a >= c {
		return a
	} else if b >= a && b >= c {
		return b
	}
	return c
}

// 2.4 Категория возраста
func ageCategory(age int) string {
	switch {
	case age <= 12:
		return "Ребенок"
	case age <= 17:
		return "Подросток"
	case age <= 64:
		return "Взрослый"
	default:
		return "Пожилой"
	}
}

// 2.5 Проверка делимости на 3 и 5
func isDivisibleBy3And5(n int) string {
	if n%3 == 0 && n%5 == 0 {
		return "Делится"
	}
	return "Не делится"
}

// 3.1 Факториал числа
func factorial(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

// 3.2 Числа Фибоначчи
func fibonacci(n int) []int {
	fib := make([]int, n)
	if n > 0 {
		fib[0] = 0
	}
	if n > 1 {
		fib[1] = 1
	}
	for i := 2; i < n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	return fib
}

// 3.3 Реверс массива
func reverseArray(arr []int) []int {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

// 3.4 Поиск простых чисел
func primeNumbers(n int) []int {
	primes := []int{}
	for i := 2; i <= n; i++ {
		isPrime := true
		for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, i)
		}
	}
	return primes
}

// 3.5 Сумма чисел в массиве
func sumArray(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func main() {
	fmt.Println("1. Сумма цифр числа:", sumOfDigits(1234))
	fmt.Println("2. Преобразование температуры:", celsiusToFahrenheit(25))
	fmt.Println("3. Удвоение массива:", doubleArray([]int{1, 2, 3, 4}))
	fmt.Println("4. Объединение строк:", joinStrings([]string{"Hello", "world"}))
	fmt.Println("5. Расстояние между точками:", distance(1, 1, 4, 5))
	fmt.Println("2.1 Четное/Нечетное:", isEven(4))
	fmt.Println("2.2 Високосный год:", isLeapYear(2020))
	fmt.Println("2.3 Наибольшее из трех:", maxOfThree(4, 9, 7))
	fmt.Println("2.4 Категория возраста:", ageCategory(25))
	fmt.Println("2.5 Делимость на 3 и 5:", isDivisibleBy3And5(15))
	fmt.Println("3.1 Факториал:", factorial(5))
	fmt.Println("3.2 Числа Фибоначчи:", fibonacci(7))
	fmt.Println("3.3 Реверс массива:", reverseArray([]int{1, 2, 3, 4, 5}))
	fmt.Println("3.4 Простые числа до 20:", primeNumbers(20))
	fmt.Println("3.5 Сумма чисел в массиве:", sumArray([]int{1, 2, 3, 4, 5}))
}
