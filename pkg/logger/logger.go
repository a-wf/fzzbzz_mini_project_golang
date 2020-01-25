// package logging

// import (
// 	"io"
// 	"os"
// 	log "github.com/sirupsen/logrus"
// )


// func Init(logLevel string, logFile string) *log.Logger{
	
// 	log.SetFormatter(&log.JSONFormatter{});
	
// 	bothOutputs := io.MultiWriter(os.Stdout, logFile);
	
// 	log.SetOutput(bothOutputs);

// 	log.SetLevel(logLevel);
// 	return log
// }


package logger

import (
	"fmt"
	"os"
	"time"
)

// CheckIfError should be used to naively panics if an error is not nil.
func CheckIfError(err error, funcName string) {
	if err == nil {
		return
	}
	dt := time.Now()
	fmt.Printf(dt.Format("01-02-2006 15:04:05 Mon")+" \x1b[31;1mERROR %s\x1b[0m\n", fmt.Sprintf("error: %s: %s",funcName,  err))
	os.Exit(1)
}

// Error should be used to describe the example commands that are about to run.
func Error(format string, args ...interface{}) {
	dt := time.Now()
	fmt.Printf(dt.Format("01-02-2006 15:04:05 Mon")+" \x1b[31;1mERROR %s\x1b[0m\n", fmt.Sprintf(format, args...))
}

// Info should be used to describe the example commands that are about to run.
func Info(format string, args ...interface{}) {
	dt := time.Now()
	fmt.Printf(dt.Format("01-02-2006 15:04:05 Mon")+" \x1b[34;1mINFO %s\x1b[0m\n", fmt.Sprintf(format, args...))
}

// Warning should be used to display a warning
func Warning(format string, args ...interface{}) {
	dt := time.Now()
	fmt.Printf(dt.Format("01-02-2006 15:04:05 Mon")+" \x1b[36;WARN %s\x1b[0m\n", fmt.Sprintf(format, args...))
}






// package logger

// import (
// 	"fmt"
// 	"os"
// 	"log"
// 	"io"
// ) 

// func Init(appName string, filePath string) *log.Logger) {
// 	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
//         if err != nil {
//                 log.Fatal(err)
//         }
// 		defer f.Close()
// 		mw := io.MultiWriter(os.Stdout, f)
// 		return log.New(f, appName , log.Lshortfile)
// }

// // CheckIfError should be used to naively panics if an error is not nil.
// func  CheckIfError(err error, funcName string) {
// 	if err == nil {
// 		return
// 	}
// 	log.SetOutput(f)

// 	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s: %s",funcName,  err))
// 	log.Println(fmt.Sprintf("error: %s: %s",funcName,  err))

// 	// os.Exit(1)
// }

// // Error should be used to describe the example commands that are about to run.
// func Error(format string, args ...interface{}) {
// 	log.SetOutput(f)
// 	fmt.Printf("\x1b[31;1mERROR %s\x1b[0m\n", fmt.Sprintf(format, args...))
// 	log.Println(fmt.Sprintf(format, args...))

// }

// // Info should be used to describe the example commands that are about to run.
// func Info(format string, args ...interface{}) {
// 	log.SetOutput(f)
// 	fmt.Printf("\x1b[34;1mINFO %s\x1b[0m\n", fmt.Sprintf(format, args...))
// 	log.Println(fmt.Sprintf(format, args...))
// }

// // Warning should be used to display a warning
// func Warning(format string, args ...interface{}) {
// 	log.SetOutput(f)
// 	fmt.Printf("\x1b[36;1mWARNING %s\x1b[0m\n", fmt.Sprintf(format, args...))
// 	log.Println(fmt.Sprintf(format, args...))
// }





// package logging

// import (
// 	"io"
// 	"os"
// 	log "github.com/sirupsen/logrus"
// )


// func Init(logLevel string, logFile string){
	
// 	log.SetFormatter(&log.JSONFormatter{});
	
// 	bothOutputs := io.MultiWriter(os.Stdout, logFile);
	
// 	log.SetOutput(bothOutputs);

// 	log.SetLevel(logLevel);
// }

