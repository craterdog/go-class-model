/*
................................................................................
.    Copyright (c) 2009-2025 Crater Dog Technologies.  All Rights Reserved.    .
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
	mod "github.com/craterdog/go-class-model/v7"
	uti "github.com/craterdog/go-missing-utilities/v7"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

var modelFiles = []string{
	"./ast/package_api.go",
	"./grammar/package_api.go",
	"./testdata/package_api.go",
}

func TestRoundTrips(t *tes.T) {
	fmt.Println("Round Trip Tests:")
	for _, modelFile := range modelFiles {
		fmt.Printf("   %v\n", modelFile)
		var source = uti.ReadFile(modelFile)
		var model = mod.ParseSource(source)
		mod.ValidateModel(model)
		var actual = mod.FormatModel(model)
		//fmt.Println(actual)
		ass.Equal(t, source, actual)
	}
	fmt.Println("Done.")
}
