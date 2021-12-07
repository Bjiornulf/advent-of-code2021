# Advent of Code 2021

I decided to solve this years advent of code in **GO** for my computer science class. Other implementations can be found [here](https://github.com/mines-nancy/advent-of-code-2021)

## Journey

### Day 1
Day 1 was easy. I had a bit of trouble finding my way arround go modules and packages, but at the end I managed to understand how it worked.

### Day 2
Day 2 wasn't complicated either. I was the first one of my class to solve the problem, though I wasn't nearly fast enough to make it to the global leaderboard.
It was also an opportunity to benchmark `fmt.Sscanf` vs `stings.SplitN` + `strconv.Atoi`. The latter combination ended up being significantly faster (~10x)
and used fewer resources.

### Day 3
Day 3 was a bit harder, because I didn't know the language quiet that good. I knew how to solve the problme since the beginning, but it took me a bit more time. I also faced problems because I didn't fully read the problem in the first place, and some corner cases made the difference. In trying to go faster, I ended up going slower

### Day 4
Simple, but hard to make readable code. The code got pretty long very quickly, and is not very readable. I don't (yet) know how to split it in digestible bite sizes

### Day 5
Day 5 took me longer than the previous days. The idea was simpler, but I could get lost in the indices. Moreover, it was very easy to produce long and unreadable code, which required some time after finding the solution to clean up.

### Day 6
The first part could be trivialy solved without thinking and just implementing the problem the way it was presented. The second part required a bit more thinking as the array would grow exponentially, and eventualy lead to a crash. The result is a much more elegant solution.

### Day 7
Extremely simple. However I'm still too slow at solving the puzzle (probably using GO doesn't help with speed)
