package computer

type Spec struct { // exported struct
	Maker string //export field
	Price int    // exported field
	model string // unexported field
}
