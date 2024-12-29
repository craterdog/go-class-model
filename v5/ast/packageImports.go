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
	col "github.com/craterdog/go-collection-framework/v5/collection"
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func PackageImportsClass() PackageImportsClassLike {
	return packageImportsClass()
}

// Constructor Methods

func (c *packageImportsClass_) PackageImports(
	importedPackages col.Sequential[ImportedPackageLike],
) PackageImportsLike {
	if uti.IsUndefined(importedPackages) {
		panic("The \"importedPackages\" attribute is required by this class.")
	}
	var instance = &packageImports_{
		// Initialize the instance attributes.
		importedPackages_: importedPackages,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *packageImports_) GetClass() PackageImportsClassLike {
	return packageImportsClass()
}

// Attribute Methods

func (v *packageImports_) GetImportedPackages() col.Sequential[ImportedPackageLike] {
	return v.importedPackages_
}

// PROTECTED INTERFACE

// Instance Structure

type packageImports_ struct {
	// Declare the instance attributes.
	importedPackages_ col.Sequential[ImportedPackageLike]
}

// Class Structure

type packageImportsClass_ struct {
	// Declare the class constants.
}

// Class Reference

func packageImportsClass() *packageImportsClass_ {
	return packageImportsClassReference_
}

var packageImportsClassReference_ = &packageImportsClass_{
	// Initialize the class constants.
}
