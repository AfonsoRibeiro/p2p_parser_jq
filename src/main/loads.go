package main

import (
	"os"

	gojq_extentions "example.com/gojq_extentions/src"

	"github.com/itchyny/gojq"
	"github.com/sirupsen/logrus"
)

// Filter Funciton
func with_function_compile_test() gojq.CompilerOption {
	return gojq.WithFunction("ctest", 1, 1, gojq_extentions.Compiled_test)
}

func load_jq(program_file string, options ...gojq.CompilerOption) *gojq.Code {
	buf, _ := os.ReadFile(program_file)

	program, err := gojq.Parse(string(buf))
	if err != nil {
		logrus.Errorf("load_jq parse %s: %+v", program_file, err)
		return nil
	}

	compiled_program, err := gojq.Compile(program, options...)
	if err != nil {
		logrus.Errorf("load_jq compile %s: %+v", program_file, err)
		return nil
	}

	return compiled_program
}

func load_filters(filters_dir string) []*gojq.Code {
	filters := make([]*gojq.Code, 0)

	files, err := os.ReadDir(filters_dir)
	if err != nil {
		logrus.Fatalf("load_filters open dir: %v", err)
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		full_filename := filters_dir + "/" + file.Name()
		if filter := load_jq(full_filename, with_function_compile_test()); filter != nil {
			filters = append(filters, filter)
		}
	}

	return filters
}
