package glog

type GRPCLogger struct{}

func (g *GRPCLogger) Fatal(args ...interface{}) {
	Error(args...)
}

func (g *GRPCLogger) Fatalf(format string, args ...interface{}) {
	Errorf(format, args...)
}

func (g *GRPCLogger) Fatalln(args ...interface{}) {
	Errorln(args...)
}

func (g *GRPCLogger) Print(args ...interface{}) {
	Info(args...)
}

func (g *GRPCLogger) Printf(format string, args ...interface{}) {
	Infof(format, args...)
}

func (g *GRPCLogger) Println(args ...interface{}) {
	Infoln(args...)
}
