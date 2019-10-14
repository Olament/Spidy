package schedule

import "crawler/controller"

/* schedule between worker and controller */
type SimpleSchedule struct {
	input chan controller.Request
}

func (schedule *SimpleSchedule) Submit(request controller.Request) {
	go func() {
		schedule.input <- request
	}()
}

func (schedule *SimpleSchedule) Init(input chan controller.Request) {
	schedule.input = input
}
