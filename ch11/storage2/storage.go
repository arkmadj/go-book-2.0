package storage

var usage = make(map[string]int64)

func bytesInUse(username string) int64 { return usage[username] }
