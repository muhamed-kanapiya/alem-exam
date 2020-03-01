package main

//last/firstruneprog
func main() {
	fmt.Println(FirstRune("firstruna"))
}
func FirstRune(str string) (res rune) {
	count := 0
	for range str {
		count++
	}
	return rune(str[count-1])
}

//rot13
func main() {
	defer z01.PrintRune('\n')
	if len(os.Args) == 2 {
		for _, r := range os.Args[1] {
			if (r >= 65 && r <= 77) || (r >= 97 && r <= 109) {
				z01.PrintRune(r + 13)
			} else if (r >= 78 && r <= 92) || (r >= 110 && r <= 122) {
				z01.PrintRune(r - 13)
			} else {
				z01.PrintRune(r)
			}
		}
	}
}

//lastword
func main() {
	defer z01.PrintRune('\n')
	if len(os.Args) == 2 {
		var start, end int
		var slice string
		str := os.Args[1]
		for i := len(str) - 1; i != 0; i-- {
			if end == 0 && str[i] != ' ' {
				end = i + 1
			}
			if end > 0 && str[i] == ' ' {
				start = i + 1
			}
		}
		slice = str[start:end]
		for _, r := range slice {
			z01.PrintRune(r)
		}
	}
}

//robotorigin
func main() {
	defer z01.PrintRune('\n')
	result := "false"
	var way int
	if len(os.Args) == 2 {
		for _, r := range os.Args[1] {
			if r == 'U' || r == 'R' {
				way++
			}
			if r == 'D' || r == 'L' {
				way--
			}
		}
		if way == 0 {
			result = "true"
		}
	}
	for _, r := range result {
		z01.PrintRune(r)
	}
}

//tabmult
func main() {
	var res string
	if len(os.Args) == 2 {
		num, _ := strconv.Atoi(os.Args[1])
		for i := 1; i <= 9; i++ {
			res = res + iToA(i) + " x " + iToA(num) + " = " + iToA(num*i) + "\n"
		}
		for _, r := range res {
			z01.PrintRune(r)
		}
	} else{
		z01.PrintRune('\n')
	}

}
func iToA(num int) (str string) {
	for num > 0 {
		str = string(num%10+48) + str
		num = num / 10
	}
	return str
}

alphamirror
func main() {
	defer z01.PrintRune('\n')
	if len(os.Args) == 2 {
		for _, r := range os.Args[1] {
			if r >= 65 && r <= 90 {
				z01.PrintRune(90 - (r - 65))
			} else if r >= 97 && r <= 122 {
				z01.PrintRune(122 - (r - 97))
			} else {
				z01.PrintRune(r)
			}
		}
	}
}

//compareprog
func main() {
	defer z01.PrintRune('\n')
	neg := "-1"
	if len(os.Args) == 3 {
		a := os.Args[1]
		b := os.Args[2]
		if a > b {
			z01.PrintRune(49)
		} else if a == b {
			z01.PrintRune(48)
		} else {
			for _, r := range neg {
				z01.PrintRune(r)
			}
		}
	}
}

doop
func main() {
	if len(os.Args) == 4 {
		var res int
		num1, f := strconv.Atoi(os.Args[1])
		num2, s := strconv.Atoi(os.Args[3])
		if f != nil || s != nil {
			fmt.Println(0)
			return
		}
		if os.Args[2] == "+" {
			res = num1 + num2
			if num1 > 0 && num2 > 0 && res < 0 {
				res = 0
			}
		}
		if os.Args[2] == "-" {
			res = num1 - num2
			if num1 < 0 && num2 < 0 && res < 0 {
				res = 0
			}
		}
		if os.Args[2] == "%" {
			if num2 == 0 {
				fmt.Println("no modulo by 0")
				return
			} else {
				res = num1 % num2
			}
		}
		if os.Args[2] == "/" {
			if num2 == 0 {
				fmt.Println("no division by 0")
				return
			} else {
				res = num1 / num2
			}
		}
		if os.Args[2] == "*" {
			res = num1 * num2
			if res/num1 != num2 {
				res = 0
			}
		}
		fmt.Println(res)
	}
}

