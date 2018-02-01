package main

import (
	"fmt"
	"strings"
	"unicode"
)

func Len() {
	fmt.Printf("\n\nlen:\n")
	fmt.Println("Len(\"Go\") = ", len("Go"))
	fmt.Println("Len(\"张三\") = ", len("张三"))
	fmt.Println("Len(\"Go张三\") = ", len("Go张三"))
}

func EqualFold() {
	fmt.Printf("\n\nstrings.EqualFold:\n")
	fmt.Println("strings.EqualFold(\"Go\", \"go\") = ", strings.EqualFold("Go", "go"))
	fmt.Println("strings.EqualFold(\"张三\", \"张三\") = ", strings.EqualFold("张三", "张三"))
	fmt.Println("strings.EqualFold(\"Go张三\", \"go张三\") = ", strings.EqualFold("Go张三", "Go张三"))
}

func HasPrefix() {
	fmt.Printf("\n\nstrings.HasPrefix:\n")
	fmt.Println("strings.HasPrefix(\"Go\", \"go\") = ", strings.HasPrefix("Go", "go"))
	fmt.Println("strings.HasPrefix(\"Go\", \"G\") = ", strings.HasPrefix("Go", "G"))
	fmt.Println("strings.HasPrefix(\"Go\", \"Go\") = ", strings.HasPrefix("Go", "Go"))
	fmt.Println("strings.HasPrefix(\"Go\", \"Goo\") = ", strings.HasPrefix("Go", "Goo"))
	fmt.Println("strings.HasPrefix(\"Go\", \"\") = ", strings.HasPrefix("Go", ""))
}

func HasSuffix() {
	fmt.Printf("\n\nstrings.HasSuffix:\n")
	fmt.Println("strings.HasSuffix(\"Go\", \"go\") = ", strings.HasSuffix("Go", "go"))
	fmt.Println("strings.HasSuffix(\"Go\", \"o\") = ", strings.HasSuffix("Go", "o"))
	fmt.Println("strings.HasSuffix(\"Go\", \"Go\") = ", strings.HasSuffix("Go", "Go"))
	fmt.Println("strings.HasSuffix(\"Go\", \"gGo\") = ", strings.HasSuffix("Go", "gGo"))
	fmt.Println("strings.HasSuffix(\"Go\", \"\") = ", strings.HasSuffix("Go", ""))
}

func Contains() {
	fmt.Printf("\n\nstrings.Contains:\n")
	fmt.Println("strings.Contains(\"Go\", \"go\") = ", strings.Contains("Go", "go"))
	fmt.Println("strings.Contains(\"Go\", \"o\") = ", strings.Contains("Go", "o"))
	fmt.Println("strings.Contains(\"Go\", \"Go\") = ", strings.Contains("Go", "Go"))
	fmt.Println("strings.Contains(\"Go\", \"gGo\") = ", strings.Contains("Go", "gGo"))
	fmt.Println("strings.Contains(\"Go\", \"\") = ", strings.Contains("Go", ""))
}

func ContainsRune() {
	fmt.Printf("\n\nstrings.ContainsRune:\n")
	fmt.Println("strings.ContainsRune(\"Go\", \"G\") = ", strings.ContainsRune("Go", 'G'))
	fmt.Println("strings.ContainsRune(\"Go\", \"g\") = ", strings.ContainsRune("Go", 'g'))
	fmt.Println("strings.ContainsRune(\"Go世界\", \"世\") = ", strings.ContainsRune("Go世界", '世'))
}

func ContainsAny() {
	fmt.Printf("\n\nstrings.ContainsAny:\n")
	fmt.Println("strings.ContainsAny(\"Go\", \"go\") = ", strings.ContainsAny("Go", "go"))
	fmt.Println("strings.ContainsAny(\"Go\", \"o\") = ", strings.ContainsAny("Go", "o"))
	fmt.Println("strings.ContainsAny(\"Go\", \"Go\") = ", strings.ContainsAny("Go", "Go"))
	fmt.Println("strings.ContainsAny(\"Go\", \"gGo\") = ", strings.ContainsAny("Go", "gGo"))
	fmt.Println("strings.ContainsAny(\"Go\", \"\") = ", strings.ContainsAny("Go", ""))
	fmt.Println("strings.ContainsAny(\"Go张三\", \"三\") = ", strings.ContainsAny("Go张三", "三"))
}

func Count() {
	fmt.Printf("\n\nstrings.Count:\n")
	fmt.Println("strings.Count(\"Go\", \"go\") = ", strings.Count("Go", "go"))
	fmt.Println("strings.Count(\"Go\", \"Go\") = ", strings.Count("Go", "Go"))
	fmt.Println("strings.Count(\"Google\", \"o\") = ", strings.Count("Google", "o"))
	fmt.Println("strings.Count(\"Google\", \"g\") = ", strings.Count("Google", "g"))
	fmt.Println("strings.Count(\"Google\", \"\") = ", strings.Count("Google", ""))
	fmt.Println("strings.Count(\"张三三\", \"三\") = ", strings.Count("张三三", "三"))
}

