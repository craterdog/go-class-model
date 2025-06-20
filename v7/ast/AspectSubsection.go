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
	fra "github.com/craterdog/go-component-framework/v7"
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func AspectSubsectionClass() AspectSubsectionClassLike {
	return aspectSubsectionClass()
}

// Constructor Methods

func (c *aspectSubsectionClass_) AspectSubsection(
	delimiter string,
	aspectInterfaces fra.ListLike[AspectInterfaceLike],
) AspectSubsectionLike {
	if uti.IsUndefined(delimiter) {
		panic("The \"delimiter\" attribute is required by this class.")
	}
	if uti.IsUndefined(aspectInterfaces) {
		panic("The \"aspectInterfaces\" attribute is required by this class.")
	}
	var instance = &aspectSubsection_{
		// Initialize the instance attributes.
		delimiter_:        delimiter,
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

func (v *aspectSubsection_) GetDelimiter() string {
	return v.delimiter_
}

func (v *aspectSubsection_) GetAspectInterfaces() fra.ListLike[AspectInterfaceLike] {
	return v.aspectInterfaces_
}

// PROTECTED INTERFACE

// Instance Structure

type aspectSubsection_ struct {
	// Declare the instance attributes.
	delimiter_        string
	aspectInterfaces_ fra.ListLike[AspectInterfaceLike]
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
