package calcadapter

/*
   #cgo CXXFLAGS:
   #cgo LDFLAGS: -L. -lsmart_calc
   #include "../../cpp/src/s21_calc_controller.h"
*/
import "C"
import (
	"bufio"
	"errors"
	"log"
	"math"
	"os"
	"slices"
	"strings"
)

type Resp struct {
	Result float64
	Err    error
}

const History = "./history.txt"
const Step = 3000
const MaxHistoryStore = 35
const MaxHistory = 25

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

type GraphSrc struct {
	Str_r         string
	Range_a       float64
	Range_b       float64
	Value_range_a float64
	Value_range_b float64
}

func GraphicCalc(data GraphSrc) ([]Point, error) {
	_, err := Calculate(data.Str_r, data.Range_a)
	if err != nil {
		var result []Point
		return result, err
	}
	str := replaceMathFunctions(data.Str_r)
	if data.Range_a == data.Range_b || data.Value_range_a == data.Value_range_b {
		return nil, errors.New("range_a = range_b or value_range_a == value_range_b")
	}
	if data.Range_b < data.Range_a {
		data.Range_a, data.Range_b = data.Range_b, data.Range_a
	}
	if data.Value_range_a < data.Value_range_b {
		data.Value_range_a, data.Value_range_b = data.Value_range_b, data.Value_range_a
	}
	diff := math.Abs(data.Range_b-data.Range_a) / Step
	var result []Point = make([]Point, Step)

	for i := range result {
		x := data.Range_a + float64(i)*diff
		result[i].X = x
		calc_y, err := Calculator(str, x)
		if err != nil {
			return nil, err
		}
		if calc_y > data.Value_range_a {
			calc_y = data.Value_range_a
		}
		if calc_y < data.Value_range_b {
			calc_y = data.Value_range_b
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
	if e == 1 {
		var tmp string
		if c.errors != nil {
			tmp = C.GoString(cstr)
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
	file, err := os.OpenFile(History, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Println("unable to open file:", err)
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		count++
	}
	if count > MaxHistoryStore {
		history_tmp, err := HistoryRead()
		if err != nil {

			return err
		}
		err = file.Truncate(0)
		if err != nil {

			return err
		}
		_, err = file.Seek(0, 0)
		if err != nil {

			return err
		}
		for i := len(history_tmp) - 1; i >= 0; i-- {
			_, err = file.WriteString(history_tmp[i] + "\n")
			if err != nil {
				log.Println("unable to write file:", err)

				return err
			}
		}
	}
	_, err = file.WriteString(text + "\n")
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
	slices.Reverse(line)
	if len(line) > MaxHistory {
		line = line[:MaxHistory]
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