//printchessboard
func main() {
	err := "Error"
	if len(os.Args) == 3 {
		x, _ := strconv.Atoi(os.Args[1])
		y, _ := strconv.Atoi(os.Args[2])
		for i := 0; i < y; i++ {
			for j := 0; j < x; j++ {
				if i%2 == 0 {
					if j%2 == 0 {
						z01.PrintRune('#')
					} else {
						z01.PrintRune(' ')
					}
				} else {
					if j%2 == 0 {
						z01.PrintRune(' ')
					} else {
						z01.PrintRune('#')
					}
				}
			}
			z01.PrintRune('\n')
		}
	} else {
		for _, r := range err {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}

//ispowerof2
func main() {
	defer z01.PrintRune('\n')
	flag := "true"
	if len(os.Args) == 2 {
		num, _ := strconv.Atoi(os.Args[1])
		for num > 1 {
			if num%2 != 0 {
				flag = "false"
				break
			}
			num /= 2
		}
	}
	for _, r := range flag {
		z01.PrintRune(r)
	}
}

//union
func main() {
	defer z01.PrintRune('\n')
	var res string
	if len(os.Args) == 3 {
		join := os.Args[1] + os.Args[2]
		for _, r := range join {
			if isUniq(res, r) {
				res += string(r)
			}
		}
		for _, r := range res {
			z01.PrintRune(r)
		}
	}
}
func isUniq(str string, k rune) bool {
	for _, r := range str {
		if r == k {
			return false
		}
	}
	return true
}

func main() {
	arg := os.Args[1:]
	if len(arg) != 2 {
		z01.PrintRune('\n')
		return
	}
	arg1 := arg[0]
	arg2 := arg[1]
	m := map[rune]int{}
	for _, value := range arg1 + arg2 {
		if m[value] < 1 {
			m[value]++
			z01.PrintRune(rune(value))
		}
	}
	z01.PrintRune('\n')
}

//inter
func main() {
	defer z01.PrintRune('\n')
	if len(os.Args) == 3 {
		var res, fin string

		for _, b := range os.Args[1] {
			for _, r := range os.Args[2] {
				if r == b {
					res += string(r)
				}
			}
		}
		for _, r := range res {
			if isUniq(r, fin) {
				fin = fin + string(r)
			}
		}
		for _, r := range fin {
			z01.PrintRune(r)
		}
	}
}

func isUniq(r rune, str string) bool {
	for _, b := range str {
		if r == b {
			return false
		}
	}
	return true
}

func main() {
	defer z01.PrintRune('\n')
	if len(os.Args) == 3 {
		str := map[rune]int{}
		for _, r := range os.Args[1] {
			for _, b := range os.Args[2] {
				if r == b && str != 1 {
					str[r]++
					fmt.Println(string(str))
				}
			}
		}
	}
}

//swapbits
func SwapBits(octet byte) byte {
	return octet<<4 + octet>>4
}

//twosums
func main() {
	case1 := []int{1, 2, 3, 4}
	out := TwoSum(case1, 5)
	fmt.Println(out)
}

func TwoSum(nums []int, target int) []int {
	res := []int{0, 0}
	for i, r := range nums {
		for j, k := range nums {
			if r+k == target && i != j {
				res[1] = i
				res[0] = j
			}
		}
	}
	return res
}

//repeatalpha
func main() {
	defer z01.PrintRune('\n')
	if len(os.Args) == 2 {
		str := os.Args[1]
		for _, r := range str {
			if r >= 65 && r <= 90 {
				for i := 0; i <= int(r)-65; i++ {
					z01.PrintRune(r)
				}
			}
			if r >= 97 && r <= 122 {
				for i := 0; i <= int(r)-97; i++ {
					z01.PrintRune(r)
				}
			} else {
				z01.PrintRune(r)
			}
		}
	}
}

//sortwordarrprog
func main() {
	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	SortWordArr(result)
	fmt.Println(result)
}
func SortWordArr(array []string) {
	count := 0
	for range array {
		count++
	}
	for i := 0; i < count; i++ {
		for j := i + 1; j < count; j++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
}

//printhex
func main() {
	defer z01.PrintRune('\n')
	if len(os.Args) == 2 {
		base := "0123456789abcdef"
		res := ""
		num, err := strconv.Atoi(os.Args[1])
		if err != nil {
			z01.PrintRune('0')
		}
		for num != 0 {
			res = string(base[num%16]) + res
			num /= 16
		}
		for _, r := range res {
			z01.PrintRune(r)
		}
	}
}

//gcd
func main() {
	if len(os.Args) == 3 {
		res := 0
		num1, _ := strconv.Atoi(os.Args[1])
		num2, _ := strconv.Atoi(os.Args[2])
		for i := num1; i > 0; i-- {
			if num1%i == 0 && num2%i == 0 {
				res = i
				break
			}
		}
		fmt.Println(res)
	} else {
		fmt.Println()
		return
	}
}

//anagram
func main() {
	test1 := IsAnagram("a", "a")
	fmt.Println(test1)

	test2 := IsAnagram("alem", "school")
	fmt.Println(test2)

	test3 := IsAnagram("abbc", "    a b bcс   ")
	fmt.Println(test3)

	test4 := IsAnagram("anna madrigal", "a man and a girl")
	fmt.Println(test4)
}

func isAnagram(str1, str2 string) bool {
	str := str1 + str2
	sum := byte(0)
	for _, v := range str {
		if v == ' ' {
			continue
		}
		sum ^= byte(v)
	}
	//fmt.Println(string(sum))
	return sum == 0
}

//balancedstring
func main() {
	if len(os.Args) == 2 {
		var count, res int
		for _, r := range os.Args[1] {
			if r == 'C' {
				count++
			} else {
				count--
			}
			if count == 0 {
				res++
			}
		}
		fmt.Println(res)
	} else {
		fmt.Println()
		return
	}
}

lcm
func main() {
	fmt.Println(Lcm(7, 34))
}
func Lcm(first, second int) int {
	res := 0
	for i := second; i <= first*second; i++ {
		if i%first == 0 && i%second == 0 {
			res = i
			break
		}
	}
	return res
}

//nauuo
func main() {
	fmt.Println(Nauuo(50, 43, 20))
	fmt.Println(Nauuo(13, 13, 0))
	fmt.Println(Nauuo(10, 9, 0))
	fmt.Println(Nauuo(5, 9, 2))
}
func Nauuo(plus, minus, rand int) string {
	var res string
	if plus == minus && rand == 0 {
		res = "0"
	} else if plus > (minus + rand) {
		res = "+"
	} else if (plus + rand) < minus {
		res = "-"
	} else {
		res = "?"
	}
	return res
}

//uniqueoccur
func main() {
	defer z01.PrintRune('\n')
	if len(os.Args) == 2 {
		str := os.Args[1]
		m := map[rune]int{}
		res := ""
		for _, r := range str {
			m[r]++
		}
		fmt.Println(m)
		for key, count := range m {
			if isUniq(m, key, count) {
				toPrint("true")
				res = "true"
			} else {
				toPrint("false")
				return
			}
		}
		toPrint(res)
	}
}
func toPrint(str string) {
	for _, r := range str {
		z01.PrintRune(r)
	}
}
func isUniq(m map[rune]int, key rune, count int) bool {
	for s, num := range m {
		if s != key {
			if num == count {
				return false
			}
		}
	}
	return true
}

//hiddenp
func main() {
	defer z01.PrintRune('\n')
	if len(os.Args) == 3 {
		str1 := os.Args[1]
		str2 := os.Args[2]
		i := 0
		for j := range str2 {
			if str1[i] == str2[j] {
				i++
				if i == len(str1) {
					z01.PrintRune('1')
					return
				}
			}
		}
		z01.PrintRune('0')
	}
}

rostring
func main() {
	defer z01.PrintRune('\n')
	if len(os.Args) == 2 {
		res := swap(splitWS(os.Args[1]))
		res1 := ""
		for i, r := range res {
			if i != len(res)-1 {
				res1 = res1 + r + " "
			} else {
				res1 = res1 + r
			}
		}
		for _, b := range res1 {
			z01.PrintRune(b)
		}
	}
}

func swap(arg []string) []string {
	counter := 0
	for i := range arg {
		counter++
		if counter != len(arg) {
			arg[i], arg[counter] = arg[counter], arg[i]
		}
	}
	return arg
}
func splitWS(str string) []string {
	var s string
	count := 0
	n := 0
	str = str + " "
	for i := range str {
		if str[i] == ' ' {
			s = str[n:i]
			n = i + 1
			if s != "" {
				count++
			}
		}
	}
	n = 0
	arr := make([]string, count)
	j := 0
	for i := range str {
		if str[i] == ' ' {
			s = str[n:i]
			n = i + 1
			if s != "" {
				arr[j] = s
				j++
			}
		}
	}
	return arr
}

//splitprog
func main() {
	str := "HelloHAhowHAareHAyou?"
	fmt.Println(Split(str, "HA"))
}
func Split(str, charset string) []string {
	if charset == "" {
		ar := make([]string, len(str))
		i := 0
		for _, v := range str {
			ar[i] = string(v)
			i++
		}
		return ar
	}
	nw := 0
	a := 0
	for i := 0; i <= len(str)-len(charset); i++ {
		if str[i:i+len(charset)] == charset {
			a++
			i = i + len(charset)
		}
	}
	ar := make([]string, a+1)
	j := 0
	for i := 0; i <= len(str)-len(charset); i++ {
		if str[i:i+len(charset)] == charset {
			ar[j] = str[nw:i]
			j++
			nw = i + len(charset)
			i = i + len(charset)
		}
	}
	ar[j] = str[nw:]
	return ar
}

//revwstr
func main() {
	defer z01.PrintRune('\n')
	if len(os.Args) == 2 {
		var res string
		str := os.Args[1]
		for i := len(str) - 1; i != 0; i-- {
			if str[i] == ' ' {
				res += str[i+1:] + " "
				str = str[:i]
			}
		}
		res += str
		for _, r := range res {
			z01.PrintRune(r)
		}
	}
}

//priorprime
func main() {
	fmt.Println(Prioprime(14))
}
func Prioprime(x int) int {
	res := 0
	for i := 2; i < x; i++ {
		if Prime(i) {
			fmt.Println(i)
			res += i
		}
	}
	return res
}
func Prime(x int) bool {
	for i := 2; i < x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

//fprime
func main() {
    arg := os.Args[1]
    divider := 2
    num, _ := strconv.Atoi(arg)
    for divider <= num {
        if num%divider == 0 {
            fmt.Print(divider)
            if divider == num {
                fmt.Println("")
                return
            }
            fmt.Print("*")
            num /= divider
            divider = 1
        }
        divider++
    }
}
