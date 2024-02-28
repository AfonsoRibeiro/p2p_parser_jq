package main

import (
	"github.com/jnovack/flag"

	p2p_parser "example.com/p2p_parser/src"
)

type opt struct {
	p2p_parser_opt p2p_parser.Otp

	filtersdir string

	prometheusport uint
	pprof_on       bool
	loglevel       string
}

func from_args() opt {

	var opt opt

	flag.StringVar(&opt.filtersdir, "filters_dir", "./filters/", "Directory of all the jq filters files")

	flag.UintVar(&opt.prometheusport, "prometheus_port", 7700, "Prometheous port")
	flag.StringVar(&opt.loglevel, "log_level", "info", "Logging level: panic - fatal - error - warn - info - debug - trace")

	opt.p2p_parser_opt = p2p_parser.From_args()

	return opt

}
