package log

func MakeLogWithCaller(skip int) Logger {
	return New(&Config{Caller: true, CallerSkip: skip, Stack: false})
}

func MakeLog() Logger {
	return New(&Config{Caller: false, Stack: false})
}

func MakeLogWithStack() Logger {
	return New(&Config{Caller: true, Stack: true})
}
