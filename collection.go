package nengine

import (
	"slices"
)

type Collection[T comparable] struct {
	Items []T
}

func (self *Collection[T]) Add(item T) {
	if !slices.Contains(self.Items, item) {
		self.Items = append(self.Items, item)
	}
}

func (self *Collection[T]) Remove(item T) {
	for i, e := range self.Items {
		if e == item {
			self.Items = append(self.Items[:i], self.Items[i+1:]...)
			return
		}
	}
}
