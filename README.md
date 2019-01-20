# BaseConv

Utility to convert numbers from base 10 integers to base X strings and back again.
Based on Django's [baseconv.py](https://github.com/django/django/blob/master/django/utils/baseconv.py) utility.

## Documentation
Read the [documentation](https://godoc.org/github.com/enricofoltran/baseconv) at godoc.org.

## Usage
```go
func Example() {
	encoded := baseconv.Base36.Encode(1234)
	fmt.Println(encoded)

	decoded, err := baseconv.Base36.Decode(encoded)
	if err != nil {
		panic(err)
	}
	fmt.Print(decoded)

	// Output:
	// ya
	// 1234
}

func Example_base11() {
	base11, err := baseconv.New("0123456789-", "$")
	if err != nil {
		panic(err)
	}

	encoded := base11.Encode(-1234)
	fmt.Println(encoded)

	decoded, err := base11.Decode(encoded)
	if err != nil {
		panic(err)
	}
	fmt.Print(decoded)

	// Output:
	// $-22
	// -1234
}
```
