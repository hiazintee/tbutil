package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"strconv"
)

// print_tb draws a string at a specific location
func print_tb(line, col int, msg string) {
	for _, c := range msg {
		termbox.SetCell(col, line, c, termbox.ColorDefault, termbox.ColorDefault)
		col++
	}
}

func printf_tb(x, y int, format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	print_tb(x, y, s)
}

func drawat(line, col int, char rune) {
	termbox.SetCell(col, line, char, termbox.ColorDefault, termbox.ColorDefault)
}

// drawline draws a horizontal line on line number y
// constisting of str.
func drawHline(line int, char rune) {
	width, _ := termbox.Size() // width, height
	for col := 0; col < width; col++ {
		drawat(line, col, char)
	}
}

// drawVline draws a vertical line on column x
// constisting of str.
func drawVline(col int, char rune) {
	_, height := termbox.Size() // width, heigth
	for line := 0; line < height-2; line++ {
		drawat(line, col, char)
	}
}

func drawBorder() {
	// get the "screen" size
	width, height := termbox.Size()
	lineChar, verticalChar, crossChar := '-', '|', '+'

	// draw border
	drawHline(0, lineChar)
	drawHline(height-2, lineChar)
	drawVline(0, verticalChar)
	drawVline(width-1, verticalChar)

	// draw coners
	drawat(0, 0, crossChar)              // left upper corner
	drawat(0, width-1, crossChar)        // right upper corner
	drawat(height-2, 0, crossChar)       // left lower corner
	drawat(height-2, width-1, crossChar) // right lower corner

}

func printTitle(title string) {
	width, _ := termbox.Size()
	print_tb(1, width/2-len(title)/2, title)
}

func printResult(results []string, skipIndex, skipColumns int) {
	width, height := termbox.Size()
	startLine, StartCol := 9, 5
	MaxLine, MaxCol := height-startLine-1, width-StartCol-1

	for curLine := startLine; curLine < MaxLine; curLine++ {
		// if skipIndex > len(results) {
		if MaxLine >= curLine && skipColumns < MaxCol && skipIndex < len(results)-2 {
			pline := results[skipIndex]
			print_tb(curLine, StartCol, pline)
			skipIndex += 1
		} else {
			msg1 := fmt.Sprintf("Index out of bound. curLine:%v skipVIndex=%v SkipColumns:%v", curLine, skipIndex, skipColumns)
			msg2 := fmt.Sprintf("MaxLine=%v MaxCol:%v len(results):%v len(results[]):%v", MaxLine, MaxCol, len(results), len(results[skipIndex]))
			print_tb(curLine, 3, msg1)
			curLine += 1
			print_tb(curLine, 3, msg2)
			curLine += 1
		}
	}
}

// prepare_main_screen draws the main screen sections within the termbox back
// buffer
func prepare_main_screen() {
	width, height := termbox.Size()
	drawBorder()
	curLine := 2

	printTitle("--= The Main Screen =--")
	curLine++
	print_tb(curLine, 1, "Current valid screens/commands: sesKill sqlTextById sqlPlanByIdChild ...")
	curLine++
	print_tb(curLine, 1, "ResLines: 23-31/2994  ResCols: 156-236/7352 WinSize: 20 lines 70 columns")
	curLine++
	print_tb(curLine, 1, "")
	curLine++

	drawHline(curLine, '-')
	drawat(curLine, 0, '+')
	drawat(curLine, width-1, '+')

	drawHline(curLine+2, '-')
	drawat(curLine+2, 0, '+')
	drawat(curLine+2, width-1, '+')

	curLine++
	tc := 1
	for i := 1; i < width-1; i++ {
		if i%10 == 0 {
			print_tb(curLine+1, i, strconv.Itoa(tc))
			tc++
		}
		print_tb(curLine, i, strconv.Itoa(i%10))
	}

	print_tb(height-1, 0, "Command: ")
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func draw_all(call string) {
	result := []string{"Results: ",
		"---------",
		"This is just a Little bit of output to demonstrate (FAKE)!",
		"",
		"SQL> select * from dual;",
		"",
		"X",
		"_",
		"",
		"",
		"SQL>"}

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	prepare_main_screen()

	printResult(result, 0, 0)

	/*
		w, h := termbox.Size()
		msg := fmt.Sprintf("width: %v heigth: %v    Message: %v", w, h, call)
		print_tb(5, 5, msg)
	*/

	termbox.Flush()
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	draw_all("")
loop:
	for {
		ev := termbox.PollEvent()

		switch ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				break loop
			}
		case termbox.EventResize:
			draw_all("Resize!!")
		}
	}
}
