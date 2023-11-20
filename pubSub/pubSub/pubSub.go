package pubSub

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"os"
	"path/filepath"
)

type Publisher struct {
	subscribers map[chan string]struct{}
}

func NewPublisher() *Publisher {
	return &Publisher{
		subscribers: make(map[chan string]struct{}),
	}
}

func (p *Publisher) Subscribe(c chan string) {
	p.subscribers[c] = struct{}{}
}

func (p *Publisher) Unsubscribe(c chan string) {
	delete(p.subscribers, c)
}

func (p *Publisher) NotifyAll(event string) {
	for c := range p.subscribers {
		c <- event
	}
}

func StartMonitoring(dirPath string, publisher *Publisher) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer watcher.Close()

	err = filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			err := watcher.Add(path)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					publisher.NotifyAll(fmt.Sprintf("File %s was modified", event.Name))
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				fmt.Println("Error:", err)
			}
		}
	}()
}
