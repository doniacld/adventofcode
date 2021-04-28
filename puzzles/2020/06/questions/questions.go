package questions

// Questions holds methods to compute the yeses sum answer
type Questions interface {
	GetYeses() int
	IncCounter([]byte)
	Reset()
}

// yesesSum holds answers and the yeses counter
type yesesSum struct {
	answers    map[byte]struct{}
	yesCounter int
}

// NewYesesSum returns a new yeses sum
func NewYesesSum() Questions {
	return &yesesSum{
		answers: make(map[byte]struct{}, 0),
	}
}

// GetYeses returns the counter of yeses
func (y *yesesSum) GetYeses() int {
	return y.yesCounter
}

// IncCounter increments the counter if the value is not seen
func (y *yesesSum) IncCounter(answers []byte) {
	for _, a := range answers {
		if _, ok := y.answers[a]; !ok {
			y.answers[a] = struct{}{}
		}
		y.yesCounter = len(y.answers)
	}
}

// Reset the yesesSum structure
func (y *yesesSum) Reset() {
	y.answers = make(map[byte]struct{}, 0)
	y.yesCounter = 0
}


type commonAnswers struct {
	answers        map[byte]int
	answersCounter int
}

// NewCommonAnswers returns the number of common answers in a group
func NewCommonAnswers() Questions {
	return &commonAnswers{
		answers: make(map[byte]int),
	}
}

// GetYeses returns the counter of yeses
func (c *commonAnswers) GetYeses() int {
	var counter int
	for _, occAnswers := range c.answers {
		if occAnswers == c.answersCounter {
			counter++
		}
	}
	return counter
}

// IncCounter increments the counter if the value is not seen
func (c *commonAnswers) IncCounter(answers []byte) {
	c.answersCounter++
	for _, a := range answers {
		 c.answers[a]++
	}
}

// Reset the yesesSum structure
func (c *commonAnswers) Reset() {
	c.answers = make(map[byte]int, 0)
	c.answersCounter = 0
}
