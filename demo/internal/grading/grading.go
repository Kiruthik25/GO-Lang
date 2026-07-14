package grading

type grader interface {
	Grade(mark int) string
	Outputvalue() int
}

type SchoolAGrader struct{}

type SchoolBGrader struct{}

func NewSchoolAGrader() grader {
	return &SchoolAGrader{}
}

func NewSchoolBGrader() grader {
	return &SchoolBGrader{}
}

func (s SchoolAGrader) Grade(mark int) string {

	switch {

	case mark >= 90:
		return "Excellent"

	default:
		return "fail"

	}
}


func (s SchoolBGrader) Grade(mark int) string {

	switch {

	case mark >= 90:
		return "OutStanding"

	default:
		return "fail"

	}
}

func (SchoolAGrader) Outputvalue() int {

	return 1

}

func (SchoolBGrader) Outputvalue() int {

	return 110

}