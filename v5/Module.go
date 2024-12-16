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
│              This "Module.go" file was automatically generated.              │
│      Updates to any part of this file—other than the Module Description      │
│             and the Global Functions sections may be overwritten.            │
└──────────────────────────────────────────────────────────────────────────────┘

Package "module" declares type aliases for the commonly used types declared in
the packages contained in this module.  It also provides constructors for each
commonly used class that is exported by the module.  Each constructor delegates
the actual construction process to its corresponding concrete class declared in
the corresponding package contained within this module.

For detailed documentation on this entire module refer to the wiki:
  - https://github.com/craterdog/go-class-model/wiki
*/
package module

import (
	ast "github.com/craterdog/go-class-model/v5/ast"
	gra "github.com/craterdog/go-class-model/v5/grammar"
	abs "github.com/craterdog/go-collection-framework/v5/collection"
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
	MultivalueLike            = ast.MultivalueLike
	NoneLike                  = ast.NoneLike
	PackageDeclarationLike    = ast.PackageDeclarationLike
	PackageHeaderLike         = ast.PackageHeaderLike
	PackageImportsLike        = ast.PackageImportsLike
	ParameterLike             = ast.ParameterLike
	PrefixLike                = ast.PrefixLike
	PrimitiveDeclarationsLike = ast.PrimitiveDeclarationsLike
	PrincipalMethodLike       = ast.PrincipalMethodLike
	PrincipalSubsectionLike   = ast.PrincipalSubsectionLike
	ResultLike                = ast.ResultLike
	SetterMethodLike          = ast.SetterMethodLike
	SuffixLike                = ast.SuffixLike
	TypeDeclarationLike       = ast.TypeDeclarationLike
	TypeSectionLike           = ast.TypeSectionLike
	ValueLike                 = ast.ValueLike
)

// Grammar

type (
	TokenType = gra.TokenType
)

const (
	ErrorToken     = gra.ErrorToken
	CommentToken   = gra.CommentToken
	DelimiterToken = gra.DelimiterToken
	NameToken      = gra.NameToken
	NewlineToken   = gra.NewlineToken
	PathToken      = gra.PathToken
	SpaceToken     = gra.SpaceToken
)

type (
	FormatterLike = gra.FormatterLike
	ParserLike    = gra.ParserLike
	ProcessorLike = gra.ProcessorLike
	ScannerLike   = gra.ScannerLike
	TokenLike     = gra.TokenLike
	ValidatorLike = gra.ValidatorLike
	VisitorLike   = gra.VisitorLike
)

type (
	Methodical = gra.Methodical
)

// CLASS CONSTRUCTORS

// Ast/Abstraction

func Abstraction(
	optionalPrefix PrefixLike,
	name string,
	optionalSuffix SuffixLike,
	optionalArguments ArgumentsLike,
) ast.AbstractionLike {
	return ast.AbstractionClass().Abstraction(
		optionalPrefix,
		name,
		optionalSuffix,
		optionalArguments,
	)
}

// Ast/AdditionalArgument

func AdditionalArgument(
	argument ArgumentLike,
) ast.AdditionalArgumentLike {
	return ast.AdditionalArgumentClass().AdditionalArgument(
		argument,
	)
}

// Ast/AdditionalConstraint

func AdditionalConstraint(
	constraint ConstraintLike,
) ast.AdditionalConstraintLike {
	return ast.AdditionalConstraintClass().AdditionalConstraint(
		constraint,
	)
}

// Ast/AdditionalValue

func AdditionalValue(
	name string,
) ast.AdditionalValueLike {
	return ast.AdditionalValueClass().AdditionalValue(
		name,
	)
}

// Ast/Argument

func Argument(
	abstraction AbstractionLike,
) ast.ArgumentLike {
	return ast.ArgumentClass().Argument(
		abstraction,
	)
}

// Ast/Arguments

func Arguments(
	argument ArgumentLike,
	additionalArguments abs.Sequential[AdditionalArgumentLike],
) ast.ArgumentsLike {
	return ast.ArgumentsClass().Arguments(
		argument,
		additionalArguments,
	)
}

// Ast/Array

func Array() ast.ArrayLike {
	return ast.ArrayClass().Array()
}

// Ast/AspectDeclaration

func AspectDeclaration(
	declaration DeclarationLike,
	aspectMethods abs.Sequential[AspectMethodLike],
) ast.AspectDeclarationLike {
	return ast.AspectDeclarationClass().AspectDeclaration(
		declaration,
		aspectMethods,
	)
}

// Ast/AspectInterface

