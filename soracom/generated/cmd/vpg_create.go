package cmd

import (

  "encoding/json"
  "io/ioutil"

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var VpgCreateCmdPrimaryServiceName string




var VpgCreateCmdBody string


func init() {
  VpgCreateCmd.Flags().StringVar(&VpgCreateCmdPrimaryServiceName, "primary-service-name", "", "")



  VpgCreateCmd.Flags().StringVar(&VpgCreateCmdBody, "body", "", TR("cli.common_params.body.short_help"))


  VpgCmd.AddCommand(VpgCreateCmd)
}

var VpgCreateCmd = &cobra.Command{
  Use: "create",
  Short: TR("Create Virtual Private Gateway"),
  Long: TR(`VPGを新規作成する
`),
  RunE: func(cmd *cobra.Command, args []string) error {
    opt := &apiClientOptions{
      Endpoint: getSpecifiedEndpoint(),
      BasePath: "/v1",
      Language: getSelectedLanguage(),
    }

    ac := newAPIClient(opt)
    if v := os.Getenv("SORACOM_VERBOSE"); v != "" {
      ac.SetVerbose(true)
    }

    
    err := authHelper(ac, cmd, args)
    if err != nil {
      cmd.SilenceUsage = true
      return err
    }
    
    param, err := collectVpgCreateCmdParams()
    if err != nil {
      return err
    }

    result, err := ac.callAPI(param)
    if err != nil {
      cmd.SilenceUsage = true
      return err
    }

    return prettyPrintJSON(result)
  },
}

func collectVpgCreateCmdParams() (*apiParams, error) {
  
  body, err := buildBodyForVpgCreateCmd()
  if err != nil {
    return nil, err
  }
  

  return &apiParams{
    method: "POST",
    path: buildPathForVpgCreateCmd("/virtual_private_gateways"),
    query: buildQueryForVpgCreateCmd(),
    contentType: "application/json",
    body: body,
  }, nil
}

func buildPathForVpgCreateCmd(path string) string {
  
  
  
  
  
  return path
}

func buildQueryForVpgCreateCmd() string {
  result := []string{}
  
  
  

  

  
  return strings.Join(result, "&")
}


func buildBodyForVpgCreateCmd() (string, error) {
  if VpgCreateCmdBody != "" {
    if strings.HasPrefix(VpgCreateCmdBody, "@") {
      fname := strings.TrimPrefix(VpgCreateCmdBody, "@")
      bytes, err := ioutil.ReadFile(fname)
      if err != nil {
        return "", err
      }
      return string(bytes), nil
    } else if VpgCreateCmdBody == "-" {
      bytes, err := ioutil.ReadAll(os.Stdin)
      if err != nil {
        return "", err
      }
      return string(bytes), nil
    } else {
      return VpgCreateCmdBody, nil
    }
  }

  result := map[string]interface{}{}
  
  if VpgCreateCmdPrimaryServiceName != "" {
    result["primaryServiceName"] = VpgCreateCmdPrimaryServiceName
  }
  
  
  

  resultBytes, err := json.Marshal(result)
  if err != nil {
    return "", err
  }
  return string(resultBytes), nil
}
