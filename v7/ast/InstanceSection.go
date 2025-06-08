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
	com "github.com/craterdog/go-component-framework/v7"
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func InstanceSectionClass() InstanceSectionClassLike {
	return instanceSectionClass()
}

// Constructor Methods

func (c *instanceSectionClass_) InstanceSection(
	delimiter string,
	instanceDeclarations com.ListLike[InstanceDeclarationLike],
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

func (v *instanceSection_) GetInstanceDeclarations() com.ListLike[InstanceDeclarationLike] {
	return v.instanceDeclarations_
}

// PROTECTED INTERFACE

// Instance Structure

type instanceSection_ struct {
	// Declare the instance attributes.
	delimiter_            string
	instanceDeclarations_ com.ListLike[InstanceDeclarationLike]
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
