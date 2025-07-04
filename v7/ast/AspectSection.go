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
	fra "github.com/craterdog/go-component-framework/v7"
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func AspectSectionClass() AspectSectionClassLike {
	return aspectSectionClass()
}

// Constructor Methods

func (c *aspectSectionClass_) AspectSection(
	delimiter string,
	aspectDeclarations fra.ListLike[AspectDeclarationLike],
) AspectSectionLike {
	if uti.IsUndefined(delimiter) {
		panic("The \"delimiter\" attribute is required by this class.")
	}
	if uti.IsUndefined(aspectDeclarations) {
		panic("The \"aspectDeclarations\" attribute is required by this class.")
	}
	var instance = &aspectSection_{
		// Initialize the instance attributes.
		delimiter_:          delimiter,
		aspectDeclarations_: aspectDeclarations,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *aspectSection_) GetClass() AspectSectionClassLike {
	return aspectSectionClass()
}

// Attribute Methods

func (v *aspectSection_) GetDelimiter() string {
	return v.delimiter_
}

func (v *aspectSection_) GetAspectDeclarations() fra.ListLike[AspectDeclarationLike] {
	return v.aspectDeclarations_
}

// PROTECTED INTERFACE

// Instance Structure

type aspectSection_ struct {
	// Declare the instance attributes.
	delimiter_          string
	aspectDeclarations_ fra.ListLike[AspectDeclarationLike]
}

// Class Structure

type aspectSectionClass_ struct {
	// Declare the class constants.
}

// Class Reference

func aspectSectionClass() *aspectSectionClass_ {
	return aspectSectionClassReference_
}

var aspectSectionClassReference_ = &aspectSectionClass_{
	// Initialize the class constants.
}
