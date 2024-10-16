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

package module_test

import (
	fmt "fmt"
	cla "github.com/craterdog/go-class-model/v5"
	ass "github.com/stretchr/testify/assert"
	osx "os"
	tes "testing"
)

var filenames = []string{
	"./ast/Package.go",
	"./grammar/Package.go",
	"./testdata/Package.go",
}

func TestRoundTrips(t *tes.T) {
	fmt.Println("Round Trip Tests:")
	for _, filename := range filenames {
		fmt.Printf("   %v\n", filename)
		// Read in the class model file.
		var bytes, err = osx.ReadFile(filename)
		if err != nil {
			panic(err)
		}
		var source = string(bytes)

		// Parse the source code for the class model.
		var model = cla.ParseSource(source)

		// Validate the class model.
		cla.ValidateModel(model)

		// Format the class model.
		var actual = cla.FormatModel(model)
		ass.Equal(t, source, actual)
	}
	fmt.Println("Done.")
}
