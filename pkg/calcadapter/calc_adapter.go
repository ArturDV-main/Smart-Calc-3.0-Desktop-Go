package calcadapter

/*
   #cgo CXXFLAGS: -std=c++17
   #cgo LDFLAGS: -L. -lsmart_calc -static-libstdc++ -static-libgcc
   #include "../../cpp/src/s21_calc_controller.h"
*/
import "C"
import (
	"bufio"
	"errors"
	"log"
	"os"
	"strings"
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
		var result []Point
		return result, err
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

	for i := range result {
		x := range_a + float64(i)*diff
		result[i].X = x
		calc_y, err := Calculator(str, x)
		if err != nil {
			return nil, err
		}
		result[i].Y = calc_y
	}

	return result, nil
}

func Calculate(str_r string, num_x float64) (float64, error) {
	str := replaceMathFunctions(str_r)
	c, err := Calculator(str, num_x)
	if err != nil {
		return 0.0, err
	}
	HistoryWrite(str_r)
	return c, nil
}

func Calculator(str string, x float64) (float64, error) {
	num := C.double(x)
	cstr := C.CString(str)
	c := C.StartCalc(cstr, num)
	e := c.err
	log.Println(e)
	if e == 1 {
		var tmp string
		if c.errors != nil {
			tmp = C.GoStringN(cstr, 5)
		}
		return 0.0, errors.New("cc-error: " + tmp)
	}
	result := float64(c.result)
	return result, nil
}

func replaceMathFunctions(input string) string {
	replacements := map[string]TrigonCode{
		"acos": ACOS,
		"asin": ASIN,
		"atan": ATAN,
	}
	replacements2 := map[string]TrigonCode{
		"cos":  COS,
		"sin":  SIN,
		"tan":  TAN,
		"sqrt": SQRT,
		"ln":   LN,
		"log":  LOG,
	}
	for funcName, code := range replacements {
		input = strings.ReplaceAll(input, funcName, string(code))
	}
	for funcName, code := range replacements2 {
		input = strings.ReplaceAll(input, funcName, string(code))
	}
	input = strings.ReplaceAll(input, " ", "")
	return input
}

func HistoryWrite(text string) error {
	f, err := os.OpenFile(History, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("unable to open file:", err)
		return err
	}
	defer f.Close()
	_, err = f.WriteString(text + "\n")
	if err != nil {
		log.Println("unable to write file:", err)
		return err
	}
	return nil
}

func HistoryRead() ([]string, error) {
	file, err := os.Open(History)
	if err != nil {
		log.Println("unable to open file:", err)
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
		log.Println("unable to delete file: ", err)
		return
	}
}
