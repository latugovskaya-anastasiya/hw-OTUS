package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)
type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	out := in
	for _, stage := range stages {
		out = runner(done, stage(out))
	}
	return out
}

func runner(done, in In) Out {
	innerOut := make(Bi)
	go func() {
		defer close(innerOut) // close when there is nothing to write
		for {
			select {
			case <-done:
				return
			case v, ok := <-in: // checking if we can read in the channel
				if !ok {
					return
				}
				// implementation of the ability to stop pipeline
				select {
				case <-done:
					return
				case innerOut <- v:
				}
			}
		}
	}()
	return innerOut
}
