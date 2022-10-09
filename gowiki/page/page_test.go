package page

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetValue(t *testing.T) {
	testCase := []struct {
		field, in, want string
	}{
		{"Meta", "view", "view"},
		{"Body", "new body", "new body"},
	}

	for _, tc := range testCase {
		p := InitPage()
		if tc.field == "Body" {
			p.SetValue(tc.field, []byte(tc.in))
		} else {
			p.SetValue(tc.field, tc.in)
		}
		newValue, _ := p.GetValue(tc.field)
		if newValue != tc.want {
			t.Errorf("setValue error for field %q actual:%q, want: %q\n", tc.field, newValue, tc.want)
		}
	}
}

func TestReflect(t *testing.T) {
	testCase := []struct {
		in, want string
	}{
		{"view", "view"},
	}

	for _, tc := range testCase {
		p := InitPage()
		pV := reflect.ValueOf(p)
		pT := reflect.TypeOf(p)
		// fmt.Printf("kind for page reflect: %v\n", pT.Kind())
		if pT.Kind() == reflect.Ptr {
			pV = pV.Elem()
			pT = pT.Elem()
		}
		fieldsNum := pT.NumField()
		for i := 0; i < fieldsNum; i++ {
			fieldName := pT.Field(i).Name
			fieldValue := pV.FieldByName(fieldName)
			t.Logf("index %d: filedName: %s FiledValue: %q, filedType: %q\n", i+1, fieldName, fieldValue, fieldValue.Kind().String())
			if fieldValue.Kind() == reflect.String {
				fieldValue.SetString(tc.in)
				t.Logf("change value filedName: %s newFiledValue: %q\n", fieldName, fieldValue)
			}
			t.Logf("filedValue type: %q\n", fieldValue.Kind().String())
		}
		if p.(*Page).Meta != tc.want {
			pr := reflect.ValueOf(p)
			fmt.Printf("kind for page reflect: %v\n", pr.Kind())
			t.Errorf("setValue error for string:%q, actual: %q\n", tc.in, p.(*Page).Meta)
		}
	}

}
