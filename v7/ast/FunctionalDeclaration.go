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

/*
┌────────────────────────────────── WARNING ───────────────────────────────────┐
│              This class file was automatically generated using:              │
│            https://github.com/craterdog/go-development-tools/wiki            │
│                                                                              │
│                     Any updates to it may be overwritten.                    │
└──────────────────────────────────────────────────────────────────────────────┘
*/

package ast

import (
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func FunctionalDeclarationClass() FunctionalDeclarationClassLike {
	return functionalDeclarationClass()
}

// Constructor Methods

func (c *functionalDeclarationClass_) FunctionalDeclaration(
	declaration DeclarationLike,
	delimiter1 string,
	delimiter2 string,
	optionalParameterList ParameterListLike,
	delimiter3 string,
	result ResultLike,
) FunctionalDeclarationLike {
	if uti.IsUndefined(declaration) {
		panic("The \"declaration\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter1) {
		panic("The \"delimiter1\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter2) {
		panic("The \"delimiter2\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter3) {
		panic("The \"delimiter3\" attribute is required by this class.")
	}
	if uti.IsUndefined(result) {
		panic("The \"result\" attribute is required by this class.")
	}
	var instance = &functionalDeclaration_{
		// Initialize the instance attributes.
		declaration_:           declaration,
		delimiter1_:            delimiter1,
		delimiter2_:            delimiter2,
		optionalParameterList_: optionalParameterList,
		delimiter3_:            delimiter3,
		result_:                result,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *functionalDeclaration_) GetClass() FunctionalDeclarationClassLike {
	return functionalDeclarationClass()
}

// Attribute Methods

func (v *functionalDeclaration_) GetDeclaration() DeclarationLike {
	return v.declaration_
}

func (v *functionalDeclaration_) GetDelimiter1() string {
	return v.delimiter1_
}

func (v *functionalDeclaration_) GetDelimiter2() string {
	return v.delimiter2_
}

func (v *functionalDeclaration_) GetOptionalParameterList() ParameterListLike {
	return v.optionalParameterList_
}

func (v *functionalDeclaration_) GetDelimiter3() string {
	return v.delimiter3_
}

func (v *functionalDeclaration_) GetResult() ResultLike {
	return v.result_
}

// PROTECTED INTERFACE

// Instance Structure

type functionalDeclaration_ struct {
	// Declare the instance attributes.
	declaration_           DeclarationLike
	delimiter1_            string
	delimiter2_            string
	optionalParameterList_ ParameterListLike
	delimiter3_            string
	result_                ResultLike
}

// Class Structure

type functionalDeclarationClass_ struct {
	// Declare the class constants.
}

// Class Reference

func functionalDeclarationClass() *functionalDeclarationClass_ {
	return functionalDeclarationClassReference_
}

var functionalDeclarationClassReference_ = &functionalDeclarationClass_{
	// Initialize the class constants.
}
