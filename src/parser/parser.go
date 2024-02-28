package parser

import (
	"encoding/json"

	"github.com/itchyny/gojq"
	"github.com/sirupsen/logrus"
)

type Parser struct {
	Filters []*gojq.Code
}

func (dn *Parser) Parse(msg []byte) [][]byte {
	parsed := make([][]byte, 0)

	for _, filter := range dn.Filters {
		//fmt.Printf("%d filter: \n", a)

		var msg_json interface{}

		if err := json.Unmarshal(msg, &msg_json); err != nil { // Before for loop but copy the json struct some how... Does it even need to be copied??
			logrus.Errorf("Parse unmarshal msg: %+v", err)
			return parsed
		}
		iter := filter.Run(msg_json)
		for {
			//fmt.Printf("%#v\n", iter)
			v, ok := iter.Next()
			if !ok {
				break
			}
			if _, ok := v.(error); ok {
				continue
			} else {
				msg, err := json.Marshal(v)
				if err != nil {
					logrus.Errorf("Parse marshal parsed msg: %+v", err)
					continue
				}
				parsed = append(parsed, msg)
			}
		}
	}

	return parsed
}
