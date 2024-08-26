package main

import (
    "fmt"
    "github.com/xeipuuv/gojsonschema"
    goxml2json "github.com/basgys/goxml2json"
    "strings"
)

func main() {
    // Exemplo de conteúdo XML
    xmlContent := `<note>
        <to>Tove</to>
        <from>Jani</from>
        <heading>Reminder</heading>
        <body>Don't forget me this weekend!</body>
    </note>`

    // Converte o XML para JSON
    xmlReader := strings.NewReader(xmlContent)
    jsonContent, err := goxml2json.Convert(xmlReader)
    if err != nil {
        fmt.Printf("Erro ao converter XML para JSON: %s\n", err)
        return
    }

    // Definição do schema JSON (análogo ao XSD)
    schemaJSON := `{
        "type": "object",
        "properties": {
            "note": {
                "type": "object",
                "properties": {
                    "to": {"type": "string"},
                    "from": {"type": "string"},
                    "heading": {"type": "string"},
                    "body": {"type": "string"}
                },
                "required": ["to", "from", "heading", "body"]
            }
        }
    }`

    // Carrega o schema JSON
    schemaLoader := gojsonschema.NewStringLoader(schemaJSON)

    // Carrega o JSON gerado
    documentLoader := gojsonschema.NewStringLoader(jsonContent.String())

    // Valida o JSON contra o schema
    result, err := gojsonschema.Validate(schemaLoader, documentLoader)
    if err != nil {
        fmt.Printf("Erro na validação: %s\n", err)
        return
    }

    // Exibe o resultado da validação
    if result.Valid() {
        fmt.Println("O XML é válido.")
    } else {
        fmt.Println("O XML não é válido. Erros:")
        for _, desc := range result.Errors() {
            fmt.Printf("- %s\n", desc)
        }
    }
}
