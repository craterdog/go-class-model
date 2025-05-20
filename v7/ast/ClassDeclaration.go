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

func ClassDeclarationClass() ClassDeclarationClassLike {
	return classDeclarationClass()
}

// Constructor Methods

func (c *classDeclarationClass_) ClassDeclaration(
	declaration DeclarationLike,
	delimiter1 string,
	delimiter2 string,
	classMethods ClassMethodsLike,
	delimiter3 string,
) ClassDeclarationLike {
	if uti.IsUndefined(declaration) {
		panic("The \"declaration\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter1) {
		panic("The \"delimiter1\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter2) {
		panic("The \"delimiter2\" attribute is required by this class.")
	}
	if uti.IsUndefined(classMethods) {
		panic("The \"classMethods\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter3) {
		panic("The \"delimiter3\" attribute is required by this class.")
	}
	var instance = &classDeclaration_{
		// Initialize the instance attributes.
		declaration_:  declaration,
		delimiter1_:   delimiter1,
		delimiter2_:   delimiter2,
		classMethods_: classMethods,
		delimiter3_:   delimiter3,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *classDeclaration_) GetClass() ClassDeclarationClassLike {
	return classDeclarationClass()
}

// Attribute Methods

func (v *classDeclaration_) GetDeclaration() DeclarationLike {
	return v.declaration_
}

func (v *classDeclaration_) GetDelimiter1() string {
	return v.delimiter1_
}

func (v *classDeclaration_) GetDelimiter2() string {
	return v.delimiter2_
}

func (v *classDeclaration_) GetClassMethods() ClassMethodsLike {
	return v.classMethods_
}

func (v *classDeclaration_) GetDelimiter3() string {
	return v.delimiter3_
}

// PROTECTED INTERFACE

// Instance Structure

type classDeclaration_ struct {
	// Declare the instance attributes.
	declaration_  DeclarationLike
	delimiter1_   string
	delimiter2_   string
	classMethods_ ClassMethodsLike
	delimiter3_   string
}

// Class Structure

type classDeclarationClass_ struct {
	// Declare the class constants.
}

// Class Reference

func classDeclarationClass() *classDeclarationClass_ {
	return classDeclarationClassReference_
}

var classDeclarationClassReference_ = &classDeclarationClass_{
	// Initialize the class constants.
}
