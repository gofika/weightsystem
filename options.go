package weightsystem

type NewOptions struct {
	minWeight float64
	maxWeight float64
}

type NewOption func(*NewOptions)

func WithMinWeight(minWeight float64) NewOption {
	return func(options *NewOptions) {
		options.minWeight = minWeight
	}
}

func WithMaxWeight(maxWeight float64) NewOption {
	return func(options *NewOptions) {
		options.maxWeight = maxWeight
	}
}

type AddItemOptions struct {
	weight float64
}

type AddItemOption func(*AddItemOptions)

func WithWeight(weight float64) AddItemOption {
	return func(options *AddItemOptions) {
		options.weight = weight
	}
}
