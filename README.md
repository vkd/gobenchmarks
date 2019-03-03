# gobenchmarks
Benchmarks of golang code

## Contents

[Reflect check type](#reflect-check-type)

## Reflect check type

```
$ go test -benchmem -run=NONE -bench=BenchmarkReflectCheckType -benchtime=10s
BenchmarkReflectCheckTypeInterfaceSwitch-4   	2000000000	         9.35 ns/op
BenchmarkReflectCheckTypeInterfaceIf-4       	2000000000	         9.35 ns/op
BenchmarkReflectCheckTypeEquals-4            	2000000000	         9.37 ns/op
```

```
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
```
