package deletedomain

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

func DeleteDomain(workdir string, gomod string) error {

	var err error

	//* Display section category
	category := section.CategorySection()
	domain := section.GetDomainSection(workdir, category)

	views.DisplayResetWithMessage("please wait", constants.ColorGreen)
	var domainCap = strings.Replace(domain, domain[0:1], strings.ToUpper(domain[0:1]), 1)

	//* Make temp folder
	tempPath := fmt.Sprintf("%s/temp", workdir)
	utils.Mkdir(tempPath)

	//* Make backup folder
	backupPath := fmt.Sprintf("%s/backup", tempPath)
	utils.Mkdir(backupPath)

	//* Backup data
	err = cp.Copy(workdir+"/internal/"+category+"/"+domain, backupPath+"/"+domain)
	if err != nil {
		return errors.New("failed to backup domain")
	}
	err = cp.Copy(workdir+"/internal/wrapper/repository/"+category+"/"+category+".go", backupPath+"/repository"+category+".go")
	if err != nil {
		return errors.New("failed to backup repository wrapper")
	}
	err = cp.Copy(workdir+"/internal/wrapper/usecase/"+category+"/"+category+".go", backupPath+"/usecase"+category+".go")
	if err != nil {
		return errors.New("failed to backup usecase wrapper")
	}
	err = cp.Copy(workdir+"/internal/wrapper/handler/"+category+"/"+category+".go", backupPath+"/handler"+category+".go")
	if err != nil {
		return errors.New("failed to backup handler wrapper")
	}

	//* Make go-tools folder
	goToolsPath := fmt.Sprintf("%s/go-tools-temp", tempPath)
	utils.Mkdir(goToolsPath)

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
		repoImportContent := constants.DeleteWrapperImportContent(repoContent, "repo", gomod, category, domain)
		repoStructContent := constants.DeleteWrapperStructContent(repoImportContent, "repo", category, domain, domainCap)
		repoFunctionContent := constants.DeleteWrapperFunctionContent(repoStructContent, "repo", category, domain, domainCap)

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
		usecaseImportContent := constants.DeleteWrapperImportContent(usecaseContent, "uc", gomod, category, domain)
		usecaseStructContent := constants.DeleteWrapperStructContent(usecaseImportContent, "uc", category, domain, domainCap)
		usecaseFunctionContent := constants.DeleteWrapperFunctionContent(usecaseStructContent, "uc", category, domain, domainCap)

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
		handlerImportContent := constants.DeleteWrapperImportContent(handlerContent, "handler", gomod, category, domain)
		handlerStructContent := constants.DeleteWrapperStructContent(handlerImportContent, "handler", category, domain, domainCap)
		handlerFunctionContent := constants.DeleteWrapperFunctionContent(handlerStructContent, "handler", category, domain, domainCap)

		err = utils.WriteFile(goToolsPath+"/handler"+category+".go", handlerFunctionContent)
		if err != nil {
			return err
		}

		return nil
	}()

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
	err = utils.RemoveContents(workdir + "/internal/" + category + "/" + domain)
	if err != nil {
		return errors.New("failed to delete domain")
	}

	if err == nil {
		err = utils.RemoveContents(tempPath)
		if err != nil {
			return errors.New("failed to delete content in temp folder")
		}
	}

	views.DisplayResetWithMessage("Thank you", constants.ColorGreen)

	return nil
}
