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
	fra "github.com/craterdog/go-essential-composites/v8"
	uti "github.com/craterdog/go-missing-utilities/v8"
)

// CLASS INTERFACE

// Access Function

func InstanceSectionClass() InstanceSectionClassLike {
	return instanceSectionClass()
}

// Constructor Methods

func (c *instanceSectionClass_) InstanceSection(
	delimiter string,
	instanceDeclarations fra.Sequential[InstanceDeclarationLike],
) InstanceSectionLike {
	if uti.IsUndefined(delimiter) {
		panic("The \"delimiter\" attribute is required by this class.")
	}
	if uti.IsUndefined(instanceDeclarations) {
		panic("The \"instanceDeclarations\" attribute is required by this class.")
	}
	var instance = &instanceSection_{
		// Initialize the instance attributes.
		delimiter_:            delimiter,
		instanceDeclarations_: instanceDeclarations,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *instanceSection_) GetClass() InstanceSectionClassLike {
	return instanceSectionClass()
}

// Attribute Methods

func (v *instanceSection_) GetDelimiter() string {
	return v.delimiter_
}

func (v *instanceSection_) GetInstanceDeclarations() fra.Sequential[InstanceDeclarationLike] {
	return v.instanceDeclarations_
}

// PROTECTED INTERFACE

// Instance Structure

type instanceSection_ struct {
	// Declare the instance attributes.
	delimiter_            string
	instanceDeclarations_ fra.Sequential[InstanceDeclarationLike]
}

// Class Structure

type instanceSectionClass_ struct {
	// Declare the class constants.
}

// Class Reference

func instanceSectionClass() *instanceSectionClass_ {
	return instanceSectionClassReference_
}

var instanceSectionClassReference_ = &instanceSectionClass_{
	// Initialize the class constants.
}
