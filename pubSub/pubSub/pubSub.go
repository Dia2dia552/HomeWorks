package pubSub

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/go-redis/redis/v8"
	"os"
	"path/filepath"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Publisher struct {
	subscribers map[chan string]struct{}
	redisClient *redis.Client
}

func NewPublisher() *Publisher {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	return &Publisher{
		subscribers: make(map[chan string]struct{}),
		redisClient: rdb,
	}
}

func (p *Publisher) GetDataFromDB() ([]User, error) {
	cachedData, err := p.redisClient.Get(context.Background(), "cachedData").Result()
	if err == nil {
		var users []User
		if err := json.Unmarshal([]byte(cachedData), &users); err == nil {
			fmt.Println("Data fetched from cache")
			return users, nil
		}
	}

	users, err := getAllUsersFromDB()
	if err != nil {
		return nil, err
	}

	jsonData, err := json.Marshal(users)
	if err != nil {
		return nil, err
	}
	p.redisClient.Set(context.Background(), "cachedData", jsonData, 0)

	fmt.Println("Data fetched from the database")
	return users, nil
}

func getAllUsersFromDB() ([]User, error) {
	db, err := sql.Open("postgres", "user=username password=password dbname=mydb sslmode=disable")
	if err != nil {
		return nil, err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	rows, err := db.Query("SELECT id, name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
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
	defer func(watcher *fsnotify.Watcher) {
		err := watcher.Close()
		if err != nil {

		}
	}(watcher)

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
