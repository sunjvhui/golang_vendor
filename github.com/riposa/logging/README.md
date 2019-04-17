# logging
A go log module which imitated python "logging" module

    type Logger struct {

        // contains filtered or unexported fields

    }
the logger struct with a series of method for logging, such as: Info(),
Warning(), Error(), Exception()

    func (l *Logger) Error(v ...interface{})

    func (l *Logger) Exception(v error)

    func (l *Logger) Info(v ...interface{})

    func (l *Logger) Warning(v ...interface{})


Logger.Init() will initialize the logger with default log format and default output

<os.Stdout> you can set the handler list by using <SetHandler>, a

rotating handler was already provided above

    func (l *Logger) Init()

    func (l *Logger) SetHandler(handler ...io.Writer)



NewRotateHandler() will return a new rotating file handler path: <string>, log dir path

filename: <string>, base name of log file ext: <string>, extension name

of log file maxFile: <int>, the max count of rotating files

    func NewRotateHandler(path string, filename string, ext string, maxFile int) *TimeRotateHandler


Type TimeRotateHandler implemented the common interface <io.Writer> to printing logs

    type TimeRotateHandler struct {

        // contains filtered or unexported fields

    }

    func (t *TimeRotateHandler) Write(p []byte) (n int, err error)
