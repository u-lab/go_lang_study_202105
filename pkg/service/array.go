package service

import "fmt"

func Play() {
	// スライス 最大10個まで値を入れることができる配列
	sliceLen := make([]float64, 8)    // 要素数8 容量8
	sliceCap := make([]float64, 0, 8) // 要素数0 容量8

	// スライスをforで回す
	// 要素数 (Speed High, エラー検知しない？)
	for idx, value := range sliceLen {
		sliceLen[idx] = value + 2.0*float64(idx)
	}
	fmt.Println(sliceLen)

	fmt.Println(sliceCap)
	// 容量 (Speed Low, エラー検知する？)
	for i := 0; i < cap(sliceCap); i++ {
		sliceCap = append(sliceCap, 2.0*float64(i))
	}
	fmt.Println(sliceCap)
	fmt.Println(sliceCap[6])
}
