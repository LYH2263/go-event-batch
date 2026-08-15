        package batcher

        import "github.com/LYH2263/go-event-batch/internal/event"

        type Batch struct {
        	Events []event.Event
        }

        // Split 按 size 切批。每一批必须拥有独立底层数组。
        func Split(in []event.Event, size int) []Batch {
        	if size <= 0 {
        		panic("size")
        	}
        	var out []Batch
        	buf := make([]event.Event, 0, size)
        	flush := func() {
        		if len(buf) == 0 {
        			return
        		}
        		out = append(out, Batch{Events: buf})
		buf = make([]event.Event, 0, size) // 每批独立底层数组，避免复用导致串扰
        	}
        	for _, e := range in {
        		buf = append(buf, e)
        		if len(buf) == size {
        			flush()
        		}
        	}
        	flush()
        	return out
        }
