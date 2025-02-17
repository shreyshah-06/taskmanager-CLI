package db

import (
	"encoding/binary"
	"time"

	"github.com/boltdb/bolt"
)

var taskBucket = []byte("tasks")

var db *bolt.DB

type Task struct {
	Key         int
	Value       string
	CompletedAt int64
}

func Init(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})
}

func CreateTask(task string) (int, error) {
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		id64, _ := b.NextSequence()
		id = int(id64)
		key := itob(id)

		taskData := encodeTask(Task{Key: id, Value: task, CompletedAt: 0})
		return b.Put(key, taskData)
	})
	if err != nil {
		return -1, err
	}
	return id, nil
}

func AllTasks() ([]Task, error) {
	var tasks []Task
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			task := decodeTask(v)
			if task.CompletedAt == 0 { // Only show pending tasks
				tasks = append(tasks, task)
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func DeleteTask(key int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		return b.Delete(itob(key))
	})
}

func CompleteTask(key int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		taskData := b.Get(itob(key))
		if taskData == nil {
			return nil
		}
		task := decodeTask(taskData)
		task.CompletedAt = time.Now().Unix()
		return b.Put(itob(key), encodeTask(task))
	})
}

func CompletedTasks() ([]Task, error) {
	var tasks []Task
	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).Unix()

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			task := decodeTask(v)
			if task.CompletedAt >= startOfDay {
				tasks = append(tasks, task)
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

// Encoding and Decoding Helpers
func encodeTask(task Task) []byte {
	b := make([]byte, 16+len(task.Value))
	binary.BigEndian.PutUint64(b[:8], uint64(task.Key))
	binary.BigEndian.PutUint64(b[8:16], uint64(task.CompletedAt))
	copy(b[16:], []byte(task.Value))
	return b
}

func decodeTask(data []byte) Task {
	key := int(binary.BigEndian.Uint64(data[:8]))
	completedAt := int64(binary.BigEndian.Uint64(data[8:16]))
	value := string(data[16:])
	return Task{Key: key, Value: value, CompletedAt: completedAt}
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}