package rc3

# A function that returns a random number between 0 and 1
func Random() float64 {
	return rand.Float64()
}

# A function that divided users into two groups
func Split(users []string) (group1, group2 []string) {
	group1 = users[:len(users)/2]
	group2 = users[len(users)/2:]
	return group1, group2
}

# A function that simulate the roll of a dice
func Roll() int {
	return rand.Intn(6) + 1
}