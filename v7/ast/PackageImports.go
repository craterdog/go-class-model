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
│                 This class file was automatically generated.                 │
│                     Any updates to it may be overwritten.                    │
└──────────────────────────────────────────────────────────────────────────────┘
*/

package ast

import (
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func PackageImportsClass() PackageImportsClassLike {
	return packageImportsClass()
}

// Constructor Methods

func (c *packageImportsClass_) PackageImports(
	delimiter1 string,
	delimiter2 string,
	optionalImportList ImportListLike,
	delimiter3 string,
) PackageImportsLike {
	if uti.IsUndefined(delimiter1) {
		panic("The \"delimiter1\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter2) {
		panic("The \"delimiter2\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter3) {
		panic("The \"delimiter3\" attribute is required by this class.")
	}
	var instance = &packageImports_{
		// Initialize the instance attributes.
		delimiter1_:         delimiter1,
		delimiter2_:         delimiter2,
		optionalImportList_: optionalImportList,
		delimiter3_:         delimiter3,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *packageImports_) GetClass() PackageImportsClassLike {
	return packageImportsClass()
}

// Attribute Methods

func (v *packageImports_) GetDelimiter1() string {
	return v.delimiter1_
}

func (v *packageImports_) GetDelimiter2() string {
	return v.delimiter2_
}

func (v *packageImports_) GetOptionalImportList() ImportListLike {
	return v.optionalImportList_
}

func (v *packageImports_) GetDelimiter3() string {
	return v.delimiter3_
}

// PROTECTED INTERFACE

// Instance Structure

type packageImports_ struct {
	// Declare the instance attributes.
	delimiter1_         string
	delimiter2_         string
	optionalImportList_ ImportListLike
	delimiter3_         string
}

// Class Structure

type packageImportsClass_ struct {
	// Declare the class constants.
}

// Class Reference

func packageImportsClass() *packageImportsClass_ {
	return packageImportsClassReference_
}

var packageImportsClassReference_ = &packageImportsClass_{
	// Initialize the class constants.
}
