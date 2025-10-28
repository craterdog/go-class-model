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
	mod "github.com/craterdog/go-class-model/v8"
	uti "github.com/craterdog/go-essential-utilities/v8"
	ass "github.com/stretchr/testify/assert"
	sts "strings"
	tes "testing"
)

var modelFiles = []string{
	"./ast/package_api.go",
	"./grammar/package_api.go",
	"./test/package_api.go",
}

func TestRoundTrips(t *tes.T) {
	fmt.Println("Round Trip Tests:")
	for _, modelFile := range modelFiles {
		fmt.Printf("   %v\n", modelFile)
		var source = uti.ReadFile(modelFile)
		var model = mod.ParseSource(source)
		mod.ValidateModel(model)
		var actual = mod.FormatModel(model)
		source = sts.ReplaceAll(source, "\t", "    ")
		actual = sts.ReplaceAll(actual, "\t", "    ")
		ass.Equal(t, source, actual)
	}
	fmt.Println("Done.")
}
