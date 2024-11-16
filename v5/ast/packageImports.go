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

func PackageImports() PackageImportsClassLike {
	return packageImportsReference()
}

// Constructor Methods

func (c *packageImportsClass_) Make(
	importedPackages abs.Sequential[ImportedPackageLike],
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

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Primary Methods

func (v *packageImports_) GetClass() PackageImportsClassLike {
	return packageImportsReference()
}

// Attribute Methods

func (v *packageImports_) GetImportedPackages() abs.Sequential[ImportedPackageLike] {
	return v.importedPackages_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type packageImports_ struct {
	// Declare the instance attributes.
	importedPackages_ abs.Sequential[ImportedPackageLike]
}

// Class Structure

type packageImportsClass_ struct {
	// Declare the class constants.
}

// Class Reference

func packageImportsReference() *packageImportsClass_ {
	return packageImportsReference_
}

var packageImportsReference_ = &packageImportsClass_{
	// Initialize the class constants.
}
