package storage

var usage = make(map[string]int64)

func bytesInUse(username string) int64 { return usage[username] }

const sender = "notifications@example.com"
const password = "correcthorsebatterystaple"
const hostname = "smtp.example.com"
