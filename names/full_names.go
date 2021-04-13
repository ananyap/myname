package names

type FullName struct {
	MyName    RealName   `json:"my_name"`
	PairsName []RealName `json:"pairs_name"`
}
