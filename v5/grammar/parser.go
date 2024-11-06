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

package grammar

import (
	fmt "fmt"
	ast "github.com/craterdog/go-class-model/v5/ast"
	col "github.com/craterdog/go-collection-framework/v4"
	abs "github.com/craterdog/go-collection-framework/v4/collection"
	uti "github.com/craterdog/go-missing-utilities/v2"
	mat "math"
	sts "strings"
)

// CLASS INTERFACE

// Access Function

func Parser() ParserClassLike {
	return parserReference()
}

// Constructor Methods

func (c *parserClass_) Make() ParserLike {
	var instance = &parser_{
		// Initialize the instance attributes.
	}
	return instance
}

// INSTANCE INTERFACE

// Primary Methods

func (v *parser_) GetClass() ParserClassLike {
	return parserReference()
}

func (v *parser_) ParseSource(
	source string,
) ast.ModelLike {
	v.source_ = source
	v.tokens_ = col.Queue[TokenLike](parserReference().queueSize_)
	v.next_ = col.Stack[TokenLike](parserReference().stackSize_)

	// The scanner runs in a separate Go routine.
	Scanner().Make(v.source_, v.tokens_)

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

	// Attempt to parse an optional Prefix rule.
	var optionalPrefix ast.PrefixLike
	optionalPrefix, _, ok = v.parsePrefix()
	if ok {
		// No additional put backs allowed at this point.
		tokens = nil
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

	// Attempt to parse an optional Suffix rule.
	var optionalSuffix ast.SuffixLike
	optionalSuffix, _, ok = v.parseSuffix()
	if ok {
		// No additional put backs allowed at this point.
		tokens = nil
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
	abstraction = ast.Abstraction().Make(
		optionalPrefix,
		name,
		optionalSuffix,
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
	additionalArgument = ast.AdditionalArgument().Make(argument)
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
	additionalConstraint = ast.AdditionalConstraint().Make(constraint)
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
	additionalValue = ast.AdditionalValue().Make(name)
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
	argument = ast.Argument().Make(abstraction)
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
	arguments = ast.Arguments().Make(
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
	array = ast.Array().Make()
	return
}

func (v *parser_) parseAspectDefinition() (
	aspectDefinition ast.AspectDefinitionLike,
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
		// This is not a single AspectDefinition rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$AspectDefinition", token)
		panic(message)
	}

	// Attempt to parse a single "interface" delimiter.
	_, token, ok = v.parseDelimiter("interface")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single AspectDefinition rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$AspectDefinition", token)
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
			// This is not a single AspectDefinition rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$AspectDefinition", token)
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
				var message = v.formatError("$AspectDefinition", token)
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
			// This is not a single AspectDefinition rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$AspectDefinition", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single AspectDefinition rule.
	ok = true
	v.remove(tokens)
	aspectDefinition = ast.AspectDefinition().Make(
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
	aspectInterface = ast.AspectInterface().Make(abstraction)
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
	aspectMethod = ast.AspectMethod().Make(method)
	return
}

func (v *parser_) parseAspectSection() (
	aspectSection ast.AspectSectionLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "// Aspect Declarations" delimiter.
	_, token, ok = v.parseDelimiter("// Aspect Declarations")
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

	// Attempt to parse multiple AspectDefinition rules.
	var aspectDefinitions = col.List[ast.AspectDefinitionLike]()
aspectDefinitionsLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var aspectDefinition ast.AspectDefinitionLike
		aspectDefinition, token, ok = v.parseAspectDefinition()
		if !ok {
			switch {
			case count >= 1:
				break aspectDefinitionsLoop
			case uti.IsDefined(tokens):
				// This is not multiple AspectDefinition rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$AspectSection", token)
				message += "1 or more AspectDefinition rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		aspectDefinitions.AppendValue(aspectDefinition)
	}

	// Found a single AspectSection rule.
	ok = true
	v.remove(tokens)
	aspectSection = ast.AspectSection().Make(aspectDefinitions)
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
	aspectSubsection = ast.AspectSubsection().Make(aspectInterfaces)
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
		attributeMethod = ast.AttributeMethod().Make(getterMethod)
		return
	}

	// Attempt to parse a single SetterMethod AttributeMethod.
	var setterMethod ast.SetterMethodLike
	setterMethod, token, ok = v.parseSetterMethod()
	if ok {
		// Found a single SetterMethod AttributeMethod.
		attributeMethod = ast.AttributeMethod().Make(setterMethod)
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
	attributeSubsection = ast.AttributeSubsection().Make(attributeMethods)
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
	channel = ast.Channel().Make()
	return
}

func (v *parser_) parseClassDefinition() (
	classDefinition ast.ClassDefinitionLike,
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
		// This is not a single ClassDefinition rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$ClassDefinition", token)
		panic(message)
	}

	// Attempt to parse a single "interface" delimiter.
	_, token, ok = v.parseDelimiter("interface")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ClassDefinition rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ClassDefinition", token)
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
			// This is not a single ClassDefinition rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ClassDefinition", token)
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
		// This is not a single ClassDefinition rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$ClassDefinition", token)
		panic(message)
	}

	// Attempt to parse a single "}" delimiter.
	_, token, ok = v.parseDelimiter("}")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ClassDefinition rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ClassDefinition", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single ClassDefinition rule.
	ok = true
	v.remove(tokens)
	classDefinition = ast.ClassDefinition().Make(
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
	classMethods = ast.ClassMethods().Make(
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

	// Attempt to parse a single "// Class Declarations" delimiter.
	_, token, ok = v.parseDelimiter("// Class Declarations")
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

	// Attempt to parse multiple ClassDefinition rules.
	var classDefinitions = col.List[ast.ClassDefinitionLike]()
classDefinitionsLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var classDefinition ast.ClassDefinitionLike
		classDefinition, token, ok = v.parseClassDefinition()
		if !ok {
			switch {
			case count >= 1:
				break classDefinitionsLoop
			case uti.IsDefined(tokens):
				// This is not multiple ClassDefinition rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$ClassSection", token)
				message += "1 or more ClassDefinition rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		classDefinitions.AppendValue(classDefinition)
	}

	// Found a single ClassSection rule.
	ok = true
	v.remove(tokens)
	classSection = ast.ClassSection().Make(classDefinitions)
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
	constantMethod = ast.ConstantMethod().Make(
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
	constantSubsection = ast.ConstantSubsection().Make(constantMethods)
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
	constraint = ast.Constraint().Make(
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
	constraints = ast.Constraints().Make(
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
	constructorMethod = ast.ConstructorMethod().Make(
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
	constructorSubsection = ast.ConstructorSubsection().Make(constructorMethods)
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
	declaration = ast.Declaration().Make(
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
	enumeration = ast.Enumeration().Make(
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
	functionMethod = ast.FunctionMethod().Make(
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
	functionSubsection = ast.FunctionSubsection().Make(functionMethods)
	return
}

func (v *parser_) parseFunctionalDefinition() (
	functionalDefinition ast.FunctionalDefinitionLike,
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
		// This is not a single FunctionalDefinition rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$FunctionalDefinition", token)
		panic(message)
	}

	// Attempt to parse a single "func" delimiter.
	_, token, ok = v.parseDelimiter("func")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single FunctionalDefinition rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$FunctionalDefinition", token)
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
			// This is not a single FunctionalDefinition rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$FunctionalDefinition", token)
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
				var message = v.formatError("$FunctionalDefinition", token)
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
			// This is not a single FunctionalDefinition rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$FunctionalDefinition", token)
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
		// This is not a single FunctionalDefinition rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$FunctionalDefinition", token)
		panic(message)
	}

	// Found a single FunctionalDefinition rule.
	ok = true
	v.remove(tokens)
	functionalDefinition = ast.FunctionalDefinition().Make(
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

	// Attempt to parse a single "// Functional Declarations" delimiter.
	_, token, ok = v.parseDelimiter("// Functional Declarations")
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

	// Attempt to parse multiple FunctionalDefinition rules.
	var functionalDefinitions = col.List[ast.FunctionalDefinitionLike]()
functionalDefinitionsLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var functionalDefinition ast.FunctionalDefinitionLike
		functionalDefinition, token, ok = v.parseFunctionalDefinition()
		if !ok {
			switch {
			case count >= 1:
				break functionalDefinitionsLoop
			case uti.IsDefined(tokens):
				// This is not multiple FunctionalDefinition rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$FunctionalSection", token)
				message += "1 or more FunctionalDefinition rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		functionalDefinitions.AppendValue(functionalDefinition)
	}

	// Found a single FunctionalSection rule.
	ok = true
	v.remove(tokens)
	functionalSection = ast.FunctionalSection().Make(functionalDefinitions)
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
	getterMethod = ast.GetterMethod().Make(
		name,
		abstraction,
	)
	return
}

func (v *parser_) parseHeader() (
	header ast.HeaderLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single comment token.
	var comment string
	comment, token, ok = v.parseToken(CommentToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Header rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Header", token)
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
			// This is not a single Header rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Header", token)
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
			// This is not a single Header rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Header", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Header rule.
	ok = true
	v.remove(tokens)
	header = ast.Header().Make(
		comment,
		name,
	)
	return
}

func (v *parser_) parseImports() (
	imports ast.ImportsLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "import" delimiter.
	_, token, ok = v.parseDelimiter("import")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Imports rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Imports", token)
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
			// This is not a single Imports rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Imports", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse multiple Module rules.
	var modules = col.List[ast.ModuleLike]()
modulesLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var module ast.ModuleLike
		module, token, ok = v.parseModule()
		if !ok {
			switch {
			case count >= 1:
				break modulesLoop
			case uti.IsDefined(tokens):
				// This is not multiple Module rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$Imports", token)
				message += "1 or more Module rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		modules.AppendValue(module)
	}

	// Attempt to parse a single ")" delimiter.
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Imports rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Imports", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Imports rule.
	ok = true
	v.remove(tokens)
	imports = ast.Imports().Make(modules)
	return
}

func (v *parser_) parseInstanceDefinition() (
	instanceDefinition ast.InstanceDefinitionLike,
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
		// This is not a single InstanceDefinition rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$InstanceDefinition", token)
		panic(message)
	}

	// Attempt to parse a single "interface" delimiter.
	_, token, ok = v.parseDelimiter("interface")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InstanceDefinition rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InstanceDefinition", token)
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
			// This is not a single InstanceDefinition rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InstanceDefinition", token)
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
		// This is not a single InstanceDefinition rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$InstanceDefinition", token)
		panic(message)
	}

	// Attempt to parse a single "}" delimiter.
	_, token, ok = v.parseDelimiter("}")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InstanceDefinition rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InstanceDefinition", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single InstanceDefinition rule.
	ok = true
	v.remove(tokens)
	instanceDefinition = ast.InstanceDefinition().Make(
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

	// Attempt to parse a single PrimarySubsection rule.
	var primarySubsection ast.PrimarySubsectionLike
	primarySubsection, token, ok = v.parsePrimarySubsection()
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
	instanceMethods = ast.InstanceMethods().Make(
		primarySubsection,
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

	// Attempt to parse a single "// Instance Declarations" delimiter.
	_, token, ok = v.parseDelimiter("// Instance Declarations")
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

	// Attempt to parse multiple InstanceDefinition rules.
	var instanceDefinitions = col.List[ast.InstanceDefinitionLike]()
instanceDefinitionsLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var instanceDefinition ast.InstanceDefinitionLike
		instanceDefinition, token, ok = v.parseInstanceDefinition()
		if !ok {
			switch {
			case count >= 1:
				break instanceDefinitionsLoop
			case uti.IsDefined(tokens):
				// This is not multiple InstanceDefinition rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$InstanceSection", token)
				message += "1 or more InstanceDefinition rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		instanceDefinitions.AppendValue(instanceDefinition)
	}

	// Found a single InstanceSection rule.
	ok = true
	v.remove(tokens)
	instanceSection = ast.InstanceSection().Make(instanceDefinitions)
	return
}

func (v *parser_) parseInterfaceDefinitions() (
	interfaceDefinitions ast.InterfaceDefinitionsLike,
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
		// This is not a single InterfaceDefinitions rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$InterfaceDefinitions", token)
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
		// This is not a single InterfaceDefinitions rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$InterfaceDefinitions", token)
		panic(message)
	}

	// Attempt to parse an optional AspectSection rule.
	var optionalAspectSection ast.AspectSectionLike
	optionalAspectSection, _, ok = v.parseAspectSection()
	if ok {
		// No additional put backs allowed at this point.
		tokens = nil
	}

	// Found a single InterfaceDefinitions rule.
	ok = true
	v.remove(tokens)
	interfaceDefinitions = ast.InterfaceDefinitions().Make(
		classSection,
		instanceSection,
		optionalAspectSection,
	)
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
	map_ = ast.Map().Make(name)
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
	method = ast.Method().Make(
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

	// Attempt to parse a single ModuleDefinition rule.
	var moduleDefinition ast.ModuleDefinitionLike
	moduleDefinition, token, ok = v.parseModuleDefinition()
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

	// Attempt to parse a single PrimitiveDefinitions rule.
	var primitiveDefinitions ast.PrimitiveDefinitionsLike
	primitiveDefinitions, token, ok = v.parsePrimitiveDefinitions()
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

	// Attempt to parse a single InterfaceDefinitions rule.
	var interfaceDefinitions ast.InterfaceDefinitionsLike
	interfaceDefinitions, token, ok = v.parseInterfaceDefinitions()
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
	model = ast.Model().Make(
		moduleDefinition,
		primitiveDefinitions,
		interfaceDefinitions,
	)
	return
}

func (v *parser_) parseModule() (
	module ast.ModuleLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single name token.
	var name string
	name, token, ok = v.parseToken(NameToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Module rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Module", token)
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
			// This is not a single Module rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Module", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Module rule.
	ok = true
	v.remove(tokens)
	module = ast.Module().Make(
		name,
		path,
	)
	return
}

func (v *parser_) parseModuleDefinition() (
	moduleDefinition ast.ModuleDefinitionLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single Notice rule.
	var notice ast.NoticeLike
	notice, token, ok = v.parseNotice()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single ModuleDefinition rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$ModuleDefinition", token)
		panic(message)
	}

	// Attempt to parse a single Header rule.
	var header ast.HeaderLike
	header, token, ok = v.parseHeader()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single ModuleDefinition rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$ModuleDefinition", token)
		panic(message)
	}

	// Attempt to parse an optional Imports rule.
	var optionalImports ast.ImportsLike
	optionalImports, _, ok = v.parseImports()
	if ok {
		// No additional put backs allowed at this point.
		tokens = nil
	}

	// Found a single ModuleDefinition rule.
	ok = true
	v.remove(tokens)
	moduleDefinition = ast.ModuleDefinition().Make(
		notice,
		header,
		optionalImports,
	)
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
	none = ast.None().Make(newline)
	return
}

func (v *parser_) parseNotice() (
	notice ast.NoticeLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single comment token.
	var comment string
	comment, token, ok = v.parseToken(CommentToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Notice rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Notice", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Notice rule.
	ok = true
	v.remove(tokens)
	notice = ast.Notice().Make(comment)
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
	parameter = ast.Parameter().Make(
		name,
		abstraction,
	)
	return
}

func (v *parser_) parseParameterized() (
	parameterized ast.ParameterizedLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "(" delimiter.
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Parameterized rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Parameterized", token)
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
				var message = v.formatError("$Parameterized", token)
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
			// This is not a single Parameterized rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Parameterized", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Parameterized rule.
	ok = true
	v.remove(tokens)
	parameterized = ast.Parameterized().Make(parameters)
	return
}

func (v *parser_) parsePrefix() (
	prefix ast.PrefixLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single Array Prefix.
	var array ast.ArrayLike
	array, token, ok = v.parseArray()
	if ok {
		// Found a single Array Prefix.
		prefix = ast.Prefix().Make(array)
		return
	}

	// Attempt to parse a single Map Prefix.
	var map_ ast.MapLike
	map_, token, ok = v.parseMap()
	if ok {
		// Found a single Map Prefix.
		prefix = ast.Prefix().Make(map_)
		return
	}

	// Attempt to parse a single Channel Prefix.
	var channel ast.ChannelLike
	channel, token, ok = v.parseChannel()
	if ok {
		// Found a single Channel Prefix.
		prefix = ast.Prefix().Make(channel)
		return
	}

	// This is not a single Prefix rule.
	return
}

func (v *parser_) parsePrimaryMethod() (
	primaryMethod ast.PrimaryMethodLike,
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
		// This is not a single PrimaryMethod rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$PrimaryMethod", token)
		panic(message)
	}

	// Found a single PrimaryMethod rule.
	ok = true
	v.remove(tokens)
	primaryMethod = ast.PrimaryMethod().Make(method)
	return
}

func (v *parser_) parsePrimarySubsection() (
	primarySubsection ast.PrimarySubsectionLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "// Primary Methods" delimiter.
	_, token, ok = v.parseDelimiter("// Primary Methods")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single PrimarySubsection rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$PrimarySubsection", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse multiple PrimaryMethod rules.
	var primaryMethods = col.List[ast.PrimaryMethodLike]()
primaryMethodsLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var primaryMethod ast.PrimaryMethodLike
		primaryMethod, token, ok = v.parsePrimaryMethod()
		if !ok {
			switch {
			case count >= 1:
				break primaryMethodsLoop
			case uti.IsDefined(tokens):
				// This is not multiple PrimaryMethod rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$PrimarySubsection", token)
				message += "1 or more PrimaryMethod rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		primaryMethods.AppendValue(primaryMethod)
	}

	// Found a single PrimarySubsection rule.
	ok = true
	v.remove(tokens)
	primarySubsection = ast.PrimarySubsection().Make(primaryMethods)
	return
}

func (v *parser_) parsePrimitiveDefinitions() (
	primitiveDefinitions ast.PrimitiveDefinitionsLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse an optional TypeSection rule.
	var optionalTypeSection ast.TypeSectionLike
	optionalTypeSection, _, ok = v.parseTypeSection()
	if ok {
		// No additional put backs allowed at this point.
		tokens = nil
	}

	// Attempt to parse an optional FunctionalSection rule.
	var optionalFunctionalSection ast.FunctionalSectionLike
	optionalFunctionalSection, _, ok = v.parseFunctionalSection()
	if ok {
		// No additional put backs allowed at this point.
		tokens = nil
	}

	// Found a single PrimitiveDefinitions rule.
	ok = true
	v.remove(tokens)
	primitiveDefinitions = ast.PrimitiveDefinitions().Make(
		optionalTypeSection,
		optionalFunctionalSection,
	)
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
		result = ast.Result().Make(none)
		return
	}

	// Attempt to parse a single Abstraction Result.
	var abstraction ast.AbstractionLike
	abstraction, token, ok = v.parseAbstraction()
	if ok {
		// Found a single Abstraction Result.
		result = ast.Result().Make(abstraction)
		return
	}

	// Attempt to parse a single Parameterized Result.
	var parameterized ast.ParameterizedLike
	parameterized, token, ok = v.parseParameterized()
	if ok {
		// Found a single Parameterized Result.
		result = ast.Result().Make(parameterized)
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
	setterMethod = ast.SetterMethod().Make(
		name,
		parameter,
	)
	return
}

func (v *parser_) parseSuffix() (
	suffix ast.SuffixLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "." delimiter.
	_, token, ok = v.parseDelimiter(".")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Suffix rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Suffix", token)
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
			// This is not a single Suffix rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Suffix", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Suffix rule.
	ok = true
	v.remove(tokens)
	suffix = ast.Suffix().Make(name)
	return
}

func (v *parser_) parseTypeDefinition() (
	typeDefinition ast.TypeDefinitionLike,
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
		// This is not a single TypeDefinition rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$TypeDefinition", token)
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
		// This is not a single TypeDefinition rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$TypeDefinition", token)
		panic(message)
	}

	// Attempt to parse an optional Enumeration rule.
	var optionalEnumeration ast.EnumerationLike
	optionalEnumeration, _, ok = v.parseEnumeration()
	if ok {
		// No additional put backs allowed at this point.
		tokens = nil
	}

	// Found a single TypeDefinition rule.
	ok = true
	v.remove(tokens)
	typeDefinition = ast.TypeDefinition().Make(
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

	// Attempt to parse a single "// Type Declarations" delimiter.
	_, token, ok = v.parseDelimiter("// Type Declarations")
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

	// Attempt to parse multiple TypeDefinition rules.
	var typeDefinitions = col.List[ast.TypeDefinitionLike]()
typeDefinitionsLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var typeDefinition ast.TypeDefinitionLike
		typeDefinition, token, ok = v.parseTypeDefinition()
		if !ok {
			switch {
			case count >= 1:
				break typeDefinitionsLoop
			case uti.IsDefined(tokens):
				// This is not multiple TypeDefinition rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$TypeSection", token)
				message += "1 or more TypeDefinition rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		typeDefinitions.AppendValue(typeDefinition)
	}

	// Found a single TypeSection rule.
	ok = true
	v.remove(tokens)
	typeSection = ast.TypeSection().Make(typeDefinitions)
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
	value = ast.Value().Make(
		name,
		abstraction,
	)
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

func (v *parser_) parseToken(tokenType TokenType) (
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
		Scanner().FormatToken(token),
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
	return parserReference().syntax_.GetValue(ruleName)
}

func (v *parser_) getNextToken() TokenLike {
	// Check for any read, but unprocessed tokens.
	if !v.next_.IsEmpty() {
		return v.next_.RemoveTop()
	}

	// Read a new token from the token stream.
	var token, ok = v.tokens_.RemoveHead() // This will wait for a token.
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
	queueSize_ uint
	stackSize_ uint
	syntax_    abs.CatalogLike[string, string]
}

// Class Reference

func parserReference() *parserClass_ {
	return parserReference_
}

var parserReference_ = &parserClass_{
	// Initialize the class constants.
	queueSize_: 16,
	stackSize_: 16,
	syntax_: col.Catalog[string, string](
		map[string]string{
			"$Model":                `ModuleDefinition PrimitiveDefinitions InterfaceDefinitions`,
			"$ModuleDefinition":     `Notice Header Imports?`,
			"$PrimitiveDefinitions": `TypeSection? FunctionalSection?`,
			"$InterfaceDefinitions": `ClassSection InstanceSection AspectSection?`,
			"$Notice":               `comment`,
			"$Header":               `comment "package" name`,
			"$Imports":              `"import" "(" Module+ ")"`,
			"$Module":               `name path`,
			"$TypeSection":          `"// Type Declarations" TypeDefinition+`,
			"$TypeDefinition":       `Declaration Abstraction Enumeration?`,
			"$Declaration":          `comment "type" name Constraints?`,
			"$Constraints":          `"[" Constraint AdditionalConstraint* "]"`,
			"$Constraint":           `name Abstraction`,
			"$AdditionalConstraint": `"," Constraint`,
			"$Abstraction":          `Prefix? name Suffix? Arguments?`,
			"$Prefix": `
  - Array
  - Map
  - Channel`,
			"$Array":                `"[" "]"`,
			"$Map":                  `"map" "[" name "]"`,
			"$Channel":              `"chan"`,
			"$Suffix":               `"." name`,
			"$Arguments":            `"[" Argument AdditionalArgument* "]"`,
			"$Argument":             `Abstraction`,
			"$AdditionalArgument":   `"," Argument`,
			"$Enumeration":          `"const" "(" Value AdditionalValue* ")"`,
			"$Value":                `name Abstraction "=" "iota"`,
			"$AdditionalValue":      `name`,
			"$FunctionalSection":    `"// Functional Declarations" FunctionalDefinition+`,
			"$FunctionalDefinition": `Declaration "func" "(" Parameter* ")" Result`,
			"$Parameter":            `name Abstraction ","`,
			"$Result": `
  - None
  - Abstraction
  - Parameterized`,
			"$None":                  `newline`,
			"$Parameterized":         `"(" Parameter+ ")"`,
			"$ClassSection":          `"// Class Declarations" ClassDefinition+`,
			"$ClassDefinition":       `Declaration "interface" "{" ClassMethods "}"`,
			"$ClassMethods":          `ConstructorSubsection ConstantSubsection? FunctionSubsection?`,
			"$ConstructorSubsection": `"// Constructor Methods" ConstructorMethod+`,
			"$ConstructorMethod":     `name "(" Parameter* ")" Abstraction`,
			"$ConstantSubsection":    `"// Constant Methods" ConstantMethod+`,
			"$ConstantMethod":        `name "(" ")" Abstraction`,
			"$FunctionSubsection":    `"// Function Methods" FunctionMethod+`,
			"$FunctionMethod":        `name "(" Parameter* ")" Result`,
			"$InstanceSection":       `"// Instance Declarations" InstanceDefinition+`,
			"$InstanceDefinition":    `Declaration "interface" "{" InstanceMethods "}"`,
			"$InstanceMethods":       `PrimarySubsection AttributeSubsection? AspectSubsection?`,
			"$PrimarySubsection":     `"// Primary Methods" PrimaryMethod+`,
			"$PrimaryMethod":         `Method`,
			"$Method":                `name "(" Parameter* ")" Result?`,
			"$AttributeSubsection":   `"// Attribute Methods" AttributeMethod+`,
			"$AttributeMethod": `
  - GetterMethod
  - SetterMethod`,
			"$GetterMethod":     `name "(" ")" Abstraction`,
			"$SetterMethod":     `name "(" Parameter ")"`,
			"$AspectSubsection": `"// Aspect Interfaces" AspectInterface+`,
			"$AspectInterface":  `Abstraction`,
			"$AspectSection":    `"// Aspect Declarations" AspectDefinition+`,
			"$AspectDefinition": `Declaration "interface" "{" AspectMethod+ "}"`,
			"$AspectMethod":     `Method`,
		},
	),
}
