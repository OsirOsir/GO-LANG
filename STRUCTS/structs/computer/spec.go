package computer

type Spec struct { // exported struct
	Maker string // exported struct
	Price int    // exported struct
	model string // unexported struct
}
