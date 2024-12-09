package calcadapter

/*
   #cgo CXXFLAGS: -std=c++17
   #cgo LDFLAGS: -L. -lsmart_calc
   #include "../../cpp/src/s21_calc_controller.h"
*/
import "C"
import (
	"bufio"
	"errors"
	"log"
	"os"
	"strings"
	"sync"
)

type Resp struct {
	Result float64
	Err    error
}

const History = "../../history.txt"
const Step = 3000

type TrigonCode rune

const (
	COS  TrigonCode = '@'
	SIN  TrigonCode = 'A'
	TAN  TrigonCode = 'B'
	ACOS TrigonCode = 'C'
	ASIN TrigonCode = 'D'
	ATAN TrigonCode = 'E'
	SQRT TrigonCode = 'F'
	LN   TrigonCode = 'G'
	LOG  TrigonCode = 'H'
)

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func GraphicCalc(str_r string, range_a float64, range_b float64) ([]Point, error) {
	_, err := Calculate(str_r, range_a)
	if err != nil {

		return nil, err
	}
	str := replaceMathFunctions(str_r)
	if range_a == range_b {
		return nil, errors.New("range_a = range_b")
	}
	if range_b < range_a {
		range_a, range_b = range_b, range_a
	}
	diff := (range_b - range_a) / Step
	var result []Point = make([]Point, Step)
	var wg sync.WaitGroup

	for i := range result {
		wg.Add(1)
		x := range_a + float64(i)*diff
		go Calc(&wg, str, C.double(x), &result[i].Y)
		result[i].X = x
	}
	wg.Wait()

	return result, nil
}

func Calc(wg *sync.WaitGroup, str_r string, num C.double, val *float64) {
	cstr := C.CString(str_r)
	if wg == nil {
		log.Println("wg is nil")

		return
	}
	defer wg.Done()
	if val == nil {
		log.Println("val is nil")

		return
	}
	result, err := Calculator(cstr, num)
	if err != nil {
		log.Println(err)

		return
	}
	*val = float64(result)
}

func Calculate(str_r string, num_x float64) (float64, error) {
	str := replaceMathFunctions(str_r)
	cstr := C.CString(str)
	num := C.double(num_x)
	c, err := Calculator(cstr, num)
	if err != nil {
		return 0.0, err
	}
	HistoryWrite(str_r)
	return c, nil
}

func Calculator(cstr *C.char, num C.double) (float64, error) {
	c := C.StartCalc(cstr, num)
	if c.err == 1 {
		return 0.0, errors.New("cc-error: " + C.GoString(c.errors))
	}
	return float64(c.result), nil
}

func replaceMathFunctions(input string) string {
	replacements := map[string]TrigonCode{
		"cos":  COS,
		"sin":  SIN,
		"tan":  TAN,
		"acos": ACOS,
		"asin": ASIN,
		"atan": ATAN,
		"sqrt": SQRT,
		"ln":   LN,
		"log":  LOG,
	}
	for funcName, code := range replacements {
		input = strings.ReplaceAll(input, funcName, string(code))
	}
	input = strings.ReplaceAll(input, " ", "")
	return input
}

func HistoryWrite(text string) error {
	f, err := os.OpenFile(History, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Unable to open file:", err)
		return err
	}
	defer f.Close()
	_, err = f.WriteString(text + "\n")
	if err != nil {
		log.Println("Unable to write file:", err)
		return err
	}
	return nil
}

func HistoryRead() ([]string, error) {
	file, err := os.Open(History)
	if err != nil {
		log.Println("Unable to open file:", err)
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var line []string
	for scanner.Scan() {
		l := scanner.Text()
		line = append(line, l)
	}
	return line, nil
}

func CleanHistory() {
	err := os.Remove(History)
	if err != nil {
		log.Println("Не удалось удалить файл:", err)
		return
	}
}