func AspectInterface(
	abstraction AbstractionLike,
) ast.AspectInterfaceLike {
	return ast.AspectInterfaceClass().AspectInterface(
		abstraction,
	)
}

// Ast/AspectMethod

func AspectMethod(
	method MethodLike,
) ast.AspectMethodLike {
	return ast.AspectMethodClass().AspectMethod(
		method,
	)
}

// Ast/AspectSection

func AspectSection(
	aspectDeclarations abs.Sequential[AspectDeclarationLike],
) ast.AspectSectionLike {
	return ast.AspectSectionClass().AspectSection(
		aspectDeclarations,
	)
}

// Ast/AspectSubsection

func AspectSubsection(
	aspectInterfaces abs.Sequential[AspectInterfaceLike],
) ast.AspectSubsectionLike {
	return ast.AspectSubsectionClass().AspectSubsection(
		aspectInterfaces,
	)
}

// Ast/AttributeMethod

func AttributeMethod(
	any_ any,
) ast.AttributeMethodLike {
	return ast.AttributeMethodClass().AttributeMethod(
		any_,
	)
}

// Ast/AttributeSubsection

func AttributeSubsection(
	attributeMethods abs.Sequential[AttributeMethodLike],
) ast.AttributeSubsectionLike {
	return ast.AttributeSubsectionClass().AttributeSubsection(
		attributeMethods,
	)
}

// Ast/Channel

func Channel() ast.ChannelLike {
	return ast.ChannelClass().Channel()
}

// Ast/ClassDeclaration

func ClassDeclaration(
	declaration DeclarationLike,
	classMethods ClassMethodsLike,
) ast.ClassDeclarationLike {
	return ast.ClassDeclarationClass().ClassDeclaration(
		declaration,
		classMethods,
	)
}

// Ast/ClassMethods

func ClassMethods(
	constructorSubsection ConstructorSubsectionLike,
	optionalConstantSubsection ConstantSubsectionLike,
	optionalFunctionSubsection FunctionSubsectionLike,
) ast.ClassMethodsLike {
	return ast.ClassMethodsClass().ClassMethods(
		constructorSubsection,
		optionalConstantSubsection,
		optionalFunctionSubsection,
	)
}

// Ast/ClassSection

func ClassSection(
	classDeclarations abs.Sequential[ClassDeclarationLike],
) ast.ClassSectionLike {
	return ast.ClassSectionClass().ClassSection(
		classDeclarations,
	)
}

// Ast/ConstantMethod

func ConstantMethod(
	name string,
	abstraction AbstractionLike,
) ast.ConstantMethodLike {
	return ast.ConstantMethodClass().ConstantMethod(
		name,
		abstraction,
	)
}

// Ast/ConstantSubsection

func ConstantSubsection(
	constantMethods abs.Sequential[ConstantMethodLike],
) ast.ConstantSubsectionLike {
	return ast.ConstantSubsectionClass().ConstantSubsection(
		constantMethods,
	)
}

// Ast/Constraint

func Constraint(
	name string,
	abstraction AbstractionLike,
) ast.ConstraintLike {
	return ast.ConstraintClass().Constraint(
		name,
		abstraction,
	)
}

// Ast/Constraints

func Constraints(
	constraint ConstraintLike,
	additionalConstraints abs.Sequential[AdditionalConstraintLike],
) ast.ConstraintsLike {
	return ast.ConstraintsClass().Constraints(
		constraint,
		additionalConstraints,
	)
}

// Ast/ConstructorMethod

func ConstructorMethod(
	name string,
	parameters abs.Sequential[ParameterLike],
	abstraction AbstractionLike,
) ast.ConstructorMethodLike {
	return ast.ConstructorMethodClass().ConstructorMethod(
		name,
		parameters,
		abstraction,
	)
}

// Ast/ConstructorSubsection

func ConstructorSubsection(
	constructorMethods abs.Sequential[ConstructorMethodLike],
) ast.ConstructorSubsectionLike {
	return ast.ConstructorSubsectionClass().ConstructorSubsection(
		constructorMethods,
	)
}

// Ast/Declaration

func Declaration(
	comment string,
	name string,
	optionalConstraints ConstraintsLike,
) ast.DeclarationLike {
	return ast.DeclarationClass().Declaration(
		comment,
		name,
		optionalConstraints,
	)
}

// Ast/Enumeration

func Enumeration(
	value ValueLike,
	additionalValues abs.Sequential[AdditionalValueLike],
) ast.EnumerationLike {
	return ast.EnumerationClass().Enumeration(
		value,
		additionalValues,
	)
}

// Ast/FunctionMethod

