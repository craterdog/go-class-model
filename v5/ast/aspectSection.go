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

func AspectSectionClass() AspectSectionClassLike {
	return aspectSectionClassReference()
}

// Constructor Methods

func (c *aspectSectionClass_) AspectSection(
	aspectDeclarations abs.Sequential[AspectDeclarationLike],
) AspectSectionLike {
	if uti.IsUndefined(aspectDeclarations) {
		panic("The \"aspectDeclarations\" attribute is required by this class.")
	}
	var instance = &aspectSection_{
		// Initialize the instance attributes.
		aspectDeclarations_: aspectDeclarations,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *aspectSection_) GetClass() AspectSectionClassLike {
	return aspectSectionClassReference()
}

// Attribute Methods

func (v *aspectSection_) GetAspectDeclarations() abs.Sequential[AspectDeclarationLike] {
	return v.aspectDeclarations_
}

// PROTECTED INTERFACE

// Instance Structure

type aspectSection_ struct {
	// Declare the instance attributes.
	aspectDeclarations_ abs.Sequential[AspectDeclarationLike]
}

// Class Structure

type aspectSectionClass_ struct {
	// Declare the class constants.
}

// Class Reference

func aspectSectionClassReference() *aspectSectionClass_ {
	return aspectSectionClassReference_
}

var aspectSectionClassReference_ = &aspectSectionClass_{
	// Initialize the class constants.
}
