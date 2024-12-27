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
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func ImportedPackageClass() ImportedPackageClassLike {
	return importedPackageClass()
}

// Constructor Methods

func (c *importedPackageClass_) ImportedPackage(
	name string,
	path string,
) ImportedPackageLike {
	if uti.IsUndefined(name) {
		panic("The \"name\" attribute is required by this class.")
	}
	if uti.IsUndefined(path) {
		panic("The \"path\" attribute is required by this class.")
	}
	var instance = &importedPackage_{
		// Initialize the instance attributes.
		name_: name,
		path_: path,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *importedPackage_) GetClass() ImportedPackageClassLike {
	return importedPackageClass()
}

// Attribute Methods

func (v *importedPackage_) GetName() string {
	return v.name_
}

func (v *importedPackage_) GetPath() string {
	return v.path_
}

// PROTECTED INTERFACE

// Instance Structure

type importedPackage_ struct {
	// Declare the instance attributes.
	name_ string
	path_ string
}

// Class Structure

type importedPackageClass_ struct {
	// Declare the class constants.
}

// Class Reference

func importedPackageClass() *importedPackageClass_ {
	return importedPackageClassReference_
}

var importedPackageClassReference_ = &importedPackageClass_{
	// Initialize the class constants.
}
