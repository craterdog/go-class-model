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

func AspectSection() AspectSectionClassLike {
	return aspectSectionReference()
}

// Constructor Methods

func (c *aspectSectionClass_) Make(
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

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Primary Methods

func (v *aspectSection_) GetClass() AspectSectionClassLike {
	return aspectSectionReference()
}

// Attribute Methods

func (v *aspectSection_) GetAspectDeclarations() abs.Sequential[AspectDeclarationLike] {
	return v.aspectDeclarations_
}

// PROTECTED INTERFACE

// Private Methods

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

func aspectSectionReference() *aspectSectionClass_ {
	return aspectSectionReference_
}

var aspectSectionReference_ = &aspectSectionClass_{
	// Initialize the class constants.
}
