package main
//名字前面加aa 为了保证第一个初始化
import (
    //https://github.com/go-yaml/yaml
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "github.com/urfave/cli"
    "path/filepath"
    "os"
    "github.com/fatih/color"
    "os/exec"
)

var config Configuration

//注：约定优于配置
//定义 struct中 首字母大写，yaml中 首字母小写 单词采用同一个
type Configuration struct {
    LogLevel string
    Compress bool `yaml:"compress"` //yaml：yaml格式 enabled：属性的为enabled
    Convert struct {
        Enable      bool
        SuffixAllow []string `yaml:",flow"`
    }
    Watermark struct {
        Enable bool
        Path   string
    }
    Security struct {
        UserPass struct {
            Enable       bool
            Password2Add string
        }
        OwnerPass struct {
            Enable            bool
            Password2Add      string
            Printing          bool
            FullPrintQuality  bool
            Modify            bool
            Annotate          bool
            FillForms         bool
            RotateInsert      bool
            ExtractGraphics   bool
            DisabilityExtract bool
        }
    }

    Enabled bool   `yaml:"enabled"` //yaml：yaml格式 enabled：属性的为enabled
    Path    string `yaml:"path"`
    Path2   string
}

func (c *Configuration) getConf(path string) *Configuration {
    yamlFile, err := ioutil.ReadFile(path)
    if err != nil {
        //fmt.Printf("yamlFile.Get err   #%v ", err)
        color.Red("Yaml config file read error #%v ", err)
    }
    err = yaml.Unmarshal(yamlFile, c)
    if err != nil {
        //fmt.Fatalf("Unmarshal: %v", err)
        color.Red("Unmarshal: %v", err)
    }
    return c
}

func (c *Configuration) writeConf(path string) (err error) {
    d, err := yaml.Marshal(&config)
    if err != nil {
        color.Red("error: %v", err)
        return
    }
    err = ioutil.WriteFile(path, d, 0644)
    if err != nil {
        color.Red("error: %v", err)
    }
    //yamlFile, err := ioutil.ReadFile(path)
    //if err != nil {
    //    fmt.Printf("yamlFile.Get err   #%v ", err)
    //}
    //err = yaml.Unmarshal(yamlFile, c)
    //if err != nil {
    //    //fmt.Fatalf("Unmarshal: %v", err)
    //    fmt.Errorf("Unmarshal: %v", err)
    //}
    return err
}

func init() {
    file1, _ := exec.LookPath("./"+CONFIGFILENAME)
    path1, _ := filepath.Abs(file1)
    //color.Green(path1)
    //fmt.Println(path1)
    config.getConf(path1)
}

//用于根据定义的结构，生成yaml模板
func createYamlFile(path string) {
    d, err := yaml.Marshal(&config)
    if err != nil {
        color.Red("error: %v", err)
    }
    //fmt.Printf("--- config dump:\n%s\n\n", string(d))
    color.Blue("--- config dump:\n%s\n\n", string(d))
    config.writeConf(path)
    //config.writeConf("conf_template.yaml")
}

//生成conf.yaml模板
func createTemplate() cli.Command {
    command := cli.Command{
        Name:        "generateTemplate",
        Aliases:     []string{"tpl"},
        Category:    "Tools",
        Usage:       "Create conf.yaml template",
        UsageText:   "Example: doc2pdf generateTemplate ./conf.yaml ",
        Description: "Create conf.yaml template",
        ArgsUsage:   " <filename>",
        //Flags: []cli.Flag{
        //    cli.BoolFlag{
        //        Name:   "show,s",
        //        Usage:  "show current password",
        //        Hidden: true,
        //    },
        //},
        Action: func(c *cli.Context) error {
            if c.NArg() > 0 {
                destPath := c.Args().First()
                f, err := os.Stat(destPath)
                canGenFlag := true
                if os.IsExist(err) {
                    //judge if this is a path
                    if f.IsDir() {
                        destPath = filepath.Join(destPath, "conf_template.yaml")
                        _, err = os.Stat(destPath)
                        if os.IsExist(err) {
                            canGenFlag = false
                        }
                    }
                    canGenFlag = false
                }
                if canGenFlag {
                    createYamlFile(destPath)
                }else{
                    color.Red("Error occurred,check args.")
                }
            } else {
                createYamlFile("conf_template.yaml")
            }
            //calcBase(c.Args().First(), c.Bool("debase"))
            return nil
        },
    }
    return command
}

//func main() {
//    fmt.Println(config)
//    createYamlFile()>
//}
