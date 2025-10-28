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

func ImportListClass() ImportListClassLike {
	return importListClass()
}

// Constructor Methods

func (c *importListClass_) ImportList(
	importedPackages fra.Sequential[ImportedPackageLike],
) ImportListLike {
	if uti.IsUndefined(importedPackages) {
		panic("The \"importedPackages\" attribute is required by this class.")
	}
	var instance = &importList_{
		// Initialize the instance attributes.
		importedPackages_: importedPackages,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *importList_) GetClass() ImportListClassLike {
	return importListClass()
}

// Attribute Methods

func (v *importList_) GetImportedPackages() fra.Sequential[ImportedPackageLike] {
	return v.importedPackages_
}

// PROTECTED INTERFACE

// Instance Structure

type importList_ struct {
	// Declare the instance attributes.
	importedPackages_ fra.Sequential[ImportedPackageLike]
}

// Class Structure

type importListClass_ struct {
	// Declare the class constants.
}

// Class Reference

func importListClass() *importListClass_ {
	return importListClassReference_
}

var importListClassReference_ = &importListClass_{
	// Initialize the class constants.
}
