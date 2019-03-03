package gobenchmarks

import (
	"reflect"
	"testing"
	"time"
)

func checkInterfaceTypeSwitch(value reflect.Value) bool {
	switch value.Interface().(type) {
	case time.Duration:
		return true
	}
	return false
}

func checkInterfaceTypeIf(value reflect.Value) bool {
	_, ok := value.Interface().(time.Duration)
	return ok
}

var typeOfTimeDuration = reflect.TypeOf(time.Duration(0))

func checkInterfaceTypeEquals(value reflect.Value) bool {
	return value.Type() == typeOfTimeDuration
}

func benchmarkReflectCheckType(b *testing.B, fn func(reflect.Value) bool) {
	var ok bool
	value := reflect.ValueOf(time.Duration(0))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ok = checkInterfaceTypeSwitch(value)
	}
	b.StopTimer()

	if ok == false {
		b.Errorf("Wrong result from checker: false (expect: true)")
	}
}

func BenchmarkReflectCheckTypeInterfaceSwitch(b *testing.B) {
	benchmarkReflectCheckType(b, checkInterfaceTypeSwitch)
}
func BenchmarkReflectCheckTypeInterfaceIf(b *testing.B) {
	benchmarkReflectCheckType(b, checkInterfaceTypeIf)
}
func BenchmarkReflectCheckTypeEquals(b *testing.B) {
	benchmarkReflectCheckType(b, checkInterfaceTypeEquals)
}
