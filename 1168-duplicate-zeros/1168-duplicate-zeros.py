class Solution(object):
    def duplicateZeros(self, arr):
        i, dlinna = 0, len(arr)
        while i<dlinna:
			    if arr[i]==0:
			 	    arr.pop()
				    arr.insert(i,0)
				    i+=1
			    i+=1