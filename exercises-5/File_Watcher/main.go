package main

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
)

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil{
		panic(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func(){
		for{
			select{
			case event,ok := <- watcher.Events:
				if !ok{
					return
				}
				log.Println("event:", event)

				if event.Op&fsnotify.Create == fsnotify.Create {
					fmt.Println("File created:", event.Name)
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					fmt.Println("File modified:", event.Name)
				}
				if event.Op&fsnotify.Remove == fsnotify.Remove {
					fmt.Println("File deleted:", event.Name)
				}

			case err,ok := <- watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
				 
			}
		}
	}()

	err = watcher.Add("D:\\go\\src\\.go\\Go_MasterClass\\exercises-5\\File_Watcher\\watched_dir")
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Watching directory:")

	<-done

}