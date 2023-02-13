package zlog

func Info(args ...interface{}) {
	defer sugaredLogger.Sync()
	sugaredLogger.Info(args)
}

func Infof(format string, args ...interface{}) {
	defer sugaredLogger.Sync()
	sugaredLogger.Infof(format, args)
}

func Error(args ...interface{}) {
	defer sugaredLogger.Sync()
	sugaredLogger.Error(args)
}

func Errorf(format string, args ...interface{}) {
	defer sugaredLogger.Sync()
	sugaredLogger.Errorf(format, args)
}

func Panic(args ...interface{}) {
	defer sugaredLogger.Sync()
	sugaredLogger.Panic(args)
}

func Panicf(format string, args ...interface{}) {
	defer sugaredLogger.Sync()
	sugaredLogger.Panicf(format, args)
}

func Fatal(args ...interface{}) {
	defer sugaredLogger.Sync()
	sugaredLogger.Fatal(args)
}

func Fatalf(format string, args ...interface{}) {
	defer sugaredLogger.Sync()
	sugaredLogger.Fatalf(format, args)
}
