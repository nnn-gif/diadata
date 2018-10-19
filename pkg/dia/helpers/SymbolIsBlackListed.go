package helpers

//SymbolIsBlackListed return true if the symbol is blacklisted
// Symbols are blacklisted when they are duplicated and there is no
// approach to resolve the conflict
func SymbolIsBlackListed(symbol string) bool {
	switch symbol {
	case "VET":
		return true
	case "ACC":
		return true
	case "ALT":
		return true
	case "ARB":
		return true
	case "AT":
		return true
	case "ATX":
		return true
	case "BBK":
		return true
	case "BET":
		return true
	case "BIT":
		return true
	case "BITS":
		return true
	case "BLZ":
		return true
	case "BTM":
		return true
	case "CAN":
		return true
	case "CAT":
		return true
	case "CBC":
		return true
	case "CEN":
		return true
	case "CMCT":
		return true
	case "CMS":
		return true
	case "CMT":
		return true
	case "CPC":
		return true
	case "CRC":
		return true
	case "DFT":
		return true
	case "EDR":
		return true
	case "ENT":
		return true
	case "ETT":
		return true
	case "EVN":
		return true
	case "EXC":
		return true
	case "FAIR":
		return true
	case "FT":
		return true
	case "GCC":
		return true
	case "GENE":
		return true
	case "GET":
		return true
	case "GOT":
		return true
	case "HC":
		return true
	case "HERO":
		return true
	case "HMC":
		return true
	case "HNC":
		return true
	case "HOT":
		return true
	case "ICN":
		return true
	case "IQ":
		return true
	case "KEY":
		return true
	case "KNC":
		return true
	case "KNT":
		return true
	case "LBTC":
		return true
	case "LNC":
		return true
	case "MAG":
		return true
	case "MORE":
		return true
	case "MTC":
		return true
	case "NET":
		return true
	case "NTK":
		return true
	case "ONG":
		return true
	case "ORS":
		return true
	case "PAI":
		return true
	case "PASS":
		return true
	case "PLC":
		return true
	case "PLY":
		return true
	case "PUT":
		return true
	case "PXC":
		return true
	case "QBT":
		return true
	case "RCN":
		return true
	case "RED":
		return true
	case "SCC":
		return true
	case "SLT":
		return true
	case "SPD":
		return true
	case "TTC":
		return true
	case "WEB":
		return true
	case "XIN":
		return true
	case "XRA":
		return true
	default:
		return false
	}
}
