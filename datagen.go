package datagen

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func ID() int64 {
	return rand.Int63()
}

func UUID() string {
	return uuid.New().String()
}

func Email() (email string) {
	domains := []string{"email.com", "hotmail.com", "gmail.com", "protonmail.com"}
	domain := domains[rand.Intn(len(domains))]
	return fmt.Sprintf("%s@%s", StringWithAlphabetic(rand.Intn(16)+5), domain)
}

func Name() (name string) {
	namesAmount := rand.Intn(3) + 1
	for i := 0; i < namesAmount; i++ {
		name += StringWithAlphabetic(rand.Intn(15) + 1)
		if i != (namesAmount - 1) {
			name += " "
		}
	}

	return name
}

func Phone() string {
	return strconv.Itoa(rand.Intn(1_00000_0000) + 10000_0000)
}

func CPF() string {

	cpfInt := make([]int, 9)
	for i := range cpfInt {
		cpfInt[i] = rand.Intn(10)
	}

	var multResult int
	for i := 0; i < 9; i++ {
		multResult += cpfInt[i] * (10 - i)
	}
	remainder := multResult % 11

	if remainder < 2 {
		cpfInt = append(cpfInt, 0)
	} else {
		cpfInt = append(cpfInt, 11-remainder)
	}

	multResult = 0
	for i := 0; i < 10; i++ {
		multResult += cpfInt[i] * (11 - i)
	}
	remainder = multResult % 11

	if remainder < 2 {
		cpfInt = append(cpfInt, 0)
	} else {
		cpfInt = append(cpfInt, 11-remainder)
	}

	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(cpfInt)), ""), "[]")
}

func CardNumber() string {
	return fmt.Sprintf("%v", rand.Intn(1_0000_0000_0000_0000)+1000_0000_0000_0000)
}

func CardSecurityCode() string {
	return fmt.Sprintf("%v", rand.Intn(1_000)+100)
}

func Month() (month string) {

	month = fmt.Sprintf("%02v", rand.Intn(12)+1)
	return month
}

func FutureYear() string {
	return fmt.Sprintf("%v", rand.Intn(100)+time.Now().Year())
}

func StringWithAlphabetic(length int) string {
	runes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	return String(runes, length)
}

func StringWithAlphanumeric(length int) string {
	runes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	return String(runes, length)
}

func String(runes []rune, length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = runes[rand.Intn(len(runes))]
	}

	return string(b)
}

func Bool() bool {
	n := rand.Intn(2)
	return n == 1
}

func Int64(min, max int64) int64 {
	return rand.Int63n(max-min) + min
}

func Gender() string {
	genders := []string{"M", "F"}
	return genders[rand.Intn(len(genders))]
}
