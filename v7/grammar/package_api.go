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
	ProcessDelimiter(
		delimiter string,
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
		index uint,
		count uint,
	)
	ProcessAbstractionSlot(
		slot uint,
	)
	PostprocessAbstraction(
		abstraction ast.AbstractionLike,
		index uint,
		count uint,
	)
	PreprocessAdditionalArgument(
		additionalArgument ast.AdditionalArgumentLike,
		index uint,
		count uint,
	)
	ProcessAdditionalArgumentSlot(
		slot uint,
	)
	PostprocessAdditionalArgument(
		additionalArgument ast.AdditionalArgumentLike,
		index uint,
		count uint,
	)
	PreprocessAdditionalConstraint(
		additionalConstraint ast.AdditionalConstraintLike,
		index uint,
		count uint,
	)
	ProcessAdditionalConstraintSlot(
		slot uint,
	)
	PostprocessAdditionalConstraint(
		additionalConstraint ast.AdditionalConstraintLike,
		index uint,
		count uint,
	)
	PreprocessAdditionalValue(
		additionalValue ast.AdditionalValueLike,
		index uint,
		count uint,
	)
	ProcessAdditionalValueSlot(
		slot uint,
	)
	PostprocessAdditionalValue(
		additionalValue ast.AdditionalValueLike,
		index uint,
		count uint,
	)
	PreprocessArgument(
		argument ast.ArgumentLike,
		index uint,
		count uint,
	)
	ProcessArgumentSlot(
		slot uint,
	)
	PostprocessArgument(
		argument ast.ArgumentLike,
		index uint,
		count uint,
	)
	PreprocessArguments(
		arguments ast.ArgumentsLike,
		index uint,
		count uint,
	)
	ProcessArgumentsSlot(
		slot uint,
	)
	PostprocessArguments(
		arguments ast.ArgumentsLike,
		index uint,
		count uint,
	)
	PreprocessArray(
		array ast.ArrayLike,
		index uint,
		count uint,
	)
	ProcessArraySlot(
		slot uint,
	)
	PostprocessArray(
		array ast.ArrayLike,
		index uint,
		count uint,
	)
	PreprocessAspectDeclaration(
		aspectDeclaration ast.AspectDeclarationLike,
		index uint,
		count uint,
	)
	ProcessAspectDeclarationSlot(
		slot uint,
	)
	PostprocessAspectDeclaration(
		aspectDeclaration ast.AspectDeclarationLike,
		index uint,
		count uint,
	)
	PreprocessAspectInterface(
		aspectInterface ast.AspectInterfaceLike,
		index uint,
		count uint,
	)
	ProcessAspectInterfaceSlot(
		slot uint,
	)
	PostprocessAspectInterface(
		aspectInterface ast.AspectInterfaceLike,
		index uint,
		count uint,
	)
	PreprocessAspectMethod(
		aspectMethod ast.AspectMethodLike,
		index uint,
		count uint,
	)
	ProcessAspectMethodSlot(
		slot uint,
	)
	PostprocessAspectMethod(
		aspectMethod ast.AspectMethodLike,
		index uint,
		count uint,
	)
	PreprocessAspectSection(
		aspectSection ast.AspectSectionLike,
		index uint,
		count uint,
	)
	ProcessAspectSectionSlot(
		slot uint,
	)
	PostprocessAspectSection(
		aspectSection ast.AspectSectionLike,
		index uint,
		count uint,
	)
	PreprocessAspectSubsection(
		aspectSubsection ast.AspectSubsectionLike,
		index uint,
		count uint,
	)
	ProcessAspectSubsectionSlot(
		slot uint,
	)
	PostprocessAspectSubsection(
		aspectSubsection ast.AspectSubsectionLike,
		index uint,
		count uint,
	)
	PreprocessAttributeMethod(
		attributeMethod ast.AttributeMethodLike,
		index uint,
		count uint,
	)
	ProcessAttributeMethodSlot(
		slot uint,
	)
	PostprocessAttributeMethod(
		attributeMethod ast.AttributeMethodLike,
		index uint,
		count uint,
	)
	PreprocessAttributeSubsection(
		attributeSubsection ast.AttributeSubsectionLike,
		index uint,
		count uint,
	)
	ProcessAttributeSubsectionSlot(
		slot uint,
	)
	PostprocessAttributeSubsection(
		attributeSubsection ast.AttributeSubsectionLike,
		index uint,
		count uint,
	)
	PreprocessChannel(
		channel ast.ChannelLike,
		index uint,
		count uint,
	)
	ProcessChannelSlot(
		slot uint,
	)
	PostprocessChannel(
		channel ast.ChannelLike,
		index uint,
		count uint,
	)
	PreprocessClassDeclaration(
		classDeclaration ast.ClassDeclarationLike,
		index uint,
		count uint,
	)
	ProcessClassDeclarationSlot(
		slot uint,
	)
	PostprocessClassDeclaration(
		classDeclaration ast.ClassDeclarationLike,
		index uint,
		count uint,
	)
	PreprocessClassMethods(
		classMethods ast.ClassMethodsLike,
		index uint,
		count uint,
	)
	ProcessClassMethodsSlot(
		slot uint,
	)
	PostprocessClassMethods(
		classMethods ast.ClassMethodsLike,
		index uint,
		count uint,
	)
	PreprocessClassSection(
		classSection ast.ClassSectionLike,
		index uint,
		count uint,
	)
	ProcessClassSectionSlot(
		slot uint,
	)
	PostprocessClassSection(
		classSection ast.ClassSectionLike,
		index uint,
		count uint,
	)
	PreprocessConstantMethod(
		constantMethod ast.ConstantMethodLike,
		index uint,
		count uint,
	)
	ProcessConstantMethodSlot(
		slot uint,
	)
	PostprocessConstantMethod(
		constantMethod ast.ConstantMethodLike,
		index uint,
		count uint,
	)
	PreprocessConstantSubsection(
		constantSubsection ast.ConstantSubsectionLike,
		index uint,
		count uint,
	)
	ProcessConstantSubsectionSlot(
		slot uint,
	)
	PostprocessConstantSubsection(
		constantSubsection ast.ConstantSubsectionLike,
		index uint,
		count uint,
	)
	PreprocessConstraint(
		constraint ast.ConstraintLike,
		index uint,
		count uint,
	)
	ProcessConstraintSlot(
		slot uint,
	)
	PostprocessConstraint(
		constraint ast.ConstraintLike,
		index uint,
		count uint,
	)
	PreprocessConstraints(
		constraints ast.ConstraintsLike,
		index uint,
		count uint,
	)
	ProcessConstraintsSlot(
		slot uint,
	)
	PostprocessConstraints(
		constraints ast.ConstraintsLike,
		index uint,
		count uint,
	)
	PreprocessConstructorMethod(
		constructorMethod ast.ConstructorMethodLike,
		index uint,
		count uint,
	)
	ProcessConstructorMethodSlot(
		slot uint,
	)
	PostprocessConstructorMethod(
		constructorMethod ast.ConstructorMethodLike,
		index uint,
		count uint,
	)
	PreprocessConstructorSubsection(
		constructorSubsection ast.ConstructorSubsectionLike,
		index uint,
		count uint,
	)
	ProcessConstructorSubsectionSlot(
		slot uint,
	)
	PostprocessConstructorSubsection(
		constructorSubsection ast.ConstructorSubsectionLike,
		index uint,
		count uint,
	)
	PreprocessDeclaration(
		declaration ast.DeclarationLike,
		index uint,
		count uint,
	)
	ProcessDeclarationSlot(
		slot uint,
	)
	PostprocessDeclaration(
		declaration ast.DeclarationLike,
		index uint,
		count uint,
	)
	PreprocessEnumeration(
		enumeration ast.EnumerationLike,
		index uint,
		count uint,
	)
	ProcessEnumerationSlot(
		slot uint,
	)
	PostprocessEnumeration(
		enumeration ast.EnumerationLike,
		index uint,
		count uint,
	)
	PreprocessFunctionMethod(
		functionMethod ast.FunctionMethodLike,
		index uint,
		count uint,
	)
	ProcessFunctionMethodSlot(
		slot uint,
	)
	PostprocessFunctionMethod(
		functionMethod ast.FunctionMethodLike,
		index uint,
		count uint,
	)
	PreprocessFunctionSubsection(
		functionSubsection ast.FunctionSubsectionLike,
		index uint,
		count uint,
	)
	ProcessFunctionSubsectionSlot(
		slot uint,
	)
	PostprocessFunctionSubsection(
		functionSubsection ast.FunctionSubsectionLike,
		index uint,
		count uint,
	)
	PreprocessFunctionalDeclaration(
		functionalDeclaration ast.FunctionalDeclarationLike,
		index uint,
		count uint,
	)
	ProcessFunctionalDeclarationSlot(
		slot uint,
	)
	PostprocessFunctionalDeclaration(
		functionalDeclaration ast.FunctionalDeclarationLike,
		index uint,
		count uint,
	)
	PreprocessFunctionalSection(
		functionalSection ast.FunctionalSectionLike,
		index uint,
		count uint,
	)
	ProcessFunctionalSectionSlot(
		slot uint,
	)
	PostprocessFunctionalSection(
		functionalSection ast.FunctionalSectionLike,
		index uint,
		count uint,
	)
	PreprocessGetterMethod(
		getterMethod ast.GetterMethodLike,
		index uint,
		count uint,
	)
	ProcessGetterMethodSlot(
		slot uint,
	)
	PostprocessGetterMethod(
		getterMethod ast.GetterMethodLike,
		index uint,
		count uint,
	)
	PreprocessImportList(
		importList ast.ImportListLike,
		index uint,
		count uint,
	)
	ProcessImportListSlot(
		slot uint,
	)
	PostprocessImportList(
		importList ast.ImportListLike,
		index uint,
		count uint,
	)
	PreprocessImportedPackage(
		importedPackage ast.ImportedPackageLike,
		index uint,
		count uint,
	)
	ProcessImportedPackageSlot(
		slot uint,
	)
	PostprocessImportedPackage(
		importedPackage ast.ImportedPackageLike,
		index uint,
		count uint,
	)
	PreprocessInstanceDeclaration(
		instanceDeclaration ast.InstanceDeclarationLike,
		index uint,
		count uint,
	)
	ProcessInstanceDeclarationSlot(
		slot uint,
	)
	PostprocessInstanceDeclaration(
		instanceDeclaration ast.InstanceDeclarationLike,
		index uint,
		count uint,
	)
	PreprocessInstanceMethods(
		instanceMethods ast.InstanceMethodsLike,
		index uint,
		count uint,
	)
	ProcessInstanceMethodsSlot(
		slot uint,
	)
	PostprocessInstanceMethods(
		instanceMethods ast.InstanceMethodsLike,
		index uint,
		count uint,
	)
	PreprocessInstanceSection(
		instanceSection ast.InstanceSectionLike,
		index uint,
		count uint,
	)
	ProcessInstanceSectionSlot(
		slot uint,
	)
	PostprocessInstanceSection(
		instanceSection ast.InstanceSectionLike,
		index uint,
		count uint,
	)
	PreprocessInterfaceDeclarations(
		interfaceDeclarations ast.InterfaceDeclarationsLike,
		index uint,
		count uint,
	)
	ProcessInterfaceDeclarationsSlot(
		slot uint,
	)
	PostprocessInterfaceDeclarations(
		interfaceDeclarations ast.InterfaceDeclarationsLike,
		index uint,
		count uint,
	)
	PreprocessLegalNotice(
		legalNotice ast.LegalNoticeLike,
		index uint,
		count uint,
	)
	ProcessLegalNoticeSlot(
		slot uint,
	)
	PostprocessLegalNotice(
		legalNotice ast.LegalNoticeLike,
		index uint,
		count uint,
	)
	PreprocessMap(
		map_ ast.MapLike,
		index uint,
		count uint,
	)
	ProcessMapSlot(
		slot uint,
	)
	PostprocessMap(
		map_ ast.MapLike,
		index uint,
		count uint,
	)
	PreprocessMethod(
		method ast.MethodLike,
		index uint,
		count uint,
	)
	ProcessMethodSlot(
		slot uint,
	)
	PostprocessMethod(
		method ast.MethodLike,
		index uint,
		count uint,
	)
	PreprocessModel(
		model ast.ModelLike,
		index uint,
		count uint,
	)
	ProcessModelSlot(
		slot uint,
	)
	PostprocessModel(
		model ast.ModelLike,
		index uint,
		count uint,
	)
	PreprocessMultivalue(
		multivalue ast.MultivalueLike,
		index uint,
		count uint,
	)
	ProcessMultivalueSlot(
		slot uint,
	)
	PostprocessMultivalue(
		multivalue ast.MultivalueLike,
		index uint,
		count uint,
	)
	PreprocessNone(
		none ast.NoneLike,
		index uint,
		count uint,
	)
	ProcessNoneSlot(
		slot uint,
	)
	PostprocessNone(
		none ast.NoneLike,
		index uint,
		count uint,
	)
	PreprocessPackageDeclaration(
		packageDeclaration ast.PackageDeclarationLike,
		index uint,
		count uint,
	)
	ProcessPackageDeclarationSlot(
		slot uint,
	)
	PostprocessPackageDeclaration(
		packageDeclaration ast.PackageDeclarationLike,
		index uint,
		count uint,
	)
	PreprocessPackageHeader(
		packageHeader ast.PackageHeaderLike,
		index uint,
		count uint,
	)
	ProcessPackageHeaderSlot(
		slot uint,
	)
	PostprocessPackageHeader(
		packageHeader ast.PackageHeaderLike,
		index uint,
		count uint,
	)
	PreprocessPackageImports(
		packageImports ast.PackageImportsLike,
		index uint,
		count uint,
	)
	ProcessPackageImportsSlot(
		slot uint,
	)
	PostprocessPackageImports(
		packageImports ast.PackageImportsLike,
		index uint,
		count uint,
	)
	PreprocessParameter(
		parameter ast.ParameterLike,
		index uint,
		count uint,
	)
	ProcessParameterSlot(
		slot uint,
	)
	PostprocessParameter(
		parameter ast.ParameterLike,
		index uint,
		count uint,
	)
	PreprocessParameterList(
		parameterList ast.ParameterListLike,
		index uint,
		count uint,
	)
	ProcessParameterListSlot(
		slot uint,
	)
	PostprocessParameterList(
		parameterList ast.ParameterListLike,
		index uint,
		count uint,
	)
	PreprocessPrimitiveDeclarations(
		primitiveDeclarations ast.PrimitiveDeclarationsLike,
		index uint,
		count uint,
	)
	ProcessPrimitiveDeclarationsSlot(
		slot uint,
	)
	PostprocessPrimitiveDeclarations(
		primitiveDeclarations ast.PrimitiveDeclarationsLike,
		index uint,
		count uint,
	)
	PreprocessPrincipalMethod(
		principalMethod ast.PrincipalMethodLike,
		index uint,
		count uint,
	)
	ProcessPrincipalMethodSlot(
		slot uint,
	)
	PostprocessPrincipalMethod(
		principalMethod ast.PrincipalMethodLike,
		index uint,
		count uint,
	)
	PreprocessPrincipalSubsection(
		principalSubsection ast.PrincipalSubsectionLike,
		index uint,
		count uint,
	)
	ProcessPrincipalSubsectionSlot(
		slot uint,
	)
	PostprocessPrincipalSubsection(
		principalSubsection ast.PrincipalSubsectionLike,
		index uint,
		count uint,
	)
	PreprocessResult(
		result ast.ResultLike,
		index uint,
		count uint,
	)
	ProcessResultSlot(
		slot uint,
	)
	PostprocessResult(
		result ast.ResultLike,
		index uint,
		count uint,
	)
	PreprocessSetterMethod(
		setterMethod ast.SetterMethodLike,
		index uint,
		count uint,
	)
	ProcessSetterMethodSlot(
		slot uint,
	)
	PostprocessSetterMethod(
		setterMethod ast.SetterMethodLike,
		index uint,
		count uint,
	)
	PreprocessTypeDeclaration(
		typeDeclaration ast.TypeDeclarationLike,
		index uint,
		count uint,
	)
	ProcessTypeDeclarationSlot(
		slot uint,
	)
	PostprocessTypeDeclaration(
		typeDeclaration ast.TypeDeclarationLike,
		index uint,
		count uint,
	)
	PreprocessTypeSection(
		typeSection ast.TypeSectionLike,
		index uint,
		count uint,
	)
	ProcessTypeSectionSlot(
		slot uint,
	)
	PostprocessTypeSection(
		typeSection ast.TypeSectionLike,
		index uint,
		count uint,
	)
	PreprocessValue(
		value ast.ValueLike,
		index uint,
		count uint,
	)
	ProcessValueSlot(
		slot uint,
	)
	PostprocessValue(
		value ast.ValueLike,
		index uint,
		count uint,
	)
	PreprocessWrapper(
		wrapper ast.WrapperLike,
		index uint,
		count uint,
	)
	ProcessWrapperSlot(
		slot uint,
	)
	PostprocessWrapper(
		wrapper ast.WrapperLike,
		index uint,
		count uint,
	)
}
