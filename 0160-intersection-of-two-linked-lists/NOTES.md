При помощи двух указателей запускаем цикл, до тех пор, пока оба указателя не будут равны, для этого по очереди двигаем их до конца списка, и в случае, если они дошли до конца, то приравниваем указатель ко второму списку, тем самым в сумме они обойдут n+m длинны двух списков и в случае наличия пересечения, будет возвращен элемент списка, в ином случае вернётся значение nil. Асимптотически это будет O(n+m), так как в худшем случае, цикл обойдёт оба списка.