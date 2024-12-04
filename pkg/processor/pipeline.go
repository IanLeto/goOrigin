package processor

type Pipeline struct {
	Processors []Processor
}

func (receiver *Pipeline) Add(p []Processor) {
	receiver.Processors = append(receiver.Processors, p...)
}
