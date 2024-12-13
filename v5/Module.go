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
the packages contained in this module.  It also provides a default constructor
for each commonly used class that is exported by the module.  Each constructor
delegates the actual construction process to its corresponding concrete class
declared in the corresponding package contained within this module.

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

// DEFAULT CONSTRUCTORS

// Ast

func Abstraction(
	optionalPrefix PrefixLike,
	name string,
	optionalSuffix SuffixLike,
	optionalArguments ArgumentsLike,
) AbstractionLike {
	return ast.AbstractionClass().Make(
		optionalPrefix,
		name,
		optionalSuffix,
		optionalArguments,
	)
}

func AdditionalArgument(
	argument ArgumentLike,
) AdditionalArgumentLike {
	return ast.AdditionalArgumentClass().Make(
		argument,
	)
}

func AdditionalConstraint(
	constraint ConstraintLike,
) AdditionalConstraintLike {
	return ast.AdditionalConstraintClass().Make(
		constraint,
	)
}

func AdditionalValue(
	name string,
) AdditionalValueLike {
	return ast.AdditionalValueClass().Make(
		name,
	)
}

func Argument(
	abstraction AbstractionLike,
) ArgumentLike {
	return ast.ArgumentClass().Make(
		abstraction,
	)
}

func Arguments(
	argument ArgumentLike,
	additionalArguments abs.Sequential[AdditionalArgumentLike],
) ArgumentsLike {
	return ast.ArgumentsClass().Make(
		argument,
		additionalArguments,
	)
}

func Array() ArrayLike {
	return ast.ArrayClass().Make()
}

func AspectDeclaration(
	declaration DeclarationLike,
	aspectMethods abs.Sequential[AspectMethodLike],
) AspectDeclarationLike {
	return ast.AspectDeclarationClass().Make(
		declaration,
		aspectMethods,
	)
}

func AspectInterface(
	abstraction AbstractionLike,
) AspectInterfaceLike {
	return ast.AspectInterfaceClass().Make(
		abstraction,
	)
}

func AspectMethod(
	method MethodLike,
) AspectMethodLike {
	return ast.AspectMethodClass().Make(
		method,
	)
}

func AspectSection(
	aspectDeclarations abs.Sequential[AspectDeclarationLike],
) AspectSectionLike {
	return ast.AspectSectionClass().Make(
		aspectDeclarations,
	)
}

func AspectSubsection(
	aspectInterfaces abs.Sequential[AspectInterfaceLike],
) AspectSubsectionLike {
	return ast.AspectSubsectionClass().Make(
		aspectInterfaces,
	)
}

func AttributeMethod(
	any_ any,
) AttributeMethodLike {
	return ast.AttributeMethodClass().Make(
		any_,
	)
}

func AttributeSubsection(
	attributeMethods abs.Sequential[AttributeMethodLike],
) AttributeSubsectionLike {
	return ast.AttributeSubsectionClass().Make(
		attributeMethods,
	)
}

func Channel() ChannelLike {
	return ast.ChannelClass().Make()
}

func ClassDeclaration(
	declaration DeclarationLike,
	classMethods ClassMethodsLike,
) ClassDeclarationLike {
	return ast.ClassDeclarationClass().Make(
		declaration,
		classMethods,
	)
}

func ClassMethods(
	constructorSubsection ConstructorSubsectionLike,
	optionalConstantSubsection ConstantSubsectionLike,
	optionalFunctionSubsection FunctionSubsectionLike,
) ClassMethodsLike {
	return ast.ClassMethodsClass().Make(
		constructorSubsection,
		optionalConstantSubsection,
		optionalFunctionSubsection,
	)
}

func ClassSection(
	classDeclarations abs.Sequential[ClassDeclarationLike],
) ClassSectionLike {
	return ast.ClassSectionClass().Make(
		classDeclarations,
	)
}

func ConstantMethod(
	name string,
	abstraction AbstractionLike,
) ConstantMethodLike {
	return ast.ConstantMethodClass().Make(
		name,
		abstraction,
	)
}

func ConstantSubsection(
	constantMethods abs.Sequential[ConstantMethodLike],
) ConstantSubsectionLike {
	return ast.ConstantSubsectionClass().Make(
		constantMethods,
	)
}

