package nengine

import (
	"slices"

	"github.com/SnareChops/nengine/types"
)

type Entities struct {
	Items []types.Entity
}

func (self *Entities) Add(item types.Entity) {
	if !slices.Contains(self.Items, item) {
		self.Items = append(self.Items, item)
	}
}

func (self *Entities) Remove(item types.Entity) {
	for i, e := range self.Items {
		if e == item {
			self.Items = append(self.Items[:i], self.Items[i+1:]...)
			return
		}
	}
}

func (self *Entities) Update(delta int) {
	for _, entity := range self.Items {
		entity.Update(delta)
	}
}
