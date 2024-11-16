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
│              THIS "Module.go" FILE WAS AUTOMATICALLY GENERATED.              │
│                    ANY UPDATES TO IT WILL BE OVERWRITTEN.                    │
└──────────────────────────────────────────────────────────────────────────────┘

Package "module" defines type aliases for the commonly used types defined in the
packages contained in this module.  It also provides a universal constructor for
each commonly used class that is exported by the module.  Each constructor
delegates the actual construction process to its corresponding concrete class
defined in the corresponding package contained within this module.

For detailed documentation on this entire module refer to the wiki:
  - github.com/craterdog/go-class-model/wiki
*/
package module

import (
	fmt "fmt"
	ast "github.com/craterdog/go-class-model/v5/ast"
	gra "github.com/craterdog/go-class-model/v5/grammar"
	abs "github.com/craterdog/go-collection-framework/v4/collection"
)

// TYPE ALIASES

// Ast

type (
	AbstractionLike           = ast.AbstractionLike
	AdditionalArgumentLike    = ast.AdditionalArgumentLike
	AdditionalConstraintLike  = ast.AdditionalConstraintLike
	AdditionalValueLike       = ast.AdditionalValueLike
	ArgumentLike              = ast.ArgumentLike
	ArgumentsLike             = ast.ArgumentsLike
	ArrayLike                 = ast.ArrayLike
	AspectDeclarationLike     = ast.AspectDeclarationLike
	AspectInterfaceLike       = ast.AspectInterfaceLike
	AspectMethodLike          = ast.AspectMethodLike
	AspectSectionLike         = ast.AspectSectionLike
	AspectSubsectionLike      = ast.AspectSubsectionLike
	AttributeMethodLike       = ast.AttributeMethodLike
	AttributeSubsectionLike   = ast.AttributeSubsectionLike
	ChannelLike               = ast.ChannelLike
	ClassDeclarationLike      = ast.ClassDeclarationLike
	ClassMethodsLike          = ast.ClassMethodsLike
	ClassSectionLike          = ast.ClassSectionLike
	ConstantMethodLike        = ast.ConstantMethodLike
	ConstantSubsectionLike    = ast.ConstantSubsectionLike
	ConstraintLike            = ast.ConstraintLike
	ConstraintsLike           = ast.ConstraintsLike
	ConstructorMethodLike     = ast.ConstructorMethodLike
	ConstructorSubsectionLike = ast.ConstructorSubsectionLike
	DeclarationLike           = ast.DeclarationLike
	EnumerationLike           = ast.EnumerationLike
	FunctionMethodLike        = ast.FunctionMethodLike
	FunctionSubsectionLike    = ast.FunctionSubsectionLike
	FunctionalDeclarationLike = ast.FunctionalDeclarationLike
	FunctionalSectionLike     = ast.FunctionalSectionLike
	GetterMethodLike          = ast.GetterMethodLike
	ImportedPackageLike       = ast.ImportedPackageLike
	InstanceDeclarationLike   = ast.InstanceDeclarationLike
	InstanceMethodsLike       = ast.InstanceMethodsLike
	InstanceSectionLike       = ast.InstanceSectionLike
	InterfaceDeclarationsLike = ast.InterfaceDeclarationsLike
	LegalNoticeLike           = ast.LegalNoticeLike
	MapLike                   = ast.MapLike
	MethodLike                = ast.MethodLike
	ModelLike                 = ast.ModelLike
	ModuleDeclarationLike     = ast.ModuleDeclarationLike
	ModuleHeaderLike          = ast.ModuleHeaderLike
	ModuleImportsLike         = ast.ModuleImportsLike
	MultivalueLike            = ast.MultivalueLike
	NoneLike                  = ast.NoneLike
	ParameterLike             = ast.ParameterLike
	PrefixLike                = ast.PrefixLike
	PrimaryMethodLike         = ast.PrimaryMethodLike
	PrimarySubsectionLike     = ast.PrimarySubsectionLike
	PrimitiveDeclarationsLike = ast.PrimitiveDeclarationsLike
	ResultLike                = ast.ResultLike
	SetterMethodLike          = ast.SetterMethodLike
	SuffixLike                = ast.SuffixLike
	TypeDeclarationLike       = ast.TypeDeclarationLike
	TypeSectionLike           = ast.TypeSectionLike
	ValueLike                 = ast.ValueLike
)

// Grammar

type (
	TokenType     = gra.TokenType
	FormatterLike = gra.FormatterLike
	ParserLike    = gra.ParserLike
	ProcessorLike = gra.ProcessorLike
	ScannerLike   = gra.ScannerLike
	TokenLike     = gra.TokenLike
	ValidatorLike = gra.ValidatorLike
	VisitorLike   = gra.VisitorLike
	Methodical    = gra.Methodical
)

// UNIVERSAL CONSTRUCTORS

// Ast