func Constraint(
	name string,
	abstraction AbstractionLike,
) ConstraintLike {
	return ast.ConstraintClass().Make(
		name,
		abstraction,
	)
}

func Constraints(
	constraint ConstraintLike,
	additionalConstraints abs.Sequential[AdditionalConstraintLike],
) ConstraintsLike {
	return ast.ConstraintsClass().Make(
		constraint,
		additionalConstraints,
	)
}

func ConstructorMethod(
	name string,
	parameters abs.Sequential[ParameterLike],
	abstraction AbstractionLike,
) ConstructorMethodLike {
	return ast.ConstructorMethodClass().Make(
		name,
		parameters,
		abstraction,
	)
}

func ConstructorSubsection(
	constructorMethods abs.Sequential[ConstructorMethodLike],
) ConstructorSubsectionLike {
	return ast.ConstructorSubsectionClass().Make(
		constructorMethods,
	)
}

func Declaration(
	comment string,
	name string,
	optionalConstraints ConstraintsLike,
) DeclarationLike {
	return ast.DeclarationClass().Make(
		comment,
		name,
		optionalConstraints,
	)
}

func Enumeration(
	value ValueLike,
	additionalValues abs.Sequential[AdditionalValueLike],
) EnumerationLike {
	return ast.EnumerationClass().Make(
		value,
		additionalValues,
	)
}

func FunctionMethod(
	name string,
	parameters abs.Sequential[ParameterLike],
	result ResultLike,
) FunctionMethodLike {
	return ast.FunctionMethodClass().Make(
		name,
		parameters,
		result,
	)
}

func FunctionSubsection(
	functionMethods abs.Sequential[FunctionMethodLike],
) FunctionSubsectionLike {
	return ast.FunctionSubsectionClass().Make(
		functionMethods,
	)
}

func FunctionalDeclaration(
	declaration DeclarationLike,
	parameters abs.Sequential[ParameterLike],
	result ResultLike,
) FunctionalDeclarationLike {
	return ast.FunctionalDeclarationClass().Make(
		declaration,
		parameters,
		result,
	)
}

func FunctionalSection(
	functionalDeclarations abs.Sequential[FunctionalDeclarationLike],
) FunctionalSectionLike {
	return ast.FunctionalSectionClass().Make(
		functionalDeclarations,
	)
}

func GetterMethod(
	name string,
	abstraction AbstractionLike,
) GetterMethodLike {
	return ast.GetterMethodClass().Make(
		name,
		abstraction,
	)
}

func ImportedPackage(
	name string,
	path string,
) ImportedPackageLike {
	return ast.ImportedPackageClass().Make(
		name,
		path,
	)
}

func InstanceDeclaration(
	declaration DeclarationLike,
	instanceMethods InstanceMethodsLike,
) InstanceDeclarationLike {
	return ast.InstanceDeclarationClass().Make(
		declaration,
		instanceMethods,
	)
}

func InstanceMethods(
	principalSubsection PrincipalSubsectionLike,
	optionalAttributeSubsection AttributeSubsectionLike,
	optionalAspectSubsection AspectSubsectionLike,
) InstanceMethodsLike {
	return ast.InstanceMethodsClass().Make(
		principalSubsection,
		optionalAttributeSubsection,
		optionalAspectSubsection,
	)
}

func InstanceSection(
	instanceDeclarations abs.Sequential[InstanceDeclarationLike],
) InstanceSectionLike {
	return ast.InstanceSectionClass().Make(
		instanceDeclarations,
	)
}

func InterfaceDeclarations(
	classSection ClassSectionLike,
	instanceSection InstanceSectionLike,
	aspectSection AspectSectionLike,
) InterfaceDeclarationsLike {
	return ast.InterfaceDeclarationsClass().Make(
		classSection,
		instanceSection,
		aspectSection,
	)
}

func LegalNotice(
	comment string,
) LegalNoticeLike {
	return ast.LegalNoticeClass().Make(
		comment,
	)
}

func Map(
	name string,
) MapLike {
	return ast.MapClass().Make(
		name,
	)
}

func Method(
	name string,
	parameters abs.Sequential[ParameterLike],
	optionalResult ResultLike,
) MethodLike {
	return ast.MethodClass().Make(
		name,
		parameters,
		optionalResult,
	)
}

