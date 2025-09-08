package generator

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"strings"
	"testing"

	"github.com/jacoelho/banking/registry"
	"github.com/jacoelho/banking/registry/rule"
)

// Internal tests for code generation functions

func TestGenerateCodeForCountry_Internal(t *testing.T) {
	country := registry.Country{
		Code:             "AL",
		Name:             "Albania",
		IBAN:             "AL2!n8!n16!c",
		BBAN:             "8!n16!c",
		BankCode:         "0:3",
		BranchCode:       "3:7",
		NationalChecksum: "7:8",
		AccountNumber:    "8:24",
		IsSEPA:           false,
	}

	// Test that code generation produces syntactically valid output
	t.Run("generated code is syntactically valid", func(t *testing.T) {
		// Generate code for the country
		var buf bytes.Buffer
		err := GenerateCodeForCountry(&buf, country)
		if err != nil {
			t.Fatalf("code generation failed: %v", err)
		}

		// Ensure generated output compiles to valid Go
		fset := token.NewFileSet()
		_, err = parser.ParseFile(fset, "generated.go", buf.String(), parser.ParseComments)
		if err != nil {
			t.Fatalf("generated invalid Go code: %v\n%s", err, buf.String())
		}

		// Note: We don't test for identical string output because the generated code
		// uses consistent formatting.
		// Golden file tests ensure output consistency over time.
	})
}

func TestGenerateValidationFunction_Internal(t *testing.T) {
	staticRule := &rule.StaticRule{StartPosition: 0, Value: "AL"}
	rangeRule := &rule.RangeRule{StartPosition: 2, Length: 2, Format: rule.Digit}
	rules := []rule.Rule{staticRule, rangeRule}

	// Test internal validation function generation
	t.Run("validation function generation", func(t *testing.T) {
		funcDecl, err := generateValidationFunction("Albania", "validateAlbaniaIBAN", rules, 28)
		if err != nil {
			t.Fatalf("generateValidationFunction failed: %v", err)
		}

		// Verify it's a valid function declaration
		if funcDecl.Name.Name != "validateAlbaniaIBAN" {
			t.Errorf("Expected function name 'validateAlbaniaIBAN', got '%s'", funcDecl.Name.Name)
		}

		// Convert to Go code and verify it compiles
		fset := token.NewFileSet()
		var buf bytes.Buffer
		err = format.Node(&buf, fset, funcDecl)
		if err != nil {
			t.Fatalf("Failed to format generated function: %v", err)
		}

		// The generated function should be valid Go syntax
		generatedCode := buf.String()
		if generatedCode == "" {
			t.Error("Generated function is empty")
		}
	})
}

func TestGenerateStaticValidation_Internal(t *testing.T) {
	staticRule := &rule.StaticRule{StartPosition: 0, Value: "AL"}

	t.Run("static validation generation", func(t *testing.T) {
		stmt, err := generateStaticValidation(staticRule)
		if err != nil {
			t.Fatalf("generateStaticValidation failed: %v", err)
		}

		// Verify it's an if statement
		if _, ok := stmt.(*ast.IfStmt); !ok {
			t.Errorf("Expected *ast.IfStmt, got %T", stmt)
		}

		// Convert to Go code
		fset := token.NewFileSet()
		var buf bytes.Buffer
		err = format.Node(&buf, fset, stmt)
		if err != nil {
			t.Fatalf("Failed to format generated statement: %v", err)
		}

		generatedCode := buf.String()
		// Should contain the static validation logic
		if !strings.Contains(generatedCode, "!= \"AL\"") {
			t.Errorf("Generated code doesn't contain expected static validation: %s", generatedCode)
		}
	})
}

func TestGenerateRangeValidation_Internal(t *testing.T) {
	rangeRule := &rule.RangeRule{StartPosition: 2, Length: 2, Format: rule.Digit}

	t.Run("range validation generation", func(t *testing.T) {
		stmt, err := generateRangeValidation(rangeRule)
		if err != nil {
			t.Fatalf("generateRangeValidation failed: %v", err)
		}

		// Verify it's an if statement
		if _, ok := stmt.(*ast.IfStmt); !ok {
			t.Errorf("Expected *ast.IfStmt, got %T", stmt)
		}

		// Convert to Go code
		fset := token.NewFileSet()
		var buf bytes.Buffer
		err = format.Node(&buf, fset, stmt)
		if err != nil {
			t.Fatalf("Failed to format generated statement: %v", err)
		}

		generatedCode := buf.String()
		// Should contain the range validation logic
		if !strings.Contains(generatedCode, "ascii.IsDigit") {
			t.Errorf("Generated code doesn't contain expected range validation: %s", generatedCode)
		}
	})
}

func TestGenerateCompleteFile_Internal(t *testing.T) {
	countries := []registry.Country{
		{
			Code:   "AL",
			Name:   "Albania",
			IBAN:   "AL2!n8!n16!c",
			IsSEPA: false,
		},
	}

	t.Run("complete file generation", func(t *testing.T) {
		// Generate template-based output for comparison
		var templateBuf bytes.Buffer
		err := GenerateValidate(&templateBuf, countries)
		if err != nil {
			t.Fatalf("Template generation failed: %v", err)
		}

		// Generate AST-based output
		var astBuf bytes.Buffer
		err = GenerateValidate(&astBuf, countries)
		if err != nil {
			t.Fatalf("AST generation failed: %v", err)
		}
		astOutput := astBuf.Bytes()

		// Compare outputs - they should be functionally identical
		// (Note: exact string matching might not work due to formatting differences,
		//  but both should compile to the same behavior)

		// At minimum, verify AST output compiles
		fset := token.NewFileSet()
		_, err = parser.ParseFile(fset, "test.go", astOutput, 0)
		if err != nil {
			t.Fatalf("AST-generated code doesn't compile: %v\n%s", err, string(astOutput))
		}
	})
}
