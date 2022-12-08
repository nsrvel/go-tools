package constants

import (
	"strings"
)

func ContentRepoImport(gomod string, category string, domainLower string) string {
	return strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(`	@domain-lower	"@gomod/internal/@category/@domain-lower/repository"`, "@category", category), "@gomod", gomod), "@domain-lower", domainLower)
}

func ContentRepoStruct(domainLower string, domainCap string) string {
	return strings.ReplaceAll(strings.ReplaceAll(`	@domain-cap	@domain-lower.Repository
`, "@domain-lower", domainLower), "@domain-cap", domainCap)
}

func ContentRepoFunc(domainLower string, domainCap string) string {
	return strings.ReplaceAll(strings.ReplaceAll(`	@domain-cap:	@domain-lower.New@domain-capRepo(dbList),
	`, "@domain-lower", domainLower), "@domain-cap", domainCap)
}

func ContentUsecaseImport(gomod string, category string, domainLower string) string {
	return strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(`	@domain-lower	"@gomod/internal/@category/@domain-lower/usecase"`, "@category", category), "@gomod", gomod), "@domain-lower", domainLower)
}

func ContentUsecaseStruct(domainLower string, domainCap string) string {
	return strings.ReplaceAll(strings.ReplaceAll(`	@domain-cap	@domain-lower.Usecase
`, "@domain-lower", domainLower), "@domain-cap", domainCap)
}

func ContentUsecaseFunc(domainLower string, domainCap string) string {
	return strings.ReplaceAll(strings.ReplaceAll(`	@domain-cap:	@domain-lower.New@domain-capUsecase(repo, conf, dbList, log),
	`, "@domain-lower", domainLower), "@domain-cap", domainCap)
}

func ContentHandlerImport(gomod string, category string, domainLower string) string {
	return strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(`	@domain-lower	"@gomod/internal/@category/@domain-lower/delivery"`, "@category", category), "@gomod", gomod), "@domain-lower", domainLower)
}

func ContentHandlerStruct(domainLower string, domainCap string) string {
	return strings.ReplaceAll(strings.ReplaceAll(`	@domain-cap	@domain-lower.@domain-capHandler
`, "@domain-lower", domainLower), "@domain-cap", domainCap)
}

func ContentHandlerFunc(domainLower string, domainCap string) string {
	return strings.ReplaceAll(strings.ReplaceAll(`	@domain-cap:	@domain-lower.New@domain-capHandler(uc, conf, log),
	`, "@domain-lower", domainLower), "@domain-cap", domainCap)
}

func UpdateWrapperImportContent(file string, layer string, gomod string, category string, domainLower string) string {
	array1 := strings.Split(file, `import (`)
	if len(array1) < 1 && array1[1] != "" {
		return ""
	}
	array2 := strings.Split(array1[1], `)`)
	if len(array2) == 0 {
		return ""
	}

	content := array2[0]
	compared := `import (` + content + `)`
	var result string

	if layer == "repo" {
		result = strings.Replace(file, compared, `import (`+content+ContentRepoImport(gomod, category, domainLower)+`
)`, 1)
	} else if layer == "uc" {
		result = strings.Replace(file, compared, `import (`+content+ContentUsecaseImport(gomod, category, domainLower)+`
)`, 1)
	} else if layer == "handler" {
		result = strings.Replace(file, compared, `import (`+content+ContentHandlerImport(gomod, category, domainLower)+`
)`, 1)
	} else {
		return ""
	}

	return result
}

func UpdateWrapperStructContent(file string, layer string, category string, domainLower string, domainCap string) string {

	var keyword string
	if category == "general" {
		category = "General"
	} else if category == "core" {
		category = "Core"
	} else if category == "cms" {
		category = "CMS"
	} else {
		return ""
	}

	if layer == "repo" {
		keyword = category + "Repository"
	} else if layer == "uc" {
		keyword = category + "Usecase"
	} else if layer == "handler" {
		keyword = category + "Handler"
	} else {
		return ""
	}

	array1 := strings.Split(file, `type `+keyword+` struct {`)
	if len(array1) < 1 && array1[1] != "" {
		return ""
	}
	array2 := strings.Split(array1[1], `}`)
	if len(array2) == 0 {
		return ""
	}

	content := array2[0]
	compared := `type ` + keyword + ` struct {` + content + `}`
	var result string

	if layer == "repo" {
		result = strings.Replace(file, compared, `type `+keyword+` struct {`+content+ContentRepoStruct(domainLower, domainCap)+`}`, 1)
	} else if layer == "uc" {
		result = strings.Replace(file, compared, `type `+keyword+` struct {`+content+ContentUsecaseStruct(domainLower, domainCap)+`}`, 1)
	} else if layer == "handler" {
		result = strings.Replace(file, compared, `type `+keyword+` struct {`+content+ContentHandlerStruct(domainLower, domainCap)+`}`, 1)
	} else {
		return ""
	}
	return result
}

func UpdateWrapperFuncContent(file string, layer string, category string, domainLower string, domainCap string) string {

	var keyword string
	if category == "general" {
		category = "General"
	} else if category == "core" {
		category = "Core"
	} else if category == "cms" {
		category = "CMS"
	} else {
		return ""
	}

	if layer == "repo" {
		keyword = category + "Repository"
	} else if layer == "uc" {
		keyword = category + "Usecase"
	} else if layer == "handler" {
		keyword = category + "Handler"
	} else {
		return ""
	}

	array1 := strings.Split(file, `return `+keyword+`{`)
	if len(array1) < 1 && array1[1] != "" {
		return ""
	}
	array2 := strings.Split(array1[1], `}`)
	if len(array2) == 0 {
		return ""
	}

	content := array2[0]
	compared := `return ` + keyword + `{` + content + `}`
	var result string

	if content == "" {
		content = `
	`
	}

	if layer == "repo" {
		result = strings.Replace(file, compared, `return `+keyword+`{`+content+ContentRepoFunc(domainLower, domainCap)+`}`, 1)
	} else if layer == "uc" {
		result = strings.Replace(file, compared, `return `+keyword+`{`+content+ContentUsecaseFunc(domainLower, domainCap)+`}`, 1)
	} else if layer == "handler" {
		result = strings.Replace(file, compared, `return `+keyword+`{`+content+ContentHandlerFunc(domainLower, domainCap)+`}`, 1)
	} else {
		return ""
	}

	return result
}
