package model

type NumberGeneratorResponse struct {
	Key   string
	Value string
}

type NumberGenerator struct {
	Prefix   string
	NxtValue int
	Length   int
}

type NumberGeneratorRequest struct {
	Prefix     string
	InitialVal int
	Length     int
}
