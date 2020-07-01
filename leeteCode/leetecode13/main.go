func romanToInt(s string) int {
	  //建立map代表各字元所代表的數
    input := map[byte]int {
        'I':1,
        'V':5,
        'X':10,
        'L':50,
        'C':100,
        'D':500,
        'M':1000,
    }    
    answer := 0
		// sign:= 1
		//如果後面的字元所代表的數比前面的大，則前面的數變成負數
		//i+1是因為slice不能越界
    for i := 0; i < len(s); i++ {
        if (i+1)!=len(s) && input[s[i]] < input[s[i+1]]{
            answer += -1*input[s[i]]
        }else{
            answer += input[s[i]]
        }

    }
    return answer
}
