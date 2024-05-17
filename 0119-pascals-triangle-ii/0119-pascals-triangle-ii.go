func getRow(rowIndex int) []int {
        	var rows = make([][]int, rowIndex+1)
	for i:=0;i<rowIndex+1;i++{
		rows[i]=make([]int,i+1)
		rows[i][0]=1
		rows[i][i]=1
	}
	for i:=2;i<rowIndex+1;i++{
		for j:=1;j<i;j++ {
			rows[i][j]=rows[i-1][j-1]+rows[i-1][j]
		}
	}
    return rows[rowIndex]
}