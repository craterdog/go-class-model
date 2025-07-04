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

func PrincipalSubsectionClass() PrincipalSubsectionClassLike {
	return principalSubsectionClass()
}

// Constructor Methods

func (c *principalSubsectionClass_) PrincipalSubsection(
	delimiter string,
	principalMethods fra.ListLike[PrincipalMethodLike],
) PrincipalSubsectionLike {
	if uti.IsUndefined(delimiter) {
		panic("The \"delimiter\" attribute is required by this class.")
	}
	if uti.IsUndefined(principalMethods) {
		panic("The \"principalMethods\" attribute is required by this class.")
	}
	var instance = &principalSubsection_{
		// Initialize the instance attributes.
		delimiter_:        delimiter,
		principalMethods_: principalMethods,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *principalSubsection_) GetClass() PrincipalSubsectionClassLike {
	return principalSubsectionClass()
}

// Attribute Methods

func (v *principalSubsection_) GetDelimiter() string {
	return v.delimiter_
}

func (v *principalSubsection_) GetPrincipalMethods() fra.ListLike[PrincipalMethodLike] {
	return v.principalMethods_
}

// PROTECTED INTERFACE

// Instance Structure

type principalSubsection_ struct {
	// Declare the instance attributes.
	delimiter_        string
	principalMethods_ fra.ListLike[PrincipalMethodLike]
}

// Class Structure

type principalSubsectionClass_ struct {
	// Declare the class constants.
}

// Class Reference

func principalSubsectionClass() *principalSubsectionClass_ {
	return principalSubsectionClassReference_
}

var principalSubsectionClassReference_ = &principalSubsectionClass_{
	// Initialize the class constants.
}
