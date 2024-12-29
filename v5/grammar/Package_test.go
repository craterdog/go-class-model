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

package grammar_test

import (
	gra "github.com/craterdog/go-class-model/v5/grammar"
	ass "github.com/stretchr/testify/assert"
	osx "os"
	tes "testing"
)

func TestRoundTrips(t *tes.T) {
	var modelFile = "./Package.go"
	var bytes, err = osx.ReadFile(modelFile)
	if err != nil {
		panic(err)
	}
	var source = string(bytes)
	var parser = gra.ParserClass().Parser()
	var model = parser.ParseSource(source)
	var validator = gra.ValidatorClass().Validator()
	validator.ValidateModel(model)
	var formatter = gra.FormatterClass().Formatter()
	var actual = formatter.FormatModel(model)
	ass.Equal(t, source, actual)
}
