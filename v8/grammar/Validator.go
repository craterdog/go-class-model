/*
................................................................................
.    Copyright (c) 2009-2026 Crater Dog Technologies™.  All Rights Reserved.   .
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
│ Updates to any section other than the Methodical Methods may be overwritten. │
└──────────────────────────────────────────────────────────────────────────────┘
*/

package grammar

import (
	fmt "fmt"
	ast "github.com/craterdog/go-class-model/v8/ast"
	com "github.com/craterdog/go-essential-composites/v8"
	sts "strings"
	utf "unicode/utf8"
)

// CLASS INTERFACE

// Access Function

func ValidatorClass() ValidatorClassLike {
	return validatorClass()
}

// Constructor Methods

func (c *validatorClass_) Validator() ValidatorLike {
	var instance = &validator_{
		// Initialize the instance attributes.

		// Initialize the inherited aspects.
		Methodical: ProcessorClass().Processor(),
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *validator_) GetClass() ValidatorClassLike {
	return validatorClass()
}

func (v *validator_) ValidateModel(
	model ast.ModelLike,
) {
	VisitorClass().Visitor(v).VisitModel(model)
}

// Methodical Methods

func (v *validator_) ProcessComment(
	comment string,
) {
	v.validateToken(comment, CommentToken)
}

func (v *validator_) ProcessName(
	name string,
) {
	v.validateToken(name, NameToken)
}

func (v *validator_) ProcessNewline(
	newline string,
) {
	v.validateToken(newline, NewlineToken)
}

func (v *validator_) ProcessPath(
	path string,
) {
	v.validateToken(path, PathToken)
}

func (v *validator_) ProcessPrefix(
	prefix string,
) {
	v.validateToken(prefix, PrefixToken)
}

func (v *validator_) PreprocessAspectSection(
	aspectSection ast.AspectSectionLike,
	index_ uint,
	count_ uint,
) {
	// We can "cheat" here because we know the code generator uses lists.
	var aspectDeclarations = aspectSection.GetAspectDeclarations().(com.ListLike[ast.AspectDeclarationLike])
	aspectDeclarations.SortValuesWithRanker(
		func(
			first ast.AspectDeclarationLike,
			second ast.AspectDeclarationLike,
		) com.Rank {
			var firstName = first.GetDeclaration().GetName()
			var secondName = second.GetDeclaration().GetName()
			switch {
			case firstName < secondName:
				return com.LesserRank
			case firstName > secondName:
				return com.GreaterRank
			default:
				return com.EqualRank
			}
		},
	)
}

func (v *validator_) PreprocessClassSection(
	classSection ast.ClassSectionLike,
	index_ uint,
	count_ uint,
) {
	// We can "cheat" here because we know the code generator uses lists.
	var classDeclarations = classSection.GetClassDeclarations().(com.ListLike[ast.ClassDeclarationLike])
	classDeclarations.SortValuesWithRanker(
		func(
			first ast.ClassDeclarationLike,
			second ast.ClassDeclarationLike,
		) com.Rank {
			var firstName = sts.TrimSuffix(
				first.GetDeclaration().GetName(),
				"ClassLike",
			)
			var secondName = sts.TrimSuffix(
				second.GetDeclaration().GetName(),
				"ClassLike",
			)
			switch {
			case firstName < secondName:
				return com.LesserRank
			case firstName > secondName:
				return com.GreaterRank
			default:
				return com.EqualRank
			}
		},
	)
}

func (v *validator_) PreprocessFunctionalSection(
	functionalSection ast.FunctionalSectionLike,
	index_ uint,
	count_ uint,
) {
	// We can "cheat" here because we know the code generator uses lists.
	var functionalDeclarations = functionalSection.GetFunctionalDeclarations().(com.ListLike[ast.FunctionalDeclarationLike])
	functionalDeclarations.SortValuesWithRanker(
		func(
			first ast.FunctionalDeclarationLike,
			second ast.FunctionalDeclarationLike,
		) com.Rank {
			var firstName = sts.TrimSuffix(
				first.GetDeclaration().GetName(),
				"Function",
			)
			var secondName = sts.TrimSuffix(
				second.GetDeclaration().GetName(),
				"Function",
			)
			switch {
			case firstName < secondName:
				return com.LesserRank
			case firstName > secondName:
				return com.GreaterRank
			default:
				return com.EqualRank
			}
		},
	)
}

func (v *validator_) PreprocessFunctionMethod(
	functionMethod ast.FunctionMethodLike,
	index_ uint,
	count_ uint,
) {
	var result = functionMethod.GetResult()
	switch result.GetAny().(type) {
	case ast.NoneLike:
		var functionName = functionMethod.GetName()
		var message = fmt.Sprintf(
			"A function method must include a result type for the function: %s",
			functionName,
		)
		panic(message)
	}
}

func (v *validator_) PreprocessImportedPackage(
	importedPackage ast.ImportedPackageLike,
	index_ uint,
	count_ uint,
) {
	var packageName = importedPackage.GetName()
	if utf.RuneCountInString(packageName) != 3 {
		var message = fmt.Sprintf(
			"An imported package name must be exactly three characters long: %s",
			packageName,
		)
		panic(message)
	}
}

func (v *validator_) PreprocessInstanceSection(
	instanceSection ast.InstanceSectionLike,
	index_ uint,
	count_ uint,
) {
	// We can "cheat" here because we know the code generator uses lists.
	var instanceDeclarations = instanceSection.GetInstanceDeclarations().(com.ListLike[ast.InstanceDeclarationLike])
	instanceDeclarations.SortValuesWithRanker(
		func(
			first ast.InstanceDeclarationLike,
			second ast.InstanceDeclarationLike,
		) com.Rank {
			var firstName = sts.TrimSuffix(
				first.GetDeclaration().GetName(),
				"Like",
			)
			var secondName = sts.TrimSuffix(
				second.GetDeclaration().GetName(),
				"Like",
			)
			switch {
			case firstName < secondName:
				return com.LesserRank
			case firstName > secondName:
				return com.GreaterRank
			default:
				return com.EqualRank
			}
		},
	)
}

// PROTECTED INTERFACE

// Private Methods

func (v *validator_) validateToken(
	tokenValue string,
	tokenType TokenType,
) {
	var scannerClass = ScannerClass()
	if !scannerClass.MatchesType(tokenValue, tokenType) {
		var message = fmt.Sprintf(
			"The following token value is not of type %v: %v",
			scannerClass.FormatType(tokenType),
			tokenValue,
		)
		panic(message)
	}
}

// Instance Structure

type validator_ struct {
	// Declare the instance attributes.

	// Declare the inherited aspects.
	Methodical
}

// Class Structure

type validatorClass_ struct {
	// Declare the class constants.
}

// Class Reference

func validatorClass() *validatorClass_ {
	return validatorClassReference_
}

var validatorClassReference_ = &validatorClass_{
	// Initialize the class constants.
}
