package kata

import (
	"math"
	"reflect"
)

var (
	validLR = map[rune]rune{
		'0': '0',
		'1': '1',
		'6': '9',
		'8': '8',
		'9': '6',
	}
	validC = map[rune]rune{
		'0': '0',
		'1': '1',
		'8': '8',
	}
	varityLR = map[rune]int{
		'0': 1,
		'1': 2,
		'2': 2,
		'3': 2,
		'4': 2,
		'5': 2,
		'6': 3,
		'7': 3,
		'8': 4,
		'9': 5,
	}
	varityC = map[rune]int{
		'0': 1,
		'1': 2,
		'2': 2,
		'3': 2,
		'4': 2,
		'5': 2,
		'6': 2,
		'7': 2,
		'8': 3,
		'9': 3,
	}
)

// firstZero - if include first zero => 1 else 0
// firstZero:
// 1 - исключить ноль
// 0 - включить ноль
func razryad(coef rune, order int, firstZero int) (result int) {
	if order == 0 {
		return varityC[coef] - firstZero
	}

	// умножаем на: половина длины, так как зеркальная часть нас не интересует
	result = (varityLR[coef] - firstZero) * int(math.Pow(float64(5), float64((order+1)/2)-1))

	// если кол-во нулей четной => число имеет нечетную длину (+ первая цифра) => имеет не посчитанный центр
	if order%2 == 0 {
		result *= 3
	}

	return
}

func checkMirrorNum(n string) bool {
	var L, mirrorL, R []rune
	var C rune

	ln := len(n)

	if ln == 0 {
		return false
	}

	if ln%2 == 0 {
		L, R = []rune(n[:ln/2]), []rune(n[ln/2:])
	} else {
		L, R = []rune(n[:ln/2]), []rune(n[ln/2+1:])
		C = rune(n[ln/2])
	}

	if ln%2 != 0 {
		if _, valid := validC[C]; !valid {
			return false
		}
	}

	if ln == 1 {
		return true
	}

	for _, v := range L {
		if mirrorN, valid := validLR[v]; valid {
			mirrorL = append([]rune{mirrorN}, mirrorL...)
			continue
		}
		return false
	}

	return reflect.DeepEqual(mirrorL, R)
}

func polozProsmotr(n string, firstZero int) (result int) {
	ln := len(n)

	//проверка частных случаев
	switch ln {
	case 0:
		return 0
	case 1:
		return varityC[rune(n[0])] - firstZero
	}

	// получение левой и правой части
	var Left, Right []rune
	var Centr string

	if ln%2 == 0 {
		Left, Right = []rune(n[:ln/2]), []rune(n[ln/2:])
	} else {
		Left, Right = []rune(n[:ln/2]), []rune(n[ln/2+1:])
	}

	// if first n is valid
	if mirrorFirstN, firstIsValid := validLR[Left[0]]; firstIsValid {
		// считаем сначала для десятка ниже
		if Left[0] != '0' && !(Left[0] == '1' && firstZero == 1) {
			result += razryad(Left[0]-1, ln-1, firstZero)
		}

		Centr = n[1 : len(n)-1]
		//подсчет вариантов в центре
		result += polozProsmotr(Centr, 0)

		//проверка центральной части на валидность
		if Centr == "" || checkMirrorNum(Centr) {
			if Centr != "" {
				result -= 1
			}
			if Right[len(Right)-1] >= mirrorFirstN {
				result += 1
			}
		}
	} else {
		result += razryad(Left[0], ln-1, firstZero)
	}

	return
}

func UpsideDown(n1, n2 string) uint64 {
	var (
		result   int
		n1IsZero int
	)

	if n1[0] == '0' {
		n1IsZero = 0
	} else {
		n1IsZero = 1
	}

	lnN1, lnN2 := len(n1), len(n2)

	if lnN2-lnN1 < 1 {
		result = polozProsmotr(n2, n1IsZero)
		result -= polozProsmotr(n1, n1IsZero)
	} else {
		result = razryad('9', len(n1)-1, n1IsZero)
		result -= polozProsmotr(n1, n1IsZero)

		if checkMirrorNum(n1) {
			result += 1
		}

		for i := 1; i < lnN2-lnN1; i++ {
			result += razryad('9', lnN1-1+i, 1)
		}

		result += polozProsmotr(n2, 1)
	}

	return uint64(result)
}
