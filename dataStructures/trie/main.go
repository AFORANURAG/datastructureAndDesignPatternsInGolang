package main

import (
	impl "trie/implementation"
)

func main() {
	h := &impl.Trie{}
	h.Insert("anurag")
	h.Insert("amit")
	h.Insert("ankit")
	h.Insert("abhishek")
	h.Insert("")

}
