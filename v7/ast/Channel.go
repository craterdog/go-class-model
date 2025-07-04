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
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func ChannelClass() ChannelClassLike {
	return channelClass()
}

// Constructor Methods

func (c *channelClass_) Channel(
	delimiter string,
) ChannelLike {
	if uti.IsUndefined(delimiter) {
		panic("The \"delimiter\" attribute is required by this class.")
	}
	var instance = &channel_{
		// Initialize the instance attributes.
		delimiter_: delimiter,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *channel_) GetClass() ChannelClassLike {
	return channelClass()
}

// Attribute Methods

func (v *channel_) GetDelimiter() string {
	return v.delimiter_
}

// PROTECTED INTERFACE

// Instance Structure

type channel_ struct {
	// Declare the instance attributes.
	delimiter_ string
}

// Class Structure

type channelClass_ struct {
	// Declare the class constants.
}

// Class Reference

func channelClass() *channelClass_ {
	return channelClassReference_
}

var channelClassReference_ = &channelClass_{
	// Initialize the class constants.
}
