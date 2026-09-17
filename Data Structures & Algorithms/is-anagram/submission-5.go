func isAnagram(s string, t string) bool {
    if len(s) != len(t){
        return false
    }
    count := [26]int{}
    for i, r:= range s{
        count[r-'a']++
        count[t[i] - 'a']--
}
    for _, val := range count{
        if val!=0{
            return false
        }
    }
    return true
}
