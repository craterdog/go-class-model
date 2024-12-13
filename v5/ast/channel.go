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

import ()

// CLASS INTERFACE

// Access Function

func ChannelClass() ChannelClassLike {
	return channelClassReference()
}

// Constructor Methods

func (c *channelClass_) Channel() ChannelLike {
	var instance = &channel_{
		// Initialize the instance attributes.
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *channel_) GetClass() ChannelClassLike {
	return channelClassReference()
}

// Attribute Methods

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type channel_ struct {
	// Declare the instance attributes.
}

// Class Structure

type channelClass_ struct {
	// Declare the class constants.
}

// Class Reference

func channelClassReference() *channelClass_ {
	return channelClassReference_
}

var channelClassReference_ = &channelClass_{
	// Initialize the class constants.
}
