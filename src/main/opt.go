package main

import (
	"github.com/jnovack/flag"

	p2p_parser "example.com/p2p_parser/src"
)

type opt struct {
	p2p_parser_opt p2p_parser.Otp

	filtersdir string

	pprofon       bool
	pprofdir      string
	pprofduration uint

	prometheusport uint
	loglevel       string
}

func from_args() opt {

	var opt opt

	flag.StringVar(&opt.filtersdir, "filters_dir", "./private_filters/filters/", "Directory of all the jq filters files")

	flag.BoolVar(&opt.pprofon, "pprof_on", false, "Profoling on?")
	flag.StringVar(&opt.pprofdir, "pprof_dir", "./pprof", "Directory for pprof file")
	flag.UintVar(&opt.pprofduration, "pprof_duration", 60*4, "Number of seconds to run pprof")

	flag.UintVar(&opt.prometheusport, "prometheus_port", 7700, "Prometheous port")
	flag.StringVar(&opt.loglevel, "log_level", "info", "Logging level: panic - fatal - error - warn - info - debug - trace")

	opt.p2p_parser_opt = p2p_parser.From_args()

	return opt

}
