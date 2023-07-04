A company has n employees with unique IDs from 0 to n-1. The head of the company has the ID headID.
n = 8       8 employees: 0, 1, 2, 3, **4**, 5, 6, 7        
                                    headID = 4

You will receive a managers array where managers[i] is the ID of the manage for employee i. Each employee has one direct manager. The company head has no manager (managers[headID] = -1). It's guaranteed the subordination relationships will have a tree structure.

n = 8       8 employees:  0, 1, 2, 3,  4, 5, 6, 7     headID = 4
                managers=[2, 3, 4, 6, -1, 4, 4, 5]

The head of the company wants to inform all employees of news. He will inform his direct subordinates who will inform 
their direct subordinates and so on until everyone knows the news.

You will receive an informTime array where informTime[i] is the time it takes for employee i to inform all their direct
subordinates.
              0, 1, 2, 3, 4, 5, 6, 7
informTime = [0, 0, 4, 0, 7, 3, 6, 0]

Return the total number of minutes it takes to inform all employees of the news. 

Here in the above example the answer to total number of minutes is: 7 + max(3, 4, 6) = 7 + 6 = 13