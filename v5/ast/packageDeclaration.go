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
│   Updates to any section other than the Private Methods may be overwritten.  │
└──────────────────────────────────────────────────────────────────────────────┘
*/

package ast

import (
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func PackageDeclaration() PackageDeclarationClassLike {
	return packageDeclarationReference()
}

// Constructor Methods

func (c *packageDeclarationClass_) Make(
	legalNotice LegalNoticeLike,
	packageHeader PackageHeaderLike,
	packageImports PackageImportsLike,
) PackageDeclarationLike {
	if uti.IsUndefined(legalNotice) {
		panic("The \"legalNotice\" attribute is required by this class.")
	}
	if uti.IsUndefined(packageHeader) {
		panic("The \"packageHeader\" attribute is required by this class.")
	}
	if uti.IsUndefined(packageImports) {
		panic("The \"packageImports\" attribute is required by this class.")
	}
	var instance = &packageDeclaration_{
		// Initialize the instance attributes.
		legalNotice_:    legalNotice,
		packageHeader_:  packageHeader,
		packageImports_: packageImports,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *packageDeclaration_) GetClass() PackageDeclarationClassLike {
	return packageDeclarationReference()
}

// Attribute Methods

func (v *packageDeclaration_) GetLegalNotice() LegalNoticeLike {
	return v.legalNotice_
}

func (v *packageDeclaration_) GetPackageHeader() PackageHeaderLike {
	return v.packageHeader_
}

func (v *packageDeclaration_) GetPackageImports() PackageImportsLike {
	return v.packageImports_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type packageDeclaration_ struct {
	// Declare the instance attributes.
	legalNotice_    LegalNoticeLike
	packageHeader_  PackageHeaderLike
	packageImports_ PackageImportsLike
}

// Class Structure

type packageDeclarationClass_ struct {
	// Declare the class constants.
}

// Class Reference

func packageDeclarationReference() *packageDeclarationClass_ {
	return packageDeclarationReference_
}

var packageDeclarationReference_ = &packageDeclarationClass_{
	// Initialize the class constants.
}
