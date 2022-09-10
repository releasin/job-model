package model

import (
	"encoding/base32"
	"github.com/pborman/uuid"
	"time"
)

var encoding = base32.NewEncoding("ybndrfg8ejkmcpqxot1uwisza345h769").WithPadding(base32.NoPadding)

// NewId is a globally unique identifier.  It is a [A-Z0-9] string 26
// characters long.  It is a UUID version 4 Guid that is zbased32 encoded
// without the padding.
func NewId() string {
	r := uuid.NewRandom()
	id := encoding.EncodeToString(r)
	return id
}

func GetMillis() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func GetMillisAdd(d time.Duration) int64 {
	return time.Now().Add(d).UnixNano() / int64(time.Millisecond)
}

func GetMillisMinus(d time.Duration) int64 {
	return time.Now().Truncate(d).UnixNano() / int64(time.Millisecond)
}
