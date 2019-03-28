package main

import ("testing"
		"bytes"
		"text/template")

func BenchmarkTest(b *testing.B) {

	b.Logf("current iteration is %d", b.N)
	testStr := "Champion is {{ .Team }}"
	data := &map[string]string{
		"Team": "Warriors",
	}

	var temp bytes.Buffer

	for i := 0; i < b.N; i++ {
		t, _ := template.New("test").Parse(testStr)
		t.Execute(&temp, data)
		temp.Reset()
	}
}

func BenchmarkTest1(b *testing.B){
	b.Logf("the current iteration for second test is %d", b.N)
	testStr := "Champion is {{ .Team }}"
	t, _ := template.New("test").Parse(testStr)
	data := &map[string]string{
		"Team": "Warriors",
	}

	var temp bytes.Buffer

	for i := 0; i < b.N; i++ {
		t.Execute(&temp, data)
		temp.Reset()
	}

}