func Index() {
	fmt.Printf("\n\nstrings.Index:\n")
	fmt.Println("strings.Index(\"Go\", \"go\") = ", strings.Index("Go", "go"))
	fmt.Println("strings.Index(\"Go\", \"Go\") = ", strings.Index("Go", "Go"))
	fmt.Println("strings.Index(\"Google\", \"o\") = ", strings.Index("Google", "o"))
	fmt.Println("strings.Index(\"Google\", \"g\") = ", strings.Index("Google", "g"))
	fmt.Println("strings.Index(\"Google\", \"\") = ", strings.Index("Google", ""))
	fmt.Println("strings.Index(\"张三三\", \"三\") = ", strings.Index("张三三", "三"))
}

func IndexByte() {
	fmt.Printf("\n\nstrings.IndexByte:\n")
	fmt.Println("strings.IndexByte(\"张三三\", '1') = ", strings.IndexByte("张三三", '1'))
}

func IndexRune() {
	fmt.Printf("\n\nstrings.IndexRune:\n")
	fmt.Println("strings.IndexRune(\"张三三\", '三') = ", strings.IndexRune("张三三", '三'))
}

func IndexAny() {
	fmt.Printf("\n\nstrings.IndexAny:\n")
	fmt.Println("strings.IndexAny(\"Google\", \"book\") = ", strings.IndexAny("Google", "book"))
	fmt.Println("strings.IndexAny(\"张三三\", \"三\") = ", strings.IndexAny("张三三", "三四五"))
}

func IndexFunc() {
	isHan := func(r rune) bool { return unicode.Is(unicode.Han, r) }
	isDigit := func(r rune) bool { return unicode.Is(unicode.Digit, r) }
	fmt.Printf("\n\nstrings.IndexFunc:\n")
	fmt.Println("strings.IndexFunc(\"Google\", \"是汉字\") = ", strings.IndexFunc("Google", isHan))
	fmt.Println("strings.IndexFunc(\"Go世界\", \"是汉字\") = ", strings.IndexFunc("Go世界", isHan))
	fmt.Println("strings.IndexFunc(\"世界123\", \"是数字\") = ", strings.IndexFunc("世界123", isDigit))
	fmt.Println("strings.IndexFunc(\"世界\", \"是数字\") = ", strings.IndexFunc("世界", isDigit))
}

func Title() {
	fmt.Printf("\n\nstrings.Title:\n")
	fmt.Println("strings.Title(\"hello world\") = ", strings.Title("hello world"))
	fmt.Println("strings.Title(\"张三 李四\") = ", strings.Title("张三 李四"))
	fmt.Println("strings.Title(\"hello, world\") = ", strings.Title("hello, world"))
}

func ToLower() {
	fmt.Printf("\n\nstrings.ToLower:\n")
	fmt.Println("strings.ToLower(\"Hello World\") = ", strings.ToLower("Hello World"))
	fmt.Println("strings.ToLower(\"张三 李四\") = ", strings.ToLower("张三 李四"))
	fmt.Println("strings.ToLower(\"HELLO, WORLD\") = ", strings.ToLower("HELLO, WORLD"))
}

func ToLowerSpecial() {
	fmt.Printf("\n\nstrings.ToLowerSpecial:\n")
	fmt.Println("strings.ToLowerSpecial(unicode.AzeriCase, \"Hello World\") = ", strings.ToLowerSpecial(unicode.AzeriCase, "Hello World"))
	fmt.Println("strings.ToLowerSpecial(unicode.AzeriCase, \"张三 李四\") = ", strings.ToLowerSpecial(unicode.AzeriCase, "张三 李四"))
	fmt.Println("strings.ToLowerSpecial(unicode.AzeriCase, \"HELLO, WORLD\") = ", strings.ToLowerSpecial(unicode.AzeriCase, "HELLO, WORLD"))
	fmt.Println("0x0049 is ", string('\u0049'))
	fmt.Println("0x0069 is ", string('\u0069'))
	fmt.Println("0x0130 is ", string('\u0130'))
	fmt.Println("0x0131 is ", string('\u0131'))
	fmt.Println("strings.ToLowerSpecial(unicode.AzeriCase, \"HIGHT\") = ", strings.ToLowerSpecial(unicode.AzeriCase, "HIGHT"))
}

func Repeat() {
	fmt.Printf("\n\nstrings.Repeat:\n")
	fmt.Println("strings.Repeat(\"abc\", 3) = ", strings.Repeat("abc", 3))
	fmt.Println("strings.Repeat(\"张三\", 3) = ", strings.Repeat("张三", 3))
}

func Replace() {
	fmt.Printf("\n\nstrings.Replace:\n")
	fmt.Println("strings.Replace(\"abcbca\", \"bc\", \"cd\", 0) = ", strings.Replace("abcbca", "bc", "cd", 0))
	fmt.Println("strings.Replace(\"abcbca\", \"bc\", \"cd\", -1) = ", strings.Replace("abcbca", "bc", "cd", -1))
	fmt.Println("strings.Replace(\"abcbca\", \"\", \"cd\", -1) = ", strings.Replace("abcbca", "", "cd", -1))
	fmt.Println("strings.Replace(\"张三张三张三\", \"张\", \"李\", 1) = ", strings.Replace("张三张三张三", "张", "李", 1))
}

