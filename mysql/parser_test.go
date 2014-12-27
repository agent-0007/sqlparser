package mysql

import (
	"reflect"
	"testing"
)

func TestParseDropTableStatement(t *testing.T) {
	parseDropTableStatement(t, "DROP TABLE hoge", &DropTableStatement{TableNames:[]TableNameIdentifier{TableNameIdentifier{Name:"hoge"}}})
	parseDropTableStatement(t, "drop table hoge,fuga", &DropTableStatement{TableNames:[]TableNameIdentifier{TableNameIdentifier{Name:"fuga"}, TableNameIdentifier{Name:"hoge"}}})
	parseDropTableStatement(t, "drop table `hoge`", &DropTableStatement{TableNames:[]TableNameIdentifier{TableNameIdentifier{Name:"hoge"}}})
	parseDropTableStatement(t, "drop table hoge.fuga", &DropTableStatement{TableNames:[]TableNameIdentifier{TableNameIdentifier{Database: "hoge", Name:"fuga"}}})
}

func parseDropTableStatement(t *testing.T, src string, expect interface{}) {
	s := new(Scanner)
	s.Init(src + ";")
	statements := Parse(s)
	if len(statements) != 1 {
		t.Errorf("Expect %q to be parsed, but %+#v", src, statements)
		return
	}
	if !reflect.DeepEqual(statements[0], expect) {
		t.Errorf("Expect %+#v to be %+#v", statements[0], expect)
		return
	}
}
