class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        if len(s) != len(t): # Check if the two strings are of the same length
            return False
            
        return sorted(s) == sorted(t)
        #Iterate through the loop and compare the two sorted strings. If one does not match the other, return false