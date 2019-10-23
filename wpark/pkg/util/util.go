package util

import (
	"encoding/json"
	stdjson "encoding/json"
	"fmt"
	"reflect"
	"time"

	"github.com/rs/xid"
	"github.com/yashmurty/wealth-park/wpark/pkg/logger"
	"go.uber.org/zap"
)

var (
	log = logger.Get("wpark-util")
)

// CreateID creates a new unique id.
func CreateID() string {
	guid := xid.New()
	return guid.String()
}

// MakeTimestamp creates a new unix millisecond timestamp.
func MakeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// ToTimestamp creates a new unix millisecond timestamp from a Time instance.
func ToTimestamp(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

// GetJSON marshals data into JSON and returns it as a string.
func GetJSON(data interface{}) string {
	j, err := json.Marshal(data)
	if err != nil {
		log.Error("could not marshal JSON", zap.Error(err))
	}
	return string(j)
}

// GetPrettyJSON marshals data into pretty JSON and returns it as a string.
func GetPrettyJSON(data interface{}) string {
	b, err := stdjson.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Error("could not marshal JSON", zap.Error(err))
	}
	return string(b)
}

// EnsureNoNilPointers checks the given structs for nil pointers.
func EnsureNoNilPointers(structPtrs ...interface{}) {
	nils := 0
	for _, sp := range structPtrs {
		s := reflect.ValueOf(sp).Elem()
		t := s.Type()
		for i := 0; i < s.NumField(); i++ {
			f := s.Field(i)
			switch f.Kind().String() {
			case "ptr", "interface":
				if f.IsNil() {
					fmt.Printf("%s.%s %s is nil\n", s.Type(), t.Field(i).Name, f.Type())
					nils++
				}
			}
		}
	}
	if nils > 0 {
		panic(fmt.Sprintf("found %d nil pointers!", nils))
	}
}
