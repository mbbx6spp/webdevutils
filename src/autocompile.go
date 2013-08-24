package main

import (
	"flag"
	"github.com/howeyc/fsnotify"
	"log"
	"os"
	"strings"
	"syscall"
	"webdevutils"
)

type command struct {
	path   string
	args   string
	stdout *os.File
	stderr *os.File
}

func newCommand(path, args, logfile string) command {
	var f *os.File
	_, err := os.Stat(logfile)
	if err != nil {
		f, err = os.Create(logfile)
		if err != nil {
			log.Println("Cannot log stdout and stderr to logfile: " + logfile)
			log.Println(err)
		}
	} else {
		f, err = os.OpenFile(logfile, os.O_APPEND, 0750)
	}
	return command{
		path:   path,
		args:   args,
		stdout: f,
		stderr: f,
	}
}

func (c command) execute(evt fsnotify.FileEvent) (*os.Process, error) {
	var procAttr os.ProcAttr
	procAttr.Files = []*os.File{nil, c.stdout, c.stderr}
	args := strings.Replace(c.args, "%{file}", evt.Name, -1)
	argList := strings.Fields(args)
	result, err := os.StartProcess(c.path, argList, &procAttr)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func signaltrap(c <-chan os.Signal) {
	select {
	case s := <-c:
		switch s {
		case syscall.SIGINT:
			log.Println("Received SIGINT. Shutting down...")
		case syscall.SIGTERM:
			log.Println("Received SIGTERM. Shutting down...")
		default:
			return
		}
	}
}

func autorunner(watcher *fsnotify.Watcher, cmd command) <-chan bool {
	c := make(chan bool)
	go func(done chan<- bool) {
		for {
			select {
			case ev := <-watcher.Event:
				log.Println("Event: ", ev)
				_, err := cmd.execute(*ev)
				if err != nil {
					log.Fatal(err)
				}
			case err := <-watcher.Error:
				log.Println("Error: ", err)
				done <- true
			}
		}
	}(c)
	return c
}

func main() {
	var dir, cmd, args, logfile string

	flag.StringVar(&dir, "directory", "",
		"Directory to watch for filesystem events")
	flag.StringVar(&cmd, "command", "",
		"Command to run every time a filesystem event occurs")
	flag.StringVar(&args, "arguments", "",
		"Arguments for the command")
	flag.StringVar(&logfile, "logfile", "autocompile.log",
		"Log file to redirect stdout and stderr steams of command with args to")
	flag.Parse()

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	done := autorunner(watcher, newCommand(cmd, args, logfile))

	err = watcher.Watch(dir)
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// ignore channel returned as we will not unregister the signal traps
	webdevutils.RegisterSignalTrap(signaltrap)

	<-done
}
