package gotools

import "fmt"

// Цветной вывод текста в терминал
func ColorPrint(val interface{}, foreground bool, background bool, color ...int) {
	// Без цвета
	if !foreground && !background {
		fmt.Print(val)
		return
	}

	// Цвет текста и фона
	if foreground && background {
		if len(color) < 6 {
			fmt.Print(val)
			return
		}
		fmt.Println(fmt.Sprintf("\033[38;2;%d;%d;%dm\033[48;2;%d;%d;%dm%v\033[0m", color[0], color[1], color[2], color[3], color[4], color[5], val))
		return
	}

	// Только цвет текста
	if foreground {
		if len(color) < 3 {
			fmt.Println(val)
			return
		}
		fmt.Println(fmt.Sprintf("\033[38;2;%d;%d;%dm%v\033[0m", color[0], color[1], color[2], val))
		return
	}

	// Только цвет фона
	if background {
		if len(color) < 3 {
			fmt.Println(val)
			return
		}
		fmt.Println(fmt.Sprintf("\033[48;2;%d;%d;%dm%v\033[0m", color[0], color[1], color[2], val))
		return
	}
}

// Цветной вывод текста в терминал с переносом на новую строку
func ColorPrintln(val interface{}, foreground bool, background bool, color ...int) {
	ColorPrint(val, foreground, background, color...)
	fmt.Println()
}
