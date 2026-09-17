func isAnagram(s string, t string) bool {
    if len(s) != len(t){
        return false
    }
    seenS:= make(map[rune]int)
    for _, r:= range s{
        seenS[r] += 1
    }
    seenT:= make(map[rune]int)
    for _, r:= range t{
        seenT[r] += 1
    }
    for r,val:= range seenS{
        if seenT[r] != val{
            return false
        }
    }
    return true
}
