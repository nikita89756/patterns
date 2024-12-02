package handler

import(
	"testing"
)

func TestParserHandlerNext(t *testing.T){
	base:=&BaseParseHandler{}
	base.SetNext(nil)
	json:=&JSONHandler{}
	json.SetNext(base)
	toml:=&TomlHandler{}
	toml.SetNext(json)
	xml:=&XmlHandler{}
	xml.SetNext(toml)

	if *xml.nextHandler!=toml{
		t.Errorf("Ожидали xml.nextHandler = YamlHandler, но имеем %v",xml.nextHandler)
	}
	if *toml.nextHandler!=json{
		t.Errorf("Ожидали yaml.nextHandler = JSONHandler, но имеем %v",toml.nextHandler)
	}
	if *json.nextHandler!=base{
		t.Errorf("Ожидали json.nextHandler = BaseParseHandler, но имеем %v",json.nextHandler)
	}
		if *base.nextHandler!=nil{
		t.Errorf("Ожидали base.nextHandler = nil, но имеем %v",base.nextHandler)
	}

}

func TestParsing(t *testing.T) {
	base:=&BaseParseHandler{}
	json:=&JSONHandler{}
	json.SetNext(base)
	toml:=&TomlHandler{}
	toml.SetNext(json)
	xml:=&XmlHandler{}
	xml.SetNext(toml)

	xmltest:=`
	<person>
		<name>John Doe</name>
		<age>30</age>
		<email>john.doe@example.com</email>
	</person>`

	tomltest:=`
	name = "John Doe"
	age = 30
	[address]
	street = "123 Main St"
	city = "Anytown"
	state = "CA"
	`

	jsontest:=`
	{
		"name": "John Doe",
		"age": 30,
		"email": "john.doe@example.com"
	}
	`

	basetest:=`HelloWorld`
	
	if str:=xml.Handle(basetest); str != "string"{
		t.Errorf("Ожидали значение string, но вернулось %s",str)
	}

	if str:=xml.Handle(jsontest); str != "json"{
		t.Errorf("Ожидали значение json, но вернулось %s",str)
	}

	if str:=xml.Handle(tomltest); str != "toml"{
		t.Errorf("Ожидали значение toml, но вернулось %s",str)
	}
	if str:=xml.Handle(xmltest); str != "xml"{
		t.Errorf("Ожидали значение xml, но вернулось %s",str)
	}
}