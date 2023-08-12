package datatypes

type ExtendedCircularBuffer struct {
	CircularBuffer
}

func (cb *ExtendedCircularBuffer) addValues(vals ...float64) {
    for _, val := range vals {
        cb.AddValue(val)
    }
}

func NewExtendedCircularBuffer (size int) ExtendedCircularBuffer {
    return ExtendedCircularBuffer {
        CircularBuffer:  NewCircularBuffer(size),
    }
}
