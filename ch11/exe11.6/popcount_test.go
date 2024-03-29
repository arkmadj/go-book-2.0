package popcount

import "testing"

func PopCountShiftMask(x uint64) int {
	count := 0
	mask := uint64(1)
	for i := 0; i < 64; i++ {
		if x&mask > 0 {
			count++
		}
		mask <<= 1
	}
	return count
}

func PopCountShiftValue(x uint64) int {
	count := 0
	mask := uint64(1)
	for i := 0; i < 64; i++ {
		if x&mask > 0 {
			count++
		}
		x >>= 1
	}
	return count
}

func PopCountClearRightmost(x uint64) int {
	count := 0
	for x != 0 {
		x &= x - 1
		count++
	}
	return count
}

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCountTable(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func bench(b *testing.B, f func(uint64) int) {
	for i := 0; i < b.N; i++ {
		f(uint64(i))
	}
}

func benchN(b *testing.B, n int, f func(uint64) int) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			f(uint64(j))
		}
	}
}

func benchTableN(b *testing.B, n int) {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			PopCountTable(uint64(j))
		}
	}
}

func BenchmarkTable(b *testing.B) {
	bench(b, PopCountTable)
}

func BenchmarkShiftMask(b *testing.B) {
	bench(b, PopCountShiftMask)
}

func BenchmarkShiftValue(b *testing.B) {
	bench(b, PopCountShiftValue)
}

func BenchmarkClearRightMost(b *testing.B) {
	bench(b, PopCountClearRightmost)
}

func BenchmarkClearRightMost1(b *testing.B) {
	benchN(b, 1, PopCountClearRightmost)
}

func BenchmarkClearRightMost10(b *testing.B) {
	benchN(b, 10, PopCountClearRightmost)
}

func BenchmarkClearRightMost100(b *testing.B) {
	benchN(b, 100, PopCountClearRightmost)
}

func BenchmarkClearRightMost1000(b *testing.B) {
	benchN(b, 1000, PopCountClearRightmost)
}

func BenchmarkClearRightMost10000(b *testing.B) {
	benchN(b, 10000, PopCountClearRightmost)
}

func BenchmarkShiftValue1(b *testing.B) {
	benchN(b, 1, PopCountShiftValue)
}

func BenchmarkShiftValue10(b *testing.B) {
	benchN(b, 10, PopCountShiftValue)
}

func BenchmarkShiftValue100(b *testing.B) {
	benchN(b, 100, PopCountShiftValue)
}

func BenchmarkShiftValue1000(b *testing.B) {
	benchN(b, 1000, PopCountShiftValue)
}

func BenchmarkShiftValue10000(b *testing.B) {
	benchN(b, 10000, PopCountShiftValue)
}

func BenchmarkTable1(b *testing.B) {
	benchN(b, 1, PopCountTable)
}

func BenchmarkTable10(b *testing.B) {
	benchN(b, 10, PopCountTable)
}

func BenchmarkTable100(b *testing.B) {
	benchN(b, 100, PopCountTable)
}

func BenchmarkTable1000(b *testing.B) {
	benchN(b, 1000, PopCountTable)
}

func BenchmarkTable10000(b *testing.B) {
	benchN(b, 10000, PopCountTable)
}

func BenchmarkShiftMask1(b *testing.B) {
	benchN(b, 1, PopCountShiftMask)
}

func BenchmarkShiftMask10(b *testing.B) {
	benchN(b, 10, PopCountShiftMask)
}

func BenchmarkShiftMask100(b *testing.B) {
	benchN(b, 100, PopCountShiftMask)
}

func BenchmarkShiftMask1000(b *testing.B) {
	benchN(b, 1000, PopCountShiftMask)
}

func BenchmarkShiftMask10000(b *testing.B) {
	benchN(b, 10000, PopCountShiftMask)
}