func FunctionMethod(
	name string,
	parameters abs.Sequential[ParameterLike],
	result ResultLike,
) ast.FunctionMethodLike {
	return ast.FunctionMethodClass().FunctionMethod(
		name,
		parameters,
		result,
	)
}

// Ast/FunctionSubsection

func FunctionSubsection(
	functionMethods abs.Sequential[FunctionMethodLike],
) ast.FunctionSubsectionLike {
	return ast.FunctionSubsectionClass().FunctionSubsection(
		functionMethods,
	)
}

// Ast/FunctionalDeclaration

func FunctionalDeclaration(
	declaration DeclarationLike,
	parameters abs.Sequential[ParameterLike],
	result ResultLike,
) ast.FunctionalDeclarationLike {
	return ast.FunctionalDeclarationClass().FunctionalDeclaration(
		declaration,
		parameters,
		result,
	)
}

// Ast/FunctionalSection

func FunctionalSection(
	functionalDeclarations abs.Sequential[FunctionalDeclarationLike],
) ast.FunctionalSectionLike {
	return ast.FunctionalSectionClass().FunctionalSection(
		functionalDeclarations,
	)
}

// Ast/GetterMethod

func GetterMethod(
	name string,
	abstraction AbstractionLike,
) ast.GetterMethodLike {
	return ast.GetterMethodClass().GetterMethod(
		name,
		abstraction,
	)
}

// Ast/ImportedPackage

func ImportedPackage(
	name string,
	path string,
) ast.ImportedPackageLike {
	return ast.ImportedPackageClass().ImportedPackage(
		name,
		path,
	)
}

// Ast/InstanceDeclaration

func InstanceDeclaration(
	declaration DeclarationLike,
	instanceMethods InstanceMethodsLike,
) ast.InstanceDeclarationLike {
	return ast.InstanceDeclarationClass().InstanceDeclaration(
		declaration,
		instanceMethods,
	)
}

// Ast/InstanceMethods

func InstanceMethods(
	principalSubsection PrincipalSubsectionLike,
	optionalAttributeSubsection AttributeSubsectionLike,
	optionalAspectSubsection AspectSubsectionLike,
) ast.InstanceMethodsLike {
	return ast.InstanceMethodsClass().InstanceMethods(
		principalSubsection,
		optionalAttributeSubsection,
		optionalAspectSubsection,
	)
}

// Ast/InstanceSection

func InstanceSection(
	instanceDeclarations abs.Sequential[InstanceDeclarationLike],
) ast.InstanceSectionLike {
	return ast.InstanceSectionClass().InstanceSection(
		instanceDeclarations,
	)
}

// Ast/InterfaceDeclarations

func InterfaceDeclarations(
	classSection ClassSectionLike,
	instanceSection InstanceSectionLike,
	aspectSection AspectSectionLike,
) ast.InterfaceDeclarationsLike {
	return ast.InterfaceDeclarationsClass().InterfaceDeclarations(
		classSection,
		instanceSection,
		aspectSection,
	)
}

// Ast/LegalNotice

func LegalNotice(
	comment string,
) ast.LegalNoticeLike {
	return ast.LegalNoticeClass().LegalNotice(
		comment,
	)
}

// Ast/Map

func Map(
	name string,
) ast.MapLike {
	return ast.MapClass().Map(
		name,
	)
}

// Ast/Method

func Method(
	name string,
	parameters abs.Sequential[ParameterLike],
	optionalResult ResultLike,
) ast.MethodLike {
	return ast.MethodClass().Method(
		name,
		parameters,
		optionalResult,
	)
}

// Ast/Model

func Model(
	packageDeclaration PackageDeclarationLike,
	primitiveDeclarations PrimitiveDeclarationsLike,
	interfaceDeclarations InterfaceDeclarationsLike,
) ast.ModelLike {
	return ast.ModelClass().Model(
		packageDeclaration,
		primitiveDeclarations,
		interfaceDeclarations,
	)
}

// Ast/Multivalue

func Multivalue(
	parameters abs.Sequential[ParameterLike],
) ast.MultivalueLike {
	return ast.MultivalueClass().Multivalue(
		parameters,
	)
}

// Ast/None

func None(
	newline string,
) ast.NoneLike {
	return ast.NoneClass().None(
		newline,
	)
}

// Ast/PackageDeclaration

func PackageDeclaration(
	legalNotice LegalNoticeLike,
	packageHeader PackageHeaderLike,
	packageImports PackageImportsLike,
) ast.PackageDeclarationLike {
	return ast.PackageDeclarationClass().PackageDeclaration(
		legalNotice,
		packageHeader,
		packageImports,
	)
}

// Ast/PackageHeader

