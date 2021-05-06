package toolkit

import (
	"math"
	"math/rand"
	"time"
)

// IsEqual64 判断两个float64是否相等
func IsEqual64(f1, f2 float64) bool {
	return math.Abs(f1-f2) <= 0.000001
}

// IsEqual32 判断两个float32是否相等
func IsEqual32(f1, f2 float32) bool {
	return IsEqual64(float64(f1), float64(f2))
}

// RandSeed 随机数据种子
func RandSeed() {
	rand.Seed(time.Now().Unix())
}

// RandRangeI32 获取[min, max)之间的int32随机数
func RandRangeI32(min int32, max int32) int32 {
	if max <= min {
		return 0
	}
	return min + rand.Int31n(max-min)
}

// RandRangeI64 获取[min, max)之间的int64随机数
func RandRangeI64(min int64, max int64) int64 {
	if max <= min {
		return 0
	}
	return min + rand.Int63n(max-min)
}

// RandRange 获取[min, max)之间的int随机数
func RandRange(min int, max int) int {
	if max <= min {
		return 0
	}
	return min + rand.Intn(max-min)
}

// RandRangef 获取[min, max)之间的float32随机数
func RandRangef(min float32, max float32) float32 {
	return min + (max-min)*rand.Float32()
}

// DJBHash 获取字符串hash值
func DJBHash(str string) int32 {
	hash := 5381
	for _, c := range str {
		hash += (hash << 5) + int(c)
	}

	return int32(hash & 0x7FFFFFFF)
}

// Remap map value(between srcMin and srcMax) from [srcMin, srcMax] to [dstMin, dstMax]
func Remap(srcMin, srcMax, dstMin, dstMax, value float64) float64 {
	return (value-srcMin)/(srcMax-srcMin)*(dstMax-dstMin) + dstMin
}

// Remap32 map value(between srcMin and srcMax) from [srcMin, srcMax] to [dstMin, dstMax]
func Remap32(srcMin, srcMax, dstMin, dstMax, value float32) float32 {
	return (value-srcMin)/(srcMax-srcMin)*(dstMax-dstMin) + dstMin
}

// Clamp value limit[min, max]
func Clamp(value, min, max float64) float64 {
	return math.Max(min, math.Min(max, value))
}

// Clamp32 value limit[min, max]
func Clamp32(value, min, max float32) float32 {
	return float32(math.Max(float64(min), math.Min(float64(max), float64(value))))
}
