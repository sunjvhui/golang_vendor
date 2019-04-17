package logging

import (
	"errors"
	"io"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"runtime/debug"
	"sort"
	"time"
)

type date struct {
	year  int
	month time.Month
	day   int
}

type simpleFileInfo struct {
	name    string
	modTime time.Time
}

type fileList []simpleFileInfo

// the Logger struct with a series of method for logging, such as: Info(), Warning(), Error(), Exception()
type Logger struct {
	Logger  *log.Logger
	Handler []io.Writer
}

type TimeRotateHandler struct {
	path     string
	filename string
	ext      string
	maxFile  int
}

func (d date) Equal(d2 date) bool {
	if d.year != d2.year {
		return false
	}
	if d.month != d2.month {
		return false
	}
	if d.day != d2.day {
		return false
	}
	return true
}

func (f fileList) Len() int {
	return len(f)
}

func (f fileList) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func (f fileList) Less(i, j int) bool {
	return f[i].modTime.Unix() > f[j].modTime.Unix()
}

// will return a new rotating file Handler
// path: <string>, log dir path
// filename: <string>, base name of log file
// ext: <string>, extension name of log file
// maxFile: <int>, the max count of rotating files
func NewRotateHandler(path string, filename string, ext string, maxFile int) *TimeRotateHandler {

	return &TimeRotateHandler{
		path:     path,
		filename: filename,
		ext:      ext,
		maxFile:  maxFile,
	}
}

func (t *TimeRotateHandler) rotate(dateString string) *os.File {
	var bakPath string
	var newPath string
	var oldPath string
	var sep string

	sep = string(os.PathSeparator)
	oldPath = t.path + sep + t.filename + "." + t.ext
	bakPath = "_" + dateString
	newPath = t.path + sep + t.filename + bakPath + "." + t.ext

	err := os.Rename(oldPath, newPath)
	if os.IsNotExist(err) {
		file, e := os.Create(oldPath)
		if os.IsPermission(e) {
			log.Panic("Have no permission on logging path")
		} else if os.IsTimeout(e) {
			log.Panic("Can not create logging file")
		}
		return file
	} else {
		file, _ := os.OpenFile(oldPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
		return file
	}
}

func (t *TimeRotateHandler) cleanUpLog() error {

	var (
		fl             fileList
		rfl            fileList
		waitRemoveList fileList
	)

	dir, err := ioutil.ReadDir(t.path)
	if err != nil {
		log.Print(err)
		return errors.New("can not read log dir path")
	}

	for _, item := range dir {
		if item.IsDir() {
			// ignore dir, do nothing
		} else {
			reg, _ := regexp.Compile(t.filename)
			result := reg.FindString(item.Name())
			if result != "" {
				if item.Name() != t.filename+"."+t.ext {
					if item.ModTime().Unix() > time.Now().Unix()+60 {
						// mod time > current time, illegal log file, will delete. (60 sec to fault-tolerant)
						rfl = append(rfl, simpleFileInfo{name: item.Name(), modTime: item.ModTime()})
					} else {
						fl = append(fl, simpleFileInfo{name: item.Name(), modTime: item.ModTime()})
					}
				}
			}
		}
	}

	sort.Sort(fl)

	if len(fl) > t.maxFile {
		waitRemoveList = append(fl[t.maxFile:], rfl...)
		log.Print(waitRemoveList)
		for _, f := range waitRemoveList {
			err := os.Remove(t.path + string(os.PathSeparator) + f.name)
			if err != nil {
				log.Print(err)
			}
		}
	}
	return nil
}

func (t *TimeRotateHandler) Write(p []byte) (n int, err error) {
	var (
		nowTimeStamp  int64
		unixTimeStamp int64
		f             *os.File
		e             error
		oldPath       string
		sep           string
		modTime       time.Time
		modDate       date
		nowDate       date
	)
	nowTimeStamp = time.Now().Unix()
	sep = string(os.PathSeparator)
	oldPath = t.path + sep + t.filename + "." + t.ext

	_, e = ioutil.ReadDir(t.path)
	if e != nil {
		e = os.MkdirAll(t.path, 0755)
		if e !=nil {
			log.Print(e)
		}
	}
	t.cleanUpLog()

	f, err = os.OpenFile(oldPath, os.O_RDONLY, 0755)
	if err == nil || !os.IsNotExist(err) {
		fileInfo, e := f.Stat()
		if e != nil {
			log.Print(e)
		}
		modTime = fileInfo.ModTime()
		unixTimeStamp = fileInfo.ModTime().Unix()
	} else {
		unixTimeStamp = 0
	}
	f.Close()

	if unixTimeStamp != 0 {
		modDate.year, modDate.month, modDate.day = time.Unix(unixTimeStamp, 0).Date()
		nowDate.year, nowDate.month, nowDate.day = time.Unix(nowTimeStamp, 0).Date()

		if modDate.Equal(nowDate) {
			f, e = os.OpenFile(oldPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
			if e != nil {
				if os.IsNotExist(e) {
					log.Panic("can not open log file")
				} else if os.IsPermission(e) {
					log.Panic("lack of permission to open log file")
				} else {
					log.Print(e)
				}
			}
		} else {
			f = t.rotate(modTime.Format("2006_01_02"))
		}
	} else {
		f, e = os.Create(oldPath)
		if e != nil {
			if os.IsExist(e) {
				log.Panic("log file already exists, and can not be opened")
			} else if os.IsPermission(e) {
				log.Panic("lack of permission to create new log file")
			} else {
				log.Print(e)
			}
		}
	}

	if f != nil {
		fileLen, _ := f.Seek(0, 2)
		f.Seek(fileLen, 0)
		n, err := f.Write(p)
		f.Close()
		return n, err
	} else {
		return 0, errors.New("can not create file descriptor")
	}
}

// will initialize the Logger with default log format and default output <os.Stdout>
// you can set the Handler list by using <SetHandler>, a rotating Handler was already provided above
func (l *Logger) Init() {
	l.Logger = log.New(os.Stdout, "[INFO]", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
}

func (l *Logger) SetHandler(handler ...io.Writer) {
	l.Handler = handler
}

func (l *Logger) Exception(v error) {
	l.Logger.SetPrefix("[ERROR]")
	for _, w := range l.Handler {
		l.Logger.SetOutput(w)
		l.Logger.Printf("\n\tException: %s\n%s", v.Error(), debug.Stack())
	}
}

func (l *Logger) log(v ...interface{}) {
	if len(v) > 1 {
		if str, ok := v[0].(string); ok {
			for _, w := range l.Handler {
				l.Logger.SetOutput(w)
				l.Logger.Printf(str, v[1:]...)
			}
		} else {
			for _, w := range l.Handler {
				l.Logger.SetOutput(w)
				l.Logger.Print(v...)
			}
		}
	} else {
		for _, w := range l.Handler {
			l.Logger.SetOutput(w)
			l.Logger.Print(v...)
		}
	}
}

func (l *Logger) Info(v ...interface{}) {
	l.Logger.SetPrefix("[INFO]")
	l.log(v...)
}

func (l *Logger) Error(v ...interface{}) {
	l.Logger.SetPrefix("[ERROR]")
	l.log(v...)
}

func (l *Logger) Warning(v ...interface{}) {
	l.Logger.SetPrefix("[WARNING]")
	l.log(v...)
}
