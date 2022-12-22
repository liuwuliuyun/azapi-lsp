package complete

import (
	"fmt"
	"github.com/Azure/azapi-lsp/internal/langserver/handlers/tfschema"
	ilsp "github.com/Azure/azapi-lsp/internal/lsp"
	"github.com/Azure/azapi-lsp/internal/parser"
	lsp "github.com/Azure/azapi-lsp/internal/protocol"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"log"
)

func CandidatesAtPos(data []byte, filename string, pos hcl.Pos, logger *log.Logger) []lsp.CompletionItem {
	file, _ := hclsyntax.ParseConfig(data, filename, hcl.InitialPos)

	body, isHcl := file.Body.(*hclsyntax.Body)
	if pos.Column != 1 {
		return nil
	}
	if !isHcl {
		logger.Printf("file is not hcl")
		return nil
	}
	block := parser.LastBlock(body, pos)
	candidateList := make([]lsp.CompletionItem, 0)
	if block != nil && len(block.Labels) != 0 {
		resourceName := fmt.Sprintf("%s.%s", block.Type, block.Labels[0])
		resource := tfschema.GetResourceSchema(resourceName)

		predictionResourceList := getNextPossibleResources(resource)

		editRange := lsp.Range{
			Start: ilsp.HCLPosToLSP(pos),
			End:   ilsp.HCLPosToLSP(pos),
		}

		candidateList = append(candidateList, tfschema.RecommendedResources(predictionResourceList, editRange)...)

	}

	return candidateList
}

func editRangeFromExprRange(expression hclsyntax.Expression, pos hcl.Pos) lsp.Range {
	expRange := expression.Range()
	if expRange.Start.Line != expRange.End.Line && expRange.End.Column == 1 && expRange.End.Line-1 == pos.Line {
		expRange.End = pos
	}
	return ilsp.HCLRangeToLSP(expRange)
}

func getNextPossibleResources(resource *tfschema.Resource) []string {
	sample1 := "azurerm_storage_account"
	sample2 := "azurerm_network_interface"
	sample3 := "azurerm_virtual_network"
	return []string{sample1, sample2, sample3}
}
