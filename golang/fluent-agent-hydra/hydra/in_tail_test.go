package hydra_test

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/fujiwara/fluent-agent-hydra/hydra"
)

var (
	EOFMarker      = "__EOF__"
	RotateMarker   = "__ROTATE__"
	TruncateMarker = "__TRUNCATE__"
	Logs           = []string{
		"single line\n",
		"multi line 1\nmulti line 2\nmultiline 3\n",
		"continuous line 1",
		"continuous line 2\n",
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa\n",   // 80 bytes == hydra.ReadBufferSize for testing
		"bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb\n",  // 81 bytes
		"ccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccc\n", // 82byte
		"dddddddddddddddddddddddddddddddddddddddd",
		"eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee\n", // continuous line 80 bytes
		RotateMarker + "\n",
		"foo\n",
		"bar\n",
		"baz\n",
		TruncateMarker + "\n",
		"FOOOO\n",
		"BAAAR\n",
		"BAZZZZZZZ\n",
		EOFMarker + "\n",
	}
)

const (
	ReadBufferSizeForTest = 80
)

func TestTrail(t *testing.T) {
	hydra.ReadBufferSize = ReadBufferSizeForTest

	tmpdir, _ := ioutil.TempDir(os.TempDir(), "hydra-test")
	file, _ := ioutil.TempFile(tmpdir, "logfile.")
	defer os.RemoveAll(tmpdir)

	go fileWriter(t, file, Logs)
}

func fileWriter(t *testing.T, file *os.File, logs []string) {
	filename := file.Name()
	time.Sleep(1 * time.Second)
	for _, line := range logs {
		if strings.Index(line, RotateMarker) != -1 {
			log.Println("fileWriter: rename file => file.old")
			os.Rename(filename, filename+".old")
			file.Close()
			file, _ = os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
			log.Println("fileWriter: re-opened file")
		} else if strings.Index(line, TruncateMarker) != -1 {
			time.Sleep(1 * time.Second)
			log.Println("fileWriter: truncate(file, 0)")
			os.Truncate(filename, 0)
			file.Seek(int64(0), os.SEEK_SET)
		}
		_, err := file.WriteString(line)
		log.Print("fileWriter: wrote ", line)
		if err != nil {
			log.Println("write failed", err)
		}
		time.Sleep(1 * time.Millisecond)
	}
	file.Close()
}
