package types

type CliCommand struct {
	Name        string
	Description string
	Callback    func(arg string) error
}

type PokeResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
