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
│            This "package_api.go" file was automatically generated.           │
│                     Any updates to it may be overwritten.                    │
└──────────────────────────────────────────────────────────────────────────────┘

Package "grammar" provides the following grammar classes that operate on the
abstract syntax tree (AST) for this module:
  - Token captures the attributes associated with a parsed token.
  - Scanner is used to scan the source byte stream and recognize matching tokens.
  - Parser is used to process the token stream and generate the AST.
  - Validator is used to validate the semantics associated with an AST.
  - Formatter is used to format an AST back into a canonical version of its source.
  - Visitor walks the AST and calls processor methods for each node in the tree.
  - Processor provides empty processor methods to be inherited by the processors.

For detailed documentation on this package refer to the wiki:
  - https://github.com/craterdog/go-class-model/wiki

This package follows the Crater Dog Technologies™ Go Coding Conventions located
here:
  - https://github.com/craterdog/go-development-tools/wiki/Coding-Conventions

Additional concrete implementations of the classes declared by this package can
be developed and used seamlessly since the interface declarations only depend on
other interfaces and intrinsic types—and the class implementations only depend
on interfaces, not on each other.
*/
package grammar

import (
	ast "github.com/craterdog/go-class-model/v7/ast"
	col "github.com/craterdog/go-collection-framework/v7"
)

// TYPE DECLARATIONS

/*
TokenType is a constrained type representing any token type recognized by a
scanner.
*/
type TokenType uint8

const (
	ErrorToken TokenType = iota
	CommentToken
	DelimiterToken
	NameToken
	NewlineToken
	PathToken
	PrefixToken
	SpaceToken
)

// FUNCTIONAL DECLARATIONS

// CLASS DECLARATIONS

/*
FormatterClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete formatter-like class.
*/
type FormatterClassLike interface {
	// Constructor Methods
	Formatter() FormatterLike
}

/*
ParserClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete parser-like class.
*/
type ParserClassLike interface {
	// Constructor Methods
	Parser() ParserLike
}

/*
ProcessorClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete processor-like class.
*/
type ProcessorClassLike interface {
	// Constructor Methods
	Processor() ProcessorLike
}

/*
ScannerClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete scanner-like class.  The following functions are supported:

FormatToken() returns a formatted string containing the attributes of the token.

FormatType() returns the string version of the token type.

MatchesType() determines whether or not a token value is of a specified type.
*/
type ScannerClassLike interface {
	// Constructor Methods
	Scanner(
		source string,
		tokens col.QueueLike[TokenLike],
	) ScannerLike

	// Function Methods
	FormatToken(
		token TokenLike,
	) string
	FormatType(
		tokenType TokenType,
	) string
	MatchesType(
		tokenValue string,
		tokenType TokenType,
	) bool
}

/*
TokenClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete visitor-like class.
*/
type TokenClassLike interface {
	// Constructor Methods
	Token(
		line uint,
		position uint,
		type_ TokenType,
		value string,
	) TokenLike
}

/*
ValidatorClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete validator-like class.
*/
type ValidatorClassLike interface {
	// Constructor Methods
	Validator() ValidatorLike
}

/*
VisitorClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete visitor-like class.
*/
type VisitorClassLike interface {
	// Constructor Methods
	Visitor(
		processor Methodical,
	) VisitorLike
}

// INSTANCE DECLARATIONS

/*
FormatterLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each
instance of a concrete formatter-like class.
*/
type FormatterLike interface {
	// Principal Methods
	GetClass() FormatterClassLike
	FormatModel(
		model ast.ModelLike,
	) string

	// Aspect Interfaces
	Methodical
}

/*
ParserLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each
instance of a concrete parser-like class.
*/
type ParserLike interface {
	// Principal Methods
	GetClass() ParserClassLike
	ParseSource(
		source string,
	) ast.ModelLike
}

/*
ProcessorLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each
instance of a concrete processor-like class.
*/
type ProcessorLike interface {
	// Principal Methods
	GetClass() ProcessorClassLike

	// Aspect Interfaces
	Methodical
}

