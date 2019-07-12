package tool

// BigNumber ...
type BigNumber struct {
	numArr []int32
}

// MaxLength 最大位数
var MaxLength = 500

// Raw ...
func (b *BigNumber) Raw(val int32) {
	b.numArr = make([]int32, MaxLength)
	b.reset()
	b.addRaw(val)
}

// FromArr ...
func (b *BigNumber) FromArr(arr []int32) {
	b.numArr = arr
	len := len(b.numArr)
	// log.Debug("FromaArr len %v", len)
	for i := len; i < MaxLength; i++ {
		b.numArr = append(b.numArr, 0)
	}
}

// ToArr ...
func (b *BigNumber) ToArr() []int32 {
	return b.numArr
}

// Add ...
func (b *BigNumber) Add(other *BigNumber) {
	for i := 0; i < MaxLength-1; i++ {
		b.numArr[i] += other.numArr[i]
		b.numArr[i+1] += int32(b.numArr[i] / 10)
		b.numArr[i] = b.numArr[i] % 10
	}
}

// Minus ...
func (b *BigNumber) Minus(other *BigNumber) {
	for i := 0; i < len(b.numArr); i++ {
		if b.numArr[i] >= other.numArr[i] {
			b.numArr[i] -= other.numArr[i]
		} else {
			b.numArr[i+1]--
			b.numArr[i] += 10
			b.numArr[i] -= b.numArr[i]
		}
	}
}

func (b *BigNumber) addRaw(val int32) {
	b.numArr[0] += val
	for i := 0; i < len(b.numArr)-1; i++ {
		b.numArr[i+1] += int32(b.numArr[i] / 10)
		b.numArr[i] = b.numArr[i] % 10
	}
}

func (b *BigNumber) reset() {
	for i := 0; i < len(b.numArr); i++ {
		b.numArr[i] = 0
	}
}
