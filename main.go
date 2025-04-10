package main

import (
    "encoding/json"
	"flag"
    "fmt"
    "io/ioutil"
    "net/http"
    "strings"
)

type USGSResponse struct {
    Value struct {
        TimeSeries []struct {
            SourceInfo struct {
                SiteName string `json:"siteName"`
                SiteCode []struct {
                    Value string `json:"value"`
                } `json:"siteCode"`
            } `json:"sourceInfo"`
            Variable struct {
                VariableCode []struct {
                    Value string `json:"value"`
                } `json:"variableCode"`
                VariableName        string `json:"variableName"`
                VariableDescription string `json:"variableDescription"`
            } `json:"variable"`
            Values []struct {
                Qualifier []struct {
                    QualifierCode string `json:"qualifierCode"`
                    QualifierDescription string `json:"qualifierDescription"`
                } `json:"qualifier"`
                Method []struct {
                    MethodDescription string `json:"methodDescription"`
                } `json:"method"`
            } `json:"values"`
        } `json:"timeSeries"`
    } `json:"value"`
}

func inspectUSGSSitePayload(siteID, paramIDs string, multiMethod bool) error {
    var url string
    if paramIDs == "" {
        url = fmt.Sprintf("https://waterservices.usgs.gov/nwis/iv/?format=json&site=%s", siteID)
    } else {
        url = fmt.Sprintf("https://waterservices.usgs.gov/nwis/iv/?format=json&site=%s&vars=%s", siteID, paramIDs)
    }
	fmt.Println("Getting timeseries data order from endpoint:", url, "\n")
    
    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return err
    }

    var data USGSResponse
    if err := json.Unmarshal(body, &data); err != nil {
        return err
    }

    if len(data.Value.TimeSeries) == 0 {
        return fmt.Errorf("no time series data found")
    }

    ts := data.Value.TimeSeries[0]
    fmt.Printf("%s (%s)\n", ts.SourceInfo.SiteName, ts.SourceInfo.SiteCode[0].Value)
    fmt.Println(strings.Repeat("=", 120))

    for i, d := range data.Value.TimeSeries {
        varID := d.Variable.VariableCode[0].Value
        varName := d.Variable.VariableName

        if (multiMethod) {
            for j, v := range d.Values {
                qualifierDescription := v.Qualifier[0].QualifierCode
                methodDescription := v.Method[0].MethodDescription
                fmt.Printf("%d | %s | %s | %d | %s | %s\n", i, varID, varName, j, qualifierDescription, methodDescription)
            }
        } else {
            fmt.Printf("%d | %s | %s\n", i, varID, varName)
        }
    }

    return nil
}

func main() {
	siteId := flag.String("site", "", "USGS site ID")
    paramIDs := flag.String("params", "", "USGS param IDs (comma separated list)")
    multiMethod := flag.Bool("multiMethod", false, "Check for multiple methods")
	flag.Parse()

    if err := inspectUSGSSitePayload(*siteId, *paramIDs, *multiMethod); err != nil {
        fmt.Printf("Error: %v\n", err)
    }
}