/*
ScannerLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each
instance of a concrete scanner-like class.
*/
type ScannerLike interface {
	// Principal Methods
	GetClass() ScannerClassLike
}

/*
TokenLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each
instance of a concrete token-like class.
*/
type TokenLike interface {
	// Principal Methods
	GetClass() TokenClassLike

	// Attribute Methods
	GetLine() uint
	GetPosition() uint
	GetType() TokenType
	GetValue() string
}

/*
ValidatorLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each
instance of a concrete validator-like class.
*/
type ValidatorLike interface {
	// Principal Methods
	GetClass() ValidatorClassLike
	ValidateModel(
		model ast.ModelLike,
	)

	// Aspect Interfaces
	Methodical
}

/*
VisitorLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each
instance of a concrete visitor-like class.
*/
type VisitorLike interface {
	// Principal Methods
	GetClass() VisitorClassLike
	VisitModel(
		model ast.ModelLike,
	)
}

// ASPECT DECLARATIONS

/*
Methodical declares the set of method signatures that must be supported by
all methodical processors.
*/
type Methodical interface {
	ProcessComment(
		comment string,
	)
	ProcessName(
		name string,
	)
	ProcessNewline(
		newline string,
	)
	ProcessPath(
		path string,
	)
	ProcessPrefix(
		prefix string,
	)
	ProcessSpace(
		space string,
	)
	PreprocessAbstraction(
		abstraction ast.AbstractionLike,
	)
	ProcessAbstractionSlot(
		slot uint,
	)
	PostprocessAbstraction(
		abstraction ast.AbstractionLike,
	)
	PreprocessAdditionalArgument(
		additionalArgument ast.AdditionalArgumentLike,
		index uint,
		size uint,
	)
	ProcessAdditionalArgumentSlot(
		slot uint,
	)
	PostprocessAdditionalArgument(
		additionalArgument ast.AdditionalArgumentLike,
		index uint,
		size uint,
	)
	PreprocessAdditionalConstraint(
		additionalConstraint ast.AdditionalConstraintLike,
		index uint,
		size uint,
	)
	ProcessAdditionalConstraintSlot(
		slot uint,
	)
	PostprocessAdditionalConstraint(
		additionalConstraint ast.AdditionalConstraintLike,
		index uint,
		size uint,
	)
	PreprocessAdditionalValue(
		additionalValue ast.AdditionalValueLike,
		index uint,
		size uint,
	)
	ProcessAdditionalValueSlot(
		slot uint,
	)
	PostprocessAdditionalValue(
		additionalValue ast.AdditionalValueLike,
		index uint,
		size uint,
	)
	PreprocessArgument(
		argument ast.ArgumentLike,
	)
	ProcessArgumentSlot(
		slot uint,
	)
	PostprocessArgument(
		argument ast.ArgumentLike,
	)
	PreprocessArguments(
		arguments ast.ArgumentsLike,
	)
	ProcessArgumentsSlot(
		slot uint,
	)
	PostprocessArguments(
		arguments ast.ArgumentsLike,
	)
	PreprocessArray(
		array ast.ArrayLike,
	)
	ProcessArraySlot(
		slot uint,
	)
	PostprocessArray(
		array ast.ArrayLike,
	)
	PreprocessAspectDeclaration(
		aspectDeclaration ast.AspectDeclarationLike,
		index uint,
		size uint,
	)
	ProcessAspectDeclarationSlot(
		slot uint,
	)
	PostprocessAspectDeclaration(
		aspectDeclaration ast.AspectDeclarationLike,
		index uint,
		size uint,
	)
	PreprocessAspectInterface(
		aspectInterface ast.AspectInterfaceLike,
		index uint,
		size uint,
	)
	ProcessAspectInterfaceSlot(
		slot uint,
	)
	PostprocessAspectInterface(
		aspectInterface ast.AspectInterfaceLike,
		index uint,
		size uint,
	)
	PreprocessAspectMethod(
		aspectMethod ast.AspectMethodLike,
		index uint,
		size uint,
	)
	ProcessAspectMethodSlot(
		slot uint,
	)
	PostprocessAspectMethod(
		aspectMethod ast.AspectMethodLike,
		index uint,
		size uint,
	)
	PreprocessAspectSection(
		aspectSection ast.AspectSectionLike,
	)
	ProcessAspectSectionSlot(
		slot uint,
	)
	PostprocessAspectSection(
		aspectSection ast.AspectSectionLike,
	)
	PreprocessAspectSubsection(
		aspectSubsection ast.AspectSubsectionLike,
	)
	ProcessAspectSubsectionSlot(
		slot uint,
	)
	PostprocessAspectSubsection(
		aspectSubsection ast.AspectSubsectionLike,
	)
	PreprocessAttributeMethod(
		attributeMethod ast.AttributeMethodLike,
		index uint,
		size uint,
	)
	ProcessAttributeMethodSlot(
		slot uint,
	)
	PostprocessAttributeMethod(
		attributeMethod ast.AttributeMethodLike,
		index uint,
		size uint,
	)
	PreprocessAttributeSubsection(
		attributeSubsection ast.AttributeSubsectionLike,
	)
	ProcessAttributeSubsectionSlot(
		slot uint,
	)
	PostprocessAttributeSubsection(
		attributeSubsection ast.AttributeSubsectionLike,
	)
	PreprocessChannel(
		channel ast.ChannelLike,
	)
	ProcessChannelSlot(
		slot uint,
	)
	PostprocessChannel(
		channel ast.ChannelLike,
	)
	PreprocessClassDeclaration(
		classDeclaration ast.ClassDeclarationLike,
		index uint,
		size uint,
	)
	ProcessClassDeclarationSlot(
		slot uint,
	)
	PostprocessClassDeclaration(
		classDeclaration ast.ClassDeclarationLike,
		index uint,
		size uint,
	)
	PreprocessClassMethods(
		classMethods ast.ClassMethodsLike,
	)
	ProcessClassMethodsSlot(
		slot uint,
	)
	PostprocessClassMethods(
		classMethods ast.ClassMethodsLike,
	)
	PreprocessClassSection(
		classSection ast.ClassSectionLike,
	)
	ProcessClassSectionSlot(
		slot uint,
	)
	PostprocessClassSection(
		classSection ast.ClassSectionLike,
	)
	PreprocessConstantMethod(
		constantMethod ast.ConstantMethodLike,
		index uint,
		size uint,
	)
	ProcessConstantMethodSlot(
		slot uint,
	)
	PostprocessConstantMethod(
		constantMethod ast.ConstantMethodLike,
		index uint,
		size uint,
	)
	PreprocessConstantSubsection(
		constantSubsection ast.ConstantSubsectionLike,
	)
	ProcessConstantSubsectionSlot(
		slot uint,
	)
	PostprocessConstantSubsection(
		constantSubsection ast.ConstantSubsectionLike,
	)
	PreprocessConstraint(
		constraint ast.ConstraintLike,
	)
	ProcessConstraintSlot(
		slot uint,
	)
	PostprocessConstraint(
		constraint ast.ConstraintLike,
	)
	PreprocessConstraints(
		constraints ast.ConstraintsLike,
	)
	ProcessConstraintsSlot(
		slot uint,
	)
	PostprocessConstraints(
		constraints ast.ConstraintsLike,
	)
	PreprocessConstructorMethod(
		constructorMethod ast.ConstructorMethodLike,
		index uint,
		size uint,
	)
	ProcessConstructorMethodSlot(
		slot uint,
	)
	PostprocessConstructorMethod(
		constructorMethod ast.ConstructorMethodLike,
		index uint,
		size uint,
	)
	PreprocessConstructorSubsection(
		constructorSubsection ast.ConstructorSubsectionLike,
	)
	ProcessConstructorSubsectionSlot(
		slot uint,
	)
	PostprocessConstructorSubsection(
		constructorSubsection ast.ConstructorSubsectionLike,
	)
	PreprocessDeclaration(
		declaration ast.DeclarationLike,
	)
	ProcessDeclarationSlot(
		slot uint,
	)
	PostprocessDeclaration(
		declaration ast.DeclarationLike,
	)
	PreprocessEnumeration(
		enumeration ast.EnumerationLike,
	)
	ProcessEnumerationSlot(
		slot uint,
	)
	PostprocessEnumeration(
		enumeration ast.EnumerationLike,
	)
	PreprocessFunctionMethod(
		functionMethod ast.FunctionMethodLike,
		index uint,
		size uint,
	)
	ProcessFunctionMethodSlot(
		slot uint,
	)
	PostprocessFunctionMethod(
		functionMethod ast.FunctionMethodLike,
		index uint,
		size uint,
	)
	PreprocessFunctionSubsection(
		functionSubsection ast.FunctionSubsectionLike,
	)
	ProcessFunctionSubsectionSlot(
		slot uint,
	)
	PostprocessFunctionSubsection(
		functionSubsection ast.FunctionSubsectionLike,
	)
	PreprocessFunctionalDeclaration(
		functionalDeclaration ast.FunctionalDeclarationLike,
		index uint,
		size uint,
	)
	ProcessFunctionalDeclarationSlot(
		slot uint,
	)
	PostprocessFunctionalDeclaration(
		functionalDeclaration ast.FunctionalDeclarationLike,
		index uint,
		size uint,
	)
	PreprocessFunctionalSection(
		functionalSection ast.FunctionalSectionLike,
	)
	ProcessFunctionalSectionSlot(
		slot uint,
	)
	PostprocessFunctionalSection(
		functionalSection ast.FunctionalSectionLike,
	)
	PreprocessGetterMethod(
		getterMethod ast.GetterMethodLike,
	)
	ProcessGetterMethodSlot(
		slot uint,
	)
	PostprocessGetterMethod(
		getterMethod ast.GetterMethodLike,
	)
	PreprocessImportedPackage(
		importedPackage ast.ImportedPackageLike,
		index uint,
		size uint,
	)
	ProcessImportedPackageSlot(
		slot uint,
	)
	PostprocessImportedPackage(
		importedPackage ast.ImportedPackageLike,
		index uint,
		size uint,
	)
	PreprocessInstanceDeclaration(
		instanceDeclaration ast.InstanceDeclarationLike,
		index uint,
		size uint,
	)
	ProcessInstanceDeclarationSlot(
		slot uint,
	)
	PostprocessInstanceDeclaration(
		instanceDeclaration ast.InstanceDeclarationLike,
		index uint,
		size uint,
	)
	PreprocessInstanceMethods(
		instanceMethods ast.InstanceMethodsLike,
	)
	ProcessInstanceMethodsSlot(
		slot uint,
	)
	PostprocessInstanceMethods(
		instanceMethods ast.InstanceMethodsLike,
	)
	PreprocessInstanceSection(
		instanceSection ast.InstanceSectionLike,
	)
	ProcessInstanceSectionSlot(
		slot uint,
	)
	PostprocessInstanceSection(
		instanceSection ast.InstanceSectionLike,
	)
	PreprocessInterfaceDeclarations(
		interfaceDeclarations ast.InterfaceDeclarationsLike,
	)
	ProcessInterfaceDeclarationsSlot(
		slot uint,
	)
	PostprocessInterfaceDeclarations(
		interfaceDeclarations ast.InterfaceDeclarationsLike,
	)
	PreprocessLegalNotice(
		legalNotice ast.LegalNoticeLike,
	)
	ProcessLegalNoticeSlot(
		slot uint,
	)
	PostprocessLegalNotice(
		legalNotice ast.LegalNoticeLike,
	)
	PreprocessMap(
		map_ ast.MapLike,
	)
	ProcessMapSlot(
		slot uint,
	)
	PostprocessMap(
		map_ ast.MapLike,
	)
	PreprocessMethod(
		method ast.MethodLike,
	)
	ProcessMethodSlot(
		slot uint,
	)
	PostprocessMethod(
		method ast.MethodLike,
	)
	PreprocessModel(
		model ast.ModelLike,
	)
	ProcessModelSlot(
		slot uint,
	)
	PostprocessModel(
		model ast.ModelLike,
	)
	PreprocessMultivalue(
		multivalue ast.MultivalueLike,
	)
	ProcessMultivalueSlot(
		slot uint,
	)
	PostprocessMultivalue(
		multivalue ast.MultivalueLike,
	)
	PreprocessNone(
		none ast.NoneLike,
	)
	ProcessNoneSlot(
		slot uint,
	)
	PostprocessNone(
		none ast.NoneLike,
	)
	PreprocessPackageDeclaration(
		packageDeclaration ast.PackageDeclarationLike,
	)
	ProcessPackageDeclarationSlot(
		slot uint,
	)
	PostprocessPackageDeclaration(
		packageDeclaration ast.PackageDeclarationLike,
	)
	PreprocessPackageHeader(
		packageHeader ast.PackageHeaderLike,
	)
	ProcessPackageHeaderSlot(
		slot uint,
	)
	PostprocessPackageHeader(
		packageHeader ast.PackageHeaderLike,
	)
	PreprocessPackageImports(
		packageImports ast.PackageImportsLike,
	)
	ProcessPackageImportsSlot(
		slot uint,
	)
	PostprocessPackageImports(
		packageImports ast.PackageImportsLike,
	)
	PreprocessParameter(
		parameter ast.ParameterLike,
		index uint,
		size uint,
	)
	ProcessParameterSlot(
		slot uint,
	)
	PostprocessParameter(
		parameter ast.ParameterLike,
		index uint,
		size uint,
	)
	PreprocessPrimitiveDeclarations(
		primitiveDeclarations ast.PrimitiveDeclarationsLike,
	)
	ProcessPrimitiveDeclarationsSlot(
		slot uint,
	)
	PostprocessPrimitiveDeclarations(
		primitiveDeclarations ast.PrimitiveDeclarationsLike,
	)
	PreprocessPrincipalMethod(
		principalMethod ast.PrincipalMethodLike,
		index uint,
		size uint,
	)
	ProcessPrincipalMethodSlot(
		slot uint,
	)
	PostprocessPrincipalMethod(
		principalMethod ast.PrincipalMethodLike,
		index uint,
		size uint,
	)
	PreprocessPrincipalSubsection(
		principalSubsection ast.PrincipalSubsectionLike,
	)
	ProcessPrincipalSubsectionSlot(
		slot uint,
	)
	PostprocessPrincipalSubsection(
		principalSubsection ast.PrincipalSubsectionLike,
	)
	PreprocessResult(
		result ast.ResultLike,
	)
	ProcessResultSlot(
		slot uint,
	)
	PostprocessResult(
		result ast.ResultLike,
	)
	PreprocessSetterMethod(
		setterMethod ast.SetterMethodLike,
	)
	ProcessSetterMethodSlot(
		slot uint,
	)
	PostprocessSetterMethod(
		setterMethod ast.SetterMethodLike,
	)
	PreprocessTypeDeclaration(
		typeDeclaration ast.TypeDeclarationLike,
		index uint,
		size uint,
	)
	ProcessTypeDeclarationSlot(
		slot uint,
	)
	PostprocessTypeDeclaration(
		typeDeclaration ast.TypeDeclarationLike,
		index uint,
		size uint,
	)
	PreprocessTypeSection(
		typeSection ast.TypeSectionLike,
	)
	ProcessTypeSectionSlot(
		slot uint,
	)
	PostprocessTypeSection(
		typeSection ast.TypeSectionLike,
	)
	PreprocessValue(
		value ast.ValueLike,
	)
	ProcessValueSlot(
		slot uint,
	)
	PostprocessValue(
		value ast.ValueLike,
	)
	PreprocessWrapper(
		wrapper ast.WrapperLike,
	)
	ProcessWrapperSlot(
		slot uint,
	)
	PostprocessWrapper(
		wrapper ast.WrapperLike,
	)
}
