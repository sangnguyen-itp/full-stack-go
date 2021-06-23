package tests

import (
	"fmt"
	"github.com/cucumber/godog"
)

type CalcSuite struct {
	calc *Calculator
}

func (cs *CalcSuite) calculatorIsCleared() error {
	cs.calc.Clear()
	return nil
}

func (cs *CalcSuite) iAdd(value int) error {
	cs.calc.Add(value)
	return nil
}

func (cs *CalcSuite) iMultiplyBy(value int) error {
	cs.calc.MultiplyBy(value)
	fmt.Println(cs.calc.Result())
	return nil
}

func (cs *CalcSuite) iPress(value int) error {
	cs.calc.Press(value)
	fmt.Println(cs.calc.Result())
	return nil
}

func (cs *CalcSuite) iSubtract(value int) error {
	cs.calc.Sub(value)
	fmt.Println(cs.calc.Result())
	return nil
}

func (cs *CalcSuite) theResultShouldBe(expectedResult int) error {
	result := cs.calc.Result()
	if result == expectedResult {
		return nil
	}
	return fmt.Errorf("%d doesn't match expectation: %d", result, expectedResult)
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	s := &CalcSuite{
		calc: new(Calculator),
	}

	ctx.Step(`^calculator is cleared$`, s.calculatorIsCleared)
	ctx.Step(`^i add (\d+)$`, s.iAdd)
	ctx.Step(`^i multiply by (\d+)$`, s.iMultiplyBy)
	ctx.Step(`^i press (\d+)$`, s.iPress)
	ctx.Step(`^i subtract (\d+)$`, s.iSubtract)
	ctx.Step(`^the result should be (\d+)$`, s.theResultShouldBe)
}
