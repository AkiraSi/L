package main

var justString string // вообще не круто создавать глобальную переменную

func someFunc() {
	v := createHugeString(1 << 10)
	if !(len(v) <= 100) { // иначе паника будет, которую не перехватывали в исходном коде
		justString = v[:100]
	} else {
		justString = v
	}
}

func createHugeString(size int) string { // имитация фукнции
	return "tyt tipa big string"
}

func main() {
	someFunc()
}
