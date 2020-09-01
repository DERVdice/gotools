package gotools

import "fmt"

// Цвета, которые можно использовать для вывода в терминал За основу взяты цвета JetBrains Darcula Theme.
// Например: Цвет текста и фона: ColorPrintln("text", true, true, append(Red, Black...)...)
var (
	Red    = []int{210, 82, 82}
	Yellow = []int{255, 198, 109}
	Green  = []int{97, 150, 71}
	Blue   = []int{104, 151, 187}
	White  = []int{255, 255, 255}
	Black  = []int{0, 0, 0}
)

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
			fmt.Print(val)
			return
		}
		fmt.Print(fmt.Sprintf("\033[38;2;%d;%d;%dm%v\033[0m", color[0], color[1], color[2], val))
		return
	}

	// Только цвет фона
	if background {
		if len(color) < 3 {
			fmt.Print(val)
			return
		}
		fmt.Print(fmt.Sprintf("\033[48;2;%d;%d;%dm%v\033[0m", color[0], color[1], color[2], val))
		return
	}
}

// Цветной вывод текста в терминал с переносом на новую строку
func ColorPrintln(val interface{}, foreground bool, background bool, color ...int) {
	ColorPrint(val, foreground, background, color...)
	fmt.Println()
}
