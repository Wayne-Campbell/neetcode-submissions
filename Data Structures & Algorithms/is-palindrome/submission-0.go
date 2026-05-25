func isPalindrome(s string) bool {
//We convert our string to a rune to make it safe for comparison, GO uses bytes when 
//accesing a string
runes := []rune(s)

//Place a pointer at both the start and end of the string
start := 0
end := len(s) - 1
fmt.Println(start)
fmt.Println(end)

//iterate through the string
for start < end {
	//first check if the current character is alphanumeric
	if !unicode.IsLetter(runes[start]) && !unicode.IsDigit(runes[start]) {
		//skip the first character
		start++
		continue
	}

	if !unicode.IsLetter(runes[end]) && !unicode.IsDigit(runes[end]) {
		//skip the last character
		end--
		continue
	}

	fmt.Println("Comparing %s and %s", unicode.ToLower(runes[start]), unicode.ToLower(runes[end]))

	//now that we know that the runes are alphanumeric, we can compare our current indices
	if unicode.ToLower(runes[start]) != unicode.ToLower(runes[end]){
		return false
	}
	start++
	end--
}
return true
}
