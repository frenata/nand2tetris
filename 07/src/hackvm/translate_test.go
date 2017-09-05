package hackvm_test

import (
	"hackvm"
	"testing"
)

func TestTranslate(t *testing.T) {
	t.Log("\n" + hackvm.Translate("push constant 17"))
	t.Log("\n" + hackvm.Translate("add"))
}
