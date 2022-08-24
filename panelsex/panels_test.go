package panelsex

import "testing"

func Test1(t *testing.T) {

}

func Sum[T int](t ...T) T {
	var sun T
	for _, v := range t {
		sun += v
	}
	return sun
}
