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

func PackageDeclarationClass() PackageDeclarationClassLike {
	return packageDeclarationClassReference()
}

// Constructor Methods

func (c *packageDeclarationClass_) PackageDeclaration(
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

// INSTANCE INTERFACE

// Principal Methods

func (v *packageDeclaration_) GetClass() PackageDeclarationClassLike {
	return packageDeclarationClassReference()
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

func packageDeclarationClassReference() *packageDeclarationClass_ {
	return packageDeclarationClassReference_
}

var packageDeclarationClassReference_ = &packageDeclarationClass_{
	// Initialize the class constants.
}
