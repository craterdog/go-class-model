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
	mod "github.com/craterdog/go-class-model/v5"
	ass "github.com/stretchr/testify/assert"
	osx "os"
	tes "testing"
)

var testDirectories = []string{
	"../../go-test-framework/v5/ast/",
	"../../go-test-framework/v5/grammar/",
	"../../go-test-framework/v5/example/",
}

func TestRoundTrips(t *tes.T) {
	fmt.Println("Round Trip Tests:")
	for _, directory := range testDirectories {
		fmt.Printf("   %v\n", directory)
		var bytes, err = osx.ReadFile(directory + "Package.go")
		if err != nil {
			panic(err)
		}
		var source = string(bytes)
		var model = mod.ParseSource(source)
		mod.ValidateModel(model)
		var actual = mod.FormatModel(model)
		ass.Equal(t, source, actual)
	}
	fmt.Println("Done.")
}
