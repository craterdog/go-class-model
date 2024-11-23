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

func LegalNotice() LegalNoticeClassLike {
	return legalNoticeReference()
}

// Constructor Methods

func (c *legalNoticeClass_) Make(
	comment string,
) LegalNoticeLike {
	if uti.IsUndefined(comment) {
		panic("The \"comment\" attribute is required by this class.")
	}
	var instance = &legalNotice_{
		// Initialize the instance attributes.
		comment_: comment,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Primary Methods

func (v *legalNotice_) GetClass() LegalNoticeClassLike {
	return legalNoticeReference()
}

// Attribute Methods

func (v *legalNotice_) GetComment() string {
	return v.comment_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type legalNotice_ struct {
	// Declare the instance attributes.
	comment_ string
}

// Class Structure

type legalNoticeClass_ struct {
	// Declare the class constants.
}

// Class Reference

func legalNoticeReference() *legalNoticeClass_ {
	return legalNoticeReference_
}

var legalNoticeReference_ = &legalNoticeClass_{
	// Initialize the class constants.
}
