package main

import (
	"flag"
	"fmt"
	"os"

	"training.go/dictionary/dictionary"
)

func main() {

	action := flag.String("action", "list", "Action to perform with db")
	d, err := dictionary.New("badger")
	handlerError(err)
	defer d.Close()

	flag.Parse()
	switch *action {
	case "list":
		actionList(d)
	case "get":
		actionGet(d, flag.Args())
	case "add":
		actionAdd(d, flag.Args())
	case "remove":
		actionRemove(d, flag.Args())
	case "define":
		actionDefine(d, flag.Args())
	default:
		fmt.Printf("action inconnu %v \n", *action)
	}
}

func actionList(d *dictionary.Dictionary) {
	words, mapEntries, err := d.List()
	handlerError(err)
	fmt.Println("liste du dictionaire")
	for _, word := range words {
		fmt.Println(mapEntries[word])
	}
}

func actionGet(d *dictionary.Dictionary, args []string) {
	word := args[0]
	entry, err := d.Read(word)
	handlerError(err)
	fmt.Println(entry)
}

func actionDefine(d *dictionary.Dictionary, args []string) {
	word := args[0]
	entry, err := d.Read(word)
	handlerError(err)
	fmt.Println(entry.Definition)
}

func actionRemove(d *dictionary.Dictionary, args []string) {
	word := args[0]
	err := d.Delete(word)
	handlerError(err)
}

func actionAdd(d *dictionary.Dictionary, args []string) {
	word := args[0]
	definition := args[1]
	if len(args) != 2 {
		fmt.Printf("2 args obligatoire pour cette commande ils y en a %v", args)
		os.Exit(1)
	}
	err := d.Add(word, definition)
	handlerError(err)
}

func handlerError(err error) {
	if err != nil {
		fmt.Printf("Erreur dans la base de données Dictionary %v", err)
		os.Exit(1)
	}
}
