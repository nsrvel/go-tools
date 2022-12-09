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
