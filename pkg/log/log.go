package log

import (
	"k8s.io/klog"
	"log"
)

// KlogWriter serves as a bridge between the standard log package and the glog package.
type KlogWriter struct{}

// Write implements the io.Writer interface.
func (writer KlogWriter) Write(data []byte) (n int, err error) {
	klog.InfoDepth(1, string(data))
	return len(data), nil
}

func InitLogs() {
	log.SetOutput(KlogWriter{})
	log.SetFlags(0)
}
