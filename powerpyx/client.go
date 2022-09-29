package powerpyx

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/go-resty/resty/v2"

	"github.com/petrovskiborislav/cyberpunk2077-todoist-quests/powerpyx/models"
)

const (
	powerpyxCyberpunk2077URL  = "https://www.powerpyx.com/cyberpunk-2077-walkthrough-all-missions/"
	storyMainMissionsCategory = "Story Main Missions"
	sideQuestsCategory        = "Side Quests"
	gigsCategory              = "Gigs"
)

// Client is an interface which defines the methods for the powerpyx client.
type Client interface {
	GetCyberpunk2077Quests() (map[string][]models.Cyberpunk2077Quest, error)
}

// NewClient returns a new instance of the powerpyx client.
func NewClient() Client {
	httpReq := resty.New().R().SetDoNotParseResponse(true)
	return &client{httpRequest: httpReq}
}

type client struct {
	httpRequest *resty.Request
}

// GetCyberpunk2077Quests returns all the Cyberpunk 2077 quests got from
// the powerpyx website and returns them as a map of categories and quests.
func (c *client) GetCyberpunk2077Quests() (map[string][]models.Cyberpunk2077Quest, error) {
	powerpyxCyberpunk2077Page, err := c.httpRequest.Get(powerpyxCyberpunk2077URL)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(powerpyxCyberpunk2077Page.RawBody())
	if err != nil {
		return nil, err
	}

	return getCyberpunk2077QuestsByCategory(doc), nil
}

func getCyberpunk2077QuestsByCategory(doc *goquery.Document) map[string][]models.Cyberpunk2077Quest {
	cyberpunk2077QuestsByCategory := make(map[string][]models.Cyberpunk2077Quest)
	doc.Find("h2").Each(func(i int, h2 *goquery.Selection) {
		heading2 := h2.Text()
		quests := make([]models.Cyberpunk2077Quest, 0)

		switch heading2 {
		case storyMainMissionsCategory, sideQuestsCategory:
			quests = extractCategoryFromHeadings("", h2)
		case gigsCategory:
			headings3 := h2.NextAllFiltered("h3")
			headings3.Each(func(i int, h3 *goquery.Selection) {
				q := extractCategoryFromHeadings(h3.Text(), h3)
				quests = append(quests, q...)
			})
		}
		cyberpunk2077QuestsByCategory[heading2] = quests
	})

	return cyberpunk2077QuestsByCategory
}

func extractCategoryFromHeadings(subcategory string, selection *goquery.Selection) []models.Cyberpunk2077Quest {
	quests := make([]models.Cyberpunk2077Quest, 0)

	selection.Next().Find("li").Each(func(i int, li *goquery.Selection) {
		nameLinkPair := make(map[string]string)
		li.Find("a").Each(func(i int, a *goquery.Selection) {
			nameLinkPair[a.Text()] = a.AttrOr("href", "")
		})
		quests = append(quests, models.Cyberpunk2077Quest{NameLinkPair: nameLinkPair, SubCategory: subcategory})
	})

	return quests
}
