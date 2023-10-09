package flags

import "flag"

type Values struct {
	CsvFileName *string
	MyBoolVar   *bool
}

var flagObj Values

func Initialize() {
	flagObj = Values{}

	/*
		We get pointer after defining a flag variable with a default value
		If user provides that flag while running program, this pointer will get that value automatically

		go run main.go -csv=whateverValue

		go run main.go -h will list all flags defined in program


		Full Details:
		https://github.com/golang/go/blob/master/src/flag/flag.go
	*/
	//flagObj.CsvFileName = flag.String("csv", "problems.csv", "CSV file containing quiz question answers")

	flagObj.MyBoolVar = flag.Bool("boolVar", false, "Trying bool variable")

	// to parse the command line into the defined flags.
	flag.Parse()
}

func Get() Values {
	return flagObj
}
