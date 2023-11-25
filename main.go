package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func getArt(char byte) []string {
	font := make(map[byte][]string, 100)
	font[32] = []string{
		"     $$",
		"    $$ ",
		"   $$  ",
		"  $$   ",
		" $$    ",
		"$$     ",
	}
	font[33] = []string{
		"    __",
		"   / /",
		"  / / ",
		" /_/  ",
		"(_)   ",
		"      ",
	}
	font[34] = []string{
		" _ _ ",
		"( | )",
		"|/|/ ",
		" $   ",
		"$    ",
		"     ",
	}
	font[35] = []string{
		"     __ __ ",
		"  __/ // /_",
		" /_  _  __/",
		"/_  _  __/ ",
		" /_//_/    ",
		"           ",
	}
	font[36] = []string{
		"     __",
		"   _/ /",
		"  / __/",
		" (_  ) ",
		"/  _/  ",
		"/_/    ",
	}
	font[37] = []string{
		"   _   __",
		"  (_)_/_/",
		"   _/_/  ",
		" _/_/_   ",
		"/_/ (_)  ",
		"         ",
	}
	font[38] = []string{
		"   ___   ",
		"  ( _ )  ",
		" / __ \\/|",
		"/ /_/  < ",
		"\\____/\\/ ",
		"         ",
	}
	font[39] = []string{
		"  _ ",
		" ( )",
		" |/ ",
		" $  ",
		"$   ",
		"    ",
	}
	font[40] = []string{
		"     __",
		"   _/_/",
		"  / /  ",
		" / /   ",
		"/ /    ",
		"|_|    ",
	}
	font[41] = []string{
		"     _ ",
		"    | |",
		"    / /",
		"   / / ",
		" _/_/  ",
		"/_/    ",
	}
	font[42] = []string{
		"       ",
		"  __/|_",
		" |    /",
		"/_ __| ",
		" |/    ",
		"       ",
	}
	font[43] = []string{
		"       ",
		"    __ ",
		" __/ /_",
		"/_  __/",
		" /_/   ",
		"       ",
	}
	font[44] = []string{
		"   ",
		"   ",
		"   ",
		" _ ",
		"( )",
		"|/ ",
	}
	font[45] = []string{
		"       ",
		"       ",
		" ______",
		"/_____/",
		"  $    ",
		"       ",
	}
	font[46] = []string{
		"   ",
		"   ",
		"   ",
		" _ ",
		"(_)",
		"   ",
	}
	font[47] = []string{
		"       __",
		"     _/_/",
		"   _/_/  ",
		" _/_/    ",
		"/_/      ",
		"         ",
	}
	font[48] = []string{
		"   ____ ",
		"  / __ \\",
		" / / / /",
		"/ /_/ / ",
		"\\____/  ",
		"        ",
	}
	font[49] = []string{
		"   ___",
		"  <  /",
		"  / / ",
		" / /  ",
		"/_/   ",
		"      ",
	}
	font[50] = []string{
		"   ___ ",
		"  /__ \\",
		"  __/ /",
		" / __/ ",
		"/____/ ",
		"       ",
	}
	font[51] = []string{
		"   _____",
		"  /__  /",
		"   /_ < ",
		" ___/ / ",
		"/____/  ",
		"        ",
	}
	font[52] = []string{
		"   __ __",
		"  / // /",
		" / // /_",
		"/__  __/",
		"  /_/   ",
		"        ",
	}
	font[53] = []string{
		"    ______",
		"   / ____/",
		"  /___ \\  ",
		" ____/ /  ",
		"/_____/   ",
		"          ",
	}
	font[54] = []string{
		"   _____",
		"  / ___/",
		" / __ \\ ",
		"/ /_/ / ",
		"\\____/  ",
		"        ",
	}
	font[55] = []string{
		" _____",
		"/__  /",
		"  / / ",
		" / /  ",
		"/_/   ",
		"      ",
	}
	font[56] = []string{
		"   ____ ",
		"  / __ \\",
		" / __  /",
		"/ /_/ / ",
		"\\____/  ",
		"        ",
	}
	font[57] = []string{
		"   ____ ",
		"  / __ \\",
		" / /_/ /",
		" \\__, / ",
		"/____/  ",
		"        ",
	}
	font[58] = []string{
		"     ",
		"   _ ",
		"  (_)",
		" _   ",
		"(_)  ",
		"     ",
	}
	font[59] = []string{
		"     ",
		"   _ ",
		"  (_)",
		" _   ",
		"( )  ",
		"|/   ",
	}
	font[60] = []string{
		"  __",
		" / /",
		"/ / ",
		"\\ \\ ",
		" \\_\\",
		"    ",
	}
	font[61] = []string{
		"       ",
		"  _____",
		" /____/",
		"/____/ ",
		"  $    ",
		"       ",
	}
	font[62] = []string{
		"__  ",
		"\\ \\ ",
		" \\ \\",
		" / /",
		"/_/ ",
		"    ",
	}
	font[63] = []string{
		"  ___ ",
		" /__ \\",
		"  / _/",
		" /_/  ",
		"(_)   ",
		"      ",
	}
	font[64] = []string{
		"   ______ ",
		"  / ____ \\",
		" / / __ `/",
		"/ / /_/ / ",
		"\\ \\__,_/  ",
		" \\____/   ",
	}
	font[65] = []string{
		"    ___ ",
		"   /   |",
		"  / /| |",
		" / ___ |",
		"/_/  |_|",
		"        ",
	}
	font[66] = []string{
		"    ____ ",
		"   / __ )",
		"  / __  |",
		" / /_/ / ",
		"/_____/  ",
		"         ",
	}
	font[67] = []string{
		"   ______",
		"  / ____/",
		" / /     ",
		"/ /___   ",
		"\\____/   ",
		"         ",
	}
	font[68] = []string{
		"    ____ ",
		"   / __ \\",
		"  / / / /",
		" / /_/ / ",
		"/_____/  ",
		"         ",
	}
	font[69] = []string{
		"    ______",
		"   / ____/",
		"  / __/   ",
		" / /___   ",
		"/_____/   ",
		"          ",
	}
	font[70] = []string{
		"    ______",
		"   / ____/",
		"  / /_    ",
		" / __/    ",
		"/_/       ",
		"          ",
	}
	font[71] = []string{
		"   ______",
		"  / ____/",
		" / / __  ",
		"/ /_/ /  ",
		"\\____/   ",
		"         ",
	}
	font[72] = []string{
		"    __  __",
		"   / / / /",
		"  / /_/ / ",
		" / __  /  ",
		"/_/ /_/   ",
		"          ",
	}
	font[73] = []string{
		"    ____",
		"   /  _/",
		"   / /  ",
		" _/ /   ",
		"/___/   ",
		"        ",
	}
	font[74] = []string{
		"       __",
		"      / /",
		" __  / / ",
		"/ /_/ /  ",
		"\\____/   ",
		"         ",
	}
	font[75] = []string{
		"    __ __",
		"   / //_/",
		"  / ,<   ",
		" / /| |  ",
		"/_/ |_|  ",
		"         ",
	}
	font[76] = []string{
		"    __ ",
		"   / / ",
		"  / /  ",
		" / /___",
		"/_____/",
		"       ",
	}
	font[77] = []string{
		"    __  ___",
		"   /  |/  /",
		"  / /|_/ / ",
		" / /  / /  ",
		"/_/  /_/   ",
		"           ",
	}
	font[78] = []string{
		"    _   __",
		"   / | / /",
		"  /  |/ / ",
		" / /|  /  ",
		"/_/ |_/   ",
		"          ",
	}
	font[79] = []string{
		"   ____ ",
		"  / __ \\",
		" / / / /",
		"/ /_/ / ",
		"\\____/  ",
		"        ",
	}
	font[80] = []string{
		"    ____ ",
		"   / __ \\",
		"  / /_/ /",
		" / ____/ ",
		"/_/      ",
		"         ",
	}
	font[81] = []string{
		"   ____ ",
		"  / __ \\",
		" / / / /",
		"/ /_/ / ",
		"\\___\\_\\ ",
		"        ",
	}
	font[82] = []string{
		"    ____ ",
		"   / __ \\",
		"  / /_/ /",
		" / _, _/ ",
		"/_/ |_|  ",
		"         ",
	}
	font[83] = []string{
		"   _____",
		"  / ___/",
		"  \\__ \\ ",
		" ___/ / ",
		"/____/  ",
		"        ",
	}
	font[84] = []string{
		"  ______",
		" /_  __/",
		"  / /   ",
		" / /    ",
		"/_/     ",
		"        ",
	}
	font[85] = []string{
		"   __  __",
		"  / / / /",
		" / / / / ",
		"/ /_/ /  ",
		"\\____/   ",
		"         ",
	}
	font[86] = []string{
		" _    __",
		"| |  / /",
		"| | / / ",
		"| |/ /  ",
		"|___/   ",
		"        ",
	}
	font[87] = []string{
		" _       __",
		"| |     / /",
		"| | /| / / ",
		"| |/ |/ /  ",
		"|__/|__/   ",
		"           ",
	}
	font[88] = []string{
		"   _  __",
		"  | |/ /",
		"  |   / ",
		" /   |  ",
		"/_/|_|  ",
		"        ",
	}
	font[89] = []string{
		"__  __",
		"\\ \\/ /",
		" \\  / ",
		" / /  ",
		"/_/   ",
		"      ",
	}
	font[90] = []string{
		" _____",
		"/__  /",
		"  / / ",
		" / /__", //
		"/____/",
		"      ",
	}
	font[91] = []string{
		"     ___",
		"    / _/",
		"   / /  ",
		"  / /   ",
		" / /    ",
		"/__/    ",
	}
	font[92] = []string{
		"__    ",
		"\\ \\   ",
		" \\ \\  ",
		"  \\ \\ ",
		"   \\_\\",
		"      ",
	}
	font[93] = []string{
		"     ___",
		"    /  /",
		"    / / ",
		"   / /  ",
		" _/ /   ",
		"/__/    ",
	}
	font[94] = []string{
		"  //|",
		" |/||",
		"  $  ",
		" $   ",
		"$    ",
		"     ",
	}
	font[95] = []string{
		"       ",
		"       ",
		"       ",
		"       ",
		" ______",
		"/_____/",
	}
	font[96] = []string{
		"  _ ",
		" ( )",
		"  V ",
		" $  ",
		"$   ",
		"    ",
	}
	font[97] = []string{
		"        ",
		"  ____ _",
		" / __ `/",
		"/ /_/ / ",
		"\\__,_/  ",
		"        ",
	}
	font[98] = []string{
		"    __  ",
		"   / /_ ",
		"  / __ \\",
		" / /_/ /",
		"/_.___/ ",
		"        ",
	}
	font[99] = []string{
		"       ",
		"  _____",
		" / ___/",
		"/ /__  ",
		"\\___/  ",
		"       ",
	}
	font[100] = []string{
		"       __",
		"  ____/ /",
		" / __  / ",
		"/ /_/ /  ",
		"\\__,_/   ",
		"         ",
	}
	font[101] = []string{
		"      ",
		"  ___ ",
		" / _ \\",
		"/  __/",
		"\\___/ ",
		"      ",
	}
	font[102] = []string{
		"    ____",
		"   / __/",
		"  / /_  ",
		" / __/  ",
		"/_/     ",
		"        ",
	}
	font[103] = []string{
		"         ",
		"   ____ _",
		"  / __ `/",
		" / /_/ / ",
		" \\__, /  ",
		"/____/   ",
	}
	font[104] = []string{
		"    __  ",
		"   / /_ ",
		"  / __ \\",
		" / / / /",
		"/_/ /_/ ",
		"        ",
	}
	font[105] = []string{
		"    _ ",
		"   (_)",
		"  / / ",
		" / /  ",
		"/_/   ",
		"      ",
	}
	font[106] = []string{
		"       _ ",
		"      (_)",
		"     / / ",
		"    / /  ",
		" __/ /   ",
		"/___/    ",
	}
	font[107] = []string{
		"    __  ",
		"   / /__",
		"  / //_/",
		" / ,<   ",
		"/_/|_|  ",
		"        ",
	}
	font[108] = []string{
		"    __",
		"   / /",
		"  / / ",
		" / /  ",
		"/_/   ",
		"      ",
	}
	font[109] = []string{
		"            ",
		"   ____ ___ ",
		"  / __ `__ \\",
		" / / / / / /",
		"/_/ /_/ /_/ ",
		"            ",
	}
	font[110] = []string{
		"        ",
		"   ____ ",
		"  / __ \\",
		" / / / /",
		"/_/ /_/ ",
		"        ",
	}
	font[111] = []string{
		"       ",
		"  ____ ",
		" / __ \\",
		"/ /_/ /",
		"\\____/ ",
		"       ",
	}
	font[112] = []string{
		"         ",
		"    ____ ",
		"   / __ \\",
		"  / /_/ /",
		" / .___/ ",
		"/_/      ",
	}
	font[113] = []string{
		"        ",
		"  ____ _",
		" / __ `/",
		"/ /_/ / ",
		"\\__, /  ",
		"  /_/   ",
	}
	font[114] = []string{
		"        ",
		"   _____",
		"  / ___/",
		" / /    ",
		"/_/     ",
		"        ",
	}
	font[115] = []string{
		"        ",
		"   _____",
		"  / ___/",
		" (__  ) ",
		"/____/  ",
		"        ",
	}
	font[116] = []string{
		"   __ ",
		"  / /_",
		" / __/",
		"/ /_  ",
		"\\__/  ",
		"      ",
	}
	font[117] = []string{
		"        ",
		"  __  __",
		" / / / /",
		"/ /_/ / ",
		"\\__,_/  ",
		"        ",
	}
	font[118] = []string{
		"       ",
		" _   __",
		"| | / /",
		"| |/ / ",
		"|___/  ",
		"       ",
	}
	font[119] = []string{
		"          ",
		" _      __",
		"| | /| / /",
		"| |/ |/ / ",
		"|__/|__/  ",
		"          ",
	}
	font[120] = []string{
		"        ",
		"   _  __",
		"  | |/_/",
		" _>  <  ",
		"/_/|_|  ",
		"        ",
	}
	font[121] = []string{
		"         ",
		"   __  __",
		"  / / / /",
		" / /_/ / ",
		" \\__, /  ",
		"/____/   ",
	}
	font[122] = []string{
		"     ",
		" ____",
		"/_  /",
		" / /_",
		"/___/",
		"     ",
	}
	font[123] = []string{
		"     __",
		"   _/_/",
		" _/_/  ",
		"< <    ",
		"/ /    ",
		"\\_\\    ",
	}
	font[124] = []string{
		"     __",
		"    / /",
		"   / / ",
		"  / /  ",
		" / /   ",
		"/_/    ",
	}
	font[125] = []string{
		"     _ ",
		"    | |",
		"    / /",
		"   _>_>",
		" _/_/  ",
		"/_/    ",
	}
	font[126] = []string{
		"  /\\//",
		" //\\/ ",
		"  $   ",
		" $    ",
		"$     ",
		"      ",
	}

	if v, ok := font[char]; ok {
		return v
	}
	return font[32]
}
func String(text string) string {
	if len(text) == 0 {
		return ""
	}

	lines := strings.Split(text, "\n")
	result := make([]string, 0)

	for _, line := range lines {
		if len(line) == 0 {
			result = append(result, "")
			continue
		}

		chars := []byte(line)
		s := getArt(chars[0])

		for i := 1; i < len(chars); i++ {
			s = fuse(s, getArt(chars[i]))
		}

		result = append(result, s...)
	}

	return strings.Replace(strings.Join(result, "\n"), "$", " ", -1)
}

func fuse(left, right []string) []string {
	var step []int

	for j := 0; j < 6; j++ {
		k := len(left[j]) + len(right[j]) -
			len(strings.TrimRight(left[j], " ")) -
			len(strings.TrimLeft(right[j], " "))
		step = append(step, k)
	}

	sort.Ints(step)
	c := step[0]
	for j := 0; j < 6; j++ {
		cc := ""
		for i := 1; i <= c; i++ {
			lc := left[j][len(left[j])-(c-i+1) : len(left[j])-(c-i)]
			rc := right[j][i-1 : i]
			if lc == " " {
				cc += rc
			} else {
				cc += lc
			}
		}
		left[j] = left[j][:len(left[j])-c] + cc + right[j][c:]
	}
	return left
}

func main() {
	fmt.Println("Enter the string:")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	text = strings.ReplaceAll(text, "\\n", "\n")

	asciiArt := String(text)

	fmt.Println(asciiArt)
}
