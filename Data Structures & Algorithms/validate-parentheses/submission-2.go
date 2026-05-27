func isValid(s string) bool {
    //we will create a map that defines the matching brackets
	m := map[rune]rune{
		')':'(',
		'}':'{',
		']':'[',
	}

	//we then have a stack that we will append to as we iterate through the string
	stack := []rune{}

	//check check if the string is an outer rune
	for _, ch := range s{
		//for character in range of s
		switch ch {
			//if the character is a opening bracket, append to list
			case '(','{','[' :
				stack = append(stack, ch)
			
			//if character is anything else, pop the stack and check if the map returns true
			default:
				//Pseudo pop function
				if len(stack) != 0 {
					val := stack[len(stack)-1] //get last element
					stack = stack[:len(stack)-1] //shrink by one element

					//get the value from the map for the current character in the string
					v,_ := m[ch]

					//if the value does not match the last item from the stack, ret false
					if v != val {
						return false
					}
				} else {
					return false
				}

		}
	}
	//Edge case for unmatched characters in stack
	if len(stack) != 0 {
		return false
	}
	return true
}
