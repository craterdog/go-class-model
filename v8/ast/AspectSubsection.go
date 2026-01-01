/*
................................................................................
.    Copyright (c) 2009-2026 Crater Dog Technologies™.  All Rights Reserved.   .
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
	com "github.com/craterdog/go-essential-composites/v8"
	uti "github.com/craterdog/go-essential-utilities/v8"
)

// CLASS INTERFACE

// Access Function

func AspectSubsectionClass() AspectSubsectionClassLike {
	return aspectSubsectionClass()
}

// Constructor Methods

func (c *aspectSubsectionClass_) AspectSubsection(
	delimiter string,
	aspectInterfaces com.Sequential[AspectInterfaceLike],
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

func (v *aspectSubsection_) GetAspectInterfaces() com.Sequential[AspectInterfaceLike] {
	return v.aspectInterfaces_
}

// PROTECTED INTERFACE

// Instance Structure

type aspectSubsection_ struct {
	// Declare the instance attributes.
	delimiter_        string
	aspectInterfaces_ com.Sequential[AspectInterfaceLike]
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
