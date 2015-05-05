package main

// https://github.com/go-fsnotify/fsnotify/blob/master/example_test.go
// https://github.com/fujiwara/fluent-agent-hydra

import (
	"bytes"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/howeyc/fsnotify"
)

const Interval = 200 * time.Millisecond
const ReadBufferSize = 2
const SEEK_TAIL = int64(-1)
const SEEK_HEAD = int64(0)

var LineSeparator = []byte{'\n'}

type File struct {
	*os.File
	Path     string
	ReadBuf  []byte
	LineBuf  []byte
	Position int64
	lastStat os.FileInfo
}

func (f *File) BufferedLineRead() error {
	s, err := f.Stat()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	f.lastStat = s
	log.Printf("size: %d", f.lastStat.Size())
	if f.lastStat.Size() == 0 {
		f.Seek(0, 0)
		f.Position = int64(0)
		log.Print("size becomes 0")
	}

	for {
		n, err := io.ReadAtLeast(f, f.ReadBuf, 1)
		if n == 0 || err == io.EOF {
			return err
		} else if err != nil {
			return err
		}
		f.Position += int64(n)
		sendBuf := make([]byte, 0)
		if f.ReadBuf[n-1] == '\n' {
			// f.readBuf is just terminated by '\n'
			if len(f.LineBuf) > 0 {
				sendBuf = append(sendBuf, f.LineBuf...)
				f.LineBuf = make([]byte, 0)
			}
			sendBuf = append(sendBuf, f.ReadBuf[0:n-1]...)
		} else {
			blockLen := bytes.LastIndex(f.ReadBuf[0:n], LineSeparator)
			if blockLen == -1 {
				// whole of f.readBuf is continuous line
				f.LineBuf = append(f.LineBuf, f.ReadBuf[0:n]...)
				continue
			} else {
				// bottom line of f.readBuf is continuous line
				if len(f.LineBuf) > 0 {
					sendBuf = append(sendBuf, f.LineBuf...)
				}
				sendBuf = append(sendBuf, f.ReadBuf[0:blockLen]...)
				f.LineBuf = make([]byte, n-blockLen-1)
				copy(f.LineBuf, f.ReadBuf[blockLen+1:n])
			}
		}
		log.Printf(
			"[pos:(%d), size:(%d)]: %s\n",
			f.Position,
			f.lastStat.Size(),
			string(sendBuf))
	}
}

func NewTargetFile(path string) (*File, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	stat, err := f.Stat()
	if err != nil {
		return nil, err
	}

	file := &File{
		f,
		path,
		make([]byte, ReadBufferSize),
		make([]byte, 0),
		0,
		stat,
	}

	startPos := int64(0)
	if startPos == SEEK_TAIL {
		// seek to end of file
		size := file.lastStat.Size()
		pos, _ := file.Seek(size, os.SEEK_SET)
		file.Position = pos
	} else {
		pos, _ := file.Seek(startPos, os.SEEK_SET)
		file.Position = pos
	}
	log.Println("[info]", file.Path, "Seeked to", file.Position)
	return file, nil
}

func main() {
	if len(os.Args) != 2 {
		log.Println("Usage: go run " + filepath.Base(os.Args[0]) + ".go [file_path]")
		os.Exit(1)
	}

	filePath := os.Args[1]

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer watcher.Close()

	file, err := NewTargetFile(filePath)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Event:
				if event.IsRename() || event.IsDelete() {
					log.Println("file closed")
					file.Close()
				}
				if event.IsModify() {
					log.Println("modified")
					err = file.BufferedLineRead()
					if err != nil {
						log.Printf("%s", err)
					}
				}
				if event.IsCreate() {
					log.Println("created")
				}
			}
		}
	}()

	err = watcher.Watch(filePath)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	<-done
}