func Abstraction(arguments ...any) AbstractionLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case PrefixLike:
			argumentTypes += "PrefixLike, "
		case string:
			argumentTypes += "string, "
		case SuffixLike:
			argumentTypes += "SuffixLike, "
		case ArgumentsLike:
			argumentTypes += "ArgumentsLike, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Abstraction constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var instance_ AbstractionLike
	switch argumentTypes {
	case "PrefixLike, string, SuffixLike, ArgumentsLike":
		var optionalPrefix = arguments[0].(PrefixLike)
		var name = arguments[1].(string)
		var optionalSuffix = arguments[2].(SuffixLike)
		var optionalArguments = arguments[3].(ArgumentsLike)
		instance_ = ast.Abstraction().Make(
			optionalPrefix,
			name,
			optionalSuffix,
			optionalArguments,
		)
	default:
		var message = fmt.Sprintf(
			"No Abstraction constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func AdditionalArgument(arguments ...any) AdditionalArgumentLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case ArgumentLike:
			argumentTypes += "ArgumentLike, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the AdditionalArgument constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var instance_ AdditionalArgumentLike
	switch argumentTypes {
	case "ArgumentLike":
		var argument = arguments[0].(ArgumentLike)
		instance_ = ast.AdditionalArgument().Make(
			argument,
		)
	default:
		var message = fmt.Sprintf(
			"No AdditionalArgument constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func AdditionalConstraint(arguments ...any) AdditionalConstraintLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case ConstraintLike:
			argumentTypes += "ConstraintLike, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the AdditionalConstraint constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var instance_ AdditionalConstraintLike
	switch argumentTypes {
	case "ConstraintLike":
		var constraint = arguments[0].(ConstraintLike)
		instance_ = ast.AdditionalConstraint().Make(
			constraint,
		)
	default:
		var message = fmt.Sprintf(
			"No AdditionalConstraint constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func AdditionalValue(arguments ...any) AdditionalValueLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			argumentTypes += "string, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the AdditionalValue constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var instance_ AdditionalValueLike
	switch argumentTypes {
	case "string":
		var name = arguments[0].(string)
		instance_ = ast.AdditionalValue().Make(
			name,
		)
	default:
		var message = fmt.Sprintf(
			"No AdditionalValue constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func Argument(arguments ...any) ArgumentLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case AbstractionLike:
			argumentTypes += "AbstractionLike, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Argument constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var instance_ ArgumentLike
	switch argumentTypes {
	case "AbstractionLike":
		var abstraction = arguments[0].(AbstractionLike)
		instance_ = ast.Argument().Make(
			abstraction,
		)
	default:
		var message = fmt.Sprintf(
			"No Argument constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func Arguments(arguments ...any) ArgumentsLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case ArgumentLike:
			argumentTypes += "ArgumentLike, "
		case abs.Sequential[AdditionalArgumentLike]:
			argumentTypes += "abs.Sequential[AdditionalArgumentLike], "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Arguments constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var instance_ ArgumentsLike
	switch argumentTypes {
	case "ArgumentLike, abs.Sequential[AdditionalArgumentLike]":
		var argument = arguments[0].(ArgumentLike)
		var additionalArguments = arguments[1].(abs.Sequential[AdditionalArgumentLike])
		instance_ = ast.Arguments().Make(
			argument,
			additionalArguments,
		)
	default:
		var message = fmt.Sprintf(
			"No Arguments constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func Array(arguments ...any) ArrayLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Array constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var instance_ ArrayLike
	switch argumentTypes {
	case "":
		instance_ = ast.Array().Make()
	default:
		var message = fmt.Sprintf(
			"No Array constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func AspectDeclaration(arguments ...any) AspectDeclarationLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case DeclarationLike:
			argumentTypes += "DeclarationLike, "
		case abs.Sequential[AspectMethodLike]:
			argumentTypes += "abs.Sequential[AspectMethodLike], "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the AspectDeclaration constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var instance_ AspectDeclarationLike
	switch argumentTypes {
	case "DeclarationLike, abs.Sequential[AspectMethodLike]":
		var declaration = arguments[0].(DeclarationLike)
		var aspectMethods = arguments[1].(abs.Sequential[AspectMethodLike])
		instance_ = ast.AspectDeclaration().Make(
			declaration,
			aspectMethods,
		)
	default:
		var message = fmt.Sprintf(
			"No AspectDeclaration constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func AspectInterface(arguments ...any) AspectInterfaceLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case AbstractionLike:
			argumentTypes += "AbstractionLike, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the AspectInterface constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var instance_ AspectInterfaceLike
	switch argumentTypes {
	case "AbstractionLike":
		var abstraction = arguments[0].(AbstractionLike)
		instance_ = ast.AspectInterface().Make(
			abstraction,
		)
	default:
		var message = fmt.Sprintf(
			"No AspectInterface constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func AspectMethod(arguments ...any) AspectMethodLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case MethodLike:
			argumentTypes += "MethodLike, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the AspectMethod constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var instance_ AspectMethodLike
	switch argumentTypes {
	case "MethodLike":
		var method = arguments[0].(MethodLike)
		instance_ = ast.AspectMethod().Make(
			method,
		)
	default:
		var message = fmt.Sprintf(
			"No AspectMethod constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func AspectSection(arguments ...any) AspectSectionLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case abs.Sequential[AspectDeclarationLike]:
			argumentTypes += "abs.Sequential[AspectDeclarationLike], "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the AspectSection constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var instance_ AspectSectionLike
	switch argumentTypes {
	case "abs.Sequential[AspectDeclarationLike]":
		var aspectDeclarations = arguments[0].(abs.Sequential[AspectDeclarationLike])
		instance_ = ast.AspectSection().Make(
			aspectDeclarations,
		)
	default:
		var message = fmt.Sprintf(
			"No AspectSection constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func AspectSubsection(arguments ...any) AspectSubsectionLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case abs.Sequential[AspectInterfaceLike]:
			argumentTypes += "abs.Sequential[AspectInterfaceLike], "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the AspectSubsection constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var instance_ AspectSubsectionLike
	switch argumentTypes {
	case "abs.Sequential[AspectInterfaceLike]":
		var aspectInterfaces = arguments[0].(abs.Sequential[AspectInterfaceLike])
		instance_ = ast.AspectSubsection().Make(
			aspectInterfaces,
		)
	default:
		var message = fmt.Sprintf(
			"No AspectSubsection constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func AttributeMethod(arguments ...any) AttributeMethodLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case any:
			argumentTypes += "any, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the AttributeMethod constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var instance_ AttributeMethodLike
	switch argumentTypes {
	case "any":
		var any_ = arguments[0]
		instance_ = ast.AttributeMethod().Make(
			any_,
		)
	default:
		var message = fmt.Sprintf(
			"No AttributeMethod constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func AttributeSubsection(arguments ...any) AttributeSubsectionLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case abs.Sequential[AttributeMethodLike]:
			argumentTypes += "abs.Sequential[AttributeMethodLike], "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the AttributeSubsection constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var instance_ AttributeSubsectionLike
	switch argumentTypes {
	case "abs.Sequential[AttributeMethodLike]":
		var attributeMethods = arguments[0].(abs.Sequential[AttributeMethodLike])
		instance_ = ast.AttributeSubsection().Make(
			attributeMethods,
		)
	default:
		var message = fmt.Sprintf(
			"No AttributeSubsection constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func Channel(arguments ...any) ChannelLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Channel constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var instance_ ChannelLike
	switch argumentTypes {
	case "":
		instance_ = ast.Channel().Make()
	default:
		var message = fmt.Sprintf(
			"No Channel constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func ClassDeclaration(arguments ...any) ClassDeclarationLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case DeclarationLike:
			argumentTypes += "DeclarationLike, "
		case ClassMethodsLike:
			argumentTypes += "ClassMethodsLike, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the ClassDeclaration constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var instance_ ClassDeclarationLike
	switch argumentTypes {
	case "DeclarationLike, ClassMethodsLike":
		var declaration = arguments[0].(DeclarationLike)
		var classMethods = arguments[1].(ClassMethodsLike)
		instance_ = ast.ClassDeclaration().Make(
			declaration,
			classMethods,
		)
	default:
		var message = fmt.Sprintf(
			"No ClassDeclaration constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func ClassMethods(arguments ...any) ClassMethodsLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case ConstructorSubsectionLike:
			argumentTypes += "ConstructorSubsectionLike, "
		case ConstantSubsectionLike:
			argumentTypes += "ConstantSubsectionLike, "
		case FunctionSubsectionLike:
			argumentTypes += "FunctionSubsectionLike, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the ClassMethods constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var instance_ ClassMethodsLike
	switch argumentTypes {
	case "ConstructorSubsectionLike, ConstantSubsectionLike, FunctionSubsectionLike":
		var constructorSubsection = arguments[0].(ConstructorSubsectionLike)
		var optionalConstantSubsection = arguments[1].(ConstantSubsectionLike)
		var optionalFunctionSubsection = arguments[2].(FunctionSubsectionLike)
		instance_ = ast.ClassMethods().Make(
			constructorSubsection,
			optionalConstantSubsection,
			optionalFunctionSubsection,
		)
	default:
		var message = fmt.Sprintf(
			"No ClassMethods constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func ClassSection(arguments ...any) ClassSectionLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case abs.Sequential[ClassDeclarationLike]:
			argumentTypes += "abs.Sequential[ClassDeclarationLike], "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the ClassSection constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var instance_ ClassSectionLike
	switch argumentTypes {
	case "abs.Sequential[ClassDeclarationLike]":
		var classDeclarations = arguments[0].(abs.Sequential[ClassDeclarationLike])
		instance_ = ast.ClassSection().Make(
			classDeclarations,
		)
	default:
		var message = fmt.Sprintf(
			"No ClassSection constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func ConstantMethod(arguments ...any) ConstantMethodLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			argumentTypes += "string, "
		case AbstractionLike:
			argumentTypes += "AbstractionLike, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the ConstantMethod constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var instance_ ConstantMethodLike
	switch argumentTypes {
	case "string, AbstractionLike":
		var name = arguments[0].(string)
		var abstraction = arguments[1].(AbstractionLike)
		instance_ = ast.ConstantMethod().Make(
			name,
			abstraction,
		)
	default:
		var message = fmt.Sprintf(
			"No ConstantMethod constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func ConstantSubsection(arguments ...any) ConstantSubsectionLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case abs.Sequential[ConstantMethodLike]:
			argumentTypes += "abs.Sequential[ConstantMethodLike], "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the ConstantSubsection constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var instance_ ConstantSubsectionLike
	switch argumentTypes {
	case "abs.Sequential[ConstantMethodLike]":
		var constantMethods = arguments[0].(abs.Sequential[ConstantMethodLike])
		instance_ = ast.ConstantSubsection().Make(
			constantMethods,
		)
	default:
		var message = fmt.Sprintf(
			"No ConstantSubsection constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func Constraint(arguments ...any) ConstraintLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			argumentTypes += "string, "
		case AbstractionLike:
			argumentTypes += "AbstractionLike, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Constraint constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var instance_ ConstraintLike
	switch argumentTypes {
	case "string, AbstractionLike":
		var name = arguments[0].(string)
		var abstraction = arguments[1].(AbstractionLike)
		instance_ = ast.Constraint().Make(
			name,
			abstraction,
		)
	default:
		var message = fmt.Sprintf(
			"No Constraint constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func Constraints(arguments ...any) ConstraintsLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case ConstraintLike:
			argumentTypes += "ConstraintLike, "
		case abs.Sequential[AdditionalConstraintLike]:
			argumentTypes += "abs.Sequential[AdditionalConstraintLike], "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Constraints constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var instance_ ConstraintsLike
	switch argumentTypes {
	case "ConstraintLike, abs.Sequential[AdditionalConstraintLike]":
		var constraint = arguments[0].(ConstraintLike)
		var additionalConstraints = arguments[1].(abs.Sequential[AdditionalConstraintLike])
		instance_ = ast.Constraints().Make(
			constraint,
			additionalConstraints,
		)
	default:
		var message = fmt.Sprintf(
			"No Constraints constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func ConstructorMethod(arguments ...any) ConstructorMethodLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			argumentTypes += "string, "
		case abs.Sequential[ParameterLike]:
			argumentTypes += "abs.Sequential[ParameterLike], "
		case AbstractionLike:
			argumentTypes += "AbstractionLike, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the ConstructorMethod constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var instance_ ConstructorMethodLike
	switch argumentTypes {
	case "string, abs.Sequential[ParameterLike], AbstractionLike":
		var name = arguments[0].(string)
		var parameters = arguments[1].(abs.Sequential[ParameterLike])
		var abstraction = arguments[2].(AbstractionLike)
		instance_ = ast.ConstructorMethod().Make(
			name,
			parameters,
			abstraction,
		)
	default:
		var message = fmt.Sprintf(
			"No ConstructorMethod constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func ConstructorSubsection(arguments ...any) ConstructorSubsectionLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case abs.Sequential[ConstructorMethodLike]:
			argumentTypes += "abs.Sequential[ConstructorMethodLike], "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the ConstructorSubsection constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var instance_ ConstructorSubsectionLike
	switch argumentTypes {
	case "abs.Sequential[ConstructorMethodLike]":
		var constructorMethods = arguments[0].(abs.Sequential[ConstructorMethodLike])
		instance_ = ast.ConstructorSubsection().Make(
			constructorMethods,
		)
	default:
		var message = fmt.Sprintf(
			"No ConstructorSubsection constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func Declaration(arguments ...any) DeclarationLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			argumentTypes += "string, "
		case ConstraintsLike:
			argumentTypes += "ConstraintsLike, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Declaration constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var instance_ DeclarationLike
	switch argumentTypes {
	case "string, string, ConstraintsLike":
		var comment = arguments[0].(string)
		var name = arguments[1].(string)
		var optionalConstraints = arguments[2].(ConstraintsLike)
		instance_ = ast.Declaration().Make(
			comment,
			name,
			optionalConstraints,
		)
	default:
		var message = fmt.Sprintf(
			"No Declaration constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func Enumeration(arguments ...any) EnumerationLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case ValueLike:
			argumentTypes += "ValueLike, "
		case abs.Sequential[AdditionalValueLike]:
			argumentTypes += "abs.Sequential[AdditionalValueLike], "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Enumeration constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var instance_ EnumerationLike
	switch argumentTypes {
	case "ValueLike, abs.Sequential[AdditionalValueLike]":
		var value = arguments[0].(ValueLike)
		var additionalValues = arguments[1].(abs.Sequential[AdditionalValueLike])
		instance_ = ast.Enumeration().Make(
			value,
			additionalValues,
		)
	default:
		var message = fmt.Sprintf(
			"No Enumeration constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func FunctionMethod(arguments ...any) FunctionMethodLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			argumentTypes += "string, "
		case abs.Sequential[ParameterLike]:
			argumentTypes += "abs.Sequential[ParameterLike], "
		case ResultLike:
			argumentTypes += "ResultLike, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the FunctionMethod constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var instance_ FunctionMethodLike
	switch argumentTypes {
	case "string, abs.Sequential[ParameterLike], ResultLike":
		var name = arguments[0].(string)
		var parameters = arguments[1].(abs.Sequential[ParameterLike])
		var result = arguments[2].(ResultLike)
		instance_ = ast.FunctionMethod().Make(
			name,
			parameters,
			result,
		)
	default:
		var message = fmt.Sprintf(
			"No FunctionMethod constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func FunctionSubsection(arguments ...any) FunctionSubsectionLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case abs.Sequential[FunctionMethodLike]:
			argumentTypes += "abs.Sequential[FunctionMethodLike], "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the FunctionSubsection constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var instance_ FunctionSubsectionLike
	switch argumentTypes {
	case "abs.Sequential[FunctionMethodLike]":
		var functionMethods = arguments[0].(abs.Sequential[FunctionMethodLike])
		instance_ = ast.FunctionSubsection().Make(
			functionMethods,
		)
	default:
		var message = fmt.Sprintf(
			"No FunctionSubsection constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func FunctionalDeclaration(arguments ...any) FunctionalDeclarationLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case DeclarationLike:
			argumentTypes += "DeclarationLike, "
		case abs.Sequential[ParameterLike]:
			argumentTypes += "abs.Sequential[ParameterLike], "
		case ResultLike:
			argumentTypes += "ResultLike, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the FunctionalDeclaration constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var instance_ FunctionalDeclarationLike
	switch argumentTypes {
	case "DeclarationLike, abs.Sequential[ParameterLike], ResultLike":
		var declaration = arguments[0].(DeclarationLike)
		var parameters = arguments[1].(abs.Sequential[ParameterLike])
		var result = arguments[2].(ResultLike)
		instance_ = ast.FunctionalDeclaration().Make(
			declaration,
			parameters,
			result,
		)
	default:
		var message = fmt.Sprintf(
			"No FunctionalDeclaration constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func FunctionalSection(arguments ...any) FunctionalSectionLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case abs.Sequential[FunctionalDeclarationLike]:
			argumentTypes += "abs.Sequential[FunctionalDeclarationLike], "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the FunctionalSection constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var instance_ FunctionalSectionLike
	switch argumentTypes {
	case "abs.Sequential[FunctionalDeclarationLike]":
		var functionalDeclarations = arguments[0].(abs.Sequential[FunctionalDeclarationLike])
		instance_ = ast.FunctionalSection().Make(
			functionalDeclarations,
		)
	default:
		var message = fmt.Sprintf(
			"No FunctionalSection constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func GetterMethod(arguments ...any) GetterMethodLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			argumentTypes += "string, "
		case AbstractionLike:
			argumentTypes += "AbstractionLike, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the GetterMethod constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var instance_ GetterMethodLike
	switch argumentTypes {
	case "string, AbstractionLike":
		var name = arguments[0].(string)
		var abstraction = arguments[1].(AbstractionLike)
		instance_ = ast.GetterMethod().Make(
			name,
			abstraction,
		)
	default:
		var message = fmt.Sprintf(
			"No GetterMethod constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func ImportedPackage(arguments ...any) ImportedPackageLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			argumentTypes += "string, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the ImportedPackage constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var instance_ ImportedPackageLike
	switch argumentTypes {
	case "string, string":
		var name = arguments[0].(string)
		var path = arguments[1].(string)
		instance_ = ast.ImportedPackage().Make(
			name,
			path,
		)
	default:
		var message = fmt.Sprintf(
			"No ImportedPackage constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func InstanceDeclaration(arguments ...any) InstanceDeclarationLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case DeclarationLike:
			argumentTypes += "DeclarationLike, "
		case InstanceMethodsLike:
			argumentTypes += "InstanceMethodsLike, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the InstanceDeclaration constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var instance_ InstanceDeclarationLike
	switch argumentTypes {
	case "DeclarationLike, InstanceMethodsLike":
		var declaration = arguments[0].(DeclarationLike)
		var instanceMethods = arguments[1].(InstanceMethodsLike)
		instance_ = ast.InstanceDeclaration().Make(
			declaration,
			instanceMethods,
		)
	default:
		var message = fmt.Sprintf(
			"No InstanceDeclaration constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func InstanceMethods(arguments ...any) InstanceMethodsLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case PrimarySubsectionLike:
			argumentTypes += "PrimarySubsectionLike, "
		case AttributeSubsectionLike:
			argumentTypes += "AttributeSubsectionLike, "
		case AspectSubsectionLike:
			argumentTypes += "AspectSubsectionLike, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the InstanceMethods constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var instance_ InstanceMethodsLike
	switch argumentTypes {
	case "PrimarySubsectionLike, AttributeSubsectionLike, AspectSubsectionLike":
		var primarySubsection = arguments[0].(PrimarySubsectionLike)
		var optionalAttributeSubsection = arguments[1].(AttributeSubsectionLike)
		var optionalAspectSubsection = arguments[2].(AspectSubsectionLike)
		instance_ = ast.InstanceMethods().Make(
			primarySubsection,
			optionalAttributeSubsection,
			optionalAspectSubsection,
		)
	default:
		var message = fmt.Sprintf(
			"No InstanceMethods constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func InstanceSection(arguments ...any) InstanceSectionLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case abs.Sequential[InstanceDeclarationLike]:
			argumentTypes += "abs.Sequential[InstanceDeclarationLike], "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the InstanceSection constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var instance_ InstanceSectionLike
	switch argumentTypes {
	case "abs.Sequential[InstanceDeclarationLike]":
		var instanceDeclarations = arguments[0].(abs.Sequential[InstanceDeclarationLike])
		instance_ = ast.InstanceSection().Make(
			instanceDeclarations,
		)
	default:
		var message = fmt.Sprintf(
			"No InstanceSection constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func InterfaceDeclarations(arguments ...any) InterfaceDeclarationsLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case ClassSectionLike:
			argumentTypes += "ClassSectionLike, "
		case InstanceSectionLike:
			argumentTypes += "InstanceSectionLike, "
		case AspectSectionLike:
			argumentTypes += "AspectSectionLike, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the InterfaceDeclarations constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var instance_ InterfaceDeclarationsLike
	switch argumentTypes {
	case "ClassSectionLike, InstanceSectionLike, AspectSectionLike":
		var classSection = arguments[0].(ClassSectionLike)
		var instanceSection = arguments[1].(InstanceSectionLike)
		var aspectSection = arguments[2].(AspectSectionLike)
		instance_ = ast.InterfaceDeclarations().Make(
			classSection,
			instanceSection,
			aspectSection,
		)
	default:
		var message = fmt.Sprintf(
			"No InterfaceDeclarations constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func LegalNotice(arguments ...any) LegalNoticeLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			argumentTypes += "string, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the LegalNotice constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var instance_ LegalNoticeLike
	switch argumentTypes {
	case "string":
		var comment = arguments[0].(string)
		instance_ = ast.LegalNotice().Make(
			comment,
		)
	default:
		var message = fmt.Sprintf(
			"No LegalNotice constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func Map(arguments ...any) MapLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			argumentTypes += "string, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Map constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var instance_ MapLike
	switch argumentTypes {
	case "string":
		var name = arguments[0].(string)
		instance_ = ast.Map().Make(
			name,
		)
	default:
		var message = fmt.Sprintf(
			"No Map constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func Method(arguments ...any) MethodLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			argumentTypes += "string, "
		case abs.Sequential[ParameterLike]:
			argumentTypes += "abs.Sequential[ParameterLike], "
		case ResultLike:
			argumentTypes += "ResultLike, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Method constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var instance_ MethodLike
	switch argumentTypes {
	case "string, abs.Sequential[ParameterLike], ResultLike":
		var name = arguments[0].(string)
		var parameters = arguments[1].(abs.Sequential[ParameterLike])
		var optionalResult = arguments[2].(ResultLike)
		instance_ = ast.Method().Make(
			name,
			parameters,
			optionalResult,
		)
	default:
		var message = fmt.Sprintf(
			"No Method constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func Model(arguments ...any) ModelLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case ModuleDeclarationLike:
			argumentTypes += "ModuleDeclarationLike, "
		case PrimitiveDeclarationsLike:
			argumentTypes += "PrimitiveDeclarationsLike, "
		case InterfaceDeclarationsLike:
			argumentTypes += "InterfaceDeclarationsLike, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Model constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var instance_ ModelLike
	switch argumentTypes {
	case "ModuleDeclarationLike, PrimitiveDeclarationsLike, InterfaceDeclarationsLike":
		var moduleDeclaration = arguments[0].(ModuleDeclarationLike)
		var primitiveDeclarations = arguments[1].(PrimitiveDeclarationsLike)
		var interfaceDeclarations = arguments[2].(InterfaceDeclarationsLike)
		instance_ = ast.Model().Make(
			moduleDeclaration,
			primitiveDeclarations,
			interfaceDeclarations,
		)
	default:
		var message = fmt.Sprintf(
			"No Model constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func ModuleDeclaration(arguments ...any) ModuleDeclarationLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case LegalNoticeLike:
			argumentTypes += "LegalNoticeLike, "
		case ModuleHeaderLike:
			argumentTypes += "ModuleHeaderLike, "
		case ModuleImportsLike:
			argumentTypes += "ModuleImportsLike, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the ModuleDeclaration constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var instance_ ModuleDeclarationLike
	switch argumentTypes {
	case "LegalNoticeLike, ModuleHeaderLike, ModuleImportsLike":
		var legalNotice = arguments[0].(LegalNoticeLike)
		var moduleHeader = arguments[1].(ModuleHeaderLike)
		var moduleImports = arguments[2].(ModuleImportsLike)
		instance_ = ast.ModuleDeclaration().Make(
			legalNotice,
			moduleHeader,
			moduleImports,
		)
	default:
		var message = fmt.Sprintf(
			"No ModuleDeclaration constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func ModuleHeader(arguments ...any) ModuleHeaderLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			argumentTypes += "string, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the ModuleHeader constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var instance_ ModuleHeaderLike
	switch argumentTypes {
	case "string, string":
		var comment = arguments[0].(string)
		var name = arguments[1].(string)
		instance_ = ast.ModuleHeader().Make(
			comment,
			name,
		)
	default:
		var message = fmt.Sprintf(
			"No ModuleHeader constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func ModuleImports(arguments ...any) ModuleImportsLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case abs.Sequential[ImportedPackageLike]:
			argumentTypes += "abs.Sequential[ImportedPackageLike], "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the ModuleImports constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var instance_ ModuleImportsLike
	switch argumentTypes {
	case "abs.Sequential[ImportedPackageLike]":
		var importedPackages = arguments[0].(abs.Sequential[ImportedPackageLike])
		instance_ = ast.ModuleImports().Make(
			importedPackages,
		)
	default:
		var message = fmt.Sprintf(
			"No ModuleImports constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func Multivalue(arguments ...any) MultivalueLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case abs.Sequential[ParameterLike]:
			argumentTypes += "abs.Sequential[ParameterLike], "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Multivalue constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var instance_ MultivalueLike
	switch argumentTypes {
	case "abs.Sequential[ParameterLike]":
		var parameters = arguments[0].(abs.Sequential[ParameterLike])
		instance_ = ast.Multivalue().Make(
			parameters,
		)
	default:
		var message = fmt.Sprintf(
			"No Multivalue constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func None(arguments ...any) NoneLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			argumentTypes += "string, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the None constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var instance_ NoneLike
	switch argumentTypes {
	case "string":
		var newline = arguments[0].(string)
		instance_ = ast.None().Make(
			newline,
		)
	default:
		var message = fmt.Sprintf(
			"No None constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func Parameter(arguments ...any) ParameterLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			argumentTypes += "string, "
		case AbstractionLike:
			argumentTypes += "AbstractionLike, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Parameter constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var instance_ ParameterLike
	switch argumentTypes {
	case "string, AbstractionLike":
		var name = arguments[0].(string)
		var abstraction = arguments[1].(AbstractionLike)
		instance_ = ast.Parameter().Make(
			name,
			abstraction,
		)
	default:
		var message = fmt.Sprintf(
			"No Parameter constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func Prefix(arguments ...any) PrefixLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case any:
			argumentTypes += "any, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Prefix constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var instance_ PrefixLike
	switch argumentTypes {
	case "any":
		var any_ = arguments[0]
		instance_ = ast.Prefix().Make(
			any_,
		)
	default:
		var message = fmt.Sprintf(
			"No Prefix constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func PrimaryMethod(arguments ...any) PrimaryMethodLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case MethodLike:
			argumentTypes += "MethodLike, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the PrimaryMethod constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var instance_ PrimaryMethodLike
	switch argumentTypes {
	case "MethodLike":
		var method = arguments[0].(MethodLike)
		instance_ = ast.PrimaryMethod().Make(
			method,
		)
	default:
		var message = fmt.Sprintf(
			"No PrimaryMethod constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func PrimarySubsection(arguments ...any) PrimarySubsectionLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case abs.Sequential[PrimaryMethodLike]:
			argumentTypes += "abs.Sequential[PrimaryMethodLike], "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the PrimarySubsection constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var instance_ PrimarySubsectionLike
	switch argumentTypes {
	case "abs.Sequential[PrimaryMethodLike]":
		var primaryMethods = arguments[0].(abs.Sequential[PrimaryMethodLike])
		instance_ = ast.PrimarySubsection().Make(
			primaryMethods,
		)
	default:
		var message = fmt.Sprintf(
			"No PrimarySubsection constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func PrimitiveDeclarations(arguments ...any) PrimitiveDeclarationsLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case TypeSectionLike:
			argumentTypes += "TypeSectionLike, "
		case FunctionalSectionLike:
			argumentTypes += "FunctionalSectionLike, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the PrimitiveDeclarations constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var instance_ PrimitiveDeclarationsLike
	switch argumentTypes {
	case "TypeSectionLike, FunctionalSectionLike":
		var typeSection = arguments[0].(TypeSectionLike)
		var functionalSection = arguments[1].(FunctionalSectionLike)
		instance_ = ast.PrimitiveDeclarations().Make(
			typeSection,
			functionalSection,
		)
	default:
		var message = fmt.Sprintf(
			"No PrimitiveDeclarations constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func Result(arguments ...any) ResultLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case any:
			argumentTypes += "any, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Result constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var instance_ ResultLike
	switch argumentTypes {
	case "any":
		var any_ = arguments[0]
		instance_ = ast.Result().Make(
			any_,
		)
	default:
		var message = fmt.Sprintf(
			"No Result constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func SetterMethod(arguments ...any) SetterMethodLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			argumentTypes += "string, "
		case ParameterLike:
			argumentTypes += "ParameterLike, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the SetterMethod constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var instance_ SetterMethodLike
	switch argumentTypes {
	case "string, ParameterLike":
		var name = arguments[0].(string)
		var parameter = arguments[1].(ParameterLike)
		instance_ = ast.SetterMethod().Make(
			name,
			parameter,
		)
	default:
		var message = fmt.Sprintf(
			"No SetterMethod constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func Suffix(arguments ...any) SuffixLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			argumentTypes += "string, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Suffix constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var instance_ SuffixLike
	switch argumentTypes {
	case "string":
		var name = arguments[0].(string)
		instance_ = ast.Suffix().Make(
			name,
		)
	default:
		var message = fmt.Sprintf(
			"No Suffix constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func TypeDeclaration(arguments ...any) TypeDeclarationLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case DeclarationLike:
			argumentTypes += "DeclarationLike, "
		case AbstractionLike:
			argumentTypes += "AbstractionLike, "
		case EnumerationLike:
			argumentTypes += "EnumerationLike, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the TypeDeclaration constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var instance_ TypeDeclarationLike
	switch argumentTypes {
	case "DeclarationLike, AbstractionLike, EnumerationLike":
		var declaration = arguments[0].(DeclarationLike)
		var abstraction = arguments[1].(AbstractionLike)
		var optionalEnumeration = arguments[2].(EnumerationLike)
		instance_ = ast.TypeDeclaration().Make(
			declaration,
			abstraction,
			optionalEnumeration,
		)
	default:
		var message = fmt.Sprintf(
			"No TypeDeclaration constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func TypeSection(arguments ...any) TypeSectionLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case abs.Sequential[TypeDeclarationLike]:
			argumentTypes += "abs.Sequential[TypeDeclarationLike], "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the TypeSection constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var instance_ TypeSectionLike
	switch argumentTypes {
	case "abs.Sequential[TypeDeclarationLike]":
		var typeDeclarations = arguments[0].(abs.Sequential[TypeDeclarationLike])
		instance_ = ast.TypeSection().Make(
			typeDeclarations,
		)
	default:
		var message = fmt.Sprintf(
			"No TypeSection constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func Value(arguments ...any) ValueLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			argumentTypes += "string, "
		case AbstractionLike:
			argumentTypes += "AbstractionLike, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Value constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var instance_ ValueLike
	switch argumentTypes {
	case "string, AbstractionLike":
		var name = arguments[0].(string)
		var abstraction = arguments[1].(AbstractionLike)
		instance_ = ast.Value().Make(
			name,
			abstraction,
		)
	default:
		var message = fmt.Sprintf(
			"No Value constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

// Grammar

func Formatter(arguments ...any) FormatterLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Formatter constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var instance_ FormatterLike
	switch argumentTypes {
	case "":
		instance_ = gra.Formatter().Make()
	default:
		var message = fmt.Sprintf(
			"No Formatter constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func Parser(arguments ...any) ParserLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Parser constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var instance_ ParserLike
	switch argumentTypes {
	case "":
		instance_ = gra.Parser().Make()
	default:
		var message = fmt.Sprintf(
			"No Parser constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func Processor(arguments ...any) ProcessorLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Processor constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var instance_ ProcessorLike
	switch argumentTypes {
	case "":
		instance_ = gra.Processor().Make()
	default:
		var message = fmt.Sprintf(
			"No Processor constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func Scanner(arguments ...any) ScannerLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			argumentTypes += "string, "
		case abs.QueueLike[TokenLike]:
			argumentTypes += "abs.QueueLike[TokenLike], "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Scanner constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var instance_ ScannerLike
	switch argumentTypes {
	case "string, abs.QueueLike[TokenLike]":
		var source = arguments[0].(string)
		var tokens = arguments[1].(abs.QueueLike[TokenLike])
		instance_ = gra.Scanner().Make(
			source,
			tokens,
		)
	default:
		var message = fmt.Sprintf(
			"No Scanner constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func Token(arguments ...any) TokenLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case uint:
			argumentTypes += "uint, "
		case TokenType:
			argumentTypes += "TokenType, "
		case string:
			argumentTypes += "string, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Token constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var instance_ TokenLike
	switch argumentTypes {
	case "uint, uint, TokenType, string":
		var line = arguments[0].(uint)
		var position = arguments[1].(uint)
		var type_ = arguments[2].(TokenType)
		var value = arguments[3].(string)
		instance_ = gra.Token().Make(
			line,
			position,
			type_,
			value,
		)
	default:
		var message = fmt.Sprintf(
			"No Token constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func Validator(arguments ...any) ValidatorLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Validator constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var instance_ ValidatorLike
	switch argumentTypes {
	case "":
		instance_ = gra.Validator().Make()
	default:
		var message = fmt.Sprintf(
			"No Validator constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func Visitor(arguments ...any) VisitorLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case Methodical:
			argumentTypes += "Methodical, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Visitor constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var instance_ VisitorLike
	switch argumentTypes {
	case "Methodical":
		var processor = arguments[0].(Methodical)
		instance_ = gra.Visitor().Make(
			processor,
		)
	default:
		var message = fmt.Sprintf(
			"No Visitor constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}
