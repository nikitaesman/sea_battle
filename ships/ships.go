package ships

const (
	StatusDead    string = "dead"
	StatusWounded string = "wounded"
)

type ShipDeck struct {
	status string
}
