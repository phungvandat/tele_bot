package main

type CoinInfo struct {
	ID               uint16
	BinanceUSDT      string
	Slug             string
	Name             string
	LastPrice        float64
	LastAlertSuccess bool
}

var mapSymbolCoins = map[string][]CoinInfo{
	"QHC": {
		{
			ID:          7243,
			Slug:        "qchi-chain",
			Name:        "QChi Chain",
			BinanceUSDT: "QHCUSDT",
		},
	},
	"CPT": {
		{
			ID:          2766,
			Slug:        "cryptaur",
			Name:        "Cryptaur",
			BinanceUSDT: "CPTUSDT",
		},
	},
	"DEV": {
		{
			ID:          5990,
			Slug:        "dev-protocol",
			Name:        "Dev Protocol",
			BinanceUSDT: "DEVUSDT",
		},
	},
	"NVDA": {
		{
			ID:          7913,
			Slug:        "nvidia-tokenized-stock-ftx",
			Name:        "NVIDIA tokenized stock FTX",
			BinanceUSDT: "NVDAUSDT",
		},
	},
	"XSUTER": {
		{
			ID:          9838,
			Slug:        "xsuter",
			Name:        "xSuter",
			BinanceUSDT: "XSUTERUSDT",
		},
	},
	"AME": {
		{
			ID:          8162,
			Slug:        "amepay",
			Name:        "AMEPAY",
			BinanceUSDT: "AMEUSDT",
		},
	},
	"SPIRIT": {
		{
			ID:          10239,
			Slug:        "spiritswap",
			Name:        "SpiritSwap",
			BinanceUSDT: "SPIRITUSDT",
		},
	},
	"BNU": {
		{
			ID:          10121,
			Slug:        "bytenext",
			Name:        "ByteNext",
			BinanceUSDT: "BNUUSDT",
		},
	},
	"XVS": {
		{
			ID:          7288,
			Slug:        "venus",
			Name:        "Venus",
			BinanceUSDT: "XVSUSDT",
		},
	},
	"MTRG": {
		{
			ID:          5919,
			Slug:        "meter-governance",
			Name:        "Meter Governance",
			BinanceUSDT: "MTRGUSDT",
		},
	},
	"RYO": {
		{
			ID:          2976,
			Slug:        "ryo-currency",
			Name:        "Ryo Currency",
			BinanceUSDT: "RYOUSDT",
		},
	},
	"TMED": {
		{
			ID:          6237,
			Slug:        "mdsquare",
			Name:        "MDsquare",
			BinanceUSDT: "TMEDUSDT",
		},
	},
	"UCO": {
		{
			ID:          6887,
			Slug:        "uniris",
			Name:        "Uniris",
			BinanceUSDT: "UCOUSDT",
		},
	},
	"CATT": {
		{
			ID:          4045,
			Slug:        "catex-token",
			Name:        "Catex Token",
			BinanceUSDT: "CATTUSDT",
		},
	},
	"TRR": {
		{
			ID:          9666,
			Slug:        "terran-coin",
			Name:        "Terran Coin",
			BinanceUSDT: "TRRUSDT",
		},
	},
	"MIN": {
		{
			ID:          3296,
			Slug:        "mindol",
			Name:        "MINDOL",
			BinanceUSDT: "MINUSDT",
		},
	},
	"POST": {
		{
			ID:          1218,
			Slug:        "postcoin",
			Name:        "PostCoin",
			BinanceUSDT: "POSTUSDT",
		},
	},
	"PUX": {
		{
			ID:          6022,
			Slug:        "polypux",
			Name:        "PolypuX",
			BinanceUSDT: "PUXUSDT",
		},
	},
	"REDI": {
		{
			ID:          8388,
			Slug:        "redi",
			Name:        "REDi",
			BinanceUSDT: "REDIUSDT",
		},
	},
	"OLYMPUS": {
		{
			ID:          10764,
			Slug:        "olympus-token",
			Name:        "OLYMPUS",
			BinanceUSDT: "OLYMPUSUSDT",
		},
	},
	"BREE": {
		{
			ID:          6502,
			Slug:        "cbdao",
			Name:        "CBDAO",
			BinanceUSDT: "BREEUSDT",
		},
	},
	"FOL": {
		{
			ID:          7687,
			Slug:        "folder-protocol",
			Name:        "Folder Protocol",
			BinanceUSDT: "FOLUSDT",
		},
	},
	"DIRKDIGGLER": {
		{
			ID:          10749,
			Slug:        "the-boogie-nights",
			Name:        "The Boogie Nights",
			BinanceUSDT: "DIRKDIGGLERUSDT",
		},
	},
	"LN": {
		{
			ID:          4512,
			Slug:        "link",
			Name:        "LINK",
			BinanceUSDT: "LNUSDT",
		},
	},
	"WGRT": {
		{
			ID:          5630,
			Slug:        "waykichain-governance-coin",
			Name:        "WaykiChain Governance Coin",
			BinanceUSDT: "WGRTUSDT",
		},
	},
	"ADAUP": {
		{
			ID:          7013,
			Slug:        "adaup",
			Name:        "ADAUP",
			BinanceUSDT: "ADAUPUSDT",
		},
	},
	"LEVL": {
		{
			ID:          4173,
			Slug:        "levolution",
			Name:        "Levolution",
			BinanceUSDT: "LEVLUSDT",
		},
	},
	"GEN": {
		{
			ID:   2726,
			Slug: "daostack",
			Name: "DAOstack",
		},
		{
			ID:   9495,
			Slug: "evolution",
			Name: "Evolution",
		},
	},
	"NTK": {
		{
			ID:   2536,
			Slug: "neurotoken",
			Name: "Neurotoken",
		},
		{
			ID:   3149,
			Slug: "netkoin",
			Name: "Netkoin",
		},
	},
	"PEEPS": {
		{
			ID:          10952,
			Slug:        "the-peoples-coin",
			Name:        "The People's Coin",
			BinanceUSDT: "PEEPSUSDT",
		},
	},
	"NDR": {
		{
			ID:          7993,
			Slug:        "node-runners",
			Name:        "Node Runners",
			BinanceUSDT: "NDRUSDT",
		},
	},
	"ANK": {
		{
			ID:   5052,
			Slug: "apple-network",
			Name: "Apple Network",
		},
		{
			ID:   6632,
			Slug: "alphalink",
			Name: "AlphaLink",
		},
	},
	"SSGT": {
		{
			ID:          10331,
			Slug:        "safeswap-governance-token",
			Name:        "Safeswap Governance Token",
			BinanceUSDT: "SSGTUSDT",
		},
	},
	"XEND": {
		{
			ID:          8519,
			Slug:        "xend-finance",
			Name:        "Xend Finance",
			BinanceUSDT: "XENDUSDT",
		},
	},
	"ICE": {
		{
			ID:          9073,
			Slug:        "popsicle-finance",
			Name:        "Popsicle Finance",
			BinanceUSDT: "ICEUSDT",
		},
	},
	"SNL": {
		{
			ID:          3977,
			Slug:        "sport-and-leisure",
			Name:        "Sport and Leisure",
			BinanceUSDT: "SNLUSDT",
		},
	},
	"CUSD": {
		{
			ID:          7236,
			Slug:        "celo-dollar",
			Name:        "Celo Dollar",
			BinanceUSDT: "CUSDUSDT",
		},
	},
	"WILLIE": {
		{
			ID:          10009,
			Slug:        "williecoin",
			Name:        "Williecoin",
			BinanceUSDT: "WILLIEUSDT",
		},
	},
	"GOD": {
		{
			ID:   2370,
			Slug: "bitcoin-god",
			Name: "Bitcoin God",
		},
		{
			ID:   10001,
			Slug: "game-of-defi",
			Name: "Game Of DeFi",
		},
	},
	"RBOYS": {
		{
			ID:          10930,
			Slug:        "rocket-boys",
			Name:        "Rocket Boys",
			BinanceUSDT: "RBOYSUSDT",
		},
	},
	"EURXB": {
		{
			ID:          9136,
			Slug:        "eurxb",
			Name:        "EURxb",
			BinanceUSDT: "EURXBUSDT",
		},
	},
	"XLMG": {
		{
			ID:          3926,
			Slug:        "stellar-gold",
			Name:        "Stellar Gold",
			BinanceUSDT: "XLMGUSDT",
		},
	},
	"LIEN": {
		{
			ID:          6705,
			Slug:        "lien",
			Name:        "Lien",
			BinanceUSDT: "LIENUSDT",
		},
	},
	"WING": {
		{
			ID:   7048,
			Slug: "wing",
			Name: "Wing",
		},
		{
			ID:   6584,
			Slug: "wingshop",
			Name: "WingShop",
		},
	},
	"ZEFI": {
		{
			ID:          8759,
			Slug:        "zcore-finance",
			Name:        "ZCore Finance",
			BinanceUSDT: "ZEFIUSDT",
		},
	},
	"ARA": {
		{
			ID:          9490,
			Slug:        "ara-blocks",
			Name:        "Ara Blocks",
			BinanceUSDT: "ARAUSDT",
		},
	},
	"HBC": {
		{
			ID:   6965,
			Slug: "hbtc-token",
			Name: "HBTC Captain Token",
		},
		{
			ID:   7073,
			Slug: "hybrid-bank-cash",
			Name: "Hybrid Bank Cash",
		},
		{
			ID:   7073,
			Slug: "hybrid-bank-cash",
			Name: "Hybrid Bank Cash",
		},
	},
	"WPP": {
		{
			ID:          3885,
			Slug:        "wpp-token",
			Name:        "WPP TOKEN",
			BinanceUSDT: "WPPUSDT",
		},
	},
	"FREE": {
		{
			ID:   3388,
			Slug: "free-coin",
			Name: "FREE Coin",
		},
		{
			ID:   9538,
			Slug: "anti-lockdown",
			Name: "Anti-Lockdown",
		},
	},
	"BABYDOGECASH": {
		{
			ID:          10882,
			Slug:        "baby-doge-cash",
			Name:        "Baby Doge Cash",
			BinanceUSDT: "BABYDOGECASHUSDT",
		},
	},
	"RIO": {
		{
			ID:          4166,
			Slug:        "realio-network",
			Name:        "Realio Network",
			BinanceUSDT: "RIOUSDT",
		},
	},
	"SOVI": {
		{
			ID:          8741,
			Slug:        "sovi-finance",
			Name:        "Sovi Finance",
			BinanceUSDT: "SOVIUSDT",
		},
	},
	"BAMBOO": {
		{
			ID:          8389,
			Slug:        "bamboo-defi",
			Name:        "BambooDeFi",
			BinanceUSDT: "BAMBOOUSDT",
		},
	},
	"YAS": {
		{
			ID:          6246,
			Slug:        "yas",
			Name:        "YAS",
			BinanceUSDT: "YASUSDT",
		},
	},
	"FXF": {
		{
			ID:          8416,
			Slug:        "finxflo",
			Name:        "Finxflo",
			BinanceUSDT: "FXFUSDT",
		},
	},
	"VIKKY": {
		{
			ID:          2965,
			Slug:        "vikkytoken",
			Name:        "VikkyToken",
			BinanceUSDT: "VIKKYUSDT",
		},
	},
	"DONI": {
		{
			ID:          10940,
			Slug:        "doni-coin",
			Name:        "Doni Coin",
			BinanceUSDT: "DONIUSDT",
		},
	},
	"CHY": {
		{
			ID:          10885,
			Slug:        "concern-poverty-chain",
			Name:        "Concern Poverty Chain",
			BinanceUSDT: "CHYUSDT",
		},
	},
	"HANDY": {
		{
			ID:          7755,
			Slug:        "handy",
			Name:        "Handy",
			BinanceUSDT: "HANDYUSDT",
		},
	},
	"NAMI": {
		{
			ID:          7534,
			Slug:        "tsunami",
			Name:        "Tsunami finance",
			BinanceUSDT: "NAMIUSDT",
		},
	},
	"LMCH": {
		{
			ID:          5418,
			Slug:        "latamcash",
			Name:        "Latamcash",
			BinanceUSDT: "LMCHUSDT",
		},
	},
	"IBVOL": {
		{
			ID:          5534,
			Slug:        "inverse-bitcoin-volatility-token",
			Name:        "Inverse Bitcoin Volatility Token",
			BinanceUSDT: "IBVOLUSDT",
		},
	},
	"THC": {
		{
			ID:          416,
			Slug:        "hempcoin",
			Name:        "HempCoin",
			BinanceUSDT: "THCUSDT",
		},
	},
	"VOYRME": {
		{
			ID:          10690,
			Slug:        "voyr",
			Name:        "VOYR",
			BinanceUSDT: "VOYRMEUSDT",
		},
	},
	"VEC2": {
		{
			ID:          1052,
			Slug:        "vector",
			Name:        "VectorAI",
			BinanceUSDT: "VEC2USDT",
		},
	},
	"EOST": {
		{
			ID:          4124,
			Slug:        "eos-trust",
			Name:        "EOS TRUST",
			BinanceUSDT: "EOSTUSDT",
		},
	},
	"EVM": {
		{
			ID:          10682,
			Slug:        "evermars",
			Name:        "EverMars",
			BinanceUSDT: "EVMUSDT",
		},
	},
	"COSBY": {
		{
			ID:          10760,
			Slug:        "the-cosby-token",
			Name:        "The Cosby Token",
			BinanceUSDT: "COSBYUSDT",
		},
	},
	"VERS": {
		{
			ID:          4750,
			Slug:        "versess-coin",
			Name:        "Versess Coin",
			BinanceUSDT: "VERSUSDT",
		},
	},
	"WET": {
		{
			ID:          3585,
			Slug:        "weshow-token",
			Name:        "WeShow Token",
			BinanceUSDT: "WETUSDT",
		},
	},
	"YFIAG": {
		{
			ID:          7862,
			Slug:        "yearnagnostic-finance",
			Name:        "YearnAgnostic Finance",
			BinanceUSDT: "YFIAGUSDT",
		},
	},
	"BLX": {
		{
			ID:   5881,
			Slug: "balloon-x",
			Name: "Balloon-X",
		},
		{
			ID:   10939,
			Slug: "bullex",
			Name: "BulleX",
		},
	},
	"CLBR": {
		{
			ID:          7539,
			Slug:        "colibri",
			Name:        "Colibri Protocol",
			BinanceUSDT: "CLBRUSDT",
		},
	},
	"BANK": {
		{
			ID:   8612,
			Slug: "float-protocol",
			Name: "Float Protocol",
		},
		{
			ID:   8112,
			Slug: "the-bank-coin",
			Name: "Bankcoin",
		},
		{
			ID:   9607,
			Slug: "bankless-dao",
			Name: "Bankless DAO",
		},
	},
	"ARTEX": {
		{
			ID:          11029,
			Slug:        "artex",
			Name:        "Artex",
			BinanceUSDT: "ARTEXUSDT",
		},
	},
	"ECHO": {
		{
			ID:          9979,
			Slug:        "echelon-dao",
			Name:        "Echelon DAO",
			BinanceUSDT: "ECHOUSDT",
		},
	},
	"BPOP": {
		{
			ID:          6388,
			Slug:        "bpop",
			Name:        "BPOP",
			BinanceUSDT: "BPOPUSDT",
		},
	},
	"PWR": {
		{
			ID:          1279,
			Slug:        "powercoin",
			Name:        "PWR Coin",
			BinanceUSDT: "PWRUSDT",
		},
	},
	"SFI": {
		{
			ID:          7617,
			Slug:        "saffron-finance",
			Name:        "saffron.finance",
			BinanceUSDT: "SFIUSDT",
		},
	},
	"GAME": {
		{
			ID:          576,
			Slug:        "gamecredits",
			Name:        "GameCredits",
			BinanceUSDT: "GAMEUSDT",
		},
	},
	"WXT": {
		{
			ID:          4090,
			Slug:        "wirex-token",
			Name:        "Wirex Token",
			BinanceUSDT: "WXTUSDT",
		},
	},
	"FLUNAR": {
		{
			ID:          9623,
			Slug:        "fairlunar",
			Name:        "FairLunar",
			BinanceUSDT: "FLUNARUSDT",
		},
	},
	"LAS": {
		{
			ID:          8960,
			Slug:        "lnasolution-coin",
			Name:        "LNAsolution Coin",
			BinanceUSDT: "LASUSDT",
		},
	},
	"FEVR": {
		{
			ID:          10803,
			Slug:        "realfevr",
			Name:        "RealFevr",
			BinanceUSDT: "FEVRUSDT",
		},
	},
	"SPORTS": {
		{
			ID:          6564,
			Slug:        "zensports",
			Name:        "ZenSports",
			BinanceUSDT: "SPORTSUSDT",
		},
	},
	"FKX": {
		{
			ID:          3241,
			Slug:        "fortknoxster",
			Name:        "FortKnoxster",
			BinanceUSDT: "FKXUSDT",
		},
	},
	"POLS": {
		{
			ID:          7208,
			Slug:        "polkastarter",
			Name:        "Polkastarter",
			BinanceUSDT: "POLSUSDT",
		},
	},
	"LOCG": {
		{
			ID:          9526,
			Slug:        "locgame",
			Name:        "LOCGame",
			BinanceUSDT: "LOCGUSDT",
		},
	},
	"DASH": {
		{
			ID:          131,
			Slug:        "dash",
			Name:        "Dash",
			BinanceUSDT: "DASHUSDT",
		},
	},
	"RIF": {
		{
			ID:          3701,
			Slug:        "rsk-infrastructure-framework",
			Name:        "RSK Infrastructure Framework",
			BinanceUSDT: "RIFUSDT",
		},
	},
	"KCASH": {
		{
			ID:          2379,
			Slug:        "kcash",
			Name:        "Kcash",
			BinanceUSDT: "KCASHUSDT",
		},
	},
	"OURO": {
		{
			ID:          4970,
			Slug:        "ouroboros",
			Name:        "Ouroboros",
			BinanceUSDT: "OUROUSDT",
		},
	},
	"FEAR": {
		{
			ID:          9866,
			Slug:        "fear-nfts",
			Name:        "Fear NFTs",
			BinanceUSDT: "FEARUSDT",
		},
	},
	"NFI": {
		{
			ID:          9275,
			Slug:        "norse-finance",
			Name:        "Norse Finance",
			BinanceUSDT: "NFIUSDT",
		},
	},
	"FO": {
		{
			ID:          4058,
			Slug:        "fibos",
			Name:        "FIBOS",
			BinanceUSDT: "FOUSDT",
		},
	},
	"MNTIS": {
		{
			ID:          8196,
			Slug:        "mantis",
			Name:        "Mantis",
			BinanceUSDT: "MNTISUSDT",
		},
	},
	"NIX": {
		{
			ID:          2991,
			Slug:        "nix",
			Name:        "NIX",
			BinanceUSDT: "NIXUSDT",
		},
	},
	"MEETONE": {
		{
			ID:          3136,
			Slug:        "meetone",
			Name:        "MEET.ONE",
			BinanceUSDT: "MEETONEUSDT",
		},
	},
	"YEED": {
		{
			ID:          3474,
			Slug:        "yeed",
			Name:        "YGGDRASH",
			BinanceUSDT: "YEEDUSDT",
		},
	},
	"BVOL": {
		{
			ID:          5542,
			Slug:        "1x-long-bitcoin-implied-volatility-token",
			Name:        "1x Long Bitcoin Implied Volatility Token",
			BinanceUSDT: "BVOLUSDT",
		},
	},
	"BDPI": {
		{
			ID:          9153,
			Slug:        "interest-bearing-dpi",
			Name:        "Interest Bearing Defi Pulse Index",
			BinanceUSDT: "BDPIUSDT",
		},
	},
	"BSBT": {
		{
			ID:          8623,
			Slug:        "bsb-token",
			Name:        "BSB Token",
			BinanceUSDT: "BSBTUSDT",
		},
	},
	"SAFERMOON": {
		{
			ID:          9761,
			Slug:        "safermoon",
			Name:        "SaferMoon",
			BinanceUSDT: "SAFERMOONUSDT",
		},
	},
	"CMIT": {
		{
			ID:          3328,
			Slug:        "cmitcoin",
			Name:        "CMITCOIN",
			BinanceUSDT: "CMITUSDT",
		},
	},
	"EXEN": {
		{
			ID:          8565,
			Slug:        "exen-coin",
			Name:        "Exen Coin",
			BinanceUSDT: "EXENUSDT",
		},
	},
	"IDNA": {
		{
			ID:          5836,
			Slug:        "idena",
			Name:        "Idena",
			BinanceUSDT: "IDNAUSDT",
		},
	},
	"DMG": {
		{
			ID:          5741,
			Slug:        "dmm-governance",
			Name:        "DMM: Governance",
			BinanceUSDT: "DMGUSDT",
		},
	},
	"MEMES": {
		{
			ID:          8991,
			Slug:        "memes-token",
			Name:        "Memes Token",
			BinanceUSDT: "MEMESUSDT",
		},
	},
	"RBC": {
		{
			ID:          7219,
			Slug:        "rubic",
			Name:        "Rubic",
			BinanceUSDT: "RBCUSDT",
		},
	},
	"MOZ": {
		{
			ID:          9802,
			Slug:        "mozik",
			Name:        "Mozik",
			BinanceUSDT: "MOZUSDT",
		},
	},
	"LITTLEDOGE": {
		{
			ID:          10956,
			Slug:        "littledoge",
			Name:        "LittleDoge",
			BinanceUSDT: "LITTLEDOGEUSDT",
		},
	},
	"VVT": {
		{
			ID:          7907,
			Slug:        "versoview",
			Name:        "VersoView",
			BinanceUSDT: "VVTUSDT",
		},
	},
	"ETHM": {
		{
			ID:          3589,
			Slug:        "ethereum-meta",
			Name:        "Ethereum Meta",
			BinanceUSDT: "ETHMUSDT",
		},
	},
	"XPR": {
		{
			ID:          5350,
			Slug:        "proton",
			Name:        "Proton",
			BinanceUSDT: "XPRUSDT",
		},
	},
	"NBOT": {
		{
			ID:          4047,
			Slug:        "naka-bodhi-token",
			Name:        "Naka Bodhi Token",
			BinanceUSDT: "NBOTUSDT",
		},
	},
	"PKB": {
		{
			ID:          934,
			Slug:        "parkbyte",
			Name:        "ParkByte",
			BinanceUSDT: "PKBUSDT",
		},
	},
	"PIPL": {
		{
			ID:          2056,
			Slug:        "piplcoin",
			Name:        "PiplCoin",
			BinanceUSDT: "PIPLUSDT",
		},
	},
	"UBT": {
		{
			ID:          2758,
			Slug:        "unibright",
			Name:        "Unibright",
			BinanceUSDT: "UBTUSDT",
		},
	},
	"SLD": {
		{
			ID:   9662,
			Slug: "shieldex",
			Name: "ShieldEX",
		},
		{
			ID:   9334,
			Slug: "safelaunchpad",
			Name: "SafeLaunchpad",
		},
	},
	"VI": {
		{
			ID:          5428,
			Slug:        "vid",
			Name:        "Vid",
			BinanceUSDT: "VIUSDT",
		},
	},
	"GB": {
		{
			ID:          1285,
			Slug:        "goldblocks",
			Name:        "GoldBlocks",
			BinanceUSDT: "GBUSDT",
		},
	},
	"UTK": {
		{
			ID:          2320,
			Slug:        "utrust",
			Name:        "Utrust",
			BinanceUSDT: "UTKUSDT",
		},
	},
	"FX1": {
		{
			ID:          6444,
			Slug:        "fanzy",
			Name:        "FANZY",
			BinanceUSDT: "FX1USDT",
		},
	},
	"ILC": {
		{
			ID:          3617,
			Slug:        "ilcoin",
			Name:        "ILCOIN",
			BinanceUSDT: "ILCUSDT",
		},
	},
	"WOMI": {
		{
			ID:          9126,
			Slug:        "wrapped-ecomi",
			Name:        "Wrapped ECOMI",
			BinanceUSDT: "WOMIUSDT",
		},
	},
	"UNSAFEMOON": {
		{
			ID:          10198,
			Slug:        "unsafemoon",
			Name:        "UnSafeMoon",
			BinanceUSDT: "UNSAFEMOONUSDT",
		},
	},
	"SHROOM": {
		{
			ID:          6891,
			Slug:        "niftyx-protocol",
			Name:        "Niftyx Protocol",
			BinanceUSDT: "SHROOMUSDT",
		},
	},
	"VOTE": {
		{
			ID:          4792,
			Slug:        "agora",
			Name:        "Agora",
			BinanceUSDT: "VOTEUSDT",
		},
	},
	"BNF": {
		{
			ID:          7436,
			Slug:        "bonfi",
			Name:        "BonFi",
			BinanceUSDT: "BNFUSDT",
		},
	},
	"KAN": {
		{
			ID:          2934,
			Slug:        "bitkan",
			Name:        "BitKan",
			BinanceUSDT: "KANUSDT",
		},
	},
	"ELTCOIN": {
		{
			ID:          2147,
			Slug:        "eltcoin",
			Name:        "ELTCOIN",
			BinanceUSDT: "ELTCOINUSDT",
		},
	},
	"BOO": {
		{
			ID:          9683,
			Slug:        "booster",
			Name:        "Booster",
			BinanceUSDT: "BOOUSDT",
		},
	},
	"TCAP": {
		{
			ID:          9678,
			Slug:        "total-crypto-market-cap-token",
			Name:        "Total Crypto Market Cap Token",
			BinanceUSDT: "TCAPUSDT",
		},
	},
	"FOR": {
		{
			ID:          4118,
			Slug:        "the-force-protocol",
			Name:        "ForTube",
			BinanceUSDT: "FORUSDT",
		},
	},
	"MSS": {
		{
			ID:          8465,
			Slug:        "monster-slayer-share",
			Name:        "Monster Slayer Share",
			BinanceUSDT: "MSSUSDT",
		},
	},
	"ARDX": {
		{
			ID:          4985,
			Slug:        "ardcoin",
			Name:        "ArdCoin",
			BinanceUSDT: "ARDXUSDT",
		},
	},
	"ARTE": {
		{
			ID:          6437,
			Slug:        "ethart",
			Name:        "ethArt",
			BinanceUSDT: "ARTEUSDT",
		},
	},
	"YFBETA": {
		{
			ID:          6886,
			Slug:        "yfbeta",
			Name:        "yfBeta",
			BinanceUSDT: "YFBETAUSDT",
		},
	},
	"CARE": {
		{
			ID:          3266,
			Slug:        "carebit",
			Name:        "Carebit",
			BinanceUSDT: "CAREUSDT",
		},
	},
	"VITE": {
		{
			ID:          2937,
			Slug:        "vite",
			Name:        "VITE",
			BinanceUSDT: "VITEUSDT",
		},
	},
	"SOGE": {
		{
			ID:          9156,
			Slug:        "space-hoge",
			Name:        "Space Hoge",
			BinanceUSDT: "SOGEUSDT",
		},
	},
	"METP": {
		{
			ID:          5448,
			Slug:        "metaprediction",
			Name:        "Metaprediction",
			BinanceUSDT: "METPUSDT",
		},
	},
	"WAIF": {
		{
			ID:          6552,
			Slug:        "waifu-token",
			Name:        "Waifu Token",
			BinanceUSDT: "WAIFUSDT",
		},
	},
	"FUND": {
		{
			ID:          3854,
			Slug:        "unification",
			Name:        "Unification",
			BinanceUSDT: "FUNDUSDT",
		},
	},
	"PFR": {
		{
			ID:          2244,
			Slug:        "payfair",
			Name:        "Payfair",
			BinanceUSDT: "PFRUSDT",
		},
	},
	"BTCN": {
		{
			ID:   3068,
			Slug: "bitcoinote",
			Name: "BitcoiNote",
		},
		{
			ID:   10471,
			Slug: "bitcoin-networks",
			Name: "Bitcoin Networks",
		},
	},
	"FOXD": {
		{
			ID:          9434,
			Slug:        "foxdcoin",
			Name:        "FoxDcoin",
			BinanceUSDT: "FOXDUSDT",
		},
	},
	"ORAI": {
		{
			ID:          7533,
			Slug:        "oraichain-token",
			Name:        "Oraichain Token",
			BinanceUSDT: "ORAIUSDT",
		},
	},
	"YEA": {
		{
			ID:          7317,
			Slug:        "yeafinance",
			Name:        "YeaFinance",
			BinanceUSDT: "YEAUSDT",
		},
	},
	"SHIBARISE": {
		{
			ID:          10855,
			Slug:        "shiba-rise",
			Name:        "SHIBA RISE",
			BinanceUSDT: "SHIBARISEUSDT",
		},
	},
	"SXUT": {
		{
			ID:          2382,
			Slug:        "spectre-utility",
			Name:        "Spectre.ai Utility Token",
			BinanceUSDT: "SXUTUSDT",
		},
	},
	"TRBT": {
		{
			ID:          7047,
			Slug:        "tribute",
			Name:        "Tribute",
			BinanceUSDT: "TRBTUSDT",
		},
	},
	"ODDZ": {
		{
			ID:          8717,
			Slug:        "oddz",
			Name:        "Oddz",
			BinanceUSDT: "ODDZUSDT",
		},
	},
	"CCO": {
		{
			ID:          2241,
			Slug:        "ccore",
			Name:        "Ccore",
			BinanceUSDT: "CCOUSDT",
		},
	},
	"SAND": {
		{
			ID:          6210,
			Slug:        "the-sandbox",
			Name:        "The Sandbox",
			BinanceUSDT: "SANDUSDT",
		},
	},
	"ZUSD": {
		{
			ID:   8772,
			Slug: "zusd",
			Name: "ZUSD",
		},
		{
			ID:   8393,
			Slug: "zytara-dollar",
			Name: "Zytara dollar",
		},
	},
	"CHX": {
		{
			ID:          2673,
			Slug:        "we-own",
			Name:        "WeOwn",
			BinanceUSDT: "CHXUSDT",
		},
	},
	"KILI": {
		{
			ID:          9320,
			Slug:        "kilimanjaro",
			Name:        "KILIMANJARO",
			BinanceUSDT: "KILIUSDT",
		},
	},
	"2LC": {
		{
			ID:          9268,
			Slug:        "2local",
			Name:        "2local",
			BinanceUSDT: "2LCUSDT",
		},
	},
	"RBTC": {
		{
			ID:          3626,
			Slug:        "rsk-smart-bitcoin",
			Name:        "RSK Smart Bitcoin",
			BinanceUSDT: "RBTCUSDT",
		},
	},
	"KYSC": {
		{
			ID:          6110,
			Slug:        "kysc-token",
			Name:        "KYSC Token",
			BinanceUSDT: "KYSCUSDT",
		},
	},
	"CAN": {
		{
			ID:   8537,
			Slug: "channels",
			Name: "Channels",
		},
		{
			ID:   2343,
			Slug: "canyacoin",
			Name: "CanYaCoin",
		},
		{
			ID:   7258,
			Slug: "coinwaycoin",
			Name: "Coinwaycoin",
		},
	},
	"NEW": {
		{
			ID:          3871,
			Slug:        "newton",
			Name:        "Newton",
			BinanceUSDT: "NEWUSDT",
		},
	},
	"CATGE": {
		{
			ID:          10274,
			Slug:        "catge-coin",
			Name:        "Catge coin",
			BinanceUSDT: "CATGEUSDT",
		},
	},
	"DCT": {
		{
			ID:          1478,
			Slug:        "decent",
			Name:        "DECENT",
			BinanceUSDT: "DCTUSDT",
		},
	},
	"PERX": {
		{
			ID:          6468,
			Slug:        "peerex",
			Name:        "PeerEx",
			BinanceUSDT: "PERXUSDT",
		},
	},
	"OG": {
		{
			ID:          5309,
			Slug:        "og-fan-token",
			Name:        "OG Fan Token",
			BinanceUSDT: "OGUSDT",
		},
	},
	"XRC": {
		{
			ID:          3839,
			Slug:        "xrhodium",
			Name:        "xRhodium",
			BinanceUSDT: "XRCUSDT",
		},
	},
	"BNZ": {
		{
			ID:          7415,
			Slug:        "bonezyard",
			Name:        "BonezYard",
			BinanceUSDT: "BNZUSDT",
		},
	},
	"YT": {
		{
			ID:          6395,
			Slug:        "cherry-token",
			Name:        "Cherry Token",
			BinanceUSDT: "YTUSDT",
		},
	},
	"ZUR": {
		{
			ID:          1250,
			Slug:        "zurcoin",
			Name:        "Zurcoin",
			BinanceUSDT: "ZURUSDT",
		},
	},
	"DEXI": {
		{
			ID:          9830,
			Slug:        "dexioprotocol",
			Name:        "Dexioprotocol",
			BinanceUSDT: "DEXIUSDT",
		},
	},
	"UBXT": {
		{
			ID:          7040,
			Slug:        "upbots",
			Name:        "UpBots",
			BinanceUSDT: "UBXTUSDT",
		},
	},
	"STND": {
		{
			ID:          9251,
			Slug:        "standard-protocol",
			Name:        "Standard Protocol",
			BinanceUSDT: "STNDUSDT",
		},
	},
	"B360": {
		{
			ID:          9128,
			Slug:        "b360",
			Name:        "B360",
			BinanceUSDT: "B360USDT",
		},
	},
	"PNIX": {
		{
			ID:          9138,
			Slug:        "phoenixdefi-finance",
			Name:        "PhoenixDefi.Finance",
			BinanceUSDT: "PNIXUSDT",
		},
	},
	"HBAR": {
		{
			ID:          4642,
			Slug:        "hedera-hashgraph",
			Name:        "Hedera Hashgraph",
			BinanceUSDT: "HBARUSDT",
		},
	},
	"MAKI": {
		{
			ID:          10232,
			Slug:        "makiswap",
			Name:        "MakiSwap",
			BinanceUSDT: "MAKIUSDT",
		},
	},
	"KIMCHI": {
		{
			ID:          6839,
			Slug:        "kimchi-finance",
			Name:        "KIMCHI.finance",
			BinanceUSDT: "KIMCHIUSDT",
		},
	},
	"SASHIMI": {
		{
			ID:          6991,
			Slug:        "sashimi",
			Name:        "Sashimi",
			BinanceUSDT: "SASHIMIUSDT",
		},
	},
	"CENNZ": {
		{
			ID:          2585,
			Slug:        "centrality",
			Name:        "Centrality",
			BinanceUSDT: "CENNZUSDT",
		},
	},
	"ULG": {
		{
			ID:          4996,
			Slug:        "ultragate",
			Name:        "Ultragate",
			BinanceUSDT: "ULGUSDT",
		},
	},
	"ROBO": {
		{
			ID:          9054,
			Slug:        "robo-token",
			Name:        "Robo Token",
			BinanceUSDT: "ROBOUSDT",
		},
	},
	"USDN": {
		{
			ID:          5068,
			Slug:        "neutrino-usd",
			Name:        "Neutrino USD",
			BinanceUSDT: "USDNUSDT",
		},
	},
	"DFD": {
		{
			ID:          7593,
			Slug:        "defidollar-dao",
			Name:        "DefiDollar DAO",
			BinanceUSDT: "DFDUSDT",
		},
	},
	"VIAGRA": {
		{
			ID:          10037,
			Slug:        "viagra-token",
			Name:        "Viagra Token",
			BinanceUSDT: "VIAGRAUSDT",
		},
	},
	"ETHBTCRSI": {
		{
			ID:          6139,
			Slug:        "eth-btc-rsi-ratio-trading-set",
			Name:        "ETH/BTC RSI Ratio Trading Set",
			BinanceUSDT: "ETHBTCRSIUSDT",
		},
	},
	"ALIAS": {
		{
			ID:          1505,
			Slug:        "alias",
			Name:        "Alias",
			BinanceUSDT: "ALIASUSDT",
		},
	},
	"UUU": {
		{
			ID:          2645,
			Slug:        "u-network",
			Name:        "U Network",
			BinanceUSDT: "UUUUSDT",
		},
	},
	"MEDI": {
		{
			ID:          8012,
			Slug:        "mediconnectuk",
			Name:        "MediconnectUk",
			BinanceUSDT: "MEDIUSDT",
		},
	},
	"DEVA": {
		{
			ID:          8515,
			Slug:        "deva-token",
			Name:        "DEVA TOKEN",
			BinanceUSDT: "DEVAUSDT",
		},
	},
	"MINTME": {
		{
			ID:          3361,
			Slug:        "mintme-com-coin",
			Name:        "MintMe.com Coin",
			BinanceUSDT: "MINTMEUSDT",
		},
	},
	"BEST": {
		{
			ID:   4361,
			Slug: "bitpanda-ecosystem-token",
			Name: "Bitpanda Ecosystem Token",
		},
		{
			ID:   9738,
			Slug: "bitcoin-and-ethereum-standard-token",
			Name: "Bitcoin and Ethereum Standard Token",
		},
	},
	"AWX": {
		{
			ID:          7301,
			Slug:        "aurusdefi",
			Name:        "AurusDeFi",
			BinanceUSDT: "AWXUSDT",
		},
	},
	"AIT": {
		{
			ID:          2407,
			Slug:        "aichain",
			Name:        "AICHAIN",
			BinanceUSDT: "AITUSDT",
		},
	},
	"WDP": {
		{
			ID:          7865,
			Slug:        "waterdrop",
			Name:        "WaterDrop",
			BinanceUSDT: "WDPUSDT",
		},
	},
	"BUND": {
		{
			ID:          7785,
			Slug:        "bundles",
			Name:        "Bundles Finance",
			BinanceUSDT: "BUNDUSDT",
		},
	},
	"NTB": {
		{
			ID:          8284,
			Slug:        "tokenasset",
			Name:        "TokenAsset",
			BinanceUSDT: "NTBUSDT",
		},
	},
	"mQQQ": {
		{
			ID:          8025,
			Slug:        "mirrored-invesco-qqq-trust",
			Name:        "Mirrored Invesco QQQ Trust",
			BinanceUSDT: "mQQQUSDT",
		},
	},
	"RENZEC": {
		{
			ID:          8904,
			Slug:        "renzec",
			Name:        "renZEC",
			BinanceUSDT: "RENZECUSDT",
		},
	},
	"$ANRX": {
		{
			ID:          8057,
			Slug:        "anrkey-x",
			Name:        "AnRKey X",
			BinanceUSDT: "$ANRXUSDT",
		},
	},
	"DHC": {
		{
			ID:          7260,
			Slug:        "deltahub-community",
			Name:        "DeltaHub Community",
			BinanceUSDT: "DHCUSDT",
		},
	},
	"CHART": {
		{
			ID:          6861,
			Slug:        "chartex",
			Name:        "ChartEx",
			BinanceUSDT: "CHARTUSDT",
		},
	},
	"GLASS": {
		{
			ID:          10113,
			Slug:        "ourglass",
			Name:        "Ourglass",
			BinanceUSDT: "GLASSUSDT",
		},
	},
	"PEPPA": {
		{
			ID:          10432,
			Slug:        "peppa-network",
			Name:        "Peppa Network",
			BinanceUSDT: "PEPPAUSDT",
		},
	},
	"NOKN": {
		{
			ID:          5801,
			Slug:        "nokencoin",
			Name:        "Nokencoin",
			BinanceUSDT: "NOKNUSDT",
		},
	},
	"OMC": {
		{
			ID:          5818,
			Slug:        "ormeus-cash",
			Name:        "Ormeus Cash",
			BinanceUSDT: "OMCUSDT",
		},
	},
	"VIDY": {
		{
			ID:          4431,
			Slug:        "vidy",
			Name:        "VIDY",
			BinanceUSDT: "VIDYUSDT",
		},
	},
	"SXAU": {
		{
			ID:          6191,
			Slug:        "sxau",
			Name:        "sXAU",
			BinanceUSDT: "SXAUUSDT",
		},
	},
	"DSS": {
		{
			ID:          7145,
			Slug:        "defi-shopping-stake",
			Name:        "Defi Shopping Stake",
			BinanceUSDT: "DSSUSDT",
		},
	},
	"XAT": {
		{
			ID:          7488,
			Slug:        "shareat",
			Name:        "ShareAt",
			BinanceUSDT: "XATUSDT",
		},
	},
	"RAIL": {
		{
			ID:          10854,
			Slug:        "railgun",
			Name:        "Railgun",
			BinanceUSDT: "RAILUSDT",
		},
	},
	"ODEX": {
		{
			ID:          3954,
			Slug:        "one-dex",
			Name:        "One DEX",
			BinanceUSDT: "ODEXUSDT",
		},
	},
	"DISTX": {
		{
			ID:          6822,
			Slug:        "distx",
			Name:        "DistX",
			BinanceUSDT: "DISTXUSDT",
		},
	},
	"DCTO": {
		{
			ID:          3738,
			Slug:        "decentralized-crypto-token",
			Name:        "Decentralized Crypto Token",
			BinanceUSDT: "DCTOUSDT",
		},
	},
	"BUY": {
		{
			ID:          6701,
			Slug:        "burency",
			Name:        "Burency",
			BinanceUSDT: "BUYUSDT",
		},
	},
	"FWT": {
		{
			ID:          7585,
			Slug:        "freeway-token",
			Name:        "Freeway Token",
			BinanceUSDT: "FWTUSDT",
		},
	},
	"MISO": {
		{
			ID:          10309,
			Slug:        "miso",
			Name:        "MISO",
			BinanceUSDT: "MISOUSDT",
		},
	},
	"ETHP": {
		{
			ID:          5778,
			Slug:        "ethplus",
			Name:        "ETHPlus",
			BinanceUSDT: "ETHPUSDT",
		},
	},
	"IPL": {
		{
			ID:          2421,
			Slug:        "insurepal",
			Name:        "VouchForMe",
			BinanceUSDT: "IPLUSDT",
		},
	},
	"mIAU": {
		{
			ID:          8024,
			Slug:        "mirrored-ishares-gold-trust",
			Name:        "Mirrored iShares Gold Trust",
			BinanceUSDT: "mIAUUSDT",
		},
	},
	"PLEX": {
		{
			ID:          7980,
			Slug:        "mineplex",
			Name:        "MinePlex",
			BinanceUSDT: "PLEXUSDT",
		},
	},
	"APM": {
		{
			ID:          5079,
			Slug:        "apm-coin",
			Name:        "apM Coin",
			BinanceUSDT: "APMUSDT",
		},
	},
	"ARCH": {
		{
			ID:          7750,
			Slug:        "archer-dao-governance-token",
			Name:        "Archer DAO Governance Token",
			BinanceUSDT: "ARCHUSDT",
		},
	},
	"SETH": {
		{
			ID:   2555,
			Slug: "sether",
			Name: "Sether",
		},
		{
			ID:   5765,
			Slug: "seth",
			Name: "sETH",
		},
	},
	"BULK": {
		{
			ID:          10235,
			Slug:        "bulk",
			Name:        "Bulk",
			BinanceUSDT: "BULKUSDT",
		},
	},
	"SPRK": {
		{
			ID:          4010,
			Slug:        "sparkster",
			Name:        "Sparkster",
			BinanceUSDT: "SPRKUSDT",
		},
	},
	"MFTU": {
		{
			ID:          3189,
			Slug:        "mainstream-for-the-underground",
			Name:        "Mainstream For The Underground",
			BinanceUSDT: "MFTUUSDT",
		},
	},
	"LTC": {
		{
			ID:          2,
			Slug:        "litecoin",
			Name:        "Litecoin",
			BinanceUSDT: "LTCUSDT",
		},
	},
	"FLUX": {
		{
			ID:   3029,
			Slug: "zel",
			Name: "Flux",
		},
		{
			ID:   5876,
			Slug: "flux",
			Name: "Datamine FLUX",
		},
		{
			ID:   9837,
			Slug: "flux-protocol",
			Name: "Flux Protocol",
		},
	},
	"ECOREAL": {
		{
			ID:          3344,
			Slug:        "ecoreal-estate",
			Name:        "Ecoreal Estate",
			BinanceUSDT: "ECOREALUSDT",
		},
	},
	"USDK": {
		{
			ID:          4064,
			Slug:        "usdk",
			Name:        "USDK",
			BinanceUSDT: "USDKUSDT",
		},
	},
	"SUSD": {
		{
			ID:          2927,
			Slug:        "susd",
			Name:        "sUSD",
			BinanceUSDT: "SUSDUSDT",
		},
	},
	"ARGUS": {
		{
			ID:          1558,
			Slug:        "argus",
			Name:        "Argus",
			BinanceUSDT: "ARGUSUSDT",
		},
	},
	"NRV": {
		{
			ID:          8755,
			Slug:        "nerve-finance",
			Name:        "Nerve Finance",
			BinanceUSDT: "NRVUSDT",
		},
	},
	"VNS": {
		{
			ID:          7078,
			Slug:        "va-na-su",
			Name:        "Va Na Su",
			BinanceUSDT: "VNSUSDT",
		},
	},
	"SCU": {
		{
			ID:          8199,
			Slug:        "securypto",
			Name:        "Securypto",
			BinanceUSDT: "SCUUSDT",
		},
	},
	"MYID": {
		{
			ID:          9187,
			Slug:        "my-identity-coin",
			Name:        "MY IDENTITY COIN",
			BinanceUSDT: "MYIDUSDT",
		},
	},
	"WRX": {
		{
			ID:          5161,
			Slug:        "wazirx",
			Name:        "WazirX",
			BinanceUSDT: "WRXUSDT",
		},
	},
	"IDLE": {
		{
			ID:          7841,
			Slug:        "idle",
			Name:        "Idle",
			BinanceUSDT: "IDLEUSDT",
		},
	},
	"MEET": {
		{
			ID:          2470,
			Slug:        "coinmeet",
			Name:        "CoinMeet",
			BinanceUSDT: "MEETUSDT",
		},
	},
	"DVG": {
		{
			ID:          8478,
			Slug:        "daoventures",
			Name:        "DAOventures",
			BinanceUSDT: "DVGUSDT",
		},
	},
	"CROW": {
		{
			ID:          8533,
			Slug:        "crow-finance",
			Name:        "Crow Finance",
			BinanceUSDT: "CROWUSDT",
		},
	},
	"ERTH": {
		{
			ID:          10171,
			Slug:        "erth-token",
			Name:        "ERTH Token",
			BinanceUSDT: "ERTHUSDT",
		},
	},
	"CVX": {
		{
			ID:          9903,
			Slug:        "convex-finance",
			Name:        "Convex Finance",
			BinanceUSDT: "CVXUSDT",
		},
	},
	"PINKM": {
		{
			ID:          9596,
			Slug:        "pinkmoon",
			Name:        "PinkMoon",
			BinanceUSDT: "PINKMUSDT",
		},
	},
	"ACE": {
		{
			ID:   9792,
			Slug: "acent",
			Name: "ACENT",
		},
		{
			ID:   5717,
			Slug: "ace-entertainment",
			Name: "ACE",
		},
	},
	"WENLAMBO": {
		{
			ID:          9898,
			Slug:        "wenlambo",
			Name:        "Wenlambo",
			BinanceUSDT: "WENLAMBOUSDT",
		},
	},
	"HERO": {
		{
			ID:   10778,
			Slug: "metahero",
			Name: "Metahero",
		},
		{
			ID:   10315,
			Slug: "farmhero",
			Name: "FarmHero",
		},
	},
	"A5T": {
		{
			ID:          7933,
			Slug:        "alpha5",
			Name:        "Alpha5",
			BinanceUSDT: "A5TUSDT",
		},
	},
	"MEC": {
		{
			ID:   37,
			Slug: "megacoin",
			Name: "Megacoin",
		},
		{
			ID:   10138,
			Slug: "mechashiba",
			Name: "Mechashiba",
		},
	},
	"MODEX": {
		{
			ID:          4916,
			Slug:        "modex",
			Name:        "Modex",
			BinanceUSDT: "MODEXUSDT",
		},
	},
	"IMPACT": {
		{
			ID:          10389,
			Slug:        "alpha-impact",
			Name:        "Alpha Impact",
			BinanceUSDT: "IMPACTUSDT",
		},
	},
	"ECELL": {
		{
			ID:          7212,
			Slug:        "consensus-cell-network",
			Name:        "CellETF",
			BinanceUSDT: "ECELLUSDT",
		},
	},
	"DRAGON": {
		{
			ID:   8664,
			Slug: "dragonfarm-finance",
			Name: "DragonFarm Finance",
		},
		{
			ID:   6985,
			Slug: "dragon-ball",
			Name: "Dragon Ball",
		},
	},
	"SXPDOWN": {
		{
			ID:          7529,
			Slug:        "sxpdown",
			Name:        "SXPDOWN",
			BinanceUSDT: "SXPDOWNUSDT",
		},
	},
	"ESD": {
		{
			ID:          7033,
			Slug:        "empty-set-dollar",
			Name:        "Empty Set Dollar",
			BinanceUSDT: "ESDUSDT",
		},
	},
	"OBEE": {
		{
			ID:          6884,
			Slug:        "obee-network",
			Name:        "Obee Network",
			BinanceUSDT: "OBEEUSDT",
		},
	},
	"DEOR": {
		{
			ID:          8680,
			Slug:        "deor",
			Name:        "DEOR",
			BinanceUSDT: "DEORUSDT",
		},
	},
	"NFTI": {
		{
			ID:          8700,
			Slug:        "nft-index",
			Name:        "NFT Index",
			BinanceUSDT: "NFTIUSDT",
		},
	},
	"CRX": {
		{
			ID:          8731,
			Slug:        "cryptex",
			Name:        "CryptEx",
			BinanceUSDT: "CRXUSDT",
		},
	},
	"PAINT": {
		{
			ID:          8647,
			Slug:        "murall",
			Name:        "MurAll",
			BinanceUSDT: "PAINTUSDT",
		},
	},
	"AAPX": {
		{
			ID:          9171,
			Slug:        "ampnet-asset-platform-and-exchange",
			Name:        "AMPnet Asset Platform and Exchange",
			BinanceUSDT: "AAPXUSDT",
		},
	},
	"OX": {
		{
			ID:          9384,
			Slug:        "orcax",
			Name:        "OrcaX",
			BinanceUSDT: "OXUSDT",
		},
	},
	"OBIC": {
		{
			ID:          6982,
			Slug:        "obic",
			Name:        "OBIC",
			BinanceUSDT: "OBICUSDT",
		},
	},
	"FIT": {
		{
			ID:   8499,
			Slug: "300fit-network",
			Name: "300FIT NETWORK",
		},
		{
			ID:   8047,
			Slug: "first-interchangeable-token",
			Name: "FIRST INTERCHANGEABLE TOKEN",
		},
		{
			ID:   6984,
			Slug: "financial-investment-token",
			Name: "FINANCIAL INVESTMENT TOKEN",
		},
	},
	"MOONS": {
		{
			ID:          7406,
			Slug:        "moontools",
			Name:        "MoonTools",
			BinanceUSDT: "MOONSUSDT",
		},
	},
	"SPWN": {
		{
			ID:          10285,
			Slug:        "bitspawn-protocol",
			Name:        "Bitspawn Protocol",
			BinanceUSDT: "SPWNUSDT",
		},
	},
	"PLAY": {
		{
			ID:   2323,
			Slug: "herocoin",
			Name: "HEROcoin",
		},
		{
			ID:   9820,
			Slug: "metaverse-nft-index",
			Name: "Metaverse NFT Index",
		},
		{
			ID:   6265,
			Slug: "play-royal",
			Name: "Play Royal",
		},
	},
	"ETH20SMACO": {
		{
			ID:          6129,
			Slug:        "eth-20-day-ma-crossover-set",
			Name:        "ETH 20 Day MA Crossover Set",
			BinanceUSDT: "ETH20SMACOUSDT",
		},
	},
	"SPANDA": {
		{
			ID:          10076,
			Slug:        "superpanda",
			Name:        "Superpanda",
			BinanceUSDT: "SPANDAUSDT",
		},
	},
	"AWF": {
		{
			ID:          10043,
			Slug:        "alphawolf-finance",
			Name:        "Alphawolf Finance",
			BinanceUSDT: "AWFUSDT",
		},
	},
	"QQQ": {
		{
			ID:          4644,
			Slug:        "poseidon-network",
			Name:        "Poseidon Network",
			BinanceUSDT: "QQQUSDT",
		},
	},
	"MIX": {
		{
			ID:          4366,
			Slug:        "mixmarvel",
			Name:        "MixMarvel",
			BinanceUSDT: "MIXUSDT",
		},
	},
	"BKKG": {
		{
			ID:          7753,
			Slug:        "biokkoin",
			Name:        "BIOKKOIN",
			BinanceUSDT: "BKKGUSDT",
		},
	},
	"BONO": {
		{
			ID:          5320,
			Slug:        "bonorum",
			Name:        "Bonorum",
			BinanceUSDT: "BONOUSDT",
		},
	},
	"BZ": {
		{
			ID:          2918,
			Slug:        "bit-z-token",
			Name:        "BitZ Token",
			BinanceUSDT: "BZUSDT",
		},
	},
	"LONG": {
		{
			ID:          7668,
			Slug:        "long-coin",
			Name:        "LONG COIN",
			BinanceUSDT: "LONGUSDT",
		},
	},
	"ZEP": {
		{
			ID:          9300,
			Slug:        "zeppelin-dao",
			Name:        "Zeppelin DAO",
			BinanceUSDT: "ZEPUSDT",
		},
	},
	"WINALAMBO": {
		{
			ID:          10745,
			Slug:        "winalambo-finance",
			Name:        "WIN A LAMBO FINANCE",
			BinanceUSDT: "WINALAMBOUSDT",
		},
	},
	"DAI": {
		{
			ID:          4943,
			Slug:        "multi-collateral-dai",
			Name:        "Dai",
			BinanceUSDT: "DAIUSDT",
		},
	},
	"IHF": {
		{
			ID:          3301,
			Slug:        "invictus-hyperion-fund",
			Name:        "Invictus Hyperion Fund",
			BinanceUSDT: "IHFUSDT",
		},
	},
	"PAC": {
		{
			ID:          1107,
			Slug:        "pac-protocol",
			Name:        "PAC Protocol",
			BinanceUSDT: "PACUSDT",
		},
	},
	"BCNA": {
		{
			ID:          4263,
			Slug:        "bitcanna",
			Name:        "BitCanna",
			BinanceUSDT: "BCNAUSDT",
		},
	},
	"STARS": {
		{
			ID:          8996,
			Slug:        "mogul-productions",
			Name:        "Mogul Productions",
			BinanceUSDT: "STARSUSDT",
		},
	},
	"METM": {
		{
			ID:          3148,
			Slug:        "metamorph",
			Name:        "MetaMorph",
			BinanceUSDT: "METMUSDT",
		},
	},
	"GRO": {
		{
			ID:          6718,
			Slug:        "growthdefi",
			Name:        "Growth DeFi",
			BinanceUSDT: "GROUSDT",
		},
	},
	"BLTG": {
		{
			ID:          3627,
			Slug:        "block-logic",
			Name:        "Block-Logic",
			BinanceUSDT: "BLTGUSDT",
		},
	},
	"STRK": {
		{
			ID:          8911,
			Slug:        "strike",
			Name:        "Strike",
			BinanceUSDT: "STRKUSDT",
		},
	},
	"CHAR": {
		{
			ID:          9361,
			Slug:        "charitas",
			Name:        "Charitas",
			BinanceUSDT: "CHARUSDT",
		},
	},
	"FULLSEND": {
		{
			ID:          10022,
			Slug:        "full-send",
			Name:        "Full Send",
			BinanceUSDT: "FULLSENDUSDT",
		},
	},
	"SOUL": {
		{
			ID:   2827,
			Slug: "phantasma",
			Name: "Phantasma",
		},
		{
			ID:   3501,
			Slug: "cryptosoul",
			Name: "CryptoSoul",
		},
		{
			ID:   8684,
			Slug: "apoyield",
			Name: "APOyield",
		},
		{
			ID:   5729,
			Slug: "chainz-arena",
			Name: "ChainZ Arena",
		},
	},
	"PERP": {
		{
			ID:          6950,
			Slug:        "perpetual-protocol",
			Name:        "Perpetual Protocol",
			BinanceUSDT: "PERPUSDT",
		},
	},
	"BASI": {
		{
			ID:          8824,
			Slug:        "asi-finance",
			Name:        "ASI.finance",
			BinanceUSDT: "BASIUSDT",
		},
	},
	"LHT": {
		{
			ID:          4230,
			Slug:        "lighthouse-token",
			Name:        "LHT",
			BinanceUSDT: "LHTUSDT",
		},
	},
	"LAYERX": {
		{
			ID:          9559,
			Slug:        "unilayerx",
			Name:        "UNILAYERX",
			BinanceUSDT: "LAYERXUSDT",
		},
	},
	"CTI": {
		{
			ID:          7860,
			Slug:        "clintex-cti",
			Name:        "ClinTex CTi",
			BinanceUSDT: "CTIUSDT",
		},
	},
	"DOBE": {
		{
			ID:          10205,
			Slug:        "dobermann",
			Name:        "Dobermann",
			BinanceUSDT: "DOBEUSDT",
		},
	},
	"ATM": {
		{
			ID:          5227,
			Slug:        "atletico-de-madrid-fan-token",
			Name:        "Atletico De Madrid Fan Token",
			BinanceUSDT: "ATMUSDT",
		},
	},
	"YAXIS": {
		{
			ID:          7222,
			Slug:        "yaxis",
			Name:        "yAxis",
			BinanceUSDT: "YAXISUSDT",
		},
	},
	"CATX": {
		{
			ID:          6703,
			Slug:        "cat-trade-protocol",
			Name:        "CAT.trade Protocol",
			BinanceUSDT: "CATXUSDT",
		},
	},
	"GS": {
		{
			ID:          9196,
			Slug:        "genesis-shards",
			Name:        "Genesis Shards",
			BinanceUSDT: "GSUSDT",
		},
	},
	"OPT": {
		{
			ID:   6694,
			Slug: "open-predict-token",
			Name: "Open Predict Token",
		},
		{
			ID:   1931,
			Slug: "opus",
			Name: "Opus",
		},
	},
	"BSHIBA": {
		{
			ID:          9829,
			Slug:        "shiba-corp",
			Name:        "Shiba Corp",
			BinanceUSDT: "BSHIBAUSDT",
		},
	},
	"NAR": {
		{
			ID:          7831,
			Slug:        "narwhalswap",
			Name:        "Narwhalswap",
			BinanceUSDT: "NARUSDT",
		},
	},
	"MGC": {
		{
			ID:          4048,
			Slug:        "mgc-token",
			Name:        "MGC Token",
			BinanceUSDT: "MGCUSDT",
		},
	},
	"ZDEX": {
		{
			ID:          6989,
			Slug:        "zeedex",
			Name:        "Zeedex",
			BinanceUSDT: "ZDEXUSDT",
		},
	},
	"NXM": {
		{
			ID:          5830,
			Slug:        "nxm",
			Name:        "NXM",
			BinanceUSDT: "NXMUSDT",
		},
	},
	"RDN": {
		{
			ID:          2161,
			Slug:        "raiden-network-token",
			Name:        "Raiden Network Token",
			BinanceUSDT: "RDNUSDT",
		},
	},
	"ROCO": {
		{
			ID:          3677,
			Slug:        "roiyal-coin",
			Name:        "ROIyal Coin",
			BinanceUSDT: "ROCOUSDT",
		},
	},
	"EOSC": {
		{
			ID:          4769,
			Slug:        "eos-force",
			Name:        "EOS Force",
			BinanceUSDT: "EOSCUSDT",
		},
	},
	"ARX": {
		{
			ID:          5005,
			Slug:        "arcs",
			Name:        "ARCS",
			BinanceUSDT: "ARXUSDT",
		},
	},
	"RX": {
		{
			ID:          9227,
			Slug:        "raven-x",
			Name:        "Raven X",
			BinanceUSDT: "RXUSDT",
		},
	},
	"LQD": {
		{
			ID:          3499,
			Slug:        "liquidity-network",
			Name:        "Liquidity Network",
			BinanceUSDT: "LQDUSDT",
		},
	},
	"BLK": {
		{
			ID:          170,
			Slug:        "blackcoin",
			Name:        "BlackCoin",
			BinanceUSDT: "BLKUSDT",
		},
	},
	"TTN": {
		{
			ID:   3913,
			Slug: "titan-coin",
			Name: "Titan Coin",
		},
		{
			ID:   8396,
			Slug: "thetoken-network",
			Name: "TheToken.Network",
		},
	},
	"DONU": {
		{
			ID:          551,
			Slug:        "donu",
			Name:        "Donu",
			BinanceUSDT: "DONUUSDT",
		},
	},
	"DUCK": {
		{
			ID:   8063,
			Slug: "duck-dao",
			Name: "Duck DAO (DLP Duck Token)",
		},
		{
			ID:   7977,
			Slug: "unit-protocol-duck",
			Name: "Unit Protocol Duck",
		},
	},
	"FED": {
		{
			ID:          4555,
			Slug:        "fedora-gold",
			Name:        "Fedora Gold",
			BinanceUSDT: "FEDUSDT",
		},
	},
	"BCH": {
		{
			ID:          1831,
			Slug:        "bitcoin-cash",
			Name:        "Bitcoin Cash",
			BinanceUSDT: "BCHUSDT",
		},
	},
	"NANJ": {
		{
			ID:          2595,
			Slug:        "nanjcoin",
			Name:        "NANJCOIN",
			BinanceUSDT: "NANJUSDT",
		},
	},
	"ZUT": {
		{
			ID:          7318,
			Slug:        "zero-utility-token",
			Name:        "Zero Utility Token",
			BinanceUSDT: "ZUTUSDT",
		},
	},
	"TTF": {
		{
			ID:          10847,
			Slug:        "turbotrix-finance",
			Name:        "TurboTrix Finance",
			BinanceUSDT: "TTFUSDT",
		},
	},
	"HELMET": {
		{
			ID:          8265,
			Slug:        "helmet-insure",
			Name:        "Helmet.insure",
			BinanceUSDT: "HELMETUSDT",
		},
	},
	"mSLV": {
		{
			ID:          8026,
			Slug:        "mirrored-ishares-silver-trust",
			Name:        "Mirrored iShares Silver Trust",
			BinanceUSDT: "mSLVUSDT",
		},
	},
	"BABYCUBAN": {
		{
			ID:          10995,
			Slug:        "baby-cuban",
			Name:        "Baby Cuban",
			BinanceUSDT: "BABYCUBANUSDT",
		},
	},
	"PHR": {
		{
			ID:          2158,
			Slug:        "phore",
			Name:        "Phore",
			BinanceUSDT: "PHRUSDT",
		},
	},
	"SNRG": {
		{
			ID:          951,
			Slug:        "synergy",
			Name:        "Synergy",
			BinanceUSDT: "SNRGUSDT",
		},
	},
	"TRADE": {
		{
			ID:   6195,
			Slug: "unitrade",
			Name: "Unitrade",
		},
		{
			ID:   10108,
			Slug: "smart-trade-coin",
			Name: "Smart Trade Coin",
		},
	},
	"MVEDA": {
		{
			ID:          7641,
			Slug:        "medicalveda",
			Name:        "Medicalveda",
			BinanceUSDT: "MVEDAUSDT",
		},
	},
	"ETHMACOAPY": {
		{
			ID:          6130,
			Slug:        "eth-20-day-ma-crossover-yield-set",
			Name:        "ETH 20 Day MA Crossover Yield Set",
			BinanceUSDT: "ETHMACOAPYUSDT",
		},
	},
	"AMD": {
		{
			ID:          7903,
			Slug:        "advanced-micro-devices-tokenized-stock-ftx",
			Name:        "Advanced Micro Devices tokenized stock FTX",
			BinanceUSDT: "AMDUSDT",
		},
	},
	"HPS": {
		{
			ID:          8198,
			Slug:        "happinesstoken",
			Name:        "HappinessToken",
			BinanceUSDT: "HPSUSDT",
		},
	},
	"EGEM": {
		{
			ID:          3132,
			Slug:        "ethergem",
			Name:        "EtherGem",
			BinanceUSDT: "EGEMUSDT",
		},
	},
	"CBET": {
		{
			ID:   10925,
			Slug: "cbet-token",
			Name: "CBET Token",
		},
		{
			ID:   5651,
			Slug: "cryptobet",
			Name: "CryptoBet",
		},
	},
	"QTUM": {
		{
			ID:          1684,
			Slug:        "qtum",
			Name:        "Qtum",
			BinanceUSDT: "QTUMUSDT",
		},
	},
	"EARN": {
		{
			ID:          7586,
			Slug:        "yearn-classic-finance",
			Name:        "Yearn Classic Finance",
			BinanceUSDT: "EARNUSDT",
		},
	},
	"VNLA": {
		{
			ID:          7776,
			Slug:        "vanilla-network",
			Name:        "Vanilla Network",
			BinanceUSDT: "VNLAUSDT",
		},
	},
	"MCS": {
		{
			ID:          10731,
			Slug:        "mcs-token",
			Name:        "MCS Token",
			BinanceUSDT: "MCSUSDT",
		},
	},
	"BEER": {
		{
			ID:          5449,
			Slug:        "beer-money",
			Name:        "Beer Money",
			BinanceUSDT: "BEERUSDT",
		},
	},
	"HEDG": {
		{
			ID:          3662,
			Slug:        "hedgetrade",
			Name:        "HedgeTrade",
			BinanceUSDT: "HEDGUSDT",
		},
	},
	"GEO": {
		{
			ID:   5590,
			Slug: "geodb",
			Name: "GeoDB",
		},
		{
			ID:   823,
			Slug: "geocoin",
			Name: "GeoCoin",
		},
	},
	"VEGA": {
		{
			ID:          10223,
			Slug:        "vegaprotocol",
			Name:        "Vega Protocol",
			BinanceUSDT: "VEGAUSDT",
		},
	},
	"CHT": {
		{
			ID:          4701,
			Slug:        "coinhe-token",
			Name:        "CoinHe Token",
			BinanceUSDT: "CHTUSDT",
		},
	},
	"vBCH": {
		{
			ID:          7974,
			Slug:        "venus-bch",
			Name:        "Venus BCH",
			BinanceUSDT: "vBCHUSDT",
		},
	},
	"BSCGOLD": {
		{
			ID:          9759,
			Slug:        "bsc-gold",
			Name:        "BSC Gold",
			BinanceUSDT: "BSCGOLDUSDT",
		},
	},
	"STMX": {
		{
			ID:          2297,
			Slug:        "stormx",
			Name:        "StormX",
			BinanceUSDT: "STMXUSDT",
		},
	},
	"vUSDC": {
		{
			ID:          7958,
			Slug:        "venus-usdc",
			Name:        "Venus USDC",
			BinanceUSDT: "vUSDCUSDT",
		},
	},
	"HGET": {
		{
			ID:          6949,
			Slug:        "hedget",
			Name:        "Hedget",
			BinanceUSDT: "HGETUSDT",
		},
	},
	"SUBX": {
		{
			ID:          10741,
			Slug:        "startup-boost-token",
			Name:        "Startup Boost Token",
			BinanceUSDT: "SUBXUSDT",
		},
	},
	"MAUSDT": {
		{
			ID:          8919,
			Slug:        "matic-aave-usdt",
			Name:        "Matic Aave Interest Bearing USDT",
			BinanceUSDT: "MAUSDTUSDT",
		},
	},
	"INSTAR": {
		{
			ID:          2558,
			Slug:        "insights-network",
			Name:        "Insights Network",
			BinanceUSDT: "INSTARUSDT",
		},
	},
	"BCDN": {
		{
			ID:          2247,
			Slug:        "blockcdn",
			Name:        "BlockCDN",
			BinanceUSDT: "BCDNUSDT",
		},
	},
	"VLC": {
		{
			ID:          2488,
			Slug:        "valuechain",
			Name:        "ValueChain",
			BinanceUSDT: "VLCUSDT",
		},
	},
	"SMARTCREDIT": {
		{
			ID:          7596,
			Slug:        "smartcredit-token",
			Name:        "SmartCredit Token",
			BinanceUSDT: "SMARTCREDITUSDT",
		},
	},
	"ETNXP": {
		{
			ID:          5669,
			Slug:        "electronero-pulse",
			Name:        "Electronero Pulse",
			BinanceUSDT: "ETNXPUSDT",
		},
	},
	"YAY": {
		{
			ID:          10104,
			Slug:        "yayswap",
			Name:        "YaySwap",
			BinanceUSDT: "YAYUSDT",
		},
	},
	"WSCRT": {
		{
			ID:          8337,
			Slug:        "secret-erc20",
			Name:        "Secret (ERC20)",
			BinanceUSDT: "WSCRTUSDT",
		},
	},
	"COCOS": {
		{
			ID:          4275,
			Slug:        "cocos-bcx",
			Name:        "Cocos-BCX",
			BinanceUSDT: "COCOSUSDT",
		},
	},
	"IG": {
		{
			ID:          3121,
			Slug:        "igtoken",
			Name:        "IGToken",
			BinanceUSDT: "IGUSDT",
		},
	},
	"BF": {
		{
			ID:          4283,
			Slug:        "bitforex-token",
			Name:        "BitForex Token",
			BinanceUSDT: "BFUSDT",
		},
	},
	"STRONG": {
		{
			ID:          6511,
			Slug:        "strong",
			Name:        "Strong",
			BinanceUSDT: "STRONGUSDT",
		},
	},
	"CPAY": {
		{
			ID:   2314,
			Slug: "cryptopay",
			Name: "Cryptopay",
		},
		{
			ID:   5355,
			Slug: "chainpay",
			Name: "Chainpay",
		},
	},
	"XLMDOWN": {
		{
			ID:          8054,
			Slug:        "xlmdown",
			Name:        "XLMDOWN",
			BinanceUSDT: "XLMDOWNUSDT",
		},
	},
	"OXO": {
		{
			ID:          9569,
			Slug:        "oxo-farm",
			Name:        "OXO.Farm",
			BinanceUSDT: "OXOUSDT",
		},
	},
	"CUST": {
		{
			ID:          4240,
			Slug:        "custody-token",
			Name:        "Custody Token",
			BinanceUSDT: "CUSTUSDT",
		},
	},
	"HOKK": {
		{
			ID:          9458,
			Slug:        "hokkaido-inu",
			Name:        "Hokkaidu Inu",
			BinanceUSDT: "HOKKUSDT",
		},
	},
	"BTB": {
		{
			ID:          3687,
			Slug:        "bitball",
			Name:        "BitBall",
			BinanceUSDT: "BTBUSDT",
		},
	},
	"ELAND": {
		{
			ID:          9492,
			Slug:        "etherland",
			Name:        "ETHERLAND",
			BinanceUSDT: "ELANDUSDT",
		},
	},
	"MBL": {
		{
			ID:          4038,
			Slug:        "moviebloc",
			Name:        "MovieBloc",
			BinanceUSDT: "MBLUSDT",
		},
	},
	"HPPOT": {
		{
			ID:          10288,
			Slug:        "healing-potion",
			Name:        "Healing Potion",
			BinanceUSDT: "HPPOTUSDT",
		},
	},
	"SUNI": {
		{
			ID:          9823,
			Slug:        "suni",
			Name:        "SUNI",
			BinanceUSDT: "SUNIUSDT",
		},
	},
	"DFIP": {
		{
			ID:          7333,
			Slug:        "defi-insurance-protocol",
			Name:        "DeFi Insurance Protocol",
			BinanceUSDT: "DFIPUSDT",
		},
	},
	"FFF": {
		{
			ID:   10431,
			Slug: "future-of-finance-fund",
			Name: "Future Of Finance Fund",
		},
		{
			ID:   6448,
			Slug: "force-for-fast",
			Name: "Force For Fast",
		},
	},
	"OBTC": {
		{
			ID:          3152,
			Slug:        "obitan-chain",
			Name:        "Obitan Chain",
			BinanceUSDT: "OBTCUSDT",
		},
	},
	"J8T": {
		{
			ID:          2568,
			Slug:        "jet8",
			Name:        "JET8",
			BinanceUSDT: "J8TUSDT",
		},
	},
	"NZL": {
		{
			ID:          3624,
			Slug:        "zealium",
			Name:        "Zealium",
			BinanceUSDT: "NZLUSDT",
		},
	},
	"ARAW": {
		{
			ID:          3792,
			Slug:        "araw",
			Name:        "ARAW",
			BinanceUSDT: "ARAWUSDT",
		},
	},
	"NSURE": {
		{
			ID:          7231,
			Slug:        "nsure-network",
			Name:        "Nsure.Network",
			BinanceUSDT: "NSUREUSDT",
		},
	},
	"vUSDT": {
		{
			ID:          7957,
			Slug:        "venus-usdt",
			Name:        "Venus USDT",
			BinanceUSDT: "vUSDTUSDT",
		},
	},
	"PEECH": {
		{
			ID:          10356,
			Slug:        "peach-finance",
			Name:        "Peach.Finance",
			BinanceUSDT: "PEECHUSDT",
		},
	},
	"ADGNZ": {
		{
			ID:          10259,
			Slug:        "degen-token-finance",
			Name:        "Degen Token Finance",
			BinanceUSDT: "ADGNZUSDT",
		},
	},
	"CORGI": {
		{
			ID:   9968,
			Slug: "corgidoge-real-estate-payment",
			Name: "Corgidoge real estate payment",
		},
		{
			ID:   9939,
			Slug: "corgi-inu",
			Name: "Corgi inu",
		},
	},
	"LAR": {
		{
			ID:          5251,
			Slug:        "linkart",
			Name:        "LinkArt",
			BinanceUSDT: "LARUSDT",
		},
	},
	"GGIVE": {
		{
			ID:          9815,
			Slug:        "globalgive",
			Name:        "GlobalGive",
			BinanceUSDT: "GGIVEUSDT",
		},
	},
	"GODL": {
		{
			ID:          10947,
			Slug:        "godl",
			Name:        "GODL",
			BinanceUSDT: "GODLUSDT",
		},
	},
	"SNET": {
		{
			ID:          3435,
			Slug:        "snetwork",
			Name:        "Snetwork",
			BinanceUSDT: "SNETUSDT",
		},
	},
	"HOO": {
		{
			ID:          7543,
			Slug:        "hoo-token",
			Name:        "Hoo Token",
			BinanceUSDT: "HOOUSDT",
		},
	},
	"NETKO": {
		{
			ID:          1582,
			Slug:        "netko",
			Name:        "Netko",
			BinanceUSDT: "NETKOUSDT",
		},
	},
	"BSE": {
		{
			ID:          7996,
			Slug:        "buy-sell",
			Name:        "Buy-Sell",
			BinanceUSDT: "BSEUSDT",
		},
	},
	"LNC": {
		{
			ID:          2677,
			Slug:        "linker-coin",
			Name:        "Linker Coin",
			BinanceUSDT: "LNCUSDT",
		},
	},
	"TN": {
		{
			ID:          4782,
			Slug:        "turtlenetwork",
			Name:        "TurtleNetwork",
			BinanceUSDT: "TNUSDT",
		},
	},
	"GBK": {
		{
			ID:          7581,
			Slug:        "goldblock",
			Name:        "Goldblock",
			BinanceUSDT: "GBKUSDT",
		},
	},
	"ARCO": {
		{
			ID:          1247,
			Slug:        "aquariuscoin",
			Name:        "AquariusCoin",
			BinanceUSDT: "ARCOUSDT",
		},
	},
	"GXX": {
		{
			ID:          977,
			Slug:        "gravitycoin",
			Name:        "GravityCoin",
			BinanceUSDT: "GXXUSDT",
		},
	},
	"TAN": {
		{
			ID:          4272,
			Slug:        "taklimakan-network",
			Name:        "Taklimakan Network",
			BinanceUSDT: "TANUSDT",
		},
	},
	"ZET": {
		{
			ID:          56,
			Slug:        "zetacoin",
			Name:        "Zetacoin",
			BinanceUSDT: "ZETUSDT",
		},
	},
	"DINGO": {
		{
			ID:          9982,
			Slug:        "dingo-token",
			Name:        "DINGO TOKEN",
			BinanceUSDT: "DINGOUSDT",
		},
	},
	"SAFERUNE": {
		{
			ID:          9734,
			Slug:        "saferune",
			Name:        "Saferune",
			BinanceUSDT: "SAFERUNEUSDT",
		},
	},
	"TIDAL": {
		{
			ID:          8912,
			Slug:        "tidal-finance",
			Name:        "Tidal Finance",
			BinanceUSDT: "TIDALUSDT",
		},
	},
	"MINI": {
		{
			ID:          6405,
			Slug:        "miniswap",
			Name:        "MiniSwap",
			BinanceUSDT: "MINIUSDT",
		},
	},
	"OTB": {
		{
			ID:          2894,
			Slug:        "otcbtc-token",
			Name:        "OTCBTC Token",
			BinanceUSDT: "OTBUSDT",
		},
	},
	"RLE": {
		{
			ID:          7990,
			Slug:        "richlab-token",
			Name:        "Richlab Token",
			BinanceUSDT: "RLEUSDT",
		},
	},
	"SWC": {
		{
			ID:          3760,
			Slug:        "scanetchain",
			Name:        "Scanetchain",
			BinanceUSDT: "SWCUSDT",
		},
	},
	"GERA": {
		{
			ID:          8270,
			Slug:        "gera-coin",
			Name:        "Gera Coin",
			BinanceUSDT: "GERAUSDT",
		},
	},
	"DMC": {
		{
			ID:          8610,
			Slug:        "decentralized-mining-exchange",
			Name:        "Decentralized Mining Exchange",
			BinanceUSDT: "DMCUSDT",
		},
	},
	"RAGE": {
		{
			ID:          8862,
			Slug:        "rage-fan",
			Name:        "Rage Fan",
			BinanceUSDT: "RAGEUSDT",
		},
	},
	"XDEF2": {
		{
			ID:          8226,
			Slug:        "xdef-finance",
			Name:        "Xdef Finance",
			BinanceUSDT: "XDEF2USDT",
		},
	},
	"SUB": {
		{
			ID:          1984,
			Slug:        "substratum",
			Name:        "Substratum",
			BinanceUSDT: "SUBUSDT",
		},
	},
	"DCTD": {
		{
			ID:          9066,
			Slug:        "dctdao",
			Name:        "DCTDAO",
			BinanceUSDT: "DCTDUSDT",
		},
	},
	"DRC": {
		{
			ID:   7380,
			Slug: "dracula-token",
			Name: "Dracula Token",
		},
		{
			ID:   7420,
			Slug: "digital-reserve-currency",
			Name: "Digital Reserve Currency",
		},
		{
			ID:   8124,
			Slug: "drc-mobility",
			Name: "DRC mobility",
		},
	},
	"UQC": {
		{
			ID:          2273,
			Slug:        "uquid-coin",
			Name:        "Uquid Coin",
			BinanceUSDT: "UQCUSDT",
		},
	},
	"FAN8": {
		{
			ID:          10701,
			Slug:        "fan8",
			Name:        "FAN8",
			BinanceUSDT: "FAN8USDT",
		},
	},
	"LINKA": {
		{
			ID:          4850,
			Slug:        "linka",
			Name:        "LINKA",
			BinanceUSDT: "LINKAUSDT",
		},
	},
	"THOR": {
		{
			ID:          8663,
			Slug:        "asgard-finance",
			Name:        "Asgard finance",
			BinanceUSDT: "THORUSDT",
		},
	},
	"VLS": {
		{
			ID:          4136,
			Slug:        "veles",
			Name:        "Veles",
			BinanceUSDT: "VLSUSDT",
		},
	},
	"NINJA": {
		{
			ID:          4609,
			Slug:        "ninjacoin",
			Name:        "Ninjacoin",
			BinanceUSDT: "NINJAUSDT",
		},
	},
	"KAWAII": {
		{
			ID:          10193,
			Slug:        "kawai-inu",
			Name:        "Kawai INU",
			BinanceUSDT: "KAWAIIUSDT",
		},
	},
	"SAFEYIELD": {
		{
			ID:          9256,
			Slug:        "safeyield",
			Name:        "SafeYield",
			BinanceUSDT: "SAFEYIELDUSDT",
		},
	},
	"MTN": {
		{
			ID:          2497,
			Slug:        "medical-chain",
			Name:        "Medicalchain",
			BinanceUSDT: "MTNUSDT",
		},
	},
	"MCO2": {
		{
			ID:          8826,
			Slug:        "moss-carbon-credit",
			Name:        "Moss Carbon Credit",
			BinanceUSDT: "MCO2USDT",
		},
	},
	"GDR": {
		{
			ID:          4881,
			Slug:        "guider",
			Name:        "Guider",
			BinanceUSDT: "GDRUSDT",
		},
	},
	"CLIMB": {
		{
			ID:          9295,
			Slug:        "climb-token-finance",
			Name:        "CLIMB TOKEN FINANCE",
			BinanceUSDT: "CLIMBUSDT",
		},
	},
	"HL": {
		{
			ID:          6258,
			Slug:        "hl-chain",
			Name:        "HL Chain",
			BinanceUSDT: "HLUSDT",
		},
	},
	"BNBBULL": {
		{
			ID:          5295,
			Slug:        "3x-long-bnb-token",
			Name:        "3X Long BNB Token",
			BinanceUSDT: "BNBBULLUSDT",
		},
	},
	"QBX": {
		{
			ID:          4100,
			Slug:        "qiibee",
			Name:        "qiibee",
			BinanceUSDT: "QBXUSDT",
		},
	},
	"ZIG": {
		{
			ID:          9260,
			Slug:        "zigcoin",
			Name:        "Zigcoin",
			BinanceUSDT: "ZIGUSDT",
		},
	},
	"KICK": {
		{
			ID:          10700,
			Slug:        "kick-token-new",
			Name:        "KickToken [new]",
			BinanceUSDT: "KICKUSDT",
		},
	},
	"GENS": {
		{
			ID:          10278,
			Slug:        "genshiro",
			Name:        "Genshiro",
			BinanceUSDT: "GENSUSDT",
		},
	},
	"THETA": {
		{
			ID:          2416,
			Slug:        "theta",
			Name:        "THETA",
			BinanceUSDT: "THETAUSDT",
		},
	},
	"CDEX": {
		{
			ID:          4801,
			Slug:        "codex",
			Name:        "Codex",
			BinanceUSDT: "CDEXUSDT",
		},
	},
	"ALGOBEAR": {
		{
			ID:          6091,
			Slug:        "3x-short-algorand-token",
			Name:        "3X Short Algorand Token",
			BinanceUSDT: "ALGOBEARUSDT",
		},
	},
	"UAXIE": {
		{
			ID:          9476,
			Slug:        "unicly-mystic-axies-collection",
			Name:        "Unicly Mystic Axies Collection",
			BinanceUSDT: "UAXIEUSDT",
		},
	},
	"SAFEP": {
		{
			ID:          9266,
			Slug:        "safe-protocol",
			Name:        "Safe Protocol",
			BinanceUSDT: "SAFEPUSDT",
		},
	},
	"$TRDL": {
		{
			ID:          8390,
			Slug:        "strudel-finance",
			Name:        "Strudel Finance",
			BinanceUSDT: "$TRDLUSDT",
		},
	},
	"RARI": {
		{
			ID:          5877,
			Slug:        "rarible",
			Name:        "Rarible",
			BinanceUSDT: "RARIUSDT",
		},
	},
	"BLINK": {
		{
			ID:   7784,
			Slug: "blink",
			Name: "BLink",
		},
		{
			ID:   4089,
			Slug: "blockmason-link",
			Name: "Blockmason Link",
		},
	},
	"NYAN-2": {
		{
			ID:          7643,
			Slug:        "nyan-v2",
			Name:        "Nyan V2",
			BinanceUSDT: "NYAN-2USDT",
		},
	},
	"UNIDOWN": {
		{
			ID:          7525,
			Slug:        "unidown",
			Name:        "UNIDOWN",
			BinanceUSDT: "UNIDOWNUSDT",
		},
	},
	"BR": {
		{
			ID:          10624,
			Slug:        "bull-run-finance",
			Name:        "Bull Run Finance",
			BinanceUSDT: "BRUSDT",
		},
	},
	"KP4R": {
		{
			ID:          7582,
			Slug:        "keep4r",
			Name:        "Keep4r",
			BinanceUSDT: "KP4RUSDT",
		},
	},
	"XDN": {
		{
			ID:          405,
			Slug:        "digitalnote",
			Name:        "DigitalNote",
			BinanceUSDT: "XDNUSDT",
		},
	},
	"DFND": {
		{
			ID:          9597,
			Slug:        "dfund",
			Name:        "dFund",
			BinanceUSDT: "DFNDUSDT",
		},
	},
	"SIMPLE": {
		{
			ID:          5319,
			Slug:        "simplechain",
			Name:        "SimpleChain",
			BinanceUSDT: "SIMPLEUSDT",
		},
	},
	"FUD": {
		{
			ID:          7335,
			Slug:        "fudfinance",
			Name:        "FUD.finance",
			BinanceUSDT: "FUDUSDT",
		},
	},
	"CRO": {
		{
			ID:          3635,
			Slug:        "crypto-com-coin",
			Name:        "Crypto.com Coin",
			BinanceUSDT: "CROUSDT",
		},
	},
	"YAE": {
		{
			ID:          9149,
			Slug:        "cryptonovae",
			Name:        "Cryptonovae",
			BinanceUSDT: "YAEUSDT",
		},
	},
	"DOOR": {
		{
			ID:          10873,
			Slug:        "door",
			Name:        "DOOR",
			BinanceUSDT: "DOORUSDT",
		},
	},
	"CFX": {
		{
			ID:          7334,
			Slug:        "conflux-network",
			Name:        "Conflux Network",
			BinanceUSDT: "CFXUSDT",
		},
	},
	"TXL": {
		{
			ID:          7024,
			Slug:        "tixl-new",
			Name:        "Tixl",
			BinanceUSDT: "TXLUSDT",
		},
	},
	"MYOBU": {
		{
			ID:          10578,
			Slug:        "myobu",
			Name:        "Myōbu",
			BinanceUSDT: "MYOBUUSDT",
		},
	},
	"EGX": {
		{
			ID:          3403,
			Slug:        "eaglex",
			Name:        "EagleX",
			BinanceUSDT: "EGXUSDT",
		},
	},
	"TRTT": {
		{
			ID:          2865,
			Slug:        "trittium",
			Name:        "Trittium",
			BinanceUSDT: "TRTTUSDT",
		},
	},
	"KOMET": {
		{
			ID:          7760,
			Slug:        "komet",
			Name:        "Komet",
			BinanceUSDT: "KOMETUSDT",
		},
	},
	"DEQ": {
		{
			ID:          8224,
			Slug:        "dequant",
			Name:        "Dequant",
			BinanceUSDT: "DEQUSDT",
		},
	},
	"SIM": {
		{
			ID:          9948,
			Slug:        "simba-inu",
			Name:        "Simba Inu",
			BinanceUSDT: "SIMUSDT",
		},
	},
	"BIP": {
		{
			ID:          4957,
			Slug:        "minter-network",
			Name:        "Minter Network",
			BinanceUSDT: "BIPUSDT",
		},
	},
	"NEU": {
		{
			ID:          2318,
			Slug:        "neumark",
			Name:        "Neumark",
			BinanceUSDT: "NEUUSDT",
		},
	},
	"DAC": {
		{
			ID:          3154,
			Slug:        "davinci-coin",
			Name:        "Davinci Coin",
			BinanceUSDT: "DACUSDT",
		},
	},
	"SAFEBTC": {
		{
			ID:          8993,
			Slug:        "safebtc",
			Name:        "SafeBTC",
			BinanceUSDT: "SAFEBTCUSDT",
		},
	},
	"ALINK": {
		{
			ID:          5751,
			Slug:        "aave-link",
			Name:        "Aave LINK",
			BinanceUSDT: "ALINKUSDT",
		},
	},
	"EZY": {
		{
			ID:          5521,
			Slug:        "ezystayz",
			Name:        "EzyStayz",
			BinanceUSDT: "EZYUSDT",
		},
	},
	"XCUR": {
		{
			ID:          7942,
			Slug:        "curate",
			Name:        "Curate",
			BinanceUSDT: "XCURUSDT",
		},
	},
	"SIERRA": {
		{
			ID:          4630,
			Slug:        "sierracoin",
			Name:        "Sierracoin",
			BinanceUSDT: "SIERRAUSDT",
		},
	},
	"ΤBTC": {
		{
			ID:          9114,
			Slug:        "t-bitcoin",
			Name:        "τBitcoin",
			BinanceUSDT: "ΤBTCUSDT",
		},
	},
	"BCZ": {
		{
			ID:          4008,
			Slug:        "bitcoin-cz",
			Name:        "Bitcoin CZ",
			BinanceUSDT: "BCZUSDT",
		},
	},
	"COLL": {
		{
			ID:          8999,
			Slug:        "collateral-pay",
			Name:        "Collateral Pay",
			BinanceUSDT: "COLLUSDT",
		},
	},
	"ZXC": {
		{
			ID:          2920,
			Slug:        "0xcert",
			Name:        "0xcert",
			BinanceUSDT: "ZXCUSDT",
		},
	},
	"NDAU": {
		{
			ID:          6524,
			Slug:        "ndau",
			Name:        "Ndau",
			BinanceUSDT: "NDAUUSDT",
		},
	},
	"ZCK": {
		{
			ID:          8687,
			Slug:        "polkazeck",
			Name:        "Polkazeck",
			BinanceUSDT: "ZCKUSDT",
		},
	},
	"XENO": {
		{
			ID:          4205,
			Slug:        "xenoverse",
			Name:        "Xenoverse",
			BinanceUSDT: "XENOUSDT",
		},
	},
	"EC": {
		{
			ID:   4695,
			Slug: "echoin",
			Name: "Echoin",
		},
		{
			ID:   8839,
			Slug: "eternal-cash",
			Name: "Eternal Cash",
		},
	},
	"ARES": {
		{
			ID:          8702,
			Slug:        "ares-protocol",
			Name:        "Ares Protocol",
			BinanceUSDT: "ARESUSDT",
		},
	},
	"MASQ": {
		{
			ID:          8376,
			Slug:        "masq",
			Name:        "MASQ",
			BinanceUSDT: "MASQUSDT",
		},
	},
	"XGC": {
		{
			ID:          10436,
			Slug:        "xiglute-coin",
			Name:        "Xiglute Coin",
			BinanceUSDT: "XGCUSDT",
		},
	},
	"EUM": {
		{
			ID:          3968,
			Slug:        "elitium",
			Name:        "Elitium",
			BinanceUSDT: "EUMUSDT",
		},
	},
	"CBR": {
		{
			ID:          5921,
			Slug:        "cybercoin",
			Name:        "Cybercoin",
			BinanceUSDT: "CBRUSDT",
		},
	},
	"MOJO": {
		{
			ID:   1212,
			Slug: "mojocoin",
			Name: "MojoCoin",
		},
		{
			ID:   9832,
			Slug: "moonjuice",
			Name: "MoonJuice",
		},
	},
	"GSPI": {
		{
			ID:          8737,
			Slug:        "gspi-governance",
			Name:        "GSPI Shopping.io Governance",
			BinanceUSDT: "GSPIUSDT",
		},
	},
	"DGD": {
		{
			ID:          1229,
			Slug:        "digixdao",
			Name:        "DigixDAO",
			BinanceUSDT: "DGDUSDT",
		},
	},
	"PUML": {
		{
			ID:          5842,
			Slug:        "puml-better-health",
			Name:        "PUML Better Health",
			BinanceUSDT: "PUMLUSDT",
		},
	},
	"AUDIO": {
		{
			ID:          7455,
			Slug:        "audius",
			Name:        "Audius",
			BinanceUSDT: "AUDIOUSDT",
		},
	},
	"FIDA": {
		{
			ID:          7978,
			Slug:        "bonfida",
			Name:        "Bonfida",
			BinanceUSDT: "FIDAUSDT",
		},
	},
	"OCC": {
		{
			ID:          9191,
			Slug:        "occamfi",
			Name:        "Occam.Fi",
			BinanceUSDT: "OCCUSDT",
		},
	},
	"LDN": {
		{
			ID:          7992,
			Slug:        "ludena-protocol",
			Name:        "Ludena Protocol",
			BinanceUSDT: "LDNUSDT",
		},
	},
	"SLNV2": {
		{
			ID:          8530,
			Slug:        "starlink",
			Name:        "StarLink",
			BinanceUSDT: "SLNV2USDT",
		},
	},
	"ITAM": {
		{
			ID:          6490,
			Slug:        "itam-games",
			Name:        "ITAM Games",
			BinanceUSDT: "ITAMUSDT",
		},
	},
	"yBAN": {
		{
			ID:          7489,
			Slug:        "bananodos",
			Name:        "BananoDOS",
			BinanceUSDT: "yBANUSDT",
		},
	},
	"BHIG": {
		{
			ID:          3820,
			Slug:        "buck-hath-coin",
			Name:        "BuckHathCoin",
			BinanceUSDT: "BHIGUSDT",
		},
	},
	"REQ": {
		{
			ID:          2071,
			Slug:        "request",
			Name:        "Request",
			BinanceUSDT: "REQUSDT",
		},
	},
	"KUS": {
		{
			ID:          10908,
			Slug:        "kuswap",
			Name:        "KuSwap",
			BinanceUSDT: "KUSUSDT",
		},
	},
	"YFPRO": {
		{
			ID:          7331,
			Slug:        "yfpro-finance",
			Name:        "YFPRO Finance",
			BinanceUSDT: "YFPROUSDT",
		},
	},
	"HTRE": {
		{
			ID:          7100,
			Slug:        "hodltree",
			Name:        "HodlTree",
			BinanceUSDT: "HTREUSDT",
		},
	},
	"ACXT": {
		{
			ID:          7854,
			Slug:        "acdx-exchange-token",
			Name:        "ACDX Exchange Governance Token",
			BinanceUSDT: "ACXTUSDT",
		},
	},
	"MCPC": {
		{
			ID:          4014,
			Slug:        "mobile-crypto-pay-coin",
			Name:        "Mobile Crypto Pay Coin",
			BinanceUSDT: "MCPCUSDT",
		},
	},
	"FCX": {
		{
			ID:          8222,
			Slug:        "fission-cash",
			Name:        "Fission Cash",
			BinanceUSDT: "FCXUSDT",
		},
	},
	"BBK": {
		{
			ID:          3051,
			Slug:        "bitblocks",
			Name:        "Bitblocks",
			BinanceUSDT: "BBKUSDT",
		},
	},
	"ESRC": {
		{
			ID:          6765,
			Slug:        "esr-coin",
			Name:        "ESR Coin",
			BinanceUSDT: "ESRCUSDT",
		},
	},
	"EXO": {
		{
			ID:   3708,
			Slug: "exosis",
			Name: "Exosis",
		},
		{
			ID:   9879,
			Slug: "exohood",
			Name: "Exohood",
		},
	},
	"QUAM": {
		{
			ID:          8864,
			Slug:        "quam-network",
			Name:        "Quam Network",
			BinanceUSDT: "QUAMUSDT",
		},
	},
	"MRX": {
		{
			ID:          1814,
			Slug:        "metrix-coin",
			Name:        "Metrix Coin",
			BinanceUSDT: "MRXUSDT",
		},
	},
	"SQR": {
		{
			ID:          3467,
			Slug:        "squorum",
			Name:        "Squorum",
			BinanceUSDT: "SQRUSDT",
		},
	},
	"WIFI": {
		{
			ID:          8140,
			Slug:        "wifi-coin",
			Name:        "Wifi Coin",
			BinanceUSDT: "WIFIUSDT",
		},
	},
	"HSN": {
		{
			ID:          4682,
			Slug:        "hyper-speed-network",
			Name:        "Hyper Speed Network",
			BinanceUSDT: "HSNUSDT",
		},
	},
	"ISLE": {
		{
			ID:          10080,
			Slug:        "island-coin",
			Name:        "Island Coin",
			BinanceUSDT: "ISLEUSDT",
		},
	},
	"FDZ": {
		{
			ID:          2626,
			Slug:        "friends",
			Name:        "Friendz",
			BinanceUSDT: "FDZUSDT",
		},
	},
	"LBY": {
		{
			ID:          8569,
			Slug:        "libonomy",
			Name:        "Libonomy",
			BinanceUSDT: "LBYUSDT",
		},
	},
	"XAMP": {
		{
			ID:          5998,
			Slug:        "antiample",
			Name:        "Antiample",
			BinanceUSDT: "XAMPUSDT",
		},
	},
	"DPET": {
		{
			ID:          9665,
			Slug:        "my-defi-pet",
			Name:        "My DeFi Pet",
			BinanceUSDT: "DPETUSDT",
		},
	},
	"ACT": {
		{
			ID:          1918,
			Slug:        "achain",
			Name:        "Achain",
			BinanceUSDT: "ACTUSDT",
		},
	},
	"NFTFY": {
		{
			ID:          9622,
			Slug:        "nftfy",
			Name:        "Nftfy",
			BinanceUSDT: "NFTFYUSDT",
		},
	},
	"BTRS": {
		{
			ID:          4257,
			Slug:        "bitball-treasure",
			Name:        "Bitball Treasure",
			BinanceUSDT: "BTRSUSDT",
		},
	},
	"DSL": {
		{
			ID:          8739,
			Slug:        "deadsoul",
			Name:        "DeadSoul",
			BinanceUSDT: "DSLUSDT",
		},
	},
	"SCP": {
		{
			ID:          4074,
			Slug:        "scprime",
			Name:        "ScPrime",
			BinanceUSDT: "SCPUSDT",
		},
	},
	"PIVX": {
		{
			ID:          1169,
			Slug:        "pivx",
			Name:        "PIVX",
			BinanceUSDT: "PIVXUSDT",
		},
	},
	"vBNB": {
		{
			ID:          7961,
			Slug:        "venus-bnb",
			Name:        "Venus BNB",
			BinanceUSDT: "vBNBUSDT",
		},
	},
	"DRS": {
		{
			ID:          5796,
			Slug:        "doctors-coin",
			Name:        "Doctors Coin",
			BinanceUSDT: "DRSUSDT",
		},
	},
	"AMON": {
		{
			ID:          4712,
			Slug:        "amond",
			Name:        "AmonD",
			BinanceUSDT: "AMONUSDT",
		},
	},
	"AIB": {
		{
			ID:          911,
			Slug:        "advanced-internet-blocks",
			Name:        "Advanced Internet Blocks",
			BinanceUSDT: "AIBUSDT",
		},
	},
	"GRG": {
		{
			ID:          4927,
			Slug:        "rigoblock",
			Name:        "RigoBlock",
			BinanceUSDT: "GRGUSDT",
		},
	},
	"NIF": {
		{
			ID:          7879,
			Slug:        "unifty",
			Name:        "Unifty",
			BinanceUSDT: "NIFUSDT",
		},
	},
	"BCUBE": {
		{
			ID:          9553,
			Slug:        "b-cube-ai",
			Name:        "B-cube.ai",
			BinanceUSDT: "BCUBEUSDT",
		},
	},
	"OMG": {
		{
			ID:          1808,
			Slug:        "omg",
			Name:        "OMG Network",
			BinanceUSDT: "OMGUSDT",
		},
	},
	"CBP": {
		{
			ID:          5781,
			Slug:        "cashbackpro",
			Name:        "CashBackPro",
			BinanceUSDT: "CBPUSDT",
		},
	},
	"DEFI": {
		{
			ID:          4276,
			Slug:        "defi",
			Name:        "Defi",
			BinanceUSDT: "DEFIUSDT",
		},
	},
	"MPAY": {
		{
			ID:          4001,
			Slug:        "menapay",
			Name:        "MenaPay",
			BinanceUSDT: "MPAYUSDT",
		},
	},
	"STR": {
		{
			ID:   3232,
			Slug: "staker",
			Name: "Staker",
		},
		{
			ID:   8909,
			Slug: "stater",
			Name: "Stater",
		},
	},
	"BPLC": {
		{
			ID:          6113,
			Slug:        "blackpearl-chain",
			Name:        "BlackPearl Token",
			BinanceUSDT: "BPLCUSDT",
		},
	},
	"NST": {
		{
			ID:   10114,
			Slug: "nft-starter",
			Name: "NFT Starter",
		},
		{
			ID:   4975,
			Slug: "newsolution",
			Name: "Newsolution",
		},
	},
	"BRIGADEIRO": {
		{
			ID:          10654,
			Slug:        "brigadeiro",
			Name:        "Brigadeiro.Finance",
			BinanceUSDT: "BRIGADEIROUSDT",
		},
	},
	"WANUSDT": {
		{
			ID:          8657,
			Slug:        "wanusdt",
			Name:        "wanUSDT",
			BinanceUSDT: "WANUSDTUSDT",
		},
	},
	"AMKR": {
		{
			ID:          5753,
			Slug:        "aave-mkr",
			Name:        "Aave MKR",
			BinanceUSDT: "AMKRUSDT",
		},
	},
	"BIXB": {
		{
			ID:          9330,
			Slug:        "bixbcoin",
			Name:        "BIXBCOIN",
			BinanceUSDT: "BIXBUSDT",
		},
	},
	"DEEZNUTS": {
		{
			ID:          9873,
			Slug:        "deez-nuts",
			Name:        "Deez Nuts",
			BinanceUSDT: "DEEZNUTSUSDT",
		},
	},
	"XGT": {
		{
			ID:          8607,
			Slug:        "xion-finance",
			Name:        "Xion Finance",
			BinanceUSDT: "XGTUSDT",
		},
	},
	"IDOGE": {
		{
			ID:          10446,
			Slug:        "influencer-doge",
			Name:        "Influencer Doge",
			BinanceUSDT: "IDOGEUSDT",
		},
	},
	"EQUAD": {
		{
			ID:          3625,
			Slug:        "quadrantprotocol",
			Name:        "QuadrantProtocol",
			BinanceUSDT: "EQUADUSDT",
		},
	},
	"SKY": {
		{
			ID:          1619,
			Slug:        "skycoin",
			Name:        "Skycoin",
			BinanceUSDT: "SKYUSDT",
		},
	},
	"ORBYT": {
		{
			ID:          6559,
			Slug:        "orbyt-token",
			Name:        "ORBYT Token",
			BinanceUSDT: "ORBYTUSDT",
		},
	},
	"SHLD": {
		{
			ID:          8450,
			Slug:        "shield-finance",
			Name:        "Shield Finance",
			BinanceUSDT: "SHLDUSDT",
		},
	},
	"DTA": {
		{
			ID:          2446,
			Slug:        "data",
			Name:        "DATA",
			BinanceUSDT: "DTAUSDT",
		},
	},
	"HOTCROSS": {
		{
			ID:          9867,
			Slug:        "hot-cross",
			Name:        "Hot Cross",
			BinanceUSDT: "HOTCROSSUSDT",
		},
	},
	"OKT": {
		{
			ID:          8267,
			Slug:        "okexchain",
			Name:        "OKExChain",
			BinanceUSDT: "OKTUSDT",
		},
	},
	"ERK": {
		{
			ID:          4966,
			Slug:        "eureka-coin",
			Name:        "Eureka Coin",
			BinanceUSDT: "ERKUSDT",
		},
	},
	"YFIIG": {
		{
			ID:          7193,
			Slug:        "yfii-gold",
			Name:        "YFII Gold",
			BinanceUSDT: "YFIIGUSDT",
		},
	},
	"RUP": {
		{
			ID:          1799,
			Slug:        "rupee",
			Name:        "Rupee",
			BinanceUSDT: "RUPUSDT",
		},
	},
	"IC": {
		{
			ID:          2395,
			Slug:        "ignition",
			Name:        "Ignition",
			BinanceUSDT: "ICUSDT",
		},
	},
	"CREP": {
		{
			ID:          5746,
			Slug:        "compound-augur",
			Name:        "Compound Augur",
			BinanceUSDT: "CREPUSDT",
		},
	},
	"SURF": {
		{
			ID:          7547,
			Slug:        "surf",
			Name:        "SURF Finance",
			BinanceUSDT: "SURFUSDT",
		},
	},
	"THE": {
		{
			ID:          5078,
			Slug:        "thenode",
			Name:        "THENODE",
			BinanceUSDT: "THEUSDT",
		},
	},
	"COSM": {
		{
			ID:          2955,
			Slug:        "cosmo-coin",
			Name:        "Cosmo Coin",
			BinanceUSDT: "COSMUSDT",
		},
	},
	"DINU": {
		{
			ID:          10021,
			Slug:        "dogey-inu",
			Name:        "Dogey-Inu",
			BinanceUSDT: "DINUUSDT",
		},
	},
	"UMX": {
		{
			ID:          8229,
			Slug:        "unimex-network",
			Name:        "UniMex Network",
			BinanceUSDT: "UMXUSDT",
		},
	},
	"AQUARI": {
		{
			ID:          9921,
			Slug:        "aquari",
			Name:        "Aquari",
			BinanceUSDT: "AQUARIUSDT",
		},
	},
	"SOLO": {
		{
			ID:          5279,
			Slug:        "sologenic",
			Name:        "Sologenic",
			BinanceUSDT: "SOLOUSDT",
		},
	},
	"SHE": {
		{
			ID:          3252,
			Slug:        "shinechain",
			Name:        "ShineChain",
			BinanceUSDT: "SHEUSDT",
		},
	},
	"COOK": {
		{
			ID:          8997,
			Slug:        "cook-protocol",
			Name:        "Cook Protocol",
			BinanceUSDT: "COOKUSDT",
		},
	},
	"WFTM": {
		{
			ID:          10240,
			Slug:        "wrapped-fantom",
			Name:        "Wrapped Fantom",
			BinanceUSDT: "WFTMUSDT",
		},
	},
	"BLOWF": {
		{
			ID:          8984,
			Slug:        "blowfish",
			Name:        "BlowFish",
			BinanceUSDT: "BLOWFUSDT",
		},
	},
	"SOC": {
		{
			ID:   2473,
			Slug: "all-sports",
			Name: "All Sports",
		},
		{
			ID:   5126,
			Slug: "soda-coin",
			Name: "Soda Coin",
		},
	},
	"EVT": {
		{
			ID:          4238,
			Slug:        "everitoken",
			Name:        "EveriToken",
			BinanceUSDT: "EVTUSDT",
		},
	},
	"HUSH": {
		{
			ID:          1466,
			Slug:        "hush",
			Name:        "Hush",
			BinanceUSDT: "HUSHUSDT",
		},
	},
	"BRT": {
		{
			ID:          10602,
			Slug:        "base-reward-token",
			Name:        "Base Reward Token",
			BinanceUSDT: "BRTUSDT",
		},
	},
	"HUB": {
		{
			ID:          7986,
			Slug:        "hub-human-trust-protocol",
			Name:        "Hub - Human Trust Protocol",
			BinanceUSDT: "HUBUSDT",
		},
	},
	"VIB": {
		{
			ID:          2019,
			Slug:        "viberate",
			Name:        "Viberate",
			BinanceUSDT: "VIBUSDT",
		},
	},
	"EXE": {
		{
			ID:          5620,
			Slug:        "8x8-protocol",
			Name:        "8X8 PROTOCOL",
			BinanceUSDT: "EXEUSDT",
		},
	},
	"SWINGBY": {
		{
			ID:          5922,
			Slug:        "swingby",
			Name:        "Swingby",
			BinanceUSDT: "SWINGBYUSDT",
		},
	},
	"AABC": {
		{
			ID:          9541,
			Slug:        "aabc-token",
			Name:        "AABC Token",
			BinanceUSDT: "AABCUSDT",
		},
	},
	"LGCY": {
		{
			ID:          6665,
			Slug:        "lgcy-network",
			Name:        "LGCY Network",
			BinanceUSDT: "LGCYUSDT",
		},
	},
	"TAXI": {
		{
			ID:          8070,
			Slug:        "taxi",
			Name:        "Taxi",
			BinanceUSDT: "TAXIUSDT",
		},
	},
	"VN": {
		{
			ID:          5828,
			Slug:        "vn-token",
			Name:        "VN Token",
			BinanceUSDT: "VNUSDT",
		},
	},
	"FIN": {
		{
			ID:          7225,
			Slug:        "definer",
			Name:        "DeFiner",
			BinanceUSDT: "FINUSDT",
		},
	},
	"GHOST": {
		{
			ID:   5471,
			Slug: "ghost",
			Name: "Ghost",
		},
		{
			ID:   5475,
			Slug: "ghostprism",
			Name: "GHOSTPRISM",
		},
	},
	"BSDS": {
		{
			ID:          8237,
			Slug:        "basis-dollar-share",
			Name:        "Basis Dollar Share",
			BinanceUSDT: "BSDSUSDT",
		},
	},
	"RDD": {
		{
			ID:          118,
			Slug:        "redd",
			Name:        "ReddCoin",
			BinanceUSDT: "RDDUSDT",
		},
	},
	"CHAIN": {
		{
			ID:          6744,
			Slug:        "chain-games",
			Name:        "Chain Games",
			BinanceUSDT: "CHAINUSDT",
		},
	},
	"TNDR": {
		{
			ID:          8951,
			Slug:        "thunderswap",
			Name:        "ThunderSwap",
			BinanceUSDT: "TNDRUSDT",
		},
	},
	"TMCN": {
		{
			ID:          7738,
			Slug:        "timecoinprotocol",
			Name:        "TimeCoinProtocol",
			BinanceUSDT: "TMCNUSDT",
		},
	},
	"DMME": {
		{
			ID:          4819,
			Slug:        "dmme",
			Name:        "DMme",
			BinanceUSDT: "DMMEUSDT",
		},
	},
	"TIC": {
		{
			ID:          3131,
			Slug:        "thingschain",
			Name:        "Thingschain",
			BinanceUSDT: "TICUSDT",
		},
	},
	"WINDY": {
		{
			ID:          9560,
			Slug:        "windswap",
			Name:        "WindSwap",
			BinanceUSDT: "WINDYUSDT",
		},
	},
	"HZN": {
		{
			ID:          9237,
			Slug:        "horizon-protocol",
			Name:        "Horizon Protocol",
			BinanceUSDT: "HZNUSDT",
		},
	},
	"DETS": {
		{
			ID:          6722,
			Slug:        "dextrust",
			Name:        "Dextrust",
			BinanceUSDT: "DETSUSDT",
		},
	},
	"OLE": {
		{
			ID:          10343,
			Slug:        "olecoin",
			Name:        "Olecoin",
			BinanceUSDT: "OLEUSDT",
		},
	},
	"TWIN": {
		{
			ID:          9253,
			Slug:        "twinci",
			Name:        "Twinci",
			BinanceUSDT: "TWINUSDT",
		},
	},
	"EDGT": {
		{
			ID:          9466,
			Slug:        "edgecoin",
			Name:        "Edgecoin",
			BinanceUSDT: "EDGTUSDT",
		},
	},
	"CAPP": {
		{
			ID:          2248,
			Slug:        "cappasity",
			Name:        "Cappasity",
			BinanceUSDT: "CAPPUSDT",
		},
	},
	"PGN": {
		{
			ID:          2988,
			Slug:        "pigeoncoin",
			Name:        "Pigeoncoin",
			BinanceUSDT: "PGNUSDT",
		},
	},
	"XFUND": {
		{
			ID:          8339,
			Slug:        "xfund",
			Name:        "xFund",
			BinanceUSDT: "XFUNDUSDT",
		},
	},
	"NOVA": {
		{
			ID:          4835,
			Slug:        "nova",
			Name:        "NOVA",
			BinanceUSDT: "NOVAUSDT",
		},
	},
	"POOLZ": {
		{
			ID:          8271,
			Slug:        "poolz-finance",
			Name:        "Poolz Finance",
			BinanceUSDT: "POOLZUSDT",
		},
	},
	"BIX": {
		{
			ID:          2307,
			Slug:        "bibox-token",
			Name:        "Bibox Token",
			BinanceUSDT: "BIXUSDT",
		},
	},
	"SLICE": {
		{
			ID:          8637,
			Slug:        "tranche-finance",
			Name:        "Tranche Finance",
			BinanceUSDT: "SLICEUSDT",
		},
	},
	"JPYC": {
		{
			ID:          9045,
			Slug:        "jpycoin",
			Name:        "JPYC",
			BinanceUSDT: "JPYCUSDT",
		},
	},
	"ADEL": {
		{
			ID:          6852,
			Slug:        "akropolis-delphi",
			Name:        "Akropolis Delphi",
			BinanceUSDT: "ADELUSDT",
		},
	},
	"MHC": {
		{
			ID:          3756,
			Slug:        "metahash",
			Name:        "#MetaHash",
			BinanceUSDT: "MHCUSDT",
		},
	},
	"GFX": {
		{
			ID:          9018,
			Slug:        "gamyfi-platform",
			Name:        "GamyFi Platform",
			BinanceUSDT: "GFXUSDT",
		},
	},
	"RATING": {
		{
			ID:          3194,
			Slug:        "dprating",
			Name:        "DPRating",
			BinanceUSDT: "RATINGUSDT",
		},
	},
	"PM": {
		{
			ID:          10302,
			Slug:        "pomskey",
			Name:        "Pomskey",
			BinanceUSDT: "PMUSDT",
		},
	},
	"POWER": {
		{
			ID:          5660,
			Slug:        "unipower",
			Name:        "UniPower",
			BinanceUSDT: "POWERUSDT",
		},
	},
	"XFUEL": {
		{
			ID:          6602,
			Slug:        "xfuel",
			Name:        "XFUEL",
			BinanceUSDT: "XFUELUSDT",
		},
	},
	"1X2": {
		{
			ID:          3767,
			Slug:        "1x2-coin",
			Name:        "1X2 COIN",
			BinanceUSDT: "1X2USDT",
		},
	},
	"PAI": {
		{
			ID:          2900,
			Slug:        "project-pai",
			Name:        "Project Pai",
			BinanceUSDT: "PAIUSDT",
		},
	},
	"TOSHI": {
		{
			ID:          8744,
			Slug:        "toshimon",
			Name:        "Toshimon",
			BinanceUSDT: "TOSHIUSDT",
		},
	},
	"HT": {
		{
			ID:          2502,
			Slug:        "huobi-token",
			Name:        "Huobi Token",
			BinanceUSDT: "HTUSDT",
		},
	},
	"DOG": {
		{
			ID:          8818,
			Slug:        "underdog",
			Name:        "UnderDog",
			BinanceUSDT: "DOGUSDT",
		},
	},
	"ECC": {
		{
			ID:          212,
			Slug:        "eccoin",
			Name:        "ECC",
			BinanceUSDT: "ECCUSDT",
		},
	},
	"ORBS": {
		{
			ID:          3835,
			Slug:        "orbs",
			Name:        "Orbs",
			BinanceUSDT: "ORBSUSDT",
		},
	},
	"NZO": {
		{
			ID:          5102,
			Slug:        "enzo",
			Name:        "Enzo",
			BinanceUSDT: "NZOUSDT",
		},
	},
	"MET": {
		{
			ID:          2873,
			Slug:        "metronome",
			Name:        "Metronome",
			BinanceUSDT: "METUSDT",
		},
	},
	"XFYI": {
		{
			ID:          7323,
			Slug:        "xcredit",
			Name:        "XCredit",
			BinanceUSDT: "XFYIUSDT",
		},
	},
	"CIX100": {
		{
			ID:          4067,
			Slug:        "cryptoindex-com-100",
			Name:        "Cryptoindex.com 100",
			BinanceUSDT: "CIX100USDT",
		},
	},
	"TCORE": {
		{
			ID:          8082,
			Slug:        "tornado",
			Name:        "Tornado",
			BinanceUSDT: "TCOREUSDT",
		},
	},
	"IFT": {
		{
			ID:   1888,
			Slug: "investfeed",
			Name: "InvestFeed",
		},
		{
			ID:   8949,
			Slug: "iftoken",
			Name: "IFToken",
		},
	},
	"START": {
		{
			ID:   389,
			Slug: "startcoin",
			Name: "Startcoin",
		},
		{
			ID:   8662,
			Slug: "bscstarter",
			Name: "Starter",
		},
	},
	"LABO": {
		{
			ID:          9312,
			Slug:        "the-lab-finance",
			Name:        "The Lab Finance",
			BinanceUSDT: "LABOUSDT",
		},
	},
	"MER": {
		{
			ID:   1590,
			Slug: "mercury",
			Name: "Mercury",
		},
		{
			ID:   9549,
			Slug: "mercurial-finance",
			Name: "Mercurial Finance",
		},
	},
	"JUS": {
		{
			ID:          6483,
			Slug:        "just-network",
			Name:        "JUST NETWORK",
			BinanceUSDT: "JUSUSDT",
		},
	},
	"SOTA": {
		{
			ID:          8576,
			Slug:        "sota-finance",
			Name:        "SOTA Finance",
			BinanceUSDT: "SOTAUSDT",
		},
	},
	"CFT": {
		{
			ID:          6029,
			Slug:        "coinbene-future-token",
			Name:        "CoinBene Future Token",
			BinanceUSDT: "CFTUSDT",
		},
	},
	"GSC": {
		{
			ID:          2737,
			Slug:        "global-social-chain",
			Name:        "Global Social Chain",
			BinanceUSDT: "GSCUSDT",
		},
	},
	"MOOCHII": {
		{
			ID:          10077,
			Slug:        "moochii",
			Name:        "MOOCHII",
			BinanceUSDT: "MOOCHIIUSDT",
		},
	},
	"DICK": {
		{
			ID:          10027,
			Slug:        "dick",
			Name:        "Dick",
			BinanceUSDT: "DICKUSDT",
		},
	},
	"DOCK": {
		{
			ID:          2675,
			Slug:        "dock",
			Name:        "Dock",
			BinanceUSDT: "DOCKUSDT",
		},
	},
	"HNT": {
		{
			ID:          5665,
			Slug:        "helium",
			Name:        "Helium",
			BinanceUSDT: "HNTUSDT",
		},
	},
	"IBNB": {
		{
			ID:          10878,
			Slug:        "i-bnb",
			Name:        "iBNB",
			BinanceUSDT: "IBNBUSDT",
		},
	},
	"PDEX": {
		{
			ID:          9017,
			Slug:        "polkadex",
			Name:        "Polkadex",
			BinanceUSDT: "PDEXUSDT",
		},
	},
	"HBTC": {
		{
			ID:          6941,
			Slug:        "huobi-btc",
			Name:        "Huobi BTC",
			BinanceUSDT: "HBTCUSDT",
		},
	},
	"STEP": {
		{
			ID:          9443,
			Slug:        "step-finance",
			Name:        "Step Finance",
			BinanceUSDT: "STEPUSDT",
		},
	},
	"SWING": {
		{
			ID:          1085,
			Slug:        "swing",
			Name:        "Swing",
			BinanceUSDT: "SWINGUSDT",
		},
	},
	"JVZ": {
		{
			ID:          7353,
			Slug:        "jiviz",
			Name:        "Jiviz",
			BinanceUSDT: "JVZUSDT",
		},
	},
	"HTML": {
		{
			ID:          2315,
			Slug:        "html-coin",
			Name:        "HTMLCOIN",
			BinanceUSDT: "HTMLUSDT",
		},
	},
	"JS": {
		{
			ID:          2100,
			Slug:        "javascript-token",
			Name:        "JavaScript Token",
			BinanceUSDT: "JSUSDT",
		},
	},
	"TM2": {
		{
			ID:          2850,
			Slug:        "traxia",
			Name:        "TRAXIA",
			BinanceUSDT: "TM2USDT",
		},
	},
	"SHPP": {
		{
			ID:          9316,
			Slug:        "shipit-pro",
			Name:        "Shipit pro",
			BinanceUSDT: "SHPPUSDT",
		},
	},
	"BSOCIAL": {
		{
			ID:          10102,
			Slug:        "banksocial",
			Name:        "BankSocial",
			BinanceUSDT: "BSOCIALUSDT",
		},
	},
	"AET": {
		{
			ID:          4920,
			Slug:        "aerotoken",
			Name:        "Aerotoken",
			BinanceUSDT: "AETUSDT",
		},
	},
	"SPOK": {
		{
			ID:          5312,
			Slug:        "spockchain-network",
			Name:        "Spockchain Network",
			BinanceUSDT: "SPOKUSDT",
		},
	},
	"RichDoge": {
		{
			ID:          10989,
			Slug:        "rich-doge-coin",
			Name:        "Rich Doge Coin",
			BinanceUSDT: "RichDogeUSDT",
		},
	},
	"IIC": {
		{
			ID:          2741,
			Slug:        "intelligent-investment-chain",
			Name:        "Intelligent Investment Chain",
			BinanceUSDT: "IICUSDT",
		},
	},
	"NCT": {
		{
			ID:   2630,
			Slug: "polyswarm",
			Name: "PolySwarm",
		},
		{
			ID:   8367,
			Slug: "name-change-token",
			Name: "Name Change Token",
		},
	},
	"COS": {
		{
			ID:          4036,
			Slug:        "contentos",
			Name:        "Contentos",
			BinanceUSDT: "COSUSDT",
		},
	},
	"PSHP": {
		{
			ID:          7407,
			Slug:        "payship",
			Name:        "Payship",
			BinanceUSDT: "PSHPUSDT",
		},
	},
	"KLT": {
		{
			ID:          9426,
			Slug:        "klend",
			Name:        "KLend",
			BinanceUSDT: "KLTUSDT",
		},
	},
	"SCSX": {
		{
			ID:          4487,
			Slug:        "secure-cash",
			Name:        "Secure Cash",
			BinanceUSDT: "SCSXUSDT",
		},
	},
	"FXT": {
		{
			ID:   2723,
			Slug: "fuzex",
			Name: "FuzeX",
		},
		{
			ID:   8854,
			Slug: "fxt-token",
			Name: "FXT Token",
		},
	},
	"NDX": {
		{
			ID:          8260,
			Slug:        "indexed-finance",
			Name:        "Indexed Finance",
			BinanceUSDT: "NDXUSDT",
		},
	},
	"ABX": {
		{
			ID:          3179,
			Slug:        "arbidex",
			Name:        "Arbidex",
			BinanceUSDT: "ABXUSDT",
		},
	},
	"ICA": {
		{
			ID:          9415,
			Slug:        "icarus-finance",
			Name:        "Icarus Finance",
			BinanceUSDT: "ICAUSDT",
		},
	},
	"CATZ": {
		{
			ID:          9692,
			Slug:        "catzcoin",
			Name:        "CatzCoin",
			BinanceUSDT: "CATZUSDT",
		},
	},
	"LVX": {
		{
			ID:          5094,
			Slug:        "level01",
			Name:        "Level01",
			BinanceUSDT: "LVXUSDT",
		},
	},
	"COUSINDOGE": {
		{
			ID:          10780,
			Slug:        "cousin-doge-coin",
			Name:        "COUSIN DOGE COIN",
			BinanceUSDT: "COUSINDOGEUSDT",
		},
	},
	"ECOM": {
		{
			ID:          3081,
			Slug:        "omnitude",
			Name:        "Omnitude",
			BinanceUSDT: "ECOMUSDT",
		},
	},
	"EMT": {
		{
			ID:          3993,
			Slug:        "emanate",
			Name:        "Emanate",
			BinanceUSDT: "EMTUSDT",
		},
	},
	"ARCONA": {
		{
			ID:          6218,
			Slug:        "arcona",
			Name:        "Arcona",
			BinanceUSDT: "ARCONAUSDT",
		},
	},
	"SIX": {
		{
			ID:          3327,
			Slug:        "six",
			Name:        "SIX",
			BinanceUSDT: "SIXUSDT",
		},
	},
	"EQO": {
		{
			ID:          9593,
			Slug:        "equos-io",
			Name:        "EQO",
			BinanceUSDT: "EQOUSDT",
		},
	},
	"BBR": {
		{
			ID:          406,
			Slug:        "boolberry",
			Name:        "Boolberry",
			BinanceUSDT: "BBRUSDT",
		},
	},
	"NUKO": {
		{
			ID:          2257,
			Slug:        "nekonium",
			Name:        "Nekonium",
			BinanceUSDT: "NUKOUSDT",
		},
	},
	"CNTR": {
		{
			ID:          7349,
			Slug:        "centaur",
			Name:        "Centaur",
			BinanceUSDT: "CNTRUSDT",
		},
	},
	"SEFI": {
		{
			ID:          9123,
			Slug:        "sefi",
			Name:        "SEFI",
			BinanceUSDT: "SEFIUSDT",
		},
	},
	"NSBT": {
		{
			ID:          7320,
			Slug:        "neutrino-system-base-token",
			Name:        "Neutrino Token",
			BinanceUSDT: "NSBTUSDT",
		},
	},
	"IMPL": {
		{
			ID:          3665,
			Slug:        "impleum",
			Name:        "Impleum",
			BinanceUSDT: "IMPLUSDT",
		},
	},
	"XNV": {
		{
			ID:          3596,
			Slug:        "nerva",
			Name:        "Nerva",
			BinanceUSDT: "XNVUSDT",
		},
	},
	"UCAP": {
		{
			ID:          7819,
			Slug:        "unicap-finance",
			Name:        "Unicap.finance",
			BinanceUSDT: "UCAPUSDT",
		},
	},
	"KIP": {
		{
			ID:          6489,
			Slug:        "khipu-token",
			Name:        "Khipu Token",
			BinanceUSDT: "KIPUSDT",
		},
	},
	"HP": {
		{
			ID:   5156,
			Slug: "heartbout-pay",
			Name: "HeartBout Pay",
		},
		{
			ID:   6033,
			Slug: "healing-plus",
			Name: "Healing Plus",
		},
	},
	"RKN": {
		{
			ID:          5072,
			Slug:        "rakon",
			Name:        "Rakon",
			BinanceUSDT: "RKNUSDT",
		},
	},
	"WAND": {
		{
			ID:          2269,
			Slug:        "wandx",
			Name:        "WandX",
			BinanceUSDT: "WANDUSDT",
		},
	},
	"HERB": {
		{
			ID:          3646,
			Slug:        "herbalist-token",
			Name:        "Herbalist Token",
			BinanceUSDT: "HERBUSDT",
		},
	},
	"ZUZ": {
		{
			ID:          9074,
			Slug:        "zuz-protocol",
			Name:        "ZUZ Protocol",
			BinanceUSDT: "ZUZUSDT",
		},
	},
	"TEAM": {
		{
			ID:          2729,
			Slug:        "tokenstars",
			Name:        "TEAM (TokenStars)",
			BinanceUSDT: "TEAMUSDT",
		},
	},
	"XLPG": {
		{
			ID:          5362,
			Slug:        "stellarpayglobal",
			Name:        "StellarPayGlobal",
			BinanceUSDT: "XLPGUSDT",
		},
	},
	"DAM": {
		{
			ID:          2260,
			Slug:        "datamine",
			Name:        "Datamine",
			BinanceUSDT: "DAMUSDT",
		},
	},
	"AAVEDOWN": {
		{
			ID:          7775,
			Slug:        "aave-down",
			Name:        "AAVEDOWN",
			BinanceUSDT: "AAVEDOWNUSDT",
		},
	},
	"VALUE": {
		{
			ID:          7404,
			Slug:        "value-defi",
			Name:        "Value Liquidity",
			BinanceUSDT: "VALUEUSDT",
		},
	},
	"SCR": {
		{
			ID:          3094,
			Slug:        "scorum-coins",
			Name:        "Scorum Coins",
			BinanceUSDT: "SCRUSDT",
		},
	},
	"ULTRA": {
		{
			ID:          10195,
			Slug:        "ultrasafe",
			Name:        "Ultrasafe",
			BinanceUSDT: "ULTRAUSDT",
		},
	},
	"ROSN": {
		{
			ID:          9783,
			Slug:        "roseon-finance",
			Name:        "Roseon Finance",
			BinanceUSDT: "ROSNUSDT",
		},
	},
	"IOP": {
		{
			ID:          1464,
			Slug:        "internet-of-people",
			Name:        "Internet of People",
			BinanceUSDT: "IOPUSDT",
		},
	},
	"AK12": {
		{
			ID:          5189,
			Slug:        "ak12",
			Name:        "AK12",
			BinanceUSDT: "AK12USDT",
		},
	},
	"XQC": {
		{
			ID:          5220,
			Slug:        "quras",
			Name:        "QURAS",
			BinanceUSDT: "XQCUSDT",
		},
	},
	"CPU": {
		{
			ID:   4178,
			Slug: "cpuchain",
			Name: "CPUchain",
		},
		{
			ID:   8295,
			Slug: "cpucoin",
			Name: "CPUcoin",
		},
	},
	"MBN": {
		{
			ID:   7059,
			Slug: "mobilian-coin",
			Name: "Mobilian Coin",
		},
		{
			ID:   4233,
			Slug: "membrana",
			Name: "Membrana",
		},
	},
	"POLYDOGE": {
		{
			ID:          10088,
			Slug:        "polydoge",
			Name:        "PolyDoge",
			BinanceUSDT: "POLYDOGEUSDT",
		},
	},
	"MRPH": {
		{
			ID:          2763,
			Slug:        "morpheus-network",
			Name:        "Morpheus.Network",
			BinanceUSDT: "MRPHUSDT",
		},
	},
	"YLD": {
		{
			ID:   8066,
			Slug: "yield-app",
			Name: "YIELD App",
		},
		{
			ID:   7991,
			Slug: "yield",
			Name: "Yield",
		},
	},
	"USDB": {
		{
			ID:          5219,
			Slug:        "usd-bancor",
			Name:        "USD Bancor",
			BinanceUSDT: "USDBUSDT",
		},
	},
	"KODA": {
		{
			ID:          9966,
			Slug:        "summit-koda-token",
			Name:        "Summit Koda Token",
			BinanceUSDT: "KODAUSDT",
		},
	},
	"SKYLARK": {
		{
			ID:          10039,
			Slug:        "skylark",
			Name:        "SKYLARK",
			BinanceUSDT: "SKYLARKUSDT",
		},
	},
	"SLW": {
		{
			ID:          10724,
			Slug:        "solarwind-token",
			Name:        "SolarWind Token",
			BinanceUSDT: "SLWUSDT",
		},
	},
	"WEST": {
		{
			ID:          5159,
			Slug:        "waves-enterprise",
			Name:        "Waves Enterprise",
			BinanceUSDT: "WESTUSDT",
		},
	},
	"PLA": {
		{
			ID:   4242,
			Slug: "planet",
			Name: "PLANET",
		},
		{
			ID:   3711,
			Slug: "plair",
			Name: "Plair",
		},
		{
			ID:   7461,
			Slug: "playdapp",
			Name: "PlayDapp",
		},
		{
			ID:   3731,
			Slug: "playchip",
			Name: "PlayChip",
		},
	},
	"SUPERBID": {
		{
			ID:          8836,
			Slug:        "superbid",
			Name:        "Superbid",
			BinanceUSDT: "SUPERBIDUSDT",
		},
	},
	"BAR": {
		{
			ID:          5225,
			Slug:        "fc-barcelona-fan-token",
			Name:        "FC Barcelona Fan Token",
			BinanceUSDT: "BARUSDT",
		},
	},
	"BTT": {
		{
			ID:          3718,
			Slug:        "bittorrent",
			Name:        "BitTorrent",
			BinanceUSDT: "BTTUSDT",
		},
	},
	"GNX": {
		{
			ID:          2291,
			Slug:        "genaro-network",
			Name:        "Genaro Network",
			BinanceUSDT: "GNXUSDT",
		},
	},
	"UNIQ": {
		{
			ID:          9287,
			Slug:        "uniqly",
			Name:        "Uniqly",
			BinanceUSDT: "UNIQUSDT",
		},
	},
	"1TRC": {
		{
			ID:          10955,
			Slug:        "1tronic-network",
			Name:        "1TRONIC Network",
			BinanceUSDT: "1TRCUSDT",
		},
	},
	"EULER": {
		{
			ID:          9368,
			Slug:        "euler-tools",
			Name:        "Euler Tools",
			BinanceUSDT: "EULERUSDT",
		},
	},
	"PI": {
		{
			ID:          2838,
			Slug:        "pchain",
			Name:        "Plian",
			BinanceUSDT: "PIUSDT",
		},
	},
	"FYD": {
		{
			ID:          4680,
			Slug:        "fydcoin",
			Name:        "FYDcoin",
			BinanceUSDT: "FYDUSDT",
		},
	},
	"CLOUT": {
		{
			ID:   10442,
			Slug: "bitclout",
			Name: "BitClout",
		},
		{
			ID:   5904,
			Slug: "blockclout",
			Name: "BLOCKCLOUT",
		},
		{
			ID:   9724,
			Slug: "clout",
			Name: "CLOUT",
		},
	},
	"SWASS": {
		{
			ID:          10194,
			Slug:        "swass-finance",
			Name:        "SWASS.FINANCE",
			BinanceUSDT: "SWASSUSDT",
		},
	},
	"BMH": {
		{
			ID:          2747,
			Slug:        "blockmesh",
			Name:        "BlockMesh",
			BinanceUSDT: "BMHUSDT",
		},
	},
	"PAMP": {
		{
			ID:          10323,
			Slug:        "pamp-cc",
			Name:        "PAMP.CC",
			BinanceUSDT: "PAMPUSDT",
		},
	},
	"UPP": {
		{
			ID:          2866,
			Slug:        "sentinel-protocol",
			Name:        "Sentinel Protocol",
			BinanceUSDT: "UPPUSDT",
		},
	},
	"vETH": {
		{
			ID:          7963,
			Slug:        "venus-eth",
			Name:        "Venus ETH",
			BinanceUSDT: "vETHUSDT",
		},
	},
	"HOGL": {
		{
			ID:          8944,
			Slug:        "hogl-finance",
			Name:        "HOGL finance",
			BinanceUSDT: "HOGLUSDT",
		},
	},
	"FNK": {
		{
			ID:          8014,
			Slug:        "fnk-wallet",
			Name:        "FNK wallet",
			BinanceUSDT: "FNKUSDT",
		},
	},
	"OAK": {
		{
			ID:          9152,
			Slug:        "octree-oak",
			Name:        "Octree Finance",
			BinanceUSDT: "OAKUSDT",
		},
	},
	"TEMCO": {
		{
			ID:          3722,
			Slug:        "temco",
			Name:        "TEMCO",
			BinanceUSDT: "TEMCOUSDT",
		},
	},
	"DLC": {
		{
			ID:          1395,
			Slug:        "dollarcoin",
			Name:        "Dollarcoin",
			BinanceUSDT: "DLCUSDT",
		},
	},
	"HYC": {
		{
			ID:          3147,
			Slug:        "hycon",
			Name:        "HYCON",
			BinanceUSDT: "HYCUSDT",
		},
	},
	"FTT": {
		{
			ID:          4195,
			Slug:        "ftx-token",
			Name:        "FTX Token",
			BinanceUSDT: "FTTUSDT",
		},
	},
	"XCASH": {
		{
			ID:          3334,
			Slug:        "x-cash",
			Name:        "X-CASH",
			BinanceUSDT: "XCASHUSDT",
		},
	},
	"PAT": {
		{
			ID:          2759,
			Slug:        "patron",
			Name:        "Patron",
			BinanceUSDT: "PATUSDT",
		},
	},
	"NIO": {
		{
			ID:          7888,
			Slug:        "nio-tokenized-stock-ftx",
			Name:        "Nio tokenized stock FTX",
			BinanceUSDT: "NIOUSDT",
		},
	},
	"BXT": {
		{
			ID:          4858,
			Slug:        "bitfxt-coin",
			Name:        "BITFXT COIN",
			BinanceUSDT: "BXTUSDT",
		},
	},
	"MFG": {
		{
			ID:          2658,
			Slug:        "smart-mfg",
			Name:        "Smart MFG",
			BinanceUSDT: "MFGUSDT",
		},
	},
	"CLEANOCEAN": {
		{
			ID:          10499,
			Slug:        "cleanocean",
			Name:        "CleanOcean",
			BinanceUSDT: "CLEANOCEANUSDT",
		},
	},
	"TKMN": {
		{
			ID:          8338,
			Slug:        "tokemon",
			Name:        "Tokemon",
			BinanceUSDT: "TKMNUSDT",
		},
	},
	"SRAT": {
		{
			ID:          10579,
			Slug:        "spacerat",
			Name:        "SpaceRat",
			BinanceUSDT: "SRATUSDT",
		},
	},
	"BTCP": {
		{
			ID:          2575,
			Slug:        "bitcoin-private",
			Name:        "Bitcoin Private",
			BinanceUSDT: "BTCPUSDT",
		},
	},
	"ENRG": {
		{
			ID:          322,
			Slug:        "energycoin",
			Name:        "Energycoin",
			BinanceUSDT: "ENRGUSDT",
		},
	},
	"DIN": {
		{
			ID:          3263,
			Slug:        "dinero",
			Name:        "Dinero",
			BinanceUSDT: "DINUSDT",
		},
	},
	"OM": {
		{
			ID:          6536,
			Slug:        "mantra-dao",
			Name:        "MANTRA DAO",
			BinanceUSDT: "OMUSDT",
		},
	},
	"RSV": {
		{
			ID:          6727,
			Slug:        "reserve",
			Name:        "Reserve",
			BinanceUSDT: "RSVUSDT",
		},
	},
	"FAIRLIFE": {
		{
			ID:          10546,
			Slug:        "fairlife",
			Name:        "FAIRLIFE",
			BinanceUSDT: "FAIRLIFEUSDT",
		},
	},
	"ZNN": {
		{
			ID:          4003,
			Slug:        "zenon",
			Name:        "Zenon",
			BinanceUSDT: "ZNNUSDT",
		},
	},
	"TAPE": {
		{
			ID:          10147,
			Slug:        "toolape",
			Name:        "ToolApe",
			BinanceUSDT: "TAPEUSDT",
		},
	},
	"SHOKK": {
		{
			ID:          9730,
			Slug:        "shikokuaido",
			Name:        "Shikokuaido",
			BinanceUSDT: "SHOKKUSDT",
		},
	},
	"NETZ": {
		{
			ID:          10997,
			Slug:        "netzcoin",
			Name:        "Netzcoin",
			BinanceUSDT: "NETZUSDT",
		},
	},
	"YEFIM": {
		{
			ID:          7400,
			Slug:        "yfi-management",
			Name:        "YFi Management",
			BinanceUSDT: "YEFIMUSDT",
		},
	},
	"LIME": {
		{
			ID:          10469,
			Slug:        "ime-lab",
			Name:        "iMe Lab",
			BinanceUSDT: "LIMEUSDT",
		},
	},
	"XMARK": {
		{
			ID:          8835,
			Slug:        "xmark",
			Name:        "xMARK",
			BinanceUSDT: "XMARKUSDT",
		},
	},
	"OPIUM": {
		{
			ID:          7230,
			Slug:        "opium",
			Name:        "Opium",
			BinanceUSDT: "OPIUMUSDT",
		},
	},
	"BANCA": {
		{
			ID:          2592,
			Slug:        "banca",
			Name:        "Banca",
			BinanceUSDT: "BANCAUSDT",
		},
	},
	"BID": {
		{
			ID:   9236,
			Slug: "topbidder",
			Name: "TopBidder",
		},
		{
			ID:   6857,
			Slug: "defi-bids",
			Name: "DeFi Bids",
		},
		{
			ID:   7200,
			Slug: "bidao",
			Name: "Bidao",
		},
		{
			ID:   5512,
			Slug: "blockidcoin",
			Name: "BLOCKIDCOIN",
		},
	},
	"50X": {
		{
			ID:          8868,
			Slug:        "50x-com",
			Name:        "50x.com",
			BinanceUSDT: "50XUSDT",
		},
	},
	"KUE": {
		{
			ID:          3369,
			Slug:        "kuende",
			Name:        "Kuende",
			BinanceUSDT: "KUEUSDT",
		},
	},
	"BSPAY": {
		{
			ID:          8956,
			Slug:        "brosispay",
			Name:        "Brosispay",
			BinanceUSDT: "BSPAYUSDT",
		},
	},
	"VXV": {
		{
			ID:          4441,
			Slug:        "vectorspace-ai",
			Name:        "Vectorspace AI",
			BinanceUSDT: "VXVUSDT",
		},
	},
	"GFUN": {
		{
			ID:          3776,
			Slug:        "goldfund",
			Name:        "GoldFund",
			BinanceUSDT: "GFUNUSDT",
		},
	},
	"MWBTC": {
		{
			ID:          8504,
			Slug:        "metawhale-btc",
			Name:        "MetaWhale BTC",
			BinanceUSDT: "MWBTCUSDT",
		},
	},
	"XIO": {
		{
			ID:          4997,
			Slug:        "blockzerolabs",
			Name:        "Blockzero Labs",
			BinanceUSDT: "XIOUSDT",
		},
	},
	"BHP": {
		{
			ID:          3020,
			Slug:        "bhp-coin",
			Name:        "BHPCoin",
			BinanceUSDT: "BHPUSDT",
		},
	},
	"EURU": {
		{
			ID:          6905,
			Slug:        "upper-euro",
			Name:        "Upper Euro",
			BinanceUSDT: "EURUUSDT",
		},
	},
	"BIFI": {
		{
			ID:   7311,
			Slug: "beefy-finance",
			Name: "Beefy.Finance",
		},
		{
			ID:   8132,
			Slug: "bifi",
			Name: "BiFi",
		},
		{
			ID:   2994,
			Slug: "bitcoin-file",
			Name: "Bitcoin File",
		},
	},
	"API3": {
		{
			ID:          7737,
			Slug:        "api3",
			Name:        "API3",
			BinanceUSDT: "API3USDT",
		},
	},
	"SDAO": {
		{
			ID:          9638,
			Slug:        "singularitydao",
			Name:        "SingularityDAO",
			BinanceUSDT: "SDAOUSDT",
		},
	},
	"BYN": {
		{
			ID:          9134,
			Slug:        "beyond-finance",
			Name:        "Beyond Finance",
			BinanceUSDT: "BYNUSDT",
		},
	},
	"WHT": {
		{
			ID:          8524,
			Slug:        "wrapped-huobi-token",
			Name:        "Wrapped Huobi Token",
			BinanceUSDT: "WHTUSDT",
		},
	},
	"TRC": {
		{
			ID:          4,
			Slug:        "terracoin",
			Name:        "Terracoin",
			BinanceUSDT: "TRCUSDT",
		},
	},
	"LTRBT": {
		{
			ID:          10051,
			Slug:        "little-rabbit",
			Name:        "LITTLE RABBIT",
			BinanceUSDT: "LTRBTUSDT",
		},
	},
	"MOCA": {
		{
			ID:          10220,
			Slug:        "museum-of-crypto-art",
			Name:        "Museum of Crypto Art",
			BinanceUSDT: "MOCAUSDT",
		},
	},
	"FLEX": {
		{
			ID:          5190,
			Slug:        "flex",
			Name:        "FLEX",
			BinanceUSDT: "FLEXUSDT",
		},
	},
	"LEAD": {
		{
			ID:          6940,
			Slug:        "lead-wallet",
			Name:        "Lead Wallet",
			BinanceUSDT: "LEADUSDT",
		},
	},
	"NEWS": {
		{
			ID:          4647,
			Slug:        "publish",
			Name:        "PUBLISH",
			BinanceUSDT: "NEWSUSDT",
		},
	},
	"CICX": {
		{
			ID:          5406,
			Slug:        "cicoin",
			Name:        "Cicoin",
			BinanceUSDT: "CICXUSDT",
		},
	},
	"YFIVE": {
		{
			ID:          6812,
			Slug:        "yfive-finance",
			Name:        "YFIVE FINANCE",
			BinanceUSDT: "YFIVEUSDT",
		},
	},
	"YVS": {
		{
			ID:          8036,
			Slug:        "yvs-finance",
			Name:        "YVS.Finance",
			BinanceUSDT: "YVSUSDT",
		},
	},
	"X2P": {
		{
			ID:          9735,
			Slug:        "xenon-pay-ii",
			Name:        "Xenon Pay",
			BinanceUSDT: "X2PUSDT",
		},
	},
	"AMI": {
		{
			ID:          9008,
			Slug:        "ammyi-coin",
			Name:        "AMMYI Coin",
			BinanceUSDT: "AMIUSDT",
		},
	},
	"OXEN": {
		{
			ID:          2748,
			Slug:        "oxen",
			Name:        "Oxen",
			BinanceUSDT: "OXENUSDT",
		},
	},
	"FEY": {
		{
			ID:          10361,
			Slug:        "feyorra",
			Name:        "Feyorra",
			BinanceUSDT: "FEYUSDT",
		},
	},
	"ATTN": {
		{
			ID:          6803,
			Slug:        "attn",
			Name:        "ATTN",
			BinanceUSDT: "ATTNUSDT",
		},
	},
	"SRPH": {
		{
			ID:          10069,
			Slug:        "seraphium",
			Name:        "Seraphium",
			BinanceUSDT: "SRPHUSDT",
		},
	},
	"DETO": {
		{
			ID:          9050,
			Slug:        "delta-exchange-token",
			Name:        "Delta Exchange Token",
			BinanceUSDT: "DETOUSDT",
		},
	},
	"BITO": {
		{
			ID:          6118,
			Slug:        "bitopro-exchange-token",
			Name:        "BitoPro Exchange Token",
			BinanceUSDT: "BITOUSDT",
		},
	},
	"AUTZ": {
		{
			ID:          10660,
			Slug:        "autz-token",
			Name:        "AUTZ Token",
			BinanceUSDT: "AUTZUSDT",
		},
	},
	"BRN": {
		{
			ID:          9161,
			Slug:        "brainaut-defi",
			Name:        "Brainaut Defi",
			BinanceUSDT: "BRNUSDT",
		},
	},
	"GEMS": {
		{
			ID:   9849,
			Slug: "safegem-finance",
			Name: "SafeGem.Finance",
		},
		{
			ID:   10034,
			Slug: "carbon-token",
			Name: "CARBON",
		},
	},
	"PAPEL": {
		{
			ID:          9359,
			Slug:        "papel-token",
			Name:        "Papel Token",
			BinanceUSDT: "PAPELUSDT",
		},
	},
	"EMB": {
		{
			ID:          9626,
			Slug:        "emblem",
			Name:        "Emblem",
			BinanceUSDT: "EMBUSDT",
		},
	},
	"PIB": {
		{
			ID:          3768,
			Slug:        "pibble",
			Name:        "PIBBLE",
			BinanceUSDT: "PIBUSDT",
		},
	},
	"DOWS": {
		{
			ID:          8643,
			Slug:        "shadows",
			Name:        "Shadows",
			BinanceUSDT: "DOWSUSDT",
		},
	},
	"NAVY": {
		{
			ID:          3805,
			Slug:        "boat-pilot-token",
			Name:        "BoatPilot Token",
			BinanceUSDT: "NAVYUSDT",
		},
	},
	"EMD": {
		{
			ID:          50,
			Slug:        "emerald",
			Name:        "Emerald Crypto",
			BinanceUSDT: "EMDUSDT",
		},
	},
	"ETX": {
		{
			ID:          3771,
			Slug:        "ethereumx",
			Name:        "EthereumX",
			BinanceUSDT: "ETXUSDT",
		},
	},
	"CERBERUS": {
		{
			ID:          10438,
			Slug:        "gocerberus",
			Name:        "GoCerberus",
			BinanceUSDT: "CERBERUSUSDT",
		},
	},
	"ZGOAT": {
		{
			ID:          9705,
			Slug:        "goat-zuckerberg",
			Name:        "GOAT Zuckerberg",
			BinanceUSDT: "ZGOATUSDT",
		},
	},
	"APACHE": {
		{
			ID:          10094,
			Slug:        "apache",
			Name:        "Apache",
			BinanceUSDT: "APACHEUSDT",
		},
	},
	"MLM": {
		{
			ID:          2516,
			Slug:        "mktcoin",
			Name:        "MktCoin",
			BinanceUSDT: "MLMUSDT",
		},
	},
	"FINIX": {
		{
			ID:          10661,
			Slug:        "definix",
			Name:        "Definix",
			BinanceUSDT: "FINIXUSDT",
		},
	},
	"MARKGOAT": {
		{
			ID:          10437,
			Slug:        "mark-goat",
			Name:        "Mark Goat",
			BinanceUSDT: "MARKGOATUSDT",
		},
	},
	"EFAR": {
		{
			ID:          7361,
			Slug:        "fridn",
			Name:        "Fridn",
			BinanceUSDT: "EFARUSDT",
		},
	},
	"GLOX": {
		{
			ID:          7886,
			Slug:        "glox-finance",
			Name:        "Glox Finance",
			BinanceUSDT: "GLOXUSDT",
		},
	},
	"ETN": {
		{
			ID:          2137,
			Slug:        "electroneum",
			Name:        "Electroneum",
			BinanceUSDT: "ETNUSDT",
		},
	},
	"QUO": {
		{
			ID:          3230,
			Slug:        "quoxent",
			Name:        "Quoxent",
			BinanceUSDT: "QUOUSDT",
		},
	},
	"DODO": {
		{
			ID:          7224,
			Slug:        "dodo",
			Name:        "DODO",
			BinanceUSDT: "DODOUSDT",
		},
	},
	"SFG": {
		{
			ID:          7187,
			Slug:        "s-finance",
			Name:        "S.Finance",
			BinanceUSDT: "SFGUSDT",
		},
	},
	"MICRO": {
		{
			ID:          3610,
			Slug:        "micromines",
			Name:        "Micromines",
			BinanceUSDT: "MICROUSDT",
		},
	},
	"SKILL": {
		{
			ID:          9654,
			Slug:        "cryptoblades",
			Name:        "CryptoBlades",
			BinanceUSDT: "SKILLUSDT",
		},
	},
	"VELO": {
		{
			ID:          7127,
			Slug:        "velo",
			Name:        "Velo",
			BinanceUSDT: "VELOUSDT",
		},
	},
	"ZPR": {
		{
			ID:          2972,
			Slug:        "zper",
			Name:        "ZPER",
			BinanceUSDT: "ZPRUSDT",
		},
	},
	"ΤDOGE": {
		{
			ID:          10181,
			Slug:        "tdoge",
			Name:        "τDoge",
			BinanceUSDT: "ΤDOGEUSDT",
		},
	},
	"BSCX": {
		{
			ID:          8190,
			Slug:        "bscex",
			Name:        "BSCEX",
			BinanceUSDT: "BSCXUSDT",
		},
	},
	"ZLA": {
		{
			ID:          2500,
			Slug:        "zilla",
			Name:        "Zilla",
			BinanceUSDT: "ZLAUSDT",
		},
	},
	"LTK": {
		{
			ID:   3807,
			Slug: "litecoin-token",
			Name: "LitecoinToken",
		},
		{
			ID:   4212,
			Slug: "linktoken",
			Name: "LinkToken",
		},
	},
	"XMX": {
		{
			ID:          2859,
			Slug:        "xmax",
			Name:        "XMax",
			BinanceUSDT: "XMXUSDT",
		},
	},
	"ORS": {
		{
			ID:   2879,
			Slug: "origin-sport",
			Name: "Origin Sport",
		},
		{
			ID:   2911,
			Slug: "ors-group",
			Name: "ORS Group",
		},
	},
	"AUR": {
		{
			ID:   10991,
			Slug: "aurum",
			Name: "Aurum",
		},
		{
			ID:   148,
			Slug: "auroracoin",
			Name: "Auroracoin",
		},
		{
			ID:   8831,
			Slug: "aurix",
			Name: "Aurix",
		},
	},
	"SATA": {
		{
			ID:          9245,
			Slug:        "signata",
			Name:        "Signata",
			BinanceUSDT: "SATAUSDT",
		},
	},
	"ALLBI": {
		{
			ID:          5024,
			Slug:        "all-best-ico",
			Name:        "ALL BEST ICO",
			BinanceUSDT: "ALLBIUSDT",
		},
	},
	"PHB": {
		{
			ID:          2112,
			Slug:        "phoenix-global",
			Name:        "Phoenix Global",
			BinanceUSDT: "PHBUSDT",
		},
	},
	"XAVA": {
		{
			ID:          9797,
			Slug:        "avalaunch",
			Name:        "Avalaunch",
			BinanceUSDT: "XAVAUSDT",
		},
	},
	"NEO": {
		{
			ID:          1376,
			Slug:        "neo",
			Name:        "Neo",
			BinanceUSDT: "NEOUSDT",
		},
	},
	"KSM": {
		{
			ID:          5034,
			Slug:        "kusama",
			Name:        "Kusama",
			BinanceUSDT: "KSMUSDT",
		},
	},
	"ETM": {
		{
			ID:          4108,
			Slug:        "en-tan-mo",
			Name:        "En-Tan-Mo",
			BinanceUSDT: "ETMUSDT",
		},
	},
	"GYA": {
		{
			ID:          9101,
			Slug:        "gya",
			Name:        "GYA",
			BinanceUSDT: "GYAUSDT",
		},
	},
	"AKITA": {
		{
			ID:          8378,
			Slug:        "akita-inu",
			Name:        "Akita Inu",
			BinanceUSDT: "AKITAUSDT",
		},
	},
	"SFC": {
		{
			ID:          10937,
			Slug:        "safecap-token",
			Name:        "SafeCap Token",
			BinanceUSDT: "SFCUSDT",
		},
	},
	"TPT": {
		{
			ID:          5947,
			Slug:        "tokenpocket",
			Name:        "TokenPocket",
			BinanceUSDT: "TPTUSDT",
		},
	},
	"BLANK": {
		{
			ID:          8695,
			Slug:        "blank-wallet",
			Name:        "Blank Wallet",
			BinanceUSDT: "BLANKUSDT",
		},
	},
	"DMST": {
		{
			ID:          5952,
			Slug:        "dmscript",
			Name:        "DMScript",
			BinanceUSDT: "DMSTUSDT",
		},
	},
	"EHRT": {
		{
			ID:          6107,
			Slug:        "eight-hours",
			Name:        "Eight Hours",
			BinanceUSDT: "EHRTUSDT",
		},
	},
	"DEALDOUGH": {
		{
			ID:          10672,
			Slug:        "dealdough-token",
			Name:        "DealDough Token",
			BinanceUSDT: "DEALDOUGHUSDT",
		},
	},
	"mGOOGL": {
		{
			ID:          8003,
			Slug:        "mirrored-google",
			Name:        "Mirrored Google",
			BinanceUSDT: "mGOOGLUSDT",
		},
	},
	"STP": {
		{
			ID:          5785,
			Slug:        "stpay",
			Name:        "STPAY",
			BinanceUSDT: "STPUSDT",
		},
	},
	"SMC": {
		{
			ID:          113,
			Slug:        "smartcoin",
			Name:        "SmartCoin",
			BinanceUSDT: "SMCUSDT",
		},
	},
	"RWN": {
		{
			ID:          5886,
			Slug:        "rowan-token",
			Name:        "Rowan Token",
			BinanceUSDT: "RWNUSDT",
		},
	},
	"CELO": {
		{
			ID:          5567,
			Slug:        "celo",
			Name:        "Celo",
			BinanceUSDT: "CELOUSDT",
		},
	},
	"CDAI": {
		{
			ID:          5263,
			Slug:        "compound-dai",
			Name:        "Compound Dai",
			BinanceUSDT: "CDAIUSDT",
		},
	},
	"C2O": {
		{
			ID:          7214,
			Slug:        "cryptowater",
			Name:        "CryptoWater",
			BinanceUSDT: "C2OUSDT",
		},
	},
	"CSP": {
		{
			ID:          3842,
			Slug:        "caspian",
			Name:        "Caspian",
			BinanceUSDT: "CSPUSDT",
		},
	},
	"MNR": {
		{
			ID:          6053,
			Slug:        "mineral",
			Name:        "Mineral",
			BinanceUSDT: "MNRUSDT",
		},
	},
	"IMO": {
		{
			ID:          9281,
			Slug:        "imo",
			Name:        "IMO",
			BinanceUSDT: "IMOUSDT",
		},
	},
	"KISSMYMOON": {
		{
			ID:          10664,
			Slug:        "kissmymoon",
			Name:        "KissMyMoon",
			BinanceUSDT: "KISSMYMOONUSDT",
		},
	},
	"BURN1": {
		{
			ID:          9952,
			Slug:        "burn1-coin",
			Name:        "Burn1 Coin",
			BinanceUSDT: "BURN1USDT",
		},
	},
	"TKN": {
		{
			ID:          1660,
			Slug:        "monolith",
			Name:        "Monolith",
			BinanceUSDT: "TKNUSDT",
		},
	},
	"CHZ006": {
		{
			ID:          8990,
			Slug:        "charizardtoken",
			Name:        "Charizard Token",
			BinanceUSDT: "CHZ006USDT",
		},
	},
	"PLU": {
		{
			ID:          1392,
			Slug:        "pluton",
			Name:        "Pluton",
			BinanceUSDT: "PLUUSDT",
		},
	},
	"VMR": {
		{
			ID:          5093,
			Slug:        "vomer",
			Name:        "VOMER",
			BinanceUSDT: "VMRUSDT",
		},
	},
	"SAFECITY": {
		{
			ID:          10473,
			Slug:        "safecity",
			Name:        "SafeCity",
			BinanceUSDT: "SAFECITYUSDT",
		},
	},
	"ATOMBULL": {
		{
			ID:          6080,
			Slug:        "3x-long-cosmos-token",
			Name:        "3X Long Cosmos Token",
			BinanceUSDT: "ATOMBULLUSDT",
		},
	},
	"DILI": {
		{
			ID:          4793,
			Slug:        "d-community",
			Name:        "D Community",
			BinanceUSDT: "DILIUSDT",
		},
	},
	"IQN": {
		{
			ID:          3336,
			Slug:        "iqeon",
			Name:        "IQeon",
			BinanceUSDT: "IQNUSDT",
		},
	},
	"KP3R": {
		{
			ID:          7535,
			Slug:        "keep3rv1",
			Name:        "Keep3rV1",
			BinanceUSDT: "KP3RUSDT",
		},
	},
	"GHOSTFACE": {
		{
			ID:          10697,
			Slug:        "ghostface",
			Name:        "GHOSTFACE",
			BinanceUSDT: "GHOSTFACEUSDT",
		},
	},
	"STREAM": {
		{
			ID:          4365,
			Slug:        "streamit-coin",
			Name:        "Streamit Coin",
			BinanceUSDT: "STREAMUSDT",
		},
	},
	"SAFEEARTH": {
		{
			ID:          9081,
			Slug:        "safeearth",
			Name:        "SafeEarth",
			BinanceUSDT: "SAFEEARTHUSDT",
		},
	},
	"RVT": {
		{
			ID:          1991,
			Slug:        "rivetz",
			Name:        "Rivetz",
			BinanceUSDT: "RVTUSDT",
		},
	},
	"MRS": {
		{
			ID:          8634,
			Slug:        "marsan-exchange-token",
			Name:        "Marsan Exchange token",
			BinanceUSDT: "MRSUSDT",
		},
	},
	"XMR": {
		{
			ID:          328,
			Slug:        "monero",
			Name:        "Monero",
			BinanceUSDT: "XMRUSDT",
		},
	},
	"BSCGIRL": {
		{
			ID:          10570,
			Slug:        "binance-smart-chain-girl",
			Name:        "Binance Smart Chain Girl",
			BinanceUSDT: "BSCGIRLUSDT",
		},
	},
	"YUKI": {
		{
			ID:          3208,
			Slug:        "yuki",
			Name:        "YUKI",
			BinanceUSDT: "YUKIUSDT",
		},
	},
	"KEN": {
		{
			ID:          6621,
			Slug:        "keysians-network",
			Name:        "Keysians Network",
			BinanceUSDT: "KENUSDT",
		},
	},
	"DFI": {
		{
			ID:          5804,
			Slug:        "defichain",
			Name:        "DeFiChain",
			BinanceUSDT: "DFIUSDT",
		},
	},
	"BZE": {
		{
			ID:          3961,
			Slug:        "bzedge",
			Name:        "BZEdge",
			BinanceUSDT: "BZEUSDT",
		},
	},
	"PAK": {
		{
			ID:          1066,
			Slug:        "pakcoin",
			Name:        "Pakcoin",
			BinanceUSDT: "PAKUSDT",
		},
	},
	"MRCR": {
		{
			ID:          10170,
			Slug:        "mercor-finance",
			Name:        "Mercor Finance",
			BinanceUSDT: "MRCRUSDT",
		},
	},
	"XAUR": {
		{
			ID:          895,
			Slug:        "xaurum",
			Name:        "Xaurum",
			BinanceUSDT: "XAURUSDT",
		},
	},
	"UNO": {
		{
			ID:   8875,
			Slug: "unore",
			Name: "Uno Re",
		},
		{
			ID:   67,
			Slug: "unobtanium",
			Name: "Unobtanium",
		},
	},
	"YFDOT": {
		{
			ID:          7410,
			Slug:        "yearn-finance-dot",
			Name:        "Yearn Finance DOT",
			BinanceUSDT: "YFDOTUSDT",
		},
	},
	"YUANG": {
		{
			ID:          10327,
			Slug:        "yuang-coin",
			Name:        "Yuang Coin",
			BinanceUSDT: "YUANGUSDT",
		},
	},
	"SINS": {
		{
			ID:          3366,
			Slug:        "safeinsure",
			Name:        "SafeInsure",
			BinanceUSDT: "SINSUSDT",
		},
	},
	"GTO": {
		{
			ID:          2289,
			Slug:        "gifto",
			Name:        "Gifto",
			BinanceUSDT: "GTOUSDT",
		},
	},
	"MELODY": {
		{
			ID:          10252,
			Slug:        "mozart-finance",
			Name:        "Mozart Finance",
			BinanceUSDT: "MELODYUSDT",
		},
	},
	"C3": {
		{
			ID:          9991,
			Slug:        "charli3",
			Name:        "Charli3",
			BinanceUSDT: "C3USDT",
		},
	},
	"HGT": {
		{
			ID:   2042,
			Slug: "hellogold",
			Name: "HelloGold",
		},
		{
			ID:   9664,
			Slug: "hypergraph",
			Name: "HyperGraph",
		},
	},
	"IDRT": {
		{
			ID:          4702,
			Slug:        "rupiah-token",
			Name:        "Rupiah Token",
			BinanceUSDT: "IDRTUSDT",
		},
	},
	"CL": {
		{
			ID:          2352,
			Slug:        "coinlancer",
			Name:        "Coinlancer",
			BinanceUSDT: "CLUSDT",
		},
	},
	"ABT": {
		{
			ID:          2545,
			Slug:        "arcblock",
			Name:        "Arcblock",
			BinanceUSDT: "ABTUSDT",
		},
	},
	"B1P": {
		{
			ID:          5383,
			Slug:        "b-one-payment",
			Name:        "B ONE PAYMENT",
			BinanceUSDT: "B1PUSDT",
		},
	},
	"BEAR": {
		{
			ID:   5158,
			Slug: "3x-short-bitcoin-token",
			Name: "3X Short Bitcoin Token",
		},
		{
			ID:   6934,
			Slug: "voytek-bear-coin",
			Name: "BEAR Coin",
		},
	},
	"LESS": {
		{
			ID:          10279,
			Slug:        "less-network",
			Name:        "Less Network",
			BinanceUSDT: "LESSUSDT",
		},
	},
	"GEAR": {
		{
			ID:          6593,
			Slug:        "bitgear",
			Name:        "Bitgear",
			BinanceUSDT: "GEARUSDT",
		},
	},
	"SLEEPY": {
		{
			ID:          9774,
			Slug:        "sleepy-sloth-finance",
			Name:        "Sleepy Sloth Finance",
			BinanceUSDT: "SLEEPYUSDT",
		},
	},
	"ICEBRK": {
		{
			ID:          10304,
			Slug:        "icebreak-r",
			Name:        "IceBreak-R",
			BinanceUSDT: "ICEBRKUSDT",
		},
	},
	"REP": {
		{
			ID:          1104,
			Slug:        "augur",
			Name:        "Augur",
			BinanceUSDT: "REPUSDT",
		},
	},
	"LOC": {
		{
			ID:          2287,
			Slug:        "lockchain",
			Name:        "LockTrip",
			BinanceUSDT: "LOCUSDT",
		},
	},
	"100X": {
		{
			ID:          9460,
			Slug:        "100xcoin",
			Name:        "100xCoin",
			BinanceUSDT: "100XUSDT",
		},
	},
	"NUT": {
		{
			ID:   10836,
			Slug: "nut-money",
			Name: "NUT MONEY",
		},
		{
			ID:   4033,
			Slug: "native-utility-token",
			Name: "Native Utility Token",
		},
	},
	"XCP": {
		{
			ID:          132,
			Slug:        "counterparty",
			Name:        "Counterparty",
			BinanceUSDT: "XCPUSDT",
		},
	},
	"G9TRO": {
		{
			ID:          9594,
			Slug:        "g9tro-crowdfunding-platform",
			Name:        "g9tro Crowdfunding Platform",
			BinanceUSDT: "G9TROUSDT",
		},
	},
	"VOT": {
		{
			ID:          2221,
			Slug:        "votecoin",
			Name:        "VoteCoin",
			BinanceUSDT: "VOTUSDT",
		},
	},
	"EQMT": {
		{
			ID:          6572,
			Slug:        "equus-mining-token",
			Name:        "Equus Mining Token",
			BinanceUSDT: "EQMTUSDT",
		},
	},
	"CGG": {
		{
			ID:          8648,
			Slug:        "chain-guardians",
			Name:        "Chain Guardians",
			BinanceUSDT: "CGGUSDT",
		},
	},
	"HLX": {
		{
			ID:          5057,
			Slug:        "helex",
			Name:        "Helex",
			BinanceUSDT: "HLXUSDT",
		},
	},
	"JORDAN": {
		{
			ID:          10915,
			Slug:        "cryingjordan-token",
			Name:        "cryingJORDAN Token",
			BinanceUSDT: "JORDANUSDT",
		},
	},
	"BTC3L": {
		{
			ID:          5736,
			Slug:        "amun-bitcoin-3x-daily-long",
			Name:        "Amun Bitcoin 3x Daily Long",
			BinanceUSDT: "BTC3LUSDT",
		},
	},
	"TEAT": {
		{
			ID:          7614,
			Slug:        "teal",
			Name:        "TEAL",
			BinanceUSDT: "TEATUSDT",
		},
	},
	"LIBREF": {
		{
			ID:          7752,
			Slug:        "librefreelencer",
			Name:        "LibreFreelencer",
			BinanceUSDT: "LIBREFUSDT",
		},
	},
	"YFD": {
		{
			ID:          7690,
			Slug:        "yfdfi-finance",
			Name:        "Your Finance Decentralized",
			BinanceUSDT: "YFDUSDT",
		},
	},
	"BASK": {
		{
			ID:          9157,
			Slug:        "basketdao",
			Name:        "BasketDAO",
			BinanceUSDT: "BASKUSDT",
		},
	},
	"LTO": {
		{
			ID:          3714,
			Slug:        "lto-network",
			Name:        "LTO Network",
			BinanceUSDT: "LTOUSDT",
		},
	},
	"LML": {
		{
			ID:          3636,
			Slug:        "lisk-machine-learning",
			Name:        "Lisk Machine Learning",
			BinanceUSDT: "LMLUSDT",
		},
	},
	"GGT": {
		{
			ID:          7658,
			Slug:        "gard-governance-token",
			Name:        "GARD Governance Token",
			BinanceUSDT: "GGTUSDT",
		},
	},
	"SPH": {
		{
			ID:          6209,
			Slug:        "spheroid-universe",
			Name:        "Spheroid Universe",
			BinanceUSDT: "SPHUSDT",
		},
	},
	"SFMS": {
		{
			ID:          10204,
			Slug:        "safemoon-swap",
			Name:        "SafeMoon.swap",
			BinanceUSDT: "SFMSUSDT",
		},
	},
	"VIT": {
		{
			ID:          11015,
			Slug:        "team-vitality-fan-token",
			Name:        "Team Vitality Fan Token",
			BinanceUSDT: "VITUSDT",
		},
	},
	"WAIV": {
		{
			ID:          10470,
			Slug:        "waivlength",
			Name:        "Waivlength",
			BinanceUSDT: "WAIVUSDT",
		},
	},
	"GZE": {
		{
			ID:          3257,
			Slug:        "gazecoin",
			Name:        "GazeCoin",
			BinanceUSDT: "GZEUSDT",
		},
	},
	"DAILY": {
		{
			ID:          8910,
			Slug:        "daily",
			Name:        "Daily",
			BinanceUSDT: "DAILYUSDT",
		},
	},
	"DJV": {
		{
			ID:          7834,
			Slug:        "dejave",
			Name:        "DEJAVE",
			BinanceUSDT: "DJVUSDT",
		},
	},
	"PUNDIX": {
		{
			ID:          9040,
			Slug:        "pundix-new",
			Name:        "Pundi X[new]",
			BinanceUSDT: "PUNDIXUSDT",
		},
	},
	"STAX": {
		{
			ID:          7502,
			Slug:        "stablexswap",
			Name:        "StableXSwap",
			BinanceUSDT: "STAXUSDT",
		},
	},
	"TUP": {
		{
			ID:          4411,
			Slug:        "tenup",
			Name:        "TenUp",
			BinanceUSDT: "TUPUSDT",
		},
	},
	"GBCR": {
		{
			ID:          7557,
			Slug:        "gold-bcr",
			Name:        "Gold BCR",
			BinanceUSDT: "GBCRUSDT",
		},
	},
	"SHARK": {
		{
			ID:          9949,
			Slug:        "baby-shark",
			Name:        "Baby Shark",
			BinanceUSDT: "SHARKUSDT",
		},
	},
	"7E": {
		{
			ID:          4063,
			Slug:        "7eleven",
			Name:        "7Eleven",
			BinanceUSDT: "7EUSDT",
		},
	},
	"BITSZ": {
		{
			ID:          9292,
			Slug:        "bitsz",
			Name:        "Bitsz",
			BinanceUSDT: "BITSZUSDT",
		},
	},
	"ELP": {
		{
			ID:          10392,
			Slug:        "the-everlasting-parachain",
			Name:        "The Everlasting Parachain",
			BinanceUSDT: "ELPUSDT",
		},
	},
	"DSCPL": {
		{
			ID:          3761,
			Slug:        "disciplina",
			Name:        "DISCIPLINA",
			BinanceUSDT: "DSCPLUSDT",
		},
	},
	"ARB": {
		{
			ID:          938,
			Slug:        "arbit",
			Name:        "ARbit",
			BinanceUSDT: "ARBUSDT",
		},
	},
	"YFID": {
		{
			ID:          7880,
			Slug:        "yfidapp",
			Name:        "YFIDapp",
			BinanceUSDT: "YFIDUSDT",
		},
	},
	"UIP": {
		{
			ID:          2454,
			Slug:        "unlimitedip",
			Name:        "UnlimitedIP",
			BinanceUSDT: "UIPUSDT",
		},
	},
	"DSWAP": {
		{
			ID:          8195,
			Slug:        "definex",
			Name:        "Definex",
			BinanceUSDT: "DSWAPUSDT",
		},
	},
	"CPX": {
		{
			ID:          9945,
			Slug:        "center-prime-token",
			Name:        "CenterPrime",
			BinanceUSDT: "CPXUSDT",
		},
	},
	"PWRB": {
		{
			ID:          5965,
			Slug:        "powerbalt",
			Name:        "PowerBalt",
			BinanceUSDT: "PWRBUSDT",
		},
	},
	"DIVI": {
		{
			ID:          3441,
			Slug:        "divi",
			Name:        "Divi",
			BinanceUSDT: "DIVIUSDT",
		},
	},
	"MODX": {
		{
			ID:          3479,
			Slug:        "model-x-coin",
			Name:        "MODEL-X-coin",
			BinanceUSDT: "MODXUSDT",
		},
	},
	"2GIVE": {
		{
			ID:          1252,
			Slug:        "2give",
			Name:        "2GIVE",
			BinanceUSDT: "2GIVEUSDT",
		},
	},
	"AQUAPIG": {
		{
			ID:          10150,
			Slug:        "aqua-pig",
			Name:        "Aqua Pig",
			BinanceUSDT: "AQUAPIGUSDT",
		},
	},
	"ELF": {
		{
			ID:          2299,
			Slug:        "aelf",
			Name:        "aelf",
			BinanceUSDT: "ELFUSDT",
		},
	},
	"CEN": {
		{
			ID:   3010,
			Slug: "coinsuper-ecosystem-network",
			Name: "Coinsuper Ecosystem Network",
		},
		{
			ID:   3229,
			Slug: "centaure",
			Name: "Centaure",
		},
	},
	"7ADD": {
		{
			ID:          7771,
			Slug:        "holdtowin",
			Name:        "HoldToWin",
			BinanceUSDT: "7ADDUSDT",
		},
	},
	"BMX": {
		{
			ID:          2933,
			Slug:        "bitmart-token",
			Name:        "BitMart Token",
			BinanceUSDT: "BMXUSDT",
		},
	},
	"MOMO": {
		{
			ID:          9271,
			Slug:        "momo-protocol",
			Name:        "Momo Protocol",
			BinanceUSDT: "MOMOUSDT",
		},
	},
	"XLA": {
		{
			ID:   2629,
			Slug: "scala",
			Name: "Scala",
		},
		{
			ID:   6528,
			Slug: "ripple-alpha",
			Name: "Ripple Alpha",
		},
	},
	"CYMT": {
		{
			ID:          3255,
			Slug:        "cybermusic",
			Name:        "CyberMusic",
			BinanceUSDT: "CYMTUSDT",
		},
	},
	"CTF": {
		{
			ID:          8791,
			Slug:        "cybertime-finance-token",
			Name:        "CyberTime Finance Token",
			BinanceUSDT: "CTFUSDT",
		},
	},
	"KCLP": {
		{
			ID:          11041,
			Slug:        "kucoin-launchpad",
			Name:        "KuCoin LaunchPad",
			BinanceUSDT: "KCLPUSDT",
		},
	},
	"BLUE": {
		{
			ID:   2076,
			Slug: "ethereum-blue",
			Name: "Blue Protocol",
		},
		{
			ID:   9162,
			Slug: "blue-swap",
			Name: "Blue Swap",
		},
	},
	"BBP": {
		{
			ID:   1916,
			Slug: "biblepay",
			Name: "BiblePay",
		},
		{
			ID:   8312,
			Slug: "bitbot-protocol",
			Name: "Bitbot Protocol",
		},
	},
	"PPT": {
		{
			ID:          1789,
			Slug:        "populous",
			Name:        "Populous",
			BinanceUSDT: "PPTUSDT",
		},
	},
	"CNN": {
		{
			ID:          2735,
			Slug:        "content-neutrality-network",
			Name:        "Content Neutrality Network",
			BinanceUSDT: "CNNUSDT",
		},
	},
	"NTRN": {
		{
			ID:          894,
			Slug:        "neutron",
			Name:        "Neutron",
			BinanceUSDT: "NTRNUSDT",
		},
	},
	"HPX": {
		{
			ID:          6062,
			Slug:        "hupayx",
			Name:        "HUPAYX",
			BinanceUSDT: "HPXUSDT",
		},
	},
	"MDM": {
		{
			ID:          4928,
			Slug:        "medium",
			Name:        "Medium",
			BinanceUSDT: "MDMUSDT",
		},
	},
	"PTA": {
		{
			ID:          8382,
			Slug:        "petrachor",
			Name:        "Petrachor",
			BinanceUSDT: "PTAUSDT",
		},
	},
	"VGD": {
		{
			ID:          9240,
			Slug:        "vangold-token",
			Name:        "Vangold Token",
			BinanceUSDT: "VGDUSDT",
		},
	},
	"LOTTO": {
		{
			ID:          8399,
			Slug:        "lotto",
			Name:        "Lotto",
			BinanceUSDT: "LOTTOUSDT",
		},
	},
	"WEB": {
		{
			ID:          3027,
			Slug:        "webcoin",
			Name:        "Webcoin",
			BinanceUSDT: "WEBUSDT",
		},
	},
	"AMN": {
		{
			ID:          2705,
			Slug:        "amon",
			Name:        "Amon",
			BinanceUSDT: "AMNUSDT",
		},
	},
	"GIN": {
		{
			ID:          2773,
			Slug:        "gincoin",
			Name:        "GINcoin",
			BinanceUSDT: "GINUSDT",
		},
	},
	"WANUNI": {
		{
			ID:          8654,
			Slug:        "wanuni",
			Name:        "wanUNI",
			BinanceUSDT: "WANUNIUSDT",
		},
	},
	"SBT": {
		{
			ID:          4913,
			Slug:        "solbit",
			Name:        "SOLBIT",
			BinanceUSDT: "SBTUSDT",
		},
	},
	"CF": {
		{
			ID:          898,
			Slug:        "californium",
			Name:        "Californium",
			BinanceUSDT: "CFUSDT",
		},
	},
	"NCC": {
		{
			ID:          3016,
			Slug:        "neurochain",
			Name:        "NeuroChain",
			BinanceUSDT: "NCCUSDT",
		},
	},
	"SPP": {
		{
			ID:          10483,
			Slug:        "shapepay",
			Name:        "ShapePay",
			BinanceUSDT: "SPPUSDT",
		},
	},
	"NFTART": {
		{
			ID:          9299,
			Slug:        "nft-art-finance",
			Name:        "NFT Art Finance",
			BinanceUSDT: "NFTARTUSDT",
		},
	},
	"CEL": {
		{
			ID:          2700,
			Slug:        "celsius",
			Name:        "Celsius",
			BinanceUSDT: "CELUSDT",
		},
	},
	"VERA": {
		{
			ID:          4830,
			Slug:        "vera",
			Name:        "VERA",
			BinanceUSDT: "VERAUSDT",
		},
	},
	"CHARIX TOKEN": {
		{
			ID:          10241,
			Slug:        "charix",
			Name:        "Charix",
			BinanceUSDT: "CHARIX TOKENUSDT",
		},
	},
	"ZYX": {
		{
			ID:          6131,
			Slug:        "zyx",
			Name:        "ZYX",
			BinanceUSDT: "ZYXUSDT",
		},
	},
	"BNKR": {
		{
			ID:          4918,
			Slug:        "bankroll-network",
			Name:        "Bankroll Network",
			BinanceUSDT: "BNKRUSDT",
		},
	},
	"SATS": {
		{
			ID:          9022,
			Slug:        "satoshi",
			Name:        "Satoshi",
			BinanceUSDT: "SATSUSDT",
		},
	},
	"GIFT": {
		{
			ID:          10283,
			Slug:        "gift-coin",
			Name:        "Gift-Coin",
			BinanceUSDT: "GIFTUSDT",
		},
	},
	"ERC20": {
		{
			ID:          2165,
			Slug:        "erc20",
			Name:        "ERC20",
			BinanceUSDT: "ERC20USDT",
		},
	},
	"SOLAPE": {
		{
			ID:          10452,
			Slug:        "solapefinance",
			Name:        "SOLAPE Finance",
			BinanceUSDT: "SOLAPEUSDT",
		},
	},
	"GINUX": {
		{
			ID:          10414,
			Slug:        "green-shiba-inu-new",
			Name:        "Green Shiba Inu (new)",
			BinanceUSDT: "GINUXUSDT",
		},
	},
	"AYA": {
		{
			ID:          3973,
			Slug:        "aryacoin",
			Name:        "Aryacoin",
			BinanceUSDT: "AYAUSDT",
		},
	},
	"SENC": {
		{
			ID:          2624,
			Slug:        "sentinel-chain",
			Name:        "Sentinel Chain",
			BinanceUSDT: "SENCUSDT",
		},
	},
	"PALG": {
		{
			ID:          9819,
			Slug:        "palgold",
			Name:        "PalGold",
			BinanceUSDT: "PALGUSDT",
		},
	},
	"RAZE": {
		{
			ID:          9173,
			Slug:        "raze-network",
			Name:        "Raze Network",
			BinanceUSDT: "RAZEUSDT",
		},
	},
	"SMTY": {
		{
			ID:          7594,
			Slug:        "smoothy",
			Name:        "Smoothy",
			BinanceUSDT: "SMTYUSDT",
		},
	},
	"DIAH": {
		{
			ID:          10459,
			Slug:        "diarrheacoin",
			Name:        "DiarrheaCoin",
			BinanceUSDT: "DIAHUSDT",
		},
	},
	"BLKC": {
		{
			ID:          10566,
			Slug:        "blackhat",
			Name:        "BlackHat",
			BinanceUSDT: "BLKCUSDT",
		},
	},
	"ALIA": {
		{
			ID:          9099,
			Slug:        "xanalia",
			Name:        "XANALIA",
			BinanceUSDT: "ALIAUSDT",
		},
	},
	"QUN": {
		{
			ID:          2375,
			Slug:        "qunqun",
			Name:        "QunQun",
			BinanceUSDT: "QUNUSDT",
		},
	},
	"APW": {
		{
			ID:          10364,
			Slug:        "apwine-finance",
			Name:        "APWine Finance",
			BinanceUSDT: "APWUSDT",
		},
	},
	"HARD": {
		{
			ID:          7576,
			Slug:        "hard-protocol",
			Name:        "HARD Protocol",
			BinanceUSDT: "HARDUSDT",
		},
	},
	"MOMMYDOGE": {
		{
			ID:          10876,
			Slug:        "mommy-doge-coin",
			Name:        "Mommy Doge Coin",
			BinanceUSDT: "MOMMYDOGEUSDT",
		},
	},
	"ELLA": {
		{
			ID:          2122,
			Slug:        "ellaism",
			Name:        "Ellaism",
			BinanceUSDT: "ELLAUSDT",
		},
	},
	"$TREAM": {
		{
			ID:          11006,
			Slug:        "world-stream-finance",
			Name:        "World Stream Finance",
			BinanceUSDT: "$TREAMUSDT",
		},
	},
	"LION": {
		{
			ID:          9598,
			Slug:        "lion-token",
			Name:        "Lion Token",
			BinanceUSDT: "LIONUSDT",
		},
	},
	"ADB": {
		{
			ID:          2501,
			Slug:        "adbank",
			Name:        "adbank",
			BinanceUSDT: "ADBUSDT",
		},
	},
	"PRESIDENTDOGE": {
		{
			ID:          10946,
			Slug:        "presidentdoge",
			Name:        "PresidentDoge",
			BinanceUSDT: "PRESIDENTDOGEUSDT",
		},
	},
	"CFL": {
		{
			ID:          3313,
			Slug:        "cryptoflow",
			Name:        "CryptoFlow",
			BinanceUSDT: "CFLUSDT",
		},
	},
	"GTM": {
		{
			ID:          3215,
			Slug:        "gentarium",
			Name:        "Gentarium",
			BinanceUSDT: "GTMUSDT",
		},
	},
	"MKR": {
		{
			ID:          1518,
			Slug:        "maker",
			Name:        "Maker",
			BinanceUSDT: "MKRUSDT",
		},
	},
	"VIKING": {
		{
			ID:          8555,
			Slug:        "viking-swap",
			Name:        "Viking Swap",
			BinanceUSDT: "VIKINGUSDT",
		},
	},
	"YOU": {
		{
			ID:   3053,
			Slug: "you-coin",
			Name: "YOU COIN",
		},
		{
			ID:   10070,
			Slug: "youswap",
			Name: "YouSwap",
		},
	},
	"ADA": {
		{
			ID:          2010,
			Slug:        "cardano",
			Name:        "Cardano",
			BinanceUSDT: "ADAUSDT",
		},
	},
	"NEKO": {
		{
			ID:          9963,
			Slug:        "neko-network",
			Name:        "Neko Network",
			BinanceUSDT: "NEKOUSDT",
		},
	},
	"EXZO": {
		{
			ID:          10528,
			Slug:        "exzocoin-2",
			Name:        "ExzoCoin 2.0",
			BinanceUSDT: "EXZOUSDT",
		},
	},
	"CATBREAD": {
		{
			ID:          10887,
			Slug:        "catbread",
			Name:        "CatBread",
			BinanceUSDT: "CATBREADUSDT",
		},
	},
	"AMPL": {
		{
			ID:          4056,
			Slug:        "ampleforth",
			Name:        "Ampleforth",
			BinanceUSDT: "AMPLUSDT",
		},
	},
	"BTSG": {
		{
			ID:          8905,
			Slug:        "bitsong",
			Name:        "BitSong",
			BinanceUSDT: "BTSGUSDT",
		},
	},
	"GUSD": {
		{
			ID:          3306,
			Slug:        "gemini-dollar",
			Name:        "Gemini Dollar",
			BinanceUSDT: "GUSDUSDT",
		},
	},
	"SAFECOMET": {
		{
			ID:          9568,
			Slug:        "safecomet",
			Name:        "SafeComet",
			BinanceUSDT: "SAFECOMETUSDT",
		},
	},
	"ICON": {
		{
			ID:          1528,
			Slug:        "iconic",
			Name:        "Iconic",
			BinanceUSDT: "ICONUSDT",
		},
	},
	"IFLT": {
		{
			ID:          1504,
			Slug:        "inflationcoin",
			Name:        "InflationCoin",
			BinanceUSDT: "IFLTUSDT",
		},
	},
	"TUBE": {
		{
			ID:          2561,
			Slug:        "bit-tube",
			Name:        "BitTube",
			BinanceUSDT: "TUBEUSDT",
		},
	},
	"QWC": {
		{
			ID:          3942,
			Slug:        "qwertycoin",
			Name:        "Qwertycoin",
			BinanceUSDT: "QWCUSDT",
		},
	},
	"DAPP": {
		{
			ID:          4026,
			Slug:        "liquid-apps",
			Name:        "LiquidApps",
			BinanceUSDT: "DAPPUSDT",
		},
	},
	"MUSUBI": {
		{
			ID:          10708,
			Slug:        "musubi",
			Name:        "Musubi",
			BinanceUSDT: "MUSUBIUSDT",
		},
	},
	"PING": {
		{
			ID:   1777,
			Slug: "cryptoping",
			Name: "CryptoPing",
		},
		{
			ID:   10774,
			Slug: "sonar",
			Name: "Sonar",
		},
	},
	"SERO": {
		{
			ID:          4078,
			Slug:        "super-zero-protocol",
			Name:        "Super Zero Protocol",
			BinanceUSDT: "SEROUSDT",
		},
	},
	"PINKE": {
		{
			ID:          9929,
			Slug:        "pinkelon",
			Name:        "PinkElon",
			BinanceUSDT: "PINKEUSDT",
		},
	},
	"BWB": {
		{
			ID:          6004,
			Slug:        "bw-token",
			Name:        "Bit World Token",
			BinanceUSDT: "BWBUSDT",
		},
	},
	"XFIT": {
		{
			ID:          9217,
			Slug:        "xfai",
			Name:        "XFai",
			BinanceUSDT: "XFITUSDT",
		},
	},
	"GLXM": {
		{
			ID:          10305,
			Slug:        "galaxium",
			Name:        "Galaxium",
			BinanceUSDT: "GLXMUSDT",
		},
	},
	"KABOSU": {
		{
			ID:          9710,
			Slug:        "kabosu",
			Name:        "Kabosu",
			BinanceUSDT: "KABOSUUSDT",
		},
	},
	"CMERGE": {
		{
			ID:          11016,
			Slug:        "coinmerge",
			Name:        "CoinMerge",
			BinanceUSDT: "CMERGEUSDT",
		},
	},
	"EVX": {
		{
			ID:          2034,
			Slug:        "everex",
			Name:        "Everex",
			BinanceUSDT: "EVXUSDT",
		},
	},
	"BOLTT": {
		{
			ID:          3789,
			Slug:        "boltt-coin",
			Name:        "Boltt Coin ",
			BinanceUSDT: "BOLTTUSDT",
		},
	},
	"TORN": {
		{
			ID:          8049,
			Slug:        "torn",
			Name:        "Tornado Cash",
			BinanceUSDT: "TORNUSDT",
		},
	},
	"FEED": {
		{
			ID:   10109,
			Slug: "feeder-finance",
			Name: "Feeder.finance",
		},
		{
			ID:   9793,
			Slug: "feed-token",
			Name: "FEED Token",
		},
	},
	"EVN": {
		{
			ID:   3261,
			Slug: "evencoin",
			Name: "EvenCoin",
		},
		{
			ID:   9580,
			Slug: "evolution-finance",
			Name: "Evolution Finance",
		},
	},
	"FWB": {
		{
			ID:          10090,
			Slug:        "friends-with-benefits-pro",
			Name:        "Friends With Benefits Pro",
			BinanceUSDT: "FWBUSDT",
		},
	},
	"KAT": {
		{
			ID:          3634,
			Slug:        "kambria",
			Name:        "Kambria",
			BinanceUSDT: "KATUSDT",
		},
	},
	"SAFEWIN": {
		{
			ID:          10536,
			Slug:        "safewin",
			Name:        "SafeWin",
			BinanceUSDT: "SAFEWINUSDT",
		},
	},
	"BOX": {
		{
			ID:   3475,
			Slug: "box-token",
			Name: "BOX Token",
		},
		{
			ID:   2945,
			Slug: "contentbox",
			Name: "ContentBox",
		},
		{
			ID:   6960,
			Slug: "defibox",
			Name: "DefiBox",
		},
	},
	"ALLOY": {
		{
			ID:          9038,
			Slug:        "hyper-alloy",
			Name:        "HyperAlloy",
			BinanceUSDT: "ALLOYUSDT",
		},
	},
	"AGAR": {
		{
			ID:          8251,
			Slug:        "agar",
			Name:        "AGAr",
			BinanceUSDT: "AGARUSDT",
		},
	},
	"PALLA": {
		{
			ID:          10669,
			Slug:        "pallapay",
			Name:        "Pallapay",
			BinanceUSDT: "PALLAUSDT",
		},
	},
	"DFS": {
		{
			ID:   6610,
			Slug: "defis-network",
			Name: "Defis Network",
		},
		{
			ID:   3329,
			Slug: "fantasy-sports",
			Name: "Fantasy Sports",
		},
	},
	"FERA": {
		{
			ID:          6612,
			Slug:        "fera",
			Name:        "Fera",
			BinanceUSDT: "FERAUSDT",
		},
	},
	"IZE": {
		{
			ID:          5677,
			Slug:        "ize",
			Name:        "IZE",
			BinanceUSDT: "IZEUSDT",
		},
	},
	"XLMUP": {
		{
			ID:          8055,
			Slug:        "xlmup",
			Name:        "XLMUP",
			BinanceUSDT: "XLMUPUSDT",
		},
	},
	"CYCLUB": {
		{
			ID:          6810,
			Slug:        "cyclub",
			Name:        "CYCLUB",
			BinanceUSDT: "CYCLUBUSDT",
		},
	},
	"ALOHA": {
		{
			ID:          8548,
			Slug:        "aloha",
			Name:        "Aloha",
			BinanceUSDT: "ALOHAUSDT",
		},
	},
	"BNTY": {
		{
			ID:          2310,
			Slug:        "bounty0x",
			Name:        "Bounty0x",
			BinanceUSDT: "BNTYUSDT",
		},
	},
	"NZDX": {
		{
			ID:          5494,
			Slug:        "etoro-new-zealand-dollar",
			Name:        "eToro New Zealand Dollar",
			BinanceUSDT: "NZDXUSDT",
		},
	},
	"DDRT": {
		{
			ID:          5697,
			Slug:        "digidinar-token",
			Name:        "DigiDinar Token",
			BinanceUSDT: "DDRTUSDT",
		},
	},
	"PEOS": {
		{
			ID:          3910,
			Slug:        "peos",
			Name:        "pEOS",
			BinanceUSDT: "PEOSUSDT",
		},
	},
	"ACAT": {
		{
			ID:          2525,
			Slug:        "alphacat",
			Name:        "Alphacat",
			BinanceUSDT: "ACATUSDT",
		},
	},
	"SNOB": {
		{
			ID:          9780,
			Slug:        "snowball-finance",
			Name:        "Snowball",
			BinanceUSDT: "SNOBUSDT",
		},
	},
	"EL": {
		{
			ID:          5382,
			Slug:        "elysia",
			Name:        "ELYSIA",
			BinanceUSDT: "ELUSDT",
		},
	},
	"BOOM": {
		{
			ID:          4128,
			Slug:        "boom",
			Name:        "BOOM",
			BinanceUSDT: "BOOMUSDT",
		},
	},
	"PC": {
		{
			ID:          3061,
			Slug:        "promotion-coin",
			Name:        "Promotion Coin",
			BinanceUSDT: "PCUSDT",
		},
	},
	"SCS": {
		{
			ID:          1651,
			Slug:        "speedcash",
			Name:        "SpeedCash",
			BinanceUSDT: "SCSUSDT",
		},
	},
	"TKO": {
		{
			ID:          9020,
			Slug:        "tokocrypto",
			Name:        "Toko Token",
			BinanceUSDT: "TKOUSDT",
		},
	},
	"MERL": {
		{
			ID:          9853,
			Slug:        "merlin",
			Name:        "Merlin",
			BinanceUSDT: "MERLUSDT",
		},
	},
	"NYE": {
		{
			ID:          4268,
			Slug:        "newyork-exchange",
			Name:        "NewYork Exchange",
			BinanceUSDT: "NYEUSDT",
		},
	},
	"USHIBA": {
		{
			ID:          9943,
			Slug:        "american-shiba",
			Name:        "American Shiba",
			BinanceUSDT: "USHIBAUSDT",
		},
	},
	"IGNIS": {
		{
			ID:          2276,
			Slug:        "ignis",
			Name:        "Ignis",
			BinanceUSDT: "IGNISUSDT",
		},
	},
	"MTH": {
		{
			ID:          1947,
			Slug:        "monetha",
			Name:        "Monetha",
			BinanceUSDT: "MTHUSDT",
		},
	},
	"MCRN": {
		{
			ID:          8880,
			Slug:        "macaronswap",
			Name:        "MacaronSwap",
			BinanceUSDT: "MCRNUSDT",
		},
	},
	"XMON": {
		{
			ID:          8509,
			Slug:        "xmon",
			Name:        "XMON",
			BinanceUSDT: "XMONUSDT",
		},
	},
	"pBTC35A": {
		{
			ID:          8252,
			Slug:        "pbtc35a",
			Name:        "pBTC35A",
			BinanceUSDT: "pBTC35AUSDT",
		},
	},
	"EGGP": {
		{
			ID:          10301,
			Slug:        "eggplant-finance",
			Name:        "Eggplant Finance",
			BinanceUSDT: "EGGPUSDT",
		},
	},
	"RPD": {
		{
			ID:          3432,
			Slug:        "rapids",
			Name:        "Rapids",
			BinanceUSDT: "RPDUSDT",
		},
	},
	"UFR": {
		{
			ID:          2178,
			Slug:        "upfiring",
			Name:        "Upfiring",
			BinanceUSDT: "UFRUSDT",
		},
	},
	"DACS": {
		{
			ID:          3054,
			Slug:        "dacsee",
			Name:        "DACSEE",
			BinanceUSDT: "DACSUSDT",
		},
	},
	"OWC": {
		{
			ID:          3763,
			Slug:        "oduwa",
			Name:        "ODUWA",
			BinanceUSDT: "OWCUSDT",
		},
	},
	"QTBK": {
		{
			ID:          6738,
			Slug:        "quantbook",
			Name:        "Quantbook",
			BinanceUSDT: "QTBKUSDT",
		},
	},
	"MOONPIRATE": {
		{
			ID:          9886,
			Slug:        "moonpirate",
			Name:        "MoonPirate",
			BinanceUSDT: "MOONPIRATEUSDT",
		},
	},
	"EPS": {
		{
			ID:   8938,
			Slug: "ellipsis",
			Name: "Ellipsis",
		},
		{
			ID:   6824,
			Slug: "epanus",
			Name: "Epanus",
		},
	},
	"SFJP": {
		{
			ID:          9880,
			Slug:        "safejupiter-sfjp",
			Name:        "SafeJupiter $SFJP",
			BinanceUSDT: "SFJPUSDT",
		},
	},
	"STO": {
		{
			ID:          4162,
			Slug:        "storeum",
			Name:        "Storeum",
			BinanceUSDT: "STOUSDT",
		},
	},
	"YFIM": {
		{
			ID:          7613,
			Slug:        "yfi-mobi",
			Name:        "Yfi.mobi",
			BinanceUSDT: "YFIMUSDT",
		},
	},
	"PCN": {
		{
			ID:          1803,
			Slug:        "peepcoin",
			Name:        "PeepCoin",
			BinanceUSDT: "PCNUSDT",
		},
	},
	"TABOO": {
		{
			ID:          10586,
			Slug:        "taboo-token",
			Name:        "TABOO TOKEN",
			BinanceUSDT: "TABOOUSDT",
		},
	},
	"CROSS": {
		{
			ID:          9442,
			Slug:        "crosspad",
			Name:        "CrossPad",
			BinanceUSDT: "CROSSUSDT",
		},
	},
	"CLO": {
		{
			ID:          2757,
			Slug:        "callisto-network",
			Name:        "Callisto Network",
			BinanceUSDT: "CLOUSDT",
		},
	},
	"ATA": {
		{
			ID:          10188,
			Slug:        "automata-network",
			Name:        "Automata Network",
			BinanceUSDT: "ATAUSDT",
		},
	},
	"KEEP": {
		{
			ID:          5566,
			Slug:        "keep-network",
			Name:        "Keep Network",
			BinanceUSDT: "KEEPUSDT",
		},
	},
	"PROMISE": {
		{
			ID:          10516,
			Slug:        "promise",
			Name:        "Promise",
			BinanceUSDT: "PROMISEUSDT",
		},
	},
	"DIT": {
		{
			ID:          3264,
			Slug:        "digital-insurance-token",
			Name:        "Digital Insurance Token",
			BinanceUSDT: "DITUSDT",
		},
	},
	"PLOT": {
		{
			ID:          7422,
			Slug:        "plotx",
			Name:        "PlotX",
			BinanceUSDT: "PLOTUSDT",
		},
	},
	"SHIELD": {
		{
			ID:          8452,
			Slug:        "shield-protocol",
			Name:        "Shield Protocol",
			BinanceUSDT: "SHIELDUSDT",
		},
	},
	"LINKBULL": {
		{
			ID:          6037,
			Slug:        "3x-long-chainlink-token",
			Name:        "3X Long Chainlink Token",
			BinanceUSDT: "LINKBULLUSDT",
		},
	},
	"WANSUSHI": {
		{
			ID:          8653,
			Slug:        "wansushi",
			Name:        "wanSUSHI",
			BinanceUSDT: "WANSUSHIUSDT",
		},
	},
	"KTS": {
		{
			ID:          4018,
			Slug:        "klimatas",
			Name:        "Klimatas",
			BinanceUSDT: "KTSUSDT",
		},
	},
	"LIKE": {
		{
			ID:          2909,
			Slug:        "likecoin",
			Name:        "LikeCoin",
			BinanceUSDT: "LIKEUSDT",
		},
	},
	"SPACETOAST": {
		{
			ID:          10573,
			Slug:        "spacetoast",
			Name:        "SpaceToast",
			BinanceUSDT: "SPACETOASTUSDT",
		},
	},
	"TASTE": {
		{
			ID:          10329,
			Slug:        "tastenft",
			Name:        "TasteNFT",
			BinanceUSDT: "TASTEUSDT",
		},
	},
	"STPL": {
		{
			ID:          7322,
			Slug:        "stream-protocol",
			Name:        "Stream Protocol",
			BinanceUSDT: "STPLUSDT",
		},
	},
	"JBX": {
		{
			ID:          6704,
			Slug:        "jboxcoin",
			Name:        "JBOX",
			BinanceUSDT: "JBXUSDT",
		},
	},
	"PKT": {
		{
			ID:          2279,
			Slug:        "playkey",
			Name:        "Playkey",
			BinanceUSDT: "PKTUSDT",
		},
	},
	"HITX": {
		{
			ID:          8233,
			Slug:        "hithotx",
			Name:        "Hithotx",
			BinanceUSDT: "HITXUSDT",
		},
	},
	"RED": {
		{
			ID:          2771,
			Slug:        "red",
			Name:        "RED",
			BinanceUSDT: "REDUSDT",
		},
	},
	"SIN": {
		{
			ID:          3514,
			Slug:        "sinovate",
			Name:        "SINOVATE",
			BinanceUSDT: "SINUSDT",
		},
	},
	"INARI": {
		{
			ID:          10809,
			Slug:        "inari",
			Name:        "Inari",
			BinanceUSDT: "INARIUSDT",
		},
	},
	"MSP": {
		{
			ID:          10553,
			Slug:        "moonship-finance",
			Name:        "Moonship Finance",
			BinanceUSDT: "MSPUSDT",
		},
	},
	"CXO": {
		{
			ID:          2490,
			Slug:        "cargox",
			Name:        "CargoX",
			BinanceUSDT: "CXOUSDT",
		},
	},
	"SLME": {
		{
			ID:          8693,
			Slug:        "slime-finance",
			Name:        "Slime Finance",
			BinanceUSDT: "SLMEUSDT",
		},
	},
	"PHOENIX": {
		{
			ID:          10906,
			Slug:        "phoenix-force",
			Name:        "PHOENIX FORCE",
			BinanceUSDT: "PHOENIXUSDT",
		},
	},
	"LEXI": {
		{
			ID:          4549,
			Slug:        "lexit",
			Name:        "LEXIT",
			BinanceUSDT: "LEXIUSDT",
		},
	},
	"WAVAX": {
		{
			ID:          9462,
			Slug:        "wavax",
			Name:        "Wrapped AVAX",
			BinanceUSDT: "WAVAXUSDT",
		},
	},
	"DRG": {
		{
			ID:          2593,
			Slug:        "dragon-coins",
			Name:        "Dragon Coins",
			BinanceUSDT: "DRGUSDT",
		},
	},
	"BULLS": {
		{
			ID:          8360,
			Slug:        "bulls",
			Name:        "BULLS",
			BinanceUSDT: "BULLSUSDT",
		},
	},
	"FIL": {
		{
			ID:          2280,
			Slug:        "filecoin",
			Name:        "Filecoin",
			BinanceUSDT: "FILUSDT",
		},
	},
	"BNBTC": {
		{
			ID:          10647,
			Slug:        "bnbitcoin",
			Name:        "BNbitcoin",
			BinanceUSDT: "BNBTCUSDT",
		},
	},
	"YIELD": {
		{
			ID:          7498,
			Slug:        "yield-protocol",
			Name:        "Yield Protocol",
			BinanceUSDT: "YIELDUSDT",
		},
	},
	"ALH": {
		{
			ID:          8674,
			Slug:        "allohash",
			Name:        "AlloHash",
			BinanceUSDT: "ALHUSDT",
		},
	},
	"UPR": {
		{
			ID:          9694,
			Slug:        "upfire",
			Name:        "Upfire",
			BinanceUSDT: "UPRUSDT",
		},
	},
	"LOON": {
		{
			ID:          8173,
			Slug:        "loon-network",
			Name:        "Loon Network",
			BinanceUSDT: "LOONUSDT",
		},
	},
	"COSMIC": {
		{
			ID:          10521,
			Slug:        "cosmicswap",
			Name:        "CosmicSwap",
			BinanceUSDT: "COSMICUSDT",
		},
	},
	"TRXUP": {
		{
			ID:          7005,
			Slug:        "trxup",
			Name:        "TRXUP",
			BinanceUSDT: "TRXUPUSDT",
		},
	},
	"DIAMONDS": {
		{
			ID:          10924,
			Slug:        "black-diamond",
			Name:        "Black Diamond",
			BinanceUSDT: "DIAMONDSUSDT",
		},
	},
	"UBU": {
		{
			ID:   8787,
			Slug: "ubu-finance",
			Name: "UBU Finance",
		},
		{
			ID:   5074,
			Slug: "ubu",
			Name: "UBU",
		},
	},
	"USDC": {
		{
			ID:          3408,
			Slug:        "usd-coin",
			Name:        "USD Coin",
			BinanceUSDT: "USDCUSDT",
		},
	},
	"LTCR": {
		{
			ID:          1155,
			Slug:        "litecred",
			Name:        "Litecred",
			BinanceUSDT: "LTCRUSDT",
		},
	},
	"MCO": {
		{
			ID:          1776,
			Slug:        "crypto-com",
			Name:        "MCO",
			BinanceUSDT: "MCOUSDT",
		},
	},
	"RAKUC": {
		{
			ID:          10638,
			Slug:        "raku-coin",
			Name:        "Raku Coin",
			BinanceUSDT: "RAKUCUSDT",
		},
	},
	"BSOV": {
		{
			ID:          4306,
			Slug:        "bitcoinsov",
			Name:        "BitcoinSoV",
			BinanceUSDT: "BSOVUSDT",
		},
	},
	"ETHPLO": {
		{
			ID:          4076,
			Slug:        "ethplode",
			Name:        "ETHplode",
			BinanceUSDT: "ETHPLOUSDT",
		},
	},
	"AMC": {
		{
			ID:   11039,
			Slug: "amc-fight-night",
			Name: "AMC FIGHT NIGHT",
		},
		{
			ID:   10992,
			Slug: "amnext",
			Name: "Amnext",
		},
		{
			ID:   8343,
			Slug: "amc-entertainment-holdings-tokenized-stock-ftx",
			Name: "AMC Entertainment Holdings tokenized stock FTX",
		},
	},
	"HEJJ": {
		{
			ID:          9992,
			Slug:        "hedge4-ai",
			Name:        "HEDGE4.Ai",
			BinanceUSDT: "HEJJUSDT",
		},
	},
	"SPIZ": {
		{
			ID:          6626,
			Slug:        "space-iz",
			Name:        "SPACE-iZ",
			BinanceUSDT: "SPIZUSDT",
		},
	},
	"ONE": {
		{
			ID:   3945,
			Slug: "harmony",
			Name: "Harmony",
		},
		{
			ID:   2324,
			Slug: "bigone-token",
			Name: "BigONE Token",
		},
		{
			ID:   10509,
			Slug: "one-token",
			Name: "One Token",
		},
	},
	"IDXM": {
		{
			ID:          2422,
			Slug:        "idex-membership",
			Name:        "IDEX Membership",
			BinanceUSDT: "IDXMUSDT",
		},
	},
	"ANV": {
		{
			ID:          7705,
			Slug:        "aniverse",
			Name:        "ANIVERSE",
			BinanceUSDT: "ANVUSDT",
		},
	},
	"CHIHUA": {
		{
			ID:          9947,
			Slug:        "chihua-token",
			Name:        "Chihua Token",
			BinanceUSDT: "CHIHUAUSDT",
		},
	},
	"DDX": {
		{
			ID:          7228,
			Slug:        "derivadao",
			Name:        "DerivaDAO",
			BinanceUSDT: "DDXUSDT",
		},
	},
	"RISKMOON": {
		{
			ID:          9278,
			Slug:        "riskmoon",
			Name:        "RiskMoon",
			BinanceUSDT: "RISKMOONUSDT",
		},
	},
	"TOKAU": {
		{
			ID:          10951,
			Slug:        "tokyo-au",
			Name:        "Tokyo AU",
			BinanceUSDT: "TOKAUUSDT",
		},
	},
	"BXY": {
		{
			ID:          4646,
			Slug:        "beaxy",
			Name:        "Beaxy",
			BinanceUSDT: "BXYUSDT",
		},
	},
	"S": {
		{
			ID:          3423,
			Slug:        "sharpay",
			Name:        "Sharpay",
			BinanceUSDT: "SUSDT",
		},
	},
	"FYZ": {
		{
			ID:          6674,
			Slug:        "fyooz",
			Name:        "Fyooz",
			BinanceUSDT: "FYZUSDT",
		},
	},
	"FOC": {
		{
			ID:          9259,
			Slug:        "theforce-trade",
			Name:        "TheForce Trade",
			BinanceUSDT: "FOCUSDT",
		},
	},
	"RON": {
		{
			ID:          10921,
			Slug:        "rise-of-nebula",
			Name:        "Rise Of Nebula",
			BinanceUSDT: "RONUSDT",
		},
	},
	"UNIC": {
		{
			ID:          9998,
			Slug:        "unicly",
			Name:        "Unicly",
			BinanceUSDT: "UNICUSDT",
		},
	},
	"ICX": {
		{
			ID:          2099,
			Slug:        "icon",
			Name:        "ICON",
			BinanceUSDT: "ICXUSDT",
		},
	},
	"PCX": {
		{
			ID:          4200,
			Slug:        "chainx",
			Name:        "ChainX",
			BinanceUSDT: "PCXUSDT",
		},
	},
	"KTN": {
		{
			ID:          9110,
			Slug:        "kattana",
			Name:        "Kattana",
			BinanceUSDT: "KTNUSDT",
		},
	},
	"XSR": {
		{
			ID:          4818,
			Slug:        "xensor",
			Name:        "Xensor",
			BinanceUSDT: "XSRUSDT",
		},
	},
	"ECO": {
		{
			ID:          4466,
			Slug:        "ormeus-ecosystem",
			Name:        "Ormeus Ecosystem",
			BinanceUSDT: "ECOUSDT",
		},
	},
	"GUM": {
		{
			ID:          8386,
			Slug:        "gourmet-galaxy",
			Name:        "Gourmet Galaxy",
			BinanceUSDT: "GUMUSDT",
		},
	},
	"SWIRL": {
		{
			ID:          9051,
			Slug:        "swirl-cash",
			Name:        "Swirl Cash",
			BinanceUSDT: "SWIRLUSDT",
		},
	},
	"URAC": {
		{
			ID:          4093,
			Slug:        "uranus",
			Name:        "Uranus",
			BinanceUSDT: "URACUSDT",
		},
	},
	"GDOGE": {
		{
			ID:          10227,
			Slug:        "gdoge-finance",
			Name:        "GDOGE Finance",
			BinanceUSDT: "GDOGEUSDT",
		},
	},
	"SATX": {
		{
			ID:          5430,
			Slug:        "satoexchange-token",
			Name:        "SatoExchange Token",
			BinanceUSDT: "SATXUSDT",
		},
	},
	"TRXBULL": {
		{
			ID:          5378,
			Slug:        "3x-long-trx-token",
			Name:        "3X Long TRX Token",
			BinanceUSDT: "TRXBULLUSDT",
		},
	},
	"XRPUP": {
		{
			ID:          7001,
			Slug:        "xrpup",
			Name:        "XRPUP",
			BinanceUSDT: "XRPUPUSDT",
		},
	},
	"NEBO": {
		{
			ID:          7597,
			Slug:        "csp-dao",
			Name:        "CSP DAO",
			BinanceUSDT: "NEBOUSDT",
		},
	},
	"SC": {
		{
			ID:          1042,
			Slug:        "siacoin",
			Name:        "Siacoin",
			BinanceUSDT: "SCUSDT",
		},
	},
	"PINKPANDA": {
		{
			ID:          10513,
			Slug:        "pinkpanda",
			Name:        "PinkPanda",
			BinanceUSDT: "PINKPANDAUSDT",
		},
	},
	"ZRC": {
		{
			ID:          1726,
			Slug:        "zrcoin",
			Name:        "ZrCoin",
			BinanceUSDT: "ZRCUSDT",
		},
	},
	"XLR": {
		{
			ID:          1606,
			Slug:        "solaris",
			Name:        "Solaris",
			BinanceUSDT: "XLRUSDT",
		},
	},
	"HER": {
		{
			ID:          2754,
			Slug:        "heronode",
			Name:        "HeroNode",
			BinanceUSDT: "HERUSDT",
		},
	},
	"FRA": {
		{
			ID:          4249,
			Slug:        "findora",
			Name:        "Findora",
			BinanceUSDT: "FRAUSDT",
		},
	},
	"LINKETHRSI": {
		{
			ID:          6158,
			Slug:        "link-eth-rsi-ratio-trading-set",
			Name:        "LINK/ETH RSI Ratio Trading Set",
			BinanceUSDT: "LINKETHRSIUSDT",
		},
	},
	"INVE": {
		{
			ID:          3597,
			Slug:        "intervalue",
			Name:        "InterValue",
			BinanceUSDT: "INVEUSDT",
		},
	},
	"ETLT": {
		{
			ID:          7630,
			Slug:        "ethereumlightning",
			Name:        "Ethereum Lightning",
			BinanceUSDT: "ETLTUSDT",
		},
	},
	"MUST": {
		{
			ID:   8294,
			Slug: "cometh",
			Name: "Cometh",
		},
		{
			ID:   9328,
			Slug: "mustangtoken",
			Name: "MustangToken",
		},
	},
	"FAT": {
		{
			ID:   3766,
			Slug: "fatcoin",
			Name: "Fatcoin",
		},
		{
			ID:   8846,
			Slug: "fatfi-protocol",
			Name: "Fatfi Protocol",
		},
	},
	"KIAN": {
		{
			ID:          9265,
			Slug:        "kianite-finance",
			Name:        "Kianite Finance",
			BinanceUSDT: "KIANUSDT",
		},
	},
	"POND": {
		{
			ID:          7497,
			Slug:        "marlin",
			Name:        "Marlin",
			BinanceUSDT: "PONDUSDT",
		},
	},
	"TVK": {
		{
			ID:          8037,
			Slug:        "terra-virtua-kolect",
			Name:        "Terra Virtua Kolect",
			BinanceUSDT: "TVKUSDT",
		},
	},
	"ADAI": {
		{
			ID:          5763,
			Slug:        "aave-dai",
			Name:        "Aave DAI",
			BinanceUSDT: "ADAIUSDT",
		},
	},
	"OLIVE": {
		{
			ID:          10685,
			Slug:        "olive-cash",
			Name:        "Olive.Cash",
			BinanceUSDT: "OLIVEUSDT",
		},
	},
	"FARA": {
		{
			ID:          9530,
			Slug:        "faraland",
			Name:        "FaraLand",
			BinanceUSDT: "FARAUSDT",
		},
	},
	"HUSD": {
		{
			ID:          4779,
			Slug:        "husd",
			Name:        "HUSD",
			BinanceUSDT: "HUSDUSDT",
		},
	},
	"OCP": {
		{
			ID:          6693,
			Slug:        "oc-protocol",
			Name:        "OC Protocol",
			BinanceUSDT: "OCPUSDT",
		},
	},
	"REAP": {
		{
			ID:          7677,
			Slug:        "reapchain",
			Name:        "ReapChain",
			BinanceUSDT: "REAPUSDT",
		},
	},
	"LIBFX": {
		{
			ID:          6714,
			Slug:        "libfx",
			Name:        "Libfx",
			BinanceUSDT: "LIBFXUSDT",
		},
	},
	"CHSB": {
		{
			ID:          2499,
			Slug:        "swissborg",
			Name:        "SwissBorg",
			BinanceUSDT: "CHSBUSDT",
		},
	},
	"DIRTY": {
		{
			ID:          10328,
			Slug:        "dirty-finance",
			Name:        "Dirty Finance",
			BinanceUSDT: "DIRTYUSDT",
		},
	},
	"BCR": {
		{
			ID:          5535,
			Slug:        "bankcoin-reserve",
			Name:        "Bankcoin Reserve",
			BinanceUSDT: "BCRUSDT",
		},
	},
	"ETH2X-FLI": {
		{
			ID:          9789,
			Slug:        "eth-2x-flexible-leverage-index",
			Name:        "ETH 2x Flexible Leverage Index",
			BinanceUSDT: "ETH2X-FLIUSDT",
		},
	},
	"GRUMPY": {
		{
			ID:          8816,
			Slug:        "grumpy-finance",
			Name:        "Grumpy.finance",
			BinanceUSDT: "GRUMPYUSDT",
		},
	},
	"SDOG": {
		{
			ID:          9680,
			Slug:        "small-dogecoin",
			Name:        "Small dogecoin",
			BinanceUSDT: "SDOGUSDT",
		},
	},
	"MKOALA": {
		{
			ID:          10137,
			Slug:        "koala-token",
			Name:        "KOALA TOKEN",
			BinanceUSDT: "MKOALAUSDT",
		},
	},
	"ALCX": {
		{
			ID:          8613,
			Slug:        "alchemix",
			Name:        "Alchemix",
			BinanceUSDT: "ALCXUSDT",
		},
	},
	"SONO": {
		{
			ID:   5420,
			Slug: "sonocoin",
			Name: "SonoCoin",
		},
		{
			ID:   2140,
			Slug: "altcommunity-coin",
			Name: "SONO",
		},
	},
	"WOZX": {
		{
			ID:          7882,
			Slug:        "efforce",
			Name:        "EFFORCE",
			BinanceUSDT: "WOZXUSDT",
		},
	},
	"GPKR": {
		{
			ID:          3184,
			Slug:        "gold-poker",
			Name:        "Gold Poker",
			BinanceUSDT: "GPKRUSDT",
		},
	},
	"YBO": {
		{
			ID:          8341,
			Slug:        "young-boys-fan-token",
			Name:        "Young Boys Fan Token",
			BinanceUSDT: "YBOUSDT",
		},
	},
	"PAL": {
		{
			ID:          8207,
			Slug:        "playandlike",
			Name:        "PlayAndLike",
			BinanceUSDT: "PALUSDT",
		},
	},
	"JGN": {
		{
			ID:          6942,
			Slug:        "juggernaut",
			Name:        "Juggernaut",
			BinanceUSDT: "JGNUSDT",
		},
	},
	"PIE": {
		{
			ID:          6243,
			Slug:        "defipie",
			Name:        "DeFiPie",
			BinanceUSDT: "PIEUSDT",
		},
	},
	"UNCX": {
		{
			ID:          7664,
			Slug:        "uncx",
			Name:        "UniCrypt",
			BinanceUSDT: "UNCXUSDT",
		},
	},
	"DEXM": {
		{
			ID:          8557,
			Slug:        "dexmex",
			Name:        "DexMex",
			BinanceUSDT: "DEXMUSDT",
		},
	},
	"FRST": {
		{
			ID:          1522,
			Slug:        "firstcoin",
			Name:        "FirstCoin",
			BinanceUSDT: "FRSTUSDT",
		},
	},
	"YFIDOWN": {
		{
			ID:          7453,
			Slug:        "yfidown",
			Name:        "YFIDOWN",
			BinanceUSDT: "YFIDOWNUSDT",
		},
	},
	"SISHI": {
		{
			ID:          9862,
			Slug:        "sishi-finance",
			Name:        "Sishi Finance",
			BinanceUSDT: "SISHIUSDT",
		},
	},
	"BBANK": {
		{
			ID:          9839,
			Slug:        "blockbank",
			Name:        "BlockBank",
			BinanceUSDT: "BBANKUSDT",
		},
	},
	"SSN": {
		{
			ID:   9779,
			Slug: "supersonic-finance",
			Name: "Supersonic Finance",
		},
		{
			ID:   5298,
			Slug: "superskynet",
			Name: "SuperSkynet",
		},
	},
	"N3RDz": {
		{
			ID:          8019,
			Slug:        "n3rd-finance",
			Name:        "N3RD Finance",
			BinanceUSDT: "N3RDzUSDT",
		},
	},
	"CHR": {
		{
			ID:          3978,
			Slug:        "chromia",
			Name:        "Chromia",
			BinanceUSDT: "CHRUSDT",
		},
	},
	"MXF": {
		{
			ID:          9093,
			Slug:        "mixty-finance",
			Name:        "Mixty Finance",
			BinanceUSDT: "MXFUSDT",
		},
	},
	"INT": {
		{
			ID:          2399,
			Slug:        "int-chain",
			Name:        "INT",
			BinanceUSDT: "INTUSDT",
		},
	},
	"FLL": {
		{
			ID:          6410,
			Slug:        "feellike",
			Name:        "Feellike",
			BinanceUSDT: "FLLUSDT",
		},
	},
	"VSD": {
		{
			ID:          8775,
			Slug:        "value-set-dollar",
			Name:        "Value Set Dollar",
			BinanceUSDT: "VSDUSDT",
		},
	},
	"AUTO": {
		{
			ID:          8387,
			Slug:        "auto",
			Name:        "Auto",
			BinanceUSDT: "AUTOUSDT",
		},
	},
	"LIMON": {
		{
			ID:          10141,
			Slug:        "limon-group",
			Name:        "LIMON.GROUP",
			BinanceUSDT: "LIMONUSDT",
		},
	},
	"FET": {
		{
			ID:          3773,
			Slug:        "fetch",
			Name:        "Fetch.ai",
			BinanceUSDT: "FETUSDT",
		},
	},
	"MRF": {
		{
			ID:          10287,
			Slug:        "moonradar",
			Name:        "MoonRadar",
			BinanceUSDT: "MRFUSDT",
		},
	},
	"KYL": {
		{
			ID:          8644,
			Slug:        "kylin",
			Name:        "Kylin",
			BinanceUSDT: "KYLUSDT",
		},
	},
	"SWRV": {
		{
			ID:          6901,
			Slug:        "swerve",
			Name:        "Swerve",
			BinanceUSDT: "SWRVUSDT",
		},
	},
	"ATMOS": {
		{
			ID:          1624,
			Slug:        "atmos",
			Name:        "Atmos",
			BinanceUSDT: "ATMOSUSDT",
		},
	},
	"EOSBEAR": {
		{
			ID:          5415,
			Slug:        "3x-short-eos-token",
			Name:        "3x Short EOS Token",
			BinanceUSDT: "EOSBEARUSDT",
		},
	},
	"XHI": {
		{
			ID:          1244,
			Slug:        "hicoin",
			Name:        "HiCoin",
			BinanceUSDT: "XHIUSDT",
		},
	},
	"BKR": {
		{
			ID:          10542,
			Slug:        "bakerdao",
			Name:        "BakerDAO",
			BinanceUSDT: "BKRUSDT",
		},
	},
	"ARTH": {
		{
			ID:          10865,
			Slug:        "arth-new",
			Name:        "ARTH [polygon]",
			BinanceUSDT: "ARTHUSDT",
		},
	},
	"EFK": {
		{
			ID:          6761,
			Slug:        "refork",
			Name:        "ReFork",
			BinanceUSDT: "EFKUSDT",
		},
	},
	"XCH": {
		{
			ID:          9258,
			Slug:        "chia-network",
			Name:        "Chia Network",
			BinanceUSDT: "XCHUSDT",
		},
	},
	"MOONSHOT": {
		{
			ID:          9140,
			Slug:        "moonshot",
			Name:        "Moonshot",
			BinanceUSDT: "MOONSHOTUSDT",
		},
	},
	"IDO": {
		{
			ID:          9365,
			Slug:        "idohunt-app",
			Name:        "IDOHunt app",
			BinanceUSDT: "IDOUSDT",
		},
	},
	"TRX": {
		{
			ID:          1958,
			Slug:        "tron",
			Name:        "TRON",
			BinanceUSDT: "TRXUSDT",
		},
	},
	"FILDA": {
		{
			ID:          8426,
			Slug:        "filda",
			Name:        "Filda",
			BinanceUSDT: "FILDAUSDT",
		},
	},
	"POG": {
		{
			ID:          10658,
			Slug:        "pogcoin",
			Name:        "PogCoin",
			BinanceUSDT: "POGUSDT",
		},
	},
	"QC": {
		{
			ID:          2319,
			Slug:        "qcash",
			Name:        "Qcash",
			BinanceUSDT: "QCUSDT",
		},
	},
	"MLGC": {
		{
			ID:          4892,
			Slug:        "marshal-lion-group-coin",
			Name:        "Marshal Lion Group Coin",
			BinanceUSDT: "MLGCUSDT",
		},
	},
	"SST": {
		{
			ID:          5987,
			Slug:        "simba-storage-token",
			Name:        "SIMBA Storage Token",
			BinanceUSDT: "SSTUSDT",
		},
	},
	"MPS": {
		{
			ID:          5540,
			Slug:        "mt-pelerin",
			Name:        "Mt Pelerin",
			BinanceUSDT: "MPSUSDT",
		},
	},
	"MOONLIGHT": {
		{
			ID:          9848,
			Slug:        "moonlight-token",
			Name:        "Moonlight Token",
			BinanceUSDT: "MOONLIGHTUSDT",
		},
	},
	"LCMS": {
		{
			ID:          9033,
			Slug:        "lcms",
			Name:        "LCMS",
			BinanceUSDT: "LCMSUSDT",
		},
	},
	"PEX": {
		{
			ID:          1209,
			Slug:        "posex",
			Name:        "PosEx",
			BinanceUSDT: "PEXUSDT",
		},
	},
	"ESK": {
		{
			ID:          6435,
			Slug:        "eska",
			Name:        "Eska",
			BinanceUSDT: "ESKUSDT",
		},
	},
	"GANJA": {
		{
			ID:          8983,
			Slug:        "trees-finance",
			Name:        "trees.finance",
			BinanceUSDT: "GANJAUSDT",
		},
	},
	"SHIKO": {
		{
			ID:          10510,
			Slug:        "shikoku-inu",
			Name:        "Shikoku Inu",
			BinanceUSDT: "SHIKOUSDT",
		},
	},
	"BNS": {
		{
			ID:          5989,
			Slug:        "bns-token",
			Name:        "BNS Token",
			BinanceUSDT: "BNSUSDT",
		},
	},
	"KVA": {
		{
			ID:          6487,
			Slug:        "kevacoin",
			Name:        "Kevacoin",
			BinanceUSDT: "KVAUSDT",
		},
	},
	"FACT": {
		{
			ID:          6811,
			Slug:        "fee-active-collateral-token",
			Name:        "Fee Active Collateral Token",
			BinanceUSDT: "FACTUSDT",
		},
	},
	"ADZ": {
		{
			ID:          1136,
			Slug:        "adzcoin",
			Name:        "Adzcoin",
			BinanceUSDT: "ADZUSDT",
		},
	},
	"AIN": {
		{
			ID:   8230,
			Slug: "ai-network",
			Name: "AI Network",
		},
		{
			ID:   8742,
			Slug: "ainori",
			Name: "AINORI",
		},
	},
	"HQT": {
		{
			ID:          3615,
			Slug:        "hyperquant",
			Name:        "HyperQuant",
			BinanceUSDT: "HQTUSDT",
		},
	},
	"CLIQ": {
		{
			ID:          7815,
			Slug:        "deficliq",
			Name:        "DefiCliq",
			BinanceUSDT: "CLIQUSDT",
		},
	},
	"BWC": {
		{
			ID:          10306,
			Slug:        "bongweedcoin",
			Name:        "BongWeedCoin",
			BinanceUSDT: "BWCUSDT",
		},
	},
	"SOCC": {
		{
			ID:          1774,
			Slug:        "socialcoin-socc",
			Name:        "SocialCoin",
			BinanceUSDT: "SOCCUSDT",
		},
	},
	"EOX": {
		{
			ID:          8325,
			Slug:        "eox",
			Name:        "EOX",
			BinanceUSDT: "EOXUSDT",
		},
	},
	"TSL": {
		{
			ID:          2215,
			Slug:        "energo",
			Name:        "Energo",
			BinanceUSDT: "TSLUSDT",
		},
	},
	"GAPT": {
		{
			ID:          9887,
			Slug:        "gaptt",
			Name:        "Gaptt",
			BinanceUSDT: "GAPTUSDT",
		},
	},
	"EVR": {
		{
			ID:          2066,
			Slug:        "everus",
			Name:        "Everus",
			BinanceUSDT: "EVRUSDT",
		},
	},
	"SARCO": {
		{
			ID:          10348,
			Slug:        "sarcophagus",
			Name:        "Sarcophagus",
			BinanceUSDT: "SARCOUSDT",
		},
	},
	"GO": {
		{
			ID:          2861,
			Slug:        "gochain",
			Name:        "GoChain",
			BinanceUSDT: "GOUSDT",
		},
	},
	"SNM": {
		{
			ID:          9931,
			Slug:        "sonm-bep20",
			Name:        "SONM (BEP-20)",
			BinanceUSDT: "SNMUSDT",
		},
	},
	"SHIP": {
		{
			ID:          2579,
			Slug:        "shipchain",
			Name:        "ShipChain",
			BinanceUSDT: "SHIPUSDT",
		},
	},
	"RISE": {
		{
			ID:   1294,
			Slug: "rise",
			Name: "Rise",
		},
		{
			ID:   10548,
			Slug: "everrise",
			Name: "EverRise",
		},
	},
	"AXPR": {
		{
			ID:          2466,
			Slug:        "axpr-token",
			Name:        "AXPR",
			BinanceUSDT: "AXPRUSDT",
		},
	},
	"SPI": {
		{
			ID:          8161,
			Slug:        "shopping",
			Name:        "Shopping",
			BinanceUSDT: "SPIUSDT",
		},
	},
	"DFY": {
		{
			ID:          9179,
			Slug:        "defi-for-you",
			Name:        "Defi For You",
			BinanceUSDT: "DFYUSDT",
		},
	},
	"BBGC": {
		{
			ID:          3832,
			Slug:        "big-bang-game-coin",
			Name:        "Big Bang Game Coin",
			BinanceUSDT: "BBGCUSDT",
		},
	},
	"MERI": {
		{
			ID:          3915,
			Slug:        "merebel",
			Name:        "Merebel",
			BinanceUSDT: "MERIUSDT",
		},
	},
	"SDX": {
		{
			ID:          9420,
			Slug:        "swapdex",
			Name:        "SwapDEX",
			BinanceUSDT: "SDXUSDT",
		},
	},
	"DOTX": {
		{
			ID:          6796,
			Slug:        "deli-of-thrones",
			Name:        "DeFi of Thrones",
			BinanceUSDT: "DOTXUSDT",
		},
	},
	"KIRO": {
		{
			ID:          7276,
			Slug:        "kirobo",
			Name:        "Kirobo",
			BinanceUSDT: "KIROUSDT",
		},
	},
	"YOOSHI": {
		{
			ID:          9892,
			Slug:        "yooshi",
			Name:        "YooShi",
			BinanceUSDT: "YOOSHIUSDT",
		},
	},
	"CAPT": {
		{
			ID:          10226,
			Slug:        "captain",
			Name:        "Captain",
			BinanceUSDT: "CAPTUSDT",
		},
	},
	"SPD": {
		{
			ID:   2616,
			Slug: "stipend",
			Name: "Stipend",
		},
		{
			ID:   2828,
			Slug: "spindle",
			Name: "SPINDLE",
		},
	},
	"XTZDOWN": {
		{
			ID:          7008,
			Slug:        "xtzdown",
			Name:        "XTZDOWN",
			BinanceUSDT: "XTZDOWNUSDT",
		},
	},
	"KNDC": {
		{
			ID:          2890,
			Slug:        "kanadecoin",
			Name:        "KanadeCoin",
			BinanceUSDT: "KNDCUSDT",
		},
	},
	"FUZE": {
		{
			ID:          4104,
			Slug:        "fuze-token",
			Name:        "FUZE Token",
			BinanceUSDT: "FUZEUSDT",
		},
	},
	"BCV": {
		{
			ID:          3066,
			Slug:        "bitcapitalvendor",
			Name:        "BitCapitalVendor",
			BinanceUSDT: "BCVUSDT",
		},
	},
	"LF": {
		{
			ID:          9160,
			Slug:        "linkflow-finance",
			Name:        "Linkflow Finance",
			BinanceUSDT: "LFUSDT",
		},
	},
	"R3FI": {
		{
			ID:          8313,
			Slug:        "r3fi-finance",
			Name:        "Recharge Finance",
			BinanceUSDT: "R3FIUSDT",
		},
	},
	"ELEPHANT": {
		{
			ID:          9936,
			Slug:        "elephant-money",
			Name:        "Elephant Money",
			BinanceUSDT: "ELEPHANTUSDT",
		},
	},
	"GOL": {
		{
			ID:          9670,
			Slug:        "gogolcoin",
			Name:        "GogolCoin",
			BinanceUSDT: "GOLUSDT",
		},
	},
	"JEJUDOGE": {
		{
			ID:   10816,
			Slug: "jejudogebsc",
			Name: "JejuDogeBSC",
		},
		{
			ID:   10584,
			Slug: "jejudoge",
			Name: "Jejudoge",
		},
	},
	"RYOSHI": {
		{
			ID:          10817,
			Slug:        "ryoshi-token",
			Name:        "Ryoshi Token",
			BinanceUSDT: "RYOSHIUSDT",
		},
	},
	"OIN": {
		{
			ID:          6870,
			Slug:        "oin-finance",
			Name:        "OIN Finance",
			BinanceUSDT: "OINUSDT",
		},
	},
	"GEEQ": {
		{
			ID:          6194,
			Slug:        "geeq",
			Name:        "Geeq",
			BinanceUSDT: "GEEQUSDT",
		},
	},
	"MT": {
		{
			ID:          2712,
			Slug:        "mytoken",
			Name:        "MyToken",
			BinanceUSDT: "MTUSDT",
		},
	},
	"ZTB": {
		{
			ID:          9372,
			Slug:        "ztb",
			Name:        "ZTB",
			BinanceUSDT: "ZTBUSDT",
		},
	},
	"SVR": {
		{
			ID:          4594,
			Slug:        "sovranocoin",
			Name:        "SovranoCoin",
			BinanceUSDT: "SVRUSDT",
		},
	},
	"GTC": {
		{
			ID:   10052,
			Slug: "gitcoin",
			Name: "Gitcoin",
		},
		{
			ID:   2336,
			Slug: "game",
			Name: "Game.com",
		},
		{
			ID:   9676,
			Slug: "gastrocoin",
			Name: "Gastrocoin",
		},
	},
	"VIDYX": {
		{
			ID:          8182,
			Slug:        "vidyx",
			Name:        "VidyX",
			BinanceUSDT: "VIDYXUSDT",
		},
	},
	"CNT": {
		{
			ID:   1546,
			Slug: "centurion",
			Name: "Centurion",
		},
		{
			ID:   9747,
			Slug: "cryption-network",
			Name: "Cryption Network",
		},
	},
	"ZOE": {
		{
			ID:          10959,
			Slug:        "zoe-cash",
			Name:        "Zoe Cash",
			BinanceUSDT: "ZOEUSDT",
		},
	},
	"KEBAB": {
		{
			ID:          8334,
			Slug:        "kebab-token",
			Name:        "Kebab Token",
			BinanceUSDT: "KEBABUSDT",
		},
	},
	"ZFL": {
		{
			ID:          5247,
			Slug:        "zuflo-coin",
			Name:        "Zuflo Coin",
			BinanceUSDT: "ZFLUSDT",
		},
	},
	"FILDOWN": {
		{
			ID:          8051,
			Slug:        "fildown",
			Name:        "FILDOWN",
			BinanceUSDT: "FILDOWNUSDT",
		},
	},
	"VAIN": {
		{
			ID:          10761,
			Slug:        "vain",
			Name:        "Vain",
			BinanceUSDT: "VAINUSDT",
		},
	},
	"BOGE": {
		{
			ID:          10400,
			Slug:        "bogecoin",
			Name:        "Bogecoin",
			BinanceUSDT: "BOGEUSDT",
		},
	},
	"VGX": {
		{
			ID:          1817,
			Slug:        "voyager-token",
			Name:        "Voyager Token",
			BinanceUSDT: "VGXUSDT",
		},
	},
	"LAYER": {
		{
			ID:          6638,
			Slug:        "unilayer",
			Name:        "UniLayer",
			BinanceUSDT: "LAYERUSDT",
		},
	},
	"LINKUP": {
		{
			ID:          7011,
			Slug:        "linkup",
			Name:        "LINKUP",
			BinanceUSDT: "LINKUPUSDT",
		},
	},
	"DMUSK": {
		{
			ID:          10018,
			Slug:        "dragonmusk",
			Name:        "Dragonmusk",
			BinanceUSDT: "DMUSKUSDT",
		},
	},
	"UDOKI": {
		{
			ID:          9474,
			Slug:        "unicly-doki-doki-collection",
			Name:        "Unicly Doki Doki Collection",
			BinanceUSDT: "UDOKIUSDT",
		},
	},
	"PKG": {
		{
			ID:          3106,
			Slug:        "pkg-token",
			Name:        "PKG Token",
			BinanceUSDT: "PKGUSDT",
		},
	},
	"DOGDEFI": {
		{
			ID:          7511,
			Slug:        "dogdeficoin",
			Name:        "DogDeFiCoin",
			BinanceUSDT: "DOGDEFIUSDT",
		},
	},
	"TSR": {
		{
			ID:          4806,
			Slug:        "tesra",
			Name:        "Tesra",
			BinanceUSDT: "TSRUSDT",
		},
	},
	"ELONONE": {
		{
			ID:          10166,
			Slug:        "astroelon",
			Name:        "AstroElon",
			BinanceUSDT: "ELONONEUSDT",
		},
	},
	"NMP": {
		{
			ID:          5703,
			Slug:        "neuromorphic-io",
			Name:        "Neuromorphic.io",
			BinanceUSDT: "NMPUSDT",
		},
	},
	"BTAP": {
		{
			ID:          9060,
			Slug:        "bta-protocol",
			Name:        "BTA Protocol",
			BinanceUSDT: "BTAPUSDT",
		},
	},
	"RAI": {
		{
			ID:   9144,
			Slug: "rai-finance",
			Name: "RAI Finance",
		},
		{
			ID:   8525,
			Slug: "rai",
			Name: "Rai Reflex Index",
		},
	},
	"AHF": {
		{
			ID:          7241,
			Slug:        "americanhorror-finance",
			Name:        "AmericanHorror.Finance",
			BinanceUSDT: "AHFUSDT",
		},
	},
	"PAX": {
		{
			ID:          3330,
			Slug:        "paxos-standard",
			Name:        "Paxos Standard",
			BinanceUSDT: "PAXUSDT",
		},
	},
	"IMP": {
		{
			ID:          3271,
			Slug:        "ether-kingdoms-token",
			Name:        "Ether Kingdoms Token",
			BinanceUSDT: "IMPUSDT",
		},
	},
	"ADI": {
		{
			ID:          2660,
			Slug:        "aditus",
			Name:        "Aditus",
			BinanceUSDT: "ADIUSDT",
		},
	},
	"BON": {
		{
			ID:          2256,
			Slug:        "bonpay",
			Name:        "Bonpay",
			BinanceUSDT: "BONUSDT",
		},
	},
	"XBC": {
		{
			ID:   293,
			Slug: "bitcoin-plus",
			Name: "Bitcoin Plus",
		},
		{
			ID:   9585,
			Slug: "xbn-community-token",
			Name: "XBN Community Token",
		},
	},
	"METIS": {
		{
			ID:          9640,
			Slug:        "metisdao",
			Name:        "Metis",
			BinanceUSDT: "METISUSDT",
		},
	},
	"erowan": {
		{
			ID:          8541,
			Slug:        "sifchain",
			Name:        "SifChain",
			BinanceUSDT: "erowanUSDT",
		},
	},
	"CHG": {
		{
			ID:          5400,
			Slug:        "charg-coin",
			Name:        "Charg Coin",
			BinanceUSDT: "CHGUSDT",
		},
	},
	"MST": {
		{
			ID:          1396,
			Slug:        "mustangcoin",
			Name:        "MustangCoin",
			BinanceUSDT: "MSTUSDT",
		},
	},
	"SACT": {
		{
			ID:          8753,
			Slug:        "srnartgallery",
			Name:        "srnArt Gallery",
			BinanceUSDT: "SACTUSDT",
		},
	},
	"ZOC": {
		{
			ID:          4546,
			Slug:        "01coin",
			Name:        "01coin",
			BinanceUSDT: "ZOCUSDT",
		},
	},
	"NANO": {
		{
			ID:          1567,
			Slug:        "nano",
			Name:        "Nano",
			BinanceUSDT: "NANOUSDT",
		},
	},
	"TROLL": {
		{
			ID:          638,
			Slug:        "trollcoin",
			Name:        "Trollcoin",
			BinanceUSDT: "TROLLUSDT",
		},
	},
	"ADP": {
		{
			ID:          8044,
			Slug:        "adappter-token",
			Name:        "Adappter Token",
			BinanceUSDT: "ADPUSDT",
		},
	},
	"EVZ": {
		{
			ID:          6430,
			Slug:        "electric-vehicle-zone",
			Name:        "Electric Vehicle Zone",
			BinanceUSDT: "EVZUSDT",
		},
	},
	"LABX": {
		{
			ID:          3751,
			Slug:        "stakinglab",
			Name:        "Stakinglab",
			BinanceUSDT: "LABXUSDT",
		},
	},
	"PWC": {
		{
			ID:          8750,
			Slug:        "prime-whiterock-company",
			Name:        "Prime Whiterock Company",
			BinanceUSDT: "PWCUSDT",
		},
	},
	"LOGE": {
		{
			ID:          10156,
			Slug:        "lunadoge",
			Name:        "LunaDoge",
			BinanceUSDT: "LOGEUSDT",
		},
	},
	"XRGE": {
		{
			ID:          9106,
			Slug:        "rougecoin",
			Name:        "RougeCoin",
			BinanceUSDT: "XRGEUSDT",
		},
	},
	"FTXT": {
		{
			ID:          3219,
			Slug:        "futurax",
			Name:        "FUTURAX",
			BinanceUSDT: "FTXTUSDT",
		},
	},
	"SNY": {
		{
			ID:          9447,
			Slug:        "synthetify",
			Name:        "Synthetify",
			BinanceUSDT: "SNYUSDT",
		},
	},
	"AAC": {
		{
			ID:          2438,
			Slug:        "acute-angle-cloud",
			Name:        "Acute Angle Cloud",
			BinanceUSDT: "AACUSDT",
		},
	},
	"BNA": {
		{
			ID:          5792,
			Slug:        "bananatok",
			Name:        "Bananatok",
			BinanceUSDT: "BNAUSDT",
		},
	},
	"GRT": {
		{
			ID:   6719,
			Slug: "the-graph",
			Name: "The Graph",
		},
		{
			ID:   5711,
			Slug: "golden-ratio-token",
			Name: "Golden Ratio Token",
		},
	},
	"INVEST": {
		{
			ID:          10468,
			Slug:        "investdex",
			Name:        "InvestDex",
			BinanceUSDT: "INVESTUSDT",
		},
	},
	"FSN": {
		{
			ID:          2530,
			Slug:        "fusion",
			Name:        "Fusion",
			BinanceUSDT: "FSNUSDT",
		},
	},
	"EVIL": {
		{
			ID:          1165,
			Slug:        "evil-coin",
			Name:        "Evil Coin",
			BinanceUSDT: "EVILUSDT",
		},
	},
	"ORT": {
		{
			ID:          10756,
			Slug:        "omni-real-estate-token",
			Name:        "Omni Real Estate Token",
			BinanceUSDT: "ORTUSDT",
		},
	},
	"RSTR": {
		{
			ID:          3407,
			Slug:        "ondori",
			Name:        "Ondori",
			BinanceUSDT: "RSTRUSDT",
		},
	},
	"ROBBIN": {
		{
			ID:          10883,
			Slug:        "robbin-hood",
			Name:        "ROBBIN HOOD",
			BinanceUSDT: "ROBBINUSDT",
		},
	},
	"RPL": {
		{
			ID:          2943,
			Slug:        "rocket-pool",
			Name:        "Rocket Pool",
			BinanceUSDT: "RPLUSDT",
		},
	},
	"UDO": {
		{
			ID:          8679,
			Slug:        "unido",
			Name:        "Unido EP",
			BinanceUSDT: "UDOUSDT",
		},
	},
	"HCA": {
		{
			ID:          4994,
			Slug:        "harcomia",
			Name:        "Harcomia",
			BinanceUSDT: "HCAUSDT",
		},
	},
	"PNG": {
		{
			ID:          8422,
			Slug:        "pangolin",
			Name:        "Pangolin",
			BinanceUSDT: "PNGUSDT",
		},
	},
	"PTOY": {
		{
			ID:          1708,
			Slug:        "patientory",
			Name:        "Patientory",
			BinanceUSDT: "PTOYUSDT",
		},
	},
	"HBO": {
		{
			ID:          8528,
			Slug:        "hashbridge-oracle",
			Name:        "HashBridge Oracle",
			BinanceUSDT: "HBOUSDT",
		},
	},
	"MAAAVE": {
		{
			ID:          8920,
			Slug:        "matic-aave-aave",
			Name:        "Matic Aave Interest Bearing AAVE",
			BinanceUSDT: "MAAAVEUSDT",
		},
	},
	"HBD": {
		{
			ID:          5375,
			Slug:        "hive-dollar",
			Name:        "Hive Dollar",
			BinanceUSDT: "HBDUSDT",
		},
	},
	"TACO": {
		{
			ID:   8632,
			Slug: "taco-finance",
			Name: "Taco Finance",
		},
		{
			ID:   6575,
			Slug: "tacos",
			Name: "Tacos",
		},
		{
			ID:   8110,
			Slug: "tacoswap",
			Name: "Tacoswap",
		},
	},
	"TRIAS": {
		{
			ID:          8936,
			Slug:        "trias-token",
			Name:        "Trias Token (new)",
			BinanceUSDT: "TRIASUSDT",
		},
	},
	"iOWN": {
		{
			ID:          5337,
			Slug:        "iown-token",
			Name:        "iOWN Token",
			BinanceUSDT: "iOWNUSDT",
		},
	},
	"DSR": {
		{
			ID:          2148,
			Slug:        "desire",
			Name:        "Desire",
			BinanceUSDT: "DSRUSDT",
		},
	},
	"UGOTCHI": {
		{
			ID:          9472,
			Slug:        "unicly-aavegotchi-astronauts-collection",
			Name:        "Unicly Aavegotchi Astronauts Collection",
			BinanceUSDT: "UGOTCHIUSDT",
		},
	},
	"KMW": {
		{
			ID:          5282,
			Slug:        "kepler-network",
			Name:        "Kepler Network",
			BinanceUSDT: "KMWUSDT",
		},
	},
	"SURE": {
		{
			ID:          5113,
			Slug:        "insure",
			Name:        "inSure DeFi",
			BinanceUSDT: "SUREUSDT",
		},
	},
	"HEAT": {
		{
			ID:          1308,
			Slug:        "heat-ledger",
			Name:        "HEAT",
			BinanceUSDT: "HEATUSDT",
		},
	},
	"KPC": {
		{
			ID:          8573,
			Slug:        "koloop-basic",
			Name:        "Koloop Basic",
			BinanceUSDT: "KPCUSDT",
		},
	},
	"CBD": {
		{
			ID:   10098,
			Slug: "greenheart-cbd",
			Name: "Greenheart CBD",
		},
		{
			ID:   10659,
			Slug: "cbd-coin",
			Name: "CBD Coin",
		},
	},
	"GMEE": {
		{
			ID:          9103,
			Slug:        "gamee",
			Name:        "GAMEE",
			BinanceUSDT: "GMEEUSDT",
		},
	},
	"ETG": {
		{
			ID:          2074,
			Slug:        "ethereum-gold",
			Name:        "Ethereum Gold",
			BinanceUSDT: "ETGUSDT",
		},
	},
	"TDX": {
		{
			ID:          2542,
			Slug:        "tidex-token",
			Name:        "Tidex Token",
			BinanceUSDT: "TDXUSDT",
		},
	},
	"XHV": {
		{
			ID:          2662,
			Slug:        "haven-protocol",
			Name:        "Haven Protocol",
			BinanceUSDT: "XHVUSDT",
		},
	},
	"EDDA": {
		{
			ID:          8789,
			Slug:        "eddaswap",
			Name:        "EDDASwap",
			BinanceUSDT: "EDDAUSDT",
		},
	},
	"FECLIPSE": {
		{
			ID:          9185,
			Slug:        "faireclipse",
			Name:        "FairEclipse",
			BinanceUSDT: "FECLIPSEUSDT",
		},
	},
	"NASH": {
		{
			ID:          3867,
			Slug:        "neoworld-cash",
			Name:        "NeoWorld Cash",
			BinanceUSDT: "NASHUSDT",
		},
	},
	"PMEER": {
		{
			ID:          5144,
			Slug:        "qitmeer",
			Name:        "Qitmeer",
			BinanceUSDT: "PMEERUSDT",
		},
	},
	"HPB": {
		{
			ID:          2345,
			Slug:        "high-performance-blockchain",
			Name:        "High Performance Blockchain",
			BinanceUSDT: "HPBUSDT",
		},
	},
	"FOTO": {
		{
			ID:          8729,
			Slug:        "unique-photo",
			Name:        "Unique Photo",
			BinanceUSDT: "FOTOUSDT",
		},
	},
	"EDAO": {
		{
			ID:          10258,
			Slug:        "elondoge-dao",
			Name:        "ElonDoge DAO",
			BinanceUSDT: "EDAOUSDT",
		},
	},
	"SSS": {
		{
			ID:          5607,
			Slug:        "simple-software-solutions",
			Name:        "Simple Software Solutions",
			BinanceUSDT: "SSSUSDT",
		},
	},
	"PFI": {
		{
			ID:   8215,
			Slug: "primefinance",
			Name: "PrimeFinance",
		},
		{
			ID:   8374,
			Slug: "protocol-finance",
			Name: "protocol finance",
		},
	},
	"PERI": {
		{
			ID:          9550,
			Slug:        "peri-finance",
			Name:        "PERI Finance",
			BinanceUSDT: "PERIUSDT",
		},
	},
	"BOTX": {
		{
			ID:          3873,
			Slug:        "botxcoin",
			Name:        "botXcoin",
			BinanceUSDT: "BOTXUSDT",
		},
	},
	"NUX": {
		{
			ID:          8458,
			Slug:        "peanut",
			Name:        "Peanut",
			BinanceUSDT: "NUXUSDT",
		},
	},
	"UPI": {
		{
			ID:          5086,
			Slug:        "pawtocol",
			Name:        "Pawtocol",
			BinanceUSDT: "UPIUSDT",
		},
	},
	"STC": {
		{
			ID:          5966,
			Slug:        "student-coin",
			Name:        "Student Coin",
			BinanceUSDT: "STCUSDT",
		},
	},
	"SWFTC": {
		{
			ID:          2341,
			Slug:        "swftcoin",
			Name:        "SwftCoin",
			BinanceUSDT: "SWFTCUSDT",
		},
	},
	"CPHR": {
		{
			ID:          10978,
			Slug:        "polkacipher",
			Name:        "PolkaCipher",
			BinanceUSDT: "CPHRUSDT",
		},
	},
	"NET": {
		{
			ID:          3788,
			Slug:        "next",
			Name:        "NEXT",
			BinanceUSDT: "NETUSDT",
		},
	},
	"PST": {
		{
			ID:          1930,
			Slug:        "primas",
			Name:        "Primas",
			BinanceUSDT: "PSTUSDT",
		},
	},
	"STEEL": {
		{
			ID:          10639,
			Slug:        "steel",
			Name:        "Steel",
			BinanceUSDT: "STEELUSDT",
		},
	},
	"SNX": {
		{
			ID:          2586,
			Slug:        "synthetix-network-token",
			Name:        "Synthetix",
			BinanceUSDT: "SNXUSDT",
		},
	},
	"ICHI": {
		{
			ID:          7726,
			Slug:        "ichi",
			Name:        "ICHI",
			BinanceUSDT: "ICHIUSDT",
		},
	},
	"DEAL": {
		{
			ID:          3439,
			Slug:        "idealcash",
			Name:        "iDealCash",
			BinanceUSDT: "DEALUSDT",
		},
	},
	"XVG": {
		{
			ID:          693,
			Slug:        "verge",
			Name:        "Verge",
			BinanceUSDT: "XVGUSDT",
		},
	},
	"TOZ": {
		{
			ID:          8522,
			Slug:        "tozex",
			Name:        "TOZEX",
			BinanceUSDT: "TOZUSDT",
		},
	},
	"BIOT": {
		{
			ID:          8034,
			Slug:        "biopassport-token",
			Name:        "BioPassport Token",
			BinanceUSDT: "BIOTUSDT",
		},
	},
	"KOBE": {
		{
			ID:          8189,
			Slug:        "shabu-shabu-finance",
			Name:        "Shabu Shabu Finance",
			BinanceUSDT: "KOBEUSDT",
		},
	},
	"NVT": {
		{
			ID:          5906,
			Slug:        "nervenetwork",
			Name:        "NerveNetwork",
			BinanceUSDT: "NVTUSDT",
		},
	},
	"UP": {
		{
			ID:   2597,
			Slug: "uptoken",
			Name: "UpToken",
		},
		{
			ID:   7077,
			Slug: "unifi-protocol",
			Name: "UniFi Protocol",
		},
	},
	"LMT": {
		{
			ID:          9482,
			Slug:        "lympo-market-token",
			Name:        "Lympo Market Token",
			BinanceUSDT: "LMTUSDT",
		},
	},
	"ETHV": {
		{
			ID:   6409,
			Slug: "ethverse",
			Name: "Ethverse",
		},
		{
			ID:   6617,
			Slug: "ethereum-vault",
			Name: "Ethereum Vault",
		},
	},
	"EBST": {
		{
			ID:          1704,
			Slug:        "eboostcoin",
			Name:        "eBoost",
			BinanceUSDT: "EBSTUSDT",
		},
	},
	"YFT": {
		{
			ID:   6247,
			Slug: "yield-farming-token",
			Name: "Yield Farming Token",
		},
		{
			ID:   7156,
			Slug: "toshify-finance",
			Name: "Toshify.finance",
		},
	},
	"RFOX": {
		{
			ID:          7654,
			Slug:        "redfox-labs",
			Name:        "RedFOX Labs",
			BinanceUSDT: "RFOXUSDT",
		},
	},
	"8PAY": {
		{
			ID:          9046,
			Slug:        "8pay",
			Name:        "8PAY",
			BinanceUSDT: "8PAYUSDT",
		},
	},
	"ASTA": {
		{
			ID:          6375,
			Slug:        "asta",
			Name:        "ASTA",
			BinanceUSDT: "ASTAUSDT",
		},
	},
	"CAVA": {
		{
			ID:          10511,
			Slug:        "cavapoo",
			Name:        "Cavapoo",
			BinanceUSDT: "CAVAUSDT",
		},
	},
	"MEM": {
		{
			ID:          10461,
			Slug:        "meme-coin",
			Name:        "Memecoin",
			BinanceUSDT: "MEMUSDT",
		},
	},
	"TKING": {
		{
			ID:          9854,
			Slug:        "tiger-king",
			Name:        "Tiger King",
			BinanceUSDT: "TKINGUSDT",
		},
	},
	"JULB": {
		{
			ID:          8812,
			Slug:        "justliquidity-binance",
			Name:        "JustLiquidity Binance",
			BinanceUSDT: "JULBUSDT",
		},
	},
	"LUA": {
		{
			ID:          7216,
			Slug:        "lua-token",
			Name:        "LuaSwap",
			BinanceUSDT: "LUAUSDT",
		},
	},
	"KNC": {
		{
			ID:          1982,
			Slug:        "kyber-network-crystal-legacy",
			Name:        "Kyber Network Crystal Legacy",
			BinanceUSDT: "KNCUSDT",
		},
	},
	"JFI": {
		{
			ID:          6898,
			Slug:        "jackpool-finance",
			Name:        "JackPool.finance",
			BinanceUSDT: "JFIUSDT",
		},
	},
	"DUN": {
		{
			ID:          5160,
			Slug:        "dune-network",
			Name:        "Dune Network",
			BinanceUSDT: "DUNUSDT",
		},
	},
	"PUSSY": {
		{
			ID:          9639,
			Slug:        "pussy-financial",
			Name:        "Pussy Financial",
			BinanceUSDT: "PUSSYUSDT",
		},
	},
	"DSD": {
		{
			ID:          8106,
			Slug:        "dynamic-set-dollar",
			Name:        "Dynamic Set Dollar",
			BinanceUSDT: "DSDUSDT",
		},
	},
	"POWR": {
		{
			ID:          2132,
			Slug:        "power-ledger",
			Name:        "Power Ledger",
			BinanceUSDT: "POWRUSDT",
		},
	},
	"CLAM": {
		{
			ID:          460,
			Slug:        "clams",
			Name:        "Clams",
			BinanceUSDT: "CLAMUSDT",
		},
	},
	"COPE": {
		{
			ID:          9015,
			Slug:        "cope",
			Name:        "Cope",
			BinanceUSDT: "COPEUSDT",
		},
	},
	"WLT": {
		{
			ID:          8721,
			Slug:        "wealthlocks",
			Name:        "Wealthlocks",
			BinanceUSDT: "WLTUSDT",
		},
	},
	"INFS": {
		{
			ID:          4903,
			Slug:        "infinity-esaham",
			Name:        "Infinity Esaham",
			BinanceUSDT: "INFSUSDT",
		},
	},
	"WTC": {
		{
			ID:          1925,
			Slug:        "waltonchain",
			Name:        "Waltonchain",
			BinanceUSDT: "WTCUSDT",
		},
	},
	"MKAT": {
		{
			ID:          10691,
			Slug:        "moonkat",
			Name:        "MoonKat",
			BinanceUSDT: "MKATUSDT",
		},
	},
	"IT": {
		{
			ID:          6473,
			Slug:        "idcm-token",
			Name:        "IDCM Token",
			BinanceUSDT: "ITUSDT",
		},
	},
	"GAT": {
		{
			ID:   6213,
			Slug: "global-aex-token",
			Name: "Global AEX Token",
		},
		{
			ID:   10255,
			Slug: "alchemytoys",
			Name: "AlchemyToys",
		},
	},
	"MSD": {
		{
			ID:          2008,
			Slug:        "msd",
			Name:        "MSD",
			BinanceUSDT: "MSDUSDT",
		},
	},
	"QNT": {
		{
			ID:          3155,
			Slug:        "quant",
			Name:        "Quant",
			BinanceUSDT: "QNTUSDT",
		},
	},
	"LOCC": {
		{
			ID:          9907,
			Slug:        "low-orbit-crypto-cannon",
			Name:        "Low Orbit Crypto Cannon",
			BinanceUSDT: "LOCCUSDT",
		},
	},
	"GLCH": {
		{
			ID:          8236,
			Slug:        "glitch",
			Name:        "Glitch",
			BinanceUSDT: "GLCHUSDT",
		},
	},
	"A": {
		{
			ID:          3869,
			Slug:        "alpha-token",
			Name:        "Alpha Token",
			BinanceUSDT: "AUSDT",
		},
	},
	"YEE": {
		{
			ID:          2437,
			Slug:        "yee",
			Name:        "YEE",
			BinanceUSDT: "YEEUSDT",
		},
	},
	"RABBIT": {
		{
			ID:          10178,
			Slug:        "rabbit-finance",
			Name:        "Rabbit Finance",
			BinanceUSDT: "RABBITUSDT",
		},
	},
	"ZLP": {
		{
			ID:          7611,
			Slug:        "zuplo",
			Name:        "Zuplo",
			BinanceUSDT: "ZLPUSDT",
		},
	},
	"VTX": {
		{
			ID:          8661,
			Slug:        "vortex-defi",
			Name:        "Vortex Defi",
			BinanceUSDT: "VTXUSDT",
		},
	},
	"YFMS": {
		{
			ID:          7093,
			Slug:        "yfmoonshot",
			Name:        "YFMoonshot",
			BinanceUSDT: "YFMSUSDT",
		},
	},
	"ENTS": {
		{
			ID:          3433,
			Slug:        "eunomia",
			Name:        "EUNOMIA",
			BinanceUSDT: "ENTSUSDT",
		},
	},
	"XWP": {
		{
			ID:          3878,
			Slug:        "swap",
			Name:        "Swap",
			BinanceUSDT: "XWPUSDT",
		},
	},
	"BTW": {
		{
			ID:          2489,
			Slug:        "bitwhite",
			Name:        "BitWhite",
			BinanceUSDT: "BTWUSDT",
		},
	},
	"FTO": {
		{
			ID:          2846,
			Slug:        "futurocoin",
			Name:        "FuturoCoin",
			BinanceUSDT: "FTOUSDT",
		},
	},
	"DTRC": {
		{
			ID:          2752,
			Slug:        "datarius-credit",
			Name:        "Datarius Credit",
			BinanceUSDT: "DTRCUSDT",
		},
	},
	"ICNQ": {
		{
			ID:          3431,
			Slug:        "iconic-token",
			Name:        "Iconic Token",
			BinanceUSDT: "ICNQUSDT",
		},
	},
	"vXVS": {
		{
			ID:          7960,
			Slug:        "venus-xvs",
			Name:        "Venus XVS",
			BinanceUSDT: "vXVSUSDT",
		},
	},
	"INK": {
		{
			ID:          2209,
			Slug:        "ink",
			Name:        "Ink",
			BinanceUSDT: "INKUSDT",
		},
	},
	"ZCX": {
		{
			ID:          9263,
			Slug:        "unizen",
			Name:        "Unizen",
			BinanceUSDT: "ZCXUSDT",
		},
	},
	"GBPU": {
		{
			ID:          6906,
			Slug:        "upper-pound",
			Name:        "Upper Pound",
			BinanceUSDT: "GBPUUSDT",
		},
	},
	"SLF": {
		{
			ID:          9684,
			Slug:        "solarfare",
			Name:        "Solarfare",
			BinanceUSDT: "SLFUSDT",
		},
	},
	"PTERIA": {
		{
			ID:          7564,
			Slug:        "pteria",
			Name:        "Pteria",
			BinanceUSDT: "PTERIAUSDT",
		},
	},
	"ZABAKU": {
		{
			ID:          10032,
			Slug:        "zabaku-inu",
			Name:        "ZABAKU INU",
			BinanceUSDT: "ZABAKUUSDT",
		},
	},
	"TYC": {
		{
			ID:          9698,
			Slug:        "tycoon",
			Name:        "Tycoon",
			BinanceUSDT: "TYCUSDT",
		},
	},
	"SEPA": {
		{
			ID:          9163,
			Slug:        "secure-pad-token",
			Name:        "Secure Pad",
			BinanceUSDT: "SEPAUSDT",
		},
	},
	"BIR": {
		{
			ID:          3285,
			Slug:        "birake",
			Name:        "Birake",
			BinanceUSDT: "BIRUSDT",
		},
	},
	"PUMP": {
		{
			ID:          10649,
			Slug:        "moonpump",
			Name:        "MoonPump",
			BinanceUSDT: "PUMPUSDT",
		},
	},
	"AGLT": {
		{
			ID:          3427,
			Slug:        "agrolot",
			Name:        "Agrolot",
			BinanceUSDT: "AGLTUSDT",
		},
	},
	"CIFI": {
		{
			ID:          10129,
			Slug:        "citizen-finance",
			Name:        "Citizen Finance",
			BinanceUSDT: "CIFIUSDT",
		},
	},
	"DEFO": {
		{
			ID:          7720,
			Slug:        "defhold",
			Name:        "DefHold",
			BinanceUSDT: "DEFOUSDT",
		},
	},
	"MYST": {
		{
			ID:          1721,
			Slug:        "mysterium",
			Name:        "Mysterium",
			BinanceUSDT: "MYSTUSDT",
		},
	},
	"AIDUS": {
		{
			ID:          3785,
			Slug:        "aidus-token",
			Name:        "AIDUS TOKEN",
			BinanceUSDT: "AIDUSUSDT",
		},
	},
	"OHM": {
		{
			ID:          9067,
			Slug:        "olympus",
			Name:        "Olympus",
			BinanceUSDT: "OHMUSDT",
		},
	},
	"COVAL": {
		{
			ID:          788,
			Slug:        "circuits-of-value",
			Name:        "Circuits of Value",
			BinanceUSDT: "COVALUSDT",
		},
	},
	"GSR": {
		{
			ID:          1846,
			Slug:        "geysercoin",
			Name:        "GeyserCoin",
			BinanceUSDT: "GSRUSDT",
		},
	},
	"GYSR": {
		{
			ID:          7661,
			Slug:        "gysr",
			Name:        "GYSR",
			BinanceUSDT: "GYSRUSDT",
		},
	},
	"ABS": {
		{
			ID:          8032,
			Slug:        "absorber-protocol",
			Name:        "Absorber Protocol",
			BinanceUSDT: "ABSUSDT",
		},
	},
	"ON": {
		{
			ID:          7021,
			Slug:        "ofin-token",
			Name:        "OFIN Token",
			BinanceUSDT: "ONUSDT",
		},
	},
	"EQL": {
		{
			ID:          2479,
			Slug:        "equal",
			Name:        "Equal",
			BinanceUSDT: "EQLUSDT",
		},
	},
	"YFP": {
		{
			ID:          6915,
			Slug:        "yearn-finance-protocol",
			Name:        "Yearn Finance Protocol",
			BinanceUSDT: "YFPUSDT",
		},
	},
	"XSGD": {
		{
			ID:          8489,
			Slug:        "xsgd",
			Name:        "XSGD",
			BinanceUSDT: "XSGDUSDT",
		},
	},
	"DGB": {
		{
			ID:          109,
			Slug:        "digibyte",
			Name:        "DigiByte",
			BinanceUSDT: "DGBUSDT",
		},
	},
	"YFPI": {
		{
			ID:          7249,
			Slug:        "yearn-finance-passive-income",
			Name:        "Yearn Finance Passive Income",
			BinanceUSDT: "YFPIUSDT",
		},
	},
	"D": {
		{
			ID:          1769,
			Slug:        "denarius-d",
			Name:        "Denarius",
			BinanceUSDT: "DUSDT",
		},
	},
	"JINDOGE": {
		{
			ID:          9701,
			Slug:        "jindoge",
			Name:        "Jindoge",
			BinanceUSDT: "JINDOGEUSDT",
		},
	},
	"QBT": {
		{
			ID:          2242,
			Slug:        "qbao",
			Name:        "Qbao",
			BinanceUSDT: "QBTUSDT",
		},
	},
	"SHA": {
		{
			ID:          3831,
			Slug:        "safe-haven",
			Name:        "Safe Haven",
			BinanceUSDT: "SHAUSDT",
		},
	},
	"RENFIL": {
		{
			ID:          7997,
			Slug:        "renfil",
			Name:        "renFIL",
			BinanceUSDT: "RENFILUSDT",
		},
	},
	"WORLD": {
		{
			ID:          8366,
			Slug:        "world-token",
			Name:        "World Token",
			BinanceUSDT: "WORLDUSDT",
		},
	},
	"BULLSHIT": {
		{
			ID:          6088,
			Slug:        "3x-long-shitcoin-index-token",
			Name:        "3X Long Shitcoin Index Token",
			BinanceUSDT: "BULLSHITUSDT",
		},
	},
	"BCN": {
		{
			ID:          372,
			Slug:        "bytecoin-bcn",
			Name:        "Bytecoin",
			BinanceUSDT: "BCNUSDT",
		},
	},
	"YYFI": {
		{
			ID:          7551,
			Slug:        "yyfi-protocol",
			Name:        "YYFI.Protocol",
			BinanceUSDT: "YYFIUSDT",
		},
	},
	"BBQ": {
		{
			ID:          10443,
			Slug:        "barbecueswap-finance",
			Name:        "BarbecueSwap Finance",
			BinanceUSDT: "BBQUSDT",
		},
	},
	"TNC": {
		{
			ID:   5524,
			Slug: "tnc-coin",
			Name: "TNC Coin",
		},
		{
			ID:   2443,
			Slug: "trinity-network-credit",
			Name: "Trinity Network Credit",
		},
	},
	"XCON": {
		{
			ID:          3932,
			Slug:        "connect-coin",
			Name:        "Connect Coin",
			BinanceUSDT: "XCONUSDT",
		},
	},
	"SHD": {
		{
			ID:   8815,
			Slug: "shadetech",
			Name: "Shadetech",
		},
		{
			ID:   9151,
			Slug: "shardingdao",
			Name: "ShardingDAO",
		},
	},
	"LUN": {
		{
			ID:          1658,
			Slug:        "lunyr",
			Name:        "Lunyr",
			BinanceUSDT: "LUNUSDT",
		},
	},
	"HEBE": {
		{
			ID:          5250,
			Slug:        "hebeblock",
			Name:        "HebeBlock",
			BinanceUSDT: "HEBEUSDT",
		},
	},
	"TOMOBULL": {
		{
			ID:          6090,
			Slug:        "3x-long-tomochain-token",
			Name:        "3X Long TomoChain Token",
			BinanceUSDT: "TOMOBULLUSDT",
		},
	},
	"ANG": {
		{
			ID:          6598,
			Slug:        "aureus-nummus-gold",
			Name:        "Aureus Nummus Gold",
			BinanceUSDT: "ANGUSDT",
		},
	},
	"GNBU": {
		{
			ID:          10632,
			Slug:        "nimbus-governance-token",
			Name:        "Nimbus Governance Token",
			BinanceUSDT: "GNBUUSDT",
		},
	},
	"STN": {
		{
			ID:   9043,
			Slug: "stone-token",
			Name: "Stone DeFi",
		},
		{
			ID:   8321,
			Slug: "sting",
			Name: "STING",
		},
	},
	"QASH": {
		{
			ID:          2213,
			Slug:        "qash",
			Name:        "QASH",
			BinanceUSDT: "QASHUSDT",
		},
	},
	"BCDT": {
		{
			ID:          3616,
			Slug:        "evidenz",
			Name:        "EvidenZ",
			BinanceUSDT: "BCDTUSDT",
		},
	},
	"SHREW": {
		{
			ID:          11010,
			Slug:        "shrew",
			Name:        "Shrew",
			BinanceUSDT: "SHREWUSDT",
		},
	},
	"VAULT": {
		{
			ID:          5003,
			Slug:        "vault",
			Name:        "VAULT",
			BinanceUSDT: "VAULTUSDT",
		},
	},
	"EVED": {
		{
			ID:          3953,
			Slug:        "evedo",
			Name:        "Evedo",
			BinanceUSDT: "EVEDUSDT",
		},
	},
	"FAN": {
		{
			ID:          10609,
			Slug:        "fanspel",
			Name:        "Fanspel",
			BinanceUSDT: "FANUSDT",
		},
	},
	"PMON": {
		{
			ID:   8968,
			Slug: "polkamon",
			Name: "Polychain Monsters",
		},
		{
			ID:   9821,
			Slug: "pocmon",
			Name: "PocMon",
		},
	},
	"TOC": {
		{
			ID:          3965,
			Slug:        "touchcon",
			Name:        "TouchCon",
			BinanceUSDT: "TOCUSDT",
		},
	},
	"PRDZ": {
		{
			ID:          7998,
			Slug:        "predictz",
			Name:        "Predictz",
			BinanceUSDT: "PRDZUSDT",
		},
	},
	"REV": {
		{
			ID:   2135,
			Slug: "revain",
			Name: "Revain",
		},
		{
			ID:   2021,
			Slug: "rchain",
			Name: "RChain",
		},
	},
	"PROM": {
		{
			ID:          4120,
			Slug:        "prometeus",
			Name:        "Prometeus",
			BinanceUSDT: "PROMUSDT",
		},
	},
	"PNODE": {
		{
			ID:          9657,
			Slug:        "pinknode",
			Name:        "Pinknode",
			BinanceUSDT: "PNODEUSDT",
		},
	},
	"SAFESTAR": {
		{
			ID:          8872,
			Slug:        "safe-star",
			Name:        "Safe Star",
			BinanceUSDT: "SAFESTARUSDT",
		},
	},
	"MCM": {
		{
			ID:          5125,
			Slug:        "mochimo",
			Name:        "Mochimo",
			BinanceUSDT: "MCMUSDT",
		},
	},
	"PRVS": {
		{
			ID:          8493,
			Slug:        "previse",
			Name:        "Previse",
			BinanceUSDT: "PRVSUSDT",
		},
	},
	"BAT": {
		{
			ID:   1697,
			Slug: "basic-attention-token",
			Name: "Basic Attention Token",
		},
		{
			ID:   10028,
			Slug: "bat-finance",
			Name: "Bat Finance",
		},
	},
	"EVY": {
		{
			ID:          3754,
			Slug:        "everycoin",
			Name:        "EveryCoin ",
			BinanceUSDT: "EVYUSDT",
		},
	},
	"SHPING": {
		{
			ID:          3422,
			Slug:        "shping",
			Name:        "SHPING",
			BinanceUSDT: "SHPINGUSDT",
		},
	},
	"STEEP": {
		{
			ID:          3395,
			Slug:        "steepcoin",
			Name:        "SteepCoin",
			BinanceUSDT: "STEEPUSDT",
		},
	},
	"SAM": {
		{
			ID:          7121,
			Slug:        "samurai",
			Name:        "Samurai",
			BinanceUSDT: "SAMUSDT",
		},
	},
	"BCPT": {
		{
			ID:          2061,
			Slug:        "blockmason",
			Name:        "Blockmason Credit Protocol",
			BinanceUSDT: "BCPTUSDT",
		},
	},
	"YOC": {
		{
			ID:          1156,
			Slug:        "yocoin",
			Name:        "Yocoin",
			BinanceUSDT: "YOCUSDT",
		},
	},
	"WICC": {
		{
			ID:          2346,
			Slug:        "waykichain",
			Name:        "WaykiChain",
			BinanceUSDT: "WICCUSDT",
		},
	},
	"URUS": {
		{
			ID:          8616,
			Slug:        "urus",
			Name:        "Aurox",
			BinanceUSDT: "URUSUSDT",
		},
	},
	"HAG": {
		{
			ID:          9846,
			Slug:        "hagglex",
			Name:        "HaggleX",
			BinanceUSDT: "HAGUSDT",
		},
	},
	"DIP": {
		{
			ID:   6588,
			Slug: "etherisc",
			Name: "Etherisc DIP Token",
		},
		{
			ID:   6775,
			Slug: "dipper-network",
			Name: "Dipper Network",
		},
	},
	"ETHBULL": {
		{
			ID:          5217,
			Slug:        "3x-long-ethereum-token",
			Name:        "3X Long Ethereum Token",
			BinanceUSDT: "ETHBULLUSDT",
		},
	},
	"LED": {
		{
			ID:          10598,
			Slug:        "ledgerscore",
			Name:        "LedgerScore",
			BinanceUSDT: "LEDUSDT",
		},
	},
	"FROGE": {
		{
			ID:          9234,
			Slug:        "froge-finance",
			Name:        "Froge Finance",
			BinanceUSDT: "FROGEUSDT",
		},
	},
	"SMDX": {
		{
			ID:          7688,
			Slug:        "somidax",
			Name:        "SOMIDAX",
			BinanceUSDT: "SMDXUSDT",
		},
	},
	"ATP": {
		{
			ID:   3620,
			Slug: "atlas-protocol",
			Name: "Atlas Protocol",
		},
		{
			ID:   7814,
			Slug: "alaya",
			Name: "Alaya",
		},
	},
	"PUGL": {
		{
			ID:          10333,
			Slug:        "puglife",
			Name:        "PUGLIFE",
			BinanceUSDT: "PUGLUSDT",
		},
	},
	"CNC": {
		{
			ID:          6018,
			Slug:        "global-china-cash",
			Name:        "Global China Cash",
			BinanceUSDT: "CNCUSDT",
		},
	},
	"ZAP": {
		{
			ID:          2363,
			Slug:        "zap",
			Name:        "Zap",
			BinanceUSDT: "ZAPUSDT",
		},
	},
	"MXW": {
		{
			ID:          6369,
			Slug:        "maxonrow",
			Name:        "Maxonrow",
			BinanceUSDT: "MXWUSDT",
		},
	},
	"SBF": {
		{
			ID:          9953,
			Slug:        "steakbankfinance",
			Name:        "SteakBankFinance",
			BinanceUSDT: "SBFUSDT",
		},
	},
	"GRN": {
		{
			ID:          2746,
			Slug:        "greenpower",
			Name:        "GreenPower",
			BinanceUSDT: "GRNUSDT",
		},
	},
	"EPIC": {
		{
			ID:          5435,
			Slug:        "epic-cash",
			Name:        "Epic Cash",
			BinanceUSDT: "EPICUSDT",
		},
	},
	"WIZARD": {
		{
			ID:          10960,
			Slug:        "wizard",
			Name:        "WIZARD",
			BinanceUSDT: "WIZARDUSDT",
		},
	},
	"LABRA": {
		{
			ID:          9519,
			Slug:        "labracoin",
			Name:        "LabraCoin",
			BinanceUSDT: "LABRAUSDT",
		},
	},
	"KONO": {
		{
			ID:          8697,
			Slug:        "konomi-network",
			Name:        "Konomi Network",
			BinanceUSDT: "KONOUSDT",
		},
	},
	"PDOGE": {
		{
			ID:          11037,
			Slug:        "polkadoge",
			Name:        "POLKADOGE",
			BinanceUSDT: "PDOGEUSDT",
		},
	},
	"WCK": {
		{
			ID:          7473,
			Slug:        "wrapped-cryptokitties",
			Name:        "Wrapped Basic CryptoKitties",
			BinanceUSDT: "WCKUSDT",
		},
	},
	"XNC": {
		{
			ID:          5060,
			Slug:        "xenioscoin",
			Name:        "XeniosCoin",
			BinanceUSDT: "XNCUSDT",
		},
	},
	"SPND": {
		{
			ID:          10369,
			Slug:        "safepanda",
			Name:        "SafePanda",
			BinanceUSDT: "SPNDUSDT",
		},
	},
	"WAXP": {
		{
			ID:          2300,
			Slug:        "wax",
			Name:        "WAX",
			BinanceUSDT: "WAXPUSDT",
		},
	},
	"ASP": {
		{
			ID:          7536,
			Slug:        "aspire",
			Name:        "Aspire",
			BinanceUSDT: "ASPUSDT",
		},
	},
	"COLX": {
		{
			ID:          2001,
			Slug:        "colossusxt",
			Name:        "ColossusXT",
			BinanceUSDT: "COLXUSDT",
		},
	},
	"PLUS1": {
		{
			ID:          3456,
			Slug:        "plusonecoin",
			Name:        "PlusOneCoin",
			BinanceUSDT: "PLUS1USDT",
		},
	},
	"VERI": {
		{
			ID:          1710,
			Slug:        "veritaseum",
			Name:        "Veritaseum",
			BinanceUSDT: "VERIUSDT",
		},
	},
	"yTSLA": {
		{
			ID:          7140,
			Slug:        "ytsla-finance",
			Name:        "yTSLA Finance",
			BinanceUSDT: "yTSLAUSDT",
		},
	},
	"CTT": {
		{
			ID:   5397,
			Slug: "castweet",
			Name: "Castweet",
		},
		{
			ID:   9503,
			Slug: "cryptotycoon",
			Name: "CryptoTycoon",
		},
	},
	"ETNX": {
		{
			ID:          4448,
			Slug:        "electronero",
			Name:        "Electronero",
			BinanceUSDT: "ETNXUSDT",
		},
	},
	"NOKU": {
		{
			ID:          3138,
			Slug:        "noku",
			Name:        "Noku",
			BinanceUSDT: "NOKUUSDT",
		},
	},
	"CREVA": {
		{
			ID:          986,
			Slug:        "crevacoin",
			Name:        "CrevaCoin",
			BinanceUSDT: "CREVAUSDT",
		},
	},
	"GSMT": {
		{
			ID:          6851,
			Slug:        "grafsound",
			Name:        "GrafSound",
			BinanceUSDT: "GSMTUSDT",
		},
	},
	"PRX": {
		{
			ID:          3668,
			Slug:        "proxynode",
			Name:        "ProxyNode",
			BinanceUSDT: "PRXUSDT",
		},
	},
	"MTCN": {
		{
			ID:          4190,
			Slug:        "multicoin",
			Name:        "Multicoin",
			BinanceUSDT: "MTCNUSDT",
		},
	},
	"BGG": {
		{
			ID:          3587,
			Slug:        "bgogo-token",
			Name:        "Bgogo Token",
			BinanceUSDT: "BGGUSDT",
		},
	},
	"XYO": {
		{
			ID:          2765,
			Slug:        "xyo",
			Name:        "XYO",
			BinanceUSDT: "XYOUSDT",
		},
	},
	"THX": {
		{
			ID:          3916,
			Slug:        "thorenext",
			Name:        "ThoreNext",
			BinanceUSDT: "THXUSDT",
		},
	},
	"BAB": {
		{
			ID:          10549,
			Slug:        "bauble",
			Name:        "Bauble",
			BinanceUSDT: "BABUSDT",
		},
	},
	"WABI": {
		{
			ID:          2267,
			Slug:        "wabi",
			Name:        "Wabi",
			BinanceUSDT: "WABIUSDT",
		},
	},
	"CBAT": {
		{
			ID:          5742,
			Slug:        "compound-basic-attention-token",
			Name:        "Compound Basic Attention Token",
			BinanceUSDT: "CBATUSDT",
		},
	},
	"PYE": {
		{
			ID:          10174,
			Slug:        "creampye",
			Name:        "CREAMPYE",
			BinanceUSDT: "PYEUSDT",
		},
	},
	"HAPPY": {
		{
			ID:          9824,
			Slug:        "thehappycoin",
			Name:        "HappyCoin",
			BinanceUSDT: "HAPPYUSDT",
		},
	},
	"DEPAY": {
		{
			ID:          8181,
			Slug:        "depay",
			Name:        "DePay",
			BinanceUSDT: "DEPAYUSDT",
		},
	},
	"CCAKE": {
		{
			ID:          8747,
			Slug:        "cheesecakeswap-token",
			Name:        "CheesecakeSwap Token",
			BinanceUSDT: "CCAKEUSDT",
		},
	},
	"XPD": {
		{
			ID:          260,
			Slug:        "petrodollar",
			Name:        "PetroDollar",
			BinanceUSDT: "XPDUSDT",
		},
	},
	"LINKDOWN": {
		{
			ID:          7012,
			Slug:        "linkdown",
			Name:        "LINKDOWN",
			BinanceUSDT: "LINKDOWNUSDT",
		},
	},
	"FRY": {
		{
			ID:          6119,
			Slug:        "foundry",
			Name:        "Foundry",
			BinanceUSDT: "FRYUSDT",
		},
	},
	"PKR": {
		{
			ID:          10427,
			Slug:        "polker",
			Name:        "Polker",
			BinanceUSDT: "PKRUSDT",
		},
	},
	"TSM": {
		{
			ID:          7893,
			Slug:        "taiwan-semiconductor-mfg-tokenized-stock-ftx",
			Name:        "Taiwan Semiconductor Mfg tokenized stock FTX",
			BinanceUSDT: "TSMUSDT",
		},
	},
	"OMI": {
		{
			ID:          6432,
			Slug:        "ecomi",
			Name:        "ECOMI",
			BinanceUSDT: "OMIUSDT",
		},
	},
	"UBN": {
		{
			ID:          4071,
			Slug:        "ubricoin",
			Name:        "Ubricoin",
			BinanceUSDT: "UBNUSDT",
		},
	},
	"MOONMOON": {
		{
			ID:          9525,
			Slug:        "moonmoon",
			Name:        "MoonMoon",
			BinanceUSDT: "MOONMOONUSDT",
		},
	},
	"BITG": {
		{
			ID:          2604,
			Slug:        "bitgreen",
			Name:        "BitGreen",
			BinanceUSDT: "BITGUSDT",
		},
	},
	"ZT": {
		{
			ID:          3458,
			Slug:        "zbg-token",
			Name:        "ZBG Token",
			BinanceUSDT: "ZTUSDT",
		},
	},
	"0xMR": {
		{
			ID:          5668,
			Slug:        "0xmonero",
			Name:        "0xMonero",
			BinanceUSDT: "0xMRUSDT",
		},
	},
	"ARI": {
		{
			ID:          9565,
			Slug:        "arise-finance",
			Name:        "Arise Finance",
			BinanceUSDT: "ARIUSDT",
		},
	},
	"INNBC": {
		{
			ID:          5016,
			Slug:        "innovative-bioresearch-coin",
			Name:        "Innovative Bioresearch Coin",
			BinanceUSDT: "INNBCUSDT",
		},
	},
	"PERL": {
		{
			ID:          4293,
			Slug:        "perlin",
			Name:        "PERL.eco",
			BinanceUSDT: "PERLUSDT",
		},
	},
	"LUXO": {
		{
			ID:          8443,
			Slug:        "luxochain",
			Name:        "LUXOCHAIN",
			BinanceUSDT: "LUXOUSDT",
		},
	},
	"ADT": {
		{
			ID:          1775,
			Slug:        "adtoken",
			Name:        "adToken",
			BinanceUSDT: "ADTUSDT",
		},
	},
	"SKYBORN": {
		{
			ID:          10184,
			Slug:        "sky-born",
			Name:        "SkyBorn",
			BinanceUSDT: "SKYBORNUSDT",
		},
	},
	"CIPHC": {
		{
			ID:          5844,
			Slug:        "cipher-core-token",
			Name:        "Cipher Core Token",
			BinanceUSDT: "CIPHCUSDT",
		},
	},
	"ZIOT": {
		{
			ID:          9226,
			Slug:        "ziot-coin",
			Name:        "ziot Coin",
			BinanceUSDT: "ZIOTUSDT",
		},
	},
	"XNK": {
		{
			ID:          2549,
			Slug:        "ink-protocol",
			Name:        "Ink Protocol",
			BinanceUSDT: "XNKUSDT",
		},
	},
	"PHX": {
		{
			ID:          10206,
			Slug:        "phoenix-protocol",
			Name:        "Phoenix Protocol",
			BinanceUSDT: "PHXUSDT",
		},
	},
	"FLM": {
		{
			ID:          7150,
			Slug:        "flamingo",
			Name:        "Flamingo",
			BinanceUSDT: "FLMUSDT",
		},
	},
	"BNTX": {
		{
			ID:   6586,
			Slug: "bintex-futures",
			Name: "Bintex Futures",
		},
		{
			ID:   7925,
			Slug: "biontech-tokenized-stock-bittrex",
			Name: "BioNTech tokenized stock Bittrex",
		},
		{
			ID:   7918,
			Slug: "biontech-tokenized-stock-ftx",
			Name: "BioNTech tokenized stock FTX",
		},
	},
	"ROT": {
		{
			ID:          7164,
			Slug:        "rotten",
			Name:        "Rotten",
			BinanceUSDT: "ROTUSDT",
		},
	},
	"LEMD": {
		{
			ID:          9455,
			Slug:        "lemond",
			Name:        "Lemond",
			BinanceUSDT: "LEMDUSDT",
		},
	},
	"SCHO": {
		{
			ID:          8837,
			Slug:        "scholarship-coin",
			Name:        "Scholarship Coin",
			BinanceUSDT: "SCHOUSDT",
		},
	},
	"BIS": {
		{
			ID:          2009,
			Slug:        "bismuth",
			Name:        "Bismuth",
			BinanceUSDT: "BISUSDT",
		},
	},
	"EFX": {
		{
			ID:          2666,
			Slug:        "effect-ai",
			Name:        "Effect.AI",
			BinanceUSDT: "EFXUSDT",
		},
	},
	"vXRP": {
		{
			ID:          7965,
			Slug:        "venus-xrp",
			Name:        "Venus XRP",
			BinanceUSDT: "vXRPUSDT",
		},
	},
	"NFTSHIBA": {
		{
			ID:          10614,
			Slug:        "nftshiba-finance",
			Name:        "NFTShiba.Finance",
			BinanceUSDT: "NFTSHIBAUSDT",
		},
	},
	"XEC": {
		{
			ID:          10791,
			Slug:        "ecash",
			Name:        "eCash",
			BinanceUSDT: "XECUSDT",
		},
	},
	"ARQ": {
		{
			ID:          3882,
			Slug:        "arqma",
			Name:        "Arqma",
			BinanceUSDT: "ARQUSDT",
		},
	},
	"SAFEMOON": {
		{
			ID:          8757,
			Slug:        "safemoon",
			Name:        "SafeMoon",
			BinanceUSDT: "SAFEMOONUSDT",
		},
	},
	"SHARD": {
		{
			ID:          3335,
			Slug:        "shard",
			Name:        "Shard",
			BinanceUSDT: "SHARDUSDT",
		},
	},
	"HBN": {
		{
			ID:          78,
			Slug:        "hobonickels",
			Name:        "HoboNickels",
			BinanceUSDT: "HBNUSDT",
		},
	},
	"USDA": {
		{
			ID:          5058,
			Slug:        "usda",
			Name:        "USDA",
			BinanceUSDT: "USDAUSDT",
		},
	},
	"DEVE": {
		{
			ID:          7766,
			Slug:        "divert-finance",
			Name:        "Divert Finance",
			BinanceUSDT: "DEVEUSDT",
		},
	},
	"ULTGG": {
		{
			ID:          10881,
			Slug:        "ultimogg",
			Name:        "UltimoGG",
			BinanceUSDT: "ULTGGUSDT",
		},
	},
	"MOONRISE": {
		{
			ID:          10681,
			Slug:        "moonrise",
			Name:        "MoonRise",
			BinanceUSDT: "MOONRISEUSDT",
		},
	},
	"PAMPTHER": {
		{
			ID:          9925,
			Slug:        "pampther",
			Name:        "Pampther",
			BinanceUSDT: "PAMPTHERUSDT",
		},
	},
	"FORINT": {
		{
			ID:          10981,
			Slug:        "forint-token",
			Name:        "Forint Token",
			BinanceUSDT: "FORINTUSDT",
		},
	},
	"PUP": {
		{
			ID:          10242,
			Slug:        "pupper",
			Name:        "Pupper",
			BinanceUSDT: "PUPUSDT",
		},
	},
	"THEMOON": {
		{
			ID:          10773,
			Slug:        "win-space-ticket",
			Name:        "Win Space Ticket",
			BinanceUSDT: "THEMOONUSDT",
		},
	},
	"GHX": {
		{
			ID:          6554,
			Slug:        "gamercoin",
			Name:        "GamerCoin",
			BinanceUSDT: "GHXUSDT",
		},
	},
	"MLN": {
		{
			ID:          1552,
			Slug:        "enzyme",
			Name:        "Enzyme",
			BinanceUSDT: "MLNUSDT",
		},
	},
	"FORM": {
		{
			ID:          10408,
			Slug:        "formation-fi",
			Name:        "Formation Fi",
			BinanceUSDT: "FORMUSDT",
		},
	},
	"YMEN": {
		{
			ID:          7039,
			Slug:        "ymen-finance",
			Name:        "Ymen.Finance",
			BinanceUSDT: "YMENUSDT",
		},
	},
	"ZB": {
		{
			ID:          3351,
			Slug:        "zb-token",
			Name:        "ZB Token",
			BinanceUSDT: "ZBUSDT",
		},
	},
	"IDK": {
		{
			ID:          5580,
			Slug:        "idk",
			Name:        "IDK",
			BinanceUSDT: "IDKUSDT",
		},
	},
	"NTON": {
		{
			ID:          6764,
			Slug:        "nton",
			Name:        "NTON",
			BinanceUSDT: "NTONUSDT",
		},
	},
	"BAGS": {
		{
			ID:          8264,
			Slug:        "basis-gold-share",
			Name:        "Basis Gold Share",
			BinanceUSDT: "BAGSUSDT",
		},
	},
	"MCHC": {
		{
			ID:          9686,
			Slug:        "my-crypto-heroes",
			Name:        "My Crypto Heroes",
			BinanceUSDT: "MCHCUSDT",
		},
	},
	"GZIL": {
		{
			ID:          8031,
			Slug:        "governance-zil",
			Name:        "governance ZIL",
			BinanceUSDT: "GZILUSDT",
		},
	},
	"GUH": {
		{
			ID:          10481,
			Slug:        "goes-up-higher",
			Name:        "Goes Up Higher",
			BinanceUSDT: "GUHUSDT",
		},
	},
	"MINDS": {
		{
			ID:          8675,
			Slug:        "minds",
			Name:        "Minds",
			BinanceUSDT: "MINDSUSDT",
		},
	},
	"EUC": {
		{
			ID:          1038,
			Slug:        "eurocoin",
			Name:        "Eurocoin",
			BinanceUSDT: "EUCUSDT",
		},
	},
	"PONZI": {
		{
			ID:          1259,
			Slug:        "ponzicoin",
			Name:        "PonziCoin",
			BinanceUSDT: "PONZIUSDT",
		},
	},
	"LEOPARD": {
		{
			ID:          10393,
			Slug:        "leopard",
			Name:        "LEOPARD",
			BinanceUSDT: "LEOPARDUSDT",
		},
	},
	"SERGS": {
		{
			ID:          7595,
			Slug:        "sergs",
			Name:        "SERGS",
			BinanceUSDT: "SERGSUSDT",
		},
	},
	"PGO": {
		{
			ID:          5381,
			Slug:        "pengolincoin",
			Name:        "PengolinCoin",
			BinanceUSDT: "PGOUSDT",
		},
	},
	"FTX": {
		{
			ID:          2667,
			Slug:        "fintrux-network",
			Name:        "FintruX Network",
			BinanceUSDT: "FTXUSDT",
		},
	},
	"AURUM": {
		{
			ID:          9224,
			Slug:        "alchemist-defi-aurum",
			Name:        "Alchemist DeFi Aurum",
			BinanceUSDT: "AURUMUSDT",
		},
	},
	"FLOW": {
		{
			ID:          4558,
			Slug:        "flow",
			Name:        "Flow",
			BinanceUSDT: "FLOWUSDT",
		},
	},
	"SEPHI": {
		{
			ID:          10758,
			Slug:        "sephiroth-inu",
			Name:        "Sephiroth Inu",
			BinanceUSDT: "SEPHIUSDT",
		},
	},
	"ONES": {
		{
			ID:          7136,
			Slug:        "oneswap-dao-token",
			Name:        "OneSwap DAO Token",
			BinanceUSDT: "ONESUSDT",
		},
	},
	"CYC": {
		{
			ID:          8590,
			Slug:        "cyclone-protocol",
			Name:        "Cyclone Protocol",
			BinanceUSDT: "CYCUSDT",
		},
	},
	"HOMI": {
		{
			ID:          5572,
			Slug:        "homihelp",
			Name:        "HOMIHELP",
			BinanceUSDT: "HOMIUSDT",
		},
	},
	"STAR": {
		{
			ID:   2295,
			Slug: "starbase",
			Name: "Starbase",
		},
		{
			ID:   10728,
			Slug: "pornstar",
			Name: "Pornstar",
		},
		{
			ID:   8593,
			Slug: "filestar",
			Name: "FileStar",
		},
	},
	"GOFI": {
		{
			ID:          8873,
			Slug:        "goswapp",
			Name:        "GoSwapp",
			BinanceUSDT: "GOFIUSDT",
		},
	},
	"LELE": {
		{
			ID:          7781,
			Slug:        "lelefoodchain",
			Name:        "LeLeFoodChain",
			BinanceUSDT: "LELEUSDT",
		},
	},
	"FEAST": {
		{
			ID:          10116,
			Slug:        "feast-finance",
			Name:        "Feast Finance",
			BinanceUSDT: "FEASTUSDT",
		},
	},
	"GLDY": {
		{
			ID:          5538,
			Slug:        "buzzshow",
			Name:        "Buzzshow",
			BinanceUSDT: "GLDYUSDT",
		},
	},
	"SHFT": {
		{
			ID:          8917,
			Slug:        "shyft-network",
			Name:        "Shyft Network",
			BinanceUSDT: "SHFTUSDT",
		},
	},
	"DOGES": {
		{
			ID:          7377,
			Slug:        "dogeswap",
			Name:        "Dogeswap",
			BinanceUSDT: "DOGESUSDT",
		},
	},
	"ALA": {
		{
			ID:          6980,
			Slug:        "aladiex",
			Name:        "AladiEx",
			BinanceUSDT: "ALAUSDT",
		},
	},
	"BTNYX": {
		{
			ID:          7513,
			Slug:        "bitonyx",
			Name:        "BitOnyx",
			BinanceUSDT: "BTNYXUSDT",
		},
	},
	"ICOB": {
		{
			ID:          1514,
			Slug:        "icobid",
			Name:        "ICOBID",
			BinanceUSDT: "ICOBUSDT",
		},
	},
	"VSP": {
		{
			ID:          8492,
			Slug:        "vesper",
			Name:        "Vesper",
			BinanceUSDT: "VSPUSDT",
		},
	},
	"LOCK": {
		{
			ID:          6566,
			Slug:        "meridian-network",
			Name:        "Meridian Network",
			BinanceUSDT: "LOCKUSDT",
		},
	},
	"GOAT": {
		{
			ID:   9602,
			Slug: "goat-coin",
			Name: "GOAT COIN",
		},
		{
			ID:   9069,
			Slug: "goatcoin",
			Name: "Goatcoin",
		},
	},
	"ZEBI": {
		{
			ID:          2685,
			Slug:        "zebi-token",
			Name:        "Zebi Token",
			BinanceUSDT: "ZEBIUSDT",
		},
	},
	"OCN": {
		{
			ID:          2458,
			Slug:        "odyssey",
			Name:        "Odyssey",
			BinanceUSDT: "OCNUSDT",
		},
	},
	"FLASH": {
		{
			ID:          1755,
			Slug:        "flash",
			Name:        "Flash",
			BinanceUSDT: "FLASHUSDT",
		},
	},
	"BLY": {
		{
			ID:          6283,
			Slug:        "blocery",
			Name:        "Blocery",
			BinanceUSDT: "BLYUSDT",
		},
	},
	"FMG": {
		{
			ID:          8902,
			Slug:        "fm-gallery",
			Name:        "FM Gallery",
			BinanceUSDT: "FMGUSDT",
		},
	},
	"XBI": {
		{
			ID:          3166,
			Slug:        "bitcoin-incognito",
			Name:        "Bitcoin Incognito",
			BinanceUSDT: "XBIUSDT",
		},
	},
	"SBTC": {
		{
			ID:   5764,
			Slug: "sbtc",
			Name: "sBTC",
		},
		{
			ID:   2282,
			Slug: "super-bitcoin",
			Name: "Super Bitcoin",
		},
	},
	"DIS": {
		{
			ID:          8310,
			Slug:        "tosdis",
			Name:        "TosDis",
			BinanceUSDT: "DISUSDT",
		},
	},
	"RICH": {
		{
			ID:   10863,
			Slug: "richcity",
			Name: "RichCity",
		},
		{
			ID:   9648,
			Slug: "richie",
			Name: "Richie",
		},
		{
			ID:   8546,
			Slug: "rich-maker",
			Name: "Rich Maker",
		},
	},
	"MDOGE": {
		{
			ID:          10911,
			Slug:        "missdoge",
			Name:        "MissDoge",
			BinanceUSDT: "MDOGEUSDT",
		},
	},
	"FST": {
		{
			ID:   3840,
			Slug: "1irstcoin",
			Name: "1irstcoin",
		},
		{
			ID:   32,
			Slug: "fastcoin",
			Name: "Fastcoin",
		},
		{
			ID:   8961,
			Slug: "futureswap",
			Name: "Futureswap",
		},
	},
	"SNO": {
		{
			ID:          3276,
			Slug:        "savenode",
			Name:        "SaveNode",
			BinanceUSDT: "SNOUSDT",
		},
	},
	"ISDT": {
		{
			ID:          7148,
			Slug:        "istardust",
			Name:        "ISTARDUST",
			BinanceUSDT: "ISDTUSDT",
		},
	},
	"SWOP": {
		{
			ID:          8732,
			Slug:        "swop",
			Name:        "Swop",
			BinanceUSDT: "SWOPUSDT",
		},
	},
	"MYX": {
		{
			ID:          6594,
			Slug:        "myx-network",
			Name:        "MYX Network",
			BinanceUSDT: "MYXUSDT",
		},
	},
	"MGB": {
		{
			ID:          9165,
			Slug:        "magic-balancer",
			Name:        "Magic Balancer",
			BinanceUSDT: "MGBUSDT",
		},
	},
	"GMR": {
		{
			ID:          9629,
			Slug:        "gmr-finance",
			Name:        "GMR Finance",
			BinanceUSDT: "GMRUSDT",
		},
	},
	"BAND": {
		{
			ID:          4679,
			Slug:        "band-protocol",
			Name:        "Band Protocol",
			BinanceUSDT: "BANDUSDT",
		},
	},
	"FSCC": {
		{
			ID:          6447,
			Slug:        "fisco",
			Name:        "Fisco Coin",
			BinanceUSDT: "FSCCUSDT",
		},
	},
	"NAX": {
		{
			ID:          5873,
			Slug:        "nextdao",
			Name:        "NextDAO",
			BinanceUSDT: "NAXUSDT",
		},
	},
	"ME": {
		{
			ID:          6026,
			Slug:        "all-me",
			Name:        "All.me",
			BinanceUSDT: "MEUSDT",
		},
	},
	"ELD": {
		{
			ID:          3746,
			Slug:        "electrum-dark-eld",
			Name:        "Electrum Dark",
			BinanceUSDT: "ELDUSDT",
		},
	},
	"MAID": {
		{
			ID:          291,
			Slug:        "maidsafecoin",
			Name:        "MaidSafeCoin",
			BinanceUSDT: "MAIDUSDT",
		},
	},
	"MONA": {
		{
			ID:   213,
			Slug: "monacoin",
			Name: "MonaCoin",
		},
		{
			ID:   7866,
			Slug: "monavale",
			Name: "Monavale",
		},
	},
	"FIO": {
		{
			ID:          5865,
			Slug:        "fio-protocol",
			Name:        "FIO Protocol",
			BinanceUSDT: "FIOUSDT",
		},
	},
	"TNB": {
		{
			ID:          2235,
			Slug:        "time-new-bank",
			Name:        "Time New Bank",
			BinanceUSDT: "TNBUSDT",
		},
	},
	"GMNG": {
		{
			ID:          7399,
			Slug:        "global-gaming",
			Name:        "Global Gaming",
			BinanceUSDT: "GMNGUSDT",
		},
	},
	"N0031": {
		{
			ID:          7575,
			Slug:        "nyfi",
			Name:        "nYFI",
			BinanceUSDT: "N0031USDT",
		},
	},
	"CVXCRV": {
		{
			ID:          10291,
			Slug:        "convex-crv",
			Name:        "Convex CRV",
			BinanceUSDT: "CVXCRVUSDT",
		},
	},
	"SHX": {
		{
			ID:          3661,
			Slug:        "stronghold-token",
			Name:        "Stronghold Token",
			BinanceUSDT: "SHXUSDT",
		},
	},
	"ROAD": {
		{
			ID:   9145,
			Slug: "yellow-road",
			Name: "Yellow Road",
		},
		{
			ID:   5028,
			Slug: "road",
			Name: "ROAD",
		},
	},
	"IMM": {
		{
			ID:          10450,
			Slug:        "imm",
			Name:        "IMM",
			BinanceUSDT: "IMMUSDT",
		},
	},
	"AC": {
		{
			ID:          7382,
			Slug:        "acoconut",
			Name:        "ACoconut",
			BinanceUSDT: "ACUSDT",
		},
	},
	"PIPI": {
		{
			ID:          9502,
			Slug:        "pippi-finance",
			Name:        "Pippi Finance",
			BinanceUSDT: "PIPIUSDT",
		},
	},
	"RAKU": {
		{
			ID:          5334,
			Slug:        "rakun",
			Name:        "RAKUN",
			BinanceUSDT: "RAKUUSDT",
		},
	},
	"LIVE": {
		{
			ID:          6889,
			Slug:        "tronbetlive",
			Name:        "TRONbetLive",
			BinanceUSDT: "LIVEUSDT",
		},
	},
	"LND": {
		{
			ID:          2686,
			Slug:        "lendingblock",
			Name:        "Lendingblock",
			BinanceUSDT: "LNDUSDT",
		},
	},
	"TBP": {
		{
			ID:          7777,
			Slug:        "tradebitpay",
			Name:        "Tradebitpay",
			BinanceUSDT: "TBPUSDT",
		},
	},
	"XGM": {
		{
			ID:          5847,
			Slug:        "defis",
			Name:        "Defis",
			BinanceUSDT: "XGMUSDT",
		},
	},
	"NIIFI": {
		{
			ID:          9825,
			Slug:        "niifi",
			Name:        "NiiFi",
			BinanceUSDT: "NIIFIUSDT",
		},
	},
	"ALPHR": {
		{
			ID:          9430,
			Slug:        "alphr-finance",
			Name:        "Alphr finance",
			BinanceUSDT: "ALPHRUSDT",
		},
	},
	"SCRIV": {
		{
			ID:          3398,
			Slug:        "scriv-network",
			Name:        "SCRIV NETWORK",
			BinanceUSDT: "SCRIVUSDT",
		},
	},
	"USDQ": {
		{
			ID:          4020,
			Slug:        "usdq",
			Name:        "USDQ",
			BinanceUSDT: "USDQUSDT",
		},
	},
	"K21": {
		{
			ID:          9205,
			Slug:        "k21",
			Name:        "K21",
			BinanceUSDT: "K21USDT",
		},
	},
	"BLZN": {
		{
			ID:          6723,
			Slug:        "blaze-network",
			Name:        "Blaze Network",
			BinanceUSDT: "BLZNUSDT",
		},
	},
	"EASY": {
		{
			ID:          7332,
			Slug:        "easyfi",
			Name:        "EasyFi",
			BinanceUSDT: "EASYUSDT",
		},
	},
	"IPM": {
		{
			ID:          7247,
			Slug:        "timers",
			Name:        "Timers",
			BinanceUSDT: "IPMUSDT",
		},
	},
	"BURNS": {
		{
			ID:          10734,
			Slug:        "mr-burn-token",
			Name:        "Mr Burn Token",
			BinanceUSDT: "BURNSUSDT",
		},
	},
	"SOCKS": {
		{
			ID:          7095,
			Slug:        "unisocks",
			Name:        "Unisocks",
			BinanceUSDT: "SOCKSUSDT",
		},
	},
	"AKT": {
		{
			ID:          7431,
			Slug:        "akash-network",
			Name:        "Akash Network",
			BinanceUSDT: "AKTUSDT",
		},
	},
	"AIOZ": {
		{
			ID:          9104,
			Slug:        "aioz-network",
			Name:        "AIOZ Network",
			BinanceUSDT: "AIOZUSDT",
		},
	},
	"MORPH": {
		{
			ID:          9408,
			Slug:        "morphose",
			Name:        "MORPHOSE",
			BinanceUSDT: "MORPHUSDT",
		},
	},
	"IBS": {
		{
			ID:          5358,
			Slug:        "ibstoken",
			Name:        "IBStoken",
			BinanceUSDT: "IBSUSDT",
		},
	},
	"TLB": {
		{
			ID:          7544,
			Slug:        "the-luxury-coin",
			Name:        "The Luxury Coin",
			BinanceUSDT: "TLBUSDT",
		},
	},
	"NOR": {
		{
			ID:          3611,
			Slug:        "noir",
			Name:        "Noir",
			BinanceUSDT: "NORUSDT",
		},
	},
	"INN": {
		{
			ID:          2160,
			Slug:        "innova",
			Name:        "Innova",
			BinanceUSDT: "INNUSDT",
		},
	},
	"LUX": {
		{
			ID:          2107,
			Slug:        "luxcoin",
			Name:        "LUXCoin",
			BinanceUSDT: "LUXUSDT",
		},
	},
	"TBX": {
		{
			ID:          2452,
			Slug:        "tokenbox",
			Name:        "Tokenbox",
			BinanceUSDT: "TBXUSDT",
		},
	},
	"NAVI": {
		{
			ID:          8340,
			Slug:        "natus-vincere-fan-token",
			Name:        "Natus Vincere Fan Token",
			BinanceUSDT: "NAVIUSDT",
		},
	},
	"AINU": {
		{
			ID:          10877,
			Slug:        "ainu-token",
			Name:        "Ainu Token",
			BinanceUSDT: "AINUUSDT",
		},
	},
	"CVR": {
		{
			ID:          8300,
			Slug:        "polkacover",
			Name:        "Polkacover",
			BinanceUSDT: "CVRUSDT",
		},
	},
	"ION": {
		{
			ID:          1281,
			Slug:        "ion",
			Name:        "ION",
			BinanceUSDT: "IONUSDT",
		},
	},
	"DC": {
		{
			ID:          8511,
			Slug:        "deep-coin",
			Name:        "DeepCoin",
			BinanceUSDT: "DCUSDT",
		},
	},
	"R2R": {
		{
			ID:          5204,
			Slug:        "citios",
			Name:        "CitiOs",
			BinanceUSDT: "R2RUSDT",
		},
	},
	"VOLTZ": {
		{
			ID:          4508,
			Slug:        "voltz",
			Name:        "Voltz",
			BinanceUSDT: "VOLTZUSDT",
		},
	},
	"GCX": {
		{
			ID:          5322,
			Slug:        "germancoin",
			Name:        "GermanCoin",
			BinanceUSDT: "GCXUSDT",
		},
	},
	"INCNT": {
		{
			ID:          1475,
			Slug:        "incent",
			Name:        "Incent",
			BinanceUSDT: "INCNTUSDT",
		},
	},
	"BCP": {
		{
			ID:          7867,
			Slug:        "bitcashpay",
			Name:        "Bitcashpay",
			BinanceUSDT: "BCPUSDT",
		},
	},
	"XPA": {
		{
			ID:          1968,
			Slug:        "xpa",
			Name:        "XPA",
			BinanceUSDT: "XPAUSDT",
		},
	},
	"CENT": {
		{
			ID:          3681,
			Slug:        "centercoin",
			Name:        "CENTERCOIN",
			BinanceUSDT: "CENTUSDT",
		},
	},
	"TMT": {
		{
			ID:          7246,
			Slug:        "tamy-token",
			Name:        "Tamy Token",
			BinanceUSDT: "TMTUSDT",
		},
	},
	"DSC": {
		{
			ID:          4836,
			Slug:        "dash-cash",
			Name:        "Dash Cash",
			BinanceUSDT: "DSCUSDT",
		},
	},
	"PURE": {
		{
			ID:          7203,
			Slug:        "puriever",
			Name:        "Puriever",
			BinanceUSDT: "PUREUSDT",
		},
	},
	"mNFLX": {
		{
			ID:          8005,
			Slug:        "mirrored-netflix",
			Name:        "Mirrored Netflix",
			BinanceUSDT: "mNFLXUSDT",
		},
	},
	"ESTI": {
		{
			ID:          6760,
			Slug:        "easticoin",
			Name:        "Easticoin",
			BinanceUSDT: "ESTIUSDT",
		},
	},
	"BAFI": {
		{
			ID:          9075,
			Slug:        "bafi-finance",
			Name:        "Bafi Finance",
			BinanceUSDT: "BAFIUSDT",
		},
	},
	"HRD": {
		{
			ID:          7801,
			Slug:        "hrdcoin",
			Name:        "HRDCOIN",
			BinanceUSDT: "HRDUSDT",
		},
	},
	"MINA": {
		{
			ID:          8646,
			Slug:        "mina",
			Name:        "Mina",
			BinanceUSDT: "MINAUSDT",
		},
	},
	"PAD": {
		{
			ID:          10610,
			Slug:        "smartpad",
			Name:        "SMARTPAD",
			BinanceUSDT: "PADUSDT",
		},
	},
	"QDX": {
		{
			ID:          10079,
			Slug:        "quidax",
			Name:        "Quidax",
			BinanceUSDT: "QDXUSDT",
		},
	},
	"BOOMC": {
		{
			ID:          10575,
			Slug:        "boomcoin",
			Name:        "BoomCoin",
			BinanceUSDT: "BOOMCUSDT",
		},
	},
	"PXG": {
		{
			ID:          3639,
			Slug:        "playgame",
			Name:        "PlayGame",
			BinanceUSDT: "PXGUSDT",
		},
	},
	"RAVEN": {
		{
			ID:          4024,
			Slug:        "raven-protocol",
			Name:        "Raven Protocol",
			BinanceUSDT: "RAVENUSDT",
		},
	},
	"INE": {
		{
			ID:          3811,
			Slug:        "intellishare",
			Name:        "IntelliShare",
			BinanceUSDT: "INEUSDT",
		},
	},
	"TEM": {
		{
			ID:          5945,
			Slug:        "temtum",
			Name:        "Temtum",
			BinanceUSDT: "TEMUSDT",
		},
	},
	"CHAINCADE": {
		{
			ID:          10875,
			Slug:        "chaincade",
			Name:        "ChainCade",
			BinanceUSDT: "CHAINCADEUSDT",
		},
	},
	"NFTP": {
		{
			ID:          8828,
			Slug:        "nft-pool",
			Name:        "NFT POOL",
			BinanceUSDT: "NFTPUSDT",
		},
	},
	"INSUR": {
		{
			ID:          8799,
			Slug:        "insurace",
			Name:        "InsurAce",
			BinanceUSDT: "INSURUSDT",
		},
	},
	"BAKE": {
		{
			ID:          7064,
			Slug:        "bakerytoken",
			Name:        "BakeryToken",
			BinanceUSDT: "BAKEUSDT",
		},
	},
	"EFL": {
		{
			ID:          234,
			Slug:        "e-gulden",
			Name:        "e-Gulden",
			BinanceUSDT: "EFLUSDT",
		},
	},
	"ELEC": {
		{
			ID:          2573,
			Slug:        "electrifyasia",
			Name:        "Electrify.Asia",
			BinanceUSDT: "ELECUSDT",
		},
	},
	"DSLA": {
		{
			ID:          5423,
			Slug:        "dsla-protocol",
			Name:        "DSLA Protocol",
			BinanceUSDT: "DSLAUSDT",
		},
	},
	"NGM": {
		{
			ID:          8279,
			Slug:        "e-money-coin",
			Name:        "e-Money",
			BinanceUSDT: "NGMUSDT",
		},
	},
	"MMON": {
		{
			ID:          10480,
			Slug:        "mammon",
			Name:        "Mammon",
			BinanceUSDT: "MMONUSDT",
		},
	},
	"UDOGE": {
		{
			ID:          10788,
			Slug:        "uncle-doge",
			Name:        "Uncle Doge",
			BinanceUSDT: "UDOGEUSDT",
		},
	},
	"GIG": {
		{
			ID:          3998,
			Slug:        "krios",
			Name:        "Krios",
			BinanceUSDT: "GIGUSDT",
		},
	},
	"YFICG": {
		{
			ID:          7517,
			Slug:        "yfi-credits-group",
			Name:        "YFI CREDITS GROUP",
			BinanceUSDT: "YFICGUSDT",
		},
	},
	"FLOAT": {
		{
			ID:          9861,
			Slug:        "float-protocol-float",
			Name:        "Float Protocol: Float",
			BinanceUSDT: "FLOATUSDT",
		},
	},
	"AIDOC": {
		{
			ID:          2357,
			Slug:        "aidoc",
			Name:        "AI Doctor",
			BinanceUSDT: "AIDOCUSDT",
		},
	},
	"xBTC": {
		{
			ID:          7168,
			Slug:        "xbtc",
			Name:        "xBTC",
			BinanceUSDT: "xBTCUSDT",
		},
	},
	"WOM": {
		{
			ID:          5328,
			Slug:        "wom-protocol",
			Name:        "WOM Protocol",
			BinanceUSDT: "WOMUSDT",
		},
	},
	"SHR": {
		{
			ID:          4197,
			Slug:        "sharetoken",
			Name:        "ShareToken",
			BinanceUSDT: "SHRUSDT",
		},
	},
	"CPR": {
		{
			ID:          4589,
			Slug:        "cipher",
			Name:        "Cipher",
			BinanceUSDT: "CPRUSDT",
		},
	},
	"HENTAI": {
		{
			ID:          10667,
			Slug:        "hentaicoin",
			Name:        "HentaiCoin",
			BinanceUSDT: "HENTAIUSDT",
		},
	},
	"ZDX": {
		{
			ID:          4184,
			Slug:        "zer-dex",
			Name:        "Zer-Dex",
			BinanceUSDT: "ZDXUSDT",
		},
	},
	"IUT": {
		{
			ID:          4737,
			Slug:        "ito-utility-token",
			Name:        "ITO Utility Token",
			BinanceUSDT: "IUTUSDT",
		},
	},
	"SAT": {
		{
			ID:   6533,
			Slug: "smartxone",
			Name: "SmartX",
		},
		{
			ID:   10176,
			Slug: "saturna",
			Name: "Saturna",
		},
	},
	"STORJ": {
		{
			ID:          1772,
			Slug:        "storj",
			Name:        "Storj",
			BinanceUSDT: "STORJUSDT",
		},
	},
	"R34P": {
		{
			ID:          8040,
			Slug:        "r34p",
			Name:        "R34P",
			BinanceUSDT: "R34PUSDT",
		},
	},
	"ROCK2": {
		{
			ID:          3065,
			Slug:        "ice-rock-mining",
			Name:        "ICE ROCK MINING",
			BinanceUSDT: "ROCK2USDT",
		},
	},
	"mMSFT": {
		{
			ID:          8017,
			Slug:        "mirrored-microsoft",
			Name:        "Mirrored Microsoft",
			BinanceUSDT: "mMSFTUSDT",
		},
	},
	"ARTX": {
		{
			ID:          9105,
			Slug:        "artx-trading",
			Name:        "ARTX Trading",
			BinanceUSDT: "ARTXUSDT",
		},
	},
	"NIA": {
		{
			ID:          8314,
			Slug:        "nydronia",
			Name:        "Nydronia",
			BinanceUSDT: "NIAUSDT",
		},
	},
	"MOF": {
		{
			ID:          2441,
			Slug:        "molecular-future",
			Name:        "Molecular Future",
			BinanceUSDT: "MOFUSDT",
		},
	},
	"CLVA": {
		{
			ID:          8228,
			Slug:        "clever-defi",
			Name:        "Clever DeFi",
			BinanceUSDT: "CLVAUSDT",
		},
	},
	"INTX": {
		{
			ID:          5914,
			Slug:        "intexcoin",
			Name:        "intexcoin",
			BinanceUSDT: "INTXUSDT",
		},
	},
	"BOXER": {
		{
			ID:          10972,
			Slug:        "boxer-inu",
			Name:        "Boxer Inu",
			BinanceUSDT: "BOXERUSDT",
		},
	},
	"CNRG": {
		{
			ID:          5376,
			Slug:        "cryptoenergy",
			Name:        "CryptoEnergy",
			BinanceUSDT: "CNRGUSDT",
		},
	},
	"DVF": {
		{
			ID:          10759,
			Slug:        "deversifi",
			Name:        "DeversiFi",
			BinanceUSDT: "DVFUSDT",
		},
	},
	"AIX": {
		{
			ID:          2367,
			Slug:        "aigang",
			Name:        "Aigang",
			BinanceUSDT: "AIXUSDT",
		},
	},
	"sBTC": {
		{
			ID:          6406,
			Slug:        "softbtc",
			Name:        "sBTC",
			BinanceUSDT: "sBTCUSDT",
		},
	},
	"OROS": {
		{
			ID:          10842,
			Slug:        "oros-finance",
			Name:        "OROS.finance",
			BinanceUSDT: "OROSUSDT",
		},
	},
	"MAUSDC": {
		{
			ID:          8918,
			Slug:        "matic-aave-usdc",
			Name:        "Matic Aave Interest Bearing USDC",
			BinanceUSDT: "MAUSDCUSDT",
		},
	},
	"QRL": {
		{
			ID:          1712,
			Slug:        "quantum-resistant-ledger",
			Name:        "Quantum Resistant Ledger",
			BinanceUSDT: "QRLUSDT",
		},
	},
	"RUC": {
		{
			ID:          7423,
			Slug:        "rush",
			Name:        "RUSH",
			BinanceUSDT: "RUCUSDT",
		},
	},
	"DYT": {
		{
			ID:          5894,
			Slug:        "doyourtip",
			Name:        "DoYourTip",
			BinanceUSDT: "DYTUSDT",
		},
	},
	"TRA": {
		{
			ID:          7637,
			Slug:        "trabzonspor-fan-token",
			Name:        "Trabzonspor Fan Token",
			BinanceUSDT: "TRAUSDT",
		},
	},
	"MNE": {
		{
			ID:          1673,
			Slug:        "minereum",
			Name:        "Minereum",
			BinanceUSDT: "MNEUSDT",
		},
	},
	"DGMT": {
		{
			ID:          6025,
			Slug:        "digimax-dgmt",
			Name:        "DigiMax DGMT",
			BinanceUSDT: "DGMTUSDT",
		},
	},
	"DEGO": {
		{
			ID:          7087,
			Slug:        "dego-finance",
			Name:        "Dego Finance",
			BinanceUSDT: "DEGOUSDT",
		},
	},
	"HUSL": {
		{
			ID:          5253,
			Slug:        "the-hustle-app",
			Name:        "The Hustle App",
			BinanceUSDT: "HUSLUSDT",
		},
	},
	"ORE": {
		{
			ID:   8929,
			Slug: "oreo",
			Name: "OREO",
		},
		{
			ID:   2411,
			Slug: "galactrum",
			Name: "Galactrum",
		},
	},
	"ZEE": {
		{
			ID:          7438,
			Slug:        "zeroswap",
			Name:        "ZeroSwap",
			BinanceUSDT: "ZEEUSDT",
		},
	},
	"WEX": {
		{
			ID:          9951,
			Slug:        "waultswap",
			Name:        "WaultSwap",
			BinanceUSDT: "WEXUSDT",
		},
	},
	"GAINZ": {
		{
			ID:          10587,
			Slug:        "gainz-token",
			Name:        "GAINZ TOKEN",
			BinanceUSDT: "GAINZUSDT",
		},
	},
	"YFIB": {
		{
			ID:          6875,
			Slug:        "yfibusiness",
			Name:        "YFIBusiness",
			BinanceUSDT: "YFIBUSDT",
		},
	},
	"BSVBEAR": {
		{
			ID:          5459,
			Slug:        "3x-short-bitcoin-sv-token",
			Name:        "3x Short Bitcoin SV Token",
			BinanceUSDT: "BSVBEARUSDT",
		},
	},
	"ROC": {
		{
			ID:          8768,
			Slug:        "roxe-cash",
			Name:        "Roxe Cash",
			BinanceUSDT: "ROCUSDT",
		},
	},
	"RSUN": {
		{
			ID:          11004,
			Slug:        "risingsun",
			Name:        "RisingSun",
			BinanceUSDT: "RSUNUSDT",
		},
	},
	"HOPR": {
		{
			ID:          6520,
			Slug:        "hopr",
			Name:        "HOPR",
			BinanceUSDT: "HOPRUSDT",
		},
	},
	"BLC": {
		{
			ID:          125,
			Slug:        "blakecoin",
			Name:        "Blakecoin",
			BinanceUSDT: "BLCUSDT",
		},
	},
	"LEMO": {
		{
			ID:          2950,
			Slug:        "lemochain",
			Name:        "LemoChain",
			BinanceUSDT: "LEMOUSDT",
		},
	},
	"BPT": {
		{
			ID:   10402,
			Slug: "blackpool",
			Name: "BlackPool",
		},
		{
			ID:   5850,
			Slug: "bitpayer-token",
			Name: "Bitpayer Token",
		},
		{
			ID:   7802,
			Slug: "bitpumps-token",
			Name: "Bitpumps Token",
		},
	},
	"ORAO": {
		{
			ID:          8895,
			Slug:        "orao-network",
			Name:        "ORAO Network",
			BinanceUSDT: "ORAOUSDT",
		},
	},
	"KISHU": {
		{
			ID:          9386,
			Slug:        "kishu-inu",
			Name:        "Kishu Inu",
			BinanceUSDT: "KISHUUSDT",
		},
	},
	"DESH": {
		{
			ID:          5808,
			Slug:        "decash",
			Name:        "DeCash",
			BinanceUSDT: "DESHUSDT",
		},
	},
	"COPS": {
		{
			ID:          8947,
			Slug:        "cops-finance",
			Name:        "COPS FINANCE",
			BinanceUSDT: "COPSUSDT",
		},
	},
	"DON": {
		{
			ID:   8814,
			Slug: "donnie-finance",
			Name: "Donnie Finance",
		},
		{
			ID:   9643,
			Slug: "don-key",
			Name: "Don-key",
		},
		{
			ID:   6732,
			Slug: "deonex-coin",
			Name: "DEONEX COIN",
		},
	},
	"XOR": {
		{
			ID:   5802,
			Slug: "sora",
			Name: "Sora",
		},
		{
			ID:   7074,
			Slug: "oracolxor",
			Name: "Oracolxor",
		},
	},
	"BAEPAY": {
		{
			ID:          7657,
			Slug:        "baepay",
			Name:        "BAEPAY",
			BinanceUSDT: "BAEPAYUSDT",
		},
	},
	"STEEM": {
		{
			ID:          1230,
			Slug:        "steem",
			Name:        "Steem",
			BinanceUSDT: "STEEMUSDT",
		},
	},
	"ZWAP": {
		{
			ID:          9107,
			Slug:        "zilswap",
			Name:        "Zilswap",
			BinanceUSDT: "ZWAPUSDT",
		},
	},
	"RVX": {
		{
			ID:          4697,
			Slug:        "rivex",
			Name:        "Rivex",
			BinanceUSDT: "RVXUSDT",
		},
	},
	"NNB": {
		{
			ID:          3937,
			Slug:        "nnb-token",
			Name:        "NNB Token",
			BinanceUSDT: "NNBUSDT",
		},
	},
	"USDJ": {
		{
			ID:          5446,
			Slug:        "usdj",
			Name:        "USDJ",
			BinanceUSDT: "USDJUSDT",
		},
	},
	"BDO": {
		{
			ID:          8219,
			Slug:        "bdollar",
			Name:        "bDollar",
			BinanceUSDT: "BDOUSDT",
		},
	},
	"IOI": {
		{
			ID:          10295,
			Slug:        "trade-race-manager",
			Name:        "IOI Token",
			BinanceUSDT: "IOIUSDT",
		},
	},
	"APN": {
		{
			ID:          8406,
			Slug:        "apron-network",
			Name:        "Apron Network",
			BinanceUSDT: "APNUSDT",
		},
	},
	"HOT": {
		{
			ID:   2682,
			Slug: "holo",
			Name: "Holo",
		},
		{
			ID:   2430,
			Slug: "hydro-protocol",
			Name: "Hydro Protocol",
		},
	},
	"WAN": {
		{
			ID:          2606,
			Slug:        "wanchain",
			Name:        "Wanchain",
			BinanceUSDT: "WANUSDT",
		},
	},
	"DFST": {
		{
			ID:          8446,
			Slug:        "defistarter",
			Name:        "DeFiStarter",
			BinanceUSDT: "DFSTUSDT",
		},
	},
	"TANGO": {
		{
			ID:          8712,
			Slug:        "keytango",
			Name:        "keyTango",
			BinanceUSDT: "TANGOUSDT",
		},
	},
	"PANDA": {
		{
			ID:          10131,
			Slug:        "hashpanda",
			Name:        "HashPanda",
			BinanceUSDT: "PANDAUSDT",
		},
	},
	"TURTLE": {
		{
			ID:          10757,
			Slug:        "turtle",
			Name:        "Turtle",
			BinanceUSDT: "TURTLEUSDT",
		},
	},
	"PEPS": {
		{
			ID:          4397,
			Slug:        "peps-coin",
			Name:        "PEPS Coin",
			BinanceUSDT: "PEPSUSDT",
		},
	},
	"TRU": {
		{
			ID:          7725,
			Slug:        "truefi-token",
			Name:        "TrueFi",
			BinanceUSDT: "TRUUSDT",
		},
	},
	"TERN": {
		{
			ID:   2876,
			Slug: "ternio",
			Name: "Ternio",
		},
		{
			ID:   7155,
			Slug: "ternio-erc20",
			Name: "Ternio-ERC20",
		},
	},
	"TIME": {
		{
			ID:          1556,
			Slug:        "chrono-tech",
			Name:        "Chrono.tech",
			BinanceUSDT: "TIMEUSDT",
		},
	},
	"DASHG": {
		{
			ID:          3613,
			Slug:        "dash-green",
			Name:        "Dash Green",
			BinanceUSDT: "DASHGUSDT",
		},
	},
	"HTDF": {
		{
			ID:          5127,
			Slug:        "orient-walt",
			Name:        "Orient Walt",
			BinanceUSDT: "HTDFUSDT",
		},
	},
	"VIG": {
		{
			ID:          6562,
			Slug:        "vig",
			Name:        "VIG",
			BinanceUSDT: "VIGUSDT",
		},
	},
	"$BOOB": {
		{
			ID:          10144,
			Slug:        "boob",
			Name:        "$BOOB",
			BinanceUSDT: "$BOOBUSDT",
		},
	},
	"JST": {
		{
			ID:          5488,
			Slug:        "just",
			Name:        "JUST",
			BinanceUSDT: "JSTUSDT",
		},
	},
	"TRST": {
		{
			ID:          1638,
			Slug:        "trust",
			Name:        "WeTrust",
			BinanceUSDT: "TRSTUSDT",
		},
	},
	"GBX": {
		{
			ID:   2200,
			Slug: "gobyte",
			Name: "GoByte",
		},
		{
			ID:   6450,
			Slug: "gbrick",
			Name: "Gbrick",
		},
	},
	"ACU": {
		{
			ID:   4260,
			Slug: "aitheon",
			Name: "Aitheon",
		},
		{
			ID:   7167,
			Slug: "acuity-token",
			Name: "Acuity Token",
		},
	},
	"AZ": {
		{
			ID:          4777,
			Slug:        "azbit",
			Name:        "Azbit",
			BinanceUSDT: "AZUSDT",
		},
	},
	"PYRO": {
		{
			ID:          5169,
			Slug:        "pyro-network",
			Name:        "PYRO Network",
			BinanceUSDT: "PYROUSDT",
		},
	},
	"DAPS": {
		{
			ID:          3345,
			Slug:        "daps-coin",
			Name:        "DAPS Coin",
			BinanceUSDT: "DAPSUSDT",
		},
	},
	"WNL": {
		{
			ID:          4138,
			Slug:        "winstars-live",
			Name:        "WinStars.live",
			BinanceUSDT: "WNLUSDT",
		},
	},
	"BAS": {
		{
			ID:          7816,
			Slug:        "basis-share",
			Name:        "Basis Share",
			BinanceUSDT: "BASUSDT",
		},
	},
	"TKNT": {
		{
			ID:          7238,
			Slug:        "tkn-token",
			Name:        "TKN Token",
			BinanceUSDT: "TKNTUSDT",
		},
	},
	"ETF": {
		{
			ID:          5441,
			Slug:        "entherfound",
			Name:        "Entherfound",
			BinanceUSDT: "ETFUSDT",
		},
	},
	"MAP": {
		{
			ID:          4956,
			Slug:        "map-protocol",
			Name:        "MAP Protocol",
			BinanceUSDT: "MAPUSDT",
		},
	},
	"VIBE": {
		{
			ID:          1983,
			Slug:        "vibe",
			Name:        "VIBE",
			BinanceUSDT: "VIBEUSDT",
		},
	},
	"TONS": {
		{
			ID:          7060,
			Slug:        "thisoption",
			Name:        "Thisoption",
			BinanceUSDT: "TONSUSDT",
		},
	},
	"XPB": {
		{
			ID:          8783,
			Slug:        "transmute-protocol",
			Name:        "Transmute Protocol",
			BinanceUSDT: "XPBUSDT",
		},
	},
	"KUV": {
		{
			ID:          4940,
			Slug:        "kuverit",
			Name:        "Kuverit",
			BinanceUSDT: "KUVUSDT",
		},
	},
	"KIDS": {
		{
			ID:          10318,
			Slug:        "save-the-kids",
			Name:        "Save The Kids",
			BinanceUSDT: "KIDSUSDT",
		},
	},
	"WEC": {
		{
			ID:   4668,
			Slug: "wave-edu-coin",
			Name: "wave edu coin",
		},
		{
			ID:   6548,
			Slug: "web-coin-pay",
			Name: "Web Coin Pay",
		},
		{
			ID:   9247,
			Slug: "whole-earth-coin",
			Name: "Whole Earth Coin",
		},
	},
	"GIX": {
		{
			ID:          7624,
			Slug:        "goldfinx",
			Name:        "GoldFinX",
			BinanceUSDT: "GIXUSDT",
		},
	},
	"SIB": {
		{
			ID:          1082,
			Slug:        "sibcoin",
			Name:        "SIBCoin",
			BinanceUSDT: "SIBUSDT",
		},
	},
	"ZYN": {
		{
			ID:          4951,
			Slug:        "zynecoin",
			Name:        "Zynecoin",
			BinanceUSDT: "ZYNUSDT",
		},
	},
	"CLC": {
		{
			ID:          4384,
			Slug:        "caluracoin",
			Name:        "CaluraCoin",
			BinanceUSDT: "CLCUSDT",
		},
	},
	"UNIT": {
		{
			ID:   921,
			Slug: "universal-currency",
			Name: "Universal Currency",
		},
		{
			ID:   9137,
			Slug: "uniti-protocol",
			Name: "UNITi Protocol",
		},
	},
	"BSL": {
		{
			ID:          10036,
			Slug:        "bsclaunch",
			Name:        "BSClaunch",
			BinanceUSDT: "BSLUSDT",
		},
	},
	"NDSK": {
		{
			ID:          9677,
			Slug:        "nadeshiko",
			Name:        "Nadeshiko",
			BinanceUSDT: "NDSKUSDT",
		},
	},
	"MAVRO": {
		{
			ID:          1599,
			Slug:        "mavro",
			Name:        "Mavro",
			BinanceUSDT: "MAVROUSDT",
		},
	},
	"VBit": {
		{
			ID:          10082,
			Slug:        "voltbit",
			Name:        "Voltbit",
			BinanceUSDT: "VBitUSDT",
		},
	},
	"CWT": {
		{
			ID:   10530,
			Slug: "crosswallet",
			Name: "CrossWallet",
		},
		{
			ID:   10807,
			Slug: "coinw-token",
			Name: "CoinW Token",
		},
	},
	"HUNNY": {
		{
			ID:          10514,
			Slug:        "pancake-hunny",
			Name:        "PANCAKE HUNNY",
			BinanceUSDT: "HUNNYUSDT",
		},
	},
	"TOTO": {
		{
			ID:          2960,
			Slug:        "tourist-token",
			Name:        "Tourist Token",
			BinanceUSDT: "TOTOUSDT",
		},
	},
	"CMK": {
		{
			ID:          10485,
			Slug:        "credmark",
			Name:        "Credmark",
			BinanceUSDT: "CMKUSDT",
		},
	},
	"WTN": {
		{
			ID:          3484,
			Slug:        "waletoken",
			Name:        "Waletoken",
			BinanceUSDT: "WTNUSDT",
		},
	},
	"IDV": {
		{
			ID:          8726,
			Slug:        "idavoll-network",
			Name:        "Idavoll Network",
			BinanceUSDT: "IDVUSDT",
		},
	},
	"HPY": {
		{
			ID:          2329,
			Slug:        "hyper-pay",
			Name:        "Hyper Pay",
			BinanceUSDT: "HPYUSDT",
		},
	},
	"PFE": {
		{
			ID:   7891,
			Slug: "pfizer-tokenized-stock-ftx",
			Name: "Pfizer tokenized stock FTX",
		},
		{
			ID:   7923,
			Slug: "pfizer-tokenized-stock-bittrex",
			Name: "Pfizer tokenized stock Bittrex",
		},
	},
	"CLA": {
		{
			ID:          6978,
			Slug:        "candelacoin",
			Name:        "Candela Coin",
			BinanceUSDT: "CLAUSDT",
		},
	},
	"LEOS": {
		{
			ID:          10762,
			Slug:        "leonicorn-swap",
			Name:        "Leonicorn Swap",
			BinanceUSDT: "LEOSUSDT",
		},
	},
	"BABYDOGE": {
		{
			ID:          10970,
			Slug:        "babydoge-coin",
			Name:        "BabyDoge ETH",
			BinanceUSDT: "BABYDOGEUSDT",
		},
	},
	"LEV": {
		{
			ID:          9505,
			Slug:        "lever-token",
			Name:        "Lever Token",
			BinanceUSDT: "LEVUSDT",
		},
	},
	"TWTR": {
		{
			ID:          7916,
			Slug:        "twitter-tokenized-stock-ftx",
			Name:        "Twitter tokenized stock FTX",
			BinanceUSDT: "TWTRUSDT",
		},
	},
	"BHC": {
		{
			ID:          7182,
			Slug:        "billionhappiness",
			Name:        "BillionHappiness",
			BinanceUSDT: "BHCUSDT",
		},
	},
	"BULLY": {
		{
			ID:          10243,
			Slug:        "pitbully",
			Name:        "PitBULLY",
			BinanceUSDT: "BULLYUSDT",
		},
	},
	"RTE": {
		{
			ID:          2880,
			Slug:        "rate3",
			Name:        "Rate3",
			BinanceUSDT: "RTEUSDT",
		},
	},
	"JADE": {
		{
			ID:          5025,
			Slug:        "jade-currency",
			Name:        "Jade Currency",
			BinanceUSDT: "JADEUSDT",
		},
	},
	"KEMA": {
		{
			ID:          5236,
			Slug:        "kemacoin",
			Name:        "Kemacoin",
			BinanceUSDT: "KEMAUSDT",
		},
	},
	"ARKK": {
		{
			ID:          7896,
			Slug:        "ark-innovation-etf-tokenized-stock-ftx",
			Name:        "ARK Innovation ETF tokenized stock FTX",
			BinanceUSDT: "ARKKUSDT",
		},
	},
	"OLY": {
		{
			ID:          8484,
			Slug:        "olyseum",
			Name:        "Olyseum",
			BinanceUSDT: "OLYUSDT",
		},
	},
	"INTO": {
		{
			ID:          10054,
			Slug:        "infiniti",
			Name:        "Infiniti",
			BinanceUSDT: "INTOUSDT",
		},
	},
	"UTP": {
		{
			ID:          10286,
			Slug:        "utopian-protocol",
			Name:        "Utopian Protocol",
			BinanceUSDT: "UTPUSDT",
		},
	},
	"aENJ": {
		{
			ID:          8594,
			Slug:        "aave-enjin",
			Name:        "Aave Enjin",
			BinanceUSDT: "aENJUSDT",
		},
	},
	"CMT": {
		{
			ID:   2246,
			Slug: "cybermiles",
			Name: "CyberMiles",
		},
		{
			ID:   1291,
			Slug: "comet",
			Name: "Comet",
		},
	},
	"INO": {
		{
			ID:          3085,
			Slug:        "ino-coin",
			Name:        "INO COIN",
			BinanceUSDT: "INOUSDT",
		},
	},
	"sLINK": {
		{
			ID:          6190,
			Slug:        "slink",
			Name:        "sLINK",
			BinanceUSDT: "sLINKUSDT",
		},
	},
	"NXT": {
		{
			ID:          66,
			Slug:        "nxt",
			Name:        "Nxt",
			BinanceUSDT: "NXTUSDT",
		},
	},
	"PCL": {
		{
			ID:          2610,
			Slug:        "peculium",
			Name:        "Peculium",
			BinanceUSDT: "PCLUSDT",
		},
	},
	"PHA": {
		{
			ID:          6841,
			Slug:        "phala-network",
			Name:        "Phala Network",
			BinanceUSDT: "PHAUSDT",
		},
	},
	"BFT": {
		{
			ID:   2605,
			Slug: "bnktothefuture",
			Name: "BnkToTheFuture",
		},
		{
			ID:   4643,
			Slug: "bitget-defi-token",
			Name: "Bitget DeFi Token",
		},
		{
			ID:   10534,
			Slug: "bitfresh",
			Name: "Bitfresh",
		},
	},
	"JCC": {
		{
			ID:          6914,
			Slug:        "junca-cash",
			Name:        "junca Cash",
			BinanceUSDT: "JCCUSDT",
		},
	},
	"ZIPT": {
		{
			ID:          2724,
			Slug:        "zippie",
			Name:        "Zippie",
			BinanceUSDT: "ZIPTUSDT",
		},
	},
	"CUR": {
		{
			ID:          5083,
			Slug:        "curio",
			Name:        "Curio",
			BinanceUSDT: "CURUSDT",
		},
	},
	"NAME": {
		{
			ID:          9264,
			Slug:        "polkadomain",
			Name:        "PolkaDomain",
			BinanceUSDT: "NAMEUSDT",
		},
	},
	"QCH": {
		{
			ID:          3337,
			Slug:        "qchi",
			Name:        "QChi",
			BinanceUSDT: "QCHUSDT",
		},
	},
	"PRINCESS": {
		{
			ID:          10849,
			Slug:        "paris-inuton",
			Name:        "Paris Inuton",
			BinanceUSDT: "PRINCESSUSDT",
		},
	},
	"LANA": {
		{
			ID:          1257,
			Slug:        "lanacoin",
			Name:        "LanaCoin",
			BinanceUSDT: "LANAUSDT",
		},
	},
	"OBOT": {
		{
			ID:          9590,
			Slug:        "obortech",
			Name:        "OBORTECH",
			BinanceUSDT: "OBOTUSDT",
		},
	},
	"CRETH2": {
		{
			ID:          8893,
			Slug:        "cream-eth2",
			Name:        "Cream ETH 2",
			BinanceUSDT: "CRETH2USDT",
		},
	},
	"YNI": {
		{
			ID:          8725,
			Slug:        "yearnyfi-network",
			Name:        "YEARNYFI NETWORK",
			BinanceUSDT: "YNIUSDT",
		},
	},
	"YOYOW": {
		{
			ID:          1899,
			Slug:        "yoyow",
			Name:        "YOYOW",
			BinanceUSDT: "YOYOWUSDT",
		},
	},
	"BSCS": {
		{
			ID:          9345,
			Slug:        "bsc-station",
			Name:        "BSC Station",
			BinanceUSDT: "BSCSUSDT",
		},
	},
	"NBX": {
		{
			ID:          4297,
			Slug:        "netbox-coin",
			Name:        "Netbox Coin",
			BinanceUSDT: "NBXUSDT",
		},
	},
	"BLVR": {
		{
			ID:          5855,
			Slug:        "believer",
			Name:        "BELIEVER",
			BinanceUSDT: "BLVRUSDT",
		},
	},
	"XDEX": {
		{
			ID:          9087,
			Slug:        "xdefi",
			Name:        "xDeFi",
			BinanceUSDT: "XDEXUSDT",
		},
	},
	"BITX": {
		{
			ID:          3011,
			Slug:        "bitscreener-token",
			Name:        "BitScreener Token",
			BinanceUSDT: "BITXUSDT",
		},
	},
	"MIS": {
		{
			ID:   8121,
			Slug: "themis-oracle",
			Name: "Themis",
		},
		{
			ID:   8141,
			Slug: "mithril-share",
			Name: "Mithril Share",
		},
	},
	"WIKEN": {
		{
			ID:          4809,
			Slug:        "project-with",
			Name:        "Project WITH",
			BinanceUSDT: "WIKENUSDT",
		},
	},
	"RBLX": {
		{
			ID:          2689,
			Slug:        "rublix",
			Name:        "Rublix",
			BinanceUSDT: "RBLXUSDT",
		},
	},
	"YFX": {
		{
			ID:          9946,
			Slug:        "your-future-exchange",
			Name:        "Your Future Exchange",
			BinanceUSDT: "YFXUSDT",
		},
	},
	"ONC": {
		{
			ID:          8159,
			Slug:        "one-cash",
			Name:        "One Cash",
			BinanceUSDT: "ONCUSDT",
		},
	},
	"BOSON": {
		{
			ID:          8827,
			Slug:        "boson-protocol",
			Name:        "Boson Protocol",
			BinanceUSDT: "BOSONUSDT",
		},
	},
	"CNB": {
		{
			ID:          5114,
			Slug:        "coinsbit-token",
			Name:        "Coinsbit Token",
			BinanceUSDT: "CNBUSDT",
		},
	},
	"TCASH": {
		{
			ID:          3980,
			Slug:        "tcash",
			Name:        "TCASH",
			BinanceUSDT: "TCASHUSDT",
		},
	},
	"CHNG": {
		{
			ID:          9071,
			Slug:        "chainge",
			Name:        "Chainge",
			BinanceUSDT: "CHNGUSDT",
		},
	},
	"MAYFI": {
		{
			ID:          8921,
			Slug:        "matic-aave-yfi",
			Name:        "Matic Aave Interest Bearing YFI",
			BinanceUSDT: "MAYFIUSDT",
		},
	},
	"CFG": {
		{
			ID:          6748,
			Slug:        "centrifuge",
			Name:        "Centrifuge",
			BinanceUSDT: "CFGUSDT",
		},
	},
	"PICKLE": {
		{
			ID:          7022,
			Slug:        "pickle-finance",
			Name:        "Pickle Finance",
			BinanceUSDT: "PICKLEUSDT",
		},
	},
	"NFTL": {
		{
			ID:          8784,
			Slug:        "nftl-token",
			Name:        "NFTL Token",
			BinanceUSDT: "NFTLUSDT",
		},
	},
	"TULIP": {
		{
			ID:          10373,
			Slug:        "solfarm",
			Name:        "SolFarm",
			BinanceUSDT: "TULIPUSDT",
		},
	},
	"FXP": {
		{
			ID:          6477,
			Slug:        "fxpay",
			Name:        "FXPay",
			BinanceUSDT: "FXPUSDT",
		},
	},
	"MVC": {
		{
			ID:          7703,
			Slug:        "mileverse",
			Name:        "MileVerse",
			BinanceUSDT: "MVCUSDT",
		},
	},
	"SNB": {
		{
			ID:          5277,
			Slug:        "synchrobitcoin",
			Name:        "SynchroBitcoin",
			BinanceUSDT: "SNBUSDT",
		},
	},
	"HYDRA": {
		{
			ID:          8245,
			Slug:        "hydra",
			Name:        "Hydra",
			BinanceUSDT: "HYDRAUSDT",
		},
	},
	"SEED": {
		{
			ID:   9377,
			Slug: "treedefi",
			Name: "TreeDefi",
		},
		{
			ID:   10895,
			Slug: "seed",
			Name: "SEED",
		},
		{
			ID:   4980,
			Slug: "sesameseed",
			Name: "Sesameseed",
		},
	},
	"FLC": {
		{
			ID:          3727,
			Slug:        "flowchain",
			Name:        "Flowchain",
			BinanceUSDT: "FLCUSDT",
		},
	},
	"PCHART": {
		{
			ID:          10375,
			Slug:        "polychart",
			Name:        "Polychart",
			BinanceUSDT: "PCHARTUSDT",
		},
	},
	"LPT": {
		{
			ID:          3640,
			Slug:        "livepeer",
			Name:        "Livepeer",
			BinanceUSDT: "LPTUSDT",
		},
	},
	"BASID": {
		{
			ID:          7157,
			Slug:        "basid-coin",
			Name:        "Basid Coin",
			BinanceUSDT: "BASIDUSDT",
		},
	},
	"ABYSS": {
		{
			ID:          2847,
			Slug:        "abyss",
			Name:        "Abyss",
			BinanceUSDT: "ABYSSUSDT",
		},
	},
	"FMT": {
		{
			ID:          9116,
			Slug:        "finminity",
			Name:        "Finminity",
			BinanceUSDT: "FMTUSDT",
		},
	},
	"KXC": {
		{
			ID:          3198,
			Slug:        "kingxchain",
			Name:        "KingXChain",
			BinanceUSDT: "KXCUSDT",
		},
	},
	"BENZ": {
		{
			ID:          3384,
			Slug:        "benz",
			Name:        "Benz",
			BinanceUSDT: "BENZUSDT",
		},
	},
	"ID": {
		{
			ID:          8495,
			Slug:        "everest",
			Name:        "Everest",
			BinanceUSDT: "IDUSDT",
		},
	},
	"AVC": {
		{
			ID:          6376,
			Slug:        "avccoin",
			Name:        "AVCCOIN",
			BinanceUSDT: "AVCUSDT",
		},
	},
	"EM": {
		{
			ID:   4215,
			Slug: "eminer",
			Name: "Eminer",
		},
		{
			ID:   6436,
			Slug: "empow",
			Name: "Empow",
		},
	},
	"ZLW": {
		{
			ID:          5614,
			Slug:        "zelwin",
			Name:        "Zelwin",
			BinanceUSDT: "ZLWUSDT",
		},
	},
	"XTRM": {
		{
			ID:          5599,
			Slug:        "xtrm-coin",
			Name:        "XTRM COIN",
			BinanceUSDT: "XTRMUSDT",
		},
	},
	"AUDT": {
		{
			ID:          8123,
			Slug:        "australian-dollar-token",
			Name:        "Australian Dollar Token",
			BinanceUSDT: "AUDTUSDT",
		},
	},
	"vBUSD": {
		{
			ID:          7959,
			Slug:        "venus-busd",
			Name:        "Venus BUSD",
			BinanceUSDT: "vBUSDUSDT",
		},
	},
	"NBR": {
		{
			ID:          3006,
			Slug:        "niobio",
			Name:        "Niobio",
			BinanceUSDT: "NBRUSDT",
		},
	},
	"MOC": {
		{
			ID:          2915,
			Slug:        "moss-coin",
			Name:        "Moss Coin",
			BinanceUSDT: "MOCUSDT",
		},
	},
	"AGOV": {
		{
			ID:          7673,
			Slug:        "agov-answer-governance",
			Name:        "AGOV (ANSWER Governance)",
			BinanceUSDT: "AGOVUSDT",
		},
	},
	"IDT": {
		{
			ID:          2406,
			Slug:        "investdigital",
			Name:        "InvestDigital",
			BinanceUSDT: "IDTUSDT",
		},
	},
	"GALA": {
		{
			ID:          7080,
			Slug:        "gala",
			Name:        "Gala",
			BinanceUSDT: "GALAUSDT",
		},
	},
	"LUD": {
		{
			ID:          6527,
			Slug:        "ludos",
			Name:        "Ludos Protocol",
			BinanceUSDT: "LUDUSDT",
		},
	},
	"SANSHU": {
		{
			ID:          9711,
			Slug:        "sanshu-inu",
			Name:        "Sanshu Inu",
			BinanceUSDT: "SANSHUUSDT",
		},
	},
	"VEE": {
		{
			ID:          2223,
			Slug:        "blockv",
			Name:        "BLOCKv",
			BinanceUSDT: "VEEUSDT",
		},
	},
	"DIESEL": {
		{
			ID:          8792,
			Slug:        "diesel",
			Name:        "DIESEL",
			BinanceUSDT: "DIESELUSDT",
		},
	},
	"HAKKA": {
		{
			ID:          6622,
			Slug:        "hakka-finance",
			Name:        "Hakka.Finance",
			BinanceUSDT: "HAKKAUSDT",
		},
	},
	"PAYB": {
		{
			ID:          8942,
			Slug:        "paybswap",
			Name:        "Paybswap",
			BinanceUSDT: "PAYBUSDT",
		},
	},
	"CREA": {
		{
			ID:          1676,
			Slug:        "crea",
			Name:        "CREA",
			BinanceUSDT: "CREAUSDT",
		},
	},
	"TCP": {
		{
			ID:   9416,
			Slug: "the-crypto-prophecies",
			Name: "The Crypto Prophecies",
		},
		{
			ID:   6931,
			Slug: "token-cashpay",
			Name: "Token CashPay",
		},
	},
	"MCDC": {
		{
			ID:          8562,
			Slug:        "mcdonalds-coin",
			Name:        "McDonalds Coin",
			BinanceUSDT: "MCDCUSDT",
		},
	},
	"DUSD": {
		{
			ID:          6881,
			Slug:        "defidollar",
			Name:        "DefiDollar",
			BinanceUSDT: "DUSDUSDT",
		},
	},
	"BFLY": {
		{
			ID:          8405,
			Slug:        "butterfly-protocol-2",
			Name:        "Butterfly Protocol",
			BinanceUSDT: "BFLYUSDT",
		},
	},
	"BUNI": {
		{
			ID:          9906,
			Slug:        "bunicorn",
			Name:        "Bunicorn",
			BinanceUSDT: "BUNIUSDT",
		},
	},
	"DTX": {
		{
			ID:          2913,
			Slug:        "databroker",
			Name:        "Databroker",
			BinanceUSDT: "DTXUSDT",
		},
	},
	"ROM": {
		{
			ID:          3670,
			Slug:        "romtoken",
			Name:        "ROMToken",
			BinanceUSDT: "ROMUSDT",
		},
	},
	"SNT": {
		{
			ID:          1759,
			Slug:        "status",
			Name:        "Status",
			BinanceUSDT: "SNTUSDT",
		},
	},
	"CORD": {
		{
			ID:          8535,
			Slug:        "cord-finance",
			Name:        "CORD.Finance",
			BinanceUSDT: "CORDUSDT",
		},
	},
	"RIFT": {
		{
			ID:          5110,
			Slug:        "rift-token",
			Name:        "RIFT Token",
			BinanceUSDT: "RIFTUSDT",
		},
	},
	"ARCT": {
		{
			ID:          2415,
			Slug:        "arbitragect",
			Name:        "ArbitrageCT",
			BinanceUSDT: "ARCTUSDT",
		},
	},
	"FTV": {
		{
			ID:          10650,
			Slug:        "futurov-governance-token",
			Name:        "Futurov Governance Token",
			BinanceUSDT: "FTVUSDT",
		},
	},
	"RVN": {
		{
			ID:          2577,
			Slug:        "ravencoin",
			Name:        "Ravencoin",
			BinanceUSDT: "RVNUSDT",
		},
	},
	"WILD": {
		{
			ID:   9674,
			Slug: "wilder-world",
			Name: "Wilder World",
		},
		{
			ID:   10998,
			Slug: "wild-credit",
			Name: "Wild Credit",
		},
	},
	"DPI": {
		{
			ID:          7055,
			Slug:        "defi-pulse-index",
			Name:        "DeFi Pulse Index",
			BinanceUSDT: "DPIUSDT",
		},
	},
	"CRD": {
		{
			ID:   3367,
			Slug: "crdnetwork",
			Name: "CRD Network",
		},
		{
			ID:   10993,
			Slug: "cryptodogs",
			Name: "CryptoDogs",
		},
	},
	"SENSI": {
		{
			ID:          10250,
			Slug:        "sensible-finance",
			Name:        "Sensible.Finance",
			BinanceUSDT: "SENSIUSDT",
		},
	},
	"BWX": {
		{
			ID:          2953,
			Slug:        "blue-whale-exchange",
			Name:        "Blue Whale EXchange",
			BinanceUSDT: "BWXUSDT",
		},
	},
	"PYR": {
		{
			ID:          9308,
			Slug:        "vulcan-forged-pyr",
			Name:        "Vulcan Forged PYR",
			BinanceUSDT: "PYRUSDT",
		},
	},
	"HAKA": {
		{
			ID:          10526,
			Slug:        "tribe-one",
			Name:        "TribeOne",
			BinanceUSDT: "HAKAUSDT",
		},
	},
	"eRSDL": {
		{
			ID:          7553,
			Slug:        "unfederalreserve",
			Name:        "unFederalReserve",
			BinanceUSDT: "eRSDLUSDT",
		},
	},
	"DANK": {
		{
			ID:          9041,
			Slug:        "mu-dank",
			Name:        "MU DANK",
			BinanceUSDT: "DANKUSDT",
		},
	},
	"VESPASHIBA": {
		{
			ID:          10795,
			Slug:        "vespa-shiba-coin",
			Name:        "VESPA SHIBA COIN",
			BinanceUSDT: "VESPASHIBAUSDT",
		},
	},
	"RBY": {
		{
			ID:          215,
			Slug:        "rubycoin",
			Name:        "Rubycoin",
			BinanceUSDT: "RBYUSDT",
		},
	},
	"TBCC": {
		{
			ID:          8487,
			Slug:        "tbcc-labs",
			Name:        "TBCC Labs",
			BinanceUSDT: "TBCCUSDT",
		},
	},
	"DOGEBACK": {
		{
			ID:          10792,
			Slug:        "doge-back",
			Name:        "Doge Back",
			BinanceUSDT: "DOGEBACKUSDT",
		},
	},
	"REVV": {
		{
			ID:          6993,
			Slug:        "revv",
			Name:        "REVV",
			BinanceUSDT: "REVVUSDT",
		},
	},
	"PCI": {
		{
			ID:          5275,
			Slug:        "payprotocol",
			Name:        "PayProtocol",
			BinanceUSDT: "PCIUSDT",
		},
	},
	"WOW": {
		{
			ID:   4978,
			Slug: "wownero",
			Name: "Wownero",
		},
		{
			ID:   8605,
			Slug: "wowswap",
			Name: "WOWswap",
		},
		{
			ID:   10064,
			Slug: "world-of-waves",
			Name: "World of Waves",
		},
	},
	"EPK": {
		{
			ID:          9537,
			Slug:        "epik-protocol",
			Name:        "EpiK Protocol",
			BinanceUSDT: "EPKUSDT",
		},
	},
	"WQT": {
		{
			ID:          9115,
			Slug:        "workquest",
			Name:        "WorkQuest",
			BinanceUSDT: "WQTUSDT",
		},
	},
	"AOS": {
		{
			ID:          6017,
			Slug:        "aos",
			Name:        "AOS",
			BinanceUSDT: "AOSUSDT",
		},
	},
	"QRK": {
		{
			ID:          53,
			Slug:        "quark",
			Name:        "Quark",
			BinanceUSDT: "QRKUSDT",
		},
	},
	"IQQ": {
		{
			ID:          9233,
			Slug:        "iqoniq-fanecosystem",
			Name:        "IQONIQ FanEcoSystem",
			BinanceUSDT: "IQQUSDT",
		},
	},
	"MEX": {
		{
			ID:          3286,
			Slug:        "mex",
			Name:        "MEX",
			BinanceUSDT: "MEXUSDT",
		},
	},
	"ATOMBEAR": {
		{
			ID:          6096,
			Slug:        "3x-short-cosmos-token",
			Name:        "3X Short Cosmos Token",
			BinanceUSDT: "ATOMBEARUSDT",
		},
	},
	"N8V": {
		{
			ID:          584,
			Slug:        "native-coin",
			Name:        "NativeCoin",
			BinanceUSDT: "N8VUSDT",
		},
	},
	"EURX": {
		{
			ID:          5507,
			Slug:        "etoro-euro",
			Name:        "eToro Euro",
			BinanceUSDT: "EURXUSDT",
		},
	},
	"BCA": {
		{
			ID:   2387,
			Slug: "bitcoin-atom",
			Name: "Bitcoin Atom",
		},
		{
			ID:   7619,
			Slug: "bitcoiva",
			Name: "Bitcoiva",
		},
	},
	"BBKFI": {
		{
			ID:          9581,
			Slug:        "bitblocks-finance",
			Name:        "BitBlocks Finance",
			BinanceUSDT: "BBKFIUSDT",
		},
	},
	"ALTBULL": {
		{
			ID:          6077,
			Slug:        "3x-long-altcoin-index-token",
			Name:        "3X Long Altcoin Index Token",
			BinanceUSDT: "ALTBULLUSDT",
		},
	},
	"PUBE": {
		{
			ID:          9669,
			Slug:        "pube-finance",
			Name:        "Pube finance",
			BinanceUSDT: "PUBEUSDT",
		},
	},
	"SWZL": {
		{
			ID:          5424,
			Slug:        "swapzilla",
			Name:        "Swapzilla",
			BinanceUSDT: "SWZLUSDT",
		},
	},
	"HYVE": {
		{
			ID:          7552,
			Slug:        "hyve",
			Name:        "Hyve",
			BinanceUSDT: "HYVEUSDT",
		},
	},
	"SWM": {
		{
			ID:          2506,
			Slug:        "swarm-network",
			Name:        "Swarm",
			BinanceUSDT: "SWMUSDT",
		},
	},
	"BA": {
		{
			ID:          9167,
			Slug:        "batorrent",
			Name:        "BaTorrent",
			BinanceUSDT: "BAUSDT",
		},
	},
	"GARD": {
		{
			ID:          2938,
			Slug:        "hashgard",
			Name:        "Hashgard",
			BinanceUSDT: "GARDUSDT",
		},
	},
	"PRS": {
		{
			ID:          2455,
			Slug:        "pressone",
			Name:        "PressOne",
			BinanceUSDT: "PRSUSDT",
		},
	},
	"VBSWAP": {
		{
			ID:          8865,
			Slug:        "vbswap",
			Name:        "vBSWAP",
			BinanceUSDT: "VBSWAPUSDT",
		},
	},
	"ELENA": {
		{
			ID:          9612,
			Slug:        "elena-protocol",
			Name:        "Elena Protocol",
			BinanceUSDT: "ELENAUSDT",
		},
	},
	"AT": {
		{
			ID:   3944,
			Slug: "artfinity",
			Name: "Artfinity",
		},
		{
			ID:   3238,
			Slug: "abcc-token",
			Name: "ABCC Token",
		},
		{
			ID:   10344,
			Slug: "at-finance",
			Name: "AT Finance",
		},
	},
	"LAUNCH": {
		{
			ID:          9279,
			Slug:        "superlauncher",
			Name:        "SuperLauncher",
			BinanceUSDT: "LAUNCHUSDT",
		},
	},
	"GENX": {
		{
			ID:          10753,
			Slug:        "evodefi",
			Name:        "Evodefi",
			BinanceUSDT: "GENXUSDT",
		},
	},
	"BPP": {
		{
			ID:          8126,
			Slug:        "bitpower",
			Name:        "Bitpower",
			BinanceUSDT: "BPPUSDT",
		},
	},
	"FILUP": {
		{
			ID:          8050,
			Slug:        "filup",
			Name:        "FILUP",
			BinanceUSDT: "FILUPUSDT",
		},
	},
	"TMN": {
		{
			ID:          4102,
			Slug:        "translateme-network-token",
			Name:        "TranslateMe Network Token",
			BinanceUSDT: "TMNUSDT",
		},
	},
	"DAPPX": {
		{
			ID:          10376,
			Slug:        "dappstore",
			Name:        "dAppstore",
			BinanceUSDT: "DAPPXUSDT",
		},
	},
	"ANI": {
		{
			ID:          8394,
			Slug:        "anime-token",
			Name:        "Anime Token",
			BinanceUSDT: "ANIUSDT",
		},
	},
	"WAB": {
		{
			ID:          2980,
			Slug:        "wabnetwork",
			Name:        "WABnetwork",
			BinanceUSDT: "WABUSDT",
		},
	},
	"LTN": {
		{
			ID:          9534,
			Slug:        "life-token",
			Name:        "Life Token",
			BinanceUSDT: "LTNUSDT",
		},
	},
	"TFUEL": {
		{
			ID:          3822,
			Slug:        "theta-fuel",
			Name:        "Theta Fuel",
			BinanceUSDT: "TFUELUSDT",
		},
	},
	"XLD": {
		{
			ID:          10766,
			Slug:        "stellar-diamond",
			Name:        "Stellar Diamond",
			BinanceUSDT: "XLDUSDT",
		},
	},
	"EKT": {
		{
			ID:          2453,
			Slug:        "educare",
			Name:        "EDUCare",
			BinanceUSDT: "EKTUSDT",
		},
	},
	"GVY": {
		{
			ID:          7905,
			Slug:        "groovy-finance",
			Name:        "Groovy Finance",
			BinanceUSDT: "GVYUSDT",
		},
	},
	"MAI": {
		{
			ID:          10533,
			Slug:        "mindsync",
			Name:        "Mindsync",
			BinanceUSDT: "MAIUSDT",
		},
	},
	"GBTC": {
		{
			ID:          8344,
			Slug:        "grayscale-bitcoin-trust-tokenized-stock-ftx",
			Name:        "Grayscale Bitcoin Trust tokenized stock FTX",
			BinanceUSDT: "GBTCUSDT",
		},
	},
	"KEYFI": {
		{
			ID:          8561,
			Slug:        "keyfi",
			Name:        "KeyFi",
			BinanceUSDT: "KEYFIUSDT",
		},
	},
	"SMART": {
		{
			ID:          1828,
			Slug:        "smartcash",
			Name:        "SmartCash",
			BinanceUSDT: "SMARTUSDT",
		},
	},
	"UVU": {
		{
			ID:          4244,
			Slug:        "ccuniverse",
			Name:        "CCUniverse",
			BinanceUSDT: "UVUUSDT",
		},
	},
	"LBXC": {
		{
			ID:          5490,
			Slug:        "lux-bio-cell",
			Name:        "Lux Bio Cell",
			BinanceUSDT: "LBXCUSDT",
		},
	},
	"OVI": {
		{
			ID:          9888,
			Slug:        "oviex",
			Name:        "Oviex",
			BinanceUSDT: "OVIUSDT",
		},
	},
	"BNSD": {
		{
			ID:          6911,
			Slug:        "bnsd-finance",
			Name:        "BNSD Finance",
			BinanceUSDT: "BNSDUSDT",
		},
	},
	"D4RK": {
		{
			ID:          3516,
			Slug:        "dark",
			Name:        "Dark",
			BinanceUSDT: "D4RKUSDT",
		},
	},
	"WPX": {
		{
			ID:          5208,
			Slug:        "wallet-plus-x",
			Name:        "Wallet Plus X",
			BinanceUSDT: "WPXUSDT",
		},
	},
	"SMG": {
		{
			ID:          9491,
			Slug:        "smaugs-nft",
			Name:        "Smaugs NFT",
			BinanceUSDT: "SMGUSDT",
		},
	},
	"BGL": {
		{
			ID:          5667,
			Slug:        "bitgesell",
			Name:        "Bitgesell",
			BinanceUSDT: "BGLUSDT",
		},
	},
	"FLUSD": {
		{
			ID:          10390,
			Slug:        "fluity-usd",
			Name:        "Fluity USD",
			BinanceUSDT: "FLUSDUSDT",
		},
	},
	"RADDIT": {
		{
			ID:          10074,
			Slug:        "radditarium-network",
			Name:        "Radditarium Network",
			BinanceUSDT: "RADDITUSDT",
		},
	},
	"HBX": {
		{
			ID:          3769,
			Slug:        "hashsbx",
			Name:        "HashBX",
			BinanceUSDT: "HBXUSDT",
		},
	},
	"KHC": {
		{
			ID:          10246,
			Slug:        "koho-chain",
			Name:        "KoHo Chain",
			BinanceUSDT: "KHCUSDT",
		},
	},
	"mTWTR": {
		{
			ID:          8018,
			Slug:        "mirrored-twitter",
			Name:        "Mirrored Twitter",
			BinanceUSDT: "mTWTRUSDT",
		},
	},
	"BANANA": {
		{
			ID:   8186,
			Slug: "banana-finance",
			Name: "Banana.finance",
		},
		{
			ID:   8497,
			Slug: "apeswap-finance",
			Name: "ApeSwap Finance",
		},
	},
	"SEELE": {
		{
			ID:          2830,
			Slug:        "seele",
			Name:        "Seele-N",
			BinanceUSDT: "SEELEUSDT",
		},
	},
	"YFIEC": {
		{
			ID:          7275,
			Slug:        "yearn-finance-ecosystem",
			Name:        "Yearn Finance Ecosystem",
			BinanceUSDT: "YFIECUSDT",
		},
	},
	"CAPS": {
		{
			ID:          9291,
			Slug:        "ternoa",
			Name:        "Ternoa",
			BinanceUSDT: "CAPSUSDT",
		},
	},
	"BTC2": {
		{
			ID:          3974,
			Slug:        "bitcoin2",
			Name:        "Bitcoin 2",
			BinanceUSDT: "BTC2USDT",
		},
	},
	"AIRT": {
		{
			ID:          10905,
			Slug:        "airnfts",
			Name:        "AirNFTs",
			BinanceUSDT: "AIRTUSDT",
		},
	},
	"CHEE": {
		{
			ID:          10010,
			Slug:        "cheecoin",
			Name:        "Cheecoin",
			BinanceUSDT: "CHEEUSDT",
		},
	},
	"AMP": {
		{
			ID:          6945,
			Slug:        "amp",
			Name:        "Amp",
			BinanceUSDT: "AMPUSDT",
		},
	},
	"BIDR": {
		{
			ID:          6855,
			Slug:        "binance-idr",
			Name:        "BIDR",
			BinanceUSDT: "BIDRUSDT",
		},
	},
	"DEGEN": {
		{
			ID:          8699,
			Slug:        "degen-index",
			Name:        "DEGEN Index",
			BinanceUSDT: "DEGENUSDT",
		},
	},
	"WAD": {
		{
			ID:          8981,
			Slug:        "wardenswap",
			Name:        "WardenSwap",
			BinanceUSDT: "WADUSDT",
		},
	},
	"BONFIRE": {
		{
			ID:          9522,
			Slug:        "bonfire",
			Name:        "Bonfire",
			BinanceUSDT: "BONFIREUSDT",
		},
	},
	"BLURT": {
		{
			ID:          7570,
			Slug:        "blurt",
			Name:        "Blurt",
			BinanceUSDT: "BLURTUSDT",
		},
	},
	"MM": {
		{
			ID:   10866,
			Slug: "million",
			Name: "Million",
		},
		{
			ID:   6944,
			Slug: "millimeter",
			Name: "MilliMeter",
		},
		{
			ID:   7875,
			Slug: "mm-token",
			Name: "MM Token",
		},
	},
	"HIP": {
		{
			ID:          10297,
			Slug:        "hippo-token",
			Name:        "HIPPO TOKEN",
			BinanceUSDT: "HIPUSDT",
		},
	},
	"ARCC": {
		{
			ID:          7843,
			Slug:        "asia-reserve-currency-coin",
			Name:        "Asia Reserve Currency Coin",
			BinanceUSDT: "ARCCUSDT",
		},
	},
	"EXRN": {
		{
			ID:          2088,
			Slug:        "exrnchain",
			Name:        "EXRNchain",
			BinanceUSDT: "EXRNUSDT",
		},
	},
	"MNST": {
		{
			ID:          9679,
			Slug:        "moonstarter",
			Name:        "MoonStarter",
			BinanceUSDT: "MNSTUSDT",
		},
	},
	"TEN": {
		{
			ID:   2576,
			Slug: "tokenomy",
			Name: "Tokenomy",
		},
		{
			ID:   8193,
			Slug: "tenet",
			Name: "Tenet",
		},
	},
	"MIDAS": {
		{
			ID:          4505,
			Slug:        "midas",
			Name:        "Midas",
			BinanceUSDT: "MIDASUSDT",
		},
	},
	"KDA": {
		{
			ID:          5647,
			Slug:        "kadena",
			Name:        "Kadena",
			BinanceUSDT: "KDAUSDT",
		},
	},
	"HALV": {
		{
			ID:          6506,
			Slug:        "halving-coin",
			Name:        "Halving Token",
			BinanceUSDT: "HALVUSDT",
		},
	},
	"EIFI": {
		{
			ID:          10648,
			Slug:        "eifi-finance",
			Name:        "EIFI FINANCE",
			BinanceUSDT: "EIFIUSDT",
		},
	},
	"MATPAD": {
		{
			ID:          10696,
			Slug:        "maticpad",
			Name:        "MaticPad",
			BinanceUSDT: "MATPADUSDT",
		},
	},
	"MAY": {
		{
			ID:          1693,
			Slug:        "theresa-may-coin",
			Name:        "Theresa May Coin",
			BinanceUSDT: "MAYUSDT",
		},
	},
	"WLEO": {
		{
			ID:          7221,
			Slug:        "wrapped-leo",
			Name:        "Wrapped LEO",
			BinanceUSDT: "WLEOUSDT",
		},
	},
	"SKB": {
		{
			ID:          2742,
			Slug:        "sakura-bloom",
			Name:        "Sakura Bloom",
			BinanceUSDT: "SKBUSDT",
		},
	},
	"CLAIM": {
		{
			ID:          9349,
			Slug:        "claim",
			Name:        "CLAIM",
			BinanceUSDT: "CLAIMUSDT",
		},
	},
	"GXC": {
		{
			ID:          1750,
			Slug:        "gxchain",
			Name:        "GXChain",
			BinanceUSDT: "GXCUSDT",
		},
	},
	"SWG": {
		{
			ID:          7467,
			Slug:        "swirge",
			Name:        "Swirge",
			BinanceUSDT: "SWGUSDT",
		},
	},
	"CDT": {
		{
			ID:          1864,
			Slug:        "blox",
			Name:        "Blox",
			BinanceUSDT: "CDTUSDT",
		},
	},
	"STV": {
		{
			ID:          8883,
			Slug:        "sint-truidense-voetbalvereniging",
			Name:        "Sint-Truidense Voetbalvereniging Fan Token",
			BinanceUSDT: "STVUSDT",
		},
	},
	"BTCX": {
		{
			ID:   10476,
			Slug: "bitcoinxgen",
			Name: "BitcoinX",
		},
		{
			ID:   4325,
			Slug: "bitscoin",
			Name: "Bitscoin",
		},
	},
	"PARA": {
		{
			ID:          8298,
			Slug:        "paralink-network",
			Name:        "Paralink Network",
			BinanceUSDT: "PARAUSDT",
		},
	},
	"PRIX": {
		{
			ID:          2184,
			Slug:        "privatix",
			Name:        "Privatix",
			BinanceUSDT: "PRIXUSDT",
		},
	},
	"KODURO": {
		{
			ID:          9835,
			Slug:        "koduro",
			Name:        "Koduro",
			BinanceUSDT: "KODUROUSDT",
		},
	},
	"CSPR": {
		{
			ID:          5899,
			Slug:        "casper",
			Name:        "Casper",
			BinanceUSDT: "CSPRUSDT",
		},
	},
	"GOLDUCK": {
		{
			ID:          9394,
			Slug:        "golden-duck",
			Name:        "Golden Duck",
			BinanceUSDT: "GOLDUCKUSDT",
		},
	},
	"CBIX": {
		{
			ID:          4142,
			Slug:        "cubiex",
			Name:        "Cubiex",
			BinanceUSDT: "CBIXUSDT",
		},
	},
	"WBB": {
		{
			ID:          831,
			Slug:        "wild-beast-block",
			Name:        "Wild Beast Block",
			BinanceUSDT: "WBBUSDT",
		},
	},
	"WINGS": {
		{
			ID:   1500,
			Slug: "wings",
			Name: "Wings",
		},
		{
			ID:   10810,
			Slug: "jetswap-finance",
			Name: "Jetswap.finance",
		},
	},
	"STACY": {
		{
			ID:          7574,
			Slug:        "stacy",
			Name:        "Stacy",
			BinanceUSDT: "STACYUSDT",
		},
	},
	"MNDAO": {
		{
			ID:          9006,
			Slug:        "moondao",
			Name:        "MoonDAO",
			BinanceUSDT: "MNDAOUSDT",
		},
	},
	"ORA": {
		{
			ID:          10656,
			Slug:        "coin-oracle",
			Name:        "COIN ORACLE",
			BinanceUSDT: "ORAUSDT",
		},
	},
	"OK": {
		{
			ID:          760,
			Slug:        "okcash",
			Name:        "OKCash",
			BinanceUSDT: "OKUSDT",
		},
	},
	"DEK": {
		{
			ID:          10172,
			Slug:        "dekbox",
			Name:        "DekBox",
			BinanceUSDT: "DEKUSDT",
		},
	},
	"DEXA": {
		{
			ID:          4917,
			Slug:        "dexa-coin",
			Name:        "DEXA COIN",
			BinanceUSDT: "DEXAUSDT",
		},
	},
	"DENA": {
		{
			ID:          8761,
			Slug:        "decentralized-nations",
			Name:        "Decentralized Nations",
			BinanceUSDT: "DENAUSDT",
		},
	},
	"BEC": {
		{
			ID:          7030,
			Slug:        "betherchip",
			Name:        "Betherchip",
			BinanceUSDT: "BECUSDT",
		},
	},
	"CZRX": {
		{
			ID:          5743,
			Slug:        "compound-0x",
			Name:        "Compound 0x",
			BinanceUSDT: "CZRXUSDT",
		},
	},
	"GLM": {
		{
			ID:          1455,
			Slug:        "golem-network-tokens",
			Name:        "Golem",
			BinanceUSDT: "GLMUSDT",
		},
	},
	"STAKE": {
		{
			ID:          5601,
			Slug:        "xdai",
			Name:        "xDai",
			BinanceUSDT: "STAKEUSDT",
		},
	},
	"SORA": {
		{
			ID:          5721,
			Slug:        "sorachancoin",
			Name:        "SorachanCoin",
			BinanceUSDT: "SORAUSDT",
		},
	},
	"TINV": {
		{
			ID:          10611,
			Slug:        "tinville",
			Name:        "Tinville",
			BinanceUSDT: "TINVUSDT",
		},
	},
	"EGCC": {
		{
			ID:          2852,
			Slug:        "engine",
			Name:        "Engine",
			BinanceUSDT: "EGCCUSDT",
		},
	},
	"SGC": {
		{
			ID:          5489,
			Slug:        "sudan-gold-coin",
			Name:        "Sudan Gold Coin",
			BinanceUSDT: "SGCUSDT",
		},
	},
	"FAST": {
		{
			ID:          8087,
			Slug:        "fastswap",
			Name:        "FastSwap",
			BinanceUSDT: "FASTUSDT",
		},
	},
	"CVNT": {
		{
			ID:          3686,
			Slug:        "content-value-network",
			Name:        "Content Value Network",
			BinanceUSDT: "CVNTUSDT",
		},
	},
	"CRUSADER": {
		{
			ID:          10976,
			Slug:        "crusaders-of-crypto",
			Name:        "Crusaders of Crypto",
			BinanceUSDT: "CRUSADERUSDT",
		},
	},
	"YOP": {
		{
			ID:          8187,
			Slug:        "yop",
			Name:        "Yield Optimization Platform & Protocol",
			BinanceUSDT: "YOPUSDT",
		},
	},
	"MONSTA": {
		{
			ID:          10366,
			Slug:        "cake-monster",
			Name:        "Cake Monster",
			BinanceUSDT: "MONSTAUSDT",
		},
	},
	"FL": {
		{
			ID:          8436,
			Slug:        "freeliquid",
			Name:        "Freeliquid",
			BinanceUSDT: "FLUSDT",
		},
	},
	"WAF": {
		{
			ID:          9283,
			Slug:        "waffle",
			Name:        "Waffle",
			BinanceUSDT: "WAFUSDT",
		},
	},
	"HVCO": {
		{
			ID:          1282,
			Slug:        "high-voltage",
			Name:        "High Voltage",
			BinanceUSDT: "HVCOUSDT",
		},
	},
	"XFII": {
		{
			ID:          7648,
			Slug:        "xfii",
			Name:        "XFII",
			BinanceUSDT: "XFIIUSDT",
		},
	},
	"WBTC": {
		{
			ID:          3717,
			Slug:        "wrapped-bitcoin",
			Name:        "Wrapped Bitcoin",
			BinanceUSDT: "WBTCUSDT",
		},
	},
	"LSV": {
		{
			ID:          5577,
			Slug:        "litecoin-sv",
			Name:        "Litecoin SV",
			BinanceUSDT: "LSVUSDT",
		},
	},
	"BOND": {
		{
			ID:   7440,
			Slug: "barnbridge",
			Name: "BarnBridge",
		},
		{
			ID:   7500,
			Slug: "bonded-finance",
			Name: "Bonded Finance",
		},
	},
	"IOEX": {
		{
			ID:          4289,
			Slug:        "ioex",
			Name:        "IOEX",
			BinanceUSDT: "IOEXUSDT",
		},
	},
	"YPIE": {
		{
			ID:          8324,
			Slug:        "piedao-yearn-ecosystem-pie",
			Name:        "PieDAO Yearn Ecosystem Pie",
			BinanceUSDT: "YPIEUSDT",
		},
	},
	"STARSHIP": {
		{
			ID:          9962,
			Slug:        "starship",
			Name:        "STARSHIP",
			BinanceUSDT: "STARSHIPUSDT",
		},
	},
	"CHINU": {
		{
			ID:          10056,
			Slug:        "chubby-inu",
			Name:        "Chubby Inu",
			BinanceUSDT: "CHINUUSDT",
		},
	},
	"VRX": {
		{
			ID:          8278,
			Slug:        "verox",
			Name:        "VEROX",
			BinanceUSDT: "VRXUSDT",
		},
	},
	"BTM": {
		{
			ID:          1866,
			Slug:        "bytom",
			Name:        "Bytom",
			BinanceUSDT: "BTMUSDT",
		},
	},
	"GLC": {
		{
			ID:          25,
			Slug:        "goldcoin",
			Name:        "Goldcoin",
			BinanceUSDT: "GLCUSDT",
		},
	},
	"BTCB": {
		{
			ID:          4023,
			Slug:        "bitcoin-bep2",
			Name:        "Bitcoin BEP2",
			BinanceUSDT: "BTCBUSDT",
		},
	},
	"MUSH": {
		{
			ID:          8502,
			Slug:        "mushroom",
			Name:        "Mushroom",
			BinanceUSDT: "MUSHUSDT",
		},
	},
	"ROPE": {
		{
			ID:          9326,
			Slug:        "rope-token",
			Name:        "ROPE Token",
			BinanceUSDT: "ROPEUSDT",
		},
	},
	"HDAO": {
		{
			ID:          5299,
			Slug:        "hyperdao",
			Name:        "HyperDAO",
			BinanceUSDT: "HDAOUSDT",
		},
	},
	"AVS": {
		{
			ID:          8832,
			Slug:        "algovest",
			Name:        "AlgoVest",
			BinanceUSDT: "AVSUSDT",
		},
	},
	"UNCL": {
		{
			ID:          7669,
			Slug:        "uncl",
			Name:        "UNCL",
			BinanceUSDT: "UNCLUSDT",
		},
	},
	"VIPS": {
		{
			ID:          2688,
			Slug:        "vipstar-coin",
			Name:        "Vipstar Coin",
			BinanceUSDT: "VIPSUSDT",
		},
	},
	"BONE": {
		{
			ID:          5915,
			Slug:        "bone",
			Name:        "Bone",
			BinanceUSDT: "BONEUSDT",
		},
	},
	"DEFLCT": {
		{
			ID:          7973,
			Slug:        "deflect",
			Name:        "Deflect",
			BinanceUSDT: "DEFLCTUSDT",
		},
	},
	"MAN": {
		{
			ID:          2474,
			Slug:        "matrix-ai-network",
			Name:        "Matrix AI Network",
			BinanceUSDT: "MANUSDT",
		},
	},
	"AQUA": {
		{
			ID:   10023,
			Slug: "planet-finance",
			Name: "Planet Finance",
		},
		{
			ID:   10254,
			Slug: "bela-aqua",
			Name: "Bela Aqua",
		},
	},
	"LEVELG": {
		{
			ID:          5578,
			Slug:        "levelg",
			Name:        "LEVELG",
			BinanceUSDT: "LEVELGUSDT",
		},
	},
	"$PEEPO": {
		{
			ID:          10146,
			Slug:        "peepocoin",
			Name:        "PeepoCoin",
			BinanceUSDT: "$PEEPOUSDT",
		},
	},
	"FTN": {
		{
			ID:          3658,
			Slug:        "fountain",
			Name:        "Fountain",
			BinanceUSDT: "FTNUSDT",
		},
	},
	"NMR": {
		{
			ID:          1732,
			Slug:        "numeraire",
			Name:        "Numeraire",
			BinanceUSDT: "NMRUSDT",
		},
	},
	"ZYD": {
		{
			ID:          1389,
			Slug:        "zayedcoin",
			Name:        "Zayedcoin",
			BinanceUSDT: "ZYDUSDT",
		},
	},
	"HAPI": {
		{
			ID:          8567,
			Slug:        "hapi-one",
			Name:        "HAPI",
			BinanceUSDT: "HAPIUSDT",
		},
	},
	"ROYA": {
		{
			ID:          7821,
			Slug:        "royale-finance",
			Name:        "Royale Finance",
			BinanceUSDT: "ROYAUSDT",
		},
	},
	"BLOC": {
		{
			ID:   3451,
			Slug: "bloc-money",
			Name: "BLOC.MONEY",
		},
		{
			ID:   3860,
			Slug: "blockcloud",
			Name: "Blockcloud",
		},
	},
	"PEFI": {
		{
			ID:          10818,
			Slug:        "penguin-finance",
			Name:        "Penguin Finance",
			BinanceUSDT: "PEFIUSDT",
		},
	},
	"ETHPY": {
		{
			ID:          7049,
			Slug:        "etherpay",
			Name:        "Etherpay",
			BinanceUSDT: "ETHPYUSDT",
		},
	},
	"MYFI": {
		{
			ID:          8953,
			Slug:        "myfinance",
			Name:        "MYFinance",
			BinanceUSDT: "MYFIUSDT",
		},
	},
	"GHST": {
		{
			ID:          7046,
			Slug:        "aavegotchi",
			Name:        "Aavegotchi",
			BinanceUSDT: "GHSTUSDT",
		},
	},
	"BUIDL": {
		{
			ID:          5704,
			Slug:        "dfohub",
			Name:        "DFOhub",
			BinanceUSDT: "BUIDLUSDT",
		},
	},
	"COTI": {
		{
			ID:          3992,
			Slug:        "coti",
			Name:        "COTI",
			BinanceUSDT: "COTIUSDT",
		},
	},
	"F1C": {
		{
			ID:          3417,
			Slug:        "future1coin",
			Name:        "Future1coin",
			BinanceUSDT: "F1CUSDT",
		},
	},
	"GRLC": {
		{
			ID:          2475,
			Slug:        "garlicoin",
			Name:        "Garlicoin",
			BinanceUSDT: "GRLCUSDT",
		},
	},
	"BET": {
		{
			ID:   1771,
			Slug: "daobet",
			Name: "DAOBet",
		},
		{
			ID:   5487,
			Slug: "earnbet",
			Name: "EarnBet",
		},
	},
	"KLV": {
		{
			ID:          6724,
			Slug:        "klever",
			Name:        "Klever",
			BinanceUSDT: "KLVUSDT",
		},
	},
	"AST": {
		{
			ID:   2058,
			Slug: "airswap",
			Name: "AirSwap",
		},
		{
			ID:   5827,
			Slug: "antiscamtoken",
			Name: "AntiscamToken",
		},
	},
	"BTSE": {
		{
			ID:          5305,
			Slug:        "btse",
			Name:        "BTSE",
			BinanceUSDT: "BTSEUSDT",
		},
	},
	"CWBTC": {
		{
			ID:          5744,
			Slug:        "compound-wrapped-btc",
			Name:        "Compound Wrapped BTC",
			BinanceUSDT: "CWBTCUSDT",
		},
	},
	"SHIBERUS": {
		{
			ID:          10874,
			Slug:        "shiberus-inu",
			Name:        "Shiberus Inu",
			BinanceUSDT: "SHIBERUSUSDT",
		},
	},
	"MWAR": {
		{
			ID:          10840,
			Slug:        "mooniwar",
			Name:        "MooniWar",
			BinanceUSDT: "MWARUSDT",
		},
	},
	"LMY": {
		{
			ID:          4577,
			Slug:        "lunchmoney",
			Name:        "LunchMoney",
			BinanceUSDT: "LMYUSDT",
		},
	},
	"DRGNBULL": {
		{
			ID:          6082,
			Slug:        "3x-long-dragon-index-token",
			Name:        "3X Long Dragon Index Token",
			BinanceUSDT: "DRGNBULLUSDT",
		},
	},
	"PLUM": {
		{
			ID:          9773,
			Slug:        "plumcake-finance",
			Name:        "PlumCake Finance",
			BinanceUSDT: "PLUMUSDT",
		},
	},
	"WLITI": {
		{
			ID:          10740,
			Slug:        "liti-capital",
			Name:        "Liti Capital",
			BinanceUSDT: "WLITIUSDT",
		},
	},
	"CEUR": {
		{
			ID:          9467,
			Slug:        "celo-euro",
			Name:        "Celo Euro",
			BinanceUSDT: "CEURUSDT",
		},
	},
	"LIMIT": {
		{
			ID:          7379,
			Slug:        "limitswap",
			Name:        "LimitSwap",
			BinanceUSDT: "LIMITUSDT",
		},
	},
	"LIMEX": {
		{
			ID:          5985,
			Slug:        "limestone-network",
			Name:        "Limestone Network",
			BinanceUSDT: "LIMEXUSDT",
		},
	},
	"WDR": {
		{
			ID:          9636,
			Slug:        "widercoin",
			Name:        "Widercoin",
			BinanceUSDT: "WDRUSDT",
		},
	},
	"FTC": {
		{
			ID:          8,
			Slug:        "feathercoin",
			Name:        "Feathercoin",
			BinanceUSDT: "FTCUSDT",
		},
	},
	"ZENI": {
		{
			ID:          1629,
			Slug:        "zennies",
			Name:        "Zennies",
			BinanceUSDT: "ZENIUSDT",
		},
	},
	"POKE": {
		{
			ID:          8303,
			Slug:        "pokeball",
			Name:        "Pokeball",
			BinanceUSDT: "POKEUSDT",
		},
	},
	"MMAON": {
		{
			ID:          9169,
			Slug:        "mmaon",
			Name:        "MMAON",
			BinanceUSDT: "MMAONUSDT",
		},
	},
	"RAISE": {
		{
			ID:          2755,
			Slug:        "raise",
			Name:        "Raise",
			BinanceUSDT: "RAISEUSDT",
		},
	},
	"LEDU": {
		{
			ID:          2562,
			Slug:        "education-ecosystem",
			Name:        "Education Ecosystem",
			BinanceUSDT: "LEDUUSDT",
		},
	},
	"EPRO": {
		{
			ID:          10607,
			Slug:        "ethereum-pro",
			Name:        "Ethereum Pro",
			BinanceUSDT: "EPROUSDT",
		},
	},
	"HINT": {
		{
			ID:          4254,
			Slug:        "hintchain",
			Name:        "Hintchain",
			BinanceUSDT: "HINTUSDT",
		},
	},
	"NOV": {
		{
			ID:          8885,
			Slug:        "novara-calcio-fan-token",
			Name:        "Novara Calcio Fan Token",
			BinanceUSDT: "NOVUSDT",
		},
	},
	"KIM": {
		{
			ID:          5286,
			Slug:        "kingmoney",
			Name:        "KingMoney",
			BinanceUSDT: "KIMUSDT",
		},
	},
	"SENSO": {
		{
			ID:          5522,
			Slug:        "senso",
			Name:        "SENSO",
			BinanceUSDT: "SENSOUSDT",
		},
	},
	"LINKETHPA": {
		{
			ID:          6140,
			Slug:        "eth-link-price-action-candlestick-set",
			Name:        "ETH/LINK Price Action Candlestick Set",
			BinanceUSDT: "LINKETHPAUSDT",
		},
	},
	"PUSH": {
		{
			ID:          9111,
			Slug:        "epns",
			Name:        "Ethereum Push Notification Service",
			BinanceUSDT: "PUSHUSDT",
		},
	},
	"BLAST": {
		{
			ID:   3387,
			Slug: "blast",
			Name: "BLAST",
		},
		{
			ID:   9967,
			Slug: "safeblast",
			Name: "SafeBlast",
		},
	},
	"TKP": {
		{
			ID:          4116,
			Slug:        "tokpie",
			Name:        "TOKPIE",
			BinanceUSDT: "TKPUSDT",
		},
	},
	"TCUB": {
		{
			ID:          10862,
			Slug:        "tiger-cub",
			Name:        "Tiger Cub",
			BinanceUSDT: "TCUBUSDT",
		},
	},
	"TOM": {
		{
			ID:          7734,
			Slug:        "tom-finance",
			Name:        "TOM Finance",
			BinanceUSDT: "TOMUSDT",
		},
	},
	"SAFEMARS": {
		{
			ID:          8966,
			Slug:        "safemars",
			Name:        "Safemars",
			BinanceUSDT: "SAFEMARSUSDT",
		},
	},
	"RANK": {
		{
			ID:          7947,
			Slug:        "rank-token",
			Name:        "Rank Token",
			BinanceUSDT: "RANKUSDT",
		},
	},
	"NFT": {
		{
			ID:   9816,
			Slug: "apenft",
			Name: "APENFT",
		},
		{
			ID:   6650,
			Slug: "nft",
			Name: "NFT",
		},
		{
			ID:   10829,
			Slug: "neftipedia",
			Name: "NEFTiPEDiA",
		},
	},
	"VLT": {
		{
			ID:   1368,
			Slug: "veltor",
			Name: "Veltor",
		},
		{
			ID:   5856,
			Slug: "bankroll-vault",
			Name: "Bankroll Vault",
		},
	},
	"FREN": {
		{
			ID:          9767,
			Slug:        "frenchie-network",
			Name:        "Frenchie Network",
			BinanceUSDT: "FRENUSDT",
		},
	},
	"WITCH": {
		{
			ID:          10984,
			Slug:        "witch-token",
			Name:        "Witch Token",
			BinanceUSDT: "WITCHUSDT",
		},
	},
	"ORION": {
		{
			ID:          10395,
			Slug:        "orion",
			Name:        "Orion",
			BinanceUSDT: "ORIONUSDT",
		},
	},
	"ESCE": {
		{
			ID:          3489,
			Slug:        "escroco-emerald",
			Name:        "Escroco Emerald",
			BinanceUSDT: "ESCEUSDT",
		},
	},
	"VIA": {
		{
			ID:          470,
			Slug:        "viacoin",
			Name:        "Viacoin",
			BinanceUSDT: "VIAUSDT",
		},
	},
	"NPXS": {
		{
			ID:          2603,
			Slug:        "pundi-x",
			Name:        "Pundi X[old]",
			BinanceUSDT: "NPXSUSDT",
		},
	},
	"CSPC": {
		{
			ID:          6412,
			Slug:        "cspc",
			Name:        "CSPC",
			BinanceUSDT: "CSPCUSDT",
		},
	},
	"AXS": {
		{
			ID:          6783,
			Slug:        "axie-infinity",
			Name:        "Axie Infinity",
			BinanceUSDT: "AXSUSDT",
		},
	},
	"BERN": {
		{
			ID:   1223,
			Slug: "berncash",
			Name: "BERNcash",
		},
		{
			ID:   10737,
			Slug: "bernard",
			Name: "BERNARD",
		},
	},
	"STM": {
		{
			ID:          5046,
			Slug:        "streamity",
			Name:        "Streamity",
			BinanceUSDT: "STMUSDT",
		},
	},
	"LCT": {
		{
			ID:          7733,
			Slug:        "light-coin-exchange-token",
			Name:        "Light Coin Exchange Token",
			BinanceUSDT: "LCTUSDT",
		},
	},
	"CRAD": {
		{
			ID:          4255,
			Slug:        "cryptoads-marketplace",
			Name:        "CryptoAds Marketplace",
			BinanceUSDT: "CRADUSDT",
		},
	},
	"ITT": {
		{
			ID:          2103,
			Slug:        "intelligent-trading-foundation",
			Name:        "Intelligent Trading Foundation",
			BinanceUSDT: "ITTUSDT",
		},
	},
	"CRYPT": {
		{
			ID:          10135,
			Slug:        "cryptonaught",
			Name:        "Cryptonaught",
			BinanceUSDT: "CRYPTUSDT",
		},
	},
	"KLP": {
		{
			ID:          6507,
			Slug:        "kulupu",
			Name:        "Kulupu",
			BinanceUSDT: "KLPUSDT",
		},
	},
	"TOTM": {
		{
			ID:          8673,
			Slug:        "totemfi",
			Name:        "TotemFi",
			BinanceUSDT: "TOTMUSDT",
		},
	},
	"HD": {
		{
			ID:          7541,
			Slug:        "hubdao",
			Name:        "HubDao",
			BinanceUSDT: "HDUSDT",
		},
	},
	"UTU": {
		{
			ID:          7587,
			Slug:        "utu-protocol",
			Name:        "UTU Protocol",
			BinanceUSDT: "UTUUSDT",
		},
	},
	"OGC": {
		{
			ID:          9257,
			Slug:        "one-get-coin",
			Name:        "One Get Coin",
			BinanceUSDT: "OGCUSDT",
		},
	},
	"$PAWG": {
		{
			ID:          10538,
			Slug:        "pawgcoin",
			Name:        "PAWGcoin",
			BinanceUSDT: "$PAWGUSDT",
		},
	},
	"WDC": {
		{
			ID:   16,
			Slug: "worldcoin",
			Name: "WorldCoin",
		},
		{
			ID:   5178,
			Slug: "wisdom-chain",
			Name: "Wisdom Chain",
		},
	},
	"SVN": {
		{
			ID:          7662,
			Slug:        "7finance",
			Name:        "7Finance",
			BinanceUSDT: "SVNUSDT",
		},
	},
	"UME": {
		{
			ID:          8852,
			Slug:        "ume-token",
			Name:        "UME Token",
			BinanceUSDT: "UMEUSDT",
		},
	},
	"MITX": {
		{
			ID:          2709,
			Slug:        "morpheus-labs",
			Name:        "Morpheus Labs",
			BinanceUSDT: "MITXUSDT",
		},
	},
	"PIGX": {
		{
			ID:          7981,
			Slug:        "pigx",
			Name:        "PIGX",
			BinanceUSDT: "PIGXUSDT",
		},
	},
	"FUN": {
		{
			ID:          1757,
			Slug:        "funtoken",
			Name:        "FUNToken",
			BinanceUSDT: "FUNUSDT",
		},
	},
	"ROUTE": {
		{
			ID:          8292,
			Slug:        "router-protocol",
			Name:        "Router Protocol",
			BinanceUSDT: "ROUTEUSDT",
		},
	},
	"ZINC": {
		{
			ID:          2883,
			Slug:        "zinc",
			Name:        "ZINC",
			BinanceUSDT: "ZINCUSDT",
		},
	},
	"BOG": {
		{
			ID:          8723,
			Slug:        "bogged-finance",
			Name:        "Bogged Finance",
			BinanceUSDT: "BOGUSDT",
		},
	},
	"10SET": {
		{
			ID:          9089,
			Slug:        "tenset",
			Name:        "Tenset",
			BinanceUSDT: "10SETUSDT",
		},
	},
	"FAI": {
		{
			ID:          8414,
			Slug:        "fairum",
			Name:        "Fairum",
			BinanceUSDT: "FAIUSDT",
		},
	},
	"NEC": {
		{
			ID:          2538,
			Slug:        "nectar",
			Name:        "Nectar",
			BinanceUSDT: "NECUSDT",
		},
	},
	"TPAY": {
		{
			ID:          2627,
			Slug:        "tokenpay",
			Name:        "TokenPay",
			BinanceUSDT: "TPAYUSDT",
		},
	},
	"YBEAR": {
		{
			ID:          9147,
			Slug:        "ybear-finance",
			Name:        "yBEAR.finance",
			BinanceUSDT: "YBEARUSDT",
		},
	},
	"ECA": {
		{
			ID:          1711,
			Slug:        "electra",
			Name:        "Electra",
			BinanceUSDT: "ECAUSDT",
		},
	},
	"TME": {
		{
			ID:          8455,
			Slug:        "tama-egg-niftygotchi",
			Name:        "TAMA EGG NiftyGotchi",
			BinanceUSDT: "TMEUSDT",
		},
	},
	"NAOS": {
		{
			ID:          9504,
			Slug:        "naos-finance",
			Name:        "NAOS Finance",
			BinanceUSDT: "NAOSUSDT",
		},
	},
	"SHIBAKEN": {
		{
			ID:          10257,
			Slug:        "shibaken-finance",
			Name:        "Shibaken Finance",
			BinanceUSDT: "SHIBAKENUSDT",
		},
	},
	"JACK": {
		{
			ID:          5486,
			Slug:        "jack-token",
			Name:        "Jack Token",
			BinanceUSDT: "JACKUSDT",
		},
	},
	"XQN": {
		{
			ID:          733,
			Slug:        "quotient",
			Name:        "Quotient",
			BinanceUSDT: "XQNUSDT",
		},
	},
	"SGO": {
		{
			ID:          10751,
			Slug:        "sportemon-go",
			Name:        "Sportemon-Go",
			BinanceUSDT: "SGOUSDT",
		},
	},
	"UGMC": {
		{
			ID:          10909,
			Slug:        "unicly-genesis-mooncats-collection",
			Name:        "Unicly Genesis MoonCats Collection",
			BinanceUSDT: "UGMCUSDT",
		},
	},
	"POLVEN": {
		{
			ID:          8974,
			Slug:        "polka-ventures",
			Name:        "Polka Ventures",
			BinanceUSDT: "POLVENUSDT",
		},
	},
	"TGIC": {
		{
			ID:          5551,
			Slug:        "the-global-index-chain",
			Name:        "The global index chain",
			BinanceUSDT: "TGICUSDT",
		},
	},
	"CHUM": {
		{
			ID:          10360,
			Slug:        "chumhum",
			Name:        "Chumhum",
			BinanceUSDT: "CHUMUSDT",
		},
	},
	"LOV": {
		{
			ID:          7634,
			Slug:        "the-lovechain",
			Name:        "The LoveChain",
			BinanceUSDT: "LOVUSDT",
		},
	},
	"SWAG": {
		{
			ID:          7428,
			Slug:        "swag-finance",
			Name:        "SWAG Finance",
			BinanceUSDT: "SWAGUSDT",
		},
	},
	"MATICBULL": {
		{
			ID:          6085,
			Slug:        "3x-long-matic-token",
			Name:        "3X Long Matic Token",
			BinanceUSDT: "MATICBULLUSDT",
		},
	},
	"POSW": {
		{
			ID:          1495,
			Slug:        "posw-coin",
			Name:        "PoSW Coin",
			BinanceUSDT: "POSWUSDT",
		},
	},
	"AIRX": {
		{
			ID:          4552,
			Slug:        "aircoins",
			Name:        "Aircoins",
			BinanceUSDT: "AIRXUSDT",
		},
	},
	"XLM": {
		{
			ID:          512,
			Slug:        "stellar",
			Name:        "Stellar",
			BinanceUSDT: "XLMUSDT",
		},
	},
	"MDO": {
		{
			ID:          8486,
			Slug:        "midas-dollar",
			Name:        "Midas Dollar",
			BinanceUSDT: "MDOUSDT",
		},
	},
	"BFF": {
		{
			ID:          5180,
			Slug:        "bitcoffeen",
			Name:        "Bitcoffeen",
			BinanceUSDT: "BFFUSDT",
		},
	},
	"NIU": {
		{
			ID:          8748,
			Slug:        "niubiswap",
			Name:        "Niubi Swap",
			BinanceUSDT: "NIUUSDT",
		},
	},
	"KUMA": {
		{
			ID:          10394,
			Slug:        "kuma-inu",
			Name:        "Kuma Inu",
			BinanceUSDT: "KUMAUSDT",
		},
	},
	"BABYWOLF": {
		{
			ID:          10769,
			Slug:        "baby-moon-wolf",
			Name:        "Baby Moon Wolf",
			BinanceUSDT: "BABYWOLFUSDT",
		},
	},
	"TARA": {
		{
			ID:          8715,
			Slug:        "taraxa",
			Name:        "Taraxa",
			BinanceUSDT: "TARAUSDT",
		},
	},
	"SHIBBY": {
		{
			ID:          10616,
			Slug:        "shibby",
			Name:        "Shibby",
			BinanceUSDT: "SHIBBYUSDT",
		},
	},
	"LINKPT": {
		{
			ID:          6159,
			Slug:        "link-profit-taker-set",
			Name:        "LINK Profit Taker Set",
			BinanceUSDT: "LINKPTUSDT",
		},
	},
	"TENA": {
		{
			ID:   3643,
			Slug: "tena-old",
			Name: "TENA [old]",
		},
		{
			ID:   9930,
			Slug: "tena-new",
			Name: "Tena [new]",
		},
	},
	"TONE": {
		{
			ID:   2578,
			Slug: "te-food",
			Name: "TE-FOOD",
		},
		{
			ID:   10169,
			Slug: "nfttone",
			Name: "NFTTONE",
		},
	},
	"FYP": {
		{
			ID:          2126,
			Slug:        "flypme",
			Name:        "FlypMe",
			BinanceUSDT: "FYPUSDT",
		},
	},
	"CSC": {
		{
			ID:          45,
			Slug:        "casinocoin",
			Name:        "CasinoCoin",
			BinanceUSDT: "CSCUSDT",
		},
	},
	"SPACE": {
		{
			ID:   11026,
			Slug: "spacelens",
			Name: "Spacelens",
		},
		{
			ID:   8952,
			Slug: "farm-space",
			Name: "Farm Space",
		},
	},
	"UNDG": {
		{
			ID:          8400,
			Slug:        "unidexgas",
			Name:        "UniDexGas",
			BinanceUSDT: "UNDGUSDT",
		},
	},
	"CAS": {
		{
			ID:          2529,
			Slug:        "cashaa",
			Name:        "Cashaa",
			BinanceUSDT: "CASUSDT",
		},
	},
	"EOSDAC": {
		{
			ID:          2644,
			Slug:        "eosdac",
			Name:        "eosDAC",
			BinanceUSDT: "EOSDACUSDT",
		},
	},
	"FUSII": {
		{
			ID:          8703,
			Slug:        "fusible",
			Name:        "Fusible",
			BinanceUSDT: "FUSIIUSDT",
		},
	},
	"XCF": {
		{
			ID:          8441,
			Slug:        "cenfura-token",
			Name:        "Cenfura Token",
			BinanceUSDT: "XCFUSDT",
		},
	},
	"SHDC": {
		{
			ID:          8740,
			Slug:        "shd-cash",
			Name:        "SHD CASH",
			BinanceUSDT: "SHDCUSDT",
		},
	},
	"ALY": {
		{
			ID:          5011,
			Slug:        "ally",
			Name:        "ALLY",
			BinanceUSDT: "ALYUSDT",
		},
	},
	"UNI": {
		{
			ID:   7083,
			Slug: "uniswap",
			Name: "Uniswap",
		},
		{
			ID:   4307,
			Slug: "unicorn-token",
			Name: "UNICORN Token",
		},
	},
	"EIDOS": {
		{
			ID:          6070,
			Slug:        "eidos",
			Name:        "EIDOS",
			BinanceUSDT: "EIDOSUSDT",
		},
	},
	"AFEN": {
		{
			ID:          9752,
			Slug:        "afen-blockchain",
			Name:        "AFEN Blockchain",
			BinanceUSDT: "AFENUSDT",
		},
	},
	"FX": {
		{
			ID:          3884,
			Slug:        "function-x",
			Name:        "Function X",
			BinanceUSDT: "FXUSDT",
		},
	},
	"SYS": {
		{
			ID:          541,
			Slug:        "syscoin",
			Name:        "Syscoin",
			BinanceUSDT: "SYSUSDT",
		},
	},
	"WHX": {
		{
			ID:          8943,
			Slug:        "whitex",
			Name:        "WHITEX",
			BinanceUSDT: "WHXUSDT",
		},
	},
	"EPAY": {
		{
			ID:          10374,
			Slug:        "ethereumpay",
			Name:        "EthereumPay",
			BinanceUSDT: "EPAYUSDT",
		},
	},
	"FOAM": {
		{
			ID:          3631,
			Slug:        "foam",
			Name:        "FOAM",
			BinanceUSDT: "FOAMUSDT",
		},
	},
	"YFI": {
		{
			ID:          5864,
			Slug:        "yearn-finance",
			Name:        "yearn.finance",
			BinanceUSDT: "YFIUSDT",
		},
	},
	"KOJI": {
		{
			ID:          10680,
			Slug:        "koji",
			Name:        "Koji",
			BinanceUSDT: "KOJIUSDT",
		},
	},
	"DEP": {
		{
			ID:   5429,
			Slug: "deapcoin",
			Name: "DEAPcoin",
		},
		{
			ID:   9164,
			Slug: "depth-token",
			Name: "Depth Token",
		},
	},
	"TESLASAFE": {
		{
			ID:          11019,
			Slug:        "teslasafe",
			Name:        "TeslaSafe",
			BinanceUSDT: "TESLASAFEUSDT",
		},
	},
	"REAU": {
		{
			ID:          9413,
			Slug:        "vira-lata-finance",
			Name:        "Vira-lata Finance",
			BinanceUSDT: "REAUUSDT",
		},
	},
	"TOOLS": {
		{
			ID:          8866,
			Slug:        "bsc-tools",
			Name:        "BSC TOOLS",
			BinanceUSDT: "TOOLSUSDT",
		},
	},
	"PHT": {
		{
			ID:   4991,
			Slug: "lightstreams",
			Name: "Lightstreams",
		},
		{
			ID:   5214,
			Slug: "phoneum",
			Name: "Phoneum",
		},
	},
	"HSC": {
		{
			ID:          2908,
			Slug:        "hashcoin",
			Name:        "HashCoin",
			BinanceUSDT: "HSCUSDT",
		},
	},
	"GOGO": {
		{
			ID:          8375,
			Slug:        "gogo-finance",
			Name:        "GOGO.finance",
			BinanceUSDT: "GOGOUSDT",
		},
	},
	"CARR": {
		{
			ID:          8806,
			Slug:        "carnomaly",
			Name:        "Carnomaly",
			BinanceUSDT: "CARRUSDT",
		},
	},
	"KET": {
		{
			ID:          9431,
			Slug:        "rowket-market",
			Name:        "Rowket",
			BinanceUSDT: "KETUSDT",
		},
	},
	"BELA": {
		{
			ID:          217,
			Slug:        "belacoin",
			Name:        "Bela",
			BinanceUSDT: "BELAUSDT",
		},
	},
	"SLOKI": {
		{
			ID:          10868,
			Slug:        "super-floki",
			Name:        "Super Floki",
			BinanceUSDT: "SLOKIUSDT",
		},
	},
	"MDS": {
		{
			ID:   2274,
			Slug: "medishares",
			Name: "MediShares",
		},
		{
			ID:   8485,
			Slug: "midas-dollar-share",
			Name: "Midas Dollar Share",
		},
	},
	"SWISE": {
		{
			ID:          10439,
			Slug:        "stakewise",
			Name:        "StakeWise",
			BinanceUSDT: "SWISEUSDT",
		},
	},
	"BTAD": {
		{
			ID:          3294,
			Slug:        "bitcoin-adult",
			Name:        "Bitcoin Adult",
			BinanceUSDT: "BTADUSDT",
		},
	},
	"BTCST": {
		{
			ID:          8891,
			Slug:        "btc-standard-hashrate-token",
			Name:        "Bitcoin Standard Hashrate Token",
			BinanceUSDT: "BTCSTUSDT",
		},
	},
	"UUNICLY": {
		{
			ID:          9469,
			Slug:        "unicly-genesis-collection",
			Name:        "Unicly Genesis Collection",
			BinanceUSDT: "UUNICLYUSDT",
		},
	},
	"GAINS": {
		{
			ID:          9125,
			Slug:        "gains-associates",
			Name:        "Gains Associates",
			BinanceUSDT: "GAINSUSDT",
		},
	},
	"IOST": {
		{
			ID:          2405,
			Slug:        "iostoken",
			Name:        "IOST",
			BinanceUSDT: "IOSTUSDT",
		},
	},
	"XRT": {
		{
			ID:          4757,
			Slug:        "robonomics-network",
			Name:        "Robonomics.network",
			BinanceUSDT: "XRTUSDT",
		},
	},
	"EMC2": {
		{
			ID:          201,
			Slug:        "einsteinium",
			Name:        "Einsteinium",
			BinanceUSDT: "EMC2USDT",
		},
	},
	"WGP": {
		{
			ID:          3912,
			Slug:        "w-green-pay",
			Name:        "W Green Pay",
			BinanceUSDT: "WGPUSDT",
		},
	},
	"LSK": {
		{
			ID:          1214,
			Slug:        "lisk",
			Name:        "Lisk",
			BinanceUSDT: "LSKUSDT",
		},
	},
	"SKM": {
		{
			ID:          2725,
			Slug:        "skrumble-network",
			Name:        "Skrumble Network",
			BinanceUSDT: "SKMUSDT",
		},
	},
	"PWAY": {
		{
			ID:          8838,
			Slug:        "pway",
			Name:        "PWAY",
			BinanceUSDT: "PWAYUSDT",
		},
	},
	"BUL": {
		{
			ID:          3690,
			Slug:        "bulleon",
			Name:        "Bulleon",
			BinanceUSDT: "BULUSDT",
		},
	},
	"ARCX": {
		{
			ID:          10515,
			Slug:        "arcx-token",
			Name:        "ARC Governance",
			BinanceUSDT: "ARCXUSDT",
		},
	},
	"NWC": {
		{
			ID:          4890,
			Slug:        "newscrypto",
			Name:        "Newscrypto",
			BinanceUSDT: "NWCUSDT",
		},
	},
	"EGLD": {
		{
			ID:          6892,
			Slug:        "elrond-egld",
			Name:        "Elrond",
			BinanceUSDT: "EGLDUSDT",
		},
	},
	"BALI": {
		{
			ID:          5480,
			Slug:        "bali-coin",
			Name:        "Bali Coin",
			BinanceUSDT: "BALIUSDT",
		},
	},
	"WASABI": {
		{
			ID:          9077,
			Slug:        "wasabix",
			Name:        "WasabiX",
			BinanceUSDT: "WASABIUSDT",
		},
	},
	"XKI": {
		{
			ID:          9908,
			Slug:        "ki-foundation",
			Name:        "Ki",
			BinanceUSDT: "XKIUSDT",
		},
	},
	"ROO": {
		{
			ID:          10574,
			Slug:        "roocoin",
			Name:        "RooCoin",
			BinanceUSDT: "ROOUSDT",
		},
	},
	"PAPP": {
		{
			ID:          10163,
			Slug:        "papp-mobile",
			Name:        "Papp Mobile",
			BinanceUSDT: "PAPPUSDT",
		},
	},
	"DCB": {
		{
			ID:          10563,
			Slug:        "decubate",
			Name:        "Decubate",
			BinanceUSDT: "DCBUSDT",
		},
	},
	"GTON": {
		{
			ID:          10006,
			Slug:        "graviton-one",
			Name:        "Graviton",
			BinanceUSDT: "GTONUSDT",
		},
	},
	"CRB": {
		{
			ID:          10345,
			Slug:        "cribnb-decentralized-renting-and-sharing",
			Name:        "Cribnb Decentralized Renting and Sharing",
			BinanceUSDT: "CRBUSDT",
		},
	},
	"AOE": {
		{
			ID:          10415,
			Slug:        "asset-of-empires",
			Name:        "Asset of Empires",
			BinanceUSDT: "AOEUSDT",
		},
	},
	"MAHA": {
		{
			ID:          8043,
			Slug:        "mahadao",
			Name:        "MahaDAO",
			BinanceUSDT: "MAHAUSDT",
		},
	},
	"STBU": {
		{
			ID:          7245,
			Slug:        "stobox-token",
			Name:        "Stobox Token",
			BinanceUSDT: "STBUUSDT",
		},
	},
	"MWAT": {
		{
			ID:          2533,
			Slug:        "restart-energy-mwat",
			Name:        "Restart Energy MWAT",
			BinanceUSDT: "MWATUSDT",
		},
	},
	"OPNN": {
		{
			ID:          4202,
			Slug:        "opennity",
			Name:        "Opennity",
			BinanceUSDT: "OPNNUSDT",
		},
	},
	"DPT": {
		{
			ID:          3920,
			Slug:        "diamond-platform-token",
			Name:        "Diamond Platform Token",
			BinanceUSDT: "DPTUSDT",
		},
	},
	"MYO": {
		{
			ID:          9577,
			Slug:        "mycro",
			Name:        "Mycro",
			BinanceUSDT: "MYOUSDT",
		},
	},
	"TAJ": {
		{
			ID:          1353,
			Slug:        "tajcoin",
			Name:        "TajCoin",
			BinanceUSDT: "TAJUSDT",
		},
	},
	"BDOG": {
		{
			ID:          9280,
			Slug:        "bulldog-token",
			Name:        "Bulldog Token",
			BinanceUSDT: "BDOGUSDT",
		},
	},
	"FCL": {
		{
			ID:          8544,
			Slug:        "fractal",
			Name:        "Fractal",
			BinanceUSDT: "FCLUSDT",
		},
	},
	"USNBT": {
		{
			ID:          626,
			Slug:        "nubits",
			Name:        "NuBits",
			BinanceUSDT: "USNBTUSDT",
		},
	},
	"TRAIL": {
		{
			ID:          10732,
			Slug:        "polkatrail",
			Name:        "PolkaTrail",
			BinanceUSDT: "TRAILUSDT",
		},
	},
	"STQ": {
		{
			ID:          2541,
			Slug:        "storiqa",
			Name:        "Storiqa",
			BinanceUSDT: "STQUSDT",
		},
	},
	"DDK": {
		{
			ID:          4180,
			Slug:        "ddkoin",
			Name:        "DDKoin",
			BinanceUSDT: "DDKUSDT",
		},
	},
	"MOLA": {
		{
			ID:          10185,
			Slug:        "moonlana",
			Name:        "Moonlana",
			BinanceUSDT: "MOLAUSDT",
		},
	},
	"AGIX": {
		{
			ID:          2424,
			Slug:        "singularitynet",
			Name:        "SingularityNET",
			BinanceUSDT: "AGIXUSDT",
		},
	},
	"SAFEZONE": {
		{
			ID:          9799,
			Slug:        "safezone",
			Name:        "SafeZone",
			BinanceUSDT: "SAFEZONEUSDT",
		},
	},
	"SHIBSC": {
		{
			ID:          10085,
			Slug:        "shiba-bsc",
			Name:        "Shiba BSC",
			BinanceUSDT: "SHIBSCUSDT",
		},
	},
	"PHOON": {
		{
			ID:          8306,
			Slug:        "typhoon-cash",
			Name:        "Typhoon Cash",
			BinanceUSDT: "PHOONUSDT",
		},
	},
	"NEWB": {
		{
			ID:          10888,
			Slug:        "newb-farm",
			Name:        "NewB.Farm",
			BinanceUSDT: "NEWBUSDT",
		},
	},
	"STACK": {
		{
			ID:   9201,
			Slug: "stackos",
			Name: "StackOs",
		},
		{
			ID:   8629,
			Slug: "stacker-ventures",
			Name: "Stacker Ventures",
		},
	},
	"JOINT": {
		{
			ID:          2745,
			Slug:        "joint-ventures",
			Name:        "Joint Ventures",
			BinanceUSDT: "JOINTUSDT",
		},
	},
	"XAUT": {
		{
			ID:          5176,
			Slug:        "tether-gold",
			Name:        "Tether Gold",
			BinanceUSDT: "XAUTUSDT",
		},
	},
	"SAFESUN": {
		{
			ID:          9294,
			Slug:        "safesun-crypto",
			Name:        "SAFESUN",
			BinanceUSDT: "SAFESUNUSDT",
		},
	},
	"DUMPDOGE": {
		{
			ID:          10892,
			Slug:        "dump-doge",
			Name:        "DUMP DOGE",
			BinanceUSDT: "DUMPDOGEUSDT",
		},
	},
	"DDIM": {
		{
			ID:          6611,
			Slug:        "duckdaodime",
			Name:        "DuckDaoDime",
			BinanceUSDT: "DDIMUSDT",
		},
	},
	"BUB": {
		{
			ID:          918,
			Slug:        "bubble",
			Name:        "Bubble",
			BinanceUSDT: "BUBUSDT",
		},
	},
	"NFTX": {
		{
			ID:          8191,
			Slug:        "nftx",
			Name:        "NFTX",
			BinanceUSDT: "NFTXUSDT",
		},
	},
	"FRMX": {
		{
			ID:          7631,
			Slug:        "frmx-token",
			Name:        "FRMx Token",
			BinanceUSDT: "FRMXUSDT",
		},
	},
	"BP": {
		{
			ID:          10904,
			Slug:        "bunnypark",
			Name:        "BunnyPark",
			BinanceUSDT: "BPUSDT",
		},
	},
	"SUREBETS": {
		{
			ID:          10505,
			Slug:        "surebets-online",
			Name:        "SureBets Online",
			BinanceUSDT: "SUREBETSUSDT",
		},
	},
	"BTZC": {
		{
			ID:          4867,
			Slug:        "beatzcoin",
			Name:        "BeatzCoin",
			BinanceUSDT: "BTZCUSDT",
		},
	},
	"GZRO": {
		{
			ID:          3488,
			Slug:        "gravity",
			Name:        "Gravity",
			BinanceUSDT: "GZROUSDT",
		},
	},
	"RVF": {
		{
			ID:          9176,
			Slug:        "rocket-vault",
			Name:        "Rocket Vault",
			BinanceUSDT: "RVFUSDT",
		},
	},
	"UFARM": {
		{
			ID:          9262,
			Slug:        "unifarm",
			Name:        "UniFarm",
			BinanceUSDT: "UFARMUSDT",
		},
	},
	"RCN": {
		{
			ID:          2096,
			Slug:        "ripio-credit-network",
			Name:        "Ripio Credit Network",
			BinanceUSDT: "RCNUSDT",
		},
	},
	"TEP": {
		{
			ID:          4677,
			Slug:        "tepleton",
			Name:        "Tepleton",
			BinanceUSDT: "TEPUSDT",
		},
	},
	"OOE": {
		{
			ID:          9938,
			Slug:        "openocean",
			Name:        "OpenOcean",
			BinanceUSDT: "OOEUSDT",
		},
	},
	"YFFII": {
		{
			ID:          6975,
			Slug:        "yffii-finance",
			Name:        "YFFII Finance",
			BinanceUSDT: "YFFIIUSDT",
		},
	},
	"DOKI": {
		{
			ID:          7376,
			Slug:        "doki-doki-finance",
			Name:        "Doki Doki Finance",
			BinanceUSDT: "DOKIUSDT",
		},
	},
	"EFG": {
		{
			ID:          7872,
			Slug:        "ecoc-financial-growth",
			Name:        "ECOC Financial Growth",
			BinanceUSDT: "EFGUSDT",
		},
	},
	"STOP": {
		{
			ID:          6766,
			Slug:        "satopay-network",
			Name:        "Satopay Network",
			BinanceUSDT: "STOPUSDT",
		},
	},
	"MOONRABBIT": {
		{
			ID:          10249,
			Slug:        "moonrabbit-money",
			Name:        "MoonRabbit",
			BinanceUSDT: "MOONRABBITUSDT",
		},
	},
	"SPRINK": {
		{
			ID:          8844,
			Slug:        "sprink",
			Name:        "SPRINK",
			BinanceUSDT: "SPRINKUSDT",
		},
	},
	"XAP": {
		{
			ID:          3159,
			Slug:        "apollon",
			Name:        "Apollon",
			BinanceUSDT: "XAPUSDT",
		},
	},
	"UNISTAKE": {
		{
			ID:          7512,
			Slug:        "unistake",
			Name:        "Unistake",
			BinanceUSDT: "UNISTAKEUSDT",
		},
	},
	"LICK": {
		{
			ID:          10812,
			Slug:        "vitalick-neuterin",
			Name:        "VITALICK NEUTERIN",
			BinanceUSDT: "LICKUSDT",
		},
	},
	"EDU": {
		{
			ID:          2734,
			Slug:        "edu-coin",
			Name:        "EduCoin",
			BinanceUSDT: "EDUUSDT",
		},
	},
	"DCD": {
		{
			ID:          7159,
			Slug:        "digital-currency-daily",
			Name:        "Digital Currency Daily",
			BinanceUSDT: "DCDUSDT",
		},
	},
	"REN": {
		{
			ID:          2539,
			Slug:        "ren",
			Name:        "Ren",
			BinanceUSDT: "RENUSDT",
		},
	},
	"XTZBULL": {
		{
			ID:          5464,
			Slug:        "3x-long-tezos-token",
			Name:        "3x Long Tezos Token",
			BinanceUSDT: "XTZBULLUSDT",
		},
	},
	"BOS": {
		{
			ID:   2095,
			Slug: "boscoin",
			Name: "BOScoin",
		},
		{
			ID:   4629,
			Slug: "boscore",
			Name: "BOSCore",
		},
	},
	"XDB": {
		{
			ID:          4566,
			Slug:        "digitalbits",
			Name:        "DigitalBits",
			BinanceUSDT: "XDBUSDT",
		},
	},
	"QAC": {
		{
			ID:          3419,
			Slug:        "quasarcoin",
			Name:        "Quasarcoin",
			BinanceUSDT: "QACUSDT",
		},
	},
	"BDX": {
		{
			ID:          3987,
			Slug:        "beldex",
			Name:        "Beldex",
			BinanceUSDT: "BDXUSDT",
		},
	},
	"CBC": {
		{
			ID:   2855,
			Slug: "casino-betting-coin",
			Name: "CBC.network",
		},
		{
			ID:   5563,
			Slug: "cryptobharatcoin",
			Name: "CryptoBharatCoin",
		},
		{
			ID:   4633,
			Slug: "cryptobosscoin",
			Name: "CryptoBossCoin",
		},
	},
	"POLC": {
		{
			ID:          8549,
			Slug:        "polkacity",
			Name:        "Polkacity",
			BinanceUSDT: "POLCUSDT",
		},
	},
	"VAL": {
		{
			ID:   1154,
			Slug: "validity",
			Name: "Validity",
		},
		{
			ID:   7876,
			Slug: "sora-validator-token",
			Name: "Sora Validator Token",
		},
		{
			ID:   9516,
			Slug: "valkyrie-network",
			Name: "Valkyrie Network",
		},
	},
	"RNBW": {
		{
			ID:          10429,
			Slug:        "halodao",
			Name:        "HaloDAO",
			BinanceUSDT: "RNBWUSDT",
		},
	},
	"ONG": {
		{
			ID:   3217,
			Slug: "ontology-gas",
			Name: "Ontology Gas",
		},
		{
			ID:   2240,
			Slug: "ongsocial",
			Name: "SoMee.Social",
		},
	},
	"WETH": {
		{
			ID:          2396,
			Slug:        "weth",
			Name:        "WETH",
			BinanceUSDT: "WETHUSDT",
		},
	},
	"MUE": {
		{
			ID:          706,
			Slug:        "monetaryunit",
			Name:        "MonetaryUnit",
			BinanceUSDT: "MUEUSDT",
		},
	},
	"EXM": {
		{
			ID:          4974,
			Slug:        "exmo-coin",
			Name:        "EXMO Coin",
			BinanceUSDT: "EXMUSDT",
		},
	},
	"PRARE": {
		{
			ID:          9544,
			Slug:        "polkarare",
			Name:        "POLKARARE",
			BinanceUSDT: "PRAREUSDT",
		},
	},
	"VENUS": {
		{
			ID:          10073,
			Slug:        "venusia",
			Name:        "Venusia",
			BinanceUSDT: "VENUSUSDT",
		},
	},
	"TENDIE": {
		{
			ID:          11034,
			Slug:        "tendieswap",
			Name:        "TendieSwap",
			BinanceUSDT: "TENDIEUSDT",
		},
	},
	"COR": {
		{
			ID:          7398,
			Slug:        "coreto",
			Name:        "Coreto",
			BinanceUSDT: "CORUSDT",
		},
	},
	"KCAL": {
		{
			ID:          7298,
			Slug:        "phantasma-energy",
			Name:        "Phantasma Energy",
			BinanceUSDT: "KCALUSDT",
		},
	},
	"RELI": {
		{
			ID:          9095,
			Slug:        "relite-finance",
			Name:        "Relite Finance",
			BinanceUSDT: "RELIUSDT",
		},
	},
	"DAG": {
		{
			ID:          2868,
			Slug:        "constellation",
			Name:        "Constellation",
			BinanceUSDT: "DAGUSDT",
		},
	},
	"HOME": {
		{
			ID:          10342,
			Slug:        "homecoin",
			Name:        "HomeCoin",
			BinanceUSDT: "HOMEUSDT",
		},
	},
	"A2A": {
		{
			ID:          8745,
			Slug:        "a2a-50x-com",
			Name:        "A2A",
			BinanceUSDT: "A2AUSDT",
		},
	},
	"AREPA": {
		{
			ID:          3133,
			Slug:        "arepacoin",
			Name:        "Arepacoin",
			BinanceUSDT: "AREPAUSDT",
		},
	},
	"DYP": {
		{
			ID:          8080,
			Slug:        "defi-yield-protocol",
			Name:        "DeFi Yield Protocol",
			BinanceUSDT: "DYPUSDT",
		},
	},
	"BNBDOWN": {
		{
			ID:          7010,
			Slug:        "bnbdown",
			Name:        "BNBDOWN",
			BinanceUSDT: "BNBDOWNUSDT",
		},
	},
	"SATT": {
		{
			ID:          7244,
			Slug:        "satt",
			Name:        "SaTT",
			BinanceUSDT: "SATTUSDT",
		},
	},
	"SOME": {
		{
			ID:          9599,
			Slug:        "mixsome",
			Name:        "Mixsome",
			BinanceUSDT: "SOMEUSDT",
		},
	},
	"CCE": {
		{
			ID:          7655,
			Slug:        "cloudcoin",
			Name:        "CloudCoin",
			BinanceUSDT: "CCEUSDT",
		},
	},
	"BBTC": {
		{
			ID:          10869,
			Slug:        "baby-bitcoin",
			Name:        "Baby Bitcoin",
			BinanceUSDT: "BBTCUSDT",
		},
	},
	"USDFL": {
		{
			ID:          8437,
			Slug:        "usdfreeliquidity",
			Name:        "USDFreeLiquidity",
			BinanceUSDT: "USDFLUSDT",
		},
	},
	"TZC": {
		{
			ID:          2002,
			Slug:        "trezarcoin",
			Name:        "TrezarCoin",
			BinanceUSDT: "TZCUSDT",
		},
	},
	"NBL": {
		{
			ID:          10982,
			Slug:        "nobility",
			Name:        "Nobility",
			BinanceUSDT: "NBLUSDT",
		},
	},
	"CITY": {
		{
			ID:   10049,
			Slug: "manchester-city-fan-token",
			Name: "Manchester City Fan Token",
		},
		{
			ID:   4545,
			Slug: "city-coin",
			Name: "City Coin",
		},
	},
	"DEXT": {
		{
			ID:          5866,
			Slug:        "dextools",
			Name:        "DEXTools",
			BinanceUSDT: "DEXTUSDT",
		},
	},
	"DUSK": {
		{
			ID:          4092,
			Slug:        "dusk-network",
			Name:        "Dusk Network",
			BinanceUSDT: "DUSKUSDT",
		},
	},
	"YAP": {
		{
			ID:          4899,
			Slug:        "yap-stone",
			Name:        "Yap Stone",
			BinanceUSDT: "YAPUSDT",
		},
	},
	"ALICE": {
		{
			ID:          8766,
			Slug:        "myneighboralice",
			Name:        "MyNeighborAlice",
			BinanceUSDT: "ALICEUSDT",
		},
	},
	"CYL": {
		{
			ID:          3580,
			Slug:        "crystal-token",
			Name:        "Crystal Token",
			BinanceUSDT: "CYLUSDT",
		},
	},
	"BAN": {
		{
			ID:          4704,
			Slug:        "banano",
			Name:        "Banano",
			BinanceUSDT: "BANUSDT",
		},
	},
	"MIMO": {
		{
			ID:          8638,
			Slug:        "mimosa",
			Name:        "MIMOSA",
			BinanceUSDT: "MIMOUSDT",
		},
	},
	"RYIU": {
		{
			ID:          8871,
			Slug:        "ryi-unity",
			Name:        "RYI Unity",
			BinanceUSDT: "RYIUUSDT",
		},
	},
	"WANBTC": {
		{
			ID:          8650,
			Slug:        "wanbtc",
			Name:        "wanBTC",
			BinanceUSDT: "WANBTCUSDT",
		},
	},
	"HYBN": {
		{
			ID:          5934,
			Slug:        "hey-bitcoin",
			Name:        "Hey Bitcoin",
			BinanceUSDT: "HYBNUSDT",
		},
	},
	"SNGLS": {
		{
			ID:          1409,
			Slug:        "singulardtv",
			Name:        "SingularDTV",
			BinanceUSDT: "SNGLSUSDT",
		},
	},
	"BCNT": {
		{
			ID:          4808,
			Slug:        "bincentive",
			Name:        "Bincentive",
			BinanceUSDT: "BCNTUSDT",
		},
	},
	"ASI": {
		{
			ID:          8876,
			Slug:        "asi-finance-token",
			Name:        "ASI finance",
			BinanceUSDT: "ASIUSDT",
		},
	},
	"LPNT": {
		{
			ID:          8501,
			Slug:        "luxurious-pro-network-token",
			Name:        "Luxurious Pro Network Token",
			BinanceUSDT: "LPNTUSDT",
		},
	},
	"HGOLD": {
		{
			ID:          8256,
			Slug:        "hollygold",
			Name:        "HollyGold",
			BinanceUSDT: "HGOLDUSDT",
		},
	},
	"XDNA": {
		{
			ID:   7262,
			Slug: "extradna",
			Name: "extraDNA",
		},
		{
			ID:   3125,
			Slug: "xdna",
			Name: "XDNA",
		},
	},
	"BEPRO": {
		{
			ID:          5062,
			Slug:        "bepro-network",
			Name:        "BEPRO Network",
			BinanceUSDT: "BEPROUSDT",
		},
	},
	"GRIN": {
		{
			ID:          3709,
			Slug:        "grin",
			Name:        "Grin",
			BinanceUSDT: "GRINUSDT",
		},
	},
	"BDOGE": {
		{
			ID:          9731,
			Slug:        "blue-eyes-white-doge",
			Name:        "Blue Eyes White Doge",
			BinanceUSDT: "BDOGEUSDT",
		},
	},
	"SHIB": {
		{
			ID:          5994,
			Slug:        "shiba-inu",
			Name:        "SHIBA INU",
			BinanceUSDT: "SHIBUSDT",
		},
	},
	"PSIX": {
		{
			ID:          8575,
			Slug:        "propersix",
			Name:        "ProperSix",
			BinanceUSDT: "PSIXUSDT",
		},
	},
	"SAFEBANK": {
		{
			ID:          10015,
			Slug:        "safebank-yes",
			Name:        "SafeBank YES",
			BinanceUSDT: "SAFEBANKUSDT",
		},
	},
	"BNK": {
		{
			ID:          2842,
			Slug:        "bankera",
			Name:        "Bankera",
			BinanceUSDT: "BNKUSDT",
		},
	},
	"REM": {
		{
			ID:          2546,
			Slug:        "remme",
			Name:        "Remme",
			BinanceUSDT: "REMUSDT",
		},
	},
	"BQT": {
		{
			ID:          3231,
			Slug:        "blockchain-quotations-index-token",
			Name:        "Blockchain Quotations Index Token",
			BinanceUSDT: "BQTUSDT",
		},
	},
	"PBTC": {
		{
			ID:          5434,
			Slug:        "ptokens-btc",
			Name:        "pTokens BTC",
			BinanceUSDT: "PBTCUSDT",
		},
	},
	"LBC": {
		{
			ID:          1298,
			Slug:        "library-credit",
			Name:        "LBRY Credits",
			BinanceUSDT: "LBCUSDT",
		},
	},
	"STPT": {
		{
			ID:          4006,
			Slug:        "standard-tokenization-protocol",
			Name:        "Standard Tokenization Protocol",
			BinanceUSDT: "STPTUSDT",
		},
	},
	"MOR": {
		{
			ID:   5304,
			Slug: "morcrypto-coin",
			Name: "MorCrypto Coin",
		},
		{
			ID:   9571,
			Slug: "mirror-farm",
			Name: "Mirror Farm",
		},
	},
	"LON": {
		{
			ID:          8083,
			Slug:        "tokenlon-network-token",
			Name:        "Tokenlon Network Token",
			BinanceUSDT: "LONUSDT",
		},
	},
	"XYZ": {
		{
			ID:          10979,
			Slug:        "universe-xyz",
			Name:        "Universe.XYZ",
			BinanceUSDT: "XYZUSDT",
		},
	},
	"LET": {
		{
			ID:          2468,
			Slug:        "linkeye",
			Name:        "LinkEye",
			BinanceUSDT: "LETUSDT",
		},
	},
	"ATB": {
		{
			ID:          1970,
			Slug:        "atbcoin",
			Name:        "ATBCoin",
			BinanceUSDT: "ATBUSDT",
		},
	},
	"IOC": {
		{
			ID:          495,
			Slug:        "iocoin",
			Name:        "I/O Coin",
			BinanceUSDT: "IOCUSDT",
		},
	},
	"CON": {
		{
			ID:   3866,
			Slug: "conun",
			Name: "CONUN",
		},
		{
			ID:   9005,
			Slug: "connectico",
			Name: "Connectico",
		},
		{
			ID:   9014,
			Slug: "converter-finance",
			Name: "Converter.Finance",
		},
	},
	"STOGE": {
		{
			ID:          9381,
			Slug:        "stoner-doge",
			Name:        "Stoner Doge Finance",
			BinanceUSDT: "STOGEUSDT",
		},
	},
	"BTY": {
		{
			ID:          6224,
			Slug:        "bityuan",
			Name:        "Bityuan",
			BinanceUSDT: "BTYUSDT",
		},
	},
	"CBM": {
		{
			ID:          4253,
			Slug:        "cryptobonusmiles",
			Name:        "CryptoBonusMiles",
			BinanceUSDT: "CBMUSDT",
		},
	},
	"LBA": {
		{
			ID:          2760,
			Slug:        "libra-credit",
			Name:        "Cred",
			BinanceUSDT: "LBAUSDT",
		},
	},
	"FOCV": {
		{
			ID:          6002,
			Slug:        "focv",
			Name:        "FOCV",
			BinanceUSDT: "FOCVUSDT",
		},
	},
	"MFI": {
		{
			ID:          8411,
			Slug:        "marginswap",
			Name:        "Marginswap",
			BinanceUSDT: "MFIUSDT",
		},
	},
	"TEL": {
		{
			ID:          2394,
			Slug:        "telcoin",
			Name:        "Telcoin",
			BinanceUSDT: "TELUSDT",
		},
	},
	"HYN": {
		{
			ID:          3695,
			Slug:        "hyperion",
			Name:        "Hyperion",
			BinanceUSDT: "HYNUSDT",
		},
	},
	"PUFFY": {
		{
			ID:          10168,
			Slug:        "puffydog-coin",
			Name:        "Puffydog Coin",
			BinanceUSDT: "PUFFYUSDT",
		},
	},
	"XBP": {
		{
			ID:          2614,
			Slug:        "blitzpredict",
			Name:        "BlitzPick",
			BinanceUSDT: "XBPUSDT",
		},
	},
	"FLG": {
		{
			ID:          4842,
			Slug:        "folgory-coin",
			Name:        "Folgory Coin",
			BinanceUSDT: "FLGUSDT",
		},
	},
	"AGA": {
		{
			ID:          6404,
			Slug:        "aga",
			Name:        "AGA Token",
			BinanceUSDT: "AGAUSDT",
		},
	},
	"YFE": {
		{
			ID:          7250,
			Slug:        "yfe-money",
			Name:        "YFE Money",
			BinanceUSDT: "YFEUSDT",
		},
	},
	"UC": {
		{
			ID:          3259,
			Slug:        "youlive-coin",
			Name:        "YouLive Coin",
			BinanceUSDT: "UCUSDT",
		},
	},
	"ARRO": {
		{
			ID:          8841,
			Slug:        "arro-social",
			Name:        "Arro Social",
			BinanceUSDT: "ARROUSDT",
		},
	},
	"DITTO": {
		{
			ID:          8086,
			Slug:        "ditto",
			Name:        "Ditto",
			BinanceUSDT: "DITTOUSDT",
		},
	},
	"MOONI": {
		{
			ID:          9465,
			Slug:        "mooni-defi",
			Name:        "Mooni DeFi",
			BinanceUSDT: "MOONIUSDT",
		},
	},
	"ENERGYX": {
		{
			ID:          10325,
			Slug:        "safe-energy",
			Name:        "Safe Energy",
			BinanceUSDT: "ENERGYXUSDT",
		},
	},
	"DUKE": {
		{
			ID:          10552,
			Slug:        "duke-inu-token",
			Name:        "DUKE INU TOKEN",
			BinanceUSDT: "DUKEUSDT",
		},
	},
	"ZCR": {
		{
			ID:          3158,
			Slug:        "zcore",
			Name:        "ZCore",
			BinanceUSDT: "ZCRUSDT",
		},
	},
	"HOMT": {
		{
			ID:          5722,
			Slug:        "homt",
			Name:        "HOMT",
			BinanceUSDT: "HOMTUSDT",
		},
	},
	"FFYI": {
		{
			ID:          6935,
			Slug:        "fiscus-fyi",
			Name:        "Fiscus.fyi",
			BinanceUSDT: "FFYIUSDT",
		},
	},
	"GNBT": {
		{
			ID:          10492,
			Slug:        "genebank-token",
			Name:        "Genebank Token",
			BinanceUSDT: "GNBTUSDT",
		},
	},
	"UCT": {
		{
			ID:   9539,
			Slug: "unitedcrowd",
			Name: "UnitedCrowd",
		},
		{
			ID:   2871,
			Slug: "ubique-chain-of-things",
			Name: "Ubique Chain Of Things",
		},
	},
	"COV": {
		{
			ID:          2342,
			Slug:        "covesting",
			Name:        "Covesting",
			BinanceUSDT: "COVUSDT",
		},
	},
	"LOG": {
		{
			ID:          859,
			Slug:        "woodcoin",
			Name:        "Woodcoin",
			BinanceUSDT: "LOGUSDT",
		},
	},
	"GUSDT": {
		{
			ID:          8118,
			Slug:        "gusd-token",
			Name:        "Global Utility Smart Digital Token",
			BinanceUSDT: "GUSDTUSDT",
		},
	},
	"FRED": {
		{
			ID:          5109,
			Slug:        "fred-energy",
			Name:        "FRED Energy",
			BinanceUSDT: "FREDUSDT",
		},
	},
	"BBT": {
		{
			ID:          10201,
			Slug:        "bitbook",
			Name:        "BitBook",
			BinanceUSDT: "BBTUSDT",
		},
	},
	"CV": {
		{
			ID:          2450,
			Slug:        "carvertical",
			Name:        "carVertical",
			BinanceUSDT: "CVUSDT",
		},
	},
	"VEST": {
		{
			ID:          3607,
			Slug:        "vestchain",
			Name:        "VestChain",
			BinanceUSDT: "VESTUSDT",
		},
	},
	"TCO": {
		{
			ID:          8556,
			Slug:        "tcoin-token",
			Name:        "Tcoin",
			BinanceUSDT: "TCOUSDT",
		},
	},
	"IBFK": {
		{
			ID:          8884,
			Slug:        "istanbul-basaksehir-fan-token",
			Name:        "İstanbul Başakşehir Fan Token",
			BinanceUSDT: "IBFKUSDT",
		},
	},
	"SAKE": {
		{
			ID:          6997,
			Slug:        "sake-token",
			Name:        "SakeToken",
			BinanceUSDT: "SAKEUSDT",
		},
	},
	"BTSC": {
		{
			ID:          6392,
			Slug:        "bts-coin",
			Name:        "BTS Coin",
			BinanceUSDT: "BTSCUSDT",
		},
	},
	"BRZE": {
		{
			ID:          3519,
			Slug:        "breezecoin",
			Name:        "Breezecoin",
			BinanceUSDT: "BRZEUSDT",
		},
	},
	"VBNT": {
		{
			ID:          8622,
			Slug:        "bancor-governance-token",
			Name:        "Bancor Governance Token",
			BinanceUSDT: "VBNTUSDT",
		},
	},
	"BUT": {
		{
			ID:          3258,
			Slug:        "bitup-token",
			Name:        "BitUP Token",
			BinanceUSDT: "BUTUSDT",
		},
	},
	"TRY": {
		{
			ID:   3970,
			Slug: "trias",
			Name: "Trias (old)",
		},
		{
			ID:   9097,
			Slug: "try-finance",
			Name: "Try.Finance",
		},
	},
	"IMG": {
		{
			ID:          4156,
			Slug:        "imagecoin",
			Name:        "ImageCoin",
			BinanceUSDT: "IMGUSDT",
		},
	},
	"SEMI": {
		{
			ID:          6823,
			Slug:        "semitoken",
			Name:        "Semitoken",
			BinanceUSDT: "SEMIUSDT",
		},
	},
	"NFTS": {
		{
			ID:          10311,
			Slug:        "nft-stars",
			Name:        "NFT STARS",
			BinanceUSDT: "NFTSUSDT",
		},
	},
	"DTH": {
		{
			ID:          2528,
			Slug:        "dether",
			Name:        "Dether",
			BinanceUSDT: "DTHUSDT",
		},
	},
	"ANCT": {
		{
			ID:          4901,
			Slug:        "anchor",
			Name:        "Anchor",
			BinanceUSDT: "ANCTUSDT",
		},
	},
	"KAWA": {
		{
			ID:          10640,
			Slug:        "kawakami-inu",
			Name:        "Kawakami Inu",
			BinanceUSDT: "KAWAUSDT",
		},
	},
	"PEPECASH": {
		{
			ID:          1405,
			Slug:        "pepe-cash",
			Name:        "Pepe Cash",
			BinanceUSDT: "PEPECASHUSDT",
		},
	},
	"$KING": {
		{
			ID:          7569,
			Slug:        "kingswap",
			Name:        "King Swap",
			BinanceUSDT: "$KINGUSDT",
		},
	},
	"XFT": {
		{
			ID:          6236,
			Slug:        "offshift",
			Name:        "Offshift",
			BinanceUSDT: "XFTUSDT",
		},
	},
	"PRE": {
		{
			ID:          2245,
			Slug:        "presearch",
			Name:        "Presearch",
			BinanceUSDT: "PREUSDT",
		},
	},
	"SAPP": {
		{
			ID:          4121,
			Slug:        "sapphire",
			Name:        "Sapphire",
			BinanceUSDT: "SAPPUSDT",
		},
	},
	"DELTA": {
		{
			ID:          8994,
			Slug:        "delta-finance",
			Name:        "Delta",
			BinanceUSDT: "DELTAUSDT",
		},
	},
	"USDX": {
		{
			ID:   6651,
			Slug: "usdx-kava",
			Name: "USDX [Kava]",
		},
		{
			ID:   4557,
			Slug: "usdx-lighthouse",
			Name: "USDX [Lighthouse]",
		},
		{
			ID:   4621,
			Slug: "dforce-usdx",
			Name: "dForce USDx",
		},
	},
	"BRZ": {
		{
			ID:          4139,
			Slug:        "brz",
			Name:        "Brazilian Digital Token",
			BinanceUSDT: "BRZUSDT",
		},
	},
	"JUI": {
		{
			ID:          6169,
			Slug:        "juiice",
			Name:        "JUIICE",
			BinanceUSDT: "JUIUSDT",
		},
	},
	"MLT": {
		{
			ID:   9764,
			Slug: "milc-platform",
			Name: "MILC Platform",
		},
		{
			ID:   10567,
			Slug: "modern-liquidity-token",
			Name: "Modern Liquidity Token",
		},
	},
	"RMT": {
		{
			ID:          2527,
			Slug:        "sureremit",
			Name:        "SureRemit",
			BinanceUSDT: "RMTUSDT",
		},
	},
	"DEFT": {
		{
			ID:          10044,
			Slug:        "defi-factory-token",
			Name:        "DeFi Factory Token",
			BinanceUSDT: "DEFTUSDT",
		},
	},
	"NOA": {
		{
			ID:          9997,
			Slug:        "noa-play",
			Name:        "NOA PLAY",
			BinanceUSDT: "NOAUSDT",
		},
	},
	"RAMP": {
		{
			ID:          7463,
			Slug:        "ramp",
			Name:        "RAMP",
			BinanceUSDT: "RAMPUSDT",
		},
	},
	"BWF": {
		{
			ID:          7176,
			Slug:        "beowulf",
			Name:        "Beowulf",
			BinanceUSDT: "BWFUSDT",
		},
	},
	"VRSC": {
		{
			ID:          5049,
			Slug:        "veruscoin",
			Name:        "VerusCoin",
			BinanceUSDT: "VRSCUSDT",
		},
	},
	"PYN": {
		{
			ID:          3323,
			Slug:        "paycent",
			Name:        "PAYCENT",
			BinanceUSDT: "PYNUSDT",
		},
	},
	"JFC": {
		{
			ID:          4568,
			Slug:        "jfin",
			Name:        "JFIN",
			BinanceUSDT: "JFCUSDT",
		},
	},
	"BVL": {
		{
			ID:          7811,
			Slug:        "bullswap-exchange",
			Name:        "Bullswap Exchange",
			BinanceUSDT: "BVLUSDT",
		},
	},
	"RITO": {
		{
			ID:          4264,
			Slug:        "ritocoin",
			Name:        "Ritocoin",
			BinanceUSDT: "RITOUSDT",
		},
	},
	"TRONX": {
		{
			ID:          8221,
			Slug:        "tronx-coin",
			Name:        "Tronx Coin",
			BinanceUSDT: "TRONXUSDT",
		},
	},
	"TRIB": {
		{
			ID:          7170,
			Slug:        "contribute",
			Name:        "Contribute",
			BinanceUSDT: "TRIBUSDT",
		},
	},
	"CFXQ": {
		{
			ID:          9070,
			Slug:        "cfx-quantum",
			Name:        "CFX Quantum",
			BinanceUSDT: "CFXQUSDT",
		},
	},
	"IXC": {
		{
			ID:          13,
			Slug:        "ixcoin",
			Name:        "Ixcoin",
			BinanceUSDT: "IXCUSDT",
		},
	},
	"KIN": {
		{
			ID:          1993,
			Slug:        "kin",
			Name:        "Kin",
			BinanceUSDT: "KINUSDT",
		},
	},
	"CTSI": {
		{
			ID:          5444,
			Slug:        "cartesi",
			Name:        "Cartesi",
			BinanceUSDT: "CTSIUSDT",
		},
	},
	"HAND": {
		{
			ID:          3181,
			Slug:        "showhand",
			Name:        "ShowHand",
			BinanceUSDT: "HANDUSDT",
		},
	},
	"SEUR": {
		{
			ID:          10419,
			Slug:        "seur",
			Name:        "sEUR",
			BinanceUSDT: "SEURUSDT",
		},
	},
	"SBEAR": {
		{
			ID:          9926,
			Slug:        "ybearswap",
			Name:        "yBEARSwap",
			BinanceUSDT: "SBEARUSDT",
		},
	},
	"SAL": {
		{
			ID:          6874,
			Slug:        "salmonswap",
			Name:        "SalmonSwap",
			BinanceUSDT: "SALUSDT",
		},
	},
	"COFI": {
		{
			ID:   2478,
			Slug: "coinfi",
			Name: "CoinFi",
		},
		{
			ID:   7381,
			Slug: "cofix",
			Name: "CoFiX",
		},
	},
	"PHTR": {
		{
			ID:          9423,
			Slug:        "phuture",
			Name:        "Phuture",
			BinanceUSDT: "PHTRUSDT",
		},
	},
	"TH": {
		{
			ID:          7636,
			Slug:        "team-heretics-fan-token",
			Name:        "Team Heretics Fan Token",
			BinanceUSDT: "THUSDT",
		},
	},
	"TAP": {
		{
			ID:          8463,
			Slug:        "tapmydata",
			Name:        "Tapmydata",
			BinanceUSDT: "TAPUSDT",
		},
	},
	"MNS": {
		{
			ID:          5702,
			Slug:        "monnos",
			Name:        "MONNOS",
			BinanceUSDT: "MNSUSDT",
		},
	},
	"LUNAR": {
		{
			ID:   10852,
			Slug: "lunarswap",
			Name: "LunarSwap",
		},
		{
			ID:   9810,
			Slug: "lunar-highway",
			Name: "Lunar Highway",
		},
	},
	"AAVE": {
		{
			ID:          7278,
			Slug:        "aave",
			Name:        "Aave",
			BinanceUSDT: "AAVEUSDT",
		},
	},
	"BNC": {
		{
			ID:          3222,
			Slug:        "bionic",
			Name:        "Bionic",
			BinanceUSDT: "BNCUSDT",
		},
	},
	"JUL": {
		{
			ID:          6937,
			Slug:        "justliquidity",
			Name:        "JustLiquidity",
			BinanceUSDT: "JULUSDT",
		},
	},
	"KOBO": {
		{
			ID:          815,
			Slug:        "kobocoin",
			Name:        "Kobocoin",
			BinanceUSDT: "KOBOUSDT",
		},
	},
	"ES": {
		{
			ID:          4860,
			Slug:        "era-swap",
			Name:        "Era Swap",
			BinanceUSDT: "ESUSDT",
		},
	},
	"AKN": {
		{
			ID:          4618,
			Slug:        "akoin",
			Name:        "Akoin",
			BinanceUSDT: "AKNUSDT",
		},
	},
	"SPA": {
		{
			ID:          6715,
			Slug:        "sperax",
			Name:        "Sperax",
			BinanceUSDT: "SPAUSDT",
		},
	},
	"BLACK": {
		{
			ID:   9450,
			Slug: "blackhole-protocol",
			Name: "BLACKHOLE PROTOCOL",
		},
		{
			ID:   3253,
			Slug: "eosblack",
			Name: "eosBLACK",
		},
	},
	"FB": {
		{
			ID:   7897,
			Slug: "facebook-tokenized-stock-ftx",
			Name: "Facebook tokenized stock FTX",
		},
		{
			ID:   7926,
			Slug: "facebook-tokenized-stock-bittrex",
			Name: "Facebook tokenized stock Bittrex",
		},
	},
	"CBSN": {
		{
			ID:          9418,
			Slug:        "blockswap-network",
			Name:        "BlockSwap Network",
			BinanceUSDT: "CBSNUSDT",
		},
	},
	"IND": {
		{
			ID:          1967,
			Slug:        "indorse-token",
			Name:        "Indorse Token",
			BinanceUSDT: "INDUSDT",
		},
	},
	"KINE": {
		{
			ID:          8790,
			Slug:        "kine",
			Name:        "KINE",
			BinanceUSDT: "KINEUSDT",
		},
	},
	"B2G": {
		{
			ID:          3684,
			Slug:        "bitcoiin",
			Name:        "Bitcoiin",
			BinanceUSDT: "B2GUSDT",
		},
	},
	"ELY": {
		{
			ID:          2944,
			Slug:        "elysian",
			Name:        "Elysian",
			BinanceUSDT: "ELYUSDT",
		},
	},
	"BRO": {
		{
			ID:          1754,
			Slug:        "bitradio",
			Name:        "Bitradio",
			BinanceUSDT: "BROUSDT",
		},
	},
	"BHAO": {
		{
			ID:          7503,
			Slug:        "bithao",
			Name:        "Bithao",
			BinanceUSDT: "BHAOUSDT",
		},
	},
	"CUSDT": {
		{
			ID:          5745,
			Slug:        "compound-usdt",
			Name:        "Compound USDT",
			BinanceUSDT: "CUSDTUSDT",
		},
	},
	"DVPN": {
		{
			ID:          2643,
			Slug:        "sentinel",
			Name:        "Sentinel",
			BinanceUSDT: "DVPNUSDT",
		},
	},
	"XIASI": {
		{
			ID:          10636,
			Slug:        "xiasi-inu",
			Name:        "Xiasi Inu",
			BinanceUSDT: "XIASIUSDT",
		},
	},
	"PMA": {
		{
			ID:          3164,
			Slug:        "pumapay",
			Name:        "PumaPay",
			BinanceUSDT: "PMAUSDT",
		},
	},
	"CRON": {
		{
			ID:          4309,
			Slug:        "cryptocean",
			Name:        "Cryptocean",
			BinanceUSDT: "CRONUSDT",
		},
	},
	"SEND": {
		{
			ID:          2255,
			Slug:        "social-send",
			Name:        "Social Send",
			BinanceUSDT: "SENDUSDT",
		},
	},
	"ELNC": {
		{
			ID:          9858,
			Slug:        "eloniumcoin",
			Name:        "EloniumCoin",
			BinanceUSDT: "ELNCUSDT",
		},
	},
	"SIG": {
		{
			ID:          8598,
			Slug:        "xsigma",
			Name:        "xSigma",
			BinanceUSDT: "SIGUSDT",
		},
	},
	"DFC": {
		{
			ID:          10802,
			Slug:        "defi-city",
			Name:        "DeFi City",
			BinanceUSDT: "DFCUSDT",
		},
	},
	"KWIK": {
		{
			ID:          10101,
			Slug:        "kwikswap",
			Name:        "Kwikswap Protocol",
			BinanceUSDT: "KWIKUSDT",
		},
	},
	"CORX": {
		{
			ID:          7698,
			Slug:        "corionx",
			Name:        "CorionX",
			BinanceUSDT: "CORXUSDT",
		},
	},
	"POT": {
		{
			ID:          122,
			Slug:        "potcoin",
			Name:        "PotCoin",
			BinanceUSDT: "POTUSDT",
		},
	},
	"BONDLY": {
		{
			ID:          7931,
			Slug:        "bondly",
			Name:        "Bondly",
			BinanceUSDT: "BONDLYUSDT",
		},
	},
	"ARMOR": {
		{
			ID:          8309,
			Slug:        "armor",
			Name:        "ARMOR",
			BinanceUSDT: "ARMORUSDT",
		},
	},
	"YLAND": {
		{
			ID:          7234,
			Slug:        "yearn-land",
			Name:        "Yearn Land",
			BinanceUSDT: "YLANDUSDT",
		},
	},
	"VSO": {
		{
			ID:          9451,
			Slug:        "verso-token",
			Name:        "Verso Token",
			BinanceUSDT: "VSOUSDT",
		},
	},
	"REST": {
		{
			ID:          8089,
			Slug:        "restore",
			Name:        "Restore",
			BinanceUSDT: "RESTUSDT",
		},
	},
	"N1": {
		{
			ID:          9346,
			Slug:        "nftify",
			Name:        "NFTify",
			BinanceUSDT: "N1USDT",
		},
	},
	"THRT": {
		{
			ID:          2899,
			Slug:        "thrive-token",
			Name:        "Thrive Token",
			BinanceUSDT: "THRTUSDT",
		},
	},
	"ZYRO": {
		{
			ID:          7044,
			Slug:        "zyro",
			Name:        "Zyro",
			BinanceUSDT: "ZYROUSDT",
		},
	},
	"ZRX": {
		{
			ID:          1896,
			Slug:        "0x",
			Name:        "0x",
			BinanceUSDT: "ZRXUSDT",
		},
	},
	"ZEC": {
		{
			ID:          1437,
			Slug:        "zcash",
			Name:        "Zcash",
			BinanceUSDT: "ZECUSDT",
		},
	},
	"BSF": {
		{
			ID:          10856,
			Slug:        "babyspacefloki",
			Name:        "BabySpaceFloki",
			BinanceUSDT: "BSFUSDT",
		},
	},
	"HAUS": {
		{
			ID:          9016,
			Slug:        "daohaus",
			Name:        "DAOhaus",
			BinanceUSDT: "HAUSUSDT",
		},
	},
	"ORB": {
		{
			ID:          80,
			Slug:        "orbitcoin",
			Name:        "Orbitcoin",
			BinanceUSDT: "ORBUSDT",
		},
	},
	"LRC": {
		{
			ID:          1934,
			Slug:        "loopring",
			Name:        "Loopring",
			BinanceUSDT: "LRCUSDT",
		},
	},
	"HOODRAT": {
		{
			ID:          10071,
			Slug:        "hoodrat",
			Name:        "Hoodrat Finance",
			BinanceUSDT: "HOODRATUSDT",
		},
	},
	"ENERGY": {
		{
			ID:          10199,
			Slug:        "energy-token",
			Name:        "ENERGY Token",
			BinanceUSDT: "ENERGYUSDT",
		},
	},
	"RVP": {
		{
			ID:          9277,
			Slug:        "revolution-populi",
			Name:        "Revolution Populi",
			BinanceUSDT: "RVPUSDT",
		},
	},
	"CUT": {
		{
			ID:          4752,
			Slug:        "cutcoin",
			Name:        "CUTcoin",
			BinanceUSDT: "CUTUSDT",
		},
	},
	"PIG": {
		{
			ID:          8829,
			Slug:        "pig-finance",
			Name:        "Pig Finance",
			BinanceUSDT: "PIGUSDT",
		},
	},
	"CAP": {
		{
			ID:   5809,
			Slug: "cap",
			Name: "Cap",
		},
		{
			ID:   8076,
			Slug: "capital-finance",
			Name: "Capital.Finance",
		},
	},
	"COMFY": {
		{
			ID:          9627,
			Slug:        "comfytoken",
			Name:        "ComfyToken",
			BinanceUSDT: "COMFYUSDT",
		},
	},
	"PVT": {
		{
			ID:          4115,
			Slug:        "pivot-token",
			Name:        "Pivot Token",
			BinanceUSDT: "PVTUSDT",
		},
	},
	"COMBO": {
		{
			ID:          8259,
			Slug:        "furucombo",
			Name:        "Furucombo",
			BinanceUSDT: "COMBOUSDT",
		},
	},
	"BOR": {
		{
			ID:          7509,
			Slug:        "boringdao",
			Name:        "BoringDAO",
			BinanceUSDT: "BORUSDT",
		},
	},
	"CANN": {
		{
			ID:          506,
			Slug:        "cannabiscoin",
			Name:        "CannabisCoin",
			BinanceUSDT: "CANNUSDT",
		},
	},
	"DHV": {
		{
			ID:          8867,
			Slug:        "dehive",
			Name:        "DeHive",
			BinanceUSDT: "DHVUSDT",
		},
	},
	"IHT": {
		{
			ID:          2552,
			Slug:        "iht-real-estate-protocol",
			Name:        "IHT Real Estate Protocol",
			BinanceUSDT: "IHTUSDT",
		},
	},
	"GETH": {
		{
			ID:          7908,
			Slug:        "guarded-ether",
			Name:        "Guarded Ether",
			BinanceUSDT: "GETHUSDT",
		},
	},
	"DOGET": {
		{
			ID:          3919,
			Slug:        "doge-token",
			Name:        "Doge Token",
			BinanceUSDT: "DOGETUSDT",
		},
	},
	"VISION": {
		{
			ID:          8045,
			Slug:        "apy-vision",
			Name:        "APY Vision",
			BinanceUSDT: "VISIONUSDT",
		},
	},
	"PLF": {
		{
			ID:          5084,
			Slug:        "playfuel",
			Name:        "PlayFuel",
			BinanceUSDT: "PLFUSDT",
		},
	},
	"UNII": {
		{
			ID:          7207,
			Slug:        "unii-finance",
			Name:        "UNII Finance",
			BinanceUSDT: "UNIIUSDT",
		},
	},
	"NODE": {
		{
			ID:   4674,
			Slug: "whole-network",
			Name: "Whole Network",
		},
		{
			ID:   10990,
			Slug: "dappnode",
			Name: "DAppNode",
		},
		{
			ID:   4773,
			Slug: "pocketnode",
			Name: "PocketNode",
		},
	},
	"NIM": {
		{
			ID:          2916,
			Slug:        "nimiq",
			Name:        "Nimiq",
			BinanceUSDT: "NIMUSDT",
		},
	},
	"NBNG": {
		{
			ID:          10716,
			Slug:        "nobunaga-token-nbng",
			Name:        "Nobunaga Token, NBNG",
			BinanceUSDT: "NBNGUSDT",
		},
	},
	"REAL": {
		{
			ID:          2030,
			Slug:        "real",
			Name:        "REAL",
			BinanceUSDT: "REALUSDT",
		},
	},
	"DICE": {
		{
			ID:   6219,
			Slug: "dice",
			Name: "Dice",
		},
		{
			ID:   1677,
			Slug: "etheroll",
			Name: "Etheroll",
		},
		{
			ID:   6501,
			Slug: "tronbetdice",
			Name: "TRONbetDice",
		},
	},
	"DAPPT": {
		{
			ID:          4176,
			Slug:        "dapp-token",
			Name:        "Dapp Token",
			BinanceUSDT: "DAPPTUSDT",
		},
	},
	"DLT": {
		{
			ID:          1949,
			Slug:        "agrello-delta",
			Name:        "Agrello",
			BinanceUSDT: "DLTUSDT",
		},
	},
	"CC": {
		{
			ID:          9564,
			Slug:        "cryptocart",
			Name:        "CryptoCart",
			BinanceUSDT: "CCUSDT",
		},
	},
	"ZIK": {
		{
			ID:          6249,
			Slug:        "ziktalk",
			Name:        "Ziktalk",
			BinanceUSDT: "ZIKUSDT",
		},
	},
	"PSL": {
		{
			ID:          8738,
			Slug:        "pastel",
			Name:        "Pastel",
			BinanceUSDT: "PSLUSDT",
		},
	},
	"DGX": {
		{
			ID:          2739,
			Slug:        "digix-gold-token",
			Name:        "Digix Gold Token",
			BinanceUSDT: "DGXUSDT",
		},
	},
	"RPT": {
		{
			ID:          8413,
			Slug:        "rug-proof",
			Name:        "Rug Proof",
			BinanceUSDT: "RPTUSDT",
		},
	},
	"ORK": {
		{
			ID:          9150,
			Slug:        "orakuru",
			Name:        "Orakuru",
			BinanceUSDT: "ORKUSDT",
		},
	},
	"WCO": {
		{
			ID:          3669,
			Slug:        "winco",
			Name:        "Winco",
			BinanceUSDT: "WCOUSDT",
		},
	},
	"UAT": {
		{
			ID:          4262,
			Slug:        "ultralpha",
			Name:        "UltrAlpha",
			BinanceUSDT: "UATUSDT",
		},
	},
	"MILF": {
		{
			ID:   9860,
			Slug: "moms-id-like-to-fund",
			Name: "Moms I'd Like to Fund",
		},
		{
			ID:   9884,
			Slug: "milf-token",
			Name: "MILF Token",
		},
	},
	"GAX": {
		{
			ID:          10004,
			Slug:        "getart",
			Name:        "GETART",
			BinanceUSDT: "GAXUSDT",
		},
	},
	"TGDY": {
		{
			ID:          10320,
			Slug:        "tegridy",
			Name:        "Tegridy",
			BinanceUSDT: "TGDYUSDT",
		},
	},
	"PICA": {
		{
			ID:          7627,
			Slug:        "picaartmoney",
			Name:        "PicaArtMoney",
			BinanceUSDT: "PICAUSDT",
		},
	},
	"BTDX": {
		{
			ID:          1381,
			Slug:        "bitcloud",
			Name:        "Bitcloud",
			BinanceUSDT: "BTDXUSDT",
		},
	},
	"MOMA": {
		{
			ID:          9440,
			Slug:        "mochi-market",
			Name:        "Mochi Market",
			BinanceUSDT: "MOMAUSDT",
		},
	},
	"TOAD": {
		{
			ID:          9983,
			Slug:        "toad-network",
			Name:        "toad.network",
			BinanceUSDT: "TOADUSDT",
		},
	},
	"HYPER": {
		{
			ID:          9309,
			Slug:        "hyperchain",
			Name:        "HyperChain",
			BinanceUSDT: "HYPERUSDT",
		},
	},
	"DVS": {
		{
			ID:   4523,
			Slug: "davies",
			Name: "Davies",
		},
		{
			ID:   6230,
			Slug: "diamond-voucher",
			Name: "Diamond Voucher",
		},
	},
	"RFCTR": {
		{
			ID:          8039,
			Slug:        "reflector-finance",
			Name:        "Reflector.Finance",
			BinanceUSDT: "RFCTRUSDT",
		},
	},
	"XLAB": {
		{
			ID:          4709,
			Slug:        "xceltoken-plus",
			Name:        "XcelToken Plus",
			BinanceUSDT: "XLABUSDT",
		},
	},
	"XIN": {
		{
			ID:   2349,
			Slug: "mixin",
			Name: "Mixin",
		},
		{
			ID:   2013,
			Slug: "infinity-economics",
			Name: "Infinity Economics",
		},
	},
	"MP3": {
		{
			ID:          8412,
			Slug:        "mp3",
			Name:        "MP3",
			BinanceUSDT: "MP3USDT",
		},
	},
	"PIKA": {
		{
			ID:          8879,
			Slug:        "pika",
			Name:        "Pika",
			BinanceUSDT: "PIKAUSDT",
		},
	},
	"HOGE": {
		{
			ID:          8438,
			Slug:        "hoge-finance",
			Name:        "Hoge Finance",
			BinanceUSDT: "HOGEUSDT",
		},
	},
	"AXIOM": {
		{
			ID:          1020,
			Slug:        "axiom",
			Name:        "Axiom",
			BinanceUSDT: "AXIOMUSDT",
		},
	},
	"PLG": {
		{
			ID:          4196,
			Slug:        "pledge-coin",
			Name:        "Pledge Coin",
			BinanceUSDT: "PLGUSDT",
		},
	},
	"TKY": {
		{
			ID:          2507,
			Slug:        "thekey",
			Name:        "THEKEY",
			BinanceUSDT: "TKYUSDT",
		},
	},
	"L2P": {
		{
			ID:          5993,
			Slug:        "lung-protocol",
			Name:        "Lung Protocol",
			BinanceUSDT: "L2PUSDT",
		},
	},
	"KIWI": {
		{
			ID:   8640,
			Slug: "kiwi-finance",
			Name: "Kiwi Finance",
		},
		{
			ID:   5969,
			Slug: "kiwi-token",
			Name: "KIWI TOKEN",
		},
	},
	"COL": {
		{
			ID:          5890,
			Slug:        "unit-protocol",
			Name:        "Unit Protocol",
			BinanceUSDT: "COLUSDT",
		},
	},
	"BYTE": {
		{
			ID:          6126,
			Slug:        "btc-network-demand-set-ii",
			Name:        "BTC Network Demand Set II",
			BinanceUSDT: "BYTEUSDT",
		},
	},
	"SPK": {
		{
			ID:          2448,
			Slug:        "sparkspay",
			Name:        "SparksPay",
			BinanceUSDT: "SPKUSDT",
		},
	},
	"ASKO": {
		{
			ID:          5833,
			Slug:        "askobar-network",
			Name:        "ASKO",
			BinanceUSDT: "ASKOUSDT",
		},
	},
	"VDX": {
		{
			ID:          3962,
			Slug:        "vodi-x",
			Name:        "Vodi X",
			BinanceUSDT: "VDXUSDT",
		},
	},
	"MOBI": {
		{
			ID:          2429,
			Slug:        "mobius",
			Name:        "Mobius",
			BinanceUSDT: "MOBIUSDT",
		},
	},
	"ROCKS": {
		{
			ID:   8105,
			Slug: "rocki",
			Name: "ROCKI",
		},
		{
			ID:   7571,
			Slug: "social-rocket",
			Name: "Social Rocket",
		},
	},
	"XMC": {
		{
			ID:          2655,
			Slug:        "monero-classic",
			Name:        "Monero Classic",
			BinanceUSDT: "XMCUSDT",
		},
	},
	"REEC": {
		{
			ID:          3904,
			Slug:        "electronic-energy-coin",
			Name:        "Renewable Electronic Energy Coin",
			BinanceUSDT: "REECUSDT",
		},
	},
	"LDFI": {
		{
			ID:          9341,
			Slug:        "lendefi",
			Name:        "Lendefi",
			BinanceUSDT: "LDFIUSDT",
		},
	},
	"VTC": {
		{
			ID:          99,
			Slug:        "vertcoin",
			Name:        "Vertcoin",
			BinanceUSDT: "VTCUSDT",
		},
	},
	"SI": {
		{
			ID:          8701,
			Slug:        "siren",
			Name:        "Siren",
			BinanceUSDT: "SIUSDT",
		},
	},
	"SALT": {
		{
			ID:   1996,
			Slug: "salt",
			Name: "SALT",
		},
		{
			ID:   8625,
			Slug: "saltswap-finance",
			Name: "SaltSwap Finance",
		},
	},
	"SAFEMOONCASH": {
		{
			ID:          10081,
			Slug:        "safemooncash",
			Name:        "SafeMoonCash",
			BinanceUSDT: "SAFEMOONCASHUSDT",
		},
	},
	"BTCR": {
		{
			ID:          5404,
			Slug:        "bitcurate",
			Name:        "Bitcurate",
			BinanceUSDT: "BTCRUSDT",
		},
	},
	"NBU": {
		{
			ID:          8939,
			Slug:        "nimbus",
			Name:        "Nimbus",
			BinanceUSDT: "NBUUSDT",
		},
	},
	"RGP": {
		{
			ID:          9225,
			Slug:        "rigel-protocol",
			Name:        "Rigel Protocol",
			BinanceUSDT: "RGPUSDT",
		},
	},
	"NOAHP": {
		{
			ID:          2599,
			Slug:        "noah-coin",
			Name:        "Noah Coin",
			BinanceUSDT: "NOAHPUSDT",
		},
	},
	"C20": {
		{
			ID:          2444,
			Slug:        "c20",
			Name:        "CRYPTO20",
			BinanceUSDT: "C20USDT",
		},
	},
	"SCONEX": {
		{
			ID:          9827,
			Slug:        "sportcash-one",
			Name:        "Sportcash One",
			BinanceUSDT: "SCONEXUSDT",
		},
	},
	"DAT": {
		{
			ID:          2283,
			Slug:        "datum",
			Name:        "Datum",
			BinanceUSDT: "DATUSDT",
		},
	},
	"UNFT": {
		{
			ID:          10008,
			Slug:        "ultra-nft",
			Name:        "Ultra NFT",
			BinanceUSDT: "UNFTUSDT",
		},
	},
	"GNY": {
		{
			ID:          3936,
			Slug:        "gny",
			Name:        "GNY",
			BinanceUSDT: "GNYUSDT",
		},
	},
	"DOUGH": {
		{
			ID:          7284,
			Slug:        "piedao-dough-v2",
			Name:        "PieDAO DOUGH v2",
			BinanceUSDT: "DOUGHUSDT",
		},
	},
	"ECOIN": {
		{
			ID:   10159,
			Slug: "e-coin-finance",
			Name: "E-coin Finance",
		},
		{
			ID:   6111,
			Slug: "ecoin-2",
			Name: "Ecoin",
		},
	},
	"FND": {
		{
			ID:          9901,
			Slug:        "fundum-capital",
			Name:        "Fundum Capital",
			BinanceUSDT: "FNDUSDT",
		},
	},
	"EA": {
		{
			ID:          6427,
			Slug:        "ea-token",
			Name:        "EA Token",
			BinanceUSDT: "EAUSDT",
		},
	},
	"DPY": {
		{
			ID:          2162,
			Slug:        "delphy",
			Name:        "Delphy",
			BinanceUSDT: "DPYUSDT",
		},
	},
	"YFII": {
		{
			ID:          5957,
			Slug:        "yearn-finance-ii",
			Name:        "DFI.Money",
			BinanceUSDT: "YFIIUSDT",
		},
	},
	"IRIS": {
		{
			ID:          3874,
			Slug:        "irisnet",
			Name:        "IRISnet",
			BinanceUSDT: "IRISUSDT",
		},
	},
	"ASTRA": {
		{
			ID:          10621,
			Slug:        "astra-coin",
			Name:        "Astra Coin",
			BinanceUSDT: "ASTRAUSDT",
		},
	},
	"DIGS": {
		{
			ID:          10980,
			Slug:        "digies-coin",
			Name:        "Digies Coin",
			BinanceUSDT: "DIGSUSDT",
		},
	},
	"mVIXY": {
		{
			ID:          8028,
			Slug:        "mirrored-proshares-vix-short-term-futures-etf",
			Name:        "Mirrored ProShares VIX",
			BinanceUSDT: "mVIXYUSDT",
		},
	},
	"AIDI": {
		{
			ID:          10692,
			Slug:        "aidi-finance",
			Name:        "Aidi Finance",
			BinanceUSDT: "AIDIUSDT",
		},
	},
	"PRUDE": {
		{
			ID:          10503,
			Slug:        "prude-token",
			Name:        "Prude Token",
			BinanceUSDT: "PRUDEUSDT",
		},
	},
	"NOIA": {
		{
			ID:          4191,
			Slug:        "syntropy",
			Name:        "Syntropy",
			BinanceUSDT: "NOIAUSDT",
		},
	},
	"DEX": {
		{
			ID:          3515,
			Slug:        "dex",
			Name:        "DEX",
			BinanceUSDT: "DEXUSDT",
		},
	},
	"AAVEUP": {
		{
			ID:          7774,
			Slug:        "aave-up",
			Name:        "AAVEUP",
			BinanceUSDT: "AAVEUPUSDT",
		},
	},
	"KIT": {
		{
			ID:          7739,
			Slug:        "dexkit",
			Name:        "DexKit",
			BinanceUSDT: "KITUSDT",
		},
	},
	"HLIX": {
		{
			ID:          5447,
			Slug:        "helix",
			Name:        "Helix",
			BinanceUSDT: "HLIXUSDT",
		},
	},
	"JNB": {
		{
			ID:          3759,
			Slug:        "jinbi-token",
			Name:        "Jinbi Token",
			BinanceUSDT: "JNBUSDT",
		},
	},
	"HAZE": {
		{
			ID:          8810,
			Slug:        "haze-finance",
			Name:        "Haze Finance",
			BinanceUSDT: "HAZEUSDT",
		},
	},
	"DOGETF": {
		{
			ID:          9729,
			Slug:        "doge-father-token",
			Name:        "Doge Father Token",
			BinanceUSDT: "DOGETFUSDT",
		},
	},
	"REI": {
		{
			ID:          10844,
			Slug:        "zerogoki",
			Name:        "Zerogoki",
			BinanceUSDT: "REIUSDT",
		},
	},
	"ETP": {
		{
			ID:          1703,
			Slug:        "metaverse",
			Name:        "Metaverse ETP",
			BinanceUSDT: "ETPUSDT",
		},
	},
	"BART": {
		{
			ID:          6592,
			Slug:        "bartertrade",
			Name:        "BarterTrade",
			BinanceUSDT: "BARTUSDT",
		},
	},
	"MNTP": {
		{
			ID:          2513,
			Slug:        "goldmint",
			Name:        "GoldMint",
			BinanceUSDT: "MNTPUSDT",
		},
	},
	"SWT": {
		{
			ID:          1562,
			Slug:        "swarm-city",
			Name:        "Swarm City",
			BinanceUSDT: "SWTUSDT",
		},
	},
	"DSYS": {
		{
			ID:          6426,
			Slug:        "dsys",
			Name:        "DSYS",
			BinanceUSDT: "DSYSUSDT",
		},
	},
	"LKT": {
		{
			ID:          10775,
			Slug:        "locklet",
			Name:        "Locklet",
			BinanceUSDT: "LKTUSDT",
		},
	},
	"TELOS": {
		{
			ID:          3482,
			Slug:        "teloscoin",
			Name:        "Teloscoin",
			BinanceUSDT: "TELOSUSDT",
		},
	},
	"CJ": {
		{
			ID:          1306,
			Slug:        "cryptojacks",
			Name:        "Cryptojacks",
			BinanceUSDT: "CJUSDT",
		},
	},
	"MVP": {
		{
			ID:          2869,
			Slug:        "merculet",
			Name:        "Merculet",
			BinanceUSDT: "MVPUSDT",
		},
	},
	"BU": {
		{
			ID:          3295,
			Slug:        "bumo",
			Name:        "BUMO",
			BinanceUSDT: "BUUSDT",
		},
	},
	"TLOS": {
		{
			ID:          4660,
			Slug:        "telos",
			Name:        "Telos",
			BinanceUSDT: "TLOSUSDT",
		},
	},
	"QLC": {
		{
			ID:          2321,
			Slug:        "qlink",
			Name:        "QLC Chain",
			BinanceUSDT: "QLCUSDT",
		},
	},
	"SLR": {
		{
			ID:          233,
			Slug:        "solarcoin",
			Name:        "SolarCoin",
			BinanceUSDT: "SLRUSDT",
		},
	},
	"GTN": {
		{
			ID:          3914,
			Slug:        "glitzkoin",
			Name:        "GlitzKoin",
			BinanceUSDT: "GTNUSDT",
		},
	},
	"EOC": {
		{
			ID:          8048,
			Slug:        "everyonescrypto",
			Name:        "Everyonescrypto",
			BinanceUSDT: "EOCUSDT",
		},
	},
	"BOOZE": {
		{
			ID:          9924,
			Slug:        "boozemoon",
			Name:        "BoozeMoon",
			BinanceUSDT: "BOOZEUSDT",
		},
	},
	"PIRL": {
		{
			ID:          2105,
			Slug:        "pirl",
			Name:        "Pirl",
			BinanceUSDT: "PIRLUSDT",
		},
	},
	"4STC": {
		{
			ID:          10488,
			Slug:        "4-stock",
			Name:        "4-Stock",
			BinanceUSDT: "4STCUSDT",
		},
	},
	"RTH": {
		{
			ID:          3279,
			Slug:        "rotharium",
			Name:        "Rotharium",
			BinanceUSDT: "RTHUSDT",
		},
	},
	"GEM": {
		{
			ID:   2537,
			Slug: "gems-protocol",
			Name: "Gems ",
		},
		{
			ID:   10540,
			Slug: "gem-protocol",
			Name: "GEM PROTOCOL",
		},
	},
	"RGT": {
		{
			ID:          7486,
			Slug:        "rari-governance-token",
			Name:        "Rari Governance Token",
			BinanceUSDT: "RGTUSDT",
		},
	},
	"PYLNT": {
		{
			ID:          2330,
			Slug:        "pylon-network",
			Name:        "Pylon Network",
			BinanceUSDT: "PYLNTUSDT",
		},
	},
	"STKR": {
		{
			ID:          8934,
			Slug:        "stakerdao",
			Name:        "StakerDAO",
			BinanceUSDT: "STKRUSDT",
		},
	},
	"GMB": {
		{
			ID:   3819,
			Slug: "gamb",
			Name: "GAMB",
		},
		{
			ID:   3757,
			Slug: "gmb",
			Name: "GMB",
		},
	},
	"GIO": {
		{
			ID:          3118,
			Slug:        "graviocoin",
			Name:        "Graviocoin",
			BinanceUSDT: "GIOUSDT",
		},
	},
	"DION": {
		{
			ID:          6211,
			Slug:        "dionpay",
			Name:        "Dionpay",
			BinanceUSDT: "DIONUSDT",
		},
	},
	"LYR": {
		{
			ID:          7762,
			Slug:        "lyra",
			Name:        "Lyra",
			BinanceUSDT: "LYRUSDT",
		},
	},
	"XPRT": {
		{
			ID:          7281,
			Slug:        "persistence",
			Name:        "Persistence",
			BinanceUSDT: "XPRTUSDT",
		},
	},
	"JUP": {
		{
			ID:          1503,
			Slug:        "jupiter",
			Name:        "Jupiter",
			BinanceUSDT: "JUPUSDT",
		},
	},
	"GST": {
		{
			ID:   5820,
			Slug: "gstcoin",
			Name: "Gstcoin",
		},
		{
			ID:   8817,
			Slug: "the-gemstone",
			Name: "The Gemstone",
		},
	},
	"PLR": {
		{
			ID:          1834,
			Slug:        "pillar",
			Name:        "Pillar",
			BinanceUSDT: "PLRUSDT",
		},
	},
	"DUCATO": {
		{
			ID:          7133,
			Slug:        "ducato-protocol-token",
			Name:        "Ducato Protocol Token",
			BinanceUSDT: "DUCATOUSDT",
		},
	},
	"HXRO": {
		{
			ID:          3748,
			Slug:        "hxro",
			Name:        "Hxro",
			BinanceUSDT: "HXROUSDT",
		},
	},
	"MOLK": {
		{
			ID:          3304,
			Slug:        "mobilinktoken",
			Name:        "MobilinkToken",
			BinanceUSDT: "MOLKUSDT",
		},
	},
	"DEFX": {
		{
			ID:          10145,
			Slug:        "definity",
			Name:        "DeFinity",
			BinanceUSDT: "DEFXUSDT",
		},
	},
	"LUNES": {
		{
			ID:          3786,
			Slug:        "lunes",
			Name:        "Lunes",
			BinanceUSDT: "LUNESUSDT",
		},
	},
	"BTCV": {
		{
			ID:   5175,
			Slug: "bitcoin-vault",
			Name: "Bitcoin Vault",
		},
		{
			ID:   4787,
			Slug: "bitcoinv",
			Name: "BitcoinV",
		},
	},
	"UBTC": {
		{
			ID:          2293,
			Slug:        "united-bitcoin",
			Name:        "United Bitcoin",
			BinanceUSDT: "UBTCUSDT",
		},
	},
	"LIBERTAS": {
		{
			ID:          5997,
			Slug:        "libertas-token",
			Name:        "Libertas Token",
			BinanceUSDT: "LIBERTASUSDT",
		},
	},
	"SNBL": {
		{
			ID:          9036,
			Slug:        "nebulaprotocol",
			Name:        "Nebulaprotocol",
			BinanceUSDT: "SNBLUSDT",
		},
	},
	"HORUS": {
		{
			ID:          2993,
			Slug:        "horuspay",
			Name:        "HorusPay",
			BinanceUSDT: "HORUSUSDT",
		},
	},
	"SAP": {
		{
			ID:          7584,
			Slug:        "swapall",
			Name:        "SwapAll",
			BinanceUSDT: "SAPUSDT",
		},
	},
	"$NOOB": {
		{
			ID:          7646,
			Slug:        "noob-finance",
			Name:        "noob.finance",
			BinanceUSDT: "$NOOBUSDT",
		},
	},
	"ACED": {
		{
			ID:          2978,
			Slug:        "aced",
			Name:        "AceD",
			BinanceUSDT: "ACEDUSDT",
		},
	},
	"WOLFY": {
		{
			ID:          10963,
			Slug:        "wolfystreetbets",
			Name:        "Wolfystreetbets",
			BinanceUSDT: "WOLFYUSDT",
		},
	},
	"KICKS": {
		{
			ID:          4273,
			Slug:        "sessia",
			Name:        "Sessia",
			BinanceUSDT: "KICKSUSDT",
		},
	},
	"WIX": {
		{
			ID:          3404,
			Slug:        "wixlar",
			Name:        "Wixlar",
			BinanceUSDT: "WIXUSDT",
		},
	},
	"VEN": {
		{
			ID:          8908,
			Slug:        "impulseven",
			Name:        "ImpulseVen",
			BinanceUSDT: "VENUSDT",
		},
	},
	"BDCC": {
		{
			ID:          4427,
			Slug:        "bdcc-bitica-coin",
			Name:        "BDCC Bitica COIN",
			BinanceUSDT: "BDCCUSDT",
		},
	},
	"UENC": {
		{
			ID:          6525,
			Slug:        "universalenergychain",
			Name:        "UniversalEnergyChain",
			BinanceUSDT: "UENCUSDT",
		},
	},
	"EYE": {
		{
			ID:          7414,
			Slug:        "beholder",
			Name:        "Behodler",
			BinanceUSDT: "EYEUSDT",
		},
	},
	"AIR": {
		{
			ID:          10715,
			Slug:        "air",
			Name:        "AirCoin",
			BinanceUSDT: "AIRUSDT",
		},
	},
	"DRK": {
		{
			ID:          10234,
			Slug:        "draken",
			Name:        "Draken",
			BinanceUSDT: "DRKUSDT",
		},
	},
	"CROAT": {
		{
			ID:          3087,
			Slug:        "croat",
			Name:        "CROAT",
			BinanceUSDT: "CROATUSDT",
		},
	},
	"YFIKING": {
		{
			ID:          6964,
			Slug:        "yfiking-finance",
			Name:        "YFIKING,FINANCE",
			BinanceUSDT: "YFIKINGUSDT",
		},
	},
	"TRUST": {
		{
			ID:   5972,
			Slug: "trustdao",
			Name: "TrustDAO",
		},
		{
			ID:   9206,
			Slug: "trustworks",
			Name: "Trustworks",
		},
	},
	"SCAP": {
		{
			ID:          5002,
			Slug:        "safecapital",
			Name:        "SafeCapital",
			BinanceUSDT: "SCAPUSDT",
		},
	},
	"CELR": {
		{
			ID:          3814,
			Slug:        "celer-network",
			Name:        "Celer Network",
			BinanceUSDT: "CELRUSDT",
		},
	},
	"CTL": {
		{
			ID:          1273,
			Slug:        "citadel",
			Name:        "Citadel",
			BinanceUSDT: "CTLUSDT",
		},
	},
	"COLLG": {
		{
			ID:          10860,
			Slug:        "collateral-pay-governance",
			Name:        "Collateral Pay Governance",
			BinanceUSDT: "COLLGUSDT",
		},
	},
	"ZILD": {
		{
			ID:          10381,
			Slug:        "zild-finance",
			Name:        "Zild Finance",
			BinanceUSDT: "ZILDUSDT",
		},
	},
	"TFBX": {
		{
			ID:          4144,
			Slug:        "truefeedback",
			Name:        "TrueFeedBack",
			BinanceUSDT: "TFBXUSDT",
		},
	},
	"LKN": {
		{
			ID:          6323,
			Slug:        "linkcoin-token",
			Name:        "LinkCoin Token",
			BinanceUSDT: "LKNUSDT",
		},
	},
	"WELL": {
		{
			ID:          7883,
			Slug:        "well-token",
			Name:        "WELL",
			BinanceUSDT: "WELLUSDT",
		},
	},
	"MSTR": {
		{
			ID:          7898,
			Slug:        "microstrategy-tokenized-stock-ftx",
			Name:        "MicroStrategy tokenized stock FTX",
			BinanceUSDT: "MSTRUSDT",
		},
	},
	"METH": {
		{
			ID:          8907,
			Slug:        "farming-bad",
			Name:        "Farming Bad",
			BinanceUSDT: "METHUSDT",
		},
	},
	"ZBTC": {
		{
			ID:          10280,
			Slug:        "zetta-bitcoin-hashrate-token",
			Name:        "Zetta Bitcoin Hashrate Token",
			BinanceUSDT: "ZBTCUSDT",
		},
	},
	"TBTC": {
		{
			ID:          5776,
			Slug:        "tbtc",
			Name:        "tBTC",
			BinanceUSDT: "TBTCUSDT",
		},
	},
	"BYTS": {
		{
			ID:          6393,
			Slug:        "bytus",
			Name:        "Bytus",
			BinanceUSDT: "BYTSUSDT",
		},
	},
	"TRXBEAR": {
		{
			ID:          5379,
			Slug:        "3x-short-trx-token",
			Name:        "3X Short TRX Token",
			BinanceUSDT: "TRXBEARUSDT",
		},
	},
	"KPAD": {
		{
			ID:          8897,
			Slug:        "kickpad",
			Name:        "KickPad",
			BinanceUSDT: "KPADUSDT",
		},
	},
	"WHEN": {
		{
			ID:          3849,
			Slug:        "when-token",
			Name:        "WHEN Token",
			BinanceUSDT: "WHENUSDT",
		},
	},
	"$MAD": {
		{
			ID:          10215,
			Slug:        "make-a-difference-token",
			Name:        "Make A Difference Token",
			BinanceUSDT: "$MADUSDT",
		},
	},
	"MOON": {
		{
			ID:   7017,
			Slug: "moonswap",
			Name: "MoonSwap",
		},
		{
			ID:   10229,
			Slug: "polywolf",
			Name: "Polywolf",
		},
	},
	"NAUT": {
		{
			ID:          8770,
			Slug:        "astronaut",
			Name:        "Astronaut",
			BinanceUSDT: "NAUTUSDT",
		},
	},
	"DXT": {
		{
			ID:          2510,
			Slug:        "datawallet",
			Name:        "Datawallet",
			BinanceUSDT: "DXTUSDT",
		},
	},
	"AMZ": {
		{
			ID:          6814,
			Slug:        "amazonascoin",
			Name:        "AmazonasCoin",
			BinanceUSDT: "AMZUSDT",
		},
	},
	"XXA": {
		{
			ID:          5474,
			Slug:        "ixinium",
			Name:        "Ixinium",
			BinanceUSDT: "XXAUSDT",
		},
	},
	"LCX": {
		{
			ID:          4950,
			Slug:        "lcx",
			Name:        "LCX",
			BinanceUSDT: "LCXUSDT",
		},
	},
	"LAT": {
		{
			ID:          9720,
			Slug:        "platon",
			Name:        "PlatON",
			BinanceUSDT: "LATUSDT",
		},
	},
	"1MIL": {
		{
			ID:          9344,
			Slug:        "1millionnfts",
			Name:        "1MillionNFTs",
			BinanceUSDT: "1MILUSDT",
		},
	},
	"UMASK": {
		{
			ID:          9477,
			Slug:        "unicly-hashmasks-collection",
			Name:        "Unicly Hashmasks Collection",
			BinanceUSDT: "UMASKUSDT",
		},
	},
	"SHVR": {
		{
			ID:          3645,
			Slug:        "shivers",
			Name:        "Shivers",
			BinanceUSDT: "SHVRUSDT",
		},
	},
	"RNT": {
		{
			ID:          2400,
			Slug:        "oneroot-network",
			Name:        "OneRoot Network",
			BinanceUSDT: "RNTUSDT",
		},
	},
	"VALOR": {
		{
			ID:          3875,
			Slug:        "valor-token",
			Name:        "Valor Token",
			BinanceUSDT: "VALORUSDT",
		},
	},
	"VOLT": {
		{
			ID:          1657,
			Slug:        "bitvolt",
			Name:        "Bitvolt",
			BinanceUSDT: "VOLTUSDT",
		},
	},
	"BNY": {
		{
			ID:          4775,
			Slug:        "bancacy",
			Name:        "Bancacy",
			BinanceUSDT: "BNYUSDT",
		},
	},
	"NFD": {
		{
			ID:          9496,
			Slug:        "nifdo-protocol",
			Name:        "NIFDO Protocol",
			BinanceUSDT: "NFDUSDT",
		},
	},
	"ETCBULL": {
		{
			ID:          6083,
			Slug:        "3x-long-ethereum-classic-token",
			Name:        "3X Long Ethereum Classic Token",
			BinanceUSDT: "ETCBULLUSDT",
		},
	},
	"NEWOS": {
		{
			ID:          3110,
			Slug:        "newstoken",
			Name:        "NewsToken",
			BinanceUSDT: "NEWOSUSDT",
		},
	},
	"SUSHIBA": {
		{
			ID:          10308,
			Slug:        "sushiba",
			Name:        "Sushiba",
			BinanceUSDT: "SUSHIBAUSDT",
		},
	},
	"DINO": {
		{
			ID:   9621,
			Slug: "dinoexchange",
			Name: "DinoExchange",
		},
		{
			ID:   10777,
			Slug: "dinoswap",
			Name: "DinoSwap",
		},
	},
	"ALU": {
		{
			ID:          9637,
			Slug:        "altura",
			Name:        "Altura",
			BinanceUSDT: "ALUUSDT",
		},
	},
	"SLM": {
		{
			ID:          8268,
			Slug:        "solomon-defi",
			Name:        "Solomon Defi",
			BinanceUSDT: "SLMUSDT",
		},
	},
	"GYEN": {
		{
			ID:          8771,
			Slug:        "gyen",
			Name:        "GYEN",
			BinanceUSDT: "GYENUSDT",
		},
	},
	"DAIN": {
		{
			ID:          9304,
			Slug:        "dain",
			Name:        "DAIN",
			BinanceUSDT: "DAINUSDT",
		},
	},
	"N0001": {
		{
			ID:          8084,
			Slug:        "nhbtc",
			Name:        "nHBTC",
			BinanceUSDT: "N0001USDT",
		},
	},
	"ARGON": {
		{
			ID:          8421,
			Slug:        "argon",
			Name:        "Argon",
			BinanceUSDT: "ARGONUSDT",
		},
	},
	"HLP": {
		{
			ID:          7263,
			Slug:        "help-coin",
			Name:        "HLP Token",
			BinanceUSDT: "HLPUSDT",
		},
	},
	"TRDG": {
		{
			ID:          10785,
			Slug:        "tardigrades-finance-eth",
			Name:        "Tardigrades.Finance (ETH)",
			BinanceUSDT: "TRDGUSDT",
		},
	},
	"SWIFT": {
		{
			ID:          3933,
			Slug:        "swiftcash",
			Name:        "SwiftCash",
			BinanceUSDT: "SWIFTUSDT",
		},
	},
	"WAR": {
		{
			ID:   9757,
			Slug: "westarter",
			Name: "WeStarter",
		},
		{
			ID:   7251,
			Slug: "yieldwars-com",
			Name: "YieldWars",
		},
		{
			ID:   8927,
			Slug: "nft-wars",
			Name: "NFT Wars",
		},
		{
			ID:   8682,
			Slug: "warrior-token-spartan-casino",
			Name: "Warrior Token",
		},
	},
	"NCP": {
		{
			ID:          2984,
			Slug:        "newton-coin-project",
			Name:        "Newton Coin Project",
			BinanceUSDT: "NCPUSDT",
		},
	},
	"UMB": {
		{
			ID:          8385,
			Slug:        "umbrella-network",
			Name:        "Umbrella Network",
			BinanceUSDT: "UMBUSDT",
		},
	},
	"BTG": {
		{
			ID:          2083,
			Slug:        "bitcoin-gold",
			Name:        "Bitcoin Gold",
			BinanceUSDT: "BTGUSDT",
		},
	},
	"NEAR": {
		{
			ID:          6535,
			Slug:        "near-protocol",
			Name:        "NEAR Protocol",
			BinanceUSDT: "NEARUSDT",
		},
	},
	"CLIT$": {
		{
			ID:          10571,
			Slug:        "clit-token-protocol",
			Name:        "CLIT TOKEN PROTOCOL",
			BinanceUSDT: "CLIT$USDT",
		},
	},
	"METRIC": {
		{
			ID:          7254,
			Slug:        "metric-exchange",
			Name:        "Metric Exchange",
			BinanceUSDT: "METRICUSDT",
		},
	},
	"BCHC": {
		{
			ID:          5803,
			Slug:        "bitcherry",
			Name:        "BitCherry",
			BinanceUSDT: "BCHCUSDT",
		},
	},
	"CTRT": {
		{
			ID:          3317,
			Slug:        "cryptrust",
			Name:        "Cryptrust",
			BinanceUSDT: "CTRTUSDT",
		},
	},
	"SSX": {
		{
			ID:          5612,
			Slug:        "somesing",
			Name:        "SOMESING",
			BinanceUSDT: "SSXUSDT",
		},
	},
	"ASY": {
		{
			ID:          5529,
			Slug:        "asyagro",
			Name:        "ASYAGRO",
			BinanceUSDT: "ASYUSDT",
		},
	},
	"ANT": {
		{
			ID:   1680,
			Slug: "aragon",
			Name: "Aragon",
		},
		{
			ID:   9380,
			Slug: "antcoin",
			Name: "ANTcoin",
		},
	},
	"TAD": {
		{
			ID:          7559,
			Slug:        "tadpole-finance",
			Name:        "Tadpole Finance",
			BinanceUSDT: "TADUSDT",
		},
	},
	"YFIE": {
		{
			ID:          6708,
			Slug:        "yfiexchange-finance",
			Name:        "YFIEXCHANGE.FINANCE",
			BinanceUSDT: "YFIEUSDT",
		},
	},
	"TTX": {
		{
			ID:          7392,
			Slug:        "talent-token",
			Name:        "Talent Token",
			BinanceUSDT: "TTXUSDT",
		},
	},
	"WAYF": {
		{
			ID:          7700,
			Slug:        "way-f-coin",
			Name:        "WAY-F coin",
			BinanceUSDT: "WAYFUSDT",
		},
	},
	"PYPL": {
		{
			ID:          7901,
			Slug:        "paypal-tokenized-stock-ftx",
			Name:        "PayPal tokenized stock FTX",
			BinanceUSDT: "PYPLUSDT",
		},
	},
	"FSW": {
		{
			ID:          6743,
			Slug:        "fsw-token",
			Name:        "Falconswap",
			BinanceUSDT: "FSWUSDT",
		},
	},
	"RENBTC": {
		{
			ID:          5777,
			Slug:        "renbtc",
			Name:        "renBTC",
			BinanceUSDT: "RENBTCUSDT",
		},
	},
	"SWAPP": {
		{
			ID:          10783,
			Slug:        "swapp-protocol",
			Name:        "SWAPP Protocol",
			BinanceUSDT: "SWAPPUSDT",
		},
	},
	"DONK": {
		{
			ID:          9340,
			Slug:        "donkey",
			Name:        "Donkey",
			BinanceUSDT: "DONKUSDT",
		},
	},
	"42": {
		{
			ID:          93,
			Slug:        "42-coin",
			Name:        "42-coin",
			BinanceUSDT: "42USDT",
		},
	},
	"PRV": {
		{
			ID:          9933,
			Slug:        "privacyswap",
			Name:        "PrivacySwap",
			BinanceUSDT: "PRVUSDT",
		},
	},
	"MORK": {
		{
			ID:          6692,
			Slug:        "mork",
			Name:        "MORK",
			BinanceUSDT: "MORKUSDT",
		},
	},
	"YELD": {
		{
			ID:          6974,
			Slug:        "yeld-finance",
			Name:        "Yeld Finance",
			BinanceUSDT: "YELDUSDT",
		},
	},
	"NILU": {
		{
			ID:          6032,
			Slug:        "nilu",
			Name:        "Nilu",
			BinanceUSDT: "NILUUSDT",
		},
	},
	"ETHPA": {
		{
			ID:          6141,
			Slug:        "eth-price-action-candlestick-set",
			Name:        "ETH Price Action Candlestick Set",
			BinanceUSDT: "ETHPAUSDT",
		},
	},
	"BURNX20": {
		{
			ID:          10072,
			Slug:        "burnx20",
			Name:        "BurnX 2.0",
			BinanceUSDT: "BURNX20USDT",
		},
	},
	"EDC": {
		{
			ID:   7521,
			Slug: "edc-blockchain",
			Name: "EDC Blockchain",
		},
		{
			ID:   8212,
			Slug: "earn-defi",
			Name: "Earn Defi Coin",
		},
	},
	"SAFU": {
		{
			ID:          9375,
			Slug:        "ceezee-safu",
			Name:        "CEEZEE SAFU",
			BinanceUSDT: "SAFUUSDT",
		},
	},
	"HODL": {
		{
			ID:   9900,
			Slug: "hodl",
			Name: "HODL",
		},
		{
			ID:   8572,
			Slug: "hodlearn-net",
			Name: "hodlearn",
		},
	},
	"ASM": {
		{
			ID:          6069,
			Slug:        "assemble-protocol",
			Name:        "Assemble Protocol",
			BinanceUSDT: "ASMUSDT",
		},
	},
	"SNTVT": {
		{
			ID:          3917,
			Slug:        "sentivate",
			Name:        "Sentivate",
			BinanceUSDT: "SNTVTUSDT",
		},
	},
	"OAX": {
		{
			ID:          1853,
			Slug:        "oax",
			Name:        "OAX",
			BinanceUSDT: "OAXUSDT",
		},
	},
	"CALI": {
		{
			ID:          9595,
			Slug:        "calicoin",
			Name:        "CaliCoin",
			BinanceUSDT: "CALIUSDT",
		},
	},
	"DEFI++": {
		{
			ID:          7874,
			Slug:        "piedao-defi",
			Name:        "PieDAO DEFI++",
			BinanceUSDT: "DEFI++USDT",
		},
	},
	"BRDG": {
		{
			ID:          2968,
			Slug:        "bridge-protocol",
			Name:        "Bridge Protocol",
			BinanceUSDT: "BRDGUSDT",
		},
	},
	"DGCL": {
		{
			ID:          8148,
			Slug:        "digicol",
			Name:        "DigiCol",
			BinanceUSDT: "DGCLUSDT",
		},
	},
	"KONJ": {
		{
			ID:          5192,
			Slug:        "konjungate",
			Name:        "KONJUNGATE",
			BinanceUSDT: "KONJUSDT",
		},
	},
	"DOW": {
		{
			ID:          3223,
			Slug:        "dowcoin",
			Name:        "DOWCOIN",
			BinanceUSDT: "DOWUSDT",
		},
	},
	"IOV": {
		{
			ID:   7271,
			Slug: "starname",
			Name: "Starname",
		},
		{
			ID:   3026,
			Slug: "iov-blockchain",
			Name: "IOV BlockChain",
		},
	},
	"LA": {
		{
			ID:          2090,
			Slug:        "latoken",
			Name:        "LATOKEN",
			BinanceUSDT: "LAUSDT",
		},
	},
	"ISAL": {
		{
			ID:          8275,
			Slug:        "isalcoin",
			Name:        "ISALCOIN",
			BinanceUSDT: "ISALUSDT",
		},
	},
	"GSX": {
		{
			ID:          10093,
			Slug:        "gold-secured-currency",
			Name:        "Gold Secured Currency",
			BinanceUSDT: "GSXUSDT",
		},
	},
	"ASA": {
		{
			ID:          3477,
			Slug:        "asura-coin",
			Name:        "Asura Coin",
			BinanceUSDT: "ASAUSDT",
		},
	},
	"APPC": {
		{
			ID:          2344,
			Slug:        "appcoins",
			Name:        "AppCoins",
			BinanceUSDT: "APPCUSDT",
		},
	},
	"VRT": {
		{
			ID:          9691,
			Slug:        "venus-reward-token",
			Name:        "Venus Reward Token",
			BinanceUSDT: "VRTUSDT",
		},
	},
	"XPAT": {
		{
			ID:          3112,
			Slug:        "bitnation",
			Name:        "Bitnation",
			BinanceUSDT: "XPATUSDT",
		},
	},
	"DMOD": {
		{
			ID:          9713,
			Slug:        "demodyfi",
			Name:        "Demodyfi",
			BinanceUSDT: "DMODUSDT",
		},
	},
	"BCX": {
		{
			ID:          2281,
			Slug:        "bitcoinx",
			Name:        "BitcoinX",
			BinanceUSDT: "BCXUSDT",
		},
	},
	"HY": {
		{
			ID:          6261,
			Slug:        "hybrix",
			Name:        "hybrix",
			BinanceUSDT: "HYUSDT",
		},
	},
	"JT": {
		{
			ID:          6262,
			Slug:        "jubi-token",
			Name:        "Jubi Token",
			BinanceUSDT: "JTUSDT",
		},
	},
	"FORESTPLUS": {
		{
			ID:          4848,
			Slug:        "the-forbidden-forest",
			Name:        "The Forbidden Forest",
			BinanceUSDT: "FORESTPLUSUSDT",
		},
	},
	"ISR": {
		{
			ID:          3466,
			Slug:        "insureum",
			Name:        "Insureum",
			BinanceUSDT: "ISRUSDT",
		},
	},
	"XMM": {
		{
			ID:          7042,
			Slug:        "momentum",
			Name:        "Momentum",
			BinanceUSDT: "XMMUSDT",
		},
	},
	"YOT": {
		{
			ID:          8247,
			Slug:        "payyoda",
			Name:        "PayYoda",
			BinanceUSDT: "YOTUSDT",
		},
	},
	"POE": {
		{
			ID:          1937,
			Slug:        "poet",
			Name:        "Po.et",
			BinanceUSDT: "POEUSDT",
		},
	},
	"CAKECRYPT": {
		{
			ID:          10066,
			Slug:        "cakecrypt",
			Name:        "CAKECRYPT",
			BinanceUSDT: "CAKECRYPTUSDT",
		},
	},
	"KT": {
		{
			ID:          3691,
			Slug:        "kuai-token",
			Name:        "Kuai Token",
			BinanceUSDT: "KTUSDT",
		},
	},
	"DCN": {
		{
			ID:          1876,
			Slug:        "dentacoin",
			Name:        "Dentacoin",
			BinanceUSDT: "DCNUSDT",
		},
	},
	"DWS": {
		{
			ID:          2919,
			Slug:        "dws",
			Name:        "DWS",
			BinanceUSDT: "DWSUSDT",
		},
	},
	"WKCS": {
		{
			ID:          11023,
			Slug:        "wrapped-kucoin-token",
			Name:        "Wrapped KuCoin Token",
			BinanceUSDT: "WKCSUSDT",
		},
	},
	"SKL": {
		{
			ID:          5691,
			Slug:        "skale-network",
			Name:        "SKALE Network",
			BinanceUSDT: "SKLUSDT",
		},
	},
	"MVL": {
		{
			ID:          2982,
			Slug:        "mvl",
			Name:        "MVL",
			BinanceUSDT: "MVLUSDT",
		},
	},
	"AION": {
		{
			ID:          2062,
			Slug:        "aion",
			Name:        "Aion",
			BinanceUSDT: "AIONUSDT",
		},
	},
	"WCELO": {
		{
			ID:          8288,
			Slug:        "wrapped-celo",
			Name:        "Wrapped Celo",
			BinanceUSDT: "WCELOUSDT",
		},
	},
	"TULIP₿": {
		{
			ID:          10477,
			Slug:        "tulips-city",
			Name:        "Tulips City",
			BinanceUSDT: "TULIP₿USDT",
		},
	},
	"ONS": {
		{
			ID:          8160,
			Slug:        "one-share",
			Name:        "One Share",
			BinanceUSDT: "ONSUSDT",
		},
	},
	"BNP": {
		{
			ID:          4872,
			Slug:        "benepit-protocol",
			Name:        "BenePit Protocol",
			BinanceUSDT: "BNPUSDT",
		},
	},
	"RAE": {
		{
			ID:          4887,
			Slug:        "receive-access-ecosystem",
			Name:        "Receive Access Ecosystem",
			BinanceUSDT: "RAEUSDT",
		},
	},
	"BSYS": {
		{
			ID:          6391,
			Slug:        "bsys",
			Name:        "BSYS",
			BinanceUSDT: "BSYSUSDT",
		},
	},
	"MYB": {
		{
			ID:          1902,
			Slug:        "mybit",
			Name:        "MyBit",
			BinanceUSDT: "MYBUSDT",
		},
	},
	"LAIKA": {
		{
			ID:   9776,
			Slug: "laikaprotocol",
			Name: "LaikaProtocol",
		},
		{
			ID:   9515,
			Slug: "laikacoin",
			Name: "LaikaCoin",
		},
	},
	"DPC": {
		{
			ID:          8093,
			Slug:        "dappcents",
			Name:        "DAPPCENTS",
			BinanceUSDT: "DPCUSDT",
		},
	},
	"SHEESHA": {
		{
			ID:          10337,
			Slug:        "sheesha-finance-bep20",
			Name:        "Sheesha Finance [BEP20]",
			BinanceUSDT: "SHEESHAUSDT",
		},
	},
	"UBER": {
		{
			ID:          7917,
			Slug:        "uber-tokenized-stock-ftx",
			Name:        "Uber tokenized stock FTX",
			BinanceUSDT: "UBERUSDT",
		},
	},
	"LYRA": {
		{
			ID:          4542,
			Slug:        "scrypta",
			Name:        "Scrypta",
			BinanceUSDT: "LYRAUSDT",
		},
	},
	"NBXC": {
		{
			ID:          4292,
			Slug:        "nibble",
			Name:        "Nibble",
			BinanceUSDT: "NBXCUSDT",
		},
	},
	"TYPH": {
		{
			ID:          8794,
			Slug:        "typhoon-network",
			Name:        "Typhoon Network",
			BinanceUSDT: "TYPHUSDT",
		},
	},
	"UMI": {
		{
			ID:          9632,
			Slug:        "umi",
			Name:        "UMI",
			BinanceUSDT: "UMIUSDT",
		},
	},
	"ROS": {
		{
			ID:          5730,
			Slug:        "ros-coin",
			Name:        "ROS Coin",
			BinanceUSDT: "ROSUSDT",
		},
	},
	"HNS": {
		{
			ID:          5221,
			Slug:        "handshake",
			Name:        "Handshake",
			BinanceUSDT: "HNSUSDT",
		},
	},
	"SAK3": {
		{
			ID:          10608,
			Slug:        "sake",
			Name:        "Sake",
			BinanceUSDT: "SAK3USDT",
		},
	},
	"VID": {
		{
			ID:          4300,
			Slug:        "videocoin",
			Name:        "VideoCoin",
			BinanceUSDT: "VIDUSDT",
		},
	},
	"YFARMER": {
		{
			ID:          7061,
			Slug:        "yfarmland-token",
			Name:        "YFarmLand Token",
			BinanceUSDT: "YFARMERUSDT",
		},
	},
	"BIDCOM": {
		{
			ID:          10321,
			Slug:        "bidcommerce",
			Name:        "Bidcommerce",
			BinanceUSDT: "BIDCOMUSDT",
		},
	},
	"YFOS": {
		{
			ID:          7319,
			Slug:        "yfos-finance",
			Name:        "YFOS.finance",
			BinanceUSDT: "YFOSUSDT",
		},
	},
	"USELESS": {
		{
			ID:          10851,
			Slug:        "useless",
			Name:        "Useless",
			BinanceUSDT: "USELESSUSDT",
		},
	},
	"MDX": {
		{
			ID:   8335,
			Slug: "mdex",
			Name: "Mdex",
		},
		{
			ID:   8266,
			Slug: "mandala-exchange-token",
			Name: "Mandala Exchange Token",
		},
	},
	"NFLX": {
		{
			ID:   7899,
			Slug: "netflix-tokenized-stock-ftx",
			Name: "Netflix tokenized stock FTX",
		},
		{
			ID:   7928,
			Slug: "netflix-tokenized-stock-bittrex",
			Name: "Netflix tokenized stock Bittrex",
		},
	},
	"CRWNY": {
		{
			ID:          9348,
			Slug:        "crowny",
			Name:        "Crowny",
			BinanceUSDT: "CRWNYUSDT",
		},
	},
	"vDOGE": {
		{
			ID:          9427,
			Slug:        "venus-dogecoin",
			Name:        "Venus Dogecoin",
			BinanceUSDT: "vDOGEUSDT",
		},
	},
	"JRT": {
		{
			ID:          5187,
			Slug:        "jarvis-network",
			Name:        "Jarvis Network",
			BinanceUSDT: "JRTUSDT",
		},
	},
	"CHEESE": {
		{
			ID:          3464,
			Slug:        "cheesecoin",
			Name:        "Cheesecoin",
			BinanceUSDT: "CHEESEUSDT",
		},
	},
	"GOOGL": {
		{
			ID:   7927,
			Slug: "google-tokenized-stock-bittrex",
			Name: "Google tokenized stock Bittrex",
		},
		{
			ID:   7914,
			Slug: "google-tokenized-stock-ftx",
			Name: "Google tokenized stock FTX",
		},
	},
	"KLKS": {
		{
			ID:          3018,
			Slug:        "kalkulus",
			Name:        "Kalkulus",
			BinanceUSDT: "KLKSUSDT",
		},
	},
	"XLT": {
		{
			ID:          6735,
			Slug:        "nexalt",
			Name:        "Nexalt",
			BinanceUSDT: "XLTUSDT",
		},
	},
	"DTOP": {
		{
			ID:          6414,
			Slug:        "data-trade-on-demand-platform",
			Name:        "DTOP Token",
			BinanceUSDT: "DTOPUSDT",
		},
	},
	"APES": {
		{
			ID:          10003,
			Slug:        "apehaven",
			Name:        "ApeHaven",
			BinanceUSDT: "APESUSDT",
		},
	},
	"KNT": {
		{
			ID:   3086,
			Slug: "kora-network-token",
			Name: "Kora Network Token",
		},
		{
			ID:   3383,
			Slug: "knekted",
			Name: "Knekted",
		},
	},
	"COMFI": {
		{
			ID:          9609,
			Slug:        "complifi",
			Name:        "CompliFi",
			BinanceUSDT: "COMFIUSDT",
		},
	},
	"ZNY": {
		{
			ID:          990,
			Slug:        "bitzeny",
			Name:        "Bitzeny",
			BinanceUSDT: "ZNYUSDT",
		},
	},
	"CSM": {
		{
			ID:   10055,
			Slug: "crust-shadow",
			Name: "Crust Shadow",
		},
		{
			ID:   2981,
			Slug: "consentium",
			Name: "Consentium",
		},
	},
	"BZRX": {
		{
			ID:          5810,
			Slug:        "bzx-protocol",
			Name:        "bZx Protocol",
			BinanceUSDT: "BZRXUSDT",
		},
	},
	"POP!": {
		{
			ID:          9520,
			Slug:        "pop",
			Name:        "POP",
			BinanceUSDT: "POP!USDT",
		},
	},
	"CUMMIES": {
		{
			ID:          9212,
			Slug:        "cumrocket",
			Name:        "CUMROCKET",
			BinanceUSDT: "CUMMIESUSDT",
		},
	},
	"BUCKS": {
		{
			ID:          9978,
			Slug:        "buckswap",
			Name:        "BuckSwap",
			BinanceUSDT: "BUCKSUSDT",
		},
	},
	"LYXe": {
		{
			ID:          5625,
			Slug:        "lukso",
			Name:        "LUKSO",
			BinanceUSDT: "LYXeUSDT",
		},
	},
	"TCAT": {
		{
			ID:          3730,
			Slug:        "the-currency-analytics",
			Name:        "The Currency Analytics",
			BinanceUSDT: "TCATUSDT",
		},
	},
	"X8X": {
		{
			ID:          3079,
			Slug:        "x8x-token",
			Name:        "X8X Token",
			BinanceUSDT: "X8XUSDT",
		},
	},
	"AFC": {
		{
			ID:          7469,
			Slug:        "apiary-fund-coin",
			Name:        "Apiary Fund Coin",
			BinanceUSDT: "AFCUSDT",
		},
	},
	"DROPS": {
		{
			ID:          9357,
			Slug:        "defidrop-launchpad",
			Name:        "DefiDrop Launchpad",
			BinanceUSDT: "DROPSUSDT",
		},
	},
	"KOK": {
		{
			ID:          5185,
			Slug:        "keystone-of-opportunity-knowledge",
			Name:        "KOK",
			BinanceUSDT: "KOKUSDT",
		},
	},
	"MAMADOGE": {
		{
			ID:          10835,
			Slug:        "mamadoge",
			Name:        "MAMADOGE",
			BinanceUSDT: "MAMADOGEUSDT",
		},
	},
	"YFET": {
		{
			ID:          7293,
			Slug:        "yfet",
			Name:        "YFET",
			BinanceUSDT: "YFETUSDT",
		},
	},
	"OGO": {
		{
			ID:          3985,
			Slug:        "origo",
			Name:        "Origo",
			BinanceUSDT: "OGOUSDT",
		},
	},
	"WWAN": {
		{
			ID:          8658,
			Slug:        "wrapped-wan",
			Name:        "Wrapped WAN",
			BinanceUSDT: "WWANUSDT",
		},
	},
	"SOP": {
		{
			ID:          2947,
			Slug:        "sopay",
			Name:        "SoPay",
			BinanceUSDT: "SOPUSDT",
		},
	},
	"HAMTARO": {
		{
			ID:          9575,
			Slug:        "hamtaro",
			Name:        "Hamtaro",
			BinanceUSDT: "HAMTAROUSDT",
		},
	},
	"BTC": {
		{
			ID:          1,
			Slug:        "bitcoin",
			Name:        "Bitcoin",
			BinanceUSDT: "BTCUSDT",
		},
	},
	"FLTY": {
		{
			ID:          9722,
			Slug:        "fluity",
			Name:        "Fluity",
			BinanceUSDT: "FLTYUSDT",
		},
	},
	"QUAN": {
		{
			ID:          3390,
			Slug:        "quantis-network",
			Name:        "Quantis Network",
			BinanceUSDT: "QUANUSDT",
		},
	},
	"MCX": {
		{
			ID:          6521,
			Slug:        "machix",
			Name:        "Machi X",
			BinanceUSDT: "MCXUSDT",
		},
	},
	"MEPAD": {
		{
			ID:          9518,
			Slug:        "memepad",
			Name:        "MemePad",
			BinanceUSDT: "MEPADUSDT",
		},
	},
	"PUG": {
		{
			ID:          10460,
			Slug:        "pug-cash",
			Name:        "Pug Cash",
			BinanceUSDT: "PUGUSDT",
		},
	},
	"CLR": {
		{
			ID:          4681,
			Slug:        "color-platform",
			Name:        "Color Platform",
			BinanceUSDT: "CLRUSDT",
		},
	},
	"BPC": {
		{
			ID:          5641,
			Slug:        "backpacker-coin",
			Name:        "BackPacker Coin",
			BinanceUSDT: "BPCUSDT",
		},
	},
	"BOOST": {
		{
			ID:          6862,
			Slug:        "boosted-finance",
			Name:        "Boosted Finance",
			BinanceUSDT: "BOOSTUSDT",
		},
	},
	"RUGBUST": {
		{
			ID:          10247,
			Slug:        "rug-busters",
			Name:        "Rug Busters",
			BinanceUSDT: "RUGBUSTUSDT",
		},
	},
	"RPEPE": {
		{
			ID:          8901,
			Slug:        "rare-pepe",
			Name:        "Rare Pepe",
			BinanceUSDT: "RPEPEUSDT",
		},
	},
	"ADABEAR": {
		{
			ID:          6040,
			Slug:        "3x-short-cardano-token",
			Name:        "3X Short Cardano Token",
			BinanceUSDT: "ADABEARUSDT",
		},
	},
	"VNDC": {
		{
			ID:          4805,
			Slug:        "vndc",
			Name:        "VNDC",
			BinanceUSDT: "VNDCUSDT",
		},
	},
	"FOLD": {
		{
			ID:          10182,
			Slug:        "manifold-finance",
			Name:        "Manifold Finance",
			BinanceUSDT: "FOLDUSDT",
		},
	},
	"XPX": {
		{
			ID:          3126,
			Slug:        "proximax",
			Name:        "ProximaX",
			BinanceUSDT: "XPXUSDT",
		},
	},
	"ORCL5": {
		{
			ID:          8889,
			Slug:        "oracle-top-5",
			Name:        "Oracle Top 5 Tokens Index",
			BinanceUSDT: "ORCL5USDT",
		},
	},
	"ESBC": {
		{
			ID:          3879,
			Slug:        "esbc",
			Name:        "ESBC",
			BinanceUSDT: "ESBCUSDT",
		},
	},
	"VEIL": {
		{
			ID:          3830,
			Slug:        "veil",
			Name:        "Veil",
			BinanceUSDT: "VEILUSDT",
		},
	},
	"NFTD": {
		{
			ID:          9249,
			Slug:        "nftd-protocol",
			Name:        "NFTD Protocol",
			BinanceUSDT: "NFTDUSDT",
		},
	},
	"AR": {
		{
			ID:          5632,
			Slug:        "arweave",
			Name:        "Arweave",
			BinanceUSDT: "ARUSDT",
		},
	},
	"DGTX": {
		{
			ID:          2772,
			Slug:        "digitex",
			Name:        "Digitex",
			BinanceUSDT: "DGTXUSDT",
		},
	},
	"KRW": {
		{
			ID:          11024,
			Slug:        "kingdefi",
			Name:        "KingDeFi",
			BinanceUSDT: "KRWUSDT",
		},
	},
	"PROB": {
		{
			ID:          4586,
			Slug:        "probit-token",
			Name:        "ProBit Token",
			BinanceUSDT: "PROBUSDT",
		},
	},
	"YEAR": {
		{
			ID:          4525,
			Slug:        "lightyears",
			Name:        "Lightyears",
			BinanceUSDT: "YEARUSDT",
		},
	},
	"RNDR": {
		{
			ID:          5690,
			Slug:        "render-token",
			Name:        "Render Token",
			BinanceUSDT: "RNDRUSDT",
		},
	},
	"AWC": {
		{
			ID:          3667,
			Slug:        "atomic-wallet-coin",
			Name:        "Atomic Wallet Coin",
			BinanceUSDT: "AWCUSDT",
		},
	},
	"XTZBEAR": {
		{
			ID:          5463,
			Slug:        "3x-short-tezos-token",
			Name:        "3x Short Tezos Token",
			BinanceUSDT: "XTZBEARUSDT",
		},
	},
	"NRG": {
		{
			ID:          3218,
			Slug:        "energi",
			Name:        "Energi",
			BinanceUSDT: "NRGUSDT",
		},
	},
	"FRN": {
		{
			ID:          1164,
			Slug:        "francs",
			Name:        "Francs",
			BinanceUSDT: "FRNUSDT",
		},
	},
	"BOTS": {
		{
			ID:          8454,
			Slug:        "botocean",
			Name:        "BotOcean",
			BinanceUSDT: "BOTSUSDT",
		},
	},
	"FONT": {
		{
			ID:          8601,
			Slug:        "font",
			Name:        "Font",
			BinanceUSDT: "FONTUSDT",
		},
	},
	"LXT": {
		{
			ID:          3070,
			Slug:        "litex",
			Name:        "Litex",
			BinanceUSDT: "LXTUSDT",
		},
	},
	"STRI": {
		{
			ID:          9395,
			Slug:        "strite",
			Name:        "Strite",
			BinanceUSDT: "STRIUSDT",
		},
	},
	"XSM": {
		{
			ID:          3696,
			Slug:        "spectrumcash",
			Name:        "SpectrumCash",
			BinanceUSDT: "XSMUSDT",
		},
	},
	"TATA": {
		{
			ID:          9972,
			Slug:        "hakunamatata",
			Name:        "HakunaMatata",
			BinanceUSDT: "TATAUSDT",
		},
	},
	"UCN": {
		{
			ID:          3302,
			Slug:        "uchain",
			Name:        "UChain",
			BinanceUSDT: "UCNUSDT",
		},
	},
	"EMON": {
		{
			ID:          9651,
			Slug:        "ethermon",
			Name:        "Ethermon",
			BinanceUSDT: "EMONUSDT",
		},
	},
	"DOGZ": {
		{
			ID:          6422,
			Slug:        "dogz",
			Name:        "Dogz",
			BinanceUSDT: "DOGZUSDT",
		},
	},
	"ERG": {
		{
			ID:          1762,
			Slug:        "ergo",
			Name:        "Ergo",
			BinanceUSDT: "ERGUSDT",
		},
	},
	"SPORE": {
		{
			ID:   9468,
			Slug: "spore",
			Name: "Spore",
		},
		{
			ID:   8589,
			Slug: "spore-engineering",
			Name: "Spore Engineering",
		},
		{
			ID:   7496,
			Slug: "enoki-finance",
			Name: "Enoki Finance",
		},
	},
	"UPT": {
		{
			ID:          5555,
			Slug:        "universal-protocol-token",
			Name:        "Universal Protocol Token",
			BinanceUSDT: "UPTUSDT",
		},
	},
	"MOONSTAR": {
		{
			ID:          9214,
			Slug:        "moonstar",
			Name:        "MoonStar",
			BinanceUSDT: "MOONSTARUSDT",
		},
	},
	"POSS": {
		{
			ID:          3583,
			Slug:        "posscoin",
			Name:        "Posscoin",
			BinanceUSDT: "POSSUSDT",
		},
	},
	"NEVA": {
		{
			ID:          1200,
			Slug:        "nevacoin",
			Name:        "NevaCoin",
			BinanceUSDT: "NEVAUSDT",
		},
	},
	"TRND": {
		{
			ID:          5967,
			Slug:        "trendering",
			Name:        "Trendering",
			BinanceUSDT: "TRNDUSDT",
		},
	},
	"PINT": {
		{
			ID:          8571,
			Slug:        "pub-finance",
			Name:        "Pub Finance",
			BinanceUSDT: "PINTUSDT",
		},
	},
	"MIST": {
		{
			ID:   9218,
			Slug: "mist",
			Name: "Mist",
		},
		{
			ID:   9131,
			Slug: "alchemist",
			Name: "Alchemist",
		},
		{
			ID:   9223,
			Slug: "alchemist-defi-mist",
			Name: "Alchemist DeFi Mist",
		},
	},
	"DVI": {
		{
			ID:          7590,
			Slug:        "dvision-network",
			Name:        "Dvision Network",
			BinanceUSDT: "DVIUSDT",
		},
	},
	"HNY": {
		{
			ID:          7972,
			Slug:        "honey-token",
			Name:        "Honey",
			BinanceUSDT: "HNYUSDT",
		},
	},
	"RISEUP": {
		{
			ID:          10679,
			Slug:        "riseup",
			Name:        "RiseUp",
			BinanceUSDT: "RISEUPUSDT",
		},
	},
	"PTF": {
		{
			ID:          7190,
			Slug:        "powertrade-fuel",
			Name:        "PowerTrade Fuel",
			BinanceUSDT: "PTFUSDT",
		},
	},
	"LOOT": {
		{
			ID:          8030,
			Slug:        "nftlootbox",
			Name:        "NFTLootBox",
			BinanceUSDT: "LOOTUSDT",
		},
	},
	"LCP": {
		{
			ID:          331,
			Slug:        "litecoin-plus",
			Name:        "Litecoin Plus",
			BinanceUSDT: "LCPUSDT",
		},
	},
	"RIGEL": {
		{
			ID:          8304,
			Slug:        "rigel-finance",
			Name:        "Rigel Finance",
			BinanceUSDT: "RIGELUSDT",
		},
	},
	"BSKT": {
		{
			ID:          8733,
			Slug:        "basketcoin",
			Name:        "BasketCoin",
			BinanceUSDT: "BSKTUSDT",
		},
	},
	"TOPC": {
		{
			ID:          2376,
			Slug:        "topchain",
			Name:        "TopChain",
			BinanceUSDT: "TOPCUSDT",
		},
	},
	"ITL": {
		{
			ID:          3476,
			Slug:        "italian-lira",
			Name:        "Italian Lira",
			BinanceUSDT: "ITLUSDT",
		},
	},
	"SFCP": {
		{
			ID:          3856,
			Slug:        "sf-capital",
			Name:        "SF Capital",
			BinanceUSDT: "SFCPUSDT",
		},
	},
	"BARMY": {
		{
			ID:          10340,
			Slug:        "bscarmy",
			Name:        "BscArmy",
			BinanceUSDT: "BARMYUSDT",
		},
	},
	"DYN": {
		{
			ID:          1587,
			Slug:        "dynamic",
			Name:        "Dynamic",
			BinanceUSDT: "DYNUSDT",
		},
	},
	"XPN": {
		{
			ID:          4760,
			Slug:        "pantheon-x",
			Name:        "PANTHEON X",
			BinanceUSDT: "XPNUSDT",
		},
	},
	"WBX": {
		{
			ID:          5450,
			Slug:        "wibx",
			Name:        "WiBX",
			BinanceUSDT: "WBXUSDT",
		},
	},
	"WPR": {
		{
			ID:          2511,
			Slug:        "wepower",
			Name:        "WePower",
			BinanceUSDT: "WPRUSDT",
		},
	},
	"YAMV2": {
		{
			ID:          6657,
			Slug:        "yam-v2",
			Name:        "YAMv2",
			BinanceUSDT: "YAMV2USDT",
		},
	},
	"STETH": {
		{
			ID:   8085,
			Slug: "steth",
			Name: "stETH (Lido)",
		},
		{
			ID:   8311,
			Slug: "stakehound-staked-ether",
			Name: "StakeHound Staked Ether",
		},
	},
	"MRC": {
		{
			ID:          8415,
			Slug:        "meroechain",
			Name:        "MeroeChain",
			BinanceUSDT: "MRCUSDT",
		},
	},
	"AUX": {
		{
			ID:          3362,
			Slug:        "auxilium",
			Name:        "Auxilium",
			BinanceUSDT: "AUXUSDT",
		},
	},
	"BPX": {
		{
			ID:   4941,
			Slug: "bispex",
			Name: "Bispex",
		},
		{
			ID:   10720,
			Slug: "black-phoenix",
			Name: "Black Phoenix",
		},
	},
	"SKC": {
		{
			ID:          5533,
			Slug:        "skinchain",
			Name:        "SKINCHAIN",
			BinanceUSDT: "SKCUSDT",
		},
	},
	"COLDKOALA": {
		{
			ID:          10383,
			Slug:        "cold-koala",
			Name:        "Cold Koala",
			BinanceUSDT: "COLDKOALAUSDT",
		},
	},
	"MSWAP": {
		{
			ID:          8188,
			Slug:        "moneyswap",
			Name:        "MoneySwap",
			BinanceUSDT: "MSWAPUSDT",
		},
	},
	"CATGIRL": {
		{
			ID:          10275,
			Slug:        "catgirl",
			Name:        "Catgirl",
			BinanceUSDT: "CATGIRLUSDT",
		},
	},
	"SPR": {
		{
			ID:          702,
			Slug:        "spreadcoin",
			Name:        "SpreadCoin",
			BinanceUSDT: "SPRUSDT",
		},
	},
	"RES": {
		{
			ID:          5556,
			Slug:        "resfinex-token",
			Name:        "Resfinex Token",
			BinanceUSDT: "RESUSDT",
		},
	},
	"SAITAMA": {
		{
			ID:          10498,
			Slug:        "saitama-inu",
			Name:        "Saitama Inu",
			BinanceUSDT: "SAITAMAUSDT",
		},
	},
	"YRISE": {
		{
			ID:          7441,
			Slug:        "yrise-finance",
			Name:        "yRise Finance",
			BinanceUSDT: "YRISEUSDT",
		},
	},
	"LKK": {
		{
			ID:          1454,
			Slug:        "lykke",
			Name:        "Lykke",
			BinanceUSDT: "LKKUSDT",
		},
	},
	"AMZN": {
		{
			ID:   7890,
			Slug: "amazon-tokenized-stock-ftx",
			Name: "Amazon tokenized stock FTX",
		},
		{
			ID:   7929,
			Slug: "amazon-tokenized-stock-bittrex",
			Name: "Amazon tokenized stock Bittrex",
		},
	},
	"IDEX": {
		{
			ID:          3928,
			Slug:        "idex",
			Name:        "IDEX",
			BinanceUSDT: "IDEXUSDT",
		},
	},
	"XAUTBULL": {
		{
			ID:          6238,
			Slug:        "3x-long-tether-gold-token",
			Name:        "3X Long Tether Gold Token",
			BinanceUSDT: "XAUTBULLUSDT",
		},
	},
	"WENMOON": {
		{
			ID:          9400,
			Slug:        "wenmoon",
			Name:        "WenMoon",
			BinanceUSDT: "WENMOONUSDT",
		},
	},
	"COVA": {
		{
			ID:          3650,
			Slug:        "cova",
			Name:        "COVA",
			BinanceUSDT: "COVAUSDT",
		},
	},
	"EGT": {
		{
			ID:   2885,
			Slug: "egretia",
			Name: "Egretia",
		},
		{
			ID:   9154,
			Slug: "elastic-governance",
			Name: "Elastic Governance",
		},
	},
	"BCD": {
		{
			ID:          2222,
			Slug:        "bitcoin-diamond",
			Name:        "Bitcoin Diamond",
			BinanceUSDT: "BCDUSDT",
		},
	},
	"BTR": {
		{
			ID:   4167,
			Slug: "bitrue-coin",
			Name: "Bitrue Coin",
		},
		{
			ID:   3256,
			Slug: "bitether",
			Name: "Bitether",
		},
	},
	"BDG": {
		{
			ID:          2374,
			Slug:        "bitdegree",
			Name:        "BitDegree",
			BinanceUSDT: "BDGUSDT",
		},
	},
	"HTZ": {
		{
			ID:          10824,
			Slug:        "hertz-network",
			Name:        "Hertz Network",
			BinanceUSDT: "HTZUSDT",
		},
	},
	"GTF": {
		{
			ID:          6457,
			Slug:        "globaltrustfund-token",
			Name:        "GLOBALTRUSTFUND TOKEN",
			BinanceUSDT: "GTFUSDT",
		},
	},
	"LTMS": {
		{
			ID:          10190,
			Slug:        "littlemouse",
			Name:        "LittleMouse",
			BinanceUSDT: "LTMSUSDT",
		},
	},
	"FARM": {
		{
			ID:          6859,
			Slug:        "harvest-finance",
			Name:        "Harvest Finance",
			BinanceUSDT: "FARMUSDT",
		},
	},
	"MEDIC": {
		{
			ID:          916,
			Slug:        "mediccoin",
			Name:        "MedicCoin",
			BinanceUSDT: "MEDICUSDT",
		},
	},
	"SKLAY": {
		{
			ID:          7881,
			Slug:        "sklay",
			Name:        "sKLAY",
			BinanceUSDT: "SKLAYUSDT",
		},
	},
	"EDOGE": {
		{
			ID:          9932,
			Slug:        "elondoge",
			Name:        "ElonDoge",
			BinanceUSDT: "EDOGEUSDT",
		},
	},
	"CFi": {
		{
			ID:          7699,
			Slug:        "cyberfi",
			Name:        "CyberFi Token",
			BinanceUSDT: "CFiUSDT",
		},
	},
	"DVP": {
		{
			ID:          4520,
			Slug:        "decentralized-vulnerability-platform",
			Name:        "Decentralized Vulnerability Platform",
			BinanceUSDT: "DVPUSDT",
		},
	},
	"LOL": {
		{
			ID:   4290,
			Slug: "emogi-network",
			Name: "EMOGI Network",
		},
		{
			ID:   4936,
			Slug: "loltoken",
			Name: "LOLTOKEN",
		},
	},
	"UNN": {
		{
			ID:          8056,
			Slug:        "union-protocol-governance-token",
			Name:        "UNION Protocol Governance Token",
			BinanceUSDT: "UNNUSDT",
		},
	},
	"XUC": {
		{
			ID:          2091,
			Slug:        "exchange-union",
			Name:        "Exchange Union",
			BinanceUSDT: "XUCUSDT",
		},
	},
	"PROT": {
		{
			ID:          9193,
			Slug:        "prostarter",
			Name:        "Prostarter",
			BinanceUSDT: "PROTUSDT",
		},
	},
	"GFARM2": {
		{
			ID:          8444,
			Slug:        "gains-farm-v2",
			Name:        "Gains Farm",
			BinanceUSDT: "GFARM2USDT",
		},
	},
	"JDB": {
		{
			ID:          7066,
			Slug:        "justdobet",
			Name:        "Justdobet",
			BinanceUSDT: "JDBUSDT",
		},
	},
	"DOGEY": {
		{
			ID:          10941,
			Slug:        "dogey-style",
			Name:        "DOGEY STYLE",
			BinanceUSDT: "DOGEYUSDT",
		},
	},
	"RBIES": {
		{
			ID:          1175,
			Slug:        "rubies",
			Name:        "Rubies",
			BinanceUSDT: "RBIESUSDT",
		},
	},
	"NMX": {
		{
			ID:          9438,
			Slug:        "nominex-token",
			Name:        "Nominex Token",
			BinanceUSDT: "NMXUSDT",
		},
	},
	"BULL": {
		{
			ID:   5157,
			Slug: "3x-long-bitcoin-token",
			Name: "3X Long Bitcoin Token",
		},
		{
			ID:   10482,
			Slug: "bull-finance",
			Name: "BULL FINANCE",
		},
		{
			ID:   5056,
			Slug: "buysell",
			Name: "BuySell",
		},
	},
	"FUSE": {
		{
			ID:          5634,
			Slug:        "fuse-network",
			Name:        "Fuse Network",
			BinanceUSDT: "FUSEUSDT",
		},
	},
	"ADABULL": {
		{
			ID:          6079,
			Slug:        "3x-long-cardano-token",
			Name:        "3X Long Cardano Token",
			BinanceUSDT: "ADABULLUSDT",
		},
	},
	"LCC": {
		{
			ID:          2540,
			Slug:        "litecoin-cash",
			Name:        "Litecoin Cash",
			BinanceUSDT: "LCCUSDT",
		},
	},
	"MOLLYDOGE ⭐": {
		{
			ID:          10938,
			Slug:        "mini-hollywood-doge",
			Name:        "Mini Hollywood Doge",
			BinanceUSDT: "MOLLYDOGE ⭐USDT",
		},
	},
	"DIA": {
		{
			ID:          6138,
			Slug:        "dia",
			Name:        "DIA",
			BinanceUSDT: "DIAUSDT",
		},
	},
	"TOL": {
		{
			ID:          3389,
			Slug:        "tolar",
			Name:        "Tolar",
			BinanceUSDT: "TOLUSDT",
		},
	},
	"POLARV3": {
		{
			ID:          9019,
			Slug:        "polar",
			Name:        "Polar",
			BinanceUSDT: "POLARV3USDT",
		},
	},
	"MILK2": {
		{
			ID:          7386,
			Slug:        "spaceswap",
			Name:        "Spaceswap MILK2",
			BinanceUSDT: "MILK2USDT",
		},
	},
	"BUSD": {
		{
			ID:          4687,
			Slug:        "binance-usd",
			Name:        "Binance USD",
			BinanceUSDT: "BUSDUSDT",
		},
	},
	"xDAI": {
		{
			ID:          8635,
			Slug:        "xdaistable",
			Name:        "xDAI",
			BinanceUSDT: "xDAIUSDT",
		},
	},
	"AMS": {
		{
			ID:          1035,
			Slug:        "amsterdamcoin",
			Name:        "AmsterdamCoin",
			BinanceUSDT: "AMSUSDT",
		},
	},
	"SMOL": {
		{
			ID:          7306,
			Slug:        "smol",
			Name:        "Smol",
			BinanceUSDT: "SMOLUSDT",
		},
	},
	"APPLEB": {
		{
			ID:          10500,
			Slug:        "appleb",
			Name:        "APPLEB",
			BinanceUSDT: "APPLEBUSDT",
		},
	},
	"BCUG": {
		{
			ID:          6789,
			Slug:        "blockchain-cuties-universe",
			Name:        "Blockchain Cuties Universe Governance",
			BinanceUSDT: "BCUGUSDT",
		},
	},
	"BTYM": {
		{
			ID:          9655,
			Slug:        "blocktyme",
			Name:        "Blocktyme",
			BinanceUSDT: "BTYMUSDT",
		},
	},
	"PZM": {
		{
			ID:          1681,
			Slug:        "prizm",
			Name:        "PRIZM",
			BinanceUSDT: "PZMUSDT",
		},
	},
	"TST": {
		{
			ID:          5393,
			Slug:        "touch-social",
			Name:        "Touch Social",
			BinanceUSDT: "TSTUSDT",
		},
	},
	"GIC": {
		{
			ID:          3104,
			Slug:        "giant-coin",
			Name:        "Giant",
			BinanceUSDT: "GICUSDT",
		},
	},
	"BCAPS": {
		{
			ID:          9221,
			Slug:        "binacaps",
			Name:        "Binacaps",
			BinanceUSDT: "BCAPSUSDT",
		},
	},
	"STK": {
		{
			ID:          2493,
			Slug:        "stk",
			Name:        "STK",
			BinanceUSDT: "STKUSDT",
		},
	},
	"BTCT": {
		{
			ID:   4159,
			Slug: "bitcoin-token",
			Name: "Bitcoin Token",
		},
		{
			ID:   9486,
			Slug: "bitcoin-trc20",
			Name: "Bitcoin TRC20",
		},
		{
			ID:   5903,
			Slug: "bitcoin-true",
			Name: "Bitcoin True",
		},
	},
	"NGC": {
		{
			ID:          2305,
			Slug:        "naga",
			Name:        "NAGA",
			BinanceUSDT: "NGCUSDT",
		},
	},
	"NSD": {
		{
			ID:          3200,
			Slug:        "nasdacoin",
			Name:        "Nasdacoin",
			BinanceUSDT: "NSDUSDT",
		},
	},
	"KANGAL": {
		{
			ID:          8543,
			Slug:        "kangal",
			Name:        "Kangal",
			BinanceUSDT: "KANGALUSDT",
		},
	},
	"NPX": {
		{
			ID:          2602,
			Slug:        "napoleonx",
			Name:        "NaPoleonX",
			BinanceUSDT: "NPXUSDT",
		},
	},
	"PEKC": {
		{
			ID:          9899,
			Slug:        "peacockcoin",
			Name:        "PEACOCKCOIN",
			BinanceUSDT: "PEKCUSDT",
		},
	},
	"FTI": {
		{
			ID:          2901,
			Slug:        "fanstime",
			Name:        "FansTime",
			BinanceUSDT: "FTIUSDT",
		},
	},
	"TBC": {
		{
			ID:          10128,
			Slug:        "terablock",
			Name:        "TeraBlock",
			BinanceUSDT: "TBCUSDT",
		},
	},
	"DFYN": {
		{
			ID:          9511,
			Slug:        "dfyn-network",
			Name:        "Dfyn Network",
			BinanceUSDT: "DFYNUSDT",
		},
	},
	"BOLI": {
		{
			ID:          1053,
			Slug:        "bolivarcoin",
			Name:        "Bolivarcoin",
			BinanceUSDT: "BOLIUSDT",
		},
	},
	"DORA": {
		{
			ID:   8800,
			Slug: "dora-factory",
			Name: "Dora Factory",
		},
		{
			ID:   10192,
			Slug: "doraemoon",
			Name: "Doraemoon",
		},
	},
	"VBK": {
		{
			ID:          3846,
			Slug:        "veriblock",
			Name:        "VeriBlock",
			BinanceUSDT: "VBKUSDT",
		},
	},
	"SWAMP": {
		{
			ID:          9082,
			Slug:        "swampy",
			Name:        "Swampy",
			BinanceUSDT: "SWAMPUSDT",
		},
	},
	"ETH": {
		{
			ID:          1027,
			Slug:        "ethereum",
			Name:        "Ethereum",
			BinanceUSDT: "ETHUSDT",
		},
	},
	"CRAFT": {
		{
			ID:          7385,
			Slug:        "decraft-finance",
			Name:        "deCraft Finance",
			BinanceUSDT: "CRAFTUSDT",
		},
	},
	"MIDBULL": {
		{
			ID:          6086,
			Slug:        "3x-long-midcap-index-token",
			Name:        "3X Long Midcap Index Token",
			BinanceUSDT: "MIDBULLUSDT",
		},
	},
	"ALT": {
		{
			ID:   10897,
			Slug: "alitas",
			Name: "Alitas",
		},
		{
			ID:   3465,
			Slug: "alt-estate-token",
			Name: "Alt.Estate token",
		},
	},
	"ZPAY": {
		{
			ID:          10929,
			Slug:        "zoidpay",
			Name:        "ZoidPay",
			BinanceUSDT: "ZPAYUSDT",
		},
	},
	"BRTR": {
		{
			ID:          6543,
			Slug:        "barter",
			Name:        "Barter",
			BinanceUSDT: "BRTRUSDT",
		},
	},
	"IBTC": {
		{
			ID:   2131,
			Slug: "ibtc",
			Name: "iBTC",
		},
		{
			ID:   6200,
			Slug: "ibtc-synthetix",
			Name: "iBTC (Synthetix)",
		},
	},
	"SBDO": {
		{
			ID:          8172,
			Slug:        "bdollar-share",
			Name:        "bDollar Share",
			BinanceUSDT: "SBDOUSDT",
		},
	},
	"POFI": {
		{
			ID:          9276,
			Slug:        "pofi",
			Name:        "Pofi",
			BinanceUSDT: "POFIUSDT",
		},
	},
	"FMTA": {
		{
			ID:          7560,
			Slug:        "fundamenta",
			Name:        "Fundamenta",
			BinanceUSDT: "FMTAUSDT",
		},
	},
	"HAPY": {
		{
			ID:          6105,
			Slug:        "hapy-coin",
			Name:        "HAPY Coin",
			BinanceUSDT: "HAPYUSDT",
		},
	},
	"DOGIRA": {
		{
			ID:          9298,
			Slug:        "dogira",
			Name:        "Dogira",
			BinanceUSDT: "DOGIRAUSDT",
		},
	},
	"FEX": {
		{
			ID:          3800,
			Slug:        "fidex-token",
			Name:        "FidexToken",
			BinanceUSDT: "FEXUSDT",
		},
	},
	"UNW": {
		{
			ID:          7635,
			Slug:        "uniworld",
			Name:        "UniWorld",
			BinanceUSDT: "UNWUSDT",
		},
	},
	"HTA": {
		{
			ID:          5365,
			Slug:        "historia",
			Name:        "Historia",
			BinanceUSDT: "HTAUSDT",
		},
	},
	"QUIN": {
		{
			ID:          3659,
			Slug:        "quinads",
			Name:        "QUINADS",
			BinanceUSDT: "QUINUSDT",
		},
	},
	"BINGUS": {
		{
			ID:          9402,
			Slug:        "bingus-token",
			Name:        "Bingus Token",
			BinanceUSDT: "BINGUSUSDT",
		},
	},
	"DEEP": {
		{
			ID:          4251,
			Slug:        "deepcloud-ai",
			Name:        "DeepCloud AI",
			BinanceUSDT: "DEEPUSDT",
		},
	},
	"BBC": {
		{
			ID:   5644,
			Slug: "blue-baikal",
			Name: "Blue Baikal",
		},
		{
			ID:   5432,
			Slug: "bigbang-core",
			Name: "BigBang Core",
		},
	},
	"YOUC": {
		{
			ID:          7321,
			Slug:        "youcash",
			Name:        "yOUcash",
			BinanceUSDT: "YOUCUSDT",
		},
	},
	"SHND": {
		{
			ID:          1106,
			Slug:        "stronghands",
			Name:        "StrongHands",
			BinanceUSDT: "SHNDUSDT",
		},
	},
	"BLOCK": {
		{
			ID:          707,
			Slug:        "blocknet",
			Name:        "Blocknet",
			BinanceUSDT: "BLOCKUSDT",
		},
	},
	"COFE": {
		{
			ID:          8859,
			Slug:        "coffeeswap",
			Name:        "CoffeeSwap",
			BinanceUSDT: "COFEUSDT",
		},
	},
	"ARK": {
		{
			ID:          1586,
			Slug:        "ark",
			Name:        "Ark",
			BinanceUSDT: "ARKUSDT",
		},
	},
	"CNZ": {
		{
			ID:          6397,
			Slug:        "coinzo-token",
			Name:        "Coinzo Token",
			BinanceUSDT: "CNZUSDT",
		},
	},
	"LGO": {
		{
			ID:          2600,
			Slug:        "lgo-token",
			Name:        "LGO Token",
			BinanceUSDT: "LGOUSDT",
		},
	},
	"DFX": {
		{
			ID:          6919,
			Slug:        "definitex",
			Name:        "Definitex",
			BinanceUSDT: "DFXUSDT",
		},
	},
	"SSF": {
		{
			ID:          10615,
			Slug:        "secretsky-finance",
			Name:        "SecretSky.finance",
			BinanceUSDT: "SSFUSDT",
		},
	},
	"BASIC": {
		{
			ID:          5481,
			Slug:        "basic",
			Name:        "BASIC",
			BinanceUSDT: "BASICUSDT",
		},
	},
	"ASAP": {
		{
			ID:          9289,
			Slug:        "chainswap",
			Name:        "Chainswap",
			BinanceUSDT: "ASAPUSDT",
		},
	},
	"BSP": {
		{
			ID:          8566,
			Slug:        "ballswap",
			Name:        "Ballswap",
			BinanceUSDT: "BSPUSDT",
		},
	},
	"BSTY": {
		{
			ID:          644,
			Slug:        "globalboost-y",
			Name:        "GlobalBoost-Y",
			BinanceUSDT: "BSTYUSDT",
		},
	},
	"KOMBAT": {
		{
			ID:          9432,
			Slug:        "crypto-kombat",
			Name:        "Crypto Kombat",
			BinanceUSDT: "KOMBATUSDT",
		},
	},
	"FRAX": {
		{
			ID:          6952,
			Slug:        "frax",
			Name:        "Frax",
			BinanceUSDT: "FRAXUSDT",
		},
	},
	"NCDT": {
		{
			ID:          6933,
			Slug:        "nuco-cloud",
			Name:        "Nuco.cloud",
			BinanceUSDT: "NCDTUSDT",
		},
	},
	"VECT": {
		{
			ID:          5686,
			Slug:        "vectorium",
			Name:        "Vectorium",
			BinanceUSDT: "VECTUSDT",
		},
	},
	"SUPER": {
		{
			ID:   8290,
			Slug: "superfarm",
			Name: "SuperFarm",
		},
		{
			ID:   341,
			Slug: "supercoin",
			Name: "SuperCoin",
		},
	},
	"ATH": {
		{
			ID:          3447,
			Slug:        "atheios",
			Name:        "Atheios",
			BinanceUSDT: "ATHUSDT",
		},
	},
	"ETGP": {
		{
			ID:          3810,
			Slug:        "ethereum-gold-project",
			Name:        "Ethereum Gold Project",
			BinanceUSDT: "ETGPUSDT",
		},
	},
	"FDR": {
		{
			ID:          6976,
			Slug:        "french-digital-reserve",
			Name:        "French Digital Reserve",
			BinanceUSDT: "FDRUSDT",
		},
	},
	"TOP": {
		{
			ID:          3826,
			Slug:        "top",
			Name:        "TOP",
			BinanceUSDT: "TOPUSDT",
		},
	},
	"AXL": {
		{
			ID:          5065,
			Slug:        "axial-entertainment-digital-asset",
			Name:        "Axial Entertainment Digital Asset",
			BinanceUSDT: "AXLUSDT",
		},
	},
	"HGH": {
		{
			ID:          4989,
			Slug:        "hgh-token",
			Name:        "HGH Token",
			BinanceUSDT: "HGHUSDT",
		},
	},
	"YFFI": {
		{
			ID:          6145,
			Slug:        "yffi-finance",
			Name:        "yffi finance",
			BinanceUSDT: "YFFIUSDT",
		},
	},
	"QBC": {
		{
			ID:          278,
			Slug:        "quebecoin",
			Name:        "Quebecoin",
			BinanceUSDT: "QBCUSDT",
		},
	},
	"NODOGE": {
		{
			ID:          10823,
			Slug:        "no-doge",
			Name:        "NO DOGE",
			BinanceUSDT: "NODOGEUSDT",
		},
	},
	"LITH": {
		{
			ID:          10592,
			Slug:        "lith-token",
			Name:        "Lith Token",
			BinanceUSDT: "LITHUSDT",
		},
	},
	"GABECOIN": {
		{
			ID:          10926,
			Slug:        "gabecoin",
			Name:        "Gabecoin",
			BinanceUSDT: "GABECOINUSDT",
		},
	},
	"PAN": {
		{
			ID:          5924,
			Slug:        "pantos",
			Name:        "Pantos",
			BinanceUSDT: "PANUSDT",
		},
	},
	"MAR": {
		{
			ID:          4510,
			Slug:        "mchain",
			Name:        "Mchain",
			BinanceUSDT: "MARUSDT",
		},
	},
	"CORE": {
		{
			ID:          7242,
			Slug:        "cvault-finance",
			Name:        "cVault.finance",
			BinanceUSDT: "COREUSDT",
		},
	},
	"CYFM": {
		{
			ID:          3146,
			Slug:        "cyberfm",
			Name:        "CyberFM",
			BinanceUSDT: "CYFMUSDT",
		},
	},
	"LIQ": {
		{
			ID:          11013,
			Slug:        "liq-protocol",
			Name:        "LIQ Protocol",
			BinanceUSDT: "LIQUSDT",
		},
	},
	"AfroX": {
		{
			ID:          5135,
			Slug:        "afrodex",
			Name:        "AfroDex",
			BinanceUSDT: "AfroXUSDT",
		},
	},
	"OKB": {
		{
			ID:          3897,
			Slug:        "okb",
			Name:        "OKB",
			BinanceUSDT: "OKBUSDT",
		},
	},
	"ARGO": {
		{
			ID:          9663,
			Slug:        "argoapp",
			Name:        "ArGo",
			BinanceUSDT: "ARGOUSDT",
		},
	},
	"XTX": {
		{
			ID:          3844,
			Slug:        "xtock",
			Name:        "Xtock",
			BinanceUSDT: "XTXUSDT",
		},
	},
	"LIQUID": {
		{
			ID:          7650,
			Slug:        "liquidefi",
			Name:        "LIQUID",
			BinanceUSDT: "LIQUIDUSDT",
		},
	},
	"CHP": {
		{
			ID:          2569,
			Slug:        "coinpoker",
			Name:        "CoinPoker",
			BinanceUSDT: "CHPUSDT",
		},
	},
	"MIOTA": {
		{
			ID:          1720,
			Slug:        "iota",
			Name:        "IOTA",
			BinanceUSDT: "MIOTAUSDT",
		},
	},
	"NAV": {
		{
			ID:          377,
			Slug:        "nav-coin",
			Name:        "Navcoin",
			BinanceUSDT: "NAVUSDT",
		},
	},
	"SYM": {
		{
			ID:          4824,
			Slug:        "symverse",
			Name:        "SymVerse",
			BinanceUSDT: "SYMUSDT",
		},
	},
	"RENASCENT": {
		{
			ID:          9748,
			Slug:        "renascent-finance",
			Name:        "Renascent Finance",
			BinanceUSDT: "RENASCENTUSDT",
		},
	},
	"LEO": {
		{
			ID:          3957,
			Slug:        "unus-sed-leo",
			Name:        "UNUS SED LEO",
			BinanceUSDT: "LEOUSDT",
		},
	},
	"DPR": {
		{
			ID:          8894,
			Slug:        "deeper-network",
			Name:        "Deeper Network",
			BinanceUSDT: "DPRUSDT",
		},
	},
	"ART": {
		{
			ID:          2064,
			Slug:        "maecenas",
			Name:        "Maecenas",
			BinanceUSDT: "ARTUSDT",
		},
	},
	"URQA": {
		{
			ID:          9053,
			Slug:        "ureeqa",
			Name:        "UREEQA",
			BinanceUSDT: "URQAUSDT",
		},
	},
	"BLCT": {
		{
			ID:          5280,
			Slug:        "bloomzed-token",
			Name:        "Bloomzed Loyalty Club Ticket",
			BinanceUSDT: "BLCTUSDT",
		},
	},
	"GASG": {
		{
			ID:          8352,
			Slug:        "gasgains",
			Name:        "Gasgains",
			BinanceUSDT: "GASGUSDT",
		},
	},
	"DOPE": {
		{
			ID:          145,
			Slug:        "dopecoin",
			Name:        "DopeCoin",
			BinanceUSDT: "DOPEUSDT",
		},
	},
	"BURGER": {
		{
			ID:          7158,
			Slug:        "burger-swap",
			Name:        "Burger Swap",
			BinanceUSDT: "BURGERUSDT",
		},
	},
	"SFUEL": {
		{
			ID:          8145,
			Slug:        "sparkpoint-fuel",
			Name:        "SparkPoint Fuel",
			BinanceUSDT: "SFUELUSDT",
		},
	},
	"INNBCL": {
		{
			ID:          3787,
			Slug:        "innovative-bioresearch-classic",
			Name:        "Innovative Bioresearch Classic",
			BinanceUSDT: "INNBCLUSDT",
		},
	},
	"PNGN": {
		{
			ID:          9404,
			Slug:        "spacepenguin",
			Name:        "SpacePenguin",
			BinanceUSDT: "PNGNUSDT",
		},
	},
	"DOTA": {
		{
			ID:          10406,
			Slug:        "dota-finance",
			Name:        "Dota Finance",
			BinanceUSDT: "DOTAUSDT",
		},
	},
	"BEAM": {
		{
			ID:          3702,
			Slug:        "beam",
			Name:        "Beam",
			BinanceUSDT: "BEAMUSDT",
		},
	},
	"CHIBI": {
		{
			ID:          10019,
			Slug:        "chibi-inu",
			Name:        "Chibi Inu",
			BinanceUSDT: "CHIBIUSDT",
		},
	},
	"BST": {
		{
			ID:   4132,
			Slug: "bitsten-token",
			Name: "Bitsten Token",
		},
		{
			ID:   3997,
			Slug: "blockstamp",
			Name: "BlockStamp",
		},
	},
	"SCV": {
		{
			ID:          8470,
			Slug:        "super-coinview",
			Name:        "Super CoinView Token",
			BinanceUSDT: "SCVUSDT",
		},
	},
	"QUICK": {
		{
			ID:          8206,
			Slug:        "quickswap",
			Name:        "QuickSwap",
			BinanceUSDT: "QUICKUSDT",
		},
	},
	"TRUBGR": {
		{
			ID:          10902,
			Slug:        "trubadger",
			Name:        "TruBadger",
			BinanceUSDT: "TRUBGRUSDT",
		},
	},
	"CLS": {
		{
			ID:          9562,
			Slug:        "coldstack",
			Name:        "Coldstack",
			BinanceUSDT: "CLSUSDT",
		},
	},
	"SWAPZ": {
		{
			ID:          10557,
			Slug:        "swapz",
			Name:        "Swapz",
			BinanceUSDT: "SWAPZUSDT",
		},
	},
	"HNST": {
		{
			ID:          4035,
			Slug:        "honest",
			Name:        "Honest",
			BinanceUSDT: "HNSTUSDT",
		},
	},
	"MOVE": {
		{
			ID:          7371,
			Slug:        "mover",
			Name:        "Mover",
			BinanceUSDT: "MOVEUSDT",
		},
	},
	"BIA": {
		{
			ID:          3872,
			Slug:        "bilaxy-token",
			Name:        "Bilaxy Token",
			BinanceUSDT: "BIAUSDT",
		},
	},
	"TREAT": {
		{
			ID:          8808,
			Slug:        "treat-dao",
			Name:        "Treat DAO",
			BinanceUSDT: "TREATUSDT",
		},
	},
	"LTX": {
		{
			ID:          7616,
			Slug:        "lattice-token",
			Name:        "Lattice Token",
			BinanceUSDT: "LTXUSDT",
		},
	},
	"G999": {
		{
			ID:          8431,
			Slug:        "g999",
			Name:        "G999",
			BinanceUSDT: "G999USDT",
		},
	},
	"PVM": {
		{
			ID:          9586,
			Slug:        "privateum-initiative",
			Name:        "PRIVATEUM INITIATIVE",
			BinanceUSDT: "PVMUSDT",
		},
	},
	"YCE": {
		{
			ID:          4381,
			Slug:        "myce",
			Name:        "MYCE",
			BinanceUSDT: "YCEUSDT",
		},
	},
	"IFEX": {
		{
			ID:          7949,
			Slug:        "interfinex",
			Name:        "Interfinex",
			BinanceUSDT: "IFEXUSDT",
		},
	},
	"DIGI": {
		{
			ID:          9551,
			Slug:        "digible",
			Name:        "Digible",
			BinanceUSDT: "DIGIUSDT",
		},
	},
	"BDK": {
		{
			ID:          5288,
			Slug:        "bidesk",
			Name:        "Bidesk",
			BinanceUSDT: "BDKUSDT",
		},
	},
	"USDEX": {
		{
			ID:          8348,
			Slug:        "usdex",
			Name:        "USDEX",
			BinanceUSDT: "USDEXUSDT",
		},
	},
	"WHITE": {
		{
			ID:          8120,
			Slug:        "whiteheart",
			Name:        "Whiteheart",
			BinanceUSDT: "WHITEUSDT",
		},
	},
	"RAPTR": {
		{
			ID:          9628,
			Slug:        "raptor-token",
			Name:        "Raptor Token",
			BinanceUSDT: "RAPTRUSDT",
		},
	},
	"LINA": {
		{
			ID:   7102,
			Slug: "linear",
			Name: "Linear",
		},
		{
			ID:   3083,
			Slug: "lina",
			Name: "LINA",
		},
	},
	"EVC": {
		{
			ID:   2237,
			Slug: "eventchain",
			Name: "EventChain",
		},
		{
			ID:   4112,
			Slug: "eco-value-coin",
			Name: "Eco Value Coin",
		},
		{
			ID:   8820,
			Slug: "evrice",
			Name: "Evrice",
		},
	},
	"TYPE": {
		{
			ID:          3505,
			Slug:        "typerium",
			Name:        "Typerium",
			BinanceUSDT: "TYPEUSDT",
		},
	},
	"BRG": {
		{
			ID:          7096,
			Slug:        "bridge-oracle",
			Name:        "Bridge Oracle",
			BinanceUSDT: "BRGUSDT",
		},
	},
	"JEM": {
		{
			ID:          7723,
			Slug:        "itchiro-games",
			Name:        "Itchiro Games",
			BinanceUSDT: "JEMUSDT",
		},
	},
	"BTRN": {
		{
			ID:          2690,
			Slug:        "biotron",
			Name:        "Biotron",
			BinanceUSDT: "BTRNUSDT",
		},
	},
	"KARMA": {
		{
			ID:   3137,
			Slug: "karma-eos",
			Name: "KARMA",
		},
		{
			ID:   5935,
			Slug: "karma-dao",
			Name: "Karma DAO",
		},
	},
	"BAX": {
		{
			ID:          2572,
			Slug:        "babb",
			Name:        "BABB",
			BinanceUSDT: "BAXUSDT",
		},
	},
	"LEM": {
		{
			ID:          9142,
			Slug:        "lemur-finance",
			Name:        "Lemur Finance",
			BinanceUSDT: "LEMUSDT",
		},
	},
	"PENDLE": {
		{
			ID:          9481,
			Slug:        "pendle",
			Name:        "Pendle",
			BinanceUSDT: "PENDLEUSDT",
		},
	},
	"UBIN": {
		{
			ID:          7515,
			Slug:        "ubiner",
			Name:        "Ubiner",
			BinanceUSDT: "UBINUSDT",
		},
	},
	"NAZ": {
		{
			ID:          7516,
			Slug:        "naz-coin",
			Name:        "Naz Coin",
			BinanceUSDT: "NAZUSDT",
		},
	},
	"PONZU": {
		{
			ID:          10859,
			Slug:        "ponzu-inu",
			Name:        "Ponzu Inu",
			BinanceUSDT: "PONZUUSDT",
		},
	},
	"SNN": {
		{
			ID:          6179,
			Slug:        "sechain",
			Name:        "SeChain",
			BinanceUSDT: "SNNUSDT",
		},
	},
	"SAFEGALAXY": {
		{
			ID:          9297,
			Slug:        "safegalaxy",
			Name:        "SafeGalaxy",
			BinanceUSDT: "SAFEGALAXYUSDT",
		},
	},
	"DEFI+L": {
		{
			ID:          7337,
			Slug:        "piedao-defi-large-cap",
			Name:        "PieDAO DEFI Large Cap",
			BinanceUSDT: "DEFI+LUSDT",
		},
	},
	"SFR": {
		{
			ID:          7205,
			Slug:        "safari",
			Name:        "Safari",
			BinanceUSDT: "SFRUSDT",
		},
	},
	"VX": {
		{
			ID:          5246,
			Slug:        "vitex-coin",
			Name:        "ViteX Coin",
			BinanceUSDT: "VXUSDT",
		},
	},
	"MFT": {
		{
			ID:          2896,
			Slug:        "mainframe",
			Name:        "Hifi Finance",
			BinanceUSDT: "MFTUSDT",
		},
	},
	"BAL": {
		{
			ID:          5728,
			Slug:        "balancer",
			Name:        "Balancer",
			BinanceUSDT: "BALUSDT",
		},
	},
	"LOA": {
		{
			ID:          5516,
			Slug:        "loa-protocol",
			Name:        "LOA Protocol",
			BinanceUSDT: "LOAUSDT",
		},
	},
	"KOMP": {
		{
			ID:          7264,
			Slug:        "kompass",
			Name:        "Kompass",
			BinanceUSDT: "KOMPUSDT",
		},
	},
	"TTT": {
		{
			ID:   5514,
			Slug: "the-transfer-token",
			Name: "The Transfer Token",
		},
		{
			ID:   4098,
			Slug: "tapcoin",
			Name: "Tapcoin",
		},
	},
	"UST": {
		{
			ID:          7129,
			Slug:        "terrausd",
			Name:        "TerraUSD",
			BinanceUSDT: "USTUSDT",
		},
	},
	"MFC": {
		{
			ID:          3837,
			Slug:        "mfcoin",
			Name:        "MFCoin",
			BinanceUSDT: "MFCUSDT",
		},
	},
	"GOLDR": {
		{
			ID:          6624,
			Slug:        "golden-ratio-coin",
			Name:        "Golden Ratio Coin",
			BinanceUSDT: "GOLDRUSDT",
		},
	},
	"CHI": {
		{
			ID:   5541,
			Slug: "xaya",
			Name: "Xaya",
		},
		{
			ID:   5709,
			Slug: "chi-gastoken",
			Name: "Chi Gastoken",
		},
	},
	"BITCI": {
		{
			ID:          8357,
			Slug:        "bitcicoin",
			Name:        "Bitcicoin",
			BinanceUSDT: "BITCIUSDT",
		},
	},
	"BTA": {
		{
			ID:   945,
			Slug: "bata",
			Name: "Bata",
		},
		{
			ID:   8957,
			Slug: "bitcoin-asset",
			Name: "Bitcoin Asset",
		},
	},
	"Sunder": {
		{
			ID:          10298,
			Slug:        "sunder-goverance-token",
			Name:        "Sunder Goverance Token",
			BinanceUSDT: "SunderUSDT",
		},
	},
	"FREL": {
		{
			ID:          9976,
			Slug:        "freela",
			Name:        "Freela",
			BinanceUSDT: "FRELUSDT",
		},
	},
	"GET": {
		{
			ID:   2354,
			Slug: "get-protocol",
			Name: "GET Protocol",
		},
		{
			ID:   3127,
			Slug: "themis",
			Name: "Themis",
		},
		{
			ID:   7160,
			Slug: "gire-token",
			Name: "Gire Token",
		},
	},
	"DOGGY": {
		{
			ID:          9693,
			Slug:        "doggy",
			Name:        "DOGGY",
			BinanceUSDT: "DOGGYUSDT",
		},
	},
	"DND": {
		{
			ID:          9578,
			Slug:        "dungeonswap",
			Name:        "Dungeonswap",
			BinanceUSDT: "DNDUSDT",
		},
	},
	"LTHN": {
		{
			ID:          2185,
			Slug:        "lethean",
			Name:        "Lethean",
			BinanceUSDT: "LTHNUSDT",
		},
	},
	"QOOB": {
		{
			ID:          7347,
			Slug:        "qoober",
			Name:        "QOOBER",
			BinanceUSDT: "QOOBUSDT",
		},
	},
	"VRA": {
		{
			ID:          3816,
			Slug:        "verasity",
			Name:        "Verasity",
			BinanceUSDT: "VRAUSDT",
		},
	},
	"BMON": {
		{
			ID:          10704,
			Slug:        "binamon",
			Name:        "Binamon",
			BinanceUSDT: "BMONUSDT",
		},
	},
	"SUN": {
		{
			ID:          10529,
			Slug:        "sun-token",
			Name:        "Sun (New)",
			BinanceUSDT: "SUNUSDT",
		},
	},
	"OCAT": {
		{
			ID:          10633,
			Slug:        "orange-cat-token",
			Name:        "Orange Cat Token",
			BinanceUSDT: "OCATUSDT",
		},
	},
	"LIGHT": {
		{
			ID:          8801,
			Slug:        "lightning",
			Name:        "Lightning",
			BinanceUSDT: "LIGHTUSDT",
		},
	},
	"PRO": {
		{
			ID:          1974,
			Slug:        "propy",
			Name:        "Propy",
			BinanceUSDT: "PROUSDT",
		},
	},
	"PORNROCKET": {
		{
			ID:          10165,
			Slug:        "pornrocket",
			Name:        "PORNROCKET",
			BinanceUSDT: "PORNROCKETUSDT",
		},
	},
	"XPM": {
		{
			ID:          42,
			Slug:        "primecoin",
			Name:        "Primecoin",
			BinanceUSDT: "XPMUSDT",
		},
	},
	"CUBE": {
		{
			ID:          5338,
			Slug:        "somnium-space-cubes",
			Name:        "Somnium Space Cubes",
			BinanceUSDT: "CUBEUSDT",
		},
	},
	"MBTC": {
		{
			ID:          10245,
			Slug:        "micro-bitcoin-finance",
			Name:        "Micro Bitcoin Finance",
			BinanceUSDT: "MBTCUSDT",
		},
	},
	"SEFA": {
		{
			ID:          5425,
			Slug:        "mesefa",
			Name:        "MESEFA",
			BinanceUSDT: "SEFAUSDT",
		},
	},
	"DMCH": {
		{
			ID:          5622,
			Slug:        "darma-cash",
			Name:        "Darma Cash",
			BinanceUSDT: "DMCHUSDT",
		},
	},
	"SPT": {
		{
			ID:          3777,
			Slug:        "spectrum",
			Name:        "Spectrum",
			BinanceUSDT: "SPTUSDT",
		},
	},
	"AKA": {
		{
			ID:          3151,
			Slug:        "akroma",
			Name:        "Akroma",
			BinanceUSDT: "AKAUSDT",
		},
	},
	"MFO": {
		{
			ID:          10411,
			Slug:        "moonfarm-finance",
			Name:        "Moonfarm Finance",
			BinanceUSDT: "MFOUSDT",
		},
	},
	"QUEENSHIBA": {
		{
			ID:          10983,
			Slug:        "queen-of-shiba",
			Name:        "Queen of Shiba",
			BinanceUSDT: "QUEENSHIBAUSDT",
		},
	},
	"OSB": {
		{
			ID:          7789,
			Slug:        "oasisbloc",
			Name:        "OASISBloc",
			BinanceUSDT: "OSBUSDT",
		},
	},
	"SAFETESLA": {
		{
			ID:          9393,
			Slug:        "safetesla",
			Name:        "Safetesla",
			BinanceUSDT: "SAFETESLAUSDT",
		},
	},
	"WGOLD": {
		{
			ID:          9332,
			Slug:        "apwars",
			Name:        "APWars",
			BinanceUSDT: "WGOLDUSDT",
		},
	},
	"PTM": {
		{
			ID:          8358,
			Slug:        "potentiam",
			Name:        "Potentiam",
			BinanceUSDT: "PTMUSDT",
		},
	},
	"UFT": {
		{
			ID:          7412,
			Slug:        "unilend",
			Name:        "UniLend",
			BinanceUSDT: "UFTUSDT",
		},
	},
	"SPARTA": {
		{
			ID:          6992,
			Slug:        "spartan-protocol",
			Name:        "Spartan Protocol",
			BinanceUSDT: "SPARTAUSDT",
		},
	},
	"KODX": {
		{
			ID:          7828,
			Slug:        "king-of-defi",
			Name:        "KING OF DEFI",
			BinanceUSDT: "KODXUSDT",
		},
	},
	"TITAN": {
		{
			ID:   7206,
			Slug: "titanswap",
			Name: "TitanSwap",
		},
		{
			ID:   10467,
			Slug: "iron-titanium-token",
			Name: "IRON Titanium Token",
		},
	},
	"SNOGE": {
		{
			ID:          10007,
			Slug:        "snoop-doge",
			Name:        "Snoop Doge",
			BinanceUSDT: "SNOGEUSDT",
		},
	},
	"CCN": {
		{
			ID:          3770,
			Slug:        "customcontractnetwork",
			Name:        "CustomContractNetwork",
			BinanceUSDT: "CCNUSDT",
		},
	},
	"ADD": {
		{
			ID:          5834,
			Slug:        "add-xyz",
			Name:        "Add.xyz",
			BinanceUSDT: "ADDUSDT",
		},
	},
	"CDL": {
		{
			ID:          4910,
			Slug:        "coindeal-token",
			Name:        "CoinDeal Token",
			BinanceUSDT: "CDLUSDT",
		},
	},
	"USF": {
		{
			ID:          8896,
			Slug:        "unslashed-finance",
			Name:        "Unslashed Finance",
			BinanceUSDT: "USFUSDT",
		},
	},
	"SWAPS": {
		{
			ID:          9412,
			Slug:        "nftswaps",
			Name:        "NFTSwaps",
			BinanceUSDT: "SWAPSUSDT",
		},
	},
	"OST": {
		{
			ID:          2296,
			Slug:        "ost",
			Name:        "OST",
			BinanceUSDT: "OSTUSDT",
		},
	},
	"HPT": {
		{
			ID:          3721,
			Slug:        "huobi-pool-token",
			Name:        "Huobi Pool Token",
			BinanceUSDT: "HPTUSDT",
		},
	},
	"MSC": {
		{
			ID:          8428,
			Slug:        "monster-slayer-finance",
			Name:        "Monster Slayer Cash",
			BinanceUSDT: "MSCUSDT",
		},
	},
	"$RICH": {
		{
			ID:          10068,
			Slug:        "richierich-coin",
			Name:        "RichieRich Coin",
			BinanceUSDT: "$RICHUSDT",
		},
	},
	"DEXG": {
		{
			ID:          6684,
			Slug:        "dextoken",
			Name:        "Dextoken",
			BinanceUSDT: "DEXGUSDT",
		},
	},
	"CHND": {
		{
			ID:          5427,
			Slug:        "cashhand",
			Name:        "Cashhand",
			BinanceUSDT: "CHNDUSDT",
		},
	},
	"FC": {
		{
			ID:          9078,
			Slug:        "fanscoin",
			Name:        "FansCoin",
			BinanceUSDT: "FCUSDT",
		},
	},
	"FLX": {
		{
			ID:   7746,
			Slug: "felixo-coin",
			Name: "Felixo Coin",
		},
		{
			ID:   9493,
			Slug: "reflexer-ungovernance-token",
			Name: "Reflexer Ungovernance Token",
		},
	},
	"UPX": {
		{
			ID:          3752,
			Slug:        "uplexa",
			Name:        "uPlexa",
			BinanceUSDT: "UPXUSDT",
		},
	},
	"SUSHI": {
		{
			ID:          6758,
			Slug:        "sushiswap",
			Name:        "SushiSwap",
			BinanceUSDT: "SUSHIUSDT",
		},
	},
	"BIGGLES": {
		{
			ID:          11028,
			Slug:        "mr-bigglesworth",
			Name:        "Mr Bigglesworth",
			BinanceUSDT: "BIGGLESUSDT",
		},
	},
	"ENJ": {
		{
			ID:          2130,
			Slug:        "enjin-coin",
			Name:        "Enjin Coin",
			BinanceUSDT: "ENJUSDT",
		},
	},
	"MoFi": {
		{
			ID:          9132,
			Slug:        "mobi-finance",
			Name:        "MobiFi",
			BinanceUSDT: "MoFiUSDT",
		},
	},
	"ADS": {
		{
			ID:          1883,
			Slug:        "adshares",
			Name:        "Adshares",
			BinanceUSDT: "ADSUSDT",
		},
	},
	"TRXDOWN": {
		{
			ID:          7004,
			Slug:        "trxdown",
			Name:        "TRXDOWN",
			BinanceUSDT: "TRXDOWNUSDT",
		},
	},
	"EHASH": {
		{
			ID:          8678,
			Slug:        "ehash",
			Name:        "EHash",
			BinanceUSDT: "EHASHUSDT",
		},
	},
	"CHAT": {
		{
			ID:          2427,
			Slug:        "chatcoin",
			Name:        "ChatCoin",
			BinanceUSDT: "CHATUSDT",
		},
	},
	"LOBS": {
		{
			ID:          3203,
			Slug:        "lobstex",
			Name:        "Lobstex",
			BinanceUSDT: "LOBSUSDT",
		},
	},
	"xMOON": {
		{
			ID:          7396,
			Slug:        "moon",
			Name:        "r/CryptoCurrency Moons",
			BinanceUSDT: "xMOONUSDT",
		},
	},
	"ONION": {
		{
			ID:          1881,
			Slug:        "deeponion",
			Name:        "DeepOnion",
			BinanceUSDT: "ONIONUSDT",
		},
	},
	"QUACK": {
		{
			ID:          10641,
			Slug:        "richquack-com",
			Name:        "RichQUACK.com",
			BinanceUSDT: "QUACKUSDT",
		},
	},
	"FNT": {
		{
			ID:          5871,
			Slug:        "falcon-project",
			Name:        "Falcon Project",
			BinanceUSDT: "FNTUSDT",
		},
	},
	"VFOX": {
		{
			ID:          10290,
			Slug:        "rfox-finance",
			Name:        "RFOX Finance",
			BinanceUSDT: "VFOXUSDT",
		},
	},
	"YFW": {
		{
			ID:          7983,
			Slug:        "yfworld",
			Name:        "YFWorld",
			BinanceUSDT: "YFWUSDT",
		},
	},
	"ETHY": {
		{
			ID:          7743,
			Slug:        "ethereum-yield",
			Name:        "Ethereum Yield",
			BinanceUSDT: "ETHYUSDT",
		},
	},
	"PAWS": {
		{
			ID:          9573,
			Slug:        "animal-adoption-advocacy",
			Name:        "Animal Adoption Advocacy",
			BinanceUSDT: "PAWSUSDT",
		},
	},
	"SIGNA": {
		{
			ID:          10776,
			Slug:        "signum",
			Name:        "Signum",
			BinanceUSDT: "SIGNAUSDT",
		},
	},
	"VOLLAR": {
		{
			ID:          3891,
			Slug:        "v-dimension",
			Name:        "V-Dimension",
			BinanceUSDT: "VOLLARUSDT",
		},
	},
	"WANLINK": {
		{
			ID:          8651,
			Slug:        "wanlink",
			Name:        "wanLINK",
			BinanceUSDT: "WANLINKUSDT",
		},
	},
	"PANTHER": {
		{
			ID:          9778,
			Slug:        "pantherswap",
			Name:        "PantherSwap",
			BinanceUSDT: "PANTHERUSDT",
		},
	},
	"OAP": {
		{
			ID:          6970,
			Slug:        "openalexa-protocol",
			Name:        "OpenAlexa Protocol",
			BinanceUSDT: "OAPUSDT",
		},
	},
	"DOS": {
		{
			ID:   3809,
			Slug: "dos-network",
			Name: "DOS Network",
		},
		{
			ID:   5942,
			Slug: "demos",
			Name: "DEMOS",
		},
	},
	"ADX": {
		{
			ID:          1768,
			Slug:        "adx-net",
			Name:        "AdEx Network",
			BinanceUSDT: "ADXUSDT",
		},
	},
	"MD+": {
		{
			ID:          8392,
			Slug:        "moondayplus",
			Name:        "MoonDayPlus",
			BinanceUSDT: "MD+USDT",
		},
	},
	"ARTHX": {
		{
			ID:          10796,
			Slug:        "arthx",
			Name:        "ARTH Shares",
			BinanceUSDT: "ARTHXUSDT",
		},
	},
	"IOUX": {
		{
			ID:          3996,
			Slug:        "iou",
			Name:        "IOU",
			BinanceUSDT: "IOUXUSDT",
		},
	},
	"PROPEL": {
		{
			ID:          4043,
			Slug:        "payrue-propel",
			Name:        "PayRue (Propel)",
			BinanceUSDT: "PROPELUSDT",
		},
	},
	"SDT": {
		{
			ID:   8299,
			Slug: "stake-dao",
			Name: "Stake DAO",
		},
		{
			ID:   6370,
			Slug: "terra-sdt",
			Name: "Terra SDT",
		},
	},
	"VARC": {
		{
			ID:          5628,
			Slug:        "varc",
			Name:        "VARC",
			BinanceUSDT: "VARCUSDT",
		},
	},
	"GT": {
		{
			ID:          4269,
			Slug:        "gatetoken",
			Name:        "GateToken",
			BinanceUSDT: "GTUSDT",
		},
	},
	"AE": {
		{
			ID:          1700,
			Slug:        "aeternity",
			Name:        "Aeternity",
			BinanceUSDT: "AEUSDT",
		},
	},
	"USDT": {
		{
			ID:          825,
			Slug:        "tether",
			Name:        "Tether",
			BinanceUSDT: "USDTUSDT",
		},
	},
	"SIL": {
		{
			ID:          10136,
			Slug:        "sil-finance",
			Name:        "SIL.FINANCE",
			BinanceUSDT: "SILUSDT",
		},
	},
	"TRV": {
		{
			ID:          4060,
			Slug:        "trustverse",
			Name:        "TrustVerse",
			BinanceUSDT: "TRVUSDT",
		},
	},
	"BLP": {
		{
			ID:          10326,
			Slug:        "bullperks",
			Name:        "BullPerks",
			BinanceUSDT: "BLPUSDT",
		},
	},
	"LYFE": {
		{
			ID:          5786,
			Slug:        "lyfe",
			Name:        "LYFE",
			BinanceUSDT: "LYFEUSDT",
		},
	},
	"BRY": {
		{
			ID:          8483,
			Slug:        "berry-data",
			Name:        "Berry Data",
			BinanceUSDT: "BRYUSDT",
		},
	},
	"BIKI": {
		{
			ID:          5325,
			Slug:        "biki",
			Name:        "BIKI",
			BinanceUSDT: "BIKIUSDT",
		},
	},
	"WON": {
		{
			ID:          7504,
			Slug:        "weblock",
			Name:        "WeBlock",
			BinanceUSDT: "WONUSDT",
		},
	},
	"YFRM": {
		{
			ID:          7530,
			Slug:        "yearn-finance-red-moon",
			Name:        "Yearn Finance Red Moon",
			BinanceUSDT: "YFRMUSDT",
		},
	},
	"SMT": {
		{
			ID:          2277,
			Slug:        "smartmesh",
			Name:        "SmartMesh",
			BinanceUSDT: "SMTUSDT",
		},
	},
	"ETHO": {
		{
			ID:          3452,
			Slug:        "ether-1",
			Name:        "Etho Protocol",
			BinanceUSDT: "ETHOUSDT",
		},
	},
	"PLE": {
		{
			ID:          9606,
			Slug:        "plethori",
			Name:        "Plethori",
			BinanceUSDT: "PLEUSDT",
		},
	},
	"ASTRO": {
		{
			ID:          6894,
			Slug:        "astrotools",
			Name:        "AstroTools",
			BinanceUSDT: "ASTROUSDT",
		},
	},
	"WHOLE": {
		{
			ID:          7633,
			Slug:        "wormhole-finance",
			Name:        "wormhole.finance",
			BinanceUSDT: "WHOLEUSDT",
		},
	},
	"NEXO": {
		{
			ID:          2694,
			Slug:        "nexo",
			Name:        "Nexo",
			BinanceUSDT: "NEXOUSDT",
		},
	},
	"BVND": {
		{
			ID:          8058,
			Slug:        "binance-vnd",
			Name:        "Binance VND",
			BinanceUSDT: "BVNDUSDT",
		},
	},
	"$BITE": {
		{
			ID:          10314,
			Slug:        "dragonbite",
			Name:        "DragonBite",
			BinanceUSDT: "$BITEUSDT",
		},
	},
	"APC": {
		{
			ID:          3512,
			Slug:        "alpha-coin",
			Name:        "Alpha Coin",
			BinanceUSDT: "APCUSDT",
		},
	},
	"SXAG": {
		{
			ID:          5863,
			Slug:        "sxag",
			Name:        "sXAG",
			BinanceUSDT: "SXAGUSDT",
		},
	},
	"CO2": {
		{
			ID:          8931,
			Slug:        "collective",
			Name:        "Collective",
			BinanceUSDT: "CO2USDT",
		},
	},
	"CNUS": {
		{
			ID:          3648,
			Slug:        "coinus",
			Name:        "CoinUs",
			BinanceUSDT: "CNUSUSDT",
		},
	},
	"WEWON": {
		{
			ID:          10645,
			Slug:        "wewon-world",
			Name:        "WeWon World",
			BinanceUSDT: "WEWONUSDT",
		},
	},
	"UGAS": {
		{
			ID:          3863,
			Slug:        "ugas",
			Name:        "UGAS",
			BinanceUSDT: "UGASUSDT",
		},
	},
	"TNS": {
		{
			ID:          2704,
			Slug:        "transcodium",
			Name:        "Transcodium",
			BinanceUSDT: "TNSUSDT",
		},
	},
	"DCL": {
		{
			ID:          8201,
			Slug:        "delphi-chain-link",
			Name:        "Delphi Chain Link",
			BinanceUSDT: "DCLUSDT",
		},
	},
	"ABUSD": {
		{
			ID:          5755,
			Slug:        "aave-busd",
			Name:        "Aave BUSD",
			BinanceUSDT: "ABUSDUSDT",
		},
	},
	"CARD": {
		{
			ID:          2891,
			Slug:        "cardstack",
			Name:        "Cardstack",
			BinanceUSDT: "CARDUSDT",
		},
	},
	"KIMOCHI": {
		{
			ID:          8507,
			Slug:        "kimochi-finance",
			Name:        "Kimochi Finance",
			BinanceUSDT: "KIMOCHIUSDT",
		},
	},
	"MSR": {
		{
			ID:          2674,
			Slug:        "masari",
			Name:        "Masari",
			BinanceUSDT: "MSRUSDT",
		},
	},
	"MINIBABYDOGE": {
		{
			ID:          10806,
			Slug:        "mini-baby-doge",
			Name:        "Mini Baby Doge",
			BinanceUSDT: "MINIBABYDOGEUSDT",
		},
	},
	"TOR": {
		{
			ID:          5518,
			Slug:        "torex",
			Name:        "Torex",
			BinanceUSDT: "TORUSDT",
		},
	},
	"ANY": {
		{
			ID:          5892,
			Slug:        "anyswap",
			Name:        "Anyswap",
			BinanceUSDT: "ANYUSDT",
		},
	},
	"IDH": {
		{
			ID:          2459,
			Slug:        "indahash",
			Name:        "indaHash",
			BinanceUSDT: "IDHUSDT",
		},
	},
	"O3": {
		{
			ID:          9588,
			Slug:        "o3swap",
			Name:        "O3Swap",
			BinanceUSDT: "O3USDT",
		},
	},
	"XNODE": {
		{
			ID:          8372,
			Slug:        "xnode",
			Name:        "XNODE",
			BinanceUSDT: "XNODEUSDT",
		},
	},
	"SPHR": {
		{
			ID:          914,
			Slug:        "sphere",
			Name:        "Sphere",
			BinanceUSDT: "SPHRUSDT",
		},
	},
	"DUO": {
		{
			ID:   1089,
			Slug: "parallelcoin",
			Name: "ParallelCoin",
		},
		{
			ID:   3905,
			Slug: "duo-network-token",
			Name: "DUO Network Token",
		},
	},
	"GERO": {
		{
			ID:          9904,
			Slug:        "gerowallet",
			Name:        "GeroWallet",
			BinanceUSDT: "GEROUSDT",
		},
	},
	"TRISM": {
		{
			ID:          8185,
			Slug:        "trism",
			Name:        "Trism",
			BinanceUSDT: "TRISMUSDT",
		},
	},
	"SET": {
		{
			ID:   3764,
			Slug: "save-environment-token",
			Name: "Save Environment Token",
		},
		{
			ID:   10157,
			Slug: "sustainable-energy-token",
			Name: "Sustainable Energy Token",
		},
	},
	"ELONGATE": {
		{
			ID:          9178,
			Slug:        "elongate",
			Name:        "ElonGate",
			BinanceUSDT: "ELONGATEUSDT",
		},
	},
	"STRIKE": {
		{
			ID:          9220,
			Slug:        "strikecoin",
			Name:        "StrikeCoin",
			BinanceUSDT: "STRIKEUSDT",
		},
	},
	"VNTW": {
		{
			ID:          9230,
			Slug:        "value-network",
			Name:        "Value Network",
			BinanceUSDT: "VNTWUSDT",
		},
	},
	"MRCH": {
		{
			ID:          8971,
			Slug:        "merchdao",
			Name:        "MerchDAO",
			BinanceUSDT: "MRCHUSDT",
		},
	},
	"ALGO": {
		{
			ID:          4030,
			Slug:        "algorand",
			Name:        "Algorand",
			BinanceUSDT: "ALGOUSDT",
		},
	},
	"SHIELDNET": {
		{
			ID:          9457,
			Slug:        "shield-network",
			Name:        "Shield Network",
			BinanceUSDT: "SHIELDNETUSDT",
		},
	},
	"2KEY": {
		{
			ID:          5587,
			Slug:        "2key-network",
			Name:        "2key.network",
			BinanceUSDT: "2KEYUSDT",
		},
	},
	"TCC": {
		{
			ID:          1894,
			Slug:        "the-champcoin",
			Name:        "The ChampCoin",
			BinanceUSDT: "TCCUSDT",
		},
	},
	"TREX": {
		{
			ID:          5455,
			Slug:        "trexcoin",
			Name:        "Trexcoin",
			BinanceUSDT: "TREXUSDT",
		},
	},
	"CAKE": {
		{
			ID:          7186,
			Slug:        "pancakeswap",
			Name:        "PancakeSwap",
			BinanceUSDT: "CAKEUSDT",
		},
	},
	"CRPT": {
		{
			ID:          2447,
			Slug:        "crpt",
			Name:        "Crypterium",
			BinanceUSDT: "CRPTUSDT",
		},
	},
	"CREX": {
		{
			ID:          4225,
			Slug:        "crex-token",
			Name:        "Crex Token",
			BinanceUSDT: "CREXUSDT",
		},
	},
	"OPTI": {
		{
			ID:          3101,
			Slug:        "optitoken",
			Name:        "OptiToken",
			BinanceUSDT: "OPTIUSDT",
		},
	},
	"QBIC": {
		{
			ID:          2460,
			Slug:        "qbic",
			Name:        "Qbic",
			BinanceUSDT: "QBICUSDT",
		},
	},
	"yPANDA": {
		{
			ID:          8471,
			Slug:        "yieldpanda-finance",
			Name:        "YieldPanda Finance",
			BinanceUSDT: "yPANDAUSDT",
		},
	},
	"SAVE": {
		{
			ID:   9857,
			Slug: "savetheworld",
			Name: "SaveTheWorld",
		},
		{
			ID:   5789,
			Slug: "savetoken",
			Name: "SaveToken",
		},
	},
	"RVC": {
		{
			ID:          5713,
			Slug:        "ravencoin-classic",
			Name:        "Ravencoin Classic",
			BinanceUSDT: "RVCUSDT",
		},
	},
	"JSHIBA": {
		{
			ID:          9685,
			Slug:        "jomon-shiba",
			Name:        "Jomon Shiba",
			BinanceUSDT: "JSHIBAUSDT",
		},
	},
	"DIGEX": {
		{
			ID:          6680,
			Slug:        "digex",
			Name:        "Digex",
			BinanceUSDT: "DIGEXUSDT",
		},
	},
	"CE": {
		{
			ID:          9329,
			Slug:        "crypto-excellence",
			Name:        "Crypto Excellence",
			BinanceUSDT: "CEUSDT",
		},
	},
	"EXMR": {
		{
			ID:          2990,
			Slug:        "exmr-fdn",
			Name:        "EXMR FDN",
			BinanceUSDT: "EXMRUSDT",
		},
	},
	"3Cs": {
		{
			ID:          5907,
			Slug:        "crypto-cricket-club",
			Name:        "Crypto Cricket Club",
			BinanceUSDT: "3CsUSDT",
		},
	},
	"MEMEX": {
		{
			ID:          10363,
			Slug:        "memex-exchange",
			Name:        "MEMEX",
			BinanceUSDT: "MEMEXUSDT",
		},
	},
	"TSHP": {
		{
			ID:          4280,
			Slug:        "12ships",
			Name:        "12Ships",
			BinanceUSDT: "TSHPUSDT",
		},
	},
	"BNBD": {
		{
			ID:          9667,
			Slug:        "bnb-diamond",
			Name:        "BNB Diamond",
			BinanceUSDT: "BNBDUSDT",
		},
	},
	"PARTY": {
		{
			ID:          6556,
			Slug:        "money-party",
			Name:        "MONEY PARTY",
			BinanceUSDT: "PARTYUSDT",
		},
	},
	"TAI": {
		{
			ID:          6866,
			Slug:        "tai",
			Name:        "TAI",
			BinanceUSDT: "TAIUSDT",
		},
	},
	"SWYFTT": {
		{
			ID:          4588,
			Slug:        "swyft",
			Name:        "SWYFT",
			BinanceUSDT: "SWYFTTUSDT",
		},
	},
	"GENE": {
		{
			ID:   3297,
			Slug: "gene-source-code-chain",
			Name: "Gene Source Code Chain",
		},
		{
			ID:   2720,
			Slug: "parkgene",
			Name: "Parkgene",
		},
	},
	"MOMAT": {
		{
			ID:          9755,
			Slug:        "moma-protocol",
			Name:        "Moma Protocol",
			BinanceUSDT: "MOMATUSDT",
		},
	},
	"DWZ": {
		{
			ID:          8641,
			Slug:        "defi-wizard",
			Name:        "DeFi Wizard",
			BinanceUSDT: "DWZUSDT",
		},
	},
	"VLK": {
		{
			ID:          9576,
			Slug:        "vulkania",
			Name:        "Vulkania",
			BinanceUSDT: "VLKUSDT",
		},
	},
	"TIGER": {
		{
			ID:          8453,
			Slug:        "tigerfinance",
			Name:        "Tigerfinance",
			BinanceUSDT: "TIGERUSDT",
		},
	},
	"XRPBULL": {
		{
			ID:          5412,
			Slug:        "3x-long-xrp-token",
			Name:        "3x Long XRP Token",
			BinanceUSDT: "XRPBULLUSDT",
		},
	},
	"XIOT": {
		{
			ID:          6667,
			Slug:        "xiotri",
			Name:        "Xiotri",
			BinanceUSDT: "XIOTUSDT",
		},
	},
	"MP4": {
		{
			ID:          8597,
			Slug:        "mp4",
			Name:        "MP4",
			BinanceUSDT: "MP4USDT",
		},
	},
	"FBT": {
		{
			ID:          6442,
			Slug:        "fanbi-token",
			Name:        "FANBI TOKEN",
			BinanceUSDT: "FBTUSDT",
		},
	},
	"WINLAMBO": {
		{
			ID:          10472,
			Slug:        "winlambo",
			Name:        "Winlambo",
			BinanceUSDT: "WINLAMBOUSDT",
		},
	},
	"MODIC": {
		{
			ID:          6176,
			Slug:        "modern-investment-coin",
			Name:        "Modern Investment Coin",
			BinanceUSDT: "MODICUSDT",
		},
	},
	"YSKF": {
		{
			ID:          8274,
			Slug:        "yearn-shark-finance",
			Name:        "Yearn Shark Finance",
			BinanceUSDT: "YSKFUSDT",
		},
	},
	"REVO": {
		{
			ID:   9200,
			Slug: "revomon",
			Name: "Revomon",
		},
		{
			ID:   9391,
			Slug: "revo-network",
			Name: "Revo Network",
		},
	},
	"CHFT": {
		{
			ID:          5513,
			Slug:        "crypto-holding-frank-token",
			Name:        "Crypto Holding Frank Token",
			BinanceUSDT: "CHFTUSDT",
		},
	},
	"FTM": {
		{
			ID:   3513,
			Slug: "fantom",
			Name: "Fantom",
		},
		{
			ID:   9990,
			Slug: "fitmin-finance",
			Name: "Fitmin Finance",
		},
	},
	"AERGO": {
		{
			ID:          3637,
			Slug:        "aergo",
			Name:        "Aergo",
			BinanceUSDT: "AERGOUSDT",
		},
	},
	"MEL": {
		{
			ID:   9583,
			Slug: "melalie",
			Name: "Melalie",
		},
		{
			ID:   9795,
			Slug: "caramel-swap",
			Name: "Caramel Swap",
		},
	},
	"MOV": {
		{
			ID:          5958,
			Slug:        "motiv-protocol",
			Name:        "MOTIV Protocol",
			BinanceUSDT: "MOVUSDT",
		},
	},
	"NYZO": {
		{
			ID:          5155,
			Slug:        "nyzo",
			Name:        "Nyzo",
			BinanceUSDT: "NYZOUSDT",
		},
	},
	"UNIUP": {
		{
			ID:          7524,
			Slug:        "uniup",
			Name:        "UNIUP",
			BinanceUSDT: "UNIUPUSDT",
		},
	},
	"XRP": {
		{
			ID:          52,
			Slug:        "xrp",
			Name:        "XRP",
			BinanceUSDT: "XRPUSDT",
		},
	},
	"ETH3S": {
		{
			ID:          5738,
			Slug:        "amun-ether-3x-daily-short",
			Name:        "Amun Ether 3x Daily Short",
			BinanceUSDT: "ETH3SUSDT",
		},
	},
	"XOV": {
		{
			ID:          3097,
			Slug:        "xovbank",
			Name:        "XOVBank",
			BinanceUSDT: "XOVUSDT",
		},
	},
	"SANDG": {
		{
			ID:          1090,
			Slug:        "save-and-gain",
			Name:        "Save and Gain",
			BinanceUSDT: "SANDGUSDT",
		},
	},
	"KRILL": {
		{
			ID:          9459,
			Slug:        "krill",
			Name:        "Krill",
			BinanceUSDT: "KRILLUSDT",
		},
	},
	"JDC": {
		{
			ID:          4929,
			Slug:        "jd-coin",
			Name:        "JD Coin",
			BinanceUSDT: "JDCUSDT",
		},
	},
	"SWPRL": {
		{
			ID:          8550,
			Slug:        "swaprol",
			Name:        "Swaprol",
			BinanceUSDT: "SWPRLUSDT",
		},
	},
	"DEXF": {
		{
			ID:          11027,
			Slug:        "dexfolio",
			Name:        "Dexfolio",
			BinanceUSDT: "DEXFUSDT",
		},
	},
	"RAD": {
		{
			ID:          6843,
			Slug:        "radicle",
			Name:        "Radicle",
			BinanceUSDT: "RADUSDT",
		},
	},
	"XANK": {
		{
			ID:          5659,
			Slug:        "xank",
			Name:        "Xank",
			BinanceUSDT: "XANKUSDT",
		},
	},
	"DDOS": {
		{
			ID:          9024,
			Slug:        "disbalancer",
			Name:        "disBalancer",
			BinanceUSDT: "DDOSUSDT",
		},
	},
	"ESWA": {
		{
			ID:          6425,
			Slug:        "easyswap",
			Name:        "EasySwap",
			BinanceUSDT: "ESWAUSDT",
		},
	},
	"DHX": {
		{
			ID:          6771,
			Slug:        "datahighway",
			Name:        "DataHighway",
			BinanceUSDT: "DHXUSDT",
		},
	},
	"SQ": {
		{
			ID:          7902,
			Slug:        "square-tokenized-stock-ftx",
			Name:        "Square tokenized stock FTX",
			BinanceUSDT: "SQUSDT",
		},
	},
	"XCO": {
		{
			ID:          837,
			Slug:        "x-coin",
			Name:        "X-Coin",
			BinanceUSDT: "XCOUSDT",
		},
	},
	"ACPT": {
		{
			ID:          7116,
			Slug:        "crypto-accept",
			Name:        "Crypto Accept",
			BinanceUSDT: "ACPTUSDT",
		},
	},
	"TIX": {
		{
			ID:          1873,
			Slug:        "blocktix",
			Name:        "Blocktix",
			BinanceUSDT: "TIXUSDT",
		},
	},
	"ETHBN": {
		{
			ID:          5826,
			Slug:        "etherbone",
			Name:        "EtherBone",
			BinanceUSDT: "ETHBNUSDT",
		},
	},
	"ARRR": {
		{
			ID:          3951,
			Slug:        "pirate-chain",
			Name:        "Pirate Chain",
			BinanceUSDT: "ARRRUSDT",
		},
	},
	"VAI": {
		{
			ID:   7824,
			Slug: "vai",
			Name: "Vai",
		},
		{
			ID:   8479,
			Slug: "vaiot",
			Name: "VAIOT",
		},
	},
	"MITH": {
		{
			ID:          2608,
			Slug:        "mithril",
			Name:        "Mithril",
			BinanceUSDT: "MITHUSDT",
		},
	},
	"SENPAI": {
		{
			ID:          9242,
			Slug:        "senpai",
			Name:        "SENPAI",
			BinanceUSDT: "SENPAIUSDT",
		},
	},
	"FCR": {
		{
			ID:          7707,
			Slug:        "fromm-car",
			Name:        "Fromm Car",
			BinanceUSDT: "FCRUSDT",
		},
	},
	"AXN": {
		{
			ID:          7676,
			Slug:        "axion",
			Name:        "Axion",
			BinanceUSDT: "AXNUSDT",
		},
	},
	"GRPL": {
		{
			ID:          8059,
			Slug:        "goldenratioperliquidity",
			Name:        "Golden Ratio Per Liquidity",
			BinanceUSDT: "GRPLUSDT",
		},
	},
	"XCAD": {
		{
			ID:          9868,
			Slug:        "xcad-network",
			Name:        "XCAD Network",
			BinanceUSDT: "XCADUSDT",
		},
	},
	"SYA": {
		{
			ID:          9635,
			Slug:        "save-your-assets",
			Name:        "Save Your Assets",
			BinanceUSDT: "SYAUSDT",
		},
	},
	"SECO": {
		{
			ID:          8220,
			Slug:        "serum-ecosystem-token",
			Name:        "Serum Ecosystem Token",
			BinanceUSDT: "SECOUSDT",
		},
	},
	"BSCPAD": {
		{
			ID:          8660,
			Slug:        "bscpad",
			Name:        "BSCPAD",
			BinanceUSDT: "BSCPADUSDT",
		},
	},
	"BASX": {
		{
			ID:          8513,
			Slug:        "basix",
			Name:        "Basix",
			BinanceUSDT: "BASXUSDT",
		},
	},
	"PRIA": {
		{
			ID:          7466,
			Slug:        "pria",
			Name:        "PRIA",
			BinanceUSDT: "PRIAUSDT",
		},
	},
	"SLS": {
		{
			ID:          1159,
			Slug:        "salus",
			Name:        "SaluS",
			BinanceUSDT: "SLSUSDT",
		},
	},
	"MPAD": {
		{
			ID:          10203,
			Slug:        "moonpad",
			Name:        "Moonpad",
			BinanceUSDT: "MPADUSDT",
		},
	},
	"DMS": {
		{
			ID:          5143,
			Slug:        "documentchain",
			Name:        "Documentchain",
			BinanceUSDT: "DMSUSDT",
		},
	},
	"mAAPL": {
		{
			ID:          8001,
			Slug:        "mirrored-apple",
			Name:        "Mirrored Apple",
			BinanceUSDT: "mAAPLUSDT",
		},
	},
	"FGC": {
		{
			ID:          2870,
			Slug:        "fantasygold",
			Name:        "FantasyGold",
			BinanceUSDT: "FGCUSDT",
		},
	},
	"HIZ": {
		{
			ID:          7554,
			Slug:        "hiz-finance",
			Name:        "Hiz Finance",
			BinanceUSDT: "HIZUSDT",
		},
	},
	"OOKS": {
		{
			ID:          8349,
			Slug:        "onooks",
			Name:        "Onooks",
			BinanceUSDT: "OOKSUSDT",
		},
	},
	"HFS": {
		{
			ID:          10354,
			Slug:        "holder-swap",
			Name:        "Holder Swap",
			BinanceUSDT: "HFSUSDT",
		},
	},
	"KGC": {
		{
			ID:          4291,
			Slug:        "krypton-galaxy-coin",
			Name:        "Krypton Galaxy Coin",
			BinanceUSDT: "KGCUSDT",
		},
	},
	"1AI": {
		{
			ID:          5105,
			Slug:        "1ai-token",
			Name:        "1AI Token",
			BinanceUSDT: "1AIUSDT",
		},
	},
	"NSFW": {
		{
			ID:          9840,
			Slug:        "xxxnifty",
			Name:        "xxxNifty",
			BinanceUSDT: "NSFWUSDT",
		},
	},
	"vLTC": {
		{
			ID:          7964,
			Slug:        "venus-ltc",
			Name:        "Venus LTC",
			BinanceUSDT: "vLTCUSDT",
		},
	},
	"XAG": {
		{
			ID:          6558,
			Slug:        "xrpalike-gene",
			Name:        "Xrpalike Gene",
			BinanceUSDT: "XAGUSDT",
		},
	},
	"ECTE": {
		{
			ID:          3741,
			Slug:        "eurocoin-token",
			Name:        "EurocoinToken",
			BinanceUSDT: "ECTEUSDT",
		},
	},
	"OCE": {
		{
			ID:          3880,
			Slug:        "oceanex-token",
			Name:        "OceanEx Token",
			BinanceUSDT: "OCEUSDT",
		},
	},
	"B2X": {
		{
			ID:          10475,
			Slug:        "b2x",
			Name:        "B2X",
			BinanceUSDT: "B2XUSDT",
		},
	},
	"VOICE": {
		{
			ID:          7518,
			Slug:        "nix-bridge-token",
			Name:        "Voice Token",
			BinanceUSDT: "VOICEUSDT",
		},
	},
	"GVT": {
		{
			ID:          2181,
			Slug:        "genesis-vision",
			Name:        "Genesis Vision",
			BinanceUSDT: "GVTUSDT",
		},
	},
	"MCC": {
		{
			ID:   5595,
			Slug: "multicoincasino",
			Name: "MultiCoinCasino",
		},
		{
			ID:   7704,
			Slug: "mining-core",
			Name: "Mining Core Coin",
		},
	},
	"HVE2": {
		{
			ID:          9027,
			Slug:        "uhive",
			Name:        "Uhive",
			BinanceUSDT: "HVE2USDT",
		},
	},
	"GAZE": {
		{
			ID:          9433,
			Slug:        "gazetv",
			Name:        "GazeTV",
			BinanceUSDT: "GAZEUSDT",
		},
	},
	"RMOON": {
		{
			ID:          9274,
			Slug:        "rocketmoon",
			Name:        "RocketMoon",
			BinanceUSDT: "RMOONUSDT",
		},
	},
	"MRI": {
		{
			ID:          3254,
			Slug:        "mirai",
			Name:        "Mirai",
			BinanceUSDT: "MRIUSDT",
		},
	},
	"HOPE": {
		{
			ID:          9480,
			Slug:        "hope-token",
			Name:        "Hope",
			BinanceUSDT: "HOPEUSDT",
		},
	},
	"PPDEX": {
		{
			ID:          7592,
			Slug:        "pepedex",
			Name:        "Pepedex",
			BinanceUSDT: "PPDEXUSDT",
		},
	},
	"CETH": {
		{
			ID:          5636,
			Slug:        "compound-ether",
			Name:        "Compound Ether",
			BinanceUSDT: "CETHUSDT",
		},
	},
	"STX": {
		{
			ID:   4847,
			Slug: "stacks",
			Name: "Stacks",
		},
		{
			ID:   1861,
			Slug: "stox",
			Name: "Stox",
		},
	},
	"MDT": {
		{
			ID:          2348,
			Slug:        "measurable-data-token",
			Name:        "Measurable Data Token",
			BinanceUSDT: "MDTUSDT",
		},
	},
	"LQT": {
		{
			ID:          10594,
			Slug:        "liquidifty",
			Name:        "Liquidifty",
			BinanceUSDT: "LQTUSDT",
		},
	},
	"EOSDT": {
		{
			ID:          4017,
			Slug:        "eosdt",
			Name:        "EOSDT",
			BinanceUSDT: "EOSDTUSDT",
		},
	},
	"GR": {
		{
			ID:          8035,
			Slug:        "grom",
			Name:        "Grom",
			BinanceUSDT: "GRUSDT",
		},
	},
	"DOTUP": {
		{
			ID:          7003,
			Slug:        "dotup",
			Name:        "DOTUP",
			BinanceUSDT: "DOTUPUSDT",
		},
	},
	"CORN": {
		{
			ID:   7806,
			Slug: "cornichon",
			Name: "Cornichon",
		},
		{
			ID:   6896,
			Slug: "corn",
			Name: "CORN",
		},
		{
			ID:   6818,
			Slug: "popcorn",
			Name: "Popcorn",
		},
	},
	"ZOOT": {
		{
			ID:          10005,
			Slug:        "zoo-token",
			Name:        "Zoo Token",
			BinanceUSDT: "ZOOTUSDT",
		},
	},
	"EMPIRE": {
		{
			ID:          10613,
			Slug:        "empire-token",
			Name:        "Empire Token",
			BinanceUSDT: "EMPIREUSDT",
		},
	},
	"CEEK": {
		{
			ID:          2856,
			Slug:        "ceek-vr",
			Name:        "CEEK VR",
			BinanceUSDT: "CEEKUSDT",
		},
	},
	"onLEXpa": {
		{
			ID:          5097,
			Slug:        "onlexpa",
			Name:        "onLEXpa",
			BinanceUSDT: "onLEXpaUSDT",
		},
	},
	"KKO": {
		{
			ID:          9718,
			Slug:        "kineko",
			Name:        "Kineko",
			BinanceUSDT: "KKOUSDT",
		},
	},
	"XRPBEAR": {
		{
			ID:          5413,
			Slug:        "3x-short-xrp-token",
			Name:        "3x Short XRP Token",
			BinanceUSDT: "XRPBEARUSDT",
		},
	},
	"HUGO": {
		{
			ID:          9282,
			Slug:        "hugo-finance",
			Name:        "Hugo Finance",
			BinanceUSDT: "HUGOUSDT",
		},
	},
	"MATTER": {
		{
			ID:          8603,
			Slug:        "antimatter",
			Name:        "AntiMatter",
			BinanceUSDT: "MATTERUSDT",
		},
	},
	"1337": {
		{
			ID:          9484,
			Slug:        "e1337",
			Name:        "E1337",
			BinanceUSDT: "1337USDT",
		},
	},
	"MTA": {
		{
			ID:          5748,
			Slug:        "meta",
			Name:        "mStable Governance Token: Meta (MTA)",
			BinanceUSDT: "MTAUSDT",
		},
	},
	"imBTC": {
		{
			ID:          5292,
			Slug:        "the-tokenized-bitcoin",
			Name:        "The Tokenized Bitcoin",
			BinanceUSDT: "imBTCUSDT",
		},
	},
	"ABST": {
		{
			ID:          4518,
			Slug:        "abitshadow-token",
			Name:        "Abitshadow Token",
			BinanceUSDT: "ABSTUSDT",
		},
	},
	"SMBSWAP": {
		{
			ID:          7764,
			Slug:        "simbcoin-swap",
			Name:        "Simbcoin Swap",
			BinanceUSDT: "SMBSWAPUSDT",
		},
	},
	"XEQ": {
		{
			ID:          4211,
			Slug:        "equilibria",
			Name:        "Equilibria",
			BinanceUSDT: "XEQUSDT",
		},
	},
	"SAFECOM": {
		{
			ID:          9379,
			Slug:        "safe-community-token",
			Name:        "SAFE Community Token",
			BinanceUSDT: "SAFECOMUSDT",
		},
	},
	"BITC": {
		{
			ID:          3907,
			Slug:        "bitcash",
			Name:        "BitCash",
			BinanceUSDT: "BITCUSDT",
		},
	},
	"CALL": {
		{
			ID:          5019,
			Slug:        "global-crypto-alliance",
			Name:        "Global Crypto Alliance",
			BinanceUSDT: "CALLUSDT",
		},
	},
	"MIL": {
		{
			ID:          11009,
			Slug:        "military-finance",
			Name:        "Military Finance",
			BinanceUSDT: "MILUSDT",
		},
	},
	"AHT": {
		{
			ID:          6641,
			Slug:        "ahatoken",
			Name:        "AhaToken",
			BinanceUSDT: "AHTUSDT",
		},
	},
	"BTCL": {
		{
			ID:          3685,
			Slug:        "btc-lite",
			Name:        "BTC Lite",
			BinanceUSDT: "BTCLUSDT",
		},
	},
	"XDC": {
		{
			ID:          2634,
			Slug:        "xinfin-network",
			Name:        "XinFin Network",
			BinanceUSDT: "XDCUSDT",
		},
	},
	"NIOX": {
		{
			ID:          2151,
			Slug:        "autonio",
			Name:        "Autonio",
			BinanceUSDT: "NIOXUSDT",
		},
	},
	"LKR": {
		{
			ID:          8970,
			Slug:        "polkalokr",
			Name:        "Polkalokr",
			BinanceUSDT: "LKRUSDT",
		},
	},
	"XEM": {
		{
			ID:          873,
			Slug:        "nem",
			Name:        "NEM",
			BinanceUSDT: "XEMUSDT",
		},
	},
	"CHS": {
		{
			ID:   9390,
			Slug: "chainsquare",
			Name: "Chainsquare",
		},
		{
			ID:   8217,
			Slug: "cheeseswap",
			Name: "CheeseSwap",
		},
	},
	"MLA": {
		{
			ID:          8619,
			Slug:        "moola",
			Name:        "Moola",
			BinanceUSDT: "MLAUSDT",
		},
	},
	"GGC": {
		{
			ID:          5090,
			Slug:        "global-game-coin",
			Name:        "Global Game Coin",
			BinanceUSDT: "GGCUSDT",
		},
	},
	"ENT": {
		{
			ID:          1474,
			Slug:        "eternity",
			Name:        "Eternity",
			BinanceUSDT: "ENTUSDT",
		},
	},
	"GLS": {
		{
			ID:          4834,
			Slug:        "golos-blockchain",
			Name:        "Golos Blockchain",
			BinanceUSDT: "GLSUSDT",
		},
	},
	"DEFI5": {
		{
			ID:          8430,
			Slug:        "defi-top-5-tokens-index",
			Name:        "DEFI Top 5 Tokens Index",
			BinanceUSDT: "DEFI5USDT",
		},
	},
	"DADDYDOGE": {
		{
			ID:          10899,
			Slug:        "daddy-doge",
			Name:        "Daddy Doge",
			BinanceUSDT: "DADDYDOGEUSDT",
		},
	},
	"VDL": {
		{
			ID:          4315,
			Slug:        "vidulum",
			Name:        "Vidulum",
			BinanceUSDT: "VDLUSDT",
		},
	},
	"KEY": {
		{
			ID:   2398,
			Slug: "selfkey",
			Name: "Selfkey",
		},
		{
			ID:   9302,
			Slug: "momo-key",
			Name: "MoMo KEY",
		},
		{
			ID:   2713,
			Slug: "key",
			Name: "KEY",
		},
	},
	"MPH": {
		{
			ID:   7742,
			Slug: "88mph",
			Name: "88mph",
		},
		{
			ID:   7217,
			Slug: "morpher",
			Name: "Morpher",
		},
	},
	"PUSD": {
		{
			ID:   5960,
			Slug: "payfrequent-usd-2",
			Name: "PayFrequent USD",
		},
		{
			ID:   7450,
			Slug: "payusd",
			Name: "PayUSD",
		},
	},
	"AM": {
		{
			ID:          10763,
			Slug:        "aston-martin-cognizant-fan-token",
			Name:        "Aston Martin Cognizant Fan Token",
			BinanceUSDT: "AMUSDT",
		},
	},
	"EXCC": {
		{
			ID:          4388,
			Slug:        "exchange-coin",
			Name:        "ExchangeCoin",
			BinanceUSDT: "EXCCUSDT",
		},
	},
	"ETCBEAR": {
		{
			ID:          6099,
			Slug:        "3x-short-ethereum-classic-token",
			Name:        "3X Short Ethereum Classic Token",
			BinanceUSDT: "ETCBEARUSDT",
		},
	},
	"DFT": {
		{
			ID:   1120,
			Slug: "draftcoin",
			Name: "DraftCoin",
		},
		{
			ID:   2878,
			Slug: "digifinextoken",
			Name: "DigiFinexToken",
		},
		{
			ID:   8758,
			Slug: "dfuture",
			Name: "dFuture",
		},
		{
			ID:   6850,
			Slug: "defiat",
			Name: "DeFiat",
		},
	},
	"DRE": {
		{
			ID:          10377,
			Slug:        "dare-token",
			Name:        "Dare Token",
			BinanceUSDT: "DREUSDT",
		},
	},
	"DFR": {
		{
			ID:          8116,
			Slug:        "diffract-finance",
			Name:        "Diffract Finance",
			BinanceUSDT: "DFRUSDT",
		},
	},
	"TKX": {
		{
			ID:          4715,
			Slug:        "tokenize-xchange",
			Name:        "Tokenize Xchange",
			BinanceUSDT: "TKXUSDT",
		},
	},
	"RYIP": {
		{
			ID:          8668,
			Slug:        "ryi-platinum",
			Name:        "RYI Platinum",
			BinanceUSDT: "RYIPUSDT",
		},
	},
	"IBANK": {
		{
			ID:          1515,
			Slug:        "ibank",
			Name:        "iBank",
			BinanceUSDT: "IBANKUSDT",
		},
	},
	"POL": {
		{
			ID:          6297,
			Slug:        "proof-of-liquidity",
			Name:        "Proof Of Liquidity",
			BinanceUSDT: "POLUSDT",
		},
	},
	"LMAO": {
		{
			ID:          10717,
			Slug:        "lucky-meow-token",
			Name:        "Lucky Meow Token",
			BinanceUSDT: "LMAOUSDT",
		},
	},
	"TROP": {
		{
			ID:          8336,
			Slug:        "interop",
			Name:        "Interop",
			BinanceUSDT: "TROPUSDT",
		},
	},
	"YFBTC": {
		{
			ID:          8464,
			Slug:        "yfbitcoin",
			Name:        "YFBitcoin",
			BinanceUSDT: "YFBTCUSDT",
		},
	},
	"BQTX": {
		{
			ID:          3929,
			Slug:        "bqt",
			Name:        "BQT",
			BinanceUSDT: "BQTXUSDT",
		},
	},
	"vLINK": {
		{
			ID:          7975,
			Slug:        "venus-link",
			Name:        "Venus LINK",
			BinanceUSDT: "vLINKUSDT",
		},
	},
	"ETCH": {
		{
			ID:          9742,
			Slug:        "elontech",
			Name:        "ElonTech",
			BinanceUSDT: "ETCHUSDT",
		},
	},
	"LAND": {
		{
			ID:          8752,
			Slug:        "landbox",
			Name:        "Landbox",
			BinanceUSDT: "LANDUSDT",
		},
	},
	"MTV": {
		{
			ID:          3853,
			Slug:        "multivac",
			Name:        "MultiVAC",
			BinanceUSDT: "MTVUSDT",
		},
	},
	"CVC": {
		{
			ID:          1816,
			Slug:        "civic",
			Name:        "Civic",
			BinanceUSDT: "CVCUSDT",
		},
	},
	"vDAI": {
		{
			ID:          8214,
			Slug:        "venus-dai",
			Name:        "Venus DAI",
			BinanceUSDT: "vDAIUSDT",
		},
	},
	"JULD": {
		{
			ID:          8164,
			Slug:        "julswap",
			Name:        "JulSwap",
			BinanceUSDT: "JULDUSDT",
		},
	},
	"CUSDC": {
		{
			ID:          5265,
			Slug:        "compound-usd-coin",
			Name:        "Compound USD Coin",
			BinanceUSDT: "CUSDCUSDT",
		},
	},
	"SAFEHAMSTERS": {
		{
			ID:          10060,
			Slug:        "safehamsters",
			Name:        "SafeHamsters",
			BinanceUSDT: "SAFEHAMSTERSUSDT",
		},
	},
	"DASS": {
		{
			ID:          10801,
			Slug:        "dashsports",
			Name:        "DashSports",
			BinanceUSDT: "DASSUSDT",
		},
	},
	"DMOON": {
		{
			ID:          10125,
			Slug:        "dragonmoon",
			Name:        "DragonMoon",
			BinanceUSDT: "DMOONUSDT",
		},
	},
	"SHIBACASH": {
		{
			ID:          9736,
			Slug:        "shibacash",
			Name:        "ShibaCash",
			BinanceUSDT: "SHIBACASHUSDT",
		},
	},
	"XPT": {
		{
			ID:   5363,
			Slug: "cryptobuyer",
			Name: "Cryptobuyer",
		},
		{
			ID:   7779,
			Slug: "xptoken-io",
			Name: "XPToken.io",
		},
	},
	"XRUNE": {
		{
			ID:          10260,
			Slug:        "thorstarter",
			Name:        "Thorstarter",
			BinanceUSDT: "XRUNEUSDT",
		},
	},
	"DVX": {
		{
			ID:          6104,
			Slug:        "derivex",
			Name:        "Derivex",
			BinanceUSDT: "DVXUSDT",
		},
	},
	"GADOSHI": {
		{
			ID:          8600,
			Slug:        "gadoshi",
			Name:        "Gadoshi",
			BinanceUSDT: "GADOSHIUSDT",
		},
	},
	"L3P": {
		{
			ID:          8845,
			Slug:        "lepricon",
			Name:        "Lepricon",
			BinanceUSDT: "L3PUSDT",
		},
	},
	"CHONK": {
		{
			ID:          7487,
			Slug:        "chonk",
			Name:        "Chonk",
			BinanceUSDT: "CHONKUSDT",
		},
	},
	"ENQ": {
		{
			ID:          4245,
			Slug:        "enecuum",
			Name:        "Enecuum",
			BinanceUSDT: "ENQUSDT",
		},
	},
	"PASTA": {
		{
			ID:          8676,
			Slug:        "pasta-finance",
			Name:        "Pasta Finance",
			BinanceUSDT: "PASTAUSDT",
		},
	},
	"BTCZ": {
		{
			ID:          2041,
			Slug:        "bitcoinz",
			Name:        "BitcoinZ",
			BinanceUSDT: "BTCZUSDT",
		},
	},
	"AZUKI": {
		{
			ID:          7647,
			Slug:        "azuki",
			Name:        "Azuki",
			BinanceUSDT: "AZUKIUSDT",
		},
	},
	"HX": {
		{
			ID:          4895,
			Slug:        "hyperexchange",
			Name:        "HyperExchange",
			BinanceUSDT: "HXUSDT",
		},
	},
	"INU": {
		{
			ID:          10900,
			Slug:        "hachikoinu",
			Name:        "Hachiko Inu",
			BinanceUSDT: "INUUSDT",
		},
	},
	"ITS": {
		{
			ID:          7852,
			Slug:        "iterationsyndicate",
			Name:        "IterationSyndicate",
			BinanceUSDT: "ITSUSDT",
		},
	},
	"DRT": {
		{
			ID:          2070,
			Slug:        "domraider",
			Name:        "DomRaider",
			BinanceUSDT: "DRTUSDT",
		},
	},
	"VRGX": {
		{
			ID:          7621,
			Slug:        "vroomgo",
			Name:        "VROOMGO",
			BinanceUSDT: "VRGXUSDT",
		},
	},
	"FSBT": {
		{
			ID:          2884,
			Slug:        "fsbt-api-token",
			Name:        "FSBT API Token",
			BinanceUSDT: "FSBTUSDT",
		},
	},
	"CORGIB": {
		{
			ID:          10251,
			Slug:        "the-corgi-of-polkabridge",
			Name:        "The Corgi of PolkaBridge",
			BinanceUSDT: "CORGIBUSDT",
		},
	},
	"$PIR": {
		{
			ID:          8013,
			Slug:        "piranhas",
			Name:        "PIRANHAS",
			BinanceUSDT: "$PIRUSDT",
		},
	},
	"SLAM": {
		{
			ID:          9584,
			Slug:        "slam-token",
			Name:        "Slam Token",
			BinanceUSDT: "SLAMUSDT",
		},
	},
	"MTTCOIN": {
		{
			ID:          7835,
			Slug:        "money-of-tomorrow-today",
			Name:        "Money of Tomorrow, Today",
			BinanceUSDT: "MTTCOINUSDT",
		},
	},
	"BCHBEAR": {
		{
			ID:          5466,
			Slug:        "3x-short-bitcoin-cash-token",
			Name:        "3x Short Bitcoin Cash Token",
			BinanceUSDT: "BCHBEARUSDT",
		},
	},
	"UBEX": {
		{
			ID:          3140,
			Slug:        "ubex",
			Name:        "Ubex",
			BinanceUSDT: "UBEXUSDT",
		},
	},
	"PIRATE": {
		{
			ID:          4460,
			Slug:        "piratecash",
			Name:        "PirateCash",
			BinanceUSDT: "PIRATEUSDT",
		},
	},
	"AWG": {
		{
			ID:          5705,
			Slug:        "aurusgold",
			Name:        "AurusGOLD",
			BinanceUSDT: "AWGUSDT",
		},
	},
	"DAWN": {
		{
			ID:          5618,
			Slug:        "dawn-protocol",
			Name:        "Dawn Protocol",
			BinanceUSDT: "DAWNUSDT",
		},
	},
	"NDS": {
		{
			ID:          9068,
			Slug:        "nodeseeds",
			Name:        "Nodeseeds",
			BinanceUSDT: "NDSUSDT",
		},
	},
	"POLT": {
		{
			ID:          9392,
			Slug:        "polkatrain",
			Name:        "Polkatrain",
			BinanceUSDT: "POLTUSDT",
		},
	},
	"ENG": {
		{
			ID:          2044,
			Slug:        "enigma",
			Name:        "Enigma",
			BinanceUSDT: "ENGUSDT",
		},
	},
	"ADC": {
		{
			ID:          948,
			Slug:        "audiocoin",
			Name:        "AudioCoin",
			BinanceUSDT: "ADCUSDT",
		},
	},
	"BUX": {
		{
			ID:   2465,
			Slug: "bux-token",
			Name: "BUX Token",
		},
		{
			ID:   5174,
			Slug: "buxcoin",
			Name: "Buxcoin",
		},
	},
	"ADADOWN": {
		{
			ID:          7014,
			Slug:        "adadown",
			Name:        "ADADOWN",
			BinanceUSDT: "ADADOWNUSDT",
		},
	},
	"DEXE": {
		{
			ID:          7326,
			Slug:        "dexe",
			Name:        "DeXe",
			BinanceUSDT: "DEXEUSDT",
		},
	},
	"HMNG": {
		{
			ID:          9737,
			Slug:        "hummingbird-finance",
			Name:        "HummingBird Finance",
			BinanceUSDT: "HMNGUSDT",
		},
	},
	"DGVC": {
		{
			ID:          6689,
			Slug:        "degenvc",
			Name:        "DegenVC",
			BinanceUSDT: "DGVCUSDT",
		},
	},
	"JAR": {
		{
			ID:          4034,
			Slug:        "jarvis",
			Name:        "Jarvis+",
			BinanceUSDT: "JARUSDT",
		},
	},
	"CSPN": {
		{
			ID:          3894,
			Slug:        "crypto-sports",
			Name:        "Crypto Sports",
			BinanceUSDT: "CSPNUSDT",
		},
	},
	"HQX": {
		{
			ID:          2564,
			Slug:        "hoqu",
			Name:        "HOQU",
			BinanceUSDT: "HQXUSDT",
		},
	},
	"PNL": {
		{
			ID:          9605,
			Slug:        "truepnl",
			Name:        "TruePNL",
			BinanceUSDT: "PNLUSDT",
		},
	},
	"DMX": {
		{
			ID:          8117,
			Slug:        "dymmax",
			Name:        "Dymmax",
			BinanceUSDT: "DMXUSDT",
		},
	},
	"ALN": {
		{
			ID:          5544,
			Slug:        "aluna-social",
			Name:        "Aluna.Social",
			BinanceUSDT: "ALNUSDT",
		},
	},
	"WHALE": {
		{
			ID:          6679,
			Slug:        "whale",
			Name:        "WHALE",
			BinanceUSDT: "WHALEUSDT",
		},
	},
	"SAFE": {
		{
			ID:   3918,
			Slug: "safe",
			Name: "Safe",
		},
		{
			ID:   3799,
			Slug: "safecoin",
			Name: "SafeCoin",
		},
	},
	"MARK": {
		{
			ID:          7765,
			Slug:        "benchmark-protocol",
			Name:        "Benchmark Protocol",
			BinanceUSDT: "MARKUSDT",
		},
	},
	"TFT": {
		{
			ID:   6500,
			Slug: "threefold",
			Name: "ThreeFold",
		},
		{
			ID:   8391,
			Slug: "the-famous-token",
			Name: "The Famous Token",
		},
	},
	"BTX": {
		{
			ID:          1654,
			Slug:        "bitcore",
			Name:        "BitCore",
			BinanceUSDT: "BTXUSDT",
		},
	},
	"STARINU": {
		{
			ID:          11014,
			Slug:        "starship-inu",
			Name:        "Starship Inu",
			BinanceUSDT: "STARINUUSDT",
		},
	},
	"AAPL": {
		{
			ID:   7924,
			Slug: "apple-tokenized-stock-bittrex",
			Name: "Apple tokenized stock Bittrex",
		},
		{
			ID:   7894,
			Slug: "apple-tokenized-stock-ftx",
			Name: "Apple tokenized stock FTX",
		},
	},
	"TOK": {
		{
			ID:          3692,
			Slug:        "tokok",
			Name:        "TOKOK",
			BinanceUSDT: "TOKUSDT",
		},
	},
	"OT": {
		{
			ID:          8404,
			Slug:        "option-token",
			Name:        "Option Token",
			BinanceUSDT: "OTUSDT",
		},
	},
	"LNT": {
		{
			ID:          7036,
			Slug:        "lottonation",
			Name:        "Lottonation",
			BinanceUSDT: "LNTUSDT",
		},
	},
	"GLOB": {
		{
			ID:          5076,
			Slug:        "global-reserve-system",
			Name:        "Global Reserve System",
			BinanceUSDT: "GLOBUSDT",
		},
	},
	"WAULT": {
		{
			ID:          8588,
			Slug:        "wault-finance",
			Name:        "Wault Finance (OLD)",
			BinanceUSDT: "WAULTUSDT",
		},
	},
	"AMB": {
		{
			ID:          2081,
			Slug:        "amber",
			Name:        "Ambrosus",
			BinanceUSDT: "AMBUSDT",
		},
	},
	"PAYX": {
		{
			ID:          2191,
			Slug:        "paypex",
			Name:        "Paypex",
			BinanceUSDT: "PAYXUSDT",
		},
	},
	"BRC": {
		{
			ID:          3653,
			Slug:        "baer-chain",
			Name:        "Baer Chain",
			BinanceUSDT: "BRCUSDT",
		},
	},
	"PIN": {
		{
			ID:          64,
			Slug:        "public-index-network",
			Name:        "Public Index Network",
			BinanceUSDT: "PINUSDT",
		},
	},
	"COCO": {
		{
			ID:          9790,
			Slug:        "coco-swap",
			Name:        "Coco Swap",
			BinanceUSDT: "COCOUSDT",
		},
	},
	"BILL": {
		{
			ID:          10568,
			Slug:        "bill-hwang-finance",
			Name:        "Bill Hwang Finance",
			BinanceUSDT: "BILLUSDT",
		},
	},
	"SPE": {
		{
			ID:          9303,
			Slug:        "save-planet-earth",
			Name:        "Save Planet Earth",
			BinanceUSDT: "SPEUSDT",
		},
	},
	"CHIMOM": {
		{
			ID:          10175,
			Slug:        "chihua-chimon",
			Name:        "Chihua Token",
			BinanceUSDT: "CHIMOMUSDT",
		},
	},
	"BABY": {
		{
			ID:   10334,
			Slug: "babyswap",
			Name: "BabySwap",
		},
		{
			ID:   9414,
			Slug: "babytoken",
			Name: "Babytoken",
		},
	},
	"aEth": {
		{
			ID:          8100,
			Slug:        "ankreth",
			Name:        "ankrETH",
			BinanceUSDT: "aEthUSDT",
		},
	},
	"Whale": {
		{
			ID:          10693,
			Slug:        "whale-fall",
			Name:        "Whale Fall",
			BinanceUSDT: "WhaleUSDT",
		},
	},
	"ZHEGIC": {
		{
			ID:          7601,
			Slug:        "zhegic",
			Name:        "zHEGIC",
			BinanceUSDT: "ZHEGICUSDT",
		},
	},
	"PCNT": {
		{
			ID:          8704,
			Slug:        "playcent",
			Name:        "Playcent",
			BinanceUSDT: "PCNTUSDT",
		},
	},
	"RFI": {
		{
			ID:          7747,
			Slug:        "reflect-finance",
			Name:        "reflect.finance",
			BinanceUSDT: "RFIUSDT",
		},
	},
	"UNIDX": {
		{
			ID:          8232,
			Slug:        "unidex",
			Name:        "UniDex",
			BinanceUSDT: "UNIDXUSDT",
		},
	},
	"TEND": {
		{
			ID:          5971,
			Slug:        "tendies",
			Name:        "Tendies",
			BinanceUSDT: "TENDUSDT",
		},
	},
	"JOYS": {
		{
			ID:          5278,
			Slug:        "joys-digital",
			Name:        "Joys Digital",
			BinanceUSDT: "JOYSUSDT",
		},
	},
	"SWAP": {
		{
			ID:          5829,
			Slug:        "trustswap",
			Name:        "TrustSwap",
			BinanceUSDT: "SWAPUSDT",
		},
	},
	"YFTE": {
		{
			ID:          8269,
			Slug:        "yftether",
			Name:        "YFTether",
			BinanceUSDT: "YFTEUSDT",
		},
	},
	"DIAMND": {
		{
			ID:          11011,
			Slug:        "projekt-diamond",
			Name:        "Projekt Diamond",
			BinanceUSDT: "DIAMNDUSDT",
		},
	},
	"ACS": {
		{
			ID:          7844,
			Slug:        "acryptos",
			Name:        "ACryptoS",
			BinanceUSDT: "ACSUSDT",
		},
	},
	"DAFT": {
		{
			ID:          8913,
			Slug:        "daftcoin",
			Name:        "DaftCoin",
			BinanceUSDT: "DAFTUSDT",
		},
	},
	"SIGN": {
		{
			ID:          3868,
			Slug:        "signature-chain",
			Name:        "Signature Chain",
			BinanceUSDT: "SIGNUSDT",
		},
	},
	"ETHYS": {
		{
			ID:          7803,
			Slug:        "ethereum-stake",
			Name:        "Ethereum Stake",
			BinanceUSDT: "ETHYSUSDT",
		},
	},
	"mTSLA": {
		{
			ID:          8004,
			Slug:        "mirrored-tesla",
			Name:        "Mirrored Tesla",
			BinanceUSDT: "mTSLAUSDT",
		},
	},
	"ISIKC": {
		{
			ID:          5468,
			Slug:        "isiklar-coin",
			Name:        "Isiklar Coin",
			BinanceUSDT: "ISIKCUSDT",
		},
	},
	"MCN": {
		{
			ID:          7482,
			Slug:        "mcnetwork",
			Name:        "McNetworkDefi",
			BinanceUSDT: "MCNUSDT",
		},
	},
	"PACT": {
		{
			ID:          10087,
			Slug:        "pact-community-token",
			Name:        "PACT community token",
			BinanceUSDT: "PACTUSDT",
		},
	},
	"HONOR": {
		{
			ID:          10620,
			Slug:        "farm-hero",
			Name:        "FarmHero",
			BinanceUSDT: "HONORUSDT",
		},
	},
	"vSXP": {
		{
			ID:          7952,
			Slug:        "vsxp",
			Name:        "Venus SXP",
			BinanceUSDT: "vSXPUSDT",
		},
	},
	"GDT": {
		{
			ID:   9812,
			Slug: "gorilla-diamond",
			Name: "Gorilla Diamond",
		},
		{
			ID:   9188,
			Slug: "globe-derivative-exchange",
			Name: "Globe Derivative Exchange",
		},
	},
	"SDFI": {
		{
			ID:          9064,
			Slug:        "sting-defi",
			Name:        "Sting Defi",
			BinanceUSDT: "SDFIUSDT",
		},
	},
	"ETHRSI6040": {
		{
			ID:          6143,
			Slug:        "eth-rsi-60-40-crossover-set",
			Name:        "ETH RSI 60/40 Crossover Set",
			BinanceUSDT: "ETHRSI6040USDT",
		},
	},
	"MNC": {
		{
			ID:   3774,
			Slug: "maincoin",
			Name: "Maincoin",
		},
		{
			ID:   3755,
			Slug: "moneynet",
			Name: "Moneynet",
		},
	},
	"HIBS": {
		{
			ID:          6470,
			Slug:        "hiblocks",
			Name:        "Hiblocks",
			BinanceUSDT: "HIBSUSDT",
		},
	},
	"AUNIT": {
		{
			ID:          3747,
			Slug:        "aunite",
			Name:        "Aunite",
			BinanceUSDT: "AUNITUSDT",
		},
	},
	"BAZT": {
		{
			ID:          5050,
			Slug:        "baz-token",
			Name:        "Baz Token",
			BinanceUSDT: "BAZTUSDT",
		},
	},
	"YFFS": {
		{
			ID:          6863,
			Slug:        "yffs",
			Name:        "YFFS Finance",
			BinanceUSDT: "YFFSUSDT",
		},
	},
	"HNZO": {
		{
			ID:          9765,
			Slug:        "hanzo-inu",
			Name:        "Hanzo Inu",
			BinanceUSDT: "HNZOUSDT",
		},
	},
	"TOSC": {
		{
			ID:          3707,
			Slug:        "t-os",
			Name:        "T.OS",
			BinanceUSDT: "TOSCUSDT",
		},
	},
	"HEZ": {
		{
			ID:          7424,
			Slug:        "hermez-network",
			Name:        "Hermez Network",
			BinanceUSDT: "HEZUSDT",
		},
	},
	"AQUAGOAT": {
		{
			ID:          9370,
			Slug:        "aquagoat-finance",
			Name:        "AquaGoat.Finance",
			BinanceUSDT: "AQUAGOATUSDT",
		},
	},
	"INJ": {
		{
			ID:          7226,
			Slug:        "injective-protocol",
			Name:        "Injective Protocol",
			BinanceUSDT: "INJUSDT",
		},
	},
	"TFI": {
		{
			ID:          10585,
			Slug:        "trustfi-network",
			Name:        "TrustFi Network",
			BinanceUSDT: "TFIUSDT",
		},
	},
	"PASS": {
		{
			ID:          3141,
			Slug:        "blockpass",
			Name:        "Blockpass",
			BinanceUSDT: "PASSUSDT",
		},
	},
	"DEC": {
		{
			ID:   5835,
			Slug: "decentr",
			Name: "Decentr",
		},
		{
			ID:   6264,
			Slug: "dark-energy-crystals",
			Name: "Dark Energy Crystals",
		},
		{
			ID:   6420,
			Slug: "distributed-energy-coin",
			Name: "Distributed Energy Coin",
		},
	},
	"SHDW": {
		{
			ID:          1878,
			Slug:        "shadow-token",
			Name:        "Shadow Token",
			BinanceUSDT: "SHDWUSDT",
		},
	},
	"IPX": {
		{
			ID:          5103,
			Slug:        "tachyon-protocol",
			Name:        "Tachyon Protocol",
			BinanceUSDT: "IPXUSDT",
		},
	},
	"SAFEARN": {
		{
			ID:          10988,
			Slug:        "safe-earn",
			Name:        "Safe Earn",
			BinanceUSDT: "SAFEARNUSDT",
		},
	},
	"YFBT": {
		{
			ID:          7237,
			Slug:        "yearn-finance-bit",
			Name:        "Yearn Finance Bit",
			BinanceUSDT: "YFBTUSDT",
		},
	},
	"THN": {
		{
			ID:          10805,
			Slug:        "throne",
			Name:        "Throne",
			BinanceUSDT: "THNUSDT",
		},
	},
	"GLITCHY": {
		{
			ID:          10595,
			Slug:        "glitchy",
			Name:        "Glitchy",
			BinanceUSDT: "GLITCHYUSDT",
		},
	},
	"TAVITT": {
		{
			ID:          7397,
			Slug:        "tavittcoin",
			Name:        "Tavittcoin",
			BinanceUSDT: "TAVITTUSDT",
		},
	},
	"LEASH": {
		{
			ID:          9286,
			Slug:        "doge-killer",
			Name:        "Doge Killer",
			BinanceUSDT: "LEASHUSDT",
		},
	},
	"MILK": {
		{
			ID:          9141,
			Slug:        "milk-token",
			Name:        "Milk Token",
			BinanceUSDT: "MILKUSDT",
		},
	},
	"XPOSE": {
		{
			ID:          9704,
			Slug:        "xpose",
			Name:        "Xpose Protocol",
			BinanceUSDT: "XPOSEUSDT",
		},
	},
	"SPDR": {
		{
			ID:          8002,
			Slug:        "spiderdao",
			Name:        "SpiderDAO",
			BinanceUSDT: "SPDRUSDT",
		},
	},
	"0xBTC": {
		{
			ID:          2837,
			Slug:        "0xbtc",
			Name:        "0xBitcoin",
			BinanceUSDT: "0xBTCUSDT",
		},
	},
	"WISE": {
		{
			ID:          8167,
			Slug:        "wise",
			Name:        "Wise Token",
			BinanceUSDT: "WISEUSDT",
		},
	},
	"NADA": {
		{
			ID:          10623,
			Slug:        "nothing",
			Name:        "Nothing",
			BinanceUSDT: "NADAUSDT",
		},
	},
	"MAG": {
		{
			ID:          2434,
			Slug:        "maggie",
			Name:        "Maggie",
			BinanceUSDT: "MAGUSDT",
		},
	},
	"XSPC": {
		{
			ID:          4013,
			Slug:        "spectre-security-coin",
			Name:        "SpectreSecurityCoin",
			BinanceUSDT: "XSPCUSDT",
		},
	},
	"SHILLING": {
		{
			ID:          10209,
			Slug:        "shilling-token",
			Name:        "Shilling Token",
			BinanceUSDT: "SHILLINGUSDT",
		},
	},
	"CARBON": {
		{
			ID:          502,
			Slug:        "carboncoin",
			Name:        "Carboncoin",
			BinanceUSDT: "CARBONUSDT",
		},
	},
	"WTF": {
		{
			ID:          6848,
			Slug:        "walnut-finance",
			Name:        "Walnut.finance",
			BinanceUSDT: "WTFUSDT",
		},
	},
	"INFTEE": {
		{
			ID:          11040,
			Slug:        "infinitee-finance",
			Name:        "Infinitee Finance",
			BinanceUSDT: "INFTEEUSDT",
		},
	},
	"HUE": {
		{
			ID:          8967,
			Slug:        "hue",
			Name:        "Hue",
			BinanceUSDT: "HUEUSDT",
		},
	},
	"CRING": {
		{
			ID:          9243,
			Slug:        "darwinia-crab-network",
			Name:        "Darwinia Crab Network",
			BinanceUSDT: "CRINGUSDT",
		},
	},
	"NOW": {
		{
			ID:          3893,
			Slug:        "now-token",
			Name:        "ChangeNOW Token",
			BinanceUSDT: "NOWUSDT",
		},
	},
	"ZUM": {
		{
			ID:   3652,
			Slug: "zumcoin",
			Name: "ZumCoin",
		},
		{
			ID:   4826,
			Slug: "zum-token",
			Name: "ZUM TOKEN",
		},
	},
	"TLM": {
		{
			ID:          9119,
			Slug:        "alien-worlds",
			Name:        "Alien Worlds",
			BinanceUSDT: "TLMUSDT",
		},
	},
	"TCH": {
		{
			ID:   3806,
			Slug: "tigercash",
			Name: "TigerCash",
		},
		{
			ID:   3056,
			Slug: "thore-cash",
			Name: "Thore Cash",
		},
		{
			ID:   4560,
			Slug: "tchain",
			Name: "Tchain",
		},
	},
	"NAFT": {
		{
			ID:          9828,
			Slug:        "nafter",
			Name:        "Nafter",
			BinanceUSDT: "NAFTUSDT",
		},
	},
	"GME": {
		{
			ID:          8342,
			Slug:        "gamestop-tokenized-stock-ftx",
			Name:        "GameStop tokenized stock FTX",
			BinanceUSDT: "GMEUSDT",
		},
	},
	"BKS": {
		{
			ID:          5685,
			Slug:        "barkis-network",
			Name:        "Barkis Network",
			BinanceUSDT: "BKSUSDT",
		},
	},
	"DIME": {
		{
			ID:          90,
			Slug:        "dimecoin",
			Name:        "Dimecoin",
			BinanceUSDT: "DIMEUSDT",
		},
	},
	"ZSC": {
		{
			ID:          2047,
			Slug:        "zeusshield",
			Name:        "Zeusshield",
			BinanceUSDT: "ZSCUSDT",
		},
	},
	"DDD": {
		{
			ID:          2428,
			Slug:        "scryinfo",
			Name:        "Scry.info",
			BinanceUSDT: "DDDUSDT",
		},
	},
	"PAYT": {
		{
			ID:          7425,
			Slug:        "payaccept",
			Name:        "PayAccept",
			BinanceUSDT: "PAYTUSDT",
		},
	},
	"NABOX": {
		{
			ID:          9653,
			Slug:        "nabox",
			Name:        "Nabox",
			BinanceUSDT: "NABOXUSDT",
		},
	},
	"DATX": {
		{
			ID:          2567,
			Slug:        "datx",
			Name:        "DATx",
			BinanceUSDT: "DATXUSDT",
		},
	},
	"BIT": {
		{
			ID:   2977,
			Slug: "bitrewards",
			Name: "BitRewards",
		},
		{
			ID:   3478,
			Slug: "bitmoney",
			Name: "BitMoney",
		},
	},
	"HUM": {
		{
			ID:          3600,
			Slug:        "humanscape",
			Name:        "Humanscape",
			BinanceUSDT: "HUMUSDT",
		},
	},
	"NSR": {
		{
			ID:          699,
			Slug:        "nushares",
			Name:        "NuShares",
			BinanceUSDT: "NSRUSDT",
		},
	},
	"CQT": {
		{
			ID:          7411,
			Slug:        "covalent",
			Name:        "Covalent",
			BinanceUSDT: "CQTUSDT",
		},
	},
	"SUTER": {
		{
			ID:          4841,
			Slug:        "suterusu",
			Name:        "suterusu",
			BinanceUSDT: "SUTERUSDT",
		},
	},
	"HAI": {
		{
			ID:          5583,
			Slug:        "hackenai",
			Name:        "Hacken Token",
			BinanceUSDT: "HAIUSDT",
		},
	},
	"VJC": {
		{
			ID:          3923,
			Slug:        "venjocoin",
			Name:        "VENJOCOIN",
			BinanceUSDT: "VJCUSDT",
		},
	},
	"MEDIBIT": {
		{
			ID:          3582,
			Slug:        "medibit",
			Name:        "MediBit",
			BinanceUSDT: "MEDIBITUSDT",
		},
	},
	"MOD": {
		{
			ID:          8494,
			Slug:        "modefi",
			Name:        "Modefi",
			BinanceUSDT: "MODUSDT",
		},
	},
	"PLUTO": {
		{
			ID:          9715,
			Slug:        "plutopepe",
			Name:        "PlutoPepe",
			BinanceUSDT: "PLUTOUSDT",
		},
	},
	"B26": {
		{
			ID:          8803,
			Slug:        "b26-finance",
			Name:        "B26 Finance",
			BinanceUSDT: "B26USDT",
		},
	},
	"FETCH": {
		{
			ID:          10112,
			Slug:        "moonretriever",
			Name:        "MoonRetriever",
			BinanceUSDT: "FETCHUSDT",
		},
	},
	"CHADS": {
		{
			ID:          7031,
			Slug:        "chads-vc",
			Name:        "CHADS VC",
			BinanceUSDT: "CHADSUSDT",
		},
	},
	"PRDX": {
		{
			ID:          6971,
			Slug:        "predix-network",
			Name:        "Predix Network",
			BinanceUSDT: "PRDXUSDT",
		},
	},
	"ATT": {
		{
			ID:          5600,
			Slug:        "attila",
			Name:        "Attila",
			BinanceUSDT: "ATTUSDT",
		},
	},
	"KP3RB": {
		{
			ID:          8134,
			Slug:        "keep3r-bsc-network",
			Name:        "Keep3r BSC Network",
			BinanceUSDT: "KP3RBUSDT",
		},
	},
	"EPAN": {
		{
			ID:          7749,
			Slug:        "paypolitan-token",
			Name:        "Paypolitan Token",
			BinanceUSDT: "EPANUSDT",
		},
	},
	"PROPHET": {
		{
			ID:          7773,
			Slug:        "prophet",
			Name:        "Prophet",
			BinanceUSDT: "PROPHETUSDT",
		},
	},
	"GOZ": {
		{
			ID:          9507,
			Slug:        "goztepe-sk-fantoken",
			Name:        "Göztepe S.K. Fan Token",
			BinanceUSDT: "GOZUSDT",
		},
	},
	"PAPADOGE": {
		{
			ID:          11045,
			Slug:        "papa-doge-coin",
			Name:        "Papa Doge Coin",
			BinanceUSDT: "PAPADOGEUSDT",
		},
	},
	"MKCY": {
		{
			ID:          7113,
			Slug:        "markaccy",
			Name:        "Markaccy",
			BinanceUSDT: "MKCYUSDT",
		},
	},
	"CHM": {
		{
			ID:          7589,
			Slug:        "cryptochrome",
			Name:        "Cryptochrome",
			BinanceUSDT: "CHMUSDT",
		},
	},
	"CNR": {
		{
			ID:          10555,
			Slug:        "canary",
			Name:        "Canary",
			BinanceUSDT: "CNRUSDT",
		},
	},
	"SHIH": {
		{
			ID:          9712,
			Slug:        "shih-tzu",
			Name:        "Shih Tzu",
			BinanceUSDT: "SHIHUSDT",
		},
	},
	"ELET": {
		{
			ID:          3931,
			Slug:        "elementeum",
			Name:        "Elementeum",
			BinanceUSDT: "ELETUSDT",
		},
	},
	"FRM": {
		{
			ID:          4228,
			Slug:        "ferrum-network",
			Name:        "Ferrum Network",
			BinanceUSDT: "FRMUSDT",
		},
	},
	"SBANK": {
		{
			ID:   11025,
			Slug: "safebank-eth",
			Name: "SafeBank ETH",
		},
		{
			ID:   10975,
			Slug: "safebank-bsc",
			Name: "SafeBank BSC",
		},
	},
	"USDL": {
		{
			ID:          7364,
			Slug:        "usdl",
			Name:        "USDL",
			BinanceUSDT: "USDLUSDT",
		},
	},
	"FLIXX": {
		{
			ID:          2231,
			Slug:        "flixxo",
			Name:        "Flixxo",
			BinanceUSDT: "FLIXXUSDT",
		},
	},
	"WHL": {
		{
			ID:          8560,
			Slug:        "whaleroom",
			Name:        "WhaleRoom",
			BinanceUSDT: "WHLUSDT",
		},
	},
	"MCONTENT": {
		{
			ID:          10954,
			Slug:        "mcontent",
			Name:        "MContent",
			BinanceUSDT: "MCONTENTUSDT",
		},
	},
	"WISB": {
		{
			ID:          10000,
			Slug:        "wise-token",
			Name:        "Wise Token",
			BinanceUSDT: "WISBUSDT",
		},
	},
	"DAV": {
		{
			ID:          3220,
			Slug:        "dav-coin",
			Name:        "DAV Coin",
			BinanceUSDT: "DAVUSDT",
		},
	},
	"DFIO": {
		{
			ID:          6672,
			Slug:        "defi-omega",
			Name:        "DeFi Omega",
			BinanceUSDT: "DFIOUSDT",
		},
	},
	"GRC": {
		{
			ID:          833,
			Slug:        "gridcoin",
			Name:        "Gridcoin",
			BinanceUSDT: "GRCUSDT",
		},
	},
	"NOLE": {
		{
			ID:          5849,
			Slug:        "nolecoin",
			Name:        "NoleCoin",
			BinanceUSDT: "NOLEUSDT",
		},
	},
	"HNC": {
		{
			ID:          1004,
			Slug:        "hnccoin",
			Name:        "HNC COIN",
			BinanceUSDT: "HNCUSDT",
		},
	},
	"SWISS": {
		{
			ID:          7714,
			Slug:        "swiss-finance",
			Name:        "swiss.finance",
			BinanceUSDT: "SWISSUSDT",
		},
	},
	"IDL": {
		{
			ID:          8142,
			Slug:        "idl-token",
			Name:        "IDL Token",
			BinanceUSDT: "IDLUSDT",
		},
	},
	"PFID": {
		{
			ID:          7043,
			Slug:        "pofid-dao",
			Name:        "Pofid Dao",
			BinanceUSDT: "PFIDUSDT",
		},
	},
	"TBAKE": {
		{
			ID:          9796,
			Slug:        "bakery-tools",
			Name:        "Bakery Tools",
			BinanceUSDT: "TBAKEUSDT",
		},
	},
	"DAO1": {
		{
			ID:          10384,
			Slug:        "dao1",
			Name:        "DAO1",
			BinanceUSDT: "DAO1USDT",
		},
	},
	"EXNT": {
		{
			ID:          6882,
			Slug:        "exnetwork-token",
			Name:        "ExNetwork Token",
			BinanceUSDT: "EXNTUSDT",
		},
	},
	"VIDYA": {
		{
			ID:          6709,
			Slug:        "vidya",
			Name:        "Vidya",
			BinanceUSDT: "VIDYAUSDT",
		},
	},
	"PNK": {
		{
			ID:          3581,
			Slug:        "kleros",
			Name:        "Kleros",
			BinanceUSDT: "PNKUSDT",
		},
	},
	"SCL": {
		{
			ID:          1969,
			Slug:        "sociall",
			Name:        "Sociall",
			BinanceUSDT: "SCLUSDT",
		},
	},
	"DML": {
		{
			ID:          2679,
			Slug:        "decentralized-machine-learning",
			Name:        "Decentralized Machine Learning",
			BinanceUSDT: "DMLUSDT",
		},
	},
	"BSY": {
		{
			ID:          5782,
			Slug:        "bestay",
			Name:        "Bestay",
			BinanceUSDT: "BSYUSDT",
		},
	},
	"NEBL": {
		{
			ID:          1955,
			Slug:        "neblio",
			Name:        "Neblio",
			BinanceUSDT: "NEBLUSDT",
		},
	},
	"L2": {
		{
			ID:          7772,
			Slug:        "leverj-gluon",
			Name:        "Leverj Gluon",
			BinanceUSDT: "L2USDT",
		},
	},
	"POLY": {
		{
			ID:          2496,
			Slug:        "polymath-network",
			Name:        "Polymath",
			BinanceUSDT: "POLYUSDT",
		},
	},
	"HAVY": {
		{
			ID:          3265,
			Slug:        "havy",
			Name:        "Havy",
			BinanceUSDT: "HAVYUSDT",
		},
	},
	"INFI": {
		{
			ID:          8305,
			Slug:        "insured-finance",
			Name:        "Insured Finance",
			BinanceUSDT: "INFIUSDT",
		},
	},
	"FSHN": {
		{
			ID:          6446,
			Slug:        "fashion-coin",
			Name:        "Fashion Coin",
			BinanceUSDT: "FSHNUSDT",
		},
	},
	"STARL": {
		{
			ID:          10821,
			Slug:        "star-link",
			Name:        "StarLink",
			BinanceUSDT: "STARLUSDT",
		},
	},
	"BCVT": {
		{
			ID:          8782,
			Slug:        "bitcoinvend",
			Name:        "BitcoinVend",
			BinanceUSDT: "BCVTUSDT",
		},
	},
	"ARIA20": {
		{
			ID:          8276,
			Slug:        "arianee-protocol",
			Name:        "Arianee",
			BinanceUSDT: "ARIA20USDT",
		},
	},
	"BAG": {
		{
			ID:   8273,
			Slug: "bag",
			Name: "Basis Gold",
		},
		{
			ID:   9322,
			Slug: "bondappetit-governance-token",
			Name: "BondAppétit Governance Token",
		},
	},
	"LOWB": {
		{
			ID:          9673,
			Slug:        "loser-coin",
			Name:        "Loser Coin",
			BinanceUSDT: "LOWBUSDT",
		},
	},
	"BAK": {
		{
			ID:          6217,
			Slug:        "tokenbacon",
			Name:        "TokenBacon",
			BinanceUSDT: "BAKUSDT",
		},
	},
	"X42": {
		{
			ID:          4097,
			Slug:        "x42-protocol",
			Name:        "x42 Protocol",
			BinanceUSDT: "X42USDT",
		},
	},
	"ZEN": {
		{
			ID:          1698,
			Slug:        "horizen",
			Name:        "Horizen",
			BinanceUSDT: "ZENUSDT",
		},
	},
	"FSC": {
		{
			ID:          5343,
			Slug:        "five-star-coin",
			Name:        "Five Star Coin",
			BinanceUSDT: "FSCUSDT",
		},
	},
	"REPO": {
		{
			ID:          2829,
			Slug:        "repo",
			Name:        "REPO",
			BinanceUSDT: "REPOUSDT",
		},
	},
	"FBN": {
		{
			ID:          3468,
			Slug:        "fivebalance",
			Name:        "Fivebalance",
			BinanceUSDT: "FBNUSDT",
		},
	},
	"FCT": {
		{
			ID:   4953,
			Slug: "firmachain",
			Name: "FirmaChain",
		},
		{
			ID:   1087,
			Slug: "factom",
			Name: "Factom",
		},
	},
	"PKF": {
		{
			ID:          8617,
			Slug:        "polkafoundry",
			Name:        "PolkaFoundry",
			BinanceUSDT: "PKFUSDT",
		},
	},
	"TBB": {
		{
			ID:          8090,
			Slug:        "trade-butler-bot",
			Name:        "Trade Butler Bot",
			BinanceUSDT: "TBBUSDT",
		},
	},
	"MXC": {
		{
			ID:          3628,
			Slug:        "mxc",
			Name:        "MXC",
			BinanceUSDT: "MXCUSDT",
		},
	},
	"KIMJ": {
		{
			ID:          10164,
			Slug:        "kimjongmoon",
			Name:        "KimJongMoon",
			BinanceUSDT: "KIMJUSDT",
		},
	},
	"NANODOGE": {
		{
			ID:          10916,
			Slug:        "nano-doge-token",
			Name:        "Nano Doge Token",
			BinanceUSDT: "NANODOGEUSDT",
		},
	},
	"DCR": {
		{
			ID:          1168,
			Slug:        "decred",
			Name:        "Decred",
			BinanceUSDT: "DCRUSDT",
		},
	},
	"REX": {
		{
			ID:          1961,
			Slug:        "imbrex",
			Name:        "imbrex",
			BinanceUSDT: "REXUSDT",
		},
	},
	"PDF": {
		{
			ID:          6737,
			Slug:        "port-of-defi-network",
			Name:        "Port of DeFi Network",
			BinanceUSDT: "PDFUSDT",
		},
	},
	"C2": {
		{
			ID:          367,
			Slug:        "coin2-1",
			Name:        "Coin2.1",
			BinanceUSDT: "C2USDT",
		},
	},
	"TREES": {
		{
			ID:          10512,
			Slug:        "safetrees",
			Name:        "SAFETREES",
			BinanceUSDT: "TREESUSDT",
		},
	},
	"BYND": {
		{
			ID:   7892,
			Slug: "beyond-meat-tokenized-stock-ftx",
			Name: "Beyond Meat tokenized stock FTX",
		},
		{
			ID:   7922,
			Slug: "beyond-meat-inc-tokenized-stock-bittrex",
			Name: "Beyond Meat Inc tokenized stock Bittrex",
		},
	},
	"NYEX": {
		{
			ID:          3129,
			Slug:        "nyerium",
			Name:        "Nyerium",
			BinanceUSDT: "NYEXUSDT",
		},
	},
	"CHOP": {
		{
			ID:          6633,
			Slug:        "porkchop",
			Name:        "Porkchop",
			BinanceUSDT: "CHOPUSDT",
		},
	},
	"ZER": {
		{
			ID:          1578,
			Slug:        "zero",
			Name:        "Zero",
			BinanceUSDT: "ZERUSDT",
		},
	},
	"RBBT": {
		{
			ID:          572,
			Slug:        "rabbitcoin",
			Name:        "RabbitCoin",
			BinanceUSDT: "RBBTUSDT",
		},
	},
	"WEBD": {
		{
			ID:          3877,
			Slug:        "webdollar",
			Name:        "WebDollar",
			BinanceUSDT: "WEBDUSDT",
		},
	},
	"ROX": {
		{
			ID:          3325,
			Slug:        "robotina",
			Name:        "Robotina",
			BinanceUSDT: "ROXUSDT",
		},
	},
	"OBR": {
		{
			ID:          7713,
			Slug:        "order-of-the-black-rose",
			Name:        "Order of the Black Rose",
			BinanceUSDT: "OBRUSDT",
		},
	},
	"CIR": {
		{
			ID:          8174,
			Slug:        "circleswap",
			Name:        "CircleSwap",
			BinanceUSDT: "CIRUSDT",
		},
	},
	"SUP": {
		{
			ID:          7375,
			Slug:        "sup",
			Name:        "SUP",
			BinanceUSDT: "SUPUSDT",
		},
	},
	"PRYZ": {
		{
			ID:          10107,
			Slug:        "pryz",
			Name:        "PRYZ",
			BinanceUSDT: "PRYZUSDT",
		},
	},
	"EMAX": {
		{
			ID:          9855,
			Slug:        "ethereummax",
			Name:        "EthereumMax",
			BinanceUSDT: "EMAXUSDT",
		},
	},
	"HMR": {
		{
			ID:          5336,
			Slug:        "homeros",
			Name:        "Homeros",
			BinanceUSDT: "HMRUSDT",
		},
	},
	"SAFEETH": {
		{
			ID:          10316,
			Slug:        "safeeth",
			Name:        "SafeETH",
			BinanceUSDT: "SAFEETHUSDT",
		},
	},
	"TENFI": {
		{
			ID:          10031,
			Slug:        "ten",
			Name:        "TEN",
			BinanceUSDT: "TENFIUSDT",
		},
	},
	"BUGG": {
		{
			ID:          10479,
			Slug:        "bugg-inu",
			Name:        "Bugg Inu",
			BinanceUSDT: "BUGGUSDT",
		},
	},
	"NUG": {
		{
			ID:          3092,
			Slug:        "nuggets",
			Name:        "Nuggets",
			BinanceUSDT: "NUGUSDT",
		},
	},
	"DERI": {
		{
			ID:          8424,
			Slug:        "deri-protocol",
			Name:        "Deri Protocol",
			BinanceUSDT: "DERIUSDT",
		},
	},
	"PATH": {
		{
			ID:          9878,
			Slug:        "pathfund",
			Name:        "PathFund",
			BinanceUSDT: "PATHUSDT",
		},
	},
	"DVT": {
		{
			ID:          4027,
			Slug:        "devault",
			Name:        "DeVault",
			BinanceUSDT: "DVTUSDT",
		},
	},
	"XMS": {
		{
			ID:          10030,
			Slug:        "mars-ecosystem-governance-token",
			Name:        "Mars Ecosystem Token",
			BinanceUSDT: "XMSUSDT",
		},
	},
	"EDN": {
		{
			ID:          3305,
			Slug:        "eden",
			Name:        "Eden",
			BinanceUSDT: "EDNUSDT",
		},
	},
	"BNBX": {
		{
			ID:          10627,
			Slug:        "bnbx-finance",
			Name:        "BNBX Finance",
			BinanceUSDT: "BNBXUSDT",
		},
	},
	"KCCPAD": {
		{
			ID:          10784,
			Slug:        "kccpad",
			Name:        "KCCPAD",
			BinanceUSDT: "KCCPADUSDT",
		},
	},
	"YFL": {
		{
			ID:          6407,
			Slug:        "yflink",
			Name:        "YF Link",
			BinanceUSDT: "YFLUSDT",
		},
	},
	"RIT20": {
		{
			ID:          8467,
			Slug:        "uberstate-inc",
			Name:        "Uberstate RIT 2.0",
			BinanceUSDT: "RIT20USDT",
		},
	},
	"DRGB": {
		{
			ID:          6423,
			Slug:        "dragonbit",
			Name:        "Dragonbit",
			BinanceUSDT: "DRGBUSDT",
		},
	},
	"CNFI": {
		{
			ID:          8227,
			Slug:        "connect-financial",
			Name:        "Connect Financial",
			BinanceUSDT: "CNFIUSDT",
		},
	},
	"HVI": {
		{
			ID:          10800,
			Slug:        "hungarian-vizsla-inu",
			Name:        "Hungarian Vizsla Inu",
			BinanceUSDT: "HVIUSDT",
		},
	},
	"CUB": {
		{
			ID:          8858,
			Slug:        "cub-finance",
			Name:        "Cub Finance",
			BinanceUSDT: "CUBUSDT",
		},
	},
	"$AVATAR": {
		{
			ID:          10495,
			Slug:        "avatar-moon",
			Name:        "Avatar Moon",
			BinanceUSDT: "$AVATARUSDT",
		},
	},
	"AGRS": {
		{
			ID:          1037,
			Slug:        "agoras-tokens",
			Name:        "Agoras Tokens",
			BinanceUSDT: "AGRSUSDT",
		},
	},
	"KBTC": {
		{
			ID:          8892,
			Slug:        "klondike-btc",
			Name:        "Klondike BTC",
			BinanceUSDT: "KBTCUSDT",
		},
	},
	"QBIT": {
		{
			ID:          10307,
			Slug:        "project-quantum",
			Name:        "Project Quantum",
			BinanceUSDT: "QBITUSDT",
		},
	},
	"ASAC": {
		{
			ID:          5276,
			Slug:        "asac-coin",
			Name:        "Asac Coin",
			BinanceUSDT: "ASACUSDT",
		},
	},
	"DERO": {
		{
			ID:          2665,
			Slug:        "dero",
			Name:        "Dero",
			BinanceUSDT: "DEROUSDT",
		},
	},
	"MOOV": {
		{
			ID:          10046,
			Slug:        "dotmoovs",
			Name:        "Dotmoovs",
			BinanceUSDT: "MOOVUSDT",
		},
	},
	"PNT": {
		{
			ID:   5794,
			Slug: "pnetwork",
			Name: "pNetwork",
		},
		{
			ID:   2691,
			Slug: "penta",
			Name: "Penta",
		},
	},
	"BMI": {
		{
			ID:          8364,
			Slug:        "bridge-mutual",
			Name:        "Bridge Mutual",
			BinanceUSDT: "BMIUSDT",
		},
	},
	"VNT": {
		{
			ID:          3988,
			Slug:        "vnt-chain",
			Name:        "VNT Chain",
			BinanceUSDT: "VNTUSDT",
		},
	},
	"KUB": {
		{
			ID:          5793,
			Slug:        "kublaicoin",
			Name:        "Kublaicoin",
			BinanceUSDT: "KUBUSDT",
		},
	},
	"ZORA": {
		{
			ID:          7826,
			Slug:        "zoracles",
			Name:        "Zoracles",
			BinanceUSDT: "ZORAUSDT",
		},
	},
	"BEET": {
		{
			ID:          3242,
			Slug:        "beetle-coin",
			Name:        "Beetle Coin",
			BinanceUSDT: "BEETUSDT",
		},
	},
	"GLT": {
		{
			ID:          1731,
			Slug:        "globaltoken",
			Name:        "GlobalToken",
			BinanceUSDT: "GLTUSDT",
		},
	},
	"DAFI": {
		{
			ID:          8874,
			Slug:        "dafi-protocol",
			Name:        "DAFI Protocol",
			BinanceUSDT: "DAFIUSDT",
		},
	},
	"BOMB": {
		{
			ID:          3956,
			Slug:        "bomb",
			Name:        "BOMB",
			BinanceUSDT: "BOMBUSDT",
		},
	},
	"NPLC": {
		{
			ID:          3649,
			Slug:        "plus-coin",
			Name:        "Plus-Coin",
			BinanceUSDT: "NPLCUSDT",
		},
	},
	"FF1": {
		{
			ID:          5457,
			Slug:        "two-prime-ff1-token",
			Name:        "Two Prime FF1 Token",
			BinanceUSDT: "FF1USDT",
		},
	},
	"MTS": {
		{
			ID:          7778,
			Slug:        "metis",
			Name:        "Metis",
			BinanceUSDT: "MTSUSDT",
		},
	},
	"HOTDOGE": {
		{
			ID:          9961,
			Slug:        "hotdoge",
			Name:        "HotDoge",
			BinanceUSDT: "HOTDOGEUSDT",
		},
	},
	"MAX": {
		{
			ID:   5067,
			Slug: "max-exchange-token",
			Name: "MAX Exchange Token",
		},
		{
			ID:   128,
			Slug: "maxcoin",
			Name: "Maxcoin",
		},
	},
	"MPG": {
		{
			ID:          3758,
			Slug:        "max-property-group",
			Name:        "Max Property Group",
			BinanceUSDT: "MPGUSDT",
		},
	},
	"DKA": {
		{
			ID:          5908,
			Slug:        "dkargo",
			Name:        "dKargo",
			BinanceUSDT: "DKAUSDT",
		},
	},
	"GOAL": {
		{
			ID:          10267,
			Slug:        "goal",
			Name:        "Goal",
			BinanceUSDT: "GOALUSDT",
		},
	},
	"TRAVEL": {
		{
			ID:          10754,
			Slug:        "travel-care",
			Name:        "Travel Care",
			BinanceUSDT: "TRAVELUSDT",
		},
	},
	"CAB": {
		{
			ID:          1210,
			Slug:        "cabbage",
			Name:        "Cabbage",
			BinanceUSDT: "CABUSDT",
		},
	},
	"BCHA": {
		{
			ID:          7686,
			Slug:        "bitcoin-cash-abc-2",
			Name:        "Bitcoin Cash ABC",
			BinanceUSDT: "BCHAUSDT",
		},
	},
	"UNFI": {
		{
			ID:   7672,
			Slug: "unifi-protocol-dao",
			Name: "Unifi Protocol DAO",
		},
		{
			ID:   7840,
			Slug: "uniswap-finance",
			Name: "Uniswap Finance",
		},
	},
	"FESS": {
		{
			ID:          5910,
			Slug:        "fesschain",
			Name:        "Fesschain",
			BinanceUSDT: "FESSUSDT",
		},
	},
	"SMARS": {
		{
			ID:          10502,
			Slug:        "smars",
			Name:        "SafeMars",
			BinanceUSDT: "SMARSUSDT",
		},
	},
	"UCX": {
		{
			ID:          6604,
			Slug:        "ucx-foundation",
			Name:        "UCX FOUNDATION",
			BinanceUSDT: "UCXUSDT",
		},
	},
	"WISHDOGE": {
		{
			ID:          10966,
			Slug:        "wish-doge-dragon",
			Name:        "Wish Doge Dragon",
			BinanceUSDT: "WISHDOGEUSDT",
		},
	},
	"OBSR": {
		{
			ID:          3698,
			Slug:        "observer",
			Name:        "Observer",
			BinanceUSDT: "OBSRUSDT",
		},
	},
	"YETH": {
		{
			ID:          7989,
			Slug:        "fyeth-finance",
			Name:        "fyeth.finance",
			BinanceUSDT: "YETHUSDT",
		},
	},
	"CSAI": {
		{
			ID:          5264,
			Slug:        "compound-sai",
			Name:        "Compound SAI",
			BinanceUSDT: "CSAIUSDT",
		},
	},
	"CIV": {
		{
			ID:          3381,
			Slug:        "civitas",
			Name:        "Civitas",
			BinanceUSDT: "CIVUSDT",
		},
	},
	"FORTH": {
		{
			ID:          9421,
			Slug:        "ampleforth-governance-token",
			Name:        "Ampleforth Governance Token",
			BinanceUSDT: "FORTHUSDT",
		},
	},
	"BTCHG": {
		{
			ID:          5732,
			Slug:        "bitcoinhedge",
			Name:        "BITCOINHEDGE",
			BinanceUSDT: "BTCHGUSDT",
		},
	},
	"XST": {
		{
			ID:   448,
			Slug: "stealth",
			Name: "Stealth",
		},
		{
			ID:   8330,
			Slug: "xstable-protocol",
			Name: "Xstable.Protocol",
		},
	},
	"BZZ": {
		{
			ID:          10293,
			Slug:        "ethereum-swarm",
			Name:        "Swarm",
			BinanceUSDT: "BZZUSDT",
		},
	},
	"NAS": {
		{
			ID:          1908,
			Slug:        "nebulas-token",
			Name:        "Nebulas",
			BinanceUSDT: "NASUSDT",
		},
	},
	"GC": {
		{
			ID:   5001,
			Slug: "gric-coin",
			Name: "Gric Coin",
		},
		{
			ID:   6095,
			Slug: "galaxy-wallet",
			Name: "Galaxy Wallet",
		},
	},
	"NVA": {
		{
			ID:          7507,
			Slug:        "neeva-defi",
			Name:        "Neeva Defi",
			BinanceUSDT: "NVAUSDT",
		},
	},
	"LNCHX": {
		{
			ID:          9771,
			Slug:        "launchx",
			Name:        "LaunchX",
			BinanceUSDT: "LNCHXUSDT",
		},
	},
	"AOA": {
		{
			ID:          2874,
			Slug:        "aurora",
			Name:        "Aurora",
			BinanceUSDT: "AOAUSDT",
		},
	},
	"XED": {
		{
			ID:          8163,
			Slug:        "exeedme",
			Name:        "Exeedme",
			BinanceUSDT: "XEDUSDT",
		},
	},
	"SHIFT": {
		{
			ID:          1050,
			Slug:        "shift",
			Name:        "Shift",
			BinanceUSDT: "SHIFTUSDT",
		},
	},
	"ETGF": {
		{
			ID:          7454,
			Slug:        "etg-finance",
			Name:        "ETG Finance",
			BinanceUSDT: "ETGFUSDT",
		},
	},
	"SFUND": {
		{
			ID:          8972,
			Slug:        "seedify-fund",
			Name:        "Seedify.fund",
			BinanceUSDT: "SFUNDUSDT",
		},
	},
	"MOONARCH": {
		{
			ID:          10117,
			Slug:        "moonarch-app",
			Name:        "Moonarch.app",
			BinanceUSDT: "MOONARCHUSDT",
		},
	},
	"HLT": {
		{
			ID:          3838,
			Slug:        "esportbits",
			Name:        "Esportbits",
			BinanceUSDT: "HLTUSDT",
		},
	},
	"KFX": {
		{
			ID:          8177,
			Slug:        "knoxfs-new",
			Name:        "KnoxFS (new)",
			BinanceUSDT: "KFXUSDT",
		},
	},
	"PLC": {
		{
			ID:          3364,
			Slug:        "platincoin",
			Name:        "PLATINCOIN",
			BinanceUSDT: "PLCUSDT",
		},
	},
	"TAC": {
		{
			ID:          3227,
			Slug:        "traceability-chain",
			Name:        "Traceability Chain",
			BinanceUSDT: "TACUSDT",
		},
	},
	"SFD": {
		{
			ID:          7270,
			Slug:        "safe-deal",
			Name:        "SAFE DEAL",
			BinanceUSDT: "SFDUSDT",
		},
	},
	"SCRT": {
		{
			ID:          5604,
			Slug:        "secret",
			Name:        "Secret",
			BinanceUSDT: "SCRTUSDT",
		},
	},
	"MANA": {
		{
			ID:          1966,
			Slug:        "decentraland",
			Name:        "Decentraland",
			BinanceUSDT: "MANAUSDT",
		},
	},
	"BITCNY": {
		{
			ID:          624,
			Slug:        "bitcny",
			Name:        "bitCNY",
			BinanceUSDT: "BITCNYUSDT",
		},
	},
	"CARDS": {
		{
			ID:          9047,
			Slug:        "card-starter",
			Name:        "CARD.STARTER",
			BinanceUSDT: "CARDSUSDT",
		},
	},
	"COAL": {
		{
			ID:   1824,
			Slug: "bitcoal",
			Name: "BitCoal",
		},
		{
			ID:   6248,
			Slug: "coalculus",
			Name: "Coalculus",
		},
	},
	"SDEFI": {
		{
			ID:          5862,
			Slug:        "sdefi",
			Name:        "sDEFI",
			BinanceUSDT: "SDEFIUSDT",
		},
	},
	"BTRL": {
		{
			ID:          4174,
			Slug:        "bitcoinregular",
			Name:        "BitcoinRegular",
			BinanceUSDT: "BTRLUSDT",
		},
	},
	"FEG": {
		{
			ID:          8397,
			Slug:        "fegtoken",
			Name:        "FEG Token",
			BinanceUSDT: "FEGUSDT",
		},
	},
	"BKBT": {
		{
			ID:          2914,
			Slug:        "beekan",
			Name:        "BeeKan",
			BinanceUSDT: "BKBTUSDT",
		},
	},
	"PET": {
		{
			ID:          8915,
			Slug:        "battle-pets",
			Name:        "Battle Pets",
			BinanceUSDT: "PETUSDT",
		},
	},
	"FUNDX": {
		{
			ID:          8681,
			Slug:        "funder-one-capital",
			Name:        "Funder One Capital",
			BinanceUSDT: "FUNDXUSDT",
		},
	},
	"GVE": {
		{
			ID:          2969,
			Slug:        "globalvillage-ecosystem",
			Name:        "Globalvillage Ecosystem",
			BinanceUSDT: "GVEUSDT",
		},
	},
	"DIAMOND": {
		{
			ID:          9650,
			Slug:        "diamondtoken",
			Name:        "DiamondToken",
			BinanceUSDT: "DIAMONDUSDT",
		},
	},
	"OBS": {
		{
			ID:   10814,
			Slug: "one-basis-cash",
			Name: "One Basis Cash",
		},
		{
			ID:   9485,
			Slug: "openbisea",
			Name: "OpenBiSea",
		},
	},
	"SOL": {
		{
			ID:          5426,
			Slug:        "solana",
			Name:        "Solana",
			BinanceUSDT: "SOLUSDT",
		},
	},
	"FOMO": {
		{
			ID:          9646,
			Slug:        "fomo-lab",
			Name:        "FOMO LAB",
			BinanceUSDT: "FOMOUSDT",
		},
	},
	"TENSHI": {
		{
			ID:          10957,
			Slug:        "tenshi",
			Name:        "Tenshi",
			BinanceUSDT: "TENSHIUSDT",
		},
	},
	"BDP": {
		{
			ID:   8708,
			Slug: "big-data-protocol",
			Name: "Big Data Protocol",
		},
		{
			ID:   4452,
			Slug: "bidipass",
			Name: "BidiPass",
		},
	},
	"ORMEUS": {
		{
			ID:          1998,
			Slug:        "ormeus-coin",
			Name:        "Ormeus Coin",
			BinanceUSDT: "ORMEUSUSDT",
		},
	},
	"KTT": {
		{
			ID:          5130,
			Slug:        "k-tune",
			Name:        "K-Tune",
			BinanceUSDT: "KTTUSDT",
		},
	},
	"ZIL": {
		{
			ID:          2469,
			Slug:        "zilliqa",
			Name:        "Zilliqa",
			BinanceUSDT: "ZILUSDT",
		},
	},
	"ELON": {
		{
			ID:          9436,
			Slug:        "dogelon",
			Name:        "Dogelon Mars",
			BinanceUSDT: "ELONUSDT",
		},
	},
	"MWG": {
		{
			ID:          8505,
			Slug:        "metawhale-gold",
			Name:        "Metawhale Gold",
			BinanceUSDT: "MWGUSDT",
		},
	},
	"MANGO": {
		{
			ID:          8856,
			Slug:        "mango-finance",
			Name:        "Mango Finance",
			BinanceUSDT: "MANGOUSDT",
		},
	},
	"CRU": {
		{
			ID:          6747,
			Slug:        "crustnetwork",
			Name:        "Crust Network",
			BinanceUSDT: "CRUUSDT",
		},
	},
	"NULS": {
		{
			ID:          2092,
			Slug:        "nuls",
			Name:        "NULS",
			BinanceUSDT: "NULSUSDT",
		},
	},
	"vFIL": {
		{
			ID:          8213,
			Slug:        "venus-filecoin",
			Name:        "Venus Filecoin",
			BinanceUSDT: "vFILUSDT",
		},
	},
	"ETHHEDGE": {
		{
			ID:          5658,
			Slug:        "1x-short-ethereum-token",
			Name:        "1X Short Ethereum Token",
			BinanceUSDT: "ETHHEDGEUSDT",
		},
	},
	"CBK": {
		{
			ID:   8107,
			Slug: "cobak-token",
			Name: "Cobak Token",
		},
		{
			ID:   10733,
			Slug: "crossing-the-yellow-blocks",
			Name: "Crossing the Yellow Blocks",
		},
	},
	"BNT": {
		{
			ID:          1727,
			Slug:        "bancor",
			Name:        "Bancor",
			BinanceUSDT: "BNTUSDT",
		},
	},
	"LUCY": {
		{
			ID:          5291,
			Slug:        "lucy",
			Name:        "LUCY",
			BinanceUSDT: "LUCYUSDT",
		},
	},
	"UCM": {
		{
			ID:   9475,
			Slug: "unicly-chris-mccann-collection",
			Name: "Unicly Chris McCann Collection",
		},
		{
			ID:   5633,
			Slug: "ucrowdme",
			Name: "UCROWDME",
		},
	},
	"BIUT": {
		{
			ID:          4749,
			Slug:        "bit-trust-system",
			Name:        "Bit Trust System",
			BinanceUSDT: "BIUTUSDT",
		},
	},
	"CMM": {
		{
			ID:          3080,
			Slug:        "commercium",
			Name:        "Commercium",
			BinanceUSDT: "CMMUSDT",
		},
	},
	"QDAO": {
		{
			ID:          4053,
			Slug:        "q-dao-governance-token",
			Name:        "Q DAO Governance token v1.0",
			BinanceUSDT: "QDAOUSDT",
		},
	},
	"FNB": {
		{
			ID:   3858,
			Slug: "fnb-protocol",
			Name: "FNB Protocol",
		},
		{
			ID:   4794,
			Slug: "finexboxtoken",
			Name: "FinexboxToken",
		},
	},
	"EFT": {
		{
			ID:          8851,
			Slug:        "eft-finance",
			Name:        "EFT.finance",
			BinanceUSDT: "EFTUSDT",
		},
	},
	"BALLBAG": {
		{
			ID:          10543,
			Slug:        "ballbag-token",
			Name:        "Ballbag Token",
			BinanceUSDT: "BALLBAGUSDT",
		},
	},
	"BOA": {
		{
			ID:          4217,
			Slug:        "bosagora",
			Name:        "BOSAGORA",
			BinanceUSDT: "BOAUSDT",
		},
	},
	"BRICK": {
		{
			ID:          8853,
			Slug:        "brickchain-finance",
			Name:        "Brickchain Finance",
			BinanceUSDT: "BRICKUSDT",
		},
	},
	"MBY": {
		{
			ID:          10583,
			Slug:        "monkey-token",
			Name:        "Monkey Token",
			BinanceUSDT: "MBYUSDT",
		},
	},
	"SAUBER": {
		{
			ID:          10793,
			Slug:        "alfa-romeo-racing-orlen-fan-token",
			Name:        "Alfa Romeo Racing ORLEN Fan Token",
			BinanceUSDT: "SAUBERUSDT",
		},
	},
	"AZX": {
		{
			ID:          9917,
			Slug:        "azeusx",
			Name:        "AzeusX",
			BinanceUSDT: "AZXUSDT",
		},
	},
	"WIN": {
		{
			ID:          4206,
			Slug:        "wink",
			Name:        "WINkLink",
			BinanceUSDT: "WINUSDT",
		},
	},
	"TVNT": {
		{
			ID:          3644,
			Slug:        "travelnote",
			Name:        "TravelNote",
			BinanceUSDT: "TVNTUSDT",
		},
	},
	"CRAZYTIME": {
		{
			ID:          10582,
			Slug:        "crazytime",
			Name:        "CrazyTime",
			BinanceUSDT: "CRAZYTIMEUSDT",
		},
	},
	"STS": {
		{
			ID:          5039,
			Slug:        "sbank",
			Name:        "SBank",
			BinanceUSDT: "STSUSDT",
		},
	},
	"SPERM": {
		{
			ID:          9902,
			Slug:        "esperm",
			Name:        "Elon Sperm",
			BinanceUSDT: "SPERMUSDT",
		},
	},
	"INB": {
		{
			ID:          3116,
			Slug:        "insight-chain",
			Name:        "Insight Chain",
			BinanceUSDT: "INBUSDT",
		},
	},
	"BBI": {
		{
			ID:          8578,
			Slug:        "bigboys-industry",
			Name:        "BigBoys Industry",
			BinanceUSDT: "BBIUSDT",
		},
	},
	"META": {
		{
			ID:          3418,
			Slug:        "metadium",
			Name:        "Metadium",
			BinanceUSDT: "METAUSDT",
		},
	},
	"APR": {
		{
			ID:          2721,
			Slug:        "apr-coin",
			Name:        "APR Coin",
			BinanceUSDT: "APRUSDT",
		},
	},
	"VPP": {
		{
			ID:          9756,
			Slug:        "virtue-poker",
			Name:        "Virtue Poker",
			BinanceUSDT: "VPPUSDT",
		},
	},
	"DAGT": {
		{
			ID:          3357,
			Slug:        "digital-asset-guarantee-token",
			Name:        "Digital Asset Guarantee Token",
			BinanceUSDT: "DAGTUSDT",
		},
	},
	"B20": {
		{
			ID:          8457,
			Slug:        "b20",
			Name:        "B20",
			BinanceUSDT: "B20USDT",
		},
	},
	"LOT": {
		{
			ID:   9446,
			Slug: "lottery-token",
			Name: "Lottery Token",
		},
		{
			ID:   4285,
			Slug: "lukki-operating-token",
			Name: "Lukki Operating Token",
		},
	},
	"IOG": {
		{
			ID:          3315,
			Slug:        "playgroundz",
			Name:        "Playgroundz",
			BinanceUSDT: "IOGUSDT",
		},
	},
	"SAMO": {
		{
			ID:          9721,
			Slug:        "somoyedcoin",
			Name:        "Samoyedcoin",
			BinanceUSDT: "SAMOUSDT",
		},
	},
	"CTASK": {
		{
			ID:          8418,
			Slug:        "cryptotask",
			Name:        "CryptoTask",
			BinanceUSDT: "CTASKUSDT",
		},
	},
	"JWL": {
		{
			ID:          3791,
			Slug:        "jewel",
			Name:        "Jewel",
			BinanceUSDT: "JWLUSDT",
		},
	},
	"GRAP": {
		{
			ID:          6664,
			Slug:        "grap",
			Name:        "GRAP",
			BinanceUSDT: "GRAPUSDT",
		},
	},
	"BCMC1": {
		{
			ID:          9318,
			Slug:        "beforecoinmarketcap",
			Name:        "BeforeCoinMarketCap",
			BinanceUSDT: "BCMC1USDT",
		},
	},
	"ACC": {
		{
			ID:          4696,
			Slug:        "asian-african-capital-chain",
			Name:        "Asian-African Capital Chain",
			BinanceUSDT: "ACCUSDT",
		},
	},
	"NRVE": {
		{
			ID:          2956,
			Slug:        "narrative",
			Name:        "Narrative",
			BinanceUSDT: "NRVEUSDT",
		},
	},
	"HDI": {
		{
			ID:          5797,
			Slug:        "heidi",
			Name:        "HEIDI",
			BinanceUSDT: "HDIUSDT",
		},
	},
	"XPTX": {
		{
			ID:          1254,
			Slug:        "platinumbar",
			Name:        "PlatinumBAR",
			BinanceUSDT: "XPTXUSDT",
		},
	},
	"ETZ": {
		{
			ID:          2843,
			Slug:        "ether-zero",
			Name:        "Ether Zero",
			BinanceUSDT: "ETZUSDT",
		},
	},
	"MORA": {
		{
			ID:          9031,
			Slug:        "meliora",
			Name:        "Meliora",
			BinanceUSDT: "MORAUSDT",
		},
	},
	"GOM2": {
		{
			ID:          5839,
			Slug:        "animalgo",
			Name:        "AnimalGo",
			BinanceUSDT: "GOM2USDT",
		},
	},
	"SUSHIUP": {
		{
			ID:          8053,
			Slug:        "sushiup",
			Name:        "SUSHIUP",
			BinanceUSDT: "SUSHIUPUSDT",
		},
	},
	"YFI3": {
		{
			ID:          7794,
			Slug:        "yfi3-money",
			Name:        "YFI3.money",
			BinanceUSDT: "YFI3USDT",
		},
	},
	"XBN": {
		{
			ID:          9385,
			Slug:        "xbn",
			Name:        "Elastic BNB",
			BinanceUSDT: "XBNUSDT",
		},
	},
	"DOOS": {
		{
			ID:          6983,
			Slug:        "doos-token",
			Name:        "DOOS TOKEN",
			BinanceUSDT: "DOOSUSDT",
		},
	},
	"COMP": {
		{
			ID:   5692,
			Slug: "compound",
			Name: "Compound",
		},
		{
			ID:   3180,
			Slug: "compound-coin",
			Name: "Compound Coin",
		},
	},
	"ULT": {
		{
			ID:   3666,
			Slug: "ultiledger",
			Name: "Ultiledger",
		},
		{
			ID:   5330,
			Slug: "shardus",
			Name: "Shardus",
		},
	},
	"DBET": {
		{
			ID:          2175,
			Slug:        "decent-bet",
			Name:        "DecentBet",
			BinanceUSDT: "DBETUSDT",
		},
	},
	"IQ": {
		{
			ID:   2930,
			Slug: "everipedia",
			Name: "Everipedia",
		},
		{
			ID:   3273,
			Slug: "iqcash",
			Name: "IQ.cash",
		},
	},
	"PCM": {
		{
			ID:          4958,
			Slug:        "precium",
			Name:        "Precium",
			BinanceUSDT: "PCMUSDT",
		},
	},
	"NDN": {
		{
			ID:          5629,
			Slug:        "ndn-link",
			Name:        "NDN Link",
			BinanceUSDT: "NDNUSDT",
		},
	},
	"BITS": {
		{
			ID:   659,
			Slug: "bitswift",
			Name: "Bitswift",
		},
		{
			ID:   276,
			Slug: "bitstar",
			Name: "Bitstar",
		},
		{
			ID:   3460,
			Slug: "bitcoinus",
			Name: "Bitcoinus",
		},
		{
			ID:   9023,
			Slug: "bit",
			Name: "Bit",
		},
	},
	"$HINA": {
		{
			ID:          10189,
			Slug:        "hina-inu",
			Name:        "Hina Inu",
			BinanceUSDT: "$HINAUSDT",
		},
	},
	"CREAM": {
		{
			ID:          6193,
			Slug:        "cream-finance",
			Name:        "Cream Finance",
			BinanceUSDT: "CREAMUSDT",
		},
	},
	"LTCU": {
		{
			ID:          1935,
			Slug:        "litecoin-ultra",
			Name:        "LiteCoin Ultra",
			BinanceUSDT: "LTCUUSDT",
		},
	},
	"BBO": {
		{
			ID:          2836,
			Slug:        "bigbom",
			Name:        "Bigbom",
			BinanceUSDT: "BBOUSDT",
		},
	},
	"TBT": {
		{
			ID:          6565,
			Slug:        "tidebit-token",
			Name:        "TideBit Token",
			BinanceUSDT: "TBTUSDT",
		},
	},
	"MLR": {
		{
			ID:          5433,
			Slug:        "mega-lottery-services-global",
			Name:        "Mega Lottery Services Global",
			BinanceUSDT: "MLRUSDT",
		},
	},
	"HNDC": {
		{
			ID:          3504,
			Slug:        "hondaiscoin",
			Name:        "HondaisCoin",
			BinanceUSDT: "HNDCUSDT",
		},
	},
	"KFC": {
		{
			ID:          7169,
			Slug:        "chicken",
			Name:        "Chicken",
			BinanceUSDT: "KFCUSDT",
		},
	},
	"DRCT": {
		{
			ID:          10385,
			Slug:        "ally-direct-token",
			Name:        "Ally Direct Token",
			BinanceUSDT: "DRCTUSDT",
		},
	},
	"AHOUSE": {
		{
			ID:          10541,
			Slug:        "animalhouse",
			Name:        "AnimalHouse",
			BinanceUSDT: "AHOUSEUSDT",
		},
	},
	"MNTT": {
		{
			ID:          9709,
			Slug:        "moontrust",
			Name:        "MoonTrust",
			BinanceUSDT: "MNTTUSDT",
		},
	},
	"EBOX": {
		{
			ID:          8930,
			Slug:        "ethbox",
			Name:        "Ethbox",
			BinanceUSDT: "EBOXUSDT",
		},
	},
	"PBOM": {
		{
			ID:          8928,
			Slug:        "pocket-bomb",
			Name:        "Pocket Bomb",
			BinanceUSDT: "PBOMUSDT",
		},
	},
	"1GOLD": {
		{
			ID:          4774,
			Slug:        "1irstgold",
			Name:        "1irstGold",
			BinanceUSDT: "1GOLDUSDT",
		},
	},
	"FINU": {
		{
			ID:          10173,
			Slug:        "fire-inu",
			Name:        "Fire Inu",
			BinanceUSDT: "FINUUSDT",
		},
	},
	"QTCON": {
		{
			ID:          4746,
			Slug:        "quiztok",
			Name:        "Quiztok",
			BinanceUSDT: "QTCONUSDT",
		},
	},
	"WCT": {
		{
			ID:          1527,
			Slug:        "waves-community-token",
			Name:        "Waves Community Token",
			BinanceUSDT: "WCTUSDT",
		},
	},
	"CTX": {
		{
			ID:          10368,
			Slug:        "cryptex-finance",
			Name:        "Cryptex Finance",
			BinanceUSDT: "CTXUSDT",
		},
	},
	"CPAC": {
		{
			ID:          10263,
			Slug:        "compact",
			Name:        "Compact",
			BinanceUSDT: "CPACUSDT",
		},
	},
	"KCS": {
		{
			ID:          2087,
			Slug:        "kucoin-token",
			Name:        "KuCoin Token",
			BinanceUSDT: "KCSUSDT",
		},
	},
	"INST": {
		{
			ID:          10508,
			Slug:        "instadapp",
			Name:        "Instadapp",
			BinanceUSDT: "INSTUSDT",
		},
	},
	"LAMB": {
		{
			ID:          3657,
			Slug:        "lambda",
			Name:        "Lambda",
			BinanceUSDT: "LAMBUSDT",
		},
	},
	"GULAG": {
		{
			ID:          10923,
			Slug:        "gulag-token",
			Name:        "Gulag Token",
			BinanceUSDT: "GULAGUSDT",
		},
	},
	"BC": {
		{
			ID:          3976,
			Slug:        "bitcoin-confidential",
			Name:        "Bitcoin Confidential",
			BinanceUSDT: "BCUSDT",
		},
	},
	"SYLO": {
		{
			ID:          5662,
			Slug:        "sylo",
			Name:        "Sylo",
			BinanceUSDT: "SYLOUSDT",
		},
	},
	"NFTB": {
		{
			ID:          9545,
			Slug:        "nftb",
			Name:        "NFTb",
			BinanceUSDT: "NFTBUSDT",
		},
	},
	"CURE": {
		{
			ID:   333,
			Slug: "curecoin",
			Name: "Curecoin",
		},
		{
			ID:   8326,
			Slug: "cure-farm",
			Name: "CURE Farm",
		},
	},
	"MIR": {
		{
			ID:   7857,
			Slug: "mirror-protocol",
			Name: "Mirror Protocol",
		},
		{
			ID:   3371,
			Slug: "mir-coin",
			Name: "MIR COIN",
		},
	},
	"Mars": {
		{
			ID:          8253,
			Slug:        "mars",
			Name:        "Mars",
			BinanceUSDT: "MarsUSDT",
		},
	},
	"XNS": {
		{
			ID:          4802,
			Slug:        "xeonbit-token",
			Name:        "Xeonbit Token",
			BinanceUSDT: "XNSUSDT",
		},
	},
	"DEUS": {
		{
			ID:          7125,
			Slug:        "deus-finance",
			Name:        "DEUS Finance",
			BinanceUSDT: "DEUSUSDT",
		},
	},
	"OCEAN": {
		{
			ID:          3911,
			Slug:        "ocean-protocol",
			Name:        "Ocean Protocol",
			BinanceUSDT: "OCEANUSDT",
		},
	},
	"GPT": {
		{
			ID:          7946,
			Slug:        "grace-period-token",
			Name:        "Grace Period Token",
			BinanceUSDT: "GPTUSDT",
		},
	},
	"NS": {
		{
			ID:          8633,
			Slug:        "nodestats",
			Name:        "Nodestats",
			BinanceUSDT: "NSUSDT",
		},
	},
	"CC10": {
		{
			ID:          8887,
			Slug:        "cryptocurrency-top-10-tokens-index",
			Name:        "Cryptocurrency Top 10 Tokens Index",
			BinanceUSDT: "CC10USDT",
		},
	},
	"PLUT": {
		{
			ID:          7968,
			Slug:        "pluto",
			Name:        "Pluto",
			BinanceUSDT: "PLUTUSDT",
		},
	},
	"PANDO": {
		{
			ID:          8711,
			Slug:        "pando",
			Name:        "Pando",
			BinanceUSDT: "PANDOUSDT",
		},
	},
	"GBYTE": {
		{
			ID:          1492,
			Slug:        "obyte",
			Name:        "Obyte",
			BinanceUSDT: "GBYTEUSDT",
		},
	},
	"ESWAP": {
		{
			ID:          9567,
			Slug:        "eswapping",
			Name:        "eSwapping",
			BinanceUSDT: "ESWAPUSDT",
		},
	},
	"TGAME": {
		{
			ID:          2929,
			Slug:        "tgame",
			Name:        "Truegame",
			BinanceUSDT: "TGAMEUSDT",
		},
	},
	"BMC": {
		{
			ID:          1976,
			Slug:        "blackmoon",
			Name:        "Blackmoon",
			BinanceUSDT: "BMCUSDT",
		},
	},
	"TSLA": {
		{
			ID:   7887,
			Slug: "tesla-tokenized-stock-ftx",
			Name: "Tesla tokenized stock FTX",
		},
		{
			ID:   7598,
			Slug: "tessla-coin",
			Name: "Tessla Coin",
		},
		{
			ID:   7919,
			Slug: "tesla-tokenized-stock-bittrex",
			Name: "Tesla tokenized stock Bittrex",
		},
	},
	"PTI": {
		{
			ID:          3859,
			Slug:        "paytomat",
			Name:        "Paytomat",
			BinanceUSDT: "PTIUSDT",
		},
	},
	"EXRD": {
		{
			ID:          7692,
			Slug:        "radix",
			Name:        "Radix",
			BinanceUSDT: "EXRDUSDT",
		},
	},
	"BADGER": {
		{
			ID:          7859,
			Slug:        "badger-dao",
			Name:        "Badger DAO",
			BinanceUSDT: "BADGERUSDT",
		},
	},
	"TUSD": {
		{
			ID:          2563,
			Slug:        "trueusd",
			Name:        "TrueUSD",
			BinanceUSDT: "TUSDUSDT",
		},
	},
	"AENS": {
		{
			ID:          3683,
			Slug:        "aen-smart-token",
			Name:        "AEN Smart Token",
			BinanceUSDT: "AENSUSDT",
		},
	},
	"BUSY": {
		{
			ID:          9002,
			Slug:        "busy",
			Name:        "Busy DAO",
			BinanceUSDT: "BUSYUSDT",
		},
	},
	"WHEAT": {
		{
			ID:          9229,
			Slug:        "wheat",
			Name:        "WHEAT Token",
			BinanceUSDT: "WHEATUSDT",
		},
	},
	"PASC": {
		{
			ID:          1473,
			Slug:        "pascal",
			Name:        "Pascal",
			BinanceUSDT: "PASCUSDT",
		},
	},
	"HLAND": {
		{
			ID:          8021,
			Slug:        "hland-token",
			Name:        "HLand Token",
			BinanceUSDT: "HLANDUSDT",
		},
	},
	"YTN": {
		{
			ID:          2290,
			Slug:        "yenten",
			Name:        "YENTEN",
			BinanceUSDT: "YTNUSDT",
		},
	},
	"BAC": {
		{
			ID:          7813,
			Slug:        "basis-cash",
			Name:        "Basis Cash",
			BinanceUSDT: "BACUSDT",
		},
	},
	"SOFI": {
		{
			ID:          6825,
			Slug:        "social-finance",
			Name:        "Social Finance",
			BinanceUSDT: "SOFIUSDT",
		},
	},
	"RUFF": {
		{
			ID:          2476,
			Slug:        "ruff",
			Name:        "Ruff",
			BinanceUSDT: "RUFFUSDT",
		},
	},
	"SNOWGE": {
		{
			ID:          9772,
			Slug:        "snowgecoin",
			Name:        "SnowgeCoin",
			BinanceUSDT: "SNOWGEUSDT",
		},
	},
	"RAY": {
		{
			ID:          8526,
			Slug:        "raydium",
			Name:        "Raydium",
			BinanceUSDT: "RAYUSDT",
		},
	},
	"XPO": {
		{
			ID:          9048,
			Slug:        "xpool",
			Name:        "Xpool",
			BinanceUSDT: "XPOUSDT",
		},
	},
	"AXIA": {
		{
			ID:          7474,
			Slug:        "axia-protocol",
			Name:        "Axia Protocol",
			BinanceUSDT: "AXIAUSDT",
		},
	},
	"PETB": {
		{
			ID:          10152,
			Slug:        "petbloc",
			Name:        "PETBloc",
			BinanceUSDT: "PETBUSDT",
		},
	},
	"KBC": {
		{
			ID:          2907,
			Slug:        "karatgold-coin",
			Name:        "Karatgold Coin",
			BinanceUSDT: "KBCUSDT",
		},
	},
	"KYTE": {
		{
			ID:          9192,
			Slug:        "kambria-yield-tuning-engine",
			Name:        "Kambria Yield Tuning Engine",
			BinanceUSDT: "KYTEUSDT",
		},
	},
	"GPO": {
		{
			ID:          6478,
			Slug:        "galaxy-pool-coin",
			Name:        "Galaxy Pool Coin",
			BinanceUSDT: "GPOUSDT",
		},
	},
	"CYTR": {
		{
			ID:          7307,
			Slug:        "cyclops-treasure",
			Name:        "Cyclops Treasure",
			BinanceUSDT: "CYTRUSDT",
		},
	},
	"INSN": {
		{
			ID:          1678,
			Slug:        "insanecoin-insn",
			Name:        "InsaneCoin",
			BinanceUSDT: "INSNUSDT",
		},
	},
	"MKMOON": {
		{
			ID:          10057,
			Slug:        "monkeycoin",
			Name:        "MonkeyCoin",
			BinanceUSDT: "MKMOONUSDT",
		},
	},
	"HTR": {
		{
			ID:          5552,
			Slug:        "hathor",
			Name:        "Hathor",
			BinanceUSDT: "HTRUSDT",
		},
	},
	"SAFESPACE": {
		{
			ID:          10075,
			Slug:        "safespace",
			Name:        "SAFESPACE",
			BinanceUSDT: "SAFESPACEUSDT",
		},
	},
	"SHOPX": {
		{
			ID:          8863,
			Slug:        "splyt",
			Name:        "Splyt",
			BinanceUSDT: "SHOPXUSDT",
		},
	},
	"DEA": {
		{
			ID:          8248,
			Slug:        "deus-finance-dea",
			Name:        "DEUS Finance DEA",
			BinanceUSDT: "DEAUSDT",
		},
	},
	"VITAE": {
		{
			ID:          3063,
			Slug:        "vitae",
			Name:        "Vitae",
			BinanceUSDT: "VITAEUSDT",
		},
	},
	"CET": {
		{
			ID:          2941,
			Slug:        "coinex-token",
			Name:        "CoinEx Token",
			BinanceUSDT: "CETUSDT",
		},
	},
	"APYS": {
		{
			ID:          8419,
			Slug:        "apyswap",
			Name:        "APYSwap",
			BinanceUSDT: "APYSUSDT",
		},
	},
	"NTVRK": {
		{
			ID:          9954,
			Slug:        "netvrk",
			Name:        "Netvrk",
			BinanceUSDT: "NTVRKUSDT",
		},
	},
	"CELEB": {
		{
			ID:          8788,
			Slug:        "celebplus",
			Name:        "CELEBPLUS",
			BinanceUSDT: "CELEBUSDT",
		},
	},
	"CHAD": {
		{
			ID:          9942,
			Slug:        "the-chad-project",
			Name:        "The Chad Token",
			BinanceUSDT: "CHADUSDT",
		},
	},
	"ETHIX": {
		{
			ID:          8442,
			Slug:        "ethichub",
			Name:        "EthicHub",
			BinanceUSDT: "ETHIXUSDT",
		},
	},
	"ASK": {
		{
			ID:          7105,
			Slug:        "permission-coin",
			Name:        "Permission Coin",
			BinanceUSDT: "ASKUSDT",
		},
	},
	"EXF": {
		{
			ID:          8563,
			Slug:        "extend-finance",
			Name:        "Extend Finance",
			BinanceUSDT: "EXFUSDT",
		},
	},
	"AICO": {
		{
			ID:          5920,
			Slug:        "aicon",
			Name:        "AICON",
			BinanceUSDT: "AICOUSDT",
		},
	},
	"BUZZ": {
		{
			ID:          1962,
			Slug:        "buzzcoin",
			Name:        "BUZZCoin",
			BinanceUSDT: "BUZZUSDT",
		},
	},
	"MORE": {
		{
			ID:          1722,
			Slug:        "more-coin",
			Name:        "More Coin",
			BinanceUSDT: "MOREUSDT",
		},
	},
	"AXIS": {
		{
			ID:   6670,
			Slug: "axis-defi",
			Name: "Axis DeFi",
		},
		{
			ID:   8849,
			Slug: "axis-token",
			Name: "AXIS Token",
		},
	},
	"wCRES": {
		{
			ID:          7818,
			Slug:        "wrapped-crescofin",
			Name:        "Wrapped CrescoFin",
			BinanceUSDT: "wCRESUSDT",
		},
	},
	"CRS": {
		{
			ID:          8608,
			Slug:        "crypto-rewards-studio",
			Name:        "Crypto Rewards Studio",
			BinanceUSDT: "CRSUSDT",
		},
	},
	"NKN": {
		{
			ID:          2780,
			Slug:        "nkn",
			Name:        "NKN",
			BinanceUSDT: "NKNUSDT",
		},
	},
	"MDTK": {
		{
			ID:          4786,
			Slug:        "mdtoken",
			Name:        "MDtoken",
			BinanceUSDT: "MDTKUSDT",
		},
	},
	"PACOCA": {
		{
			ID:          10522,
			Slug:        "pacoca",
			Name:        "Pacoca",
			BinanceUSDT: "PACOCAUSDT",
		},
	},
	"LVH": {
		{
			ID:          6495,
			Slug:        "lovehearts",
			Name:        "LoveHearts",
			BinanceUSDT: "LVHUSDT",
		},
	},
	"LOVE": {
		{
			ID:          6615,
			Slug:        "love-coin",
			Name:        "Love Coin",
			BinanceUSDT: "LOVEUSDT",
		},
	},
	"DMT": {
		{
			ID:   2503,
			Slug: "dmarket",
			Name: "DMarket",
		},
		{
			ID:   9186,
			Slug: "dark-matter",
			Name: "Dark Matter",
		},
	},
	"KAREN": {
		{
			ID:          10378,
			Slug:        "karencoin",
			Name:        "KarenCoin",
			BinanceUSDT: "KARENUSDT",
		},
	},
	"AGOLP": {
		{
			ID:          9407,
			Slug:        "algoil",
			Name:        "AlgOil",
			BinanceUSDT: "AGOLPUSDT",
		},
	},
	"ALPA": {
		{
			ID:          7618,
			Slug:        "alpaca-city",
			Name:        "Alpaca City",
			BinanceUSDT: "ALPAUSDT",
		},
	},
	"UPUNK": {
		{
			ID:          9473,
			Slug:        "unicly-cryptopunks-collection",
			Name:        "Unicly CryptoPunks Collection",
			BinanceUSDT: "UPUNKUSDT",
		},
	},
	"SZC": {
		{
			ID:          7988,
			Slug:        "zugacoin",
			Name:        "Zugacoin",
			BinanceUSDT: "SZCUSDT",
		},
	},
	"MOTO": {
		{
			ID:          360,
			Slug:        "motocoin",
			Name:        "Motocoin",
			BinanceUSDT: "MOTOUSDT",
		},
	},
	"MIB": {
		{
			ID:          3210,
			Slug:        "mib-coin",
			Name:        "MIB Coin",
			BinanceUSDT: "MIBUSDT",
		},
	},
	"ZOO": {
		{
			ID:   9488,
			Slug: "zookeeper",
			Name: "ZooKeeper",
		},
		{
			ID:   11020,
			Slug: "zoo-crypto-world",
			Name: "ZOO - Crypto World",
		},
	},
	"PRQ": {
		{
			ID:          5410,
			Slug:        "parsiq",
			Name:        "PARSIQ",
			BinanceUSDT: "PRQUSDT",
		},
	},
	"OPM": {
		{
			ID:          6888,
			Slug:        "omega-protocol-money",
			Name:        "Omega Protocol Money",
			BinanceUSDT: "OPMUSDT",
		},
	},
	"ELA": {
		{
			ID:          2492,
			Slug:        "elastos",
			Name:        "Elastos",
			BinanceUSDT: "ELAUSDT",
		},
	},
	"VLD": {
		{
			ID:          3492,
			Slug:        "vetri",
			Name:        "Vetri",
			BinanceUSDT: "VLDUSDT",
		},
	},
	"HATCH": {
		{
			ID:          7109,
			Slug:        "hatch-dao",
			Name:        "Hatch DAO",
			BinanceUSDT: "HATCHUSDT",
		},
	},
	"XMY": {
		{
			ID:          182,
			Slug:        "myriad",
			Name:        "Myriad",
			BinanceUSDT: "XMYUSDT",
		},
	},
	"PREMIA": {
		{
			ID:          8476,
			Slug:        "premia",
			Name:        "Premia",
			BinanceUSDT: "PREMIAUSDT",
		},
	},
	"REDPANDA": {
		{
			ID:          9980,
			Slug:        "redpanda",
			Name:        "Redpanda Earth",
			BinanceUSDT: "REDPANDAUSDT",
		},
	},
	"CRV": {
		{
			ID:          6538,
			Slug:        "curve-dao-token",
			Name:        "Curve DAO Token",
			BinanceUSDT: "CRVUSDT",
		},
	},
	"ARDR": {
		{
			ID:          1320,
			Slug:        "ardor",
			Name:        "Ardor",
			BinanceUSDT: "ARDRUSDT",
		},
	},
	"DONUT": {
		{
			ID:          6156,
			Slug:        "donut",
			Name:        "Donut",
			BinanceUSDT: "DONUTUSDT",
		},
	},
	"KMD": {
		{
			ID:          1521,
			Slug:        "komodo",
			Name:        "Komodo",
			BinanceUSDT: "KMDUSDT",
		},
	},
	"ENCX": {
		{
			ID:          6954,
			Slug:        "enceladus-network",
			Name:        "Enceladus Network",
			BinanceUSDT: "ENCXUSDT",
		},
	},
	"SLIM": {
		{
			ID:          9741,
			Slug:        "solanium",
			Name:        "Solanium",
			BinanceUSDT: "SLIMUSDT",
		},
	},
	"ACDC": {
		{
			ID:          3004,
			Slug:        "volt",
			Name:        "Volt",
			BinanceUSDT: "ACDCUSDT",
		},
	},
	"VLX": {
		{
			ID:          4747,
			Slug:        "velas",
			Name:        "Velas",
			BinanceUSDT: "VLXUSDT",
		},
	},
	"UMBR": {
		{
			ID:          8780,
			Slug:        "umbria-network",
			Name:        "Umbria Network",
			BinanceUSDT: "UMBRUSDT",
		},
	},
	"NFTBS": {
		{
			ID:          11042,
			Slug:        "nftbooks",
			Name:        "NFTBooks",
			BinanceUSDT: "NFTBSUSDT",
		},
	},
	"HUNT": {
		{
			ID:          5380,
			Slug:        "hunt",
			Name:        "HUNT",
			BinanceUSDT: "HUNTUSDT",
		},
	},
	"DF": {
		{
			ID:          4758,
			Slug:        "dforce",
			Name:        "dForce",
			BinanceUSDT: "DFUSDT",
		},
	},
	"DAWG": {
		{
			ID:          10912,
			Slug:        "inumaki",
			Name:        "Inumaki",
			BinanceUSDT: "DAWGUSDT",
		},
	},
	"HPAY": {
		{
			ID:          8532,
			Slug:        "hyper-credit-network",
			Name:        "Hyper Credit Network",
			BinanceUSDT: "HPAYUSDT",
		},
	},
	"TRIO": {
		{
			ID:          2661,
			Slug:        "tripio",
			Name:        "Tripio",
			BinanceUSDT: "TRIOUSDT",
		},
	},
	"BCEO": {
		{
			ID:          3888,
			Slug:        "bitceo",
			Name:        "bitCEO",
			BinanceUSDT: "BCEOUSDT",
		},
	},
	"USDS": {
		{
			ID:          3719,
			Slug:        "stableusd",
			Name:        "Stably USD",
			BinanceUSDT: "USDSUSDT",
		},
	},
	"FLS": {
		{
			ID:          4491,
			Slug:        "flits",
			Name:        "Flits",
			BinanceUSDT: "FLSUSDT",
		},
	},
	"SPHTX": {
		{
			ID:          2309,
			Slug:        "sophiatx",
			Name:        "SophiaTX",
			BinanceUSDT: "SPHTXUSDT",
		},
	},
	"VEX": {
		{
			ID:   2998,
			Slug: "vexanium",
			Name: "Vexanium",
		},
		{
			ID:   9798,
			Slug: "velorex",
			Name: "VELOREX",
		},
	},
	"ARTY": {
		{
			ID:          10790,
			Slug:        "artys-world",
			Name:        "Arty's World",
			BinanceUSDT: "ARTYUSDT",
		},
	},
	"JAGUAR": {
		{
			ID:          9610,
			Slug:        "jaguarswap",
			Name:        "JaguarSwap",
			BinanceUSDT: "JAGUARUSDT",
		},
	},
	"NFXC": {
		{
			ID:          5545,
			Slug:        "nfx-coin",
			Name:        "NFX Coin",
			BinanceUSDT: "NFXCUSDT",
		},
	},
	"PKOIN": {
		{
			ID:          5925,
			Slug:        "pocketnet",
			Name:        "Pkoin",
			BinanceUSDT: "PKOINUSDT",
		},
	},
	"CLUB": {
		{
			ID:          1135,
			Slug:        "clubcoin",
			Name:        "ClubCoin",
			BinanceUSDT: "CLUBUSDT",
		},
	},
	"YUI": {
		{
			ID:          7709,
			Slug:        "yui-token",
			Name:        "YUI Token",
			BinanceUSDT: "YUIUSDT",
		},
	},
	"PMGT": {
		{
			ID:          5203,
			Slug:        "perth-mint-gold-token",
			Name:        "Perth Mint Gold Token",
			BinanceUSDT: "PMGTUSDT",
		},
	},
	"LUSD": {
		{
			ID:          9566,
			Slug:        "liquity-usd",
			Name:        "Liquity USD",
			BinanceUSDT: "LUSDUSDT",
		},
	},
	"HERPES": {
		{
			ID:          10879,
			Slug:        "herpes",
			Name:        "Herpes",
			BinanceUSDT: "HERPESUSDT",
		},
	},
	"BTS": {
		{
			ID:   463,
			Slug: "bitshares",
			Name: "BitShares",
		},
		{
			ID:   8205,
			Slug: "bat-true-share",
			Name: "Bat True Share",
		},
	},
	"CXT": {
		{
			ID:          1630,
			Slug:        "coinonat",
			Name:        "Coinonat",
			BinanceUSDT: "CXTUSDT",
		},
	},
	"wxDai": {
		{
			ID:          9021,
			Slug:        "wxdai",
			Name:        "Wrapped XDAI",
			BinanceUSDT: "wxDaiUSDT",
		},
	},
	"KALM": {
		{
			ID:          10099,
			Slug:        "kalmar",
			Name:        "Kalmar",
			BinanceUSDT: "KALMUSDT",
		},
	},
	"UFO": {
		{
			ID:   10729,
			Slug: "the-truth",
			Name: "The Truth",
		},
		{
			ID:   168,
			Slug: "uniform-fiscal-object",
			Name: "Uniform Fiscal Object",
		},
		{
			ID:   5403,
			Slug: "unknown-fair-object",
			Name: "Unknown Fair Object",
		},
	},
	"ENOL": {
		{
			ID:          7782,
			Slug:        "ethanol",
			Name:        "Ethanol",
			BinanceUSDT: "ENOLUSDT",
		},
	},
	"EDR": {
		{
			ID:          2835,
			Slug:        "endor-protocol",
			Name:        "Endor Protocol",
			BinanceUSDT: "EDRUSDT",
		},
	},
	"VSN": {
		{
			ID:   5973,
			Slug: "vision-network",
			Name: "Vision Network",
		},
		{
			ID:   6734,
			Slug: "vision",
			Name: "Vision",
		},
	},
	"SXTZ": {
		{
			ID:          6192,
			Slug:        "sxtz",
			Name:        "sXTZ",
			BinanceUSDT: "SXTZUSDT",
		},
	},
	"VBZRX": {
		{
			ID:          5879,
			Slug:        "vbzrx",
			Name:        "bZx Vesting Token",
			BinanceUSDT: "VBZRXUSDT",
		},
	},
	"CAROM": {
		{
			ID:          9248,
			Slug:        "carillonium-finance",
			Name:        "Carillonium finance",
			BinanceUSDT: "CAROMUSDT",
		},
	},
	"EVE": {
		{
			ID:          2464,
			Slug:        "devery",
			Name:        "Devery",
			BinanceUSDT: "EVEUSDT",
		},
	},
	"DGN": {
		{
			ID:          8667,
			Slug:        "degen-protocol",
			Name:        "Degen Protocol",
			BinanceUSDT: "DGNUSDT",
		},
	},
	"VENA": {
		{
			ID:          4495,
			Slug:        "vena",
			Name:        "VENA",
			BinanceUSDT: "VENAUSDT",
		},
	},
	"VRAP": {
		{
			ID:          8763,
			Slug:        "veraswap",
			Name:        "VeraSwap",
			BinanceUSDT: "VRAPUSDT",
		},
	},
	"MEDIADOGE": {
		{
			ID:          10811,
			Slug:        "the-mediadoge",
			Name:        "The MEDIADOGE",
			BinanceUSDT: "MEDIADOGEUSDT",
		},
	},
	"ZANO": {
		{
			ID:          4691,
			Slug:        "zano",
			Name:        "Zano",
			BinanceUSDT: "ZANOUSDT",
		},
	},
	"ICP": {
		{
			ID:          8916,
			Slug:        "internet-computer",
			Name:        "Internet Computer",
			BinanceUSDT: "ICPUSDT",
		},
	},
	"HIVE": {
		{
			ID:          5370,
			Slug:        "hive-blockchain",
			Name:        "Hive",
			BinanceUSDT: "HIVEUSDT",
		},
	},
	"ANW": {
		{
			ID:          6120,
			Slug:        "anchor-neural-world",
			Name:        "Anchor Neural World",
			BinanceUSDT: "ANWUSDT",
		},
	},
	"ASAFE": {
		{
			ID:          1439,
			Slug:        "allsafe",
			Name:        "AllSafe",
			BinanceUSDT: "ASAFEUSDT",
		},
	},
	"TESLF": {
		{
			ID:          10035,
			Slug:        "teslafan",
			Name:        "Teslafan",
			BinanceUSDT: "TESLFUSDT",
		},
	},
	"OLT": {
		{
			ID:          2921,
			Slug:        "oneledger",
			Name:        "OneLedger",
			BinanceUSDT: "OLTUSDT",
		},
	},
	"DEB": {
		{
			ID:          2584,
			Slug:        "debitum-network",
			Name:        "Debitum",
			BinanceUSDT: "DEBUSDT",
		},
	},
	"BBS": {
		{
			ID:          3093,
			Slug:        "bbscoin",
			Name:        "BBSCoin",
			BinanceUSDT: "BBSUSDT",
		},
	},
	"NII": {
		{
			ID:          4865,
			Slug:        "nahmii",
			Name:        "Nahmii",
			BinanceUSDT: "NIIUSDT",
		},
	},
	"SBS": {
		{
			ID:          8283,
			Slug:        "staysbase",
			Name:        "StaysBASE",
			BinanceUSDT: "SBSUSDT",
		},
	},
	"UTED": {
		{
			ID:          7462,
			Slug:        "united",
			Name:        "United",
			BinanceUSDT: "UTEDUSDT",
		},
	},
	"XPC": {
		{
			ID:          3750,
			Slug:        "experience-chain",
			Name:        "eXPerience Chain",
			BinanceUSDT: "XPCUSDT",
		},
	},
	"XCM": {
		{
			ID:          4105,
			Slug:        "coinmetro-token",
			Name:        "CoinMetro Token",
			BinanceUSDT: "XCMUSDT",
		},
	},
	"STOPELON": {
		{
			ID:          10426,
			Slug:        "stopelon",
			Name:        "Stopelon",
			BinanceUSDT: "STOPELONUSDT",
		},
	},
	"KEK": {
		{
			ID:          8135,
			Slug:        "cryptokek",
			Name:        "CryptoKek",
			BinanceUSDT: "KEKUSDT",
		},
	},
	"RAINI": {
		{
			ID:          9061,
			Slug:        "rainicorn",
			Name:        "Rainicorn",
			BinanceUSDT: "RAINIUSDT",
		},
	},
	"BTC3S": {
		{
			ID:          5737,
			Slug:        "amun-bitcoin-3x-daily-short",
			Name:        "Amun Bitcoin 3x Daily Short",
			BinanceUSDT: "BTC3SUSDT",
		},
	},
	"AER": {
		{
			ID:   4359,
			Slug: "aeryus",
			Name: "Aeryus",
		},
		{
			ID:   10208,
			Slug: "aerdrop",
			Name: "Aerdrop",
		},
	},
	"BRD": {
		{
			ID:          2306,
			Slug:        "bread",
			Name:        "Bread",
			BinanceUSDT: "BRDUSDT",
		},
	},
	"XMV": {
		{
			ID:          3902,
			Slug:        "monero-v",
			Name:        "MoneroV ",
			BinanceUSDT: "XMVUSDT",
		},
	},
	"NYAN": {
		{
			ID:   7020,
			Slug: "nyan-finance",
			Name: "Nyan Finance",
		},
		{
			ID:   8473,
			Slug: "yieldnyan",
			Name: "YieldNyan",
		},
	},
	"CONI": {
		{
			ID:          2895,
			Slug:        "coni",
			Name:        "Coni",
			BinanceUSDT: "CONIUSDT",
		},
	},
	"UDT": {
		{
			ID:          9364,
			Slug:        "unlock-protocol",
			Name:        "Unlock Protocol",
			BinanceUSDT: "UDTUSDT",
		},
	},
	"BOB": {
		{
			ID:          2889,
			Slug:        "bobs-repair",
			Name:        "Bob's Repair",
			BinanceUSDT: "BOBUSDT",
		},
	},
	"KWATT": {
		{
			ID:          3209,
			Slug:        "4new",
			Name:        "4NEW",
			BinanceUSDT: "KWATTUSDT",
		},
	},
	"MARTK": {
		{
			ID:          5520,
			Slug:        "martkist",
			Name:        "Martkist",
			BinanceUSDT: "MARTKUSDT",
		},
	},
	"EXOR": {
		{
			ID:          4126,
			Slug:        "exor",
			Name:        "EXOR",
			BinanceUSDT: "EXORUSDT",
		},
	},
	"HUSKY": {
		{
			ID:          9387,
			Slug:        "husky",
			Name:        "Husky",
			BinanceUSDT: "HUSKYUSDT",
		},
	},
	"ASNX": {
		{
			ID:          5752,
			Slug:        "aave-snx",
			Name:        "Aave SNX",
			BinanceUSDT: "ASNXUSDT",
		},
	},
	"COW": {
		{
			ID:          10011,
			Slug:        "coinwind",
			Name:        "CoinWind",
			BinanceUSDT: "COWUSDT",
		},
	},
	"BITN": {
		{
			ID:          4914,
			Slug:        "bitcoin-and-company-network",
			Name:        "Bitcoin & Company Network",
			BinanceUSDT: "BITNUSDT",
		},
	},
	"YUMMY": {
		{
			ID:          9859,
			Slug:        "yummy",
			Name:        "YUMMY",
			BinanceUSDT: "YUMMYUSDT",
		},
	},
	"QCX": {
		{
			ID:          3883,
			Slug:        "quickx-protocol",
			Name:        "QuickX Protocol",
			BinanceUSDT: "QCXUSDT",
		},
	},
	"NEWW": {
		{
			ID:          9084,
			Slug:        "newv-finance",
			Name:        "New Ventures",
			BinanceUSDT: "NEWWUSDT",
		},
	},
	"AWS": {
		{
			ID:          9547,
			Slug:        "aurussilver",
			Name:        "AurusSILVER",
			BinanceUSDT: "AWSUSDT",
		},
	},
	"BAO": {
		{
			ID:          8168,
			Slug:        "bao-finance",
			Name:        "Bao Finance",
			BinanceUSDT: "BAOUSDT",
		},
	},
	"BBYXRP": {
		{
			ID:          10872,
			Slug:        "babyxrp",
			Name:        "BABYXRP",
			BinanceUSDT: "BBYXRPUSDT",
		},
	},
	"FAIR": {
		{
			ID:   2366,
			Slug: "fairgame",
			Name: "FairGame",
		},
		{
			ID:   224,
			Slug: "faircoin",
			Name: "FairCoin",
		},
	},
	"FFA": {
		{
			ID:          9659,
			Slug:        "cryptofifa",
			Name:        "Cryptofifa",
			BinanceUSDT: "FFAUSDT",
		},
	},
	"SUP8EME": {
		{
			ID:          7701,
			Slug:        "sup8eme",
			Name:        "SUP8EME",
			BinanceUSDT: "SUP8EMEUSDT",
		},
	},
	"SWTH": {
		{
			ID:          2620,
			Slug:        "switcheo",
			Name:        "Switcheo",
			BinanceUSDT: "SWTHUSDT",
		},
	},
	"SHAKE": {
		{
			ID:          7390,
			Slug:        "shake",
			Name:        "Spaceswap SHAKE",
			BinanceUSDT: "SHAKEUSDT",
		},
	},
	"FLA": {
		{
			ID:          7609,
			Slug:        "fiola",
			Name:        "Fiola",
			BinanceUSDT: "FLAUSDT",
		},
	},
	"XVIX": {
		{
			ID:          7969,
			Slug:        "xvix",
			Name:        "XVIX",
			BinanceUSDT: "XVIXUSDT",
		},
	},
	"ACH": {
		{
			ID:          6958,
			Slug:        "alchemy-pay",
			Name:        "Alchemy Pay",
			BinanceUSDT: "ACHUSDT",
		},
	},
	"KUBO": {
		{
			ID:          3901,
			Slug:        "kubocoin",
			Name:        "KuboCoin",
			BinanceUSDT: "KUBOUSDT",
		},
	},
	"AMM": {
		{
			ID:          2286,
			Slug:        "micromoney",
			Name:        "MicroMoney",
			BinanceUSDT: "AMMUSDT",
		},
	},
	"OPCT": {
		{
			ID:          3632,
			Slug:        "opacity",
			Name:        "Opacity",
			BinanceUSDT: "OPCTUSDT",
		},
	},
	"ETHUP": {
		{
			ID:          7016,
			Slug:        "ethup",
			Name:        "ETHUP",
			BinanceUSDT: "ETHUPUSDT",
		},
	},
	"XBY": {
		{
			ID:          1636,
			Slug:        "xtrabytes",
			Name:        "XTRABYTES",
			BinanceUSDT: "XBYUSDT",
		},
	},
	"KEYT": {
		{
			ID:          5530,
			Slug:        "rebit",
			Name:        "REBIT",
			BinanceUSDT: "KEYTUSDT",
		},
	},
	"MINTY": {
		{
			ID:          8694,
			Slug:        "minty-art",
			Name:        "Minty Art",
			BinanceUSDT: "MINTYUSDT",
		},
	},
	"OWN": {
		{
			ID:          3120,
			Slug:        "owndata",
			Name:        "OWNDATA",
			BinanceUSDT: "OWNUSDT",
		},
	},
	"DYNMT": {
		{
			ID:          4193,
			Slug:        "dynamite",
			Name:        "Dynamite",
			BinanceUSDT: "DYNMTUSDT",
		},
	},
	"AEON": {
		{
			ID:          1026,
			Slug:        "aeon",
			Name:        "Aeon",
			BinanceUSDT: "AEONUSDT",
		},
	},
	"WXMR": {
		{
			ID:          8235,
			Slug:        "wrapped-monero",
			Name:        "Wrapped Monero",
			BinanceUSDT: "WXMRUSDT",
		},
	},
	"ISLA": {
		{
			ID:          5389,
			Slug:        "insula",
			Name:        "Insula",
			BinanceUSDT: "ISLAUSDT",
		},
	},
	"MPT": {
		{
			ID:   6769,
			Slug: "money-plant-token",
			Name: "Money Plant Token",
		},
		{
			ID:   8769,
			Slug: "meetple",
			Name: "MeetPle",
		},
	},
	"SOAR": {
		{
			ID:          8240,
			Slug:        "soar-fi",
			Name:        "SOAR.FI",
			BinanceUSDT: "SOARUSDT",
		},
	},
	"PYX": {
		{
			ID:          9168,
			Slug:        "pyxis-network",
			Name:        "PYXIS Network",
			BinanceUSDT: "PYXUSDT",
		},
	},
	"POLA": {
		{
			ID:          7740,
			Slug:        "polaris-share",
			Name:        "Polaris Share",
			BinanceUSDT: "POLAUSDT",
		},
	},
	"BETH": {
		{
			ID:          8353,
			Slug:        "beacon-eth",
			Name:        "Beacon ETH",
			BinanceUSDT: "BETHUSDT",
		},
	},
	"PWAR": {
		{
			ID:          10748,
			Slug:        "polkawar",
			Name:        "PolkaWar",
			BinanceUSDT: "PWARUSDT",
		},
	},
	"WBIND": {
		{
			ID:          7403,
			Slug:        "wrapped-bind",
			Name:        "Wrapped BIND",
			BinanceUSDT: "WBINDUSDT",
		},
	},
	"eMTRG": {
		{
			ID:          6628,
			Slug:        "meter-governance-mapped-by-meter-io",
			Name:        "Meter Governance mapped by Meter.io",
			BinanceUSDT: "eMTRGUSDT",
		},
	},
	"LYM": {
		{
			ID:          2554,
			Slug:        "lympo",
			Name:        "Lympo",
			BinanceUSDT: "LYMUSDT",
		},
	},
	"XTA": {
		{
			ID:          3682,
			Slug:        "italo",
			Name:        "Italo",
			BinanceUSDT: "XTAUSDT",
		},
	},
	"LBD": {
		{
			ID:          8138,
			Slug:        "linkbased",
			Name:        "LinkBased",
			BinanceUSDT: "LBDUSDT",
		},
	},
	"ENTONE": {
		{
			ID:          5811,
			Slug:        "entone",
			Name:        "ENTONE",
			BinanceUSDT: "ENTONEUSDT",
		},
	},
	"IMS": {
		{
			ID:          1194,
			Slug:        "independent-money-system",
			Name:        "Independent Money System",
			BinanceUSDT: "IMSUSDT",
		},
	},
	"GFI": {
		{
			ID:   10324,
			Slug: "gravity-finance",
			Name: "Gravity Finance",
		},
		{
			ID:   9049,
			Slug: "gorilla-fi",
			Name: "Gorilla-Fi",
		},
	},
	"EURS": {
		{
			ID:          2989,
			Slug:        "stasis-euro",
			Name:        "STASIS EURO",
			BinanceUSDT: "EURSUSDT",
		},
	},
	"YFB2": {
		{
			ID:          7612,
			Slug:        "yearn-finance-bit2",
			Name:        "Yearn Finance Bit2",
			BinanceUSDT: "YFB2USDT",
		},
	},
	"GCC": {
		{
			ID:   1531,
			Slug: "global-cryptocurrency",
			Name: "Global Cryptocurrency",
		},
		{
			ID:   1033,
			Slug: "guccionecoin",
			Name: "GuccioneCoin",
		},
	},
	"GOVI": {
		{
			ID:          8408,
			Slug:        "govi",
			Name:        "Govi",
			BinanceUSDT: "GOVIUSDT",
		},
	},
	"CCA": {
		{
			ID:          4122,
			Slug:        "counos-coin",
			Name:        "Counos Coin",
			BinanceUSDT: "CCAUSDT",
		},
	},
	"PEPE": {
		{
			ID:          10268,
			Slug:        "pepemoon",
			Name:        "PepeMoon",
			BinanceUSDT: "PEPEUSDT",
		},
	},
	"MOK": {
		{
			ID:          10819,
			Slug:        "mocktailswap",
			Name:        "MocktailSwap",
			BinanceUSDT: "MOKUSDT",
		},
	},
	"KOMBAI": {
		{
			ID:          10319,
			Slug:        "kombai-inu",
			Name:        "Kombai Inu",
			BinanceUSDT: "KOMBAIUSDT",
		},
	},
	"SUR": {
		{
			ID:          1933,
			Slug:        "suretly",
			Name:        "Suretly",
			BinanceUSDT: "SURUSDT",
		},
	},
	"BNBBEAR": {
		{
			ID:          5294,
			Slug:        "3x-short-bnb-token",
			Name:        "3X Short BNB Token",
			BinanceUSDT: "BNBBEARUSDT",
		},
	},
	"SNOW": {
		{
			ID:          7367,
			Slug:        "snowswap",
			Name:        "SnowSwap",
			BinanceUSDT: "SNOWUSDT",
		},
	},
	"SOLVE": {
		{
			ID:          3724,
			Slug:        "solve",
			Name:        "SOLVE",
			BinanceUSDT: "SOLVEUSDT",
		},
	},
	"RAK": {
		{
			ID:          7632,
			Slug:        "rake-finance",
			Name:        "Rake Finance",
			BinanceUSDT: "RAKUSDT",
		},
	},
	"TMTG": {
		{
			ID:          3356,
			Slug:        "the-midas-touch-gold",
			Name:        "The Midas Touch Gold",
			BinanceUSDT: "TMTGUSDT",
		},
	},
	"OTO": {
		{
			ID:          3850,
			Slug:        "otocash",
			Name:        "OTOCASH",
			BinanceUSDT: "OTOUSDT",
		},
	},
	"CRFI": {
		{
			ID:          9535,
			Slug:        "crossfi",
			Name:        "CrossFi",
			BinanceUSDT: "CRFIUSDT",
		},
	},
	"WARP": {
		{
			ID:          8439,
			Slug:        "warp-finance",
			Name:        "Warp Finance",
			BinanceUSDT: "WARPUSDT",
		},
	},
	"RCUBE": {
		{
			ID:          10590,
			Slug:        "retro-defi-rcube",
			Name:        "RETRO DEFI - RCUBE",
			BinanceUSDT: "RCUBEUSDT",
		},
	},
	"KSS": {
		{
			ID:          6655,
			Slug:        "krosscoin",
			Name:        "Krosscoin",
			BinanceUSDT: "KSSUSDT",
		},
	},
	"PAXEX": {
		{
			ID:          3338,
			Slug:        "paxex",
			Name:        "PAXEX",
			BinanceUSDT: "PAXEXUSDT",
		},
	},
	"HDAC": {
		{
			ID:          2999,
			Slug:        "hdac",
			Name:        "Hdac",
			BinanceUSDT: "HDACUSDT",
		},
	},
	"TEST": {
		{
			ID:          9911,
			Slug:        "test-token",
			Name:        "Test Token",
			BinanceUSDT: "TESTUSDT",
		},
	},
	"NOIZ": {
		{
			ID:          4125,
			Slug:        "noizchain",
			Name:        "NOIZ",
			BinanceUSDT: "NOIZUSDT",
		},
	},
	"DOTDOWN": {
		{
			ID:          7006,
			Slug:        "dotdown",
			Name:        "DOTDOWN",
			BinanceUSDT: "DOTDOWNUSDT",
		},
	},
	"VEO": {
		{
			ID:          3716,
			Slug:        "amoveo",
			Name:        "Amoveo",
			BinanceUSDT: "VEOUSDT",
		},
	},
	"RAM": {
		{
			ID:          8798,
			Slug:        "ramifi-protocol",
			Name:        "Ramifi Protocol",
			BinanceUSDT: "RAMUSDT",
		},
	},
	"BabyDoge": {
		{
			ID:          10407,
			Slug:        "baby-doge-coin",
			Name:        "Baby Doge Coin",
			BinanceUSDT: "BabyDogeUSDT",
		},
	},
	"GGTK": {
		{
			ID:          8156,
			Slug:        "ggdapp",
			Name:        "GGDApp",
			BinanceUSDT: "GGTKUSDT",
		},
	},
	"NLG": {
		{
			ID:          254,
			Slug:        "gulden",
			Name:        "Gulden",
			BinanceUSDT: "NLGUSDT",
		},
	},
	"KAR": {
		{
			ID:          10042,
			Slug:        "karura",
			Name:        "Karura",
			BinanceUSDT: "KARUSDT",
		},
	},
	"SFT": {
		{
			ID:          1172,
			Slug:        "safex-token",
			Name:        "Safex Token",
			BinanceUSDT: "SFTUSDT",
		},
	},
	"YFSI": {
		{
			ID:          7343,
			Slug:        "yfscience",
			Name:        "Yfscience",
			BinanceUSDT: "YFSIUSDT",
		},
	},
	"OMNI": {
		{
			ID:          83,
			Slug:        "omni",
			Name:        "Omni",
			BinanceUSDT: "OMNIUSDT",
		},
	},
	"CCAI": {
		{
			ID:          10935,
			Slug:        "cryptocurrencies-ai",
			Name:        "Cryptocurrencies.ai",
			BinanceUSDT: "CCAIUSDT",
		},
	},
	"ELK": {
		{
			ID:          10095,
			Slug:        "elk-finance",
			Name:        "Elk Finance",
			BinanceUSDT: "ELKUSDT",
		},
	},
	"SAFEBULL": {
		{
			ID:          10644,
			Slug:        "safebull",
			Name:        "SafeBull",
			BinanceUSDT: "SAFEBULLUSDT",
		},
	},
	"PORTAL": {
		{
			ID:          6496,
			Slug:        "portal",
			Name:        "Portal",
			BinanceUSDT: "PORTALUSDT",
		},
	},
	"POP": {
		{
			ID:   275,
			Slug: "popularcoin",
			Name: "PopularCoin",
		},
		{
			ID:   7363,
			Slug: "pop-network-token",
			Name: "POP Network Token",
		},
	},
	"ETHA": {
		{
			ID:          8709,
			Slug:        "etha-lend",
			Name:        "ETHA Lend",
			BinanceUSDT: "ETHAUSDT",
		},
	},
	"PNIXS": {
		{
			ID:          9641,
			Slug:        "phoenxidefi-finance",
			Name:        "PhoenxiDefi Finance",
			BinanceUSDT: "PNIXSUSDT",
		},
	},
	"TKS": {
		{
			ID:          1588,
			Slug:        "tokes",
			Name:        "Tokes",
			BinanceUSDT: "TKSUSDT",
		},
	},
	"PAPER": {
		{
			ID:          8847,
			Slug:        "fomo-app",
			Name:        "Fomo App",
			BinanceUSDT: "PAPERUSDT",
		},
	},
	"IOTX": {
		{
			ID:          2777,
			Slug:        "iotex",
			Name:        "IoTeX",
			BinanceUSDT: "IOTXUSDT",
		},
	},
	"ETHRSIAPY": {
		{
			ID:          6144,
			Slug:        "eth-rsi-60-40-yield-set",
			Name:        "ETH RSI 60/40 Yield Set",
			BinanceUSDT: "ETHRSIAPYUSDT",
		},
	},
	"DRM": {
		{
			ID:   316,
			Slug: "dreamcoin",
			Name: "Dreamcoin",
		},
		{
			ID:   5837,
			Slug: "dodreamchain",
			Name: "DoDreamChain",
		},
	},
	"TOWER": {
		{
			ID:          8620,
			Slug:        "tower-token",
			Name:        "Tower",
			BinanceUSDT: "TOWERUSDT",
		},
	},
	"SUKU": {
		{
			ID:          6180,
			Slug:        "suku",
			Name:        "SUKU",
			BinanceUSDT: "SUKUUSDT",
		},
	},
	"TKINU": {
		{
			ID:          10024,
			Slug:        "tsuki-inu",
			Name:        "Tsuki Inu",
			BinanceUSDT: "TKINUUSDT",
		},
	},
	"ZORT": {
		{
			ID:          10830,
			Slug:        "zort",
			Name:        "ZORT",
			BinanceUSDT: "ZORTUSDT",
		},
	},
	"ONT": {
		{
			ID:          2566,
			Slug:        "ontology",
			Name:        "Ontology",
			BinanceUSDT: "ONTUSDT",
		},
	},
	"ELX": {
		{
			ID:          8331,
			Slug:        "energy-ledger",
			Name:        "Energy Ledger",
			BinanceUSDT: "ELXUSDT",
		},
	},
	"VIDT": {
		{
			ID:          3845,
			Slug:        "vidt-datalink",
			Name:        "VIDT Datalink",
			BinanceUSDT: "VIDTUSDT",
		},
	},
	"ROOK": {
		{
			ID:          7678,
			Slug:        "keeperdao",
			Name:        "KeeperDAO",
			BinanceUSDT: "ROOKUSDT",
		},
	},
	"FIRE": {
		{
			ID:   5933,
			Slug: "fireball",
			Name: "Fireball",
		},
		{
			ID:   10089,
			Slug: "fire-token",
			Name: "Fire Token",
		},
		{
			ID:   8129,
			Slug: "fire-protocol",
			Name: "Fire Protocol",
		},
	},
	"BPCN": {
		{
			ID:          6242,
			Slug:        "blipcoin",
			Name:        "BlipCoin",
			BinanceUSDT: "BPCNUSDT",
		},
	},
	"XCN": {
		{
			ID:          501,
			Slug:        "cryptonite",
			Name:        "Cryptonite",
			BinanceUSDT: "XCNUSDT",
		},
	},
	"NPC": {
		{
			ID:          4110,
			Slug:        "npcoin",
			Name:        "NPCoin",
			BinanceUSDT: "NPCUSDT",
		},
	},
	"YAM": {
		{
			ID:   7131,
			Slug: "yam",
			Name: "YAM",
		},
		{
			ID:   6539,
			Slug: "yamv1",
			Name: "YAM v1",
		},
	},
	"GROOT": {
		{
			ID:          8604,
			Slug:        "growth-root-token",
			Name:        "growth Root Token",
			BinanceUSDT: "GROOTUSDT",
		},
	},
	"FOMP": {
		{
			ID:          7936,
			Slug:        "fompound",
			Name:        "FOMPOUND",
			BinanceUSDT: "FOMPUSDT",
		},
	},
	"MEGA": {
		{
			ID:          7540,
			Slug:        "megacryptopolis",
			Name:        "MegaCryptoPolis",
			BinanceUSDT: "MEGAUSDT",
		},
	},
	"MCH": {
		{
			ID:          4844,
			Slug:        "meconcash",
			Name:        "MeconCash",
			BinanceUSDT: "MCHUSDT",
		},
	},
	"EBK": {
		{
			ID:          4778,
			Slug:        "ebakus",
			Name:        "ebakus",
			BinanceUSDT: "EBKUSDT",
		},
	},
	"INX": {
		{
			ID:          5840,
			Slug:        "insight-protocol",
			Name:        "Insight Protocol",
			BinanceUSDT: "INXUSDT",
		},
	},
	"SRK": {
		{
			ID:          3935,
			Slug:        "sparkpoint",
			Name:        "SparkPoint",
			BinanceUSDT: "SRKUSDT",
		},
	},
	"MX": {
		{
			ID:          4041,
			Slug:        "mx-token",
			Name:        "MX Token",
			BinanceUSDT: "MXUSDT",
		},
	},
	"ECU": {
		{
			ID:   5478,
			Slug: "ecosc",
			Name: "ECOSC",
		},
		{
			ID:   5884,
			Slug: "decurian",
			Name: "Decurian",
		},
	},
	"CYF": {
		{
			ID:          7179,
			Slug:        "cy-finance",
			Name:        "CY Finance",
			BinanceUSDT: "CYFUSDT",
		},
	},
	"SENSE": {
		{
			ID:          2402,
			Slug:        "sense",
			Name:        "Sense",
			BinanceUSDT: "SENSEUSDT",
		},
	},
	"HDP.ф": {
		{
			ID:          4571,
			Slug:        "hedpay",
			Name:        "HEdpAY",
			BinanceUSDT: "HDP.фUSDT",
		},
	},
	"ZUC": {
		{
			ID:          4250,
			Slug:        "zeuxcoin",
			Name:        "ZeuxCoin",
			BinanceUSDT: "ZUCUSDT",
		},
	},
	"WOOFY": {
		{
			ID:          9719,
			Slug:        "woofy",
			Name:        "Woofy",
			BinanceUSDT: "WOOFYUSDT",
		},
	},
	"SVT": {
		{
			ID:          10493,
			Slug:        "spacevikings",
			Name:        "SpaceVikings",
			BinanceUSDT: "SVTUSDT",
		},
	},
	"BLZ": {
		{
			ID:          2505,
			Slug:        "bluzelle",
			Name:        "Bluzelle",
			BinanceUSDT: "BLZUSDT",
		},
	},
	"CONV": {
		{
			ID:          8716,
			Slug:        "convergence",
			Name:        "Convergence",
			BinanceUSDT: "CONVUSDT",
		},
	},
	"INEX": {
		{
			ID:          5300,
			Slug:        "inex-project",
			Name:        "Inex Project",
			BinanceUSDT: "INEXUSDT",
		},
	},
	"XBE": {
		{
			ID:          9092,
			Slug:        "xbe-token",
			Name:        "XBE Token",
			BinanceUSDT: "XBEUSDT",
		},
	},
	"PBT": {
		{
			ID:          1841,
			Slug:        "primalbase",
			Name:        "Primalbase Token",
			BinanceUSDT: "PBTUSDT",
		},
	},
	"SPICE": {
		{
			ID:   4259,
			Slug: "spice",
			Name: "Spice",
		},
		{
			ID:   8023,
			Slug: "scifi-finance",
			Name: "SPICE",
		},
	},
	"CPC": {
		{
			ID:   2482,
			Slug: "cpchain",
			Name: "CPChain",
		},
		{
			ID:   9091,
			Slug: "cpcoin",
			Name: "CPCoin",
		},
	},
	"CBIX-P": {
		{
			ID:          9512,
			Slug:        "cubiex-power",
			Name:        "Cubiex Power",
			BinanceUSDT: "CBIX-PUSDT",
		},
	},
	"OLXA": {
		{
			ID:          3823,
			Slug:        "olxa",
			Name:        "OLXA",
			BinanceUSDT: "OLXAUSDT",
		},
	},
	"TWERK": {
		{
			ID:          9893,
			Slug:        "twerk-finance",
			Name:        "Twerk Finance",
			BinanceUSDT: "TWERKUSDT",
		},
	},
	"GRID": {
		{
			ID:          2134,
			Slug:        "grid",
			Name:        "Grid+",
			BinanceUSDT: "GRIDUSDT",
		},
	},
	"IDX": {
		{
			ID:          5898,
			Slug:        "index-chain",
			Name:        "Index Chain",
			BinanceUSDT: "IDXUSDT",
		},
	},
	"VCCO": {
		{
			ID:          7341,
			Slug:        "vera-cruz-coin",
			Name:        "Vera Cruz Coin",
			BinanceUSDT: "VCCOUSDT",
		},
	},
	"ROSE": {
		{
			ID:          7653,
			Slug:        "oasis-network",
			Name:        "Oasis Network",
			BinanceUSDT: "ROSEUSDT",
		},
	},
	"VD": {
		{
			ID:          4119,
			Slug:        "vindax-coin",
			Name:        "VinDax Coin",
			BinanceUSDT: "VDUSDT",
		},
	},
	"FIRO": {
		{
			ID:          1414,
			Slug:        "firo",
			Name:        "Firo",
			BinanceUSDT: "FIROUSDT",
		},
	},
	"VBETH": {
		{
			ID:          8370,
			Slug:        "venus-beth",
			Name:        "Venus BETH",
			BinanceUSDT: "VBETHUSDT",
		},
	},
	"FOXX": {
		{
			ID:          10387,
			Slug:        "star-foxx",
			Name:        "Star Foxx",
			BinanceUSDT: "FOXXUSDT",
		},
	},
	"ESS": {
		{
			ID:          2906,
			Slug:        "essentia",
			Name:        "Essentia",
			BinanceUSDT: "ESSUSDT",
		},
	},
	"IBP": {
		{
			ID:          6625,
			Slug:        "innovation-blockchain-payment",
			Name:        "Innovation Blockchain Payment",
			BinanceUSDT: "IBPUSDT",
		},
	},
	"GDILDO": {
		{
			ID:          10618,
			Slug:        "green-dildo-finance",
			Name:        "Green Dildo Finance",
			BinanceUSDT: "GDILDOUSDT",
		},
	},
	"WAXE": {
		{
			ID:          8136,
			Slug:        "waxe",
			Name:        "WAXE",
			BinanceUSDT: "WAXEUSDT",
		},
	},
	"I9C": {
		{
			ID:          6966,
			Slug:        "i9-coin",
			Name:        "i9 Coin",
			BinanceUSDT: "I9CUSDT",
		},
	},
	"LTCBULL": {
		{
			ID:          5462,
			Slug:        "3x-long-litecoin-token",
			Name:        "3x Long Litecoin Token",
			BinanceUSDT: "LTCBULLUSDT",
		},
	},
	"ELE": {
		{
			ID:          10271,
			Slug:        "eleven-finance",
			Name:        "Eleven Finance",
			BinanceUSDT: "ELEUSDT",
		},
	},
	"LUNA": {
		{
			ID:   4172,
			Slug: "terra-luna",
			Name: "Terra",
		},
		{
			ID:   1496,
			Slug: "luna-coin",
			Name: "Luna Coin",
		},
	},
	"DOGE": {
		{
			ID:          74,
			Slug:        "dogecoin",
			Name:        "Dogecoin",
			BinanceUSDT: "DOGEUSDT",
		},
	},
	"SLT": {
		{
			ID:          2471,
			Slug:        "smartlands-network",
			Name:        "Smartlands Network",
			BinanceUSDT: "SLTUSDT",
		},
	},
	"AUSCM": {
		{
			ID:          7583,
			Slug:        "auric-network",
			Name:        "Auric Network",
			BinanceUSDT: "AUSCMUSDT",
		},
	},
	"ZEFU": {
		{
			ID:          7430,
			Slug:        "zenfuse",
			Name:        "Zenfuse",
			BinanceUSDT: "ZEFUUSDT",
		},
	},
	"CHZ": {
		{
			ID:          4066,
			Slug:        "chiliz",
			Name:        "Chiliz",
			BinanceUSDT: "CHZUSDT",
		},
	},
	"777": {
		{
			ID:          7305,
			Slug:        "jackpot",
			Name:        "Jackpot",
			BinanceUSDT: "777USDT",
		},
	},
	"TRAC": {
		{
			ID:          2467,
			Slug:        "origintrail",
			Name:        "OriginTrail",
			BinanceUSDT: "TRACUSDT",
		},
	},
	"INDEX": {
		{
			ID:          7336,
			Slug:        "index-cooperative",
			Name:        "Index Cooperative",
			BinanceUSDT: "INDEXUSDT",
		},
	},
	"OUSD": {
		{
			ID:          7189,
			Slug:        "origin-dollar",
			Name:        "Origin Dollar",
			BinanceUSDT: "OUSDUSDT",
		},
	},
	"dART": {
		{
			ID:          9098,
			Slug:        "dart-insurance",
			Name:        "dART Insurance",
			BinanceUSDT: "dARTUSDT",
		},
	},
	"ABL": {
		{
			ID:          3156,
			Slug:        "airbloc",
			Name:        "Airbloc",
			BinanceUSDT: "ABLUSDT",
		},
	},
	"OVR": {
		{
			ID:          8144,
			Slug:        "ovr",
			Name:        "OVR",
			BinanceUSDT: "OVRUSDT",
		},
	},
	"KRL": {
		{
			ID:          2949,
			Slug:        "kryll",
			Name:        "Kryll",
			BinanceUSDT: "KRLUSDT",
		},
	},
	"CXC": {
		{
			ID:          4955,
			Slug:        "capital-x-cell",
			Name:        "CAPITAL X CELL",
			BinanceUSDT: "CXCUSDT",
		},
	},
	"WOOF": {
		{
			ID:          10857,
			Slug:        "shibance",
			Name:        "Shibance",
			BinanceUSDT: "WOOFUSDT",
		},
	},
	"AAB": {
		{
			ID:          5509,
			Slug:        "aax-token",
			Name:        "AAX Token",
			BinanceUSDT: "AABUSDT",
		},
	},
	"MBONK": {
		{
			ID:          6116,
			Slug:        "bonk",
			Name:        "megaBONK",
			BinanceUSDT: "MBONKUSDT",
		},
	},
	"CWS": {
		{
			ID:          8365,
			Slug:        "crowns",
			Name:        "Crowns",
			BinanceUSDT: "CWSUSDT",
		},
	},
	"COM": {
		{
			ID:          6608,
			Slug:        "community-token",
			Name:        "Community Token",
			BinanceUSDT: "COMUSDT",
		},
	},
	"MDU": {
		{
			ID:          6028,
			Slug:        "mdu",
			Name:        "MDUKEY",
			BinanceUSDT: "MDUUSDT",
		},
	},
	"LVN": {
		{
			ID:          5853,
			Slug:        "livenpay",
			Name:        "LivenPay",
			BinanceUSDT: "LVNUSDT",
		},
	},
	"SLTH": {
		{
			ID:          10596,
			Slug:        "slothi",
			Name:        "SLOTHI",
			BinanceUSDT: "SLTHUSDT",
		},
	},
	"CTK": {
		{
			ID:          4807,
			Slug:        "certik",
			Name:        "CertiK",
			BinanceUSDT: "CTKUSDT",
		},
	},
	"KPOP": {
		{
			ID:          10310,
			Slug:        "kpop-fan-token",
			Name:        "KPOP Fan Token",
			BinanceUSDT: "KPOPUSDT",
		},
	},
	"MALINK": {
		{
			ID:          8923,
			Slug:        "matic-aave-link",
			Name:        "Matic Aave Interest Bearing LINK",
			BinanceUSDT: "MALINKUSDT",
		},
	},
	"IDEA": {
		{
			ID:          7681,
			Slug:        "ideaology",
			Name:        "Ideaology",
			BinanceUSDT: "IDEAUSDT",
		},
	},
	"MCAN": {
		{
			ID:          8559,
			Slug:        "medican-coin",
			Name:        "Medican Coin",
			BinanceUSDT: "MCANUSDT",
		},
	},
	"PIS": {
		{
			ID:          8169,
			Slug:        "polkainsure-finance",
			Name:        "Polkainsure Finance",
			BinanceUSDT: "PISUSDT",
		},
	},
	"MQL": {
		{
			ID:          6804,
			Slug:        "miraqle",
			Name:        "MiraQle",
			BinanceUSDT: "MQLUSDT",
		},
	},
	"zUSD": {
		{
			ID:          10845,
			Slug:        "zerogoki-usd",
			Name:        "Zerogoki USD",
			BinanceUSDT: "zUSDUSDT",
		},
	},
	"USDU": {
		{
			ID:          6907,
			Slug:        "upper-dollar",
			Name:        "Upper Dollar",
			BinanceUSDT: "USDUUSDT",
		},
	},
	"OKS": {
		{
			ID:   5775,
			Slug: "oikos",
			Name: "Oikos",
		},
		{
			ID:   5272,
			Slug: "okschain",
			Name: "Okschain",
		},
	},
	"MASH": {
		{
			ID:          9352,
			Slug:        "marshmallowdefi",
			Name:        "Marshmallowdefi",
			BinanceUSDT: "MASHUSDT",
		},
	},
	"CRBN": {
		{
			ID:          7809,
			Slug:        "carbon",
			Name:        "Carbon",
			BinanceUSDT: "CRBNUSDT",
		},
	},
	"HC": {
		{
			ID:          1903,
			Slug:        "hypercash",
			Name:        "HyperCash",
			BinanceUSDT: "HCUSDT",
		},
	},
	"LIT": {
		{
			ID:   6833,
			Slug: "litentry",
			Name: "Litentry",
		},
		{
			ID:   3870,
			Slug: "lition",
			Name: "Lition",
		},
	},
	"PRISM": {
		{
			ID:          9788,
			Slug:        "prism-network",
			Name:        "Prism Network",
			BinanceUSDT: "PRISMUSDT",
		},
	},
	"ASR": {
		{
			ID:          5229,
			Slug:        "as-roma-fan-token",
			Name:        "AS Roma Fan Token",
			BinanceUSDT: "ASRUSDT",
		},
	},
	"MAGI": {
		{
			ID:          8914,
			Slug:        "magikarp-finance",
			Name:        "Magikarp Finance",
			BinanceUSDT: "MAGIUSDT",
		},
	},
	"WAVES": {
		{
			ID:          1274,
			Slug:        "waves",
			Name:        "Waves",
			BinanceUSDT: "WAVESUSDT",
		},
	},
	"LST": {
		{
			ID:   5009,
			Slug: "luckyseventoken",
			Name: "LuckySevenToken",
		},
		{
			ID:   7623,
			Slug: "libartysharetoken",
			Name: "Libartysharetoken",
		},
	},
	"DBC": {
		{
			ID:          2316,
			Slug:        "deepbrain-chain",
			Name:        "DeepBrain Chain",
			BinanceUSDT: "DBCUSDT",
		},
	},
	"DIC": {
		{
			ID:          4385,
			Slug:        "daikicoin",
			Name:        "Daikicoin",
			BinanceUSDT: "DICUSDT",
		},
	},
	"GTH": {
		{
			ID:          7041,
			Slug:        "gather",
			Name:        "Gather",
			BinanceUSDT: "GTHUSDT",
		},
	},
	"TENT": {
		{
			ID:          2912,
			Slug:        "tent",
			Name:        "TENT",
			BinanceUSDT: "TENTUSDT",
		},
	},
	"MOONDAY": {
		{
			ID:          7481,
			Slug:        "moonday-finance",
			Name:        "Moonday Finance",
			BinanceUSDT: "MOONDAYUSDT",
		},
	},
	"ZTNZ": {
		{
			ID:          8165,
			Slug:        "ztranzit-coin",
			Name:        "Ztranzit Coin",
			BinanceUSDT: "ZTNZUSDT",
		},
	},
	"XPY": {
		{
			ID:          764,
			Slug:        "paycoin2",
			Name:        "PayCoin",
			BinanceUSDT: "XPYUSDT",
		},
	},
	"BAFE": {
		{
			ID:          9367,
			Slug:        "bafe-io",
			Name:        "Bafe io",
			BinanceUSDT: "BAFEUSDT",
		},
	},
	"VOX": {
		{
			ID:          7465,
			Slug:        "vox-finance",
			Name:        "Vox.Finance",
			BinanceUSDT: "VOXUSDT",
		},
	},
	"ZOM": {
		{
			ID:          7063,
			Slug:        "zoom-protocol",
			Name:        "Zoom Protocol",
			BinanceUSDT: "ZOMUSDT",
		},
	},
	"SFP": {
		{
			ID:          8119,
			Slug:        "safepal",
			Name:        "SafePal",
			BinanceUSDT: "SFPUSDT",
		},
	},
	"BCHBULL": {
		{
			ID:          5467,
			Slug:        "3x-long-bitcoin-cash-token",
			Name:        "3x Long Bitcoin Cash Token",
			BinanceUSDT: "BCHBULLUSDT",
		},
	},
	"MINIDOGE": {
		{
			ID:          10798,
			Slug:        "minidoge",
			Name:        "MiniDOGE",
			BinanceUSDT: "MINIDOGEUSDT",
		},
	},
	"MEDIA": {
		{
			ID:          9524,
			Slug:        "media-network",
			Name:        "Media Network",
			BinanceUSDT: "MEDIAUSDT",
		},
	},
	"ZASH": {
		{
			ID:          5610,
			Slug:        "zimbocash",
			Name:        "ZIMBOCASH",
			BinanceUSDT: "ZASHUSDT",
		},
	},
	"TRP": {
		{
			ID:          3939,
			Slug:        "tronipay",
			Name:        "Tronipay",
			BinanceUSDT: "TRPUSDT",
		},
	},
	"BEZOGE": {
		{
			ID:          9996,
			Slug:        "bezoge-earth",
			Name:        "Bezoge Earth",
			BinanceUSDT: "BEZOGEUSDT",
		},
	},
	"YFIA": {
		{
			ID:          7442,
			Slug:        "yfia",
			Name:        "YFIA",
			BinanceUSDT: "YFIAUSDT",
		},
	},
	"BLAZR": {
		{
			ID:          1623,
			Slug:        "blazercoin",
			Name:        "BlazerCoin",
			BinanceUSDT: "BLAZRUSDT",
		},
	},
	"OPEN": {
		{
			ID:   2762,
			Slug: "open-platform",
			Name: "Open Platform",
		},
		{
			ID:   7783,
			Slug: "open-governance-token",
			Name: "Open Governance Token",
		},
	},
	"DX": {
		{
			ID:          3139,
			Slug:        "dxchain-token",
			Name:        "DxChain Token",
			BinanceUSDT: "DXUSDT",
		},
	},
	"CNS": {
		{
			ID:          5963,
			Slug:        "centric-swap",
			Name:        "Centric Swap",
			BinanceUSDT: "CNSUSDT",
		},
	},
	"LNOT": {
		{
			ID:          7118,
			Slug:        "livenodes-token",
			Name:        "Livenodes Token",
			BinanceUSDT: "LNOTUSDT",
		},
	},
	"LCS": {
		{
			ID:          2970,
			Slug:        "local-coin-swap",
			Name:        "LocalCoinSwap",
			BinanceUSDT: "LCSUSDT",
		},
	},
	"PLNC": {
		{
			ID:          263,
			Slug:        "plncoin",
			Name:        "PLNcoin",
			BinanceUSDT: "PLNCUSDT",
		},
	},
	"SBNB": {
		{
			ID:          6204,
			Slug:        "sbnb",
			Name:        "sBNB",
			BinanceUSDT: "SBNBUSDT",
		},
	},
	"CELL": {
		{
			ID:          8992,
			Slug:        "cellframe",
			Name:        "Cellframe",
			BinanceUSDT: "CELLUSDT",
		},
	},
	"CLU": {
		{
			ID:          9984,
			Slug:        "clucoin",
			Name:        "CluCoin",
			BinanceUSDT: "CLUUSDT",
		},
	},
	"PAY": {
		{
			ID:          1758,
			Slug:        "tenx",
			Name:        "TenX",
			BinanceUSDT: "PAYUSDT",
		},
	},
	"TERA": {
		{
			ID:          3948,
			Slug:        "tera",
			Name:        "TERA",
			BinanceUSDT: "TERAUSDT",
		},
	},
	"POLICEDOGE": {
		{
			ID:          10913,
			Slug:        "policedoge",
			Name:        "PoliceDOGE",
			BinanceUSDT: "POLICEDOGEUSDT",
		},
	},
	"MNP": {
		{
			ID:          3348,
			Slug:        "mnpcoin",
			Name:        "MNPCoin",
			BinanceUSDT: "MNPUSDT",
		},
	},
	"POC": {
		{
			ID:          8319,
			Slug:        "poc-blockchain",
			Name:        "POC Blockchain",
			BinanceUSDT: "POCUSDT",
		},
	},
	"WFX": {
		{
			ID:          3990,
			Slug:        "webflix-token",
			Name:        "Webflix Token",
			BinanceUSDT: "WFXUSDT",
		},
	},
	"AUCTION": {
		{
			ID:          8602,
			Slug:        "bounce-token",
			Name:        "Bounce Token",
			BinanceUSDT: "AUCTIONUSDT",
		},
	},
	"WIKI": {
		{
			ID:          3090,
			Slug:        "wiki-token",
			Name:        "Wiki Token",
			BinanceUSDT: "WIKIUSDT",
		},
	},
	"BNX": {
		{
			ID:          9891,
			Slug:        "binaryx",
			Name:        "BinaryX",
			BinanceUSDT: "BNXUSDT",
		},
	},
	"COMOS": {
		{
			ID:          9122,
			Slug:        "comos-finance",
			Name:        "COMOS Finance",
			BinanceUSDT: "COMOSUSDT",
		},
	},
	"MES": {
		{
			ID:          4870,
			Slug:        "meschain",
			Name:        "MesChain",
			BinanceUSDT: "MESUSDT",
		},
	},
	"MUSD": {
		{
			ID:          5747,
			Slug:        "mstable-usd",
			Name:        "mStable USD",
			BinanceUSDT: "MUSDUSDT",
		},
	},
	"FRC": {
		{
			ID:          10,
			Slug:        "freicoin",
			Name:        "Freicoin",
			BinanceUSDT: "FRCUSDT",
		},
	},
	"CTRFI": {
		{
			ID:          10038,
			Slug:        "chester-moon",
			Name:        "Chester.Moon",
			BinanceUSDT: "CTRFIUSDT",
		},
	},
	"SAITO": {
		{
			ID:          9194,
			Slug:        "saito",
			Name:        "Saito",
			BinanceUSDT: "SAITOUSDT",
		},
	},
	"SXC": {
		{
			ID:   6013,
			Slug: "sxc",
			Name: "SXC Token",
		},
		{
			ID:   10388,
			Slug: "supremex",
			Name: "SupremeX",
		},
	},
	"BNANA": {
		{
			ID:          3742,
			Slug:        "chimpion",
			Name:        "Chimpion",
			BinanceUSDT: "BNANAUSDT",
		},
	},
	"PHIBA": {
		{
			ID:          9768,
			Slug:        "papa-shiba",
			Name:        "Papa Shiba",
			BinanceUSDT: "PHIBAUSDT",
		},
	},
	"DCNTR": {
		{
			ID:          6609,
			Slug:        "decentrahub-coin",
			Name:        "Decentrahub Coin",
			BinanceUSDT: "DCNTRUSDT",
		},
	},
	"BGTT": {
		{
			ID:          6943,
			Slug:        "baguette-token",
			Name:        "Baguette Token",
			BinanceUSDT: "BGTTUSDT",
		},
	},
	"ALLEY": {
		{
			ID:          9527,
			Slug:        "nft-alley",
			Name:        "NFT Alley",
			BinanceUSDT: "ALLEYUSDT",
		},
	},
	"BNBUP": {
		{
			ID:          7009,
			Slug:        "bnbup",
			Name:        "BNBUP",
			BinanceUSDT: "BNBUPUSDT",
		},
	},
	"COT": {
		{
			ID:          3779,
			Slug:        "cotrader",
			Name:        "CoTrader",
			BinanceUSDT: "COTUSDT",
		},
	},
	"ROOT": {
		{
			ID:          8077,
			Slug:        "rootkit-finance",
			Name:        "Rootkit Finance",
			BinanceUSDT: "ROOTUSDT",
		},
	},
	"BERRY": {
		{
			ID:   2628,
			Slug: "rentberry",
			Name: "Rentberry",
		},
		{
			ID:   6257,
			Slug: "berry",
			Name: "Berry",
		},
		{
			ID:   9055,
			Slug: "berryswap",
			Name: "BerrySwap",
		},
	},
	"BXC": {
		{
			ID:   4703,
			Slug: "bonuscloud",
			Name: "BonusCloud",
		},
		{
			ID:   5168,
			Slug: "bitcoin-classic",
			Name: "Bitcoin Classic",
		},
	},
	"STARK": {
		{
			ID:          7610,
			Slug:        "stark-chain",
			Name:        "STARK CHAIN",
			BinanceUSDT: "STARKUSDT",
		},
	},
	"DENT": {
		{
			ID:          1886,
			Slug:        "dent",
			Name:        "Dent",
			BinanceUSDT: "DENTUSDT",
		},
	},
	"LTCDOWN": {
		{
			ID:          7527,
			Slug:        "ltcdown",
			Name:        "LTCDOWN",
			BinanceUSDT: "LTCDOWNUSDT",
		},
	},
	"HTBULL": {
		{
			ID:          6084,
			Slug:        "3x-long-huobi-token-token",
			Name:        "3X Long Huobi Token Token",
			BinanceUSDT: "HTBULLUSDT",
		},
	},
	"CKB": {
		{
			ID:          4948,
			Slug:        "nervos-network",
			Name:        "Nervos Network",
			BinanceUSDT: "CKBUSDT",
		},
	},
	"CWAP": {
		{
			ID:          9817,
			Slug:        "defire",
			Name:        "DeFIRE",
			BinanceUSDT: "CWAPUSDT",
		},
	},
	"SHO": {
		{
			ID:          6540,
			Slug:        "showcase",
			Name:        "Showcase",
			BinanceUSDT: "SHOUSDT",
		},
	},
	"CLV": {
		{
			ID:          8384,
			Slug:        "clover",
			Name:        "Clover Finance",
			BinanceUSDT: "CLVUSDT",
		},
	},
	"BITT": {
		{
			ID:          8517,
			Slug:        "bittoken",
			Name:        "BiTToken",
			BinanceUSDT: "BITTUSDT",
		},
	},
	"GST2": {
		{
			ID:          5716,
			Slug:        "gas-token-two",
			Name:        "Gas Token Two",
			BinanceUSDT: "GST2USDT",
		},
	},
	"FME": {
		{
			ID:          5955,
			Slug:        "fme",
			Name:        "FME",
			BinanceUSDT: "FMEUSDT",
		},
	},
	"BNB": {
		{
			ID:          1839,
			Slug:        "binance-coin",
			Name:        "Binance Coin",
			BinanceUSDT: "BNBUSDT",
		},
	},
	"MCAFEE": {
		{
			ID:          10619,
			Slug:        "the-last-mcafee-token",
			Name:        "The Last McAfee Token",
			BinanceUSDT: "MCAFEEUSDT",
		},
	},
	"PEARL": {
		{
			ID:          6829,
			Slug:        "pearl",
			Name:        "Pearl",
			BinanceUSDT: "PEARLUSDT",
		},
	},
	"BIZZ": {
		{
			ID:          5437,
			Slug:        "bizzcoin",
			Name:        "BIZZCOIN",
			BinanceUSDT: "BIZZUSDT",
		},
	},
	"DEXTF": {
		{
			ID:          8691,
			Slug:        "dextf-protocol",
			Name:        "DEXTF Protocol",
			BinanceUSDT: "DEXTFUSDT",
		},
	},
	"FTS": {
		{
			ID:          9592,
			Slug:        "fortress-lending",
			Name:        "Fortress Lending",
			BinanceUSDT: "FTSUSDT",
		},
	},
	"JINU": {
		{
			ID:          10445,
			Slug:        "jomon-inu",
			Name:        "Jomon Inu",
			BinanceUSDT: "JINUUSDT",
		},
	},
	"BASE": {
		{
			ID:          7838,
			Slug:        "base-protocol",
			Name:        "Base Protocol",
			BinanceUSDT: "BASEUSDT",
		},
	},
	"MOB": {
		{
			ID:          7878,
			Slug:        "mobilecoin",
			Name:        "MobileCoin",
			BinanceUSDT: "MOBUSDT",
		},
	},
	"ATFI": {
		{
			ID:          9844,
			Slug:        "atlantic-finance-token",
			Name:        "Atlantic Finance Token",
			BinanceUSDT: "ATFIUSDT",
		},
	},
	"JUR": {
		{
			ID:          6482,
			Slug:        "jur",
			Name:        "Jur",
			BinanceUSDT: "JURUSDT",
		},
	},
	"ANKR": {
		{
			ID:          3783,
			Slug:        "ankr",
			Name:        "Ankr",
			BinanceUSDT: "ANKRUSDT",
		},
	},
	"WAULTX": {
		{
			ID:          9478,
			Slug:        "wault-finance-new",
			Name:        "Wault [New]",
			BinanceUSDT: "WAULTXUSDT",
		},
	},
	"DISC": {
		{
			ID:          9582,
			Slug:        "discas-vision",
			Name:        "DisCas Vision",
			BinanceUSDT: "DISCUSDT",
		},
	},
	"MOGX": {
		{
			ID:          4800,
			Slug:        "mogu",
			Name:        "Mogu",
			BinanceUSDT: "MOGXUSDT",
		},
	},
	"IETH": {
		{
			ID:   2104,
			Slug: "iethereum",
			Name: "iEthereum",
		},
		{
			ID:   6188,
			Slug: "ieth",
			Name: "iETH",
		},
	},
	"IONX": {
		{
			ID:          10264,
			Slug:        "charged-particles",
			Name:        "Charged Particles",
			BinanceUSDT: "IONXUSDT",
		},
	},
	"TON": {
		{
			ID:   6731,
			Slug: "tokamak-network",
			Name: "Tokamak Network",
		},
		{
			ID:   6890,
			Slug: "tontoken",
			Name: "TON Token",
		},
		{
			ID:   7505,
			Slug: "ton-crystal",
			Name: "TON Crystal",
		},
	},
	"RET": {
		{
			ID:          3280,
			Slug:        "realtract",
			Name:        "RealTract",
			BinanceUSDT: "RETUSDT",
		},
	},
	"SCT": {
		{
			ID:          8429,
			Slug:        "clash-token",
			Name:        "Clash Token",
			BinanceUSDT: "SCTUSDT",
		},
	},
	"NEWTON": {
		{
			ID:          6685,
			Slug:        "newtonium",
			Name:        "Newtonium",
			BinanceUSDT: "NEWTONUSDT",
		},
	},
	"BEZ": {
		{
			ID:          2551,
			Slug:        "bezop",
			Name:        "Bezop",
			BinanceUSDT: "BEZUSDT",
		},
	},
	"PEG": {
		{
			ID:   4979,
			Slug: "pegnet",
			Name: "PegNet",
		},
		{
			ID:   9882,
			Slug: "pegazus-finance",
			Name: "Pegazus finance",
		},
	},
	"PRT": {
		{
			ID:          7348,
			Slug:        "portion",
			Name:        "Portion",
			BinanceUSDT: "PRTUSDT",
		},
	},
	"EYES": {
		{
			ID:          5673,
			Slug:        "eyes-protocol",
			Name:        "EYES Protocol",
			BinanceUSDT: "EYESUSDT",
		},
	},
	"TUSC": {
		{
			ID:          5131,
			Slug:        "the-universal-settlement-coin",
			Name:        "The Universal Settlement Coin",
			BinanceUSDT: "TUSCUSDT",
		},
	},
	"ARION": {
		{
			ID:          3165,
			Slug:        "arion",
			Name:        "Arion",
			BinanceUSDT: "ARIONUSDT",
		},
	},
	"ANC": {
		{
			ID:   8857,
			Slug: "anchor-protocol",
			Name: "Anchor Protocol",
		},
		{
			ID:   43,
			Slug: "anoncoin",
			Name: "Anoncoin",
		},
	},
	"PSY": {
		{
			ID:          8347,
			Slug:        "psychic",
			Name:        "Psychic",
			BinanceUSDT: "PSYUSDT",
		},
	},
	"CVCC": {
		{
			ID:          4152,
			Slug:        "cryptoverificationcoin",
			Name:        "CryptoVerificationCoin",
			BinanceUSDT: "CVCCUSDT",
		},
	},
	"KAIINU": {
		{
			ID:          10330,
			Slug:        "kai-inu",
			Name:        "KAI INU",
			BinanceUSDT: "KAIINUUSDT",
		},
	},
	"SRM": {
		{
			ID:          6187,
			Slug:        "serum",
			Name:        "Serum",
			BinanceUSDT: "SRMUSDT",
		},
	},
	"AV": {
		{
			ID:          1146,
			Slug:        "avatarcoin",
			Name:        "AvatarCoin",
			BinanceUSDT: "AVUSDT",
		},
	},
	"PIZZA": {
		{
			ID:   5831,
			Slug: "pizza",
			Name: "Pizza",
		},
		{
			ID:   8459,
			Slug: "pizzaswap",
			Name: "PizzaSwap",
		},
		{
			ID:   10282,
			Slug: "safepizza",
			Name: "SafePizza",
		},
	},
	"ROOM": {
		{
			ID:          8351,
			Slug:        "optionroom",
			Name:        "OptionRoom",
			BinanceUSDT: "ROOMUSDT",
		},
	},
	"SMGM": {
		{
			ID:          10867,
			Slug:        "smegmars",
			Name:        "SMEGMARS",
			BinanceUSDT: "SMGMUSDT",
		},
	},
	"EWT": {
		{
			ID:          5268,
			Slug:        "energy-web-token",
			Name:        "Energy Web Token",
			BinanceUSDT: "EWTUSDT",
		},
	},
	"AID": {
		{
			ID:          2462,
			Slug:        "aidcoin",
			Name:        "AidCoin",
			BinanceUSDT: "AIDUSDT",
		},
	},
	"XSUSHI": {
		{
			ID:          8899,
			Slug:        "xsushi",
			Name:        "xSUSHI",
			BinanceUSDT: "XSUSHIUSDT",
		},
	},
	"LOF": {
		{
			ID:          10012,
			Slug:        "lonelyfans",
			Name:        "Lonelyfans",
			BinanceUSDT: "LOFUSDT",
		},
	},
	"BSCV": {
		{
			ID:          8843,
			Slug:        "bscview",
			Name:        "BSCView",
			BinanceUSDT: "BSCVUSDT",
		},
	},
	"KKC": {
		{
			ID:          3421,
			Slug:        "kabberry-coin",
			Name:        "Kabberry Coin",
			BinanceUSDT: "KKCUSDT",
		},
	},
	"THOREUM": {
		{
			ID:          10787,
			Slug:        "thoreum",
			Name:        "Thoreum",
			BinanceUSDT: "THOREUMUSDT",
		},
	},
	"NMC": {
		{
			ID:          3,
			Slug:        "namecoin",
			Name:        "Namecoin",
			BinanceUSDT: "NMCUSDT",
		},
	},
	"MTHD": {
		{
			ID:          9108,
			Slug:        "method-finance",
			Name:        "Method Finance",
			BinanceUSDT: "MTHDUSDT",
		},
	},
	"KASH": {
		{
			ID:          5995,
			Slug:        "kids-cash",
			Name:        "Kids Cash",
			BinanceUSDT: "KASHUSDT",
		},
	},
	"KRB": {
		{
			ID:          1340,
			Slug:        "karbo",
			Name:        "Karbo",
			BinanceUSDT: "KRBUSDT",
		},
	},
	"AXE": {
		{
			ID:          3898,
			Slug:        "axe",
			Name:        "Axe",
			BinanceUSDT: "AXEUSDT",
		},
	},
	"ZERO": {
		{
			ID:          8293,
			Slug:        "zero-exchange",
			Name:        "Zero Exchange",
			BinanceUSDT: "ZEROUSDT",
		},
	},
	"KSP": {
		{
			ID:          8296,
			Slug:        "klayswap-protocol",
			Name:        "KLAYswap Protocol",
			BinanceUSDT: "KSPUSDT",
		},
	},
	"PSI": {
		{
			ID:          8243,
			Slug:        "passive-income",
			Name:        "Passive Income",
			BinanceUSDT: "PSIUSDT",
		},
	},
	"KDOGE": {
		{
			ID:          10261,
			Slug:        "kingdoge",
			Name:        "KINGDOGE",
			BinanceUSDT: "KDOGEUSDT",
		},
	},
	"BT": {
		{
			ID:          8558,
			Slug:        "bt-finance",
			Name:        "BT.Finance",
			BinanceUSDT: "BTUSDT",
		},
	},
	"SMRAT": {
		{
			ID:          9284,
			Slug:        "secured-moonrat-token",
			Name:        "Secured MoonRat Token",
			BinanceUSDT: "SMRATUSDT",
		},
	},
	"BTD": {
		{
			ID:          8204,
			Slug:        "bat-true-dollar",
			Name:        "Bolt Dollar",
			BinanceUSDT: "BTDUSDT",
		},
	},
	"ZCL": {
		{
			ID:          1447,
			Slug:        "zclassic",
			Name:        "ZClassic",
			BinanceUSDT: "ZCLUSDT",
		},
	},
	"TRB": {
		{
			ID:          4944,
			Slug:        "tellor",
			Name:        "Tellor",
			BinanceUSDT: "TRBUSDT",
		},
	},
	"GASP": {
		{
			ID:          7625,
			Slug:        "gasp",
			Name:        "gAsp",
			BinanceUSDT: "GASPUSDT",
		},
	},
	"ESPRO": {
		{
			ID:          8760,
			Slug:        "esportspro",
			Name:        "EsportsPro",
			BinanceUSDT: "ESPROUSDT",
		},
	},
	"SYNC": {
		{
			ID:          7812,
			Slug:        "sync-network",
			Name:        "SYNC Network",
			BinanceUSDT: "SYNCUSDT",
		},
	},
	"BOLT": {
		{
			ID:          3843,
			Slug:        "bolt",
			Name:        "BOLT",
			BinanceUSDT: "BOLTUSDT",
		},
	},
	"WATER": {
		{
			ID:          8860,
			Slug:        "water-finance-token",
			Name:        "WaterDefi",
			BinanceUSDT: "WATERUSDT",
		},
	},
	"COVER": {
		{
			ID:          8175,
			Slug:        "cover-protocol-new",
			Name:        "COVER Protocol",
			BinanceUSDT: "COVERUSDT",
		},
	},
	"GCN": {
		{
			ID:          730,
			Slug:        "gcn-coin",
			Name:        "GCN Coin",
			BinanceUSDT: "GCNUSDT",
		},
	},
	"UDOO": {
		{
			ID:          3608,
			Slug:        "hyprr",
			Name:        "Howdoo",
			BinanceUSDT: "UDOOUSDT",
		},
	},
	"XUEZ": {
		{
			ID:          3798,
			Slug:        "xuez",
			Name:        "Xuez",
			BinanceUSDT: "XUEZUSDT",
		},
	},
	"IRD": {
		{
			ID:          3457,
			Slug:        "iridium",
			Name:        "Iridium",
			BinanceUSDT: "IRDUSDT",
		},
	},
	"WGR": {
		{
			ID:          1779,
			Slug:        "wagerr",
			Name:        "Wagerr",
			BinanceUSDT: "WGRUSDT",
		},
	},
	"DOGEDAO": {
		{
			ID:          9804,
			Slug:        "dogedao-finance",
			Name:        "DogeDao Finance",
			BinanceUSDT: "DOGEDAOUSDT",
		},
	},
	"WXC": {
		{
			ID:          3363,
			Slug:        "wxcoins",
			Name:        "WXCOINS",
			BinanceUSDT: "WXCUSDT",
		},
	},
	"SDS": {
		{
			ID:          3119,
			Slug:        "alchemint-standards",
			Name:        "Alchemint Standards",
			BinanceUSDT: "SDSUSDT",
		},
	},
	"DAD": {
		{
			ID:          4862,
			Slug:        "dad",
			Name:        "DAD",
			BinanceUSDT: "DADUSDT",
		},
	},
	"MTI": {
		{
			ID:          7615,
			Slug:        "mti-finance",
			Name:        "MTI Finance",
			BinanceUSDT: "MTIUSDT",
		},
	},
	"STRAX": {
		{
			ID:          1343,
			Slug:        "stratis",
			Name:        "Stratis",
			BinanceUSDT: "STRAXUSDT",
		},
	},
	"SWACE": {
		{
			ID:          4199,
			Slug:        "swace",
			Name:        "Swace",
			BinanceUSDT: "SWACEUSDT",
		},
	},
	"SONG": {
		{
			ID:          857,
			Slug:        "songcoin",
			Name:        "SongCoin",
			BinanceUSDT: "SONGUSDT",
		},
	},
	"MSB": {
		{
			ID:          7591,
			Slug:        "misbloc",
			Name:        "Misbloc",
			BinanceUSDT: "MSBUSDT",
		},
	},
	"XHT": {
		{
			ID:          10423,
			Slug:        "hollaex-token",
			Name:        "HollaEx Token",
			BinanceUSDT: "XHTUSDT",
		},
	},
	"RNO": {
		{
			ID:          3679,
			Slug:        "earneo",
			Name:        "Earneo",
			BinanceUSDT: "RNOUSDT",
		},
	},
	"CRP": {
		{
			ID:          6865,
			Slug:        "utopia",
			Name:        "Crypton",
			BinanceUSDT: "CRPUSDT",
		},
	},
	"GDC": {
		{
			ID:          4678,
			Slug:        "global-digital-content",
			Name:        "Global Digital Content",
			BinanceUSDT: "GDCUSDT",
		},
	},
	"HEGIC": {
		{
			ID:          6929,
			Slug:        "hegic",
			Name:        "Hegic",
			BinanceUSDT: "HEGICUSDT",
		},
	},
	"BPS": {
		{
			ID:          5815,
			Slug:        "bitcoinpos",
			Name:        "BitcoinPoS",
			BinanceUSDT: "BPSUSDT",
		},
	},
	"PXC": {
		{
			ID:   35,
			Slug: "phoenixcoin",
			Name: "Phoenixcoin",
		},
		{
			ID:   2670,
			Slug: "pixie-coin",
			Name: "Pixie Coin",
		},
	},
	"BECN": {
		{
			ID:          3656,
			Slug:        "beacon",
			Name:        "Beacon",
			BinanceUSDT: "BECNUSDT",
		},
	},
	"CARMA": {
		{
			ID:          10637,
			Slug:        "carma-coin",
			Name:        "CARMA COIN",
			BinanceUSDT: "CARMAUSDT",
		},
	},
	"HZT": {
		{
			ID:          6214,
			Slug:        "black-diamond-rating",
			Name:        "Black Diamond Rating",
			BinanceUSDT: "HZTUSDT",
		},
	},
	"LQTY": {
		{
			ID:          7429,
			Slug:        "liquity",
			Name:        "Liquity",
			BinanceUSDT: "LQTYUSDT",
		},
	},
	"808TA": {
		{
			ID:          5108,
			Slug:        "808ta",
			Name:        "808TA",
			BinanceUSDT: "808TAUSDT",
		},
	},
	"NFTBOX": {
		{
			ID:          10063,
			Slug:        "nftbox-fun",
			Name:        "NFTBOX.fun",
			BinanceUSDT: "NFTBOXUSDT",
		},
	},
	"PART": {
		{
			ID:          1826,
			Slug:        "particl",
			Name:        "Particl",
			BinanceUSDT: "PARTUSDT",
		},
	},
	"NPXSXEM": {
		{
			ID:          3096,
			Slug:        "pundi-x-nem",
			Name:        "Pundi X NEM",
			BinanceUSDT: "NPXSXEMUSDT",
		},
	},
	"ABNB": {
		{
			ID:          7932,
			Slug:        "airbnb-tokenized-stock-ftx",
			Name:        "Airbnb tokenized stock FTX",
			BinanceUSDT: "ABNBUSDT",
		},
	},
	"GENIUS": {
		{
			ID:          10948,
			Slug:        "genius-coin",
			Name:        "Genius Coin",
			BinanceUSDT: "GENIUSUSDT",
		},
	},
	"YSR": {
		{
			ID:          5779,
			Slug:        "ystar",
			Name:        "Ystar",
			BinanceUSDT: "YSRUSDT",
		},
	},
	"GAMESAFE": {
		{
			ID:          9909,
			Slug:        "gamesafe-io",
			Name:        "Gamesafe.io",
			BinanceUSDT: "GAMESAFEUSDT",
		},
	},
	"COMB": {
		{
			ID:   7718,
			Slug: "combo",
			Name: "Combo",
		},
		{
			ID:   7178,
			Slug: "combine-finance",
			Name: "Combine.finance",
		},
	},
	"MNSTP": {
		{
			ID:          9213,
			Slug:        "moon-stop",
			Name:        "Moon Stop",
			BinanceUSDT: "MNSTPUSDT",
		},
	},
	"MRNA": {
		{
			ID:          7900,
			Slug:        "moderna-tokenized-stock-ftx",
			Name:        "Moderna tokenized stock FTX",
			BinanceUSDT: "MRNAUSDT",
		},
	},
	"IDALL": {
		{
			ID:          7514,
			Slug:        "idall",
			Name:        "IDall",
			BinanceUSDT: "IDALLUSDT",
		},
	},
	"YFIII": {
		{
			ID:   7143,
			Slug: "dify-finance",
			Name: "DiFy.Finance",
		},
		{
			ID:   7629,
			Slug: "yfiii",
			Name: "YFIII",
		},
	},
	"TRYB": {
		{
			ID:          5181,
			Slug:        "bilira",
			Name:        "BiLira",
			BinanceUSDT: "TRYBUSDT",
		},
	},
	"CZZ": {
		{
			ID:          10083,
			Slug:        "classzz",
			Name:        "ClassZZ",
			BinanceUSDT: "CZZUSDT",
		},
	},
	"CELC": {
		{
			ID:          5388,
			Slug:        "celcoin",
			Name:        "CelCoin",
			BinanceUSDT: "CELCUSDT",
		},
	},
	"BZNT": {
		{
			ID:          2727,
			Slug:        "bezant",
			Name:        "Bezant",
			BinanceUSDT: "BZNTUSDT",
		},
	},
	"EAURIC": {
		{
			ID:          7758,
			Slug:        "eauric",
			Name:        "Eauric",
			BinanceUSDT: "EAURICUSDT",
		},
	},
	"POLIS": {
		{
			ID:          2359,
			Slug:        "polis",
			Name:        "Polis",
			BinanceUSDT: "POLISUSDT",
		},
	},
	"RC20": {
		{
			ID:          3790,
			Slug:        "robocalls",
			Name:        "RoboCalls",
			BinanceUSDT: "RC20USDT",
		},
	},
	"VIDZ": {
		{
			ID:          1511,
			Slug:        "purevidz",
			Name:        "PureVidz",
			BinanceUSDT: "VIDZUSDT",
		},
	},
	"GTX": {
		{
			ID:          5366,
			Slug:        "goaltime-n",
			Name:        "GoalTime N",
			BinanceUSDT: "GTXUSDT",
		},
	},
	"ALGOBULL": {
		{
			ID:          6074,
			Slug:        "3x-long-algorand-token",
			Name:        "3X Long Algorand Token",
			BinanceUSDT: "ALGOBULLUSDT",
		},
	},
	"PLURA": {
		{
			ID:          3324,
			Slug:        "pluracoin",
			Name:        "PluraCoin",
			BinanceUSDT: "PLURAUSDT",
		},
	},
	"CRE": {
		{
			ID:   3946,
			Slug: "carry",
			Name: "Carry",
		},
		{
			ID:   2719,
			Slug: "cybereits",
			Name: "Cybereits",
		},
	},
	"UBQ": {
		{
			ID:          588,
			Slug:        "ubiq",
			Name:        "Ubiq",
			BinanceUSDT: "UBQUSDT",
		},
	},
	"EXRT": {
		{
			ID:          8807,
			Slug:        "exrt-network",
			Name:        "EXRT Network",
			BinanceUSDT: "EXRTUSDT",
		},
	},
	"LRG": {
		{
			ID:          5329,
			Slug:        "largo-coin",
			Name:        "Largo Coin",
			BinanceUSDT: "LRGUSDT",
		},
	},
	"FLY": {
		{
			ID:          9120,
			Slug:        "franklin",
			Name:        "Franklin",
			BinanceUSDT: "FLYUSDT",
		},
	},
	"YO": {
		{
			ID:          4229,
			Slug:        "yobit-token",
			Name:        "Yobit Token",
			BinanceUSDT: "YOUSDT",
		},
	},
	"CCF": {
		{
			ID:          10755,
			Slug:        "cashcow-finance",
			Name:        "Cashcow Finance",
			BinanceUSDT: "CCFUSDT",
		},
	},
	"SMOKE": {
		{
			ID:          8503,
			Slug:        "the-smokehouse",
			Name:        "The Smokehouse",
			BinanceUSDT: "SMOKEUSDT",
		},
	},
	"STZEN": {
		{
			ID:          8945,
			Slug:        "stakedzen",
			Name:        "StakedZEN",
			BinanceUSDT: "STZENUSDT",
		},
	},
	"PPY": {
		{
			ID:          1719,
			Slug:        "peerplays-ppy",
			Name:        "Peerplays",
			BinanceUSDT: "PPYUSDT",
		},
	},
	"ETNA": {
		{
			ID:          8962,
			Slug:        "etna-network",
			Name:        "ETNA Network",
			BinanceUSDT: "ETNAUSDT",
		},
	},
	"EXCL": {
		{
			ID:          633,
			Slug:        "exclusivecoin",
			Name:        "ExclusiveCoin",
			BinanceUSDT: "EXCLUSDT",
		},
	},
	"YTA": {
		{
			ID:          4133,
			Slug:        "yottachain",
			Name:        "YottaChain",
			BinanceUSDT: "YTAUSDT",
		},
	},
	"SKEY": {
		{
			ID:          8133,
			Slug:        "smartkey",
			Name:        "SmartKey",
			BinanceUSDT: "SKEYUSDT",
		},
	},
	"SPC": {
		{
			ID:          2410,
			Slug:        "spacechain",
			Name:        "SpaceChain",
			BinanceUSDT: "SPCUSDT",
		},
	},
	"PTE": {
		{
			ID:          10478,
			Slug:        "peet-defi-new",
			Name:        "Peet DeFi [new]",
			BinanceUSDT: "PTEUSDT",
		},
	},
	"POOL": {
		{
			ID:          8508,
			Slug:        "pooltogether",
			Name:        "PoolTogether",
			BinanceUSDT: "POOLUSDT",
		},
	},
	"BST1": {
		{
			ID:          7670,
			Slug:        "blueshare-token",
			Name:        "Blueshare Token",
			BinanceUSDT: "BST1USDT",
		},
	},
	"EDG": {
		{
			ID:   5274,
			Slug: "edgeware",
			Name: "Edgeware",
		},
		{
			ID:   1596,
			Slug: "edgeless",
			Name: "Edgeless",
		},
	},
	"BXX": {
		{
			ID:          10949,
			Slug:        "baanx",
			Name:        "Baanx",
			BinanceUSDT: "BXXUSDT",
		},
	},
	"CUDOS": {
		{
			ID:          8258,
			Slug:        "cudos",
			Name:        "CUDOS",
			BinanceUSDT: "CUDOSUSDT",
		},
	},
	"MTDR": {
		{
			ID:          10078,
			Slug:        "matador-token",
			Name:        "Matador Token",
			BinanceUSDT: "MTDRUSDT",
		},
	},
	"AMA": {
		{
			ID:   5469,
			Slug: "amaten",
			Name: "AMATEN",
		},
		{
			ID:   9935,
			Slug: "mrweb-finance",
			Name: "Mrweb Finance",
		},
	},
	"XFC": {
		{
			ID:          3663,
			Slug:        "footballcoin",
			Name:        "Footballcoin",
			BinanceUSDT: "XFCUSDT",
		},
	},
	"KAVA": {
		{
			ID:          4846,
			Slug:        "kava",
			Name:        "Kava.io",
			BinanceUSDT: "KAVAUSDT",
		},
	},
	"XBTC21": {
		{
			ID:          1248,
			Slug:        "bitcoin-21",
			Name:        "Bitcoin 21",
			BinanceUSDT: "XBTC21USDT",
		},
	},
	"JET": {
		{
			ID:          1787,
			Slug:        "jetcoin",
			Name:        "Jetcoin",
			BinanceUSDT: "JETUSDT",
		},
	},
	"SFEX": {
		{
			ID:          10434,
			Slug:        "safelaunch",
			Name:        "SafeX",
			BinanceUSDT: "SFEXUSDT",
		},
	},
	"LABS": {
		{
			ID:          8813,
			Slug:        "labs-group",
			Name:        "LABS Group",
			BinanceUSDT: "LABSUSDT",
		},
	},
	"BHD": {
		{
			ID:          3966,
			Slug:        "bitcoinhd",
			Name:        "BitcoinHD",
			BinanceUSDT: "BHDUSDT",
		},
	},
	"YLDY": {
		{
			ID:          10820,
			Slug:        "yieldly",
			Name:        "Yieldly",
			BinanceUSDT: "YLDYUSDT",
		},
	},
	"UNIFI": {
		{
			ID:          7122,
			Slug:        "unifi-defi",
			Name:        "UNIFI DeFi",
			BinanceUSDT: "UNIFIUSDT",
		},
	},
	"BBOO": {
		{
			ID:          8350,
			Slug:        "panda-yield",
			Name:        "Panda Yield",
			BinanceUSDT: "BBOOUSDT",
		},
	},
	"SATOZ": {
		{
			ID:          9241,
			Slug:        "satozhi",
			Name:        "Satozhi",
			BinanceUSDT: "SATOZUSDT",
		},
	},
	"CONX": {
		{
			ID:          1632,
			Slug:        "concoin",
			Name:        "Concoin",
			BinanceUSDT: "CONXUSDT",
		},
	},
	"FSHIB": {
		{
			ID:          10752,
			Slug:        "floki-shiba",
			Name:        "Floki Shiba",
			BinanceUSDT: "FSHIBUSDT",
		},
	},
	"EC2": {
		{
			ID:          9360,
			Slug:        "employmentcoin",
			Name:        "EmploymentCoin",
			BinanceUSDT: "EC2USDT",
		},
	},
	"FXS": {
		{
			ID:          6953,
			Slug:        "frax-share",
			Name:        "Frax Share",
			BinanceUSDT: "FXSUSDT",
		},
	},
	"UCR": {
		{
			ID:          7199,
			Slug:        "ultra-clear",
			Name:        "Ultra Clear",
			BinanceUSDT: "UCRUSDT",
		},
	},
	"ELONPEG": {
		{
			ID:          10671,
			Slug:        "elonpeg",
			Name:        "ElonPeg",
			BinanceUSDT: "ELONPEGUSDT",
		},
	},
	"KAI": {
		{
			ID:          5453,
			Slug:        "kardiachain",
			Name:        "KardiaChain",
			BinanceUSDT: "KAIUSDT",
		},
	},
	"MWC": {
		{
			ID:          5031,
			Slug:        "mimblewimblecoin",
			Name:        "MimbleWimbleCoin",
			BinanceUSDT: "MWCUSDT",
		},
	},
	"RMPL": {
		{
			ID:          6503,
			Slug:        "rmpl",
			Name:        "RMPL",
			BinanceUSDT: "RMPLUSDT",
		},
	},
	"PAND": {
		{
			ID:          10014,
			Slug:        "panda-finance",
			Name:        "Panda Finance",
			BinanceUSDT: "PANDUSDT",
		},
	},
	"FLOKI": {
		{
			ID:   10804,
			Slug: "floki-inu",
			Name: "Floki Inu",
		},
		{
			ID:   10901,
			Slug: "shiba-floki",
			Name: "Shiba Floki",
		},
	},
	"FLETA": {
		{
			ID:          4103,
			Slug:        "fleta",
			Name:        "FLETA",
			BinanceUSDT: "FLETAUSDT",
		},
	},
	"INRT": {
		{
			ID:          5416,
			Slug:        "inrtoken",
			Name:        "INRToken",
			BinanceUSDT: "INRTUSDT",
		},
	},
	"ALPACA": {
		{
			ID:          8707,
			Slug:        "alpaca-finance",
			Name:        "Alpaca Finance",
			BinanceUSDT: "ALPACAUSDT",
		},
	},
	"PAID": {
		{
			ID:          8329,
			Slug:        "paid-network",
			Name:        "PAID Network",
			BinanceUSDT: "PAIDUSDT",
		},
	},
	"PSWAP": {
		{
			ID:   6754,
			Slug: "polkaswap",
			Name: "Polkaswap",
		},
		{
			ID:   10699,
			Slug: "porkswap",
			Name: "PorkSwap",
		},
	},
	"YFIUP": {
		{
			ID:          7452,
			Slug:        "yfiup",
			Name:        "YFIUP",
			BinanceUSDT: "YFIUPUSDT",
		},
	},
	"BLU": {
		{
			ID:          290,
			Slug:        "bluecoin",
			Name:        "BlueCoin",
			BinanceUSDT: "BLUUSDT",
		},
	},
	"ICH": {
		{
			ID:   5560,
			Slug: "idea-chain-coin",
			Name: "Idea Chain Coin",
		},
		{
			ID:   7162,
			Slug: "icherry-finance",
			Name: "iCherry Finance",
		},
	},
	"PHI": {
		{
			ID:          2676,
			Slug:        "phi-token",
			Name:        "PHI Token",
			BinanceUSDT: "PHIUSDT",
		},
	},
	"ECOC": {
		{
			ID:          5477,
			Slug:        "ecochain",
			Name:        "ECOChain",
			BinanceUSDT: "ECOCUSDT",
		},
	},
	"NTR": {
		{
			ID:          3950,
			Slug:        "netrum",
			Name:        "Netrum",
			BinanceUSDT: "NTRUSDT",
		},
	},
	"IFUND": {
		{
			ID:          8767,
			Slug:        "unifund",
			Name:        "Unifund",
			BinanceUSDT: "IFUNDUSDT",
		},
	},
	"BTCDOWN": {
		{
			ID:          5609,
			Slug:        "btcdown",
			Name:        "BTCDOWN",
			BinanceUSDT: "BTCDOWNUSDT",
		},
	},
	"KAU": {
		{
			ID:          7346,
			Slug:        "kauri-crypto",
			Name:        "Kauri",
			BinanceUSDT: "KAUUSDT",
		},
	},
	"RPZX": {
		{
			ID:          4298,
			Slug:        "rapidz",
			Name:        "Rapidz",
			BinanceUSDT: "RPZXUSDT",
		},
	},
	"ZCN": {
		{
			ID:          2882,
			Slug:        "0chain",
			Name:        "0Chain",
			BinanceUSDT: "ZCNUSDT",
		},
	},
	"TFC": {
		{
			ID:   9872,
			Slug: "thefutbolcoin",
			Name: "TheFutbolCoin",
		},
		{
			ID:   9625,
			Slug: "triforce-protocol",
			Name: "Triforce Protocol",
		},
	},
	"EGG": {
		{
			ID:   4467,
			Slug: "nestree",
			Name: "Nestree",
		},
		{
			ID:   8449,
			Slug: "goose-finance",
			Name: "Goose Finance",
		},
		{
			ID:   7665,
			Slug: "nestegg-coin",
			Name: "NestEGG Coin",
		},
	},
	"BALPHA": {
		{
			ID:          8710,
			Slug:        "balpha",
			Name:        "bAlpha",
			BinanceUSDT: "BALPHAUSDT",
		},
	},
	"VGW": {
		{
			ID:          3735,
			Slug:        "vegawallet-token",
			Name:        "VegaWallet Token",
			BinanceUSDT: "VGWUSDT",
		},
	},
	"MKB": {
		{
			ID:          9727,
			Slug:        "maker-basic-mkb",
			Name:        "Maker Basic-MKB",
			BinanceUSDT: "MKBUSDT",
		},
	},
	"MBS": {
		{
			ID:          9405,
			Slug:        "moonboys",
			Name:        "MoonBoys",
			BinanceUSDT: "MBSUSDT",
		},
	},
	"VIBRA": {
		{
			ID:          10143,
			Slug:        "vibraniums",
			Name:        "Vibraniums",
			BinanceUSDT: "VIBRAUSDT",
		},
	},
	"HORD": {
		{
			ID:          9198,
			Slug:        "hord",
			Name:        "Hord",
			BinanceUSDT: "HORDUSDT",
		},
	},
	"XRA": {
		{
			ID:   978,
			Slug: "ratecoin",
			Name: "Ratecoin",
		},
		{
			ID:   3234,
			Slug: "xriba",
			Name: "Xriba",
		},
	},
	"1MT": {
		{
			ID:          4222,
			Slug:        "1million-token",
			Name:        "1Million Token",
			BinanceUSDT: "1MTUSDT",
		},
	},
	"FINE": {
		{
			ID:          9269,
			Slug:        "refinable",
			Name:        "Refinable",
			BinanceUSDT: "FINEUSDT",
		},
	},
	"AUC": {
		{
			ID:          2653,
			Slug:        "auctus",
			Name:        "Auctus",
			BinanceUSDT: "AUCUSDT",
		},
	},
	"PHNX": {
		{
			ID:          5674,
			Slug:        "phoenixdao",
			Name:        "PhoenixDAO",
			BinanceUSDT: "PHNXUSDT",
		},
	},
	"DREP": {
		{
			ID:          9148,
			Slug:        "drep-new",
			Name:        "Drep [new]",
			BinanceUSDT: "DREPUSDT",
		},
	},
	"TRUMP": {
		{
			ID:          1185,
			Slug:        "trumpcoin",
			Name:        "TrumpCoin",
			BinanceUSDT: "TRUMPUSDT",
		},
	},
	"TUNE": {
		{
			ID:          6330,
			Slug:        "tune-token",
			Name:        "TUNE TOKEN",
			BinanceUSDT: "TUNEUSDT",
		},
	},
	"OPAL": {
		{
			ID:          597,
			Slug:        "opal",
			Name:        "Opal",
			BinanceUSDT: "OPALUSDT",
		},
	},
	"FAR": {
		{
			ID:   7691,
			Slug: "farmland-protocol",
			Name: "Farmland Protocol",
		},
		{
			ID:   8932,
			Slug: "farswap",
			Name: "FarSwap",
		},
	},
	"QTV": {
		{
			ID:          5695,
			Slug:        "quish-coin",
			Name:        "Quish Coin",
			BinanceUSDT: "QTVUSDT",
		},
	},
	"CCXX": {
		{
			ID:          5482,
			Slug:        "counos-x",
			Name:        "Counos X",
			BinanceUSDT: "CCXXUSDT",
		},
	},
	"GOC": {
		{
			ID:          3052,
			Slug:        "gocrypto-token",
			Name:        "GoCrypto Token",
			BinanceUSDT: "GOCUSDT",
		},
	},
	"JMC": {
		{
			ID:          4324,
			Slug:        "junsonmingchncoin",
			Name:        "Junsonmingchncoin",
			BinanceUSDT: "JMCUSDT",
		},
	},
	"AQT": {
		{
			ID:          7460,
			Slug:        "alpha-quark-token",
			Name:        "Alpha Quark Token",
			BinanceUSDT: "AQTUSDT",
		},
	},
	"PPBLZ": {
		{
			ID:          7435,
			Slug:        "pepemon-pepeballs",
			Name:        "Pepemon Pepeballs",
			BinanceUSDT: "PPBLZUSDT",
		},
	},
	"CATS": {
		{
			ID:          10663,
			Slug:        "catoshi-nakamoto",
			Name:        "Catoshi Nakamoto",
			BinanceUSDT: "CATSUSDT",
		},
	},
	"PCH": {
		{
			ID:          2902,
			Slug:        "popchain",
			Name:        "POPCHAIN",
			BinanceUSDT: "PCHUSDT",
		},
	},
	"TAT": {
		{
			ID:          7537,
			Slug:        "tatcoin",
			Name:        "Tatcoin",
			BinanceUSDT: "TATUSDT",
		},
	},
	"POLR": {
		{
			ID:          10561,
			Slug:        "polystarter-net",
			Name:        "Polystarter.net",
			BinanceUSDT: "POLRUSDT",
		},
	},
	"PPAY": {
		{
			ID:          7870,
			Slug:        "plasma-finance",
			Name:        "Plasma Finance",
			BinanceUSDT: "PPAYUSDT",
		},
	},
	"EDI": {
		{
			ID:          5165,
			Slug:        "freight-trust-clearing-network",
			Name:        "Freight Trust & Clearing Network",
			BinanceUSDT: "EDIUSDT",
		},
	},
	"GRS": {
		{
			ID:          258,
			Slug:        "groestlcoin",
			Name:        "Groestlcoin",
			BinanceUSDT: "GRSUSDT",
		},
	},
	"INTRATIO": {
		{
			ID:          6155,
			Slug:        "intelligent-ratio-set",
			Name:        "Intelligent Ratio Set",
			BinanceUSDT: "INTRATIOUSDT",
		},
	},
	"OMN": {
		{
			ID:          10944,
			Slug:        "omni-token",
			Name:        "OMNI - People Driven",
			BinanceUSDT: "OMNUSDT",
		},
	},
	"DAB": {
		{
			ID:          4284,
			Slug:        "dabanking",
			Name:        "DABANKING",
			BinanceUSDT: "DABUSDT",
		},
	},
	"YFOX": {
		{
			ID:          7286,
			Slug:        "yfox-finance",
			Name:        "YFOX FINANCE",
			BinanceUSDT: "YFOXUSDT",
		},
	},
	"VIN": {
		{
			ID:          3082,
			Slug:        "vinchain",
			Name:        "VINchain",
			BinanceUSDT: "VINUSDT",
		},
	},
	"JSB": {
		{
			ID:          7787,
			Slug:        "jsb-foundation",
			Name:        "JSB FOUNDATION",
			BinanceUSDT: "JSBUSDT",
		},
	},
	"YFN": {
		{
			ID:          7478,
			Slug:        "yearn-finance-network",
			Name:        "Yearn Finance Network",
			BinanceUSDT: "YFNUSDT",
		},
	},
	"DOGEFI": {
		{
			ID:          6623,
			Slug:        "dogefi",
			Name:        "DOGEFI",
			BinanceUSDT: "DOGEFIUSDT",
		},
	},
	"AIM": {
		{
			ID:          5918,
			Slug:        "modihost",
			Name:        "ModiHost",
			BinanceUSDT: "AIMUSDT",
		},
	},
	"FootballStars": {
		{
			ID:          9910,
			Slug:        "football-stars",
			Name:        "Football Stars",
			BinanceUSDT: "FootballStarsUSDT",
		},
	},
	"EVA": {
		{
			ID:          10686,
			Slug:        "evanesco-network",
			Name:        "Evanesco Network",
			BinanceUSDT: "EVAUSDT",
		},
	},
	"ASS": {
		{
			ID:          9456,
			Slug:        "australian-safe-shepherd",
			Name:        "Australian Safe Shepherd",
			BinanceUSDT: "ASSUSDT",
		},
	},
	"RICK": {
		{
			ID:          6475,
			Slug:        "infinite-ricks",
			Name:        "Infinite Ricks",
			BinanceUSDT: "RICKUSDT",
		},
	},
	"PNX": {
		{
			ID:          2205,
			Slug:        "phantomx",
			Name:        "Phantomx",
			BinanceUSDT: "PNXUSDT",
		},
	},
	"AKRO": {
		{
			ID:          4134,
			Slug:        "akropolis",
			Name:        "Akropolis",
			BinanceUSDT: "AKROUSDT",
		},
	},
	"CATE": {
		{
			ID:   8950,
			Slug: "cash-tech",
			Name: "Cash Tech",
		},
		{
			ID:   9656,
			Slug: "catecoin",
			Name: "CateCoin",
		},
	},
	"DCH": {
		{
			ID:          10020,
			Slug:        "dechart",
			Name:        "DeChart",
			BinanceUSDT: "DCHUSDT",
		},
	},
	"vBTC": {
		{
			ID:          7962,
			Slug:        "venus-btc",
			Name:        "Venus BTC",
			BinanceUSDT: "vBTCUSDT",
		},
	},
	"GSWAP": {
		{
			ID:          7588,
			Slug:        "gameswap",
			Name:        "Gameswap",
			BinanceUSDT: "GSWAPUSDT",
		},
	},
	"XIF": {
		{
			ID:          7724,
			Slug:        "x-infinity",
			Name:        "X Infinity",
			BinanceUSDT: "XIFUSDT",
		},
	},
	"GOMA": {
		{
			ID:          10858,
			Slug:        "goma-finance",
			Name:        "GOMA Finance",
			BinanceUSDT: "GOMAUSDT",
		},
	},
	"FENIX": {
		{
			ID:          10124,
			Slug:        "fenix-finance",
			Name:        "Fenix Finance",
			BinanceUSDT: "FENIXUSDT",
		},
	},
	"ITEN": {
		{
			ID:          7154,
			Slug:        "iten",
			Name:        "ITEN",
			BinanceUSDT: "ITENUSDT",
		},
	},
	"HEX": {
		{
			ID:          5015,
			Slug:        "hex",
			Name:        "HEX",
			BinanceUSDT: "HEXUSDT",
		},
	},
	"NORD": {
		{
			ID:          8143,
			Slug:        "nord-finance",
			Name:        "Nord Finance",
			BinanceUSDT: "NORDUSDT",
		},
	},
	"JIND": {
		{
			ID:          10127,
			Slug:        "jindo-inu",
			Name:        "JINDO INU",
			BinanceUSDT: "JINDUSDT",
		},
	},
	"PROPS": {
		{
			ID:          5880,
			Slug:        "props",
			Name:        "Props Token",
			BinanceUSDT: "PROPSUSDT",
		},
	},
	"TCAKE": {
		{
			ID:          9063,
			Slug:        "pancaketools",
			Name:        "Tcake",
			BinanceUSDT: "TCAKEUSDT",
		},
	},
	"ZMT": {
		{
			ID:          8146,
			Slug:        "zipmex",
			Name:        "Zipmex",
			BinanceUSDT: "ZMTUSDT",
		},
	},
	"EARNX": {
		{
			ID:          9389,
			Slug:        "earnx",
			Name:        "EarnX",
			BinanceUSDT: "EARNXUSDT",
		},
	},
	"ALP": {
		{
			ID:          4820,
			Slug:        "alp-coin",
			Name:        "ALP Coin",
			BinanceUSDT: "ALPUSDT",
		},
	},
	"WINATESLA": {
		{
			ID:          10934,
			Slug:        "win-a-tesla",
			Name:        "WIN A TESLA",
			BinanceUSDT: "WINATESLAUSDT",
		},
	},
	"HTN": {
		{
			ID:          6464,
			Slug:        "heartnumber",
			Name:        "Heart Number",
			BinanceUSDT: "HTNUSDT",
		},
	},
	"COIN": {
		{
			ID:   6920,
			Slug: "coin-artist",
			Name: "Coin Artist",
		},
		{
			ID:   8068,
			Slug: "coinbase-pre-ipo-tokenized-stock-ftx",
			Name: "Coinbase tokenized stock FTX",
		},
	},
	"1INCH": {
		{
			ID:          8104,
			Slug:        "1inch",
			Name:        "1inch",
			BinanceUSDT: "1INCHUSDT",
		},
	},
	"$BABYDOGEINU": {
		{
			ID:          10914,
			Slug:        "baby-doge-inu",
			Name:        "BABY DOGE INU",
			BinanceUSDT: "$BABYDOGEINUUSDT",
		},
	},
	"SKRT": {
		{
			ID:          10554,
			Slug:        "sekuritance",
			Name:        "Sekuritance",
			BinanceUSDT: "SKRTUSDT",
		},
	},
	"EOSBULL": {
		{
			ID:          5414,
			Slug:        "3x-long-eos-token",
			Name:        "3x Long EOS Token",
			BinanceUSDT: "EOSBULLUSDT",
		},
	},
	"WANUSDC": {
		{
			ID:          8655,
			Slug:        "wanusdc",
			Name:        "wanUSDC",
			BinanceUSDT: "WANUSDCUSDT",
		},
	},
	"REW": {
		{
			ID:          5816,
			Slug:        "rewardiqa",
			Name:        "Rewardiqa",
			BinanceUSDT: "REWUSDT",
		},
	},
	"EXT": {
		{
			ID:          3074,
			Slug:        "experience-token",
			Name:        "Experience Token",
			BinanceUSDT: "EXTUSDT",
		},
	},
	"ALD": {
		{
			ID:          6519,
			Slug:        "aludra-network",
			Name:        "Aludra Network",
			BinanceUSDT: "ALDUSDT",
		},
	},
	"IRON": {
		{
			ID:          10484,
			Slug:        "iron-finance",
			Name:        "Iron",
			BinanceUSDT: "IRONUSDT",
		},
	},
	"TRONPAD": {
		{
			ID:          10277,
			Slug:        "tronpad",
			Name:        "TRONPAD",
			BinanceUSDT: "TRONPADUSDT",
		},
	},
	"UNC": {
		{
			ID:          5948,
			Slug:        "unicrypt",
			Name:        "Unicrypt",
			BinanceUSDT: "UNCUSDT",
		},
	},
	"GAL": {
		{
			ID:          5228,
			Slug:        "galatasaray-fan-token",
			Name:        "Galatasaray Fan Token",
			BinanceUSDT: "GALUSDT",
		},
	},
	"WATCH": {
		{
			ID:          8621,
			Slug:        "yieldwatch",
			Name:        "yieldwatch",
			BinanceUSDT: "WATCHUSDT",
		},
	},
	"DZI": {
		{
			ID:          6662,
			Slug:        "definition",
			Name:        "DeFinition",
			BinanceUSDT: "DZIUSDT",
		},
	},
	"TIPS": {
		{
			ID:          87,
			Slug:        "fedoracoin",
			Name:        "FedoraCoin",
			BinanceUSDT: "TIPSUSDT",
		},
	},
	"CTC": {
		{
			ID:   5198,
			Slug: "creditcoin",
			Name: "Creditcoin",
		},
		{
			ID:   8046,
			Slug: "cybertronchain",
			Name: "Cybertronchain",
		},
		{
			ID:   6036,
			Slug: "culture-ticket-chain",
			Name: "Culture Ticket Chain",
		},
	},
	"LTFM": {
		{
			ID:          10707,
			Slug:        "little-fish-moon-token",
			Name:        "Little Fish Moon Token",
			BinanceUSDT: "LTFMUSDT",
		},
	},
	"EXP": {
		{
			ID:   1070,
			Slug: "expanse",
			Name: "Expanse",
		},
		{
			ID:   7099,
			Slug: "exchange-payment-coin",
			Name: "Exchange Payment Coin",
		},
	},
	"TERC": {
		{
			ID:          6508,
			Slug:        "troneuroperewardcoin",
			Name:        "TronEuropeRewardCoin",
			BinanceUSDT: "TERCUSDT",
		},
	},
	"DNT": {
		{
			ID:          1856,
			Slug:        "district0x",
			Name:        "district0x",
			BinanceUSDT: "DNTUSDT",
		},
	},
	"XTP": {
		{
			ID:          5070,
			Slug:        "tap",
			Name:        "Tap",
			BinanceUSDT: "XTPUSDT",
		},
	},
	"FDOTA": {
		{
			ID:          10846,
			Slug:        "fomodota",
			Name:        "FomoDota",
			BinanceUSDT: "FDOTAUSDT",
		},
	},
	"ZCH": {
		{
			ID:          5316,
			Slug:        "0cash",
			Name:        "0cash",
			BinanceUSDT: "ZCHUSDT",
		},
	},
	"SG": {
		{
			ID:          6245,
			Slug:        "socialgood",
			Name:        "SocialGood",
			BinanceUSDT: "SGUSDT",
		},
	},
	"QBZ": {
		{
			ID:          5267,
			Slug:        "queenbee",
			Name:        "QUEENBEE",
			BinanceUSDT: "QBZUSDT",
		},
	},
	"HEPA": {
		{
			ID:          10504,
			Slug:        "hepa-finance",
			Name:        "Hepa Finance",
			BinanceUSDT: "HEPAUSDT",
		},
	},
	"UNIUSD": {
		{
			ID:          5694,
			Slug:        "unidollar",
			Name:        "UniDollar",
			BinanceUSDT: "UNIUSDUSDT",
		},
	},
	"$KEI": {
		{
			ID:          10048,
			Slug:        "keisuke-inu",
			Name:        "Keisuke Inu",
			BinanceUSDT: "$KEIUSDT",
		},
	},
	"WMATIC": {
		{
			ID:          8925,
			Slug:        "wmatic",
			Name:        "Wrapped Matic",
			BinanceUSDT: "WMATICUSDT",
		},
	},
	"SCC": {
		{
			ID:   3772,
			Slug: "stem-cell-coin",
			Name: "STEM CELL COIN",
		},
		{
			ID:   3986,
			Slug: "stakecubecoin",
			Name: "StakeCubeCoin",
		},
		{
			ID:   3128,
			Slug: "siacashcoin",
			Name: "SiaCashCoin",
		},
		{
			ID:   2618,
			Slug: "stockchain",
			Name: "StockChain",
		},
		{
			ID:   6266,
			Slug: "scc-digforit",
			Name: "SCC DIGforIT",
		},
	},
	"RFR": {
		{
			ID:   2553,
			Slug: "refereum",
			Name: "Refereum",
		},
		{
			ID:   8041,
			Slug: "refract",
			Name: "Refract",
		},
	},
	"BKING": {
		{
			ID:          10550,
			Slug:        "king-arthur",
			Name:        "King Arthur",
			BinanceUSDT: "BKINGUSDT",
		},
	},
	"CROX": {
		{
			ID:          9548,
			Slug:        "croxswap",
			Name:        "CroxSwap",
			BinanceUSDT: "CROXUSDT",
		},
	},
	"PSG": {
		{
			ID:          5226,
			Slug:        "paris-saint-germain-fan-token",
			Name:        "Paris Saint-Germain Fan Token",
			BinanceUSDT: "PSGUSDT",
		},
	},
	"FUZZ": {
		{
			ID:          1241,
			Slug:        "fuzzballs",
			Name:        "FuzzBalls",
			BinanceUSDT: "FUZZUSDT",
		},
	},
	"FIS": {
		{
			ID:          5882,
			Slug:        "stafi",
			Name:        "Stafi",
			BinanceUSDT: "FISUSDT",
		},
	},
	"JNTR": {
		{
			ID:          7545,
			Slug:        "jointer",
			Name:        "Jointer",
			BinanceUSDT: "JNTRUSDT",
		},
	},
	"AXO": {
		{
			ID:          10581,
			Slug:        "axolotl-finance",
			Name:        "Axolotl Finance",
			BinanceUSDT: "AXOUSDT",
		},
	},
	"CVT": {
		{
			ID:   2642,
			Slug: "cybervein",
			Name: "CyberVein",
		},
		{
			ID:   8369,
			Slug: "civitas-protocol",
			Name: "Civitas Protocol",
		},
	},
	"GDAO": {
		{
			ID:          7694,
			Slug:        "governor-dao",
			Name:        "Governor DAO",
			BinanceUSDT: "GDAOUSDT",
		},
	},
	"UCA": {
		{
			ID:          5479,
			Slug:        "uca-coin",
			Name:        "UCA Coin",
			BinanceUSDT: "UCAUSDT",
		},
	},
	"CORAL": {
		{
			ID:          7628,
			Slug:        "coral-swap",
			Name:        "Coral Swap",
			BinanceUSDT: "CORALUSDT",
		},
	},
	"RNB": {
		{
			ID:          9766,
			Slug:        "rentible",
			Name:        "Rentible",
			BinanceUSDT: "RNBUSDT",
		},
	},
	"ZEUS": {
		{
			ID:          3414,
			Slug:        "zeusnetwork",
			Name:        "ZeusNetwork",
			BinanceUSDT: "ZEUSUSDT",
		},
	},
	"SSP": {
		{
			ID:          2862,
			Slug:        "smartshare",
			Name:        "Smartshare",
			BinanceUSDT: "SSPUSDT",
		},
	},
	"LEAN": {
		{
			ID:          10091,
			Slug:        "lean-protocol",
			Name:        "Lean",
			BinanceUSDT: "LEANUSDT",
		},
	},
	"KALLY": {
		{
			ID:          9631,
			Slug:        "polkally",
			Name:        "Polkally",
			BinanceUSDT: "KALLYUSDT",
		},
	},
	"WEXPOLY": {
		{
			ID:          10725,
			Slug:        "waultswap-polygon",
			Name:        "WaultSwap Polygon",
			BinanceUSDT: "WEXPOLYUSDT",
		},
	},
	"PDO": {
		{
			ID:          9210,
			Slug:        "pollo",
			Name:        "Pollo Dollar",
			BinanceUSDT: "PDOUSDT",
		},
	},
	"NIRX": {
		{
			ID:          4971,
			Slug:        "nairax",
			Name:        "NairaX",
			BinanceUSDT: "NIRXUSDT",
		},
	},
	"ABAT": {
		{
			ID:          5754,
			Slug:        "aave-bat",
			Name:        "Aave BAT",
			BinanceUSDT: "ABATUSDT",
		},
	},
	"CPI": {
		{
			ID:          6654,
			Slug:        "crypto-price-index",
			Name:        "Crypto Price Index",
			BinanceUSDT: "CPIUSDT",
		},
	},
	"MLTP": {
		{
			ID:          10576,
			Slug:        "moonlift-protocol",
			Name:        "MoonLift Protocol",
			BinanceUSDT: "MLTPUSDT",
		},
	},
	"MBX": {
		{
			ID:          7437,
			Slug:        "mobiepay",
			Name:        "MobieCoin",
			BinanceUSDT: "MBXUSDT",
		},
	},
	"YUSRA": {
		{
			ID:          6726,
			Slug:        "yusra",
			Name:        "YUSRA",
			BinanceUSDT: "YUSRAUSDT",
		},
	},
	"DEFIT": {
		{
			ID:          9155,
			Slug:        "digital-fitness",
			Name:        "Digital Fitness",
			BinanceUSDT: "DEFITUSDT",
		},
	},
	"ERECT": {
		{
			ID:          10535,
			Slug:        "evererected",
			Name:        "EVERERECTED",
			BinanceUSDT: "ERECTUSDT",
		},
	},
	"DIREWOLF": {
		{
			ID:          10266,
			Slug:        "direwolf",
			Name:        "Direwolf",
			BinanceUSDT: "DIREWOLFUSDT",
		},
	},
	"PCHF": {
		{
			ID:          10706,
			Slug:        "peachfolio",
			Name:        "peachfolio",
			BinanceUSDT: "PCHFUSDT",
		},
	},
	"HELP": {
		{
			ID:   3459,
			Slug: "gohelpfund",
			Name: "GoHelpFund",
		},
		{
			ID:   4231,
			Slug: "helpico",
			Name: "Helpico",
		},
	},
	"NYC": {
		{
			ID:          298,
			Slug:        "newyorkcoin",
			Name:        "NewYorkCoin",
			BinanceUSDT: "NYCUSDT",
		},
	},
	"PZAP": {
		{
			ID:          9896,
			Slug:        "polyzap-finance",
			Name:        "PolyZap Finance",
			BinanceUSDT: "PZAPUSDT",
		},
	},
	"MDAO": {
		{
			ID:          9228,
			Slug:        "martian-dao",
			Name:        "Martian DAO",
			BinanceUSDT: "MDAOUSDT",
		},
	},
	"FXC": {
		{
			ID:          6773,
			Slug:        "futurexcrypto",
			Name:        "FUTUREXCRYPTO",
			BinanceUSDT: "FXCUSDT",
		},
	},
	"MED": {
		{
			ID:          2303,
			Slug:        "medibloc",
			Name:        "MediBloc",
			BinanceUSDT: "MEDUSDT",
		},
	},
	"SRX": {
		{
			ID:          10894,
			Slug:        "storx-network",
			Name:        "StorX Network",
			BinanceUSDT: "SRXUSDT",
		},
	},
	"PHO": {
		{
			ID:   175,
			Slug: "photon",
			Name: "Photon",
		},
		{
			ID:   8379,
			Slug: "phoswap",
			Name: "Phoswap",
		},
	},
	"OCTA": {
		{
			ID:          9273,
			Slug:        "octans",
			Name:        "Octans",
			BinanceUSDT: "OCTAUSDT",
		},
	},
	"NOTSAFEMOON": {
		{
			ID:          9750,
			Slug:        "notsafemoon",
			Name:        "NotSafeMoon",
			BinanceUSDT: "NOTSAFEMOONUSDT",
		},
	},
	"PDOG": {
		{
			ID:          10626,
			Slug:        "polkadog",
			Name:        "Polkadog",
			BinanceUSDT: "PDOGUSDT",
		},
	},
	"EROTICA": {
		{
			ID:          10197,
			Slug:        "erotica-token",
			Name:        "Erotica",
			BinanceUSDT: "EROTICAUSDT",
		},
	},
	"TRUE": {
		{
			ID:          2457,
			Slug:        "truechain",
			Name:        "TrueChain",
			BinanceUSDT: "TRUEUSDT",
		},
	},
	"KP2R": {
		{
			ID:          7751,
			Slug:        "kp2r-network",
			Name:        "KP2R.Network",
			BinanceUSDT: "KP2RUSDT",
		},
	},
	"METAMOON": {
		{
			ID:          9688,
			Slug:        "metamoon",
			Name:        "MetaMoon",
			BinanceUSDT: "METAMOONUSDT",
		},
	},
	"BABA": {
		{
			ID:   7889,
			Slug: "alibaba-tokenized-stock-ftx",
			Name: "Alibaba tokenized stock FTX",
		},
		{
			ID:   7921,
			Slug: "alibaba-tokenized-stock-bittrex",
			Name: "Alibaba tokenized stock Bittrex",
		},
	},
	"RFUEL": {
		{
			ID:          6537,
			Slug:        "rio-defi",
			Name:        "RioDeFi",
			BinanceUSDT: "RFUELUSDT",
		},
	},
	"PPP": {
		{
			ID:          2036,
			Slug:        "paypie",
			Name:        "PayPie",
			BinanceUSDT: "PPPUSDT",
		},
	},
	"KGO": {
		{
			ID:          8877,
			Slug:        "kiwigo",
			Name:        "KIWIGO",
			BinanceUSDT: "KGOUSDT",
		},
	},
	"MERCI": {
		{
			ID:          5646,
			Slug:        "merci",
			Name:        "MERCI",
			BinanceUSDT: "MERCIUSDT",
		},
	},
	"TARM": {
		{
			ID:          5564,
			Slug:        "armtoken",
			Name:        "ARMTOKEN",
			BinanceUSDT: "TARMUSDT",
		},
	},
	"MTR": {
		{
			ID:          6627,
			Slug:        "meter-stable",
			Name:        "Meter Stable",
			BinanceUSDT: "MTRUSDT",
		},
	},
	"ORC": {
		{
			ID:   5326,
			Slug: "orbit-chain",
			Name: "Orbit Chain",
		},
		{
			ID:   7103,
			Slug: "oracle-system",
			Name: "Oracle System",
		},
	},
	"REL": {
		{
			ID:   5807,
			Slug: "release-project",
			Name: "Release Project",
		},
		{
			ID:   6541,
			Slug: "relevant",
			Name: "Relevant",
		},
	},
	"$WNTR": {
		{
			ID:          9981,
			Slug:        "weentar",
			Name:        "Weentar",
			BinanceUSDT: "$WNTRUSDT",
		},
	},
	"JASMY": {
		{
			ID:          8425,
			Slug:        "jasmy",
			Name:        "Jasmy",
			BinanceUSDT: "JASMYUSDT",
		},
	},
	"YFIS": {
		{
			ID:          6807,
			Slug:        "yfiscurity",
			Name:        "YFISCURITY",
			BinanceUSDT: "YFISUSDT",
		},
	},
	"SALE": {
		{
			ID:          6742,
			Slug:        "dxsale-network",
			Name:        "DxSale Network",
			BinanceUSDT: "SALEUSDT",
		},
	},
	"MOCHI": {
		{
			ID:          8881,
			Slug:        "mochiswap",
			Name:        "MOCHISWAP",
			BinanceUSDT: "MOCHIUSDT",
		},
	},
	"GSE": {
		{
			ID:          3123,
			Slug:        "gsenetwork",
			Name:        "GSENetwork",
			BinanceUSDT: "GSEUSDT",
		},
	},
	"HARE": {
		{
			ID:          10675,
			Slug:        "hare-token",
			Name:        "Hare Token",
			BinanceUSDT: "HAREUSDT",
		},
	},
	"SXI": {
		{
			ID:          9624,
			Slug:        "safexi",
			Name:        "SafeXI",
			BinanceUSDT: "SXIUSDT",
		},
	},
	"MTLX": {
		{
			ID:          7256,
			Slug:        "mettalex",
			Name:        "Mettalex",
			BinanceUSDT: "MTLXUSDT",
		},
	},
	"MIMATIC": {
		{
			ID:          10238,
			Slug:        "mimatic",
			Name:        "miMatic",
			BinanceUSDT: "MIMATICUSDT",
		},
	},
	"WSPP": {
		{
			ID:          10841,
			Slug:        "wolf-safe-poor-people",
			Name:        "Wolf Safe Poor People",
			BinanceUSDT: "WSPPUSDT",
		},
	},
	"NAN": {
		{
			ID:          6529,
			Slug:        "nantrade",
			Name:        "NanTrade",
			BinanceUSDT: "NANUSDT",
		},
	},
	"UTNP": {
		{
			ID:          2524,
			Slug:        "universa",
			Name:        "Universa",
			BinanceUSDT: "UTNPUSDT",
		},
	},
	"INV": {
		{
			ID:          8720,
			Slug:        "inverse-finance",
			Name:        "Inverse Finance",
			BinanceUSDT: "INVUSDT",
		},
	},
	"SCA": {
		{
			ID:          9118,
			Slug:        "scaleswap",
			Name:        "Scaleswap",
			BinanceUSDT: "SCAUSDT",
		},
	},
	"LTCUP": {
		{
			ID:          7526,
			Slug:        "ltcup",
			Name:        "LTCUP",
			BinanceUSDT: "LTCUPUSDT",
		},
	},
	"SLP": {
		{
			ID:          5824,
			Slug:        "small-love-potion",
			Name:        "Small Love Potion",
			BinanceUSDT: "SLPUSDT",
		},
	},
	"PolyMoon": {
		{
			ID:          9531,
			Slug:        "polymoon",
			Name:        "PolyMoon",
			BinanceUSDT: "PolyMoonUSDT",
		},
	},
	"LINK": {
		{
			ID:          1975,
			Slug:        "chainlink",
			Name:        "Chainlink",
			BinanceUSDT: "LINKUSDT",
		},
	},
	"RBT": {
		{
			ID:   703,
			Slug: "rimbit",
			Name: "Rimbit",
		},
		{
			ID:   10822,
			Slug: "robust-protocol",
			Name: "Robust Protocol",
		},
		{
			ID:   9690,
			Slug: "rabbit-token",
			Name: "Rabbit token",
		},
	},
	"RULER": {
		{
			ID:          8698,
			Slug:        "ruler-protocol",
			Name:        "Ruler Protocol",
			BinanceUSDT: "RULERUSDT",
		},
	},
	"YUP": {
		{
			ID:          7689,
			Slug:        "yup-token",
			Name:        "Yup",
			BinanceUSDT: "YUPUSDT",
		},
	},
	"RUGZ": {
		{
			ID:          7649,
			Slug:        "rugz",
			Name:        "pulltherug.finance",
			BinanceUSDT: "RUGZUSDT",
		},
	},
	"GOF": {
		{
			ID:          7034,
			Slug:        "golff",
			Name:        "Golff",
			BinanceUSDT: "GOFUSDT",
		},
	},
	"CTCN": {
		{
			ID:          5313,
			Slug:        "contracoin",
			Name:        "CONTRACOIN",
			BinanceUSDT: "CTCNUSDT",
		},
	},
	"IGG": {
		{
			ID:          4054,
			Slug:        "ig-gold",
			Name:        "IG Gold",
			BinanceUSDT: "IGGUSDT",
		},
	},
	"UBE": {
		{
			ID:          10808,
			Slug:        "ubeswap",
			Name:        "Ubeswap",
			BinanceUSDT: "UBEUSDT",
		},
	},
	"JUV": {
		{
			ID:          5224,
			Slug:        "juventus-fan-token",
			Name:        "Juventus Fan Token",
			BinanceUSDT: "JUVUSDT",
		},
	},
	"OWL": {
		{
			ID:          7330,
			Slug:        "owl-token-stealthswap",
			Name:        "OWL Token (StealthSwap)",
			BinanceUSDT: "OWLUSDT",
		},
	},
	"GNT": {
		{
			ID:          9533,
			Slug:        "greentrust",
			Name:        "GreenTrust",
			BinanceUSDT: "GNTUSDT",
		},
	},
	"HB": {
		{
			ID:          3171,
			Slug:        "heartbout",
			Name:        "HeartBout",
			BinanceUSDT: "HBUSDT",
		},
	},
	"PROS": {
		{
			ID:   8255,
			Slug: "prosper",
			Name: "Prosper",
		},
		{
			ID:   9807,
			Slug: "proswap",
			Name: "ProSwap",
		},
	},
	"CNTM": {
		{
			ID:          6039,
			Slug:        "connectome",
			Name:        "Connectome",
			BinanceUSDT: "CNTMUSDT",
		},
	},
	"ENV": {
		{
			ID:          9296,
			Slug:        "env-finance",
			Name:        "ENV Finance",
			BinanceUSDT: "ENVUSDT",
		},
	},
	"BXA": {
		{
			ID:          6225,
			Slug:        "blockchain-exchange-alliance",
			Name:        "Blockchain Exchange Alliance",
			BinanceUSDT: "BXAUSDT",
		},
	},
	"FLOT": {
		{
			ID:          3247,
			Slug:        "fire-lotto",
			Name:        "Fire Lotto",
			BinanceUSDT: "FLOTUSDT",
		},
	},
	"WOWS": {
		{
			ID:          8496,
			Slug:        "wolves-of-wall-street",
			Name:        "Wolves of Wall Street",
			BinanceUSDT: "WOWSUSDT",
		},
	},
	"AGVC": {
		{
			ID:          3664,
			Slug:        "agavecoin",
			Name:        "AgaveCoin",
			BinanceUSDT: "AGVCUSDT",
		},
	},
	"MAC": {
		{
			ID:          3895,
			Slug:        "matrexcoin",
			Name:        "Matrexcoin",
			BinanceUSDT: "MACUSDT",
		},
	},
	"UOP": {
		{
			ID:          8150,
			Slug:        "utopia-genesis-foundation",
			Name:        "Utopia Genesis Foundation",
			BinanceUSDT: "UOPUSDT",
		},
	},
	"DOV": {
		{
			ID:          2110,
			Slug:        "dovu",
			Name:        "Dovu",
			BinanceUSDT: "DOVUSDT",
		},
	},
	"VISR": {
		{
			ID:          9170,
			Slug:        "visor-finance",
			Name:        "Visor.Finance",
			BinanceUSDT: "VISRUSDT",
		},
	},
	"BELT": {
		{
			ID:          8730,
			Slug:        "belt",
			Name:        "Belt Finance",
			BinanceUSDT: "BELTUSDT",
		},
	},
	"CS": {
		{
			ID:          2556,
			Slug:        "credits",
			Name:        "Credits",
			BinanceUSDT: "CSUSDT",
		},
	},
	"PRCY": {
		{
			ID:          8554,
			Slug:        "prcy-coin",
			Name:        "PRCY Coin",
			BinanceUSDT: "PRCYUSDT",
		},
	},
	"NEAL": {
		{
			ID:          3960,
			Slug:        "coineal-token",
			Name:        "Coineal Token",
			BinanceUSDT: "NEALUSDT",
		},
	},
	"PTON": {
		{
			ID:          3813,
			Slug:        "pton",
			Name:        "PTON",
			BinanceUSDT: "PTONUSDT",
		},
	},
	"MUSE": {
		{
			ID:          7805,
			Slug:        "muse",
			Name:        "Muse",
			BinanceUSDT: "MUSEUSDT",
		},
	},
	"KENU": {
		{
			ID:          10635,
			Slug:        "ken-inu",
			Name:        "Ken Inu",
			BinanceUSDT: "KENUUSDT",
		},
	},
	"XAUTBEAR": {
		{
			ID:          6240,
			Slug:        "3x-short-tether-gold-token",
			Name:        "3X Short Tether Gold Token",
			BinanceUSDT: "XAUTBEARUSDT",
		},
	},
	"XC": {
		{
			ID:          9682,
			Slug:        "xcom",
			Name:        "XCOM",
			BinanceUSDT: "XCUSDT",
		},
	},
	"NEX": {
		{
			ID:          3829,
			Slug:        "nash-exchange",
			Name:        "Nash",
			BinanceUSDT: "NEXUSDT",
		},
	},
	"FIC": {
		{
			ID:          8383,
			Slug:        "filecash",
			Name:        "Filecash",
			BinanceUSDT: "FICUSDT",
		},
	},
	"RINGX": {
		{
			ID:          4894,
			Slug:        "ring-x-platform",
			Name:        "RING X PLATFORM",
			BinanceUSDT: "RINGXUSDT",
		},
	},
	"ATRI": {
		{
			ID:          7395,
			Slug:        "atari-token",
			Name:        "Atari Token",
			BinanceUSDT: "ATRIUSDT",
		},
	},
	"XRE": {
		{
			ID:          10154,
			Slug:        "xre-global",
			Name:        "XRE Global",
			BinanceUSDT: "XREUSDT",
		},
	},
	"SUSHIBULL": {
		{
			ID:          7084,
			Slug:        "3x-long-sushi-token",
			Name:        "3X Long Sushi Token",
			BinanceUSDT: "SUSHIBULLUSDT",
		},
	},
	"MVI": {
		{
			ID:          9207,
			Slug:        "metaverse-index",
			Name:        "Metaverse Index",
			BinanceUSDT: "MVIUSDT",
		},
	},
	"SVX": {
		{
			ID:          9109,
			Slug:        "savix",
			Name:        "Savix",
			BinanceUSDT: "SVXUSDT",
		},
	},
	"CRM": {
		{
			ID:          1850,
			Slug:        "cream",
			Name:        "Cream",
			BinanceUSDT: "CRMUSDT",
		},
	},
	"GRIMM": {
		{
			ID:          4866,
			Slug:        "grimm",
			Name:        "Grimm",
			BinanceUSDT: "GRIMMUSDT",
		},
	},
	"ELT": {
		{
			ID:          7884,
			Slug:        "elite-swap",
			Name:        "Elite Swap",
			BinanceUSDT: "ELTUSDT",
		},
	},
	"WSTA": {
		{
			ID:          9620,
			Slug:        "wrapped-statera",
			Name:        "Wrapped Statera",
			BinanceUSDT: "WSTAUSDT",
		},
	},
	"TAU": {
		{
			ID:          2337,
			Slug:        "lamden",
			Name:        "Lamden",
			BinanceUSDT: "TAUUSDT",
		},
	},
	"HUP": {
		{
			ID:          10678,
			Slug:        "hup-life",
			Name:        "HUP.LIFE",
			BinanceUSDT: "HUPUSDT",
		},
	},
	"CRL": {
		{
			ID:          8756,
			Slug:        "coralfarm",
			Name:        "CoralFarm",
			BinanceUSDT: "CRLUSDT",
		},
	},
	"IBEX": {
		{
			ID:          10256,
			Slug:        "ibex",
			Name:        "IBEX",
			BinanceUSDT: "IBEXUSDT",
		},
	},
	"BTCBAM": {
		{
			ID:          9270,
			Slug:        "bitcoin-bam",
			Name:        "Bitcoin Bam",
			BinanceUSDT: "BTCBAMUSDT",
		},
	},
	"FEI": {
		{
			ID:          8642,
			Slug:        "fei-protocol",
			Name:        "Fei Protocol",
			BinanceUSDT: "FEIUSDT",
		},
	},
	"CHORD": {
		{
			ID:          9199,
			Slug:        "chord-protocol",
			Name:        "Chord Protocol",
			BinanceUSDT: "CHORDUSDT",
		},
	},
	"RARE": {
		{
			ID:          8125,
			Slug:        "unique-one",
			Name:        "Unique One",
			BinanceUSDT: "RAREUSDT",
		},
	},
	"ENTRC": {
		{
			ID:          3876,
			Slug:        "entercoin",
			Name:        "EnterCoin",
			BinanceUSDT: "ENTRCUSDT",
		},
	},
	"OXY": {
		{
			ID:          8029,
			Slug:        "oxygen",
			Name:        "Oxygen",
			BinanceUSDT: "OXYUSDT",
		},
	},
	"IOOX": {
		{
			ID:          5733,
			Slug:        "ioox-system",
			Name:        "IOOX System",
			BinanceUSDT: "IOOXUSDT",
		},
	},
	"YLFI": {
		{
			ID:          7788,
			Slug:        "yearn-loans-finance",
			Name:        "Yearn Loans Finance",
			BinanceUSDT: "YLFIUSDT",
		},
	},
	"TCR": {
		{
			ID:          7717,
			Slug:        "tecracoin",
			Name:        "TecraCoin",
			BinanceUSDT: "TCRUSDT",
		},
	},
	"PXI": {
		{
			ID:          656,
			Slug:        "prime-xi",
			Name:        "Prime-XI",
			BinanceUSDT: "PXIUSDT",
		},
	},
	"WHIRL": {
		{
			ID:          8778,
			Slug:        "whirl-finance",
			Name:        "Whirl Finance",
			BinanceUSDT: "WHIRLUSDT",
		},
	},
	"ATL": {
		{
			ID:          2136,
			Slug:        "atlant",
			Name:        "ATLANT",
			BinanceUSDT: "ATLUSDT",
		},
	},
	"WRC": {
		{
			ID:          2288,
			Slug:        "worldcore",
			Name:        "Worldcore",
			BinanceUSDT: "WRCUSDT",
		},
	},
	"MERO": {
		{
			ID:          9059,
			Slug:        "mero",
			Name:        "Mero",
			BinanceUSDT: "MEROUSDT",
		},
	},
	"MOAC": {
		{
			ID:          2403,
			Slug:        "moac",
			Name:        "MOAC",
			BinanceUSDT: "MOACUSDT",
		},
	},
	"USE": {
		{
			ID:   9611,
			Slug: "elanausd",
			Name: "ElenaUSD",
		},
		{
			ID:   3249,
			Slug: "usechain-token",
			Name: "Usechain Token",
		},
	},
	"ENB": {
		{
			ID:          7716,
			Slug:        "earnbase",
			Name:        "Earnbase",
			BinanceUSDT: "ENBUSDT",
		},
	},
	"SPRKL": {
		{
			ID:          3780,
			Slug:        "sparkle-loyalty",
			Name:        "Sparkle Loyalty",
			BinanceUSDT: "SPRKLUSDT",
		},
	},
	"P2P": {
		{
			ID:          7075,
			Slug:        "p2p",
			Name:        "P2P",
			BinanceUSDT: "P2PUSDT",
		},
	},
	"KRT": {
		{
			ID:          5115,
			Slug:        "terra-krw",
			Name:        "TerraKRW",
			BinanceUSDT: "KRTUSDT",
		},
	},
	"ETH3L": {
		{
			ID:          5739,
			Slug:        "amun-ether-3x-daily-long",
			Name:        "Amun Ether 3x Daily Long",
			BinanceUSDT: "ETH3LUSDT",
		},
	},
	"KTLYO": {
		{
			ID:          7885,
			Slug:        "katalyo",
			Name:        "Katalyo",
			BinanceUSDT: "KTLYOUSDT",
		},
	},
	"QARK": {
		{
			ID:          5858,
			Slug:        "qanplatform",
			Name:        "QANplatform",
			BinanceUSDT: "QARKUSDT",
		},
	},
	"WINR": {
		{
			ID:          7494,
			Slug:        "justbet",
			Name:        "JustBet",
			BinanceUSDT: "WINRUSDT",
		},
	},
	"UPBNB": {
		{
			ID:          9671,
			Slug:        "upbnb",
			Name:        "upBNB",
			BinanceUSDT: "UPBNBUSDT",
		},
	},
	"TOS": {
		{
			ID:          2987,
			Slug:        "thingsoperatingsystem",
			Name:        "ThingsOperatingSystem",
			BinanceUSDT: "TOSUSDT",
		},
	},
	"MINT": {
		{
			ID:   8423,
			Slug: "public-mint",
			Name: "Public Mint",
		},
		{
			ID:   10977,
			Slug: "mint-club",
			Name: "Mint Club",
		},
		{
			ID:   141,
			Slug: "mintcoin",
			Name: "MintCoin",
		},
	},
	"BMXX": {
		{
			ID:          8618,
			Slug:        "bmultiplier",
			Name:        "Multiplier",
			BinanceUSDT: "BMXXUSDT",
		},
	},
	"LID": {
		{
			ID:          5986,
			Slug:        "liquidity-dividends-protocol",
			Name:        "Liquidity Dividends Protocol",
			BinanceUSDT: "LIDUSDT",
		},
	},
	"FRONT": {
		{
			ID:          5893,
			Slug:        "frontier",
			Name:        "Frontier",
			BinanceUSDT: "FRONTUSDT",
		},
	},
	"OXT": {
		{
			ID:          5026,
			Slug:        "orchid",
			Name:        "Orchid",
			BinanceUSDT: "OXTUSDT",
		},
	},
	"BILI": {
		{
			ID:   7895,
			Slug: "billibilli-inc-tokenized-stock-ftx",
			Name: "Billibilli Inc tokenized stock FTX",
		},
		{
			ID:   7930,
			Slug: "billibilli-tokenized-stock-bittrex",
			Name: "Billibilli tokenized stock Bittrex",
		},
	},
	"BLES": {
		{
			ID:          9026,
			Slug:        "blind-boxes",
			Name:        "Blind Boxes",
			BinanceUSDT: "BLESUSDT",
		},
	},
	"PIPT": {
		{
			ID:          7368,
			Slug:        "power-index-pool-token",
			Name:        "Power Index Pool Token",
			BinanceUSDT: "PIPTUSDT",
		},
	},
	"WCC": {
		{
			ID:          5233,
			Slug:        "wincash",
			Name:        "WinCash",
			BinanceUSDT: "WCCUSDT",
		},
	},
	"MARO": {
		{
			ID:          3175,
			Slug:        "maro",
			Name:        "Maro",
			BinanceUSDT: "MAROUSDT",
		},
	},
	"ONX": {
		{
			ID:   1747,
			Slug: "onix",
			Name: "Onix",
		},
		{
			ID:   8071,
			Slug: "onx-finance",
			Name: "OnX Finance",
		},
	},
	"HEDGE": {
		{
			ID:          5656,
			Slug:        "1x-short-bitcoin-token",
			Name:        "1x Short Bitcoin Token",
			BinanceUSDT: "HEDGEUSDT",
		},
	},
	"STAK": {
		{
			ID:   2332,
			Slug: "straks",
			Name: "STRAKS",
		},
		{
			ID:   9441,
			Slug: "jigstack",
			Name: "Jigstack",
		},
	},
	"WOLF": {
		{
			ID:          9158,
			Slug:        "moonwolf",
			Name:        "moonwolf.io",
			BinanceUSDT: "WOLFUSDT",
		},
	},
	"REEF": {
		{
			ID:          6951,
			Slug:        "reef",
			Name:        "Reef",
			BinanceUSDT: "REEFUSDT",
		},
	},
	"OFT": {
		{
			ID:          8095,
			Slug:        "orient",
			Name:        "Orient",
			BinanceUSDT: "OFTUSDT",
		},
	},
	"ALM": {
		{
			ID:          10428,
			Slug:        "alium-finance",
			Name:        "Alium Finance",
			BinanceUSDT: "ALMUSDT",
		},
	},
	"SPY": {
		{
			ID:   7915,
			Slug: "spdr-sp-500-etf-tokenized-stock-ftx",
			Name: "SPDR S&P 500 ETF tokenized stock FTX",
		},
		{
			ID:   7920,
			Slug: "spdr-sp-500-etf-tokenized-stock-bittrex",
			Name: "SPDR S&P 500 ETF tokenized stock Bittrex",
		},
	},
	"AXI": {
		{
			ID:          7141,
			Slug:        "axioms",
			Name:        "Axioms",
			BinanceUSDT: "AXIUSDT",
		},
	},
	"BABYBNB": {
		{
			ID:          10698,
			Slug:        "babybnb",
			Name:        "Babybnb",
			BinanceUSDT: "BABYBNBUSDT",
		},
	},
	"NFUP": {
		{
			ID:          8007,
			Slug:        "natural-farm-union-protocol",
			Name:        "Natural Farm Union Protocol",
			BinanceUSDT: "NFUPUSDT",
		},
	},
	"TRI": {
		{
			ID:          7995,
			Slug:        "trinity-protocol",
			Name:        "Trinity Protocol",
			BinanceUSDT: "TRIUSDT",
		},
	},
	"TWI": {
		{
			ID:          7620,
			Slug:        "trade-win",
			Name:        "Trade.win",
			BinanceUSDT: "TWIUSDT",
		},
	},
	"ATN": {
		{
			ID:          2380,
			Slug:        "atn",
			Name:        "ATN",
			BinanceUSDT: "ATNUSDT",
		},
	},
	"BCI": {
		{
			ID:          2702,
			Slug:        "bitcoin-interest",
			Name:        "Bitcoin Interest",
			BinanceUSDT: "BCIUSDT",
		},
	},
	"DOOGEE": {
		{
			ID:          10920,
			Slug:        "doogee-io",
			Name:        "Doogee.io",
			BinanceUSDT: "DOOGEEUSDT",
		},
	},
	"ALL": {
		{
			ID:          8882,
			Slug:        "alliance-fan-token",
			Name:        "Alliance Fan Token",
			BinanceUSDT: "ALLUSDT",
		},
	},
	"CHEX": {
		{
			ID:          8534,
			Slug:        "chex-token",
			Name:        "Chintai",
			BinanceUSDT: "CHEXUSDT",
		},
	},
	"SHARE": {
		{
			ID:          6868,
			Slug:        "seigniorage-shares",
			Name:        "Seigniorage Shares",
			BinanceUSDT: "SHAREUSDT",
		},
	},
	"HAM": {
		{
			ID:          10336,
			Slug:        "hamster",
			Name:        "Hamster",
			BinanceUSDT: "HAMUSDT",
		},
	},
	"SAFELIGHT": {
		{
			ID:          9208,
			Slug:        "safelight",
			Name:        "SafeLight",
			BinanceUSDT: "SAFELIGHTUSDT",
		},
	},
	"EMRALS": {
		{
			ID:          5259,
			Slug:        "emrals",
			Name:        "Emrals",
			BinanceUSDT: "EMRALSUSDT",
		},
	},
	"UHP": {
		{
			ID:          5949,
			Slug:        "ulgen-hash-power",
			Name:        "Ulgen Hash Power",
			BinanceUSDT: "UHPUSDT",
		},
	},
	"WEBN": {
		{
			ID:          3744,
			Slug:        "webn-token",
			Name:        "WEBN token",
			BinanceUSDT: "WEBNUSDT",
		},
	},
	"HYDRO": {
		{
			ID:          2698,
			Slug:        "hydro",
			Name:        "Hydro",
			BinanceUSDT: "HYDROUSDT",
		},
	},
	"ISP": {
		{
			ID:          9865,
			Slug:        "ispolink",
			Name:        "Ispolink",
			BinanceUSDT: "ISPUSDT",
		},
	},
	"PGT": {
		{
			ID:          7315,
			Slug:        "polyient-games-governance-token",
			Name:        "Polyient Games Governance Token",
			BinanceUSDT: "PGTUSDT",
		},
	},
	"HMQ": {
		{
			ID:          1669,
			Slug:        "humaniq",
			Name:        "Humaniq",
			BinanceUSDT: "HMQUSDT",
		},
	},
	"FDO": {
		{
			ID:          8176,
			Slug:        "firdaos",
			Name:        "Firdaos",
			BinanceUSDT: "FDOUSDT",
		},
	},
	"SKIN": {
		{
			ID:          1830,
			Slug:        "skincoin",
			Name:        "SkinCoin",
			BinanceUSDT: "SKINUSDT",
		},
	},
	"NBC": {
		{
			ID:          3095,
			Slug:        "niobium-coin",
			Name:        "Niobium Coin",
			BinanceUSDT: "NBCUSDT",
		},
	},
	"MAZ": {
		{
			ID:          6711,
			Slug:        "mazzuma",
			Name:        "Mazzuma",
			BinanceUSDT: "MAZUSDT",
		},
	},
	"FOX": {
		{
			ID:   8200,
			Slug: "fox-token",
			Name: "Shapeshift FOX Token",
		},
		{
			ID:   9239,
			Slug: "fox-finance",
			Name: "Fox Finance",
		},
	},
	"DEGOV": {
		{
			ID:          7839,
			Slug:        "degov",
			Name:        "Degov",
			BinanceUSDT: "DEGOVUSDT",
		},
	},
	"GLQ": {
		{
			ID:          9029,
			Slug:        "graphlinq-protocol",
			Name:        "Graphlinq Protocol",
			BinanceUSDT: "GLQUSDT",
		},
	},
	"MONK": {
		{
			ID:          2230,
			Slug:        "monk",
			Name:        "MONK",
			BinanceUSDT: "MONKUSDT",
		},
	},
	"AZZR": {
		{
			ID:          7413,
			Slug:        "azzure",
			Name:        "Azzure",
			BinanceUSDT: "AZZRUSDT",
		},
	},
	"IMT": {
		{
			ID:   3243,
			Slug: "moneytoken",
			Name: "Moneytoken",
		},
		{
			ID:   5806,
			Slug: "imsmart",
			Name: "Imsmart",
		},
	},
	"BFC": {
		{
			ID:   7817,
			Slug: "bifrost",
			Name: "Bifrost (BFC)",
		},
		{
			ID:   7432,
			Slug: "bit-financial",
			Name: "Bit Financial",
		},
	},
	"KEL": {
		{
			ID:          9261,
			Slug:        "kelvpn",
			Name:        "KelVPN",
			BinanceUSDT: "KELUSDT",
		},
	},
	"LIB": {
		{
			ID:          6700,
			Slug:        "libera",
			Name:        "Libera",
			BinanceUSDT: "LIBUSDT",
		},
	},
	"MUNCH": {
		{
			ID:          9272,
			Slug:        "munch-token",
			Name:        "Munch Token",
			BinanceUSDT: "MUNCHUSDT",
		},
	},
	"OCB": {
		{
			ID:          6869,
			Slug:        "blockmax",
			Name:        "BLOCKMAX",
			BinanceUSDT: "OCBUSDT",
		},
	},
	"HUA": {
		{
			ID:          9941,
			Slug:        "chihuahua",
			Name:        "Chihuahua",
			BinanceUSDT: "HUAUSDT",
		},
	},
	"AXEL": {
		{
			ID:          6216,
			Slug:        "axel",
			Name:        "AXEL",
			BinanceUSDT: "AXELUSDT",
		},
	},
	"SYL": {
		{
			ID:          9180,
			Slug:        "xsl-labs",
			Name:        "SYL",
			BinanceUSDT: "SYLUSDT",
		},
	},
	"TRAT": {
		{
			ID:          3925,
			Slug:        "tratok",
			Name:        "Tratok",
			BinanceUSDT: "TRATUSDT",
		},
	},
	"MXT": {
		{
			ID:   1266,
			Slug: "martexcoin",
			Name: "MarteXcoin",
		},
		{
			ID:   6607,
			Slug: "mixtrust",
			Name: "MixTrust",
		},
	},
	"TRIX": {
		{
			ID:          6801,
			Slug:        "triumphx",
			Name:        "TriumphX",
			BinanceUSDT: "TRIXUSDT",
		},
	},
	"PAXG": {
		{
			ID:          4705,
			Slug:        "pax-gold",
			Name:        "PAX Gold",
			BinanceUSDT: "PAXGUSDT",
		},
	},
	"RLY": {
		{
			ID:          8075,
			Slug:        "rally",
			Name:        "Rally",
			BinanceUSDT: "RLYUSDT",
		},
	},
	"CVAG": {
		{
			ID:          10918,
			Slug:        "new-crypto-village-accelerator",
			Name:        "Crypto Village Accelerator",
			BinanceUSDT: "CVAGUSDT",
		},
	},
	"BLV": {
		{
			ID:          7309,
			Slug:        "bellevue-network",
			Name:        "Bellevue Network",
			BinanceUSDT: "BLVUSDT",
		},
	},
	"NXS": {
		{
			ID:          789,
			Slug:        "nexus",
			Name:        "Nexus",
			BinanceUSDT: "NXSUSDT",
		},
	},
	"TCT": {
		{
			ID:   2364,
			Slug: "tokenclub",
			Name: "TokenClub",
		},
		{
			ID:   10721,
			Slug: "tacocat-token",
			Name: "TacoCat Token",
		},
	},
	"SETS": {
		{
			ID:          9785,
			Slug:        "sensitrust",
			Name:        "Sensitrust",
			BinanceUSDT: "SETSUSDT",
		},
	},
	"STONK": {
		{
			ID:          5843,
			Slug:        "stonk",
			Name:        "STONK",
			BinanceUSDT: "STONKUSDT",
		},
	},
	"DOT": {
		{
			ID:          6636,
			Slug:        "polkadot-new",
			Name:        "Polkadot",
			BinanceUSDT: "DOTUSDT",
		},
	},
	"XT": {
		{
			ID:          5999,
			Slug:        "xtcom-token",
			Name:        "XT.com Token",
			BinanceUSDT: "XTUSDT",
		},
	},
	"OXB": {
		{
			ID:          8649,
			Slug:        "oxbull-tech",
			Name:        "Oxbull.tech",
			BinanceUSDT: "OXBUSDT",
		},
	},
	"CIPX": {
		{
			ID:          4708,
			Slug:        "colletrix",
			Name:        "Colletrix",
			BinanceUSDT: "CIPXUSDT",
		},
	},
	"MANNA": {
		{
			ID:          1019,
			Slug:        "manna",
			Name:        "Manna",
			BinanceUSDT: "MANNAUSDT",
		},
	},
	"LTP": {
		{
			ID:          7790,
			Slug:        "lifetioncoin",
			Name:        "LifetionCoin",
			BinanceUSDT: "LTPUSDT",
		},
	},
	"AKNC": {
		{
			ID:          5750,
			Slug:        "aave-knc",
			Name:        "Aave KNC",
			BinanceUSDT: "AKNCUSDT",
		},
	},
	"VACAY": {
		{
			ID:          10815,
			Slug:        "vacay",
			Name:        "Vacay",
			BinanceUSDT: "VACAYUSDT",
		},
	},
	"MCT": {
		{
			ID:   3002,
			Slug: "master-contract-token",
			Name: "Master Contract Token",
		},
		{
			ID:   9700,
			Slug: "microtuber",
			Name: "MicroTuber",
		},
		{
			ID:   8448,
			Slug: "mcobit",
			Name: "MCOBIT",
		},
	},
	"OPS": {
		{
			ID:          10494,
			Slug:        "octopus-protocol",
			Name:        "Octopus Protocol",
			BinanceUSDT: "OPSUSDT",
		},
	},
	"FOTA": {
		{
			ID:          2472,
			Slug:        "fortuna",
			Name:        "Fortuna",
			BinanceUSDT: "FOTAUSDT",
		},
	},
	"BTO": {
		{
			ID:          2392,
			Slug:        "bottos",
			Name:        "Bottos",
			BinanceUSDT: "BTOUSDT",
		},
	},
	"ETI": {
		{
			ID:          3599,
			Slug:        "etherinc",
			Name:        "EtherInc",
			BinanceUSDT: "ETIUSDT",
		},
	},
	"BTU": {
		{
			ID:   3737,
			Slug: "btu-protocol",
			Name: "BTU Protocol",
		},
		{
			ID:   1022,
			Slug: "bittup",
			Name: "BITTUP",
		},
	},
	"TOMO": {
		{
			ID:          2570,
			Slug:        "tomochain",
			Name:        "TomoChain",
			BinanceUSDT: "TOMOUSDT",
		},
	},
	"BIRD": {
		{
			ID:   7795,
			Slug: "bird-money",
			Name: "Bird.Money",
		},
		{
			ID:   4207,
			Slug: "birdchain",
			Name: "Birdchain",
		},
		{
			ID:   9739,
			Slug: "bird-finance",
			Name: "Bird Finance",
		},
		{
			ID:   10727,
			Slug: "bird-finance-heco",
			Name: "Bird Finance(HECO)",
		},
	},
	"DFSOCIAL": {
		{
			ID:          8128,
			Slug:        "defisocial-gaming",
			Name:        "DFSocial Gaming",
			BinanceUSDT: "DFSOCIALUSDT",
		},
	},
	"CNX": {
		{
			ID:          5332,
			Slug:        "cofinex-coin",
			Name:        "Cofinex Coin",
			BinanceUSDT: "CNXUSDT",
		},
	},
	"VIVID": {
		{
			ID:          3008,
			Slug:        "vivid-coin",
			Name:        "Vivid Coin",
			BinanceUSDT: "VIVIDUSDT",
		},
	},
	"MOONTOKEN": {
		{
			ID:          9399,
			Slug:        "moon-token",
			Name:        "MoonToken",
			BinanceUSDT: "MOONTOKENUSDT",
		},
	},
	"CND": {
		{
			ID:          2043,
			Slug:        "cindicator",
			Name:        "Cindicator",
			BinanceUSDT: "CNDUSDT",
		},
	},
	"PLAT": {
		{
			ID:          3633,
			Slug:        "bitguild-plat",
			Name:        "BitGuild PLAT",
			BinanceUSDT: "PLATUSDT",
		},
	},
	"LINKBEAR": {
		{
			ID:          6094,
			Slug:        "3x-short-chainlink-token",
			Name:        "3X Short Chainlink Token",
			BinanceUSDT: "LINKBEARUSDT",
		},
	},
	"BIOS": {
		{
			ID:          10139,
			Slug:        "0x-nodes",
			Name:        "0x_nodes",
			BinanceUSDT: "BIOSUSDT",
		},
	},
	"PDATA": {
		{
			ID:          4088,
			Slug:        "pdata",
			Name:        "PDATA",
			BinanceUSDT: "PDATAUSDT",
		},
	},
	"BZX": {
		{
			ID:          3497,
			Slug:        "bitcoin-zero",
			Name:        "Bitcoin Zero",
			BinanceUSDT: "BZXUSDT",
		},
	},
	"CRT": {
		{
			ID:          6872,
			Slug:        "carrot",
			Name:        "Carrot",
			BinanceUSDT: "CRTUSDT",
		},
	},
	"TRO": {
		{
			ID:          8636,
			Slug:        "trodl",
			Name:        "Trodl",
			BinanceUSDT: "TROUSDT",
		},
	},
	"YFI2": {
		{
			ID:          6973,
			Slug:        "yearn2-finance",
			Name:        "YEARN2.FINANCE",
			BinanceUSDT: "YFI2USDT",
		},
	},
	"PXT": {
		{
			ID:          4192,
			Slug:        "populous-xbrl-token",
			Name:        "Populous XBRL Token",
			BinanceUSDT: "PXTUSDT",
		},
	},
	"LHB": {
		{
			ID:          8427,
			Slug:        "lendhub",
			Name:        "Lendhub",
			BinanceUSDT: "LHBUSDT",
		},
	},
	"IRA": {
		{
			ID:          4885,
			Slug:        "diligence",
			Name:        "Diligence",
			BinanceUSDT: "IRAUSDT",
		},
	},
	"SID": {
		{
			ID:          10559,
			Slug:        "shield-finance-protocol",
			Name:        "Shield Token",
			BinanceUSDT: "SIDUSDT",
		},
	},
	"GALI": {
		{
			ID:          3793,
			Slug:        "galilel",
			Name:        "Galilel",
			BinanceUSDT: "GALIUSDT",
		},
	},
	"YFIH2": {
		{
			ID:          10545,
			Slug:        "h2finance",
			Name:        "H2Finance",
			BinanceUSDT: "YFIH2USDT",
		},
	},
	"DNA": {
		{
			ID:   5234,
			Slug: "metaverse-dualchain-network-architecture",
			Name: "Metaverse Dualchain Network Architecture",
		},
		{
			ID:   2208,
			Slug: "encrypgen",
			Name: "EncrypGen",
		},
	},
	"GROW": {
		{
			ID:          9813,
			Slug:        "growingfi",
			Name:        "GrowingFi",
			BinanceUSDT: "GROWUSDT",
		},
	},
	"ERO": {
		{
			ID:          2249,
			Slug:        "eroscoin",
			Name:        "Eroscoin",
			BinanceUSDT: "EROUSDT",
		},
	},
	"EVAI": {
		{
			ID:          9805,
			Slug:        "evai-io",
			Name:        "Evai.io",
			BinanceUSDT: "EVAIUSDT",
		},
	},
	"TRCL": {
		{
			ID:          5800,
			Slug:        "treecle",
			Name:        "Treecle",
			BinanceUSDT: "TRCLUSDT",
		},
	},
	"FCH": {
		{
			ID:   5270,
			Slug: "freecash",
			Name: "Freecash",
		},
		{
			ID:   5661,
			Slug: "fanaticos-cash",
			Name: "Fanaticos Cash",
		},
	},
	"OCTAX": {
		{
			ID:          10588,
			Slug:        "octax-finance",
			Name:        "OctaX Finance",
			BinanceUSDT: "OCTAXUSDT",
		},
	},
	"YFO": {
		{
			ID:          8398,
			Slug:        "yfione",
			Name:        "YFIONE",
			BinanceUSDT: "YFOUSDT",
		},
	},
	"PUNK": {
		{
			ID:          8490,
			Slug:        "punk",
			Name:        "Punk",
			BinanceUSDT: "PUNKUSDT",
		},
	},
	"ETL": {
		{
			ID:          10832,
			Slug:        "etherlite",
			Name:        "Etherlite",
			BinanceUSDT: "ETLUSDT",
		},
	},
	"STAC": {
		{
			ID:          2565,
			Slug:        "startercoin",
			Name:        "StarterCoin",
			BinanceUSDT: "STACUSDT",
		},
	},
	"TDP": {
		{
			ID:          3469,
			Slug:        "truedeck",
			Name:        "TrueDeck",
			BinanceUSDT: "TDPUSDT",
		},
	},
	"SOLDIER": {
		{
			ID:          10067,
			Slug:        "space-soldier",
			Name:        "Space Soldier",
			BinanceUSDT: "SOLDIERUSDT",
		},
	},
	"HUNGRY": {
		{
			ID:          9231,
			Slug:        "hungry-bear",
			Name:        "Hungry Bear",
			BinanceUSDT: "HUNGRYUSDT",
		},
	},
	"MARSM": {
		{
			ID:          9293,
			Slug:        "marsmission-protocol",
			Name:        "MarsMission Protocol",
			BinanceUSDT: "MARSMUSDT",
		},
	},
	"RWS": {
		{
			ID:          5936,
			Slug:        "robonomics-web-services",
			Name:        "Robonomics Web Services",
			BinanceUSDT: "RWSUSDT",
		},
	},
	"SME": {
		{
			ID:          10153,
			Slug:        "safememe",
			Name:        "SafeMeme",
			BinanceUSDT: "SMEUSDT",
		},
	},
	"HACHIKO": {
		{
			ID:          9988,
			Slug:        "hachiko-inu",
			Name:        "Hachiko Inu",
			BinanceUSDT: "HACHIKOUSDT",
		},
	},
	"HDFL": {
		{
			ID:          10628,
			Slug:        "hyper-deflate",
			Name:        "Hyper Deflate",
			BinanceUSDT: "HDFLUSDT",
		},
	},
	"DG": {
		{
			ID:   7798,
			Slug: "decentral-games",
			Name: "Decentral Games",
		},
		{
			ID:   8833,
			Slug: "degate",
			Name: "DeGate",
		},
	},
	"MELLO": {
		{
			ID:          9339,
			Slug:        "mello-token",
			Name:        "Mello Token",
			BinanceUSDT: "MELLOUSDT",
		},
	},
	"KIND": {
		{
			ID:   3078,
			Slug: "kind-ads-token",
			Name: "Kind Ads Token",
		},
		{
			ID:   8802,
			Slug: "kindcow-finance",
			Name: "Kindcow Finance",
		},
	},
	"ITC": {
		{
			ID:          2251,
			Slug:        "iot-chain",
			Name:        "IoT Chain",
			BinanceUSDT: "ITCUSDT",
		},
	},
	"BDT": {
		{
			ID:   8285,
			Slug: "block-duelers",
			Name: "Block Duelers NFT Battles",
		},
		{
			ID:   9452,
			Slug: "bandot-protocol",
			Name: "Bandot Protocol",
		},
		{
			ID:   7763,
			Slug: "blackdragon",
			Name: "BlackDragon",
		},
	},
	"BSVBULL": {
		{
			ID:          5460,
			Slug:        "3x-long-bitcoin-sv-token",
			Name:        "3x Long Bitcoin SV Token",
			BinanceUSDT: "BSVBULLUSDT",
		},
	},
	"EXY": {
		{
			ID:          2547,
			Slug:        "experty",
			Name:        "Experty",
			BinanceUSDT: "EXYUSDT",
		},
	},
	"RAGNA": {
		{
			ID:          3374,
			Slug:        "ragnarok",
			Name:        "Ragnarok",
			BinanceUSDT: "RAGNAUSDT",
		},
	},
	"POOCOIN": {
		{
			ID:          8978,
			Slug:        "poocoin",
			Name:        "PooCoin",
			BinanceUSDT: "POOCOINUSDT",
		},
	},
	"SPKL": {
		{
			ID:          7198,
			Slug:        "spoklottery",
			Name:        "SpokLottery",
			BinanceUSDT: "SPKLUSDT",
		},
	},
	"ATCC": {
		{
			ID:          1751,
			Slug:        "atc-coin",
			Name:        "ATC Coin",
			BinanceUSDT: "ATCCUSDT",
		},
	},
	"BDCASH": {
		{
			ID:          7277,
			Slug:        "bigdatacash",
			Name:        "BigdataCash",
			BinanceUSDT: "BDCASHUSDT",
		},
	},
	"SYN": {
		{
			ID:          7520,
			Slug:        "synlev",
			Name:        "SynLev",
			BinanceUSDT: "SYNUSDT",
		},
	},
	"SBGO": {
		{
			ID:          9552,
			Slug:        "bingo-cash-finance",
			Name:        "Bingo Share",
			BinanceUSDT: "SBGOUSDT",
		},
	},
	"KVI": {
		{
			ID:          6031,
			Slug:        "kvi",
			Name:        "KVI",
			BinanceUSDT: "KVIUSDT",
		},
	},
	"YFS": {
		{
			ID:          9181,
			Slug:        "yfsfinance",
			Name:        "YFS.FINANCE",
			BinanceUSDT: "YFSUSDT",
		},
	},
	"ARC": {
		{
			ID:   1434,
			Slug: "arcticcoin",
			Name: "Advanced Technology Coin",
		},
		{
			ID:   6873,
			Slug: "arcx",
			Name: "ARCx (old)",
		},
	},
	"DATP": {
		{
			ID:          3454,
			Slug:        "decentralized-asset-trading-platform",
			Name:        "Decentralized Asset Trading Platform",
			BinanceUSDT: "DATPUSDT",
		},
	},
	"DXH": {
		{
			ID:          9317,
			Slug:        "daxhund",
			Name:        "Daxhund",
			BinanceUSDT: "DXHUSDT",
		},
	},
	"SWSH": {
		{
			ID:          7191,
			Slug:        "swapship",
			Name:        "SwapShip",
			BinanceUSDT: "SWSHUSDT",
		},
	},
	"FLO": {
		{
			ID:          8774,
			Slug:        "flourmix",
			Name:        "FlourMix",
			BinanceUSDT: "FLOUSDT",
		},
	},
	"BNOX": {
		{
			ID:          5648,
			Slug:        "blocknotex",
			Name:        "BlockNoteX",
			BinanceUSDT: "BNOXUSDT",
		},
	},
	"ZP": {
		{
			ID:          3186,
			Slug:        "zen-protocol",
			Name:        "Zen Protocol",
			BinanceUSDT: "ZPUSDT",
		},
	},
	"XMCT": {
		{
			ID:          2923,
			Slug:        "xmct",
			Name:        "XMCT",
			BinanceUSDT: "XMCTUSDT",
		},
	},
	"THR": {
		{
			ID:          3144,
			Slug:        "thorecoin",
			Name:        "ThoreCoin",
			BinanceUSDT: "THRUSDT",
		},
	},
	"WAGE": {
		{
			ID:          3416,
			Slug:        "digiwage",
			Name:        "Digiwage",
			BinanceUSDT: "WAGEUSDT",
		},
	},
	"HXN": {
		{
			ID:          8184,
			Slug:        "havens-nook",
			Name:        "Havens Nook",
			BinanceUSDT: "HXNUSDT",
		},
	},
	"DTC": {
		{
			ID:          69,
			Slug:        "datacoin",
			Name:        "Datacoin",
			BinanceUSDT: "DTCUSDT",
		},
	},
	"EMRX": {
		{
			ID:          4490,
			Slug:        "emirex-token",
			Name:        "Emirex Token",
			BinanceUSDT: "EMRXUSDT",
		},
	},
	"VTHO": {
		{
			ID:          3012,
			Slug:        "vethor-token",
			Name:        "VeThor Token",
			BinanceUSDT: "VTHOUSDT",
		},
	},
	"CAI": {
		{
			ID:          7639,
			Slug:        "club-atletico-independiente",
			Name:        "Club Atletico Independiente",
			BinanceUSDT: "CAIUSDT",
		},
	},
	"YFA": {
		{
			ID:          6909,
			Slug:        "yfa-finance",
			Name:        "YFA Finance",
			BinanceUSDT: "YFAUSDT",
		},
	},
	"BKK": {
		{
			ID:          5154,
			Slug:        "bkex-token",
			Name:        "BKEX Token",
			BinanceUSDT: "BKKUSDT",
		},
	},
	"NEWINU": {
		{
			ID:          10371,
			Slug:        "newinu",
			Name:        "Newinu",
			BinanceUSDT: "NEWINUUSDT",
		},
	},
	"LPL": {
		{
			ID:          9062,
			Slug:        "linkpool",
			Name:        "LinkPool",
			BinanceUSDT: "LPLUSDT",
		},
	},
	"TLW": {
		{
			ID:          5399,
			Slug:        "tilwiki",
			Name:        "TILWIKI",
			BinanceUSDT: "TLWUSDT",
		},
	},
	"BUILD": {
		{
			ID:          6956,
			Slug:        "build-finance",
			Name:        "BUILD Finance",
			BinanceUSDT: "BUILDUSDT",
		},
	},
	"GOM": {
		{
			ID:          4623,
			Slug:        "gomics",
			Name:        "Gomics",
			BinanceUSDT: "GOMUSDT",
		},
	},
	"IFX24": {
		{
			ID:          5170,
			Slug:        "ifx24",
			Name:        "IFX24",
			BinanceUSDT: "IFX24USDT",
		},
	},
	"MOFI": {
		{
			ID:          9238,
			Slug:        "mofi-finance",
			Name:        "Mofi Finance",
			BinanceUSDT: "MOFIUSDT",
		},
	},
	"AITRA": {
		{
			ID:          7255,
			Slug:        "aitra",
			Name:        "Aitra",
			BinanceUSDT: "AITRAUSDT",
		},
	},
	"DACC": {
		{
			ID:          2986,
			Slug:        "dacc",
			Name:        "DACC",
			BinanceUSDT: "DACCUSDT",
		},
	},
	"DHOLD": {
		{
			ID:          10230,
			Slug:        "diamondhold",
			Name:        "DiamondHold",
			BinanceUSDT: "DHOLDUSDT",
		},
	},
	"ETHBACK": {
		{
			ID:          11002,
			Slug:        "etherback",
			Name:        "EtherBack",
			BinanceUSDT: "ETHBACKUSDT",
		},
	},
	"APRIL": {
		{
			ID:          10367,
			Slug:        "april",
			Name:        "April",
			BinanceUSDT: "APRILUSDT",
		},
	},
	"FCD": {
		{
			ID:          7640,
			Slug:        "future-cash-digital",
			Name:        "Future-Cash Digital",
			BinanceUSDT: "FCDUSDT",
		},
	},
	"ADMONKEY": {
		{
			ID:          11012,
			Slug:        "admonkey",
			Name:        "AdMonkey",
			BinanceUSDT: "ADMONKEYUSDT",
		},
	},
	"CUMINU": {
		{
			ID:          10061,
			Slug:        "cuminu",
			Name:        "CumInu",
			BinanceUSDT: "CUMINUUSDT",
		},
	},
	"FR": {
		{
			ID:          7712,
			Slug:        "freedom-reserve",
			Name:        "Freedom Reserve",
			BinanceUSDT: "FRUSDT",
		},
	},
	"XNO": {
		{
			ID:          8368,
			Slug:        "xeno-token",
			Name:        "Xeno Token",
			BinanceUSDT: "XNOUSDT",
		},
	},
	"WASP": {
		{
			ID:          8147,
			Slug:        "wanswap",
			Name:        "WanSwap",
			BinanceUSDT: "WASPUSDT",
		},
	},
	"MMO": {
		{
			ID:          3449,
			Slug:        "mmocoin",
			Name:        "MMOCoin",
			BinanceUSDT: "MMOUSDT",
		},
	},
	"VANCII": {
		{
			ID:          8671,
			Slug:        "vanci-finance",
			Name:        "VANCI FINANCE",
			BinanceUSDT: "VANCIIUSDT",
		},
	},
	"ABBC": {
		{
			ID:          3437,
			Slug:        "abbc-coin",
			Name:        "ABBC Coin",
			BinanceUSDT: "ABBCUSDT",
		},
	},
	"GEX": {
		{
			ID:          4150,
			Slug:        "globex",
			Name:        "GLOBEX",
			BinanceUSDT: "GEXUSDT",
		},
	},
	"BITB": {
		{
			ID:          819,
			Slug:        "bean-cash",
			Name:        "Bean Cash",
			BinanceUSDT: "BITBUSDT",
		},
	},
	"MGP": {
		{
			ID:   6526,
			Slug: "mangochain",
			Name: "MangoChain",
		},
		{
			ID:   10718,
			Slug: "micro-gaming-protocol",
			Name: "Micro Gaming Protocol",
		},
	},
	"APY": {
		{
			ID:          7227,
			Slug:        "apy-finance",
			Name:        "APY.Finance",
			BinanceUSDT: "APYUSDT",
		},
	},
	"ARPA": {
		{
			ID:          4039,
			Slug:        "arpa-chain",
			Name:        "ARPA Chain",
			BinanceUSDT: "ARPAUSDT",
		},
	},
	"MARS": {
		{
			ID:   7579,
			Slug: "mars-network",
			Name: "Mars Network",
		},
		{
			ID:   154,
			Slug: "marscoin",
			Name: "Marscoin",
		},
	},
	"B21": {
		{
			ID:          8060,
			Slug:        "b21-invest",
			Name:        "B21 Invest",
			BinanceUSDT: "B21USDT",
		},
	},
	"MATIC": {
		{
			ID:          3890,
			Slug:        "polygon",
			Name:        "Polygon",
			BinanceUSDT: "MATICUSDT",
		},
	},
	"TRTL": {
		{
			ID:          2958,
			Slug:        "turtlecoin",
			Name:        "TurtleCoin",
			BinanceUSDT: "TRTLUSDT",
		},
	},
	"MPL": {
		{
			ID:          9417,
			Slug:        "maple",
			Name:        "Maple",
			BinanceUSDT: "MPLUSDT",
		},
	},
	"SPAZ": {
		{
			ID:          4797,
			Slug:        "swapcoinz",
			Name:        "Swapcoinz",
			BinanceUSDT: "SPAZUSDT",
		},
	},
	"LOUVRE": {
		{
			ID:          10782,
			Slug:        "louvre-finance",
			Name:        "Louvre Finance",
			BinanceUSDT: "LOUVREUSDT",
		},
	},
	"WICKED": {
		{
			ID:          10422,
			Slug:        "the-witcher-fans",
			Name:        "The Witcher Fans",
			BinanceUSDT: "WICKEDUSDT",
		},
	},
	"PPC": {
		{
			ID:   5,
			Slug: "peercoin",
			Name: "Peercoin",
		},
		{
			ID:   6253,
			Slug: "philips-pay-coin",
			Name: "PHILLIPS PAY COIN",
		},
	},
	"REDFEG": {
		{
			ID:          11033,
			Slug:        "redfeg",
			Name:        "RedFEG",
			BinanceUSDT: "REDFEGUSDT",
		},
	},
	"WOO": {
		{
			ID:          7501,
			Slug:        "wootrade",
			Name:        "Wootrade",
			BinanceUSDT: "WOOUSDT",
		},
	},
	"KWH": {
		{
			ID:          3001,
			Slug:        "kwhcoin",
			Name:        "KWHCoin",
			BinanceUSDT: "KWHUSDT",
		},
	},
	"YFUEL": {
		{
			ID:          6913,
			Slug:        "yfuel",
			Name:        "YFUEL",
			BinanceUSDT: "YFUELUSDT",
		},
	},
	"ESH": {
		{
			ID:          4096,
			Slug:        "switch",
			Name:        "Switch",
			BinanceUSDT: "ESHUSDT",
		},
	},
	"VYBE": {
		{
			ID:          7025,
			Slug:        "vybe",
			Name:        "Vybe",
			BinanceUSDT: "VYBEUSDT",
		},
	},
	"XEP": {
		{
			ID:          8216,
			Slug:        "electra-protocol",
			Name:        "Electra Protocol",
			BinanceUSDT: "XEPUSDT",
		},
	},
	"SPIKE": {
		{
			ID:          4094,
			Slug:        "spiking",
			Name:        "Spiking",
			BinanceUSDT: "SPIKEUSDT",
		},
	},
	"COSHI": {
		{
			ID:          9574,
			Slug:        "coshi-inu",
			Name:        "CoShi Inu",
			BinanceUSDT: "COSHIUSDT",
		},
	},
	"NRP": {
		{
			ID:          3397,
			Slug:        "neural-protocol",
			Name:        "Neural Protocol",
			BinanceUSDT: "NRPUSDT",
		},
	},
	"RAZOR": {
		{
			ID:          8409,
			Slug:        "razor-network",
			Name:        "Razor Network",
			BinanceUSDT: "RAZORUSDT",
		},
	},
	"KZC": {
		{
			ID:          2386,
			Slug:        "kz-cash",
			Name:        "KZ Cash",
			BinanceUSDT: "KZCUSDT",
		},
	},
	"BSV": {
		{
			ID:          3602,
			Slug:        "bitcoin-sv",
			Name:        "Bitcoin SV",
			BinanceUSDT: "BSVUSDT",
		},
	},
	"AVAX": {
		{
			ID:          5805,
			Slug:        "avalanche",
			Name:        "Avalanche",
			BinanceUSDT: "AVAXUSDT",
		},
	},
	"PERA": {
		{
			ID:          10225,
			Slug:        "pera-finance",
			Name:        "Pera Finance",
			BinanceUSDT: "PERAUSDT",
		},
	},
	"N1CE": {
		{
			ID:          10689,
			Slug:        "n1ce",
			Name:        "N1CE",
			BinanceUSDT: "N1CEUSDT",
		},
	},
	"WSOTE": {
		{
			ID:          8381,
			Slug:        "soteria",
			Name:        "Soteria",
			BinanceUSDT: "WSOTEUSDT",
		},
	},
	"KSF": {
		{
			ID:          9871,
			Slug:        "kesef-finance",
			Name:        "Kesef Finance",
			BinanceUSDT: "KSFUSDT",
		},
	},
	"CVP": {
		{
			ID:          6669,
			Slug:        "powerpool",
			Name:        "PowerPool",
			BinanceUSDT: "CVPUSDT",
		},
	},
	"FACE": {
		{
			ID:          2775,
			Slug:        "faceter",
			Name:        "Faceter",
			BinanceUSDT: "FACEUSDT",
		},
	},
	"POX": {
		{
			ID:          6682,
			Slug:        "pollux-coin",
			Name:        "Pollux Coin",
			BinanceUSDT: "POXUSDT",
		},
	},
	"888": {
		{
			ID:          6051,
			Slug:        "888tron",
			Name:        "888tron",
			BinanceUSDT: "888USDT",
		},
	},
	"ETHDOWN": {
		{
			ID:          10853,
			Slug:        "eth-down",
			Name:        "ETHDOWN",
			BinanceUSDT: "ETHDOWNUSDT",
		},
	},
	"PINK": {
		{
			ID:   313,
			Slug: "pinkcoin",
			Name: "Pinkcoin",
		},
		{
			ID:   9740,
			Slug: "dot-finance",
			Name: "Dot Finance",
		},
	},
	"GOSS": {
		{
			ID:          3332,
			Slug:        "gossipcoin",
			Name:        "Gossip Coin",
			BinanceUSDT: "GOSSUSDT",
		},
	},
	"GMT": {
		{
			ID:   10180,
			Slug: "gomining-token",
			Name: "GoMining token",
		},
		{
			ID:   9203,
			Slug: "gambit-finance",
			Name: "Gambit",
		},
	},
	"DJ15": {
		{
			ID:          8011,
			Slug:        "davincij15-token",
			Name:        "Davincij15 Token",
			BinanceUSDT: "DJ15USDT",
		},
	},
	"VRC": {
		{
			ID:          323,
			Slug:        "vericoin",
			Name:        "VeriCoin",
			BinanceUSDT: "VRCUSDT",
		},
	},
	"WOOP": {
		{
			ID:          8937,
			Slug:        "woonkly-power",
			Name:        "Woonkly Power",
			BinanceUSDT: "WOOPUSDT",
		},
	},
	"CAD": {
		{
			ID:          8564,
			Slug:        "candy-protocol",
			Name:        "Candy Protocol",
			BinanceUSDT: "CADUSDT",
		},
	},
	"MOTA": {
		{
			ID:          4028,
			Slug:        "motacoin",
			Name:        "MotaCoin",
			BinanceUSDT: "MOTAUSDT",
		},
	},
	"EMC": {
		{
			ID:   558,
			Slug: "emercoin",
			Name: "Emercoin",
		},
		{
			ID:   7663,
			Slug: "edumetrix",
			Name: "EduMetrix Coin",
		},
	},
	"MCASH": {
		{
			ID:          4224,
			Slug:        "mcashchain",
			Name:        "Mcashchain",
			BinanceUSDT: "MCASHUSDT",
		},
	},
	"FRMS": {
		{
			ID:          6916,
			Slug:        "the-forms",
			Name:        "The Forms",
			BinanceUSDT: "FRMSUSDT",
		},
	},
	"APLP": {
		{
			ID:          7837,
			Slug:        "apple-finance",
			Name:        "Apple Finance",
			BinanceUSDT: "APLPUSDT",
		},
	},
	"CHIK": {
		{
			ID:          8518,
			Slug:        "chickenkebab-finance",
			Name:        "Chickenkebab Finance",
			BinanceUSDT: "CHIKUSDT",
		},
	},
	"SIAM": {
		{
			ID:          10711,
			Slug:        "siamese-neko",
			Name:        "Siamese Neko",
			BinanceUSDT: "SIAMUSDT",
		},
	},
	"POA": {
		{
			ID:          2548,
			Slug:        "poa",
			Name:        "POA",
			BinanceUSDT: "POAUSDT",
		},
	},
	"APL": {
		{
			ID:   2992,
			Slug: "apollo-currency",
			Name: "Apollo Currency",
		},
		{
			ID:   7638,
			Slug: "apollon-limassol",
			Name: "Apollon Limassol",
		},
	},
	"AVA": {
		{
			ID:          2776,
			Slug:        "travala",
			Name:        "Travala.com",
			BinanceUSDT: "AVAUSDT",
		},
	},
	"SFX": {
		{
			ID:          4183,
			Slug:        "safex-cash",
			Name:        "Safex Cash",
			BinanceUSDT: "SFXUSDT",
		},
	},
	"PDAO": {
		{
			ID:          8735,
			Slug:        "panda-dao",
			Name:        "Panda Dao",
			BinanceUSDT: "PDAOUSDT",
		},
	},
	"ZETH": {
		{
			ID:          10284,
			Slug:        "zetta-ethereum-hashrate-token",
			Name:        "Zetta Ethereum Hashrate Token",
			BinanceUSDT: "ZETHUSDT",
		},
	},
	"BAST": {
		{
			ID:          6821,
			Slug:        "bast",
			Name:        "Bast",
			BinanceUSDT: "BASTUSDT",
		},
	},
	"SLV": {
		{
			ID:          4079,
			Slug:        "silverway",
			Name:        "Silverway",
			BinanceUSDT: "SLVUSDT",
		},
	},
	"MDA": {
		{
			ID:          1954,
			Slug:        "moeda-loyalty-points",
			Name:        "Moeda Loyalty Points",
			BinanceUSDT: "MDAUSDT",
		},
	},
	"CTXC": {
		{
			ID:          2638,
			Slug:        "cortex",
			Name:        "Cortex",
			BinanceUSDT: "CTXCUSDT",
		},
	},
	"NYANTE": {
		{
			ID:          8067,
			Slug:        "nyantereum",
			Name:        "Nyantereum International",
			BinanceUSDT: "NYANTEUSDT",
		},
	},
	"YOK": {
		{
			ID:          6998,
			Slug:        "yokcoin",
			Name:        "YOKcoin",
			BinanceUSDT: "YOKUSDT",
		},
	},
	"NTX": {
		{
			ID:          8500,
			Slug:        "nitroex",
			Name:        "Nitroex",
			BinanceUSDT: "NTXUSDT",
		},
	},
	"DIGG": {
		{
			ID:          8307,
			Slug:        "digg",
			Name:        "DIGG",
			BinanceUSDT: "DIGGUSDT",
		},
	},
	"TINIDAWG": {
		{
			ID:          10694,
			Slug:        "minidog-finance",
			Name:        "MiniDog Finance",
			BinanceUSDT: "TINIDAWGUSDT",
		},
	},
	"NOTE": {
		{
			ID:          184,
			Slug:        "dnotes",
			Name:        "DNotes",
			BinanceUSDT: "NOTEUSDT",
		},
	},
	"HIT": {
		{
			ID:   3182,
			Slug: "hitchain",
			Name: "HitChain",
		},
		{
			ID:   10506,
			Slug: "hitbtc-token",
			Name: "HitBTC Token",
		},
	},
	"ADK": {
		{
			ID:          1706,
			Slug:        "aidos-kuneen",
			Name:        "Aidos Kuneen",
			BinanceUSDT: "ADKUSDT",
		},
	},
	"MAS": {
		{
			ID:          3352,
			Slug:        "midasprotocol",
			Name:        "MidasProtocol",
			BinanceUSDT: "MASUSDT",
		},
	},
	"DXF": {
		{
			ID:          8079,
			Slug:        "dexfin",
			Name:        "Dexfin",
			BinanceUSDT: "DXFUSDT",
		},
	},
	"xSAT": {
		{
			ID:          9411,
			Slug:        "satisfinance-token",
			Name:        "SatisFinance Token",
			BinanceUSDT: "xSATUSDT",
		},
	},
	"CNNS": {
		{
			ID:          3934,
			Slug:        "cnns",
			Name:        "CNNS",
			BinanceUSDT: "CNNSUSDT",
		},
	},
	"4ART": {
		{
			ID:          5912,
			Slug:        "4artechnologies",
			Name:        "4ART Coin",
			BinanceUSDT: "4ARTUSDT",
		},
	},
	"UFC": {
		{
			ID:          5901,
			Slug:        "union-fair-coin",
			Name:        "Union Fair Coin",
			BinanceUSDT: "UFCUSDT",
		},
	},
	"WCCX": {
		{
			ID:          7600,
			Slug:        "wrapped-conceal",
			Name:        "Wrapped Conceal",
			BinanceUSDT: "WCCXUSDT",
		},
	},
	"NANOX": {
		{
			ID:          1691,
			Slug:        "project-x",
			Name:        "Project-X",
			BinanceUSDT: "NANOXUSDT",
		},
	},
	"UFFYI": {
		{
			ID:          7273,
			Slug:        "unlimited-fiscusfyi",
			Name:        "Unlimited FiscusFYI",
			BinanceUSDT: "UFFYIUSDT",
		},
	},
	"GOT": {
		{
			ID:   3251,
			Slug: "parkingo",
			Name: "ParkinGo",
		},
		{
			ID:   2898,
			Slug: "gonetwork",
			Name: "GoNetwork",
		},
	},
	"VSYS": {
		{
			ID:          3704,
			Slug:        "v-systems",
			Name:        "v.systems",
			BinanceUSDT: "VSYSUSDT",
		},
	},
	"mUSO": {
		{
			ID:          8027,
			Slug:        "mirrored-united-states-oil-fund",
			Name:        "Mirrored United States Oil Fund",
			BinanceUSDT: "mUSOUSDT",
		},
	},
	"1WO": {
		{
			ID:          2601,
			Slug:        "1world",
			Name:        "1World",
			BinanceUSDT: "1WOUSDT",
		},
	},
	"CAVO": {
		{
			ID:          8286,
			Slug:        "excavo-finance",
			Name:        "Excavo Finance",
			BinanceUSDT: "CAVOUSDT",
		},
	},
	"YEC": {
		{
			ID:          4160,
			Slug:        "ycash",
			Name:        "Ycash",
			BinanceUSDT: "YECUSDT",
		},
	},
	"YI12": {
		{
			ID:          6912,
			Slug:        "yield-stake-finance",
			Name:        "Yield Stake Finance",
			BinanceUSDT: "YI12USDT",
		},
	},
	"UNDO": {
		{
			ID:          10589,
			Slug:        "undotoken",
			Name:        "UndoToken",
			BinanceUSDT: "UNDOUSDT",
		},
	},
	"GNO": {
		{
			ID:          1659,
			Slug:        "gnosis-gno",
			Name:        "Gnosis",
			BinanceUSDT: "GNOUSDT",
		},
	},
	"CBT": {
		{
			ID:          9342,
			Slug:        "community-business-token",
			Name:        "Community Business Token",
			BinanceUSDT: "CBTUSDT",
		},
	},
	"SHMN": {
		{
			ID:          3480,
			Slug:        "stronghands-masternode",
			Name:        "StrongHands Masternode",
			BinanceUSDT: "SHMNUSDT",
		},
	},
	"YMAX": {
		{
			ID:          7124,
			Slug:        "ymax",
			Name:        "YMAX",
			BinanceUSDT: "YMAXUSDT",
		},
	},
	"ECOS": {
		{
			ID:          5885,
			Slug:        "ecodollar",
			Name:        "EcoDollar",
			BinanceUSDT: "ECOSUSDT",
		},
	},
	"CBANK": {
		{
			ID:          8113,
			Slug:        "cryptobank",
			Name:        "CryptoBank",
			BinanceUSDT: "CBANKUSDT",
		},
	},
	"BCU": {
		{
			ID:          8777,
			Slug:        "biscuit-farm-finance",
			Name:        "Biscuit Farm Finance",
			BinanceUSDT: "BCUUSDT",
		},
	},
	"VDG": {
		{
			ID:          3205,
			Slug:        "veridocglobal",
			Name:        "VeriDocGlobal",
			BinanceUSDT: "VDGUSDT",
		},
	},
	"DINK": {
		{
			ID:          10710,
			Slug:        "dink-doink",
			Name:        "Dink Doink",
			BinanceUSDT: "DINKUSDT",
		},
	},
	"TCFX": {
		{
			ID:          7175,
			Slug:        "tcbcoin",
			Name:        "Tcbcoin",
			BinanceUSDT: "TCFXUSDT",
		},
	},
	"LBK": {
		{
			ID:   5445,
			Slug: "lbk",
			Name: "LBK",
		},
		{
			ID:   5255,
			Slug: "legalblock",
			Name: "LegalBlock",
		},
	},
	"OKBBULL": {
		{
			ID:          6087,
			Slug:        "3x-long-okb-token",
			Name:        "3X Long OKB Token",
			BinanceUSDT: "OKBBULLUSDT",
		},
	},
	"XGG": {
		{
			ID:          8685,
			Slug:        "10xgg",
			Name:        "10x.gg",
			BinanceUSDT: "XGGUSDT",
		},
	},
	"ANJ": {
		{
			ID:          5523,
			Slug:        "aragon-court",
			Name:        "Aragon Court",
			BinanceUSDT: "ANJUSDT",
		},
	},
	"WIZ": {
		{
			ID:          3331,
			Slug:        "crowdwiz",
			Name:        "CrowdWiz",
			BinanceUSDT: "WIZUSDT",
		},
	},
	"SAINT": {
		{
			ID:          10317,
			Slug:        "saint-token",
			Name:        "Saint Token",
			BinanceUSDT: "SAINTUSDT",
		},
	},
	"CSTL": {
		{
			ID:          3311,
			Slug:        "castle",
			Name:        "Castle",
			BinanceUSDT: "CSTLUSDT",
		},
	},
	"CCX": {
		{
			ID:          3732,
			Slug:        "conceal",
			Name:        "Conceal",
			BinanceUSDT: "CCXUSDT",
		},
	},
	"CCT": {
		{
			ID:          5054,
			Slug:        "coupon-chain",
			Name:        "Coupon Chain",
			BinanceUSDT: "CCTUSDT",
		},
	},
	"SADA": {
		{
			ID:          5769,
			Slug:        "sada",
			Name:        "sADA",
			BinanceUSDT: "SADAUSDT",
		},
	},
	"SBE": {
		{
			ID:          4538,
			Slug:        "sombe",
			Name:        "Sombe",
			BinanceUSDT: "SBEUSDT",
		},
	},
	"XTK": {
		{
			ID:          8599,
			Slug:        "xtoken",
			Name:        "xToken",
			BinanceUSDT: "XTKUSDT",
		},
	},
	"TRXC": {
		{
			ID:          3354,
			Slug:        "tronclassic",
			Name:        "TRONCLASSIC",
			BinanceUSDT: "TRXCUSDT",
		},
	},
	"SMLY": {
		{
			ID:          799,
			Slug:        "smileycoin",
			Name:        "SmileyCoin",
			BinanceUSDT: "SMLYUSDT",
		},
	},
	"DOP": {
		{
			ID:   9675,
			Slug: "drops",
			Name: "Drops Ownership Power",
		},
		{
			ID:   9489,
			Slug: "dopple-finance",
			Name: "Dopple Finance",
		},
	},
	"SCORGI": {
		{
			ID:          9716,
			Slug:        "spacecorgi",
			Name:        "SpaceCorgi",
			BinanceUSDT: "SCORGIUSDT",
		},
	},
	"PMP": {
		{
			ID:          9090,
			Slug:        "pumpy-farm",
			Name:        "Pumpy farm",
			BinanceUSDT: "PMPUSDT",
		},
	},
	"ARNX": {
		{
			ID:          2153,
			Slug:        "aeron",
			Name:        "Aeron",
			BinanceUSDT: "ARNXUSDT",
		},
	},
	"BTCUP": {
		{
			ID:          5608,
			Slug:        "btcup",
			Name:        "BTCUP",
			BinanceUSDT: "BTCUPUSDT",
		},
	},
	"FOXT": {
		{
			ID:          2966,
			Slug:        "fox-trading",
			Name:        "Fox Trading",
			BinanceUSDT: "FOXTUSDT",
		},
	},
	"STOS": {
		{
			ID:          9760,
			Slug:        "stratos",
			Name:        "Stratos",
			BinanceUSDT: "STOSUSDT",
		},
	},
	"APE": {
		{
			ID:          7257,
			Slug:        "apecoin",
			Name:        "APEcoin",
			BinanceUSDT: "APEUSDT",
		},
	},
	"PBR": {
		{
			ID:          8320,
			Slug:        "polkabridge",
			Name:        "PolkaBridge",
			BinanceUSDT: "PBRUSDT",
		},
	},
	"UOS": {
		{
			ID:          4189,
			Slug:        "ultra",
			Name:        "Ultra",
			BinanceUSDT: "UOSUSDT",
		},
	},
	"ROOBEE": {
		{
			ID:          4804,
			Slug:        "roobee",
			Name:        "ROOBEE",
			BinanceUSDT: "ROOBEEUSDT",
		},
	},
	"BEL": {
		{
			ID:          6928,
			Slug:        "bella-protocol",
			Name:        "Bella Protocol",
			BinanceUSDT: "BELUSDT",
		},
	},
	"KIF": {
		{
			ID:          6883,
			Slug:        "kittenfinance",
			Name:        "KittenFinance",
			BinanceUSDT: "KIFUSDT",
		},
	},
	"PEAK": {
		{
			ID:          5354,
			Slug:        "peakdefi",
			Name:        "PEAKDEFI",
			BinanceUSDT: "PEAKUSDT",
		},
	},
	"GMAT": {
		{
			ID:          4182,
			Slug:        "gowithmi",
			Name:        "GoWithMi",
			BinanceUSDT: "GMATUSDT",
		},
	},
	"MACH": {
		{
			ID:          4984,
			Slug:        "mach-project",
			Name:        "MACH Project",
			BinanceUSDT: "MACHUSDT",
		},
	},
	"MARSH": {
		{
			ID:          8963,
			Slug:        "unmarshal-token",
			Name:        "UnMarshal",
			BinanceUSDT: "MARSHUSDT",
		},
	},
	"RNX": {
		{
			ID:          6531,
			Slug:        "roonex",
			Name:        "ROONEX",
			BinanceUSDT: "RNXUSDT",
		},
	},
	"MASS": {
		{
			ID:          5548,
			Slug:        "massnet",
			Name:        "Massnet",
			BinanceUSDT: "MASSUSDT",
		},
	},
	"VINCI": {
		{
			ID:          4946,
			Slug:        "vinci",
			Name:        "Vinci",
			BinanceUSDT: "VINCIUSDT",
		},
	},
	"COB": {
		{
			ID:          2006,
			Slug:        "cobinhood",
			Name:        "Cobinhood",
			BinanceUSDT: "COBUSDT",
		},
	},
	"ELG": {
		{
			ID:          10312,
			Slug:        "escointoken",
			Name:        "EscoinToken",
			BinanceUSDT: "ELGUSDT",
		},
	},
	"LKM": {
		{
			ID:          9410,
			Slug:        "lokum-finance",
			Name:        "Lokum Finance",
			BinanceUSDT: "LKMUSDT",
		},
	},
	"CGT": {
		{
			ID:   5719,
			Slug: "cache-gold",
			Name: "CACHE Gold",
		},
		{
			ID:   8131,
			Slug: "curio-governance",
			Name: "Curio Governance",
		},
	},
	"VRO": {
		{
			ID:          5539,
			Slug:        "veraone",
			Name:        "VeraOne",
			BinanceUSDT: "VROUSDT",
		},
	},
	"IX": {
		{
			ID:          6326,
			Slug:        "x-block",
			Name:        "X-Block",
			BinanceUSDT: "IXUSDT",
		},
	},
	"BAAS": {
		{
			ID:          3142,
			Slug:        "baasid",
			Name:        "BaaSid",
			BinanceUSDT: "BAASUSDT",
		},
	},
	"MSHLD": {
		{
			ID:          9746,
			Slug:        "moonshield",
			Name:        "Moonshield",
			BinanceUSDT: "MSHLDUSDT",
		},
	},
	"MOX": {
		{
			ID:          3688,
			Slug:        "mox",
			Name:        "MoX",
			BinanceUSDT: "MOXUSDT",
		},
	},
	"IMGC": {
		{
			ID:          4344,
			Slug:        "imagecash",
			Name:        "ImageCash",
			BinanceUSDT: "IMGCUSDT",
		},
	},
	"SPANK": {
		{
			ID:          2219,
			Slug:        "spankchain",
			Name:        "SpankChain",
			BinanceUSDT: "SPANKUSDT",
		},
	},
	"HOP": {
		{
			ID:          10050,
			Slug:        "hoppy",
			Name:        "HOPPY",
			BinanceUSDT: "HOPUSDT",
		},
	},
	"VTUBE": {
		{
			ID:          9745,
			Slug:        "vtube-token",
			Name:        "VTube Token",
			BinanceUSDT: "VTUBEUSDT",
		},
	},
	"ILV": {
		{
			ID:          8719,
			Slug:        "illuvium",
			Name:        "Illuvium",
			BinanceUSDT: "ILVUSDT",
		},
	},
	"KUN": {
		{
			ID:   3444,
			Slug: "kun",
			Name: "KUN",
		},
		{
			ID:   7721,
			Slug: "qian-kun",
			Name: "KUN",
		},
	},
	"ATUSD": {
		{
			ID:          5749,
			Slug:        "aave-tusd",
			Name:        "Aave TUSD",
			BinanceUSDT: "ATUSDUSDT",
		},
	},
	"HANA": {
		{
			ID:          5099,
			Slug:        "hanacoin",
			Name:        "Hanacoin",
			BinanceUSDT: "HANAUSDT",
		},
	},
	"LYNC": {
		{
			ID:          7329,
			Slug:        "lync-network",
			Name:        "LYNC Network",
			BinanceUSDT: "LYNCUSDT",
		},
	},
	"TRYON": {
		{
			ID:          10341,
			Slug:        "stellar-invictus-gaming",
			Name:        "Stellar Invictus Gaming",
			BinanceUSDT: "TRYONUSDT",
		},
	},
	"FVT": {
		{
			ID:          8149,
			Slug:        "finance-vote",
			Name:        "Finance.Vote",
			BinanceUSDT: "FVTUSDT",
		},
	},
	"DOGEFATHER": {
		{
			ID:          9732,
			Slug:        "dogefather",
			Name:        "Dogefather",
			BinanceUSDT: "DOGEFATHERUSDT",
		},
	},
	"DGC": {
		{
			ID:          18,
			Slug:        "digitalcoin",
			Name:        "Digitalcoin",
			BinanceUSDT: "DGCUSDT",
		},
	},
	"DDS": {
		{
			ID:          9032,
			Slug:        "dds-store",
			Name:        "DDS.Store",
			BinanceUSDT: "DDSUSDT",
		},
	},
	"TIKI": {
		{
			ID:          10684,
			Slug:        "tiki-token",
			Name:        "Tiki Token",
			BinanceUSDT: "TIKIUSDT",
		},
	},
	"EJECT": {
		{
			ID:          10417,
			Slug:        "eject",
			Name:        "Eject",
			BinanceUSDT: "EJECTUSDT",
		},
	},
	"DVC": {
		{
			ID:          5902,
			Slug:        "dragonvein",
			Name:        "DragonVein",
			BinanceUSDT: "DVCUSDT",
		},
	},
	"KOKOMO": {
		{
			ID:          10709,
			Slug:        "kokomoswap",
			Name:        "KokomoSwap",
			BinanceUSDT: "KOKOMOUSDT",
		},
	},
	"MIC": {
		{
			ID:          8137,
			Slug:        "mith-cash",
			Name:        "MITH Cash",
			BinanceUSDT: "MICUSDT",
		},
	},
	"MANY": {
		{
			ID:          9124,
			Slug:        "manyswap",
			Name:        "Manyswap",
			BinanceUSDT: "MANYUSDT",
		},
	},
	"XXT": {
		{
			ID:          10965,
			Slug:        "xxt-token",
			Name:        "XXT-Token",
			BinanceUSDT: "XXTUSDT",
		},
	},
	"URG": {
		{
			ID:          10719,
			Slug:        "urgaming",
			Name:        "UrGaming",
			BinanceUSDT: "URGUSDT",
		},
	},
	"REFI": {
		{
			ID:          9065,
			Slug:        "realfinance-network",
			Name:        "Realfinance Network",
			BinanceUSDT: "REFIUSDT",
		},
	},
	"SWFL": {
		{
			ID:          6799,
			Slug:        "swapfolio",
			Name:        "Swapfolio",
			BinanceUSDT: "SWFLUSDT",
		},
	},
	"BEE": {
		{
			ID:   8529,
			Slug: "beeswap",
			Name: "BeeSwap",
		},
		{
			ID:   5550,
			Slug: "beeex",
			Name: "BeeEx",
		},
		{
			ID:   10466,
			Slug: "bees",
			Name: "Bees",
		},
	},
	"XTZUP": {
		{
			ID:          7007,
			Slug:        "xtzup",
			Name:        "XTZUP",
			BinanceUSDT: "XTZUPUSDT",
		},
	},
	"APIX": {
		{
			ID:          5258,
			Slug:        "apix",
			Name:        "APIX",
			BinanceUSDT: "APIXUSDT",
		},
	},
	"PoSH": {
		{
			ID:          7112,
			Slug:        "shill-win",
			Name:        "Shill & Win",
			BinanceUSDT: "PoSHUSDT",
		},
	},
	"YCC": {
		{
			ID:          3060,
			Slug:        "yuan-chain-coin",
			Name:        "Yuan Chain Coin",
			BinanceUSDT: "YCCUSDT",
		},
	},
	"BABYCAKE": {
		{
			ID:          10971,
			Slug:        "baby-cake",
			Name:        "Baby Cake",
			BinanceUSDT: "BABYCAKEUSDT",
		},
	},
	"wSIENNA": {
		{
			ID:          9449,
			Slug:        "sienna-erc20",
			Name:        "Sienna (ERC20)",
			BinanceUSDT: "wSIENNAUSDT",
		},
	},
	"PTN": {
		{
			ID:          3595,
			Slug:        "palletone",
			Name:        "PalletOne",
			BinanceUSDT: "PTNUSDT",
		},
	},
	"TPAD": {
		{
			ID:          9613,
			Slug:        "trustpad",
			Name:        "Trustpad",
			BinanceUSDT: "TPADUSDT",
		},
	},
	"HFI": {
		{
			ID:   8540,
			Slug: "hecofi",
			Name: "HecoFi",
		},
		{
			ID:   10355,
			Slug: "holder-finance",
			Name: "Holder Finance",
		},
	},
	"GPYX": {
		{
			ID:          3349,
			Slug:        "goldenpyrex",
			Name:        "GoldenPyrex",
			BinanceUSDT: "GPYXUSDT",
		},
	},
	"DINA": {
		{
			ID:          10413,
			Slug:        "dina",
			Name:        "Dina",
			BinanceUSDT: "DINAUSDT",
		},
	},
	"EDRC": {
		{
			ID:          1216,
			Slug:        "edrcoin",
			Name:        "EDRCoin",
			BinanceUSDT: "EDRCUSDT",
		},
	},
	"WSG": {
		{
			ID:          10040,
			Slug:        "wall-street-games",
			Name:        "Wall Street Games",
			BinanceUSDT: "WSGUSDT",
		},
	},
	"DMD": {
		{
			ID:   77,
			Slug: "diamond",
			Name: "Diamond",
		},
		{
			ID:   6963,
			Slug: "dmd",
			Name: "DMD",
		},
	},
	"D100": {
		{
			ID:          8551,
			Slug:        "defi100",
			Name:        "DeFi100",
			BinanceUSDT: "D100USDT",
		},
	},
	"CT": {
		{
			ID:          8779,
			Slug:        "communitytoken",
			Name:        "CommunityToken",
			BinanceUSDT: "CTUSDT",
		},
	},
	"BURN": {
		{
			ID:          4069,
			Slug:        "blockburn",
			Name:        "Blockburn",
			BinanceUSDT: "BURNUSDT",
		},
	},
	"CASH": {
		{
			ID:          5038,
			Slug:        "litecash",
			Name:        "Litecash",
			BinanceUSDT: "CASHUSDT",
		},
	},
	"TCGCOIN": {
		{
			ID:          9971,
			Slug:        "tcgcoin",
			Name:        "TCGcoin",
			BinanceUSDT: "TCGCOINUSDT",
		},
	},
	"CVN": {
		{
			ID:          1810,
			Slug:        "cvcoin",
			Name:        "CVCoin",
			BinanceUSDT: "CVNUSDT",
		},
	},
	"RM": {
		{
			ID:          5814,
			Slug:        "rivermount",
			Name:        "Rivermount",
			BinanceUSDT: "RMUSDT",
		},
	},
	"HNB": {
		{
			ID:          3947,
			Slug:        "hashnet-biteco",
			Name:        "HashNet BitEco",
			BinanceUSDT: "HNBUSDT",
		},
	},
	"ANON": {
		{
			ID:   3343,
			Slug: "anon",
			Name: "ANON",
		},
		{
			ID:   10162,
			Slug: "anonymousbsc",
			Name: "AnonymousBSC",
		},
	},
	"CRDT": {
		{
			ID:          5473,
			Slug:        "crdt",
			Name:        "CRDT",
			BinanceUSDT: "CRDTUSDT",
		},
	},
	"ARTEON": {
		{
			ID:          9591,
			Slug:        "arteon",
			Name:        "Arteon",
			BinanceUSDT: "ARTEONUSDT",
		},
	},
	"UNL": {
		{
			ID:          8103,
			Slug:        "unilock-network",
			Name:        "unilock.network",
			BinanceUSDT: "UNLUSDT",
		},
	},
	"AMO": {
		{
			ID:          3260,
			Slug:        "amo-coin",
			Name:        "AMO Coin",
			BinanceUSDT: "AMOUSDT",
		},
	},
	"MIKS": {
		{
			ID:          6675,
			Slug:        "miks-coin",
			Name:        "MIKS COIN",
			BinanceUSDT: "MIKSUSDT",
		},
	},
	"DAILYS": {
		{
			ID:          8840,
			Slug:        "dailyswap-token",
			Name:        "DailySwap Token",
			BinanceUSDT: "DAILYSUSDT",
		},
	},
	"KSEED": {
		{
			ID:          7135,
			Slug:        "kush-finance",
			Name:        "Kush Finance",
			BinanceUSDT: "KSEEDUSDT",
		},
	},
	"DTEP": {
		{
			ID:          4277,
			Slug:        "decoin",
			Name:        "DECOIN",
			BinanceUSDT: "DTEPUSDT",
		},
	},
	"SOJU": {
		{
			ID:          8781,
			Slug:        "soju-finance",
			Name:        "Soju Finance",
			BinanceUSDT: "SOJUUSDT",
		},
	},
	"SLRS": {
		{
			ID:          9989,
			Slug:        "solrise-finance",
			Name:        "Solrise Finance",
			BinanceUSDT: "SLRSUSDT",
		},
	},
	"PAR": {
		{
			ID:   4051,
			Slug: "parachute",
			Name: "Parachute",
		},
		{
			ID:   8665,
			Slug: "parallel",
			Name: "Parallel",
		},
	},
	"CHESS": {
		{
			ID:   1297,
			Slug: "chesscoin",
			Name: "ChessCoin",
		},
		{
			ID:   10974,
			Slug: "tranchess",
			Name: "Tranchess",
		},
	},
	"YSEC": {
		{
			ID:          7572,
			Slug:        "yearn-secure",
			Name:        "Yearn Secure",
			BinanceUSDT: "YSECUSDT",
		},
	},
	"PQT": {
		{
			ID:          8804,
			Slug:        "prediqt",
			Name:        "PREDIQT",
			BinanceUSDT: "PQTUSDT",
		},
	},
	"TRIBE": {
		{
			ID:          9025,
			Slug:        "tribe",
			Name:        "Tribe",
			BinanceUSDT: "TRIBEUSDT",
		},
	},
	"FRIES": {
		{
			ID:          7351,
			Slug:        "fryworld",
			Name:        "fry.world",
			BinanceUSDT: "FRIESUSDT",
		},
	},
	"arNXM": {
		{
			ID:          8328,
			Slug:        "armor-nxm",
			Name:        "Armor NXM",
			BinanceUSDT: "arNXMUSDT",
		},
	},
	"OCTO": {
		{
			ID:          7202,
			Slug:        "octofi",
			Name:        "OctoFi",
			BinanceUSDT: "OCTOUSDT",
		},
	},
	"VSL": {
		{
			ID:          1483,
			Slug:        "vslice",
			Name:        "vSlice",
			BinanceUSDT: "VSLUSDT",
		},
	},
	"DAY": {
		{
			ID:          1985,
			Slug:        "chronologic",
			Name:        "Chronologic",
			BinanceUSDT: "DAYUSDT",
		},
	},
	"SUMO": {
		{
			ID:          1694,
			Slug:        "sumokoin",
			Name:        "Sumokoin",
			BinanceUSDT: "SUMOUSDT",
		},
	},
	"BITTO": {
		{
			ID:          4534,
			Slug:        "bitto",
			Name:        "BITTO",
			BinanceUSDT: "BITTOUSDT",
		},
	},
	"HBT": {
		{
			ID:   2031,
			Slug: "hubii-network",
			Name: "Hubii Network",
		},
		{
			ID:   8946,
			Slug: "habitat",
			Name: "Habitat",
		},
	},
	"MCB": {
		{
			ID:          5956,
			Slug:        "mcdex",
			Name:        "MCDEX",
			BinanceUSDT: "MCBUSDT",
		},
	},
	"VANCAT": {
		{
			ID:          8955,
			Slug:        "vancat",
			Name:        "Vancat",
			BinanceUSDT: "VANCATUSDT",
		},
	},
	"XGS": {
		{
			ID:          3377,
			Slug:        "genesisx",
			Name:        "GenesisX",
			BinanceUSDT: "XGSUSDT",
		},
	},
	"PNY": {
		{
			ID:          3481,
			Slug:        "peony",
			Name:        "Peony",
			BinanceUSDT: "PNYUSDT",
		},
	},
	"BUTTER": {
		{
			ID:   9725,
			Slug: "butter-token",
			Name: "Butter TOken",
		},
		{
			ID:   10531,
			Slug: "butterswap",
			Name: "ButterSwap",
		},
	},
	"RAP": {
		{
			ID:          7832,
			Slug:        "realpay",
			Name:        "REALPAY",
			BinanceUSDT: "RAPUSDT",
		},
	},
	"WBNB": {
		{
			ID:          7192,
			Slug:        "wbnb",
			Name:        "Wrapped BNB",
			BinanceUSDT: "WBNBUSDT",
		},
	},
	"POLK": {
		{
			ID:          8579,
			Slug:        "polkamarkets",
			Name:        "Polkamarkets",
			BinanceUSDT: "POLKUSDT",
		},
	},
	"BXK": {
		{
			ID:          4049,
			Slug:        "bitbook-gambling",
			Name:        "Bitbook Gambling",
			BinanceUSDT: "BXKUSDT",
		},
	},
	"MAYP": {
		{
			ID:          4077,
			Slug:        "maya-preferred",
			Name:        "Maya Preferred",
			BinanceUSDT: "MAYPUSDT",
		},
	},
	"NU": {
		{
			ID:          4761,
			Slug:        "nucypher",
			Name:        "NuCypher",
			BinanceUSDT: "NUUSDT",
		},
	},
	"CLOAK": {
		{
			ID:          362,
			Slug:        "cloakcoin",
			Name:        "CloakCoin",
			BinanceUSDT: "CLOAKUSDT",
		},
	},
	"FWATCH": {
		{
			ID:          10151,
			Slug:        "foliowatch",
			Name:        "Foliowatch",
			BinanceUSDT: "FWATCHUSDT",
		},
	},
	"GXT": {
		{
			ID:   7310,
			Slug: "gem-exchange-and-trading",
			Name: "Gem Exchange And Trading",
		},
		{
			ID:   4693,
			Slug: "global-x-change-token",
			Name: "Global X Change Token",
		},
	},
	"vDOT": {
		{
			ID:          7976,
			Slug:        "venus-dot",
			Name:        "Venus DOT",
			BinanceUSDT: "vDOTUSDT",
		},
	},
	"GHD": {
		{
			ID:          7126,
			Slug:        "giftedhands",
			Name:        "Giftedhands",
			BinanceUSDT: "GHDUSDT",
		},
	},
	"ACSI": {
		{
			ID:          10797,
			Slug:        "acryptosi",
			Name:        "ACryptoSI",
			BinanceUSDT: "ACSIUSDT",
		},
	},
	"UTL": {
		{
			ID:          8728,
			Slug:        "utile-network",
			Name:        "Utile Network",
			BinanceUSDT: "UTLUSDT",
		},
	},
	"BORA": {
		{
			ID:          3801,
			Slug:        "bora",
			Name:        "BORA",
			BinanceUSDT: "BORAUSDT",
		},
	},
	"MTX": {
		{
			ID:          2325,
			Slug:        "matryx",
			Name:        "Matryx",
			BinanceUSDT: "MTXUSDT",
		},
	},
	"SHB": {
		{
			ID:          3604,
			Slug:        "skyhub-coin",
			Name:        "SkyHub Coin",
			BinanceUSDT: "SHBUSDT",
		},
	},
	"yVault LP-yCurve(YYCRV)": {
		{
			ID:          6686,
			Slug:        "yvault-lp-ycurve",
			Name:        "yVault LP-yCurve",
			BinanceUSDT: "yVault LP-yCurve(YYCRV)USDT",
		},
	},
	"ZDR": {
		{
			ID:          3899,
			Slug:        "zloadr",
			Name:        "Zloadr",
			BinanceUSDT: "ZDRUSDT",
		},
	},
	"GRFT": {
		{
			ID:          2571,
			Slug:        "graft",
			Name:        "Graft",
			BinanceUSDT: "GRFTUSDT",
		},
	},
	"COLD": {
		{
			ID:          9751,
			Slug:        "cold-finance",
			Name:        "COLD FINANCE",
			BinanceUSDT: "COLDUSDT",
		},
	},
	"LANC": {
		{
			ID:          10666,
			Slug:        "lanceria",
			Name:        "Lanceria",
			BinanceUSDT: "LANCUSDT",
		},
	},
	"ADL": {
		{
			ID:          1725,
			Slug:        "adelphoi",
			Name:        "Adelphoi",
			BinanceUSDT: "ADLUSDT",
		},
	},
	"ZEON": {
		{
			ID:          3795,
			Slug:        "zeon",
			Name:        "ZEON",
			BinanceUSDT: "ZEONUSDT",
		},
	},
	"BOUTS": {
		{
			ID:          2717,
			Slug:        "boutspro",
			Name:        "BoutsPro",
			BinanceUSDT: "BOUTSUSDT",
		},
	},
	"QSP": {
		{
			ID:          2212,
			Slug:        "quantstamp",
			Name:        "Quantstamp",
			BinanceUSDT: "QSPUSDT",
		},
	},
	"MTL": {
		{
			ID:          1788,
			Slug:        "metal",
			Name:        "Metal",
			BinanceUSDT: "MTLUSDT",
		},
	},
	"KEANU": {
		{
			ID:          10196,
			Slug:        "keanu-inu",
			Name:        "Keanu Inu",
			BinanceUSDT: "KEANUUSDT",
		},
	},
	"DT": {
		{
			ID:   6416,
			Slug: "dcoin-token",
			Name: "Dcoin Token",
		},
		{
			ID:   3470,
			Slug: "dragon-token",
			Name: "Dragon Token",
		},
	},
	"ALG": {
		{
			ID:   5508,
			Slug: "algory-project",
			Name: "Algory Project",
		},
		{
			ID:   6600,
			Slug: "bitalgo",
			Name: "Bitalgo",
		},
	},
	"LBTC": {
		{
			ID:          2335,
			Slug:        "lightning-bitcoin",
			Name:        "Lightning Bitcoin",
			BinanceUSDT: "LBTCUSDT",
		},
	},
	"MXX": {
		{
			ID:          6583,
			Slug:        "multiplier",
			Name:        "Multiplier",
			BinanceUSDT: "MXXUSDT",
		},
	},
	"RAMEN": {
		{
			ID:          8547,
			Slug:        "ramenswap",
			Name:        "RamenSwap",
			BinanceUSDT: "RAMENUSDT",
		},
	},
	"ATYNE": {
		{
			ID:          10100,
			Slug:        "aerotyne",
			Name:        "Aerotyne",
			BinanceUSDT: "ATYNEUSDT",
		},
	},
	"FISH": {
		{
			ID:          10134,
			Slug:        "polycat-finance",
			Name:        "Polycat Finance",
			BinanceUSDT: "FISHUSDT",
		},
	},
	"DATA": {
		{
			ID:          2143,
			Slug:        "streamr",
			Name:        "Streamr",
			BinanceUSDT: "DATAUSDT",
		},
	},
	"mAMZN": {
		{
			ID:          8016,
			Slug:        "mirrored-amazon",
			Name:        "Mirrored Amazon",
			BinanceUSDT: "mAMZNUSDT",
		},
	},
	"NEXT": {
		{
			ID:          3651,
			Slug:        "next-coin",
			Name:        "NEXT",
			BinanceUSDT: "NEXTUSDT",
		},
	},
	"SAV3": {
		{
			ID:          7680,
			Slug:        "sav3token",
			Name:        "Sav3Token",
			BinanceUSDT: "SAV3USDT",
		},
	},
	"ATD": {
		{
			ID:          8926,
			Slug:        "a2dao",
			Name:        "A2DAO",
			BinanceUSDT: "ATDUSDT",
		},
	},
	"ADM": {
		{
			ID:          3703,
			Slug:        "adamant-messenger",
			Name:        "ADAMANT Messenger",
			BinanceUSDT: "ADMUSDT",
		},
	},
	"HTB": {
		{
			ID:          4506,
			Slug:        "hotbit-token",
			Name:        "Hotbit Token",
			BinanceUSDT: "HTBUSDT",
		},
	},
	"DNS": {
		{
			ID:          8194,
			Slug:        "bitdns",
			Name:        "BitDNS",
			BinanceUSDT: "DNSUSDT",
		},
	},
	"AG8": {
		{
			ID:          5536,
			Slug:        "atromg8",
			Name:        "AtromG8",
			BinanceUSDT: "AG8USDT",
		},
	},
	"LUCK": {
		{
			ID:          9222,
			Slug:        "lucktogether",
			Name:        "LuckTogether",
			BinanceUSDT: "LUCKUSDT",
		},
	},
	"TTC": {
		{
			ID:          8359,
			Slug:        "ttcrypto",
			Name:        "TTCRYPTO",
			BinanceUSDT: "TTCUSDT",
		},
	},
	"INXT": {
		{
			ID:          2022,
			Slug:        "internxt",
			Name:        "Internxt",
			BinanceUSDT: "INXTUSDT",
		},
	},
	"IMX": {
		{
			ID:          9532,
			Slug:        "impermax",
			Name:        "Impermax",
			BinanceUSDT: "IMXUSDT",
		},
	},
	"SHENG": {
		{
			ID:          6256,
			Slug:        "sheng",
			Name:        "SHENG",
			BinanceUSDT: "SHENGUSDT",
		},
	},
	"WCFG": {
		{
			ID:          10898,
			Slug:        "wrapped-centrifuge",
			Name:        "Wrapped Centrifuge",
			BinanceUSDT: "WCFGUSDT",
		},
	},
	"FORS": {
		{
			ID:          6918,
			Slug:        "foresight",
			Name:        "Foresight",
			BinanceUSDT: "FORSUSDT",
		},
	},
	"LDO": {
		{
			ID:          8000,
			Slug:        "lido-dao",
			Name:        "Lido DAO Token",
			BinanceUSDT: "LDOUSDT",
		},
	},
	"EGR": {
		{
			ID:          5075,
			Slug:        "egoras",
			Name:        "Egoras",
			BinanceUSDT: "EGRUSDT",
		},
	},
	"XCHF": {
		{
			ID:          4075,
			Slug:        "cryptofranc",
			Name:        "CryptoFranc",
			BinanceUSDT: "XCHFUSDT",
		},
	},
	"KELPIE": {
		{
			ID:          10736,
			Slug:        "kelpie-inu",
			Name:        "Kelpie Inu",
			BinanceUSDT: "KELPIEUSDT",
		},
	},
	"SOV": {
		{
			ID:          8669,
			Slug:        "sovryn",
			Name:        "Sovryn",
			BinanceUSDT: "SOVUSDT",
		},
	},
	"PLTC": {
		{
			ID:          3753,
			Slug:        "platoncoin",
			Name:        "PlatonCoin",
			BinanceUSDT: "PLTCUSDT",
		},
	},
	"FRENS": {
		{
			ID:          6551,
			Slug:        "frens-community",
			Name:        "Frens Community",
			BinanceUSDT: "FRENSUSDT",
		},
	},
	"ORN": {
		{
			ID:          5631,
			Slug:        "orion-protocol",
			Name:        "Orion Protocol",
			BinanceUSDT: "ORNUSDT",
		},
	},
	"ILG": {
		{
			ID:          9918,
			Slug:        "ilgon",
			Name:        "ILGON",
			BinanceUSDT: "ILGUSDT",
		},
	},
	"NAFTY": {
		{
			ID:          10864,
			Slug:        "nafty",
			Name:        "NAFTY",
			BinanceUSDT: "NAFTYUSDT",
		},
	},
	"ABET": {
		{
			ID:          4502,
			Slug:        "altbet",
			Name:        "Altbet",
			BinanceUSDT: "ABETUSDT",
		},
	},
	"NCAT": {
		{
			ID:          8959,
			Slug:        "ncat-token",
			Name:        "NCAT Token",
			BinanceUSDT: "NCATUSDT",
		},
	},
	"QUAI": {
		{
			ID:          8373,
			Slug:        "quai-dao",
			Name:        "QUAI DAO",
			BinanceUSDT: "QUAIUSDT",
		},
	},
	"XRPDOWN": {
		{
			ID:          7002,
			Slug:        "xrpdown",
			Name:        "XRPDOWN",
			BinanceUSDT: "XRPDOWNUSDT",
		},
	},
	"AGVE": {
		{
			ID:          9453,
			Slug:        "agave",
			Name:        "Agave",
			BinanceUSDT: "AGVEUSDT",
		},
	},
	"UTT": {
		{
			ID:          2371,
			Slug:        "uttoken",
			Name:        "United Traders Token",
			BinanceUSDT: "UTTUSDT",
		},
	},
	"BB": {
		{
			ID:          8345,
			Slug:        "blackberry-tokenized-stock-ftx",
			Name:        "BlackBerry tokenized stock FTX",
			BinanceUSDT: "BBUSDT",
		},
	},
	"GRIMEX": {
		{
			ID:          10158,
			Slug:        "spacegrime",
			Name:        "SpaceGrime",
			BinanceUSDT: "GRIMEXUSDT",
		},
	},
	"KDAG": {
		{
			ID:          5626,
			Slug:        "king-dag",
			Name:        "King DAG",
			BinanceUSDT: "KDAGUSDT",
		},
	},
	"KLAY": {
		{
			ID:          4256,
			Slug:        "klaytn",
			Name:        "Klaytn",
			BinanceUSDT: "KLAYUSDT",
		},
	},
	"MTC": {
		{
			ID:   2711,
			Slug: "doc-com",
			Name: "DOC.COM",
		},
		{
			ID:   6498,
			Slug: "metacoin",
			Name: "Metacoin",
		},
	},
	"NYB": {
		{
			ID:          6981,
			Slug:        "new-year-bull",
			Name:        "New Year Bull",
			BinanceUSDT: "NYBUSDT",
		},
	},
	"VOCO": {
		{
			ID:          3509,
			Slug:        "provoco-token",
			Name:        "Provoco Token",
			BinanceUSDT: "VOCOUSDT",
		},
	},
	"OSM": {
		{
			ID:          10177,
			Slug:        "supermoon",
			Name:        "Supermoon",
			BinanceUSDT: "OSMUSDT",
		},
	},
	"LSS": {
		{
			ID:          10103,
			Slug:        "lossless",
			Name:        "Lossless",
			BinanceUSDT: "LSSUSDT",
		},
	},
	"PRQBOOST": {
		{
			ID:          8488,
			Slug:        "parsiq-boost",
			Name:        "Parsiq Boost",
			BinanceUSDT: "PRQBOOSTUSDT",
		},
	},
	"TAIYO": {
		{
			ID:          11030,
			Slug:        "taiyo",
			Name:        "TAIYO",
			BinanceUSDT: "TAIYOUSDT",
		},
	},
	"KOIN": {
		{
			ID:          8282,
			Slug:        "koinos",
			Name:        "Koinos",
			BinanceUSDT: "KOINUSDT",
		},
	},
	"FSXU": {
		{
			ID:          8811,
			Slug:        "flashx-ultra",
			Name:        "FlashX Ultra",
			BinanceUSDT: "FSXUUSDT",
		},
	},
	"XYM": {
		{
			ID:          8677,
			Slug:        "symbol",
			Name:        "Symbol",
			BinanceUSDT: "XYMUSDT",
		},
	},
	"QRX": {
		{
			ID:          7510,
			Slug:        "quiverx",
			Name:        "QuiverX",
			BinanceUSDT: "QRXUSDT",
		},
	},
	"XSH": {
		{
			ID:          2144,
			Slug:        "shield-xsh",
			Name:        "SHIELD",
			BinanceUSDT: "XSHUSDT",
		},
	},
	"ACES": {
		{
			ID:          1351,
			Slug:        "aces",
			Name:        "Aces",
			BinanceUSDT: "ACESUSDT",
		},
	},
	"MATH": {
		{
			ID:          5616,
			Slug:        "math",
			Name:        "MATH",
			BinanceUSDT: "MATHUSDT",
		},
	},
	"VAIP": {
		{
			ID:          5619,
			Slug:        "vehicle-data-artificial-intelligence-platform",
			Name:        "VEHICLE DATA ARTIFICIAL INTELLIGENCE PLATFORM",
			BinanceUSDT: "VAIPUSDT",
		},
	},
	"BABYSHIB": {
		{
			ID:          10886,
			Slug:        "babyshibby-inu",
			Name:        "BabyShibby Inu",
			BinanceUSDT: "BABYSHIBUSDT",
		},
	},
	"OIL": {
		{
			ID:   8645,
			Slug: "oiler-network",
			Name: "Oiler Network",
		},
		{
			ID:   8527,
			Slug: "crudeoil-finance",
			Name: "Crudeoil Finance",
		},
	},
	"MTF": {
		{
			ID:          9373,
			Slug:        "milktea-finance",
			Name:        "Milktea.finance",
			BinanceUSDT: "MTFUSDT",
		},
	},
	"CAT": {
		{
			ID:          6649,
			Slug:        "cat-token",
			Name:        "Cat Token",
			BinanceUSDT: "CATUSDT",
		},
	},
	"BSD": {
		{
			ID:   366,
			Slug: "bitsend",
			Name: "BitSend",
		},
		{
			ID:   8074,
			Slug: "basis-dollar",
			Name: "Basis Dollar",
		},
	},
	"DXD": {
		{
			ID:          5589,
			Slug:        "dxdao",
			Name:        "DXdao",
			BinanceUSDT: "DXDUSDT",
		},
	},
	"CHOW": {
		{
			ID:          8433,
			Slug:        "chow-chow",
			Name:        "Chow Chow",
			BinanceUSDT: "CHOWUSDT",
		},
	},
	"GIVE": {
		{
			ID:   10838,
			Slug: "give-token",
			Name: "GIVE Token",
		},
		{
			ID:   10253,
			Slug: "give-global",
			Name: "GIVE GLOBAL",
		},
	},
	"XIV": {
		{
			ID:          8746,
			Slug:        "project-inverse",
			Name:        "Project Inverse",
			BinanceUSDT: "XIVUSDT",
		},
	},
	"ZKS": {
		{
			ID:          8202,
			Slug:        "zkswap",
			Name:        "ZKSwap",
			BinanceUSDT: "ZKSUSDT",
		},
	},
	"GOLD": {
		{
			ID:   8281,
			Slug: "golden-goose",
			Name: "Golden Goose",
		},
		{
			ID:   4114,
			Slug: "golden-token",
			Name: "Golden Token",
		},
		{
			ID:   10642,
			Slug: "goldfarm",
			Name: "GoldFarm",
		},
	},
	"KONG": {
		{
			ID:          8998,
			Slug:        "kong-defi",
			Name:        "Kong Defi",
			BinanceUSDT: "KONGUSDT",
		},
	},
	"LEC": {
		{
			ID:          10779,
			Slug:        "love-earth-coin",
			Name:        "LOVE EARTH COIN",
			BinanceUSDT: "LECUSDT",
		},
	},
	"UBX": {
		{
			ID:          7622,
			Slug:        "ubix-network",
			Name:        "UBIX.Network",
			BinanceUSDT: "UBXUSDT",
		},
	},
	"SXP": {
		{
			ID:          4279,
			Slug:        "swipe",
			Name:        "Swipe",
			BinanceUSDT: "SXPUSDT",
		},
	},
	"PUT": {
		{
			ID:   1299,
			Slug: "putincoin",
			Name: "PutinCoin",
		},
		{
			ID:   2419,
			Slug: "profile-utility-token",
			Name: "Profile Utility Token",
		},
	},
	"XWC": {
		{
			ID:          268,
			Slug:        "whitecoin",
			Name:        "WhiteCoin",
			BinanceUSDT: "XWCUSDT",
		},
	},
	"EQZ": {
		{
			ID:          9083,
			Slug:        "equalizer",
			Name:        "Equalizer",
			BinanceUSDT: "EQZUSDT",
		},
	},
	"GNC": {
		{
			ID:          7447,
			Slug:        "galaxy-network",
			Name:        "GALAXY NETWORK",
			BinanceUSDT: "GNCUSDT",
		},
	},
	"ECP": {
		{
			ID:          9378,
			Slug:        "eclipse-ecp",
			Name:        "Eclipse",
			BinanceUSDT: "ECPUSDT",
		},
	},
	"YLC": {
		{
			ID:          3162,
			Slug:        "yolocash",
			Name:        "YoloCash",
			BinanceUSDT: "YLCUSDT",
		},
	},
	"VNXLU": {
		{
			ID:          4430,
			Slug:        "vnx-exchange",
			Name:        "VNX Exchange",
			BinanceUSDT: "VNXLUUSDT",
		},
	},
	"NBTC": {
		{
			ID:          5888,
			Slug:        "neobitcoin",
			Name:        "NEOBITCOIN",
			BinanceUSDT: "NBTCUSDT",
		},
	},
	"DRUNK": {
		{
			ID:          10964,
			Slug:        "drunkdoge",
			Name:        "DrunkDoge",
			BinanceUSDT: "DRUNKUSDT",
		},
	},
	"SAFEICARUS": {
		{
			ID:          9630,
			Slug:        "safeicarus",
			Name:        "Safeicarus",
			BinanceUSDT: "SAFEICARUSUSDT",
		},
	},
	"GAS": {
		{
			ID:          1785,
			Slug:        "gas",
			Name:        "Gas",
			BinanceUSDT: "GASUSDT",
		},
	},
	"VKF": {
		{
			ID:          7644,
			Slug:        "vkf-platform",
			Name:        "VKF Platform",
			BinanceUSDT: "VKFUSDT",
		},
	},
	"CREDIT": {
		{
			ID:   4165,
			Slug: "credit",
			Name: "TerraCredit",
		},
		{
			ID:   6668,
			Slug: "proxi",
			Name: "PROXI",
		},
	},
	"CLT": {
		{
			ID:          5401,
			Slug:        "coinloan",
			Name:        "CoinLoan",
			BinanceUSDT: "CLTUSDT",
		},
	},
	"X-TOKEN": {
		{
			ID:          9927,
			Slug:        "x-token",
			Name:        "X-Token",
			BinanceUSDT: "X-TOKENUSDT",
		},
	},
	"UMA": {
		{
			ID:          5617,
			Slug:        "uma",
			Name:        "UMA",
			BinanceUSDT: "UMAUSDT",
		},
	},
	"RSR": {
		{
			ID:          3964,
			Slug:        "reserve-rights",
			Name:        "Reserve Rights",
			BinanceUSDT: "RSRUSDT",
		},
	},
	"STARLINKDOGE": {
		{
			ID:          10871,
			Slug:        "baby-starlink-doge",
			Name:        "Baby Starlink Doge",
			BinanceUSDT: "STARLINKDOGEUSDT",
		},
	},
	"NUTS": {
		{
			ID:          6986,
			Slug:        "squirrel-finance",
			Name:        "Squirrel Finance",
			BinanceUSDT: "NUTSUSDT",
		},
	},
	"TT": {
		{
			ID:          3930,
			Slug:        "thunder-token",
			Name:        "Thunder Token",
			BinanceUSDT: "TTUSDT",
		},
	},
	"COIL": {
		{
			ID:          6645,
			Slug:        "coil",
			Name:        "COIL",
			BinanceUSDT: "COILUSDT",
		},
	},
	"TAG": {
		{
			ID:          61,
			Slug:        "tagcoin",
			Name:        "TagCoin",
			BinanceUSDT: "TAGUSDT",
		},
	},
	"BLOSM": {
		{
			ID:          9733,
			Slug:        "blossomcoin",
			Name:        "BlossomCoin",
			BinanceUSDT: "BLOSMUSDT",
		},
	},
	"CWV": {
		{
			ID:          3609,
			Slug:        "cwv-chain",
			Name:        "CWV Chain",
			BinanceUSDT: "CWVUSDT",
		},
	},
	"BUP": {
		{
			ID:          7761,
			Slug:        "buildup",
			Name:        "BuildUp",
			BinanceUSDT: "BUPUSDT",
		},
	},
	"BMP": {
		{
			ID:          7732,
			Slug:        "brother-music-platform",
			Name:        "Brother Music Platform",
			BinanceUSDT: "BMPUSDT",
		},
	},
	"YNK": {
		{
			ID:          7252,
			Slug:        "yoink",
			Name:        "Yoink",
			BinanceUSDT: "YNKUSDT",
		},
	},
	"BTCDOM": {
		{
			ID:          7421,
			Slug:        "bitfinex-bitcoin-dominance-perps",
			Name:        "Bitfinex Bitcoin Dominance Perps",
			BinanceUSDT: "BTCDOMUSDT",
		},
	},
	"FNSP": {
		{
			ID:          7134,
			Slug:        "finswap",
			Name:        "Finswap",
			BinanceUSDT: "FNSPUSDT",
		},
	},
	"2CRZ": {
		{
			ID:          10418,
			Slug:        "2crazynft",
			Name:        "2crazyNFT",
			BinanceUSDT: "2CRZUSDT",
		},
	},
	"RUSH": {
		{
			ID:   9920,
			Slug: "rush-coin",
			Name: "RUSH COIN",
		},
		{
			ID:   10126,
			Slug: "rushmoon",
			Name: "RushMoon",
		},
	},
	"BSW": {
		{
			ID:          10746,
			Slug:        "biswap",
			Name:        "Biswap",
			BinanceUSDT: "BSWUSDT",
		},
	},
	"KDG": {
		{
			ID:          5407,
			Slug:        "kingdom-game-4",
			Name:        "Kingdom Game 4.0",
			BinanceUSDT: "KDGUSDT",
		},
	},
	"VBIT": {
		{
			ID:          6947,
			Slug:        "valobit",
			Name:        "Valobit",
			BinanceUSDT: "VBITUSDT",
		},
	},
	"HVN": {
		{
			ID:          1950,
			Slug:        "hiveterminal-token",
			Name:        "Hiveterminal Token",
			BinanceUSDT: "HVNUSDT",
		},
	},
	"HYPE": {
		{
			ID:          8130,
			Slug:        "hype",
			Name:        "Supreme Finance",
			BinanceUSDT: "HYPEUSDT",
		},
	},
	"TUDA": {
		{
			ID:          4301,
			Slug:        "tutors-diary",
			Name:        "Tutor's Diary",
			BinanceUSDT: "TUDAUSDT",
		},
	},
	"ERN": {
		{
			ID:          8615,
			Slug:        "ethernity-chain",
			Name:        "Ethernity Chain",
			BinanceUSDT: "ERNUSDT",
		},
	},
	"SND": {
		{
			ID:          3523,
			Slug:        "snodecoin",
			Name:        "SnodeCoin",
			BinanceUSDT: "SNDUSDT",
		},
	},
	"SHIBAPUP": {
		{
			ID:          9708,
			Slug:        "shibapup",
			Name:        "ShibaPup",
			BinanceUSDT: "SHIBAPUPUSDT",
		},
	},
	"NMT": {
		{
			ID:          10033,
			Slug:        "nftmart-token",
			Name:        "NFTMart Token",
			BinanceUSDT: "NMTUSDT",
		},
	},
	"ALI": {
		{
			ID:          3248,
			Slug:        "ailink-token",
			Name:        "AiLink Token",
			BinanceUSDT: "ALIUSDT",
		},
	},
	"DACXI": {
		{
			ID:          10372,
			Slug:        "dacxi",
			Name:        "Dacxi",
			BinanceUSDT: "DACXIUSDT",
		},
	},
	"ODE": {
		{
			ID:   2631,
			Slug: "odem",
			Name: "ODEM",
		},
		{
			ID:   5367,
			Slug: "ode",
			Name: "ODE",
		},
	},
	"VLAD": {
		{
			ID:          9096,
			Slug:        "vlad-finance",
			Name:        "Vlad Finance",
			BinanceUSDT: "VLADUSDT",
		},
	},
	"XDAG": {
		{
			ID:          4424,
			Slug:        "xdag",
			Name:        "XDAG",
			BinanceUSDT: "XDAGUSDT",
		},
	},
	"HOD": {
		{
			ID:          10412,
			Slug:        "hodooi",
			Name:        "HoDooi",
			BinanceUSDT: "HODUSDT",
		},
	},
	"WIS": {
		{
			ID:          7697,
			Slug:        "experty-wisdom-token",
			Name:        "Experty Wisdom Token",
			BinanceUSDT: "WISUSDT",
		},
	},
	"CEDS": {
		{
			ID:          6745,
			Slug:        "cedars",
			Name:        "CEDARS",
			BinanceUSDT: "CEDSUSDT",
		},
	},
	"VET": {
		{
			ID:          3077,
			Slug:        "vechain",
			Name:        "VeChain",
			BinanceUSDT: "VETUSDT",
		},
	},
	"DEM": {
		{
			ID:          72,
			Slug:        "deutsche-emark",
			Name:        "Deutsche eMark",
			BinanceUSDT: "DEMUSDT",
		},
	},
	"IFC": {
		{
			ID:          41,
			Slug:        "infinitecoin",
			Name:        "Infinitecoin",
			BinanceUSDT: "IFCUSDT",
		},
	},
	"pETH18C": {
		{
			ID:          9034,
			Slug:        "peth18c",
			Name:        "pETH18C",
			BinanceUSDT: "pETH18CUSDT",
		},
	},
	"SMI": {
		{
			ID:          9958,
			Slug:        "safemoon-inu",
			Name:        "SafeMoon Inu",
			BinanceUSDT: "SMIUSDT",
		},
	},
	"KICH": {
		{
			ID:          10927,
			Slug:        "kichicoin",
			Name:        "KichiCoin",
			BinanceUSDT: "KICHUSDT",
		},
	},
	"SX": {
		{
			ID:          8377,
			Slug:        "sportx",
			Name:        "SportX",
			BinanceUSDT: "SXUSDT",
		},
	},
	"ZPT": {
		{
			ID:          2481,
			Slug:        "zeepin",
			Name:        "Zeepin",
			BinanceUSDT: "ZPTUSDT",
		},
	},
	"MLK": {
		{
			ID:          5266,
			Slug:        "milk-alliance",
			Name:        "MiL.k",
			BinanceUSDT: "MLKUSDT",
		},
	},
	"BDAY": {
		{
			ID:          8462,
			Slug:        "birthday-cake",
			Name:        "Birthday Cake",
			BinanceUSDT: "BDAYUSDT",
		},
	},
	"SCAT": {
		{
			ID:          8819,
			Slug:        "sad-cat-token",
			Name:        "Sad Cat Token",
			BinanceUSDT: "SCATUSDT",
		},
	},
	"TROY": {
		{
			ID:          5007,
			Slug:        "troy",
			Name:        "TROY",
			BinanceUSDT: "TROYUSDT",
		},
	},
	"DXC": {
		{
			ID:          5364,
			Slug:        "mydexpay",
			Name:        "Dexchain",
			BinanceUSDT: "DXCUSDT",
		},
	},
	"SBD": {
		{
			ID:          1312,
			Slug:        "steem-dollars",
			Name:        "Steem Dollars",
			BinanceUSDT: "SBDUSDT",
		},
	},
	"OFC": {
		{
			ID:          10794,
			Slug:        "ofc-coin",
			Name:        "$OFC Coin",
			BinanceUSDT: "OFCUSDT",
		},
	},
	"ATOM": {
		{
			ID:          3794,
			Slug:        "cosmos",
			Name:        "Cosmos",
			BinanceUSDT: "ATOMUSDT",
		},
	},
	"UT": {
		{
			ID:          3233,
			Slug:        "ulord",
			Name:        "Ulord",
			BinanceUSDT: "UTUSDT",
		},
	},
	"DAOFI": {
		{
			ID:          7419,
			Slug:        "daofi",
			Name:        "DAOFi",
			BinanceUSDT: "DAOFIUSDT",
		},
	},
	"BABYBITC": {
		{
			ID:          11032,
			Slug:        "babybitcoin",
			Name:        "BabyBitcoin",
			BinanceUSDT: "BABYBITCUSDT",
		},
	},
	"WG0": {
		{
			ID:          7472,
			Slug:        "wrapped-gen-0-cryptokitties",
			Name:        "Wrapped Gen-0 CryptoKitties",
			BinanceUSDT: "WG0USDT",
		},
	},
	"SAN": {
		{
			ID:          1807,
			Slug:        "santiment",
			Name:        "Santiment Network Token",
			BinanceUSDT: "SANUSDT",
		},
	},
	"NFY": {
		{
			ID:          7389,
			Slug:        "non-fungible-yearn",
			Name:        "Non-Fungible Yearn",
			BinanceUSDT: "NFYUSDT",
		},
	},
	"ZOOM": {
		{
			ID:          5926,
			Slug:        "coinzoom",
			Name:        "CoinZoom",
			BinanceUSDT: "ZOOMUSDT",
		},
	},
	"YFIX": {
		{
			ID:          7166,
			Slug:        "yfix-finance",
			Name:        "YFIX Finance",
			BinanceUSDT: "YFIXUSDT",
		},
	},
	"EUNO": {
		{
			ID:          3071,
			Slug:        "euno",
			Name:        "EUNO",
			BinanceUSDT: "EUNOUSDT",
		},
	},
	"VANY": {
		{
			ID:          5554,
			Slug:        "vanywhere",
			Name:        "Vanywhere",
			BinanceUSDT: "VANYUSDT",
		},
	},
	"YF-DAI": {
		{
			ID:          6938,
			Slug:        "yfdai-finance",
			Name:        "YFDAI.FINANCE",
			BinanceUSDT: "YF-DAIUSDT",
		},
	},
	"HYMETEOR": {
		{
			ID:          9699,
			Slug:        "hypermeteor",
			Name:        "HyperMeteor",
			BinanceUSDT: "HYMETEORUSDT",
		},
	},
	"MNG": {
		{
			ID:          10456,
			Slug:        "moon-nation-game",
			Name:        "Moon Nation Game",
			BinanceUSDT: "MNGUSDT",
		},
	},
	"ALEPH": {
		{
			ID:          5821,
			Slug:        "aleph-im",
			Name:        "Aleph.im",
			BinanceUSDT: "ALEPHUSDT",
		},
	},
	"KAM": {
		{
			ID:          5107,
			Slug:        "bitkam",
			Name:        "BitKAM",
			BinanceUSDT: "KAMUSDT",
		},
	},
	"FRIDGE": {
		{
			ID:          9042,
			Slug:        "fridge-token",
			Name:        "Fridge Token",
			BinanceUSDT: "FRIDGEUSDT",
		},
	},
	"ACM": {
		{
			ID:   8538,
			Slug: "ac-milan-fan-token",
			Name: "AC Milan Fan Token",
		},
		{
			ID:   3386,
			Slug: "actinium",
			Name: "Actinium",
		},
	},
	"KEX": {
		{
			ID:          6930,
			Slug:        "kira-network",
			Name:        "Kira Network",
			BinanceUSDT: "KEXUSDT",
		},
	},
	"HOGT": {
		{
			ID:          10605,
			Slug:        "hogt",
			Name:        "HOGT",
			BinanceUSDT: "HOGTUSDT",
		},
	},
	"UNIFY": {
		{
			ID:          1736,
			Slug:        "unify",
			Name:        "Unify",
			BinanceUSDT: "UNIFYUSDT",
		},
	},
	"PLBT": {
		{
			ID:          1784,
			Slug:        "polybius",
			Name:        "Polybius",
			BinanceUSDT: "PLBTUSDT",
		},
	},
	"TMC": {
		{
			ID:          7235,
			Slug:        "trading-membership-community",
			Name:        "Trading Membership Community",
			BinanceUSDT: "TMCUSDT",
		},
	},
	"ZNT": {
		{
			ID:          3446,
			Slug:        "zenswap-network-token",
			Name:        "Zenswap Network Token",
			BinanceUSDT: "ZNTUSDT",
		},
	},
	"SS": {
		{
			ID:          2699,
			Slug:        "sharder",
			Name:        "Sharder",
			BinanceUSDT: "SSUSDT",
		},
	},
	"XTZ": {
		{
			ID:          2011,
			Slug:        "tezos",
			Name:        "Tezos",
			BinanceUSDT: "XTZUSDT",
		},
	},
	"STBZ": {
		{
			ID:          7366,
			Slug:        "stabilize",
			Name:        "Stabilize",
			BinanceUSDT: "STBZUSDT",
		},
	},
	"IXT": {
		{
			ID:          1845,
			Slug:        "ixledger",
			Name:        "IXT",
			BinanceUSDT: "IXTUSDT",
		},
	},
	"DAX": {
		{
			ID:          2696,
			Slug:        "daex",
			Name:        "DAEX",
			BinanceUSDT: "DAXUSDT",
		},
	},
	"TFF": {
		{
			ID:          8870,
			Slug:        "tutti-frutti",
			Name:        "Tutti Frutti",
			BinanceUSDT: "TFFUSDT",
		},
	},
	"X": {
		{
			ID:          10213,
			Slug:        "x-by-spacegrime",
			Name:        "X (By SpaceGrime)",
			BinanceUSDT: "XUSDT",
		},
	},
	"PYRK": {
		{
			ID:          5591,
			Slug:        "pyrk",
			Name:        "Pyrk",
			BinanceUSDT: "PYRKUSDT",
		},
	},
	"ALPHA": {
		{
			ID:          7232,
			Slug:        "alpha-finance-lab",
			Name:        "Alpha Finance Lab",
			BinanceUSDT: "ALPHAUSDT",
		},
	},
	"AC3": {
		{
			ID:          2722,
			Slug:        "ac3",
			Name:        "AC3",
			BinanceUSDT: "AC3USDT",
		},
	},
	"STA": {
		{
			ID:   5868,
			Slug: "statera",
			Name: "STATERA",
		},
		{
			ID:   6867,
			Slug: "stable-asset",
			Name: "STABLE ASSET",
		},
	},
	"WVG0": {
		{
			ID:          7471,
			Slug:        "wrapped-virgin-gen-0-cryptokitties",
			Name:        "Wrapped Virgin Gen-0 CryptoKitties",
			BinanceUSDT: "WVG0USDT",
		},
	},
	"SEM": {
		{
			ID:          3023,
			Slug:        "semux",
			Name:        "Semux",
			BinanceUSDT: "SEMUSDT",
		},
	},
	"POODL": {
		{
			ID:          8823,
			Slug:        "poodle",
			Name:        "Poodl Token",
			BinanceUSDT: "POODLUSDT",
		},
	},
	"ORO": {
		{
			ID:          7684,
			Slug:        "oro",
			Name:        "ORO",
			BinanceUSDT: "OROUSDT",
		},
	},
	"AAA": {
		{
			ID:          3287,
			Slug:        "abulaba",
			Name:        "Abulaba",
			BinanceUSDT: "AAAUSDT",
		},
	},
	"ARO": {
		{
			ID:          3024,
			Slug:        "arionum",
			Name:        "Arionum",
			BinanceUSDT: "AROUSDT",
		},
	},
	"SRSB": {
		{
			ID:          10025,
			Slug:        "sirius-bond",
			Name:        "Sirius Bond",
			BinanceUSDT: "SRSBUSDT",
		},
	},
	"VIP": {
		{
			ID:          719,
			Slug:        "limitless-vip",
			Name:        "Limitless VIP",
			BinanceUSDT: "VIPUSDT",
		},
	},
	"XFI": {
		{
			ID:   7038,
			Slug: "xfinance",
			Name: "Xfinance",
		},
		{
			ID:   7223,
			Slug: "dfinance",
			Name: "Dfinance",
		},
	},
	"BREW": {
		{
			ID:          8481,
			Slug:        "cafeswap-token",
			Name:        "CafeSwap Token",
			BinanceUSDT: "BREWUSDT",
		},
	},
	"EOS": {
		{
			ID:          1765,
			Slug:        "eos",
			Name:        "EOS",
			BinanceUSDT: "EOSUSDT",
		},
	},
	"XAS": {
		{
			ID:          1609,
			Slug:        "asch",
			Name:        "Asch",
			BinanceUSDT: "XASUSDT",
		},
	},
	"WNXM": {
		{
			ID:          5939,
			Slug:        "wrapped-nxm",
			Name:        "Wrapped NXM",
			BinanceUSDT: "WNXMUSDT",
		},
	},
	"ZCRT": {
		{
			ID:          5594,
			Slug:        "zcore-token",
			Name:        "ZCore Token",
			BinanceUSDT: "ZCRTUSDT",
		},
	},
	"DGP": {
		{
			ID:          7864,
			Slug:        "dgpayment",
			Name:        "DGPayment",
			BinanceUSDT: "DGPUSDT",
		},
	},
	"mFB": {
		{
			ID:          8585,
			Slug:        "mirrored-facebook",
			Name:        "Mirrored Facebook Inc",
			BinanceUSDT: "mFBUSDT",
		},
	},
	"NBS": {
		{
			ID:          7110,
			Slug:        "new-bitshares",
			Name:        "New BitShares",
			BinanceUSDT: "NBSUSDT",
		},
	},
	"xBLZD": {
		{
			ID:          8964,
			Slug:        "blizzard-money",
			Name:        "Blizzard.money",
			BinanceUSDT: "xBLZDUSDT",
		},
	},
	"ADAX": {
		{
			ID:          10833,
			Slug:        "adax",
			Name:        "ADAX",
			BinanceUSDT: "ADAXUSDT",
		},
	},
	"RTF": {
		{
			ID:          10097,
			Slug:        "regiment-finance",
			Name:        "Regiment Finance",
			BinanceUSDT: "RTFUSDT",
		},
	},
	"HYFI": {
		{
			ID:          9235,
			Slug:        "hyper-finance",
			Name:        "Hyper Finance",
			BinanceUSDT: "HYFIUSDT",
		},
	},
	"SMOON": {
		{
			ID:          9969,
			Slug:        "saylormoon",
			Name:        "SaylorMoon",
			BinanceUSDT: "SMOONUSDT",
		},
	},
	"RUPEE": {
		{
			ID:          8736,
			Slug:        "hyruleswap",
			Name:        "HyruleSwap",
			BinanceUSDT: "RUPEEUSDT",
		},
	},
	"SKI": {
		{
			ID:          5623,
			Slug:        "skillchain",
			Name:        "Skillchain",
			BinanceUSDT: "SKIUSDT",
		},
	},
	"TWT": {
		{
			ID:          5964,
			Slug:        "trust-wallet-token",
			Name:        "Trust Wallet Token",
			BinanceUSDT: "TWTUSDT",
		},
	},
	"NLC2": {
		{
			ID:          1382,
			Slug:        "nolimitcoin",
			Name:        "NoLimitCoin",
			BinanceUSDT: "NLC2USDT",
		},
	},
	"DAVP": {
		{
			ID:          4999,
			Slug:        "davion",
			Name:        "Davion",
			BinanceUSDT: "DAVPUSDT",
		},
	},
	"NKC": {
		{
			ID:          2477,
			Slug:        "nework",
			Name:        "Nework",
			BinanceUSDT: "NKCUSDT",
		},
	},
	"GCR": {
		{
			ID:   1044,
			Slug: "global-currency-reserve",
			Name: "Global Currency Reserve",
		},
		{
			ID:   7855,
			Slug: "gold-coin-reserve",
			Name: "Gold Coin Reserve",
		},
	},
	"BTMX": {
		{
			ID:          3673,
			Slug:        "bitmax-token",
			Name:        "ASD",
			BinanceUSDT: "BTMXUSDT",
		},
	},
	"FNX": {
		{
			ID:          5712,
			Slug:        "finnexus",
			Name:        "FinNexus",
			BinanceUSDT: "FNXUSDT",
		},
	},
	"BUNNY": {
		{
			ID:   7791,
			Slug: "pancakebunny",
			Name: "Pancake Bunny",
		},
		{
			ID:   8878,
			Slug: "rocket-bunny",
			Name: "Rocket Bunny",
		},
	},
	"ZPAE": {
		{
			ID:          5663,
			Slug:        "zelaapayae",
			Name:        "ZelaaPayAE",
			BinanceUSDT: "ZPAEUSDT",
		},
	},
	"DOGEMOON": {
		{
			ID:          9618,
			Slug:        "dogemoon",
			Name:        "DogeMoon",
			BinanceUSDT: "DOGEMOONUSDT",
		},
	},
	"PROUD": {
		{
			ID:          1398,
			Slug:        "proud-money",
			Name:        "PROUD Money",
			BinanceUSDT: "PROUDUSDT",
		},
	},
	"ETHBEAR": {
		{
			ID:          5216,
			Slug:        "3x-short-ethereum-token",
			Name:        "3X Short Ethereum Token",
			BinanceUSDT: "ETHBEARUSDT",
		},
	},
	"CLBK": {
		{
			ID:          3712,
			Slug:        "cloudbric",
			Name:        "Cloudbric",
			BinanceUSDT: "CLBKUSDT",
		},
	},
	"BFI": {
		{
			ID:   8796,
			Slug: "bearn",
			Name: "Bearn",
		},
		{
			ID:   7642,
			Slug: "bitdefi",
			Name: "BitDEFi",
		},
	},
	"ROVER": {
		{
			ID:          9702,
			Slug:        "rover-token",
			Name:        "Rover Inu Token",
			BinanceUSDT: "ROVERUSDT",
		},
	},
	"COP": {
		{
			ID:          9763,
			Slug:        "copiosa-coin",
			Name:        "Copiosa Coin",
			BinanceUSDT: "COPUSDT",
		},
	},
	"RING": {
		{
			ID:          5798,
			Slug:        "darwinia-network",
			Name:        "Darwinia Network",
			BinanceUSDT: "RINGUSDT",
		},
	},
	"HBDC": {
		{
			ID:          6542,
			Slug:        "happy-birthday-coin",
			Name:        "happy birthday coin",
			BinanceUSDT: "HBDCUSDT",
		},
	},
	"GAIA": {
		{
			ID:          9800,
			Slug:        "gaiadao",
			Name:        "GaiaDAO",
			BinanceUSDT: "GAIAUSDT",
		},
	},
	"KURT": {
		{
			ID:          1468,
			Slug:        "kurrent",
			Name:        "Kurrent",
			BinanceUSDT: "KURTUSDT",
		},
	},
	"SCH": {
		{
			ID:          4864,
			Slug:        "schilling-coin",
			Name:        "Schilling-Coin",
			BinanceUSDT: "SCHUSDT",
		},
	},
	"PRY": {
		{
			ID:          8241,
			Slug:        "prophecy",
			Name:        "Prophecy",
			BinanceUSDT: "PRYUSDT",
		},
	},
	"CMCT": {
		{
			ID:   3429,
			Slug: "cyber-movie-chain",
			Name: "Cyber Movie Chain",
		},
		{
			ID:   2708,
			Slug: "crowd-machine",
			Name: "Crowd Machine",
		},
	},
	"C98": {
		{
			ID:          10903,
			Slug:        "coin98",
			Name:        "Coin98",
			BinanceUSDT: "C98USDT",
		},
	},
	"BSC": {
		{
			ID:   993,
			Slug: "bowscoin",
			Name: "BowsCoin",
		},
		{
			ID:   7748,
			Slug: "bsc-farm",
			Name: "BSC FARM",
		},
		{
			ID:   5725,
			Slug: "bitsonic",
			Name: "Bitsonic",
		},
	},
	"vADA": {
		{
			ID:          9428,
			Slug:        "venus-cardano",
			Name:        "Venus Cardano",
			BinanceUSDT: "vADAUSDT",
		},
	},
	"NAT": {
		{
			ID:          3889,
			Slug:        "natmin-pure-escrow",
			Name:        "Natmin Pure Escrow",
			BinanceUSDT: "NATUSDT",
		},
	},
	"ELAMA": {
		{
			ID:          3734,
			Slug:        "elamachain",
			Name:        "Elamachain",
			BinanceUSDT: "ELAMAUSDT",
		},
	},
	"SGE": {
		{
			ID:          10735,
			Slug:        "society-of-galactic-exploration",
			Name:        "SOCIETY OF GALACTIC EXPLORATION",
			BinanceUSDT: "SGEUSDT",
		},
	},
	"PXL": {
		{
			ID:   4091,
			Slug: "pixel",
			Name: "PIXEL",
		},
		{
			ID:   9649,
			Slug: "pixl",
			Name: "PIXL",
		},
	},
	"LTCBEAR": {
		{
			ID:          5461,
			Slug:        "3x-short-litecoin-token",
			Name:        "3x Short Litecoin Token",
			BinanceUSDT: "LTCBEARUSDT",
		},
	},
	"DOGEBULL": {
		{
			ID:          6081,
			Slug:        "3x-long-dogecoin-token",
			Name:        "3X Long Dogecoin Token",
			BinanceUSDT: "DOGEBULLUSDT",
		},
	},
	"YETU": {
		{
			ID:          9139,
			Slug:        "yetucoin",
			Name:        "Yetucoin",
			BinanceUSDT: "YETUUSDT",
		},
	},
	"DOGS": {
		{
			ID:          10489,
			Slug:        "doggy-swap",
			Name:        "Doggy Swap",
			BinanceUSDT: "DOGSUSDT",
		},
	},
	"CARAT": {
		{
			ID:          3347,
			Slug:        "carat",
			Name:        "CARAT",
			BinanceUSDT: "CARATUSDT",
		},
	},
	"THS": {
		{
			ID:          6696,
			Slug:        "the-hash-speed",
			Name:        "The Hash Speed",
			BinanceUSDT: "THSUSDT",
		},
	},
	"CORA": {
		{
			ID:          9579,
			Slug:        "corra-finance",
			Name:        "Corra.Finance",
			BinanceUSDT: "CORAUSDT",
		},
	},
	"TSX": {
		{
			ID:          9563,
			Slug:        "tradestars",
			Name:        "TradeStars",
			BinanceUSDT: "TSXUSDT",
		},
	},
	"MANDI": {
		{
			ID:          6532,
			Slug:        "mandi-token",
			Name:        "Mandi Token",
			BinanceUSDT: "MANDIUSDT",
		},
	},
	"UWL": {
		{
			ID:          7727,
			Slug:        "uniwhales",
			Name:        "UniWhales",
			BinanceUSDT: "UWLUSDT",
		},
	},
	"GLDR": {
		{
			ID:          6809,
			Slug:        "goldergames",
			Name:        "GolderGames",
			BinanceUSDT: "GLDRUSDT",
		},
	},
	"MUTE": {
		{
			ID:          8795,
			Slug:        "mute",
			Name:        "Mute",
			BinanceUSDT: "MUTEUSDT",
		},
	},
	"LAVA": {
		{
			ID:   8469,
			Slug: "lavaswap",
			Name: "LavaSwap",
		},
		{
			ID:   10786,
			Slug: "lavacake-finance",
			Name: "LavaCake Finance",
		},
	},
	"GM": {
		{
			ID:          5698,
			Slug:        "gm-holding",
			Name:        "GM Holding",
			BinanceUSDT: "GMUSDT",
		},
	},
	"CUE": {
		{
			ID:          9121,
			Slug:        "cue-protocol",
			Name:        "CUE Protocol",
			BinanceUSDT: "CUEUSDT",
		},
	},
	"DHT": {
		{
			ID:          7094,
			Slug:        "dhedge-dao",
			Name:        "dHedge DAO",
			BinanceUSDT: "DHTUSDT",
		},
	},
	"CXN": {
		{
			ID:          7215,
			Slug:        "cxn-network",
			Name:        "CXN Network",
			BinanceUSDT: "CXNUSDT",
		},
	},
	"BHIBA": {
		{
			ID:          9769,
			Slug:        "baby-shiba",
			Name:        "Baby Shiba",
			BinanceUSDT: "BHIBAUSDT",
		},
	},
	"CLX": {
		{
			ID:          5179,
			Slug:        "celeum",
			Name:        "Celeum",
			BinanceUSDT: "CLXUSDT",
		},
	},
	"ZNZ": {
		{
			ID:          4286,
			Slug:        "zenzo",
			Name:        "ZENZO",
			BinanceUSDT: "ZNZUSDT",
		},
	},
	"MBC": {
		{
			ID:   3507,
			Slug: "microbitcoin",
			Name: "MicroBitcoin",
		},
		{
			ID:   5996,
			Slug: "marblecoin",
			Name: "Marblecoin",
		},
	},
	"GAP": {
		{
			ID:   4694,
			Slug: "gaps",
			Name: "GAPS",
		},
		{
			ID:   10617,
			Slug: "global-adversity-project",
			Name: "Global Adversity Project",
		},
	},
	"AOG": {
		{
			ID:          3316,
			Slug:        "smartofgiving",
			Name:        "smARTOFGIVING",
			BinanceUSDT: "AOGUSDT",
		},
	},
	"QTF": {
		{
			ID:          8531,
			Slug:        "quantfury-token",
			Name:        "Quantfury Token",
			BinanceUSDT: "QTFUSDT",
		},
	},
	"DMB": {
		{
			ID:          1687,
			Slug:        "digital-money-bits",
			Name:        "Digital Money Bits",
			BinanceUSDT: "DMBUSDT",
		},
	},
	"SGT": {
		{
			ID:   8445,
			Slug: "sharedstake",
			Name: "SharedStake",
		},
		{
			ID:   5889,
			Slug: "snglsdao",
			Name: "snglsDAO",
		},
		{
			ID:   10296,
			Slug: "spacegoat",
			Name: "SpaceGoat",
		},
	},
	"FOUR": {
		{
			ID:          5409,
			Slug:        "4thpillar-technologies",
			Name:        "4THPILLAR TECHNOLOGIES",
			BinanceUSDT: "FOURUSDT",
		},
	},
	"TX": {
		{
			ID:          1032,
			Slug:        "transfercoin",
			Name:        "TransferCoin",
			BinanceUSDT: "TXUSDT",
		},
	},
	"XUSD": {
		{
			ID:          8316,
			Slug:        "xusd-stable",
			Name:        "XUSD Stable",
			BinanceUSDT: "XUSDUSDT",
		},
	},
	"EDS": {
		{
			ID:          3076,
			Slug:        "endorsit",
			Name:        "Endorsit",
			BinanceUSDT: "EDSUSDT",
		},
	},
	"DOVE": {
		{
			ID:          9355,
			Slug:        "doveswap-finance",
			Name:        "DoveSwap Finance",
			BinanceUSDT: "DOVEUSDT",
		},
	},
	"DMTC": {
		{
			ID:          4983,
			Slug:        "demeter-chain",
			Name:        "Demeter Chain",
			BinanceUSDT: "DMTCUSDT",
		},
	},
	"KBOT": {
		{
			ID:          4404,
			Slug:        "korbot",
			Name:        "Korbot",
			BinanceUSDT: "KBOTUSDT",
		},
	},
	"PFL": {
		{
			ID:          9172,
			Slug:        "professional-fighters-league-fan-token",
			Name:        "Professional Fighters League Fan Token",
			BinanceUSDT: "PFLUSDT",
		},
	},
	"ALBT": {
		{
			ID:          6957,
			Slug:        "allianceblock",
			Name:        "AllianceBlock",
			BinanceUSDT: "ALBTUSDT",
		},
	},
	"TIP": {
		{
			ID:          8332,
			Slug:        "technology-innovation-project",
			Name:        "TECHNOLOGY INNOVATION PROJECT",
			BinanceUSDT: "TIPUSDT",
		},
	},
	"CNNC": {
		{
			ID:          1674,
			Slug:        "cannation",
			Name:        "Cannation",
			BinanceUSDT: "CNNCUSDT",
		},
	},
	"EOSDOWN": {
		{
			ID:          7000,
			Slug:        "eosdown",
			Name:        "EOSDOWN",
			BinanceUSDT: "EOSDOWNUSDT",
		},
	},
	"MASK": {
		{
			ID:   8536,
			Slug: "mask-network",
			Name: "Mask Network",
		},
		{
			ID:   8410,
			Slug: "nftx-hashmasks-index",
			Name: "NFTX Hashmasks Index",
		},
	},
	"USDF": {
		{
			ID:          6653,
			Slug:        "folgoryusd",
			Name:        "FolgoryUSD",
			BinanceUSDT: "USDFUSDT",
		},
	},
	"FJC": {
		{
			ID:          960,
			Slug:        "fujicoin",
			Name:        "FujiCoin",
			BinanceUSDT: "FJCUSDT",
		},
	},
	"UPDOG": {
		{
			ID:          9877,
			Slug:        "updog",
			Name:        "UPDOG",
			BinanceUSDT: "UPDOGUSDT",
		},
	},
	"AMLT": {
		{
			ID:          2607,
			Slug:        "amlt",
			Name:        "AMLT",
			BinanceUSDT: "AMLTUSDT",
		},
	},
	"BIST": {
		{
			ID:          9889,
			Slug:        "bistroo",
			Name:        "Bistroo",
			BinanceUSDT: "BISTUSDT",
		},
	},
	"ibETH": {
		{
			ID:          7675,
			Slug:        "ibeth",
			Name:        "Interest Bearing ETH",
			BinanceUSDT: "ibETHUSDT",
		},
	},
	"CASHDOG": {
		{
			ID:          10600,
			Slug:        "cashdog",
			Name:        "CashDog",
			BinanceUSDT: "CASHDOGUSDT",
		},
	},
	"RUNE": {
		{
			ID:   4157,
			Slug: "thorchain",
			Name: "THORChain",
		},
		{
			ID:   9905,
			Slug: "rune",
			Name: "Rune",
		},
		{
			ID:   8272,
			Slug: "thorchain-erc20",
			Name: "THORChain (ERC20)",
		},
	},
	"RLC": {
		{
			ID:          1637,
			Slug:        "rlc",
			Name:        "iExec RLC",
			BinanceUSDT: "RLCUSDT",
		},
	},
	"MBOX": {
		{
			ID:          9175,
			Slug:        "mobox",
			Name:        "MOBOX",
			BinanceUSDT: "MBOXUSDT",
		},
	},
	"LPOOL": {
		{
			ID:          8545,
			Slug:        "launchpool",
			Name:        "Launchpool",
			BinanceUSDT: "LPOOLUSDT",
		},
	},
	"1UP": {
		{
			ID:          4213,
			Slug:        "uptrennd",
			Name:        "Uptrennd",
			BinanceUSDT: "1UPUSDT",
		},
	},
	"WANEOS": {
		{
			ID:          8652,
			Slug:        "waneos",
			Name:        "wanEOS",
			BinanceUSDT: "WANEOSUSDT",
		},
	},
	"GMC": {
		{
			ID:          7161,
			Slug:        "gokumarket-credit",
			Name:        "GokuMarket Credit",
			BinanceUSDT: "GMCUSDT",
		},
	},
	"$COIN": {
		{
			ID:          7796,
			Slug:        "coin-defi",
			Name:        "COIN",
			BinanceUSDT: "$COINUSDT",
		},
	},
	"SNC": {
		{
			ID:          1786,
			Slug:        "suncontract",
			Name:        "SunContract",
			BinanceUSDT: "SNCUSDT",
		},
	},
	"USDP": {
		{
			ID:          8886,
			Slug:        "usdp",
			Name:        "USDP Stablecoin",
			BinanceUSDT: "USDPUSDT",
		},
	},
	"BOOMB": {
		{
			ID:          10517,
			Slug:        "boombaby-io",
			Name:        "BoomBaby.io",
			BinanceUSDT: "BOOMBUSDT",
		},
	},
	"CAJ": {
		{
			ID:          3715,
			Slug:        "cajutel",
			Name:        "Cajutel",
			BinanceUSDT: "CAJUSDT",
		},
	},
	"MEXC": {
		{
			ID:          4676,
			Slug:        "mexc-token",
			Name:        "MEXC Token",
			BinanceUSDT: "MEXCUSDT",
		},
	},
	"QI": {
		{
			ID:          8510,
			Slug:        "qiswap",
			Name:        "QiSwap",
			BinanceUSDT: "QIUSDT",
		},
	},
	"ETC": {
		{
			ID:          1321,
			Slug:        "ethereum-classic",
			Name:        "Ethereum Classic",
			BinanceUSDT: "ETCUSDT",
		},
	},
	"DCY": {
		{
			ID:          1745,
			Slug:        "dinastycoin",
			Name:        "Dinastycoin",
			BinanceUSDT: "DCYUSDT",
		},
	},
	"WEMIX": {
		{
			ID:          7548,
			Slug:        "wemix",
			Name:        "WEMIX",
			BinanceUSDT: "WEMIXUSDT",
		},
	},
	"KLEE": {
		{
			ID:          10262,
			Slug:        "kleekai",
			Name:        "KleeKai",
			BinanceUSDT: "KLEEUSDT",
		},
	},
	"GLEEC": {
		{
			ID:          5200,
			Slug:        "gleec",
			Name:        "Gleec",
			BinanceUSDT: "GLEECUSDT",
		},
	},
	"DDR": {
		{
			ID:          5900,
			Slug:        "digidinar",
			Name:        "DigiDinar",
			BinanceUSDT: "DDRUSDT",
		},
	},
	"MAPS": {
		{
			ID:          8166,
			Slug:        "maps",
			Name:        "MAPS",
			BinanceUSDT: "MAPSUSDT",
		},
	},
	"DAO": {
		{
			ID:          8420,
			Slug:        "dao-maker",
			Name:        "DAO Maker",
			BinanceUSDT: "DAOUSDT",
		},
	},
	"FUEL": {
		{
			ID:   2120,
			Slug: "etherparty",
			Name: "Etherparty",
		},
		{
			ID:   8659,
			Slug: "jetfuel-finance",
			Name: "Jetfuel Finance",
		},
	},
	"WSB": {
		{
			ID:          9749,
			Slug:        "wallstreetbets-dapp",
			Name:        "WallStreetBets DApp",
			BinanceUSDT: "WSBUSDT",
		},
	},
	"GIGA": {
		{
			ID:          9897,
			Slug:        "gigapool",
			Name:        "GigaPool",
			BinanceUSDT: "GIGAUSDT",
		},
	},
	"MEME": {
		{
			ID:   6597,
			Slug: "degenerator-meme",
			Name: "Meme",
		},
		{
			ID:   1191,
			Slug: "memetic",
			Name: "Memetic / PepeCoin",
		},
	},
	"SRN": {
		{
			ID:          2313,
			Slug:        "sirin-labs-token",
			Name:        "SIRIN LABS Token",
			BinanceUSDT: "SRNUSDT",
		},
	},
	"CAMP": {
		{
			ID:          7475,
			Slug:        "camp",
			Name:        "Camp",
			BinanceUSDT: "CAMPUSDT",
		},
	},
	"WISH": {
		{
			ID:          2236,
			Slug:        "mywish",
			Name:        "MyWish",
			BinanceUSDT: "WISHUSDT",
		},
	},
	"QKC": {
		{
			ID:          2840,
			Slug:        "quarkchain",
			Name:        "QuarkChain",
			BinanceUSDT: "QKCUSDT",
		},
	},
	"mBABA": {
		{
			ID:          8006,
			Slug:        "mirrored-alibaba",
			Name:        "Mirrored Alibaba",
			BinanceUSDT: "mBABAUSDT",
		},
	},
	"GSHIBA": {
		{
			ID:          11043,
			Slug:        "gambler-shiba",
			Name:        "Gambler Shiba",
			BinanceUSDT: "GSHIBAUSDT",
		},
	},
	"DOGEC": {
		{
			ID:          3672,
			Slug:        "dogecash",
			Name:        "DogeCash",
			BinanceUSDT: "DOGECUSDT",
		},
	},
	"OCT": {
		{
			ID:   1838,
			Slug: "oraclechain",
			Name: "OracleChain",
		},
		{
			ID:   8773,
			Slug: "octree",
			Name: "Octree",
		},
	},
	"FN": {
		{
			ID:          4817,
			Slug:        "filenet",
			Name:        "Filenet",
			BinanceUSDT: "FNUSDT",
		},
	},
	"LCG": {
		{
			ID:          7448,
			Slug:        "lcg",
			Name:        "LCG",
			BinanceUSDT: "LCGUSDT",
		},
	},
	"SXPUP": {
		{
			ID:          7528,
			Slug:        "sxpup",
			Name:        "SXPUP",
			BinanceUSDT: "SXPUPUSDT",
		},
	},
	"ODIN": {
		{
			ID:          9546,
			Slug:        "odin-protocol",
			Name:        "ODIN PROTOCOL",
			BinanceUSDT: "ODINUSDT",
		},
	},
	"PIT": {
		{
			ID:          9177,
			Slug:        "pitbull",
			Name:        "Pitbull",
			BinanceUSDT: "PITUSDT",
		},
	},
	"ACOIN": {
		{
			ID:   601,
			Slug: "acoin",
			Name: "Acoin",
		},
		{
			ID:   5465,
			Slug: "alchemy",
			Name: "Alchemy",
		},
	},
	"S4F": {
		{
			ID:          3733,
			Slug:        "s4fe",
			Name:        "S4FE",
			BinanceUSDT: "S4FUSDT",
		},
	},
	"DFGL": {
		{
			ID:          7483,
			Slug:        "defi-gold",
			Name:        "DeFi Gold",
			BinanceUSDT: "DFGLUSDT",
		},
	},
	"GUAP": {
		{
			ID:          5088,
			Slug:        "guapcoin",
			Name:        "Guapcoin",
			BinanceUSDT: "GUAPUSDT",
		},
	},
	"PTD": {
		{
			ID:   151,
			Slug: "pesetacoin",
			Name: "Peseta Digital",
		},
		{
			ID:   9039,
			Slug: "pilot",
			Name: "Pilot",
		},
	},
	"EKO": {
		{
			ID:          2391,
			Slug:        "echolink",
			Name:        "EchoLink",
			BinanceUSDT: "EKOUSDT",
		},
	},
	"XSN": {
		{
			ID:          2633,
			Slug:        "stakenet",
			Name:        "Stakenet",
			BinanceUSDT: "XSNUSDT",
		},
	},
	"FMA": {
		{
			ID:          5857,
			Slug:        "flama",
			Name:        "FLAMA",
			BinanceUSDT: "FMAUSDT",
		},
	},
	"TWEE": {
		{
			ID:          5285,
			Slug:        "tweebaa",
			Name:        "Tweebaa",
			BinanceUSDT: "TWEEUSDT",
		},
	},
	"JOB": {
		{
			ID:          4287,
			Slug:        "jobchain",
			Name:        "Jobchain",
			BinanceUSDT: "JOBUSDT",
		},
	},
	"CRW": {
		{
			ID:          720,
			Slug:        "crown",
			Name:        "Crown",
			BinanceUSDT: "CRWUSDT",
		},
	},
	"KTON": {
		{
			ID:          5931,
			Slug:        "darwinia-commitment-token",
			Name:        "Darwinia Commitment Token",
			BinanceUSDT: "KTONUSDT",
		},
	},
	"TOKO": {
		{
			ID:          4299,
			Slug:        "tokoin",
			Name:        "Tokoin",
			BinanceUSDT: "TOKOUSDT",
		},
	},
	"NEST": {
		{
			ID:          5841,
			Slug:        "nest-protocol",
			Name:        "NEST Protocol",
			BinanceUSDT: "NESTUSDT",
		},
	},
	"ILK": {
		{
			ID:          4754,
			Slug:        "inlock",
			Name:        "INLOCK",
			BinanceUSDT: "ILKUSDT",
		},
	},
	"MGO": {
		{
			ID:          1715,
			Slug:        "mobilego",
			Name:        "MobileGo",
			BinanceUSDT: "MGOUSDT",
		},
	},
	"DRGN": {
		{
			ID:          2243,
			Slug:        "dragonchain",
			Name:        "Dragonchain",
			BinanceUSDT: "DRGNUSDT",
		},
	},
	"PIGGY": {
		{
			ID:          10212,
			Slug:        "piggy-bank-token",
			Name:        "Piggy Bank Token",
			BinanceUSDT: "PIGGYUSDT",
		},
	},
	"FLP": {
		{
			ID:          2707,
			Slug:        "flip",
			Name:        "FLIP",
			BinanceUSDT: "FLPUSDT",
		},
	},
	"KSC": {
		{
			ID:          6493,
			Slug:        "kstarcoin",
			Name:        "KStarCoin",
			BinanceUSDT: "KSCUSDT",
		},
	},
	"IONC": {
		{
			ID:          3506,
			Slug:        "ionchain",
			Name:        "IONChain",
			BinanceUSDT: "IONCUSDT",
		},
	},
	"TOPB": {
		{
			ID:          6027,
			Slug:        "topb",
			Name:        "TOPBTC Token",
			BinanceUSDT: "TOPBUSDT",
		},
	},
	"SUSHIBEAR": {
		{
			ID:          7085,
			Slug:        "3x-short-sushi-token",
			Name:        "3X Short Sushi Token",
			BinanceUSDT: "SUSHIBEARUSDT",
		},
	},
	"NCASH": {
		{
			ID:          2544,
			Slug:        "nucleus-vision",
			Name:        "Nucleus Vision",
			BinanceUSDT: "NCASHUSDT",
		},
	},
	"OGN": {
		{
			ID:          5117,
			Slug:        "origin-protocol",
			Name:        "Origin Protocol",
			BinanceUSDT: "OGNUSDT",
		},
	},
	"VKNF": {
		{
			ID:          8611,
			Slug:        "vkenaf",
			Name:        "VKENAF",
			BinanceUSDT: "VKNFUSDT",
		},
	},
	"POLX": {
		{
			ID:          9615,
			Slug:        "polylastic",
			Name:        "Polylastic",
			BinanceUSDT: "POLXUSDT",
		},
	},
	"LOOM": {
		{
			ID:          2588,
			Slug:        "loom-network",
			Name:        "Loom Network",
			BinanceUSDT: "LOOMUSDT",
		},
	},
	"GIV": {
		{
			ID:          6075,
			Slug:        "givly-coin",
			Name:        "GIVLY Coin",
			BinanceUSDT: "GIVUSDT",
		},
	},
	"LEX": {
		{
			ID:          7578,
			Slug:        "elxis",
			Name:        "Elxis",
			BinanceUSDT: "LEXUSDT",
		},
	},
}
