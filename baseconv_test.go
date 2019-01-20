// Copyright 2019 Enrico Foltran. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package baseconv_test

import (
	"fmt"
	"testing"

	"github.com/enricofoltran/baseconv"
)

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

func TestBaseConv(t *testing.T) {
	nums := []int64{-10000000000, 10000000000}
	for i := -100; i <= 100; i++ {
		nums = append(nums, int64(i))
	}

	testcases := []struct {
		name      string
		converter *baseconv.BaseConverter
		nums      []int64
	}{
		{"Base2", baseconv.Base2, nums},
		{"Base16", baseconv.Base16, nums},
		{"Base36", baseconv.Base36, nums},
		{"Base56", baseconv.Base56, nums},
		{"Base62", baseconv.Base62, nums},
		{"Base64", baseconv.Base64, nums},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			for _, num := range tt.nums {
				decoded, err := tt.converter.Decode(tt.converter.Encode(num))
				if err != nil {
					t.Errorf("Decode(Encode(%d)) err = %v", num, err)
				}
				if num != decoded {
					t.Errorf("Decode(Encode(%d)); want=%d, got=%d", num, num, decoded)
				}
			}
		})
	}
}

func TestBase11(t *testing.T) {
	base11, err := baseconv.New("0123456789-", "$")
	if err != nil {
		t.Fatalf("New(0123456789-, $) err = %v", err)
	}

	t.Run("Encode", func(t *testing.T) {
		want := "-22"
		if got := base11.Encode(1234); got != want {
			t.Errorf("Encode(1234) = %s; want = %s", got, want)
		}

		want = "$-22"
		if got := base11.Encode(-1234); got != want {
			t.Errorf("Encode(1234) = %s; want = %s", got, want)
		}
	})

	t.Run("Decode", func(t *testing.T) {
		want := int64(1234)
		if got, _ := base11.Decode("-22"); got != want {
			t.Errorf("Decode(-22) = %d; want = %d", got, want)
		}

		want = int64(-1234)
		if got, _ := base11.Decode("$-22"); got != want {
			t.Errorf("Decode($-22) = %d; want = %d", got, want)
		}
	})
}

func TestBase20(t *testing.T) {
	base20, err := baseconv.New("0123456789abcdefghij", "")
	if err != nil {
		t.Fatalf("New(0123456789abcdefghij, '') err = %v", err)
	}

	t.Run("Encode", func(t *testing.T) {
		want := "31e"
		if got := base20.Encode(1234); got != want {
			t.Errorf("Encode(1234) = %s; want = %s", got, want)
		}

		want = "-31e"
		if got := base20.Encode(-1234); got != want {
			t.Errorf("Encode(1234) = %s; want = %s", got, want)
		}
	})

	t.Run("Decode", func(t *testing.T) {
		want := int64(1234)
		if got, _ := base20.Decode("31e"); got != want {
			t.Errorf("Decode(31e) = %d; want = %d", got, want)
		}

		want = int64(-1234)
		if got, _ := base20.Decode("-31e"); got != want {
			t.Errorf("Decode(-31e) = %d; want = %d", got, want)
		}
	})
}

func TestBase64(t *testing.T) {
	t.Run("Encode", func(t *testing.T) {
		want := "JI"
		if got := baseconv.Base64.Encode(1234); got != want {
			t.Errorf("Encode(1234) = %s; want = %s", got, want)
		}

		want = "$JI"
		if got := baseconv.Base64.Encode(-1234); got != want {
			t.Errorf("Encode(1234) = %s; want = %s", got, want)
		}
	})

	t.Run("Decode", func(t *testing.T) {
		want := int64(1234)
		if got, _ := baseconv.Base64.Decode("JI"); got != want {
			t.Errorf("Decode(JI) = %d; want = %d", got, want)
		}

		want = int64(-1234)
		if got, _ := baseconv.Base64.Decode("$JI"); got != want {
			t.Errorf("Decode($JI) = %d; want = %d", got, want)
		}
	})
}

func TestBase7(t *testing.T) {
	base7, err := baseconv.New("cjdhel3", "g")
	if err != nil {
		t.Fatalf("New(cjdhel3, g) err = %v", err)
	}

	t.Run("Encode", func(t *testing.T) {
		want := "hejd"
		if got := base7.Encode(1234); got != want {
			t.Errorf("Encode(1234) = %s; want = %s", got, want)
		}

		want = "ghejd"
		if got := base7.Encode(-1234); got != want {
			t.Errorf("Encode(1234) = %s; want = %s", got, want)
		}
	})

	t.Run("Decode", func(t *testing.T) {
		want := int64(1234)
		if got, _ := base7.Decode("hejd"); got != want {
			t.Errorf("Decode(hejd) = %d; want = %d", got, want)
		}

		want = int64(-1234)
		if got, _ := base7.Decode("ghejd"); got != want {
			t.Errorf("Decode(ghejd) = %d; want = %d", got, want)
		}
	})
}

func TestErrAlphabetEmptyString(t *testing.T) {
	_, err := baseconv.New("", "")
	if err != baseconv.ErrAlphabetEmptyString {
		t.Fatalf("New('', '') err = %v; want = ErrAlphabetEmptyString", err)
	}
}

func TestErrSignCharInBase(t *testing.T) {
	_, err := baseconv.New("abc", "a")
	if err != baseconv.ErrSignCharInBase {
		t.Fatalf("New('abc', 'a') err = %v; want = ErrSignCharInBase", err)
	}

	_, err = baseconv.New("abc", "d")
	if err != nil {
		t.Fatalf("New('abc', 'd') err = %v; want = nil", err)
	}
}

func TestString(t *testing.T) {
	base7, err := baseconv.New("cjdhel3", "g")
	if err != nil {
		t.Fatalf("New(cjdhel3, g) err = %v", err)
	}

	want := "<BaseConverter: base7 (cjdhel3)>"
	if got := fmt.Sprint(base7); got != want {
		t.Fatalf("fmt.Sprint(base7) = %s; want = %s", got, want)
	}
}
