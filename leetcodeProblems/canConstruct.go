package leetcodeProblems

func CanConstruct(ransomNote string, magazine string) bool {
	mapRansom := map[string]int{}
	mapMagazine := map[string]int{}
    answer := true
    for i := 0; i < len(ransomNote); i++{
        mapRansom[string(ransomNote[i])]++
    }
    
    for i := 0; i < len(magazine); i++{
        mapMagazine[string(magazine[i])]++
    }
    
    for i :=0 ;i  < len(ransomNote); i++{
        for j := 0; j < len(magazine);j++{
            if ransomNote[i] == magazine[i] && mapRansom[string(ransomNote[i])]  <= mapMagazine[string(magazine[i])]{
                answer = true
            } else{
                answer = false
            }
        }
    }


    return answer
}
