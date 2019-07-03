package ryutil

import (
	"fmt"
	"os"
	"log"
	"io"
	"time"
)

/*
var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)
*/
func Info(str string) {
	//	fmt.Print(time.Now().Format("2006-01-02 15:04:05.000"), " INFO ▶ ")
	//	fmt.Println(fmt.Sprintf(str))
	fileWriter("INFO ▶ " + fmt.Sprintf(str))
}

func Infof(format string, a ...interface{}) {
	//fmt.Print(time.Now().Format("2006-01-02 15:04:05.000"), " INFO ▶ ")
	//fmt.Println(fmt.Sprintf(format, a))
	fileWriter("INFO ▶ " + fmt.Sprintf(format, a))
}
func Error(str string) {
	//fmt.Print(time.Now().Format("2006-01-02 15:04:05.000"), " ERROR ▶ ")
	//fmt.Println(fmt.Sprintf(str))
	fileWriter("ERROR ▶ " + fmt.Sprintf(str))
}

func Errorf(format string, a ...interface{}) {
	//fmt.Print(time.Now().Format("2006-01-02 15:04:05.000"), " ERROR ▶ ")
	//fmt.Println(fmt.Sprintf(format, a))
	fileWriter("ERROR ▶ " + fmt.Sprintf(format, a))
}
func Debugf(tag string, format string, a ...interface{}) {
	//fmt.Print(time.Now().Format("2006-01-02 15:04:05.000"), fmt.Sprintf(" %s ▶", tag))
	//fmt.Println(fmt.Sprintf(format, a))result is ->
	fileWriter(fmt.Sprintf("%s ▶", tag) + fmt.Sprintf(format, a))
}

func Warning(format interface{}, a ... interface{}) {

}

func Debug(format interface{}) {

}

func fileWriter(line string) {
	f, err := os.OpenFile("daily.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {

		if os.IsNotExist(err) {
			log.Fatal("File does not exist.")
		} else {
			log.Fatalf("error opening file: %v", err)
		}

	}
	defer f.Close()
	wrt := io.MultiWriter(os.Stdout, f)
	log.SetOutput(wrt)
	log.Println(line)
}

func CutLogs() {
	preparedtime := time.Now().Format("2006-01-02-15-04")
	fi, err := os.Stat("daily.log")
	if err != nil {
		return
	}
	// get the size
	size_big := fi.Size()

	size_big_flo := float64(size_big)
	kilobytes := size_big_flo / 1024.0
	megabytes := kilobytes / 1024.0

	Infof("mega -> %f", megabytes)

	if megabytes > 6.3 {
		Info("==========================END FILE HERE==========================")
		// Open original file
		originalFile, err := os.Open("daily.log")
		if err != nil {
			log.Fatal(err)
		}
		defer originalFile.Close()

		// Create new file
		newFile, err := os.Create("log/c" + preparedtime + ".log")
		if err != nil {
			log.Fatal(err)

		}
		defer newFile.Close()

		// Copy the bytes to destination from source
		bytesWritten, err := io.Copy(newFile, originalFile)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Copied %d bytes.", bytesWritten)

		// Commit the file contents
		// Flushes memory to disk
		err = newFile.Sync()
		if err != nil {
			log.Fatal(err)
		}

		err = os.Truncate("daily.log", 0)
		if err != nil {
			log.Fatal(err)
		}

	}

}
