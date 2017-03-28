package controller

import (
	"bytes"
	"fmt"
	"github.com/howeyc/fsnotify"
	"log"
	"os/exec"
)

type monitor struct {
	watch *fsnotify.Watcher
}

func main() {
	M, err := NewMonitor()
	if err != nil {
		log.Println(err)
		return
	}
	M.Do()
	M.watch.Watch("./")
	select {}
}
func NewMonitor() (monitor, error) {
	Mon, err := fsnotify.NewWatcher()
	return monitor{Mon}, err
}
func (self monitor) Do() {
	go func() {
		for {
			select {
			case w := <-self.watch.Event:
				log.Println(w)
				if w.IsModify() {
					fmt.Println("is modify")
					go KillServer()
					go StartServer()
					continue
				}
				if w.IsDelete() {
					log.Println("文件被删除.")
					continue
				}
				if w.IsRename() {
					w = <-self.watch.Event
					log.Println(w)
					self.watch.RemoveWatch(w.Name)
					log.Println(w.Name, " 被重命名.")
				}
			case err := <-self.watch.Error:
				log.Fatalln(err)
			}
		}
	}()
}

func KillServer() {
	// cmd1 := exec.Command("/bin/sh", "-c", "kill -HUP `ps -ef |grep myserver|awk '{print $2}'`")
	cmd1 := exec.Command("killall", "-q", "-w", "-9", "myserver")
	var out bytes.Buffer
	cmd1.Stdout = &out
	err := cmd1.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Status of server: %q\n", out.String())
	StartServer()
}

func StartServer() {
	cmd2 := exec.Command("go", "run", "myserver.go")
	var out bytes.Buffer
	cmd2.Stdout = &out
	if err := cmd2.Run(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("status of server", out.String())

}
