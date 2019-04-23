package main

import (
	"bufio"
	"encoding/xml"
	"errors"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

var apiKey string
var format string
var minThumbnailWidth int

// Photo structure.
type Photo struct {
	Title     string
	Link      string
	Sizes     []Size
	Thumbnail Size
}

// Size structure.
type Size struct {
	Label  string
	Width  int
	Height int
	Source string
}

func main() {
	flag.Parse()

	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func loadFormat() error {
	if !strings.HasPrefix(format, "@") {
		return nil
	}

	file, err := ioutil.ReadFile(format[1:])

	if err != nil {
		return err
	}

	format = string(file)
	return nil
}

func run() error {
	if err := loadFormat(); err != nil {
		return err
	}

	apiKey = os.Getenv("FLICKR_API_KEY")

	if apiKey == "" {
		return errors.New(`missing environment variable: FLICKR_API_KEY
Get your key at the API Explorer: https://www.flickr.com/services/api/explore/`)
	}

	s := bufio.NewScanner(os.Stdin)
	ids := []string{}

	for s.Scan() {
		switch id, err := extractID(s.Text()); {
		case err != nil:
			return err
		default:
			ids = append(ids, id)
		}
	}

	if err := s.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	t, err := template.New("").Funcs(templateFuncs).Parse(format)

	if err != nil {
		return err
	}

	var infos = []Photo{}

	for _, id := range ids {
		info, err := getInfo(id)

		if err != nil {
			return err
		}

		infos = append(infos, *info)
	}

	switch format {
	case "":
		fmt.Printf("%+v\n", infos)
	default:
		t.Execute(os.Stdout, infos)
		fmt.Println("")
	}

	return nil
}

var templateFuncs = map[string]interface{}{
	"inc": func(i int) int {
		return i + 1
	},
}

func extractID(link string) (string, error) {
	parts := strings.Split(link, "/")

	for i := len(parts) - 1; i > 0; i-- {
		if _, err := strconv.ParseInt(parts[i], 10, 64); err == nil {
			return parts[i], nil
		}
	}

	return "", fmt.Errorf("can't find photo '%s'", link)
}

type infoContainer struct {
	XMLName xml.Name `xml:"rsp"`
	Stat    string   `xml:"stat,attr"`
	Photo   info     `xml:"photo"`
}

type info struct {
	XMLName xml.Name `xml:"photo"`
	ID      string   `xml:"id,attr"`
	Title   string   `xml:"title"`
	Owner   owner    `xml:"owner"`
}

type owner struct {
	XMLName   xml.Name `xml:"owner"`
	PathAlias string   `xml:"path_alias,attr"`
}

type sizesContainer struct {
	XMLName xml.Name `xml:"rsp"`
	Stat    string   `xml:"stat,attr"`
	Sizes   sizes    `xml:"sizes"`
}

type sizes struct {
	XMLName xml.Name `xml:"sizes"`
	Sizes   []size   `xml:"size"`
}

type size struct {
	XMLName xml.Name `xml:"size"`
	Label   string   `xml:"label,attr"`
	Width   int      `xml:"width,attr"`
	Height  int      `xml:"height,attr"`
	Source  string   `xml:"source,attr"`
}

// https://www.flickr.com/services/api/flickr.photos.getInfo.html
func getInfo(id string) (*Photo, error) {
	var q = url.Values{}

	q.Set("method", "flickr.photos.getInfo")
	q.Set("api_key", apiKey)
	q.Set("format", "rest")
	q.Set("photo_id", id)

	var u = fmt.Sprintf("https://api.flickr.com/services/rest/?%s", q.Encode())
	resp, err := http.Get(u)

	if err != nil {
		return nil, err
	}

	var i infoContainer

	xmlBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	if err = xml.Unmarshal(xmlBody, &i); err != nil {
		return nil, err
	}

	if i.Stat != "ok" {
		return nil, fmt.Errorf("response error failure: %s\n%s", i.Stat, string(xmlBody))
	}

	sizes, err := getSize(id)

	if err != nil {
		return nil, fmt.Errorf("unable to get sizes for image %s: %v", id, err)
	}

	return &Photo{
		Title:     i.Photo.Title,
		Link:      fmt.Sprintf("https://www.flickr.com/photos/%s/%s", i.Photo.Owner.PathAlias, i.Photo.ID),
		Sizes:     sizes,
		Thumbnail: pickThumbnail(sizes),
	}, nil
}

func pickThumbnail(sizes []Size) (choosen Size) {
	for _, s := range sizes {
		if choosen.Width < minThumbnailWidth {
			choosen = s
		}
	}

	return choosen
}

// https://www.flickr.com/services/api/explore/flickr.photos.getSizes
func getSize(id string) ([]Size, error) {
	var q = url.Values{}

	q.Set("method", "flickr.photos.getSizes")
	q.Set("api_key", apiKey)
	q.Set("format", "rest")
	q.Set("photo_id", id)

	var u = fmt.Sprintf("https://api.flickr.com/services/rest/?%s", q.Encode())
	resp, err := http.Get(u)

	if err != nil {
		return nil, err
	}

	var s sizesContainer

	xmlBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	if err = xml.Unmarshal(xmlBody, &s); err != nil {
		return nil, err
	}

	if s.Stat != "ok" {
		return nil, fmt.Errorf("response error failure: %s\n%s", s.Stat, string(xmlBody))
	}

	var sizes = []Size{}

	for _, ss := range s.Sizes.Sizes {
		sizes = append(sizes, Size{
			Label:  ss.Label,
			Width:  ss.Width,
			Height: ss.Height,
			Source: ss.Source,
		})
	}

	return sizes, nil
}

func init() {
	flag.StringVar(&format, "format", "", "Format the output using the given go template or use @file.html")
	flag.IntVar(&minThumbnailWidth, "min-thumbnail-width", 1600, "Minimum width for the thumbnail")
}
