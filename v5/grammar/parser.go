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
│                 This class file was automatically generated.                 │
│                     Any updates to it may be overwritten.                    │
└──────────────────────────────────────────────────────────────────────────────┘
*/

package grammar

import (
	fmt "fmt"
	ast "github.com/craterdog/go-class-model/v5/ast"
	col "github.com/craterdog/go-collection-framework/v5"
	abs "github.com/craterdog/go-collection-framework/v5/collection"
	uti "github.com/craterdog/go-missing-utilities/v2"
	mat "math"
	sts "strings"
)

// CLASS INTERFACE

// Access Function

func ParserClass() ParserClassLike {
	return parserClass()
}

// Constructor Methods

func (c *parserClass_) Parser() ParserLike {
	var instance = &parser_{
		// Initialize the instance attributes.
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *parser_) GetClass() ParserClassLike {
	return parserClass()
}

func (v *parser_) ParseSource(
	source string,
) ast.ModelLike {
	v.source_ = sts.ReplaceAll(source, "\t", "    ")
	v.tokens_ = col.Queue[TokenLike]()
	v.next_ = col.Stack[TokenLike]()

	// The scanner runs in a separate Go routine.
	ScannerClass().Scanner(v.source_, v.tokens_)

	// Attempt to parse the model.
	var model, token, ok = v.parseModel()
	if !ok {
		var message = v.formatError("$Model", token)
		panic(message)
	}
	return model
}

// PROTECTED INTERFACE

// Private Methods

func (v *parser_) parseAbstraction() (
	abstraction ast.AbstractionLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse an optional Wrapper rule.
	var optionalWrapper ast.WrapperLike
	optionalWrapper, _, ok = v.parseWrapper()
	if ok {
		// No additional put backs allowed at this point.
		tokens = nil
	}

	// Attempt to parse an optional prefix token.
	var optionalPrefix string
	optionalPrefix, token, ok = v.parseToken(PrefixToken)
	if ok && uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single name token.
	var name string
	name, token, ok = v.parseToken(NameToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Abstraction rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Abstraction", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse an optional Arguments rule.
	var optionalArguments ast.ArgumentsLike
	optionalArguments, _, ok = v.parseArguments()
	if ok {
		// No additional put backs allowed at this point.
		tokens = nil
	}

	// Found a single Abstraction rule.
	ok = true
	v.remove(tokens)
	abstraction = ast.AbstractionClass().Abstraction(
		optionalWrapper,
		optionalPrefix,
		name,
		optionalArguments,
	)
	return
}

func (v *parser_) parseAdditionalArgument() (
	additionalArgument ast.AdditionalArgumentLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "," delimiter.
	_, token, ok = v.parseDelimiter(",")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single AdditionalArgument rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$AdditionalArgument", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Argument rule.
	var argument ast.ArgumentLike
	argument, token, ok = v.parseArgument()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single AdditionalArgument rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$AdditionalArgument", token)
		panic(message)
	}

	// Found a single AdditionalArgument rule.
	ok = true
	v.remove(tokens)
	additionalArgument = ast.AdditionalArgumentClass().AdditionalArgument(argument)
	return
}

func (v *parser_) parseAdditionalConstraint() (
	additionalConstraint ast.AdditionalConstraintLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "," delimiter.
	_, token, ok = v.parseDelimiter(",")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single AdditionalConstraint rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$AdditionalConstraint", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Constraint rule.
	var constraint ast.ConstraintLike
	constraint, token, ok = v.parseConstraint()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single AdditionalConstraint rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$AdditionalConstraint", token)
		panic(message)
	}

	// Found a single AdditionalConstraint rule.
	ok = true
	v.remove(tokens)
	additionalConstraint = ast.AdditionalConstraintClass().AdditionalConstraint(constraint)
	return
}

func (v *parser_) parseAdditionalValue() (
	additionalValue ast.AdditionalValueLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single name token.
	var name string
	name, token, ok = v.parseToken(NameToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single AdditionalValue rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$AdditionalValue", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single AdditionalValue rule.
	ok = true
	v.remove(tokens)
	additionalValue = ast.AdditionalValueClass().AdditionalValue(name)
	return
}

func (v *parser_) parseArgument() (
	argument ast.ArgumentLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single Abstraction rule.
	var abstraction ast.AbstractionLike
	abstraction, token, ok = v.parseAbstraction()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Argument rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Argument", token)
		panic(message)
	}

	// Found a single Argument rule.
	ok = true
	v.remove(tokens)
	argument = ast.ArgumentClass().Argument(abstraction)
	return
}

func (v *parser_) parseArguments() (
	arguments ast.ArgumentsLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "[" delimiter.
	_, token, ok = v.parseDelimiter("[")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Arguments rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Arguments", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Argument rule.
	var argument ast.ArgumentLike
	argument, token, ok = v.parseArgument()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Arguments rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Arguments", token)
		panic(message)
	}

	// Attempt to parse multiple AdditionalArgument rules.
	var additionalArguments = col.List[ast.AdditionalArgumentLike]()
additionalArgumentsLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var additionalArgument ast.AdditionalArgumentLike
		additionalArgument, token, ok = v.parseAdditionalArgument()
		if !ok {
			switch {
			case count >= 0:
				break additionalArgumentsLoop
			case uti.IsDefined(tokens):
				// This is not multiple AdditionalArgument rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$Arguments", token)
				message += "0 or more AdditionalArgument rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		additionalArguments.AppendValue(additionalArgument)
	}

	// Attempt to parse a single "]" delimiter.
	_, token, ok = v.parseDelimiter("]")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Arguments rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Arguments", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Arguments rule.
	ok = true
	v.remove(tokens)
	arguments = ast.ArgumentsClass().Arguments(
		argument,
		additionalArguments,
	)
	return
}

func (v *parser_) parseArray() (
	array ast.ArrayLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "[" delimiter.
	_, token, ok = v.parseDelimiter("[")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Array rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Array", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single "]" delimiter.
	_, token, ok = v.parseDelimiter("]")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Array rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Array", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Array rule.
	ok = true
	v.remove(tokens)
	array = ast.ArrayClass().Array()
	return
}

func (v *parser_) parseAspectDeclaration() (
	aspectDeclaration ast.AspectDeclarationLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single Declaration rule.
	var declaration ast.DeclarationLike
	declaration, token, ok = v.parseDeclaration()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single AspectDeclaration rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$AspectDeclaration", token)
		panic(message)
	}

	// Attempt to parse a single "interface" delimiter.
	_, token, ok = v.parseDelimiter("interface")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single AspectDeclaration rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$AspectDeclaration", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single "{" delimiter.
	_, token, ok = v.parseDelimiter("{")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single AspectDeclaration rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$AspectDeclaration", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse multiple AspectMethod rules.
	var aspectMethods = col.List[ast.AspectMethodLike]()
aspectMethodsLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var aspectMethod ast.AspectMethodLike
		aspectMethod, token, ok = v.parseAspectMethod()
		if !ok {
			switch {
			case count >= 1:
				break aspectMethodsLoop
			case uti.IsDefined(tokens):
				// This is not multiple AspectMethod rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$AspectDeclaration", token)
				message += "1 or more AspectMethod rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		aspectMethods.AppendValue(aspectMethod)
	}

	// Attempt to parse a single "}" delimiter.
	_, token, ok = v.parseDelimiter("}")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single AspectDeclaration rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$AspectDeclaration", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single AspectDeclaration rule.
	ok = true
	v.remove(tokens)
	aspectDeclaration = ast.AspectDeclarationClass().AspectDeclaration(
		declaration,
		aspectMethods,
	)
	return
}

func (v *parser_) parseAspectInterface() (
	aspectInterface ast.AspectInterfaceLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single Abstraction rule.
	var abstraction ast.AbstractionLike
	abstraction, token, ok = v.parseAbstraction()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single AspectInterface rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$AspectInterface", token)
		panic(message)
	}

	// Found a single AspectInterface rule.
	ok = true
	v.remove(tokens)
	aspectInterface = ast.AspectInterfaceClass().AspectInterface(abstraction)
	return
}

