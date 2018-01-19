package main

import (
    "fmt"
    "os"
    pdfcore "github.com/unidoc/unidoc/pdf/core"
    pdf "github.com/unidoc/unidoc/pdf/model"
)

//testEncrypt return if the pdf encrypted
//return true means the file is Encrypted,if errors happen the bool value is false
func testEncrypt(inputPath string) (bool, error) {
    f, err := os.Open(inputPath)
    if err != nil {
        return false, err
    }

    defer f.Close()

    pdfReader, err := pdf.NewPdfReader(f)
    if err != nil {
        return false, err
    }
    isEncrypted, err := pdfReader.IsEncrypted()
    if err != nil {
        return false, err
    }
    if isEncrypted {
        log.Infof("The PDF is already locked")
    }
    return isEncrypted, err
}

// printAccessInfo
// inputPath the input file
// password the password specified
func printAccessInfo(inputPath string, password string) (error) {
    f, err := os.Open(inputPath)
    if err != nil {
        return err
    }

    defer f.Close()

    pdfReader, err := pdf.NewPdfReader(f)
    if err != nil {
        return err
    }

    canView, perms, err := pdfReader.CheckAccessRights([]byte(password))
    if err != nil {
        return err
    }

    if !canView {
        log.Infof("%s - Cannot view - No access with the specified password", inputPath)
        //return nil
    }

    log.Infof("Input file %s", inputPath)
    log.Infof("Access Permissions: %+v", perms)
    log.Infof("--------")

    // Print a text summary of the flags.
    booltext := map[bool]string{false: "No", true: "Yes"}
    log.Infof("Printing allowed? - %s", booltext[perms.Printing])
    if perms.Printing {
        log.Infof("Full print quality (otherwise print in low res)? - %s", booltext[perms.FullPrintQuality])
    }
    log.Infof("Modifications allowed? - %s", booltext[perms.Modify])
    log.Infof("Allow extracting graphics? %s", booltext[perms.ExtractGraphics])
    log.Infof("Can annotate? - %s", booltext[perms.Annotate])
    if perms.Annotate {
        log.Infof("Can fill forms? - Yes")
    } else {
        log.Infof("Can fill forms? - %s", booltext[perms.FillForms])
    }
    log.Infof("Extract text, graphics for users with disabilities? - %s", booltext[perms.DisabilityExtract])

    return nil
}

func addPassword(inputfilepath string, outputPath string, userPass string, ownerPass string) error {
    pdfWriter := pdf.NewPdfWriter()

    permissions := pdfcore.AccessPermissions{}
    // Allow printing with low quality
    permissions.Printing = false
    permissions.FullPrintQuality = false
    // Allow modifications.
    permissions.Modify = false
    // Allow annotations.
    permissions.Annotate = false
    permissions.FillForms = false
    // Allow modifying page order, rotating pages etc.
    permissions.RotateInsert = false
    // Allow extracting graphics.
    permissions.ExtractGraphics = false
    // Allow extracting graphics (accessibility)
    permissions.DisabilityExtract = false

    encryptOptions := &pdf.EncryptOptions{}
    encryptOptions.Permissions = permissions

    //err := pdfWriter.Encrypt([]byte(ownerPass+"A"), []byte(ownerPass+"B"), encryptOptions)
    err := pdfWriter.Encrypt([]byte(userPass), []byte(ownerPass), encryptOptions)
    if err != nil {
        return err
    }

    f, err := os.Open(inputfilepath)
    if err != nil {
        return err
    }

    defer f.Close()

    pdfReader, err := pdf.NewPdfReader(f)
    if err != nil {
        return err
    }

    isEncrypted, err := pdfReader.IsEncrypted()
    if err != nil {
        return err
    }
    if isEncrypted {
        return fmt.Errorf("The PDF is already locked (need to unlock first)")
    }

    numPages, err := pdfReader.GetNumPages()
    if err != nil {
        return err
    }

    for i := 0; i < numPages; i++ {
        pageNum := i + 1

        page, err := pdfReader.GetPage(pageNum)
        if err != nil {
            return err
        }

        err = pdfWriter.AddPage(page)
        if err != nil {
            return err
        }
    }

    fWrite, err := os.Create(outputPath)
    if err != nil {
        return err
    }

    defer fWrite.Close()

    err = pdfWriter.Write(fWrite)
    if err != nil {
        return err
    }

    return nil
}
