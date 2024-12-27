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
│                 This class file was automatically generated.                 │
│                     Any updates to it may be overwritten.                    │
└──────────────────────────────────────────────────────────────────────────────┘
*/

package ast

import (
	abs "github.com/craterdog/go-collection-framework/v5/collection"
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func ClassSectionClass() ClassSectionClassLike {
	return classSectionClass()
}

// Constructor Methods

func (c *classSectionClass_) ClassSection(
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

// INSTANCE INTERFACE

// Principal Methods

func (v *classSection_) GetClass() ClassSectionClassLike {
	return classSectionClass()
}

// Attribute Methods

func (v *classSection_) GetClassDeclarations() abs.Sequential[ClassDeclarationLike] {
	return v.classDeclarations_
}

// PROTECTED INTERFACE

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

func classSectionClass() *classSectionClass_ {
	return classSectionClassReference_
}

var classSectionClassReference_ = &classSectionClass_{
	// Initialize the class constants.
}
