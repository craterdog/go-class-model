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
Package "module" defines type aliases for the commonly used types defined in the
packages contained in this module.  It also provides a universal constructor for
the commonly used classes that are exported by this module.  Each constructor
delegates the actual construction process to its corresponding concrete class
defined in a package contained within this module.

For detailed documentation on this entire module refer to the wiki:
  - https://github.com/craterdog/go-class-model/wiki
*/
package module

import (
	fmt "fmt"
	ast "github.com/craterdog/go-class-model/v5/ast"
	gra "github.com/craterdog/go-class-model/v5/grammar"
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// TYPE ALIASES

// AST

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
	HeaderLike                = ast.HeaderLike
	ImportsLike               = ast.ImportsLike
	InstanceDeclarationLike   = ast.InstanceDeclarationLike
	InstanceMethodsLike       = ast.InstanceMethodsLike
	InstanceSectionLike       = ast.InstanceSectionLike
	InterfaceDeclarationsLike = ast.InterfaceDeclarationsLike
	MapLike                   = ast.MapLike
	MethodLike                = ast.MethodLike
	ModelLike                 = ast.ModelLike
	ModuleLike                = ast.ModuleLike
	ModuleDeclarationLike     = ast.ModuleDeclarationLike
	NoneLike                  = ast.NoneLike
	NoticeLike                = ast.NoticeLike
	ParameterLike             = ast.ParameterLike
	ParameterizedLike         = ast.ParameterizedLike
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
	FormatterLike = gra.FormatterLike
	ParserLike    = gra.ParserLike
	ValidatorLike = gra.ValidatorLike
	VisitorLike   = gra.VisitorLike
)

// Aspects

type (
	Methodical = gra.Methodical
)

// UNIVERSAL CONSTRUCTORS

func Formatter(args ...any) FormatterLike {
	if len(args) > 0 {
		panic("The \"formatter\" constructor does not take any arguments.")
	}
	var formatter = gra.Formatter().Make()
	return formatter
}

func Parser(args ...any) ParserLike {
	if len(args) > 0 {
		panic("The \"parser\" constructor does not take any arguments.")
	}
	var parser = gra.Parser().Make()
	return parser
}

func Validator(args ...any) ValidatorLike {
	if len(args) > 0 {
		panic("The \"validator\" constructor does not take any arguments.")
	}
	var validator = gra.Validator().Make()
	return validator
}

func Visitor(args ...any) VisitorLike {
	// Initialize the possible arguments.
	var processor Methodical

	// Process the actual arguments.
	for _, arg := range args {
		switch actual := arg.(type) {
		case Methodical:
			processor = actual
		default:
			if uti.IsDefined(arg) {
				var message = fmt.Sprintf(
					"An unknown argument type was passed into the \"visitor\" constructor: %T\n",
					actual,
				)
				panic(message)
			}
		}
	}

	// Call the constructor.
	var visitor = gra.Visitor().Make(
		processor,
	)
	return visitor
}

// GLOBAL FUNCTIONS

func FormatModel(model ModelLike) string {
	var formatter = Formatter()
	var source = formatter.FormatModel(model)
	return source
}

func ParseSource(source string) ModelLike {
	var parser = Parser()
	var model = parser.ParseSource(source)
	return model
}

func ValidateModel(model ModelLike) {
	var validator = Validator()
	validator.ValidateModel(model)
}
