package gophersdk

import (
	"github.com/carolynvs/gopherdome-rest"
)

func GetGopher(name string) gopherest.Gopher {
	return gopherest.GetResource("gophers", name)
}