func PackageHeader(
	comment string,
	name string,
) ast.PackageHeaderLike {
	return ast.PackageHeaderClass().PackageHeader(
		comment,
		name,
	)
}

// Ast/PackageImports

func PackageImports(
	importedPackages abs.Sequential[ImportedPackageLike],
) ast.PackageImportsLike {
	return ast.PackageImportsClass().PackageImports(
		importedPackages,
	)
}

// Ast/Parameter

func Parameter(
	name string,
	abstraction AbstractionLike,
) ast.ParameterLike {
	return ast.ParameterClass().Parameter(
		name,
		abstraction,
	)
}

// Ast/Prefix

func Prefix(
	any_ any,
) ast.PrefixLike {
	return ast.PrefixClass().Prefix(
		any_,
	)
}

// Ast/PrimitiveDeclarations

func PrimitiveDeclarations(
	typeSection TypeSectionLike,
	functionalSection FunctionalSectionLike,
) ast.PrimitiveDeclarationsLike {
	return ast.PrimitiveDeclarationsClass().PrimitiveDeclarations(
		typeSection,
		functionalSection,
	)
}

// Ast/PrincipalMethod

func PrincipalMethod(
	method MethodLike,
) ast.PrincipalMethodLike {
	return ast.PrincipalMethodClass().PrincipalMethod(
		method,
	)
}

// Ast/PrincipalSubsection

func PrincipalSubsection(
	principalMethods abs.Sequential[PrincipalMethodLike],
) ast.PrincipalSubsectionLike {
	return ast.PrincipalSubsectionClass().PrincipalSubsection(
		principalMethods,
	)
}

// Ast/Result

func Result(
	any_ any,
) ast.ResultLike {
	return ast.ResultClass().Result(
		any_,
	)
}

// Ast/SetterMethod

func SetterMethod(
	name string,
	parameter ParameterLike,
) ast.SetterMethodLike {
	return ast.SetterMethodClass().SetterMethod(
		name,
		parameter,
	)
}

// Ast/Suffix

func Suffix(
	name string,
) ast.SuffixLike {
	return ast.SuffixClass().Suffix(
		name,
	)
}

// Ast/TypeDeclaration

func TypeDeclaration(
	declaration DeclarationLike,
	abstraction AbstractionLike,
	optionalEnumeration EnumerationLike,
) ast.TypeDeclarationLike {
	return ast.TypeDeclarationClass().TypeDeclaration(
		declaration,
		abstraction,
		optionalEnumeration,
	)
}

// Ast/TypeSection

func TypeSection(
	typeDeclarations abs.Sequential[TypeDeclarationLike],
) ast.TypeSectionLike {
	return ast.TypeSectionClass().TypeSection(
		typeDeclarations,
	)
}

// Ast/Value

func Value(
	name string,
	abstraction AbstractionLike,
) ast.ValueLike {
	return ast.ValueClass().Value(
		name,
		abstraction,
	)
}

// Grammar/Formatter

func Formatter() gra.FormatterLike {
	return gra.FormatterClass().Formatter()
}

// Grammar/Parser

func Parser() gra.ParserLike {
	return gra.ParserClass().Parser()
}

// Grammar/Processor

func Processor() gra.ProcessorLike {
	return gra.ProcessorClass().Processor()
}

// Grammar/Scanner

func Scanner(
	source string,
	tokens abs.QueueLike[TokenLike],
) gra.ScannerLike {
	return gra.ScannerClass().Scanner(
		source,
		tokens,
	)
}

// Grammar/Token

func Token(
	line uint,
	position uint,
	type_ TokenType,
	value string,
) gra.TokenLike {
	return gra.TokenClass().Token(
		line,
		position,
		type_,
		value,
	)
}

// Grammar/Validator

func Validator() gra.ValidatorLike {
	return gra.ValidatorClass().Validator()
}

// Grammar/Visitor

func Visitor(
	processor Methodical,
) gra.VisitorLike {
	return gra.VisitorClass().Visitor(
		processor,
	)
}

// GLOBAL FUNCTIONS

func FormatModel(
	model ModelLike,
) string {
	var formatter = Formatter()
	return formatter.FormatModel(model)
}

func MatchesType(
	tokenValue string,
	tokenType TokenType,
) bool {
	var scannerClass = gra.ScannerClass()
	return scannerClass.MatchesType(tokenValue, tokenType)
}

func ParseSource(
	source string,
) ModelLike {
	var parser = Parser()
	return parser.ParseSource(source)
}

func ValidateModel(
	model ModelLike,
) {
	var validator = Validator()
	validator.ValidateModel(model)
}
