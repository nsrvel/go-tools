package createdomain

import (
	"errors"
	"fmt"
	"strings"

	"github.com/nsrvel/go-tools/components/section"
	"github.com/nsrvel/go-tools/constants"
	"github.com/nsrvel/go-tools/utils"
	"github.com/nsrvel/go-tools/views"
	cp "github.com/otiai10/copy"
)

func CreateDomain(workdir string, gomod string) error {

	var err error

	//* Display section category
	category := section.CategorySection()
	domain := section.GetNewDomainSection(workdir, category)

	views.DisplayResetWithMessage("please wait", constants.ColorGreen)

	//* Make temp folder
	tempPath := fmt.Sprintf("%s/temp", workdir)
	utils.Mkdir(tempPath)

	//* Make go-tools folder
	goToolsPath := fmt.Sprintf("%s/go-tools-temp", tempPath)
	utils.Mkdir(goToolsPath)

	var domainCap = strings.Replace(domain, domain[0:1], strings.ToUpper(domain[0:1]), 1)

	err = func() error {

		//* Make domain folder
		domainPath := fmt.Sprintf("%s/%s", goToolsPath, domain)
		err = utils.Mkdir(domainPath)
		if err != nil {
			return err
		}

		err = utils.Mkdir(domainPath + "/models")
		if err != nil {
			return err
		}
		err = utils.WriteFile(domainPath+"/models/db_scan.go", constants.ContentPackage("models"))
		if err != nil {
			return err
		}
		err = utils.WriteFile(domainPath+"/models/request.go", constants.ContentPackage("models"))
		if err != nil {
			return err
		}
		err = utils.WriteFile(domainPath+"/models/response.go", constants.ContentPackage("models"))
		if err != nil {
			return err
		}

		err = utils.Mkdir(domainPath + "/mock")
		if err != nil {
			return err
		}

		err = utils.Mkdir(domainPath + "/repository")
		if err != nil {
			return err
		}
		contentRepositoryFile := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(constants.ContentRepositoryFile, "domain-lower", domain), "@domain-cap", domainCap), "@gomod", gomod)
		err = utils.WriteFile(domainPath+"/repository/pg_repository.go", contentRepositoryFile)
		if err != nil {
			return err
		}
		err = utils.WriteFile(domainPath+"/repository/pg_repository_test.go", constants.ContentPackage("repository"))
		if err != nil {
			return err
		}
		err = utils.WriteFile(domainPath+"/repository/sql_queries.go", constants.ContentPackage("repository"))
		if err != nil {
			return err
		}

		err = utils.Mkdir(domainPath + "/usecase")
		if err != nil {
			return err
		}
		contentUsecaseFile := strings.ReplaceAll(strings.ReplaceAll(constants.ContentUsecaseFile, "@domain-cap", domainCap), "@gomod", gomod)
		err = utils.WriteFile(domainPath+"/usecase/usecase.go", contentUsecaseFile)
		if err != nil {
			return err
		}
		err = utils.WriteFile(domainPath+"/usecase/usecase_test.go", constants.ContentPackage("usecase"))
		if err != nil {
			return err
		}

		err = utils.Mkdir(domainPath + "/delivery")
		if err != nil {
			return err
		}
		contentHandlerFile := strings.ReplaceAll(strings.ReplaceAll(constants.ContentHandlerFile, "@domain-cap", domainCap), "@gomod", gomod)
		err = utils.WriteFile(domainPath+"/delivery/handler.go", contentHandlerFile)
		if err != nil {
			return err
		}
		err = utils.WriteFile(domainPath+"/delivery/handler_test.go", constants.ContentPackage("delivery"))
		if err != nil {
			return err
		}

		contentRoutesFile := strings.ReplaceAll(strings.ReplaceAll(constants.ContentRoutesFile, "@domain-lower", domain), "@gomod", gomod)
		err = utils.WriteFile(domainPath+"/routes.go", contentRoutesFile)
		if err != nil {
			return err
		}

		return nil
	}()

	if err != nil {
		err := utils.RemoveContents(goToolsPath)
		if err != nil {
			return errors.New("failed to delete content in temp folder")
		}
		return errors.New("failed to create domain content")
	}

	repoPath := workdir + "/internal/wrapper/repository/" + category
	usecasePath := workdir + "/internal/wrapper/usecase/" + category
	handlerPath := workdir + "/internal/wrapper/handler/" + category

	err = func() error {

		//* Repository
		isExist1, err := utils.CheckPathIfExist(repoPath + "/" + category + ".go")
		if !isExist1 {
			return err
		}
		repoContent, err := utils.ReadFile(repoPath + "/" + category + ".go")
		if err != nil {
			return err
		}
		repoImportContent := constants.UpdateWrapperImportContent(repoContent, "repo", gomod, category, domain)
		repoStructContent := constants.UpdateWrapperStructContent(repoImportContent, "repo", category, domain, domainCap)
		repoFunctionContent := constants.UpdateWrapperFuncContent(repoStructContent, "repo", category, domain, domainCap)

		err = utils.WriteFile(goToolsPath+"/repository"+category+".go", repoFunctionContent)
		if err != nil {
			return err
		}

		//* Usecase
		isExist2, err := utils.CheckPathIfExist(usecasePath + "/" + category + ".go")
		if !isExist2 {
			return err
		}
		usecaseContent, err := utils.ReadFile(usecasePath + "/" + category + ".go")
		if err != nil {
			return err
		}
		usecaseImportContent := constants.UpdateWrapperImportContent(usecaseContent, "uc", gomod, category, domain)
		usecaseStructContent := constants.UpdateWrapperStructContent(usecaseImportContent, "uc", category, domain, domainCap)
		usecaseFunctionContent := constants.UpdateWrapperFuncContent(usecaseStructContent, "uc", category, domain, domainCap)

		err = utils.WriteFile(goToolsPath+"/usecase"+category+".go", usecaseFunctionContent)
		if err != nil {
			return err
		}

		//* Handler
		isExist3, err := utils.CheckPathIfExist(handlerPath + "/" + category + ".go")
		if !isExist3 {
			return err
		}
		handlerContent, err := utils.ReadFile(handlerPath + "/" + category + ".go")
		if err != nil {
			return err
		}
		handlerImportContent := constants.UpdateWrapperImportContent(handlerContent, "handler", gomod, category, domain)
		handlerStructContent := constants.UpdateWrapperStructContent(handlerImportContent, "handler", category, domain, domainCap)
		handlerFunctionContent := constants.UpdateWrapperFuncContent(handlerStructContent, "handler", category, domain, domainCap)

		err = utils.WriteFile(goToolsPath+"/handler"+category+".go", handlerFunctionContent)
		if err != nil {
			return err
		}

		return nil
	}()

	if err != nil {
		err := utils.RemoveContents(goToolsPath)
		if err != nil {
			return errors.New("failed to delete content in temp folder")
		}
		return errors.New("failed to update wrapper content")
	}

	//* Copy data
	err = cp.Copy(goToolsPath+"/"+domain, workdir+"/internal/"+category+"/"+domain)
	if err != nil {
		return errors.New("failed to copy domain")
	}
	err = cp.Copy(goToolsPath+"/repository"+category+".go", workdir+"/internal/wrapper/repository/"+category+"/"+category+".go")
	if err != nil {
		return errors.New("failed to copy repository wrapper")
	}
	err = cp.Copy(goToolsPath+"/usecase"+category+".go", workdir+"/internal/wrapper/usecase/"+category+"/"+category+".go")
	if err != nil {
		return errors.New("failed to copy usecase wrapper")
	}
	err = cp.Copy(goToolsPath+"/handler"+category+".go", workdir+"/internal/wrapper/handler/"+category+"/"+category+".go")
	if err != nil {
		return errors.New("failed to copy handler wrapper")
	}

	err = utils.RemoveContents(tempPath)
	if err != nil {
		return errors.New("failed to delete content in temp folder")
	}

	views.DisplayResetWithMessage("Thank you", constants.ColorGreen)

	return nil
}
