package overpass

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseOverpassJson(t *testing.T) {
	t.Run("Nepal", func(t *testing.T) {
		overpassJson := `
		{
			"version": 0.6,
			"generator": "Overpass API 0.7.57.1 74a55df1",
			"osm3s": {
				"timestamp_osm_base": "2021-11-21T08:51:55Z",
				"timestamp_areas_base": "2021-11-21T08:27:30Z",
				"copyright": "The data included in this document is from www.openstreetmap.org. The data is made available under ODbL."
			},
			"elements": [{
				"type": "area",
				"id": 3604583291,
				"tags": {
					"ISO3166-2": "NP-4",
					"admin_level": "4",
					"boundary": "historic",
					"int_name": "Eastern Development Region",
					"is_in:continent": "Asia",
					"is_in:country": "Nepal",
					"is_in:country_code": "NP",
					"name": "पुर्वाञ्चल विकास क्षेत्र",
					"name:ar": "المنطقة التنموية الشرقية",
					"name:de": "Entwicklungsregion Ost",
					"name:en": "Eastern Development Region",
					"name:es": "Región de desarrollo Este",
					"name:fr": "Région de développement Est",
					"name:ja": "東部開発区域",
					"name:ne": "पुर्वाञ्चल विकास क्षेत्र",
					"name:pl": "Eastern Development Region",
					"name:ru": "Восточный регион",
					"name:ur": "مشرقی ترقیاتی علاقہ",
					"name:zh": "东部经济发展区",
					"note": "outdated",
					"type": "boundary",
					"wikidata": "Q28576",
					"wikipedia": "en:Eastern Development Region, Nepal"
				}
			}]
		}`

		expectedLocalizedNames := map[string]string{
			"ar": "المنطقة التنموية الشرقية",
			"de": "Entwicklungsregion Ost",
			"en": "Eastern Development Region",
			"es": "Región de desarrollo Este",
			"fr": "Région de développement Est",
			"ja": "東部開発区域",
			"ne": "पुर्वाञ्चल विकास क्षेत्र",
			"pl": "Eastern Development Region",
			"ru": "Восточный регион",
			"ur": "مشرقی ترقیاتی علاقہ",
			"zh": "东部经济发展区",
		}

		var r OverpassResponse

		if err := json.Unmarshal([]byte(overpassJson), &r); err != nil {
			t.Fatal(err)
		}

		assert.NotNil(t, r)
		assert.Len(t, r.Elements, 1)

		overpassArea := r.Elements[0]
		assert.Equal(t, "area", overpassArea.Type)
		assert.Equal(t, "4", overpassArea.AdministrativeLevel())
		assert.Equal(t, "पुर्वाञ्चल विकास क्षेत्र", overpassArea.Name())
		assert.Equal(t, "Eastern Development Region", overpassArea.InternationalName())
		assert.Equal(t, expectedLocalizedNames, overpassArea.LocalizedNames())
	})
}

