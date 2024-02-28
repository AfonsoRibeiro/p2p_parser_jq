package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"time"

	"github.com/sirupsen/logrus"

	p2p_parser "example.com/p2p_parser/src"
	"example.com/p2p_parser_jq/src/parser"
)

func logging(level string) {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02T15:04:05.000Z07:00",
	})
	l, err := logrus.ParseLevel(level)
	if err != nil {
		logrus.Errorf("Failed parse log level. Reason: %+v", err)
	} else {
		logrus.SetLevel(l)
	}
}

func main() {
	opt := from_args()
	logging(opt.loglevel)
	logrus.Infof("%+v", opt)

	setup_prometheus(opt.prometheusport)

	parser := &parser.Parser{
		Filters: load_filters(opt.filtersdir),
	}

	if len(parser.Filters) == 0 {
		logrus.Panicln("No valid filters.")
	}

	if opt.pprof_on {
		f, err := os.Create(fmt.Sprintf("profiles/%s.prof", time.Now().Format("2006-01-02_15:04:05")))
		if err != nil {
			logrus.Fatal("Failed to open file for profiling", err)
		}

		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	p2p_parser.P2P_parser(opt.p2p_parser_opt, parser)
}