func Model(
	packageDeclaration PackageDeclarationLike,
	primitiveDeclarations PrimitiveDeclarationsLike,
	interfaceDeclarations InterfaceDeclarationsLike,
) ModelLike {
	return ast.ModelClass().Make(
		packageDeclaration,
		primitiveDeclarations,
		interfaceDeclarations,
	)
}

func Multivalue(
	parameters abs.Sequential[ParameterLike],
) MultivalueLike {
	return ast.MultivalueClass().Make(
		parameters,
	)
}

func None(
	newline string,
) NoneLike {
	return ast.NoneClass().Make(
		newline,
	)
}

func PackageDeclaration(
	legalNotice LegalNoticeLike,
	packageHeader PackageHeaderLike,
	packageImports PackageImportsLike,
) PackageDeclarationLike {
	return ast.PackageDeclarationClass().Make(
		legalNotice,
		packageHeader,
		packageImports,
	)
}

func PackageHeader(
	comment string,
	name string,
) PackageHeaderLike {
	return ast.PackageHeaderClass().Make(
		comment,
		name,
	)
}

func PackageImports(
	importedPackages abs.Sequential[ImportedPackageLike],
) PackageImportsLike {
	return ast.PackageImportsClass().Make(
		importedPackages,
	)
}

func Parameter(
	name string,
	abstraction AbstractionLike,
) ParameterLike {
	return ast.ParameterClass().Make(
		name,
		abstraction,
	)
}

func Prefix(
	any_ any,
) PrefixLike {
	return ast.PrefixClass().Make(
		any_,
	)
}

func PrimitiveDeclarations(
	typeSection TypeSectionLike,
	functionalSection FunctionalSectionLike,
) PrimitiveDeclarationsLike {
	return ast.PrimitiveDeclarationsClass().Make(
		typeSection,
		functionalSection,
	)
}

func PrincipalMethod(
	method MethodLike,
) PrincipalMethodLike {
	return ast.PrincipalMethodClass().Make(
		method,
	)
}

func PrincipalSubsection(
	principalMethods abs.Sequential[PrincipalMethodLike],
) PrincipalSubsectionLike {
	return ast.PrincipalSubsectionClass().Make(
		principalMethods,
	)
}

func Result(
	any_ any,
) ResultLike {
	return ast.ResultClass().Make(
		any_,
	)
}

func SetterMethod(
	name string,
	parameter ParameterLike,
) SetterMethodLike {
	return ast.SetterMethodClass().Make(
		name,
		parameter,
	)
}

func Suffix(
	name string,
) SuffixLike {
	return ast.SuffixClass().Make(
		name,
	)
}

func TypeDeclaration(
	declaration DeclarationLike,
	abstraction AbstractionLike,
	optionalEnumeration EnumerationLike,
) TypeDeclarationLike {
	return ast.TypeDeclarationClass().Make(
		declaration,
		abstraction,
		optionalEnumeration,
	)
}

func TypeSection(
	typeDeclarations abs.Sequential[TypeDeclarationLike],
) TypeSectionLike {
	return ast.TypeSectionClass().Make(
		typeDeclarations,
	)
}

func Value(
	name string,
	abstraction AbstractionLike,
) ValueLike {
	return ast.ValueClass().Make(
		name,
		abstraction,
	)
}

// Grammar

func Formatter() FormatterLike {
	return gra.FormatterClass().Make()
}

func Parser() ParserLike {
	return gra.ParserClass().Make()
}

func Processor() ProcessorLike {
	return gra.ProcessorClass().Make()
}

func Scanner(
	source string,
	tokens abs.QueueLike[TokenLike],
) ScannerLike {
	return gra.ScannerClass().Make(
		source,
		tokens,
	)
}

func Token(
	line uint,
	position uint,
	type_ TokenType,
	value string,
) TokenLike {
	return gra.TokenClass().Make(
		line,
		position,
		type_,
		value,
	)
}

func Validator() ValidatorLike {
	return gra.ValidatorClass().Make()
}

func Visitor(
	processor Methodical,
) VisitorLike {
	return gra.VisitorClass().Make(
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
