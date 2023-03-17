You are given an array of positive integers where each integer represents the height of a vertical line on a chart.

Find two lines which together with the x-axis forms a container that would hold the greatest amount of water.

Return the area of water it would hold, after water filling the max area:

Area: = L * W = 8 * 3 = 24

Constraints or questions to ask:
- does the left and right sides count as walls? No, sides cannot be used to form a container
- does the thickness of the lines affect the area? No, assume they take no space.
- does a higher line inside our container affect our area? No, lines inside a container don't affect the area.