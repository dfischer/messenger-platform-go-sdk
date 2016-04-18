package template

import "testing"

func TestGenericType(t *testing.T) {
	template := &GenericTemplate{}
	if template.Type() != TemplateTypeGeneric {
		t.Error("Generic template returned invalid type")
	}
}
