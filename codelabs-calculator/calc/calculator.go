package calc

var (
	Operations map[string]operation = map[string]operation{
		"+": &SumType{input: 2},
		"-": &SubtractType{input: 2},
		"/": &DivideType{input: 2},
		"*": &MultiplyType{input: 2},
	}
)

type calculator struct {
	input int64
	output int64
	operator func()
}

type operation interface {
	Calculate() (int64, error)
}


type SumType calculator
type MultiplyType calculator
type DivideType calculator
type SubtractType calculator



func (s *SumType) Calculate() (int64, error) {
	var err error 
	s.output, err = Sum(s.input, s.input)
	if err != nil {
		return s.output, err
	}
	return s.output, nil
}

func (s *MultiplyType) Calculate() (int64, error) {
	var err error 
	s.output, err = Multiply(s.input, s.input)
	if err != nil {
		return s.output, err
	}
	return s.output, nil
}

func (s *DivideType) Calculate() (int64, error) {
	var err error 
	s.output, err = Divide(s.input, s.input)
	if err != nil {
		return s.output, err
	}
	return s.output, nil
}

func (s *SubtractType) Calculate() (int64, error) {
	var err error 
	s.output, err = Subtract(s.input, s.input)
	if err != nil {
		return s.output, err
	}
	return s.output, nil
}

func CheckOperation(op string) (operation, error) {
	v := Operations[op] 
	return v,nil
}
