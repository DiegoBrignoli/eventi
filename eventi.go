package eventi

import (
	"errors"
	"github.com/KiritoNya/htmlutils" //Diego Brignoli
	"github.com/DiegoBrignoli/month"
	strip "github.com/grokify/html-strip-tags-go"
	"golang.org/x/net/html"
	"strconv"
	"strings"
)

//Tutti e 4
type Event struct {
	Day         int //come fare "public Day"
	Month       month.Month
	Year        int
	Descrizione []string
	htmlNode    *html.Node //come fare "private htmlNode"
}

//Tutti e 4
const BaseUrl = "https://it.m.wikipedia.org/wiki/"

//Tutti e 4
func New(node *html.Node) *Event {
	return &Event{htmlNode: node}
}

//Norbis Andrea
func (e *Event) GetYear() error {

	tagsA, err := htmlutils.GetGeneralTags(e.htmlNode, "a")
	if err != nil {
		return err
	}

	data := string(htmlutils.GetNodeText(tagsA[0], "a"))
	data = strings.Replace(data, " ", "", -1) //Tolgo eventuali spazi che mandano in errore la atoi

	if strings.Contains(data, "a.C") {
		e.Year = -1
	} else {
		e.Year, err = strconv.Atoi(data)
		if err != nil {

			//Prendo il titolo che corrisponde (non sempre) alla data
			dataByte, err := htmlutils.GetValueAttr(tagsA[0], "a", "title")
			if err != nil {
				return err
			}

			data = strings.Replace(string(dataByte[0]), " ", "", -1) //Tolgo eventuali spazi che mandano in errore la atoi

			//Non gestisco i a.C
			if strings.Contains(data, "a.C") {
				e.Year = -1
			} else {

				//Riprovo l'atoi con il nuovo modo
				e.Year, err = strconv.Atoi(data)
				if err != nil {
					return errors.New("atoi error")
				}

			}
		}
	}
	return nil
}

//Federico Centurioni
func (e *Event) GetDescription() error {

	/*<ul>
		<li>
			<a href="/wiki/Helena_(Montana)" title="Helena (Montana)">Helena</a>
			" ( "
			<a href="/wiki/Montana" title="Montana">Montana</a>
			") viene fondata, dopo che quattro cercatori scoprono l'"
			<a href="/wiki/Oro" title="Oro">oro</a>
			" a ""
			<i><a href="/w/index.php?title=Last_Chance_Gulch&amp;action=edit&amp;redlink=1" class="new" title="Last Chance Gulch (la pagina non esiste)">Last Chance Gulch</a></i>"
		</li>
		<li>
			<a href="/wiki/Busto_Arsizio" title="Busto Arsizio">Busto Arsizio</a>
			" viene insignita del "
			<a href="/wiki/Titolo_di_citt%C3%A0" title="Titolo di città">titolo di città</a>
		</li>
		<li>
			"Fine della "
			<a href="/wiki/Seconda_guerra_dello_Schleswig" title="Seconda guerra dello Schleswig">Seconda guerra dello Schleswig</a>
			": il "
			<a href="/wiki/Holstein-Gottorp" title="Holstein-Gottorp">duca</a>
			<a href="/wiki/Federico_VIII_di_Schleswig-Holstein-Sonderburg-Augustenburg" title="Federico VIII di Schleswig-Holstein-Sonderburg-Augustenburg">Federico</a>
			" e la corona danese, riconoscono l'annessione a "
			<a href="/wiki/Regno_di_Prussia" title="Regno di Prussia">Prussia</a>
			" e "
			<a href="/wiki/Impero_austriaco" title="Impero austriaco">Austria</a>
			" di "
			<a href="/wiki/Schleswig_(regione)" title="Schleswig (regione)">Schleswig</a>
			", "
			<a href="/wiki/Holstein" title="Holstein">Holstein</a>
			" e "
			<a href="/wiki/Lauenburg/Elbe" title="Lauenburg/Elbe">Lauenburg</a>
		</li>
	</ul>*/

	numTagUl, _ := htmlutils.TagCount(e.htmlNode, "ul")

	if numTagUl > 0 { //Se l'item ha una lista di descrizioni

		tagsLi, err := htmlutils.GetGeneralTags(e.htmlNode, "li")
		if err != nil {
			return err
		}

		tagsLi = tagsLi[1:] //Tolgo se stesso che è anche lui un tag "li"

		for _, li := range tagsLi { //foreach
			//var str string
			stripped := strip.StripTags(htmlutils.RenderNode(li)) // 1189 - &138 successo qualcosa
			stripped = html.UnescapeString(stripped) // 1189 - è successo qualcosa
			e.Descrizione = append(e.Descrizione, stripped)
		}

	} else { //Se l'item ha solo una descrizione

		/*<li>
			<a href="/wiki/1137" title="1137">1137</a>
			" – Nella "
		    <a href="/wiki/Battaglia_di_Rignano" title="Battaglia di Rignano">battaglia di Rignano</a>
		    " le armate di "
		    <a href="/wiki/Ruggero_II_di_Sicilia" title="Ruggero II di Sicilia">Ruggero II di Sicilia</a>
		    " vengono sconfitte da una coalizione normanno-tedesca guidata da "
		    <a href="/wiki/Rainulfo_di_Alife" title="Rainulfo di Alife">Rainulfo di Alife</a>
		</li>*/

		var str string

		stripped := strip.StripTags(htmlutils.RenderNode(e.htmlNode))
		stripped = html.UnescapeString(stripped) // 1189 - è successo qualcosa
		stripped = strings.Replace(stripped, "-", "@#", -1)
		stripped = strings.Replace(stripped, "–", "@#", -1)
		slice := strings.Split(stripped, " @# ")
		slice = slice[1:]
		for _, element := range slice {
			str += element
			str = strings.Replace(str, "@#", "-", -1)
		}
		e.Descrizione = append(e.Descrizione, str)

	}

	return nil
}
