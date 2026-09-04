package main

import (
	"GOCups/utils"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
)

const defaultPort = 631

var y = color.New(color.FgYellow).SprintFunc()
var g = color.New(color.FgGreen).SprintFunc()

type Client struct {
	UrlConf     string
	UrlPrinters string
	Logger      utils.Logger
}

type PrinterInfo struct {
	Name        string
	Description string
	Location    string
	Model       string
	Status      string
	url         string
}

type JobInfo struct {
	ID    string
	Name  string
	User  string
	Size  string
	Pages int
	State string
}

func NewClient() *Client {
	return &Client{
		UrlConf:     "/admin/conf/cupsd.conf",
		UrlPrinters: "/printers",
		Logger:      *utils.NewPrint(),
	}
}

func (client *Client) printJobs(printer *PrinterInfo) {
	body, err := utils.Get(fmt.Sprintf("%s?which_jobs=all", printer.url))
	if err != nil {
		logger.Error(err)
	} else {
		jobs := client.processJob(body)
		fmt.Printf("%s: %s\n", y("Jobs"), g(len(jobs)))
		if len(jobs) > 0 {
			table := tablewriter.NewTable(os.Stdout)
			table.Header("ID", "Nombre", "Usuario", "Tamaño", "Paginas", "Estado")
			table.Bulk(jobs)
			table.Render()
		}

		logger.Info("======================================================")

	}
}

func (client *Client) PrintInformation(printers *[]PrinterInfo) {

	for _, printer := range *printers {
		fmt.Printf("%s: %s\n", y("Nombre"), g(printer.Name))
		fmt.Printf("%s: %s\n", y("Descripción"), g(printer.Description))
		fmt.Printf("%s: %s\n", y("Ubicación"), g(printer.Location))
		fmt.Printf("%s: %s\n", y("Modelo/Marca"), g(printer.Model))
		fmt.Printf("%s: %s\n", y("Estado"), g(printer.Status))
		client.printJobs(&printer)

	}
}

func (client *Client) EnumPrinters(host string, optionalPort ...int) error {
	port := defaultPort
	if len(optionalPort) > 0 && optionalPort[0] > 0 {
		port = optionalPort[0]
	}

	if host == "" {
		return fmt.Errorf("La IP del host no debe estar vacia")
	}

	client.UrlPrinters = fmt.Sprintf("http://%s:%d%s", host, port, client.UrlPrinters)

	client.Logger.Info("Enumerando impresoras de %s:%d", host, port)

	body, err := utils.Get(client.UrlPrinters)

	if err != nil {
		return err
	}

	printers := client.processPrinters(body)
	client.PrintInformation(&printers)

	return nil
}

func (c *Client) processJob(data []byte) []JobInfo {
	var jobs []JobInfo
	reader := bytes.NewReader(data)
	doc, err := goquery.NewDocumentFromReader(reader)

	if err != nil {
		logger.Error(err)
	}

	cups_body := doc.Find(".list")
	if cups_body.Length() > 0 {

		cups_body.Find("tbody tr").Each(func(i int, row *goquery.Selection) {
			cells := row.Find("td")
			if cells.Length() == 0 {
				return
			}

			rawPages := strings.TrimSpace(cells.Eq(4).Text())
			pages, err := strconv.Atoi(rawPages)

			if err != nil {
				logger.Error(err)
			}

			job := JobInfo{
				ID:    strings.TrimSpace(cells.Eq(0).Text()),
				Name:  strings.TrimSpace(cells.Eq(1).Text()),
				User:  strings.TrimSpace(cells.Eq(2).Text()),
				Size:  strings.TrimSpace(cells.Eq(3).Text()),
				State: strings.TrimSpace(cells.Eq(5).Text()),
				Pages: pages,
			}

			jobs = append(jobs, job)

		})
	}

	return jobs

}

func (c *Client) processPrinters(data []byte) []PrinterInfo {
	var printers []PrinterInfo
	reader := bytes.NewReader(data)

	doc, err := goquery.NewDocumentFromReader(reader)

	if err != nil {
		logger.Error(err)
	}

	cups_body := doc.Find(".list")
	if cups_body.Length() > 0 {

		cups_body.Find("tbody tr").Each(func(i int, row *goquery.Selection) {

			cells := row.Find("td")

			if cells.Length() == 0 {
				return
			}

			printername := strings.TrimSpace(cells.Eq(0).Text())
			printer := PrinterInfo{
				Name:        printername,
				Description: strings.TrimSpace(cells.Eq(1).Text()),
				Location:    strings.TrimSpace(cells.Eq(2).Text()),
				Model:       strings.TrimSpace(cells.Eq(3).Text()),
				Status:      strings.TrimSpace(cells.Eq(4).Text()),
				url:         fmt.Sprintf("%s/%s", c.UrlPrinters, printername),
			}

			printers = append(printers, printer)

		})
	}

	return printers
}
