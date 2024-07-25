package main

import "fmt"

const pi = 3.141592
const e = 2.718281
const (
	IronMan int = iota
	DrStrange
	Thor
	Hulk
)

const GlobalLimit = 100
const MaxCacheSize int = 2 * GlobalLimit
const (
	CacheKeyBook = "book_"
	CacheKeyCD   = "cd_"
)

var cache map[string]string

func cacheGet(key string) string {
	return cache[key]
}

func cacheSet(key, val string) {
	if len(cache)+1 >= MaxCacheSize {
		fmt.Println("The cache has reached the maximum size")
		return
	}
	cache[key] = val
}

func SetBook(isbn string, name string) {
	cacheSet(CacheKeyBook+isbn, name)
}

func GetBook(isbn string) string {
	return cacheGet(CacheKeyBook + isbn)
}

func SetCD(sku string, title string) {
	cacheSet(CacheKeyCD+sku, title)
}

func GetCD(sku string) string {
	return cacheGet(CacheKeyCD + sku)
}

func main() {

	fmt.Println("pi:", pi, "e:", e)
	fmt.Println(IronMan, DrStrange, Thor, Hulk)
	fmt.Printf("-------------------------------------------\n")

	cache = make(map[string]string)

	SetBook("1234-5678", "3 Body Problem")
	SetCD("1234-5678", "How I met your mother")

	fmt.Println("Book:", GetBook("1234-5678"))
	fmt.Println("CD:", GetCD("1234-5678"))
}
