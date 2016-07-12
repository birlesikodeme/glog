package glog

type Logger struct{}

func (g *Logger) Fatal(args ...interface{}) {
	Error(args...)
}

func (g *Logger) Fatalf(format string, args ...interface{}) {
	Errorf(format, args...)
}

func (g *Logger) Fatalln(args ...interface{}) {
	Errorln(args...)
}

func (g *Logger) Print(args ...interface{}) {
	Info(args...)
}

func (g *Logger) Printf(format string, args ...interface{}) {
	Infof(format, args...)
}

func (g *Logger) Println(args ...interface{}) {
	Infoln(args...)
}
