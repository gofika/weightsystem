package weightsystem

import (
	"math/rand/v2"
	"sort"
)

// WeightedItem represents an item with its associated weight
type WeightedItem[T comparable] struct {
	Item   T
	Weight float64
}

// WeightSystem is a generic weight system for any comparable type T
type WeightSystem[T comparable] struct {
	weights     map[T]float64
	minWeight   float64
	maxWeight   float64
	totalWeight float64
	avgWeight   float64
}

// New creates a new WeightSystem with the given items and weights
// Options can be passed to set the initial weight, min weight, and max weight
// If no options are passed, min weight will be 1, and max weight will be 1000
func New[T comparable](opts ...NewOption) *WeightSystem[T] {
	options := &NewOptions{
		minWeight: 1,
		maxWeight: 1000,
	}
	for _, opt := range opts {
		opt(options)
	}
	options.minWeight = max(1, options.minWeight)
	options.maxWeight = max(options.minWeight, options.maxWeight)
	ws := &WeightSystem[T]{
		weights:   make(map[T]float64),
		minWeight: options.minWeight,
		maxWeight: options.maxWeight,
	}
	ws.updateAvgWeight()
	return ws
}

// MinWeight returns the minimum weight allowed in the system
func (ws *WeightSystem[T]) MinWeight() float64 {
	return ws.minWeight
}

// MaxWeight returns the maximum weight allowed in the system
func (ws *WeightSystem[T]) MaxWeight() float64 {
	return ws.maxWeight
}

// updateAvgWeight recalculates the average weight of the system
func (ws *WeightSystem[T]) updateAvgWeight() {
	if len(ws.weights) > 0 {
		ws.avgWeight = ws.totalWeight / float64(len(ws.weights))
	} else {
		ws.avgWeight = (ws.minWeight + ws.maxWeight) / 2
	}
}

// AdjustWeight adjusts the weight of an item based on success or failure
func (ws *WeightSystem[T]) AdjustWeight(item T, success bool) {
	oldWeight := ws.weights[item]
	if success {
		ws.weights[item] = min(oldWeight*1.1, ws.maxWeight)
	} else {
		ws.weights[item] = max(oldWeight*0.9, ws.minWeight)
	}
	ws.totalWeight += ws.weights[item] - oldWeight
	ws.updateAvgWeight()
}

// GetItem returns a randomly selected item based on weights
func (ws *WeightSystem[T]) GetItem() T {
	r := rand.Float64() * ws.totalWeight
	for item, weight := range ws.weights {
		r -= weight
		if r <= 0 {
			return item
		}
	}
	// This should never happen, but return random item if it does
	var item T
	for item = range ws.weights {
		break
	}
	return item
}

// Weights returns a clone of the current weights map
func (ws *WeightSystem[T]) Weights() map[T]float64 {
	clone := make(map[T]float64, len(ws.weights))
	for item, weight := range ws.weights {
		clone[item] = weight
	}
	return clone
}

// AddItem adds a new item to the system with the given weight
func (ws *WeightSystem[T]) AddItem(item T, opts ...AddItemOption) {
	if _, exists := ws.weights[item]; exists {
		return
	}
	options := &AddItemOptions{
		weight: ws.avgWeight,
	}
	for _, opt := range opts {
		opt(options)
	}

	ws.weights[item] = max(min(options.weight, ws.maxWeight), ws.minWeight)
	ws.totalWeight += ws.weights[item]
	ws.updateAvgWeight()
}

// AddItems adds multiple items to the system with the same weight
func (ws *WeightSystem[T]) AddItems(items []T, opts ...AddItemOption) {
	for _, item := range items {
		ws.AddItem(item, opts...)
	}
}

// RemoveItem removes an item from the system
func (ws *WeightSystem[T]) RemoveItem(item T) {
	weight, exists := ws.weights[item]
	if !exists {
		return
	}
	delete(ws.weights, item)
	ws.totalWeight -= weight
	ws.updateAvgWeight()
}

// SortedWeights returns a sorted slice of WeightedItems, from highest weight to lowest
func (ws *WeightSystem[T]) SortedWeights() []WeightedItem[T] {
	items := make([]WeightedItem[T], 0, len(ws.weights))
	for item, weight := range ws.weights {
		items = append(items, WeightedItem[T]{Item: item, Weight: weight})
	}
	sort.Slice(items, func(i, j int) bool {
		return items[i].Weight > items[j].Weight
	})
	return items
}
