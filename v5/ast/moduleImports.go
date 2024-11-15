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

package ast

import (
	abs "github.com/craterdog/go-collection-framework/v4/collection"
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func ModuleImports() ModuleImportsClassLike {
	return moduleImportsReference()
}

// Constructor Methods

func (c *moduleImportsClass_) Make(
	importedPackages abs.Sequential[ImportedPackageLike],
) ModuleImportsLike {
	if uti.IsUndefined(importedPackages) {
		panic("The \"importedPackages\" attribute is required by this class.")
	}
	var instance = &moduleImports_{
		// Initialize the instance attributes.
		importedPackages_: importedPackages,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Primary Methods

func (v *moduleImports_) GetClass() ModuleImportsClassLike {
	return moduleImportsReference()
}

// Attribute Methods

func (v *moduleImports_) GetImportedPackages() abs.Sequential[ImportedPackageLike] {
	return v.importedPackages_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type moduleImports_ struct {
	// Declare the instance attributes.
	importedPackages_ abs.Sequential[ImportedPackageLike]
}

// Class Structure

type moduleImportsClass_ struct {
	// Declare the class constants.
}

// Class Reference

func moduleImportsReference() *moduleImportsClass_ {
	return moduleImportsReference_
}

var moduleImportsReference_ = &moduleImportsClass_{
	// Initialize the class constants.
}
