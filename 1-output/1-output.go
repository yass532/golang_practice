package main

import (
	"fmt" // <- 出力で使用
	"os"
)

// main関数の定義
func main() {

	// 標準出力 Print()
	fmt.Print("\n<< Print() >>\n") // 「\n」で改行
	fmt.Print("Hello, World!!\n")
	fmt.Print("\n") // 改行

	// Printf()
	// - フォーマット（型）の指定が可能（「%d」「%f」「%s」など）
	fmt.Printf("<< Printf() >>\n")
	fmt.Printf("%d %s %d %s %f\n", 1, "/", 3, "=", 0.333333)
	fmt.Printf("\n") // 改行

	// Println()
	// - オペランド（要素と要素の間）の間に半角スペース
	// - 文字列の最後に改行が追加
	fmt.Println("<< Println() >>")
	fmt.Println("要素", "要素")
	fmt.Println("") // 改行

	// FPrint()
	// 書き込み先（出力先）を指定可能
	fmt.Fprint(os.Stdout, "<< F >>\n")
	fmt.Fprint(os.Stdout, "Fprint()   : Hello, world!!\n")
	fmt.Fprintf(os.Stdout, "Fprintf()  : %d %s %d %s %f\n", 1, "/", 3, "=", 0.333333)
	fmt.Fprintln(os.Stdout, "Fprintln() :", "要素", "要素")
	fmt.Fprint(os.Stdout, "\n") // 改行

	// SPrint()
	// 出力はせずに、文字列を返す
	output1 := fmt.Sprint("<< S >>\n")
	output2 := fmt.Sprint("Sprint()   : Hello, world!!\n")
	output3 := fmt.Sprintf("Sprintf()  : %d %s %d %s %f\n", 1, "/", 3, "=", 0.333333)
	output4 := fmt.Sprintln("Sprintln() :", "要素", "要素")
	output5 := fmt.Sprint("\n") // 改行
	fmt.Print(output1)
	fmt.Print(output2)
	fmt.Print(output3)
	fmt.Print(output4)
	fmt.Print(output5)
}