func (v *parser_) parseAspectMethod() (
	aspectMethod ast.AspectMethodLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single Method rule.
	var method ast.MethodLike
	method, token, ok = v.parseMethod()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single AspectMethod rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$AspectMethod", token)
		panic(message)
	}

	// Found a single AspectMethod rule.
	ok = true
	v.remove(tokens)
	aspectMethod = ast.AspectMethodClass().AspectMethod(method)
	return
}

func (v *parser_) parseAspectSection() (
	aspectSection ast.AspectSectionLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "// ASPECT DECLARATIONS" delimiter.
	_, token, ok = v.parseDelimiter("// ASPECT DECLARATIONS")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single AspectSection rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$AspectSection", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse multiple AspectDeclaration rules.
	var aspectDeclarations = col.List[ast.AspectDeclarationLike]()
aspectDeclarationsLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var aspectDeclaration ast.AspectDeclarationLike
		aspectDeclaration, token, ok = v.parseAspectDeclaration()
		if !ok {
			switch {
			case count >= 0:
				break aspectDeclarationsLoop
			case uti.IsDefined(tokens):
				// This is not multiple AspectDeclaration rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$AspectSection", token)
				message += "0 or more AspectDeclaration rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		aspectDeclarations.AppendValue(aspectDeclaration)
	}

	// Found a single AspectSection rule.
	ok = true
	v.remove(tokens)
	aspectSection = ast.AspectSectionClass().AspectSection(aspectDeclarations)
	return
}

func (v *parser_) parseAspectSubsection() (
	aspectSubsection ast.AspectSubsectionLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "// Aspect Interfaces" delimiter.
	_, token, ok = v.parseDelimiter("// Aspect Interfaces")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single AspectSubsection rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$AspectSubsection", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse multiple AspectInterface rules.
	var aspectInterfaces = col.List[ast.AspectInterfaceLike]()
aspectInterfacesLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var aspectInterface ast.AspectInterfaceLike
		aspectInterface, token, ok = v.parseAspectInterface()
		if !ok {
			switch {
			case count >= 1:
				break aspectInterfacesLoop
			case uti.IsDefined(tokens):
				// This is not multiple AspectInterface rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$AspectSubsection", token)
				message += "1 or more AspectInterface rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		aspectInterfaces.AppendValue(aspectInterface)
	}

	// Found a single AspectSubsection rule.
	ok = true
	v.remove(tokens)
	aspectSubsection = ast.AspectSubsectionClass().AspectSubsection(aspectInterfaces)
	return
}

func (v *parser_) parseAttributeMethod() (
	attributeMethod ast.AttributeMethodLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single GetterMethod AttributeMethod.
	var getterMethod ast.GetterMethodLike
	getterMethod, token, ok = v.parseGetterMethod()
	if ok {
		// Found a single GetterMethod AttributeMethod.
		attributeMethod = ast.AttributeMethodClass().AttributeMethod(getterMethod)
		return
	}

	// Attempt to parse a single SetterMethod AttributeMethod.
	var setterMethod ast.SetterMethodLike
	setterMethod, token, ok = v.parseSetterMethod()
	if ok {
		// Found a single SetterMethod AttributeMethod.
		attributeMethod = ast.AttributeMethodClass().AttributeMethod(setterMethod)
		return
	}

	// This is not a single AttributeMethod rule.
	return
}

func (v *parser_) parseAttributeSubsection() (
	attributeSubsection ast.AttributeSubsectionLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "// Attribute Methods" delimiter.
	_, token, ok = v.parseDelimiter("// Attribute Methods")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single AttributeSubsection rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$AttributeSubsection", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse multiple AttributeMethod rules.
	var attributeMethods = col.List[ast.AttributeMethodLike]()
attributeMethodsLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var attributeMethod ast.AttributeMethodLike
		attributeMethod, token, ok = v.parseAttributeMethod()
		if !ok {
			switch {
			case count >= 1:
				break attributeMethodsLoop
			case uti.IsDefined(tokens):
				// This is not multiple AttributeMethod rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$AttributeSubsection", token)
				message += "1 or more AttributeMethod rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		attributeMethods.AppendValue(attributeMethod)
	}

	// Found a single AttributeSubsection rule.
	ok = true
	v.remove(tokens)
	attributeSubsection = ast.AttributeSubsectionClass().AttributeSubsection(attributeMethods)
	return
}

func (v *parser_) parseChannel() (
	channel ast.ChannelLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "chan" delimiter.
	_, token, ok = v.parseDelimiter("chan")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Channel rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Channel", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Channel rule.
	ok = true
	v.remove(tokens)
	channel = ast.ChannelClass().Channel()
	return
}

func (v *parser_) parseClassDeclaration() (
	classDeclaration ast.ClassDeclarationLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single Declaration rule.
	var declaration ast.DeclarationLike
	declaration, token, ok = v.parseDeclaration()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single ClassDeclaration rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$ClassDeclaration", token)
		panic(message)
	}

	// Attempt to parse a single "interface" delimiter.
	_, token, ok = v.parseDelimiter("interface")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ClassDeclaration rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ClassDeclaration", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single "{" delimiter.
	_, token, ok = v.parseDelimiter("{")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ClassDeclaration rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ClassDeclaration", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ClassMethods rule.
	var classMethods ast.ClassMethodsLike
	classMethods, token, ok = v.parseClassMethods()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single ClassDeclaration rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$ClassDeclaration", token)
		panic(message)
	}

	// Attempt to parse a single "}" delimiter.
	_, token, ok = v.parseDelimiter("}")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ClassDeclaration rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ClassDeclaration", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single ClassDeclaration rule.
	ok = true
	v.remove(tokens)
	classDeclaration = ast.ClassDeclarationClass().ClassDeclaration(
		declaration,
		classMethods,
	)
	return
}

func (v *parser_) parseClassMethods() (
	classMethods ast.ClassMethodsLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single ConstructorSubsection rule.
	var constructorSubsection ast.ConstructorSubsectionLike
	constructorSubsection, token, ok = v.parseConstructorSubsection()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single ClassMethods rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$ClassMethods", token)
		panic(message)
	}

	// Attempt to parse an optional ConstantSubsection rule.
	var optionalConstantSubsection ast.ConstantSubsectionLike
	optionalConstantSubsection, _, ok = v.parseConstantSubsection()
	if ok {
		// No additional put backs allowed at this point.
		tokens = nil
	}

	// Attempt to parse an optional FunctionSubsection rule.
	var optionalFunctionSubsection ast.FunctionSubsectionLike
	optionalFunctionSubsection, _, ok = v.parseFunctionSubsection()
	if ok {
		// No additional put backs allowed at this point.
		tokens = nil
	}

	// Found a single ClassMethods rule.
	ok = true
	v.remove(tokens)
	classMethods = ast.ClassMethodsClass().ClassMethods(
		constructorSubsection,
		optionalConstantSubsection,
		optionalFunctionSubsection,
	)
	return
}

func (v *parser_) parseClassSection() (
	classSection ast.ClassSectionLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "// CLASS DECLARATIONS" delimiter.
	_, token, ok = v.parseDelimiter("// CLASS DECLARATIONS")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ClassSection rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ClassSection", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse multiple ClassDeclaration rules.
	var classDeclarations = col.List[ast.ClassDeclarationLike]()
classDeclarationsLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var classDeclaration ast.ClassDeclarationLike
		classDeclaration, token, ok = v.parseClassDeclaration()
		if !ok {
			switch {
			case count >= 1:
				break classDeclarationsLoop
			case uti.IsDefined(tokens):
				// This is not multiple ClassDeclaration rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$ClassSection", token)
				message += "1 or more ClassDeclaration rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		classDeclarations.AppendValue(classDeclaration)
	}

	// Found a single ClassSection rule.
	ok = true
	v.remove(tokens)
	classSection = ast.ClassSectionClass().ClassSection(classDeclarations)
	return
}

func (v *parser_) parseConstantMethod() (
	constantMethod ast.ConstantMethodLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single name token.
	var name string
	name, token, ok = v.parseToken(NameToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ConstantMethod rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ConstantMethod", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single "(" delimiter.
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ConstantMethod rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ConstantMethod", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ")" delimiter.
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ConstantMethod rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ConstantMethod", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Abstraction rule.
	var abstraction ast.AbstractionLike
	abstraction, token, ok = v.parseAbstraction()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single ConstantMethod rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$ConstantMethod", token)
		panic(message)
	}

	// Found a single ConstantMethod rule.
	ok = true
	v.remove(tokens)
	constantMethod = ast.ConstantMethodClass().ConstantMethod(
		name,
		abstraction,
	)
	return
}

func (v *parser_) parseConstantSubsection() (
	constantSubsection ast.ConstantSubsectionLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "// Constant Methods" delimiter.
	_, token, ok = v.parseDelimiter("// Constant Methods")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ConstantSubsection rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ConstantSubsection", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse multiple ConstantMethod rules.
	var constantMethods = col.List[ast.ConstantMethodLike]()
constantMethodsLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var constantMethod ast.ConstantMethodLike
		constantMethod, token, ok = v.parseConstantMethod()
		if !ok {
			switch {
			case count >= 1:
				break constantMethodsLoop
			case uti.IsDefined(tokens):
				// This is not multiple ConstantMethod rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$ConstantSubsection", token)
				message += "1 or more ConstantMethod rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		constantMethods.AppendValue(constantMethod)
	}

	// Found a single ConstantSubsection rule.
	ok = true
	v.remove(tokens)
	constantSubsection = ast.ConstantSubsectionClass().ConstantSubsection(constantMethods)
	return
}

func (v *parser_) parseConstraint() (
	constraint ast.ConstraintLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single name token.
	var name string
	name, token, ok = v.parseToken(NameToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Constraint rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Constraint", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Abstraction rule.
	var abstraction ast.AbstractionLike
	abstraction, token, ok = v.parseAbstraction()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Constraint rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Constraint", token)
		panic(message)
	}

	// Found a single Constraint rule.
	ok = true
	v.remove(tokens)
	constraint = ast.ConstraintClass().Constraint(
		name,
		abstraction,
	)
	return
}

func (v *parser_) parseConstraints() (
	constraints ast.ConstraintsLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "[" delimiter.
	_, token, ok = v.parseDelimiter("[")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Constraints rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Constraints", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Constraint rule.
	var constraint ast.ConstraintLike
	constraint, token, ok = v.parseConstraint()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Constraints rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Constraints", token)
		panic(message)
	}

	// Attempt to parse multiple AdditionalConstraint rules.
	var additionalConstraints = col.List[ast.AdditionalConstraintLike]()
additionalConstraintsLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var additionalConstraint ast.AdditionalConstraintLike
		additionalConstraint, token, ok = v.parseAdditionalConstraint()
		if !ok {
			switch {
			case count >= 0:
				break additionalConstraintsLoop
			case uti.IsDefined(tokens):
				// This is not multiple AdditionalConstraint rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$Constraints", token)
				message += "0 or more AdditionalConstraint rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		additionalConstraints.AppendValue(additionalConstraint)
	}

	// Attempt to parse a single "]" delimiter.
	_, token, ok = v.parseDelimiter("]")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Constraints rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Constraints", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Constraints rule.
	ok = true
	v.remove(tokens)
	constraints = ast.ConstraintsClass().Constraints(
		constraint,
		additionalConstraints,
	)
	return
}

func (v *parser_) parseConstructorMethod() (
	constructorMethod ast.ConstructorMethodLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single name token.
	var name string
	name, token, ok = v.parseToken(NameToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ConstructorMethod rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ConstructorMethod", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single "(" delimiter.
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ConstructorMethod rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ConstructorMethod", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse multiple Parameter rules.
	var parameters = col.List[ast.ParameterLike]()
parametersLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var parameter ast.ParameterLike
		parameter, token, ok = v.parseParameter()
		if !ok {
			switch {
			case count >= 0:
				break parametersLoop
			case uti.IsDefined(tokens):
				// This is not multiple Parameter rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$ConstructorMethod", token)
				message += "0 or more Parameter rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		parameters.AppendValue(parameter)
	}

	// Attempt to parse a single ")" delimiter.
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ConstructorMethod rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ConstructorMethod", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Abstraction rule.
	var abstraction ast.AbstractionLike
	abstraction, token, ok = v.parseAbstraction()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single ConstructorMethod rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$ConstructorMethod", token)
		panic(message)
	}

	// Found a single ConstructorMethod rule.
	ok = true
	v.remove(tokens)
	constructorMethod = ast.ConstructorMethodClass().ConstructorMethod(
		name,
		parameters,
		abstraction,
	)
	return
}

func (v *parser_) parseConstructorSubsection() (
	constructorSubsection ast.ConstructorSubsectionLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "// Constructor Methods" delimiter.
	_, token, ok = v.parseDelimiter("// Constructor Methods")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ConstructorSubsection rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ConstructorSubsection", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse multiple ConstructorMethod rules.
	var constructorMethods = col.List[ast.ConstructorMethodLike]()
constructorMethodsLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var constructorMethod ast.ConstructorMethodLike
		constructorMethod, token, ok = v.parseConstructorMethod()
		if !ok {
			switch {
			case count >= 1:
				break constructorMethodsLoop
			case uti.IsDefined(tokens):
				// This is not multiple ConstructorMethod rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$ConstructorSubsection", token)
				message += "1 or more ConstructorMethod rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		constructorMethods.AppendValue(constructorMethod)
	}

	// Found a single ConstructorSubsection rule.
	ok = true
	v.remove(tokens)
	constructorSubsection = ast.ConstructorSubsectionClass().ConstructorSubsection(constructorMethods)
	return
}

func (v *parser_) parseDeclaration() (
	declaration ast.DeclarationLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single comment token.
	var comment string
	comment, token, ok = v.parseToken(CommentToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Declaration rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Declaration", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single "type" delimiter.
	_, token, ok = v.parseDelimiter("type")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Declaration rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Declaration", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single name token.
	var name string
	name, token, ok = v.parseToken(NameToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Declaration rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Declaration", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse an optional Constraints rule.
	var optionalConstraints ast.ConstraintsLike
	optionalConstraints, _, ok = v.parseConstraints()
	if ok {
		// No additional put backs allowed at this point.
		tokens = nil
	}

	// Found a single Declaration rule.
	ok = true
	v.remove(tokens)
	declaration = ast.DeclarationClass().Declaration(
		comment,
		name,
		optionalConstraints,
	)
	return
}

func (v *parser_) parseEnumeration() (
	enumeration ast.EnumerationLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "const" delimiter.
	_, token, ok = v.parseDelimiter("const")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Enumeration rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Enumeration", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single "(" delimiter.
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Enumeration rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Enumeration", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Value rule.
	var value ast.ValueLike
	value, token, ok = v.parseValue()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Enumeration rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Enumeration", token)
		panic(message)
	}

	// Attempt to parse multiple AdditionalValue rules.
	var additionalValues = col.List[ast.AdditionalValueLike]()
additionalValuesLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var additionalValue ast.AdditionalValueLike
		additionalValue, token, ok = v.parseAdditionalValue()
		if !ok {
			switch {
			case count >= 0:
				break additionalValuesLoop
			case uti.IsDefined(tokens):
				// This is not multiple AdditionalValue rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$Enumeration", token)
				message += "0 or more AdditionalValue rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		additionalValues.AppendValue(additionalValue)
	}

	// Attempt to parse a single ")" delimiter.
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Enumeration rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Enumeration", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Enumeration rule.
	ok = true
	v.remove(tokens)
	enumeration = ast.EnumerationClass().Enumeration(
		value,
		additionalValues,
	)
	return
}

func (v *parser_) parseFunctionMethod() (
	functionMethod ast.FunctionMethodLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single name token.
	var name string
	name, token, ok = v.parseToken(NameToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single FunctionMethod rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$FunctionMethod", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single "(" delimiter.
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single FunctionMethod rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$FunctionMethod", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse multiple Parameter rules.
	var parameters = col.List[ast.ParameterLike]()
parametersLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var parameter ast.ParameterLike
		parameter, token, ok = v.parseParameter()
		if !ok {
			switch {
			case count >= 0:
				break parametersLoop
			case uti.IsDefined(tokens):
				// This is not multiple Parameter rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$FunctionMethod", token)
				message += "0 or more Parameter rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		parameters.AppendValue(parameter)
	}

	// Attempt to parse a single ")" delimiter.
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single FunctionMethod rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$FunctionMethod", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Result rule.
	var result ast.ResultLike
	result, token, ok = v.parseResult()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single FunctionMethod rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$FunctionMethod", token)
		panic(message)
	}

	// Found a single FunctionMethod rule.
	ok = true
	v.remove(tokens)
	functionMethod = ast.FunctionMethodClass().FunctionMethod(
		name,
		parameters,
		result,
	)
	return
}

func (v *parser_) parseFunctionSubsection() (
	functionSubsection ast.FunctionSubsectionLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "// Function Methods" delimiter.
	_, token, ok = v.parseDelimiter("// Function Methods")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single FunctionSubsection rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$FunctionSubsection", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse multiple FunctionMethod rules.
	var functionMethods = col.List[ast.FunctionMethodLike]()
functionMethodsLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var functionMethod ast.FunctionMethodLike
		functionMethod, token, ok = v.parseFunctionMethod()
		if !ok {
			switch {
			case count >= 1:
				break functionMethodsLoop
			case uti.IsDefined(tokens):
				// This is not multiple FunctionMethod rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$FunctionSubsection", token)
				message += "1 or more FunctionMethod rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		functionMethods.AppendValue(functionMethod)
	}

	// Found a single FunctionSubsection rule.
	ok = true
	v.remove(tokens)
	functionSubsection = ast.FunctionSubsectionClass().FunctionSubsection(functionMethods)
	return
}

func (v *parser_) parseFunctionalDeclaration() (
	functionalDeclaration ast.FunctionalDeclarationLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single Declaration rule.
	var declaration ast.DeclarationLike
	declaration, token, ok = v.parseDeclaration()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single FunctionalDeclaration rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$FunctionalDeclaration", token)
		panic(message)
	}

	// Attempt to parse a single "func" delimiter.
	_, token, ok = v.parseDelimiter("func")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single FunctionalDeclaration rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$FunctionalDeclaration", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single "(" delimiter.
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single FunctionalDeclaration rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$FunctionalDeclaration", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse multiple Parameter rules.
	var parameters = col.List[ast.ParameterLike]()
parametersLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var parameter ast.ParameterLike
		parameter, token, ok = v.parseParameter()
		if !ok {
			switch {
			case count >= 0:
				break parametersLoop
			case uti.IsDefined(tokens):
				// This is not multiple Parameter rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$FunctionalDeclaration", token)
				message += "0 or more Parameter rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		parameters.AppendValue(parameter)
	}

	// Attempt to parse a single ")" delimiter.
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single FunctionalDeclaration rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$FunctionalDeclaration", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Result rule.
	var result ast.ResultLike
	result, token, ok = v.parseResult()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single FunctionalDeclaration rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$FunctionalDeclaration", token)
		panic(message)
	}

	// Found a single FunctionalDeclaration rule.
	ok = true
	v.remove(tokens)
	functionalDeclaration = ast.FunctionalDeclarationClass().FunctionalDeclaration(
		declaration,
		parameters,
		result,
	)
	return
}

func (v *parser_) parseFunctionalSection() (
	functionalSection ast.FunctionalSectionLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "// FUNCTIONAL DECLARATIONS" delimiter.
	_, token, ok = v.parseDelimiter("// FUNCTIONAL DECLARATIONS")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single FunctionalSection rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$FunctionalSection", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse multiple FunctionalDeclaration rules.
	var functionalDeclarations = col.List[ast.FunctionalDeclarationLike]()
functionalDeclarationsLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var functionalDeclaration ast.FunctionalDeclarationLike
		functionalDeclaration, token, ok = v.parseFunctionalDeclaration()
		if !ok {
			switch {
			case count >= 0:
				break functionalDeclarationsLoop
			case uti.IsDefined(tokens):
				// This is not multiple FunctionalDeclaration rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$FunctionalSection", token)
				message += "0 or more FunctionalDeclaration rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		functionalDeclarations.AppendValue(functionalDeclaration)
	}

	// Found a single FunctionalSection rule.
	ok = true
	v.remove(tokens)
	functionalSection = ast.FunctionalSectionClass().FunctionalSection(functionalDeclarations)
	return
}

func (v *parser_) parseGetterMethod() (
	getterMethod ast.GetterMethodLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single name token.
	var name string
	name, token, ok = v.parseToken(NameToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single GetterMethod rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$GetterMethod", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single "(" delimiter.
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single GetterMethod rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$GetterMethod", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ")" delimiter.
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single GetterMethod rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$GetterMethod", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Abstraction rule.
	var abstraction ast.AbstractionLike
	abstraction, token, ok = v.parseAbstraction()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single GetterMethod rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$GetterMethod", token)
		panic(message)
	}

	// Found a single GetterMethod rule.
	ok = true
	v.remove(tokens)
	getterMethod = ast.GetterMethodClass().GetterMethod(
		name,
		abstraction,
	)
	return
}

func (v *parser_) parseImportedPackage() (
	importedPackage ast.ImportedPackageLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single name token.
	var name string
	name, token, ok = v.parseToken(NameToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ImportedPackage rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ImportedPackage", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single path token.
	var path string
	path, token, ok = v.parseToken(PathToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ImportedPackage rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ImportedPackage", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single ImportedPackage rule.
	ok = true
	v.remove(tokens)
	importedPackage = ast.ImportedPackageClass().ImportedPackage(
		name,
		path,
	)
	return
}

func (v *parser_) parseInstanceDeclaration() (
	instanceDeclaration ast.InstanceDeclarationLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single Declaration rule.
	var declaration ast.DeclarationLike
	declaration, token, ok = v.parseDeclaration()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single InstanceDeclaration rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$InstanceDeclaration", token)
		panic(message)
	}

	// Attempt to parse a single "interface" delimiter.
	_, token, ok = v.parseDelimiter("interface")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InstanceDeclaration rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InstanceDeclaration", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single "{" delimiter.
	_, token, ok = v.parseDelimiter("{")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InstanceDeclaration rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InstanceDeclaration", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single InstanceMethods rule.
	var instanceMethods ast.InstanceMethodsLike
	instanceMethods, token, ok = v.parseInstanceMethods()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single InstanceDeclaration rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$InstanceDeclaration", token)
		panic(message)
	}

	// Attempt to parse a single "}" delimiter.
	_, token, ok = v.parseDelimiter("}")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InstanceDeclaration rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InstanceDeclaration", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single InstanceDeclaration rule.
	ok = true
	v.remove(tokens)
	instanceDeclaration = ast.InstanceDeclarationClass().InstanceDeclaration(
		declaration,
		instanceMethods,
	)
	return
}

func (v *parser_) parseInstanceMethods() (
	instanceMethods ast.InstanceMethodsLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single PrincipalSubsection rule.
	var principalSubsection ast.PrincipalSubsectionLike
	principalSubsection, token, ok = v.parsePrincipalSubsection()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single InstanceMethods rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$InstanceMethods", token)
		panic(message)
	}

	// Attempt to parse an optional AttributeSubsection rule.
	var optionalAttributeSubsection ast.AttributeSubsectionLike
	optionalAttributeSubsection, _, ok = v.parseAttributeSubsection()
	if ok {
		// No additional put backs allowed at this point.
		tokens = nil
	}

	// Attempt to parse an optional AspectSubsection rule.
	var optionalAspectSubsection ast.AspectSubsectionLike
	optionalAspectSubsection, _, ok = v.parseAspectSubsection()
	if ok {
		// No additional put backs allowed at this point.
		tokens = nil
	}

	// Found a single InstanceMethods rule.
	ok = true
	v.remove(tokens)
	instanceMethods = ast.InstanceMethodsClass().InstanceMethods(
		principalSubsection,
		optionalAttributeSubsection,
		optionalAspectSubsection,
	)
	return
}

func (v *parser_) parseInstanceSection() (
	instanceSection ast.InstanceSectionLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "// INSTANCE DECLARATIONS" delimiter.
	_, token, ok = v.parseDelimiter("// INSTANCE DECLARATIONS")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InstanceSection rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InstanceSection", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse multiple InstanceDeclaration rules.
	var instanceDeclarations = col.List[ast.InstanceDeclarationLike]()
instanceDeclarationsLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var instanceDeclaration ast.InstanceDeclarationLike
		instanceDeclaration, token, ok = v.parseInstanceDeclaration()
		if !ok {
			switch {
			case count >= 1:
				break instanceDeclarationsLoop
			case uti.IsDefined(tokens):
				// This is not multiple InstanceDeclaration rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$InstanceSection", token)
				message += "1 or more InstanceDeclaration rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		instanceDeclarations.AppendValue(instanceDeclaration)
	}

	// Found a single InstanceSection rule.
	ok = true
	v.remove(tokens)
	instanceSection = ast.InstanceSectionClass().InstanceSection(instanceDeclarations)
	return
}

func (v *parser_) parseInterfaceDeclarations() (
	interfaceDeclarations ast.InterfaceDeclarationsLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single ClassSection rule.
	var classSection ast.ClassSectionLike
	classSection, token, ok = v.parseClassSection()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single InterfaceDeclarations rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$InterfaceDeclarations", token)
		panic(message)
	}

	// Attempt to parse a single InstanceSection rule.
	var instanceSection ast.InstanceSectionLike
	instanceSection, token, ok = v.parseInstanceSection()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single InterfaceDeclarations rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$InterfaceDeclarations", token)
		panic(message)
	}

	// Attempt to parse a single AspectSection rule.
	var aspectSection ast.AspectSectionLike
	aspectSection, token, ok = v.parseAspectSection()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single InterfaceDeclarations rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$InterfaceDeclarations", token)
		panic(message)
	}

	// Found a single InterfaceDeclarations rule.
	ok = true
	v.remove(tokens)
	interfaceDeclarations = ast.InterfaceDeclarationsClass().InterfaceDeclarations(
		classSection,
		instanceSection,
		aspectSection,
	)
	return
}

func (v *parser_) parseLegalNotice() (
	legalNotice ast.LegalNoticeLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single comment token.
	var comment string
	comment, token, ok = v.parseToken(CommentToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single LegalNotice rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$LegalNotice", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single LegalNotice rule.
	ok = true
	v.remove(tokens)
	legalNotice = ast.LegalNoticeClass().LegalNotice(comment)
	return
}

func (v *parser_) parseMap() (
	map_ ast.MapLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "map" delimiter.
	_, token, ok = v.parseDelimiter("map")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Map rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Map", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single "[" delimiter.
	_, token, ok = v.parseDelimiter("[")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Map rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Map", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single name token.
	var name string
	name, token, ok = v.parseToken(NameToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Map rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Map", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single "]" delimiter.
	_, token, ok = v.parseDelimiter("]")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Map rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Map", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Map rule.
	ok = true
	v.remove(tokens)
	map_ = ast.MapClass().Map(name)
	return
}

func (v *parser_) parseMethod() (
	method ast.MethodLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single name token.
	var name string
	name, token, ok = v.parseToken(NameToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Method rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Method", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single "(" delimiter.
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Method rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Method", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse multiple Parameter rules.
	var parameters = col.List[ast.ParameterLike]()
parametersLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var parameter ast.ParameterLike
		parameter, token, ok = v.parseParameter()
		if !ok {
			switch {
			case count >= 0:
				break parametersLoop
			case uti.IsDefined(tokens):
				// This is not multiple Parameter rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$Method", token)
				message += "0 or more Parameter rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		parameters.AppendValue(parameter)
	}

	// Attempt to parse a single ")" delimiter.
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Method rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Method", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse an optional Result rule.
	var optionalResult ast.ResultLike
	optionalResult, _, ok = v.parseResult()
	if ok {
		// No additional put backs allowed at this point.
		tokens = nil
	}

	// Found a single Method rule.
	ok = true
	v.remove(tokens)
	method = ast.MethodClass().Method(
		name,
		parameters,
		optionalResult,
	)
	return
}

func (v *parser_) parseModel() (
	model ast.ModelLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single PackageDeclaration rule.
	var packageDeclaration ast.PackageDeclarationLike
	packageDeclaration, token, ok = v.parsePackageDeclaration()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Model rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Model", token)
		panic(message)
	}

	// Attempt to parse a single PrimitiveDeclarations rule.
	var primitiveDeclarations ast.PrimitiveDeclarationsLike
	primitiveDeclarations, token, ok = v.parsePrimitiveDeclarations()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Model rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Model", token)
		panic(message)
	}

	// Attempt to parse a single InterfaceDeclarations rule.
	var interfaceDeclarations ast.InterfaceDeclarationsLike
	interfaceDeclarations, token, ok = v.parseInterfaceDeclarations()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Model rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Model", token)
		panic(message)
	}

	// Found a single Model rule.
	ok = true
	v.remove(tokens)
	model = ast.ModelClass().Model(
		packageDeclaration,
		primitiveDeclarations,
		interfaceDeclarations,
	)
	return
}

func (v *parser_) parseMultivalue() (
	multivalue ast.MultivalueLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "(" delimiter.
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Multivalue rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Multivalue", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse multiple Parameter rules.
	var parameters = col.List[ast.ParameterLike]()
parametersLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var parameter ast.ParameterLike
		parameter, token, ok = v.parseParameter()
		if !ok {
			switch {
			case count >= 1:
				break parametersLoop
			case uti.IsDefined(tokens):
				// This is not multiple Parameter rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$Multivalue", token)
				message += "1 or more Parameter rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		parameters.AppendValue(parameter)
	}

	// Attempt to parse a single ")" delimiter.
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Multivalue rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Multivalue", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Multivalue rule.
	ok = true
	v.remove(tokens)
	multivalue = ast.MultivalueClass().Multivalue(parameters)
	return
}

func (v *parser_) parseNone() (
	none ast.NoneLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single newline token.
	var newline string
	newline, token, ok = v.parseToken(NewlineToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single None rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$None", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single None rule.
	ok = true
	v.remove(tokens)
	none = ast.NoneClass().None(newline)
	return
}

func (v *parser_) parsePackageDeclaration() (
	packageDeclaration ast.PackageDeclarationLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single LegalNotice rule.
	var legalNotice ast.LegalNoticeLike
	legalNotice, token, ok = v.parseLegalNotice()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single PackageDeclaration rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$PackageDeclaration", token)
		panic(message)
	}

	// Attempt to parse a single PackageHeader rule.
	var packageHeader ast.PackageHeaderLike
	packageHeader, token, ok = v.parsePackageHeader()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single PackageDeclaration rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$PackageDeclaration", token)
		panic(message)
	}

	// Attempt to parse a single PackageImports rule.
	var packageImports ast.PackageImportsLike
	packageImports, token, ok = v.parsePackageImports()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single PackageDeclaration rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$PackageDeclaration", token)
		panic(message)
	}

	// Found a single PackageDeclaration rule.
	ok = true
	v.remove(tokens)
	packageDeclaration = ast.PackageDeclarationClass().PackageDeclaration(
		legalNotice,
		packageHeader,
		packageImports,
	)
	return
}

func (v *parser_) parsePackageHeader() (
	packageHeader ast.PackageHeaderLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single comment token.
	var comment string
	comment, token, ok = v.parseToken(CommentToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single PackageHeader rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$PackageHeader", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single "package" delimiter.
	_, token, ok = v.parseDelimiter("package")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single PackageHeader rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$PackageHeader", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single name token.
	var name string
	name, token, ok = v.parseToken(NameToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single PackageHeader rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$PackageHeader", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single PackageHeader rule.
	ok = true
	v.remove(tokens)
	packageHeader = ast.PackageHeaderClass().PackageHeader(
		comment,
		name,
	)
	return
}

func (v *parser_) parsePackageImports() (
	packageImports ast.PackageImportsLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "import" delimiter.
	_, token, ok = v.parseDelimiter("import")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single PackageImports rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$PackageImports", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single "(" delimiter.
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single PackageImports rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$PackageImports", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse multiple ImportedPackage rules.
	var importedPackages = col.List[ast.ImportedPackageLike]()
importedPackagesLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var importedPackage ast.ImportedPackageLike
		importedPackage, token, ok = v.parseImportedPackage()
		if !ok {
			switch {
			case count >= 0:
				break importedPackagesLoop
			case uti.IsDefined(tokens):
				// This is not multiple ImportedPackage rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$PackageImports", token)
				message += "0 or more ImportedPackage rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		importedPackages.AppendValue(importedPackage)
	}

	// Attempt to parse a single ")" delimiter.
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single PackageImports rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$PackageImports", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single PackageImports rule.
	ok = true
	v.remove(tokens)
	packageImports = ast.PackageImportsClass().PackageImports(importedPackages)
	return
}

func (v *parser_) parseParameter() (
	parameter ast.ParameterLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single name token.
	var name string
	name, token, ok = v.parseToken(NameToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Parameter rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Parameter", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Abstraction rule.
	var abstraction ast.AbstractionLike
	abstraction, token, ok = v.parseAbstraction()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Parameter rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Parameter", token)
		panic(message)
	}

	// Attempt to parse a single "," delimiter.
	_, token, ok = v.parseDelimiter(",")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Parameter rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Parameter", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Parameter rule.
	ok = true
	v.remove(tokens)
	parameter = ast.ParameterClass().Parameter(
		name,
		abstraction,
	)
	return
}

func (v *parser_) parsePrimitiveDeclarations() (
	primitiveDeclarations ast.PrimitiveDeclarationsLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single TypeSection rule.
	var typeSection ast.TypeSectionLike
	typeSection, token, ok = v.parseTypeSection()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single PrimitiveDeclarations rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$PrimitiveDeclarations", token)
		panic(message)
	}

	// Attempt to parse a single FunctionalSection rule.
	var functionalSection ast.FunctionalSectionLike
	functionalSection, token, ok = v.parseFunctionalSection()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single PrimitiveDeclarations rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$PrimitiveDeclarations", token)
		panic(message)
	}

	// Found a single PrimitiveDeclarations rule.
	ok = true
	v.remove(tokens)
	primitiveDeclarations = ast.PrimitiveDeclarationsClass().PrimitiveDeclarations(
		typeSection,
		functionalSection,
	)
	return
}

func (v *parser_) parsePrincipalMethod() (
	principalMethod ast.PrincipalMethodLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single Method rule.
	var method ast.MethodLike
	method, token, ok = v.parseMethod()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single PrincipalMethod rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$PrincipalMethod", token)
		panic(message)
	}

	// Found a single PrincipalMethod rule.
	ok = true
	v.remove(tokens)
	principalMethod = ast.PrincipalMethodClass().PrincipalMethod(method)
	return
}

func (v *parser_) parsePrincipalSubsection() (
	principalSubsection ast.PrincipalSubsectionLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "// Principal Methods" delimiter.
	_, token, ok = v.parseDelimiter("// Principal Methods")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single PrincipalSubsection rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$PrincipalSubsection", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse multiple PrincipalMethod rules.
	var principalMethods = col.List[ast.PrincipalMethodLike]()
principalMethodsLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var principalMethod ast.PrincipalMethodLike
		principalMethod, token, ok = v.parsePrincipalMethod()
		if !ok {
			switch {
			case count >= 1:
				break principalMethodsLoop
			case uti.IsDefined(tokens):
				// This is not multiple PrincipalMethod rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$PrincipalSubsection", token)
				message += "1 or more PrincipalMethod rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		principalMethods.AppendValue(principalMethod)
	}

	// Found a single PrincipalSubsection rule.
	ok = true
	v.remove(tokens)
	principalSubsection = ast.PrincipalSubsectionClass().PrincipalSubsection(principalMethods)
	return
}

func (v *parser_) parseResult() (
	result ast.ResultLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single None Result.
	var none ast.NoneLike
	none, token, ok = v.parseNone()
	if ok {
		// Found a single None Result.
		result = ast.ResultClass().Result(none)
		return
	}

	// Attempt to parse a single Abstraction Result.
	var abstraction ast.AbstractionLike
	abstraction, token, ok = v.parseAbstraction()
	if ok {
		// Found a single Abstraction Result.
		result = ast.ResultClass().Result(abstraction)
		return
	}

	// Attempt to parse a single Multivalue Result.
	var multivalue ast.MultivalueLike
	multivalue, token, ok = v.parseMultivalue()
	if ok {
		// Found a single Multivalue Result.
		result = ast.ResultClass().Result(multivalue)
		return
	}

	// This is not a single Result rule.
	return
}

func (v *parser_) parseSetterMethod() (
	setterMethod ast.SetterMethodLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single name token.
	var name string
	name, token, ok = v.parseToken(NameToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single SetterMethod rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$SetterMethod", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single "(" delimiter.
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single SetterMethod rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$SetterMethod", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Parameter rule.
	var parameter ast.ParameterLike
	parameter, token, ok = v.parseParameter()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single SetterMethod rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$SetterMethod", token)
		panic(message)
	}

	// Attempt to parse a single ")" delimiter.
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single SetterMethod rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$SetterMethod", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single SetterMethod rule.
	ok = true
	v.remove(tokens)
	setterMethod = ast.SetterMethodClass().SetterMethod(
		name,
		parameter,
	)
	return
}

func (v *parser_) parseTypeDeclaration() (
	typeDeclaration ast.TypeDeclarationLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single Declaration rule.
	var declaration ast.DeclarationLike
	declaration, token, ok = v.parseDeclaration()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single TypeDeclaration rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$TypeDeclaration", token)
		panic(message)
	}

	// Attempt to parse a single Abstraction rule.
	var abstraction ast.AbstractionLike
	abstraction, token, ok = v.parseAbstraction()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single TypeDeclaration rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$TypeDeclaration", token)
		panic(message)
	}

	// Attempt to parse an optional Enumeration rule.
	var optionalEnumeration ast.EnumerationLike
	optionalEnumeration, _, ok = v.parseEnumeration()
	if ok {
		// No additional put backs allowed at this point.
		tokens = nil
	}

	// Found a single TypeDeclaration rule.
	ok = true
	v.remove(tokens)
	typeDeclaration = ast.TypeDeclarationClass().TypeDeclaration(
		declaration,
		abstraction,
		optionalEnumeration,
	)
	return
}

func (v *parser_) parseTypeSection() (
	typeSection ast.TypeSectionLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "// TYPE DECLARATIONS" delimiter.
	_, token, ok = v.parseDelimiter("// TYPE DECLARATIONS")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single TypeSection rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$TypeSection", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse multiple TypeDeclaration rules.
	var typeDeclarations = col.List[ast.TypeDeclarationLike]()
typeDeclarationsLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var typeDeclaration ast.TypeDeclarationLike
		typeDeclaration, token, ok = v.parseTypeDeclaration()
		if !ok {
			switch {
			case count >= 0:
				break typeDeclarationsLoop
			case uti.IsDefined(tokens):
				// This is not multiple TypeDeclaration rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$TypeSection", token)
				message += "0 or more TypeDeclaration rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		typeDeclarations.AppendValue(typeDeclaration)
	}

	// Found a single TypeSection rule.
	ok = true
	v.remove(tokens)
	typeSection = ast.TypeSectionClass().TypeSection(typeDeclarations)
	return
}

func (v *parser_) parseValue() (
	value ast.ValueLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single name token.
	var name string
	name, token, ok = v.parseToken(NameToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Value rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Value", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Abstraction rule.
	var abstraction ast.AbstractionLike
	abstraction, token, ok = v.parseAbstraction()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Value rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Value", token)
		panic(message)
	}

	// Attempt to parse a single "=" delimiter.
	_, token, ok = v.parseDelimiter("=")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Value rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Value", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single "iota" delimiter.
	_, token, ok = v.parseDelimiter("iota")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Value rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Value", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Value rule.
	ok = true
	v.remove(tokens)
	value = ast.ValueClass().Value(
		name,
		abstraction,
	)
	return
}

func (v *parser_) parseWrapper() (
	wrapper ast.WrapperLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single Array Wrapper.
	var array ast.ArrayLike
	array, token, ok = v.parseArray()
	if ok {
		// Found a single Array Wrapper.
		wrapper = ast.WrapperClass().Wrapper(array)
		return
	}

	// Attempt to parse a single Map Wrapper.
	var map_ ast.MapLike
	map_, token, ok = v.parseMap()
	if ok {
		// Found a single Map Wrapper.
		wrapper = ast.WrapperClass().Wrapper(map_)
		return
	}

	// Attempt to parse a single Channel Wrapper.
	var channel ast.ChannelLike
	channel, token, ok = v.parseChannel()
	if ok {
		// Found a single Channel Wrapper.
		wrapper = ast.WrapperClass().Wrapper(channel)
		return
	}

	// This is not a single Wrapper rule.
	return
}

func (v *parser_) parseDelimiter(
	expectedValue string,
) (
	value string,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single delimiter.
	value, token, ok = v.parseToken(DelimiterToken)
	if ok {
		if value == expectedValue {
			// Found the desired delimiter.
			return
		}
		v.next_.AddValue(token)
		ok = false
	}

	// This is not the desired delimiter.
	return
}

func (v *parser_) parseToken(
	tokenType TokenType,
) (
	value string,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a specific token type.
	var tokens = col.List[TokenLike]()
	token = v.getNextToken()
	for token != nil {
		tokens.AppendValue(token)
		switch token.GetType() {
		case tokenType:
			// Found the desired token type.
			value = token.GetValue()
			ok = true
			return
		case SpaceToken, NewlineToken:
			// Ignore any unrequested whitespace.
			token = v.getNextToken()
		default:
			// This is not the desired token type.
			v.putBack(tokens)
			return
		}
	}

	// We are at the end-of-file marker.
	return
}

func (v *parser_) formatError(
	ruleName string,
	token TokenLike,
) string {
	// Format the error message.
	var message = fmt.Sprintf(
		"An unexpected token was received by the parser: %v\n",
		ScannerClass().FormatToken(token),
	)
	var line = token.GetLine()
	var lines = sts.Split(v.source_, "\n")

	// Append the source lines with the error in it.
	message += "\033[36m"
	for index := line - 3; index < line; index++ {
		if index > 1 {
			message += fmt.Sprintf("%04d: ", index) + string(lines[index-1]) + "\n"
		}
	}
	message += fmt.Sprintf("%04d: ", line) + string(lines[line-1]) + "\n"

	// Append an arrow pointing to the error.
	message += " \033[32m>>>─"
	var count uint
	for count < token.GetPosition() {
		message += "─"
		count++
	}
	message += "⌃\033[36m\n"

	// Append the following source line for context.
	if line < uint(len(lines)) {
		message += fmt.Sprintf("%04d: ", line+1) + string(lines[line]) + "\n"
	}
	message += "\033[0m\n"
	if uti.IsDefined(ruleName) {
		message += "Was expecting:\n"
		message += fmt.Sprintf(
			"  \033[32m%v: \033[33m%v\033[0m\n\n",
			ruleName,
			v.getDefinition(ruleName),
		)
	}
	return message
}

func (v *parser_) getDefinition(
	ruleName string,
) string {
	var syntax = parserClass().syntax_
	var definition = syntax.GetValue(ruleName)
	return definition
}

func (v *parser_) getNextToken() TokenLike {
	// Check for any read, but unprocessed tokens.
	if !v.next_.IsEmpty() {
		return v.next_.RemoveLast()
	}

	// Read a new token from the token stream.
	var token, ok = v.tokens_.RemoveFirst() // This will wait for a token.
	if !ok {
		// The token channel has been closed.
		return nil
	}

	// Check for an error token.
	if token.GetType() == ErrorToken {
		var message = v.formatError("", token)
		panic(message)
	}

	return token
}

func (v *parser_) putBack(
	tokens abs.Sequential[TokenLike],
) {
	var iterator = tokens.GetIterator()
	for iterator.ToEnd(); iterator.HasPrevious(); {
		var token = iterator.GetPrevious()
		v.next_.AddValue(token)
	}
}

func (v *parser_) remove(
	tokens abs.Sequential[TokenLike],
) {
	// NOTE: This method does nothing but must exist to satisfy the lint
	// check on the generated parser code.
}

// Instance Structure

type parser_ struct {
	// Declare the instance attributes.
	source_ string                   // The original source code.
	tokens_ abs.QueueLike[TokenLike] // A queue of unread tokens from the scanner.
	next_   abs.StackLike[TokenLike] // A stack of read, but unprocessed tokens.
}

// Class Structure

type parserClass_ struct {
	// Declare the class constants.
	syntax_ abs.CatalogLike[string, string]
}

// Class Reference

func parserClass() *parserClass_ {
	return parserClassReference_
}

var parserClassReference_ = &parserClass_{
	// Initialize the class constants.
	syntax_: col.CatalogFromMap[string, string](
		map[string]string{
			"$Model":                 `PackageDeclaration PrimitiveDeclarations InterfaceDeclarations`,
			"$PackageDeclaration":    `LegalNotice PackageHeader PackageImports`,
			"$PrimitiveDeclarations": `TypeSection FunctionalSection`,
			"$InterfaceDeclarations": `ClassSection InstanceSection AspectSection`,
			"$LegalNotice":           `comment`,
			"$PackageHeader":         `comment "package" name`,
			"$PackageImports":        `"import" "(" ImportedPackage* ")"`,
			"$ImportedPackage":       `name path`,
			"$TypeSection":           `"// TYPE DECLARATIONS" TypeDeclaration*`,
			"$TypeDeclaration":       `Declaration Abstraction Enumeration?`,
			"$Declaration":           `comment "type" name Constraints?`,
			"$Constraints":           `"[" Constraint AdditionalConstraint* "]"`,
			"$Constraint":            `name Abstraction`,
			"$AdditionalConstraint":  `"," Constraint`,
			"$Abstraction":           `Wrapper? prefix? name Arguments?`,
			"$Wrapper": `
  - Array
  - Map
  - Channel`,
			"$Array":                 `"[" "]"`,
			"$Map":                   `"map" "[" name "]"`,
			"$Channel":               `"chan"`,
			"$Arguments":             `"[" Argument AdditionalArgument* "]"`,
			"$Argument":              `Abstraction`,
			"$AdditionalArgument":    `"," Argument`,
			"$Enumeration":           `"const" "(" Value AdditionalValue* ")"`,
			"$Value":                 `name Abstraction "=" "iota"`,
			"$AdditionalValue":       `name`,
			"$FunctionalSection":     `"// FUNCTIONAL DECLARATIONS" FunctionalDeclaration*`,
			"$FunctionalDeclaration": `Declaration "func" "(" Parameter* ")" Result`,
			"$Parameter":             `name Abstraction ","`,
			"$Result": `
  - None
  - Abstraction
  - Multivalue`,
			"$None":                  `newline`,
			"$Multivalue":            `"(" Parameter+ ")"`,
			"$ClassSection":          `"// CLASS DECLARATIONS" ClassDeclaration+`,
			"$ClassDeclaration":      `Declaration "interface" "{" ClassMethods "}"`,
			"$ClassMethods":          `ConstructorSubsection ConstantSubsection? FunctionSubsection?`,
			"$ConstructorSubsection": `"// Constructor Methods" ConstructorMethod+`,
			"$ConstructorMethod":     `name "(" Parameter* ")" Abstraction`,
			"$ConstantSubsection":    `"// Constant Methods" ConstantMethod+`,
			"$ConstantMethod":        `name "(" ")" Abstraction`,
			"$FunctionSubsection":    `"// Function Methods" FunctionMethod+`,
			"$FunctionMethod":        `name "(" Parameter* ")" Result`,
			"$InstanceSection":       `"// INSTANCE DECLARATIONS" InstanceDeclaration+`,
			"$InstanceDeclaration":   `Declaration "interface" "{" InstanceMethods "}"`,
			"$InstanceMethods":       `PrincipalSubsection AttributeSubsection? AspectSubsection?`,
			"$PrincipalSubsection":   `"// Principal Methods" PrincipalMethod+`,
			"$PrincipalMethod":       `Method`,
			"$Method":                `name "(" Parameter* ")" Result?`,
			"$AttributeSubsection":   `"// Attribute Methods" AttributeMethod+`,
			"$AttributeMethod": `
  - GetterMethod
  - SetterMethod`,
			"$GetterMethod":      `name "(" ")" Abstraction`,
			"$SetterMethod":      `name "(" Parameter ")"`,
			"$AspectSubsection":  `"// Aspect Interfaces" AspectInterface+`,
			"$AspectInterface":   `Abstraction`,
			"$AspectSection":     `"// ASPECT DECLARATIONS" AspectDeclaration*`,
			"$AspectDeclaration": `Declaration "interface" "{" AspectMethod+ "}"`,
			"$AspectMethod":      `Method`,
		},
	),
}
