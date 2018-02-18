package omnigo

type OmniInfo struct {
	OmnicoreVersionInt int
	OmnicoreVersion    string
	MasterCoreVersion  string
	BitcoinCoreVersion string
	Block              int
	Blocktime          int
	BlockTransactions  int
	TotalTransactions  int
	Alerts             []Alert
}

type Alert struct {
	TypeInt   int
	TokenType string
	Expiry    string
	Message   string
}

type Balance struct {
	PropertyID int
	Balance    string
	Reserved   string
}
