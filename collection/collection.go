package collection

import (
	"fmt"
	"strings"
)

type Collection[T comparable] struct {
	arr []T
}

func Collect[T comparable](t []T) *Collection[T] {
	return &Collection[T]{
		arr: t,
	}
}

func (c *Collection[T]) Map(f func(T, int64) T) *Collection[T] {
	var tmp []T
	for k, v := range c.arr {
		tmp = append(tmp, f(v, int64(k)))
	}
	c.arr = tmp
	return c
}

func (c *Collection[T]) Filter(f func(T, int64) bool) *Collection[T] {
	var tmp []T
	for k, v := range c.arr {
		if f(v, int64(k)) {
			tmp = append(tmp, v)
		}
	}
	c.arr = tmp
	return c
}

func (c *Collection[T]) Shift() *Collection[T] {
	c.arr = c.arr[1:]
	return c
}

func (c *Collection[T]) Pop() *Collection[T] {
	c.arr = c.arr[:len(c.arr)-1]
	return c
}

func (c *Collection[T]) Reverse() *Collection[T] {
	var result []T
	for i := len(c.arr) - 1; i >= 0; i-- {
		result = append(result, c.arr[i])
	}
	c.arr = result
	return c
}

func (c *Collection[T]) Join(sep string) string {
	sb := strings.Builder{}
	for k, v := range c.arr {
		sb.WriteString(fmt.Sprintf("%v", v))
		if len(c.arr) > k+1 {
			sb.WriteString(sep)
		}
	}
	return sb.String()
}

func (c *Collection[T]) Remove(index int64) *Collection[T] {
	c.arr = append(c.arr[:index], c.arr[index+1:]...)
	return c
}

func (c *Collection[T]) ToList() []T {
	return c.arr
}

func (c *Collection[T]) First() T {
	return c.arr[0]
}

func (c *Collection[T]) Last() T {
	return c.arr[len(c.arr)-1]
}

func (c *Collection[T]) Get(index int64) T {
	return c.arr[index]
}

func (c *Collection[T]) Contains(value T) bool {
	for _, v := range c.arr {
		if value == v {
			return true
		}
	}
	return false
}

func Map[T any, V any](arr []T, f func(T, int64) V) []V {
	var tmp []V
	for k, v := range arr {
		tmp = append(tmp, f(v, int64(k)))
	}
	return tmp
}

func Filter[T interface{}](t []T, f func(T, int) bool) []T {
	var o []T
	for k, v := range t {
		if f(v, k) {
			o = append(o, v)
		}
	}
	return o
}

func Concat[T interface{}](t1 []T, t2 []T) []T {
	return append(t1[:], t2[:]...)
}

func Shift[T interface{}](arr []T) []T {
	if len(arr) > 0 {
		return arr[1:]
	}
	return arr
}

func Remove[T interface{}](arr []T, index int) []T {
	if index < len(arr) {
		arr = append(arr[:index], arr[index+1:]...)
	}
	return arr
}

func Pop[T interface{}](arr []T) []T {
	if len(arr) > 0 {
		return arr[:len(arr)-1]
	}
	return arr
}

func Set[T any](arr []T) []T {
	var sets []T
	for _, v := range arr {
		exist := false
		for _, v2 := range sets {
			if fmt.Sprintf("%v", v) == fmt.Sprintf("%v", v2) {
				exist = true
				break
			}
		}
		if !exist {
			sets = append(sets, v)
		}
	}
	return sets
}

func Join[T interface{}](tab []T, sep string) string {
	sb := strings.Builder{}
	for k, v := range tab {
		sb.WriteString(fmt.Sprintf("%v", v))
		if len(tab) > k+1 {
			sb.WriteString(sep)
		}
	}
	return sb.String()
}

func Reverse[T interface{}](arr []T) []T {
	var result []T
	for i := len(arr) - 1; i >= 0; i-- {
		result = append(result, arr[i])
	}
	return result
}

func Of[T comparable](t ...T) *Collection[T] {
	return &Collection[T]{
		arr: t,
	}
}

func ListOf[T comparable](items ...T) []T {
	return items
}
