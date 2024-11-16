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

/*
┌────────────────────────────────── WARNING ───────────────────────────────────┐
│                 THIS CLASS FILE WAS AUTOMATICALLY GENERATED.                 │
│                    ANY UPDATES TO IT WILL BE OVERWRITTEN.                    │
└──────────────────────────────────────────────────────────────────────────────┘
*/

package ast

import (
	abs "github.com/craterdog/go-collection-framework/v4/collection"
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func ClassSection() ClassSectionClassLike {
	return classSectionReference()
}

// Constructor Methods

func (c *classSectionClass_) Make(
	classDeclarations abs.Sequential[ClassDeclarationLike],
) ClassSectionLike {
	if uti.IsUndefined(classDeclarations) {
		panic("The \"classDeclarations\" attribute is required by this class.")
	}
	var instance = &classSection_{
		// Initialize the instance attributes.
		classDeclarations_: classDeclarations,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Primary Methods

func (v *classSection_) GetClass() ClassSectionClassLike {
	return classSectionReference()
}

// Attribute Methods

func (v *classSection_) GetClassDeclarations() abs.Sequential[ClassDeclarationLike] {
	return v.classDeclarations_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type classSection_ struct {
	// Declare the instance attributes.
	classDeclarations_ abs.Sequential[ClassDeclarationLike]
}

// Class Structure

type classSectionClass_ struct {
	// Declare the class constants.
}

// Class Reference

func classSectionReference() *classSectionClass_ {
	return classSectionReference_
}

var classSectionReference_ = &classSectionClass_{
	// Initialize the class constants.
}
