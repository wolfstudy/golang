package main

//求数据的最大位
func maxBit (values []int) (ret int)  {
	ret = 1 //保存最大位数
	var key int =10
	iLen := len(values)

	for i := 0; i < iLen ; i++ {
		for values[i] >= key {
			key = key * 10
			ret++
		}
	}
	return
	
}

func radixSort(values []int) {
	key := maxBit(values)
	iLen := len(values)
	tmp := make([]int, iLen, iLen)
	count := new([10]int)
	radix := 1
	var i, j, k int
	for i = 0; i < key; i++ {
		//进行key次排序
		for j = 0; j < 10; j++ {
			count[j] = 0
		}
		for j = 0; j < iLen; j++ {
			k = (values[j] / radix) % 10
			count[k]++
		}
		for j = 1; j < 10; j++ {
			//将tmp中的为准依次分配给每个桶
			count[j] = count[j-1] + count[j]
		}
		for j = iLen - 1; j >= 0; j-- {
			k = (values[j] / radix) % 10
			tmp[count[k]-1] = values[j]
			count[k]--
		}
		for j = 0; j < iLen; j++ {
			values[j] = tmp[j]
		}
			radix = radix * 10
	}
}

