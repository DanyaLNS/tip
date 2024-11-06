package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"unicode"
)


// 1. Перевод чисел из одной системы счисления в другую
func convertBase(number string, fromBase, toBase int) (string, error) {
	num, err := strconv.ParseInt(number, fromBase, 64)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(num, toBase), nil
}

// 2. Решение квадратного уравнения
func solveQuadratic(a, b, c float64) (complex128, complex128) {
	discriminant := b*b - 4*a*c
	sqrtDiscriminant := complex(math.Sqrt(math.Abs(discriminant)), 0)
	if discriminant < 0 {
		sqrtDiscriminant = complex(0, math.Sqrt(-discriminant))
	}
	root1 := (-complex(b, 0) + sqrtDiscriminant) / (2 * complex(a, 0))
	root2 := (-complex(b, 0) - sqrtDiscriminant) / (2 * complex(a, 0))
	return root1, root2
}

// 3. Сортировка чисел по модулю
func sortByAbsolute(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		return math.Abs(float64(arr[i])) < math.Abs(float64(arr[j]))
	})
	return arr
}

// 4. Слияние двух отсортированных массивов
func mergeSortedArrays(arr1, arr2 []int) []int {
	i, j := 0, 0
	result := []int{}
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] <= arr2[j] {
			result = append(result, arr1[i])
			i++
		} else {
			result = append(result, arr2[j])
			j++
		}
	}
	result = append(result, arr1[i:]...)
	result = append(result, arr2[j:]...)
	return result
}

// 5. Нахождение подстроки в строке без использования встроенных функций
func findSubstring(s, sub string) int {
	for i := 0; i <= len(s)-len(sub); i++ {
		match := true
		for j := 0; j < len(sub); j++ {
			if s[i+j] != sub[j] {
				match = false
				break
			}
		}
		if match {
			return i
		}
	}
	return -1
}

// 1. Калькулятор с расширенными операциями
func calculator(a, b float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("деление на ноль")
		}
		return a / b, nil
	case "^":
		return math.Pow(a, b), nil
	case "%":
		return float64(int(a) % int(b)), nil
	default:
		return 0, fmt.Errorf("неизвестная операция")
	}
}

// 2. Проверка палиндрома
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	filtered := ""
	for _, ch := range s {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
			filtered += string(ch)
		}
	}
	for i := 0; i < len(filtered)/2; i++ {
		if filtered[i] != filtered[len(filtered)-1-i] {
			return false
		}
	}
	return true
}

// 3. Нахождение пересечения трех отрезков
func hasIntersection(a1, a2, b1, b2, c1, c2 int) bool {
	start := max(a1, b1, c1)
	end := min(a2, b2, c2)
	return start <= end
}

func max(a, b, c int) int {
	if a > b {
		if a > c {
			return a
		}
		return c
	}
	if b > c {
		return b
	}
	return c
}

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}

// 4. Выбор самого длинного слова в предложении
func longestWord(sentence string) string {
	words := strings.FieldsFunc(sentence, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r)
	})
	longest := ""
	for _, word := range words {
		if len(word) > len(longest) {
			longest = word
		}
	}
	return longest
}

// 5. Проверка високосного года
func isLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || (year%400 == 0)
}

// 1. Числа Фибоначчи до определенного числа
func fibonacciUntil(n int) []int {
	fib := []int{0, 1}
	for {
		next := fib[len(fib)-1] + fib[len(fib)-2]
		if next > n {
			break
		}
		fib = append(fib, next)
	}
	return fib
}

// 2. Определение простых чисел в диапазоне
func primesInRange(start, end int) []int {
	primes := []int{}
	for i := start; i <= end; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// 3. Числа Армстронга в заданном диапазоне
func armstrongNumbersInRange(start, end int) []int {
	armstrongs := []int{}
	for i := start; i <= end; i++ {
		if isArmstrong(i) {
			armstrongs = append(armstrongs, i)
		}
	}
	return armstrongs
}

func isArmstrong(n int) bool {
	sum, temp, digits := 0, n, len(strconv.Itoa(n))
	for temp > 0 {
		digit := temp % 10
		sum += int(math.Pow(float64(digit), float64(digits)))
		temp /= 10
	}
	return sum == n
}

// 4. Реверс строки
func reverseString(s string) string {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}

// 5. Нахождение наибольшего общего делителя (НОД)
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	res, err := convertBase("101", 2, 10)
	if err != nil {
		fmt.Println("Ошибка при переводе системы счисления: ", err)
	} else {
		fmt.Println("Перевод системы счисления:", res)
	}

	a, b := solveQuadratic(1, -3, 2)
	fmt.Println("Корни квадратного уравнения:", a, b)
	fmt.Println("Сортировка по модулю:", sortByAbsolute([]int{-5, 1, -3, 2, 4}))
	fmt.Println("Слияние массивов:", mergeSortedArrays([]int{1, 3, 5}, []int{2, 4, 6}))
	fmt.Println("Нахождение подстроки:", findSubstring("hello world", "world"))

	cRes, err := calculator(10, 3, "%")
	if err != nil {
		fmt.Println("Ошибка при выполнении вычислений: ", err)
	} else {
		fmt.Println("Расширенный калькулятор:", cRes)
	}

	fmt.Println("Проверка палиндрома:", isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println("Пересечение отрезков:", hasIntersection(1, 5, 2, 6, 3, 7))
	fmt.Println("Самое длинное слово:", longestWord("Очень интересный текст, в котором нужно найти длиннейшее слово!"))
	fmt.Println("Високосный год:", isLeapYear(2024))
	fmt.Println("Числа Фибоначчи:", fibonacciUntil(20))
	fmt.Println("Простые числа в диапазоне:", primesInRange(10, 50))
	fmt.Println("Числа Армстронга:", armstrongNumbersInRange(100, 999))
	fmt.Println("Реверс строки:", reverseString("hello"))
	fmt.Println("НОД:", gcd(48, 18))
}
