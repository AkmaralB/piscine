package piscine

import (
	"github.com/01-edu/z01"
)

func EightQueens() {
	for a := '1'; a <= '8'; a++ {
		for b := '1'; b <= '8'; b++ {
			if b != a && b != a+1 && b != a-1 {
				for c := '1'; c <= '8'; c++ {
					if c != a && c != a+2 && c != a-2 &&
						c != b && c != b+1 && c != b-1 {
						for d := '1'; d <= '8'; d++ {
							if d != a && d != b && d != c &&
								d != a+3 && d != b+2 && d != c+1 &&
								d != a-3 && d != b-2 && d != c-1 {
								for e := '1'; e <= '8'; e++ {
									if e != a && e != b && e != c && e != d &&
										e != a+4 && e != b+3 && e != c+2 && e != d+1 &&
										e != a-4 && e != b-3 && e != c-2 && e != d-1 {
										for f := '1'; f <= '8'; f++ {
											if f != a && f != b && f != c && f != d && f != e &&
												f != a+5 && f != b+4 && f != c+3 && f != d+2 && f != e+1 &&
												f != a-5 && f != b-4 && f != c-3 && f != d-2 && f != e-1 {
												for g := '1'; g <= '8'; g++ {
													if g != a && g != b && g != c && g != d && g != e && g != f &&
														g != a+6 && g != b+5 && g != c+4 && g != d+3 && g != e+2 && g != f+1 &&
														g != a-6 && g != b-5 && g != c-4 && g != d-3 && g != e-2 && g != f-1 {
														for h := '1'; h <= '8'; h++ {
															if h != a && h != b && h != c && h != d && h != e && h != f && h != g &&
																h != a+7 && h != b+6 && h != c+5 && h != d+4 && h != e+3 && h != f+2 && h != g+1 &&
																h != a-7 && h != b-6 && h != c-5 && h != d-4 && h != e-3 && h != f-2 && h != g-1 {
																z01.PrintRune(a)
																z01.PrintRune(b)
																z01.PrintRune(c)
																z01.PrintRune(d)
																z01.PrintRune(e)
																z01.PrintRune(f)
																z01.PrintRune(g)
																z01.PrintRune(h)
																z01.PrintRune('\n')
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
