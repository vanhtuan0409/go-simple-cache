package main

import (
	"fmt"

	"github.com/vanhtuan0409/go-simple-cache"
)

func main() {
	lruCache := cache.NewLRUCache(2)
	lruCache.Set("x", 1)
	lruCache.Set("y", 2)
	lruCache.Get("x")
	lruCache.Get("x")
	lruCache.Get("x")
	lruCache.Get("y")
	lruCache.Set("z", 3)
	fmt.Println("Result:")
	printValue(lruCache, "x")
	printValue(lruCache, "y")
	printValue(lruCache, "z")

	fmt.Println("=====")

	lfuCache := cache.NewLFUCache(2)
	lfuCache.Set("x", 1)
	lfuCache.Set("y", 2)
	lfuCache.Get("x")
	lfuCache.Get("x")
	lfuCache.Get("x")
	lfuCache.Get("y")
	lfuCache.Set("z", 3)
	fmt.Println("Result:")
	printValue(lfuCache, "x")
	printValue(lfuCache, "y")
	printValue(lfuCache, "z")
}

func printValue(cache cache.Cache, key string) {
	value, err := cache.Get(key)
	if err == nil {
		fmt.Printf("Value of %s is %v\n", key, value)
	} else {
		fmt.Printf("Error when getting key %s from cache: %v\n", key, err)
	}
}
