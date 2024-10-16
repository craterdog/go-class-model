/*
................................................................................
.    Copyright (c) 2009-2024 Crater Dog Technologies.  All Rights Reserved.    .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

package grammar_test

import (
	fmt "fmt"
	gra "github.com/craterdog/go-class-model/v5/grammar"
	ass "github.com/stretchr/testify/assert"
	osx "os"
	tes "testing"
)

var inputDirectory = "../"
var outputDirectory = "../../../go-test-framework/v5/"
var testDirectories = []string{
	"ast/",
	"grammar/",
	"testdata/",
}

func TestRoundTrips(t *tes.T) {
	fmt.Println("Round Trip Tests:")
	for _, testDirectory := range testDirectories {
		var directory = inputDirectory + testDirectory
		fmt.Printf("   %v\n", directory)
		var bytes, err = osx.ReadFile(directory + "Package.go")
		if err != nil {
			panic(err)
		}
		var source = string(bytes)
		var parser = gra.Parser().Make()
		var model = parser.ParseSource(source)
		var validator = gra.Validator().Make()
		validator.ValidateModel(model)
		var formatter = gra.Formatter().Make()
		var actual = formatter.FormatModel(model)
		ass.Equal(t, source, actual)
		directory = outputDirectory + testDirectory
		var filename = directory + "Package.go"
		err = osx.WriteFile(filename, bytes, 0644)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("Done.")
}
