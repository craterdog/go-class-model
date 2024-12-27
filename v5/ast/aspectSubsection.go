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
	col "github.com/craterdog/go-collection-framework/v5/collection"
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func AspectSubsectionClass() AspectSubsectionClassLike {
	return aspectSubsectionClass()
}

// Constructor Methods

func (c *aspectSubsectionClass_) AspectSubsection(
	aspectInterfaces col.Sequential[AspectInterfaceLike],
) AspectSubsectionLike {
	if uti.IsUndefined(aspectInterfaces) {
		panic("The \"aspectInterfaces\" attribute is required by this class.")
	}
	var instance = &aspectSubsection_{
		// Initialize the instance attributes.
		aspectInterfaces_: aspectInterfaces,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *aspectSubsection_) GetClass() AspectSubsectionClassLike {
	return aspectSubsectionClass()
}

// Attribute Methods

func (v *aspectSubsection_) GetAspectInterfaces() col.Sequential[AspectInterfaceLike] {
	return v.aspectInterfaces_
}

// PROTECTED INTERFACE

// Instance Structure

type aspectSubsection_ struct {
	// Declare the instance attributes.
	aspectInterfaces_ col.Sequential[AspectInterfaceLike]
}

// Class Structure

type aspectSubsectionClass_ struct {
	// Declare the class constants.
}

// Class Reference

func aspectSubsectionClass() *aspectSubsectionClass_ {
	return aspectSubsectionClassReference_
}

var aspectSubsectionClassReference_ = &aspectSubsectionClass_{
	// Initialize the class constants.
}
