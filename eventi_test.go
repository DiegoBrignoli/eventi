package eventi

import (
	"github.com/DiegoBrignoli/htmlutils" //Diego Brignoli
	"github.com/DiegoBrignoli/month"
	"golang.org/x/net/html"
	"net/http"
	"strconv"
	"testing"
)

//Tutti e 4
var test = Event{ //Come se fosse un oggetto
	Day:   1,
	Month: month.Dicembre,
	Year:  1299,
	Descrizione: []string{
		"Nella battaglia di Falconara l'esercito di Federico III di Sicilia sconfigge quello del Regno di Napoli comandato da Filippo I d'Angiò",
	},
}

//Gabriel Cenci + Norbis Andrea
func TestEvent_GetYear(t *testing.T) {
	resp, err := request(strconv.Itoa(test.Day), test.Month.String())
	if err != nil {
		t.Fatal("Error to do HTTP request", err)
	}

	sections, err := htmlutils.QuerySelector(resp, "section", "id", "mf-section-1")
	if err != nil {
		t.Fatal("Error to extract 'section'", err)
	}

	tagsUl, err := htmlutils.GetGeneralTags(sections[0], "ul")
	if err != nil {
		t.Fatal("Error to extract 'ul'", err)
	}

	tagLi, err := htmlutils.GetGeneralTags(tagsUl[0], "li")
	if err != nil {
		t.Fatal("Error to extract 'li'", err)
	}

	//create object event
	e := New(tagLi[0])
	err = e.GetYear()

	if e.Year != test.Year { //Se quello che ho ottenuto è diverso da quello che mi aspetto
		t.Error("Error not obtain", test.Year, "but obtain", e.Year)
	} else {
		t.Log("Year [OK]")
	}
}

//Federico Centurioni
func TestEvent_GetDescription(t *testing.T) {
	resp, err := request(strconv.Itoa(test.Day), test.Month.String())
	if err != nil {
		t.Fatal("Error to do HTTP request", err)
	}

	sections, err := htmlutils.QuerySelector(resp, "section", "id", "mf-section-1")
	if err != nil {
		t.Fatal("Error to extract 'section'", err)
	}

	tagsUl, err := htmlutils.GetGeneralTags(sections[0], "ul")
	if err != nil {
		t.Fatal("Error to extract 'ul'", err)
	}

	tagLi, err := htmlutils.GetGeneralTags(tagsUl[0], "li")
	if err != nil {
		t.Fatal("Error to extract 'li'", err)
	}

	//create object event
	e := New(tagLi[0])
	e.GetDescription()

	for i, descrizione := range e.Descrizione {
		if descrizione != test.Descrizione[0] { //Se quello che ho ottenuto è diverso da quello che mi aspetto
			t.Error("Error not obtain", test.Descrizione[i], "but obtain", descrizione)
		} else {
			t.Log("Description", i, "[OK]")
		}
	}
}

//Gabriel Cenci
func request(day, month string) (*html.Node, error) {

	resp, err := http.Get(BaseUrl + day + "_" + month) //https://www.wikipedia.org/wiki/31_ottobre
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseGlobal, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	return responseGlobal, nil
}


