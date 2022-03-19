package main

func maximumSubsequenceCount(text string, pattern string) int64 {
	runes := []rune(pattern)
	if runes[0] == runes[1] {
		var c int64 = 1
		for _, char := range text {
			if char == runes[0] {
				c++
			}
		}
		return c * (c - 1) / 2
	} else {
		var a_count int64 = 0
		var b_count int64 = 0
		var c_count int64 = 0
		var sum int64 = 0
		for _, char := range text {
			if char == runes[0] {
				if b_count > 0 {
					sum = sum + b_count*a_count
					c_count += b_count
					b_count = 0
				}
				a_count++
			} else if char == runes[1] {
				b_count++
			}
		}
		l := c_count + sum + (1+a_count)*b_count
		r := sum + a_count*(b_count+1)
		if l > r {
			return l
		} else {
			return r
		}
	}
}

func main() {
	println(maximumSubsequenceCount("vnedkpkkyxelxqptfwuzcjhqmwagvrglkeivowvbjdoyydnjrqrqejoyptzoklaxcjxbrrfmpdxckfjzahparhpanwqfjrpbslsyiwbldnpjqishlsuagevjmiyktgofvnyncizswldwnngnkifmaxbmospdeslxirofgqouaapfgltgqxdhurxljcepdpndqqgfwkfiqrwuwxfamciyweehktaegynfumwnhrgrhcluenpnoieqdivznrjljcotysnlylyswvdlkgsvrotavnkifwmnvgagjykxgwaimavqsxuitknmbxppgzfwtjdvegapcplreokicxcsbdrsyfpustpxxssnouifkypwqrywprjlyddrggkcglbgcrbihgpxxosmejchmzkydhquevpschkpyulqxgduqkqgwnsowxrmgqbmltrltzqmmpjilpfxocflpkwithsjlljxdygfvstvwqsyxlkknmgpppupgjvfgmxnwmvrfuwcrsadomyddazlonjyjdeswwznkaeaasyvurpgyvjsiltiykwquesfjmuswjlrphsdthmuqkrhynmqnfqdlwnwesdmiiqvcpingbcgcsvqmsmskesrajqwmgtdoktreqssutpudfykriqhblntfabspbeddpdkownehqszbmddizdgtqmobirwbopmoqzwydnpqnvkwadajbecmajilzkfwjnpfyamudpppuxhlcngkign",
		"rr"))
}
