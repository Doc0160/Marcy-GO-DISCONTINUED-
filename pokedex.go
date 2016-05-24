package main
import(
	"encoding/json"
	"slack"
	"strconv"
)
func do_pkdx(ct *CT, s Slack.OMNI) {
	e, _ := explode_cmd(s.Text)
	if len(e) > 1 {
		nb, err := strconv.Atoi(e[1])
		if err == nil {
			Typing(ct.Websocket, s)
			x, err := pkdx(ct, nb)
			Typing(ct.Websocket, s)
			if err == nil {
				Typing(ct.Websocket, s)
				ct.Slack.API_CALL("chat.postMessage", map[string]interface{}{
					"as_user": "true",
					"channel": s.Channel,
					"attachments": []map[string]interface{}{
						map[string]interface{}{
							"title":      "[" + toX(strconv.Itoa(nb), 3, "0", true) + "]" + x.General.NameFr,
							"color":      "#ff0505",
							"image_url":  "http://www.pokemontrash.com/pokedex/images/sugimori/" + toX(strconv.Itoa(nb), 3, "0", true) + ".png",
							"title_link": "http://www.pokemontrash.com/pokedex/" + strconv.Itoa(nb) + "-" + x.General.NameFr + ".html",
						},
						map[string]interface{}{
							"color": "#ff0505",
							"fields": []map[string]interface{}{
								map[string]interface{}{
									"title": "Description",
									"value": x.General.Description,
									"short": false,
								},
								map[string]interface{}{
									"title": "Height",
									"value": x.General.Height,
									"short": true,
								},
								map[string]interface{}{
									"title": "Weight",
									"value": x.General.Weight,
									"short": true,
								},
								map[string]interface{}{
									"title": "Type 1",
									"value": x.Types.Type1,
									"short": true,
								},
								map[string]interface{}{
									"title": func() string {
										if x.Types.Type2 != "" {
											return "Type 2"
										} else {
											return ""
										}
									}(),
									"value": func() string {
										if x.Types.Type2 != "" {
											return x.Types.Type2
										} else {
											return ""
										}
									}(),
									"short": true,
								},
								map[string]interface{}{
									"title": "Generation",
									"value": strconv.Itoa(x.General.Generation),
									"short": true,
								},
								map[string]interface{}{
									"title": "EV yield",
									"value": x.General.EVYield,
									"short": true,
								},
								map[string]interface{}{
									"title": "Genus",
									"value": x.General.Genus,
									"short": true,
								},
								map[string]interface{}{
									"title": "Stage",
									"value": x.General.Stage,
									"short": true,
								},
							},
						},
						map[string]interface{}{
							"title": "Stats",
							"color": "#ff0505",
							"fields": []map[string]interface{}{
								map[string]interface{}{
									"title": "PV",
									"value": x.Stats.Hp,
									"short": true,
								},
								map[string]interface{}{
									"title": "Attaque",
									"value": x.Stats.Attack,
									"short": true,
								},
								map[string]interface{}{
									"title": "Defense",
									"value": x.Stats.Defense,
									"short": true,
								},
								map[string]interface{}{
									"title": "Attaque Spé.",
									"value": x.Stats.Sattack,
									"short": true,
								},
								map[string]interface{}{
									"title": "Defense Spé.",
									"value": x.Stats.Sdefense,
									"short": true,
								},
								map[string]interface{}{
									"title": "Vitesse",
									"value": x.Stats.Speed,
									"short": true,
								},
							},
						},
						map[string]interface{}{
							"title": "Noms",
							"color": "#ff0505",
							"fields": []map[string]interface{}{
								map[string]interface{}{
									"title": "Fr",
									"value": x.General.NameFr,
									"short": true,
								},
								map[string]interface{}{
									"title": "En",
									"value": x.General.NameEn,
									"short": true,
								},
								map[string]interface{}{
									"title": "De",
									"value": x.General.NameDe,
									"short": true,
								},
								map[string]interface{}{
									"title": "Es",
									"value": x.General.NameEs,
									"short": true,
								},
								map[string]interface{}{
									"title": "It",
									"value": x.General.NameIt,
									"short": true,
								},
								map[string]interface{}{
									"title": "Jp",
									"value": x.General.NameJp,
									"short": true,
								},
							},
						},
					},
				})
			} else {
				if nb > 400 || nb < 1 {
					Message(ct.Websocket, s, "Le numéro que vous essayez de joindre n'existe pas ...")
				} else {
					Message(ct.Websocket, s, "Y'a une couille dans le paté !\n"+err.Error())
				}
			}
		} else {
			Message(ct.Websocket, s, "Veuillez présciser un numéro correct de pokemon à consulter.")
		}
	} else {
		Message(ct.Websocket, s, "Veuillez présciser le numéro du pokemon à consulter.")
	}
}
func pkdx(ct *CT, num int) (Pokedex, error) {
	var a Pokedex
	r, err := commonHttpRequest(ct, "http://www.pokemontrash.com/api/fr/pokemon/"+strconv.Itoa(num))
	if err != nil {
		return a, err
	} else {
		err := json.NewDecoder(*r).Decode(&a)
		defer (*r).Close()
		if err != nil {
			return a, err
		} else {
			return a, nil
		}
	}
}
type Pokedex struct {
	General struct {
		NameFr      string  `json:"name-fr"`
		NameEn      string  `json:"name-en"`
		NameDe      string  `json:"name-de"`
		NameEs      string  `json:"name-es"`
		NameIt      string  `json:"name-it"`
		NameJp      string  `json:"name-jp"`
		Genus       string  `json:"genus"`
		Pokedex     string  `json:"pokedex"`
		Generation  int     `json:"generation"`
		Weight      float64 `json:"weight"`
		Height      float64 `json:"height"`
		EVYield     string  `json:"EV-yield"`
		Stage       int     `json:"stage"`
		Hatch       int     `json:"hatch"`
		Exp         string  `json:"exp"`
		Description string  `json:"description"`
	} `json:"general"`
	Stats struct {
		Hp       string `json:"hp"`
		Attack   string `json:"attack"`
		Defense  string `json:"defense"`
		Sattack  string `json:"sattack"`
		Sdefense string `json:"sdefense"`
		Speed    string `json:"speed"`
	} `json:"stats"`
	Evolution struct {
		Stage1 struct {
			Evolution1 struct {
				ID    string `json:"id"`
				Name  string `json:"name"`
				Type1 string `json:"type-1"`
				Type2 string `json:"type-2"`
			} `json:"evolution-1"`
		} `json:"stage-1"`
		Stage2 struct {
			Evolution1 struct {
				ID    string `json:"id"`
				Name  string `json:"name"`
				Type1 string `json:"type-1"`
				Type2 string `json:"type-2"`
				How   string `json:"how"`
			} `json:"evolution-1"`
		} `json:"stage-2"`
		Stage3 struct {
			Evolution1 struct {
				ID    string `json:"id"`
				Name  string `json:"name"`
				Type1 string `json:"type-1"`
				Type2 string `json:"type-2"`
				How   string `json:"how"`
			} `json:"evolution-1"`
		} `json:"stage-3"`
	} `json:"evolution"`
	Types struct {
		Type1 string `json:"type-1"`
		Type2 string `json:"type-2"`
	} `json:"types"`
	Ability struct {
		Ability1 string `json:"ability-1"`
		Ability2 string `json:"ability-2"`
	} `json:"ability"`
	URI struct {
		PokemonID string `json:"pokemon-id"`
		Ability1  string `json:"ability-1"`
		Ability2  string `json:"ability-2"`
		Type1     string `json:"type-1"`
		Type2     string `json:"type-2"`
	} `json:"uri"`
}