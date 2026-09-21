

type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var encoded strings.Builder
	for _, str := range strs {
		length := len(str)
		encoded.WriteString(fmt.Sprintf("%d#%s", length, str))
	}
	return encoded.String()
}

func (s *Solution) Decode(encoded string) []string {
	var decoded []string
	encodedBytes := []byte(encoded)
	intLength := 0
	for i := 0; i < len(encodedBytes); {
		if encodedBytes[i] != '#' {
			intLength = intLength*10 + int(encodedBytes[i]-'0')
			i++
		} else {
			decoded = append(decoded, string(encodedBytes[i+1:i+1+intLength]))
			i += 1 + intLength
			intLength = 0
		}
	}
	return decoded
}
