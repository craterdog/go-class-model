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
		index_ uint,
		count_ uint,
	)
	ProcessAbstractionSlot(
		slot uint,
	)
	PostprocessAbstraction(
		abstraction ast.AbstractionLike,
		index_ uint,
		count_ uint,
	)
	PreprocessAdditionalArgument(
		additionalArgument ast.AdditionalArgumentLike,
		index_ uint,
		count_ uint,
	)
	ProcessAdditionalArgumentSlot(
		slot uint,
	)
	PostprocessAdditionalArgument(
		additionalArgument ast.AdditionalArgumentLike,
		index_ uint,
		count_ uint,
	)
	PreprocessAdditionalConstraint(
		additionalConstraint ast.AdditionalConstraintLike,
		index_ uint,
		count_ uint,
	)
	ProcessAdditionalConstraintSlot(
		slot uint,
	)
	PostprocessAdditionalConstraint(
		additionalConstraint ast.AdditionalConstraintLike,
		index_ uint,
		count_ uint,
	)
	PreprocessAdditionalValue(
		additionalValue ast.AdditionalValueLike,
		index_ uint,
		count_ uint,
	)
	ProcessAdditionalValueSlot(
		slot uint,
	)
	PostprocessAdditionalValue(
		additionalValue ast.AdditionalValueLike,
		index_ uint,
		count_ uint,
	)
	PreprocessArgument(
		argument ast.ArgumentLike,
		index_ uint,
		count_ uint,
	)
	ProcessArgumentSlot(
		slot uint,
	)
	PostprocessArgument(
		argument ast.ArgumentLike,
		index_ uint,
		count_ uint,
	)
	PreprocessArguments(
		arguments ast.ArgumentsLike,
		index_ uint,
		count_ uint,
	)
	ProcessArgumentsSlot(
		slot uint,
	)
	PostprocessArguments(
		arguments ast.ArgumentsLike,
		index_ uint,
		count_ uint,
	)
	PreprocessArray(
		array ast.ArrayLike,
		index_ uint,
		count_ uint,
	)
	ProcessArraySlot(
		slot uint,
	)
	PostprocessArray(
		array ast.ArrayLike,
		index_ uint,
		count_ uint,
	)
	PreprocessAspectDeclaration(
		aspectDeclaration ast.AspectDeclarationLike,
		index_ uint,
		count_ uint,
	)
	ProcessAspectDeclarationSlot(
		slot uint,
	)
	PostprocessAspectDeclaration(
		aspectDeclaration ast.AspectDeclarationLike,
		index_ uint,
		count_ uint,
	)
	PreprocessAspectInterface(
		aspectInterface ast.AspectInterfaceLike,
		index_ uint,
		count_ uint,
	)
	ProcessAspectInterfaceSlot(
		slot uint,
	)
	PostprocessAspectInterface(
		aspectInterface ast.AspectInterfaceLike,
		index_ uint,
		count_ uint,
	)
	PreprocessAspectMethod(
		aspectMethod ast.AspectMethodLike,
		index_ uint,
		count_ uint,
	)
	ProcessAspectMethodSlot(
		slot uint,
	)
	PostprocessAspectMethod(
		aspectMethod ast.AspectMethodLike,
		index_ uint,
		count_ uint,
	)
	PreprocessAspectSection(
		aspectSection ast.AspectSectionLike,
		index_ uint,
		count_ uint,
	)
	ProcessAspectSectionSlot(
		slot uint,
	)
	PostprocessAspectSection(
		aspectSection ast.AspectSectionLike,
		index_ uint,
		count_ uint,
	)
	PreprocessAspectSubsection(
		aspectSubsection ast.AspectSubsectionLike,
		index_ uint,
		count_ uint,
	)
	ProcessAspectSubsectionSlot(
		slot uint,
	)
	PostprocessAspectSubsection(
		aspectSubsection ast.AspectSubsectionLike,
		index_ uint,
		count_ uint,
	)
	PreprocessAttributeMethod(
		attributeMethod ast.AttributeMethodLike,
		index_ uint,
		count_ uint,
	)
	ProcessAttributeMethodSlot(
		slot uint,
	)
	PostprocessAttributeMethod(
		attributeMethod ast.AttributeMethodLike,
		index_ uint,
		count_ uint,
	)
	PreprocessAttributeSubsection(
		attributeSubsection ast.AttributeSubsectionLike,
		index_ uint,
		count_ uint,
	)
	ProcessAttributeSubsectionSlot(
		slot uint,
	)
	PostprocessAttributeSubsection(
		attributeSubsection ast.AttributeSubsectionLike,
		index_ uint,
		count_ uint,
	)
	PreprocessChannel(
		channel ast.ChannelLike,
		index_ uint,
		count_ uint,
	)
	ProcessChannelSlot(
		slot uint,
	)
	PostprocessChannel(
		channel ast.ChannelLike,
		index_ uint,
		count_ uint,
	)
	PreprocessClassDeclaration(
		classDeclaration ast.ClassDeclarationLike,
		index_ uint,
		count_ uint,
	)
	ProcessClassDeclarationSlot(
		slot uint,
	)
	PostprocessClassDeclaration(
		classDeclaration ast.ClassDeclarationLike,
		index_ uint,
		count_ uint,
	)
	PreprocessClassMethods(
		classMethods ast.ClassMethodsLike,
		index_ uint,
		count_ uint,
	)
	ProcessClassMethodsSlot(
		slot uint,
	)
	PostprocessClassMethods(
		classMethods ast.ClassMethodsLike,
		index_ uint,
		count_ uint,
	)
	PreprocessClassSection(
		classSection ast.ClassSectionLike,
		index_ uint,
		count_ uint,
	)
	ProcessClassSectionSlot(
		slot uint,
	)
	PostprocessClassSection(
		classSection ast.ClassSectionLike,
		index_ uint,
		count_ uint,
	)
	PreprocessConstantMethod(
		constantMethod ast.ConstantMethodLike,
		index_ uint,
		count_ uint,
	)
	ProcessConstantMethodSlot(
		slot uint,
	)
	PostprocessConstantMethod(
		constantMethod ast.ConstantMethodLike,
		index_ uint,
		count_ uint,
	)
	PreprocessConstantSubsection(
		constantSubsection ast.ConstantSubsectionLike,
		index_ uint,
		count_ uint,
	)
	ProcessConstantSubsectionSlot(
		slot uint,
	)
	PostprocessConstantSubsection(
		constantSubsection ast.ConstantSubsectionLike,
		index_ uint,
		count_ uint,
	)
	PreprocessConstraint(
		constraint ast.ConstraintLike,
		index_ uint,
		count_ uint,
	)
	ProcessConstraintSlot(
		slot uint,
	)
	PostprocessConstraint(
		constraint ast.ConstraintLike,
		index_ uint,
		count_ uint,
	)
	PreprocessConstraints(
		constraints ast.ConstraintsLike,
		index_ uint,
		count_ uint,
	)
	ProcessConstraintsSlot(
		slot uint,
	)
	PostprocessConstraints(
		constraints ast.ConstraintsLike,
		index_ uint,
		count_ uint,
	)
	PreprocessConstructorMethod(
		constructorMethod ast.ConstructorMethodLike,
		index_ uint,
		count_ uint,
	)
	ProcessConstructorMethodSlot(
		slot uint,
	)
	PostprocessConstructorMethod(
		constructorMethod ast.ConstructorMethodLike,
		index_ uint,
		count_ uint,
	)
	PreprocessConstructorSubsection(
		constructorSubsection ast.ConstructorSubsectionLike,
		index_ uint,
		count_ uint,
	)
	ProcessConstructorSubsectionSlot(
		slot uint,
	)
	PostprocessConstructorSubsection(
		constructorSubsection ast.ConstructorSubsectionLike,
		index_ uint,
		count_ uint,
	)
	PreprocessDeclaration(
		declaration ast.DeclarationLike,
		index_ uint,
		count_ uint,
	)
	ProcessDeclarationSlot(
		slot uint,
	)
	PostprocessDeclaration(
		declaration ast.DeclarationLike,
		index_ uint,
		count_ uint,
	)
	PreprocessEnumeration(
		enumeration ast.EnumerationLike,
		index_ uint,
		count_ uint,
	)
	ProcessEnumerationSlot(
		slot uint,
	)
	PostprocessEnumeration(
		enumeration ast.EnumerationLike,
		index_ uint,
		count_ uint,
	)
	PreprocessFunctionMethod(
		functionMethod ast.FunctionMethodLike,
		index_ uint,
		count_ uint,
	)
	ProcessFunctionMethodSlot(
		slot uint,
	)
	PostprocessFunctionMethod(
		functionMethod ast.FunctionMethodLike,
		index_ uint,
		count_ uint,
	)
	PreprocessFunctionSubsection(
		functionSubsection ast.FunctionSubsectionLike,
		index_ uint,
		count_ uint,
	)
	ProcessFunctionSubsectionSlot(
		slot uint,
	)
	PostprocessFunctionSubsection(
		functionSubsection ast.FunctionSubsectionLike,
		index_ uint,
		count_ uint,
	)
	PreprocessFunctionalDeclaration(
		functionalDeclaration ast.FunctionalDeclarationLike,
		index_ uint,
		count_ uint,
	)
	ProcessFunctionalDeclarationSlot(
		slot uint,
	)
	PostprocessFunctionalDeclaration(
		functionalDeclaration ast.FunctionalDeclarationLike,
		index_ uint,
		count_ uint,
	)
	PreprocessFunctionalSection(
		functionalSection ast.FunctionalSectionLike,
		index_ uint,
		count_ uint,
	)
	ProcessFunctionalSectionSlot(
		slot uint,
	)
	PostprocessFunctionalSection(
		functionalSection ast.FunctionalSectionLike,
		index_ uint,
		count_ uint,
	)
	PreprocessGetterMethod(
		getterMethod ast.GetterMethodLike,
		index_ uint,
		count_ uint,
	)
	ProcessGetterMethodSlot(
		slot uint,
	)
	PostprocessGetterMethod(
		getterMethod ast.GetterMethodLike,
		index_ uint,
		count_ uint,
	)
	PreprocessImportList(
		importList ast.ImportListLike,
		index_ uint,
		count_ uint,
	)
	ProcessImportListSlot(
		slot uint,
	)
	PostprocessImportList(
		importList ast.ImportListLike,
		index_ uint,
		count_ uint,
	)
	PreprocessImportedPackage(
		importedPackage ast.ImportedPackageLike,
		index_ uint,
		count_ uint,
	)
	ProcessImportedPackageSlot(
		slot uint,
	)
	PostprocessImportedPackage(
		importedPackage ast.ImportedPackageLike,
		index_ uint,
		count_ uint,
	)
	PreprocessInstanceDeclaration(
		instanceDeclaration ast.InstanceDeclarationLike,
		index_ uint,
		count_ uint,
	)
	ProcessInstanceDeclarationSlot(
		slot uint,
	)
	PostprocessInstanceDeclaration(
		instanceDeclaration ast.InstanceDeclarationLike,
		index_ uint,
		count_ uint,
	)
	PreprocessInstanceMethods(
		instanceMethods ast.InstanceMethodsLike,
		index_ uint,
		count_ uint,
	)
	ProcessInstanceMethodsSlot(
		slot uint,
	)
	PostprocessInstanceMethods(
		instanceMethods ast.InstanceMethodsLike,
		index_ uint,
		count_ uint,
	)
	PreprocessInstanceSection(
		instanceSection ast.InstanceSectionLike,
		index_ uint,
		count_ uint,
	)
	ProcessInstanceSectionSlot(
		slot uint,
	)
	PostprocessInstanceSection(
		instanceSection ast.InstanceSectionLike,
		index_ uint,
		count_ uint,
	)
	PreprocessInterfaceDeclarations(
		interfaceDeclarations ast.InterfaceDeclarationsLike,
		index_ uint,
		count_ uint,
	)
	ProcessInterfaceDeclarationsSlot(
		slot uint,
	)
	PostprocessInterfaceDeclarations(
		interfaceDeclarations ast.InterfaceDeclarationsLike,
		index_ uint,
		count_ uint,
	)
	PreprocessLegalNotice(
		legalNotice ast.LegalNoticeLike,
		index_ uint,
		count_ uint,
	)
	ProcessLegalNoticeSlot(
		slot uint,
	)
	PostprocessLegalNotice(
		legalNotice ast.LegalNoticeLike,
		index_ uint,
		count_ uint,
	)
	PreprocessMap(
		map_ ast.MapLike,
		index_ uint,
		count_ uint,
	)
	ProcessMapSlot(
		slot uint,
	)
	PostprocessMap(
		map_ ast.MapLike,
		index_ uint,
		count_ uint,
	)
	PreprocessMethod(
		method ast.MethodLike,
		index_ uint,
		count_ uint,
	)
	ProcessMethodSlot(
		slot uint,
	)
	PostprocessMethod(
		method ast.MethodLike,
		index_ uint,
		count_ uint,
	)
	PreprocessModel(
		model ast.ModelLike,
		index_ uint,
		count_ uint,
	)
	ProcessModelSlot(
		slot uint,
	)
	PostprocessModel(
		model ast.ModelLike,
		index_ uint,
		count_ uint,
	)
	PreprocessMultivalue(
		multivalue ast.MultivalueLike,
		index_ uint,
		count_ uint,
	)
	ProcessMultivalueSlot(
		slot uint,
	)
	PostprocessMultivalue(
		multivalue ast.MultivalueLike,
		index_ uint,
		count_ uint,
	)
	PreprocessNone(
		none ast.NoneLike,
		index_ uint,
		count_ uint,
	)
	ProcessNoneSlot(
		slot uint,
	)
	PostprocessNone(
		none ast.NoneLike,
		index_ uint,
		count_ uint,
	)
	PreprocessPackageDeclaration(
		packageDeclaration ast.PackageDeclarationLike,
		index_ uint,
		count_ uint,
	)
	ProcessPackageDeclarationSlot(
		slot uint,
	)
	PostprocessPackageDeclaration(
		packageDeclaration ast.PackageDeclarationLike,
		index_ uint,
		count_ uint,
	)
	PreprocessPackageHeader(
		packageHeader ast.PackageHeaderLike,
		index_ uint,
		count_ uint,
	)
	ProcessPackageHeaderSlot(
		slot uint,
	)
	PostprocessPackageHeader(
		packageHeader ast.PackageHeaderLike,
		index_ uint,
		count_ uint,
	)
	PreprocessPackageImports(
		packageImports ast.PackageImportsLike,
		index_ uint,
		count_ uint,
	)
	ProcessPackageImportsSlot(
		slot uint,
	)
	PostprocessPackageImports(
		packageImports ast.PackageImportsLike,
		index_ uint,
		count_ uint,
	)
	PreprocessParameter(
		parameter ast.ParameterLike,
		index_ uint,
		count_ uint,
	)
	ProcessParameterSlot(
		slot uint,
	)
	PostprocessParameter(
		parameter ast.ParameterLike,
		index_ uint,
		count_ uint,
	)
	PreprocessParameterList(
		parameterList ast.ParameterListLike,
		index_ uint,
		count_ uint,
	)
	ProcessParameterListSlot(
		slot uint,
	)
	PostprocessParameterList(
		parameterList ast.ParameterListLike,
		index_ uint,
		count_ uint,
	)
	PreprocessPrimitiveDeclarations(
		primitiveDeclarations ast.PrimitiveDeclarationsLike,
		index_ uint,
		count_ uint,
	)
	ProcessPrimitiveDeclarationsSlot(
		slot uint,
	)
	PostprocessPrimitiveDeclarations(
		primitiveDeclarations ast.PrimitiveDeclarationsLike,
		index_ uint,
		count_ uint,
	)
	PreprocessPrincipalMethod(
		principalMethod ast.PrincipalMethodLike,
		index_ uint,
		count_ uint,
	)
	ProcessPrincipalMethodSlot(
		slot uint,
	)
	PostprocessPrincipalMethod(
		principalMethod ast.PrincipalMethodLike,
		index_ uint,
		count_ uint,
	)
	PreprocessPrincipalSubsection(
		principalSubsection ast.PrincipalSubsectionLike,
		index_ uint,
		count_ uint,
	)
	ProcessPrincipalSubsectionSlot(
		slot uint,
	)
	PostprocessPrincipalSubsection(
		principalSubsection ast.PrincipalSubsectionLike,
		index_ uint,
		count_ uint,
	)
	PreprocessResult(
		result ast.ResultLike,
		index_ uint,
		count_ uint,
	)
	ProcessResultSlot(
		slot uint,
	)
	PostprocessResult(
		result ast.ResultLike,
		index_ uint,
		count_ uint,
	)
	PreprocessSetterMethod(
		setterMethod ast.SetterMethodLike,
		index_ uint,
		count_ uint,
	)
	ProcessSetterMethodSlot(
		slot uint,
	)
	PostprocessSetterMethod(
		setterMethod ast.SetterMethodLike,
		index_ uint,
		count_ uint,
	)
	PreprocessTypeDeclaration(
		typeDeclaration ast.TypeDeclarationLike,
		index_ uint,
		count_ uint,
	)
	ProcessTypeDeclarationSlot(
		slot uint,
	)
	PostprocessTypeDeclaration(
		typeDeclaration ast.TypeDeclarationLike,
		index_ uint,
		count_ uint,
	)
	PreprocessTypeSection(
		typeSection ast.TypeSectionLike,
		index_ uint,
		count_ uint,
	)
	ProcessTypeSectionSlot(
		slot uint,
	)
	PostprocessTypeSection(
		typeSection ast.TypeSectionLike,
		index_ uint,
		count_ uint,
	)
	PreprocessValue(
		value ast.ValueLike,
		index_ uint,
		count_ uint,
	)
	ProcessValueSlot(
		slot uint,
	)
	PostprocessValue(
		value ast.ValueLike,
		index_ uint,
		count_ uint,
	)
	PreprocessWrapper(
		wrapper ast.WrapperLike,
		index_ uint,
		count_ uint,
	)
	ProcessWrapperSlot(
		slot uint,
	)
	PostprocessWrapper(
		wrapper ast.WrapperLike,
		index_ uint,
		count_ uint,
	)
}
