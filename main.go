package main

import (
	"bytes"
	"fmt"
	"net/smtp"

	"os"
	"text/template"
  // "time"

	"github.com/dustin/go-humanize"
	// "golang.org/x/text/date"
)

type Payroll struct {
  CheckDate string
  Href string
  GrandTotal string
  DirectDeposit string
  SafeHarbor string
  EEDeferral string
  EERoth string
  RetirementTotal string
  EETax string
  EROasdi string
  ERMed string
  TaxWithholdTotal string
  Garnishments string
  PayRates []string
  Benefits []string
  SigningBonuses []string 
  OtherBonuses []string
  BridgeLoanPmtCount string
  BridgeLoanBalance string
  NewWorkers []string
  Terminations []string
  ExpenseReimbAmount string
  BankChanges []string
  TaxChanges []string
  AddressChanges []string
}

var tpl *template.Template

func main() {
checkDate := "12.15.21"
href := "https://ventussolutions.sharepoint.com/:b:/r/sites/AccountingDocs/Payroll%20Register%20Approvals/12.15.21%20PreProcess%20Register.pdf?csf=1&web=1&e=o55iPE" //updated for 12.15
dd:= 169915.63
safe:=4593.05
eeDef:= 15142.98
eeRoth:= 2350.00
eeTax:= 40952.85
erSoc:= 7899.57
erMed:= 2308.19
garn:=774.25
count:= findCount(checkDate)         
exp:= 2195.90
payrates:=[]string{"none"}
benefits:=[]string{"none"}
signing:=[]string{"none"}
bonuses:=[]string{"none"}
workers:=[]string{"none"}
terms:=[]string{"none"}
bank:=[]string{"none"}
taxChanges:= []string{"Craig Welling – changed additional amount withheld for taxes"}
address:=[]string{"Steve Harrison – both W2 & 1099","Joe Monaghan"}
loanBalance:= 36340 - (1651.81 * count) 
taxTotal:= eeTax + erSoc + erMed
retTotal:= safe + eeDef + eeRoth
grand:= dd+ retTotal+taxTotal+garn

  //Collect Payroll Info
 p:= Payroll{
  CheckDate: checkDate,
  Href: href,
  DirectDeposit:      ma(dd),
  SafeHarbor:         ma(safe),
  EEDeferral:         ma(eeDef), 
  EERoth:             ma(eeRoth), 
  EETax:              ma(eeTax),
  EROasdi:            ma(erSoc),
  ERMed:              ma(erMed),
  Garnishments:       ma(garn),  
  PayRates:           payrates,
  Benefits:           benefits,
  SigningBonuses:    signing,
  OtherBonuses:      bonuses,
  BridgeLoanPmtCount: ma(count),
  BridgeLoanBalance: ma(loanBalance),
  NewWorkers:         workers,
  Terminations:       terms,
  ExpenseReimbAmount: ma(exp),
  BankChanges:        bank,
  TaxChanges:         taxChanges,
  AddressChanges:     address,
  RetirementTotal:    ma(retTotal),
  TaxWithholdTotal:   ma(taxTotal),
  GrandTotal:         ma(grand),
 }  

  // Sender data.
  from := "assia.marcus@gmail.com"
  password := "@Race11c"

  // Receiver email address.
  to := []string{
    "marcus@monumentcpa.com",
    // "marcus@massia.net",
  }

  // smtp server configuration.
  smtpHost := "smtp.gmail.com"
  smtpPort := "587"

  // Authentication.
  auth := smtp.PlainAuth("", from, password, smtpHost)



  tpl = template.Must(template.ParseGlob("templates/*"))

  var body bytes.Buffer

  mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
  body.Write([]byte(fmt.Sprintf("Subject: CEO ACTION NEEDED: " + checkDate + " Payroll \n%s\n\n", mimeHeaders)))

  tpl.Execute(&body,  p)

  tpl.ExecuteTemplate(os.Stdout, "index.html", p)


  // Sending email
  err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println("Email Sent!")
}

func ma (num float64) string{
  return humanize.FormatFloat("#,###.##", num)
}

func findCount(chDate string) float64{
  var answer int
  paydates:=[]string{
    "10.15.21",
    "10.29.21",
    "11.15.21",
    "11.30.21",
    "12.15.21",
    "12.31.21",
  }
  for i, v := range paydates {
    if v == chDate{
      answer = i + 17
    }
  }
  return float64(answer)
}