func TestFindStateAndCity(t *testing.T) {
	t.Run("Barcelona", func(t *testing.T) {
		overpassJson := `
		{
			"version": 0.6,
			"generator": "Overpass API 0.7.57.1 74a55df1",
			"osm3s": {
				"timestamp_osm_base": "2021-12-01T13:10:56Z",
				"timestamp_areas_base": "2021-12-01T12:17:03Z",
				"copyright": "The data included in this document is from www.openstreetmap.org. The data is made available under ODbL."
			},
			"elements": [{
				"type": "info",
				"id": 1,
				"tags": {
					"approximate": "1"
				}
			}, {
				"type": "relation",
				"id": 347950,
				"tags": {
					"admin_level": "8",
					"alt_name:gl": "Barna",
					"boundary": "administrative",
					"capital": "4",
					"idee:name": "Barcelona",
					"ine:municipio": "08019",
					"name": "Barcelona",
					"name:ca": "Barcelona",
					"name:cs": "Barcelona",
					"name:el": "Βαρκελώνη",
					"name:eo": "Barcelono",
					"name:fr": "Barcelone",
					"name:gl": "Barcelona",
					"name:he": "ברצלונה",
					"name:ko": "바르셀로나",
					"name:lt": "Barselona",
					"name:oc": "Barcelona",
					"name:pl": "Barcelona",
					"name:ru": "Барселона",
					"name:uk": "Барселона",
					"name:zh": "巴塞罗那",
					"name:zh-Hans": "巴塞罗那",
					"name:zh-Hant": "巴塞隆納",
					"place": "city",
					"population": "1619337",
					"population:date": "2009",
					"postal_code": "08001",
					"source": "BDLL25, EGRN, Instituto Geográfico Nacional",
					"source:name:oc": "Lo Congrès",
					"type": "boundary",
					"wikidata": "Q1492",
					"wikipedia": "ca:Barcelona"
				}
			}, {
				"type": "relation",
				"id": 349035,
				"tags": {
					"ISO3166-2": "ES-B",
					"admin_level": "6",
					"boundary": "administrative",
					"ine:provincia": "08",
					"name": "Barcelona",
					"name:ar": "برشلونة",
					"name:be": "Барселона",
					"name:ca": "Barcelona",
					"name:el": "Βαρκελώνη",
					"name:es": "Barcelona",
					"name:eu": "Bartzelona",
					"name:fr": "Barcelone",
					"name:gl": "Barcelona",
					"name:lt": "Barselona",
					"name:oc": "Barcelona",
					"name:ru": "Барселона",
					"note": "Comarcas not fully within the province: Berguedà, Osona, Selva",
					"official_name:be": "Правінцыя Барселона",
					"ref:nuts": "ES511",
					"ref:nuts:3": "ES511",
					"short_name": "BCN",
					"source": "BDLL25, EGRN, Instituto Geográfico Nacional",
					"source:name:oc": "ieo-bdtopoc",
					"type": "boundary",
					"website": "http://www.diba.es/",
					"wikidata": "Q81949",
					"wikipedia": "ca:Província de Barcelona"
				}
			}, {
				"type": "relation",
				"id": 349053,
				"tags": {
					"ISO3166-2": "ES-CT",
					"admin_level": "4",
					"alt_name:el": "Καταλονία",
					"alt_name:eo": "Katalunujo",
					"alt_name:gl": "Catalunha",
					"border_type": "region",
					"boundary": "administrative",
					"full_name:cs": "Autonomní společenství Katalánsko",
					"ine:ccaa": "09",
					"name": "Catalunya",
					"name:ar": "كتالونيا",
					"name:be": "Каталонія",
					"name:br": "Katalonia",
					"name:ca": "Catalunya",
					"name:cs": "Katalánsko",
					"name:de": "Katalonien",
					"name:el": "Καταλωνία",
					"name:en": "Catalonia",
					"name:eo": "Katalunio",
					"name:es": "Cataluña",
					"name:eu": "Katalunia",
					"name:fi": "Katalonia",
					"name:fr": "Catalogne",
					"name:fy": "Kataloanje",
					"name:gl": "Cataluña",
					"name:hr": "Katalonija",
					"name:hu": "Katalónia",
					"name:io": "Katalunia",
					"name:is": "Katalónía",
					"name:ja": "カタルーニャ州",
					"name:lt": "Katalonija",
					"name:mk": "Каталонија",
					"name:nl": "Catalonië",
					"name:oc": "Catalonha",
					"name:pl": "Katalonia",
					"name:pt": "Catalunha",
					"name:ru": "Каталония",
					"name:sk": "Katalánsko",
					"name:sl": "Katalonija",
					"name:sv": "Katalonien",
					"name:tok": "ma Katala",
					"name:tzl": "Ceatalognh",
					"name:vo": "Katalonän",
					"name:zh-Hans": "加泰罗尼亚",
					"name:zh-Hant": "加泰羅尼亞",
					"note": "Comarcas covered by multiple provinces: Berguedà, Cerdanya, Osona, Selva",
					"official_name": "Catalunya",
					"official_name:ca": "Catalunya",
					"official_name:es": "Cataluña",
					"ref:nuts": "ES51",
					"ref:nuts:2": "ES51",
					"short_name": "CAT",
					"source": "BDLL25, EGRN, Instituto Geográfico Nacional",
					"source:name:br": "ofis publik ar brezhoneg",
					"source:name:oc": "ieo-bdtopoc",
					"type": "boundary",
					"wikidata": "Q5705",
					"wikipedia": "ca:Catalunya"
				}
			}, {
				"type": "relation",
				"id": 2417889,
				"tags": {
					"addr:country_code": "ES",
					"admin_level": "7",
					"boundary": "administrative",
					"idescat:comarca": "13",
					"name": "Barcelonès",
					"name:an": "Barcelonés",
					"name:ar": "بارسلونس",
					"name:ca": "Barcelonès",
					"name:es": "Barcelonés",
					"name:fr": "Barcelonais",
					"name:gl": "Barcelonés",
					"name:ja": "バルサルネス",
					"name:kk": "Барселонес",
					"name:oc": "Barcelonés",
					"name:ru": "Барселонес",
					"name:uk": "Барсалунес",
					"population": "2246280",
					"population:date": "2012",
					"short_name": "BCN",
					"source:population": "IDESCAT",
					"type": "boundary",
					"website": "http://www.ccbcnes.org/",
					"wikidata": "Q15348",
					"wikipedia": "ca:Barcelonès"
				}
			}]
		}`

		var r OverpassResponse

		if err := json.Unmarshal([]byte(overpassJson), &r); err != nil {
			t.Fatal(err)
		}

		assert.NotNil(t, r)

		assert.Equal(t, "Barcelona", r.Elements.FindCity())
		assert.Equal(t, "Catalunya", r.Elements.FindState())
	})
}

func TestFindState(t *testing.T) {
	t.Run("Khumjung, Nepal", func(t *testing.T) {
		state := FindState("39e9ac0d2c4c")

		// For Nepal the correct boundary is admin_level 3
		assert.Equal(t, "कोशी प्रदेश", state)
	})

	t.Run("Paris, France", func(t *testing.T) {
		state := FindState("47e66fe289fc")

		// For France the correct boundary is admin_level 4 (and NOT 3)
		assert.Equal(t, "Île-de-France", state)
	})

	t.Run("Thingvellir, Iceland", func(t *testing.T) {
		state := FindState("48d680def0a4")

		// For Iceland the correct boundary is admin_level 5
		assert.Equal(t, "Suðurland", state)
	})

	t.Run("Rambla de Mar, Barcelona, Spain", func(t *testing.T) {
		state := FindState("12a4a253f84c")

		// This location does not have data in OSM, so it is used to test the location approximation
		// assert.Equal(t, "Barcelona", city)
		assert.Equal(t, "Catalunya", state)
	})

	t.Run("Luxor, Egypt", func(t *testing.T) {
		state := FindState("14491621627")

		// For Egypt the correct boundary is admin_level 4
		assert.Equal(t, "الأقصر", state)
	})

	t.Run("Sofia, Bulgaria", func(t *testing.T) {
		state := FindState("40aa835dc0d")

		// For Bulgaria the correct boundary is admin_level 6
		assert.Equal(t, "София-град", state)
	})
}