func Map() {
	mapping := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return r + 32
		case r >= 'a' && r <= 'z':
			return r - 32
		}
		return r
	}
	fmt.Printf("\n\nstrings.Map:\n")
	fmt.Println("strings.Map(\"Hello World\", 大小写互换) = ", strings.Map(mapping, "Hello World"))
}

func Trim() {
	fmt.Printf("\n\nstrings.Trim:\n")
	fmt.Println("strings.Trim(\" !!! Achtung! Achtung! !!! \", \"! \") = ", strings.Trim(" !!! Achtung! Achtung! !!! ", "! "))
	fmt.Println("strings.Trim(\" !!! Achtung! Achtung! !!! \", \" !\") = ", strings.Trim(" !!! Achtung! Achtung! !!! ", "! "))
	fmt.Println("strings.Trim(\" !a!! Achtung! Achtung! !a!! \", \" !\") = ", strings.Trim(" !a!! Achtung! Achtung! !a!! ", "! "))
	fmt.Println("strings.Trim(\" !!!  ! !!! \", \" !\") = ", strings.Trim(" !!!  ! !!! ", "! "))
}

func TrimSpace() {
	fmt.Printf("\n\nstrings.TrimSpace:\n")
	fmt.Println("strings.TrimSpace(\" \\t\\n a lone gopher \\n\\t\\r\\n\") = ", strings.TrimSpace(" \t\n a lone gopher \n\t\r\n"))
}

func TrimPrefix() {
	fmt.Printf("\n\nstrings.TrimPrefix:\n")
	fmt.Println("strings.TrimPrefix(\"hello, world\", \"hello\") = ", strings.TrimPrefix("hello, world", "hello"))
	fmt.Println("strings.TrimPrefix(\"hello,hello,world\", \"hello,\") = ", strings.TrimPrefix("hello,hello,world", "hello,"))
	fmt.Println("strings.TrimLeft(\"hello,hello,world\", \"hello,\") = ", strings.TrimLeft("hello,hello,world", "hello,"))
}

func Fields() {
	fmt.Printf("\n\nstrings.Fields:\n")
	fmt.Println("strings.Fields(\" \\t\\n a lone gopher \\n\\t\\r\\n\") = ", strings.Fields(" \t\n a lone gopher \n\t\r\n"))
}

func Split() {
	fmt.Printf("\n\nstrings.Split:\n")
	fmt.Println("strings.Split(\"a,b,c\", \",\") = ", strings.Split("a,b,c", ","))
	fmt.Println("strings.Split(\"a,b,c\", \"a\") = ", strings.Split("a,b,c", "a"))
	fmt.Println("strings.Split(\"a,b,c\", \"\") = ", strings.Split("a,b,c", ""))
}

func SplitN() {
	fmt.Printf("\n\nstrings.SplitN:\n")
	fmt.Println("strings.SplitN(\"a,b,c\", \",\", 2) = ", strings.SplitN("a,b,c", ",", 2))
	fmt.Println("strings.SplitN(\"a,b,c\", \",\", -1) = ", strings.SplitN("a,b,c", ",", -1))
	fmt.Println("strings.SplitN(\"a,b,c\", \",\", 0) = ", strings.SplitN("a,b,c", ",", 0))
}

func SplitAfter() {
	fmt.Printf("\n\nstrings.SplitAfter:\n")
	fmt.Println("strings.SplitAfter(\"a,b,c\", \",\") = ", strings.SplitAfter("a,b,c", ","))
	fmt.Println("strings.SplitAfter(\"a,b,c\", \"a\") = ", strings.SplitAfter("a,b,c", "a"))
	fmt.Println("strings.SplitAfter(\"a,b,c\", \"\") = ", strings.SplitAfter("a,b,c", ""))
}

func Join() {
	fmt.Printf("\n\nstrings.Join:\n")
	fmt.Println("strings.Join([]string{\"a\", \"b\", \"c\"}, \",\") = ", strings.Join([]string{"a", "b", "c"}, ","))
	fmt.Println("strings.Join([]string{\"a\", \"b\", \"c\"}, \"\") = ", strings.Join([]string{"a", "b", "c"}, ""))
}

func main() {
	Len()
	EqualFold()
	HasPrefix()
	HasSuffix()
	Contains()
	ContainsRune()
	ContainsAny()
	Count()
	Index()
	IndexByte()
	IndexRune()
	IndexAny()
	IndexFunc()
	Title()
	ToLower()
	ToLowerSpecial()
	Repeat()
	Replace()
	Map()
	Trim()
	TrimSpace()
	TrimPrefix()
	Fields()
	Split()
	SplitN()
	SplitAfter()
	Join()
}
