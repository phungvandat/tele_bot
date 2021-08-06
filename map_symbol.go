package main

type CoinInfo struct {
	ID               uint16
	Slug             string
	Name             string
	LastPrice        float64
	LastAlertSuccess bool
}

var mapSymbolCoins = map[string][]CoinInfo{
	"SKILL": {
		{
			ID:   9654,
			Slug: "cryptoblades",
			Name: "CryptoBlades",
		},
	},
	"RVT": {
		{
			ID:   1991,
			Slug: "rivetz",
			Name: "Rivetz",
		},
	},
	"BBR": {
		{
			ID:   406,
			Slug: "boolberry",
			Name: "Boolberry",
		},
	},
	"BAFI": {
		{
			ID:   9075,
			Slug: "bafi-finance",
			Name: "Bafi Finance",
		},
	},
	"EXZO": {
		{
			ID:   10528,
			Slug: "exzocoin-2",
			Name: "ExzoCoin 2.0",
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
	"DGTX": {
		{
			ID:   2772,
			Slug: "digitex",
			Name: "Digitex",
		},
	},
	"ETGP": {
		{
			ID:   3810,
			Slug: "ethereum-gold-project",
			Name: "Ethereum Gold Project",
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
	"TONS": {
		{
			ID:   7060,
			Slug: "thisoption",
			Name: "Thisoption",
		},
	},
	"INNBCL": {
		{
			ID:   3787,
			Slug: "innovative-bioresearch-classic",
			Name: "Innovative Bioresearch Classic",
		},
	},
	"BTC3L": {
		{
			ID:   5736,
			Slug: "amun-bitcoin-3x-daily-long",
			Name: "Amun Bitcoin 3x Daily Long",
		},
	},
	"NAFTY": {
		{
			ID:   10864,
			Slug: "nafty",
			Name: "NAFTY",
		},
	},
	"VANCAT": {
		{
			ID:   8955,
			Slug: "vancat",
			Name: "Vancat",
		},
	},
	"INXT": {
		{
			ID:   2022,
			Slug: "internxt",
			Name: "Internxt",
		},
	},
	"YAY": {
		{
			ID:   10104,
			Slug: "yayswap",
			Name: "YaySwap",
		},
	},
	"EHRT": {
		{
			ID:   6107,
			Slug: "eight-hours",
			Name: "Eight Hours",
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
	"NUKO": {
		{
			ID:   2257,
			Slug: "nekonium",
			Name: "Nekonium",
		},
	},
	"ZILD": {
		{
			ID:   10381,
			Slug: "zild-finance",
			Name: "Zild Finance",
		},
	},
	"SNOB": {
		{
			ID:   9780,
			Slug: "snowball-finance",
			Name: "Snowball",
		},
	},
	"HERPES": {
		{
			ID:   10879,
			Slug: "herpes",
			Name: "Herpes",
		},
	},
	"PRQ": {
		{
			ID:   5410,
			Slug: "parsiq",
			Name: "PARSIQ",
		},
	},
	"INSUR": {
		{
			ID:   8799,
			Slug: "insurace",
			Name: "InsurAce",
		},
	},
	"ISLA": {
		{
			ID:   5389,
			Slug: "insula",
			Name: "Insula",
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
	"QUEENSHIBA": {
		{
			ID:   10983,
			Slug: "queen-of-shiba",
			Name: "Queen of Shiba",
		},
	},
	"ALTBULL": {
		{
			ID:   6077,
			Slug: "3x-long-altcoin-index-token",
			Name: "3X Long Altcoin Index Token",
		},
	},
	"SMDX": {
		{
			ID:   7688,
			Slug: "somidax",
			Name: "SOMIDAX",
		},
	},
	"AGVC": {
		{
			ID:   3664,
			Slug: "agavecoin",
			Name: "AgaveCoin",
		},
	},
	"DIGG": {
		{
			ID:   8307,
			Slug: "digg",
			Name: "DIGG",
		},
	},
	"QUN": {
		{
			ID:   2375,
			Slug: "qunqun",
			Name: "QunQun",
		},
	},
	"BOLT": {
		{
			ID:   3843,
			Slug: "bolt",
			Name: "BOLT",
		},
	},
	"FUZZ": {
		{
			ID:   1241,
			Slug: "fuzzballs",
			Name: "FuzzBalls",
		},
	},
	"888": {
		{
			ID:   6051,
			Slug: "888tron",
			Name: "888tron",
		},
	},
	"ACES": {
		{
			ID:   1351,
			Slug: "aces",
			Name: "Aces",
		},
	},
	"ATL": {
		{
			ID:   2136,
			Slug: "atlant",
			Name: "ATLANT",
		},
	},
	"YAP": {
		{
			ID:   4899,
			Slug: "yap-stone",
			Name: "Yap Stone",
		},
	},
	"MWAR": {
		{
			ID:   10840,
			Slug: "mooniwar",
			Name: "MooniWar",
		},
	},
	"PET": {
		{
			ID:   8915,
			Slug: "battle-pets",
			Name: "Battle Pets",
		},
	},
	"TME": {
		{
			ID:   8455,
			Slug: "tama-egg-niftygotchi",
			Name: "TAMA EGG NiftyGotchi",
		},
	},
	"HOME": {
		{
			ID:   10342,
			Slug: "homecoin",
			Name: "HomeCoin",
		},
	},
	"ONT": {
		{
			ID:   2566,
			Slug: "ontology",
			Name: "Ontology",
		},
	},
	"FTXT": {
		{
			ID:   3219,
			Slug: "futurax",
			Name: "FUTURAX",
		},
	},
	"IOUX": {
		{
			ID:   3996,
			Slug: "iou",
			Name: "IOU",
		},
	},
	"GBCR": {
		{
			ID:   7557,
			Slug: "gold-bcr",
			Name: "Gold BCR",
		},
	},
	"ALEPH": {
		{
			ID:   5821,
			Slug: "aleph-im",
			Name: "Aleph.im",
		},
	},
	"TX": {
		{
			ID:   1032,
			Slug: "transfercoin",
			Name: "TransferCoin",
		},
	},
	"DDOS": {
		{
			ID:   9024,
			Slug: "disbalancer",
			Name: "disBalancer",
		},
	},
	"VOTE": {
		{
			ID:   4792,
			Slug: "agora",
			Name: "Agora",
		},
	},
	"XTZDOWN": {
		{
			ID:   7008,
			Slug: "xtzdown",
			Name: "XTZDOWN",
		},
	},
	"BSCX": {
		{
			ID:   8190,
			Slug: "bscex",
			Name: "BSCEX",
		},
	},
	"PUML": {
		{
			ID:   5842,
			Slug: "puml-better-health",
			Name: "PUML Better Health",
		},
	},
	"MLR": {
		{
			ID:   5433,
			Slug: "mega-lottery-services-global",
			Name: "Mega Lottery Services Global",
		},
	},
	"TBB": {
		{
			ID:   8090,
			Slug: "trade-butler-bot",
			Name: "Trade Butler Bot",
		},
	},
	"MISO": {
		{
			ID:   10309,
			Slug: "miso",
			Name: "MISO",
		},
	},
	"ALICE": {
		{
			ID:   8766,
			Slug: "myneighboralice",
			Name: "MyNeighborAlice",
		},
	},
	"APIX": {
		{
			ID:   5258,
			Slug: "apix",
			Name: "APIX",
		},
	},
	"WAND": {
		{
			ID:   2269,
			Slug: "wandx",
			Name: "WandX",
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
	"DWZ": {
		{
			ID:   8641,
			Slug: "defi-wizard",
			Name: "DeFi Wizard",
		},
	},
	"WANSUSHI": {
		{
			ID:   8653,
			Slug: "wansushi",
			Name: "wanSUSHI",
		},
	},
	"GOM": {
		{
			ID:   4623,
			Slug: "gomics",
			Name: "Gomics",
		},
	},
	"TROP": {
		{
			ID:   8336,
			Slug: "interop",
			Name: "Interop",
		},
	},
	"VESPASHIBA": {
		{
			ID:   10795,
			Slug: "vespa-shiba-coin",
			Name: "VESPA SHIBA COIN",
		},
	},
	"AXEL": {
		{
			ID:   6216,
			Slug: "axel",
			Name: "AXEL",
		},
	},
	"FLTY": {
		{
			ID:   9722,
			Slug: "fluity",
			Name: "Fluity",
		},
	},
	"RYO": {
		{
			ID:   2976,
			Slug: "ryo-currency",
			Name: "Ryo Currency",
		},
	},
	"BNF": {
		{
			ID:   7436,
			Slug: "bonfi",
			Name: "BonFi",
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
	"SLNV2": {
		{
			ID:   8530,
			Slug: "starlink",
			Name: "StarLink",
		},
	},
	"HYPER": {
		{
			ID:   9309,
			Slug: "hyperchain",
			Name: "HyperChain",
		},
	},
	"USDP": {
		{
			ID:   8886,
			Slug: "usdp",
			Name: "USDP Stablecoin",
		},
	},
	"PROPS": {
		{
			ID:   5880,
			Slug: "props",
			Name: "Props Token",
		},
	},
	"BOR": {
		{
			ID:   7509,
			Slug: "boringdao",
			Name: "BoringDAO",
		},
	},
	"UBN": {
		{
			ID:   4071,
			Slug: "ubricoin",
			Name: "Ubricoin",
		},
	},
	"pBTC35A": {
		{
			ID:   8252,
			Slug: "pbtc35a",
			Name: "pBTC35A",
		},
	},
	"YEFIM": {
		{
			ID:   7400,
			Slug: "yfi-management",
			Name: "YFi Management",
		},
	},
	"PROM": {
		{
			ID:   4120,
			Slug: "prometeus",
			Name: "Prometeus",
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
	"BCP": {
		{
			ID:   7867,
			Slug: "bitcashpay",
			Name: "Bitcashpay",
		},
	},
	"WXC": {
		{
			ID:   3363,
			Slug: "wxcoins",
			Name: "WXCOINS",
		},
	},
	"GST2": {
		{
			ID:   5716,
			Slug: "gas-token-two",
			Name: "Gas Token Two",
		},
	},
	"YELD": {
		{
			ID:   6974,
			Slug: "yeld-finance",
			Name: "Yeld Finance",
		},
	},
	"PTM": {
		{
			ID:   8358,
			Slug: "potentiam",
			Name: "Potentiam",
		},
	},
	"VTUBE": {
		{
			ID:   9745,
			Slug: "vtube-token",
			Name: "VTube Token",
		},
	},
	"SPIRIT": {
		{
			ID:   10239,
			Slug: "spiritswap",
			Name: "SpiritSwap",
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
	"EVX": {
		{
			ID:   2034,
			Slug: "everex",
			Name: "Everex",
		},
	},
	"ETL": {
		{
			ID:   10832,
			Slug: "etherlite",
			Name: "Etherlite",
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
	"COVAL": {
		{
			ID:   788,
			Slug: "circuits-of-value",
			Name: "Circuits of Value",
		},
	},
	"VNXLU": {
		{
			ID:   4430,
			Slug: "vnx-exchange",
			Name: "VNX Exchange",
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
	"ETHP": {
		{
			ID:   5778,
			Slug: "ethplus",
			Name: "ETHPlus",
		},
	},
	"QQQ": {
		{
			ID:   4644,
			Slug: "poseidon-network",
			Name: "Poseidon Network",
		},
	},
	"1WO": {
		{
			ID:   2601,
			Slug: "1world",
			Name: "1World",
		},
	},
	"FMA": {
		{
			ID:   5857,
			Slug: "flama",
			Name: "FLAMA",
		},
	},
	"REAP": {
		{
			ID:   7677,
			Slug: "reapchain",
			Name: "ReapChain",
		},
	},
	"VRC": {
		{
			ID:   323,
			Slug: "vericoin",
			Name: "VeriCoin",
		},
	},
	"LUN": {
		{
			ID:   1658,
			Slug: "lunyr",
			Name: "Lunyr",
		},
	},
	"ARKK": {
		{
			ID:   7896,
			Slug: "ark-innovation-etf-tokenized-stock-ftx",
			Name: "ARK Innovation ETF tokenized stock FTX",
		},
	},
	"BNU": {
		{
			ID:   10121,
			Slug: "bytenext",
			Name: "ByteNext",
		},
	},
	"MOB": {
		{
			ID:   7878,
			Slug: "mobilecoin",
			Name: "MobileCoin",
		},
	},
	"GNT": {
		{
			ID:   9533,
			Slug: "greentrust",
			Name: "GreenTrust",
		},
	},
	"CARAT": {
		{
			ID:   3347,
			Slug: "carat",
			Name: "CARAT",
		},
	},
	"CAVO": {
		{
			ID:   8286,
			Slug: "excavo-finance",
			Name: "Excavo Finance",
		},
	},
	"FUSE": {
		{
			ID:   5634,
			Slug: "fuse-network",
			Name: "Fuse Network",
		},
	},
	"EDR": {
		{
			ID:   2835,
			Slug: "endor-protocol",
			Name: "Endor Protocol",
		},
	},
	"XCHF": {
		{
			ID:   4075,
			Slug: "cryptofranc",
			Name: "CryptoFranc",
		},
	},
	"BZE": {
		{
			ID:   3961,
			Slug: "bzedge",
			Name: "BZEdge",
		},
	},
	"WLITI": {
		{
			ID:   10740,
			Slug: "liti-capital",
			Name: "Liti Capital",
		},
	},
	"HIVE": {
		{
			ID:   5370,
			Slug: "hive-blockchain",
			Name: "Hive",
		},
	},
	"VELO": {
		{
			ID:   7127,
			Slug: "velo",
			Name: "Velo",
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
	"HACHIKO": {
		{
			ID:   9988,
			Slug: "hachiko-inu",
			Name: "Hachiko Inu",
		},
	},
	"YTA": {
		{
			ID:   4133,
			Slug: "yottachain",
			Name: "YottaChain",
		},
	},
	"DOGET": {
		{
			ID:   3919,
			Slug: "doge-token",
			Name: "Doge Token",
		},
	},
	"URG": {
		{
			ID:   10719,
			Slug: "urgaming",
			Name: "UrGaming",
		},
	},
	"ROX": {
		{
			ID:   3325,
			Slug: "robotina",
			Name: "Robotina",
		},
	},
	"MLA": {
		{
			ID:   8619,
			Slug: "moola",
			Name: "Moola",
		},
	},
	"SNET": {
		{
			ID:   3435,
			Slug: "snetwork",
			Name: "Snetwork",
		},
	},
	"DZI": {
		{
			ID:   6662,
			Slug: "definition",
			Name: "DeFinition",
		},
	},
	"ELET": {
		{
			ID:   3931,
			Slug: "elementeum",
			Name: "Elementeum",
		},
	},
	"PANTHER": {
		{
			ID:   9778,
			Slug: "pantherswap",
			Name: "PantherSwap",
		},
	},
	"TCO": {
		{
			ID:   8556,
			Slug: "tcoin-token",
			Name: "Tcoin",
		},
	},
	"CMIT": {
		{
			ID:   3328,
			Slug: "cmitcoin",
			Name: "CMITCOIN",
		},
	},
	"ADGNZ": {
		{
			ID:   10259,
			Slug: "degen-token-finance",
			Name: "Degen Token Finance",
		},
	},
	"CUSDT": {
		{
			ID:   5745,
			Slug: "compound-usdt",
			Name: "Compound USDT",
		},
	},
	"LTCUP": {
		{
			ID:   7526,
			Slug: "ltcup",
			Name: "LTCUP",
		},
	},
	"USELESS": {
		{
			ID:   10851,
			Slug: "useless",
			Name: "Useless",
		},
	},
	"DIGS": {
		{
			ID:   10980,
			Slug: "digies-coin",
			Name: "Digies Coin",
		},
	},
	"XTRM": {
		{
			ID:   5599,
			Slug: "xtrm-coin",
			Name: "XTRM COIN",
		},
	},
	"HYPE": {
		{
			ID:   8130,
			Slug: "hype",
			Name: "Supreme Finance",
		},
	},
	"ZDR": {
		{
			ID:   3899,
			Slug: "zloadr",
			Name: "Zloadr",
		},
	},
	"XFC": {
		{
			ID:   3663,
			Slug: "footballcoin",
			Name: "Footballcoin",
		},
	},
	"MONK": {
		{
			ID:   2230,
			Slug: "monk",
			Name: "MONK",
		},
	},
	"CVNT": {
		{
			ID:   3686,
			Slug: "content-value-network",
			Name: "Content Value Network",
		},
	},
	"YAMV2": {
		{
			ID:   6657,
			Slug: "yam-v2",
			Name: "YAMv2",
		},
	},
	"USDC": {
		{
			ID:   3408,
			Slug: "usd-coin",
			Name: "USD Coin",
		},
	},
	"RAVEN": {
		{
			ID:   4024,
			Slug: "raven-protocol",
			Name: "Raven Protocol",
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
	"POST": {
		{
			ID:   1218,
			Slug: "postcoin",
			Name: "PostCoin",
		},
	},
	"MINI": {
		{
			ID:   6405,
			Slug: "miniswap",
			Name: "MiniSwap",
		},
	},
	"UME": {
		{
			ID:   8852,
			Slug: "ume-token",
			Name: "UME Token",
		},
	},
	"DCN": {
		{
			ID:   1876,
			Slug: "dentacoin",
			Name: "Dentacoin",
		},
	},
	"KAT": {
		{
			ID:   3634,
			Slug: "kambria",
			Name: "Kambria",
		},
	},
	"DDS": {
		{
			ID:   9032,
			Slug: "dds-store",
			Name: "DDS.Store",
		},
	},
	"PLG": {
		{
			ID:   4196,
			Slug: "pledge-coin",
			Name: "Pledge Coin",
		},
	},
	"VAIN": {
		{
			ID:   10761,
			Slug: "vain",
			Name: "Vain",
		},
	},
	"VAULT": {
		{
			ID:   5003,
			Slug: "vault",
			Name: "VAULT",
		},
	},
	"PHI": {
		{
			ID:   2676,
			Slug: "phi-token",
			Name: "PHI Token",
		},
	},
	"ART": {
		{
			ID:   2064,
			Slug: "maecenas",
			Name: "Maecenas",
		},
	},
	"BNOX": {
		{
			ID:   5648,
			Slug: "blocknotex",
			Name: "BlockNoteX",
		},
	},
	"USDU": {
		{
			ID:   6907,
			Slug: "upper-dollar",
			Name: "Upper Dollar",
		},
	},
	"SHIBBY": {
		{
			ID:   10616,
			Slug: "shibby",
			Name: "Shibby",
		},
	},
	"SKI": {
		{
			ID:   5623,
			Slug: "skillchain",
			Name: "Skillchain",
		},
	},
	"STOGE": {
		{
			ID:   9381,
			Slug: "stoner-doge",
			Name: "Stoner Doge Finance",
		},
	},
	"NAV": {
		{
			ID:   377,
			Slug: "nav-coin",
			Name: "Navcoin",
		},
	},
	"CRPT": {
		{
			ID:   2447,
			Slug: "crpt",
			Name: "Crypterium",
		},
	},
	"NFTI": {
		{
			ID:   8700,
			Slug: "nft-index",
			Name: "NFT Index",
		},
	},
	"YFN": {
		{
			ID:   7478,
			Slug: "yearn-finance-network",
			Name: "Yearn Finance Network",
		},
	},
	"FBT": {
		{
			ID:   6442,
			Slug: "fanbi-token",
			Name: "FANBI TOKEN",
		},
	},
	"MIDAS": {
		{
			ID:   4505,
			Slug: "midas",
			Name: "Midas",
		},
	},
	"HQX": {
		{
			ID:   2564,
			Slug: "hoqu",
			Name: "HOQU",
		},
	},
	"ITAM": {
		{
			ID:   6490,
			Slug: "itam-games",
			Name: "ITAM Games",
		},
	},
	"WENMOON": {
		{
			ID:   9400,
			Slug: "wenmoon",
			Name: "WenMoon",
		},
	},
	"SXI": {
		{
			ID:   9624,
			Slug: "safexi",
			Name: "SafeXI",
		},
	},
	"KDA": {
		{
			ID:   5647,
			Slug: "kadena",
			Name: "Kadena",
		},
	},
	"NUX": {
		{
			ID:   8458,
			Slug: "peanut",
			Name: "Peanut",
		},
	},
	"TIKI": {
		{
			ID:   10684,
			Slug: "tiki-token",
			Name: "Tiki Token",
		},
	},
	"LOBS": {
		{
			ID:   3203,
			Slug: "lobstex",
			Name: "Lobstex",
		},
	},
	"SXUT": {
		{
			ID:   2382,
			Slug: "spectre-utility",
			Name: "Spectre.ai Utility Token",
		},
	},
	"WAF": {
		{
			ID:   9283,
			Slug: "waffle",
			Name: "Waffle",
		},
	},
	"MANGO": {
		{
			ID:   8856,
			Slug: "mango-finance",
			Name: "Mango Finance",
		},
	},
	"DMCH": {
		{
			ID:   5622,
			Slug: "darma-cash",
			Name: "Darma Cash",
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
	"COSM": {
		{
			ID:   2955,
			Slug: "cosmo-coin",
			Name: "Cosmo Coin",
		},
	},
	"CYL": {
		{
			ID:   3580,
			Slug: "crystal-token",
			Name: "Crystal Token",
		},
	},
	"PATH": {
		{
			ID:   9878,
			Slug: "pathfund",
			Name: "PathFund",
		},
	},
	"XLD": {
		{
			ID:   10766,
			Slug: "stellar-diamond",
			Name: "Stellar Diamond",
		},
	},
	"MP3": {
		{
			ID:   8412,
			Slug: "mp3",
			Name: "MP3",
		},
	},
	"GUSD": {
		{
			ID:   3306,
			Slug: "gemini-dollar",
			Name: "Gemini Dollar",
		},
	},
	"BCN": {
		{
			ID:   372,
			Slug: "bytecoin-bcn",
			Name: "Bytecoin",
		},
	},
	"DSLA": {
		{
			ID:   5423,
			Slug: "dsla-protocol",
			Name: "DSLA Protocol",
		},
	},
	"DHT": {
		{
			ID:   7094,
			Slug: "dhedge-dao",
			Name: "dHedge DAO",
		},
	},
	"REDPANDA": {
		{
			ID:   9980,
			Slug: "redpanda",
			Name: "Redpanda Earth",
		},
	},
	"CALI": {
		{
			ID:   9595,
			Slug: "calicoin",
			Name: "CaliCoin",
		},
	},
	"DMTC": {
		{
			ID:   4983,
			Slug: "demeter-chain",
			Name: "Demeter Chain",
		},
	},
	"IDXM": {
		{
			ID:   2422,
			Slug: "idex-membership",
			Name: "IDEX Membership",
		},
	},
	"OMG": {
		{
			ID:   1808,
			Slug: "omg",
			Name: "OMG Network",
		},
	},
	"COV": {
		{
			ID:   2342,
			Slug: "covesting",
			Name: "Covesting",
		},
	},
	"LTCR": {
		{
			ID:   1155,
			Slug: "litecred",
			Name: "Litecred",
		},
	},
	"SUSHIBULL": {
		{
			ID:   7084,
			Slug: "3x-long-sushi-token",
			Name: "3X Long Sushi Token",
		},
	},
	"ZCX": {
		{
			ID:   9263,
			Slug: "unizen",
			Name: "Unizen",
		},
	},
	"LITH": {
		{
			ID:   10592,
			Slug: "lith-token",
			Name: "Lith Token",
		},
	},
	"HEBE": {
		{
			ID:   5250,
			Slug: "hebeblock",
			Name: "HebeBlock",
		},
	},
	"NEU": {
		{
			ID:   2318,
			Slug: "neumark",
			Name: "Neumark",
		},
	},
	"SNX": {
		{
			ID:   2586,
			Slug: "synthetix-network-token",
			Name: "Synthetix",
		},
	},
	"IUT": {
		{
			ID:   4737,
			Slug: "ito-utility-token",
			Name: "ITO Utility Token",
		},
	},
	"BONE": {
		{
			ID:   5915,
			Slug: "bone",
			Name: "Bone",
		},
	},
	"PLE": {
		{
			ID:   9606,
			Slug: "plethori",
			Name: "Plethori",
		},
	},
	"AURUM": {
		{
			ID:   9224,
			Slug: "alchemist-defi-aurum",
			Name: "Alchemist DeFi Aurum",
		},
	},
	"PQT": {
		{
			ID:   8804,
			Slug: "prediqt",
			Name: "PREDIQT",
		},
	},
	"HNDC": {
		{
			ID:   3504,
			Slug: "hondaiscoin",
			Name: "HondaisCoin",
		},
	},
	"IDLE": {
		{
			ID:   7841,
			Slug: "idle",
			Name: "Idle",
		},
	},
	"AOG": {
		{
			ID:   3316,
			Slug: "smartofgiving",
			Name: "smARTOFGIVING",
		},
	},
	"VOX": {
		{
			ID:   7465,
			Slug: "vox-finance",
			Name: "Vox.Finance",
		},
	},
	"SHDC": {
		{
			ID:   8740,
			Slug: "shd-cash",
			Name: "SHD CASH",
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
	"BRY": {
		{
			ID:   8483,
			Slug: "berry-data",
			Name: "Berry Data",
		},
	},
	"HUSKY": {
		{
			ID:   9387,
			Slug: "husky",
			Name: "Husky",
		},
	},
	"MEET": {
		{
			ID:   2470,
			Slug: "coinmeet",
			Name: "CoinMeet",
		},
	},
	"KAM": {
		{
			ID:   5107,
			Slug: "bitkam",
			Name: "BitKAM",
		},
	},
	"TWEE": {
		{
			ID:   5285,
			Slug: "tweebaa",
			Name: "Tweebaa",
		},
	},
	"SBNB": {
		{
			ID:   6204,
			Slug: "sbnb",
			Name: "sBNB",
		},
	},
	"ALPHR": {
		{
			ID:   9430,
			Slug: "alphr-finance",
			Name: "Alphr finance",
		},
	},
	"REAL": {
		{
			ID:   2030,
			Slug: "real",
			Name: "REAL",
		},
	},
	"BSHIBA": {
		{
			ID:   9829,
			Slug: "shiba-corp",
			Name: "Shiba Corp",
		},
	},
	"FRM": {
		{
			ID:   4228,
			Slug: "ferrum-network",
			Name: "Ferrum Network",
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
	"TWTR": {
		{
			ID:   7916,
			Slug: "twitter-tokenized-stock-ftx",
			Name: "Twitter tokenized stock FTX",
		},
	},
	"LIEN": {
		{
			ID:   6705,
			Slug: "lien",
			Name: "Lien",
		},
	},
	"PEEPS": {
		{
			ID:   10952,
			Slug: "the-peoples-coin",
			Name: "The People's Coin",
		},
	},
	"EOC": {
		{
			ID:   8048,
			Slug: "everyonescrypto",
			Name: "Everyonescrypto",
		},
	},
	"HUP": {
		{
			ID:   10678,
			Slug: "hup-life",
			Name: "HUP.LIFE",
		},
	},
	"SAV3": {
		{
			ID:   7680,
			Slug: "sav3token",
			Name: "Sav3Token",
		},
	},
	"WEST": {
		{
			ID:   5159,
			Slug: "waves-enterprise",
			Name: "Waves Enterprise",
		},
	},
	"KONO": {
		{
			ID:   8697,
			Slug: "konomi-network",
			Name: "Konomi Network",
		},
	},
	"PORNROCKET": {
		{
			ID:   10165,
			Slug: "pornrocket",
			Name: "PORNROCKET",
		},
	},
	"$TREAM": {
		{
			ID:   11006,
			Slug: "world-stream-finance",
			Name: "World Stream Finance",
		},
	},
	"DMC": {
		{
			ID:   8610,
			Slug: "decentralized-mining-exchange",
			Name: "Decentralized Mining Exchange",
		},
	},
	"SAL": {
		{
			ID:   6874,
			Slug: "salmonswap",
			Name: "SalmonSwap",
		},
	},
	"WVG0": {
		{
			ID:   7471,
			Slug: "wrapped-virgin-gen-0-cryptokitties",
			Name: "Wrapped Virgin Gen-0 CryptoKitties",
		},
	},
	"QTF": {
		{
			ID:   8531,
			Slug: "quantfury-token",
			Name: "Quantfury Token",
		},
	},
	"SAFEWIN": {
		{
			ID:   10536,
			Slug: "safewin",
			Name: "SafeWin",
		},
	},
	"NAMI": {
		{
			ID:   7534,
			Slug: "tsunami",
			Name: "Tsunami finance",
		},
	},
	"BTCP": {
		{
			ID:   2575,
			Slug: "bitcoin-private",
			Name: "Bitcoin Private",
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
	"GIFT": {
		{
			ID:   10283,
			Slug: "gift-coin",
			Name: "Gift-Coin",
		},
	},
	"PROPHET": {
		{
			ID:   7773,
			Slug: "prophet",
			Name: "Prophet",
		},
	},
	"ENRG": {
		{
			ID:   322,
			Slug: "energycoin",
			Name: "Energycoin",
		},
	},
	"TRUBGR": {
		{
			ID:   10902,
			Slug: "trubadger",
			Name: "TruBadger",
		},
	},
	"EXEN": {
		{
			ID:   8565,
			Slug: "exen-coin",
			Name: "Exen Coin",
		},
	},
	"FLOW": {
		{
			ID:   4558,
			Slug: "flow",
			Name: "Flow",
		},
	},
	"ASTA": {
		{
			ID:   6375,
			Slug: "asta",
			Name: "ASTA",
		},
	},
	"LAUNCH": {
		{
			ID:   9279,
			Slug: "superlauncher",
			Name: "SuperLauncher",
		},
	},
	"BCMC1": {
		{
			ID:   9318,
			Slug: "beforecoinmarketcap",
			Name: "BeforeCoinMarketCap",
		},
	},
	"SHIBACASH": {
		{
			ID:   9736,
			Slug: "shibacash",
			Name: "ShibaCash",
		},
	},
	"OAP": {
		{
			ID:   6970,
			Slug: "openalexa-protocol",
			Name: "OpenAlexa Protocol",
		},
	},
	"KIRO": {
		{
			ID:   7276,
			Slug: "kirobo",
			Name: "Kirobo",
		},
	},
	"xMOON": {
		{
			ID:   7396,
			Slug: "moon",
			Name: "r/CryptoCurrency Moons",
		},
	},
	"YYFI": {
		{
			ID:   7551,
			Slug: "yyfi-protocol",
			Name: "YYFI.Protocol",
		},
	},
	"$PEEPO": {
		{
			ID:   10146,
			Slug: "peepocoin",
			Name: "PeepoCoin",
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
	"FX": {
		{
			ID:   3884,
			Slug: "function-x",
			Name: "Function X",
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
	"THN": {
		{
			ID:   10805,
			Slug: "throne",
			Name: "Throne",
		},
	},
	"ZHEGIC": {
		{
			ID:   7601,
			Slug: "zhegic",
			Name: "zHEGIC",
		},
	},
	"ATA": {
		{
			ID:   10188,
			Slug: "automata-network",
			Name: "Automata Network",
		},
	},
	"BZRX": {
		{
			ID:   5810,
			Slug: "bzx-protocol",
			Name: "bZx Protocol",
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
	"WASABI": {
		{
			ID:   9077,
			Slug: "wasabix",
			Name: "WasabiX",
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
	"PKB": {
		{
			ID:   934,
			Slug: "parkbyte",
			Name: "ParkByte",
		},
	},
	"CAMP": {
		{
			ID:   7475,
			Slug: "camp",
			Name: "Camp",
		},
	},
	"NBTC": {
		{
			ID:   5888,
			Slug: "neobitcoin",
			Name: "NEOBITCOIN",
		},
	},
	"BTMX": {
		{
			ID:   3673,
			Slug: "bitmax-token",
			Name: "ASD",
		},
	},
	"ABST": {
		{
			ID:   4518,
			Slug: "abitshadow-token",
			Name: "Abitshadow Token",
		},
	},
	"ETCBULL": {
		{
			ID:   6083,
			Slug: "3x-long-ethereum-classic-token",
			Name: "3X Long Ethereum Classic Token",
		},
	},
	"GBPU": {
		{
			ID:   6906,
			Slug: "upper-pound",
			Name: "Upper Pound",
		},
	},
	"I9C": {
		{
			ID:   6966,
			Slug: "i9-coin",
			Name: "i9 Coin",
		},
	},
	"MSP": {
		{
			ID:   10553,
			Slug: "moonship-finance",
			Name: "Moonship Finance",
		},
	},
	"BRIGADEIRO": {
		{
			ID:   10654,
			Slug: "brigadeiro",
			Name: "Brigadeiro.Finance",
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
	"GRC": {
		{
			ID:   833,
			Slug: "gridcoin",
			Name: "Gridcoin",
		},
	},
	"BBS": {
		{
			ID:   3093,
			Slug: "bbscoin",
			Name: "BBSCoin",
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
	"ZOM": {
		{
			ID:   7063,
			Slug: "zoom-protocol",
			Name: "Zoom Protocol",
		},
	},
	"ERG": {
		{
			ID:   1762,
			Slug: "ergo",
			Name: "Ergo",
		},
	},
	"EURS": {
		{
			ID:   2989,
			Slug: "stasis-euro",
			Name: "STASIS EURO",
		},
	},
	"TFBX": {
		{
			ID:   4144,
			Slug: "truefeedback",
			Name: "TrueFeedBack",
		},
	},
	"NDX": {
		{
			ID:   8260,
			Slug: "indexed-finance",
			Name: "Indexed Finance",
		},
	},
	"VRO": {
		{
			ID:   5539,
			Slug: "veraone",
			Name: "VeraOne",
		},
	},
	"LSK": {
		{
			ID:   1214,
			Slug: "lisk",
			Name: "Lisk",
		},
	},
	"IBFK": {
		{
			ID:   8884,
			Slug: "istanbul-basaksehir-fan-token",
			Name: "İstanbul Başakşehir Fan Token",
		},
	},
	"MATICBULL": {
		{
			ID:   6085,
			Slug: "3x-long-matic-token",
			Name: "3X Long Matic Token",
		},
	},
	"HGH": {
		{
			ID:   4989,
			Slug: "hgh-token",
			Name: "HGH Token",
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
	"DND": {
		{
			ID:   9578,
			Slug: "dungeonswap",
			Name: "Dungeonswap",
		},
	},
	"CUB": {
		{
			ID:   8858,
			Slug: "cub-finance",
			Name: "Cub Finance",
		},
	},
	"DEALDOUGH": {
		{
			ID:   10672,
			Slug: "dealdough-token",
			Name: "DealDough Token",
		},
	},
	"BITCNY": {
		{
			ID:   624,
			Slug: "bitcny",
			Name: "bitCNY",
		},
	},
	"SEELE": {
		{
			ID:   2830,
			Slug: "seele",
			Name: "Seele-N",
		},
	},
	"LAR": {
		{
			ID:   5251,
			Slug: "linkart",
			Name: "LinkArt",
		},
	},
	"SOVI": {
		{
			ID:   8741,
			Slug: "sovi-finance",
			Name: "Sovi Finance",
		},
	},
	"KONJ": {
		{
			ID:   5192,
			Slug: "konjungate",
			Name: "KONJUNGATE",
		},
	},
	"TUDA": {
		{
			ID:   4301,
			Slug: "tutors-diary",
			Name: "Tutor's Diary",
		},
	},
	"TLM": {
		{
			ID:   9119,
			Slug: "alien-worlds",
			Name: "Alien Worlds",
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
	"RMT": {
		{
			ID:   2527,
			Slug: "sureremit",
			Name: "SureRemit",
		},
	},
	"PRVS": {
		{
			ID:   8493,
			Slug: "previse",
			Name: "Previse",
		},
	},
	"PIGGY": {
		{
			ID:   10212,
			Slug: "piggy-bank-token",
			Name: "Piggy Bank Token",
		},
	},
	"RICK": {
		{
			ID:   6475,
			Slug: "infinite-ricks",
			Name: "Infinite Ricks",
		},
	},
	"NGM": {
		{
			ID:   8279,
			Slug: "e-money-coin",
			Name: "e-Money",
		},
	},
	"SCRIV": {
		{
			ID:   3398,
			Slug: "scriv-network",
			Name: "SCRIV NETWORK",
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
	"VCCO": {
		{
			ID:   7341,
			Slug: "vera-cruz-coin",
			Name: "Vera Cruz Coin",
		},
	},
	"YFMS": {
		{
			ID:   7093,
			Slug: "yfmoonshot",
			Name: "YFMoonshot",
		},
	},
	"CRU": {
		{
			ID:   6747,
			Slug: "crustnetwork",
			Name: "Crust Network",
		},
	},
	"XPTX": {
		{
			ID:   1254,
			Slug: "platinumbar",
			Name: "PlatinumBAR",
		},
	},
	"PNGN": {
		{
			ID:   9404,
			Slug: "spacepenguin",
			Name: "SpacePenguin",
		},
	},
	"BOLTT": {
		{
			ID:   3789,
			Slug: "boltt-coin",
			Name: "Boltt Coin ",
		},
	},
	"WWAN": {
		{
			ID:   8658,
			Slug: "wrapped-wan",
			Name: "Wrapped WAN",
		},
	},
	"$WNTR": {
		{
			ID:   9981,
			Slug: "weentar",
			Name: "Weentar",
		},
	},
	"TCAKE": {
		{
			ID:   9063,
			Slug: "pancaketools",
			Name: "Tcake",
		},
	},
	"ARTY": {
		{
			ID:   10790,
			Slug: "artys-world",
			Name: "Arty's World",
		},
	},
	"BDO": {
		{
			ID:   8219,
			Slug: "bdollar",
			Name: "bDollar",
		},
	},
	"PAMP": {
		{
			ID:   10323,
			Slug: "pamp-cc",
			Name: "PAMP.CC",
		},
	},
	"PUMP": {
		{
			ID:   10649,
			Slug: "moonpump",
			Name: "MoonPump",
		},
	},
	"ROSE": {
		{
			ID:   7653,
			Slug: "oasis-network",
			Name: "Oasis Network",
		},
	},
	"KP3R": {
		{
			ID:   7535,
			Slug: "keep3rv1",
			Name: "Keep3rV1",
		},
	},
	"KEX": {
		{
			ID:   6930,
			Slug: "kira-network",
			Name: "Kira Network",
		},
	},
	"GRFT": {
		{
			ID:   2571,
			Slug: "graft",
			Name: "Graft",
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
	"EL": {
		{
			ID:   5382,
			Slug: "elysia",
			Name: "ELYSIA",
		},
	},
	"HEAT": {
		{
			ID:   1308,
			Slug: "heat-ledger",
			Name: "HEAT",
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
	"TURTLE": {
		{
			ID:   10757,
			Slug: "turtle",
			Name: "Turtle",
		},
	},
	"NYC": {
		{
			ID:   298,
			Slug: "newyorkcoin",
			Name: "NewYorkCoin",
		},
	},
	"A2A": {
		{
			ID:   8745,
			Slug: "a2a-50x-com",
			Name: "A2A",
		},
	},
	"UHP": {
		{
			ID:   5949,
			Slug: "ulgen-hash-power",
			Name: "Ulgen Hash Power",
		},
	},
	"MTN": {
		{
			ID:   2497,
			Slug: "medical-chain",
			Name: "Medicalchain",
		},
	},
	"HERB": {
		{
			ID:   3646,
			Slug: "herbalist-token",
			Name: "Herbalist Token",
		},
	},
	"CCF": {
		{
			ID:   10755,
			Slug: "cashcow-finance",
			Name: "Cashcow Finance",
		},
	},
	"BAMBOO": {
		{
			ID:   8389,
			Slug: "bamboo-defi",
			Name: "BambooDeFi",
		},
	},
	"PDO": {
		{
			ID:   9210,
			Slug: "pollo",
			Name: "Pollo Dollar",
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
	"SDS": {
		{
			ID:   3119,
			Slug: "alchemint-standards",
			Name: "Alchemint Standards",
		},
	},
	"AUDIO": {
		{
			ID:   7455,
			Slug: "audius",
			Name: "Audius",
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
	"BOB": {
		{
			ID:   2889,
			Slug: "bobs-repair",
			Name: "Bob's Repair",
		},
	},
	"THOREUM": {
		{
			ID:   10787,
			Slug: "thoreum",
			Name: "Thoreum",
		},
	},
	"USDL": {
		{
			ID:   7364,
			Slug: "usdl",
			Name: "USDL",
		},
	},
	"CRB": {
		{
			ID:   10345,
			Slug: "cribnb-decentralized-renting-and-sharing",
			Name: "Cribnb Decentralized Renting and Sharing",
		},
	},
	"SEPHI": {
		{
			ID:   10758,
			Slug: "sephiroth-inu",
			Name: "Sephiroth Inu",
		},
	},
	"ANW": {
		{
			ID:   6120,
			Slug: "anchor-neural-world",
			Name: "Anchor Neural World",
		},
	},
	"RINGX": {
		{
			ID:   4894,
			Slug: "ring-x-platform",
			Name: "RING X PLATFORM",
		},
	},
	"FRMX": {
		{
			ID:   7631,
			Slug: "frmx-token",
			Name: "FRMx Token",
		},
	},
	"SUP8EME": {
		{
			ID:   7701,
			Slug: "sup8eme",
			Name: "SUP8EME",
		},
	},
	"YFA": {
		{
			ID:   6909,
			Slug: "yfa-finance",
			Name: "YFA Finance",
		},
	},
	"SPAZ": {
		{
			ID:   4797,
			Slug: "swapcoinz",
			Name: "Swapcoinz",
		},
	},
	"POLX": {
		{
			ID:   9615,
			Slug: "polylastic",
			Name: "Polylastic",
		},
	},
	"PAPEL": {
		{
			ID:   9359,
			Slug: "papel-token",
			Name: "Papel Token",
		},
	},
	"BTY": {
		{
			ID:   6224,
			Slug: "bityuan",
			Name: "Bityuan",
		},
	},
	"MLN": {
		{
			ID:   1552,
			Slug: "enzyme",
			Name: "Enzyme",
		},
	},
	"PI": {
		{
			ID:   2838,
			Slug: "pchain",
			Name: "Plian",
		},
	},
	"PRT": {
		{
			ID:   7348,
			Slug: "portion",
			Name: "Portion",
		},
	},
	"WSPP": {
		{
			ID:   10841,
			Slug: "wolf-safe-poor-people",
			Name: "Wolf Safe Poor People",
		},
	},
	"IZE": {
		{
			ID:   5677,
			Slug: "ize",
			Name: "IZE",
		},
	},
	"vLTC": {
		{
			ID:   7964,
			Slug: "venus-ltc",
			Name: "Venus LTC",
		},
	},
	"XNO": {
		{
			ID:   8368,
			Slug: "xeno-token",
			Name: "Xeno Token",
		},
	},
	"CYF": {
		{
			ID:   7179,
			Slug: "cy-finance",
			Name: "CY Finance",
		},
	},
	"XFT": {
		{
			ID:   6236,
			Slug: "offshift",
			Name: "Offshift",
		},
	},
	"DSR": {
		{
			ID:   2148,
			Slug: "desire",
			Name: "Desire",
		},
	},
	"NTVRK": {
		{
			ID:   9954,
			Slug: "netvrk",
			Name: "Netvrk",
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
	"OCTA": {
		{
			ID:   9273,
			Slug: "octans",
			Name: "Octans",
		},
	},
	"BLV": {
		{
			ID:   7309,
			Slug: "bellevue-network",
			Name: "Bellevue Network",
		},
	},
	"TOWER": {
		{
			ID:   8620,
			Slug: "tower-token",
			Name: "Tower",
		},
	},
	"OUSD": {
		{
			ID:   7189,
			Slug: "origin-dollar",
			Name: "Origin Dollar",
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
	"DOGEFI": {
		{
			ID:   6623,
			Slug: "dogefi",
			Name: "DOGEFI",
		},
	},
	"ELX": {
		{
			ID:   8331,
			Slug: "energy-ledger",
			Name: "Energy Ledger",
		},
	},
	"EMB": {
		{
			ID:   9626,
			Slug: "emblem",
			Name: "Emblem",
		},
	},
	"ON": {
		{
			ID:   7021,
			Slug: "ofin-token",
			Name: "OFIN Token",
		},
	},
	"UFT": {
		{
			ID:   7412,
			Slug: "unilend",
			Name: "UniLend",
		},
	},
	"GSR": {
		{
			ID:   1846,
			Slug: "geysercoin",
			Name: "GeyserCoin",
		},
	},
	"PINKPANDA": {
		{
			ID:   10513,
			Slug: "pinkpanda",
			Name: "PinkPanda",
		},
	},
	"HEDGE": {
		{
			ID:   5656,
			Slug: "1x-short-bitcoin-token",
			Name: "1x Short Bitcoin Token",
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
	"XUSD": {
		{
			ID:   8316,
			Slug: "xusd-stable",
			Name: "XUSD Stable",
		},
	},
	"PSY": {
		{
			ID:   8347,
			Slug: "psychic",
			Name: "Psychic",
		},
	},
	"BTM": {
		{
			ID:   1866,
			Slug: "bytom",
			Name: "Bytom",
		},
	},
	"ARCX": {
		{
			ID:   10515,
			Slug: "arcx-token",
			Name: "ARC Governance",
		},
	},
	"TULIP₿": {
		{
			ID:   10477,
			Slug: "tulips-city",
			Name: "Tulips City",
		},
	},
	"XNK": {
		{
			ID:   2549,
			Slug: "ink-protocol",
			Name: "Ink Protocol",
		},
	},
	"BLES": {
		{
			ID:   9026,
			Slug: "blind-boxes",
			Name: "Blind Boxes",
		},
	},
	"TENFI": {
		{
			ID:   10031,
			Slug: "ten",
			Name: "TEN",
		},
	},
	"DEXG": {
		{
			ID:   6684,
			Slug: "dextoken",
			Name: "Dextoken",
		},
	},
	"HOGT": {
		{
			ID:   10605,
			Slug: "hogt",
			Name: "HOGT",
		},
	},
	"DENT": {
		{
			ID:   1886,
			Slug: "dent",
			Name: "Dent",
		},
	},
	"BORA": {
		{
			ID:   3801,
			Slug: "bora",
			Name: "BORA",
		},
	},
	"APM": {
		{
			ID:   5079,
			Slug: "apm-coin",
			Name: "apM Coin",
		},
	},
	"BTX": {
		{
			ID:   1654,
			Slug: "bitcore",
			Name: "BitCore",
		},
	},
	"QCH": {
		{
			ID:   3337,
			Slug: "qchi",
			Name: "QChi",
		},
	},
	"SBGO": {
		{
			ID:   9552,
			Slug: "bingo-cash-finance",
			Name: "Bingo Share",
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
	"SHA": {
		{
			ID:   3831,
			Slug: "safe-haven",
			Name: "Safe Haven",
		},
	},
	"CTRT": {
		{
			ID:   3317,
			Slug: "cryptrust",
			Name: "Cryptrust",
		},
	},
	"VBK": {
		{
			ID:   3846,
			Slug: "veriblock",
			Name: "VeriBlock",
		},
	},
	"ABET": {
		{
			ID:   4502,
			Slug: "altbet",
			Name: "Altbet",
		},
	},
	"CANN": {
		{
			ID:   506,
			Slug: "cannabiscoin",
			Name: "CannabisCoin",
		},
	},
	"CFL": {
		{
			ID:   3313,
			Slug: "cryptoflow",
			Name: "CryptoFlow",
		},
	},
	"RAGE": {
		{
			ID:   8862,
			Slug: "rage-fan",
			Name: "Rage Fan",
		},
	},
	"PKG": {
		{
			ID:   3106,
			Slug: "pkg-token",
			Name: "PKG Token",
		},
	},
	"SEM": {
		{
			ID:   3023,
			Slug: "semux",
			Name: "Semux",
		},
	},
	"PAMPTHER": {
		{
			ID:   9925,
			Slug: "pampther",
			Name: "Pampther",
		},
	},
	"JSHIBA": {
		{
			ID:   9685,
			Slug: "jomon-shiba",
			Name: "Jomon Shiba",
		},
	},
	"ONION": {
		{
			ID:   1881,
			Slug: "deeponion",
			Name: "DeepOnion",
		},
	},
	"GRIMM": {
		{
			ID:   4866,
			Slug: "grimm",
			Name: "Grimm",
		},
	},
	"RNBW": {
		{
			ID:   10429,
			Slug: "halodao",
			Name: "HaloDAO",
		},
	},
	"DUN": {
		{
			ID:   5160,
			Slug: "dune-network",
			Name: "Dune Network",
		},
	},
	"FAN8": {
		{
			ID:   10701,
			Slug: "fan8",
			Name: "FAN8",
		},
	},
	"XCASH": {
		{
			ID:   3334,
			Slug: "x-cash",
			Name: "X-CASH",
		},
	},
	"DEFI++": {
		{
			ID:   7874,
			Slug: "piedao-defi",
			Name: "PieDAO DEFI++",
		},
	},
	"ZP": {
		{
			ID:   3186,
			Slug: "zen-protocol",
			Name: "Zen Protocol",
		},
	},
	"ECOM": {
		{
			ID:   3081,
			Slug: "omnitude",
			Name: "Omnitude",
		},
	},
	"CTX": {
		{
			ID:   10368,
			Slug: "cryptex-finance",
			Name: "Cryptex Finance",
		},
	},
	"DIGEX": {
		{
			ID:   6680,
			Slug: "digex",
			Name: "Digex",
		},
	},
	"BUSD": {
		{
			ID:   4687,
			Slug: "binance-usd",
			Name: "Binance USD",
		},
	},
	"XOV": {
		{
			ID:   3097,
			Slug: "xovbank",
			Name: "XOVBank",
		},
	},
	"KLKS": {
		{
			ID:   3018,
			Slug: "kalkulus",
			Name: "Kalkulus",
		},
	},
	"YFPI": {
		{
			ID:   7249,
			Slug: "yearn-finance-passive-income",
			Name: "Yearn Finance Passive Income",
		},
	},
	"ETGF": {
		{
			ID:   7454,
			Slug: "etg-finance",
			Name: "ETG Finance",
		},
	},
	"TRIB": {
		{
			ID:   7170,
			Slug: "contribute",
			Name: "Contribute",
		},
	},
	"JBX": {
		{
			ID:   6704,
			Slug: "jboxcoin",
			Name: "JBOX",
		},
	},
	"EDAO": {
		{
			ID:   10258,
			Slug: "elondoge-dao",
			Name: "ElonDoge DAO",
		},
	},
	"SENSE": {
		{
			ID:   2402,
			Slug: "sense",
			Name: "Sense",
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
	"PRINCESS": {
		{
			ID:   10849,
			Slug: "paris-inuton",
			Name: "Paris Inuton",
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
	"FONT": {
		{
			ID:   8601,
			Slug: "font",
			Name: "Font",
		},
	},
	"IMX": {
		{
			ID:   9532,
			Slug: "impermax",
			Name: "Impermax",
		},
	},
	"ATOMBEAR": {
		{
			ID:   6096,
			Slug: "3x-short-cosmos-token",
			Name: "3X Short Cosmos Token",
		},
	},
	"GHOSTFACE": {
		{
			ID:   10697,
			Slug: "ghostface",
			Name: "GHOSTFACE",
		},
	},
	"FRIDGE": {
		{
			ID:   9042,
			Slug: "fridge-token",
			Name: "Fridge Token",
		},
	},
	"AZ": {
		{
			ID:   4777,
			Slug: "azbit",
			Name: "Azbit",
		},
	},
	"YLDY": {
		{
			ID:   10820,
			Slug: "yieldly",
			Name: "Yieldly",
		},
	},
	"AWX": {
		{
			ID:   7301,
			Slug: "aurusdefi",
			Name: "AurusDeFi",
		},
	},
	"XTX": {
		{
			ID:   3844,
			Slug: "xtock",
			Name: "Xtock",
		},
	},
	"ERN": {
		{
			ID:   8615,
			Slug: "ethernity-chain",
			Name: "Ethernity Chain",
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
	"MNTP": {
		{
			ID:   2513,
			Slug: "goldmint",
			Name: "GoldMint",
		},
	},
	"GYSR": {
		{
			ID:   7661,
			Slug: "gysr",
			Name: "GYSR",
		},
	},
	"NIRX": {
		{
			ID:   4971,
			Slug: "nairax",
			Name: "NairaX",
		},
	},
	"FINE": {
		{
			ID:   9269,
			Slug: "refinable",
			Name: "Refinable",
		},
	},
	"IDNA": {
		{
			ID:   5836,
			Slug: "idena",
			Name: "Idena",
		},
	},
	"HVCO": {
		{
			ID:   1282,
			Slug: "high-voltage",
			Name: "High Voltage",
		},
	},
	"WEX": {
		{
			ID:   9951,
			Slug: "waultswap",
			Name: "WaultSwap",
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
	"UNIFI": {
		{
			ID:   7122,
			Slug: "unifi-defi",
			Name: "UNIFI DeFi",
		},
	},
	"NXT": {
		{
			ID:   66,
			Slug: "nxt",
			Name: "Nxt",
		},
	},
	"SUR": {
		{
			ID:   1933,
			Slug: "suretly",
			Name: "Suretly",
		},
	},
	"DRUNK": {
		{
			ID:   10964,
			Slug: "drunkdoge",
			Name: "DrunkDoge",
		},
	},
	"TEMCO": {
		{
			ID:   3722,
			Slug: "temco",
			Name: "TEMCO",
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
	"GTM": {
		{
			ID:   3215,
			Slug: "gentarium",
			Name: "Gentarium",
		},
	},
	"NAN": {
		{
			ID:   6529,
			Slug: "nantrade",
			Name: "NanTrade",
		},
	},
	"ONES": {
		{
			ID:   7136,
			Slug: "oneswap-dao-token",
			Name: "OneSwap DAO Token",
		},
	},
	"WQT": {
		{
			ID:   9115,
			Slug: "workquest",
			Name: "WorkQuest",
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
	"ONS": {
		{
			ID:   8160,
			Slug: "one-share",
			Name: "One Share",
		},
	},
	"BAK": {
		{
			ID:   6217,
			Slug: "tokenbacon",
			Name: "TokenBacon",
		},
	},
	"HMQ": {
		{
			ID:   1669,
			Slug: "humaniq",
			Name: "Humaniq",
		},
	},
	"SNRG": {
		{
			ID:   951,
			Slug: "synergy",
			Name: "Synergy",
		},
	},
	"PUX": {
		{
			ID:   6022,
			Slug: "polypux",
			Name: "PolypuX",
		},
	},
	"POODL": {
		{
			ID:   8823,
			Slug: "poodle",
			Name: "Poodl Token",
		},
	},
	"SRPH": {
		{
			ID:   10069,
			Slug: "seraphium",
			Name: "Seraphium",
		},
	},
	"RAINI": {
		{
			ID:   9061,
			Slug: "rainicorn",
			Name: "Rainicorn",
		},
	},
	"FOC": {
		{
			ID:   9259,
			Slug: "theforce-trade",
			Name: "TheForce Trade",
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
	"TOSC": {
		{
			ID:   3707,
			Slug: "t-os",
			Name: "T.OS",
		},
	},
	"DAGT": {
		{
			ID:   3357,
			Slug: "digital-asset-guarantee-token",
			Name: "Digital Asset Guarantee Token",
		},
	},
	"SXTZ": {
		{
			ID:   6192,
			Slug: "sxtz",
			Name: "sXTZ",
		},
	},
	"ZMT": {
		{
			ID:   8146,
			Slug: "zipmex",
			Name: "Zipmex",
		},
	},
	"APPC": {
		{
			ID:   2344,
			Slug: "appcoins",
			Name: "AppCoins",
		},
	},
	"AMB": {
		{
			ID:   2081,
			Slug: "amber",
			Name: "Ambrosus",
		},
	},
	"DMG": {
		{
			ID:   5741,
			Slug: "dmm-governance",
			Name: "DMM: Governance",
		},
	},
	"EQL": {
		{
			ID:   2479,
			Slug: "equal",
			Name: "Equal",
		},
	},
	"WHT": {
		{
			ID:   8524,
			Slug: "wrapped-huobi-token",
			Name: "Wrapped Huobi Token",
		},
	},
	"REST": {
		{
			ID:   8089,
			Slug: "restore",
			Name: "Restore",
		},
	},
	"DEFT": {
		{
			ID:   10044,
			Slug: "defi-factory-token",
			Name: "DeFi Factory Token",
		},
	},
	"AMO": {
		{
			ID:   3260,
			Slug: "amo-coin",
			Name: "AMO Coin",
		},
	},
	"FLIXX": {
		{
			ID:   2231,
			Slug: "flixxo",
			Name: "Flixxo",
		},
	},
	"POOL": {
		{
			ID:   8508,
			Slug: "pooltogether",
			Name: "PoolTogether",
		},
	},
	"LMCH": {
		{
			ID:   5418,
			Slug: "latamcash",
			Name: "Latamcash",
		},
	},
	"TINIDAWG": {
		{
			ID:   10694,
			Slug: "minidog-finance",
			Name: "MiniDog Finance",
		},
	},
	"LKN": {
		{
			ID:   6323,
			Slug: "linkcoin-token",
			Name: "LinkCoin Token",
		},
	},
	"XNS": {
		{
			ID:   4802,
			Slug: "xeonbit-token",
			Name: "Xeonbit Token",
		},
	},
	"ZOOM": {
		{
			ID:   5926,
			Slug: "coinzoom",
			Name: "CoinZoom",
		},
	},
	"SCH": {
		{
			ID:   4864,
			Slug: "schilling-coin",
			Name: "Schilling-Coin",
		},
	},
	"FAI": {
		{
			ID:   8414,
			Slug: "fairum",
			Name: "Fairum",
		},
	},
	"DXD": {
		{
			ID:   5589,
			Slug: "dxdao",
			Name: "DXdao",
		},
	},
	"SWM": {
		{
			ID:   2506,
			Slug: "swarm-network",
			Name: "Swarm",
		},
	},
	"DAT": {
		{
			ID:   2283,
			Slug: "datum",
			Name: "Datum",
		},
	},
	"MOONMOON": {
		{
			ID:   9525,
			Slug: "moonmoon",
			Name: "MoonMoon",
		},
	},
	"WTF": {
		{
			ID:   6848,
			Slug: "walnut-finance",
			Name: "Walnut.finance",
		},
	},
	"LTHN": {
		{
			ID:   2185,
			Slug: "lethean",
			Name: "Lethean",
		},
	},
	"ROYA": {
		{
			ID:   7821,
			Slug: "royale-finance",
			Name: "Royale Finance",
		},
	},
	"SWOP": {
		{
			ID:   8732,
			Slug: "swop",
			Name: "Swop",
		},
	},
	"BIA": {
		{
			ID:   3872,
			Slug: "bilaxy-token",
			Name: "Bilaxy Token",
		},
	},
	"OBR": {
		{
			ID:   7713,
			Slug: "order-of-the-black-rose",
			Name: "Order of the Black Rose",
		},
	},
	"CE": {
		{
			ID:   9329,
			Slug: "crypto-excellence",
			Name: "Crypto Excellence",
		},
	},
	"PHX": {
		{
			ID:   10206,
			Slug: "phoenix-protocol",
			Name: "Phoenix Protocol",
		},
	},
	"ENQ": {
		{
			ID:   4245,
			Slug: "enecuum",
			Name: "Enecuum",
		},
	},
	"AOS": {
		{
			ID:   6017,
			Slug: "aos",
			Name: "AOS",
		},
	},
	"KDOGE": {
		{
			ID:   10261,
			Slug: "kingdoge",
			Name: "KINGDOGE",
		},
	},
	"CYFM": {
		{
			ID:   3146,
			Slug: "cyberfm",
			Name: "CyberFM",
		},
	},
	"YFBTC": {
		{
			ID:   8464,
			Slug: "yfbitcoin",
			Name: "YFBitcoin",
		},
	},
	"USDB": {
		{
			ID:   5219,
			Slug: "usd-bancor",
			Name: "USD Bancor",
		},
	},
	"CXN": {
		{
			ID:   7215,
			Slug: "cxn-network",
			Name: "CXN Network",
		},
	},
	"CVC": {
		{
			ID:   1816,
			Slug: "civic",
			Name: "Civic",
		},
	},
	"MWAT": {
		{
			ID:   2533,
			Slug: "restart-energy-mwat",
			Name: "Restart Energy MWAT",
		},
	},
	"EXNT": {
		{
			ID:   6882,
			Slug: "exnetwork-token",
			Name: "ExNetwork Token",
		},
	},
	"INNBC": {
		{
			ID:   5016,
			Slug: "innovative-bioresearch-coin",
			Name: "Innovative Bioresearch Coin",
		},
	},
	"FUNDX": {
		{
			ID:   8681,
			Slug: "funder-one-capital",
			Name: "Funder One Capital",
		},
	},
	"DOOR": {
		{
			ID:   10873,
			Slug: "door",
			Name: "DOOR",
		},
	},
	"GYEN": {
		{
			ID:   8771,
			Slug: "gyen",
			Name: "GYEN",
		},
	},
	"POOLZ": {
		{
			ID:   8271,
			Slug: "poolz-finance",
			Name: "Poolz Finance",
		},
	},
	"MTX": {
		{
			ID:   2325,
			Slug: "matryx",
			Name: "Matryx",
		},
	},
	"CORAL": {
		{
			ID:   7628,
			Slug: "coral-swap",
			Name: "Coral Swap",
		},
	},
	"TRX": {
		{
			ID:   1958,
			Slug: "tron",
			Name: "TRON",
		},
	},
	"XPRT": {
		{
			ID:   7281,
			Slug: "persistence",
			Name: "Persistence",
		},
	},
	"BTRN": {
		{
			ID:   2690,
			Slug: "biotron",
			Name: "Biotron",
		},
	},
	"ADADOWN": {
		{
			ID:   7014,
			Slug: "adadown",
			Name: "ADADOWN",
		},
	},
	"SAFU": {
		{
			ID:   9375,
			Slug: "ceezee-safu",
			Name: "CEEZEE SAFU",
		},
	},
	"PUBE": {
		{
			ID:   9669,
			Slug: "pube-finance",
			Name: "Pube finance",
		},
	},
	"GMNG": {
		{
			ID:   7399,
			Slug: "global-gaming",
			Name: "Global Gaming",
		},
	},
	"VN": {
		{
			ID:   5828,
			Slug: "vn-token",
			Name: "VN Token",
		},
	},
	"TIME": {
		{
			ID:   1556,
			Slug: "chrono-tech",
			Name: "Chrono.tech",
		},
	},
	"CARDS": {
		{
			ID:   9047,
			Slug: "card-starter",
			Name: "CARD.STARTER",
		},
	},
	"BKING": {
		{
			ID:   10550,
			Slug: "king-arthur",
			Name: "King Arthur",
		},
	},
	"SHIBAPUP": {
		{
			ID:   9708,
			Slug: "shibapup",
			Name: "ShibaPup",
		},
	},
	"ARES": {
		{
			ID:   8702,
			Slug: "ares-protocol",
			Name: "Ares Protocol",
		},
	},
	"CRP": {
		{
			ID:   6865,
			Slug: "utopia",
			Name: "Crypton",
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
	"ABNB": {
		{
			ID:   7932,
			Slug: "airbnb-tokenized-stock-ftx",
			Name: "Airbnb tokenized stock FTX",
		},
	},
	"C2O": {
		{
			ID:   7214,
			Slug: "cryptowater",
			Name: "CryptoWater",
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
	"mAAPL": {
		{
			ID:   8001,
			Slug: "mirrored-apple",
			Name: "Mirrored Apple",
		},
	},
	"BTCL": {
		{
			ID:   3685,
			Slug: "btc-lite",
			Name: "BTC Lite",
		},
	},
	"SGE": {
		{
			ID:   10735,
			Slug: "society-of-galactic-exploration",
			Name: "SOCIETY OF GALACTIC EXPLORATION",
		},
	},
	"NEWOS": {
		{
			ID:   3110,
			Slug: "newstoken",
			Name: "NewsToken",
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
	"B20": {
		{
			ID:   8457,
			Slug: "b20",
			Name: "B20",
		},
	},
	"SVT": {
		{
			ID:   10493,
			Slug: "spacevikings",
			Name: "SpaceVikings",
		},
	},
	"ARRO": {
		{
			ID:   8841,
			Slug: "arro-social",
			Name: "Arro Social",
		},
	},
	"PRDX": {
		{
			ID:   6971,
			Slug: "predix-network",
			Name: "Predix Network",
		},
	},
	"JAGUAR": {
		{
			ID:   9610,
			Slug: "jaguarswap",
			Name: "JaguarSwap",
		},
	},
	"AOE": {
		{
			ID:   10415,
			Slug: "asset-of-empires",
			Name: "Asset of Empires",
		},
	},
	"TDX": {
		{
			ID:   2542,
			Slug: "tidex-token",
			Name: "Tidex Token",
		},
	},
	"CHND": {
		{
			ID:   5427,
			Slug: "cashhand",
			Name: "Cashhand",
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
	"ATUSD": {
		{
			ID:   5749,
			Slug: "aave-tusd",
			Name: "Aave TUSD",
		},
	},
	"RPL": {
		{
			ID:   2943,
			Slug: "rocket-pool",
			Name: "Rocket Pool",
		},
	},
	"CXO": {
		{
			ID:   2490,
			Slug: "cargox",
			Name: "CargoX",
		},
	},
	"TORN": {
		{
			ID:   8049,
			Slug: "torn",
			Name: "Tornado Cash",
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
	"XTA": {
		{
			ID:   3682,
			Slug: "italo",
			Name: "Italo",
		},
	},
	"BCR": {
		{
			ID:   5535,
			Slug: "bankcoin-reserve",
			Name: "Bankcoin Reserve",
		},
	},
	"BNP": {
		{
			ID:   4872,
			Slug: "benepit-protocol",
			Name: "BenePit Protocol",
		},
	},
	"AVA": {
		{
			ID:   2776,
			Slug: "travala",
			Name: "Travala.com",
		},
	},
	"XPM": {
		{
			ID:   42,
			Slug: "primecoin",
			Name: "Primecoin",
		},
	},
	"STQ": {
		{
			ID:   2541,
			Slug: "storiqa",
			Name: "Storiqa",
		},
	},
	"ENTS": {
		{
			ID:   3433,
			Slug: "eunomia",
			Name: "EUNOMIA",
		},
	},
	"PRISM": {
		{
			ID:   9788,
			Slug: "prism-network",
			Name: "Prism Network",
		},
	},
	"MAUSDC": {
		{
			ID:   8918,
			Slug: "matic-aave-usdc",
			Name: "Matic Aave Interest Bearing USDC",
		},
	},
	"CPAC": {
		{
			ID:   10263,
			Slug: "compact",
			Name: "Compact",
		},
	},
	"CHFT": {
		{
			ID:   5513,
			Slug: "crypto-holding-frank-token",
			Name: "Crypto Holding Frank Token",
		},
	},
	"KEL": {
		{
			ID:   9261,
			Slug: "kelvpn",
			Name: "KelVPN",
		},
	},
	"FLM": {
		{
			ID:   7150,
			Slug: "flamingo",
			Name: "Flamingo",
		},
	},
	"D": {
		{
			ID:   1769,
			Slug: "denarius-d",
			Name: "Denarius",
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
	"OKBBULL": {
		{
			ID:   6087,
			Slug: "3x-long-okb-token",
			Name: "3X Long OKB Token",
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
	"DEXE": {
		{
			ID:   7326,
			Slug: "dexe",
			Name: "DeXe",
		},
	},
	"COB": {
		{
			ID:   2006,
			Slug: "cobinhood",
			Name: "Cobinhood",
		},
	},
	"DEUS": {
		{
			ID:   7125,
			Slug: "deus-finance",
			Name: "DEUS Finance",
		},
	},
	"SMI": {
		{
			ID:   9958,
			Slug: "safemoon-inu",
			Name: "SafeMoon Inu",
		},
	},
	"UUU": {
		{
			ID:   2645,
			Slug: "u-network",
			Name: "U Network",
		},
	},
	"TRBT": {
		{
			ID:   7047,
			Slug: "tribute",
			Name: "Tribute",
		},
	},
	"PCX": {
		{
			ID:   4200,
			Slug: "chainx",
			Name: "ChainX",
		},
	},
	"BOOM": {
		{
			ID:   4128,
			Slug: "boom",
			Name: "BOOM",
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
	"SND": {
		{
			ID:   3523,
			Slug: "snodecoin",
			Name: "SnodeCoin",
		},
	},
	"XDC": {
		{
			ID:   2634,
			Slug: "xinfin-network",
			Name: "XinFin Network",
		},
	},
	"ELF": {
		{
			ID:   2299,
			Slug: "aelf",
			Name: "aelf",
		},
	},
	"RGT": {
		{
			ID:   7486,
			Slug: "rari-governance-token",
			Name: "Rari Governance Token",
		},
	},
	"ANY": {
		{
			ID:   5892,
			Slug: "anyswap",
			Name: "Anyswap",
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
	"SAFECOMET": {
		{
			ID:   9568,
			Slug: "safecomet",
			Name: "SafeComet",
		},
	},
	"WELL": {
		{
			ID:   7883,
			Slug: "well-token",
			Name: "WELL",
		},
	},
	"DOW": {
		{
			ID:   3223,
			Slug: "dowcoin",
			Name: "DOWCOIN",
		},
	},
	"CRL": {
		{
			ID:   8756,
			Slug: "coralfarm",
			Name: "CoralFarm",
		},
	},
	"KVI": {
		{
			ID:   6031,
			Slug: "kvi",
			Name: "KVI",
		},
	},
	"$AVATAR": {
		{
			ID:   10495,
			Slug: "avatar-moon",
			Name: "Avatar Moon",
		},
	},
	"SUKU": {
		{
			ID:   6180,
			Slug: "suku",
			Name: "SUKU",
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
	"SPT": {
		{
			ID:   3777,
			Slug: "spectrum",
			Name: "Spectrum",
		},
	},
	"FILDA": {
		{
			ID:   8426,
			Slug: "filda",
			Name: "Filda",
		},
	},
	"GDC": {
		{
			ID:   4678,
			Slug: "global-digital-content",
			Name: "Global Digital Content",
		},
	},
	"BTG": {
		{
			ID:   2083,
			Slug: "bitcoin-gold",
			Name: "Bitcoin Gold",
		},
	},
	"SYS": {
		{
			ID:   541,
			Slug: "syscoin",
			Name: "Syscoin",
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
	"ZOE": {
		{
			ID:   10959,
			Slug: "zoe-cash",
			Name: "Zoe Cash",
		},
	},
	"MCDC": {
		{
			ID:   8562,
			Slug: "mcdonalds-coin",
			Name: "McDonalds Coin",
		},
	},
	"ACC": {
		{
			ID:   4696,
			Slug: "asian-african-capital-chain",
			Name: "Asian-African Capital Chain",
		},
	},
	"LBD": {
		{
			ID:   8138,
			Slug: "linkbased",
			Name: "LinkBased",
		},
	},
	"CLC": {
		{
			ID:   4384,
			Slug: "caluracoin",
			Name: "CaluraCoin",
		},
	},
	"NIU": {
		{
			ID:   8748,
			Slug: "niubiswap",
			Name: "Niubi Swap",
		},
	},
	"CENT": {
		{
			ID:   3681,
			Slug: "centercoin",
			Name: "CENTERCOIN",
		},
	},
	"VERI": {
		{
			ID:   1710,
			Slug: "veritaseum",
			Name: "Veritaseum",
		},
	},
	"ICHI": {
		{
			ID:   7726,
			Slug: "ichi",
			Name: "ICHI",
		},
	},
	"TIX": {
		{
			ID:   1873,
			Slug: "blocktix",
			Name: "Blocktix",
		},
	},
	"DRK": {
		{
			ID:   10234,
			Slug: "draken",
			Name: "Draken",
		},
	},
	"UFC": {
		{
			ID:   5901,
			Slug: "union-fair-coin",
			Name: "Union Fair Coin",
		},
	},
	"EKT": {
		{
			ID:   2453,
			Slug: "educare",
			Name: "EDUCare",
		},
	},
	"WIKEN": {
		{
			ID:   4809,
			Slug: "project-with",
			Name: "Project WITH",
		},
	},
	"TOMOBULL": {
		{
			ID:   6090,
			Slug: "3x-long-tomochain-token",
			Name: "3X Long TomoChain Token",
		},
	},
	"BNBBEAR": {
		{
			ID:   5294,
			Slug: "3x-short-bnb-token",
			Name: "3X Short BNB Token",
		},
	},
	"xBTC": {
		{
			ID:   7168,
			Slug: "xbtc",
			Name: "xBTC",
		},
	},
	"BTT": {
		{
			ID:   3718,
			Slug: "bittorrent",
			Name: "BitTorrent",
		},
	},
	"VIDT": {
		{
			ID:   3845,
			Slug: "vidt-datalink",
			Name: "VIDT Datalink",
		},
	},
	"DTEP": {
		{
			ID:   4277,
			Slug: "decoin",
			Name: "DECOIN",
		},
	},
	"BIKI": {
		{
			ID:   5325,
			Slug: "biki",
			Name: "BIKI",
		},
	},
	"GRLC": {
		{
			ID:   2475,
			Slug: "garlicoin",
			Name: "Garlicoin",
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
	"RBLX": {
		{
			ID:   2689,
			Slug: "rublix",
			Name: "Rublix",
		},
	},
	"KISHU": {
		{
			ID:   9386,
			Slug: "kishu-inu",
			Name: "Kishu Inu",
		},
	},
	"CEUR": {
		{
			ID:   9467,
			Slug: "celo-euro",
			Name: "Celo Euro",
		},
	},
	"ETHBN": {
		{
			ID:   5826,
			Slug: "etherbone",
			Name: "EtherBone",
		},
	},
	"MIX": {
		{
			ID:   4366,
			Slug: "mixmarvel",
			Name: "MixMarvel",
		},
	},
	"BZNT": {
		{
			ID:   2727,
			Slug: "bezant",
			Name: "Bezant",
		},
	},
	"BECN": {
		{
			ID:   3656,
			Slug: "beacon",
			Name: "Beacon",
		},
	},
	"N1": {
		{
			ID:   9346,
			Slug: "nftify",
			Name: "NFTify",
		},
	},
	"XGS": {
		{
			ID:   3377,
			Slug: "genesisx",
			Name: "GenesisX",
		},
	},
	"EPRO": {
		{
			ID:   10607,
			Slug: "ethereum-pro",
			Name: "Ethereum Pro",
		},
	},
	"VBZRX": {
		{
			ID:   5879,
			Slug: "vbzrx",
			Name: "bZx Vesting Token",
		},
	},
	"CBR": {
		{
			ID:   5921,
			Slug: "cybercoin",
			Name: "Cybercoin",
		},
	},
	"MITX": {
		{
			ID:   2709,
			Slug: "morpheus-labs",
			Name: "Morpheus Labs",
		},
	},
	"SSX": {
		{
			ID:   5612,
			Slug: "somesing",
			Name: "SOMESING",
		},
	},
	"SLP": {
		{
			ID:   5824,
			Slug: "small-love-potion",
			Name: "Small Love Potion",
		},
	},
	"RAY": {
		{
			ID:   8526,
			Slug: "raydium",
			Name: "Raydium",
		},
	},
	"RAE": {
		{
			ID:   4887,
			Slug: "receive-access-ecosystem",
			Name: "Receive Access Ecosystem",
		},
	},
	"PACOCA": {
		{
			ID:   10522,
			Slug: "pacoca",
			Name: "Pacoca",
		},
	},
	"FORINT": {
		{
			ID:   10981,
			Slug: "forint-token",
			Name: "Forint Token",
		},
	},
	"AWG": {
		{
			ID:   5705,
			Slug: "aurusgold",
			Name: "AurusGOLD",
		},
	},
	"ESWA": {
		{
			ID:   6425,
			Slug: "easyswap",
			Name: "EasySwap",
		},
	},
	"SAKE": {
		{
			ID:   6997,
			Slug: "sake-token",
			Name: "SakeToken",
		},
	},
	"SKM": {
		{
			ID:   2725,
			Slug: "skrumble-network",
			Name: "Skrumble Network",
		},
	},
	"MNE": {
		{
			ID:   1673,
			Slug: "minereum",
			Name: "Minereum",
		},
	},
	"GVE": {
		{
			ID:   2969,
			Slug: "globalvillage-ecosystem",
			Name: "Globalvillage Ecosystem",
		},
	},
	"ZENI": {
		{
			ID:   1629,
			Slug: "zennies",
			Name: "Zennies",
		},
	},
	"SCAP": {
		{
			ID:   5002,
			Slug: "safecapital",
			Name: "SafeCapital",
		},
	},
	"BCZ": {
		{
			ID:   4008,
			Slug: "bitcoin-cz",
			Name: "Bitcoin CZ",
		},
	},
	"WATER": {
		{
			ID:   8860,
			Slug: "water-finance-token",
			Name: "WaterDefi",
		},
	},
	"DAWN": {
		{
			ID:   5618,
			Slug: "dawn-protocol",
			Name: "Dawn Protocol",
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
	"CLVA": {
		{
			ID:   8228,
			Slug: "clever-defi",
			Name: "Clever DeFi",
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
	"RENFIL": {
		{
			ID:   7997,
			Slug: "renfil",
			Name: "renFIL",
		},
	},
	"CBT": {
		{
			ID:   9342,
			Slug: "community-business-token",
			Name: "Community Business Token",
		},
	},
	"DROPS": {
		{
			ID:   9357,
			Slug: "defidrop-launchpad",
			Name: "DefiDrop Launchpad",
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
	"MAY": {
		{
			ID:   1693,
			Slug: "theresa-may-coin",
			Name: "Theresa May Coin",
		},
	},
	"MCO2": {
		{
			ID:   8826,
			Slug: "moss-carbon-credit",
			Name: "Moss Carbon Credit",
		},
	},
	"GLITCHY": {
		{
			ID:   10595,
			Slug: "glitchy",
			Name: "Glitchy",
		},
	},
	"YFIX": {
		{
			ID:   7166,
			Slug: "yfix-finance",
			Name: "YFIX Finance",
		},
	},
	"RON": {
		{
			ID:   10921,
			Slug: "rise-of-nebula",
			Name: "Rise Of Nebula",
		},
	},
	"WPX": {
		{
			ID:   5208,
			Slug: "wallet-plus-x",
			Name: "Wallet Plus X",
		},
	},
	"BTSC": {
		{
			ID:   6392,
			Slug: "bts-coin",
			Name: "BTS Coin",
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
	"NWC": {
		{
			ID:   4890,
			Slug: "newscrypto",
			Name: "Newscrypto",
		},
	},
	"DAD": {
		{
			ID:   4862,
			Slug: "dad",
			Name: "DAD",
		},
	},
	"IMPACT": {
		{
			ID:   10389,
			Slug: "alpha-impact",
			Name: "Alpha Impact",
		},
	},
	"vDOT": {
		{
			ID:   7976,
			Slug: "venus-dot",
			Name: "Venus DOT",
		},
	},
	"RAZE": {
		{
			ID:   9173,
			Slug: "raze-network",
			Name: "Raze Network",
		},
	},
	"EMD": {
		{
			ID:   50,
			Slug: "emerald",
			Name: "Emerald Crypto",
		},
	},
	"TERC": {
		{
			ID:   6508,
			Slug: "troneuroperewardcoin",
			Name: "TronEuropeRewardCoin",
		},
	},
	"TFF": {
		{
			ID:   8870,
			Slug: "tutti-frutti",
			Name: "Tutti Frutti",
		},
	},
	"ZPR": {
		{
			ID:   2972,
			Slug: "zper",
			Name: "ZPER",
		},
	},
	"BP": {
		{
			ID:   10904,
			Slug: "bunnypark",
			Name: "BunnyPark",
		},
	},
	"PDOGE": {
		{
			ID:   11037,
			Slug: "polkadoge",
			Name: "POLKADOGE",
		},
	},
	"DXC": {
		{
			ID:   5364,
			Slug: "mydexpay",
			Name: "Dexchain",
		},
	},
	"SMOKE": {
		{
			ID:   8503,
			Slug: "the-smokehouse",
			Name: "The Smokehouse",
		},
	},
	"KBOT": {
		{
			ID:   4404,
			Slug: "korbot",
			Name: "Korbot",
		},
	},
	"ZEON": {
		{
			ID:   3795,
			Slug: "zeon",
			Name: "ZEON",
		},
	},
	"USDK": {
		{
			ID:   4064,
			Slug: "usdk",
			Name: "USDK",
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
	"MNSTP": {
		{
			ID:   9213,
			Slug: "moon-stop",
			Name: "Moon Stop",
		},
	},
	"SAFEICARUS": {
		{
			ID:   9630,
			Slug: "safeicarus",
			Name: "Safeicarus",
		},
	},
	"SPRK": {
		{
			ID:   4010,
			Slug: "sparkster",
			Name: "Sparkster",
		},
	},
	"SAFECOM": {
		{
			ID:   9379,
			Slug: "safe-community-token",
			Name: "SAFE Community Token",
		},
	},
	"NAZ": {
		{
			ID:   7516,
			Slug: "naz-coin",
			Name: "Naz Coin",
		},
	},
	"LOG": {
		{
			ID:   859,
			Slug: "woodcoin",
			Name: "Woodcoin",
		},
	},
	"DVF": {
		{
			ID:   10759,
			Slug: "deversifi",
			Name: "DeversiFi",
		},
	},
	"KET": {
		{
			ID:   9431,
			Slug: "rowket-market",
			Name: "Rowket",
		},
	},
	"HLAND": {
		{
			ID:   8021,
			Slug: "hland-token",
			Name: "HLand Token",
		},
	},
	"MYFI": {
		{
			ID:   8953,
			Slug: "myfinance",
			Name: "MYFinance",
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
	"DOGEC": {
		{
			ID:   3672,
			Slug: "dogecash",
			Name: "DogeCash",
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
	"DUSD": {
		{
			ID:   6881,
			Slug: "defidollar",
			Name: "DefiDollar",
		},
	},
	"SMG": {
		{
			ID:   9491,
			Slug: "smaugs-nft",
			Name: "Smaugs NFT",
		},
	},
	"TNS": {
		{
			ID:   2704,
			Slug: "transcodium",
			Name: "Transcodium",
		},
	},
	"MEDIC": {
		{
			ID:   916,
			Slug: "mediccoin",
			Name: "MedicCoin",
		},
	},
	"YOOSHI": {
		{
			ID:   9892,
			Slug: "yooshi",
			Name: "YooShi",
		},
	},
	"HUNNY": {
		{
			ID:   10514,
			Slug: "pancake-hunny",
			Name: "PANCAKE HUNNY",
		},
	},
	"SWAMP": {
		{
			ID:   9082,
			Slug: "swampy",
			Name: "Swampy",
		},
	},
	"REDI": {
		{
			ID:   8388,
			Slug: "redi",
			Name: "REDi",
		},
	},
	"TRXBEAR": {
		{
			ID:   5379,
			Slug: "3x-short-trx-token",
			Name: "3X Short TRX Token",
		},
	},
	"DEXA": {
		{
			ID:   4917,
			Slug: "dexa-coin",
			Name: "DEXA COIN",
		},
	},
	"SFX": {
		{
			ID:   4183,
			Slug: "safex-cash",
			Name: "Safex Cash",
		},
	},
	"NFXC": {
		{
			ID:   5545,
			Slug: "nfx-coin",
			Name: "NFX Coin",
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
	"STORJ": {
		{
			ID:   1772,
			Slug: "storj",
			Name: "Storj",
		},
	},
	"AIT": {
		{
			ID:   2407,
			Slug: "aichain",
			Name: "AICHAIN",
		},
	},
	"CC": {
		{
			ID:   9564,
			Slug: "cryptocart",
			Name: "CryptoCart",
		},
	},
	"FNSP": {
		{
			ID:   7134,
			Slug: "finswap",
			Name: "Finswap",
		},
	},
	"BOOZE": {
		{
			ID:   9924,
			Slug: "boozemoon",
			Name: "BoozeMoon",
		},
	},
	"GNX": {
		{
			ID:   2291,
			Slug: "genaro-network",
			Name: "Genaro Network",
		},
	},
	"WRC": {
		{
			ID:   2288,
			Slug: "worldcore",
			Name: "Worldcore",
		},
	},
	"TAJ": {
		{
			ID:   1353,
			Slug: "tajcoin",
			Name: "TajCoin",
		},
	},
	"BVOL": {
		{
			ID:   5542,
			Slug: "1x-long-bitcoin-implied-volatility-token",
			Name: "1x Long Bitcoin Implied Volatility Token",
		},
	},
	"WCCX": {
		{
			ID:   7600,
			Slug: "wrapped-conceal",
			Name: "Wrapped Conceal",
		},
	},
	"NAUT": {
		{
			ID:   8770,
			Slug: "astronaut",
			Name: "Astronaut",
		},
	},
	"CHT": {
		{
			ID:   4701,
			Slug: "coinhe-token",
			Name: "CoinHe Token",
		},
	},
	"CELR": {
		{
			ID:   3814,
			Slug: "celer-network",
			Name: "Celer Network",
		},
	},
	"NAX": {
		{
			ID:   5873,
			Slug: "nextdao",
			Name: "NextDAO",
		},
	},
	"DDK": {
		{
			ID:   4180,
			Slug: "ddkoin",
			Name: "DDKoin",
		},
	},
	"CBM": {
		{
			ID:   4253,
			Slug: "cryptobonusmiles",
			Name: "CryptoBonusMiles",
		},
	},
	"BBOO": {
		{
			ID:   8350,
			Slug: "panda-yield",
			Name: "Panda Yield",
		},
	},
	"JOINT": {
		{
			ID:   2745,
			Slug: "joint-ventures",
			Name: "Joint Ventures",
		},
	},
	"HOO": {
		{
			ID:   7543,
			Slug: "hoo-token",
			Name: "Hoo Token",
		},
	},
	"SPA": {
		{
			ID:   6715,
			Slug: "sperax",
			Name: "Sperax",
		},
	},
	"BLOSM": {
		{
			ID:   9733,
			Slug: "blossomcoin",
			Name: "BlossomCoin",
		},
	},
	"CAPT": {
		{
			ID:   10226,
			Slug: "captain",
			Name: "Captain",
		},
	},
	"NIM": {
		{
			ID:   2916,
			Slug: "nimiq",
			Name: "Nimiq",
		},
	},
	"MTH": {
		{
			ID:   1947,
			Slug: "monetha",
			Name: "Monetha",
		},
	},
	"LEAD": {
		{
			ID:   6940,
			Slug: "lead-wallet",
			Name: "Lead Wallet",
		},
	},
	"SMTY": {
		{
			ID:   7594,
			Slug: "smoothy",
			Name: "Smoothy",
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
	"PYN": {
		{
			ID:   3323,
			Slug: "paycent",
			Name: "PAYCENT",
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
	"BOGE": {
		{
			ID:   10400,
			Slug: "bogecoin",
			Name: "Bogecoin",
		},
	},
	"SNOW": {
		{
			ID:   7367,
			Slug: "snowswap",
			Name: "SnowSwap",
		},
	},
	"GVY": {
		{
			ID:   7905,
			Slug: "groovy-finance",
			Name: "Groovy Finance",
		},
	},
	"BUGG": {
		{
			ID:   10479,
			Slug: "bugg-inu",
			Name: "Bugg Inu",
		},
	},
	"SUREBETS": {
		{
			ID:   10505,
			Slug: "surebets-online",
			Name: "SureBets Online",
		},
	},
	"KCASH": {
		{
			ID:   2379,
			Slug: "kcash",
			Name: "Kcash",
		},
	},
	"RMPL": {
		{
			ID:   6503,
			Slug: "rmpl",
			Name: "RMPL",
		},
	},
	"AUDT": {
		{
			ID:   8123,
			Slug: "australian-dollar-token",
			Name: "Australian Dollar Token",
		},
	},
	"XFUEL": {
		{
			ID:   6602,
			Slug: "xfuel",
			Name: "XFUEL",
		},
	},
	"UNCX": {
		{
			ID:   7664,
			Slug: "uncx",
			Name: "UniCrypt",
		},
	},
	"NAOS": {
		{
			ID:   9504,
			Slug: "naos-finance",
			Name: "NAOS Finance",
		},
	},
	"YTN": {
		{
			ID:   2290,
			Slug: "yenten",
			Name: "YENTEN",
		},
	},
	"BILL": {
		{
			ID:   10568,
			Slug: "bill-hwang-finance",
			Name: "Bill Hwang Finance",
		},
	},
	"CFX": {
		{
			ID:   7334,
			Slug: "conflux-network",
			Name: "Conflux Network",
		},
	},
	"MXC": {
		{
			ID:   3628,
			Slug: "mxc",
			Name: "MXC",
		},
	},
	"UDO": {
		{
			ID:   8679,
			Slug: "unido",
			Name: "Unido EP",
		},
	},
	"N0031": {
		{
			ID:   7575,
			Slug: "nyfi",
			Name: "nYFI",
		},
	},
	"YETH": {
		{
			ID:   7989,
			Slug: "fyeth-finance",
			Name: "fyeth.finance",
		},
	},
	"PCH": {
		{
			ID:   2902,
			Slug: "popchain",
			Name: "POPCHAIN",
		},
	},
	"NAFT": {
		{
			ID:   9828,
			Slug: "nafter",
			Name: "Nafter",
		},
	},
	"HDAC": {
		{
			ID:   2999,
			Slug: "hdac",
			Name: "Hdac",
		},
	},
	"BAN": {
		{
			ID:   4704,
			Slug: "banano",
			Name: "Banano",
		},
	},
	"NEWS": {
		{
			ID:   4647,
			Slug: "publish",
			Name: "PUBLISH",
		},
	},
	"MAHA": {
		{
			ID:   8043,
			Slug: "mahadao",
			Name: "MahaDAO",
		},
	},
	"ETX": {
		{
			ID:   3771,
			Slug: "ethereumx",
			Name: "EthereumX",
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
	"ESD": {
		{
			ID:   7033,
			Slug: "empty-set-dollar",
			Name: "Empty Set Dollar",
		},
	},
	"QUAN": {
		{
			ID:   3390,
			Slug: "quantis-network",
			Name: "Quantis Network",
		},
	},
	"BAO": {
		{
			ID:   8168,
			Slug: "bao-finance",
			Name: "Bao Finance",
		},
	},
	"WIX": {
		{
			ID:   3404,
			Slug: "wixlar",
			Name: "Wixlar",
		},
	},
	"POG": {
		{
			ID:   10658,
			Slug: "pogcoin",
			Name: "PogCoin",
		},
	},
	"TVK": {
		{
			ID:   8037,
			Slug: "terra-virtua-kolect",
			Name: "Terra Virtua Kolect",
		},
	},
	"UIP": {
		{
			ID:   2454,
			Slug: "unlimitedip",
			Name: "UnlimitedIP",
		},
	},
	"ZEFU": {
		{
			ID:   7430,
			Slug: "zenfuse",
			Name: "Zenfuse",
		},
	},
	"UTU": {
		{
			ID:   7587,
			Slug: "utu-protocol",
			Name: "UTU Protocol",
		},
	},
	"KIMCHI": {
		{
			ID:   6839,
			Slug: "kimchi-finance",
			Name: "KIMCHI.finance",
		},
	},
	"MOONS": {
		{
			ID:   7406,
			Slug: "moontools",
			Name: "MoonTools",
		},
	},
	"DCT": {
		{
			ID:   1478,
			Slug: "decent",
			Name: "DECENT",
		},
	},
	"KWH": {
		{
			ID:   3001,
			Slug: "kwhcoin",
			Name: "KWHCoin",
		},
	},
	"ULTRA": {
		{
			ID:   10195,
			Slug: "ultrasafe",
			Name: "Ultrasafe",
		},
	},
	"SPE": {
		{
			ID:   9303,
			Slug: "save-planet-earth",
			Name: "Save Planet Earth",
		},
	},
	"PDAO": {
		{
			ID:   8735,
			Slug: "panda-dao",
			Name: "Panda Dao",
		},
	},
	"WAYF": {
		{
			ID:   7700,
			Slug: "way-f-coin",
			Name: "WAY-F coin",
		},
	},
	"YFS": {
		{
			ID:   9181,
			Slug: "yfsfinance",
			Name: "YFS.FINANCE",
		},
	},
	"MEDI": {
		{
			ID:   8012,
			Slug: "mediconnectuk",
			Name: "MediconnectUk",
		},
	},
	"ATT": {
		{
			ID:   5600,
			Slug: "attila",
			Name: "Attila",
		},
	},
	"TXL": {
		{
			ID:   7024,
			Slug: "tixl-new",
			Name: "Tixl",
		},
	},
	"EXCC": {
		{
			ID:   4388,
			Slug: "exchange-coin",
			Name: "ExchangeCoin",
		},
	},
	"CIR": {
		{
			ID:   8174,
			Slug: "circleswap",
			Name: "CircleSwap",
		},
	},
	"COLD": {
		{
			ID:   9751,
			Slug: "cold-finance",
			Name: "COLD FINANCE",
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
	"BIDR": {
		{
			ID:   6855,
			Slug: "binance-idr",
			Name: "BIDR",
		},
	},
	"ASAFE": {
		{
			ID:   1439,
			Slug: "allsafe",
			Name: "AllSafe",
		},
	},
	"YFIUP": {
		{
			ID:   7452,
			Slug: "yfiup",
			Name: "YFIUP",
		},
	},
	"IIC": {
		{
			ID:   2741,
			Slug: "intelligent-investment-chain",
			Name: "Intelligent Investment Chain",
		},
	},
	"SWAPP": {
		{
			ID:   10783,
			Slug: "swapp-protocol",
			Name: "SWAPP Protocol",
		},
	},
	"SOP": {
		{
			ID:   2947,
			Slug: "sopay",
			Name: "SoPay",
		},
	},
	"DETO": {
		{
			ID:   9050,
			Slug: "delta-exchange-token",
			Name: "Delta Exchange Token",
		},
	},
	"SKY": {
		{
			ID:   1619,
			Slug: "skycoin",
			Name: "Skycoin",
		},
	},
	"EQZ": {
		{
			ID:   9083,
			Slug: "equalizer",
			Name: "Equalizer",
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
	"GENX": {
		{
			ID:   10753,
			Slug: "evodefi",
			Name: "Evodefi",
		},
	},
	"MRC": {
		{
			ID:   8415,
			Slug: "meroechain",
			Name: "MeroeChain",
		},
	},
	"JULB": {
		{
			ID:   8812,
			Slug: "justliquidity-binance",
			Name: "JustLiquidity Binance",
		},
	},
	"ETN": {
		{
			ID:   2137,
			Slug: "electroneum",
			Name: "Electroneum",
		},
	},
	"TRIBE": {
		{
			ID:   9025,
			Slug: "tribe",
			Name: "Tribe",
		},
	},
	"HGET": {
		{
			ID:   6949,
			Slug: "hedget",
			Name: "Hedget",
		},
	},
	"PNY": {
		{
			ID:   3481,
			Slug: "peony",
			Name: "Peony",
		},
	},
	"LINKA": {
		{
			ID:   4850,
			Slug: "linka",
			Name: "LINKA",
		},
	},
	"BTZC": {
		{
			ID:   4867,
			Slug: "beatzcoin",
			Name: "BeatzCoin",
		},
	},
	"COVA": {
		{
			ID:   3650,
			Slug: "cova",
			Name: "COVA",
		},
	},
	"TOTO": {
		{
			ID:   2960,
			Slug: "tourist-token",
			Name: "Tourist Token",
		},
	},
	"ETH2X-FLI": {
		{
			ID:   9789,
			Slug: "eth-2x-flexible-leverage-index",
			Name: "ETH 2x Flexible Leverage Index",
		},
	},
	"EOSBEAR": {
		{
			ID:   5415,
			Slug: "3x-short-eos-token",
			Name: "3x Short EOS Token",
		},
	},
	"CATBREAD": {
		{
			ID:   10887,
			Slug: "catbread",
			Name: "CatBread",
		},
	},
	"MNTIS": {
		{
			ID:   8196,
			Slug: "mantis",
			Name: "Mantis",
		},
	},
	"NFUP": {
		{
			ID:   8007,
			Slug: "natural-farm-union-protocol",
			Name: "Natural Farm Union Protocol",
		},
	},
	"ERC20": {
		{
			ID:   2165,
			Slug: "erc20",
			Name: "ERC20",
		},
	},
	"ARQ": {
		{
			ID:   3882,
			Slug: "arqma",
			Name: "Arqma",
		},
	},
	"YFIAG": {
		{
			ID:   7862,
			Slug: "yearnagnostic-finance",
			Name: "YearnAgnostic Finance",
		},
	},
	"RBY": {
		{
			ID:   215,
			Slug: "rubycoin",
			Name: "Rubycoin",
		},
	},
	"LVH": {
		{
			ID:   6495,
			Slug: "lovehearts",
			Name: "LoveHearts",
		},
	},
	"MITH": {
		{
			ID:   2608,
			Slug: "mithril",
			Name: "Mithril",
		},
	},
	"NUG": {
		{
			ID:   3092,
			Slug: "nuggets",
			Name: "Nuggets",
		},
	},
	"X42": {
		{
			ID:   4097,
			Slug: "x42-protocol",
			Name: "x42 Protocol",
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
	"ADABULL": {
		{
			ID:   6079,
			Slug: "3x-long-cardano-token",
			Name: "3X Long Cardano Token",
		},
	},
	"REFI": {
		{
			ID:   9065,
			Slug: "realfinance-network",
			Name: "Realfinance Network",
		},
	},
	"HINT": {
		{
			ID:   4254,
			Slug: "hintchain",
			Name: "Hintchain",
		},
	},
	"TAT": {
		{
			ID:   7537,
			Slug: "tatcoin",
			Name: "Tatcoin",
		},
	},
	"PRE": {
		{
			ID:   2245,
			Slug: "presearch",
			Name: "Presearch",
		},
	},
	"MFTU": {
		{
			ID:   3189,
			Slug: "mainstream-for-the-underground",
			Name: "Mainstream For The Underground",
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
	"TRI": {
		{
			ID:   7995,
			Slug: "trinity-protocol",
			Name: "Trinity Protocol",
		},
	},
	"ASNX": {
		{
			ID:   5752,
			Slug: "aave-snx",
			Name: "Aave SNX",
		},
	},
	"TREAT": {
		{
			ID:   8808,
			Slug: "treat-dao",
			Name: "Treat DAO",
		},
	},
	"KWIK": {
		{
			ID:   10101,
			Slug: "kwikswap",
			Name: "Kwikswap Protocol",
		},
	},
	"B2G": {
		{
			ID:   3684,
			Slug: "bitcoiin",
			Name: "Bitcoiin",
		},
	},
	"TKINU": {
		{
			ID:   10024,
			Slug: "tsuki-inu",
			Name: "Tsuki Inu",
		},
	},
	"UCR": {
		{
			ID:   7199,
			Slug: "ultra-clear",
			Name: "Ultra Clear",
		},
	},
	"XFIT": {
		{
			ID:   9217,
			Slug: "xfai",
			Name: "XFai",
		},
	},
	"YUKI": {
		{
			ID:   3208,
			Slug: "yuki",
			Name: "YUKI",
		},
	},
	"vDAI": {
		{
			ID:   8214,
			Slug: "venus-dai",
			Name: "Venus DAI",
		},
	},
	"MYB": {
		{
			ID:   1902,
			Slug: "mybit",
			Name: "MyBit",
		},
	},
	"WANUSDT": {
		{
			ID:   8657,
			Slug: "wanusdt",
			Name: "wanUSDT",
		},
	},
	"WINLAMBO": {
		{
			ID:   10472,
			Slug: "winlambo",
			Name: "Winlambo",
		},
	},
	"DEXM": {
		{
			ID:   8557,
			Slug: "dexmex",
			Name: "DexMex",
		},
	},
	"EVM": {
		{
			ID:   10682,
			Slug: "evermars",
			Name: "EverMars",
		},
	},
	"PLAT": {
		{
			ID:   3633,
			Slug: "bitguild-plat",
			Name: "BitGuild PLAT",
		},
	},
	"TVNT": {
		{
			ID:   3644,
			Slug: "travelnote",
			Name: "TravelNote",
		},
	},
	"NOLE": {
		{
			ID:   5849,
			Slug: "nolecoin",
			Name: "NoleCoin",
		},
	},
	"MKR": {
		{
			ID:   1518,
			Slug: "maker",
			Name: "Maker",
		},
	},
	"SINS": {
		{
			ID:   3366,
			Slug: "safeinsure",
			Name: "SafeInsure",
		},
	},
	"AGRS": {
		{
			ID:   1037,
			Slug: "agoras-tokens",
			Name: "Agoras Tokens",
		},
	},
	"GME": {
		{
			ID:   8342,
			Slug: "gamestop-tokenized-stock-ftx",
			Name: "GameStop tokenized stock FTX",
		},
	},
	"NAR": {
		{
			ID:   7831,
			Slug: "narwhalswap",
			Name: "Narwhalswap",
		},
	},
	"TEST": {
		{
			ID:   9911,
			Slug: "test-token",
			Name: "Test Token",
		},
	},
	"KIMJ": {
		{
			ID:   10164,
			Slug: "kimjongmoon",
			Name: "KimJongMoon",
		},
	},
	"STEEP": {
		{
			ID:   3395,
			Slug: "steepcoin",
			Name: "SteepCoin",
		},
	},
	"CAVA": {
		{
			ID:   10511,
			Slug: "cavapoo",
			Name: "Cavapoo",
		},
	},
	"CRV": {
		{
			ID:   6538,
			Slug: "curve-dao-token",
			Name: "Curve DAO Token",
		},
	},
	"CORX": {
		{
			ID:   7698,
			Slug: "corionx",
			Name: "CorionX",
		},
	},
	"OPS": {
		{
			ID:   10494,
			Slug: "octopus-protocol",
			Name: "Octopus Protocol",
		},
	},
	"GPO": {
		{
			ID:   6478,
			Slug: "galaxy-pool-coin",
			Name: "Galaxy Pool Coin",
		},
	},
	"MOBI": {
		{
			ID:   2429,
			Slug: "mobius",
			Name: "Mobius",
		},
	},
	"XRPDOWN": {
		{
			ID:   7002,
			Slug: "xrpdown",
			Name: "XRPDOWN",
		},
	},
	"HUGO": {
		{
			ID:   9282,
			Slug: "hugo-finance",
			Name: "Hugo Finance",
		},
	},
	"HCA": {
		{
			ID:   4994,
			Slug: "harcomia",
			Name: "Harcomia",
		},
	},
	"LINKETHPA": {
		{
			ID:   6140,
			Slug: "eth-link-price-action-candlestick-set",
			Name: "ETH/LINK Price Action Candlestick Set",
		},
	},
	"FACT": {
		{
			ID:   6811,
			Slug: "fee-active-collateral-token",
			Name: "Fee Active Collateral Token",
		},
	},
	"BPP": {
		{
			ID:   8126,
			Slug: "bitpower",
			Name: "Bitpower",
		},
	},
	"DBC": {
		{
			ID:   2316,
			Slug: "deepbrain-chain",
			Name: "DeepBrain Chain",
		},
	},
	"PIRL": {
		{
			ID:   2105,
			Slug: "pirl",
			Name: "Pirl",
		},
	},
	"UNIQ": {
		{
			ID:   9287,
			Slug: "uniqly",
			Name: "Uniqly",
		},
	},
	"FOMP": {
		{
			ID:   7936,
			Slug: "fompound",
			Name: "FOMPOUND",
		},
	},
	"mVIXY": {
		{
			ID:   8028,
			Slug: "mirrored-proshares-vix-short-term-futures-etf",
			Name: "Mirrored ProShares VIX",
		},
	},
	"BTDX": {
		{
			ID:   1381,
			Slug: "bitcloud",
			Name: "Bitcloud",
		},
	},
	"KP4R": {
		{
			ID:   7582,
			Slug: "keep4r",
			Name: "Keep4r",
		},
	},
	"SWFL": {
		{
			ID:   6799,
			Slug: "swapfolio",
			Name: "Swapfolio",
		},
	},
	"AWC": {
		{
			ID:   3667,
			Slug: "atomic-wallet-coin",
			Name: "Atomic Wallet Coin",
		},
	},
	"DIGI": {
		{
			ID:   9551,
			Slug: "digible",
			Name: "Digible",
		},
	},
	"ECOC": {
		{
			ID:   5477,
			Slug: "ecochain",
			Name: "ECOChain",
		},
	},
	"BNBX": {
		{
			ID:   10627,
			Slug: "bnbx-finance",
			Name: "BNBX Finance",
		},
	},
	"NANO": {
		{
			ID:   1567,
			Slug: "nano",
			Name: "Nano",
		},
	},
	"BTCZ": {
		{
			ID:   2041,
			Slug: "bitcoinz",
			Name: "BitcoinZ",
		},
	},
	"MODIC": {
		{
			ID:   6176,
			Slug: "modern-investment-coin",
			Name: "Modern Investment Coin",
		},
	},
	"BLTG": {
		{
			ID:   3627,
			Slug: "block-logic",
			Name: "Block-Logic",
		},
	},
	"FISH": {
		{
			ID:   10134,
			Slug: "polycat-finance",
			Name: "Polycat Finance",
		},
	},
	"HNC": {
		{
			ID:   1004,
			Slug: "hnccoin",
			Name: "HNC COIN",
		},
	},
	"SWISS": {
		{
			ID:   7714,
			Slug: "swiss-finance",
			Name: "swiss.finance",
		},
	},
	"MXF": {
		{
			ID:   9093,
			Slug: "mixty-finance",
			Name: "Mixty Finance",
		},
	},
	"GAPT": {
		{
			ID:   9887,
			Slug: "gaptt",
			Name: "Gaptt",
		},
	},
	"ISIKC": {
		{
			ID:   5468,
			Slug: "isiklar-coin",
			Name: "Isiklar Coin",
		},
	},
	"PNX": {
		{
			ID:   2205,
			Slug: "phantomx",
			Name: "Phantomx",
		},
	},
	"DEGEN": {
		{
			ID:   8699,
			Slug: "degen-index",
			Name: "DEGEN Index",
		},
	},
	"SASHIMI": {
		{
			ID:   6991,
			Slug: "sashimi",
			Name: "Sashimi",
		},
	},
	"TCGCOIN": {
		{
			ID:   9971,
			Slug: "tcgcoin",
			Name: "TCGcoin",
		},
	},
	"NMC": {
		{
			ID:   3,
			Slug: "namecoin",
			Name: "Namecoin",
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
	"FIN": {
		{
			ID:   7225,
			Slug: "definer",
			Name: "DeFiner",
		},
	},
	"RBOYS": {
		{
			ID:   10930,
			Slug: "rocket-boys",
			Name: "Rocket Boys",
		},
	},
	"PNODE": {
		{
			ID:   9657,
			Slug: "pinknode",
			Name: "Pinknode",
		},
	},
	"DEVA": {
		{
			ID:   8515,
			Slug: "deva-token",
			Name: "DEVA TOKEN",
		},
	},
	"XPB": {
		{
			ID:   8783,
			Slug: "transmute-protocol",
			Name: "Transmute Protocol",
		},
	},
	"SCRT": {
		{
			ID:   5604,
			Slug: "secret",
			Name: "Secret",
		},
	},
	"TMTG": {
		{
			ID:   3356,
			Slug: "the-midas-touch-gold",
			Name: "The Midas Touch Gold",
		},
	},
	"NBX": {
		{
			ID:   4297,
			Slug: "netbox-coin",
			Name: "Netbox Coin",
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
	"RUP": {
		{
			ID:   1799,
			Slug: "rupee",
			Name: "Rupee",
		},
	},
	"EFT": {
		{
			ID:   8851,
			Slug: "eft-finance",
			Name: "EFT.finance",
		},
	},
	"SWZL": {
		{
			ID:   5424,
			Slug: "swapzilla",
			Name: "Swapzilla",
		},
	},
	"ACED": {
		{
			ID:   2978,
			Slug: "aced",
			Name: "AceD",
		},
	},
	"CARMA": {
		{
			ID:   10637,
			Slug: "carma-coin",
			Name: "CARMA COIN",
		},
	},
	"OLE": {
		{
			ID:   10343,
			Slug: "olecoin",
			Name: "Olecoin",
		},
	},
	"BTRL": {
		{
			ID:   4174,
			Slug: "bitcoinregular",
			Name: "BitcoinRegular",
		},
	},
	"HUM": {
		{
			ID:   3600,
			Slug: "humanscape",
			Name: "Humanscape",
		},
	},
	"GRIN": {
		{
			ID:   3709,
			Slug: "grin",
			Name: "Grin",
		},
	},
	"UDOO": {
		{
			ID:   3608,
			Slug: "hyprr",
			Name: "Howdoo",
		},
	},
	"SNGLS": {
		{
			ID:   1409,
			Slug: "singulardtv",
			Name: "SingularDTV",
		},
	},
	"MLM": {
		{
			ID:   2516,
			Slug: "mktcoin",
			Name: "MktCoin",
		},
	},
	"JFC": {
		{
			ID:   4568,
			Slug: "jfin",
			Name: "JFIN",
		},
	},
	"LCC": {
		{
			ID:   2540,
			Slug: "litecoin-cash",
			Name: "Litecoin Cash",
		},
	},
	"NFTB": {
		{
			ID:   9545,
			Slug: "nftb",
			Name: "NFTb",
		},
	},
	"NIX": {
		{
			ID:   2991,
			Slug: "nix",
			Name: "NIX",
		},
	},
	"LION": {
		{
			ID:   9598,
			Slug: "lion-token",
			Name: "Lion Token",
		},
	},
	"4STC": {
		{
			ID:   10488,
			Slug: "4-stock",
			Name: "4-Stock",
		},
	},
	"EAURIC": {
		{
			ID:   7758,
			Slug: "eauric",
			Name: "Eauric",
		},
	},
	"STZEN": {
		{
			ID:   8945,
			Slug: "stakedzen",
			Name: "StakedZEN",
		},
	},
	"XMCT": {
		{
			ID:   2923,
			Slug: "xmct",
			Name: "XMCT",
		},
	},
	"DVG": {
		{
			ID:   8478,
			Slug: "daoventures",
			Name: "DAOventures",
		},
	},
	"KCCPAD": {
		{
			ID:   10784,
			Slug: "kccpad",
			Name: "KCCPAD",
		},
	},
	"DX": {
		{
			ID:   3139,
			Slug: "dxchain-token",
			Name: "DxChain Token",
		},
	},
	"GBYTE": {
		{
			ID:   1492,
			Slug: "obyte",
			Name: "Obyte",
		},
	},
	"EXMR": {
		{
			ID:   2990,
			Slug: "exmr-fdn",
			Name: "EXMR FDN",
		},
	},
	"FDO": {
		{
			ID:   8176,
			Slug: "firdaos",
			Name: "Firdaos",
		},
	},
	"KOBE": {
		{
			ID:   8189,
			Slug: "shabu-shabu-finance",
			Name: "Shabu Shabu Finance",
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
	"EYES": {
		{
			ID:   5673,
			Slug: "eyes-protocol",
			Name: "EYES Protocol",
		},
	},
	"NOKN": {
		{
			ID:   5801,
			Slug: "nokencoin",
			Name: "Nokencoin",
		},
	},
	"ZEN": {
		{
			ID:   1698,
			Slug: "horizen",
			Name: "Horizen",
		},
	},
	"ACH": {
		{
			ID:   6958,
			Slug: "alchemy-pay",
			Name: "Alchemy Pay",
		},
	},
	"BITG": {
		{
			ID:   2604,
			Slug: "bitgreen",
			Name: "BitGreen",
		},
	},
	"KIT": {
		{
			ID:   7739,
			Slug: "dexkit",
			Name: "DexKit",
		},
	},
	"LMY": {
		{
			ID:   4577,
			Slug: "lunchmoney",
			Name: "LunchMoney",
		},
	},
	"CXC": {
		{
			ID:   4955,
			Slug: "capital-x-cell",
			Name: "CAPITAL X CELL",
		},
	},
	"JINDOGE": {
		{
			ID:   9701,
			Slug: "jindoge",
			Name: "Jindoge",
		},
	},
	"GALA": {
		{
			ID:   7080,
			Slug: "gala",
			Name: "Gala",
		},
	},
	"SOTA": {
		{
			ID:   8576,
			Slug: "sota-finance",
			Name: "SOTA Finance",
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
	"GFUN": {
		{
			ID:   3776,
			Slug: "goldfund",
			Name: "GoldFund",
		},
	},
	"MAI": {
		{
			ID:   10533,
			Slug: "mindsync",
			Name: "Mindsync",
		},
	},
	"WAXP": {
		{
			ID:   2300,
			Slug: "wax",
			Name: "WAX",
		},
	},
	"EUNO": {
		{
			ID:   3071,
			Slug: "euno",
			Name: "EUNO",
		},
	},
	"CELEB": {
		{
			ID:   8788,
			Slug: "celebplus",
			Name: "CELEBPLUS",
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
	"LINKPT": {
		{
			ID:   6159,
			Slug: "link-profit-taker-set",
			Name: "LINK Profit Taker Set",
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
	"MRF": {
		{
			ID:   10287,
			Slug: "moonradar",
			Name: "MoonRadar",
		},
	},
	"ADAX": {
		{
			ID:   10833,
			Slug: "adax",
			Name: "ADAX",
		},
	},
	"ETCBEAR": {
		{
			ID:   6099,
			Slug: "3x-short-ethereum-classic-token",
			Name: "3X Short Ethereum Classic Token",
		},
	},
	"FSC": {
		{
			ID:   5343,
			Slug: "five-star-coin",
			Name: "Five Star Coin",
		},
	},
	"ARDX": {
		{
			ID:   4985,
			Slug: "ardcoin",
			Name: "ArdCoin",
		},
	},
	"GSWAP": {
		{
			ID:   7588,
			Slug: "gameswap",
			Name: "Gameswap",
		},
	},
	"BLU": {
		{
			ID:   290,
			Slug: "bluecoin",
			Name: "BlueCoin",
		},
	},
	"ANV": {
		{
			ID:   7705,
			Slug: "aniverse",
			Name: "ANIVERSE",
		},
	},
	"YFB2": {
		{
			ID:   7612,
			Slug: "yearn-finance-bit2",
			Name: "Yearn Finance Bit2",
		},
	},
	"RUGBUST": {
		{
			ID:   10247,
			Slug: "rug-busters",
			Name: "Rug Busters",
		},
	},
	"GTO": {
		{
			ID:   2289,
			Slug: "gifto",
			Name: "Gifto",
		},
	},
	"AZUKI": {
		{
			ID:   7647,
			Slug: "azuki",
			Name: "Azuki",
		},
	},
	"EURXB": {
		{
			ID:   9136,
			Slug: "eurxb",
			Name: "EURxb",
		},
	},
	"VITE": {
		{
			ID:   2937,
			Slug: "vite",
			Name: "VITE",
		},
	},
	"RVN": {
		{
			ID:   2577,
			Slug: "ravencoin",
			Name: "Ravencoin",
		},
	},
	"AREPA": {
		{
			ID:   3133,
			Slug: "arepacoin",
			Name: "Arepacoin",
		},
	},
	"DLC": {
		{
			ID:   1395,
			Slug: "dollarcoin",
			Name: "Dollarcoin",
		},
	},
	"MST": {
		{
			ID:   1396,
			Slug: "mustangcoin",
			Name: "MustangCoin",
		},
	},
	"BTCDOWN": {
		{
			ID:   5609,
			Slug: "btcdown",
			Name: "BTCDOWN",
		},
	},
	"HRD": {
		{
			ID:   7801,
			Slug: "hrdcoin",
			Name: "HRDCOIN",
		},
	},
	"IDH": {
		{
			ID:   2459,
			Slug: "indahash",
			Name: "indaHash",
		},
	},
	"XRUNE": {
		{
			ID:   10260,
			Slug: "thorstarter",
			Name: "Thorstarter",
		},
	},
	"ΤDOGE": {
		{
			ID:   10181,
			Slug: "tdoge",
			Name: "τDoge",
		},
	},
	"FOLD": {
		{
			ID:   10182,
			Slug: "manifold-finance",
			Name: "Manifold Finance",
		},
	},
	"REP": {
		{
			ID:   1104,
			Slug: "augur",
			Name: "Augur",
		},
	},
	"KOK": {
		{
			ID:   5185,
			Slug: "keystone-of-opportunity-knowledge",
			Name: "KOK",
		},
	},
	"DAC": {
		{
			ID:   3154,
			Slug: "davinci-coin",
			Name: "Davinci Coin",
		},
	},
	"AC": {
		{
			ID:   7382,
			Slug: "acoconut",
			Name: "ACoconut",
		},
	},
	"NEC": {
		{
			ID:   2538,
			Slug: "nectar",
			Name: "Nectar",
		},
	},
	"OPM": {
		{
			ID:   6888,
			Slug: "omega-protocol-money",
			Name: "Omega Protocol Money",
		},
	},
	"KAN": {
		{
			ID:   2934,
			Slug: "bitkan",
			Name: "BitKan",
		},
	},
	"WOOP": {
		{
			ID:   8937,
			Slug: "woonkly-power",
			Name: "Woonkly Power",
		},
	},
	"PKT": {
		{
			ID:   2279,
			Slug: "playkey",
			Name: "Playkey",
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
	"BELA": {
		{
			ID:   217,
			Slug: "belacoin",
			Name: "Bela",
		},
	},
	"BSKT": {
		{
			ID:   8733,
			Slug: "basketcoin",
			Name: "BasketCoin",
		},
	},
	"COL": {
		{
			ID:   5890,
			Slug: "unit-protocol",
			Name: "Unit Protocol",
		},
	},
	"CCAI": {
		{
			ID:   10935,
			Slug: "cryptocurrencies-ai",
			Name: "Cryptocurrencies.ai",
		},
	},
	"DNT": {
		{
			ID:   1856,
			Slug: "district0x",
			Name: "district0x",
		},
	},
	"AMN": {
		{
			ID:   2705,
			Slug: "amon",
			Name: "Amon",
		},
	},
	"KP3RB": {
		{
			ID:   8134,
			Slug: "keep3r-bsc-network",
			Name: "Keep3r BSC Network",
		},
	},
	"LTCBEAR": {
		{
			ID:   5461,
			Slug: "3x-short-litecoin-token",
			Name: "3x Short Litecoin Token",
		},
	},
	"XC": {
		{
			ID:   9682,
			Slug: "xcom",
			Name: "XCOM",
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
	"GHD": {
		{
			ID:   7126,
			Slug: "giftedhands",
			Name: "Giftedhands",
		},
	},
	"NTRN": {
		{
			ID:   894,
			Slug: "neutron",
			Name: "Neutron",
		},
	},
	"CRAZYTIME": {
		{
			ID:   10582,
			Slug: "crazytime",
			Name: "CrazyTime",
		},
	},
	"EJECT": {
		{
			ID:   10417,
			Slug: "eject",
			Name: "Eject",
		},
	},
	"ARDR": {
		{
			ID:   1320,
			Slug: "ardor",
			Name: "Ardor",
		},
	},
	"TATA": {
		{
			ID:   9972,
			Slug: "hakunamatata",
			Name: "HakunaMatata",
		},
	},
	"UTNP": {
		{
			ID:   2524,
			Slug: "universa",
			Name: "Universa",
		},
	},
	"VANY": {
		{
			ID:   5554,
			Slug: "vanywhere",
			Name: "Vanywhere",
		},
	},
	"MWBTC": {
		{
			ID:   8504,
			Slug: "metawhale-btc",
			Name: "MetaWhale BTC",
		},
	},
	"DAO": {
		{
			ID:   8420,
			Slug: "dao-maker",
			Name: "DAO Maker",
		},
	},
	"RVF": {
		{
			ID:   9176,
			Slug: "rocket-vault",
			Name: "Rocket Vault",
		},
	},
	"SHMN": {
		{
			ID:   3480,
			Slug: "stronghands-masternode",
			Name: "StrongHands Masternode",
		},
	},
	"KAU": {
		{
			ID:   7346,
			Slug: "kauri-crypto",
			Name: "Kauri",
		},
	},
	"MERO": {
		{
			ID:   9059,
			Slug: "mero",
			Name: "Mero",
		},
	},
	"AMKR": {
		{
			ID:   5753,
			Slug: "aave-mkr",
			Name: "Aave MKR",
		},
	},
	"LTC": {
		{
			ID:   2,
			Slug: "litecoin",
			Name: "Litecoin",
		},
	},
	"OCN": {
		{
			ID:   2458,
			Slug: "odyssey",
			Name: "Odyssey",
		},
	},
	"MOAC": {
		{
			ID:   2403,
			Slug: "moac",
			Name: "MOAC",
		},
	},
	"FOTA": {
		{
			ID:   2472,
			Slug: "fortuna",
			Name: "Fortuna",
		},
	},
	"B21": {
		{
			ID:   8060,
			Slug: "b21-invest",
			Name: "B21 Invest",
		},
	},
	"NANODOGE": {
		{
			ID:   10916,
			Slug: "nano-doge-token",
			Name: "Nano Doge Token",
		},
	},
	"DENA": {
		{
			ID:   8761,
			Slug: "decentralized-nations",
			Name: "Decentralized Nations",
		},
	},
	"PYPL": {
		{
			ID:   7901,
			Slug: "paypal-tokenized-stock-ftx",
			Name: "PayPal tokenized stock FTX",
		},
	},
	"PZAP": {
		{
			ID:   9896,
			Slug: "polyzap-finance",
			Name: "PolyZap Finance",
		},
	},
	"TWT": {
		{
			ID:   5964,
			Slug: "trust-wallet-token",
			Name: "Trust Wallet Token",
		},
	},
	"XWP": {
		{
			ID:   3878,
			Slug: "swap",
			Name: "Swap",
		},
	},
	"OURO": {
		{
			ID:   4970,
			Slug: "ouroboros",
			Name: "Ouroboros",
		},
	},
	"AINU": {
		{
			ID:   10877,
			Slug: "ainu-token",
			Name: "Ainu Token",
		},
	},
	"KSM": {
		{
			ID:   5034,
			Slug: "kusama",
			Name: "Kusama",
		},
	},
	"SFI": {
		{
			ID:   7617,
			Slug: "saffron-finance",
			Name: "saffron.finance",
		},
	},
	"VRT": {
		{
			ID:   9691,
			Slug: "venus-reward-token",
			Name: "Venus Reward Token",
		},
	},
	"SCP": {
		{
			ID:   4074,
			Slug: "scprime",
			Name: "ScPrime",
		},
	},
	"LOUVRE": {
		{
			ID:   10782,
			Slug: "louvre-finance",
			Name: "Louvre Finance",
		},
	},
	"PASTA": {
		{
			ID:   8676,
			Slug: "pasta-finance",
			Name: "Pasta Finance",
		},
	},
	"ALY": {
		{
			ID:   5011,
			Slug: "ally",
			Name: "ALLY",
		},
	},
	"FMTA": {
		{
			ID:   7560,
			Slug: "fundamenta",
			Name: "Fundamenta",
		},
	},
	"GAINS": {
		{
			ID:   9125,
			Slug: "gains-associates",
			Name: "Gains Associates",
		},
	},
	"TOK": {
		{
			ID:   3692,
			Slug: "tokok",
			Name: "TOKOK",
		},
	},
	"CTL": {
		{
			ID:   1273,
			Slug: "citadel",
			Name: "Citadel",
		},
	},
	"PINKM": {
		{
			ID:   9596,
			Slug: "pinkmoon",
			Name: "PinkMoon",
		},
	},
	"AUX": {
		{
			ID:   3362,
			Slug: "auxilium",
			Name: "Auxilium",
		},
	},
	"XPAT": {
		{
			ID:   3112,
			Slug: "bitnation",
			Name: "Bitnation",
		},
	},
	"KEBAB": {
		{
			ID:   8334,
			Slug: "kebab-token",
			Name: "Kebab Token",
		},
	},
	"DMME": {
		{
			ID:   4819,
			Slug: "dmme",
			Name: "DMme",
		},
	},
	"HAPY": {
		{
			ID:   6105,
			Slug: "hapy-coin",
			Name: "HAPY Coin",
		},
	},
	"CHAD": {
		{
			ID:   9942,
			Slug: "the-chad-project",
			Name: "The Chad Token",
		},
	},
	"DUMPDOGE": {
		{
			ID:   10892,
			Slug: "dump-doge",
			Name: "DUMP DOGE",
		},
	},
	"PAT": {
		{
			ID:   2759,
			Slug: "patron",
			Name: "Patron",
		},
	},
	"MPG": {
		{
			ID:   3758,
			Slug: "max-property-group",
			Name: "Max Property Group",
		},
	},
	"QBT": {
		{
			ID:   2242,
			Slug: "qbao",
			Name: "Qbao",
		},
	},
	"BZX": {
		{
			ID:   3497,
			Slug: "bitcoin-zero",
			Name: "Bitcoin Zero",
		},
	},
	"PAYX": {
		{
			ID:   2191,
			Slug: "paypex",
			Name: "Paypex",
		},
	},
	"NYEX": {
		{
			ID:   3129,
			Slug: "nyerium",
			Name: "Nyerium",
		},
	},
	"IOI": {
		{
			ID:   10295,
			Slug: "trade-race-manager",
			Name: "IOI Token",
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
	"YFFI": {
		{
			ID:   6145,
			Slug: "yffi-finance",
			Name: "yffi finance",
		},
	},
	"VENA": {
		{
			ID:   4495,
			Slug: "vena",
			Name: "VENA",
		},
	},
	"SPDR": {
		{
			ID:   8002,
			Slug: "spiderdao",
			Name: "SpiderDAO",
		},
	},
	"KAREN": {
		{
			ID:   10378,
			Slug: "karencoin",
			Name: "KarenCoin",
		},
	},
	"SAFEMOON": {
		{
			ID:   8757,
			Slug: "safemoon",
			Name: "SafeMoon",
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
	"DINU": {
		{
			ID:   10021,
			Slug: "dogey-inu",
			Name: "Dogey-Inu",
		},
	},
	"BKS": {
		{
			ID:   5685,
			Slug: "barkis-network",
			Name: "Barkis Network",
		},
	},
	"POT": {
		{
			ID:   122,
			Slug: "potcoin",
			Name: "PotCoin",
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
	"SATOZ": {
		{
			ID:   9241,
			Slug: "satozhi",
			Name: "Satozhi",
		},
	},
	"MERL": {
		{
			ID:   9853,
			Slug: "merlin",
			Name: "Merlin",
		},
	},
	"NOR": {
		{
			ID:   3611,
			Slug: "noir",
			Name: "Noir",
		},
	},
	"AKA": {
		{
			ID:   3151,
			Slug: "akroma",
			Name: "Akroma",
		},
	},
	"SPP": {
		{
			ID:   10483,
			Slug: "shapepay",
			Name: "ShapePay",
		},
	},
	"CROX": {
		{
			ID:   9548,
			Slug: "croxswap",
			Name: "CroxSwap",
		},
	},
	"QNT": {
		{
			ID:   3155,
			Slug: "quant",
			Name: "Quant",
		},
	},
	"XTZBEAR": {
		{
			ID:   5463,
			Slug: "3x-short-tezos-token",
			Name: "3x Short Tezos Token",
		},
	},
	"RISKMOON": {
		{
			ID:   9278,
			Slug: "riskmoon",
			Name: "RiskMoon",
		},
	},
	"MTI": {
		{
			ID:   7615,
			Slug: "mti-finance",
			Name: "MTI Finance",
		},
	},
	"ORBYT": {
		{
			ID:   6559,
			Slug: "orbyt-token",
			Name: "ORBYT Token",
		},
	},
	"ELA": {
		{
			ID:   2492,
			Slug: "elastos",
			Name: "Elastos",
		},
	},
	"CRON": {
		{
			ID:   4309,
			Slug: "cryptocean",
			Name: "Cryptocean",
		},
	},
	"ASP": {
		{
			ID:   7536,
			Slug: "aspire",
			Name: "Aspire",
		},
	},
	"WENLAMBO": {
		{
			ID:   9898,
			Slug: "wenlambo",
			Name: "Wenlambo",
		},
	},
	"YEC": {
		{
			ID:   4160,
			Slug: "ycash",
			Name: "Ycash",
		},
	},
	"MASQ": {
		{
			ID:   8376,
			Slug: "masq",
			Name: "MASQ",
		},
	},
	"POLIS": {
		{
			ID:   2359,
			Slug: "polis",
			Name: "Polis",
		},
	},
	"XUEZ": {
		{
			ID:   3798,
			Slug: "xuez",
			Name: "Xuez",
		},
	},
	"ETHIX": {
		{
			ID:   8442,
			Slug: "ethichub",
			Name: "EthicHub",
		},
	},
	"PHR": {
		{
			ID:   2158,
			Slug: "phore",
			Name: "Phore",
		},
	},
	"KUBO": {
		{
			ID:   3901,
			Slug: "kubocoin",
			Name: "KuboCoin",
		},
	},
	"SWT": {
		{
			ID:   1562,
			Slug: "swarm-city",
			Name: "Swarm City",
		},
	},
	"HMNG": {
		{
			ID:   9737,
			Slug: "hummingbird-finance",
			Name: "HummingBird Finance",
		},
	},
	"dART": {
		{
			ID:   9098,
			Slug: "dart-insurance",
			Name: "dART Insurance",
		},
	},
	"ODDZ": {
		{
			ID:   8717,
			Slug: "oddz",
			Name: "Oddz",
		},
	},
	"ONC": {
		{
			ID:   8159,
			Slug: "one-cash",
			Name: "One Cash",
		},
	},
	"XYZ": {
		{
			ID:   10979,
			Slug: "universe-xyz",
			Name: "Universe.XYZ",
		},
	},
	"GNC": {
		{
			ID:   7447,
			Slug: "galaxy-network",
			Name: "GALAXY NETWORK",
		},
	},
	"ELEPHANT": {
		{
			ID:   9936,
			Slug: "elephant-money",
			Name: "Elephant Money",
		},
	},
	"MCRN": {
		{
			ID:   8880,
			Slug: "macaronswap",
			Name: "MacaronSwap",
		},
	},
	"3Cs": {
		{
			ID:   5907,
			Slug: "crypto-cricket-club",
			Name: "Crypto Cricket Club",
		},
	},
	"BZ": {
		{
			ID:   2918,
			Slug: "bit-z-token",
			Name: "BitZ Token",
		},
	},
	"VISION": {
		{
			ID:   8045,
			Slug: "apy-vision",
			Name: "APY Vision",
		},
	},
	"MNS": {
		{
			ID:   5702,
			Slug: "monnos",
			Name: "MONNOS",
		},
	},
	"MACH": {
		{
			ID:   4984,
			Slug: "mach-project",
			Name: "MACH Project",
		},
	},
	"ABX": {
		{
			ID:   3179,
			Slug: "arbidex",
			Name: "Arbidex",
		},
	},
	"AGLT": {
		{
			ID:   3427,
			Slug: "agrolot",
			Name: "Agrolot",
		},
	},
	"CTSI": {
		{
			ID:   5444,
			Slug: "cartesi",
			Name: "Cartesi",
		},
	},
	"PEFI": {
		{
			ID:   10818,
			Slug: "penguin-finance",
			Name: "Penguin Finance",
		},
	},
	"FRA": {
		{
			ID:   4249,
			Slug: "findora",
			Name: "Findora",
		},
	},
	"ACSI": {
		{
			ID:   10797,
			Slug: "acryptosi",
			Name: "ACryptoSI",
		},
	},
	"MYX": {
		{
			ID:   6594,
			Slug: "myx-network",
			Name: "MYX Network",
		},
	},
	"PXT": {
		{
			ID:   4192,
			Slug: "populous-xbrl-token",
			Name: "Populous XBRL Token",
		},
	},
	"DESH": {
		{
			ID:   5808,
			Slug: "decash",
			Name: "DeCash",
		},
	},
	"NYE": {
		{
			ID:   4268,
			Slug: "newyork-exchange",
			Name: "NewYork Exchange",
		},
	},
	"FUND": {
		{
			ID:   3854,
			Slug: "unification",
			Name: "Unification",
		},
	},
	"WETH": {
		{
			ID:   2396,
			Slug: "weth",
			Name: "WETH",
		},
	},
	"DAWG": {
		{
			ID:   10912,
			Slug: "inumaki",
			Name: "Inumaki",
		},
	},
	"EQMT": {
		{
			ID:   6572,
			Slug: "equus-mining-token",
			Name: "Equus Mining Token",
		},
	},
	"XCP": {
		{
			ID:   132,
			Slug: "counterparty",
			Name: "Counterparty",
		},
	},
	"RDN": {
		{
			ID:   2161,
			Slug: "raiden-network-token",
			Name: "Raiden Network Token",
		},
	},
	"PUSH": {
		{
			ID:   9111,
			Slug: "epns",
			Name: "Ethereum Push Notification Service",
		},
	},
	"EARNX": {
		{
			ID:   9389,
			Slug: "earnx",
			Name: "EarnX",
		},
	},
	"ETZ": {
		{
			ID:   2843,
			Slug: "ether-zero",
			Name: "Ether Zero",
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
	"PIPL": {
		{
			ID:   2056,
			Slug: "piplcoin",
			Name: "PiplCoin",
		},
	},
	"POKE": {
		{
			ID:   8303,
			Slug: "pokeball",
			Name: "Pokeball",
		},
	},
	"CTCN": {
		{
			ID:   5313,
			Slug: "contracoin",
			Name: "CONTRACOIN",
		},
	},
	"NSD": {
		{
			ID:   3200,
			Slug: "nasdacoin",
			Name: "Nasdacoin",
		},
	},
	"ETHDOWN": {
		{
			ID:   10853,
			Slug: "eth-down",
			Name: "ETHDOWN",
		},
	},
	"IDK": {
		{
			ID:   5580,
			Slug: "idk",
			Name: "IDK",
		},
	},
	"MEGA": {
		{
			ID:   7540,
			Slug: "megacryptopolis",
			Name: "MegaCryptoPolis",
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
	"DMB": {
		{
			ID:   1687,
			Slug: "digital-money-bits",
			Name: "Digital Money Bits",
		},
	},
	"RPEPE": {
		{
			ID:   8901,
			Slug: "rare-pepe",
			Name: "Rare Pepe",
		},
	},
	"DITTO": {
		{
			ID:   8086,
			Slug: "ditto",
			Name: "Ditto",
		},
	},
	"DMOD": {
		{
			ID:   9713,
			Slug: "demodyfi",
			Name: "Demodyfi",
		},
	},
	"LIQ": {
		{
			ID:   11013,
			Slug: "liq-protocol",
			Name: "LIQ Protocol",
		},
	},
	"POP!": {
		{
			ID:   9520,
			Slug: "pop",
			Name: "POP",
		},
	},
	"HTA": {
		{
			ID:   5365,
			Slug: "historia",
			Name: "Historia",
		},
	},
	"ENTONE": {
		{
			ID:   5811,
			Slug: "entone",
			Name: "ENTONE",
		},
	},
	"NDR": {
		{
			ID:   7993,
			Slug: "node-runners",
			Name: "Node Runners",
		},
	},
	"BSTY": {
		{
			ID:   644,
			Slug: "globalboost-y",
			Name: "GlobalBoost-Y",
		},
	},
	"INFTEE": {
		{
			ID:   11040,
			Slug: "infinitee-finance",
			Name: "Infinitee Finance",
		},
	},
	"BBYXRP": {
		{
			ID:   10872,
			Slug: "babyxrp",
			Name: "BABYXRP",
		},
	},
	"BFLY": {
		{
			ID:   8405,
			Slug: "butterfly-protocol-2",
			Name: "Butterfly Protocol",
		},
	},
	"ACXT": {
		{
			ID:   7854,
			Slug: "acdx-exchange-token",
			Name: "ACDX Exchange Governance Token",
		},
	},
	"PAND": {
		{
			ID:   10014,
			Slug: "panda-finance",
			Name: "Panda Finance",
		},
	},
	"COSBY": {
		{
			ID:   10760,
			Slug: "the-cosby-token",
			Name: "The Cosby Token",
		},
	},
	"NEAR": {
		{
			ID:   6535,
			Slug: "near-protocol",
			Name: "NEAR Protocol",
		},
	},
	"AKRO": {
		{
			ID:   4134,
			Slug: "akropolis",
			Name: "Akropolis",
		},
	},
	"PNG": {
		{
			ID:   8422,
			Slug: "pangolin",
			Name: "Pangolin",
		},
	},
	"TRUE": {
		{
			ID:   2457,
			Slug: "truechain",
			Name: "TrueChain",
		},
	},
	"URQA": {
		{
			ID:   9053,
			Slug: "ureeqa",
			Name: "UREEQA",
		},
	},
	"LANA": {
		{
			ID:   1257,
			Slug: "lanacoin",
			Name: "LanaCoin",
		},
	},
	"GUH": {
		{
			ID:   10481,
			Slug: "goes-up-higher",
			Name: "Goes Up Higher",
		},
	},
	"QTBK": {
		{
			ID:   6738,
			Slug: "quantbook",
			Name: "Quantbook",
		},
	},
	"WCK": {
		{
			ID:   7473,
			Slug: "wrapped-cryptokitties",
			Name: "Wrapped Basic CryptoKitties",
		},
	},
	"STS": {
		{
			ID:   5039,
			Slug: "sbank",
			Name: "SBank",
		},
	},
	"ROPE": {
		{
			ID:   9326,
			Slug: "rope-token",
			Name: "ROPE Token",
		},
	},
	"ABS": {
		{
			ID:   8032,
			Slug: "absorber-protocol",
			Name: "Absorber Protocol",
		},
	},
	"INT": {
		{
			ID:   2399,
			Slug: "int-chain",
			Name: "INT",
		},
	},
	"IC": {
		{
			ID:   2395,
			Slug: "ignition",
			Name: "Ignition",
		},
	},
	"DRGNBULL": {
		{
			ID:   6082,
			Slug: "3x-long-dragon-index-token",
			Name: "3X Long Dragon Index Token",
		},
	},
	"IBEX": {
		{
			ID:   10256,
			Slug: "ibex",
			Name: "IBEX",
		},
	},
	"UPP": {
		{
			ID:   2866,
			Slug: "sentinel-protocol",
			Name: "Sentinel Protocol",
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
	"GRPL": {
		{
			ID:   8059,
			Slug: "goldenratioperliquidity",
			Name: "Golden Ratio Per Liquidity",
		},
	},
	"FDR": {
		{
			ID:   6976,
			Slug: "french-digital-reserve",
			Name: "French Digital Reserve",
		},
	},
	"CHX": {
		{
			ID:   2673,
			Slug: "we-own",
			Name: "WeOwn",
		},
	},
	"BRDG": {
		{
			ID:   2968,
			Slug: "bridge-protocol",
			Name: "Bridge Protocol",
		},
	},
	"BQTX": {
		{
			ID:   3929,
			Slug: "bqt",
			Name: "BQT",
		},
	},
	"MMAON": {
		{
			ID:   9169,
			Slug: "mmaon",
			Name: "MMAON",
		},
	},
	"SHENG": {
		{
			ID:   6256,
			Slug: "sheng",
			Name: "SHENG",
		},
	},
	"LRC": {
		{
			ID:   1934,
			Slug: "loopring",
			Name: "Loopring",
		},
	},
	"PAC": {
		{
			ID:   1107,
			Slug: "pac-protocol",
			Name: "PAC Protocol",
		},
	},
	"CWV": {
		{
			ID:   3609,
			Slug: "cwv-chain",
			Name: "CWV Chain",
		},
	},
	"TROLL": {
		{
			ID:   638,
			Slug: "trollcoin",
			Name: "Trollcoin",
		},
	},
	"MTHD": {
		{
			ID:   9108,
			Slug: "method-finance",
			Name: "Method Finance",
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
	"LTCDOWN": {
		{
			ID:   7527,
			Slug: "ltcdown",
			Name: "LTCDOWN",
		},
	},
	"KYL": {
		{
			ID:   8644,
			Slug: "kylin",
			Name: "Kylin",
		},
	},
	"QRL": {
		{
			ID:   1712,
			Slug: "quantum-resistant-ledger",
			Name: "Quantum Resistant Ledger",
		},
	},
	"XGG": {
		{
			ID:   8685,
			Slug: "10xgg",
			Name: "10x.gg",
		},
	},
	"SHIBAKEN": {
		{
			ID:   10257,
			Slug: "shibaken-finance",
			Name: "Shibaken Finance",
		},
	},
	"DONK": {
		{
			ID:   9340,
			Slug: "donkey",
			Name: "Donkey",
		},
	},
	"KEN": {
		{
			ID:   6621,
			Slug: "keysians-network",
			Name: "Keysians Network",
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
	"PIB": {
		{
			ID:   3768,
			Slug: "pibble",
			Name: "PIBBLE",
		},
	},
	"OST": {
		{
			ID:   2296,
			Slug: "ost",
			Name: "OST",
		},
	},
	"TPAY": {
		{
			ID:   2627,
			Slug: "tokenpay",
			Name: "TokenPay",
		},
	},
	"MOLLYDOGE ⭐": {
		{
			ID:   10938,
			Slug: "mini-hollywood-doge",
			Name: "Mini Hollywood Doge",
		},
	},
	"RARE": {
		{
			ID:   8125,
			Slug: "unique-one",
			Name: "Unique One",
		},
	},
	"PCNT": {
		{
			ID:   8704,
			Slug: "playcent",
			Name: "Playcent",
		},
	},
	"vBCH": {
		{
			ID:   7974,
			Slug: "venus-bch",
			Name: "Venus BCH",
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
			ID:   551,
			Slug: "donu",
			Name: "Donu",
		},
	},
	"HPX": {
		{
			ID:   6062,
			Slug: "hupayx",
			Name: "HUPAYX",
		},
	},
	"DAIN": {
		{
			ID:   9304,
			Slug: "dain",
			Name: "DAIN",
		},
	},
	"MATIC": {
		{
			ID:   3890,
			Slug: "polygon",
			Name: "Polygon",
		},
	},
	"ELTCOIN": {
		{
			ID:   2147,
			Slug: "eltcoin",
			Name: "ELTCOIN",
		},
	},
	"ETHBEAR": {
		{
			ID:   5216,
			Slug: "3x-short-ethereum-token",
			Name: "3X Short Ethereum Token",
		},
	},
	"2CRZ": {
		{
			ID:   10418,
			Slug: "2crazynft",
			Name: "2crazyNFT",
		},
	},
	"TERA": {
		{
			ID:   3948,
			Slug: "tera",
			Name: "TERA",
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
	"FSW": {
		{
			ID:   6743,
			Slug: "fsw-token",
			Name: "Falconswap",
		},
	},
	"BHC": {
		{
			ID:   7182,
			Slug: "billionhappiness",
			Name: "BillionHappiness",
		},
	},
	"BCNT": {
		{
			ID:   4808,
			Slug: "bincentive",
			Name: "Bincentive",
		},
	},
	"HITX": {
		{
			ID:   8233,
			Slug: "hithotx",
			Name: "Hithotx",
		},
	},
	"HUE": {
		{
			ID:   8967,
			Slug: "hue",
			Name: "Hue",
		},
	},
	"MGC": {
		{
			ID:   4048,
			Slug: "mgc-token",
			Name: "MGC Token",
		},
	},
	"BSCGOLD": {
		{
			ID:   9759,
			Slug: "bsc-gold",
			Name: "BSC Gold",
		},
	},
	"KAVA": {
		{
			ID:   4846,
			Slug: "kava",
			Name: "Kava.io",
		},
	},
	"MCO": {
		{
			ID:   1776,
			Slug: "crypto-com",
			Name: "MCO",
		},
	},
	"DVX": {
		{
			ID:   6104,
			Slug: "derivex",
			Name: "Derivex",
		},
	},
	"HB": {
		{
			ID:   3171,
			Slug: "heartbout",
			Name: "HeartBout",
		},
	},
	"DOTUP": {
		{
			ID:   7003,
			Slug: "dotup",
			Name: "DOTUP",
		},
	},
	"CCE": {
		{
			ID:   7655,
			Slug: "cloudcoin",
			Name: "CloudCoin",
		},
	},
	"AUNIT": {
		{
			ID:   3747,
			Slug: "aunite",
			Name: "Aunite",
		},
	},
	"PIKA": {
		{
			ID:   8879,
			Slug: "pika",
			Name: "Pika",
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
	"O3": {
		{
			ID:   9588,
			Slug: "o3swap",
			Name: "O3Swap",
		},
	},
	"TAI": {
		{
			ID:   6866,
			Slug: "tai",
			Name: "TAI",
		},
	},
	"HDP.ф": {
		{
			ID:   4571,
			Slug: "hedpay",
			Name: "HEdpAY",
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
	"SUSHI": {
		{
			ID:   6758,
			Slug: "sushiswap",
			Name: "SushiSwap",
		},
	},
	"KCS": {
		{
			ID:   2087,
			Slug: "kucoin-token",
			Name: "KuCoin Token",
		},
	},
	"BTW": {
		{
			ID:   2489,
			Slug: "bitwhite",
			Name: "BitWhite",
		},
	},
	"LEOPARD": {
		{
			ID:   10393,
			Slug: "leopard",
			Name: "LEOPARD",
		},
	},
	"DOGEY": {
		{
			ID:   10941,
			Slug: "dogey-style",
			Name: "DOGEY STYLE",
		},
	},
	"CFT": {
		{
			ID:   6029,
			Slug: "coinbene-future-token",
			Name: "CoinBene Future Token",
		},
	},
	"BAB": {
		{
			ID:   10549,
			Slug: "bauble",
			Name: "Bauble",
		},
	},
	"STRAX": {
		{
			ID:   1343,
			Slug: "stratis",
			Name: "Stratis",
		},
	},
	"SLIM": {
		{
			ID:   9741,
			Slug: "solanium",
			Name: "Solanium",
		},
	},
	"RPZX": {
		{
			ID:   4298,
			Slug: "rapidz",
			Name: "Rapidz",
		},
	},
	"PIT": {
		{
			ID:   9177,
			Slug: "pitbull",
			Name: "Pitbull",
		},
	},
	"BHIBA": {
		{
			ID:   9769,
			Slug: "baby-shiba",
			Name: "Baby Shiba",
		},
	},
	"PERX": {
		{
			ID:   6468,
			Slug: "peerex",
			Name: "PeerEx",
		},
	},
	"BNTY": {
		{
			ID:   2310,
			Slug: "bounty0x",
			Name: "Bounty0x",
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
	"BENZ": {
		{
			ID:   3384,
			Slug: "benz",
			Name: "Benz",
		},
	},
	"BUP": {
		{
			ID:   7761,
			Slug: "buildup",
			Name: "BuildUp",
		},
	},
	"XNODE": {
		{
			ID:   8372,
			Slug: "xnode",
			Name: "XNODE",
		},
	},
	"URAC": {
		{
			ID:   4093,
			Slug: "uranus",
			Name: "Uranus",
		},
	},
	"NZO": {
		{
			ID:   5102,
			Slug: "enzo",
			Name: "Enzo",
		},
	},
	"XIF": {
		{
			ID:   7724,
			Slug: "x-infinity",
			Name: "X Infinity",
		},
	},
	"HEPA": {
		{
			ID:   10504,
			Slug: "hepa-finance",
			Name: "Hepa Finance",
		},
	},
	"DIRKDIGGLER": {
		{
			ID:   10749,
			Slug: "the-boogie-nights",
			Name: "The Boogie Nights",
		},
	},
	"DDD": {
		{
			ID:   2428,
			Slug: "scryinfo",
			Name: "Scry.info",
		},
	},
	"CORGIB": {
		{
			ID:   10251,
			Slug: "the-corgi-of-polkabridge",
			Name: "The Corgi of PolkaBridge",
		},
	},
	"BTC3S": {
		{
			ID:   5737,
			Slug: "amun-bitcoin-3x-daily-short",
			Name: "Amun Bitcoin 3x Daily Short",
		},
	},
	"IX": {
		{
			ID:   6326,
			Slug: "x-block",
			Name: "X-Block",
		},
	},
	"DKA": {
		{
			ID:   5908,
			Slug: "dkargo",
			Name: "dKargo",
		},
	},
	"SAITAMA": {
		{
			ID:   10498,
			Slug: "saitama-inu",
			Name: "Saitama Inu",
		},
	},
	"MDAO": {
		{
			ID:   9228,
			Slug: "martian-dao",
			Name: "Martian DAO",
		},
	},
	"CAD": {
		{
			ID:   8564,
			Slug: "candy-protocol",
			Name: "Candy Protocol",
		},
	},
	"OTO": {
		{
			ID:   3850,
			Slug: "otocash",
			Name: "OTOCASH",
		},
	},
	"INE": {
		{
			ID:   3811,
			Slug: "intellishare",
			Name: "IntelliShare",
		},
	},
	"ROCO": {
		{
			ID:   3677,
			Slug: "roiyal-coin",
			Name: "ROIyal Coin",
		},
	},
	"BXX": {
		{
			ID:   10949,
			Slug: "baanx",
			Name: "Baanx",
		},
	},
	"AZZR": {
		{
			ID:   7413,
			Slug: "azzure",
			Name: "Azzure",
		},
	},
	"ADK": {
		{
			ID:   1706,
			Slug: "aidos-kuneen",
			Name: "Aidos Kuneen",
		},
	},
	"BMI": {
		{
			ID:   8364,
			Slug: "bridge-mutual",
			Name: "Bridge Mutual",
		},
	},
	"ALLEY": {
		{
			ID:   9527,
			Slug: "nft-alley",
			Name: "NFT Alley",
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
	"ECP": {
		{
			ID:   9378,
			Slug: "eclipse-ecp",
			Name: "Eclipse",
		},
	},
	"LOVE": {
		{
			ID:   6615,
			Slug: "love-coin",
			Name: "Love Coin",
		},
	},
	"DCL": {
		{
			ID:   8201,
			Slug: "delphi-chain-link",
			Name: "Delphi Chain Link",
		},
	},
	"AM": {
		{
			ID:   10763,
			Slug: "aston-martin-cognizant-fan-token",
			Name: "Aston Martin Cognizant Fan Token",
		},
	},
	"VNT": {
		{
			ID:   3988,
			Slug: "vnt-chain",
			Name: "VNT Chain",
		},
	},
	"ZIK": {
		{
			ID:   6249,
			Slug: "ziktalk",
			Name: "Ziktalk",
		},
	},
	"KICH": {
		{
			ID:   10927,
			Slug: "kichicoin",
			Name: "KichiCoin",
		},
	},
	"SAMO": {
		{
			ID:   9721,
			Slug: "somoyedcoin",
			Name: "Samoyedcoin",
		},
	},
	"HPAY": {
		{
			ID:   8532,
			Slug: "hyper-credit-network",
			Name: "Hyper Credit Network",
		},
	},
	"ADAI": {
		{
			ID:   5763,
			Slug: "aave-dai",
			Name: "Aave DAI",
		},
	},
	"BURN1": {
		{
			ID:   9952,
			Slug: "burn1-coin",
			Name: "Burn1 Coin",
		},
	},
	"WORLD": {
		{
			ID:   8366,
			Slug: "world-token",
			Name: "World Token",
		},
	},
	"ZRC": {
		{
			ID:   1726,
			Slug: "zrcoin",
			Name: "ZrCoin",
		},
	},
	"VI": {
		{
			ID:   5428,
			Slug: "vid",
			Name: "Vid",
		},
	},
	"SQR": {
		{
			ID:   3467,
			Slug: "squorum",
			Name: "Squorum",
		},
	},
	"RNB": {
		{
			ID:   9766,
			Slug: "rentible",
			Name: "Rentible",
		},
	},
	"MEDIBIT": {
		{
			ID:   3582,
			Slug: "medibit",
			Name: "MediBit",
		},
	},
	"RFUEL": {
		{
			ID:   6537,
			Slug: "rio-defi",
			Name: "RioDeFi",
		},
	},
	"DAFI": {
		{
			ID:   8874,
			Slug: "dafi-protocol",
			Name: "DAFI Protocol",
		},
	},
	"BNX": {
		{
			ID:   9891,
			Slug: "binaryx",
			Name: "BinaryX",
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
	"GOLDR": {
		{
			ID:   6624,
			Slug: "golden-ratio-coin",
			Name: "Golden Ratio Coin",
		},
	},
	"ETH": {
		{
			ID:   1027,
			Slug: "ethereum",
			Name: "Ethereum",
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
	"GAME": {
		{
			ID:   576,
			Slug: "gamecredits",
			Name: "GameCredits",
		},
	},
	"PPP": {
		{
			ID:   2036,
			Slug: "paypie",
			Name: "PayPie",
		},
	},
	"THS": {
		{
			ID:   6696,
			Slug: "the-hash-speed",
			Name: "The Hash Speed",
		},
	},
	"LIBFX": {
		{
			ID:   6714,
			Slug: "libfx",
			Name: "Libfx",
		},
	},
	"QAC": {
		{
			ID:   3419,
			Slug: "quasarcoin",
			Name: "Quasarcoin",
		},
	},
	"SKYBORN": {
		{
			ID:   10184,
			Slug: "sky-born",
			Name: "SkyBorn",
		},
	},
	"DONUT": {
		{
			ID:   6156,
			Slug: "donut",
			Name: "Donut",
		},
	},
	"BIZZ": {
		{
			ID:   5437,
			Slug: "bizzcoin",
			Name: "BIZZCOIN",
		},
	},
	"OGC": {
		{
			ID:   9257,
			Slug: "one-get-coin",
			Name: "One Get Coin",
		},
	},
	"FIS": {
		{
			ID:   5882,
			Slug: "stafi",
			Name: "Stafi",
		},
	},
	"XPX": {
		{
			ID:   3126,
			Slug: "proximax",
			Name: "ProximaX",
		},
	},
	"TIPS": {
		{
			ID:   87,
			Slug: "fedoracoin",
			Name: "FedoraCoin",
		},
	},
	"COLL": {
		{
			ID:   8999,
			Slug: "collateral-pay",
			Name: "Collateral Pay",
		},
	},
	"HALV": {
		{
			ID:   6506,
			Slug: "halving-coin",
			Name: "Halving Token",
		},
	},
	"ESWAP": {
		{
			ID:   9567,
			Slug: "eswapping",
			Name: "eSwapping",
		},
	},
	"NSR": {
		{
			ID:   699,
			Slug: "nushares",
			Name: "NuShares",
		},
	},
	"CNS": {
		{
			ID:   5963,
			Slug: "centric-swap",
			Name: "Centric Swap",
		},
	},
	"VD": {
		{
			ID:   4119,
			Slug: "vindax-coin",
			Name: "VinDax Coin",
		},
	},
	"UNFT": {
		{
			ID:   10008,
			Slug: "ultra-nft",
			Name: "Ultra NFT",
		},
	},
	"SYL": {
		{
			ID:   9180,
			Slug: "xsl-labs",
			Name: "SYL",
		},
	},
	"HPPOT": {
		{
			ID:   10288,
			Slug: "healing-potion",
			Name: "Healing Potion",
		},
	},
	"PPY": {
		{
			ID:   1719,
			Slug: "peerplays-ppy",
			Name: "Peerplays",
		},
	},
	"yBAN": {
		{
			ID:   7489,
			Slug: "bananodos",
			Name: "BananoDOS",
		},
	},
	"SPERM": {
		{
			ID:   9902,
			Slug: "esperm",
			Name: "Elon Sperm",
		},
	},
	"STRK": {
		{
			ID:   8911,
			Slug: "strike",
			Name: "Strike",
		},
	},
	"PCM": {
		{
			ID:   4958,
			Slug: "precium",
			Name: "Precium",
		},
	},
	"UNDG": {
		{
			ID:   8400,
			Slug: "unidexgas",
			Name: "UniDexGas",
		},
	},
	"OLY": {
		{
			ID:   8484,
			Slug: "olyseum",
			Name: "Olyseum",
		},
	},
	"ORK": {
		{
			ID:   9150,
			Slug: "orakuru",
			Name: "Orakuru",
		},
	},
	"DGVC": {
		{
			ID:   6689,
			Slug: "degenvc",
			Name: "DegenVC",
		},
	},
	"TAVITT": {
		{
			ID:   7397,
			Slug: "tavittcoin",
			Name: "Tavittcoin",
		},
	},
	"0xBTC": {
		{
			ID:   2837,
			Slug: "0xbtc",
			Name: "0xBitcoin",
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
	"MATTER": {
		{
			ID:   8603,
			Slug: "antimatter",
			Name: "AntiMatter",
		},
	},
	"KAR": {
		{
			ID:   10042,
			Slug: "karura",
			Name: "Karura",
		},
	},
	"PIGX": {
		{
			ID:   7981,
			Slug: "pigx",
			Name: "PIGX",
		},
	},
	"QOOB": {
		{
			ID:   7347,
			Slug: "qoober",
			Name: "QOOBER",
		},
	},
	"RENBTC": {
		{
			ID:   5777,
			Slug: "renbtc",
			Name: "renBTC",
		},
	},
	"UMI": {
		{
			ID:   9632,
			Slug: "umi",
			Name: "UMI",
		},
	},
	"WOM": {
		{
			ID:   5328,
			Slug: "wom-protocol",
			Name: "WOM Protocol",
		},
	},
	"CRFI": {
		{
			ID:   9535,
			Slug: "crossfi",
			Name: "CrossFi",
		},
	},
	"ELY": {
		{
			ID:   2944,
			Slug: "elysian",
			Name: "Elysian",
		},
	},
	"ALLOY": {
		{
			ID:   9038,
			Slug: "hyper-alloy",
			Name: "HyperAlloy",
		},
	},
	"UVU": {
		{
			ID:   4244,
			Slug: "ccuniverse",
			Name: "CCUniverse",
		},
	},
	"BDPI": {
		{
			ID:   9153,
			Slug: "interest-bearing-dpi",
			Name: "Interest Bearing Defi Pulse Index",
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
	"EVA": {
		{
			ID:   10686,
			Slug: "evanesco-network",
			Name: "Evanesco Network",
		},
	},
	"LABO": {
		{
			ID:   9312,
			Slug: "the-lab-finance",
			Name: "The Lab Finance",
		},
	},
	"ZNN": {
		{
			ID:   4003,
			Slug: "zenon",
			Name: "Zenon",
		},
	},
	"NEBL": {
		{
			ID:   1955,
			Slug: "neblio",
			Name: "Neblio",
		},
	},
	"OBEE": {
		{
			ID:   6884,
			Slug: "obee-network",
			Name: "Obee Network",
		},
	},
	"ARNX": {
		{
			ID:   2153,
			Slug: "aeron",
			Name: "Aeron",
		},
	},
	"TTF": {
		{
			ID:   10847,
			Slug: "turbotrix-finance",
			Name: "TurboTrix Finance",
		},
	},
	"SPRINK": {
		{
			ID:   8844,
			Slug: "sprink",
			Name: "SPRINK",
		},
	},
	"BMP": {
		{
			ID:   7732,
			Slug: "brother-music-platform",
			Name: "Brother Music Platform",
		},
	},
	"UMBR": {
		{
			ID:   8780,
			Slug: "umbria-network",
			Name: "Umbria Network",
		},
	},
	"ADX": {
		{
			ID:   1768,
			Slug: "adx-net",
			Name: "AdEx Network",
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
	"WAB": {
		{
			ID:   2980,
			Slug: "wabnetwork",
			Name: "WABnetwork",
		},
	},
	"TAPE": {
		{
			ID:   10147,
			Slug: "toolape",
			Name: "ToolApe",
		},
	},
	"LOCC": {
		{
			ID:   9907,
			Slug: "low-orbit-crypto-cannon",
			Name: "Low Orbit Crypto Cannon",
		},
	},
	"INTX": {
		{
			ID:   5914,
			Slug: "intexcoin",
			Name: "intexcoin",
		},
	},
	"NEST": {
		{
			ID:   5841,
			Slug: "nest-protocol",
			Name: "NEST Protocol",
		},
	},
	"STV": {
		{
			ID:   8883,
			Slug: "sint-truidense-voetbalvereniging",
			Name: "Sint-Truidense Voetbalvereniging Fan Token",
		},
	},
	"THRT": {
		{
			ID:   2899,
			Slug: "thrive-token",
			Name: "Thrive Token",
		},
	},
	"CIFI": {
		{
			ID:   10129,
			Slug: "citizen-finance",
			Name: "Citizen Finance",
		},
	},
	"SIGNA": {
		{
			ID:   10776,
			Slug: "signum",
			Name: "Signum",
		},
	},
	"FRIES": {
		{
			ID:   7351,
			Slug: "fryworld",
			Name: "fry.world",
		},
	},
	"TKX": {
		{
			ID:   4715,
			Slug: "tokenize-xchange",
			Name: "Tokenize Xchange",
		},
	},
	"KFX": {
		{
			ID:   8177,
			Slug: "knoxfs-new",
			Name: "KnoxFS (new)",
		},
	},
	"VFOX": {
		{
			ID:   10290,
			Slug: "rfox-finance",
			Name: "RFOX Finance",
		},
	},
	"MVL": {
		{
			ID:   2982,
			Slug: "mvl",
			Name: "MVL",
		},
	},
	"NDN": {
		{
			ID:   5629,
			Slug: "ndn-link",
			Name: "NDN Link",
		},
	},
	"EGGP": {
		{
			ID:   10301,
			Slug: "eggplant-finance",
			Name: "Eggplant Finance",
		},
	},
	"SFP": {
		{
			ID:   8119,
			Slug: "safepal",
			Name: "SafePal",
		},
	},
	"FIRO": {
		{
			ID:   1414,
			Slug: "firo",
			Name: "Firo",
		},
	},
	"XHI": {
		{
			ID:   1244,
			Slug: "hicoin",
			Name: "HiCoin",
		},
	},
	"SLOKI": {
		{
			ID:   10868,
			Slug: "super-floki",
			Name: "Super Floki",
		},
	},
	"LIQUID": {
		{
			ID:   7650,
			Slug: "liquidefi",
			Name: "LIQUID",
		},
	},
	"DOGDEFI": {
		{
			ID:   7511,
			Slug: "dogdeficoin",
			Name: "DogDeFiCoin",
		},
	},
	"SOLAPE": {
		{
			ID:   10452,
			Slug: "solapefinance",
			Name: "SOLAPE Finance",
		},
	},
	"DUCATO": {
		{
			ID:   7133,
			Slug: "ducato-protocol-token",
			Name: "Ducato Protocol Token",
		},
	},
	"SRN": {
		{
			ID:   2313,
			Slug: "sirin-labs-token",
			Name: "SIRIN LABS Token",
		},
	},
	"BOMB": {
		{
			ID:   3956,
			Slug: "bomb",
			Name: "BOMB",
		},
	},
	"JUP": {
		{
			ID:   1503,
			Slug: "jupiter",
			Name: "Jupiter",
		},
	},
	"ORMEUS": {
		{
			ID:   1998,
			Slug: "ormeus-coin",
			Name: "Ormeus Coin",
		},
	},
	"SNY": {
		{
			ID:   9447,
			Slug: "synthetify",
			Name: "Synthetify",
		},
	},
	"KBTC": {
		{
			ID:   8892,
			Slug: "klondike-btc",
			Name: "Klondike BTC",
		},
	},
	"RUC": {
		{
			ID:   7423,
			Slug: "rush",
			Name: "RUSH",
		},
	},
	"RLC": {
		{
			ID:   1637,
			Slug: "rlc",
			Name: "iExec RLC",
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
	"NXS": {
		{
			ID:   789,
			Slug: "nexus",
			Name: "Nexus",
		},
	},
	"BIX": {
		{
			ID:   2307,
			Slug: "bibox-token",
			Name: "Bibox Token",
		},
	},
	"SS": {
		{
			ID:   2699,
			Slug: "sharder",
			Name: "Sharder",
		},
	},
	"BCNA": {
		{
			ID:   4263,
			Slug: "bitcanna",
			Name: "BitCanna",
		},
	},
	"XLT": {
		{
			ID:   6735,
			Slug: "nexalt",
			Name: "Nexalt",
		},
	},
	"OWN": {
		{
			ID:   3120,
			Slug: "owndata",
			Name: "OWNDATA",
		},
	},
	"ZFL": {
		{
			ID:   5247,
			Slug: "zuflo-coin",
			Name: "Zuflo Coin",
		},
	},
	"APW": {
		{
			ID:   10364,
			Slug: "apwine-finance",
			Name: "APWine Finance",
		},
	},
	"OCEAN": {
		{
			ID:   3911,
			Slug: "ocean-protocol",
			Name: "Ocean Protocol",
		},
	},
	"VRA": {
		{
			ID:   3816,
			Slug: "verasity",
			Name: "Verasity",
		},
	},
	"vLINK": {
		{
			ID:   7975,
			Slug: "venus-link",
			Name: "Venus LINK",
		},
	},
	"DYT": {
		{
			ID:   5894,
			Slug: "doyourtip",
			Name: "DoYourTip",
		},
	},
	"LOC": {
		{
			ID:   2287,
			Slug: "lockchain",
			Name: "LockTrip",
		},
	},
	"UGMC": {
		{
			ID:   10909,
			Slug: "unicly-genesis-mooncats-collection",
			Name: "Unicly Genesis MoonCats Collection",
		},
	},
	"vXRP": {
		{
			ID:   7965,
			Slug: "venus-xrp",
			Name: "Venus XRP",
		},
	},
	"ARA": {
		{
			ID:   9490,
			Slug: "ara-blocks",
			Name: "Ara Blocks",
		},
	},
	"WICKED": {
		{
			ID:   10422,
			Slug: "the-witcher-fans",
			Name: "The Witcher Fans",
		},
	},
	"ANJ": {
		{
			ID:   5523,
			Slug: "aragon-court",
			Name: "Aragon Court",
		},
	},
	"DOOGEE": {
		{
			ID:   10920,
			Slug: "doogee-io",
			Name: "Doogee.io",
		},
	},
	"AIB": {
		{
			ID:   911,
			Slug: "advanced-internet-blocks",
			Name: "Advanced Internet Blocks",
		},
	},
	"BCU": {
		{
			ID:   8777,
			Slug: "biscuit-farm-finance",
			Name: "Biscuit Farm Finance",
		},
	},
	"KUB": {
		{
			ID:   5793,
			Slug: "kublaicoin",
			Name: "Kublaicoin",
		},
	},
	"NRG": {
		{
			ID:   3218,
			Slug: "energi",
			Name: "Energi",
		},
	},
	"CVN": {
		{
			ID:   1810,
			Slug: "cvcoin",
			Name: "CVCoin",
		},
	},
	"VIDYA": {
		{
			ID:   6709,
			Slug: "vidya",
			Name: "Vidya",
		},
	},
	"WON": {
		{
			ID:   7504,
			Slug: "weblock",
			Name: "WeBlock",
		},
	},
	"EGX": {
		{
			ID:   3403,
			Slug: "eaglex",
			Name: "EagleX",
		},
	},
	"XCAD": {
		{
			ID:   9868,
			Slug: "xcad-network",
			Name: "XCAD Network",
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
	"MFO": {
		{
			ID:   10411,
			Slug: "moonfarm-finance",
			Name: "Moonfarm Finance",
		},
	},
	"THR": {
		{
			ID:   3144,
			Slug: "thorecoin",
			Name: "ThoreCoin",
		},
	},
	"NDS": {
		{
			ID:   9068,
			Slug: "nodeseeds",
			Name: "Nodeseeds",
		},
	},
	"VMR": {
		{
			ID:   5093,
			Slug: "vomer",
			Name: "VOMER",
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
	"BEC": {
		{
			ID:   7030,
			Slug: "betherchip",
			Name: "Betherchip",
		},
	},
	"CZRX": {
		{
			ID:   5743,
			Slug: "compound-0x",
			Name: "Compound 0x",
		},
	},
	"GINUX": {
		{
			ID:   10414,
			Slug: "green-shiba-inu-new",
			Name: "Green Shiba Inu (new)",
		},
	},
	"RYIP": {
		{
			ID:   8668,
			Slug: "ryi-platinum",
			Name: "RYI Platinum",
		},
	},
	"FC": {
		{
			ID:   9078,
			Slug: "fanscoin",
			Name: "FansCoin",
		},
	},
	"LINKDOWN": {
		{
			ID:   7012,
			Slug: "linkdown",
			Name: "LINKDOWN",
		},
	},
	"BABYBITC": {
		{
			ID:   11032,
			Slug: "babybitcoin",
			Name: "BabyBitcoin",
		},
	},
	"ZUZ": {
		{
			ID:   9074,
			Slug: "zuz-protocol",
			Name: "ZUZ Protocol",
		},
	},
	"ELG": {
		{
			ID:   10312,
			Slug: "escointoken",
			Name: "EscoinToken",
		},
	},
	"WAD": {
		{
			ID:   8981,
			Slug: "wardenswap",
			Name: "WardenSwap",
		},
	},
	"ID": {
		{
			ID:   8495,
			Slug: "everest",
			Name: "Everest",
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
	"xDAI": {
		{
			ID:   8635,
			Slug: "xdaistable",
			Name: "xDAI",
		},
	},
	"LTN": {
		{
			ID:   9534,
			Slug: "life-token",
			Name: "Life Token",
		},
	},
	"ABBC": {
		{
			ID:   3437,
			Slug: "abbc-coin",
			Name: "ABBC Coin",
		},
	},
	"AMS": {
		{
			ID:   1035,
			Slug: "amsterdamcoin",
			Name: "AmsterdamCoin",
		},
	},
	"IOEX": {
		{
			ID:   4289,
			Slug: "ioex",
			Name: "IOEX",
		},
	},
	"QTUM": {
		{
			ID:   1684,
			Slug: "qtum",
			Name: "Qtum",
		},
	},
	"TLOS": {
		{
			ID:   4660,
			Slug: "telos",
			Name: "Telos",
		},
	},
	"QRK": {
		{
			ID:   53,
			Slug: "quark",
			Name: "Quark",
		},
	},
	"YFL": {
		{
			ID:   6407,
			Slug: "yflink",
			Name: "YF Link",
		},
	},
	"CHOW": {
		{
			ID:   8433,
			Slug: "chow-chow",
			Name: "Chow Chow",
		},
	},
	"MOONPIRATE": {
		{
			ID:   9886,
			Slug: "moonpirate",
			Name: "MoonPirate",
		},
	},
	"GOAL": {
		{
			ID:   10267,
			Slug: "goal",
			Name: "Goal",
		},
	},
	"MCN": {
		{
			ID:   7482,
			Slug: "mcnetwork",
			Name: "McNetworkDefi",
		},
	},
	"THOR": {
		{
			ID:   8663,
			Slug: "asgard-finance",
			Name: "Asgard finance",
		},
	},
	"MEMEX": {
		{
			ID:   10363,
			Slug: "memex-exchange",
			Name: "MEMEX",
		},
	},
	"XSGD": {
		{
			ID:   8489,
			Slug: "xsgd",
			Name: "XSGD",
		},
	},
	"HLP": {
		{
			ID:   7263,
			Slug: "help-coin",
			Name: "HLP Token",
		},
	},
	"LINKBEAR": {
		{
			ID:   6094,
			Slug: "3x-short-chainlink-token",
			Name: "3X Short Chainlink Token",
		},
	},
	"PTN": {
		{
			ID:   3595,
			Slug: "palletone",
			Name: "PalletOne",
		},
	},
	"OPNN": {
		{
			ID:   4202,
			Slug: "opennity",
			Name: "Opennity",
		},
	},
	"COS": {
		{
			ID:   4036,
			Slug: "contentos",
			Name: "Contentos",
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
	"XIO": {
		{
			ID:   4997,
			Slug: "blockzerolabs",
			Name: "Blockzero Labs",
		},
	},
	"CYMT": {
		{
			ID:   3255,
			Slug: "cybermusic",
			Name: "CyberMusic",
		},
	},
	"NBL": {
		{
			ID:   10982,
			Slug: "nobility",
			Name: "Nobility",
		},
	},
	"PVM": {
		{
			ID:   9586,
			Slug: "privateum-initiative",
			Name: "PRIVATEUM INITIATIVE",
		},
	},
	"WLEO": {
		{
			ID:   7221,
			Slug: "wrapped-leo",
			Name: "Wrapped LEO",
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
	"CVR": {
		{
			ID:   8300,
			Slug: "polkacover",
			Name: "Polkacover",
		},
	},
	"CVCC": {
		{
			ID:   4152,
			Slug: "cryptoverificationcoin",
			Name: "CryptoVerificationCoin",
		},
	},
	"AAA": {
		{
			ID:   3287,
			Slug: "abulaba",
			Name: "Abulaba",
		},
	},
	"XTZ": {
		{
			ID:   2011,
			Slug: "tezos",
			Name: "Tezos",
		},
	},
	"BNT": {
		{
			ID:   1727,
			Slug: "bancor",
			Name: "Bancor",
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
	"YLC": {
		{
			ID:   3162,
			Slug: "yolocash",
			Name: "YoloCash",
		},
	},
	"EIFI": {
		{
			ID:   10648,
			Slug: "eifi-finance",
			Name: "EIFI FINANCE",
		},
	},
	"PROPEL": {
		{
			ID:   4043,
			Slug: "payrue-propel",
			Name: "PayRue (Propel)",
		},
	},
	"HDFL": {
		{
			ID:   10628,
			Slug: "hyper-deflate",
			Name: "Hyper Deflate",
		},
	},
	"DEFLCT": {
		{
			ID:   7973,
			Slug: "deflect",
			Name: "Deflect",
		},
	},
	"ACDC": {
		{
			ID:   3004,
			Slug: "volt",
			Name: "Volt",
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
	"CLS": {
		{
			ID:   9562,
			Slug: "coldstack",
			Name: "Coldstack",
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
	"WCO": {
		{
			ID:   3669,
			Slug: "winco",
			Name: "Winco",
		},
	},
	"ETHBACK": {
		{
			ID:   11002,
			Slug: "etherback",
			Name: "EtherBack",
		},
	},
	"XAUT": {
		{
			ID:   5176,
			Slug: "tether-gold",
			Name: "Tether Gold",
		},
	},
	"EDS": {
		{
			ID:   3076,
			Slug: "endorsit",
			Name: "Endorsit",
		},
	},
	"DTOP": {
		{
			ID:   6414,
			Slug: "data-trade-on-demand-platform",
			Name: "DTOP Token",
		},
	},
	"TEAM": {
		{
			ID:   2729,
			Slug: "tokenstars",
			Name: "TEAM (TokenStars)",
		},
	},
	"OBOT": {
		{
			ID:   9590,
			Slug: "obortech",
			Name: "OBORTECH",
		},
	},
	"CSAI": {
		{
			ID:   5264,
			Slug: "compound-sai",
			Name: "Compound SAI",
		},
	},
	"ARCH": {
		{
			ID:   7750,
			Slug: "archer-dao-governance-token",
			Name: "Archer DAO Governance Token",
		},
	},
	"FORM": {
		{
			ID:   10408,
			Slug: "formation-fi",
			Name: "Formation Fi",
		},
	},
	"PYRO": {
		{
			ID:   5169,
			Slug: "pyro-network",
			Name: "PYRO Network",
		},
	},
	"DGMT": {
		{
			ID:   6025,
			Slug: "digimax-dgmt",
			Name: "DigiMax DGMT",
		},
	},
	"ASI": {
		{
			ID:   8876,
			Slug: "asi-finance-token",
			Name: "ASI finance",
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
	"UBT": {
		{
			ID:   2758,
			Slug: "unibright",
			Name: "Unibright",
		},
	},
	"AIM": {
		{
			ID:   5918,
			Slug: "modihost",
			Name: "ModiHost",
		},
	},
	"HLIX": {
		{
			ID:   5447,
			Slug: "helix",
			Name: "Helix",
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
	"JVZ": {
		{
			ID:   7353,
			Slug: "jiviz",
			Name: "Jiviz",
		},
	},
	"GAS": {
		{
			ID:   1785,
			Slug: "gas",
			Name: "Gas",
		},
	},
	"BTO": {
		{
			ID:   2392,
			Slug: "bottos",
			Name: "Bottos",
		},
	},
	"PYLNT": {
		{
			ID:   2330,
			Slug: "pylon-network",
			Name: "Pylon Network",
		},
	},
	"RBIES": {
		{
			ID:   1175,
			Slug: "rubies",
			Name: "Rubies",
		},
	},
	"NMX": {
		{
			ID:   9438,
			Slug: "nominex-token",
			Name: "Nominex Token",
		},
	},
	"BSVBULL": {
		{
			ID:   5460,
			Slug: "3x-long-bitcoin-sv-token",
			Name: "3x Long Bitcoin SV Token",
		},
	},
	"50X": {
		{
			ID:   8868,
			Slug: "50x-com",
			Name: "50x.com",
		},
	},
	"NOTE": {
		{
			ID:   184,
			Slug: "dnotes",
			Name: "DNotes",
		},
	},
	"ATD": {
		{
			ID:   8926,
			Slug: "a2dao",
			Name: "A2DAO",
		},
	},
	"IQQ": {
		{
			ID:   9233,
			Slug: "iqoniq-fanecosystem",
			Name: "IQONIQ FanEcoSystem",
		},
	},
	"ETHPY": {
		{
			ID:   7049,
			Slug: "etherpay",
			Name: "Etherpay",
		},
	},
	"HFS": {
		{
			ID:   10354,
			Slug: "holder-swap",
			Name: "Holder Swap",
		},
	},
	"SNL": {
		{
			ID:   3977,
			Slug: "sport-and-leisure",
			Name: "Sport and Leisure",
		},
	},
	"APRIL": {
		{
			ID:   10367,
			Slug: "april",
			Name: "April",
		},
	},
	"PRY": {
		{
			ID:   8241,
			Slug: "prophecy",
			Name: "Prophecy",
		},
	},
	"HAMTARO": {
		{
			ID:   9575,
			Slug: "hamtaro",
			Name: "Hamtaro",
		},
	},
	"VRAP": {
		{
			ID:   8763,
			Slug: "veraswap",
			Name: "VeraSwap",
		},
	},
	"SWPRL": {
		{
			ID:   8550,
			Slug: "swaprol",
			Name: "Swaprol",
		},
	},
	"DWS": {
		{
			ID:   2919,
			Slug: "dws",
			Name: "DWS",
		},
	},
	"BXT": {
		{
			ID:   4858,
			Slug: "bitfxt-coin",
			Name: "BITFXT COIN",
		},
	},
	"FESS": {
		{
			ID:   5910,
			Slug: "fesschain",
			Name: "Fesschain",
		},
	},
	"APC": {
		{
			ID:   3512,
			Slug: "alpha-coin",
			Name: "Alpha Coin",
		},
	},
	"BONFIRE": {
		{
			ID:   9522,
			Slug: "bonfire",
			Name: "Bonfire",
		},
	},
	"DIC": {
		{
			ID:   4385,
			Slug: "daikicoin",
			Name: "Daikicoin",
		},
	},
	"ZETH": {
		{
			ID:   10284,
			Slug: "zetta-ethereum-hashrate-token",
			Name: "Zetta Ethereum Hashrate Token",
		},
	},
	"NETKO": {
		{
			ID:   1582,
			Slug: "netko",
			Name: "Netko",
		},
	},
	"FEAR": {
		{
			ID:   9866,
			Slug: "fear-nfts",
			Name: "Fear NFTs",
		},
	},
	"GOL": {
		{
			ID:   9670,
			Slug: "gogolcoin",
			Name: "GogolCoin",
		},
	},
	"WDP": {
		{
			ID:   7865,
			Slug: "waterdrop",
			Name: "WaterDrop",
		},
	},
	"yPANDA": {
		{
			ID:   8471,
			Slug: "yieldpanda-finance",
			Name: "YieldPanda Finance",
		},
	},
	"ARRR": {
		{
			ID:   3951,
			Slug: "pirate-chain",
			Name: "Pirate Chain",
		},
	},
	"CAS": {
		{
			ID:   2529,
			Slug: "cashaa",
			Name: "Cashaa",
		},
	},
	"UBX": {
		{
			ID:   7622,
			Slug: "ubix-network",
			Name: "UBIX.Network",
		},
	},
	"SFUND": {
		{
			ID:   8972,
			Slug: "seedify-fund",
			Name: "Seedify.fund",
		},
	},
	"LHB": {
		{
			ID:   8427,
			Slug: "lendhub",
			Name: "Lendhub",
		},
	},
	"NS": {
		{
			ID:   8633,
			Slug: "nodestats",
			Name: "Nodestats",
		},
	},
	"SBD": {
		{
			ID:   1312,
			Slug: "steem-dollars",
			Name: "Steem Dollars",
		},
	},
	"YEE": {
		{
			ID:   2437,
			Slug: "yee",
			Name: "YEE",
		},
	},
	"HELMET": {
		{
			ID:   8265,
			Slug: "helmet-insure",
			Name: "Helmet.insure",
		},
	},
	"SHVR": {
		{
			ID:   3645,
			Slug: "shivers",
			Name: "Shivers",
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
	"QUACK": {
		{
			ID:   10641,
			Slug: "richquack-com",
			Name: "RichQUACK.com",
		},
	},
	"BNZ": {
		{
			ID:   7415,
			Slug: "bonezyard",
			Name: "BonezYard",
		},
	},
	"GFARM2": {
		{
			ID:   8444,
			Slug: "gains-farm-v2",
			Name: "Gains Farm",
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
	"KASH": {
		{
			ID:   5995,
			Slug: "kids-cash",
			Name: "Kids Cash",
		},
	},
	"SAPP": {
		{
			ID:   4121,
			Slug: "sapphire",
			Name: "Sapphire",
		},
	},
	"AXPR": {
		{
			ID:   2466,
			Slug: "axpr-token",
			Name: "AXPR",
		},
	},
	"DYP": {
		{
			ID:   8080,
			Slug: "defi-yield-protocol",
			Name: "DeFi Yield Protocol",
		},
	},
	"ROSN": {
		{
			ID:   9783,
			Slug: "roseon-finance",
			Name: "Roseon Finance",
		},
	},
	"SUSHIBA": {
		{
			ID:   10308,
			Slug: "sushiba",
			Name: "Sushiba",
		},
	},
	"SENPAI": {
		{
			ID:   9242,
			Slug: "senpai",
			Name: "SENPAI",
		},
	},
	"CENNZ": {
		{
			ID:   2585,
			Slug: "centrality",
			Name: "Centrality",
		},
	},
	"SUNI": {
		{
			ID:   9823,
			Slug: "suni",
			Name: "SUNI",
		},
	},
	"KIMOCHI": {
		{
			ID:   8507,
			Slug: "kimochi-finance",
			Name: "Kimochi Finance",
		},
	},
	"REVV": {
		{
			ID:   6993,
			Slug: "revv",
			Name: "REVV",
		},
	},
	"LET": {
		{
			ID:   2468,
			Slug: "linkeye",
			Name: "LinkEye",
		},
	},
	"XIV": {
		{
			ID:   8746,
			Slug: "project-inverse",
			Name: "Project Inverse",
		},
	},
	"UCN": {
		{
			ID:   3302,
			Slug: "uchain",
			Name: "UChain",
		},
	},
	"RIT20": {
		{
			ID:   8467,
			Slug: "uberstate-inc",
			Name: "Uberstate RIT 2.0",
		},
	},
	"BBTC": {
		{
			ID:   10869,
			Slug: "baby-bitcoin",
			Name: "Baby Bitcoin",
		},
	},
	"COLLG": {
		{
			ID:   10860,
			Slug: "collateral-pay-governance",
			Name: "Collateral Pay Governance",
		},
	},
	"NTB": {
		{
			ID:   8284,
			Slug: "tokenasset",
			Name: "TokenAsset",
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
	"YSKF": {
		{
			ID:   8274,
			Slug: "yearn-shark-finance",
			Name: "Yearn Shark Finance",
		},
	},
	"XT": {
		{
			ID:   5999,
			Slug: "xtcom-token",
			Name: "XT.com Token",
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
	"FSXU": {
		{
			ID:   8811,
			Slug: "flashx-ultra",
			Name: "FlashX Ultra",
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
	"mQQQ": {
		{
			ID:   8025,
			Slug: "mirrored-invesco-qqq-trust",
			Name: "Mirrored Invesco QQQ Trust",
		},
	},
	"SWTH": {
		{
			ID:   2620,
			Slug: "switcheo",
			Name: "Switcheo",
		},
	},
	"CIV": {
		{
			ID:   3381,
			Slug: "civitas",
			Name: "Civitas",
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
	"SKIN": {
		{
			ID:   1830,
			Slug: "skincoin",
			Name: "SkinCoin",
		},
	},
	"STO": {
		{
			ID:   4162,
			Slug: "storeum",
			Name: "Storeum",
		},
	},
	"MOONSTAR": {
		{
			ID:   9214,
			Slug: "moonstar",
			Name: "MoonStar",
		},
	},
	"GLXM": {
		{
			ID:   10305,
			Slug: "galaxium",
			Name: "Galaxium",
		},
	},
	"ANG": {
		{
			ID:   6598,
			Slug: "aureus-nummus-gold",
			Name: "Aureus Nummus Gold",
		},
	},
	"AMD": {
		{
			ID:   7903,
			Slug: "advanced-micro-devices-tokenized-stock-ftx",
			Name: "Advanced Micro Devices tokenized stock FTX",
		},
	},
	"CLBR": {
		{
			ID:   7539,
			Slug: "colibri",
			Name: "Colibri Protocol",
		},
	},
	"ABAT": {
		{
			ID:   5754,
			Slug: "aave-bat",
			Name: "Aave BAT",
		},
	},
	"SMOL": {
		{
			ID:   7306,
			Slug: "smol",
			Name: "Smol",
		},
	},
	"KELPIE": {
		{
			ID:   10736,
			Slug: "kelpie-inu",
			Name: "Kelpie Inu",
		},
	},
	"SPI": {
		{
			ID:   8161,
			Slug: "shopping",
			Name: "Shopping",
		},
	},
	"NTON": {
		{
			ID:   6764,
			Slug: "nton",
			Name: "NTON",
		},
	},
	"MASH": {
		{
			ID:   9352,
			Slug: "marshmallowdefi",
			Name: "Marshmallowdefi",
		},
	},
	"OSB": {
		{
			ID:   7789,
			Slug: "oasisbloc",
			Name: "OASISBloc",
		},
	},
	"KODA": {
		{
			ID:   9966,
			Slug: "summit-koda-token",
			Name: "Summit Koda Token",
		},
	},
	"WSCRT": {
		{
			ID:   8337,
			Slug: "secret-erc20",
			Name: "Secret (ERC20)",
		},
	},
	"PCHART": {
		{
			ID:   10375,
			Slug: "polychart",
			Name: "Polychart",
		},
	},
	"SAFETESLA": {
		{
			ID:   9393,
			Slug: "safetesla",
			Name: "Safetesla",
		},
	},
	"ELK": {
		{
			ID:   10095,
			Slug: "elk-finance",
			Name: "Elk Finance",
		},
	},
	"CHOP": {
		{
			ID:   6633,
			Slug: "porkchop",
			Name: "Porkchop",
		},
	},
	"SLF": {
		{
			ID:   9684,
			Slug: "solarfare",
			Name: "Solarfare",
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
	"CRX": {
		{
			ID:   8731,
			Slug: "cryptex",
			Name: "CryptEx",
		},
	},
	"TIC": {
		{
			ID:   3131,
			Slug: "thingschain",
			Name: "Thingschain",
		},
	},
	"BKK": {
		{
			ID:   5154,
			Slug: "bkex-token",
			Name: "BKEX Token",
		},
	},
	"LBXC": {
		{
			ID:   5490,
			Slug: "lux-bio-cell",
			Name: "Lux Bio Cell",
		},
	},
	"HARD": {
		{
			ID:   7576,
			Slug: "hard-protocol",
			Name: "HARD Protocol",
		},
	},
	"COCOS": {
		{
			ID:   4275,
			Slug: "cocos-bcx",
			Name: "Cocos-BCX",
		},
	},
	"PRCY": {
		{
			ID:   8554,
			Slug: "prcy-coin",
			Name: "PRCY Coin",
		},
	},
	"RFI": {
		{
			ID:   7747,
			Slug: "reflect-finance",
			Name: "reflect.finance",
		},
	},
	"ION": {
		{
			ID:   1281,
			Slug: "ion",
			Name: "ION",
		},
	},
	"DRGB": {
		{
			ID:   6423,
			Slug: "dragonbit",
			Name: "Dragonbit",
		},
	},
	"USDFL": {
		{
			ID:   8437,
			Slug: "usdfreeliquidity",
			Name: "USDFreeLiquidity",
		},
	},
	"GROW": {
		{
			ID:   9813,
			Slug: "growingfi",
			Name: "GrowingFi",
		},
	},
	"BLCT": {
		{
			ID:   5280,
			Slug: "bloomzed-token",
			Name: "Bloomzed Loyalty Club Ticket",
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
	"MAGI": {
		{
			ID:   8914,
			Slug: "magikarp-finance",
			Name: "Magikarp Finance",
		},
	},
	"MINA": {
		{
			ID:   8646,
			Slug: "mina",
			Name: "Mina",
		},
	},
	"XNC": {
		{
			ID:   5060,
			Slug: "xenioscoin",
			Name: "XeniosCoin",
		},
	},
	"CRBN": {
		{
			ID:   7809,
			Slug: "carbon",
			Name: "Carbon",
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
	"EOSDAC": {
		{
			ID:   2644,
			Slug: "eosdac",
			Name: "eosDAC",
		},
	},
	"ALL": {
		{
			ID:   8882,
			Slug: "alliance-fan-token",
			Name: "Alliance Fan Token",
		},
	},
	"DHX": {
		{
			ID:   6771,
			Slug: "datahighway",
			Name: "DataHighway",
		},
	},
	"YFD": {
		{
			ID:   7690,
			Slug: "yfdfi-finance",
			Name: "Your Finance Decentralized",
		},
	},
	"NFTSHIBA": {
		{
			ID:   10614,
			Slug: "nftshiba-finance",
			Name: "NFTShiba.Finance",
		},
	},
	"IG": {
		{
			ID:   3121,
			Slug: "igtoken",
			Name: "IGToken",
		},
	},
	"WIFI": {
		{
			ID:   8140,
			Slug: "wifi-coin",
			Name: "Wifi Coin",
		},
	},
	"MINDS": {
		{
			ID:   8675,
			Slug: "minds",
			Name: "Minds",
		},
	},
	"FEAST": {
		{
			ID:   10116,
			Slug: "feast-finance",
			Name: "Feast Finance",
		},
	},
	"BKR": {
		{
			ID:   10542,
			Slug: "bakerdao",
			Name: "BakerDAO",
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
	"MAN": {
		{
			ID:   2474,
			Slug: "matrix-ai-network",
			Name: "Matrix AI Network",
		},
	},
	"INN": {
		{
			ID:   2160,
			Slug: "innova",
			Name: "Innova",
		},
	},
	"DOBE": {
		{
			ID:   10205,
			Slug: "dobermann",
			Name: "Dobermann",
		},
	},
	"MKB": {
		{
			ID:   9727,
			Slug: "maker-basic-mkb",
			Name: "Maker Basic-MKB",
		},
	},
	"GLCH": {
		{
			ID:   8236,
			Slug: "glitch",
			Name: "Glitch",
		},
	},
	"DVC": {
		{
			ID:   5902,
			Slug: "dragonvein",
			Name: "DragonVein",
		},
	},
	"SKLAY": {
		{
			ID:   7881,
			Slug: "sklay",
			Name: "sKLAY",
		},
	},
	"PUSSY": {
		{
			ID:   9639,
			Slug: "pussy-financial",
			Name: "Pussy Financial",
		},
	},
	"VLK": {
		{
			ID:   9576,
			Slug: "vulkania",
			Name: "Vulkania",
		},
	},
	"ILC": {
		{
			ID:   3617,
			Slug: "ilcoin",
			Name: "ILCOIN",
		},
	},
	"NRP": {
		{
			ID:   3397,
			Slug: "neural-protocol",
			Name: "Neural Protocol",
		},
	},
	"SWYFTT": {
		{
			ID:   4588,
			Slug: "swyft",
			Name: "SWYFT",
		},
	},
	"HAM": {
		{
			ID:   10336,
			Slug: "hamster",
			Name: "Hamster",
		},
	},
	"DSD": {
		{
			ID:   8106,
			Slug: "dynamic-set-dollar",
			Name: "Dynamic Set Dollar",
		},
	},
	"NOIZ": {
		{
			ID:   4125,
			Slug: "noizchain",
			Name: "NOIZ",
		},
	},
	"CRUSADER": {
		{
			ID:   10976,
			Slug: "crusaders-of-crypto",
			Name: "Crusaders of Crypto",
		},
	},
	"POLR": {
		{
			ID:   10561,
			Slug: "polystarter-net",
			Name: "Polystarter.net",
		},
	},
	"ZERO": {
		{
			ID:   8293,
			Slug: "zero-exchange",
			Name: "Zero Exchange",
		},
	},
	"QLC": {
		{
			ID:   2321,
			Slug: "qlink",
			Name: "QLC Chain",
		},
	},
	"ULG": {
		{
			ID:   4996,
			Slug: "ultragate",
			Name: "Ultragate",
		},
	},
	"NFTART": {
		{
			ID:   9299,
			Slug: "nft-art-finance",
			Name: "NFT Art Finance",
		},
	},
	"BIGGLES": {
		{
			ID:   11028,
			Slug: "mr-bigglesworth",
			Name: "Mr Bigglesworth",
		},
	},
	"JCC": {
		{
			ID:   6914,
			Slug: "junca-cash",
			Name: "junca Cash",
		},
	},
	"CREX": {
		{
			ID:   4225,
			Slug: "crex-token",
			Name: "Crex Token",
		},
	},
	"DYN": {
		{
			ID:   1587,
			Slug: "dynamic",
			Name: "Dynamic",
		},
	},
	"HDAO": {
		{
			ID:   5299,
			Slug: "hyperdao",
			Name: "HyperDAO",
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
	"YBEAR": {
		{
			ID:   9147,
			Slug: "ybear-finance",
			Name: "yBEAR.finance",
		},
	},
	"ROS": {
		{
			ID:   5730,
			Slug: "ros-coin",
			Name: "ROS Coin",
		},
	},
	"MEDIADOGE": {
		{
			ID:   10811,
			Slug: "the-mediadoge",
			Name: "The MEDIADOGE",
		},
	},
	"BAL": {
		{
			ID:   5728,
			Slug: "balancer",
			Name: "Balancer",
		},
	},
	"COW": {
		{
			ID:   10011,
			Slug: "coinwind",
			Name: "CoinWind",
		},
	},
	"SYA": {
		{
			ID:   9635,
			Slug: "save-your-assets",
			Name: "Save Your Assets",
		},
	},
	"SAFEP": {
		{
			ID:   9266,
			Slug: "safe-protocol",
			Name: "Safe Protocol",
		},
	},
	"ILK": {
		{
			ID:   4754,
			Slug: "inlock",
			Name: "INLOCK",
		},
	},
	"VEN": {
		{
			ID:   8908,
			Slug: "impulseven",
			Name: "ImpulseVen",
		},
	},
	"AICO": {
		{
			ID:   5920,
			Slug: "aicon",
			Name: "AICON",
		},
	},
	"BST1": {
		{
			ID:   7670,
			Slug: "blueshare-token",
			Name: "Blueshare Token",
		},
	},
	"SAFESUN": {
		{
			ID:   9294,
			Slug: "safesun-crypto",
			Name: "SAFESUN",
		},
	},
	"WOZX": {
		{
			ID:   7882,
			Slug: "efforce",
			Name: "EFFORCE",
		},
	},
	"DOTDOWN": {
		{
			ID:   7006,
			Slug: "dotdown",
			Name: "DOTDOWN",
		},
	},
	"RABBIT": {
		{
			ID:   10178,
			Slug: "rabbit-finance",
			Name: "Rabbit Finance",
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
	"ALM": {
		{
			ID:   10428,
			Slug: "alium-finance",
			Name: "Alium Finance",
		},
	},
	"R2R": {
		{
			ID:   5204,
			Slug: "citios",
			Name: "CitiOs",
		},
	},
	"KENU": {
		{
			ID:   10635,
			Slug: "ken-inu",
			Name: "Ken Inu",
		},
	},
	"POLA": {
		{
			ID:   7740,
			Slug: "polaris-share",
			Name: "Polaris Share",
		},
	},
	"ELONGATE": {
		{
			ID:   9178,
			Slug: "elongate",
			Name: "ElonGate",
		},
	},
	"PNL": {
		{
			ID:   9605,
			Slug: "truepnl",
			Name: "TruePNL",
		},
	},
	"CHIMOM": {
		{
			ID:   10175,
			Slug: "chihua-chimon",
			Name: "Chihua Token",
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
	"ZCL": {
		{
			ID:   1447,
			Slug: "zclassic",
			Name: "ZClassic",
		},
	},
	"TFI": {
		{
			ID:   10585,
			Slug: "trustfi-network",
			Name: "TrustFi Network",
		},
	},
	"EXT": {
		{
			ID:   3074,
			Slug: "experience-token",
			Name: "Experience Token",
		},
	},
	"DOTA": {
		{
			ID:   10406,
			Slug: "dota-finance",
			Name: "Dota Finance",
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
	"DOVE": {
		{
			ID:   9355,
			Slug: "doveswap-finance",
			Name: "DoveSwap Finance",
		},
	},
	"YLFI": {
		{
			ID:   7788,
			Slug: "yearn-loans-finance",
			Name: "Yearn Loans Finance",
		},
	},
	"STMX": {
		{
			ID:   2297,
			Slug: "stormx",
			Name: "StormX",
		},
	},
	"vBNB": {
		{
			ID:   7961,
			Slug: "venus-bnb",
			Name: "Venus BNB",
		},
	},
	"PERL": {
		{
			ID:   4293,
			Slug: "perlin",
			Name: "PERL.eco",
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
	"LTCBULL": {
		{
			ID:   5462,
			Slug: "3x-long-litecoin-token",
			Name: "3x Long Litecoin Token",
		},
	},
	"SUSHIBEAR": {
		{
			ID:   7085,
			Slug: "3x-short-sushi-token",
			Name: "3X Short Sushi Token",
		},
	},
	"CHARIX TOKEN": {
		{
			ID:   10241,
			Slug: "charix",
			Name: "Charix",
		},
	},
	"DFI": {
		{
			ID:   5804,
			Slug: "defichain",
			Name: "DeFiChain",
		},
	},
	"DYNMT": {
		{
			ID:   4193,
			Slug: "dynamite",
			Name: "Dynamite",
		},
	},
	"GLM": {
		{
			ID:   1455,
			Slug: "golem-network-tokens",
			Name: "Golem",
		},
	},
	"TKN": {
		{
			ID:   1660,
			Slug: "monolith",
			Name: "Monolith",
		},
	},
	"SDFI": {
		{
			ID:   9064,
			Slug: "sting-defi",
			Name: "Sting Defi",
		},
	},
	"MAMADOGE": {
		{
			ID:   10835,
			Slug: "mamadoge",
			Name: "MAMADOGE",
		},
	},
	"$TRDL": {
		{
			ID:   8390,
			Slug: "strudel-finance",
			Name: "Strudel Finance",
		},
	},
	"RCUBE": {
		{
			ID:   10590,
			Slug: "retro-defi-rcube",
			Name: "RETRO DEFI - RCUBE",
		},
	},
	"SNBL": {
		{
			ID:   9036,
			Slug: "nebulaprotocol",
			Name: "Nebulaprotocol",
		},
	},
	"IOST": {
		{
			ID:   2405,
			Slug: "iostoken",
			Name: "IOST",
		},
	},
	"ASR": {
		{
			ID:   5229,
			Slug: "as-roma-fan-token",
			Name: "AS Roma Fan Token",
		},
	},
	"CHP": {
		{
			ID:   2569,
			Slug: "coinpoker",
			Name: "CoinPoker",
		},
	},
	"XPN": {
		{
			ID:   4760,
			Slug: "pantheon-x",
			Name: "PANTHEON X",
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
	"ADZ": {
		{
			ID:   1136,
			Slug: "adzcoin",
			Name: "Adzcoin",
		},
	},
	"STREAM": {
		{
			ID:   4365,
			Slug: "streamit-coin",
			Name: "Streamit Coin",
		},
	},
	"BVL": {
		{
			ID:   7811,
			Slug: "bullswap-exchange",
			Name: "Bullswap Exchange",
		},
	},
	"PUGL": {
		{
			ID:   10333,
			Slug: "puglife",
			Name: "PUGLIFE",
		},
	},
	"TRAVEL": {
		{
			ID:   10754,
			Slug: "travel-care",
			Name: "Travel Care",
		},
	},
	"GGC": {
		{
			ID:   5090,
			Slug: "global-game-coin",
			Name: "Global Game Coin",
		},
	},
	"XHV": {
		{
			ID:   2662,
			Slug: "haven-protocol",
			Name: "Haven Protocol",
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
	"SMARS": {
		{
			ID:   10502,
			Slug: "smars",
			Name: "SafeMars",
		},
	},
	"SUP": {
		{
			ID:   7375,
			Slug: "sup",
			Name: "SUP",
		},
	},
	"BIXB": {
		{
			ID:   9330,
			Slug: "bixbcoin",
			Name: "BIXBCOIN",
		},
	},
	"CROW": {
		{
			ID:   8533,
			Slug: "crow-finance",
			Name: "Crow Finance",
		},
	},
	"BTC": {
		{
			ID:   1,
			Slug: "bitcoin",
			Name: "Bitcoin",
		},
	},
	"CTK": {
		{
			ID:   4807,
			Slug: "certik",
			Name: "CertiK",
		},
	},
	"JULD": {
		{
			ID:   8164,
			Slug: "julswap",
			Name: "JulSwap",
		},
	},
	"OIN": {
		{
			ID:   6870,
			Slug: "oin-finance",
			Name: "OIN Finance",
		},
	},
	"RITO": {
		{
			ID:   4264,
			Slug: "ritocoin",
			Name: "Ritocoin",
		},
	},
	"NFTBS": {
		{
			ID:   11042,
			Slug: "nftbooks",
			Name: "NFTBooks",
		},
	},
	"HARE": {
		{
			ID:   10675,
			Slug: "hare-token",
			Name: "Hare Token",
		},
	},
	"TPAD": {
		{
			ID:   9613,
			Slug: "trustpad",
			Name: "Trustpad",
		},
	},
	"BCDN": {
		{
			ID:   2247,
			Slug: "blockcdn",
			Name: "BlockCDN",
		},
	},
	"LNOT": {
		{
			ID:   7118,
			Slug: "livenodes-token",
			Name: "Livenodes Token",
		},
	},
	"HNS": {
		{
			ID:   5221,
			Slug: "handshake",
			Name: "Handshake",
		},
	},
	"PYR": {
		{
			ID:   9308,
			Slug: "vulcan-forged-pyr",
			Name: "Vulcan Forged PYR",
		},
	},
	"HOGE": {
		{
			ID:   8438,
			Slug: "hoge-finance",
			Name: "Hoge Finance",
		},
	},
	"ZT": {
		{
			ID:   3458,
			Slug: "zbg-token",
			Name: "ZBG Token",
		},
	},
	"ACT": {
		{
			ID:   1918,
			Slug: "achain",
			Name: "Achain",
		},
	},
	"PPAY": {
		{
			ID:   7870,
			Slug: "plasma-finance",
			Name: "Plasma Finance",
		},
	},
	"AHOUSE": {
		{
			ID:   10541,
			Slug: "animalhouse",
			Name: "AnimalHouse",
		},
	},
	"MIKS": {
		{
			ID:   6675,
			Slug: "miks-coin",
			Name: "MIKS COIN",
		},
	},
	"BULLY": {
		{
			ID:   10243,
			Slug: "pitbully",
			Name: "PitBULLY",
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
	"SFC": {
		{
			ID:   10937,
			Slug: "safecap-token",
			Name: "SafeCap Token",
		},
	},
	"WBX": {
		{
			ID:   5450,
			Slug: "wibx",
			Name: "WiBX",
		},
	},
	"OLIVE": {
		{
			ID:   10685,
			Slug: "olive-cash",
			Name: "Olive.Cash",
		},
	},
	"RAP": {
		{
			ID:   7832,
			Slug: "realpay",
			Name: "REALPAY",
		},
	},
	"CYTR": {
		{
			ID:   7307,
			Slug: "cyclops-treasure",
			Name: "Cyclops Treasure",
		},
	},
	"BAEPAY": {
		{
			ID:   7657,
			Slug: "baepay",
			Name: "BAEPAY",
		},
	},
	"ZIPT": {
		{
			ID:   2724,
			Slug: "zippie",
			Name: "Zippie",
		},
	},
	"CHAR": {
		{
			ID:   9361,
			Slug: "charitas",
			Name: "Charitas",
		},
	},
	"MNTT": {
		{
			ID:   9709,
			Slug: "moontrust",
			Name: "MoonTrust",
		},
	},
	"arNXM": {
		{
			ID:   8328,
			Slug: "armor-nxm",
			Name: "Armor NXM",
		},
	},
	"GADOSHI": {
		{
			ID:   8600,
			Slug: "gadoshi",
			Name: "Gadoshi",
		},
	},
	"CSPR": {
		{
			ID:   5899,
			Slug: "casper",
			Name: "Casper",
		},
	},
	"ECO": {
		{
			ID:   4466,
			Slug: "ormeus-ecosystem",
			Name: "Ormeus Ecosystem",
		},
	},
	"OOE": {
		{
			ID:   9938,
			Slug: "openocean",
			Name: "OpenOcean",
		},
	},
	"BTCR": {
		{
			ID:   5404,
			Slug: "bitcurate",
			Name: "Bitcurate",
		},
	},
	"ADS": {
		{
			ID:   1883,
			Slug: "adshares",
			Name: "Adshares",
		},
	},
	"FTI": {
		{
			ID:   2901,
			Slug: "fanstime",
			Name: "FansTime",
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
	"Sunder": {
		{
			ID:   10298,
			Slug: "sunder-goverance-token",
			Name: "Sunder Goverance Token",
		},
	},
	"YFW": {
		{
			ID:   7983,
			Slug: "yfworld",
			Name: "YFWorld",
		},
	},
	"LUCY": {
		{
			ID:   5291,
			Slug: "lucy",
			Name: "LUCY",
		},
	},
	"EVR": {
		{
			ID:   2066,
			Slug: "everus",
			Name: "Everus",
		},
	},
	"BYTE": {
		{
			ID:   6126,
			Slug: "btc-network-demand-set-ii",
			Name: "BTC Network Demand Set II",
		},
	},
	"DEXT": {
		{
			ID:   5866,
			Slug: "dextools",
			Name: "DEXTools",
		},
	},
	"LCX": {
		{
			ID:   4950,
			Slug: "lcx",
			Name: "LCX",
		},
	},
	"ENG": {
		{
			ID:   2044,
			Slug: "enigma",
			Name: "Enigma",
		},
	},
	"SWAPZ": {
		{
			ID:   10557,
			Slug: "swapz",
			Name: "Swapz",
		},
	},
	"GENIUS": {
		{
			ID:   10948,
			Slug: "genius-coin",
			Name: "Genius Coin",
		},
	},
	"NZDX": {
		{
			ID:   5494,
			Slug: "etoro-new-zealand-dollar",
			Name: "eToro New Zealand Dollar",
		},
	},
	"ALD": {
		{
			ID:   6519,
			Slug: "aludra-network",
			Name: "Aludra Network",
		},
	},
	"CUMMIES": {
		{
			ID:   9212,
			Slug: "cumrocket",
			Name: "CUMROCKET",
		},
	},
	"JOB": {
		{
			ID:   4287,
			Slug: "jobchain",
			Name: "Jobchain",
		},
	},
	"ABL": {
		{
			ID:   3156,
			Slug: "airbloc",
			Name: "Airbloc",
		},
	},
	"HNST": {
		{
			ID:   4035,
			Slug: "honest",
			Name: "Honest",
		},
	},
	"TYPE": {
		{
			ID:   3505,
			Slug: "typerium",
			Name: "Typerium",
		},
	},
	"OMC": {
		{
			ID:   5818,
			Slug: "ormeus-cash",
			Name: "Ormeus Cash",
		},
	},
	"GIN": {
		{
			ID:   2773,
			Slug: "gincoin",
			Name: "GINcoin",
		},
	},
	"ECELL": {
		{
			ID:   7212,
			Slug: "consensus-cell-network",
			Name: "CellETF",
		},
	},
	"HAKA": {
		{
			ID:   10526,
			Slug: "tribe-one",
			Name: "TribeOne",
		},
	},
	"DEFIT": {
		{
			ID:   9155,
			Slug: "digital-fitness",
			Name: "Digital Fitness",
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
	"SUSD": {
		{
			ID:   2927,
			Slug: "susd",
			Name: "sUSD",
		},
	},
	"COVER": {
		{
			ID:   8175,
			Slug: "cover-protocol-new",
			Name: "COVER Protocol",
		},
	},
	"1X2": {
		{
			ID:   3767,
			Slug: "1x2-coin",
			Name: "1X2 COIN",
		},
	},
	"BSL": {
		{
			ID:   10036,
			Slug: "bsclaunch",
			Name: "BSClaunch",
		},
	},
	"NPLC": {
		{
			ID:   3649,
			Slug: "plus-coin",
			Name: "Plus-Coin",
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
	"MYOBU": {
		{
			ID:   10578,
			Slug: "myobu",
			Name: "Myōbu",
		},
	},
	"TSR": {
		{
			ID:   4806,
			Slug: "tesra",
			Name: "Tesra",
		},
	},
	"BEL": {
		{
			ID:   6928,
			Slug: "bella-protocol",
			Name: "Bella Protocol",
		},
	},
	"ZIG": {
		{
			ID:   9260,
			Slug: "zigcoin",
			Name: "Zigcoin",
		},
	},
	"RULER": {
		{
			ID:   8698,
			Slug: "ruler-protocol",
			Name: "Ruler Protocol",
		},
	},
	"LAT": {
		{
			ID:   9720,
			Slug: "platon",
			Name: "PlatON",
		},
	},
	"XDEF2": {
		{
			ID:   8226,
			Slug: "xdef-finance",
			Name: "Xdef Finance",
		},
	},
	"mGOOGL": {
		{
			ID:   8003,
			Slug: "mirrored-google",
			Name: "Mirrored Google",
		},
	},
	"vETH": {
		{
			ID:   7963,
			Slug: "venus-eth",
			Name: "Venus ETH",
		},
	},
	"SHARD": {
		{
			ID:   3335,
			Slug: "shard",
			Name: "Shard",
		},
	},
	"NSBT": {
		{
			ID:   7320,
			Slug: "neutrino-system-base-token",
			Name: "Neutrino Token",
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
	"PKOIN": {
		{
			ID:   5925,
			Slug: "pocketnet",
			Name: "Pkoin",
		},
	},
	"NFI": {
		{
			ID:   9275,
			Slug: "norse-finance",
			Name: "Norse Finance",
		},
	},
	"DDRT": {
		{
			ID:   5697,
			Slug: "digidinar-token",
			Name: "DigiDinar Token",
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
	"DEAL": {
		{
			ID:   3439,
			Slug: "idealcash",
			Name: "iDealCash",
		},
	},
	"MOLA": {
		{
			ID:   10185,
			Slug: "moonlana",
			Name: "Moonlana",
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
	"BRD": {
		{
			ID:   2306,
			Slug: "bread",
			Name: "Bread",
		},
	},
	"WISB": {
		{
			ID:   10000,
			Slug: "wise-token",
			Name: "Wise Token",
		},
	},
	"$PIR": {
		{
			ID:   8013,
			Slug: "piranhas",
			Name: "PIRANHAS",
		},
	},
	"CLBK": {
		{
			ID:   3712,
			Slug: "cloudbric",
			Name: "Cloudbric",
		},
	},
	"STM": {
		{
			ID:   5046,
			Slug: "streamity",
			Name: "Streamity",
		},
	},
	"MINIDOGE": {
		{
			ID:   10798,
			Slug: "minidoge",
			Name: "MiniDOGE",
		},
	},
	"SHARK": {
		{
			ID:   9949,
			Slug: "baby-shark",
			Name: "Baby Shark",
		},
	},
	"GZE": {
		{
			ID:   3257,
			Slug: "gazecoin",
			Name: "GazeCoin",
		},
	},
	"BTCHG": {
		{
			ID:   5732,
			Slug: "bitcoinhedge",
			Name: "BITCOINHEDGE",
		},
	},
	"SXP": {
		{
			ID:   4279,
			Slug: "swipe",
			Name: "Swipe",
		},
	},
	"SHR": {
		{
			ID:   4197,
			Slug: "sharetoken",
			Name: "ShareToken",
		},
	},
	"TAD": {
		{
			ID:   7559,
			Slug: "tadpole-finance",
			Name: "Tadpole Finance",
		},
	},
	"FTN": {
		{
			ID:   3658,
			Slug: "fountain",
			Name: "Fountain",
		},
	},
	"TMT": {
		{
			ID:   7246,
			Slug: "tamy-token",
			Name: "Tamy Token",
		},
	},
	"SEND": {
		{
			ID:   2255,
			Slug: "social-send",
			Name: "Social Send",
		},
	},
	"SERO": {
		{
			ID:   4078,
			Slug: "super-zero-protocol",
			Name: "Super Zero Protocol",
		},
	},
	"FGC": {
		{
			ID:   2870,
			Slug: "fantasygold",
			Name: "FantasyGold",
		},
	},
	"BCEO": {
		{
			ID:   3888,
			Slug: "bitceo",
			Name: "bitCEO",
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
	"GLC": {
		{
			ID:   25,
			Slug: "goldcoin",
			Name: "Goldcoin",
		},
	},
	"POLT": {
		{
			ID:   9392,
			Slug: "polkatrain",
			Name: "Polkatrain",
		},
	},
	"SXAG": {
		{
			ID:   5863,
			Slug: "sxag",
			Name: "sXAG",
		},
	},
	"FARA": {
		{
			ID:   9530,
			Slug: "faraland",
			Name: "FaraLand",
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
	"LEMD": {
		{
			ID:   9455,
			Slug: "lemond",
			Name: "Lemond",
		},
	},
	"HPS": {
		{
			ID:   8198,
			Slug: "happinesstoken",
			Name: "HappinessToken",
		},
	},
	"POLICEDOGE": {
		{
			ID:   10913,
			Slug: "policedoge",
			Name: "PoliceDOGE",
		},
	},
	"ITEN": {
		{
			ID:   7154,
			Slug: "iten",
			Name: "ITEN",
		},
	},
	"OBTC": {
		{
			ID:   3152,
			Slug: "obitan-chain",
			Name: "Obitan Chain",
		},
	},
	"WBNB": {
		{
			ID:   7192,
			Slug: "wbnb",
			Name: "Wrapped BNB",
		},
	},
	"BTB": {
		{
			ID:   3687,
			Slug: "bitball",
			Name: "BitBall",
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
	"DIN": {
		{
			ID:   3263,
			Slug: "dinero",
			Name: "Dinero",
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
	"NOA": {
		{
			ID:   9997,
			Slug: "noa-play",
			Name: "NOA PLAY",
		},
	},
	"UBQ": {
		{
			ID:   588,
			Slug: "ubiq",
			Name: "Ubiq",
		},
	},
	"VLD": {
		{
			ID:   3492,
			Slug: "vetri",
			Name: "Vetri",
		},
	},
	"CPR": {
		{
			ID:   4589,
			Slug: "cipher",
			Name: "Cipher",
		},
	},
	"HYDRA": {
		{
			ID:   8245,
			Slug: "hydra",
			Name: "Hydra",
		},
	},
	"KOKOMO": {
		{
			ID:   10709,
			Slug: "kokomoswap",
			Name: "KokomoSwap",
		},
	},
	"LOOM": {
		{
			ID:   2588,
			Slug: "loom-network",
			Name: "Loom Network",
		},
	},
	"SANDG": {
		{
			ID:   1090,
			Slug: "save-and-gain",
			Name: "Save and Gain",
		},
	},
	"KTT": {
		{
			ID:   5130,
			Slug: "k-tune",
			Name: "K-Tune",
		},
	},
	"CHIBI": {
		{
			ID:   10019,
			Slug: "chibi-inu",
			Name: "Chibi Inu",
		},
	},
	"SPARTA": {
		{
			ID:   6992,
			Slug: "spartan-protocol",
			Name: "Spartan Protocol",
		},
	},
	"CSP": {
		{
			ID:   3842,
			Slug: "caspian",
			Name: "Caspian",
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
	"XTK": {
		{
			ID:   8599,
			Slug: "xtoken",
			Name: "xToken",
		},
	},
	"MOGX": {
		{
			ID:   4800,
			Slug: "mogu",
			Name: "Mogu",
		},
	},
	"LIMIT": {
		{
			ID:   7379,
			Slug: "limitswap",
			Name: "LimitSwap",
		},
	},
	"BUCKS": {
		{
			ID:   9978,
			Slug: "buckswap",
			Name: "BuckSwap",
		},
	},
	"HORUS": {
		{
			ID:   2993,
			Slug: "horuspay",
			Name: "HorusPay",
		},
	},
	"UFARM": {
		{
			ID:   9262,
			Slug: "unifarm",
			Name: "UniFarm",
		},
	},
	"YPIE": {
		{
			ID:   8324,
			Slug: "piedao-yearn-ecosystem-pie",
			Name: "PieDAO Yearn Ecosystem Pie",
		},
	},
	"MOTA": {
		{
			ID:   4028,
			Slug: "motacoin",
			Name: "MotaCoin",
		},
	},
	"XPA": {
		{
			ID:   1968,
			Slug: "xpa",
			Name: "XPA",
		},
	},
	"DUKE": {
		{
			ID:   10552,
			Slug: "duke-inu-token",
			Name: "DUKE INU TOKEN",
		},
	},
	"RCN": {
		{
			ID:   2096,
			Slug: "ripio-credit-network",
			Name: "Ripio Credit Network",
		},
	},
	"IRA": {
		{
			ID:   4885,
			Slug: "diligence",
			Name: "Diligence",
		},
	},
	"TUSC": {
		{
			ID:   5131,
			Slug: "the-universal-settlement-coin",
			Name: "The Universal Settlement Coin",
		},
	},
	"SAM": {
		{
			ID:   7121,
			Slug: "samurai",
			Name: "Samurai",
		},
	},
	"RBBT": {
		{
			ID:   572,
			Slug: "rabbitcoin",
			Name: "RabbitCoin",
		},
	},
	"VTC": {
		{
			ID:   99,
			Slug: "vertcoin",
			Name: "Vertcoin",
		},
	},
	"LBC": {
		{
			ID:   1298,
			Slug: "library-credit",
			Name: "LBRY Credits",
		},
	},
	"NOV": {
		{
			ID:   8885,
			Slug: "novara-calcio-fan-token",
			Name: "Novara Calcio Fan Token",
		},
	},
	"KUV": {
		{
			ID:   4940,
			Slug: "kuverit",
			Name: "Kuverit",
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
	"CHIHUA": {
		{
			ID:   9947,
			Slug: "chihua-token",
			Name: "Chihua Token",
		},
	},
	"SCHO": {
		{
			ID:   8837,
			Slug: "scholarship-coin",
			Name: "Scholarship Coin",
		},
	},
	"BTC2": {
		{
			ID:   3974,
			Slug: "bitcoin2",
			Name: "Bitcoin 2",
		},
	},
	"SATA": {
		{
			ID:   9245,
			Slug: "signata",
			Name: "Signata",
		},
	},
	"CROAT": {
		{
			ID:   3087,
			Slug: "croat",
			Name: "CROAT",
		},
	},
	"N0001": {
		{
			ID:   8084,
			Slug: "nhbtc",
			Name: "nHBTC",
		},
	},
	"ELONONE": {
		{
			ID:   10166,
			Slug: "astroelon",
			Name: "AstroElon",
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
	"MELODY": {
		{
			ID:   10252,
			Slug: "mozart-finance",
			Name: "Mozart Finance",
		},
	},
	"1337": {
		{
			ID:   9484,
			Slug: "e1337",
			Name: "E1337",
		},
	},
	"NEX": {
		{
			ID:   3829,
			Slug: "nash-exchange",
			Name: "Nash",
		},
	},
	"GAX": {
		{
			ID:   10004,
			Slug: "getart",
			Name: "GETART",
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
	"MATH": {
		{
			ID:   5616,
			Slug: "math",
			Name: "MATH",
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
	"KRILL": {
		{
			ID:   9459,
			Slug: "krill",
			Name: "Krill",
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
	"FINIX": {
		{
			ID:   10661,
			Slug: "definix",
			Name: "Definix",
		},
	},
	"MCPC": {
		{
			ID:   4014,
			Slug: "mobile-crypto-pay-coin",
			Name: "Mobile Crypto Pay Coin",
		},
	},
	"OLYMPUS": {
		{
			ID:   10764,
			Slug: "olympus-token",
			Name: "OLYMPUS",
		},
	},
	"CRING": {
		{
			ID:   9243,
			Slug: "darwinia-crab-network",
			Name: "Darwinia Crab Network",
		},
	},
	"INB": {
		{
			ID:   3116,
			Slug: "insight-chain",
			Name: "Insight Chain",
		},
	},
	"PNK": {
		{
			ID:   3581,
			Slug: "kleros",
			Name: "Kleros",
		},
	},
	"IMG": {
		{
			ID:   4156,
			Slug: "imagecoin",
			Name: "ImageCoin",
		},
	},
	"7E": {
		{
			ID:   4063,
			Slug: "7eleven",
			Name: "7Eleven",
		},
	},
	"TGDY": {
		{
			ID:   10320,
			Slug: "tegridy",
			Name: "Tegridy",
		},
	},
	"PLC": {
		{
			ID:   3364,
			Slug: "platincoin",
			Name: "PLATINCOIN",
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
	"TKP": {
		{
			ID:   4116,
			Slug: "tokpie",
			Name: "TOKPIE",
		},
	},
	"FootballStars": {
		{
			ID:   9910,
			Slug: "football-stars",
			Name: "Football Stars",
		},
	},
	"BT": {
		{
			ID:   8558,
			Slug: "bt-finance",
			Name: "BT.Finance",
		},
	},
	"GT": {
		{
			ID:   4269,
			Slug: "gatetoken",
			Name: "GateToken",
		},
	},
	"CNFI": {
		{
			ID:   8227,
			Slug: "connect-financial",
			Name: "Connect Financial",
		},
	},
	"PolyMoon": {
		{
			ID:   9531,
			Slug: "polymoon",
			Name: "PolyMoon",
		},
	},
	"NYB": {
		{
			ID:   6981,
			Slug: "new-year-bull",
			Name: "New Year Bull",
		},
	},
	"PLUM": {
		{
			ID:   9773,
			Slug: "plumcake-finance",
			Name: "PlumCake Finance",
		},
	},
	"KPOP": {
		{
			ID:   10310,
			Slug: "kpop-fan-token",
			Name: "KPOP Fan Token",
		},
	},
	"IOG": {
		{
			ID:   3315,
			Slug: "playgroundz",
			Name: "Playgroundz",
		},
	},
	"DDX": {
		{
			ID:   7228,
			Slug: "derivadao",
			Name: "DerivaDAO",
		},
	},
	"CORA": {
		{
			ID:   9579,
			Slug: "corra-finance",
			Name: "Corra.Finance",
		},
	},
	"BRICK": {
		{
			ID:   8853,
			Slug: "brickchain-finance",
			Name: "Brickchain Finance",
		},
	},
	"MTDR": {
		{
			ID:   10078,
			Slug: "matador-token",
			Name: "Matador Token",
		},
	},
	"QTV": {
		{
			ID:   5695,
			Slug: "quish-coin",
			Name: "Quish Coin",
		},
	},
	"ERECT": {
		{
			ID:   10535,
			Slug: "evererected",
			Name: "EVERERECTED",
		},
	},
	"AEON": {
		{
			ID:   1026,
			Slug: "aeon",
			Name: "Aeon",
		},
	},
	"APN": {
		{
			ID:   8406,
			Slug: "apron-network",
			Name: "Apron Network",
		},
	},
	"PTERIA": {
		{
			ID:   7564,
			Slug: "pteria",
			Name: "Pteria",
		},
	},
	"BOOMB": {
		{
			ID:   10517,
			Slug: "boombaby-io",
			Name: "BoomBaby.io",
		},
	},
	"SBDO": {
		{
			ID:   8172,
			Slug: "bdollar-share",
			Name: "bDollar Share",
		},
	},
	"SUBX": {
		{
			ID:   10741,
			Slug: "startup-boost-token",
			Name: "Startup Boost Token",
		},
	},
	"HUB": {
		{
			ID:   7986,
			Slug: "hub-human-trust-protocol",
			Name: "Hub - Human Trust Protocol",
		},
	},
	"SISHI": {
		{
			ID:   9862,
			Slug: "sishi-finance",
			Name: "Sishi Finance",
		},
	},
	"YRISE": {
		{
			ID:   7441,
			Slug: "yrise-finance",
			Name: "yRise Finance",
		},
	},
	"LTX": {
		{
			ID:   7616,
			Slug: "lattice-token",
			Name: "Lattice Token",
		},
	},
	"WOLF": {
		{
			ID:   9158,
			Slug: "moonwolf",
			Name: "moonwolf.io",
		},
	},
	"CCX": {
		{
			ID:   3732,
			Slug: "conceal",
			Name: "Conceal",
		},
	},
	"KALM": {
		{
			ID:   10099,
			Slug: "kalmar",
			Name: "Kalmar",
		},
	},
	"GRG": {
		{
			ID:   4927,
			Slug: "rigoblock",
			Name: "RigoBlock",
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
	"BITSZ": {
		{
			ID:   9292,
			Slug: "bitsz",
			Name: "Bitsz",
		},
	},
	"VARC": {
		{
			ID:   5628,
			Slug: "varc",
			Name: "VARC",
		},
	},
	"CHIK": {
		{
			ID:   8518,
			Slug: "chickenkebab-finance",
			Name: "Chickenkebab Finance",
		},
	},
	"XQC": {
		{
			ID:   5220,
			Slug: "quras",
			Name: "QURAS",
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
	"CCXX": {
		{
			ID:   5482,
			Slug: "counos-x",
			Name: "Counos X",
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
	"AUC": {
		{
			ID:   2653,
			Slug: "auctus",
			Name: "Auctus",
		},
	},
	"PASS": {
		{
			ID:   3141,
			Slug: "blockpass",
			Name: "Blockpass",
		},
	},
	"GB": {
		{
			ID:   1285,
			Slug: "goldblocks",
			Name: "GoldBlocks",
		},
	},
	"FILUP": {
		{
			ID:   8050,
			Slug: "filup",
			Name: "FILUP",
		},
	},
	"XGC": {
		{
			ID:   10436,
			Slug: "xiglute-coin",
			Name: "Xiglute Coin",
		},
	},
	"SAFEBTC": {
		{
			ID:   8993,
			Slug: "safebtc",
			Name: "SafeBTC",
		},
	},
	"QHC": {
		{
			ID:   7243,
			Slug: "qchi-chain",
			Name: "QChi Chain",
		},
	},
	"G9TRO": {
		{
			ID:   9594,
			Slug: "g9tro-crowdfunding-platform",
			Name: "g9tro Crowdfunding Platform",
		},
	},
	"YFUEL": {
		{
			ID:   6913,
			Slug: "yfuel",
			Name: "YFUEL",
		},
	},
	"YOP": {
		{
			ID:   8187,
			Slug: "yop",
			Name: "Yield Optimization Platform & Protocol",
		},
	},
	"IRD": {
		{
			ID:   3457,
			Slug: "iridium",
			Name: "Iridium",
		},
	},
	"CLX": {
		{
			ID:   5179,
			Slug: "celeum",
			Name: "Celeum",
		},
	},
	"LIB": {
		{
			ID:   6700,
			Slug: "libera",
			Name: "Libera",
		},
	},
	"SACT": {
		{
			ID:   8753,
			Slug: "srnartgallery",
			Name: "srnArt Gallery",
		},
	},
	"SLR": {
		{
			ID:   233,
			Slug: "solarcoin",
			Name: "SolarCoin",
		},
	},
	"EGEM": {
		{
			ID:   3132,
			Slug: "ethergem",
			Name: "EtherGem",
		},
	},
	"WSB": {
		{
			ID:   9749,
			Slug: "wallstreetbets-dapp",
			Name: "WallStreetBets DApp",
		},
	},
	"SMBSWAP": {
		{
			ID:   7764,
			Slug: "simbcoin-swap",
			Name: "Simbcoin Swap",
		},
	},
	"HNT": {
		{
			ID:   5665,
			Slug: "helium",
			Name: "Helium",
		},
	},
	"WICC": {
		{
			ID:   2346,
			Slug: "waykichain",
			Name: "WaykiChain",
		},
	},
	"ATB": {
		{
			ID:   1970,
			Slug: "atbcoin",
			Name: "ATBCoin",
		},
	},
	"YFFII": {
		{
			ID:   6975,
			Slug: "yffii-finance",
			Name: "YFFII Finance",
		},
	},
	"ARION": {
		{
			ID:   3165,
			Slug: "arion",
			Name: "Arion",
		},
	},
	"JASMY": {
		{
			ID:   8425,
			Slug: "jasmy",
			Name: "Jasmy",
		},
	},
	"TESLF": {
		{
			ID:   10035,
			Slug: "teslafan",
			Name: "Teslafan",
		},
	},
	"HBO": {
		{
			ID:   8528,
			Slug: "hashbridge-oracle",
			Name: "HashBridge Oracle",
		},
	},
	"TBC": {
		{
			ID:   10128,
			Slug: "terablock",
			Name: "TeraBlock",
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
	"PLF": {
		{
			ID:   5084,
			Slug: "playfuel",
			Name: "PlayFuel",
		},
	},
	"XRT": {
		{
			ID:   4757,
			Slug: "robonomics-network",
			Name: "Robonomics.network",
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
	"BALLBAG": {
		{
			ID:   10543,
			Slug: "ballbag-token",
			Name: "Ballbag Token",
		},
	},
	"AVC": {
		{
			ID:   6376,
			Slug: "avccoin",
			Name: "AVCCOIN",
		},
	},
	"VOCO": {
		{
			ID:   3509,
			Slug: "provoco-token",
			Name: "Provoco Token",
		},
	},
	"RAISE": {
		{
			ID:   2755,
			Slug: "raise",
			Name: "Raise",
		},
	},
	"CXT": {
		{
			ID:   1630,
			Slug: "coinonat",
			Name: "Coinonat",
		},
	},
	"RAKUC": {
		{
			ID:   10638,
			Slug: "raku-coin",
			Name: "Raku Coin",
		},
	},
	"FME": {
		{
			ID:   5955,
			Slug: "fme",
			Name: "FME",
		},
	},
	"UNII": {
		{
			ID:   7207,
			Slug: "unii-finance",
			Name: "UNII Finance",
		},
	},
	"APES": {
		{
			ID:   10003,
			Slug: "apehaven",
			Name: "ApeHaven",
		},
	},
	"MARSH": {
		{
			ID:   8963,
			Slug: "unmarshal-token",
			Name: "UnMarshal",
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
	"COPS": {
		{
			ID:   8947,
			Slug: "cops-finance",
			Name: "COPS FINANCE",
		},
	},
	"UCAP": {
		{
			ID:   7819,
			Slug: "unicap-finance",
			Name: "Unicap.finance",
		},
	},
	"TRP": {
		{
			ID:   3939,
			Slug: "tronipay",
			Name: "Tronipay",
		},
	},
	"WISE": {
		{
			ID:   8167,
			Slug: "wise",
			Name: "Wise Token",
		},
	},
	"SAFEBULL": {
		{
			ID:   10644,
			Slug: "safebull",
			Name: "SafeBull",
		},
	},
	"YAS": {
		{
			ID:   6246,
			Slug: "yas",
			Name: "YAS",
		},
	},
	"CCT": {
		{
			ID:   5054,
			Slug: "coupon-chain",
			Name: "Coupon Chain",
		},
	},
	"BMXX": {
		{
			ID:   8618,
			Slug: "bmultiplier",
			Name: "Multiplier",
		},
	},
	"IFX24": {
		{
			ID:   5170,
			Slug: "ifx24",
			Name: "IFX24",
		},
	},
	"SAFESTAR": {
		{
			ID:   8872,
			Slug: "safe-star",
			Name: "Safe Star",
		},
	},
	"BASK": {
		{
			ID:   9157,
			Slug: "basketdao",
			Name: "BasketDAO",
		},
	},
	"ELNC": {
		{
			ID:   9858,
			Slug: "eloniumcoin",
			Name: "EloniumCoin",
		},
	},
	"RC20": {
		{
			ID:   3790,
			Slug: "robocalls",
			Name: "RoboCalls",
		},
	},
	"LN": {
		{
			ID:   4512,
			Slug: "link",
			Name: "LINK",
		},
	},
	"FCR": {
		{
			ID:   7707,
			Slug: "fromm-car",
			Name: "Fromm Car",
		},
	},
	"LONG": {
		{
			ID:   7668,
			Slug: "long-coin",
			Name: "LONG COIN",
		},
	},
	"ZEFI": {
		{
			ID:   8759,
			Slug: "zcore-finance",
			Name: "ZCore Finance",
		},
	},
	"BITO": {
		{
			ID:   6118,
			Slug: "bitopro-exchange-token",
			Name: "BitoPro Exchange Token",
		},
	},
	"TRXC": {
		{
			ID:   3354,
			Slug: "tronclassic",
			Name: "TRONCLASSIC",
		},
	},
	"QUIN": {
		{
			ID:   3659,
			Slug: "quinads",
			Name: "QUINADS",
		},
	},
	"ZUC": {
		{
			ID:   4250,
			Slug: "zeuxcoin",
			Name: "ZeuxCoin",
		},
	},
	"WINDY": {
		{
			ID:   9560,
			Slug: "windswap",
			Name: "WindSwap",
		},
	},
	"WXT": {
		{
			ID:   4090,
			Slug: "wirex-token",
			Name: "Wirex Token",
		},
	},
	"METIS": {
		{
			ID:   9640,
			Slug: "metisdao",
			Name: "Metis",
		},
	},
	"ICNQ": {
		{
			ID:   3431,
			Slug: "iconic-token",
			Name: "Iconic Token",
		},
	},
	"LINKUP": {
		{
			ID:   7011,
			Slug: "linkup",
			Name: "LINKUP",
		},
	},
	"EA": {
		{
			ID:   6427,
			Slug: "ea-token",
			Name: "EA Token",
		},
	},
	"SERGS": {
		{
			ID:   7595,
			Slug: "sergs",
			Name: "SERGS",
		},
	},
	"DIESEL": {
		{
			ID:   8792,
			Slug: "diesel",
			Name: "DIESEL",
		},
	},
	"BNANA": {
		{
			ID:   3742,
			Slug: "chimpion",
			Name: "Chimpion",
		},
	},
	"VIDY": {
		{
			ID:   4431,
			Slug: "vidy",
			Name: "VIDY",
		},
	},
	"SWINGBY": {
		{
			ID:   5922,
			Slug: "swingby",
			Name: "Swingby",
		},
	},
	"DFY": {
		{
			ID:   9179,
			Slug: "defi-for-you",
			Name: "Defi For You",
		},
	},
	"C3": {
		{
			ID:   9991,
			Slug: "charli3",
			Name: "Charli3",
		},
	},
	"SLAM": {
		{
			ID:   9584,
			Slug: "slam-token",
			Name: "Slam Token",
		},
	},
	"ZOOT": {
		{
			ID:   10005,
			Slug: "zoo-token",
			Name: "Zoo Token",
		},
	},
	"YUP": {
		{
			ID:   7689,
			Slug: "yup-token",
			Name: "Yup",
		},
	},
	"vADA": {
		{
			ID:   9428,
			Slug: "venus-cardano",
			Name: "Venus Cardano",
		},
	},
	"BNB": {
		{
			ID:   1839,
			Slug: "binance-coin",
			Name: "Binance Coin",
		},
	},
	"EMC2": {
		{
			ID:   201,
			Slug: "einsteinium",
			Name: "Einsteinium",
		},
	},
	"GSC": {
		{
			ID:   2737,
			Slug: "global-social-chain",
			Name: "Global Social Chain",
		},
	},
	"AUTZ": {
		{
			ID:   10660,
			Slug: "autz-token",
			Name: "AUTZ Token",
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
	"DELTA": {
		{
			ID:   8994,
			Slug: "delta-finance",
			Name: "Delta",
		},
	},
	"MKCY": {
		{
			ID:   7113,
			Slug: "markaccy",
			Name: "Markaccy",
		},
	},
	"SLRS": {
		{
			ID:   9989,
			Slug: "solrise-finance",
			Name: "Solrise Finance",
		},
	},
	"GDR": {
		{
			ID:   4881,
			Slug: "guider",
			Name: "Guider",
		},
	},
	"BDX": {
		{
			ID:   3987,
			Slug: "beldex",
			Name: "Beldex",
		},
	},
	"PAL": {
		{
			ID:   8207,
			Slug: "playandlike",
			Name: "PlayAndLike",
		},
	},
	"MANDI": {
		{
			ID:   6532,
			Slug: "mandi-token",
			Name: "Mandi Token",
		},
	},
	"VAIP": {
		{
			ID:   5619,
			Slug: "vehicle-data-artificial-intelligence-platform",
			Name: "VEHICLE DATA ARTIFICIAL INTELLIGENCE PLATFORM",
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
	"DASHG": {
		{
			ID:   3613,
			Slug: "dash-green",
			Name: "Dash Green",
		},
	},
	"ABUSD": {
		{
			ID:   5755,
			Slug: "aave-busd",
			Name: "Aave BUSD",
		},
	},
	"TRTL": {
		{
			ID:   2958,
			Slug: "turtlecoin",
			Name: "TurtleCoin",
		},
	},
	"NIOX": {
		{
			ID:   2151,
			Slug: "autonio",
			Name: "Autonio",
		},
	},
	"REPO": {
		{
			ID:   2829,
			Slug: "repo",
			Name: "REPO",
		},
	},
	"TRC": {
		{
			ID:   4,
			Slug: "terracoin",
			Name: "Terracoin",
		},
	},
	"ERTH": {
		{
			ID:   10171,
			Slug: "erth-token",
			Name: "ERTH Token",
		},
	},
	"YI12": {
		{
			ID:   6912,
			Slug: "yield-stake-finance",
			Name: "Yield Stake Finance",
		},
	},
	"BAKE": {
		{
			ID:   7064,
			Slug: "bakerytoken",
			Name: "BakeryToken",
		},
	},
	"LGO": {
		{
			ID:   2600,
			Slug: "lgo-token",
			Name: "LGO Token",
		},
	},
	"HZN": {
		{
			ID:   9237,
			Slug: "horizon-protocol",
			Name: "Horizon Protocol",
		},
	},
	"WOMI": {
		{
			ID:   9126,
			Slug: "wrapped-ecomi",
			Name: "Wrapped ECOMI",
		},
	},
	"SQ": {
		{
			ID:   7902,
			Slug: "square-tokenized-stock-ftx",
			Name: "Square tokenized stock FTX",
		},
	},
	"PIPT": {
		{
			ID:   7368,
			Slug: "power-index-pool-token",
			Name: "Power Index Pool Token",
		},
	},
	"SCL": {
		{
			ID:   1969,
			Slug: "sociall",
			Name: "Sociall",
		},
	},
	"MNR": {
		{
			ID:   6053,
			Slug: "mineral",
			Name: "Mineral",
		},
	},
	"BLURT": {
		{
			ID:   7570,
			Slug: "blurt",
			Name: "Blurt",
		},
	},
	"MBS": {
		{
			ID:   9405,
			Slug: "moonboys",
			Name: "MoonBoys",
		},
	},
	"VLX": {
		{
			ID:   4747,
			Slug: "velas",
			Name: "Velas",
		},
	},
	"EVED": {
		{
			ID:   3953,
			Slug: "evedo",
			Name: "Evedo",
		},
	},
	"JUS": {
		{
			ID:   6483,
			Slug: "just-network",
			Name: "JUST NETWORK",
		},
	},
	"BOTX": {
		{
			ID:   3873,
			Slug: "botxcoin",
			Name: "botXcoin",
		},
	},
	"HUNT": {
		{
			ID:   5380,
			Slug: "hunt",
			Name: "HUNT",
		},
	},
	"DIREWOLF": {
		{
			ID:   10266,
			Slug: "direwolf",
			Name: "Direwolf",
		},
	},
	"DSYS": {
		{
			ID:   6426,
			Slug: "dsys",
			Name: "DSYS",
		},
	},
	"UNIDX": {
		{
			ID:   8232,
			Slug: "unidex",
			Name: "UniDex",
		},
	},
	"ROOK": {
		{
			ID:   7678,
			Slug: "keeperdao",
			Name: "KeeperDAO",
		},
	},
	"OM": {
		{
			ID:   6536,
			Slug: "mantra-dao",
			Name: "MANTRA DAO",
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
	"MERI": {
		{
			ID:   3915,
			Slug: "merebel",
			Name: "Merebel",
		},
	},
	"POLVEN": {
		{
			ID:   8974,
			Slug: "polka-ventures",
			Name: "Polka Ventures",
		},
	},
	"NMT": {
		{
			ID:   10033,
			Slug: "nftmart-token",
			Name: "NFTMart Token",
		},
	},
	"MONSTA": {
		{
			ID:   10366,
			Slug: "cake-monster",
			Name: "Cake Monster",
		},
	},
	"MOONLIGHT": {
		{
			ID:   9848,
			Slug: "moonlight-token",
			Name: "Moonlight Token",
		},
	},
	"SGC": {
		{
			ID:   5489,
			Slug: "sudan-gold-coin",
			Name: "Sudan Gold Coin",
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
	"BSBT": {
		{
			ID:   8623,
			Slug: "bsb-token",
			Name: "BSB Token",
		},
	},
	"DAILY": {
		{
			ID:   8910,
			Slug: "daily",
			Name: "Daily",
		},
	},
	"GASG": {
		{
			ID:   8352,
			Slug: "gasgains",
			Name: "Gasgains",
		},
	},
	"XWC": {
		{
			ID:   268,
			Slug: "whitecoin",
			Name: "WhiteCoin",
		},
	},
	"TOP": {
		{
			ID:   3826,
			Slug: "top",
			Name: "TOP",
		},
	},
	"XDN": {
		{
			ID:   405,
			Slug: "digitalnote",
			Name: "DigitalNote",
		},
	},
	"MCAN": {
		{
			ID:   8559,
			Slug: "medican-coin",
			Name: "Medican Coin",
		},
	},
	"KIP": {
		{
			ID:   6489,
			Slug: "khipu-token",
			Name: "Khipu Token",
		},
	},
	"CLIT$": {
		{
			ID:   10571,
			Slug: "clit-token-protocol",
			Name: "CLIT TOKEN PROTOCOL",
		},
	},
	"KABOSU": {
		{
			ID:   9710,
			Slug: "kabosu",
			Name: "Kabosu",
		},
	},
	"DCY": {
		{
			ID:   1745,
			Slug: "dinastycoin",
			Name: "Dinastycoin",
		},
	},
	"SOLVE": {
		{
			ID:   3724,
			Slug: "solve",
			Name: "SOLVE",
		},
	},
	"CAJ": {
		{
			ID:   3715,
			Slug: "cajutel",
			Name: "Cajutel",
		},
	},
	"CTF": {
		{
			ID:   8791,
			Slug: "cybertime-finance-token",
			Name: "CyberTime Finance Token",
		},
	},
	"DIA": {
		{
			ID:   6138,
			Slug: "dia",
			Name: "DIA",
		},
	},
	"PRX": {
		{
			ID:   3668,
			Slug: "proxynode",
			Name: "ProxyNode",
		},
	},
	"AZX": {
		{
			ID:   9917,
			Slug: "azeusx",
			Name: "AzeusX",
		},
	},
	"7ADD": {
		{
			ID:   7771,
			Slug: "holdtowin",
			Name: "HoldToWin",
		},
	},
	"HPT": {
		{
			ID:   3721,
			Slug: "huobi-pool-token",
			Name: "Huobi Pool Token",
		},
	},
	"PFL": {
		{
			ID:   9172,
			Slug: "professional-fighters-league-fan-token",
			Name: "Professional Fighters League Fan Token",
		},
	},
	"XTP": {
		{
			ID:   5070,
			Slug: "tap",
			Name: "Tap",
		},
	},
	"MLTP": {
		{
			ID:   10576,
			Slug: "moonlift-protocol",
			Name: "MoonLift Protocol",
		},
	},
	"HAPI": {
		{
			ID:   8567,
			Slug: "hapi-one",
			Name: "HAPI",
		},
	},
	"LUA": {
		{
			ID:   7216,
			Slug: "lua-token",
			Name: "LuaSwap",
		},
	},
	"FAN": {
		{
			ID:   10609,
			Slug: "fanspel",
			Name: "Fanspel",
		},
	},
	"SAK3": {
		{
			ID:   10608,
			Slug: "sake",
			Name: "Sake",
		},
	},
	"PHIBA": {
		{
			ID:   9768,
			Slug: "papa-shiba",
			Name: "Papa Shiba",
		},
	},
	"HOPE": {
		{
			ID:   9480,
			Slug: "hope-token",
			Name: "Hope",
		},
	},
	"XAG": {
		{
			ID:   6558,
			Slug: "xrpalike-gene",
			Name: "Xrpalike Gene",
		},
	},
	"STARINU": {
		{
			ID:   11014,
			Slug: "starship-inu",
			Name: "Starship Inu",
		},
	},
	"NBU": {
		{
			ID:   8939,
			Slug: "nimbus",
			Name: "Nimbus",
		},
	},
	"BA": {
		{
			ID:   9167,
			Slug: "batorrent",
			Name: "BaTorrent",
		},
	},
	"BBGC": {
		{
			ID:   3832,
			Slug: "big-bang-game-coin",
			Name: "Big Bang Game Coin",
		},
	},
	"HYVE": {
		{
			ID:   7552,
			Slug: "hyve",
			Name: "Hyve",
		},
	},
	"GIG": {
		{
			ID:   3998,
			Slug: "krios",
			Name: "Krios",
		},
	},
	"HYC": {
		{
			ID:   3147,
			Slug: "hycon",
			Name: "HYCON",
		},
	},
	"PGO": {
		{
			ID:   5381,
			Slug: "pengolincoin",
			Name: "PengolinCoin",
		},
	},
	"DIAH": {
		{
			ID:   10459,
			Slug: "diarrheacoin",
			Name: "DiarrheaCoin",
		},
	},
	"MOONI": {
		{
			ID:   9465,
			Slug: "mooni-defi",
			Name: "Mooni DeFi",
		},
	},
	"UNSAFEMOON": {
		{
			ID:   10198,
			Slug: "unsafemoon",
			Name: "UnSafeMoon",
		},
	},
	"NMR": {
		{
			ID:   1732,
			Slug: "numeraire",
			Name: "Numeraire",
		},
	},
	"BITCI": {
		{
			ID:   8357,
			Slug: "bitcicoin",
			Name: "Bitcicoin",
		},
	},
	"BCD": {
		{
			ID:   2222,
			Slug: "bitcoin-diamond",
			Name: "Bitcoin Diamond",
		},
	},
	"ERO": {
		{
			ID:   2249,
			Slug: "eroscoin",
			Name: "Eroscoin",
		},
	},
	"HVI": {
		{
			ID:   10800,
			Slug: "hungarian-vizsla-inu",
			Name: "Hungarian Vizsla Inu",
		},
	},
	"X": {
		{
			ID:   10213,
			Slug: "x-by-spacegrime",
			Name: "X (By SpaceGrime)",
		},
	},
	"aEth": {
		{
			ID:   8100,
			Slug: "ankreth",
			Name: "ankrETH",
		},
	},
	"ANCT": {
		{
			ID:   4901,
			Slug: "anchor",
			Name: "Anchor",
		},
	},
	"CONX": {
		{
			ID:   1632,
			Slug: "concoin",
			Name: "Concoin",
		},
	},
	"INST": {
		{
			ID:   10508,
			Slug: "instadapp",
			Name: "Instadapp",
		},
	},
	"COFE": {
		{
			ID:   8859,
			Slug: "coffeeswap",
			Name: "CoffeeSwap",
		},
	},
	"AVAX": {
		{
			ID:   5805,
			Slug: "avalanche",
			Name: "Avalanche",
		},
	},
	"MINTME": {
		{
			ID:   3361,
			Slug: "mintme-com-coin",
			Name: "MintMe.com Coin",
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
	"CUBE": {
		{
			ID:   5338,
			Slug: "somnium-space-cubes",
			Name: "Somnium Space Cubes",
		},
	},
	"BSW": {
		{
			ID:   10746,
			Slug: "biswap",
			Name: "Biswap",
		},
	},
	"JMC": {
		{
			ID:   4324,
			Slug: "junsonmingchncoin",
			Name: "Junsonmingchncoin",
		},
	},
	"GGIVE": {
		{
			ID:   9815,
			Slug: "globalgive",
			Name: "GlobalGive",
		},
	},
	"WAN": {
		{
			ID:   2606,
			Slug: "wanchain",
			Name: "Wanchain",
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
	"FREN": {
		{
			ID:   9767,
			Slug: "frenchie-network",
			Name: "Frenchie Network",
		},
	},
	"AMPL": {
		{
			ID:   4056,
			Slug: "ampleforth",
			Name: "Ampleforth",
		},
	},
	"MTA": {
		{
			ID:   5748,
			Slug: "meta",
			Name: "mStable Governance Token: Meta (MTA)",
		},
	},
	"LEASH": {
		{
			ID:   9286,
			Slug: "doge-killer",
			Name: "Doge Killer",
		},
	},
	"AHF": {
		{
			ID:   7241,
			Slug: "americanhorror-finance",
			Name: "AmericanHorror.Finance",
		},
	},
	"AABC": {
		{
			ID:   9541,
			Slug: "aabc-token",
			Name: "AABC Token",
		},
	},
	"FBN": {
		{
			ID:   3468,
			Slug: "fivebalance",
			Name: "Fivebalance",
		},
	},
	"PIPI": {
		{
			ID:   9502,
			Slug: "pippi-finance",
			Name: "Pippi Finance",
		},
	},
	"NEAL": {
		{
			ID:   3960,
			Slug: "coineal-token",
			Name: "Coineal Token",
		},
	},
	"FORTH": {
		{
			ID:   9421,
			Slug: "ampleforth-governance-token",
			Name: "Ampleforth Governance Token",
		},
	},
	"FCX": {
		{
			ID:   8222,
			Slug: "fission-cash",
			Name: "Fission Cash",
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
	"QBC": {
		{
			ID:   278,
			Slug: "quebecoin",
			Name: "Quebecoin",
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
	"DEFO": {
		{
			ID:   7720,
			Slug: "defhold",
			Name: "DefHold",
		},
	},
	"FLL": {
		{
			ID:   6410,
			Slug: "feellike",
			Name: "Feellike",
		},
	},
	"LKT": {
		{
			ID:   10775,
			Slug: "locklet",
			Name: "Locklet",
		},
	},
	"OSM": {
		{
			ID:   10177,
			Slug: "supermoon",
			Name: "Supermoon",
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
	"PPDEX": {
		{
			ID:   7592,
			Slug: "pepedex",
			Name: "Pepedex",
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
	"LTRBT": {
		{
			ID:   10051,
			Slug: "little-rabbit",
			Name: "LITTLE RABBIT",
		},
	},
	"DAFT": {
		{
			ID:   8913,
			Slug: "daftcoin",
			Name: "DaftCoin",
		},
	},
	"BDOGE": {
		{
			ID:   9731,
			Slug: "blue-eyes-white-doge",
			Name: "Blue Eyes White Doge",
		},
	},
	"GDOGE": {
		{
			ID:   10227,
			Slug: "gdoge-finance",
			Name: "GDOGE Finance",
		},
	},
	"WASP": {
		{
			ID:   8147,
			Slug: "wanswap",
			Name: "WanSwap",
		},
	},
	"WOLFY": {
		{
			ID:   10963,
			Slug: "wolfystreetbets",
			Name: "Wolfystreetbets",
		},
	},
	"USHIBA": {
		{
			ID:   9943,
			Slug: "american-shiba",
			Name: "American Shiba",
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
	"MRNA": {
		{
			ID:   7900,
			Slug: "moderna-tokenized-stock-ftx",
			Name: "Moderna tokenized stock FTX",
		},
	},
	"ZIL": {
		{
			ID:   2469,
			Slug: "zilliqa",
			Name: "Zilliqa",
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
	"BBK": {
		{
			ID:   3051,
			Slug: "bitblocks",
			Name: "Bitblocks",
		},
	},
	"DATP": {
		{
			ID:   3454,
			Slug: "decentralized-asset-trading-platform",
			Name: "Decentralized Asset Trading Platform",
		},
	},
	"IMS": {
		{
			ID:   1194,
			Slug: "independent-money-system",
			Name: "Independent Money System",
		},
	},
	"GHX": {
		{
			ID:   6554,
			Slug: "gamercoin",
			Name: "GamerCoin",
		},
	},
	"HZT": {
		{
			ID:   6214,
			Slug: "black-diamond-rating",
			Name: "Black Diamond Rating",
		},
	},
	"ZTB": {
		{
			ID:   9372,
			Slug: "ztb",
			Name: "ZTB",
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
	"SNM": {
		{
			ID:   9931,
			Slug: "sonm-bep20",
			Name: "SONM (BEP-20)",
		},
	},
	"KIF": {
		{
			ID:   6883,
			Slug: "kittenfinance",
			Name: "KittenFinance",
		},
	},
	"ZYRO": {
		{
			ID:   7044,
			Slug: "zyro",
			Name: "Zyro",
		},
	},
	"wxDai": {
		{
			ID:   9021,
			Slug: "wxdai",
			Name: "Wrapped XDAI",
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
	"BEAM": {
		{
			ID:   3702,
			Slug: "beam",
			Name: "Beam",
		},
	},
	"CNTM": {
		{
			ID:   6039,
			Slug: "connectome",
			Name: "Connectome",
		},
	},
	"MARKGOAT": {
		{
			ID:   10437,
			Slug: "mark-goat",
			Name: "Mark Goat",
		},
	},
	"SKEY": {
		{
			ID:   8133,
			Slug: "smartkey",
			Name: "SmartKey",
		},
	},
	"WSG": {
		{
			ID:   10040,
			Slug: "wall-street-games",
			Name: "Wall Street Games",
		},
	},
	"SENSO": {
		{
			ID:   5522,
			Slug: "senso",
			Name: "SENSO",
		},
	},
	"ETI": {
		{
			ID:   3599,
			Slug: "etherinc",
			Name: "EtherInc",
		},
	},
	"XSUSHI": {
		{
			ID:   8899,
			Slug: "xsushi",
			Name: "xSUSHI",
		},
	},
	"RAKU": {
		{
			ID:   5334,
			Slug: "rakun",
			Name: "RAKUN",
		},
	},
	"RM": {
		{
			ID:   5814,
			Slug: "rivermount",
			Name: "Rivermount",
		},
	},
	"SIGN": {
		{
			ID:   3868,
			Slug: "signature-chain",
			Name: "Signature Chain",
		},
	},
	"POSS": {
		{
			ID:   3583,
			Slug: "posscoin",
			Name: "Posscoin",
		},
	},
	"ETHBULL": {
		{
			ID:   5217,
			Slug: "3x-long-ethereum-token",
			Name: "3X Long Ethereum Token",
		},
	},
	"YFIS": {
		{
			ID:   6807,
			Slug: "yfiscurity",
			Name: "YFISCURITY",
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
	"NEWINU": {
		{
			ID:   10371,
			Slug: "newinu",
			Name: "Newinu",
		},
	},
	"SAFERMOON": {
		{
			ID:   9761,
			Slug: "safermoon",
			Name: "SaferMoon",
		},
	},
	"IDL": {
		{
			ID:   8142,
			Slug: "idl-token",
			Name: "IDL Token",
		},
	},
	"USDQ": {
		{
			ID:   4020,
			Slug: "usdq",
			Name: "USDQ",
		},
	},
	"YFIKING": {
		{
			ID:   6964,
			Slug: "yfiking-finance",
			Name: "YFIKING,FINANCE",
		},
	},
	"SKYLARK": {
		{
			ID:   10039,
			Slug: "skylark",
			Name: "SKYLARK",
		},
	},
	"HTML": {
		{
			ID:   2315,
			Slug: "html-coin",
			Name: "HTMLCOIN",
		},
	},
	"CUSDC": {
		{
			ID:   5265,
			Slug: "compound-usd-coin",
			Name: "Compound USD Coin",
		},
	},
	"ATFI": {
		{
			ID:   9844,
			Slug: "atlantic-finance-token",
			Name: "Atlantic Finance Token",
		},
	},
	"wCRES": {
		{
			ID:   7818,
			Slug: "wrapped-crescofin",
			Name: "Wrapped CrescoFin",
		},
	},
	"SAP": {
		{
			ID:   7584,
			Slug: "swapall",
			Name: "SwapAll",
		},
	},
	"YT": {
		{
			ID:   6395,
			Slug: "cherry-token",
			Name: "Cherry Token",
		},
	},
	"BRC": {
		{
			ID:   3653,
			Slug: "baer-chain",
			Name: "Baer Chain",
		},
	},
	"OPTI": {
		{
			ID:   3101,
			Slug: "optitoken",
			Name: "OptiToken",
		},
	},
	"SETS": {
		{
			ID:   9785,
			Slug: "sensitrust",
			Name: "Sensitrust",
		},
	},
	"ZABAKU": {
		{
			ID:   10032,
			Slug: "zabaku-inu",
			Name: "ZABAKU INU",
		},
	},
	"CHZ006": {
		{
			ID:   8990,
			Slug: "charizardtoken",
			Name: "Charizard Token",
		},
	},
	"PAX": {
		{
			ID:   3330,
			Slug: "paxos-standard",
			Name: "Paxos Standard",
		},
	},
	"QC": {
		{
			ID:   2319,
			Slug: "qcash",
			Name: "Qcash",
		},
	},
	"STBU": {
		{
			ID:   7245,
			Slug: "stobox-token",
			Name: "Stobox Token",
		},
	},
	"TMN": {
		{
			ID:   4102,
			Slug: "translateme-network-token",
			Name: "TranslateMe Network Token",
		},
	},
	"YFIEC": {
		{
			ID:   7275,
			Slug: "yearn-finance-ecosystem",
			Name: "Yearn Finance Ecosystem",
		},
	},
	"WABI": {
		{
			ID:   2267,
			Slug: "wabi",
			Name: "Wabi",
		},
	},
	"BGL": {
		{
			ID:   5667,
			Slug: "bitgesell",
			Name: "Bitgesell",
		},
	},
	"VKF": {
		{
			ID:   7644,
			Slug: "vkf-platform",
			Name: "VKF Platform",
		},
	},
	"MD+": {
		{
			ID:   8392,
			Slug: "moondayplus",
			Name: "MoonDayPlus",
		},
	},
	"HXRO": {
		{
			ID:   3748,
			Slug: "hxro",
			Name: "Hxro",
		},
	},
	"KLV": {
		{
			ID:   6724,
			Slug: "klever",
			Name: "Klever",
		},
	},
	"EUM": {
		{
			ID:   3968,
			Slug: "elitium",
			Name: "Elitium",
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
	"AIDI": {
		{
			ID:   10692,
			Slug: "aidi-finance",
			Name: "Aidi Finance",
		},
	},
	"GIGA": {
		{
			ID:   9897,
			Slug: "gigapool",
			Name: "GigaPool",
		},
	},
	"SXAU": {
		{
			ID:   6191,
			Slug: "sxau",
			Name: "sXAU",
		},
	},
	"BMX": {
		{
			ID:   2933,
			Slug: "bitmart-token",
			Name: "BitMart Token",
		},
	},
	"ESS": {
		{
			ID:   2906,
			Slug: "essentia",
			Name: "Essentia",
		},
	},
	"POSW": {
		{
			ID:   1495,
			Slug: "posw-coin",
			Name: "PoSW Coin",
		},
	},
	"IDRT": {
		{
			ID:   4702,
			Slug: "rupiah-token",
			Name: "Rupiah Token",
		},
	},
	"ECTE": {
		{
			ID:   3741,
			Slug: "eurocoin-token",
			Name: "EurocoinToken",
		},
	},
	"YSR": {
		{
			ID:   5779,
			Slug: "ystar",
			Name: "Ystar",
		},
	},
	"FMT": {
		{
			ID:   9116,
			Slug: "finminity",
			Name: "Finminity",
		},
	},
	"LPT": {
		{
			ID:   3640,
			Slug: "livepeer",
			Name: "Livepeer",
		},
	},
	"DAPS": {
		{
			ID:   3345,
			Slug: "daps-coin",
			Name: "DAPS Coin",
		},
	},
	"LIMEX": {
		{
			ID:   5985,
			Slug: "limestone-network",
			Name: "Limestone Network",
		},
	},
	"IMO": {
		{
			ID:   9281,
			Slug: "imo",
			Name: "IMO",
		},
	},
	"DVI": {
		{
			ID:   7590,
			Slug: "dvision-network",
			Name: "Dvision Network",
		},
	},
	"IDEX": {
		{
			ID:   3928,
			Slug: "idex",
			Name: "IDEX",
		},
	},
	"CARD": {
		{
			ID:   2891,
			Slug: "cardstack",
			Name: "Cardstack",
		},
	},
	"LCMS": {
		{
			ID:   9033,
			Slug: "lcms",
			Name: "LCMS",
		},
	},
	"WISH": {
		{
			ID:   2236,
			Slug: "mywish",
			Name: "MyWish",
		},
	},
	"BOXER": {
		{
			ID:   10972,
			Slug: "boxer-inu",
			Name: "Boxer Inu",
		},
	},
	"WOOF": {
		{
			ID:   10857,
			Slug: "shibance",
			Name: "Shibance",
		},
	},
	"MDM": {
		{
			ID:   4928,
			Slug: "medium",
			Name: "Medium",
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
	"SOV": {
		{
			ID:   8669,
			Slug: "sovryn",
			Name: "Sovryn",
		},
	},
	"LPNT": {
		{
			ID:   8501,
			Slug: "luxurious-pro-network-token",
			Name: "Luxurious Pro Network Token",
		},
	},
	"HBN": {
		{
			ID:   78,
			Slug: "hobonickels",
			Name: "HoboNickels",
		},
	},
	"SFJP": {
		{
			ID:   9880,
			Slug: "safejupiter-sfjp",
			Name: "SafeJupiter $SFJP",
		},
	},
	"PINKE": {
		{
			ID:   9929,
			Slug: "pinkelon",
			Name: "PinkElon",
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
	"LQTY": {
		{
			ID:   7429,
			Slug: "liquity",
			Name: "Liquity",
		},
	},
	"FLP": {
		{
			ID:   2707,
			Slug: "flip",
			Name: "FLIP",
		},
	},
	"AIRX": {
		{
			ID:   4552,
			Slug: "aircoins",
			Name: "Aircoins",
		},
	},
	"NDSK": {
		{
			ID:   9677,
			Slug: "nadeshiko",
			Name: "Nadeshiko",
		},
	},
	"TRB": {
		{
			ID:   4944,
			Slug: "tellor",
			Name: "Tellor",
		},
	},
	"URUS": {
		{
			ID:   8616,
			Slug: "urus",
			Name: "Aurox",
		},
	},
	"LIME": {
		{
			ID:   10469,
			Slug: "ime-lab",
			Name: "iMe Lab",
		},
	},
	"DXH": {
		{
			ID:   9317,
			Slug: "daxhund",
			Name: "Daxhund",
		},
	},
	"XVG": {
		{
			ID:   693,
			Slug: "verge",
			Name: "Verge",
		},
	},
	"WIN": {
		{
			ID:   4206,
			Slug: "wink",
			Name: "WINkLink",
		},
	},
	"EMON": {
		{
			ID:   9651,
			Slug: "ethermon",
			Name: "Ethermon",
		},
	},
	"BCHBULL": {
		{
			ID:   5467,
			Slug: "3x-long-bitcoin-cash-token",
			Name: "3x Long Bitcoin Cash Token",
		},
	},
	"PoSH": {
		{
			ID:   7112,
			Slug: "shill-win",
			Name: "Shill & Win",
		},
	},
	"PHOENIX": {
		{
			ID:   10906,
			Slug: "phoenix-force",
			Name: "PHOENIX FORCE",
		},
	},
	"HX": {
		{
			ID:   4895,
			Slug: "hyperexchange",
			Name: "HyperExchange",
		},
	},
	"ATM": {
		{
			ID:   5227,
			Slug: "atletico-de-madrid-fan-token",
			Name: "Atletico De Madrid Fan Token",
		},
	},
	"GABECOIN": {
		{
			ID:   10926,
			Slug: "gabecoin",
			Name: "Gabecoin",
		},
	},
	"CC10": {
		{
			ID:   8887,
			Slug: "cryptocurrency-top-10-tokens-index",
			Name: "Cryptocurrency Top 10 Tokens Index",
		},
	},
	"SPKL": {
		{
			ID:   7198,
			Slug: "spoklottery",
			Name: "SpokLottery",
		},
	},
	"KICK": {
		{
			ID:   10700,
			Slug: "kick-token-new",
			Name: "KickToken [new]",
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
	"RET": {
		{
			ID:   3280,
			Slug: "realtract",
			Name: "RealTract",
		},
	},
	"CREP": {
		{
			ID:   5746,
			Slug: "compound-augur",
			Name: "Compound Augur",
		},
	},
	"NFTFY": {
		{
			ID:   9622,
			Slug: "nftfy",
			Name: "Nftfy",
		},
	},
	"REN": {
		{
			ID:   2539,
			Slug: "ren",
			Name: "Ren",
		},
	},
	"BLZ": {
		{
			ID:   2505,
			Slug: "bluzelle",
			Name: "Bluzelle",
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
	"TEND": {
		{
			ID:   5971,
			Slug: "tendies",
			Name: "Tendies",
		},
	},
	"AERGO": {
		{
			ID:   3637,
			Slug: "aergo",
			Name: "Aergo",
		},
	},
	"VNLA": {
		{
			ID:   7776,
			Slug: "vanilla-network",
			Name: "Vanilla Network",
		},
	},
	"MEXC": {
		{
			ID:   4676,
			Slug: "mexc-token",
			Name: "MEXC Token",
		},
	},
	"HVE2": {
		{
			ID:   9027,
			Slug: "uhive",
			Name: "Uhive",
		},
	},
	"AIDUS": {
		{
			ID:   3785,
			Slug: "aidus-token",
			Name: "AIDUS TOKEN",
		},
	},
	"AION": {
		{
			ID:   2062,
			Slug: "aion",
			Name: "Aion",
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
	"SAFEZONE": {
		{
			ID:   9799,
			Slug: "safezone",
			Name: "SafeZone",
		},
	},
	"NCAT": {
		{
			ID:   8959,
			Slug: "ncat-token",
			Name: "NCAT Token",
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
	"LUD": {
		{
			ID:   6527,
			Slug: "ludos",
			Name: "Ludos Protocol",
		},
	},
	"XEP": {
		{
			ID:   8216,
			Slug: "electra-protocol",
			Name: "Electra Protocol",
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
	"TGAME": {
		{
			ID:   2929,
			Slug: "tgame",
			Name: "Truegame",
		},
	},
	"IMP": {
		{
			ID:   3271,
			Slug: "ether-kingdoms-token",
			Name: "Ether Kingdoms Token",
		},
	},
	"GSPI": {
		{
			ID:   8737,
			Slug: "gspi-governance",
			Name: "GSPI Shopping.io Governance",
		},
	},
	"RVX": {
		{
			ID:   4697,
			Slug: "rivex",
			Name: "Rivex",
		},
	},
	"METRIC": {
		{
			ID:   7254,
			Slug: "metric-exchange",
			Name: "Metric Exchange",
		},
	},
	"KEEP": {
		{
			ID:   5566,
			Slug: "keep-network",
			Name: "Keep Network",
		},
	},
	"ROUTE": {
		{
			ID:   8292,
			Slug: "router-protocol",
			Name: "Router Protocol",
		},
	},
	"YF-DAI": {
		{
			ID:   6938,
			Slug: "yfdai-finance",
			Name: "YFDAI.FINANCE",
		},
	},
	"FEG": {
		{
			ID:   8397,
			Slug: "fegtoken",
			Name: "FEG Token",
		},
	},
	"ARCC": {
		{
			ID:   7843,
			Slug: "asia-reserve-currency-coin",
			Name: "Asia Reserve Currency Coin",
		},
	},
	"XFII": {
		{
			ID:   7648,
			Slug: "xfii",
			Name: "XFII",
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
	"USDF": {
		{
			ID:   6653,
			Slug: "folgoryusd",
			Name: "FolgoryUSD",
		},
	},
	"PCHF": {
		{
			ID:   10706,
			Slug: "peachfolio",
			Name: "peachfolio",
		},
	},
	"XLMDOWN": {
		{
			ID:   8054,
			Slug: "xlmdown",
			Name: "XLMDOWN",
		},
	},
	"GTH": {
		{
			ID:   7041,
			Slug: "gather",
			Name: "Gather",
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
	"WBTC": {
		{
			ID:   3717,
			Slug: "wrapped-bitcoin",
			Name: "Wrapped Bitcoin",
		},
	},
	"RDD": {
		{
			ID:   118,
			Slug: "redd",
			Name: "ReddCoin",
		},
	},
	"PAI": {
		{
			ID:   2900,
			Slug: "project-pai",
			Name: "Project Pai",
		},
	},
	"EMPIRE": {
		{
			ID:   10613,
			Slug: "empire-token",
			Name: "Empire Token",
		},
	},
	"FREL": {
		{
			ID:   9976,
			Slug: "freela",
			Name: "Freela",
		},
	},
	"MOV": {
		{
			ID:   5958,
			Slug: "motiv-protocol",
			Name: "MOTIV Protocol",
		},
	},
	"PTI": {
		{
			ID:   3859,
			Slug: "paytomat",
			Name: "Paytomat",
		},
	},
	"RADDIT": {
		{
			ID:   10074,
			Slug: "radditarium-network",
			Name: "Radditarium Network",
		},
	},
	"CGG": {
		{
			ID:   8648,
			Slug: "chain-guardians",
			Name: "Chain Guardians",
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
	"PRUDE": {
		{
			ID:   10503,
			Slug: "prude-token",
			Name: "Prude Token",
		},
	},
	"BOA": {
		{
			ID:   4217,
			Slug: "bosagora",
			Name: "BOSAGORA",
		},
	},
	"WARP": {
		{
			ID:   8439,
			Slug: "warp-finance",
			Name: "Warp Finance",
		},
	},
	"KDG": {
		{
			ID:   5407,
			Slug: "kingdom-game-4",
			Name: "Kingdom Game 4.0",
		},
	},
	"BLOWF": {
		{
			ID:   8984,
			Slug: "blowfish",
			Name: "BlowFish",
		},
	},
	"ZCRT": {
		{
			ID:   5594,
			Slug: "zcore-token",
			Name: "ZCore Token",
		},
	},
	"FOAM": {
		{
			ID:   3631,
			Slug: "foam",
			Name: "FOAM",
		},
	},
	"WCFG": {
		{
			ID:   10898,
			Slug: "wrapped-centrifuge",
			Name: "Wrapped Centrifuge",
		},
	},
	"EVAI": {
		{
			ID:   9805,
			Slug: "evai-io",
			Name: "Evai.io",
		},
	},
	"ASTRO": {
		{
			ID:   6894,
			Slug: "astrotools",
			Name: "AstroTools",
		},
	},
	"SPRKL": {
		{
			ID:   3780,
			Slug: "sparkle-loyalty",
			Name: "Sparkle Loyalty",
		},
	},
	"SUSHIUP": {
		{
			ID:   8053,
			Slug: "sushiup",
			Name: "SUSHIUP",
		},
	},
	"1GOLD": {
		{
			ID:   4774,
			Slug: "1irstgold",
			Name: "1irstGold",
		},
	},
	"WHOLE": {
		{
			ID:   7633,
			Slug: "wormhole-finance",
			Name: "wormhole.finance",
		},
	},
	"VGX": {
		{
			ID:   1817,
			Slug: "voyager-token",
			Name: "Voyager Token",
		},
	},
	"BEZOGE": {
		{
			ID:   9996,
			Slug: "bezoge-earth",
			Name: "Bezoge Earth",
		},
	},
	"FDOTA": {
		{
			ID:   10846,
			Slug: "fomodota",
			Name: "FomoDota",
		},
	},
	"MBOX": {
		{
			ID:   9175,
			Slug: "mobox",
			Name: "MOBOX",
		},
	},
	"SOCKS": {
		{
			ID:   7095,
			Slug: "unisocks",
			Name: "Unisocks",
		},
	},
	"EXY": {
		{
			ID:   2547,
			Slug: "experty",
			Name: "Experty",
		},
	},
	"XPC": {
		{
			ID:   3750,
			Slug: "experience-chain",
			Name: "eXPerience Chain",
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
	"ZYD": {
		{
			ID:   1389,
			Slug: "zayedcoin",
			Name: "Zayedcoin",
		},
	},
	"TBCC": {
		{
			ID:   8487,
			Slug: "tbcc-labs",
			Name: "TBCC Labs",
		},
	},
	"YCC": {
		{
			ID:   3060,
			Slug: "yuan-chain-coin",
			Name: "Yuan Chain Coin",
		},
	},
	"VIBRA": {
		{
			ID:   10143,
			Slug: "vibraniums",
			Name: "Vibraniums",
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
	"ETHRSI6040": {
		{
			ID:   6143,
			Slug: "eth-rsi-60-40-crossover-set",
			Name: "ETH RSI 60/40 Crossover Set",
		},
	},
	"BWF": {
		{
			ID:   7176,
			Slug: "beowulf",
			Name: "Beowulf",
		},
	},
	"ETHA": {
		{
			ID:   8709,
			Slug: "etha-lend",
			Name: "ETHA Lend",
		},
	},
	"INSN": {
		{
			ID:   1678,
			Slug: "insanecoin-insn",
			Name: "InsaneCoin",
		},
	},
	"BSE": {
		{
			ID:   7996,
			Slug: "buy-sell",
			Name: "Buy-Sell",
		},
	},
	"SHIBSC": {
		{
			ID:   10085,
			Slug: "shiba-bsc",
			Name: "Shiba BSC",
		},
	},
	"KAWA": {
		{
			ID:   10640,
			Slug: "kawakami-inu",
			Name: "Kawakami Inu",
		},
	},
	"ETHBTCRSI": {
		{
			ID:   6139,
			Slug: "eth-btc-rsi-ratio-trading-set",
			Name: "ETH/BTC RSI Ratio Trading Set",
		},
	},
	"KGC": {
		{
			ID:   4291,
			Slug: "krypton-galaxy-coin",
			Name: "Krypton Galaxy Coin",
		},
	},
	"SPWN": {
		{
			ID:   10285,
			Slug: "bitspawn-protocol",
			Name: "Bitspawn Protocol",
		},
	},
	"LVX": {
		{
			ID:   5094,
			Slug: "level01",
			Name: "Level01",
		},
	},
	"TULIP": {
		{
			ID:   10373,
			Slug: "solfarm",
			Name: "SolFarm",
		},
	},
	"POLS": {
		{
			ID:   7208,
			Slug: "polkastarter",
			Name: "Polkastarter",
		},
	},
	"L2": {
		{
			ID:   7772,
			Slug: "leverj-gluon",
			Name: "Leverj Gluon",
		},
	},
	"KOMBAT": {
		{
			ID:   9432,
			Slug: "crypto-kombat",
			Name: "Crypto Kombat",
		},
	},
	"DOGETF": {
		{
			ID:   9729,
			Slug: "doge-father-token",
			Name: "Doge Father Token",
		},
	},
	"FF1": {
		{
			ID:   5457,
			Slug: "two-prime-ff1-token",
			Name: "Two Prime FF1 Token",
		},
	},
	"NNB": {
		{
			ID:   3937,
			Slug: "nnb-token",
			Name: "NNB Token",
		},
	},
	"CRWNY": {
		{
			ID:   9348,
			Slug: "crowny",
			Name: "Crowny",
		},
	},
	"ZYX": {
		{
			ID:   6131,
			Slug: "zyx",
			Name: "ZYX",
		},
	},
	"REEF": {
		{
			ID:   6951,
			Slug: "reef",
			Name: "Reef",
		},
	},
	"BIUT": {
		{
			ID:   4749,
			Slug: "bit-trust-system",
			Name: "Bit Trust System",
		},
	},
	"PEPECASH": {
		{
			ID:   1405,
			Slug: "pepe-cash",
			Name: "Pepe Cash",
		},
	},
	"DERI": {
		{
			ID:   8424,
			Slug: "deri-protocol",
			Name: "Deri Protocol",
		},
	},
	"PEPS": {
		{
			ID:   4397,
			Slug: "peps-coin",
			Name: "PEPS Coin",
		},
	},
	"FLUNAR": {
		{
			ID:   9623,
			Slug: "fairlunar",
			Name: "FairLunar",
		},
	},
	"IBP": {
		{
			ID:   6625,
			Slug: "innovation-blockchain-payment",
			Name: "Innovation Blockchain Payment",
		},
	},
	"FXC": {
		{
			ID:   6773,
			Slug: "futurexcrypto",
			Name: "FUTUREXCRYPTO",
		},
	},
	"MTCN": {
		{
			ID:   4190,
			Slug: "multicoin",
			Name: "Multicoin",
		},
	},
	"QBIT": {
		{
			ID:   10307,
			Slug: "project-quantum",
			Name: "Project Quantum",
		},
	},
	"MOMAT": {
		{
			ID:   9755,
			Slug: "moma-protocol",
			Name: "Moma Protocol",
		},
	},
	"AKT": {
		{
			ID:   7431,
			Slug: "akash-network",
			Name: "Akash Network",
		},
	},
	"GTN": {
		{
			ID:   3914,
			Slug: "glitzkoin",
			Name: "GlitzKoin",
		},
	},
	"VALOR": {
		{
			ID:   3875,
			Slug: "valor-token",
			Name: "Valor Token",
		},
	},
	"MCASH": {
		{
			ID:   4224,
			Slug: "mcashchain",
			Name: "Mcashchain",
		},
	},
	"TN": {
		{
			ID:   4782,
			Slug: "turtlenetwork",
			Name: "TurtleNetwork",
		},
	},
	"BITB": {
		{
			ID:   819,
			Slug: "bean-cash",
			Name: "Bean Cash",
		},
	},
	"CAT": {
		{
			ID:   6649,
			Slug: "cat-token",
			Name: "Cat Token",
		},
	},
	"EDDA": {
		{
			ID:   8789,
			Slug: "eddaswap",
			Name: "EDDASwap",
		},
	},
	"PRYZ": {
		{
			ID:   10107,
			Slug: "pryz",
			Name: "PRYZ",
		},
	},
	"PAN": {
		{
			ID:   5924,
			Slug: "pantos",
			Name: "Pantos",
		},
	},
	"SHOKK": {
		{
			ID:   9730,
			Slug: "shikokuaido",
			Name: "Shikokuaido",
		},
	},
	"HEZ": {
		{
			ID:   7424,
			Slug: "hermez-network",
			Name: "Hermez Network",
		},
	},
	"QUO": {
		{
			ID:   3230,
			Slug: "quoxent",
			Name: "Quoxent",
		},
	},
	"erowan": {
		{
			ID:   8541,
			Slug: "sifchain",
			Name: "SifChain",
		},
	},
	"SWSH": {
		{
			ID:   7191,
			Slug: "swapship",
			Name: "SwapShip",
		},
	},
	"LCT": {
		{
			ID:   7733,
			Slug: "light-coin-exchange-token",
			Name: "Light Coin Exchange Token",
		},
	},
	"AAVE": {
		{
			ID:   7278,
			Slug: "aave",
			Name: "Aave",
		},
	},
	"BRO": {
		{
			ID:   1754,
			Slug: "bitradio",
			Name: "Bitradio",
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
	"CHAIN": {
		{
			ID:   6744,
			Slug: "chain-games",
			Name: "Chain Games",
		},
	},
	"AQUAGOAT": {
		{
			ID:   9370,
			Slug: "aquagoat-finance",
			Name: "AquaGoat.Finance",
		},
	},
	"GIX": {
		{
			ID:   7624,
			Slug: "goldfinx",
			Name: "GoldFinX",
		},
	},
	"AXI": {
		{
			ID:   7141,
			Slug: "axioms",
			Name: "Axioms",
		},
	},
	"COR": {
		{
			ID:   7398,
			Slug: "coreto",
			Name: "Coreto",
		},
	},
	"DHC": {
		{
			ID:   7260,
			Slug: "deltahub-community",
			Name: "DeltaHub Community",
		},
	},
	"IONX": {
		{
			ID:   10264,
			Slug: "charged-particles",
			Name: "Charged Particles",
		},
	},
	"FLY": {
		{
			ID:   9120,
			Slug: "franklin",
			Name: "Franklin",
		},
	},
	"SIAM": {
		{
			ID:   10711,
			Slug: "siamese-neko",
			Name: "Siamese Neko",
		},
	},
	"ALLBI": {
		{
			ID:   5024,
			Slug: "all-best-ico",
			Name: "ALL BEST ICO",
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
	"HOODRAT": {
		{
			ID:   10071,
			Slug: "hoodrat",
			Name: "Hoodrat Finance",
		},
	},
	"VSD": {
		{
			ID:   8775,
			Slug: "value-set-dollar",
			Name: "Value Set Dollar",
		},
	},
	"MAG": {
		{
			ID:   2434,
			Slug: "maggie",
			Name: "Maggie",
		},
	},
	"ZNZ": {
		{
			ID:   4286,
			Slug: "zenzo",
			Name: "ZENZO",
		},
	},
	"INFI": {
		{
			ID:   8305,
			Slug: "insured-finance",
			Name: "Insured Finance",
		},
	},
	"BEET": {
		{
			ID:   3242,
			Slug: "beetle-coin",
			Name: "Beetle Coin",
		},
	},
	"PNIX": {
		{
			ID:   9138,
			Slug: "phoenixdefi-finance",
			Name: "PhoenixDefi.Finance",
		},
	},
	"TAC": {
		{
			ID:   3227,
			Slug: "traceability-chain",
			Name: "Traceability Chain",
		},
	},
	"NEXO": {
		{
			ID:   2694,
			Slug: "nexo",
			Name: "Nexo",
		},
	},
	"FIC": {
		{
			ID:   8383,
			Slug: "filecash",
			Name: "Filecash",
		},
	},
	"LTFM": {
		{
			ID:   10707,
			Slug: "little-fish-moon-token",
			Name: "Little Fish Moon Token",
		},
	},
	"ENERGY": {
		{
			ID:   10199,
			Slug: "energy-token",
			Name: "ENERGY Token",
		},
	},
	"MIMO": {
		{
			ID:   8638,
			Slug: "mimosa",
			Name: "MIMOSA",
		},
	},
	"YFET": {
		{
			ID:   7293,
			Slug: "yfet",
			Name: "YFET",
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
	"TIGER": {
		{
			ID:   8453,
			Slug: "tigerfinance",
			Name: "Tigerfinance",
		},
	},
	"LEVELG": {
		{
			ID:   5578,
			Slug: "levelg",
			Name: "LEVELG",
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
	"NYZO": {
		{
			ID:   5155,
			Slug: "nyzo",
			Name: "Nyzo",
		},
	},
	"SHLD": {
		{
			ID:   8450,
			Slug: "shield-finance",
			Name: "Shield Finance",
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
	"HAKKA": {
		{
			ID:   6622,
			Slug: "hakka-finance",
			Name: "Hakka.Finance",
		},
	},
	"CO2": {
		{
			ID:   8931,
			Slug: "collective",
			Name: "Collective",
		},
	},
	"NEVA": {
		{
			ID:   1200,
			Slug: "nevacoin",
			Name: "NevaCoin",
		},
	},
	"BTCUP": {
		{
			ID:   5608,
			Slug: "btcup",
			Name: "BTCUP",
		},
	},
	"DEXTF": {
		{
			ID:   8691,
			Slug: "dextf-protocol",
			Name: "DEXTF Protocol",
		},
	},
	"ZLP": {
		{
			ID:   7611,
			Slug: "zuplo",
			Name: "Zuplo",
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
	"ASM": {
		{
			ID:   6069,
			Slug: "assemble-protocol",
			Name: "Assemble Protocol",
		},
	},
	"HANA": {
		{
			ID:   5099,
			Slug: "hanacoin",
			Name: "Hanacoin",
		},
	},
	"ALGO": {
		{
			ID:   4030,
			Slug: "algorand",
			Name: "Algorand",
		},
	},
	"ATRI": {
		{
			ID:   7395,
			Slug: "atari-token",
			Name: "Atari Token",
		},
	},
	"MFG": {
		{
			ID:   2658,
			Slug: "smart-mfg",
			Name: "Smart MFG",
		},
	},
	"KNDC": {
		{
			ID:   2890,
			Slug: "kanadecoin",
			Name: "KanadeCoin",
		},
	},
	"NOTSAFEMOON": {
		{
			ID:   9750,
			Slug: "notsafemoon",
			Name: "NotSafeMoon",
		},
	},
	"FINU": {
		{
			ID:   10173,
			Slug: "fire-inu",
			Name: "Fire Inu",
		},
	},
	"SHIELDNET": {
		{
			ID:   9457,
			Slug: "shield-network",
			Name: "Shield Network",
		},
	},
	"LXT": {
		{
			ID:   3070,
			Slug: "litex",
			Name: "Litex",
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
	"TCAT": {
		{
			ID:   3730,
			Slug: "the-currency-analytics",
			Name: "The Currency Analytics",
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
	"TNDR": {
		{
			ID:   8951,
			Slug: "thunderswap",
			Name: "ThunderSwap",
		},
	},
	"ZCH": {
		{
			ID:   5316,
			Slug: "0cash",
			Name: "0cash",
		},
	},
	"MCX": {
		{
			ID:   6521,
			Slug: "machix",
			Name: "Machi X",
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
	"ADP": {
		{
			ID:   8044,
			Slug: "adappter-token",
			Name: "Adappter Token",
		},
	},
	"MOONARCH": {
		{
			ID:   10117,
			Slug: "moonarch-app",
			Name: "Moonarch.app",
		},
	},
	"SX": {
		{
			ID:   8377,
			Slug: "sportx",
			Name: "SportX",
		},
	},
	"NCASH": {
		{
			ID:   2544,
			Slug: "nucleus-vision",
			Name: "Nucleus Vision",
		},
	},
	"SARCO": {
		{
			ID:   10348,
			Slug: "sarcophagus",
			Name: "Sarcophagus",
		},
	},
	"ENV": {
		{
			ID:   9296,
			Slug: "env-finance",
			Name: "ENV Finance",
		},
	},
	"DDR": {
		{
			ID:   5900,
			Slug: "digidinar",
			Name: "DigiDinar",
		},
	},
	"JDB": {
		{
			ID:   7066,
			Slug: "justdobet",
			Name: "Justdobet",
		},
	},
	"CHG": {
		{
			ID:   5400,
			Slug: "charg-coin",
			Name: "Charg Coin",
		},
	},
	"SOGE": {
		{
			ID:   9156,
			Slug: "space-hoge",
			Name: "Space Hoge",
		},
	},
	"EXRN": {
		{
			ID:   2088,
			Slug: "exrnchain",
			Name: "EXRNchain",
		},
	},
	"MSR": {
		{
			ID:   2674,
			Slug: "masari",
			Name: "Masari",
		},
	},
	"BEZ": {
		{
			ID:   2551,
			Slug: "bezop",
			Name: "Bezop",
		},
	},
	"BABYCAKE": {
		{
			ID:   10971,
			Slug: "baby-cake",
			Name: "Baby Cake",
		},
	},
	"LEXI": {
		{
			ID:   4549,
			Slug: "lexit",
			Name: "LEXIT",
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
	"MSC": {
		{
			ID:   8428,
			Slug: "monster-slayer-finance",
			Name: "Monster Slayer Cash",
		},
	},
	"GOZ": {
		{
			ID:   9507,
			Slug: "goztepe-sk-fantoken",
			Name: "Göztepe S.K. Fan Token",
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
	"UMX": {
		{
			ID:   8229,
			Slug: "unimex-network",
			Name: "UniMex Network",
		},
	},
	"DJV": {
		{
			ID:   7834,
			Slug: "dejave",
			Name: "DEJAVE",
		},
	},
	"ZET": {
		{
			ID:   56,
			Slug: "zetacoin",
			Name: "Zetacoin",
		},
	},
	"$KING": {
		{
			ID:   7569,
			Slug: "kingswap",
			Name: "King Swap",
		},
	},
	"GAMESAFE": {
		{
			ID:   9909,
			Slug: "gamesafe-io",
			Name: "Gamesafe.io",
		},
	},
	"WAVES": {
		{
			ID:   1274,
			Slug: "waves",
			Name: "Waves",
		},
	},
	"RVP": {
		{
			ID:   9277,
			Slug: "revolution-populi",
			Name: "Revolution Populi",
		},
	},
	"K21": {
		{
			ID:   9205,
			Slug: "k21",
			Name: "K21",
		},
	},
	"CBIX-P": {
		{
			ID:   9512,
			Slug: "cubiex-power",
			Name: "Cubiex Power",
		},
	},
	"JNTR": {
		{
			ID:   7545,
			Slug: "jointer",
			Name: "Jointer",
		},
	},
	"COMFI": {
		{
			ID:   9609,
			Slug: "complifi",
			Name: "CompliFi",
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
	"GAINZ": {
		{
			ID:   10587,
			Slug: "gainz-token",
			Name: "GAINZ TOKEN",
		},
	},
	"OGN": {
		{
			ID:   5117,
			Slug: "origin-protocol",
			Name: "Origin Protocol",
		},
	},
	"MDT": {
		{
			ID:   2348,
			Slug: "measurable-data-token",
			Name: "Measurable Data Token",
		},
	},
	"COLX": {
		{
			ID:   2001,
			Slug: "colossusxt",
			Name: "ColossusXT",
		},
	},
	"MAC": {
		{
			ID:   3895,
			Slug: "matrexcoin",
			Name: "Matrexcoin",
		},
	},
	"STAC": {
		{
			ID:   2565,
			Slug: "startercoin",
			Name: "StarterCoin",
		},
	},
	"KODX": {
		{
			ID:   7828,
			Slug: "king-of-defi",
			Name: "KING OF DEFI",
		},
	},
	"MAYFI": {
		{
			ID:   8921,
			Slug: "matic-aave-yfi",
			Name: "Matic Aave Interest Bearing YFI",
		},
	},
	"FRMS": {
		{
			ID:   6916,
			Slug: "the-forms",
			Name: "The Forms",
		},
	},
	"KILI": {
		{
			ID:   9320,
			Slug: "kilimanjaro",
			Name: "KILIMANJARO",
		},
	},
	"MOD": {
		{
			ID:   8494,
			Slug: "modefi",
			Name: "Modefi",
		},
	},
	"AAC": {
		{
			ID:   2438,
			Slug: "acute-angle-cloud",
			Name: "Acute Angle Cloud",
		},
	},
	"CCO": {
		{
			ID:   2241,
			Slug: "ccore",
			Name: "Ccore",
		},
	},
	"HTBULL": {
		{
			ID:   6084,
			Slug: "3x-long-huobi-token-token",
			Name: "3X Long Huobi Token Token",
		},
	},
	"NULS": {
		{
			ID:   2092,
			Slug: "nuls",
			Name: "NULS",
		},
	},
	"EURU": {
		{
			ID:   6905,
			Slug: "upper-euro",
			Name: "Upper Euro",
		},
	},
	"IBVOL": {
		{
			ID:   5534,
			Slug: "inverse-bitcoin-volatility-token",
			Name: "Inverse Bitcoin Volatility Token",
		},
	},
	"XDAG": {
		{
			ID:   4424,
			Slug: "xdag",
			Name: "XDAG",
		},
	},
	"XSPC": {
		{
			ID:   4013,
			Slug: "spectre-security-coin",
			Name: "SpectreSecurityCoin",
		},
	},
	"POLC": {
		{
			ID:   8549,
			Slug: "polkacity",
			Name: "Polkacity",
		},
	},
	"CFG": {
		{
			ID:   6748,
			Slug: "centrifuge",
			Name: "Centrifuge",
		},
	},
	"PNIXS": {
		{
			ID:   9641,
			Slug: "phoenxidefi-finance",
			Name: "PhoenxiDefi Finance",
		},
	},
	"AIRT": {
		{
			ID:   10905,
			Slug: "airnfts",
			Name: "AirNFTs",
		},
	},
	"WLT": {
		{
			ID:   8721,
			Slug: "wealthlocks",
			Name: "Wealthlocks",
		},
	},
	"SURF": {
		{
			ID:   7547,
			Slug: "surf",
			Name: "SURF Finance",
		},
	},
	"TRAT": {
		{
			ID:   3925,
			Slug: "tratok",
			Name: "Tratok",
		},
	},
	"SIN": {
		{
			ID:   3514,
			Slug: "sinovate",
			Name: "SINOVATE",
		},
	},
	"BNC": {
		{
			ID:   3222,
			Slug: "bionic",
			Name: "Bionic",
		},
	},
	"IBANK": {
		{
			ID:   1515,
			Slug: "ibank",
			Name: "iBank",
		},
	},
	"NKC": {
		{
			ID:   2477,
			Slug: "nework",
			Name: "Nework",
		},
	},
	"EARN": {
		{
			ID:   7586,
			Slug: "yearn-classic-finance",
			Name: "Yearn Classic Finance",
		},
	},
	"TEM": {
		{
			ID:   5945,
			Slug: "temtum",
			Name: "Temtum",
		},
	},
	"NLC2": {
		{
			ID:   1382,
			Slug: "nolimitcoin",
			Name: "NoLimitCoin",
		},
	},
	"GOVI": {
		{
			ID:   8408,
			Slug: "govi",
			Name: "Govi",
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
	"NASH": {
		{
			ID:   3867,
			Slug: "neoworld-cash",
			Name: "NeoWorld Cash",
		},
	},
	"ARO": {
		{
			ID:   3024,
			Slug: "arionum",
			Name: "Arionum",
		},
	},
	"NII": {
		{
			ID:   4865,
			Slug: "nahmii",
			Name: "Nahmii",
		},
	},
	"SI": {
		{
			ID:   8701,
			Slug: "siren",
			Name: "Siren",
		},
	},
	"CNRG": {
		{
			ID:   5376,
			Slug: "cryptoenergy",
			Name: "CryptoEnergy",
		},
	},
	"GCN": {
		{
			ID:   730,
			Slug: "gcn-coin",
			Name: "GCN Coin",
		},
	},
	"REI": {
		{
			ID:   10844,
			Slug: "zerogoki",
			Name: "Zerogoki",
		},
	},
	"VOLLAR": {
		{
			ID:   3891,
			Slug: "v-dimension",
			Name: "V-Dimension",
		},
	},
	"KP2R": {
		{
			ID:   7751,
			Slug: "kp2r-network",
			Name: "KP2R.Network",
		},
	},
	"DAILYS": {
		{
			ID:   8840,
			Slug: "dailyswap-token",
			Name: "DailySwap Token",
		},
	},
	"HTDF": {
		{
			ID:   5127,
			Slug: "orient-walt",
			Name: "Orient Walt",
		},
	},
	"ARGON": {
		{
			ID:   8421,
			Slug: "argon",
			Name: "Argon",
		},
	},
	"PTON": {
		{
			ID:   3813,
			Slug: "pton",
			Name: "PTON",
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
	"BULLSHIT": {
		{
			ID:   6088,
			Slug: "3x-long-shitcoin-index-token",
			Name: "3X Long Shitcoin Index Token",
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
	"OK": {
		{
			ID:   760,
			Slug: "okcash",
			Name: "OKCash",
		},
	},
	"WGOLD": {
		{
			ID:   9332,
			Slug: "apwars",
			Name: "APWars",
		},
	},
	"FSHIB": {
		{
			ID:   10752,
			Slug: "floki-shiba",
			Name: "Floki Shiba",
		},
	},
	"SALE": {
		{
			ID:   6742,
			Slug: "dxsale-network",
			Name: "DxSale Network",
		},
	},
	"CHM": {
		{
			ID:   7589,
			Slug: "cryptochrome",
			Name: "Cryptochrome",
		},
	},
	"RNX": {
		{
			ID:   6531,
			Slug: "roonex",
			Name: "ROONEX",
		},
	},
	"CORD": {
		{
			ID:   8535,
			Slug: "cord-finance",
			Name: "CORD.Finance",
		},
	},
	"ZKS": {
		{
			ID:   8202,
			Slug: "zkswap",
			Name: "ZKSwap",
		},
	},
	"HEX": {
		{
			ID:   5015,
			Slug: "hex",
			Name: "HEX",
		},
	},
	"AE": {
		{
			ID:   1700,
			Slug: "aeternity",
			Name: "Aeternity",
		},
	},
	"BUZZ": {
		{
			ID:   1962,
			Slug: "buzzcoin",
			Name: "BUZZCoin",
		},
	},
	"HLX": {
		{
			ID:   5057,
			Slug: "helex",
			Name: "Helex",
		},
	},
	"BDOG": {
		{
			ID:   9280,
			Slug: "bulldog-token",
			Name: "Bulldog Token",
		},
	},
	"vBUSD": {
		{
			ID:   7959,
			Slug: "venus-busd",
			Name: "Venus BUSD",
		},
	},
	"GHST": {
		{
			ID:   7046,
			Slug: "aavegotchi",
			Name: "Aavegotchi",
		},
	},
	"XCO": {
		{
			ID:   837,
			Slug: "x-coin",
			Name: "X-Coin",
		},
	},
	"BULLS": {
		{
			ID:   8360,
			Slug: "bulls",
			Name: "BULLS",
		},
	},
	"DOGIRA": {
		{
			ID:   9298,
			Slug: "dogira",
			Name: "Dogira",
		},
	},
	"UNIC": {
		{
			ID:   9998,
			Slug: "unicly",
			Name: "Unicly",
		},
	},
	"SAFEBANK": {
		{
			ID:   10015,
			Slug: "safebank-yes",
			Name: "SafeBank YES",
		},
	},
	"GO": {
		{
			ID:   2861,
			Slug: "gochain",
			Name: "GoChain",
		},
	},
	"METAMOON": {
		{
			ID:   9688,
			Slug: "metamoon",
			Name: "MetaMoon",
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
	"NIF": {
		{
			ID:   7879,
			Slug: "unifty",
			Name: "Unifty",
		},
	},
	"SPIKE": {
		{
			ID:   4094,
			Slug: "spiking",
			Name: "Spiking",
		},
	},
	"KSEED": {
		{
			ID:   7135,
			Slug: "kush-finance",
			Name: "Kush Finance",
		},
	},
	"DOGEBULL": {
		{
			ID:   6081,
			Slug: "3x-long-dogecoin-token",
			Name: "3X Long Dogecoin Token",
		},
	},
	"YFIM": {
		{
			ID:   7613,
			Slug: "yfi-mobi",
			Name: "Yfi.mobi",
		},
	},
	"IBNB": {
		{
			ID:   10878,
			Slug: "i-bnb",
			Name: "iBNB",
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
	"VVT": {
		{
			ID:   7907,
			Slug: "versoview",
			Name: "VersoView",
		},
	},
	"TOS": {
		{
			ID:   2987,
			Slug: "thingsoperatingsystem",
			Name: "ThingsOperatingSystem",
		},
	},
	"LKK": {
		{
			ID:   1454,
			Slug: "lykke",
			Name: "Lykke",
		},
	},
	"FCD": {
		{
			ID:   7640,
			Slug: "future-cash-digital",
			Name: "Future-Cash Digital",
		},
	},
	"WHITE": {
		{
			ID:   8120,
			Slug: "whiteheart",
			Name: "Whiteheart",
		},
	},
	"ZNY": {
		{
			ID:   990,
			Slug: "bitzeny",
			Name: "Bitzeny",
		},
	},
	"CPI": {
		{
			ID:   6654,
			Slug: "crypto-price-index",
			Name: "Crypto Price Index",
		},
	},
	"DMX": {
		{
			ID:   8117,
			Slug: "dymmax",
			Name: "Dymmax",
		},
	},
	"UBER": {
		{
			ID:   7917,
			Slug: "uber-tokenized-stock-ftx",
			Name: "Uber tokenized stock FTX",
		},
	},
	"KUE": {
		{
			ID:   3369,
			Slug: "kuende",
			Name: "Kuende",
		},
	},
	"IPL": {
		{
			ID:   2421,
			Slug: "insurepal",
			Name: "VouchForMe",
		},
	},
	"GBK": {
		{
			ID:   7581,
			Slug: "goldblock",
			Name: "Goldblock",
		},
	},
	"BTRS": {
		{
			ID:   4257,
			Slug: "bitball-treasure",
			Name: "Bitball Treasure",
		},
	},
	"VLS": {
		{
			ID:   4136,
			Slug: "veles",
			Name: "Veles",
		},
	},
	"GENS": {
		{
			ID:   10278,
			Slug: "genshiro",
			Name: "Genshiro",
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
	"ATOMBULL": {
		{
			ID:   6080,
			Slug: "3x-long-cosmos-token",
			Name: "3X Long Cosmos Token",
		},
	},
	"YFICG": {
		{
			ID:   7517,
			Slug: "yfi-credits-group",
			Name: "YFI CREDITS GROUP",
		},
	},
	"STPT": {
		{
			ID:   4006,
			Slug: "standard-tokenization-protocol",
			Name: "Standard Tokenization Protocol",
		},
	},
	"DCNTR": {
		{
			ID:   6609,
			Slug: "decentrahub-coin",
			Name: "Decentrahub Coin",
		},
	},
	"onLEXpa": {
		{
			ID:   5097,
			Slug: "onlexpa",
			Name: "onLEXpa",
		},
	},
	"PUP": {
		{
			ID:   10242,
			Slug: "pupper",
			Name: "Pupper",
		},
	},
	"EDI": {
		{
			ID:   5165,
			Slug: "freight-trust-clearing-network",
			Name: "Freight Trust & Clearing Network",
		},
	},
	"DFX": {
		{
			ID:   6919,
			Slug: "definitex",
			Name: "Definitex",
		},
	},
	"CQT": {
		{
			ID:   7411,
			Slug: "covalent",
			Name: "Covalent",
		},
	},
	"ETP": {
		{
			ID:   1703,
			Slug: "metaverse",
			Name: "Metaverse ETP",
		},
	},
	"OKT": {
		{
			ID:   8267,
			Slug: "okexchain",
			Name: "OKExChain",
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
	"ZAP": {
		{
			ID:   2363,
			Slug: "zap",
			Name: "Zap",
		},
	},
	"CPT": {
		{
			ID:   2766,
			Slug: "cryptaur",
			Name: "Cryptaur",
		},
	},
	"STKR": {
		{
			ID:   8934,
			Slug: "stakerdao",
			Name: "StakerDAO",
		},
	},
	"LYRA": {
		{
			ID:   4542,
			Slug: "scrypta",
			Name: "Scrypta",
		},
	},
	"WTN": {
		{
			ID:   3484,
			Slug: "waletoken",
			Name: "Waletoken",
		},
	},
	"MNP": {
		{
			ID:   3348,
			Slug: "mnpcoin",
			Name: "MNPCoin",
		},
	},
	"LANC": {
		{
			ID:   10666,
			Slug: "lanceria",
			Name: "Lanceria",
		},
	},
	"YCE": {
		{
			ID:   4381,
			Slug: "myce",
			Name: "MYCE",
		},
	},
	"AfroX": {
		{
			ID:   5135,
			Slug: "afrodex",
			Name: "AfroDex",
		},
	},
	"UTL": {
		{
			ID:   8728,
			Slug: "utile-network",
			Name: "Utile Network",
		},
	},
	"DEGO": {
		{
			ID:   7087,
			Slug: "dego-finance",
			Name: "Dego Finance",
		},
	},
	"LND": {
		{
			ID:   2686,
			Slug: "lendingblock",
			Name: "Lendingblock",
		},
	},
	"CTASK": {
		{
			ID:   8418,
			Slug: "cryptotask",
			Name: "CryptoTask",
		},
	},
	"NAT": {
		{
			ID:   3889,
			Slug: "natmin-pure-escrow",
			Name: "Natmin Pure Escrow",
		},
	},
	"CMK": {
		{
			ID:   10485,
			Slug: "credmark",
			Name: "Credmark",
		},
	},
	"PTA": {
		{
			ID:   8382,
			Slug: "petrachor",
			Name: "Petrachor",
		},
	},
	"PICKLE": {
		{
			ID:   7022,
			Slug: "pickle-finance",
			Name: "Pickle Finance",
		},
	},
	"CRAD": {
		{
			ID:   4255,
			Slug: "cryptoads-marketplace",
			Name: "CryptoAds Marketplace",
		},
	},
	"KWATT": {
		{
			ID:   3209,
			Slug: "4new",
			Name: "4NEW",
		},
	},
	"YUMMY": {
		{
			ID:   9859,
			Slug: "yummy",
			Name: "YUMMY",
		},
	},
	"808TA": {
		{
			ID:   5108,
			Slug: "808ta",
			Name: "808TA",
		},
	},
	"DPT": {
		{
			ID:   3920,
			Slug: "diamond-platform-token",
			Name: "Diamond Platform Token",
		},
	},
	"BSOV": {
		{
			ID:   4306,
			Slug: "bitcoinsov",
			Name: "BitcoinSoV",
		},
	},
	"EVY": {
		{
			ID:   3754,
			Slug: "everycoin",
			Name: "EveryCoin ",
		},
	},
	"BBKFI": {
		{
			ID:   9581,
			Slug: "bitblocks-finance",
			Name: "BitBlocks Finance",
		},
	},
	"XAP": {
		{
			ID:   3159,
			Slug: "apollon",
			Name: "Apollon",
		},
	},
	"SANSHU": {
		{
			ID:   9711,
			Slug: "sanshu-inu",
			Name: "Sanshu Inu",
		},
	},
	"ZDEX": {
		{
			ID:   6989,
			Slug: "zeedex",
			Name: "Zeedex",
		},
	},
	"GIV": {
		{
			ID:   6075,
			Slug: "givly-coin",
			Name: "GIVLY Coin",
		},
	},
	"HOGL": {
		{
			ID:   8944,
			Slug: "hogl-finance",
			Name: "HOGL finance",
		},
	},
	"WHX": {
		{
			ID:   8943,
			Slug: "whitex",
			Name: "WHITEX",
		},
	},
	"D100": {
		{
			ID:   8551,
			Slug: "defi100",
			Name: "DeFi100",
		},
	},
	"BSPAY": {
		{
			ID:   8956,
			Slug: "brosispay",
			Name: "Brosispay",
		},
	},
	"TARM": {
		{
			ID:   5564,
			Slug: "armtoken",
			Name: "ARMTOKEN",
		},
	},
	"LAS": {
		{
			ID:   8960,
			Slug: "lnasolution-coin",
			Name: "LNAsolution Coin",
		},
	},
	"HAND": {
		{
			ID:   3181,
			Slug: "showhand",
			Name: "ShowHand",
		},
	},
	"JS": {
		{
			ID:   2100,
			Slug: "javascript-token",
			Name: "JavaScript Token",
		},
	},
	"OHM": {
		{
			ID:   9067,
			Slug: "olympus",
			Name: "Olympus",
		},
	},
	"DEA": {
		{
			ID:   8248,
			Slug: "deus-finance-dea",
			Name: "DEUS Finance DEA",
		},
	},
	"LA": {
		{
			ID:   2090,
			Slug: "latoken",
			Name: "LATOKEN",
		},
	},
	"FLETA": {
		{
			ID:   4103,
			Slug: "fleta",
			Name: "FLETA",
		},
	},
	"TRTT": {
		{
			ID:   2865,
			Slug: "trittium",
			Name: "Trittium",
		},
	},
	"EPIC": {
		{
			ID:   5435,
			Slug: "epic-cash",
			Name: "Epic Cash",
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
	"ENT": {
		{
			ID:   1474,
			Slug: "eternity",
			Name: "Eternity",
		},
	},
	"IOP": {
		{
			ID:   1464,
			Slug: "internet-of-people",
			Name: "Internet of People",
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
	"VOLT": {
		{
			ID:   1657,
			Slug: "bitvolt",
			Name: "Bitvolt",
		},
	},
	"REDFEG": {
		{
			ID:   11033,
			Slug: "redfeg",
			Name: "RedFEG",
		},
	},
	"CHAINCADE": {
		{
			ID:   10875,
			Slug: "chaincade",
			Name: "ChainCade",
		},
	},
	"YNK": {
		{
			ID:   7252,
			Slug: "yoink",
			Name: "Yoink",
		},
	},
	"$NOOB": {
		{
			ID:   7646,
			Slug: "noob-finance",
			Name: "noob.finance",
		},
	},
	"DMUSK": {
		{
			ID:   10018,
			Slug: "dragonmusk",
			Name: "Dragonmusk",
		},
	},
	"BNKR": {
		{
			ID:   4918,
			Slug: "bankroll-network",
			Name: "Bankroll Network",
		},
	},
	"ENJ": {
		{
			ID:   2130,
			Slug: "enjin-coin",
			Name: "Enjin Coin",
		},
	},
	"POA": {
		{
			ID:   2548,
			Slug: "poa",
			Name: "POA",
		},
	},
	"ICOB": {
		{
			ID:   1514,
			Slug: "icobid",
			Name: "ICOBID",
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
	"CRW": {
		{
			ID:   720,
			Slug: "crown",
			Name: "Crown",
		},
	},
	"UGOTCHI": {
		{
			ID:   9472,
			Slug: "unicly-aavegotchi-astronauts-collection",
			Name: "Unicly Aavegotchi Astronauts Collection",
		},
	},
	"ESH": {
		{
			ID:   4096,
			Slug: "switch",
			Name: "Switch",
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
	"KAIINU": {
		{
			ID:   10330,
			Slug: "kai-inu",
			Name: "KAI INU",
		},
	},
	"YFIB": {
		{
			ID:   6875,
			Slug: "yfibusiness",
			Name: "YFIBusiness",
		},
	},
	"BTCB": {
		{
			ID:   4023,
			Slug: "bitcoin-bep2",
			Name: "Bitcoin BEP2",
		},
	},
	"PLR": {
		{
			ID:   1834,
			Slug: "pillar",
			Name: "Pillar",
		},
	},
	"MTLX": {
		{
			ID:   7256,
			Slug: "mettalex",
			Name: "Mettalex",
		},
	},
	"DIME": {
		{
			ID:   90,
			Slug: "dimecoin",
			Name: "Dimecoin",
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
	"GOSS": {
		{
			ID:   3332,
			Slug: "gossipcoin",
			Name: "Gossip Coin",
		},
	},
	"HOPR": {
		{
			ID:   6520,
			Slug: "hopr",
			Name: "HOPR",
		},
	},
	"LKR": {
		{
			ID:   8970,
			Slug: "polkalokr",
			Name: "Polkalokr",
		},
	},
	"HSN": {
		{
			ID:   4682,
			Slug: "hyper-speed-network",
			Name: "Hyper Speed Network",
		},
	},
	"PICA": {
		{
			ID:   7627,
			Slug: "picaartmoney",
			Name: "PicaArtMoney",
		},
	},
	"WISHDOGE": {
		{
			ID:   10966,
			Slug: "wish-doge-dragon",
			Name: "Wish Doge Dragon",
		},
	},
	"PEAK": {
		{
			ID:   5354,
			Slug: "peakdefi",
			Name: "PEAKDEFI",
		},
	},
	"IOC": {
		{
			ID:   495,
			Slug: "iocoin",
			Name: "I/O Coin",
		},
	},
	"BUIDL": {
		{
			ID:   5704,
			Slug: "dfohub",
			Name: "DFOhub",
		},
	},
	"MINIBABYDOGE": {
		{
			ID:   10806,
			Slug: "mini-baby-doge",
			Name: "Mini Baby Doge",
		},
	},
	"LOA": {
		{
			ID:   5516,
			Slug: "loa-protocol",
			Name: "LOA Protocol",
		},
	},
	"SORA": {
		{
			ID:   5721,
			Slug: "sorachancoin",
			Name: "SorachanCoin",
		},
	},
	"SHILLING": {
		{
			ID:   10209,
			Slug: "shilling-token",
			Name: "Shilling Token",
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
	"ZPT": {
		{
			ID:   2481,
			Slug: "zeepin",
			Name: "Zeepin",
		},
	},
	"WCELO": {
		{
			ID:   8288,
			Slug: "wrapped-celo",
			Name: "Wrapped Celo",
		},
	},
	"CHZ": {
		{
			ID:   4066,
			Slug: "chiliz",
			Name: "Chiliz",
		},
	},
	"IGNIS": {
		{
			ID:   2276,
			Slug: "ignis",
			Name: "Ignis",
		},
	},
	"GEEQ": {
		{
			ID:   6194,
			Slug: "geeq",
			Name: "Geeq",
		},
	},
	"SMC": {
		{
			ID:   113,
			Slug: "smartcoin",
			Name: "SmartCoin",
		},
	},
	"A": {
		{
			ID:   3869,
			Slug: "alpha-token",
			Name: "Alpha Token",
		},
	},
	"MXW": {
		{
			ID:   6369,
			Slug: "maxonrow",
			Name: "Maxonrow",
		},
	},
	"R34P": {
		{
			ID:   8040,
			Slug: "r34p",
			Name: "R34P",
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
	"ORION": {
		{
			ID:   10395,
			Slug: "orion",
			Name: "Orion",
		},
	},
	"DFST": {
		{
			ID:   8446,
			Slug: "defistarter",
			Name: "DeFiStarter",
		},
	},
	"MCM": {
		{
			ID:   5125,
			Slug: "mochimo",
			Name: "Mochimo",
		},
	},
	"EUC": {
		{
			ID:   1038,
			Slug: "eurocoin",
			Name: "Eurocoin",
		},
	},
	"ZTNZ": {
		{
			ID:   8165,
			Slug: "ztranzit-coin",
			Name: "Ztranzit Coin",
		},
	},
	"PHTR": {
		{
			ID:   9423,
			Slug: "phuture",
			Name: "Phuture",
		},
	},
	"MTTCOIN": {
		{
			ID:   7835,
			Slug: "money-of-tomorrow-today",
			Name: "Money of Tomorrow, Today",
		},
	},
	"ROVER": {
		{
			ID:   9702,
			Slug: "rover-token",
			Name: "Rover Inu Token",
		},
	},
	"SID": {
		{
			ID:   10559,
			Slug: "shield-finance-protocol",
			Name: "Shield Token",
		},
	},
	"REQ": {
		{
			ID:   2071,
			Slug: "request",
			Name: "Request",
		},
	},
	"CUDOS": {
		{
			ID:   8258,
			Slug: "cudos",
			Name: "CUDOS",
		},
	},
	"TWERK": {
		{
			ID:   9893,
			Slug: "twerk-finance",
			Name: "Twerk Finance",
		},
	},
	"TCUB": {
		{
			ID:   10862,
			Slug: "tiger-cub",
			Name: "Tiger Cub",
		},
	},
	"XEQ": {
		{
			ID:   4211,
			Slug: "equilibria",
			Name: "Equilibria",
		},
	},
	"DATA": {
		{
			ID:   2143,
			Slug: "streamr",
			Name: "Streamr",
		},
	},
	"MOMA": {
		{
			ID:   9440,
			Slug: "mochi-market",
			Name: "Mochi Market",
		},
	},
	"ZXC": {
		{
			ID:   2920,
			Slug: "0xcert",
			Name: "0xcert",
		},
	},
	"CHART": {
		{
			ID:   6861,
			Slug: "chartex",
			Name: "ChartEx",
		},
	},
	"STRI": {
		{
			ID:   9395,
			Slug: "strite",
			Name: "Strite",
		},
	},
	"PONZU": {
		{
			ID:   10859,
			Slug: "ponzu-inu",
			Name: "Ponzu Inu",
		},
	},
	"GASP": {
		{
			ID:   7625,
			Slug: "gasp",
			Name: "gAsp",
		},
	},
	"BHAO": {
		{
			ID:   7503,
			Slug: "bithao",
			Name: "Bithao",
		},
	},
	"OXB": {
		{
			ID:   8649,
			Slug: "oxbull-tech",
			Name: "Oxbull.tech",
		},
	},
	"DOKI": {
		{
			ID:   7376,
			Slug: "doki-doki-finance",
			Name: "Doki Doki Finance",
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
	"SAFEHAMSTERS": {
		{
			ID:   10060,
			Slug: "safehamsters",
			Name: "SafeHamsters",
		},
	},
	"FXP": {
		{
			ID:   6477,
			Slug: "fxpay",
			Name: "FXPay",
		},
	},
	"ALINK": {
		{
			ID:   5751,
			Slug: "aave-link",
			Name: "Aave LINK",
		},
	},
	"LBTC": {
		{
			ID:   2335,
			Slug: "lightning-bitcoin",
			Name: "Lightning Bitcoin",
		},
	},
	"LUX": {
		{
			ID:   2107,
			Slug: "luxcoin",
			Name: "LUXCoin",
		},
	},
	"DCTD": {
		{
			ID:   9066,
			Slug: "dctdao",
			Name: "DCTDAO",
		},
	},
	"CLEANOCEAN": {
		{
			ID:   10499,
			Slug: "cleanocean",
			Name: "CleanOcean",
		},
	},
	"THEMOON": {
		{
			ID:   10773,
			Slug: "win-space-ticket",
			Name: "Win Space Ticket",
		},
	},
	"BART": {
		{
			ID:   6592,
			Slug: "bartertrade",
			Name: "BarterTrade",
		},
	},
	"X8X": {
		{
			ID:   3079,
			Slug: "x8x-token",
			Name: "X8X Token",
		},
	},
	"AIX": {
		{
			ID:   2367,
			Slug: "aigang",
			Name: "Aigang",
		},
	},
	"HLT": {
		{
			ID:   3838,
			Slug: "esportbits",
			Name: "Esportbits",
		},
	},
	"ALBT": {
		{
			ID:   6957,
			Slug: "allianceblock",
			Name: "AllianceBlock",
		},
	},
	"GNY": {
		{
			ID:   3936,
			Slug: "gny",
			Name: "GNY",
		},
	},
	"AIOZ": {
		{
			ID:   9104,
			Slug: "aioz-network",
			Name: "AIOZ Network",
		},
	},
	"UPI": {
		{
			ID:   5086,
			Slug: "pawtocol",
			Name: "Pawtocol",
		},
	},
	"RUGZ": {
		{
			ID:   7649,
			Slug: "rugz",
			Name: "pulltherug.finance",
		},
	},
	"GAZE": {
		{
			ID:   9433,
			Slug: "gazetv",
			Name: "GazeTV",
		},
	},
	"DSL": {
		{
			ID:   8739,
			Slug: "deadsoul",
			Name: "DeadSoul",
		},
	},
	"MED": {
		{
			ID:   2303,
			Slug: "medibloc",
			Name: "MediBloc",
		},
	},
	"PUNDIX": {
		{
			ID:   9040,
			Slug: "pundix-new",
			Name: "Pundi X[new]",
		},
	},
	"TOC": {
		{
			ID:   3965,
			Slug: "touchcon",
			Name: "TouchCon",
		},
	},
	"DVT": {
		{
			ID:   4027,
			Slug: "devault",
			Name: "DeVault",
		},
	},
	"STEEM": {
		{
			ID:   1230,
			Slug: "steem",
			Name: "Steem",
		},
	},
	"mMSFT": {
		{
			ID:   8017,
			Slug: "mirrored-microsoft",
			Name: "Mirrored Microsoft",
		},
	},
	"SHAKE": {
		{
			ID:   7390,
			Slug: "shake",
			Name: "Spaceswap SHAKE",
		},
	},
	"EPAN": {
		{
			ID:   7749,
			Slug: "paypolitan-token",
			Name: "Paypolitan Token",
		},
	},
	"TRXDOWN": {
		{
			ID:   7004,
			Slug: "trxdown",
			Name: "TRXDOWN",
		},
	},
	"MDOGE": {
		{
			ID:   10911,
			Slug: "missdoge",
			Name: "MissDoge",
		},
	},
	"SAFELIGHT": {
		{
			ID:   9208,
			Slug: "safelight",
			Name: "SafeLight",
		},
	},
	"ZANO": {
		{
			ID:   4691,
			Slug: "zano",
			Name: "Zano",
		},
	},
	"BHP": {
		{
			ID:   3020,
			Slug: "bhp-coin",
			Name: "BHPCoin",
		},
	},
	"BDK": {
		{
			ID:   5288,
			Slug: "bidesk",
			Name: "Bidesk",
		},
	},
	"ZPAE": {
		{
			ID:   5663,
			Slug: "zelaapayae",
			Name: "ZelaaPayAE",
		},
	},
	"UNIFY": {
		{
			ID:   1736,
			Slug: "unify",
			Name: "Unify",
		},
	},
	"CFXQ": {
		{
			ID:   9070,
			Slug: "cfx-quantum",
			Name: "CFX Quantum",
		},
	},
	"FECLIPSE": {
		{
			ID:   9185,
			Slug: "faireclipse",
			Name: "FairEclipse",
		},
	},
	"RFCTR": {
		{
			ID:   8039,
			Slug: "reflector-finance",
			Name: "Reflector.Finance",
		},
	},
	"USDA": {
		{
			ID:   5058,
			Slug: "usda",
			Name: "USDA",
		},
	},
	"MIC": {
		{
			ID:   8137,
			Slug: "mith-cash",
			Name: "MITH Cash",
		},
	},
	"TKS": {
		{
			ID:   1588,
			Slug: "tokes",
			Name: "Tokes",
		},
	},
	"TDP": {
		{
			ID:   3469,
			Slug: "truedeck",
			Name: "TrueDeck",
		},
	},
	"MEM": {
		{
			ID:   10461,
			Slug: "meme-coin",
			Name: "Memecoin",
		},
	},
	"SENSI": {
		{
			ID:   10250,
			Slug: "sensible-finance",
			Name: "Sensible.Finance",
		},
	},
	"BSVBEAR": {
		{
			ID:   5459,
			Slug: "3x-short-bitcoin-sv-token",
			Name: "3x Short Bitcoin SV Token",
		},
	},
	"FUN": {
		{
			ID:   1757,
			Slug: "funtoken",
			Name: "FUNToken",
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
	"TCORE": {
		{
			ID:   8082,
			Slug: "tornado",
			Name: "Tornado",
		},
	},
	"BSF": {
		{
			ID:   10856,
			Slug: "babyspacefloki",
			Name: "BabySpaceFloki",
		},
	},
	"LOF": {
		{
			ID:   10012,
			Slug: "lonelyfans",
			Name: "Lonelyfans",
		},
	},
	"HEJJ": {
		{
			ID:   9992,
			Slug: "hedge4-ai",
			Name: "HEDGE4.Ai",
		},
	},
	"CLUB": {
		{
			ID:   1135,
			Slug: "clubcoin",
			Name: "ClubCoin",
		},
	},
	"LINKETHRSI": {
		{
			ID:   6158,
			Slug: "link-eth-rsi-ratio-trading-set",
			Name: "LINK/ETH RSI Ratio Trading Set",
		},
	},
	"SPH": {
		{
			ID:   6209,
			Slug: "spheroid-universe",
			Name: "Spheroid Universe",
		},
	},
	"EDOGE": {
		{
			ID:   9932,
			Slug: "elondoge",
			Name: "ElonDoge",
		},
	},
	"COMFY": {
		{
			ID:   9627,
			Slug: "comfytoken",
			Name: "ComfyToken",
		},
	},
	"PARA": {
		{
			ID:   8298,
			Slug: "paralink-network",
			Name: "Paralink Network",
		},
	},
	"SMRAT": {
		{
			ID:   9284,
			Slug: "secured-moonrat-token",
			Name: "Secured MoonRat Token",
		},
	},
	"MYO": {
		{
			ID:   9577,
			Slug: "mycro",
			Name: "Mycro",
		},
	},
	"BKKG": {
		{
			ID:   7753,
			Slug: "biokkoin",
			Name: "BIOKKOIN",
		},
	},
	"SWRV": {
		{
			ID:   6901,
			Slug: "swerve",
			Name: "Swerve",
		},
	},
	"BRG": {
		{
			ID:   7096,
			Slug: "bridge-oracle",
			Name: "Bridge Oracle",
		},
	},
	"IXC": {
		{
			ID:   13,
			Slug: "ixcoin",
			Name: "Ixcoin",
		},
	},
	"USDS": {
		{
			ID:   3719,
			Slug: "stableusd",
			Name: "Stably USD",
		},
	},
	"MODX": {
		{
			ID:   3479,
			Slug: "model-x-coin",
			Name: "MODEL-X-coin",
		},
	},
	"UNIUP": {
		{
			ID:   7524,
			Slug: "uniup",
			Name: "UNIUP",
		},
	},
	"PROMISE": {
		{
			ID:   10516,
			Slug: "promise",
			Name: "Promise",
		},
	},
	"DVPN": {
		{
			ID:   2643,
			Slug: "sentinel",
			Name: "Sentinel",
		},
	},
	"KTLYO": {
		{
			ID:   7885,
			Slug: "katalyo",
			Name: "Katalyo",
		},
	},
	"ETHYS": {
		{
			ID:   7803,
			Slug: "ethereum-stake",
			Name: "Ethereum Stake",
		},
	},
	"WCC": {
		{
			ID:   5233,
			Slug: "wincash",
			Name: "WinCash",
		},
	},
	"NANOX": {
		{
			ID:   1691,
			Slug: "project-x",
			Name: "Project-X",
		},
	},
	"PACT": {
		{
			ID:   10087,
			Slug: "pact-community-token",
			Name: "PACT community token",
		},
	},
	"HORD": {
		{
			ID:   9198,
			Slug: "hord",
			Name: "Hord",
		},
	},
	"MXX": {
		{
			ID:   6583,
			Slug: "multiplier",
			Name: "Multiplier",
		},
	},
	"ADB": {
		{
			ID:   2501,
			Slug: "adbank",
			Name: "adbank",
		},
	},
	"VTX": {
		{
			ID:   8661,
			Slug: "vortex-defi",
			Name: "Vortex Defi",
		},
	},
	"SG": {
		{
			ID:   6245,
			Slug: "socialgood",
			Name: "SocialGood",
		},
	},
	"RYOSHI": {
		{
			ID:   10817,
			Slug: "ryoshi-token",
			Name: "Ryoshi Token",
		},
	},
	"KFC": {
		{
			ID:   7169,
			Slug: "chicken",
			Name: "Chicken",
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
	"LAYER": {
		{
			ID:   6638,
			Slug: "unilayer",
			Name: "UniLayer",
		},
	},
	"MBX": {
		{
			ID:   7437,
			Slug: "mobiepay",
			Name: "MobieCoin",
		},
	},
	"XAUTBEAR": {
		{
			ID:   6240,
			Slug: "3x-short-tether-gold-token",
			Name: "3X Short Tether Gold Token",
		},
	},
	"SATS": {
		{
			ID:   9022,
			Slug: "satoshi",
			Name: "Satoshi",
		},
	},
	"EVT": {
		{
			ID:   4238,
			Slug: "everitoken",
			Name: "EveriToken",
		},
	},
	"OFC": {
		{
			ID:   10794,
			Slug: "ofc-coin",
			Name: "$OFC Coin",
		},
	},
	"KIM": {
		{
			ID:   5286,
			Slug: "kingmoney",
			Name: "KingMoney",
		},
	},
	"FLASH": {
		{
			ID:   1755,
			Slug: "flash",
			Name: "Flash",
		},
	},
	"0xMR": {
		{
			ID:   5668,
			Slug: "0xmonero",
			Name: "0xMonero",
		},
	},
	"KIDS": {
		{
			ID:   10318,
			Slug: "save-the-kids",
			Name: "Save The Kids",
		},
	},
	"BAND": {
		{
			ID:   4679,
			Slug: "band-protocol",
			Name: "Band Protocol",
		},
	},
	"CONV": {
		{
			ID:   8716,
			Slug: "convergence",
			Name: "Convergence",
		},
	},
	"GOC": {
		{
			ID:   3052,
			Slug: "gocrypto-token",
			Name: "GoCrypto Token",
		},
	},
	"2KEY": {
		{
			ID:   5587,
			Slug: "2key-network",
			Name: "2key.network",
		},
	},
	"FSHN": {
		{
			ID:   6446,
			Slug: "fashion-coin",
			Name: "Fashion Coin",
		},
	},
	"BBANK": {
		{
			ID:   9839,
			Slug: "blockbank",
			Name: "BlockBank",
		},
	},
	"CHEX": {
		{
			ID:   8534,
			Slug: "chex-token",
			Name: "Chintai",
		},
	},
	"MKMOON": {
		{
			ID:   10057,
			Slug: "monkeycoin",
			Name: "MonkeyCoin",
		},
	},
	"VIA": {
		{
			ID:   470,
			Slug: "viacoin",
			Name: "Viacoin",
		},
	},
	"MoFi": {
		{
			ID:   9132,
			Slug: "mobi-finance",
			Name: "MobiFi",
		},
	},
	"MCH": {
		{
			ID:   4844,
			Slug: "meconcash",
			Name: "MeconCash",
		},
	},
	"VNS": {
		{
			ID:   7078,
			Slug: "va-na-su",
			Name: "Va Na Su",
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
	"MSTR": {
		{
			ID:   7898,
			Slug: "microstrategy-tokenized-stock-ftx",
			Name: "MicroStrategy tokenized stock FTX",
		},
	},
	"METP": {
		{
			ID:   5448,
			Slug: "metaprediction",
			Name: "Metaprediction",
		},
	},
	"BPCN": {
		{
			ID:   6242,
			Slug: "blipcoin",
			Name: "BlipCoin",
		},
	},
	"FTT": {
		{
			ID:   4195,
			Slug: "ftx-token",
			Name: "FTX Token",
		},
	},
	"ARMOR": {
		{
			ID:   8309,
			Slug: "armor",
			Name: "ARMOR",
		},
	},
	"BLKC": {
		{
			ID:   10566,
			Slug: "blackhat",
			Name: "BlackHat",
		},
	},
	"SUMO": {
		{
			ID:   1694,
			Slug: "sumokoin",
			Name: "Sumokoin",
		},
	},
	"CELL": {
		{
			ID:   8992,
			Slug: "cellframe",
			Name: "Cellframe",
		},
	},
	"CETH": {
		{
			ID:   5636,
			Slug: "compound-ether",
			Name: "Compound Ether",
		},
	},
	"GRID": {
		{
			ID:   2134,
			Slug: "grid",
			Name: "Grid+",
		},
	},
	"MMON": {
		{
			ID:   10480,
			Slug: "mammon",
			Name: "Mammon",
		},
	},
	"DSWAP": {
		{
			ID:   8195,
			Slug: "definex",
			Name: "Definex",
		},
	},
	"META": {
		{
			ID:   3418,
			Slug: "metadium",
			Name: "Metadium",
		},
	},
	"NSURE": {
		{
			ID:   7231,
			Slug: "nsure-network",
			Name: "Nsure.Network",
		},
	},
	"ZLA": {
		{
			ID:   2500,
			Slug: "zilla",
			Name: "Zilla",
		},
	},
	"INX": {
		{
			ID:   5840,
			Slug: "insight-protocol",
			Name: "Insight Protocol",
		},
	},
	"NBXC": {
		{
			ID:   4292,
			Slug: "nibble",
			Name: "Nibble",
		},
	},
	"BLVR": {
		{
			ID:   5855,
			Slug: "believer",
			Name: "BELIEVER",
		},
	},
	"FRY": {
		{
			ID:   6119,
			Slug: "foundry",
			Name: "Foundry",
		},
	},
	"INTO": {
		{
			ID:   10054,
			Slug: "infiniti",
			Name: "Infiniti",
		},
	},
	"YOUC": {
		{
			ID:   7321,
			Slug: "youcash",
			Name: "yOUcash",
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
	"GUSDT": {
		{
			ID:   8118,
			Slug: "gusd-token",
			Name: "Global Utility Smart Digital Token",
		},
	},
	"UPR": {
		{
			ID:   9694,
			Slug: "upfire",
			Name: "Upfire",
		},
	},
	"ICA": {
		{
			ID:   9415,
			Slug: "icarus-finance",
			Name: "Icarus Finance",
		},
	},
	"GOLDUCK": {
		{
			ID:   9394,
			Slug: "golden-duck",
			Name: "Golden Duck",
		},
	},
	"DOGZ": {
		{
			ID:   6422,
			Slug: "dogz",
			Name: "Dogz",
		},
	},
	"VET": {
		{
			ID:   3077,
			Slug: "vechain",
			Name: "VeChain",
		},
	},
	"DFSOCIAL": {
		{
			ID:   8128,
			Slug: "defisocial-gaming",
			Name: "DFSocial Gaming",
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
	"BNY": {
		{
			ID:   4775,
			Slug: "bancacy",
			Name: "Bancacy",
		},
	},
	"BAGS": {
		{
			ID:   8264,
			Slug: "basis-gold-share",
			Name: "Basis Gold Share",
		},
	},
	"ZORT": {
		{
			ID:   10830,
			Slug: "zort",
			Name: "ZORT",
		},
	},
	"FWB": {
		{
			ID:   10090,
			Slug: "friends-with-benefits-pro",
			Name: "Friends With Benefits Pro",
		},
	},
	"MEMES": {
		{
			ID:   8991,
			Slug: "memes-token",
			Name: "Memes Token",
		},
	},
	"TRIX": {
		{
			ID:   6801,
			Slug: "triumphx",
			Name: "TriumphX",
		},
	},
	"NETZ": {
		{
			ID:   10997,
			Slug: "netzcoin",
			Name: "Netzcoin",
		},
	},
	"DILI": {
		{
			ID:   4793,
			Slug: "d-community",
			Name: "D Community",
		},
	},
	"GUAP": {
		{
			ID:   5088,
			Slug: "guapcoin",
			Name: "Guapcoin",
		},
	},
	"WANLINK": {
		{
			ID:   8651,
			Slug: "wanlink",
			Name: "wanLINK",
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
	"vXVS": {
		{
			ID:   7960,
			Slug: "venus-xvs",
			Name: "Venus XVS",
		},
	},
	"JGN": {
		{
			ID:   6942,
			Slug: "juggernaut",
			Name: "Juggernaut",
		},
	},
	"INK": {
		{
			ID:   2209,
			Slug: "ink",
			Name: "Ink",
		},
	},
	"SZC": {
		{
			ID:   7988,
			Slug: "zugacoin",
			Name: "Zugacoin",
		},
	},
	"OXO": {
		{
			ID:   9569,
			Slug: "oxo-farm",
			Name: "OXO.Farm",
		},
	},
	"MAAAVE": {
		{
			ID:   8920,
			Slug: "matic-aave-aave",
			Name: "Matic Aave Interest Bearing AAVE",
		},
	},
	"RING": {
		{
			ID:   5798,
			Slug: "darwinia-network",
			Name: "Darwinia Network",
		},
	},
	"FLS": {
		{
			ID:   4491,
			Slug: "flits",
			Name: "Flits",
		},
	},
	"YFDOT": {
		{
			ID:   7410,
			Slug: "yearn-finance-dot",
			Name: "Yearn Finance DOT",
		},
	},
	"YFOS": {
		{
			ID:   7319,
			Slug: "yfos-finance",
			Name: "YFOS.finance",
		},
	},
	"CRS": {
		{
			ID:   8608,
			Slug: "crypto-rewards-studio",
			Name: "Crypto Rewards Studio",
		},
	},
	"WBIND": {
		{
			ID:   7403,
			Slug: "wrapped-bind",
			Name: "Wrapped BIND",
		},
	},
	"THC": {
		{
			ID:   416,
			Slug: "hempcoin",
			Name: "HempCoin",
		},
	},
	"GIC": {
		{
			ID:   3104,
			Slug: "giant-coin",
			Name: "Giant",
		},
	},
	"GETH": {
		{
			ID:   7908,
			Slug: "guarded-ether",
			Name: "Guarded Ether",
		},
	},
	"CCA": {
		{
			ID:   4122,
			Slug: "counos-coin",
			Name: "Counos Coin",
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
	"TINV": {
		{
			ID:   10611,
			Slug: "tinville",
			Name: "Tinville",
		},
	},
	"OMI": {
		{
			ID:   6432,
			Slug: "ecomi",
			Name: "ECOMI",
		},
	},
	"MORPH": {
		{
			ID:   9408,
			Slug: "morphose",
			Name: "MORPHOSE",
		},
	},
	"XCON": {
		{
			ID:   3932,
			Slug: "connect-coin",
			Name: "Connect Coin",
		},
	},
	"EHASH": {
		{
			ID:   8678,
			Slug: "ehash",
			Name: "EHash",
		},
	},
	"ORO": {
		{
			ID:   7684,
			Slug: "oro",
			Name: "ORO",
		},
	},
	"PDOG": {
		{
			ID:   10626,
			Slug: "polkadog",
			Name: "Polkadog",
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
	"FEVR": {
		{
			ID:   10803,
			Slug: "realfevr",
			Name: "RealFevr",
		},
	},
	"STACY": {
		{
			ID:   7574,
			Slug: "stacy",
			Name: "Stacy",
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
	"TYPH": {
		{
			ID:   8794,
			Slug: "typhoon-network",
			Name: "Typhoon Network",
		},
	},
	"MORE": {
		{
			ID:   1722,
			Slug: "more-coin",
			Name: "More Coin",
		},
	},
	"ISAL": {
		{
			ID:   8275,
			Slug: "isalcoin",
			Name: "ISALCOIN",
		},
	},
	"MCAFEE": {
		{
			ID:   10619,
			Slug: "the-last-mcafee-token",
			Name: "The Last McAfee Token",
		},
	},
	"KOMET": {
		{
			ID:   7760,
			Slug: "komet",
			Name: "Komet",
		},
	},
	"LINK": {
		{
			ID:   1975,
			Slug: "chainlink",
			Name: "Chainlink",
		},
	},
	"VIKKY": {
		{
			ID:   2965,
			Slug: "vikkytoken",
			Name: "VikkyToken",
		},
	},
	"SFMS": {
		{
			ID:   10204,
			Slug: "safemoon-swap",
			Name: "SafeMoon.swap",
		},
	},
	"BHIG": {
		{
			ID:   3820,
			Slug: "buck-hath-coin",
			Name: "BuckHathCoin",
		},
	},
	"mSLV": {
		{
			ID:   8026,
			Slug: "mirrored-ishares-silver-trust",
			Name: "Mirrored iShares Silver Trust",
		},
	},
	"UNISTAKE": {
		{
			ID:   7512,
			Slug: "unistake",
			Name: "Unistake",
		},
	},
	"SUPERBID": {
		{
			ID:   8836,
			Slug: "superbid",
			Name: "Superbid",
		},
	},
	"GLASS": {
		{
			ID:   10113,
			Slug: "ourglass",
			Name: "Ourglass",
		},
	},
	"SCA": {
		{
			ID:   9118,
			Slug: "scaleswap",
			Name: "Scaleswap",
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
	"BSCPAD": {
		{
			ID:   8660,
			Slug: "bscpad",
			Name: "BSCPAD",
		},
	},
	"BASX": {
		{
			ID:   8513,
			Slug: "basix",
			Name: "Basix",
		},
	},
	"LTP": {
		{
			ID:   7790,
			Slug: "lifetioncoin",
			Name: "LifetionCoin",
		},
	},
	"DNS": {
		{
			ID:   8194,
			Slug: "bitdns",
			Name: "BitDNS",
		},
	},
	"PHB": {
		{
			ID:   2112,
			Slug: "phoenix-global",
			Name: "Phoenix Global",
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
	"HTZ": {
		{
			ID:   10824,
			Slug: "hertz-network",
			Name: "Hertz Network",
		},
	},
	"FJC": {
		{
			ID:   960,
			Slug: "fujicoin",
			Name: "FujiCoin",
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
	"SHDW": {
		{
			ID:   1878,
			Slug: "shadow-token",
			Name: "Shadow Token",
		},
	},
	"TM2": {
		{
			ID:   2850,
			Slug: "traxia",
			Name: "TRAXIA",
		},
	},
	"BWB": {
		{
			ID:   6004,
			Slug: "bw-token",
			Name: "Bit World Token",
		},
	},
	"YLAND": {
		{
			ID:   7234,
			Slug: "yearn-land",
			Name: "Yearn Land",
		},
	},
	"BASID": {
		{
			ID:   7157,
			Slug: "basid-coin",
			Name: "Basid Coin",
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
	"BRZ": {
		{
			ID:   4139,
			Slug: "brz",
			Name: "Brazilian Digital Token",
		},
	},
	"MVI": {
		{
			ID:   9207,
			Slug: "metaverse-index",
			Name: "Metaverse Index",
		},
	},
	"BNA": {
		{
			ID:   5792,
			Slug: "bananatok",
			Name: "Bananatok",
		},
	},
	"SRX": {
		{
			ID:   10894,
			Slug: "storx-network",
			Name: "StorX Network",
		},
	},
	"PORTAL": {
		{
			ID:   6496,
			Slug: "portal",
			Name: "Portal",
		},
	},
	"PVT": {
		{
			ID:   4115,
			Slug: "pivot-token",
			Name: "Pivot Token",
		},
	},
	"BNBUP": {
		{
			ID:   7009,
			Slug: "bnbup",
			Name: "BNBUP",
		},
	},
	"SPORTS": {
		{
			ID:   6564,
			Slug: "zensports",
			Name: "ZenSports",
		},
	},
	"SEMI": {
		{
			ID:   6823,
			Slug: "semitoken",
			Name: "Semitoken",
		},
	},
	"LICK": {
		{
			ID:   10812,
			Slug: "vitalick-neuterin",
			Name: "VITALICK NEUTERIN",
		},
	},
	"FOXT": {
		{
			ID:   2966,
			Slug: "fox-trading",
			Name: "Fox Trading",
		},
	},
	"XMV": {
		{
			ID:   3902,
			Slug: "monero-v",
			Name: "MoneroV ",
		},
	},
	"LPL": {
		{
			ID:   9062,
			Slug: "linkpool",
			Name: "LinkPool",
		},
	},
	"BLZN": {
		{
			ID:   6723,
			Slug: "blaze-network",
			Name: "Blaze Network",
		},
	},
	"XMS": {
		{
			ID:   10030,
			Slug: "mars-ecosystem-governance-token",
			Name: "Mars Ecosystem Token",
		},
	},
	"CWAP": {
		{
			ID:   9817,
			Slug: "defire",
			Name: "DeFIRE",
		},
	},
	"LID": {
		{
			ID:   5986,
			Slug: "liquidity-dividends-protocol",
			Name: "Liquidity Dividends Protocol",
		},
	},
	"AXIOM": {
		{
			ID:   1020,
			Slug: "axiom",
			Name: "Axiom",
		},
	},
	"TBP": {
		{
			ID:   7777,
			Slug: "tradebitpay",
			Name: "Tradebitpay",
		},
	},
	"VID": {
		{
			ID:   4300,
			Slug: "videocoin",
			Name: "VideoCoin",
		},
	},
	"TOL": {
		{
			ID:   3389,
			Slug: "tolar",
			Name: "Tolar",
		},
	},
	"IDALL": {
		{
			ID:   7514,
			Slug: "idall",
			Name: "IDall",
		},
	},
	"PERA": {
		{
			ID:   10225,
			Slug: "pera-finance",
			Name: "Pera Finance",
		},
	},
	"LOON": {
		{
			ID:   8173,
			Slug: "loon-network",
			Name: "Loon Network",
		},
	},
	"RWN": {
		{
			ID:   5886,
			Slug: "rowan-token",
			Name: "Rowan Token",
		},
	},
	"ACPT": {
		{
			ID:   7116,
			Slug: "crypto-accept",
			Name: "Crypto Accept",
		},
	},
	"XLM": {
		{
			ID:   512,
			Slug: "stellar",
			Name: "Stellar",
		},
	},
	"AOA": {
		{
			ID:   2874,
			Slug: "aurora",
			Name: "Aurora",
		},
	},
	"BTSG": {
		{
			ID:   8905,
			Slug: "bitsong",
			Name: "BitSong",
		},
	},
	"TOPC": {
		{
			ID:   2376,
			Slug: "topchain",
			Name: "TopChain",
		},
	},
	"CLIMB": {
		{
			ID:   9295,
			Slug: "climb-token-finance",
			Name: "CLIMB TOKEN FINANCE",
		},
	},
	"SRM": {
		{
			ID:   6187,
			Slug: "serum",
			Name: "Serum",
		},
	},
	"RBC": {
		{
			ID:   7219,
			Slug: "rubic",
			Name: "Rubic",
		},
	},
	"VIB": {
		{
			ID:   2019,
			Slug: "viberate",
			Name: "Viberate",
		},
	},
	"DTA": {
		{
			ID:   2446,
			Slug: "data",
			Name: "DATA",
		},
	},
	"AC3": {
		{
			ID:   2722,
			Slug: "ac3",
			Name: "AC3",
		},
	},
	"GTX": {
		{
			ID:   5366,
			Slug: "goaltime-n",
			Name: "GoalTime N",
		},
	},
	"GMEE": {
		{
			ID:   9103,
			Slug: "gamee",
			Name: "GAMEE",
		},
	},
	"GOF": {
		{
			ID:   7034,
			Slug: "golff",
			Name: "Golff",
		},
	},
	"GRUMPY": {
		{
			ID:   8816,
			Slug: "grumpy-finance",
			Name: "Grumpy.finance",
		},
	},
	"ZSC": {
		{
			ID:   2047,
			Slug: "zeusshield",
			Name: "Zeusshield",
		},
	},
	"LUXO": {
		{
			ID:   8443,
			Slug: "luxochain",
			Name: "LUXOCHAIN",
		},
	},
	"FETCH": {
		{
			ID:   10112,
			Slug: "moonretriever",
			Name: "MoonRetriever",
		},
	},
	"CUSD": {
		{
			ID:   7236,
			Slug: "celo-dollar",
			Name: "Celo Dollar",
		},
	},
	"YAXIS": {
		{
			ID:   7222,
			Slug: "yaxis",
			Name: "yAxis",
		},
	},
	"AXO": {
		{
			ID:   10581,
			Slug: "axolotl-finance",
			Name: "Axolotl Finance",
		},
	},
	"CELO": {
		{
			ID:   5567,
			Slug: "celo",
			Name: "Celo",
		},
	},
	"COMBO": {
		{
			ID:   8259,
			Slug: "furucombo",
			Name: "Furucombo",
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
	"KTS": {
		{
			ID:   4018,
			Slug: "klimatas",
			Name: "Klimatas",
		},
	},
	"2LC": {
		{
			ID:   9268,
			Slug: "2local",
			Name: "2local",
		},
	},
	"SHO": {
		{
			ID:   6540,
			Slug: "showcase",
			Name: "Showcase",
		},
	},
	"NBC": {
		{
			ID:   3095,
			Slug: "niobium-coin",
			Name: "Niobium Coin",
		},
	},
	"GOMA": {
		{
			ID:   10858,
			Slug: "goma-finance",
			Name: "GOMA Finance",
		},
	},
	"MOCA": {
		{
			ID:   10220,
			Slug: "museum-of-crypto-art",
			Name: "Museum of Crypto Art",
		},
	},
	"SSF": {
		{
			ID:   10615,
			Slug: "secretsky-finance",
			Name: "SecretSky.finance",
		},
	},
	"FN": {
		{
			ID:   4817,
			Slug: "filenet",
			Name: "Filenet",
		},
	},
	"XED": {
		{
			ID:   8163,
			Slug: "exeedme",
			Name: "Exeedme",
		},
	},
	"SVR": {
		{
			ID:   4594,
			Slug: "sovranocoin",
			Name: "SovranoCoin",
		},
	},
	"UST": {
		{
			ID:   7129,
			Slug: "terrausd",
			Name: "TerraUSD",
		},
	},
	"LTO": {
		{
			ID:   3714,
			Slug: "lto-network",
			Name: "LTO Network",
		},
	},
	"HYN": {
		{
			ID:   3695,
			Slug: "hyperion",
			Name: "Hyperion",
		},
	},
	"ORT": {
		{
			ID:   10756,
			Slug: "omni-real-estate-token",
			Name: "Omni Real Estate Token",
		},
	},
	"XPO": {
		{
			ID:   9048,
			Slug: "xpool",
			Name: "Xpool",
		},
	},
	"PAD": {
		{
			ID:   10610,
			Slug: "smartpad",
			Name: "SMARTPAD",
		},
	},
	"CAROM": {
		{
			ID:   9248,
			Slug: "carillonium-finance",
			Name: "Carillonium finance",
		},
	},
	"YFI3": {
		{
			ID:   7794,
			Slug: "yfi3-money",
			Name: "YFI3.money",
		},
	},
	"KIAN": {
		{
			ID:   9265,
			Slug: "kianite-finance",
			Name: "Kianite Finance",
		},
	},
	"EFK": {
		{
			ID:   6761,
			Slug: "refork",
			Name: "ReFork",
		},
	},
	"yTSLA": {
		{
			ID:   7140,
			Slug: "ytsla-finance",
			Name: "yTSLA Finance",
		},
	},
	"ETM": {
		{
			ID:   4108,
			Slug: "en-tan-mo",
			Name: "En-Tan-Mo",
		},
	},
	"AGAR": {
		{
			ID:   8251,
			Slug: "agar",
			Name: "AGAr",
		},
	},
	"GARD": {
		{
			ID:   2938,
			Slug: "hashgard",
			Name: "Hashgard",
		},
	},
	"XTZUP": {
		{
			ID:   7007,
			Slug: "xtzup",
			Name: "XTZUP",
		},
	},
	"$PAWG": {
		{
			ID:   10538,
			Slug: "pawgcoin",
			Name: "PAWGcoin",
		},
	},
	"KOMP": {
		{
			ID:   7264,
			Slug: "kompass",
			Name: "Kompass",
		},
	},
	"ASAC": {
		{
			ID:   5276,
			Slug: "asac-coin",
			Name: "Asac Coin",
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
	"XCH": {
		{
			ID:   9258,
			Slug: "chia-network",
			Name: "Chia Network",
		},
	},
	"N1CE": {
		{
			ID:   10689,
			Slug: "n1ce",
			Name: "N1CE",
		},
	},
	"LOV": {
		{
			ID:   7634,
			Slug: "the-lovechain",
			Name: "The LoveChain",
		},
	},
	"AKNC": {
		{
			ID:   5750,
			Slug: "aave-knc",
			Name: "Aave KNC",
		},
	},
	"HUSH": {
		{
			ID:   1466,
			Slug: "hush",
			Name: "Hush",
		},
	},
	"MARTK": {
		{
			ID:   5520,
			Slug: "martkist",
			Name: "Martkist",
		},
	},
	"BABYCUBAN": {
		{
			ID:   10995,
			Slug: "baby-cuban",
			Name: "Baby Cuban",
		},
	},
	"GDILDO": {
		{
			ID:   10618,
			Slug: "green-dildo-finance",
			Name: "Green Dildo Finance",
		},
	},
	"PRIX": {
		{
			ID:   2184,
			Slug: "privatix",
			Name: "Privatix",
		},
	},
	"SECO": {
		{
			ID:   8220,
			Slug: "serum-ecosystem-token",
			Name: "Serum Ecosystem Token",
		},
	},
	"KSF": {
		{
			ID:   9871,
			Slug: "kesef-finance",
			Name: "Kesef Finance",
		},
	},
	"CBP": {
		{
			ID:   5781,
			Slug: "cashbackpro",
			Name: "CashBackPro",
		},
	},
	"WANUNI": {
		{
			ID:   8654,
			Slug: "wanuni",
			Name: "wanUNI",
		},
	},
	"FTO": {
		{
			ID:   2846,
			Slug: "futurocoin",
			Name: "FuturoCoin",
		},
	},
	"BCH": {
		{
			ID:   1831,
			Slug: "bitcoin-cash",
			Name: "Bitcoin Cash",
		},
	},
	"TUBE": {
		{
			ID:   2561,
			Slug: "bit-tube",
			Name: "BitTube",
		},
	},
	"CUR": {
		{
			ID:   5083,
			Slug: "curio",
			Name: "Curio",
		},
	},
	"KICKS": {
		{
			ID:   4273,
			Slug: "sessia",
			Name: "Sessia",
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
	"MVEDA": {
		{
			ID:   7641,
			Slug: "medicalveda",
			Name: "Medicalveda",
		},
	},
	"JWL": {
		{
			ID:   3791,
			Slug: "jewel",
			Name: "Jewel",
		},
	},
	"ME": {
		{
			ID:   6026,
			Slug: "all-me",
			Name: "All.me",
		},
	},
	"TKMN": {
		{
			ID:   8338,
			Slug: "tokemon",
			Name: "Tokemon",
		},
	},
	"ATOM": {
		{
			ID:   3794,
			Slug: "cosmos",
			Name: "Cosmos",
		},
	},
	"QBIC": {
		{
			ID:   2460,
			Slug: "qbic",
			Name: "Qbic",
		},
	},
	"UDOGE": {
		{
			ID:   10788,
			Slug: "uncle-doge",
			Name: "Uncle Doge",
		},
	},
	"CVP": {
		{
			ID:   6669,
			Slug: "powerpool",
			Name: "PowerPool",
		},
	},
	"VITAE": {
		{
			ID:   3063,
			Slug: "vitae",
			Name: "Vitae",
		},
	},
	"ADM": {
		{
			ID:   3703,
			Slug: "adamant-messenger",
			Name: "ADAMANT Messenger",
		},
	},
	"DAB": {
		{
			ID:   4284,
			Slug: "dabanking",
			Name: "DABANKING",
		},
	},
	"VBSWAP": {
		{
			ID:   8865,
			Slug: "vbswap",
			Name: "vBSWAP",
		},
	},
	"BCAPS": {
		{
			ID:   9221,
			Slug: "binacaps",
			Name: "Binacaps",
		},
	},
	"GS": {
		{
			ID:   9196,
			Slug: "genesis-shards",
			Name: "Genesis Shards",
		},
	},
	"RAPTR": {
		{
			ID:   9628,
			Slug: "raptor-token",
			Name: "Raptor Token",
		},
	},
	"vFIL": {
		{
			ID:   8213,
			Slug: "venus-filecoin",
			Name: "Venus Filecoin",
		},
	},
	"OLXA": {
		{
			ID:   3823,
			Slug: "olxa",
			Name: "OLXA",
		},
	},
	"TMED": {
		{
			ID:   6237,
			Slug: "mdsquare",
			Name: "MDsquare",
		},
	},
	"EXRT": {
		{
			ID:   8807,
			Slug: "exrt-network",
			Name: "EXRT Network",
		},
	},
	"IT": {
		{
			ID:   6473,
			Slug: "idcm-token",
			Name: "IDCM Token",
		},
	},
	"PCL": {
		{
			ID:   2610,
			Slug: "peculium",
			Name: "Peculium",
		},
	},
	"EBST": {
		{
			ID:   1704,
			Slug: "eboostcoin",
			Name: "eBoost",
		},
	},
	"NUTS": {
		{
			ID:   6986,
			Slug: "squirrel-finance",
			Name: "Squirrel Finance",
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
	"XPY": {
		{
			ID:   764,
			Slug: "paycoin2",
			Name: "PayCoin",
		},
	},
	"ATTN": {
		{
			ID:   6803,
			Slug: "attn",
			Name: "ATTN",
		},
	},
	"GBTC": {
		{
			ID:   8344,
			Slug: "grayscale-bitcoin-trust-tokenized-stock-ftx",
			Name: "Grayscale Bitcoin Trust tokenized stock FTX",
		},
	},
	"SBS": {
		{
			ID:   8283,
			Slug: "staysbase",
			Name: "StaysBASE",
		},
	},
	"SPK": {
		{
			ID:   2448,
			Slug: "sparkspay",
			Name: "SparksPay",
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
	"FLO": {
		{
			ID:   8774,
			Slug: "flourmix",
			Name: "FlourMix",
		},
	},
	"vDOGE": {
		{
			ID:   9427,
			Slug: "venus-dogecoin",
			Name: "Venus Dogecoin",
		},
	},
	"$RICH": {
		{
			ID:   10068,
			Slug: "richierich-coin",
			Name: "RichieRich Coin",
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
	"JNB": {
		{
			ID:   3759,
			Slug: "jinbi-token",
			Name: "Jinbi Token",
		},
	},
	"WANEOS": {
		{
			ID:   8652,
			Slug: "waneos",
			Name: "wanEOS",
		},
	},
	"EMAX": {
		{
			ID:   9855,
			Slug: "ethereummax",
			Name: "EthereumMax",
		},
	},
	"XCM": {
		{
			ID:   4105,
			Slug: "coinmetro-token",
			Name: "CoinMetro Token",
		},
	},
	"STOPELON": {
		{
			ID:   10426,
			Slug: "stopelon",
			Name: "Stopelon",
		},
	},
	"WSOTE": {
		{
			ID:   8381,
			Slug: "soteria",
			Name: "Soteria",
		},
	},
	"UAT": {
		{
			ID:   4262,
			Slug: "ultralpha",
			Name: "UltrAlpha",
		},
	},
	"vUSDT": {
		{
			ID:   7957,
			Slug: "venus-usdt",
			Name: "Venus USDT",
		},
	},
	"SWACE": {
		{
			ID:   4199,
			Slug: "swace",
			Name: "Swace",
		},
	},
	"777": {
		{
			ID:   7305,
			Slug: "jackpot",
			Name: "Jackpot",
		},
	},
	"DEQ": {
		{
			ID:   8224,
			Slug: "dequant",
			Name: "Dequant",
		},
	},
	"ROBO": {
		{
			ID:   9054,
			Slug: "robo-token",
			Name: "Robo Token",
		},
	},
	"ETHM": {
		{
			ID:   3589,
			Slug: "ethereum-meta",
			Name: "Ethereum Meta",
		},
	},
	"BLC": {
		{
			ID:   125,
			Slug: "blakecoin",
			Name: "Blakecoin",
		},
	},
	"MBTC": {
		{
			ID:   10245,
			Slug: "micro-bitcoin-finance",
			Name: "Micro Bitcoin Finance",
		},
	},
	"OX": {
		{
			ID:   9384,
			Slug: "orcax",
			Name: "OrcaX",
		},
	},
	"VOLTZ": {
		{
			ID:   4508,
			Slug: "voltz",
			Name: "Voltz",
		},
	},
	"CATX": {
		{
			ID:   6703,
			Slug: "cat-trade-protocol",
			Name: "CAT.trade Protocol",
		},
	},
	"PRS": {
		{
			ID:   2455,
			Slug: "pressone",
			Name: "PressOne",
		},
	},
	"BIS": {
		{
			ID:   2009,
			Slug: "bismuth",
			Name: "Bismuth",
		},
	},
	"SHND": {
		{
			ID:   1106,
			Slug: "stronghands",
			Name: "StrongHands",
		},
	},
	"WG0": {
		{
			ID:   7472,
			Slug: "wrapped-gen-0-cryptokitties",
			Name: "Wrapped Gen-0 CryptoKitties",
		},
	},
	"APR": {
		{
			ID:   2721,
			Slug: "apr-coin",
			Name: "APR Coin",
		},
	},
	"CPHR": {
		{
			ID:   10978,
			Slug: "polkacipher",
			Name: "PolkaCipher",
		},
	},
	"ZEP": {
		{
			ID:   9300,
			Slug: "zeppelin-dao",
			Name: "Zeppelin DAO",
		},
	},
	"QUAI": {
		{
			ID:   8373,
			Slug: "quai-dao",
			Name: "QUAI DAO",
		},
	},
	"SRAT": {
		{
			ID:   10579,
			Slug: "spacerat",
			Name: "SpaceRat",
		},
	},
	"PIS": {
		{
			ID:   8169,
			Slug: "polkainsure-finance",
			Name: "Polkainsure Finance",
		},
	},
	"BONO": {
		{
			ID:   5320,
			Slug: "bonorum",
			Name: "Bonorum",
		},
	},
	"$BOOB": {
		{
			ID:   10144,
			Slug: "boob",
			Name: "$BOOB",
		},
	},
	"INEX": {
		{
			ID:   5300,
			Slug: "inex-project",
			Name: "Inex Project",
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
	"EBK": {
		{
			ID:   4778,
			Slug: "ebakus",
			Name: "ebakus",
		},
	},
	"KUS": {
		{
			ID:   10908,
			Slug: "kuswap",
			Name: "KuSwap",
		},
	},
	"NFTS": {
		{
			ID:   10311,
			Slug: "nft-stars",
			Name: "NFT STARS",
		},
	},
	"PIVX": {
		{
			ID:   1169,
			Slug: "pivx",
			Name: "PIVX",
		},
	},
	"TENT": {
		{
			ID:   2912,
			Slug: "tent",
			Name: "TENT",
		},
	},
	"ALP": {
		{
			ID:   4820,
			Slug: "alp-coin",
			Name: "ALP Coin",
		},
	},
	"TRYB": {
		{
			ID:   5181,
			Slug: "bilira",
			Name: "BiLira",
		},
	},
	"PALLA": {
		{
			ID:   10669,
			Slug: "pallapay",
			Name: "Pallapay",
		},
	},
	"MOOCHII": {
		{
			ID:   10077,
			Slug: "moochii",
			Name: "MOOCHII",
		},
	},
	"HAUS": {
		{
			ID:   9016,
			Slug: "daohaus",
			Name: "DAOhaus",
		},
	},
	"GYA": {
		{
			ID:   9101,
			Slug: "gya",
			Name: "GYA",
		},
	},
	"CNX": {
		{
			ID:   5332,
			Slug: "cofinex-coin",
			Name: "Cofinex Coin",
		},
	},
	"ENERGYX": {
		{
			ID:   10325,
			Slug: "safe-energy",
			Name: "Safe Energy",
		},
	},
	"YMAX": {
		{
			ID:   7124,
			Slug: "ymax",
			Name: "YMAX",
		},
	},
	"BGG": {
		{
			ID:   3587,
			Slug: "bgogo-token",
			Name: "Bgogo Token",
		},
	},
	"SWISE": {
		{
			ID:   10439,
			Slug: "stakewise",
			Name: "StakeWise",
		},
	},
	"EKO": {
		{
			ID:   2391,
			Slug: "echolink",
			Name: "EchoLink",
		},
	},
	"MALINK": {
		{
			ID:   8923,
			Slug: "matic-aave-link",
			Name: "Matic Aave Interest Bearing LINK",
		},
	},
	"NINJA": {
		{
			ID:   4609,
			Slug: "ninjacoin",
			Name: "Ninjacoin",
		},
	},
	"CERBERUS": {
		{
			ID:   10438,
			Slug: "gocerberus",
			Name: "GoCerberus",
		},
	},
	"LIKE": {
		{
			ID:   2909,
			Slug: "likecoin",
			Name: "LikeCoin",
		},
	},
	"NIO": {
		{
			ID:   7888,
			Slug: "nio-tokenized-stock-ftx",
			Name: "Nio tokenized stock FTX",
		},
	},
	"RGP": {
		{
			ID:   9225,
			Slug: "rigel-protocol",
			Name: "Rigel Protocol",
		},
	},
	"KCAL": {
		{
			ID:   7298,
			Slug: "phantasma-energy",
			Name: "Phantasma Energy",
		},
	},
	"ISLE": {
		{
			ID:   10080,
			Slug: "island-coin",
			Name: "Island Coin",
		},
	},
	"IFEX": {
		{
			ID:   7949,
			Slug: "interfinex",
			Name: "Interfinex",
		},
	},
	"MCS": {
		{
			ID:   10731,
			Slug: "mcs-token",
			Name: "MCS Token",
		},
	},
	"GPT": {
		{
			ID:   7946,
			Slug: "grace-period-token",
			Name: "Grace Period Token",
		},
	},
	"ITL": {
		{
			ID:   3476,
			Slug: "italian-lira",
			Name: "Italian Lira",
		},
	},
	"BAZT": {
		{
			ID:   5050,
			Slug: "baz-token",
			Name: "Baz Token",
		},
	},
	"STRIKE": {
		{
			ID:   9220,
			Slug: "strikecoin",
			Name: "StrikeCoin",
		},
	},
	"KONG": {
		{
			ID:   8998,
			Slug: "kong-defi",
			Name: "Kong Defi",
		},
	},
	"JST": {
		{
			ID:   5488,
			Slug: "just",
			Name: "JUST",
		},
	},
	"WHALE": {
		{
			ID:   6679,
			Slug: "whale",
			Name: "WHALE",
		},
	},
	"ISR": {
		{
			ID:   3466,
			Slug: "insureum",
			Name: "Insureum",
		},
	},
	"LSS": {
		{
			ID:   10103,
			Slug: "lossless",
			Name: "Lossless",
		},
	},
	"SIB": {
		{
			ID:   1082,
			Slug: "sibcoin",
			Name: "SIBCoin",
		},
	},
	"DTC": {
		{
			ID:   69,
			Slug: "datacoin",
			Name: "Datacoin",
		},
	},
	"XTZBULL": {
		{
			ID:   5464,
			Slug: "3x-long-tezos-token",
			Name: "3x Long Tezos Token",
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
	"TT": {
		{
			ID:   3930,
			Slug: "thunder-token",
			Name: "Thunder Token",
		},
	},
	"MRX": {
		{
			ID:   1814,
			Slug: "metrix-coin",
			Name: "Metrix Coin",
		},
	},
	"XMC": {
		{
			ID:   2655,
			Slug: "monero-classic",
			Name: "Monero Classic",
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
	"UWL": {
		{
			ID:   7727,
			Slug: "uniwhales",
			Name: "UniWhales",
		},
	},
	"TKNT": {
		{
			ID:   7238,
			Slug: "tkn-token",
			Name: "TKN Token",
		},
	},
	"BAAS": {
		{
			ID:   3142,
			Slug: "baasid",
			Name: "BaaSid",
		},
	},
	"MAS": {
		{
			ID:   3352,
			Slug: "midasprotocol",
			Name: "MidasProtocol",
		},
	},
	"SXPDOWN": {
		{
			ID:   7529,
			Slug: "sxpdown",
			Name: "SXPDOWN",
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
	"PLUS1": {
		{
			ID:   3456,
			Slug: "plusonecoin",
			Name: "PlusOneCoin",
		},
	},
	"PMEER": {
		{
			ID:   5144,
			Slug: "qitmeer",
			Name: "Qitmeer",
		},
	},
	"RED": {
		{
			ID:   2771,
			Slug: "red",
			Name: "RED",
		},
	},
	"PCI": {
		{
			ID:   5275,
			Slug: "payprotocol",
			Name: "PayProtocol",
		},
	},
	"BABYSHIB": {
		{
			ID:   10886,
			Slug: "babyshibby-inu",
			Name: "BabyShibby Inu",
		},
	},
	"ADABEAR": {
		{
			ID:   6040,
			Slug: "3x-short-cardano-token",
			Name: "3X Short Cardano Token",
		},
	},
	"PUFFY": {
		{
			ID:   10168,
			Slug: "puffydog-coin",
			Name: "Puffydog Coin",
		},
	},
	"TASTE": {
		{
			ID:   10329,
			Slug: "tastenft",
			Name: "TasteNFT",
		},
	},
	"COSMIC": {
		{
			ID:   10521,
			Slug: "cosmicswap",
			Name: "CosmicSwap",
		},
	},
	"UQC": {
		{
			ID:   2273,
			Slug: "uquid-coin",
			Name: "Uquid Coin",
		},
	},
	"STAKE": {
		{
			ID:   5601,
			Slug: "xdai",
			Name: "xDai",
		},
	},
	"VEE": {
		{
			ID:   2223,
			Slug: "blockv",
			Name: "BLOCKv",
		},
	},
	"MHC": {
		{
			ID:   3756,
			Slug: "metahash",
			Name: "#MetaHash",
		},
	},
	"WAVAX": {
		{
			ID:   9462,
			Slug: "wavax",
			Name: "Wrapped AVAX",
		},
	},
	"WKCS": {
		{
			ID:   11023,
			Slug: "wrapped-kucoin-token",
			Name: "Wrapped KuCoin Token",
		},
	},
	"CVXCRV": {
		{
			ID:   10291,
			Slug: "convex-crv",
			Name: "Convex CRV",
		},
	},
	"mAMZN": {
		{
			ID:   8016,
			Slug: "mirrored-amazon",
			Name: "Mirrored Amazon",
		},
	},
	"SLICE": {
		{
			ID:   8637,
			Slug: "tranche-finance",
			Name: "Tranche Finance",
		},
	},
	"PIN": {
		{
			ID:   64,
			Slug: "public-index-network",
			Name: "Public Index Network",
		},
	},
	"1MIL": {
		{
			ID:   9344,
			Slug: "1millionnfts",
			Name: "1MillionNFTs",
		},
	},
	"JUR": {
		{
			ID:   6482,
			Slug: "jur",
			Name: "Jur",
		},
	},
	"FTX": {
		{
			ID:   2667,
			Slug: "fintrux-network",
			Name: "FintruX Network",
		},
	},
	"ZINC": {
		{
			ID:   2883,
			Slug: "zinc",
			Name: "ZINC",
		},
	},
	"HNB": {
		{
			ID:   3947,
			Slug: "hashnet-biteco",
			Name: "HashNet BitEco",
		},
	},
	"HQT": {
		{
			ID:   3615,
			Slug: "hyperquant",
			Name: "HyperQuant",
		},
	},
	"BUND": {
		{
			ID:   7785,
			Slug: "bundles",
			Name: "Bundles Finance",
		},
	},
	"PSHP": {
		{
			ID:   7407,
			Slug: "payship",
			Name: "Payship",
		},
	},
	"CREAM": {
		{
			ID:   6193,
			Slug: "cream-finance",
			Name: "Cream Finance",
		},
	},
	"OGO": {
		{
			ID:   3985,
			Slug: "origo",
			Name: "Origo",
		},
	},
	"DOPE": {
		{
			ID:   145,
			Slug: "dopecoin",
			Name: "DopeCoin",
		},
	},
	"BURN": {
		{
			ID:   4069,
			Slug: "blockburn",
			Name: "Blockburn",
		},
	},
	"BZZ": {
		{
			ID:   10293,
			Slug: "ethereum-swarm",
			Name: "Swarm",
		},
	},
	"ISP": {
		{
			ID:   9865,
			Slug: "ispolink",
			Name: "Ispolink",
		},
	},
	"DGCL": {
		{
			ID:   8148,
			Slug: "digicol",
			Name: "DigiCol",
		},
	},
	"SMT": {
		{
			ID:   2277,
			Slug: "smartmesh",
			Name: "SmartMesh",
		},
	},
	"SOAR": {
		{
			ID:   8240,
			Slug: "soar-fi",
			Name: "SOAR.FI",
		},
	},
	"NPXS": {
		{
			ID:   2603,
			Slug: "pundi-x",
			Name: "Pundi X[old]",
		},
	},
	"ECOS": {
		{
			ID:   5885,
			Slug: "ecodollar",
			Name: "EcoDollar",
		},
	},
	"COT": {
		{
			ID:   3779,
			Slug: "cotrader",
			Name: "CoTrader",
		},
	},
	"CHAT": {
		{
			ID:   2427,
			Slug: "chatcoin",
			Name: "ChatCoin",
		},
	},
	"STK": {
		{
			ID:   2493,
			Slug: "stk",
			Name: "STK",
		},
	},
	"FOUR": {
		{
			ID:   5409,
			Slug: "4thpillar-technologies",
			Name: "4THPILLAR TECHNOLOGIES",
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
	"LABX": {
		{
			ID:   3751,
			Slug: "stakinglab",
			Name: "Stakinglab",
		},
	},
	"Whale": {
		{
			ID:   10693,
			Slug: "whale-fall",
			Name: "Whale Fall",
		},
	},
	"X2P": {
		{
			ID:   9735,
			Slug: "xenon-pay-ii",
			Name: "Xenon Pay",
		},
	},
	"COCO": {
		{
			ID:   9790,
			Slug: "coco-swap",
			Name: "Coco Swap",
		},
	},
	"EIDOS": {
		{
			ID:   6070,
			Slug: "eidos",
			Name: "EIDOS",
		},
	},
	"YFPRO": {
		{
			ID:   7331,
			Slug: "yfpro-finance",
			Name: "YFPRO Finance",
		},
	},
	"NMP": {
		{
			ID:   5703,
			Slug: "neuromorphic-io",
			Name: "Neuromorphic.io",
		},
	},
	"FLOAT": {
		{
			ID:   9861,
			Slug: "float-protocol-float",
			Name: "Float Protocol: Float",
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
	"ZGOAT": {
		{
			ID:   9705,
			Slug: "goat-zuckerberg",
			Name: "GOAT Zuckerberg",
		},
	},
	"DRS": {
		{
			ID:   5796,
			Slug: "doctors-coin",
			Name: "Doctors Coin",
		},
	},
	"KCLP": {
		{
			ID:   11041,
			Slug: "kucoin-launchpad",
			Name: "KuCoin LaunchPad",
		},
	},
	"BINGUS": {
		{
			ID:   9402,
			Slug: "bingus-token",
			Name: "Bingus Token",
		},
	},
	"IGG": {
		{
			ID:   4054,
			Slug: "ig-gold",
			Name: "IG Gold",
		},
	},
	"CNC": {
		{
			ID:   6018,
			Slug: "global-china-cash",
			Name: "Global China Cash",
		},
	},
	"MEDIA": {
		{
			ID:   9524,
			Slug: "media-network",
			Name: "Media Network",
		},
	},
	"YOK": {
		{
			ID:   6998,
			Slug: "yokcoin",
			Name: "YOKcoin",
		},
	},
	"COMOS": {
		{
			ID:   9122,
			Slug: "comos-finance",
			Name: "COMOS Finance",
		},
	},
	"SHIB": {
		{
			ID:   5994,
			Slug: "shiba-inu",
			Name: "SHIBA INU",
		},
	},
	"PEOS": {
		{
			ID:   3910,
			Slug: "peos",
			Name: "pEOS",
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
	"SADA": {
		{
			ID:   5769,
			Slug: "sada",
			Name: "sADA",
		},
	},
	"OT": {
		{
			ID:   8404,
			Slug: "option-token",
			Name: "Option Token",
		},
	},
	"ARGUS": {
		{
			ID:   1558,
			Slug: "argus",
			Name: "Argus",
		},
	},
	"TOKAU": {
		{
			ID:   10951,
			Slug: "tokyo-au",
			Name: "Tokyo AU",
		},
	},
	"GANJA": {
		{
			ID:   8983,
			Slug: "trees-finance",
			Name: "trees.finance",
		},
	},
	"ELT": {
		{
			ID:   7884,
			Slug: "elite-swap",
			Name: "Elite Swap",
		},
	},
	"UPDOG": {
		{
			ID:   9877,
			Slug: "updog",
			Name: "UPDOG",
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
	"ELAMA": {
		{
			ID:   3734,
			Slug: "elamachain",
			Name: "Elamachain",
		},
	},
	"EXE": {
		{
			ID:   5620,
			Slug: "8x8-protocol",
			Name: "8X8 PROTOCOL",
		},
	},
	"NVT": {
		{
			ID:   5906,
			Slug: "nervenetwork",
			Name: "NerveNetwork",
		},
	},
	"PHNX": {
		{
			ID:   5674,
			Slug: "phoenixdao",
			Name: "PhoenixDAO",
		},
	},
	"AK12": {
		{
			ID:   5189,
			Slug: "ak12",
			Name: "AK12",
		},
	},
	"DSC": {
		{
			ID:   4836,
			Slug: "dash-cash",
			Name: "Dash Cash",
		},
	},
	"SDEFI": {
		{
			ID:   5862,
			Slug: "sdefi",
			Name: "sDEFI",
		},
	},
	"NILU": {
		{
			ID:   6032,
			Slug: "nilu",
			Name: "Nilu",
		},
	},
	"UTK": {
		{
			ID:   2320,
			Slug: "utrust",
			Name: "Utrust",
		},
	},
	"UBXT": {
		{
			ID:   7040,
			Slug: "upbots",
			Name: "UpBots",
		},
	},
	"BTAD": {
		{
			ID:   3294,
			Slug: "bitcoin-adult",
			Name: "Bitcoin Adult",
		},
	},
	"TSHP": {
		{
			ID:   4280,
			Slug: "12ships",
			Name: "12Ships",
		},
	},
	"RYIU": {
		{
			ID:   8871,
			Slug: "ryi-unity",
			Name: "RYI Unity",
		},
	},
	"TRAC": {
		{
			ID:   2467,
			Slug: "origintrail",
			Name: "OriginTrail",
		},
	},
	"AAVEDOWN": {
		{
			ID:   7775,
			Slug: "aave-down",
			Name: "AAVEDOWN",
		},
	},
	"UTED": {
		{
			ID:   7462,
			Slug: "united",
			Name: "United",
		},
	},
	"CPX": {
		{
			ID:   9945,
			Slug: "center-prime-token",
			Name: "CenterPrime",
		},
	},
	"mTWTR": {
		{
			ID:   8018,
			Slug: "mirrored-twitter",
			Name: "Mirrored Twitter",
		},
	},
	"DIT": {
		{
			ID:   3264,
			Slug: "digital-insurance-token",
			Name: "Digital Insurance Token",
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
	"HUSL": {
		{
			ID:   5253,
			Slug: "the-hustle-app",
			Name: "The Hustle App",
		},
	},
	"RARI": {
		{
			ID:   5877,
			Slug: "rarible",
			Name: "Rarible",
		},
	},
	"LEDU": {
		{
			ID:   2562,
			Slug: "education-ecosystem",
			Name: "Education Ecosystem",
		},
	},
	"FENIX": {
		{
			ID:   10124,
			Slug: "fenix-finance",
			Name: "Fenix Finance",
		},
	},
	"BTCST": {
		{
			ID:   8891,
			Slug: "btc-standard-hashrate-token",
			Name: "Bitcoin Standard Hashrate Token",
		},
	},
	"QARK": {
		{
			ID:   5858,
			Slug: "qanplatform",
			Name: "QANplatform",
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
	"WSTA": {
		{
			ID:   9620,
			Slug: "wrapped-statera",
			Name: "Wrapped Statera",
		},
	},
	"SOLDIER": {
		{
			ID:   10067,
			Slug: "space-soldier",
			Name: "Space Soldier",
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
	"PDEX": {
		{
			ID:   9017,
			Slug: "polkadex",
			Name: "Polkadex",
		},
	},
	"CLOAK": {
		{
			ID:   362,
			Slug: "cloakcoin",
			Name: "CloakCoin",
		},
	},
	"ALN": {
		{
			ID:   5544,
			Slug: "aluna-social",
			Name: "Aluna.Social",
		},
	},
	"BTD": {
		{
			ID:   8204,
			Slug: "bat-true-dollar",
			Name: "Bolt Dollar",
		},
	},
	"POFI": {
		{
			ID:   9276,
			Slug: "pofi",
			Name: "Pofi",
		},
	},
	"DPC": {
		{
			ID:   8093,
			Slug: "dappcents",
			Name: "DAPPCENTS",
		},
	},
	"HOMI": {
		{
			ID:   5572,
			Slug: "homihelp",
			Name: "HOMIHELP",
		},
	},
	"BSCV": {
		{
			ID:   8843,
			Slug: "bscview",
			Name: "BSCView",
		},
	},
	"FUSII": {
		{
			ID:   8703,
			Slug: "fusible",
			Name: "Fusible",
		},
	},
	"LF": {
		{
			ID:   9160,
			Slug: "linkflow-finance",
			Name: "Linkflow Finance",
		},
	},
	"PEPPA": {
		{
			ID:   10432,
			Slug: "peppa-network",
			Name: "Peppa Network",
		},
	},
	"MANY": {
		{
			ID:   9124,
			Slug: "manyswap",
			Name: "Manyswap",
		},
	},
	"AXN": {
		{
			ID:   7676,
			Slug: "axion",
			Name: "Axion",
		},
	},
	"XKI": {
		{
			ID:   9908,
			Slug: "ki-foundation",
			Name: "Ki",
		},
	},
	"WOO": {
		{
			ID:   7501,
			Slug: "wootrade",
			Name: "Wootrade",
		},
	},
	"PHA": {
		{
			ID:   6841,
			Slug: "phala-network",
			Name: "Phala Network",
		},
	},
	"TBX": {
		{
			ID:   2452,
			Slug: "tokenbox",
			Name: "Tokenbox",
		},
	},
	"PEX": {
		{
			ID:   1209,
			Slug: "posex",
			Name: "PosEx",
		},
	},
	"SCS": {
		{
			ID:   1651,
			Slug: "speedcash",
			Name: "SpeedCash",
		},
	},
	"DFYN": {
		{
			ID:   9511,
			Slug: "dfyn-network",
			Name: "Dfyn Network",
		},
	},
	"VSP": {
		{
			ID:   8492,
			Slug: "vesper",
			Name: "Vesper",
		},
	},
	"XSN": {
		{
			ID:   2633,
			Slug: "stakenet",
			Name: "Stakenet",
		},
	},
	"CWS": {
		{
			ID:   8365,
			Slug: "crowns",
			Name: "Crowns",
		},
	},
	"DEFI5": {
		{
			ID:   8430,
			Slug: "defi-top-5-tokens-index",
			Name: "DEFI Top 5 Tokens Index",
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
	"MET": {
		{
			ID:   2873,
			Slug: "metronome",
			Name: "Metronome",
		},
	},
	"MGO": {
		{
			ID:   1715,
			Slug: "mobilego",
			Name: "MobileGo",
		},
	},
	"CMERGE": {
		{
			ID:   11016,
			Slug: "coinmerge",
			Name: "CoinMerge",
		},
	},
	"MELLO": {
		{
			ID:   9339,
			Slug: "mello-token",
			Name: "Mello Token",
		},
	},
	"X-TOKEN": {
		{
			ID:   9927,
			Slug: "x-token",
			Name: "X-Token",
		},
	},
	"XYO": {
		{
			ID:   2765,
			Slug: "xyo",
			Name: "XYO",
		},
	},
	"TRST": {
		{
			ID:   1638,
			Slug: "trust",
			Name: "WeTrust",
		},
	},
	"MTS": {
		{
			ID:   7778,
			Slug: "metis",
			Name: "Metis",
		},
	},
	"AGOLP": {
		{
			ID:   9407,
			Slug: "algoil",
			Name: "AlgOil",
		},
	},
	"VSYS": {
		{
			ID:   3704,
			Slug: "v-systems",
			Name: "v.systems",
		},
	},
	"NAS": {
		{
			ID:   1908,
			Slug: "nebulas-token",
			Name: "Nebulas",
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
	"2GIVE": {
		{
			ID:   1252,
			Slug: "2give",
			Name: "2GIVE",
		},
	},
	"ESK": {
		{
			ID:   6435,
			Slug: "eska",
			Name: "Eska",
		},
	},
	"PLUTO": {
		{
			ID:   9715,
			Slug: "plutopepe",
			Name: "PlutoPepe",
		},
	},
	"NRV": {
		{
			ID:   8755,
			Slug: "nerve-finance",
			Name: "Nerve Finance",
		},
	},
	"mIAU": {
		{
			ID:   8024,
			Slug: "mirrored-ishares-gold-trust",
			Name: "Mirrored iShares Gold Trust",
		},
	},
	"YFFS": {
		{
			ID:   6863,
			Slug: "yffs",
			Name: "YFFS Finance",
		},
	},
	"XRE": {
		{
			ID:   10154,
			Slug: "xre-global",
			Name: "XRE Global",
		},
	},
	"BPOP": {
		{
			ID:   6388,
			Slug: "bpop",
			Name: "BPOP",
		},
	},
	"XMM": {
		{
			ID:   7042,
			Slug: "momentum",
			Name: "Momentum",
		},
	},
	"RAGNA": {
		{
			ID:   3374,
			Slug: "ragnarok",
			Name: "Ragnarok",
		},
	},
	"WEMIX": {
		{
			ID:   7548,
			Slug: "wemix",
			Name: "WEMIX",
		},
	},
	"PWAY": {
		{
			ID:   8838,
			Slug: "pway",
			Name: "PWAY",
		},
	},
	"AMLT": {
		{
			ID:   2607,
			Slug: "amlt",
			Name: "AMLT",
		},
	},
	"ADD": {
		{
			ID:   5834,
			Slug: "add-xyz",
			Name: "Add.xyz",
		},
	},
	"BWX": {
		{
			ID:   2953,
			Slug: "blue-whale-exchange",
			Name: "Blue Whale EXchange",
		},
	},
	"TRUMP": {
		{
			ID:   1185,
			Slug: "trumpcoin",
			Name: "TrumpCoin",
		},
	},
	"JDC": {
		{
			ID:   4929,
			Slug: "jd-coin",
			Name: "JD Coin",
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
	"PAXG": {
		{
			ID:   4705,
			Slug: "pax-gold",
			Name: "PAX Gold",
		},
	},
	"ZCR": {
		{
			ID:   3158,
			Slug: "zcore",
			Name: "ZCore",
		},
	},
	"MBY": {
		{
			ID:   10583,
			Slug: "monkey-token",
			Name: "Monkey Token",
		},
	},
	"RMOON": {
		{
			ID:   9274,
			Slug: "rocketmoon",
			Name: "RocketMoon",
		},
	},
	"ARK": {
		{
			ID:   1586,
			Slug: "ark",
			Name: "Ark",
		},
	},
	"QASH": {
		{
			ID:   2213,
			Slug: "qash",
			Name: "QASH",
		},
	},
	"INFS": {
		{
			ID:   4903,
			Slug: "infinity-esaham",
			Name: "Infinity Esaham",
		},
	},
	"BOUTS": {
		{
			ID:   2717,
			Slug: "boutspro",
			Name: "BoutsPro",
		},
	},
	"INRT": {
		{
			ID:   5416,
			Slug: "inrtoken",
			Name: "INRToken",
		},
	},
	"KEANU": {
		{
			ID:   10196,
			Slug: "keanu-inu",
			Name: "Keanu Inu",
		},
	},
	"UNN": {
		{
			ID:   8056,
			Slug: "union-protocol-governance-token",
			Name: "UNION Protocol Governance Token",
		},
	},
	"WDR": {
		{
			ID:   9636,
			Slug: "widercoin",
			Name: "Widercoin",
		},
	},
	"RLY": {
		{
			ID:   8075,
			Slug: "rally",
			Name: "Rally",
		},
	},
	"XLR": {
		{
			ID:   1606,
			Slug: "solaris",
			Name: "Solaris",
		},
	},
	"XBE": {
		{
			ID:   9092,
			Slug: "xbe-token",
			Name: "XBE Token",
		},
	},
	"SAFEEARTH": {
		{
			ID:   9081,
			Slug: "safeearth",
			Name: "SafeEarth",
		},
	},
	"PURE": {
		{
			ID:   7203,
			Slug: "puriever",
			Name: "Puriever",
		},
	},
	"WINATESLA": {
		{
			ID:   10934,
			Slug: "win-a-tesla",
			Name: "WIN A TESLA",
		},
	},
	"MNDAO": {
		{
			ID:   9006,
			Slug: "moondao",
			Name: "MoonDAO",
		},
	},
	"EFAR": {
		{
			ID:   7361,
			Slug: "fridn",
			Name: "Fridn",
		},
	},
	"UTP": {
		{
			ID:   10286,
			Slug: "utopian-protocol",
			Name: "Utopian Protocol",
		},
	},
	"GLS": {
		{
			ID:   4834,
			Slug: "golos-blockchain",
			Name: "Golos Blockchain",
		},
	},
	"ALI": {
		{
			ID:   3248,
			Slug: "ailink-token",
			Name: "AiLink Token",
		},
	},
	"GPKR": {
		{
			ID:   3184,
			Slug: "gold-poker",
			Name: "Gold Poker",
		},
	},
	"XDEX": {
		{
			ID:   9087,
			Slug: "xdefi",
			Name: "xDeFi",
		},
	},
	"BAR": {
		{
			ID:   5225,
			Slug: "fc-barcelona-fan-token",
			Name: "FC Barcelona Fan Token",
		},
	},
	"LGCY": {
		{
			ID:   6665,
			Slug: "lgcy-network",
			Name: "LGCY Network",
		},
	},
	"BNK": {
		{
			ID:   2842,
			Slug: "bankera",
			Name: "Bankera",
		},
	},
	"KANGAL": {
		{
			ID:   8543,
			Slug: "kangal",
			Name: "Kangal",
		},
	},
	"XBP": {
		{
			ID:   2614,
			Slug: "blitzpredict",
			Name: "BlitzPick",
		},
	},
	"MTR": {
		{
			ID:   6627,
			Slug: "meter-stable",
			Name: "Meter Stable",
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
	"DCB": {
		{
			ID:   10563,
			Slug: "decubate",
			Name: "Decubate",
		},
	},
	"BCHBEAR": {
		{
			ID:   5466,
			Slug: "3x-short-bitcoin-cash-token",
			Name: "3x Short Bitcoin Cash Token",
		},
	},
	"LNT": {
		{
			ID:   7036,
			Slug: "lottonation",
			Name: "Lottonation",
		},
	},
	"AXL": {
		{
			ID:   5065,
			Slug: "axial-entertainment-digital-asset",
			Name: "Axial Entertainment Digital Asset",
		},
	},
	"HTR": {
		{
			ID:   5552,
			Slug: "hathor",
			Name: "Hathor",
		},
	},
	"FLG": {
		{
			ID:   4842,
			Slug: "folgory-coin",
			Name: "Folgory Coin",
		},
	},
	"VOT": {
		{
			ID:   2221,
			Slug: "votecoin",
			Name: "VoteCoin",
		},
	},
	"VIVID": {
		{
			ID:   3008,
			Slug: "vivid-coin",
			Name: "Vivid Coin",
		},
	},
	"SBE": {
		{
			ID:   4538,
			Slug: "sombe",
			Name: "Sombe",
		},
	},
	"QSP": {
		{
			ID:   2212,
			Slug: "quantstamp",
			Name: "Quantstamp",
		},
	},
	"BLY": {
		{
			ID:   6283,
			Slug: "blocery",
			Name: "Blocery",
		},
	},
	"NPX": {
		{
			ID:   2602,
			Slug: "napoleonx",
			Name: "NaPoleonX",
		},
	},
	"SHIP": {
		{
			ID:   2579,
			Slug: "shipchain",
			Name: "ShipChain",
		},
	},
	"TRCL": {
		{
			ID:   5800,
			Slug: "treecle",
			Name: "Treecle",
		},
	},
	"LEMO": {
		{
			ID:   2950,
			Slug: "lemochain",
			Name: "LemoChain",
		},
	},
	"ZPAY": {
		{
			ID:   10929,
			Slug: "zoidpay",
			Name: "ZoidPay",
		},
	},
	"AET": {
		{
			ID:   4920,
			Slug: "aerotoken",
			Name: "Aerotoken",
		},
	},
	"CIPX": {
		{
			ID:   4708,
			Slug: "colletrix",
			Name: "Colletrix",
		},
	},
	"LYXe": {
		{
			ID:   5625,
			Slug: "lukso",
			Name: "LUKSO",
		},
	},
	"vSXP": {
		{
			ID:   7952,
			Slug: "vsxp",
			Name: "Venus SXP",
		},
	},
	"TH": {
		{
			ID:   7636,
			Slug: "team-heretics-fan-token",
			Name: "Team Heretics Fan Token",
		},
	},
	"PAK": {
		{
			ID:   1066,
			Slug: "pakcoin",
			Name: "Pakcoin",
		},
	},
	"MVC": {
		{
			ID:   7703,
			Slug: "mileverse",
			Name: "MileVerse",
		},
	},
	"ALGOBULL": {
		{
			ID:   6074,
			Slug: "3x-long-algorand-token",
			Name: "3X Long Algorand Token",
		},
	},
	"DAVP": {
		{
			ID:   4999,
			Slug: "davion",
			Name: "Davion",
		},
	},
	"SRK": {
		{
			ID:   3935,
			Slug: "sparkpoint",
			Name: "SparkPoint",
		},
	},
	"VDG": {
		{
			ID:   3205,
			Slug: "veridocglobal",
			Name: "VeriDocGlobal",
		},
	},
	"LTCU": {
		{
			ID:   1935,
			Slug: "litecoin-ultra",
			Name: "LiteCoin Ultra",
		},
	},
	"POC": {
		{
			ID:   8319,
			Slug: "poc-blockchain",
			Name: "POC Blockchain",
		},
	},
	"NFTBOX": {
		{
			ID:   10063,
			Slug: "nftbox-fun",
			Name: "NFTBOX.fun",
		},
	},
	"SHROOM": {
		{
			ID:   6891,
			Slug: "niftyx-protocol",
			Name: "Niftyx Protocol",
		},
	},
	"CHSB": {
		{
			ID:   2499,
			Slug: "swissborg",
			Name: "SwissBorg",
		},
	},
	"AGA": {
		{
			ID:   6404,
			Slug: "aga",
			Name: "AGA Token",
		},
	},
	"FUD": {
		{
			ID:   7335,
			Slug: "fudfinance",
			Name: "FUD.finance",
		},
	},
	"RPD": {
		{
			ID:   3432,
			Slug: "rapids",
			Name: "Rapids",
		},
	},
	"DINGO": {
		{
			ID:   9982,
			Slug: "dingo-token",
			Name: "DINGO TOKEN",
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
	"DOWS": {
		{
			ID:   8643,
			Slug: "shadows",
			Name: "Shadows",
		},
	},
	"FYD": {
		{
			ID:   4680,
			Slug: "fydcoin",
			Name: "FYDcoin",
		},
	},
	"JPYC": {
		{
			ID:   9045,
			Slug: "jpycoin",
			Name: "JPYC",
		},
	},
	"USDT": {
		{
			ID:   825,
			Slug: "tether",
			Name: "Tether",
		},
	},
	"SKL": {
		{
			ID:   5691,
			Slug: "skale-network",
			Name: "SKALE Network",
		},
	},
	"HBTC": {
		{
			ID:   6941,
			Slug: "huobi-btc",
			Name: "Huobi BTC",
		},
	},
	"SPC": {
		{
			ID:   2410,
			Slug: "spacechain",
			Name: "SpaceChain",
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
	"TELOS": {
		{
			ID:   3482,
			Slug: "teloscoin",
			Name: "Teloscoin",
		},
	},
	"KXC": {
		{
			ID:   3198,
			Slug: "kingxchain",
			Name: "KingXChain",
		},
	},
	"CIPHC": {
		{
			ID:   5844,
			Slug: "cipher-core-token",
			Name: "Cipher Core Token",
		},
	},
	"ATMOS": {
		{
			ID:   1624,
			Slug: "atmos",
			Name: "Atmos",
		},
	},
	"MANNA": {
		{
			ID:   1019,
			Slug: "manna",
			Name: "Manna",
		},
	},
	"TMCN": {
		{
			ID:   7738,
			Slug: "timecoinprotocol",
			Name: "TimeCoinProtocol",
		},
	},
	"VECT": {
		{
			ID:   5686,
			Slug: "vectorium",
			Name: "Vectorium",
		},
	},
	"PIE": {
		{
			ID:   6243,
			Slug: "defipie",
			Name: "DeFiPie",
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
	"ABYSS": {
		{
			ID:   2847,
			Slug: "abyss",
			Name: "Abyss",
		},
	},
	"LBA": {
		{
			ID:   2760,
			Slug: "libra-credit",
			Name: "Cred",
		},
	},
	"NCP": {
		{
			ID:   2984,
			Slug: "newton-coin-project",
			Name: "Newton Coin Project",
		},
	},
	"RIFT": {
		{
			ID:   5110,
			Slug: "rift-token",
			Name: "RIFT Token",
		},
	},
	"CROSS": {
		{
			ID:   9442,
			Slug: "crosspad",
			Name: "CrossPad",
		},
	},
	"MOTO": {
		{
			ID:   360,
			Slug: "motocoin",
			Name: "Motocoin",
		},
	},
	"MPS": {
		{
			ID:   5540,
			Slug: "mt-pelerin",
			Name: "Mt Pelerin",
		},
	},
	"mTSLA": {
		{
			ID:   8004,
			Slug: "mirrored-tesla",
			Name: "Mirrored Tesla",
		},
	},
	"SONG": {
		{
			ID:   857,
			Slug: "songcoin",
			Name: "SongCoin",
		},
	},
	"INU": {
		{
			ID:   10900,
			Slug: "hachikoinu",
			Name: "Hachiko Inu",
		},
	},
	"TIP": {
		{
			ID:   8332,
			Slug: "technology-innovation-project",
			Name: "TECHNOLOGY INNOVATION PROJECT",
		},
	},
	"DC": {
		{
			ID:   8511,
			Slug: "deep-coin",
			Name: "DeepCoin",
		},
	},
	"CNTR": {
		{
			ID:   7349,
			Slug: "centaur",
			Name: "Centaur",
		},
	},
	"TZC": {
		{
			ID:   2002,
			Slug: "trezarcoin",
			Name: "TrezarCoin",
		},
	},
	"SATX": {
		{
			ID:   5430,
			Slug: "satoexchange-token",
			Name: "SatoExchange Token",
		},
	},
	"DOGEMOON": {
		{
			ID:   9618,
			Slug: "dogemoon",
			Name: "DogeMoon",
		},
	},
	"CATZ": {
		{
			ID:   9692,
			Slug: "catzcoin",
			Name: "CatzCoin",
		},
	},
	"CDAI": {
		{
			ID:   5263,
			Slug: "compound-dai",
			Name: "Compound Dai",
		},
	},
	"WIZARD": {
		{
			ID:   10960,
			Slug: "wizard",
			Name: "WIZARD",
		},
	},
	"VNDC": {
		{
			ID:   4805,
			Slug: "vndc",
			Name: "VNDC",
		},
	},
	"EXOR": {
		{
			ID:   4126,
			Slug: "exor",
			Name: "EXOR",
		},
	},
	"YSEC": {
		{
			ID:   7572,
			Slug: "yearn-secure",
			Name: "Yearn Secure",
		},
	},
	"SSGT": {
		{
			ID:   10331,
			Slug: "safeswap-governance-token",
			Name: "Safeswap Governance Token",
		},
	},
	"AGVE": {
		{
			ID:   9453,
			Slug: "agave",
			Name: "Agave",
		},
	},
	"OKB": {
		{
			ID:   3897,
			Slug: "okb",
			Name: "OKB",
		},
	},
	"DREP": {
		{
			ID:   9148,
			Slug: "drep-new",
			Name: "Drep [new]",
		},
	},
	"MUSE": {
		{
			ID:   7805,
			Slug: "muse",
			Name: "Muse",
		},
	},
	"HTRE": {
		{
			ID:   7100,
			Slug: "hodltree",
			Name: "HodlTree",
		},
	},
	"AG8": {
		{
			ID:   5536,
			Slug: "atromg8",
			Name: "AtromG8",
		},
	},
	"WGRT": {
		{
			ID:   5630,
			Slug: "waykichain-governance-coin",
			Name: "WaykiChain Governance Coin",
		},
	},
	"RAM": {
		{
			ID:   8798,
			Slug: "ramifi-protocol",
			Name: "Ramifi Protocol",
		},
	},
	"VYBE": {
		{
			ID:   7025,
			Slug: "vybe",
			Name: "Vybe",
		},
	},
	"WTC": {
		{
			ID:   1925,
			Slug: "waltonchain",
			Name: "Waltonchain",
		},
	},
	"QWC": {
		{
			ID:   3942,
			Slug: "qwertycoin",
			Name: "Qwertycoin",
		},
	},
	"GPYX": {
		{
			ID:   3349,
			Slug: "goldenpyrex",
			Name: "GoldenPyrex",
		},
	},
	"KKC": {
		{
			ID:   3421,
			Slug: "kabberry-coin",
			Name: "Kabberry Coin",
		},
	},
	"BCVT": {
		{
			ID:   8782,
			Slug: "bitcoinvend",
			Name: "BitcoinVend",
		},
	},
	"FFA": {
		{
			ID:   9659,
			Slug: "cryptofifa",
			Name: "Cryptofifa",
		},
	},
	"SHIKO": {
		{
			ID:   10510,
			Slug: "shikoku-inu",
			Name: "Shikoku Inu",
		},
	},
	"YUI": {
		{
			ID:   7709,
			Slug: "yui-token",
			Name: "YUI Token",
		},
	},
	"MOLK": {
		{
			ID:   3304,
			Slug: "mobilinktoken",
			Name: "MobilinkToken",
		},
	},
	"ALIAS": {
		{
			ID:   1505,
			Slug: "alias",
			Name: "Alias",
		},
	},
	"KSP": {
		{
			ID:   8296,
			Slug: "klayswap-protocol",
			Name: "KLAYswap Protocol",
		},
	},
	"HEGIC": {
		{
			ID:   6929,
			Slug: "hegic",
			Name: "Hegic",
		},
	},
	"FXF": {
		{
			ID:   8416,
			Slug: "finxflo",
			Name: "Finxflo",
		},
	},
	"GSX": {
		{
			ID:   10093,
			Slug: "gold-secured-currency",
			Name: "Gold Secured Currency",
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
	"MOONTOKEN": {
		{
			ID:   9399,
			Slug: "moon-token",
			Name: "MoonToken",
		},
	},
	"CAKECRYPT": {
		{
			ID:   10066,
			Slug: "cakecrypt",
			Name: "CAKECRYPT",
		},
	},
	"ETCH": {
		{
			ID:   9742,
			Slug: "elontech",
			Name: "ElonTech",
		},
	},
	"GRS": {
		{
			ID:   258,
			Slug: "groestlcoin",
			Name: "Groestlcoin",
		},
	},
	"FACE": {
		{
			ID:   2775,
			Slug: "faceter",
			Name: "Faceter",
		},
	},
	"JRT": {
		{
			ID:   5187,
			Slug: "jarvis-network",
			Name: "Jarvis Network",
		},
	},
	"BBO": {
		{
			ID:   2836,
			Slug: "bigbom",
			Name: "Bigbom",
		},
	},
	"FMG": {
		{
			ID:   8902,
			Slug: "fm-gallery",
			Name: "FM Gallery",
		},
	},
	"CBAT": {
		{
			ID:   5742,
			Slug: "compound-basic-attention-token",
			Name: "Compound Basic Attention Token",
		},
	},
	"ECOREAL": {
		{
			ID:   3344,
			Slug: "ecoreal-estate",
			Name: "Ecoreal Estate",
		},
	},
	"FED": {
		{
			ID:   4555,
			Slug: "fedora-gold",
			Name: "Fedora Gold",
		},
	},
	"YBO": {
		{
			ID:   8341,
			Slug: "young-boys-fan-token",
			Name: "Young Boys Fan Token",
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
	"SEFA": {
		{
			ID:   5425,
			Slug: "mesefa",
			Name: "MESEFA",
		},
	},
	"MAZ": {
		{
			ID:   6711,
			Slug: "mazzuma",
			Name: "Mazzuma",
		},
	},
	"NADA": {
		{
			ID:   10623,
			Slug: "nothing",
			Name: "Nothing",
		},
	},
	"EGLD": {
		{
			ID:   6892,
			Slug: "elrond-egld",
			Name: "Elrond",
		},
	},
	"XCN": {
		{
			ID:   501,
			Slug: "cryptonite",
			Name: "Cryptonite",
		},
	},
	"NRVE": {
		{
			ID:   2956,
			Slug: "narrative",
			Name: "Narrative",
		},
	},
	"HTB": {
		{
			ID:   4506,
			Slug: "hotbit-token",
			Name: "Hotbit Token",
		},
	},
	"DJ15": {
		{
			ID:   8011,
			Slug: "davincij15-token",
			Name: "Davincij15 Token",
		},
	},
	"XYM": {
		{
			ID:   8677,
			Slug: "symbol",
			Name: "Symbol",
		},
	},
	"EXRD": {
		{
			ID:   7692,
			Slug: "radix",
			Name: "Radix",
		},
	},
	"BTSE": {
		{
			ID:   5305,
			Slug: "btse",
			Name: "BTSE",
		},
	},
	"STARS": {
		{
			ID:   8996,
			Slug: "mogul-productions",
			Name: "Mogul Productions",
		},
	},
	"POOCOIN": {
		{
			ID:   8978,
			Slug: "poocoin",
			Name: "PooCoin",
		},
	},
	"HONOR": {
		{
			ID:   10620,
			Slug: "farm-hero",
			Name: "FarmHero",
		},
	},
	"KBC": {
		{
			ID:   2907,
			Slug: "karatgold-coin",
			Name: "Karatgold Coin",
		},
	},
	"KAWAII": {
		{
			ID:   10193,
			Slug: "kawai-inu",
			Name: "Kawai INU",
		},
	},
	"OMNI": {
		{
			ID:   83,
			Slug: "omni",
			Name: "Omni",
		},
	},
	"YFO": {
		{
			ID:   8398,
			Slug: "yfione",
			Name: "YFIONE",
		},
	},
	"ARTEX": {
		{
			ID:   11029,
			Slug: "artex",
			Name: "Artex",
		},
	},
	"FL": {
		{
			ID:   8436,
			Slug: "freeliquid",
			Name: "Freeliquid",
		},
	},
	"DXT": {
		{
			ID:   2510,
			Slug: "datawallet",
			Name: "Datawallet",
		},
	},
	"SIL": {
		{
			ID:   10136,
			Slug: "sil-finance",
			Name: "SIL.FINANCE",
		},
	},
	"XXT": {
		{
			ID:   10965,
			Slug: "xxt-token",
			Name: "XXT-Token",
		},
	},
	"HY": {
		{
			ID:   6261,
			Slug: "hybrix",
			Name: "hybrix",
		},
	},
	"PYRK": {
		{
			ID:   5591,
			Slug: "pyrk",
			Name: "Pyrk",
		},
	},
	"TREX": {
		{
			ID:   5455,
			Slug: "trexcoin",
			Name: "Trexcoin",
		},
	},
	"KRW": {
		{
			ID:   11024,
			Slug: "kingdefi",
			Name: "KingDeFi",
		},
	},
	"DASS": {
		{
			ID:   10801,
			Slug: "dashsports",
			Name: "DashSports",
		},
	},
	"DRGN": {
		{
			ID:   2243,
			Slug: "dragonchain",
			Name: "Dragonchain",
		},
	},
	"IDV": {
		{
			ID:   8726,
			Slug: "idavoll-network",
			Name: "Idavoll Network",
		},
	},
	"WAIF": {
		{
			ID:   6552,
			Slug: "waifu-token",
			Name: "Waifu Token",
		},
	},
	"LOCK": {
		{
			ID:   6566,
			Slug: "meridian-network",
			Name: "Meridian Network",
		},
	},
	"MOX": {
		{
			ID:   3688,
			Slug: "mox",
			Name: "MoX",
		},
	},
	"1TRC": {
		{
			ID:   10955,
			Slug: "1tronic-network",
			Name: "1TRONIC Network",
		},
	},
	"XFYI": {
		{
			ID:   7323,
			Slug: "xcredit",
			Name: "XCredit",
		},
	},
	"PAPER": {
		{
			ID:   8847,
			Slug: "fomo-app",
			Name: "Fomo App",
		},
	},
	"JAR": {
		{
			ID:   4034,
			Slug: "jarvis",
			Name: "Jarvis+",
		},
	},
	"MAR": {
		{
			ID:   4510,
			Slug: "mchain",
			Name: "Mchain",
		},
	},
	"CEL": {
		{
			ID:   2700,
			Slug: "celsius",
			Name: "Celsius",
		},
	},
	"FVT": {
		{
			ID:   8149,
			Slug: "finance-vote",
			Name: "Finance.Vote",
		},
	},
	"LQT": {
		{
			ID:   10594,
			Slug: "liquidifty",
			Name: "Liquidifty",
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
	"NFD": {
		{
			ID:   9496,
			Slug: "nifdo-protocol",
			Name: "NIFDO Protocol",
		},
	},
	"CBSN": {
		{
			ID:   9418,
			Slug: "blockswap-network",
			Name: "BlockSwap Network",
		},
	},
	"HIZ": {
		{
			ID:   7554,
			Slug: "hiz-finance",
			Name: "Hiz Finance",
		},
	},
	"SATT": {
		{
			ID:   7244,
			Slug: "satt",
			Name: "SaTT",
		},
	},
	"HER": {
		{
			ID:   2754,
			Slug: "heronode",
			Name: "HeroNode",
		},
	},
	"GERO": {
		{
			ID:   9904,
			Slug: "gerowallet",
			Name: "GeroWallet",
		},
	},
	"MOMMYDOGE": {
		{
			ID:   10876,
			Slug: "mommy-doge-coin",
			Name: "Mommy Doge Coin",
		},
	},
	"YFOX": {
		{
			ID:   7286,
			Slug: "yfox-finance",
			Name: "YFOX FINANCE",
		},
	},
	"POLARV3": {
		{
			ID:   9019,
			Slug: "polar",
			Name: "Polar",
		},
	},
	"SOLO": {
		{
			ID:   5279,
			Slug: "sologenic",
			Name: "Sologenic",
		},
	},
	"BXY": {
		{
			ID:   4646,
			Slug: "beaxy",
			Name: "Beaxy",
		},
	},
	"BNS": {
		{
			ID:   5989,
			Slug: "bns-token",
			Name: "BNS Token",
		},
	},
	"WEBN": {
		{
			ID:   3744,
			Slug: "webn-token",
			Name: "WEBN token",
		},
	},
	"CET": {
		{
			ID:   2941,
			Slug: "coinex-token",
			Name: "CoinEx Token",
		},
	},
	"BAC": {
		{
			ID:   7813,
			Slug: "basis-cash",
			Name: "Basis Cash",
		},
	},
	"BOLI": {
		{
			ID:   1053,
			Slug: "bolivarcoin",
			Name: "Bolivarcoin",
		},
	},
	"XRPUP": {
		{
			ID:   7001,
			Slug: "xrpup",
			Name: "XRPUP",
		},
	},
	"XUC": {
		{
			ID:   2091,
			Slug: "exchange-union",
			Name: "Exchange Union",
		},
	},
	"HOD": {
		{
			ID:   10412,
			Slug: "hodooi",
			Name: "HoDooi",
		},
	},
	"PKR": {
		{
			ID:   10427,
			Slug: "polker",
			Name: "Polker",
		},
	},
	"SHE": {
		{
			ID:   3252,
			Slug: "shinechain",
			Name: "ShineChain",
		},
	},
	"DIRTY": {
		{
			ID:   10328,
			Slug: "dirty-finance",
			Name: "Dirty Finance",
		},
	},
	"HOMT": {
		{
			ID:   5722,
			Slug: "homt",
			Name: "HOMT",
		},
	},
	"MLK": {
		{
			ID:   5266,
			Slug: "milk-alliance",
			Name: "MiL.k",
		},
	},
	"MUSH": {
		{
			ID:   8502,
			Slug: "mushroom",
			Name: "Mushroom",
		},
	},
	"RTH": {
		{
			ID:   3279,
			Slug: "rotharium",
			Name: "Rotharium",
		},
	},
	"FUZE": {
		{
			ID:   4104,
			Slug: "fuze-token",
			Name: "FUZE Token",
		},
	},
	"BABYBNB": {
		{
			ID:   10698,
			Slug: "babybnb",
			Name: "Babybnb",
		},
	},
	"LKM": {
		{
			ID:   9410,
			Slug: "lokum-finance",
			Name: "Lokum Finance",
		},
	},
	"DEV": {
		{
			ID:   5990,
			Slug: "dev-protocol",
			Name: "Dev Protocol",
		},
	},
	"UPX": {
		{
			ID:   3752,
			Slug: "uplexa",
			Name: "uPlexa",
		},
	},
	"wSIENNA": {
		{
			ID:   9449,
			Slug: "sienna-erc20",
			Name: "Sienna (ERC20)",
		},
	},
	"OLT": {
		{
			ID:   2921,
			Slug: "oneledger",
			Name: "OneLedger",
		},
	},
	"DEOR": {
		{
			ID:   8680,
			Slug: "deor",
			Name: "DEOR",
		},
	},
	"FAST": {
		{
			ID:   8087,
			Slug: "fastswap",
			Name: "FastSwap",
		},
	},
	"eMTRG": {
		{
			ID:   6628,
			Slug: "meter-governance-mapped-by-meter-io",
			Name: "Meter Governance mapped by Meter.io",
		},
	},
	"PTE": {
		{
			ID:   10478,
			Slug: "peet-defi-new",
			Name: "Peet DeFi [new]",
		},
	},
	"EMRX": {
		{
			ID:   4490,
			Slug: "emirex-token",
			Name: "Emirex Token",
		},
	},
	"ZYN": {
		{
			ID:   4951,
			Slug: "zynecoin",
			Name: "Zynecoin",
		},
	},
	"WAGE": {
		{
			ID:   3416,
			Slug: "digiwage",
			Name: "Digiwage",
		},
	},
	"XBTC21": {
		{
			ID:   1248,
			Slug: "bitcoin-21",
			Name: "Bitcoin 21",
		},
	},
	"JT": {
		{
			ID:   6262,
			Slug: "jubi-token",
			Name: "Jubi Token",
		},
	},
	"BDAY": {
		{
			ID:   8462,
			Slug: "birthday-cake",
			Name: "Birthday Cake",
		},
	},
	"YFIA": {
		{
			ID:   7442,
			Slug: "yfia",
			Name: "YFIA",
		},
	},
	"KIN": {
		{
			ID:   1993,
			Slug: "kin",
			Name: "Kin",
		},
	},
	"MWC": {
		{
			ID:   5031,
			Slug: "mimblewimblecoin",
			Name: "MimbleWimbleCoin",
		},
	},
	"VPP": {
		{
			ID:   9756,
			Slug: "virtue-poker",
			Name: "Virtue Poker",
		},
	},
	"N8V": {
		{
			ID:   584,
			Slug: "native-coin",
			Name: "NativeCoin",
		},
	},
	"KTON": {
		{
			ID:   5931,
			Slug: "darwinia-commitment-token",
			Name: "Darwinia Commitment Token",
		},
	},
	"DAPPT": {
		{
			ID:   4176,
			Slug: "dapp-token",
			Name: "Dapp Token",
		},
	},
	"P2P": {
		{
			ID:   7075,
			Slug: "p2p",
			Name: "P2P",
		},
	},
	"FRENS": {
		{
			ID:   6551,
			Slug: "frens-community",
			Name: "Frens Community",
		},
	},
	"MOC": {
		{
			ID:   2915,
			Slug: "moss-coin",
			Name: "Moss Coin",
		},
	},
	"ZCN": {
		{
			ID:   2882,
			Slug: "0chain",
			Name: "0Chain",
		},
	},
	"DDIM": {
		{
			ID:   6611,
			Slug: "duckdaodime",
			Name: "DuckDaoDime",
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
	"BREW": {
		{
			ID:   8481,
			Slug: "cafeswap-token",
			Name: "CafeSwap Token",
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
	"ARCO": {
		{
			ID:   1247,
			Slug: "aquariuscoin",
			Name: "AquariusCoin",
		},
	},
	"BITN": {
		{
			ID:   4914,
			Slug: "bitcoin-and-company-network",
			Name: "Bitcoin & Company Network",
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
	"RAZOR": {
		{
			ID:   8409,
			Slug: "razor-network",
			Name: "Razor Network",
		},
	},
	"EOX": {
		{
			ID:   8325,
			Slug: "eox",
			Name: "EOX",
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
	"SAN": {
		{
			ID:   1807,
			Slug: "santiment",
			Name: "Santiment Network Token",
		},
	},
	"BCV": {
		{
			ID:   3066,
			Slug: "bitcapitalvendor",
			Name: "BitCapitalVendor",
		},
	},
	"100X": {
		{
			ID:   9460,
			Slug: "100xcoin",
			Name: "100xCoin",
		},
	},
	"PEKC": {
		{
			ID:   9899,
			Slug: "peacockcoin",
			Name: "PEACOCKCOIN",
		},
	},
	"B26": {
		{
			ID:   8803,
			Slug: "b26-finance",
			Name: "B26 Finance",
		},
	},
	"GIO": {
		{
			ID:   3118,
			Slug: "graviocoin",
			Name: "Graviocoin",
		},
	},
	"NEXT": {
		{
			ID:   3651,
			Slug: "next-coin",
			Name: "NEXT",
		},
	},
	"TYC": {
		{
			ID:   9698,
			Slug: "tycoon",
			Name: "Tycoon",
		},
	},
	"SAFERUNE": {
		{
			ID:   9734,
			Slug: "saferune",
			Name: "Saferune",
		},
	},
	"BRN": {
		{
			ID:   9161,
			Slug: "brainaut-defi",
			Name: "Brainaut Defi",
		},
	},
	"EFL": {
		{
			ID:   234,
			Slug: "e-gulden",
			Name: "e-Gulden",
		},
	},
	"PONZI": {
		{
			ID:   1259,
			Slug: "ponzicoin",
			Name: "PonziCoin",
		},
	},
	"SDX": {
		{
			ID:   9420,
			Slug: "swapdex",
			Name: "SwapDEX",
		},
	},
	"PAYB": {
		{
			ID:   8942,
			Slug: "paybswap",
			Name: "Paybswap",
		},
	},
	"LUCK": {
		{
			ID:   9222,
			Slug: "lucktogether",
			Name: "LuckTogether",
		},
	},
	"PWC": {
		{
			ID:   8750,
			Slug: "prime-whiterock-company",
			Name: "Prime Whiterock Company",
		},
	},
	"FIDA": {
		{
			ID:   7978,
			Slug: "bonfida",
			Name: "Bonfida",
		},
	},
	"SHARE": {
		{
			ID:   6868,
			Slug: "seigniorage-shares",
			Name: "Seigniorage Shares",
		},
	},
	"TRR": {
		{
			ID:   9666,
			Slug: "terran-coin",
			Name: "Terran Coin",
		},
	},
	"VKNF": {
		{
			ID:   8611,
			Slug: "vkenaf",
			Name: "VKENAF",
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
	"TCC": {
		{
			ID:   1894,
			Slug: "the-champcoin",
			Name: "The ChampCoin",
		},
	},
	"TENDIE": {
		{
			ID:   11034,
			Slug: "tendieswap",
			Name: "TendieSwap",
		},
	},
	"BIOS": {
		{
			ID:   10139,
			Slug: "0x-nodes",
			Name: "0x_nodes",
		},
	},
	"PPT": {
		{
			ID:   1789,
			Slug: "populous",
			Name: "Populous",
		},
	},
	"DGD": {
		{
			ID:   1229,
			Slug: "digixdao",
			Name: "DigixDAO",
		},
	},
	"RELI": {
		{
			ID:   9095,
			Slug: "relite-finance",
			Name: "Relite Finance",
		},
	},
	"SEPA": {
		{
			ID:   9163,
			Slug: "secure-pad-token",
			Name: "Secure Pad",
		},
	},
	"ARTE": {
		{
			ID:   6437,
			Slug: "ethart",
			Name: "ethArt",
		},
	},
	"A5T": {
		{
			ID:   7933,
			Slug: "alpha5",
			Name: "Alpha5",
		},
	},
	"GCX": {
		{
			ID:   5322,
			Slug: "germancoin",
			Name: "GermanCoin",
		},
	},
	"DCH": {
		{
			ID:   10020,
			Slug: "dechart",
			Name: "DeChart",
		},
	},
	"LOGE": {
		{
			ID:   10156,
			Slug: "lunadoge",
			Name: "LunaDoge",
		},
	},
	"TRIO": {
		{
			ID:   2661,
			Slug: "tripio",
			Name: "Tripio",
		},
	},
	"Mars": {
		{
			ID:   8253,
			Slug: "mars",
			Name: "Mars",
		},
	},
	"TGIC": {
		{
			ID:   5551,
			Slug: "the-global-index-chain",
			Name: "The global index chain",
		},
	},
	"XVS": {
		{
			ID:   7288,
			Slug: "venus",
			Name: "Venus",
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
	"BUY": {
		{
			ID:   6701,
			Slug: "burency",
			Name: "Burency",
		},
	},
	"AMM": {
		{
			ID:   2286,
			Slug: "micromoney",
			Name: "MicroMoney",
		},
	},
	"CMM": {
		{
			ID:   3080,
			Slug: "commercium",
			Name: "Commercium",
		},
	},
	"DION": {
		{
			ID:   6211,
			Slug: "dionpay",
			Name: "Dionpay",
		},
	},
	"DEFI": {
		{
			ID:   4276,
			Slug: "defi",
			Name: "Defi",
		},
	},
	"WXMR": {
		{
			ID:   8235,
			Slug: "wrapped-monero",
			Name: "Wrapped Monero",
		},
	},
	"LMAO": {
		{
			ID:   10717,
			Slug: "lucky-meow-token",
			Name: "Lucky Meow Token",
		},
	},
	"TKO": {
		{
			ID:   9020,
			Slug: "tokocrypto",
			Name: "Toko Token",
		},
	},
	"APY": {
		{
			ID:   7227,
			Slug: "apy-finance",
			Name: "APY.Finance",
		},
	},
	"AIDOC": {
		{
			ID:   2357,
			Slug: "aidoc",
			Name: "AI Doctor",
		},
	},
	"DATX": {
		{
			ID:   2567,
			Slug: "datx",
			Name: "DATx",
		},
	},
	"ELD": {
		{
			ID:   3746,
			Slug: "electrum-dark-eld",
			Name: "Electrum Dark",
		},
	},
	"MAUSDT": {
		{
			ID:   8919,
			Slug: "matic-aave-usdt",
			Name: "Matic Aave Interest Bearing USDT",
		},
	},
	"TPT": {
		{
			ID:   5947,
			Slug: "tokenpocket",
			Name: "TokenPocket",
		},
	},
	"GRIMEX": {
		{
			ID:   10158,
			Slug: "spacegrime",
			Name: "SpaceGrime",
		},
	},
	"TRONX": {
		{
			ID:   8221,
			Slug: "tronx-coin",
			Name: "Tronx Coin",
		},
	},
	"CNR": {
		{
			ID:   10555,
			Slug: "canary",
			Name: "Canary",
		},
	},
	"PLTC": {
		{
			ID:   3753,
			Slug: "platoncoin",
			Name: "PlatonCoin",
		},
	},
	"BGTT": {
		{
			ID:   6943,
			Slug: "baguette-token",
			Name: "Baguette Token",
		},
	},
	"SFCP": {
		{
			ID:   3856,
			Slug: "sf-capital",
			Name: "SF Capital",
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
	"NET": {
		{
			ID:   3788,
			Slug: "next",
			Name: "NEXT",
		},
	},
	"SEFI": {
		{
			ID:   9123,
			Slug: "sefi",
			Name: "SEFI",
		},
	},
	"EOSDT": {
		{
			ID:   4017,
			Slug: "eosdt",
			Name: "EOSDT",
		},
	},
	"ATCC": {
		{
			ID:   1751,
			Slug: "atc-coin",
			Name: "ATC Coin",
		},
	},
	"FTS": {
		{
			ID:   9592,
			Slug: "fortress-lending",
			Name: "Fortress Lending",
		},
	},
	"BASIC": {
		{
			ID:   5481,
			Slug: "basic",
			Name: "BASIC",
		},
	},
	"PMGT": {
		{
			ID:   5203,
			Slug: "perth-mint-gold-token",
			Name: "Perth Mint Gold Token",
		},
	},
	"RSV": {
		{
			ID:   6727,
			Slug: "reserve",
			Name: "Reserve",
		},
	},
	"NBR": {
		{
			ID:   3006,
			Slug: "niobio",
			Name: "Niobio",
		},
	},
	"KMW": {
		{
			ID:   5282,
			Slug: "kepler-network",
			Name: "Kepler Network",
		},
	},
	"BSV": {
		{
			ID:   3602,
			Slug: "bitcoin-sv",
			Name: "Bitcoin SV",
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
	"BBT": {
		{
			ID:   10201,
			Slug: "bitbook",
			Name: "BitBook",
		},
	},
	"JUL": {
		{
			ID:   6937,
			Slug: "justliquidity",
			Name: "JustLiquidity",
		},
	},
	"DOGGY": {
		{
			ID:   9693,
			Slug: "doggy",
			Name: "DOGGY",
		},
	},
	"DPY": {
		{
			ID:   2162,
			Slug: "delphy",
			Name: "Delphy",
		},
	},
	"CJ": {
		{
			ID:   1306,
			Slug: "cryptojacks",
			Name: "Cryptojacks",
		},
	},
	"UUNICLY": {
		{
			ID:   9469,
			Slug: "unicly-genesis-collection",
			Name: "Unicly Genesis Collection",
		},
	},
	"MSB": {
		{
			ID:   7591,
			Slug: "misbloc",
			Name: "Misbloc",
		},
	},
	"NDAU": {
		{
			ID:   6524,
			Slug: "ndau",
			Name: "Ndau",
		},
	},
	"IMGC": {
		{
			ID:   4344,
			Slug: "imagecash",
			Name: "ImageCash",
		},
	},
	"ETHUP": {
		{
			ID:   7016,
			Slug: "ethup",
			Name: "ETHUP",
		},
	},
	"HOKK": {
		{
			ID:   9458,
			Slug: "hokkaido-inu",
			Name: "Hokkaidu Inu",
		},
	},
	"BXA": {
		{
			ID:   6225,
			Slug: "blockchain-exchange-alliance",
			Name: "Blockchain Exchange Alliance",
		},
	},
	"JOYS": {
		{
			ID:   5278,
			Slug: "joys-digital",
			Name: "Joys Digital",
		},
	},
	"DONI": {
		{
			ID:   10940,
			Slug: "doni-coin",
			Name: "Doni Coin",
		},
	},
	"HXN": {
		{
			ID:   8184,
			Slug: "havens-nook",
			Name: "Havens Nook",
		},
	},
	"PLBT": {
		{
			ID:   1784,
			Slug: "polybius",
			Name: "Polybius",
		},
	},
	"KLEE": {
		{
			ID:   10262,
			Slug: "kleekai",
			Name: "KleeKai",
		},
	},
	"HBAR": {
		{
			ID:   4642,
			Slug: "hedera-hashgraph",
			Name: "Hedera Hashgraph",
		},
	},
	"HBX": {
		{
			ID:   3769,
			Slug: "hashsbx",
			Name: "HashBX",
		},
	},
	"INVE": {
		{
			ID:   3597,
			Slug: "intervalue",
			Name: "InterValue",
		},
	},
	"MFC": {
		{
			ID:   3837,
			Slug: "mfcoin",
			Name: "MFCoin",
		},
	},
	"STOS": {
		{
			ID:   9760,
			Slug: "stratos",
			Name: "Stratos",
		},
	},
	"AXIA": {
		{
			ID:   7474,
			Slug: "axia-protocol",
			Name: "Axia Protocol",
		},
	},
	"FORS": {
		{
			ID:   6918,
			Slug: "foresight",
			Name: "Foresight",
		},
	},
	"SDAO": {
		{
			ID:   9638,
			Slug: "singularitydao",
			Name: "SingularityDAO",
		},
	},
	"MNST": {
		{
			ID:   9679,
			Slug: "moonstarter",
			Name: "MoonStarter",
		},
	},
	"PYE": {
		{
			ID:   10174,
			Slug: "creampye",
			Name: "CREAMPYE",
		},
	},
	"IFUND": {
		{
			ID:   8767,
			Slug: "unifund",
			Name: "Unifund",
		},
	},
	"AMI": {
		{
			ID:   9008,
			Slug: "ammyi-coin",
			Name: "AMMYI Coin",
		},
	},
	"CNN": {
		{
			ID:   2735,
			Slug: "content-neutrality-network",
			Name: "Content Neutrality Network",
		},
	},
	"HYDRO": {
		{
			ID:   2698,
			Slug: "hydro",
			Name: "Hydro",
		},
	},
	"STEP": {
		{
			ID:   9443,
			Slug: "step-finance",
			Name: "Step Finance",
		},
	},
	"ZEUS": {
		{
			ID:   3414,
			Slug: "zeusnetwork",
			Name: "ZeusNetwork",
		},
	},
	"CRDT": {
		{
			ID:   5473,
			Slug: "crdt",
			Name: "CRDT",
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
	"UTT": {
		{
			ID:   2371,
			Slug: "uttoken",
			Name: "United Traders Token",
		},
	},
	"LEAN": {
		{
			ID:   10091,
			Slug: "lean-protocol",
			Name: "Lean",
		},
	},
	"RKN": {
		{
			ID:   5072,
			Slug: "rakon",
			Name: "Rakon",
		},
	},
	"FIO": {
		{
			ID:   5865,
			Slug: "fio-protocol",
			Name: "FIO Protocol",
		},
	},
	"AUTO": {
		{
			ID:   8387,
			Slug: "auto",
			Name: "Auto",
		},
	},
	"PTF": {
		{
			ID:   7190,
			Slug: "powertrade-fuel",
			Name: "PowerTrade Fuel",
		},
	},
	"BCPT": {
		{
			ID:   2061,
			Slug: "blockmason",
			Name: "Blockmason Credit Protocol",
		},
	},
	"NVDA": {
		{
			ID:   7913,
			Slug: "nvidia-tokenized-stock-ftx",
			Name: "NVIDIA tokenized stock FTX",
		},
	},
	"XAUTBULL": {
		{
			ID:   6238,
			Slug: "3x-long-tether-gold-token",
			Name: "3X Long Tether Gold Token",
		},
	},
	"REEC": {
		{
			ID:   3904,
			Slug: "electronic-energy-coin",
			Name: "Renewable Electronic Energy Coin",
		},
	},
	"COSHI": {
		{
			ID:   9574,
			Slug: "coshi-inu",
			Name: "CoShi Inu",
		},
	},
	"sBTC": {
		{
			ID:   6406,
			Slug: "softbtc",
			Name: "sBTC",
		},
	},
	"ZRX": {
		{
			ID:   1896,
			Slug: "0x",
			Name: "0x",
		},
	},
	"WGR": {
		{
			ID:   1779,
			Slug: "wagerr",
			Name: "Wagerr",
		},
	},
	"ZER": {
		{
			ID:   1578,
			Slug: "zero",
			Name: "Zero",
		},
	},
	"PAYT": {
		{
			ID:   7425,
			Slug: "payaccept",
			Name: "PayAccept",
		},
	},
	"LDN": {
		{
			ID:   7992,
			Slug: "ludena-protocol",
			Name: "Ludena Protocol",
		},
	},
	"MES": {
		{
			ID:   4870,
			Slug: "meschain",
			Name: "MesChain",
		},
	},
	"BOG": {
		{
			ID:   8723,
			Slug: "bogged-finance",
			Name: "Bogged Finance",
		},
	},
	"XPR": {
		{
			ID:   5350,
			Slug: "proton",
			Name: "Proton",
		},
	},
	"BITX": {
		{
			ID:   3011,
			Slug: "bitscreener-token",
			Name: "BitScreener Token",
		},
	},
	"METM": {
		{
			ID:   3148,
			Slug: "metamorph",
			Name: "MetaMorph",
		},
	},
	"MMO": {
		{
			ID:   3449,
			Slug: "mmocoin",
			Name: "MMOCoin",
		},
	},
	"AQUAPIG": {
		{
			ID:   10150,
			Slug: "aqua-pig",
			Name: "Aqua Pig",
		},
	},
	"BONDLY": {
		{
			ID:   7931,
			Slug: "bondly",
			Name: "Bondly",
		},
	},
	"ENTRC": {
		{
			ID:   3876,
			Slug: "entercoin",
			Name: "EnterCoin",
		},
	},
	"QCX": {
		{
			ID:   3883,
			Slug: "quickx-protocol",
			Name: "QuickX Protocol",
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
	"BAFE": {
		{
			ID:   9367,
			Slug: "bafe-io",
			Name: "Bafe io",
		},
	},
	"OCTAX": {
		{
			ID:   10588,
			Slug: "octax-finance",
			Name: "OctaX Finance",
		},
	},
	"CUST": {
		{
			ID:   4240,
			Slug: "custody-token",
			Name: "Custody Token",
		},
	},
	"TOM": {
		{
			ID:   7734,
			Slug: "tom-finance",
			Name: "TOM Finance",
		},
	},
	"CT": {
		{
			ID:   8779,
			Slug: "communitytoken",
			Name: "CommunityToken",
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
	"CIX100": {
		{
			ID:   4067,
			Slug: "cryptoindex-com-100",
			Name: "Cryptoindex.com 100",
		},
	},
	"BLANK": {
		{
			ID:   8695,
			Slug: "blank-wallet",
			Name: "Blank Wallet",
		},
	},
	"CFi": {
		{
			ID:   7699,
			Slug: "cyberfi",
			Name: "CyberFi Token",
		},
	},
	"DEM": {
		{
			ID:   72,
			Slug: "deutsche-emark",
			Name: "Deutsche eMark",
		},
	},
	"BMON": {
		{
			ID:   10704,
			Slug: "binamon",
			Name: "Binamon",
		},
	},
	"POLYDOGE": {
		{
			ID:   10088,
			Slug: "polydoge",
			Name: "PolyDoge",
		},
	},
	"YFIE": {
		{
			ID:   6708,
			Slug: "yfiexchange-finance",
			Name: "YFIEXCHANGE.FINANCE",
		},
	},
	"HATCH": {
		{
			ID:   7109,
			Slug: "hatch-dao",
			Name: "Hatch DAO",
		},
	},
	"INSTAR": {
		{
			ID:   2558,
			Slug: "insights-network",
			Name: "Insights Network",
		},
	},
	"EZY": {
		{
			ID:   5521,
			Slug: "ezystayz",
			Name: "EzyStayz",
		},
	},
	"TIDAL": {
		{
			ID:   8912,
			Slug: "tidal-finance",
			Name: "Tidal Finance",
		},
	},
	"ADC": {
		{
			ID:   948,
			Slug: "audiocoin",
			Name: "AudioCoin",
		},
	},
	"EBOX": {
		{
			ID:   8930,
			Slug: "ethbox",
			Name: "Ethbox",
		},
	},
	"FNK": {
		{
			ID:   8014,
			Slug: "fnk-wallet",
			Name: "FNK wallet",
		},
	},
	"ETLT": {
		{
			ID:   7630,
			Slug: "ethereumlightning",
			Name: "Ethereum Lightning",
		},
	},
	"SHPP": {
		{
			ID:   9316,
			Slug: "shipit-pro",
			Name: "Shipit pro",
		},
	},
	"YOT": {
		{
			ID:   8247,
			Slug: "payyoda",
			Name: "PayYoda",
		},
	},
	"ETF": {
		{
			ID:   5441,
			Slug: "entherfound",
			Name: "Entherfound",
		},
	},
	"KDAG": {
		{
			ID:   5626,
			Slug: "king-dag",
			Name: "King DAG",
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
	"FLOT": {
		{
			ID:   3247,
			Slug: "fire-lotto",
			Name: "Fire Lotto",
		},
	},
	"XLMUP": {
		{
			ID:   8055,
			Slug: "xlmup",
			Name: "XLMUP",
		},
	},
	"BNBD": {
		{
			ID:   9667,
			Slug: "bnb-diamond",
			Name: "BNB Diamond",
		},
	},
	"DOUGH": {
		{
			ID:   7284,
			Slug: "piedao-dough-v2",
			Name: "PieDAO DOUGH v2",
		},
	},
	"WIZ": {
		{
			ID:   3331,
			Slug: "crowdwiz",
			Name: "CrowdWiz",
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
	"FNX": {
		{
			ID:   5712,
			Slug: "finnexus",
			Name: "FinNexus",
		},
	},
	"XRPBULL": {
		{
			ID:   5412,
			Slug: "3x-long-xrp-token",
			Name: "3x Long XRP Token",
		},
	},
	"TOSHI": {
		{
			ID:   8744,
			Slug: "toshimon",
			Name: "Toshimon",
		},
	},
	"AQUARI": {
		{
			ID:   9921,
			Slug: "aquari",
			Name: "Aquari",
		},
	},
	"EURX": {
		{
			ID:   5507,
			Slug: "etoro-euro",
			Name: "eToro Euro",
		},
	},
	"TRISM": {
		{
			ID:   8185,
			Slug: "trism",
			Name: "Trism",
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
	"MARK": {
		{
			ID:   7765,
			Slug: "benchmark-protocol",
			Name: "Benchmark Protocol",
		},
	},
	"IDEA": {
		{
			ID:   7681,
			Slug: "ideaology",
			Name: "Ideaology",
		},
	},
	"GMAT": {
		{
			ID:   4182,
			Slug: "gowithmi",
			Name: "GoWithMi",
		},
	},
	"BDCC": {
		{
			ID:   4427,
			Slug: "bdcc-bitica-coin",
			Name: "BDCC Bitica COIN",
		},
	},
	"MOONSHOT": {
		{
			ID:   9140,
			Slug: "moonshot",
			Name: "Moonshot",
		},
	},
	"WIS": {
		{
			ID:   7697,
			Slug: "experty-wisdom-token",
			Name: "Experty Wisdom Token",
		},
	},
	"BOO": {
		{
			ID:   9683,
			Slug: "booster",
			Name: "Booster",
		},
	},
	"PARTY": {
		{
			ID:   6556,
			Slug: "money-party",
			Name: "MONEY PARTY",
		},
	},
	"FRN": {
		{
			ID:   1164,
			Slug: "francs",
			Name: "Francs",
		},
	},
	"UNDO": {
		{
			ID:   10589,
			Slug: "undotoken",
			Name: "UndoToken",
		},
	},
	"SWING": {
		{
			ID:   1085,
			Slug: "swing",
			Name: "Swing",
		},
	},
	"BBI": {
		{
			ID:   8578,
			Slug: "bigboys-industry",
			Name: "BigBoys Industry",
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
	"SAFEMOONCASH": {
		{
			ID:   10081,
			Slug: "safemooncash",
			Name: "SafeMoonCash",
		},
	},
	"KPAD": {
		{
			ID:   8897,
			Slug: "kickpad",
			Name: "KickPad",
		},
	},
	"GLDY": {
		{
			ID:   5538,
			Slug: "buzzshow",
			Name: "Buzzshow",
		},
	},
	"TRND": {
		{
			ID:   5967,
			Slug: "trendering",
			Name: "Trendering",
		},
	},
	"MOMO": {
		{
			ID:   9271,
			Slug: "momo-protocol",
			Name: "Momo Protocol",
		},
	},
	"ORA": {
		{
			ID:   10656,
			Slug: "coin-oracle",
			Name: "COIN ORACLE",
		},
	},
	"FIL": {
		{
			ID:   2280,
			Slug: "filecoin",
			Name: "Filecoin",
		},
	},
	"SSS": {
		{
			ID:   5607,
			Slug: "simple-software-solutions",
			Name: "Simple Software Solutions",
		},
	},
	"ALPA": {
		{
			ID:   7618,
			Slug: "alpaca-city",
			Name: "Alpaca City",
		},
	},
	"YFRM": {
		{
			ID:   7530,
			Slug: "yearn-finance-red-moon",
			Name: "Yearn Finance Red Moon",
		},
	},
	"KLT": {
		{
			ID:   9426,
			Slug: "klend",
			Name: "KLend",
		},
	},
	"YFBETA": {
		{
			ID:   6886,
			Slug: "yfbeta",
			Name: "yfBeta",
		},
	},
	"LIMON": {
		{
			ID:   10141,
			Slug: "limon-group",
			Name: "LIMON.GROUP",
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
	"RVC": {
		{
			ID:   5713,
			Slug: "ravencoin-classic",
			Name: "Ravencoin Classic",
		},
	},
	"TUP": {
		{
			ID:   4411,
			Slug: "tenup",
			Name: "TenUp",
		},
	},
	"FRC": {
		{
			ID:   10,
			Slug: "freicoin",
			Name: "Freicoin",
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
	"SCV": {
		{
			ID:   8470,
			Slug: "super-coinview",
			Name: "Super CoinView Token",
		},
	},
	"DGB": {
		{
			ID:   109,
			Slug: "digibyte",
			Name: "DigiByte",
		},
	},
	"INJ": {
		{
			ID:   7226,
			Slug: "injective-protocol",
			Name: "Injective Protocol",
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
	"USDJ": {
		{
			ID:   5446,
			Slug: "usdj",
			Name: "USDJ",
		},
	},
	"DTH": {
		{
			ID:   2528,
			Slug: "dether",
			Name: "Dether",
		},
	},
	"CDL": {
		{
			ID:   4910,
			Slug: "coindeal-token",
			Name: "CoinDeal Token",
		},
	},
	"DEVE": {
		{
			ID:   7766,
			Slug: "divert-finance",
			Name: "Divert Finance",
		},
	},
	"NVA": {
		{
			ID:   7507,
			Slug: "neeva-defi",
			Name: "Neeva Defi",
		},
	},
	"VERA": {
		{
			ID:   4830,
			Slug: "vera",
			Name: "VERA",
		},
	},
	"OFT": {
		{
			ID:   8095,
			Slug: "orient",
			Name: "Orient",
		},
	},
	"GXC": {
		{
			ID:   1750,
			Slug: "gxchain",
			Name: "GXChain",
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
	"XCUR": {
		{
			ID:   7942,
			Slug: "curate",
			Name: "Curate",
		},
	},
	"UDT": {
		{
			ID:   9364,
			Slug: "unlock-protocol",
			Name: "Unlock Protocol",
		},
	},
	"CHUM": {
		{
			ID:   10360,
			Slug: "chumhum",
			Name: "Chumhum",
		},
	},
	"SAINT": {
		{
			ID:   10317,
			Slug: "saint-token",
			Name: "Saint Token",
		},
	},
	"CAPS": {
		{
			ID:   9291,
			Slug: "ternoa",
			Name: "Ternoa",
		},
	},
	"APYS": {
		{
			ID:   8419,
			Slug: "apyswap",
			Name: "APYSwap",
		},
	},
	"PWR": {
		{
			ID:   1279,
			Slug: "powercoin",
			Name: "PWR Coin",
		},
	},
	"CRETH2": {
		{
			ID:   8893,
			Slug: "cream-eth2",
			Name: "Cream ETH 2",
		},
	},
	"MSWAP": {
		{
			ID:   8188,
			Slug: "moneyswap",
			Name: "MoneySwap",
		},
	},
	"WANBTC": {
		{
			ID:   8650,
			Slug: "wanbtc",
			Name: "wanBTC",
		},
	},
	"NFTL": {
		{
			ID:   8784,
			Slug: "nftl-token",
			Name: "NFTL Token",
		},
	},
	"J8T": {
		{
			ID:   2568,
			Slug: "jet8",
			Name: "JET8",
		},
	},
	"ASKO": {
		{
			ID:   5833,
			Slug: "askobar-network",
			Name: "ASKO",
		},
	},
	"VALUE": {
		{
			ID:   7404,
			Slug: "value-defi",
			Name: "Value Liquidity",
		},
	},
	"IRON": {
		{
			ID:   10484,
			Slug: "iron-finance",
			Name: "Iron",
		},
	},
	"MIOTA": {
		{
			ID:   1720,
			Slug: "iota",
			Name: "IOTA",
		},
	},
	"EDN": {
		{
			ID:   3305,
			Slug: "eden",
			Name: "Eden",
		},
	},
	"IBS": {
		{
			ID:   5358,
			Slug: "ibstoken",
			Name: "IBStoken",
		},
	},
	"TKING": {
		{
			ID:   9854,
			Slug: "tiger-king",
			Name: "Tiger King",
		},
	},
	"TESLASAFE": {
		{
			ID:   11019,
			Slug: "teslasafe",
			Name: "TeslaSafe",
		},
	},
	"BXK": {
		{
			ID:   4049,
			Slug: "bitbook-gambling",
			Name: "Bitbook Gambling",
		},
	},
	"ELENA": {
		{
			ID:   9612,
			Slug: "elena-protocol",
			Name: "Elena Protocol",
		},
	},
	"mNFLX": {
		{
			ID:   8005,
			Slug: "mirrored-netflix",
			Name: "Mirrored Netflix",
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
	"COOK": {
		{
			ID:   8997,
			Slug: "cook-protocol",
			Name: "Cook Protocol",
		},
	},
	"EPAY": {
		{
			ID:   10374,
			Slug: "ethereumpay",
			Name: "EthereumPay",
		},
	},
	"EROTICA": {
		{
			ID:   10197,
			Slug: "erotica-token",
			Name: "Erotica",
		},
	},
	"CEDS": {
		{
			ID:   6745,
			Slug: "cedars",
			Name: "CEDARS",
		},
	},
	"INCNT": {
		{
			ID:   1475,
			Slug: "incent",
			Name: "Incent",
		},
	},
	"DEEZNUTS": {
		{
			ID:   9873,
			Slug: "deez-nuts",
			Name: "Deez Nuts",
		},
	},
	"DINK": {
		{
			ID:   10710,
			Slug: "dink-doink",
			Name: "Dink Doink",
		},
	},
	"SCAT": {
		{
			ID:   8819,
			Slug: "sad-cat-token",
			Name: "Sad Cat Token",
		},
	},
	"UFFYI": {
		{
			ID:   7273,
			Slug: "unlimited-fiscusfyi",
			Name: "Unlimited FiscusFYI",
		},
	},
	"YFTE": {
		{
			ID:   8269,
			Slug: "yftether",
			Name: "YFTether",
		},
	},
	"YFII": {
		{
			ID:   5957,
			Slug: "yearn-finance-ii",
			Name: "DFI.Money",
		},
	},
	"ITS": {
		{
			ID:   7852,
			Slug: "iterationsyndicate",
			Name: "IterationSyndicate",
		},
	},
	"WPR": {
		{
			ID:   2511,
			Slug: "wepower",
			Name: "WePower",
		},
	},
	"CARBON": {
		{
			ID:   502,
			Slug: "carboncoin",
			Name: "Carboncoin",
		},
	},
	"CVX": {
		{
			ID:   9903,
			Slug: "convex-finance",
			Name: "Convex Finance",
		},
	},
	"PC": {
		{
			ID:   3061,
			Slug: "promotion-coin",
			Name: "Promotion Coin",
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
	"KOBO": {
		{
			ID:   815,
			Slug: "kobocoin",
			Name: "Kobocoin",
		},
	},
	"XGT": {
		{
			ID:   8607,
			Slug: "xion-finance",
			Name: "Xion Finance",
		},
	},
	"YNI": {
		{
			ID:   8725,
			Slug: "yearnyfi-network",
			Name: "YEARNYFI NETWORK",
		},
	},
	"ALPHA": {
		{
			ID:   7232,
			Slug: "alpha-finance-lab",
			Name: "Alpha Finance Lab",
		},
	},
	"GEX": {
		{
			ID:   4150,
			Slug: "globex",
			Name: "GLOBEX",
		},
	},
	"GLOB": {
		{
			ID:   5076,
			Slug: "global-reserve-system",
			Name: "Global Reserve System",
		},
	},
	"KUMA": {
		{
			ID:   10394,
			Slug: "kuma-inu",
			Name: "Kuma Inu",
		},
	},
	"SUB": {
		{
			ID:   1984,
			Slug: "substratum",
			Name: "Substratum",
		},
	},
	"MUE": {
		{
			ID:   706,
			Slug: "monetaryunit",
			Name: "MonetaryUnit",
		},
	},
	"ETHY": {
		{
			ID:   7743,
			Slug: "ethereum-yield",
			Name: "Ethereum Yield",
		},
	},
	"DACXI": {
		{
			ID:   10372,
			Slug: "dacxi",
			Name: "Dacxi",
		},
	},
	"BIST": {
		{
			ID:   9889,
			Slug: "bistroo",
			Name: "Bistroo",
		},
	},
	"FX1": {
		{
			ID:   6444,
			Slug: "fanzy",
			Name: "FANZY",
		},
	},
	"QKC": {
		{
			ID:   2840,
			Slug: "quarkchain",
			Name: "QuarkChain",
		},
	},
	"MVP": {
		{
			ID:   2869,
			Slug: "merculet",
			Name: "Merculet",
		},
	},
	"ADAUP": {
		{
			ID:   7013,
			Slug: "adaup",
			Name: "ADAUP",
		},
	},
	"PBOM": {
		{
			ID:   8928,
			Slug: "pocket-bomb",
			Name: "Pocket Bomb",
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
	"PWRB": {
		{
			ID:   5965,
			Slug: "powerbalt",
			Name: "PowerBalt",
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
	"VLC": {
		{
			ID:   2488,
			Slug: "valuechain",
			Name: "ValueChain",
		},
	},
	"KYSC": {
		{
			ID:   6110,
			Slug: "kysc-token",
			Name: "KYSC Token",
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
	"YOC": {
		{
			ID:   1156,
			Slug: "yocoin",
			Name: "Yocoin",
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
	"RAIL": {
		{
			ID:   10854,
			Slug: "railgun",
			Name: "Railgun",
		},
	},
	"DEK": {
		{
			ID:   10172,
			Slug: "dekbox",
			Name: "DekBox",
		},
	},
	"MAYP": {
		{
			ID:   4077,
			Slug: "maya-preferred",
			Name: "Maya Preferred",
		},
	},
	"MRPH": {
		{
			ID:   2763,
			Slug: "morpheus-network",
			Name: "Morpheus.Network",
		},
	},
	"DAV": {
		{
			ID:   3220,
			Slug: "dav-coin",
			Name: "DAV Coin",
		},
	},
	"NANJ": {
		{
			ID:   2595,
			Slug: "nanjcoin",
			Name: "NANJCOIN",
		},
	},
	"LYNC": {
		{
			ID:   7329,
			Slug: "lync-network",
			Name: "LYNC Network",
		},
	},
	"MIB": {
		{
			ID:   3210,
			Slug: "mib-coin",
			Name: "MIB Coin",
		},
	},
	"ALU": {
		{
			ID:   9637,
			Slug: "altura",
			Name: "Altura",
		},
	},
	"HIP": {
		{
			ID:   10297,
			Slug: "hippo-token",
			Name: "HIPPO TOKEN",
		},
	},
	"ODEX": {
		{
			ID:   3954,
			Slug: "one-dex",
			Name: "One DEX",
		},
	},
	"BABYWOLF": {
		{
			ID:   10769,
			Slug: "baby-moon-wolf",
			Name: "Baby Moon Wolf",
		},
	},
	"HBD": {
		{
			ID:   5375,
			Slug: "hive-dollar",
			Name: "Hive Dollar",
		},
	},
	"ARCT": {
		{
			ID:   2415,
			Slug: "arbitragect",
			Name: "ArbitrageCT",
		},
	},
	"EXF": {
		{
			ID:   8563,
			Slug: "extend-finance",
			Name: "Extend Finance",
		},
	},
	"DEEP": {
		{
			ID:   4251,
			Slug: "deepcloud-ai",
			Name: "DeepCloud AI",
		},
	},
	"GTF": {
		{
			ID:   6457,
			Slug: "globaltrustfund-token",
			Name: "GLOBALTRUSTFUND TOKEN",
		},
	},
	"FRED": {
		{
			ID:   5109,
			Slug: "fred-energy",
			Name: "FRED Energy",
		},
	},
	"ETHMACOAPY": {
		{
			ID:   6130,
			Slug: "eth-20-day-ma-crossover-yield-set",
			Name: "ETH 20 Day MA Crossover Yield Set",
		},
	},
	"XRP": {
		{
			ID:   52,
			Slug: "xrp",
			Name: "XRP",
		},
	},
	"DLT": {
		{
			ID:   1949,
			Slug: "agrello-delta",
			Name: "Agrello",
		},
	},
	"DIAMND": {
		{
			ID:   11011,
			Slug: "projekt-diamond",
			Name: "Projekt Diamond",
		},
	},
	"CLV": {
		{
			ID:   8384,
			Slug: "clover",
			Name: "Clover Finance",
		},
	},
	"MIL": {
		{
			ID:   11009,
			Slug: "military-finance",
			Name: "Military Finance",
		},
	},
	"REAU": {
		{
			ID:   9413,
			Slug: "vira-lata-finance",
			Name: "Vira-lata Finance",
		},
	},
	"KPC": {
		{
			ID:   8573,
			Slug: "koloop-basic",
			Name: "Koloop Basic",
		},
	},
	"DFIP": {
		{
			ID:   7333,
			Slug: "defi-insurance-protocol",
			Name: "DeFi Insurance Protocol",
		},
	},
	"OCP": {
		{
			ID:   6693,
			Slug: "oc-protocol",
			Name: "OC Protocol",
		},
	},
	"ZCK": {
		{
			ID:   8687,
			Slug: "polkazeck",
			Name: "Polkazeck",
		},
	},
	"DTRC": {
		{
			ID:   2752,
			Slug: "datarius-credit",
			Name: "Datarius Credit",
		},
	},
	"TRONPAD": {
		{
			ID:   10277,
			Slug: "tronpad",
			Name: "TRONPAD",
		},
	},
	"XIASI": {
		{
			ID:   10636,
			Slug: "xiasi-inu",
			Name: "Xiasi Inu",
		},
	},
	"MORA": {
		{
			ID:   9031,
			Slug: "meliora",
			Name: "Meliora",
		},
	},
	"SYM": {
		{
			ID:   4824,
			Slug: "symverse",
			Name: "SymVerse",
		},
	},
	"ASY": {
		{
			ID:   5529,
			Slug: "asyagro",
			Name: "ASYAGRO",
		},
	},
	"SLM": {
		{
			ID:   8268,
			Slug: "solomon-defi",
			Name: "Solomon Defi",
		},
	},
	"KNC": {
		{
			ID:   1982,
			Slug: "kyber-network-crystal-legacy",
			Name: "Kyber Network Crystal Legacy",
		},
	},
	"vBTC": {
		{
			ID:   7962,
			Slug: "venus-btc",
			Name: "Venus BTC",
		},
	},
	"OXEN": {
		{
			ID:   2748,
			Slug: "oxen",
			Name: "Oxen",
		},
	},
	"EFX": {
		{
			ID:   2666,
			Slug: "effect-ai",
			Name: "Effect.AI",
		},
	},
	"VEIL": {
		{
			ID:   3830,
			Slug: "veil",
			Name: "Veil",
		},
	},
	"ESBC": {
		{
			ID:   3879,
			Slug: "esbc",
			Name: "ESBC",
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
	"FRONT": {
		{
			ID:   5893,
			Slug: "frontier",
			Name: "Frontier",
		},
	},
	"YOYOW": {
		{
			ID:   1899,
			Slug: "yoyow",
			Name: "YOYOW",
		},
	},
	"ELP": {
		{
			ID:   10392,
			Slug: "the-everlasting-parachain",
			Name: "The Everlasting Parachain",
		},
	},
	"CUT": {
		{
			ID:   4752,
			Slug: "cutcoin",
			Name: "CUTcoin",
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
	"GXX": {
		{
			ID:   977,
			Slug: "gravitycoin",
			Name: "GravityCoin",
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
	"UNW": {
		{
			ID:   7635,
			Slug: "uniworld",
			Name: "UniWorld",
		},
	},
	"HOTDOGE": {
		{
			ID:   9961,
			Slug: "hotdoge",
			Name: "HotDoge",
		},
	},
	"SIM": {
		{
			ID:   9948,
			Slug: "simba-inu",
			Name: "Simba Inu",
		},
	},
	"RTF": {
		{
			ID:   10097,
			Slug: "regiment-finance",
			Name: "Regiment Finance",
		},
	},
	"HYFI": {
		{
			ID:   9235,
			Slug: "hyper-finance",
			Name: "Hyper Finance",
		},
	},
	"KALLY": {
		{
			ID:   9631,
			Slug: "polkally",
			Name: "Polkally",
		},
	},
	"YFARMER": {
		{
			ID:   7061,
			Slug: "yfarmland-token",
			Name: "YFarmLand Token",
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
	"EASY": {
		{
			ID:   7332,
			Slug: "easyfi",
			Name: "EasyFi",
		},
	},
	"KKO": {
		{
			ID:   9718,
			Slug: "kineko",
			Name: "Kineko",
		},
	},
	"XENO": {
		{
			ID:   4205,
			Slug: "xenoverse",
			Name: "Xenoverse",
		},
	},
	"MUTE": {
		{
			ID:   8795,
			Slug: "mute",
			Name: "Mute",
		},
	},
	"1AI": {
		{
			ID:   5105,
			Slug: "1ai-token",
			Name: "1AI Token",
		},
	},
	"QUICK": {
		{
			ID:   8206,
			Slug: "quickswap",
			Name: "QuickSwap",
		},
	},
	"EXM": {
		{
			ID:   4974,
			Slug: "exmo-coin",
			Name: "EXMO Coin",
		},
	},
	"CZZ": {
		{
			ID:   10083,
			Slug: "classzz",
			Name: "ClassZZ",
		},
	},
	"TOPB": {
		{
			ID:   6027,
			Slug: "topb",
			Name: "TOPBTC Token",
		},
	},
	"APLP": {
		{
			ID:   7837,
			Slug: "apple-finance",
			Name: "Apple Finance",
		},
	},
	"DRE": {
		{
			ID:   10377,
			Slug: "dare-token",
			Name: "Dare Token",
		},
	},
	"42": {
		{
			ID:   93,
			Slug: "42-coin",
			Name: "42-coin",
		},
	},
	"TARA": {
		{
			ID:   8715,
			Slug: "taraxa",
			Name: "Taraxa",
		},
	},
	"VBit": {
		{
			ID:   10082,
			Slug: "voltbit",
			Name: "Voltbit",
		},
	},
	"CRO": {
		{
			ID:   3635,
			Slug: "crypto-com-coin",
			Name: "Crypto.com Coin",
		},
	},
	"GOM2": {
		{
			ID:   5839,
			Slug: "animalgo",
			Name: "AnimalGo",
		},
	},
	"LUNES": {
		{
			ID:   3786,
			Slug: "lunes",
			Name: "Lunes",
		},
	},
	"YO": {
		{
			ID:   4229,
			Slug: "yobit-token",
			Name: "Yobit Token",
		},
	},
	"xSAT": {
		{
			ID:   9411,
			Slug: "satisfinance-token",
			Name: "SatisFinance Token",
		},
	},
	"HAZE": {
		{
			ID:   8810,
			Slug: "haze-finance",
			Name: "Haze Finance",
		},
	},
	"COUSINDOGE": {
		{
			ID:   10780,
			Slug: "cousin-doge-coin",
			Name: "COUSIN DOGE COIN",
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
	"CWBTC": {
		{
			ID:   5744,
			Slug: "compound-wrapped-btc",
			Name: "Compound Wrapped BTC",
		},
	},
	"CDEX": {
		{
			ID:   4801,
			Slug: "codex",
			Name: "Codex",
		},
	},
	"PDATA": {
		{
			ID:   4088,
			Slug: "pdata",
			Name: "PDATA",
		},
	},
	"vUSDC": {
		{
			ID:   7958,
			Slug: "venus-usdc",
			Name: "Venus USDC",
		},
	},
	"GLT": {
		{
			ID:   1731,
			Slug: "globaltoken",
			Name: "GlobalToken",
		},
	},
	"PRESIDENTDOGE": {
		{
			ID:   10946,
			Slug: "presidentdoge",
			Name: "PresidentDoge",
		},
	},
	"FOCV": {
		{
			ID:   6002,
			Slug: "focv",
			Name: "FOCV",
		},
	},
	"SHFT": {
		{
			ID:   8917,
			Slug: "shyft-network",
			Name: "Shyft Network",
		},
	},
	"RAD": {
		{
			ID:   6843,
			Slug: "radicle",
			Name: "Radicle",
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
	"SFT": {
		{
			ID:   1172,
			Slug: "safex-token",
			Name: "Safex Token",
		},
	},
	"XBY": {
		{
			ID:   1636,
			Slug: "xtrabytes",
			Name: "XTRABYTES",
		},
	},
	"WINALAMBO": {
		{
			ID:   10745,
			Slug: "winalambo-finance",
			Name: "WIN A LAMBO FINANCE",
		},
	},
	"PROB": {
		{
			ID:   4586,
			Slug: "probit-token",
			Name: "ProBit Token",
		},
	},
	"SHPING": {
		{
			ID:   3422,
			Slug: "shping",
			Name: "SHPING",
		},
	},
	"VACAY": {
		{
			ID:   10815,
			Slug: "vacay",
			Name: "Vacay",
		},
	},
	"TSX": {
		{
			ID:   9563,
			Slug: "tradestars",
			Name: "TradeStars",
		},
	},
	"UCX": {
		{
			ID:   6604,
			Slug: "ucx-foundation",
			Name: "UCX FOUNDATION",
		},
	},
	"IDOGE": {
		{
			ID:   10446,
			Slug: "influencer-doge",
			Name: "Influencer Doge",
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
	"SOFI": {
		{
			ID:   6825,
			Slug: "social-finance",
			Name: "Social Finance",
		},
	},
	"LEO": {
		{
			ID:   3957,
			Slug: "unus-sed-leo",
			Name: "UNUS SED LEO",
		},
	},
	"VBIT": {
		{
			ID:   6947,
			Slug: "valobit",
			Name: "Valobit",
		},
	},
	"MARO": {
		{
			ID:   3175,
			Slug: "maro",
			Name: "Maro",
		},
	},
	"PART": {
		{
			ID:   1826,
			Slug: "particl",
			Name: "Particl",
		},
	},
	"GUM": {
		{
			ID:   8386,
			Slug: "gourmet-galaxy",
			Name: "Gourmet Galaxy",
		},
	},
	"DMOON": {
		{
			ID:   10125,
			Slug: "dragonmoon",
			Name: "DragonMoon",
		},
	},
	"EMRALS": {
		{
			ID:   5259,
			Slug: "emrals",
			Name: "Emrals",
		},
	},
	"MTF": {
		{
			ID:   9373,
			Slug: "milktea-finance",
			Name: "Milktea.finance",
		},
	},
	"FYZ": {
		{
			ID:   6674,
			Slug: "fyooz",
			Name: "Fyooz",
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
	"MAP": {
		{
			ID:   4956,
			Slug: "map-protocol",
			Name: "MAP Protocol",
		},
	},
	"OBSR": {
		{
			ID:   3698,
			Slug: "observer",
			Name: "Observer",
		},
	},
	"ROO": {
		{
			ID:   10574,
			Slug: "roocoin",
			Name: "RooCoin",
		},
	},
	"HBDC": {
		{
			ID:   6542,
			Slug: "happy-birthday-coin",
			Name: "happy birthday coin",
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
	"TAP": {
		{
			ID:   8463,
			Slug: "tapmydata",
			Name: "Tapmydata",
		},
	},
	"EPK": {
		{
			ID:   9537,
			Slug: "epik-protocol",
			Name: "EpiK Protocol",
		},
	},
	"S4F": {
		{
			ID:   3733,
			Slug: "s4fe",
			Name: "S4FE",
		},
	},
	"SIG": {
		{
			ID:   8598,
			Slug: "xsigma",
			Name: "xSigma",
		},
	},
	"SFD": {
		{
			ID:   7270,
			Slug: "safe-deal",
			Name: "SAFE DEAL",
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
	"LCG": {
		{
			ID:   7448,
			Slug: "lcg",
			Name: "LCG",
		},
	},
	"TEAT": {
		{
			ID:   7614,
			Slug: "teal",
			Name: "TEAL",
		},
	},
	"IDO": {
		{
			ID:   9365,
			Slug: "idohunt-app",
			Name: "IDOHunt app",
		},
	},
	"GR": {
		{
			ID:   8035,
			Slug: "grom",
			Name: "Grom",
		},
	},
	"SMOON": {
		{
			ID:   9969,
			Slug: "saylormoon",
			Name: "SaylorMoon",
		},
	},
	"VGD": {
		{
			ID:   9240,
			Slug: "vangold-token",
			Name: "Vangold Token",
		},
	},
	"PCN": {
		{
			ID:   1803,
			Slug: "peepcoin",
			Name: "PeepCoin",
		},
	},
	"BADGER": {
		{
			ID:   7859,
			Slug: "badger-dao",
			Name: "Badger DAO",
		},
	},
	"VRSC": {
		{
			ID:   5049,
			Slug: "veruscoin",
			Name: "VerusCoin",
		},
	},
	"BOSON": {
		{
			ID:   8827,
			Slug: "boson-protocol",
			Name: "Boson Protocol",
		},
	},
	"HPB": {
		{
			ID:   2345,
			Slug: "high-performance-blockchain",
			Name: "High Performance Blockchain",
		},
	},
	"ATN": {
		{
			ID:   2380,
			Slug: "atn",
			Name: "ATN",
		},
	},
	"DFC": {
		{
			ID:   10802,
			Slug: "defi-city",
			Name: "DeFi City",
		},
	},
	"CHY": {
		{
			ID:   10885,
			Slug: "concern-poverty-chain",
			Name: "Concern Poverty Chain",
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
	"MGB": {
		{
			ID:   9165,
			Slug: "magic-balancer",
			Name: "Magic Balancer",
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
	"JINU": {
		{
			ID:   10445,
			Slug: "jomon-inu",
			Name: "Jomon Inu",
		},
	},
	"UMA": {
		{
			ID:   5617,
			Slug: "uma",
			Name: "UMA",
		},
	},
	"XSM": {
		{
			ID:   3696,
			Slug: "spectrumcash",
			Name: "SpectrumCash",
		},
	},
	"MQL": {
		{
			ID:   6804,
			Slug: "miraqle",
			Name: "MiraQle",
		},
	},
	"DGP": {
		{
			ID:   7864,
			Slug: "dgpayment",
			Name: "DGPayment",
		},
	},
	"PUNK": {
		{
			ID:   8490,
			Slug: "punk",
			Name: "Punk",
		},
	},
	"THETA": {
		{
			ID:   2416,
			Slug: "theta",
			Name: "THETA",
		},
	},
	"ES": {
		{
			ID:   4860,
			Slug: "era-swap",
			Name: "Era Swap",
		},
	},
	"XLAB": {
		{
			ID:   4709,
			Slug: "xceltoken-plus",
			Name: "XcelToken Plus",
		},
	},
	"FDZ": {
		{
			ID:   2626,
			Slug: "friends",
			Name: "Friendz",
		},
	},
	"GALI": {
		{
			ID:   3793,
			Slug: "galilel",
			Name: "Galilel",
		},
	},
	"XBN": {
		{
			ID:   9385,
			Slug: "xbn",
			Name: "Elastic BNB",
		},
	},
	"USDN": {
		{
			ID:   5068,
			Slug: "neutrino-usd",
			Name: "Neutrino USD",
		},
	},
	"RBTC": {
		{
			ID:   3626,
			Slug: "rsk-smart-bitcoin",
			Name: "RSK Smart Bitcoin",
		},
	},
	"SNN": {
		{
			ID:   6179,
			Slug: "sechain",
			Name: "SeChain",
		},
	},
	"ELE": {
		{
			ID:   10271,
			Slug: "eleven-finance",
			Name: "Eleven Finance",
		},
	},
	"XLMG": {
		{
			ID:   3926,
			Slug: "stellar-gold",
			Name: "Stellar Gold",
		},
	},
	"ALGOBEAR": {
		{
			ID:   6091,
			Slug: "3x-short-algorand-token",
			Name: "3X Short Algorand Token",
		},
	},
	"LED": {
		{
			ID:   10598,
			Slug: "ledgerscore",
			Name: "LedgerScore",
		},
	},
	"DICK": {
		{
			ID:   10027,
			Slug: "dick",
			Name: "Dick",
		},
	},
	"KMD": {
		{
			ID:   1521,
			Slug: "komodo",
			Name: "Komodo",
		},
	},
	"10SET": {
		{
			ID:   9089,
			Slug: "tenset",
			Name: "Tenset",
		},
	},
	"RAK": {
		{
			ID:   7632,
			Slug: "rake-finance",
			Name: "Rake Finance",
		},
	},
	"CHEE": {
		{
			ID:   10010,
			Slug: "cheecoin",
			Name: "Cheecoin",
		},
	},
	"OCTO": {
		{
			ID:   7202,
			Slug: "octofi",
			Name: "OctoFi",
		},
	},
	"DRG": {
		{
			ID:   2593,
			Slug: "dragon-coins",
			Name: "Dragon Coins",
		},
	},
	"DRT": {
		{
			ID:   2070,
			Slug: "domraider",
			Name: "DomRaider",
		},
	},
	"XCF": {
		{
			ID:   8441,
			Slug: "cenfura-token",
			Name: "Cenfura Token",
		},
	},
	"SPOK": {
		{
			ID:   5312,
			Slug: "spockchain-network",
			Name: "Spockchain Network",
		},
	},
	"SAFEYIELD": {
		{
			ID:   9256,
			Slug: "safeyield",
			Name: "SafeYield",
		},
	},
	"UOP": {
		{
			ID:   8150,
			Slug: "utopia-genesis-foundation",
			Name: "Utopia Genesis Foundation",
		},
	},
	"EMT": {
		{
			ID:   3993,
			Slug: "emanate",
			Name: "Emanate",
		},
	},
	"DAPP": {
		{
			ID:   4026,
			Slug: "liquid-apps",
			Name: "LiquidApps",
		},
	},
	"NAVY": {
		{
			ID:   3805,
			Slug: "boat-pilot-token",
			Name: "BoatPilot Token",
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
	"HUNGRY": {
		{
			ID:   9231,
			Slug: "hungry-bear",
			Name: "Hungry Bear",
		},
	},
	"CICX": {
		{
			ID:   5406,
			Slug: "cicoin",
			Name: "Cicoin",
		},
	},
	"PBR": {
		{
			ID:   8320,
			Slug: "polkabridge",
			Name: "PolkaBridge",
		},
	},
	"BTCBAM": {
		{
			ID:   9270,
			Slug: "bitcoin-bam",
			Name: "Bitcoin Bam",
		},
	},
	"DINA": {
		{
			ID:   10413,
			Slug: "dina",
			Name: "Dina",
		},
	},
	"ARTEON": {
		{
			ID:   9591,
			Slug: "arteon",
			Name: "Arteon",
		},
	},
	"HNY": {
		{
			ID:   7972,
			Slug: "honey-token",
			Name: "Honey",
		},
	},
	"PRDZ": {
		{
			ID:   7998,
			Slug: "predictz",
			Name: "Predictz",
		},
	},
	"RSTR": {
		{
			ID:   3407,
			Slug: "ondori",
			Name: "Ondori",
		},
	},
	"HTN": {
		{
			ID:   6464,
			Slug: "heartnumber",
			Name: "Heart Number",
		},
	},
	"UPT": {
		{
			ID:   5555,
			Slug: "universal-protocol-token",
			Name: "Universal Protocol Token",
		},
	},
	"FTV": {
		{
			ID:   10650,
			Slug: "futurov-governance-token",
			Name: "Futurov Governance Token",
		},
	},
	"CHINU": {
		{
			ID:   10056,
			Slug: "chubby-inu",
			Name: "Chubby Inu",
		},
	},
	"POND": {
		{
			ID:   7497,
			Slug: "marlin",
			Name: "Marlin",
		},
	},
	"SKB": {
		{
			ID:   2742,
			Slug: "sakura-bloom",
			Name: "Sakura Bloom",
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
	"QBZ": {
		{
			ID:   5267,
			Slug: "queenbee",
			Name: "QUEENBEE",
		},
	},
	"DOGEBACK": {
		{
			ID:   10792,
			Slug: "doge-back",
			Name: "Doge Back",
		},
	},
	"ZWAP": {
		{
			ID:   9107,
			Slug: "zilswap",
			Name: "Zilswap",
		},
	},
	"GLDR": {
		{
			ID:   6809,
			Slug: "goldergames",
			Name: "GolderGames",
		},
	},
	"TOKO": {
		{
			ID:   4299,
			Slug: "tokoin",
			Name: "Tokoin",
		},
	},
	"LOOT": {
		{
			ID:   8030,
			Slug: "nftlootbox",
			Name: "NFTLootBox",
		},
	},
	"ORB": {
		{
			ID:   80,
			Slug: "orbitcoin",
			Name: "Orbitcoin",
		},
	},
	"YEAR": {
		{
			ID:   4525,
			Slug: "lightyears",
			Name: "Lightyears",
		},
	},
	"LVN": {
		{
			ID:   5853,
			Slug: "livenpay",
			Name: "LivenPay",
		},
	},
	"SPACETOAST": {
		{
			ID:   10573,
			Slug: "spacetoast",
			Name: "SpaceToast",
		},
	},
	"VTHO": {
		{
			ID:   3012,
			Slug: "vethor-token",
			Name: "VeThor Token",
		},
	},
	"SMARTCREDIT": {
		{
			ID:   7596,
			Slug: "smartcredit-token",
			Name: "SmartCredit Token",
		},
	},
	"LITTLEDOGE": {
		{
			ID:   10956,
			Slug: "littledoge",
			Name: "LittleDoge",
		},
	},
	"TSM": {
		{
			ID:   7893,
			Slug: "taiwan-semiconductor-mfg-tokenized-stock-ftx",
			Name: "Taiwan Semiconductor Mfg tokenized stock FTX",
		},
	},
	"KISSMYMOON": {
		{
			ID:   10664,
			Slug: "kissmymoon",
			Name: "KissMyMoon",
		},
	},
	"SIX": {
		{
			ID:   3327,
			Slug: "six",
			Name: "SIX",
		},
	},
	"EVZ": {
		{
			ID:   6430,
			Slug: "electric-vehicle-zone",
			Name: "Electric Vehicle Zone",
		},
	},
	"MPAY": {
		{
			ID:   4001,
			Slug: "menapay",
			Name: "MenaPay",
		},
	},
	"FEY": {
		{
			ID:   10361,
			Slug: "feyorra",
			Name: "Feyorra",
		},
	},
	"SDOG": {
		{
			ID:   9680,
			Slug: "small-dogecoin",
			Name: "Small dogecoin",
		},
	},
	"HL": {
		{
			ID:   6258,
			Slug: "hl-chain",
			Name: "HL Chain",
		},
	},
	"DERO": {
		{
			ID:   2665,
			Slug: "dero",
			Name: "Dero",
		},
	},
	"TRU": {
		{
			ID:   7725,
			Slug: "truefi-token",
			Name: "TrueFi",
		},
		// Not on Binance yet
		// {
		// 	ID:   7296,
		// 	Slug: "truebit",
		// 	Name: "Truebit",
		// },
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
	"mFB": {
		{
			ID:   8585,
			Slug: "mirrored-facebook",
			Name: "Mirrored Facebook Inc",
		},
	},
	"DAPPX": {
		{
			ID:   10376,
			Slug: "dappstore",
			Name: "dAppstore",
		},
	},
	"HENTAI": {
		{
			ID:   10667,
			Slug: "hentaicoin",
			Name: "HentaiCoin",
		},
	},
	"GMC": {
		{
			ID:   7161,
			Slug: "gokumarket-credit",
			Name: "GokuMarket Credit",
		},
	},
	"PRV": {
		{
			ID:   9933,
			Slug: "privacyswap",
			Name: "PrivacySwap",
		},
	},
	"BRTR": {
		{
			ID:   6543,
			Slug: "barter",
			Name: "Barter",
		},
	},
	"aENJ": {
		{
			ID:   8594,
			Slug: "aave-enjin",
			Name: "Aave Enjin",
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
	"WFTM": {
		{
			ID:   10240,
			Slug: "wrapped-fantom",
			Name: "Wrapped Fantom",
		},
	},
	"JUV": {
		{
			ID:   5224,
			Slug: "juventus-fan-token",
			Name: "Juventus Fan Token",
		},
	},
	"TRV": {
		{
			ID:   4060,
			Slug: "trustverse",
			Name: "TrustVerse",
		},
	},
	"PLU": {
		{
			ID:   1392,
			Slug: "pluton",
			Name: "Pluton",
		},
	},
	"UNIUSD": {
		{
			ID:   5694,
			Slug: "unidollar",
			Name: "UniDollar",
		},
	},
	"ENCX": {
		{
			ID:   6954,
			Slug: "enceladus-network",
			Name: "Enceladus Network",
		},
	},
	"MAID": {
		{
			ID:   291,
			Slug: "maidsafecoin",
			Name: "MaidSafeCoin",
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
	"ADT": {
		{
			ID:   1775,
			Slug: "adtoken",
			Name: "adToken",
		},
	},
	"TOOLS": {
		{
			ID:   8866,
			Slug: "bsc-tools",
			Name: "BSC TOOLS",
		},
	},
	"SOJU": {
		{
			ID:   8781,
			Slug: "soju-finance",
			Name: "Soju Finance",
		},
	},
	"ELONPEG": {
		{
			ID:   10671,
			Slug: "elonpeg",
			Name: "ElonPeg",
		},
	},
	"UBE": {
		{
			ID:   10808,
			Slug: "ubeswap",
			Name: "Ubeswap",
		},
	},
	"RTE": {
		{
			ID:   2880,
			Slug: "rate3",
			Name: "Rate3",
		},
	},
	"CNNS": {
		{
			ID:   3934,
			Slug: "cnns",
			Name: "CNNS",
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
	"IFC": {
		{
			ID:   41,
			Slug: "infinitecoin",
			Name: "Infinitecoin",
		},
	},
	"DISTX": {
		{
			ID:   6822,
			Slug: "distx",
			Name: "DistX",
		},
	},
	"OCC": {
		{
			ID:   9191,
			Slug: "occamfi",
			Name: "Occam.Fi",
		},
	},
	"SBT": {
		{
			ID:   4913,
			Slug: "solbit",
			Name: "SOLBIT",
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
	"PAINT": {
		{
			ID:   8647,
			Slug: "murall",
			Name: "MurAll",
		},
	},
	"MILK2": {
		{
			ID:   7386,
			Slug: "spaceswap",
			Name: "Spaceswap MILK2",
		},
	},
	"MICRO": {
		{
			ID:   3610,
			Slug: "micromines",
			Name: "Micromines",
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
	"WPP": {
		{
			ID:   3885,
			Slug: "wpp-token",
			Name: "WPP TOKEN",
		},
	},
	"KEYT": {
		{
			ID:   5530,
			Slug: "rebit",
			Name: "REBIT",
		},
	},
	"RX": {
		{
			ID:   9227,
			Slug: "raven-x",
			Name: "Raven X",
		},
	},
	"YUANG": {
		{
			ID:   10327,
			Slug: "yuang-coin",
			Name: "Yuang Coin",
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
	"ACS": {
		{
			ID:   7844,
			Slug: "acryptos",
			Name: "ACryptoS",
		},
	},
	"ARTX": {
		{
			ID:   9105,
			Slug: "artx-trading",
			Name: "ARTX Trading",
		},
	},
	"HNZO": {
		{
			ID:   9765,
			Slug: "hanzo-inu",
			Name: "Hanzo Inu",
		},
	},
	"TCR": {
		{
			ID:   7717,
			Slug: "tecracoin",
			Name: "TecraCoin",
		},
	},
	"LEVL": {
		{
			ID:   4173,
			Slug: "levolution",
			Name: "Levolution",
		},
	},
	"LML": {
		{
			ID:   3636,
			Slug: "lisk-machine-learning",
			Name: "Lisk Machine Learning",
		},
	},
	"YFSI": {
		{
			ID:   7343,
			Slug: "yfscience",
			Name: "Yfscience",
		},
	},
	"NBNG": {
		{
			ID:   10716,
			Slug: "nobunaga-token-nbng",
			Name: "Nobunaga Token, NBNG",
		},
	},
	"WILLIE": {
		{
			ID:   10009,
			Slug: "williecoin",
			Name: "Williecoin",
		},
	},
	"PEECH": {
		{
			ID:   10356,
			Slug: "peach-finance",
			Name: "Peach.Finance",
		},
	},
	"ORAI": {
		{
			ID:   7533,
			Slug: "oraichain-token",
			Name: "Oraichain Token",
		},
	},
	"ELEC": {
		{
			ID:   2573,
			Slug: "electrifyasia",
			Name: "Electrify.Asia",
		},
	},
	"DADDYDOGE": {
		{
			ID:   10899,
			Slug: "daddy-doge",
			Name: "Daddy Doge",
		},
	},
	"TRXUP": {
		{
			ID:   7005,
			Slug: "trxup",
			Name: "TRXUP",
		},
	},
	"MEPAD": {
		{
			ID:   9518,
			Slug: "memepad",
			Name: "MemePad",
		},
	},
	"HAG": {
		{
			ID:   9846,
			Slug: "hagglex",
			Name: "HaggleX",
		},
	},
	"RAMEN": {
		{
			ID:   8547,
			Slug: "ramenswap",
			Name: "RamenSwap",
		},
	},
	"CATGIRL": {
		{
			ID:   10275,
			Slug: "catgirl",
			Name: "Catgirl",
		},
	},
	"COM": {
		{
			ID:   6608,
			Slug: "community-token",
			Name: "Community Token",
		},
	},
	"ENOL": {
		{
			ID:   7782,
			Slug: "ethanol",
			Name: "Ethanol",
		},
	},
	"SOL": {
		{
			ID:   5426,
			Slug: "solana",
			Name: "Solana",
		},
	},
	"HT": {
		{
			ID:   2502,
			Slug: "huobi-token",
			Name: "Huobi Token",
		},
	},
	"TAU": {
		{
			ID:   2337,
			Slug: "lamden",
			Name: "Lamden",
		},
	},
	"AYA": {
		{
			ID:   3973,
			Slug: "aryacoin",
			Name: "Aryacoin",
		},
	},
	"ibETH": {
		{
			ID:   7675,
			Slug: "ibeth",
			Name: "Interest Bearing ETH",
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
	"SWAG": {
		{
			ID:   7428,
			Slug: "swag-finance",
			Name: "SWAG Finance",
		},
	},
	"4ART": {
		{
			ID:   5912,
			Slug: "4artechnologies",
			Name: "4ART Coin",
		},
	},
	"BABYDOGECASH": {
		{
			ID:   10882,
			Slug: "baby-doge-cash",
			Name: "Baby Doge Cash",
		},
	},
	"ZOC": {
		{
			ID:   4546,
			Slug: "01coin",
			Name: "01coin",
		},
	},
	"TCFX": {
		{
			ID:   7175,
			Slug: "tcbcoin",
			Name: "Tcbcoin",
		},
	},
	"1INCH": {
		{
			ID:   8104,
			Slug: "1inch",
			Name: "1inch",
		},
	},
	"NOIA": {
		{
			ID:   4191,
			Slug: "syntropy",
			Name: "Syntropy",
		},
	},
	"BCUG": {
		{
			ID:   6789,
			Slug: "blockchain-cuties-universe",
			Name: "Blockchain Cuties Universe Governance",
		},
	},
	"JET": {
		{
			ID:   1787,
			Slug: "jetcoin",
			Name: "Jetcoin",
		},
	},
	"ROCK2": {
		{
			ID:   3065,
			Slug: "ice-rock-mining",
			Name: "ICE ROCK MINING",
		},
	},
	"QDX": {
		{
			ID:   10079,
			Slug: "quidax",
			Name: "Quidax",
		},
	},
	"YETU": {
		{
			ID:   9139,
			Slug: "yetucoin",
			Name: "Yetucoin",
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
	"AMZ": {
		{
			ID:   6814,
			Slug: "amazonascoin",
			Name: "AmazonasCoin",
		},
	},
	"HYMETEOR": {
		{
			ID:   9699,
			Slug: "hypermeteor",
			Name: "HyperMeteor",
		},
	},
	"ETHPA": {
		{
			ID:   6141,
			Slug: "eth-price-action-candlestick-set",
			Name: "ETH Price Action Candlestick Set",
		},
	},
	"BEER": {
		{
			ID:   5449,
			Slug: "beer-money",
			Name: "Beer Money",
		},
	},
	"UCO": {
		{
			ID:   6887,
			Slug: "uniris",
			Name: "Uniris",
		},
	},
	"MOOV": {
		{
			ID:   10046,
			Slug: "dotmoovs",
			Name: "Dotmoovs",
		},
	},
	"XQN": {
		{
			ID:   733,
			Slug: "quotient",
			Name: "Quotient",
		},
	},
	"OAX": {
		{
			ID:   1853,
			Slug: "oax",
			Name: "OAX",
		},
	},
	"BSOCIAL": {
		{
			ID:   10102,
			Slug: "banksocial",
			Name: "BankSocial",
		},
	},
	"GRAP": {
		{
			ID:   6664,
			Slug: "grap",
			Name: "GRAP",
		},
	},
	"MSHLD": {
		{
			ID:   9746,
			Slug: "moonshield",
			Name: "Moonshield",
		},
	},
	"VANCII": {
		{
			ID:   8671,
			Slug: "vanci-finance",
			Name: "VANCI FINANCE",
		},
	},
	"ECA": {
		{
			ID:   1711,
			Slug: "electra",
			Name: "Electra",
		},
	},
	"MKOALA": {
		{
			ID:   10137,
			Slug: "koala-token",
			Name: "KOALA TOKEN",
		},
	},
	"HUSD": {
		{
			ID:   4779,
			Slug: "husd",
			Name: "HUSD",
		},
	},
	"CHADS": {
		{
			ID:   7031,
			Slug: "chads-vc",
			Name: "CHADS VC",
		},
	},
	"OOKS": {
		{
			ID:   8349,
			Slug: "onooks",
			Name: "Onooks",
		},
	},
	"USF": {
		{
			ID:   8896,
			Slug: "unslashed-finance",
			Name: "Unslashed Finance",
		},
	},
	"ARIA20": {
		{
			ID:   8276,
			Slug: "arianee-protocol",
			Name: "Arianee",
		},
	},
	"ETNA": {
		{
			ID:   8962,
			Slug: "etna-network",
			Name: "ETNA Network",
		},
	},
	"SNO": {
		{
			ID:   3276,
			Slug: "savenode",
			Name: "SaveNode",
		},
	},
	"WITCH": {
		{
			ID:   10984,
			Slug: "witch-token",
			Name: "Witch Token",
		},
	},
	"TUSD": {
		{
			ID:   2563,
			Slug: "trueusd",
			Name: "TrueUSD",
		},
	},
	"BMC": {
		{
			ID:   1976,
			Slug: "blackmoon",
			Name: "Blackmoon",
		},
	},
	"SAFEGALAXY": {
		{
			ID:   9297,
			Slug: "safegalaxy",
			Name: "SafeGalaxy",
		},
	},
	"ESRC": {
		{
			ID:   6765,
			Slug: "esr-coin",
			Name: "ESR Coin",
		},
	},
	"CRAFT": {
		{
			ID:   7385,
			Slug: "decraft-finance",
			Name: "deCraft Finance",
		},
	},
	"SVN": {
		{
			ID:   7662,
			Slug: "7finance",
			Name: "7Finance",
		},
	},
	"GULAG": {
		{
			ID:   10923,
			Slug: "gulag-token",
			Name: "Gulag Token",
		},
	},
	"SLV": {
		{
			ID:   4079,
			Slug: "silverway",
			Name: "Silverway",
		},
	},
	"DOGS": {
		{
			ID:   10489,
			Slug: "doggy-swap",
			Name: "Doggy Swap",
		},
	},
	"RIO": {
		{
			ID:   4166,
			Slug: "realio-network",
			Name: "Realio Network",
		},
	},
	"zUSD": {
		{
			ID:   10845,
			Slug: "zerogoki-usd",
			Name: "Zerogoki USD",
		},
	},
	"KVA": {
		{
			ID:   6487,
			Slug: "kevacoin",
			Name: "Kevacoin",
		},
	},
	"$COIN": {
		{
			ID:   7796,
			Slug: "coin-defi",
			Name: "COIN",
		},
	},
	"AAPX": {
		{
			ID:   9171,
			Slug: "ampnet-asset-platform-and-exchange",
			Name: "AMPnet Asset Platform and Exchange",
		},
	},
	"HPY": {
		{
			ID:   2329,
			Slug: "hyper-pay",
			Name: "Hyper Pay",
		},
	},
	"XDB": {
		{
			ID:   4566,
			Slug: "digitalbits",
			Name: "DigitalBits",
		},
	},
	"KTN": {
		{
			ID:   9110,
			Slug: "kattana",
			Name: "Kattana",
		},
	},
	"ETG": {
		{
			ID:   2074,
			Slug: "ethereum-gold",
			Name: "Ethereum Gold",
		},
	},
	"WHEN": {
		{
			ID:   3849,
			Slug: "when-token",
			Name: "WHEN Token",
		},
	},
	"SOME": {
		{
			ID:   9599,
			Slug: "mixsome",
			Name: "Mixsome",
		},
	},
	"LIBREF": {
		{
			ID:   7752,
			Slug: "librefreelencer",
			Name: "LibreFreelencer",
		},
	},
	"POLY": {
		{
			ID:   2496,
			Slug: "polymath-network",
			Name: "Polymath",
		},
	},
	"DFND": {
		{
			ID:   9597,
			Slug: "dfund",
			Name: "dFund",
		},
	},
	"MOFI": {
		{
			ID:   9238,
			Slug: "mofi-finance",
			Name: "Mofi Finance",
		},
	},
	"PSIX": {
		{
			ID:   8575,
			Slug: "propersix",
			Name: "ProperSix",
		},
	},
	"UAXIE": {
		{
			ID:   9476,
			Slug: "unicly-mystic-axies-collection",
			Name: "Unicly Mystic Axies Collection",
		},
	},
	"YFIVE": {
		{
			ID:   6812,
			Slug: "yfive-finance",
			Name: "YFIVE FINANCE",
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
	"AENS": {
		{
			ID:   3683,
			Slug: "aen-smart-token",
			Name: "AEN Smart Token",
		},
	},
	"CLIQ": {
		{
			ID:   7815,
			Slug: "deficliq",
			Name: "DefiCliq",
		},
	},
	"KLAY": {
		{
			ID:   4256,
			Slug: "klaytn",
			Name: "Klaytn",
		},
	},
	"AQT": {
		{
			ID:   7460,
			Slug: "alpha-quark-token",
			Name: "Alpha Quark Token",
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
	"EVE": {
		{
			ID:   2464,
			Slug: "devery",
			Name: "Devery",
		},
	},
	"LIVE": {
		{
			ID:   6889,
			Slug: "tronbetlive",
			Name: "TRONbetLive",
		},
	},
	"MBL": {
		{
			ID:   4038,
			Slug: "moviebloc",
			Name: "MovieBloc",
		},
	},
	"SNTVT": {
		{
			ID:   3917,
			Slug: "sentivate",
			Name: "Sentivate",
		},
	},
	"BCDT": {
		{
			ID:   3616,
			Slug: "evidenz",
			Name: "EvidenZ",
		},
	},
	"AITRA": {
		{
			ID:   7255,
			Slug: "aitra",
			Name: "Aitra",
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
	"ELAND": {
		{
			ID:   9492,
			Slug: "etherland",
			Name: "ETHERLAND",
		},
	},
	"FOMO": {
		{
			ID:   9646,
			Slug: "fomo-lab",
			Name: "FOMO LAB",
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
	"SHEESHA": {
		{
			ID:   10337,
			Slug: "sheesha-finance-bep20",
			Name: "Sheesha Finance [BEP20]",
		},
	},
	"SHX": {
		{
			ID:   3661,
			Slug: "stronghold-token",
			Name: "Stronghold Token",
		},
	},
	"SWIRL": {
		{
			ID:   9051,
			Slug: "swirl-cash",
			Name: "Swirl Cash",
		},
	},
	"METH": {
		{
			ID:   8907,
			Slug: "farming-bad",
			Name: "Farming Bad",
		},
	},
	"DEGOV": {
		{
			ID:   7839,
			Slug: "degov",
			Name: "Degov",
		},
	},
	"FKX": {
		{
			ID:   3241,
			Slug: "fortknoxster",
			Name: "FortKnoxster",
		},
	},
	"NOW": {
		{
			ID:   3893,
			Slug: "now-token",
			Name: "ChangeNOW Token",
		},
	},
	"XMON": {
		{
			ID:   8509,
			Slug: "xmon",
			Name: "XMON",
		},
	},
	"GFX": {
		{
			ID:   9018,
			Slug: "gamyfi-platform",
			Name: "GamyFi Platform",
		},
	},
	"XAVA": {
		{
			ID:   9797,
			Slug: "avalaunch",
			Name: "Avalaunch",
		},
	},
	"SFEX": {
		{
			ID:   10434,
			Slug: "safelaunch",
			Name: "SafeX",
		},
	},
	"PSI": {
		{
			ID:   8243,
			Slug: "passive-income",
			Name: "Passive Income",
		},
	},
	"SWAP": {
		{
			ID:   5829,
			Slug: "trustswap",
			Name: "TrustSwap",
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
	"PLURA": {
		{
			ID:   3324,
			Slug: "pluracoin",
			Name: "PluraCoin",
		},
	},
	"ESTI": {
		{
			ID:   6760,
			Slug: "easticoin",
			Name: "Easticoin",
		},
	},
	"ATYNE": {
		{
			ID:   10100,
			Slug: "aerotyne",
			Name: "Aerotyne",
		},
	},
	"DEX": {
		{
			ID:   3515,
			Slug: "dex",
			Name: "DEX",
		},
	},
	"GNBT": {
		{
			ID:   10492,
			Slug: "genebank-token",
			Name: "Genebank Token",
		},
	},
	"MFI": {
		{
			ID:   8411,
			Slug: "marginswap",
			Name: "Marginswap",
		},
	},
	"OTB": {
		{
			ID:   2894,
			Slug: "otcbtc-token",
			Name: "OTCBTC Token",
		},
	},
	"TFUEL": {
		{
			ID:   3822,
			Slug: "theta-fuel",
			Name: "Theta Fuel",
		},
	},
	"NKN": {
		{
			ID:   2780,
			Slug: "nkn",
			Name: "NKN",
		},
	},
	"SUN": {
		{
			ID:   10529,
			Slug: "sun-token",
			Name: "Sun (New)",
		},
	},
	"SUTER": {
		{
			ID:   4841,
			Slug: "suterusu",
			Name: "suterusu",
		},
	},
	"FWATCH": {
		{
			ID:   10151,
			Slug: "foliowatch",
			Name: "Foliowatch",
		},
	},
	"NEWW": {
		{
			ID:   9084,
			Slug: "newv-finance",
			Name: "New Ventures",
		},
	},
	"POWR": {
		{
			ID:   2132,
			Slug: "power-ledger",
			Name: "Power Ledger",
		},
	},
	"CNNC": {
		{
			ID:   1674,
			Slug: "cannation",
			Name: "Cannation",
		},
	},
	"EOSBULL": {
		{
			ID:   5414,
			Slug: "3x-long-eos-token",
			Name: "3x Long EOS Token",
		},
	},
	"DSCPL": {
		{
			ID:   3761,
			Slug: "disciplina",
			Name: "DISCIPLINA",
		},
	},
	"FULLSEND": {
		{
			ID:   10022,
			Slug: "full-send",
			Name: "Full Send",
		},
	},
	"EDU": {
		{
			ID:   2734,
			Slug: "edu-coin",
			Name: "EduCoin",
		},
	},
	"VXV": {
		{
			ID:   4441,
			Slug: "vectorspace-ai",
			Name: "Vectorspace AI",
		},
	},
	"PAPP": {
		{
			ID:   10163,
			Slug: "papp-mobile",
			Name: "Papp Mobile",
		},
	},
	"MTL": {
		{
			ID:   1788,
			Slug: "metal",
			Name: "Metal",
		},
	},
	"RFOX": {
		{
			ID:   7654,
			Slug: "redfox-labs",
			Name: "RedFOX Labs",
		},
	},
	"DUSK": {
		{
			ID:   4092,
			Slug: "dusk-network",
			Name: "Dusk Network",
		},
	},
	"EOSC": {
		{
			ID:   4769,
			Slug: "eos-force",
			Name: "EOS Force",
		},
	},
	"IPX": {
		{
			ID:   5103,
			Slug: "tachyon-protocol",
			Name: "Tachyon Protocol",
		},
	},
	"PKF": {
		{
			ID:   8617,
			Slug: "polkafoundry",
			Name: "PolkaFoundry",
		},
	},
	"GRO": {
		{
			ID:   6718,
			Slug: "growthdefi",
			Name: "Growth DeFi",
		},
	},
	"DHV": {
		{
			ID:   8867,
			Slug: "dehive",
			Name: "DeHive",
		},
	},
	"LCP": {
		{
			ID:   331,
			Slug: "litecoin-plus",
			Name: "Litecoin Plus",
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
	"BNBBULL": {
		{
			ID:   5295,
			Slug: "3x-long-bnb-token",
			Name: "3X Long BNB Token",
		},
	},
	"APPLEB": {
		{
			ID:   10500,
			Slug: "appleb",
			Name: "APPLEB",
		},
	},
	"ARTHX": {
		{
			ID:   10796,
			Slug: "arthx",
			Name: "ARTH Shares",
		},
	},
	"ANKR": {
		{
			ID:   3783,
			Slug: "ankr",
			Name: "Ankr",
		},
	},
	"UPUNK": {
		{
			ID:   9473,
			Slug: "unicly-cryptopunks-collection",
			Name: "Unicly CryptoPunks Collection",
		},
	},
	"OPCT": {
		{
			ID:   3632,
			Slug: "opacity",
			Name: "Opacity",
		},
	},
	"PGN": {
		{
			ID:   2988,
			Slug: "pigeoncoin",
			Name: "Pigeoncoin",
		},
	},
	"PXI": {
		{
			ID:   656,
			Slug: "prime-xi",
			Name: "Prime-XI",
		},
	},
	"TABOO": {
		{
			ID:   10586,
			Slug: "taboo-token",
			Name: "TABOO TOKEN",
		},
	},
	"SAFEMARS": {
		{
			ID:   8966,
			Slug: "safemars",
			Name: "Safemars",
		},
	},
	"LESS": {
		{
			ID:   10279,
			Slug: "less-network",
			Name: "Less Network",
		},
	},
	"BU": {
		{
			ID:   3295,
			Slug: "bumo",
			Name: "BUMO",
		},
	},
	"RAMP": {
		{
			ID:   7463,
			Slug: "ramp",
			Name: "RAMP",
		},
	},
	"STC": {
		{
			ID:   5966,
			Slug: "student-coin",
			Name: "Student Coin",
		},
	},
	"LMT": {
		{
			ID:   9482,
			Slug: "lympo-market-token",
			Name: "Lympo Market Token",
		},
	},
	"JADE": {
		{
			ID:   5025,
			Slug: "jade-currency",
			Name: "Jade Currency",
		},
	},
	"8PAY": {
		{
			ID:   9046,
			Slug: "8pay",
			Name: "8PAY",
		},
	},
	"WANUSDC": {
		{
			ID:   8655,
			Slug: "wanusdc",
			Name: "wanUSDC",
		},
	},
	"ETH20SMACO": {
		{
			ID:   6129,
			Slug: "eth-20-day-ma-crossover-set",
			Name: "ETH 20 Day MA Crossover Set",
		},
	},
	"EFG": {
		{
			ID:   7872,
			Slug: "ecoc-financial-growth",
			Name: "ECOC Financial Growth",
		},
	},
	"N3RDz": {
		{
			ID:   8019,
			Slug: "n3rd-finance",
			Name: "N3RD Finance",
		},
	},
	"BRZE": {
		{
			ID:   3519,
			Slug: "breezecoin",
			Name: "Breezecoin",
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
	"RLE": {
		{
			ID:   7990,
			Slug: "richlab-token",
			Name: "Richlab Token",
		},
	},
	"PDF": {
		{
			ID:   6737,
			Slug: "port-of-defi-network",
			Name: "Port of DeFi Network",
		},
	},
	"VIT": {
		{
			ID:   11015,
			Slug: "team-vitality-fan-token",
			Name: "Team Vitality Fan Token",
		},
	},
	"DACS": {
		{
			ID:   3054,
			Slug: "dacsee",
			Name: "DACSEE",
		},
	},
	"UBTC": {
		{
			ID:   2293,
			Slug: "united-bitcoin",
			Name: "United Bitcoin",
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
	"ALH": {
		{
			ID:   8674,
			Slug: "allohash",
			Name: "AlloHash",
		},
	},
	"BBQ": {
		{
			ID:   10443,
			Slug: "barbecueswap-finance",
			Name: "BarbecueSwap Finance",
		},
	},
	"BWC": {
		{
			ID:   10306,
			Slug: "bongweedcoin",
			Name: "BongWeedCoin",
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
	"ROC": {
		{
			ID:   8768,
			Slug: "roxe-cash",
			Name: "Roxe Cash",
		},
	},
	"CLU": {
		{
			ID:   9984,
			Slug: "clucoin",
			Name: "CluCoin",
		},
	},
	"VBNT": {
		{
			ID:   8622,
			Slug: "bancor-governance-token",
			Name: "Bancor Governance Token",
		},
	},
	"ETNXP": {
		{
			ID:   5669,
			Slug: "electronero-pulse",
			Name: "Electronero Pulse",
		},
	},
	"ORAO": {
		{
			ID:   8895,
			Slug: "orao-network",
			Name: "ORAO Network",
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
	"RUPEE": {
		{
			ID:   8736,
			Slug: "hyruleswap",
			Name: "HyruleSwap",
		},
	},
	"PANDO": {
		{
			ID:   8711,
			Slug: "pando",
			Name: "Pando",
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
	"RATING": {
		{
			ID:   3194,
			Slug: "dprating",
			Name: "DPRating",
		},
	},
	"FRST": {
		{
			ID:   1522,
			Slug: "firstcoin",
			Name: "FirstCoin",
		},
	},
	"XFUND": {
		{
			ID:   8339,
			Slug: "xfund",
			Name: "xFund",
		},
	},
	"BUILD": {
		{
			ID:   6956,
			Slug: "build-finance",
			Name: "BUILD Finance",
		},
	},
	"YEA": {
		{
			ID:   7317,
			Slug: "yeafinance",
			Name: "YeaFinance",
		},
	},
	"ZUR": {
		{
			ID:   1250,
			Slug: "zurcoin",
			Name: "Zurcoin",
		},
	},
	"UFR": {
		{
			ID:   2178,
			Slug: "upfiring",
			Name: "Upfiring",
		},
	},
	"RENZEC": {
		{
			ID:   8904,
			Slug: "renzec",
			Name: "renZEC",
		},
	},
	"MIN": {
		{
			ID:   3296,
			Slug: "mindol",
			Name: "MINDOL",
		},
	},
	"DFGL": {
		{
			ID:   7483,
			Slug: "defi-gold",
			Name: "DeFi Gold",
		},
	},
	"VERS": {
		{
			ID:   4750,
			Slug: "versess-coin",
			Name: "Versess Coin",
		},
	},
	"MEETONE": {
		{
			ID:   3136,
			Slug: "meetone",
			Name: "MEET.ONE",
		},
	},
	"NPC": {
		{
			ID:   4110,
			Slug: "npcoin",
			Name: "NPCoin",
		},
	},
	"SWG": {
		{
			ID:   7467,
			Slug: "swirge",
			Name: "Swirge",
		},
	},
	"NBS": {
		{
			ID:   7110,
			Slug: "new-bitshares",
			Name: "New BitShares",
		},
	},
	"DFD": {
		{
			ID:   7593,
			Slug: "defidollar-dao",
			Name: "DefiDollar DAO",
		},
	},
	"EYE": {
		{
			ID:   7414,
			Slug: "beholder",
			Name: "Behodler",
		},
	},
	"NGC": {
		{
			ID:   2305,
			Slug: "naga",
			Name: "NAGA",
		},
	},
	"FLUSD": {
		{
			ID:   10390,
			Slug: "fluity-usd",
			Name: "Fluity USD",
		},
	},
	"DAI": {
		{
			ID:   4943,
			Slug: "multi-collateral-dai",
			Name: "Dai",
		},
	},
	"TBT": {
		{
			ID:   6565,
			Slug: "tidebit-token",
			Name: "TideBit Token",
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
	"VIN": {
		{
			ID:   3082,
			Slug: "vinchain",
			Name: "VINchain",
		},
	},
	"AAB": {
		{
			ID:   5509,
			Slug: "aax-token",
			Name: "AAX Token",
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
	"COIL": {
		{
			ID:   6645,
			Slug: "coil",
			Name: "COIL",
		},
	},
	"WAIV": {
		{
			ID:   10470,
			Slug: "waivlength",
			Name: "Waivlength",
		},
	},
	"LYM": {
		{
			ID:   2554,
			Slug: "lympo",
			Name: "Lympo",
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
	"UMASK": {
		{
			ID:   9477,
			Slug: "unicly-hashmasks-collection",
			Name: "Unicly Hashmasks Collection",
		},
	},
	"HVN": {
		{
			ID:   1950,
			Slug: "hiveterminal-token",
			Name: "Hiveterminal Token",
		},
	},
	"DEB": {
		{
			ID:   2584,
			Slug: "debitum-network",
			Name: "Debitum",
		},
	},
	"LOCG": {
		{
			ID:   9526,
			Slug: "locgame",
			Name: "LOCGame",
		},
	},
	"SXPUP": {
		{
			ID:   7528,
			Slug: "sxpup",
			Name: "SXPUP",
		},
	},
	"DOGEFATHER": {
		{
			ID:   9732,
			Slug: "dogefather",
			Name: "Dogefather",
		},
	},
	"DEFI+L": {
		{
			ID:   7337,
			Slug: "piedao-defi-large-cap",
			Name: "PieDAO DEFI Large Cap",
		},
	},
	"MARSM": {
		{
			ID:   9293,
			Slug: "marsmission-protocol",
			Name: "MarsMission Protocol",
		},
	},
	"BUB": {
		{
			ID:   918,
			Slug: "bubble",
			Name: "Bubble",
		},
	},
	"PERP": {
		{
			ID:   6950,
			Slug: "perpetual-protocol",
			Name: "Perpetual Protocol",
		},
	},
	"UMB": {
		{
			ID:   8385,
			Slug: "umbrella-network",
			Name: "Umbrella Network",
		},
	},
	"VIPS": {
		{
			ID:   2688,
			Slug: "vipstar-coin",
			Name: "Vipstar Coin",
		},
	},
	"PGT": {
		{
			ID:   7315,
			Slug: "polyient-games-governance-token",
			Name: "Polyient Games Governance Token",
		},
	},
	"TAG": {
		{
			ID:   61,
			Slug: "tagcoin",
			Name: "TagCoin",
		},
	},
	"BITT": {
		{
			ID:   8517,
			Slug: "bittoken",
			Name: "BiTToken",
		},
	},
	"XVIX": {
		{
			ID:   7969,
			Slug: "xvix",
			Name: "XVIX",
		},
	},
	"BabyDoge": {
		{
			ID:   10407,
			Slug: "baby-doge-coin",
			Name: "Baby Doge Coin",
		},
	},
	"PIG": {
		{
			ID:   8829,
			Slug: "pig-finance",
			Name: "Pig Finance",
		},
	},
	"OCB": {
		{
			ID:   6869,
			Slug: "blockmax",
			Name: "BLOCKMAX",
		},
	},
	"YFIH2": {
		{
			ID:   10545,
			Slug: "h2finance",
			Name: "H2Finance",
		},
	},
	"BALI": {
		{
			ID:   5480,
			Slug: "bali-coin",
			Name: "Bali Coin",
		},
	},
	"BTYM": {
		{
			ID:   9655,
			Slug: "blocktyme",
			Name: "Blocktyme",
		},
	},
	"MDTK": {
		{
			ID:   4786,
			Slug: "mdtoken",
			Name: "MDtoken",
		},
	},
	"BCUBE": {
		{
			ID:   9553,
			Slug: "b-cube-ai",
			Name: "B-cube.ai",
		},
	},
	"DCR": {
		{
			ID:   1168,
			Slug: "decred",
			Name: "Decred",
		},
	},
	"CDT": {
		{
			ID:   1864,
			Slug: "blox",
			Name: "Blox",
		},
	},
	"OPIUM": {
		{
			ID:   7230,
			Slug: "opium",
			Name: "Opium",
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
	"DRCT": {
		{
			ID:   10385,
			Slug: "ally-direct-token",
			Name: "Ally Direct Token",
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
	"PROUD": {
		{
			ID:   1398,
			Slug: "proud-money",
			Name: "PROUD Money",
		},
	},
	"KSS": {
		{
			ID:   6655,
			Slug: "krosscoin",
			Name: "Krosscoin",
		},
	},
	"SAUBER": {
		{
			ID:   10793,
			Slug: "alfa-romeo-racing-orlen-fan-token",
			Name: "Alfa Romeo Racing ORLEN Fan Token",
		},
	},
	"OG": {
		{
			ID:   5309,
			Slug: "og-fan-token",
			Name: "OG Fan Token",
		},
	},
	"PLOT": {
		{
			ID:   7422,
			Slug: "plotx",
			Name: "PlotX",
		},
	},
	"IHT": {
		{
			ID:   2552,
			Slug: "iht-real-estate-protocol",
			Name: "IHT Real Estate Protocol",
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
	"SHIELD": {
		{
			ID:   8452,
			Slug: "shield-protocol",
			Name: "Shield Protocol",
		},
	},
	"STBZ": {
		{
			ID:   7366,
			Slug: "stabilize",
			Name: "Stabilize",
		},
	},
	"JSB": {
		{
			ID:   7787,
			Slug: "jsb-foundation",
			Name: "JSB FOUNDATION",
		},
	},
	"OWL": {
		{
			ID:   7330,
			Slug: "owl-token-stealthswap",
			Name: "OWL Token (StealthSwap)",
		},
	},
	"ZIOT": {
		{
			ID:   9226,
			Slug: "ziot-coin",
			Name: "ziot Coin",
		},
	},
	"WNXM": {
		{
			ID:   5939,
			Slug: "wrapped-nxm",
			Name: "Wrapped NXM",
		},
	},
	"CND": {
		{
			ID:   2043,
			Slug: "cindicator",
			Name: "Cindicator",
		},
	},
	"ADI": {
		{
			ID:   2660,
			Slug: "aditus",
			Name: "Aditus",
		},
	},
	"IMPL": {
		{
			ID:   3665,
			Slug: "impleum",
			Name: "Impleum",
		},
	},
	"FROGE": {
		{
			ID:   9234,
			Slug: "froge-finance",
			Name: "Froge Finance",
		},
	},
	"FWT": {
		{
			ID:   7585,
			Slug: "freeway-token",
			Name: "Freeway Token",
		},
	},
	"CAI": {
		{
			ID:   7639,
			Slug: "club-atletico-independiente",
			Name: "Club Atletico Independiente",
		},
	},
	"TWIN": {
		{
			ID:   9253,
			Slug: "twinci",
			Name: "Twinci",
		},
	},
	"DCTO": {
		{
			ID:   3738,
			Slug: "decentralized-crypto-token",
			Name: "Decentralized Crypto Token",
		},
	},
	"CAB": {
		{
			ID:   1210,
			Slug: "cabbage",
			Name: "Cabbage",
		},
	},
	"MKAT": {
		{
			ID:   10691,
			Slug: "moonkat",
			Name: "MoonKat",
		},
	},
	"SPANDA": {
		{
			ID:   10076,
			Slug: "superpanda",
			Name: "Superpanda",
		},
	},
	"SIERRA": {
		{
			ID:   4630,
			Slug: "sierracoin",
			Name: "Sierracoin",
		},
	},
	"BF": {
		{
			ID:   4283,
			Slug: "bitforex-token",
			Name: "BitForex Token",
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
	"ALPACA": {
		{
			ID:   8707,
			Slug: "alpaca-finance",
			Name: "Alpaca Finance",
		},
	},
	"CHEESE": {
		{
			ID:   3464,
			Slug: "cheesecoin",
			Name: "Cheesecoin",
		},
	},
	"RPT": {
		{
			ID:   8413,
			Slug: "rug-proof",
			Name: "Rug Proof",
		},
	},
	"UNC": {
		{
			ID:   5948,
			Slug: "unicrypt",
			Name: "Unicrypt",
		},
	},
	"XRC": {
		{
			ID:   3839,
			Slug: "xrhodium",
			Name: "xRhodium",
		},
	},
	"FILDOWN": {
		{
			ID:   8051,
			Slug: "fildown",
			Name: "FILDOWN",
		},
	},
	"TENSHI": {
		{
			ID:   10957,
			Slug: "tenshi",
			Name: "Tenshi",
		},
	},
	"CLAM": {
		{
			ID:   460,
			Slug: "clams",
			Name: "Clams",
		},
	},
	"BASI": {
		{
			ID:   8824,
			Slug: "asi-finance",
			Name: "ASI.finance",
		},
	},
	"API3": {
		{
			ID:   7737,
			Slug: "api3",
			Name: "API3",
		},
	},
	"CYCLUB": {
		{
			ID:   6810,
			Slug: "cyclub",
			Name: "CYCLUB",
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
	"HAVY": {
		{
			ID:   3265,
			Slug: "havy",
			Name: "Havy",
		},
	},
	"JORDAN": {
		{
			ID:   10915,
			Slug: "cryingjordan-token",
			Name: "cryingJORDAN Token",
		},
	},
	"BLAZR": {
		{
			ID:   1623,
			Slug: "blazercoin",
			Name: "BlazerCoin",
		},
	},
	"ICP": {
		{
			ID:   8916,
			Slug: "internet-computer",
			Name: "Internet Computer",
		},
	},
	"OXY": {
		{
			ID:   8029,
			Slug: "oxygen",
			Name: "Oxygen",
		},
	},
	"MX": {
		{
			ID:   4041,
			Slug: "mx-token",
			Name: "MX Token",
		},
	},
	"BANCA": {
		{
			ID:   2592,
			Slug: "banca",
			Name: "Banca",
		},
	},
	"CREA": {
		{
			ID:   1676,
			Slug: "crea",
			Name: "CREA",
		},
	},
	"LQD": {
		{
			ID:   3499,
			Slug: "liquidity-network",
			Name: "Liquidity Network",
		},
	},
	"SIMPLE": {
		{
			ID:   5319,
			Slug: "simplechain",
			Name: "SimpleChain",
		},
	},
	"PETB": {
		{
			ID:   10152,
			Slug: "petbloc",
			Name: "PETBloc",
		},
	},
	"mUSO": {
		{
			ID:   8027,
			Slug: "mirrored-united-states-oil-fund",
			Name: "Mirrored United States Oil Fund",
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
	"SHIBARISE": {
		{
			ID:   10855,
			Slug: "shiba-rise",
			Name: "SHIBA RISE",
		},
	},
	"DASH": {
		{
			ID:   131,
			Slug: "dash",
			Name: "Dash",
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
	"BSDS": {
		{
			ID:   8237,
			Slug: "basis-dollar-share",
			Name: "Basis Dollar Share",
		},
	},
	"XLPG": {
		{
			ID:   5362,
			Slug: "stellarpayglobal",
			Name: "StellarPayGlobal",
		},
	},
	"PEPE": {
		{
			ID:   10268,
			Slug: "pepemoon",
			Name: "PepeMoon",
		},
	},
	"ORN": {
		{
			ID:   5631,
			Slug: "orion-protocol",
			Name: "Orion Protocol",
		},
	},
	"NFTX": {
		{
			ID:   8191,
			Slug: "nftx",
			Name: "NFTX",
		},
	},
	"ZEE": {
		{
			ID:   7438,
			Slug: "zeroswap",
			Name: "ZeroSwap",
		},
	},
	"OWC": {
		{
			ID:   3763,
			Slug: "oduwa",
			Name: "ODUWA",
		},
	},
	"BPLC": {
		{
			ID:   6113,
			Slug: "blackpearl-chain",
			Name: "BlackPearl Token",
		},
	},
	"STARSHIP": {
		{
			ID:   9962,
			Slug: "starship",
			Name: "STARSHIP",
		},
	},
	"TTC": {
		{
			ID:   8359,
			Slug: "ttcrypto",
			Name: "TTCRYPTO",
		},
	},
	"IRIS": {
		{
			ID:   3874,
			Slug: "irisnet",
			Name: "IRISnet",
		},
	},
	"TNB": {
		{
			ID:   2235,
			Slug: "time-new-bank",
			Name: "Time New Bank",
		},
	},
	"DMST": {
		{
			ID:   5952,
			Slug: "dmscript",
			Name: "DMScript",
		},
	},
	"WINR": {
		{
			ID:   7494,
			Slug: "justbet",
			Name: "JustBet",
		},
	},
	"VSL": {
		{
			ID:   1483,
			Slug: "vslice",
			Name: "vSlice",
		},
	},
	"BITC": {
		{
			ID:   3907,
			Slug: "bitcash",
			Name: "BitCash",
		},
	},
	"DEXF": {
		{
			ID:   11027,
			Slug: "dexfolio",
			Name: "Dexfolio",
		},
	},
	"YFX": {
		{
			ID:   9946,
			Slug: "your-future-exchange",
			Name: "Your Future Exchange",
		},
	},
	"PAWS": {
		{
			ID:   9573,
			Slug: "animal-adoption-advocacy",
			Name: "Animal Adoption Advocacy",
		},
	},
	"SPND": {
		{
			ID:   10369,
			Slug: "safepanda",
			Name: "SafePanda",
		},
	},
	"WHL": {
		{
			ID:   8560,
			Slug: "whaleroom",
			Name: "WhaleRoom",
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
	"WGP": {
		{
			ID:   3912,
			Slug: "w-green-pay",
			Name: "W Green Pay",
		},
	},
	"FSCC": {
		{
			ID:   6447,
			Slug: "fisco",
			Name: "Fisco Coin",
		},
	},
	"MAKI": {
		{
			ID:   10232,
			Slug: "makiswap",
			Name: "MakiSwap",
		},
	},
	"CL": {
		{
			ID:   2352,
			Slug: "coinlancer",
			Name: "Coinlancer",
		},
	},
	"GLEEC": {
		{
			ID:   5200,
			Slug: "gleec",
			Name: "Gleec",
		},
	},
	"KRB": {
		{
			ID:   1340,
			Slug: "karbo",
			Name: "Karbo",
		},
	},
	"AWF": {
		{
			ID:   10043,
			Slug: "alphawolf-finance",
			Name: "Alphawolf Finance",
		},
	},
	"NORD": {
		{
			ID:   8143,
			Slug: "nord-finance",
			Name: "Nord Finance",
		},
	},
	"EXCL": {
		{
			ID:   633,
			Slug: "exclusivecoin",
			Name: "ExclusiveCoin",
		},
	},
	"ELLA": {
		{
			ID:   2122,
			Slug: "ellaism",
			Name: "Ellaism",
		},
	},
	"BUSY": {
		{
			ID:   9002,
			Slug: "busy",
			Name: "Busy DAO",
		},
	},
	"DEFX": {
		{
			ID:   10145,
			Slug: "definity",
			Name: "DeFinity",
		},
	},
	"GOFI": {
		{
			ID:   8873,
			Slug: "goswapp",
			Name: "GoSwapp",
		},
	},
	"MAVRO": {
		{
			ID:   1599,
			Slug: "mavro",
			Name: "Mavro",
		},
	},
	"DODO": {
		{
			ID:   7224,
			Slug: "dodo",
			Name: "DODO",
		},
	},
	"VDL": {
		{
			ID:   4315,
			Slug: "vidulum",
			Name: "Vidulum",
		},
	},
	"EGCC": {
		{
			ID:   2852,
			Slug: "engine",
			Name: "Engine",
		},
	},
	"WHIRL": {
		{
			ID:   8778,
			Slug: "whirl-finance",
			Name: "Whirl Finance",
		},
	},
	"DOV": {
		{
			ID:   2110,
			Slug: "dovu",
			Name: "Dovu",
		},
	},
	"BLOCK": {
		{
			ID:   707,
			Slug: "blocknet",
			Name: "Blocknet",
		},
	},
	"MCONTENT": {
		{
			ID:   10954,
			Slug: "mcontent",
			Name: "MContent",
		},
	},
	"$MAD": {
		{
			ID:   10215,
			Slug: "make-a-difference-token",
			Name: "Make A Difference Token",
		},
	},
	"XEM": {
		{
			ID:   873,
			Slug: "nem",
			Name: "NEM",
		},
	},
	"GAL": {
		{
			ID:   5228,
			Slug: "galatasaray-fan-token",
			Name: "Galatasaray Fan Token",
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
	"TUNE": {
		{
			ID:   6330,
			Slug: "tune-token",
			Name: "TUNE TOKEN",
		},
	},
	"CTRFI": {
		{
			ID:   10038,
			Slug: "chester-moon",
			Name: "Chester.Moon",
		},
	},
	"ADA": {
		{
			ID:   2010,
			Slug: "cardano",
			Name: "Cardano",
		},
	},
	"YFI": {
		{
			ID:   5864,
			Slug: "yearn-finance",
			Name: "yearn.finance",
		},
	},
	"RISEUP": {
		{
			ID:   10679,
			Slug: "riseup",
			Name: "RiseUp",
		},
	},
	"ADMONKEY": {
		{
			ID:   11012,
			Slug: "admonkey",
			Name: "AdMonkey",
		},
	},
	"CLA": {
		{
			ID:   6978,
			Slug: "candelacoin",
			Name: "Candela Coin",
		},
	},
	"BDCASH": {
		{
			ID:   7277,
			Slug: "bigdatacash",
			Name: "BigdataCash",
		},
	},
	"SCSX": {
		{
			ID:   4487,
			Slug: "secure-cash",
			Name: "Secure Cash",
		},
	},
	"RANK": {
		{
			ID:   7947,
			Slug: "rank-token",
			Name: "Rank Token",
		},
	},
	"CBIX": {
		{
			ID:   4142,
			Slug: "cubiex",
			Name: "Cubiex",
		},
	},
	"MRCH": {
		{
			ID:   8971,
			Slug: "merchdao",
			Name: "MerchDAO",
		},
	},
	"MOCHI": {
		{
			ID:   8881,
			Slug: "mochiswap",
			Name: "MOCHISWAP",
		},
	},
	"SNC": {
		{
			ID:   1786,
			Slug: "suncontract",
			Name: "SunContract",
		},
	},
	"WET": {
		{
			ID:   3585,
			Slug: "weshow-token",
			Name: "WeShow Token",
		},
	},
	"PFID": {
		{
			ID:   7043,
			Slug: "pofid-dao",
			Name: "Pofid Dao",
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
	"STAX": {
		{
			ID:   7502,
			Slug: "stablexswap",
			Name: "StableXSwap",
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
	"OPAL": {
		{
			ID:   597,
			Slug: "opal",
			Name: "Opal",
		},
	},
	"ETH3S": {
		{
			ID:   5738,
			Slug: "amun-ether-3x-daily-short",
			Name: "Amun Ether 3x Daily Short",
		},
	},
	"SLW": {
		{
			ID:   10724,
			Slug: "solarwind-token",
			Name: "SolarWind Token",
		},
	},
	"BCHA": {
		{
			ID:   7686,
			Slug: "bitcoin-cash-abc-2",
			Name: "Bitcoin Cash ABC",
		},
	},
	"LON": {
		{
			ID:   8083,
			Slug: "tokenlon-network-token",
			Name: "Tokenlon Network Token",
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
	"SHREW": {
		{
			ID:   11010,
			Slug: "shrew",
			Name: "Shrew",
		},
	},
	"XRPBEAR": {
		{
			ID:   5413,
			Slug: "3x-short-xrp-token",
			Name: "3x Short XRP Token",
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
	"TRYON": {
		{
			ID:   10341,
			Slug: "stellar-invictus-gaming",
			Name: "Stellar Invictus Gaming",
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
	"CKB": {
		{
			ID:   4948,
			Slug: "nervos-network",
			Name: "Nervos Network",
		},
	},
	"EWT": {
		{
			ID:   5268,
			Slug: "energy-web-token",
			Name: "Energy Web Token",
		},
	},
	"FRAX": {
		{
			ID:   6952,
			Slug: "frax",
			Name: "Frax",
		},
	},
	"NEW": {
		{
			ID:   3871,
			Slug: "newton",
			Name: "Newton",
		},
	},
	"EDRC": {
		{
			ID:   1216,
			Slug: "edrcoin",
			Name: "EDRCoin",
		},
	},
	"WAULTX": {
		{
			ID:   9478,
			Slug: "wault-finance-new",
			Name: "Wault [New]",
		},
	},
	"RichDoge": {
		{
			ID:   10989,
			Slug: "rich-doge-coin",
			Name: "Rich Doge Coin",
		},
	},
	"RIF": {
		{
			ID:   3701,
			Slug: "rsk-infrastructure-framework",
			Name: "RSK Infrastructure Framework",
		},
	},
	"LTMS": {
		{
			ID:   10190,
			Slug: "littlemouse",
			Name: "LittleMouse",
		},
	},
	"IOOX": {
		{
			ID:   5733,
			Slug: "ioox-system",
			Name: "IOOX System",
		},
	},
	"BPC": {
		{
			ID:   5641,
			Slug: "backpacker-coin",
			Name: "BackPacker Coin",
		},
	},
	"SMLY": {
		{
			ID:   799,
			Slug: "smileycoin",
			Name: "SmileyCoin",
		},
	},
	"XXA": {
		{
			ID:   5474,
			Slug: "ixinium",
			Name: "Ixinium",
		},
	},
	"IPM": {
		{
			ID:   7247,
			Slug: "timers",
			Name: "Timers",
		},
	},
	"LUSD": {
		{
			ID:   9566,
			Slug: "liquity-usd",
			Name: "Liquity USD",
		},
	},
	"KT": {
		{
			ID:   3691,
			Slug: "kuai-token",
			Name: "Kuai Token",
		},
	},
	"OVR": {
		{
			ID:   8144,
			Slug: "ovr",
			Name: "OVR",
		},
	},
	"ULTGG": {
		{
			ID:   10881,
			Slug: "ultimogg",
			Name: "UltimoGG",
		},
	},
	"NFTP": {
		{
			ID:   8828,
			Slug: "nft-pool",
			Name: "NFT POOL",
		},
	},
	"NPXSXEM": {
		{
			ID:   3096,
			Slug: "pundi-x-nem",
			Name: "Pundi X NEM",
		},
	},
	"VEC2": {
		{
			ID:   1052,
			Slug: "vector",
			Name: "VectorAI",
		},
	},
	"BLP": {
		{
			ID:   10326,
			Slug: "bullperks",
			Name: "BullPerks",
		},
	},
	"BUNI": {
		{
			ID:   9906,
			Slug: "bunicorn",
			Name: "Bunicorn",
		},
	},
	"BNBTC": {
		{
			ID:   10647,
			Slug: "bnbitcoin",
			Name: "BNbitcoin",
		},
	},
	"AV": {
		{
			ID:   1146,
			Slug: "avatarcoin",
			Name: "AvatarCoin",
		},
	},
	"DAG": {
		{
			ID:   2868,
			Slug: "constellation",
			Name: "Constellation",
		},
	},
	"KRT": {
		{
			ID:   5115,
			Slug: "terra-krw",
			Name: "TerraKRW",
		},
	},
	"FLC": {
		{
			ID:   3727,
			Slug: "flowchain",
			Name: "Flowchain",
		},
	},
	"OVI": {
		{
			ID:   9888,
			Slug: "oviex",
			Name: "Oviex",
		},
	},
	"BOTS": {
		{
			ID:   8454,
			Slug: "botocean",
			Name: "BotOcean",
		},
	},
	"KEK": {
		{
			ID:   8135,
			Slug: "cryptokek",
			Name: "CryptoKek",
		},
	},
	"SPANK": {
		{
			ID:   2219,
			Slug: "spankchain",
			Name: "SpankChain",
		},
	},
	"HEDG": {
		{
			ID:   3662,
			Slug: "hedgetrade",
			Name: "HedgeTrade",
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
	"CATGE": {
		{
			ID:   10274,
			Slug: "catge-coin",
			Name: "Catge coin",
		},
	},
	"D4RK": {
		{
			ID:   3516,
			Slug: "dark",
			Name: "Dark",
		},
	},
	"SAND": {
		{
			ID:   6210,
			Slug: "the-sandbox",
			Name: "The Sandbox",
		},
	},
	"POL": {
		{
			ID:   6297,
			Slug: "proof-of-liquidity",
			Name: "Proof Of Liquidity",
		},
	},
	"SFR": {
		{
			ID:   7205,
			Slug: "safari",
			Name: "Safari",
		},
	},
	"$KEI": {
		{
			ID:   10048,
			Slug: "keisuke-inu",
			Name: "Keisuke Inu",
		},
	},
	"EOS": {
		{
			ID:   1765,
			Slug: "eos",
			Name: "EOS",
		},
	},
	"BC": {
		{
			ID:   3976,
			Slug: "bitcoin-confidential",
			Name: "Bitcoin Confidential",
		},
	},
	"HSC": {
		{
			ID:   2908,
			Slug: "hashcoin",
			Name: "HashCoin",
		},
	},
	"SPR": {
		{
			ID:   702,
			Slug: "spreadcoin",
			Name: "SpreadCoin",
		},
	},
	"RWS": {
		{
			ID:   5936,
			Slug: "robonomics-web-services",
			Name: "Robonomics Web Services",
		},
	},
	"WNL": {
		{
			ID:   4138,
			Slug: "winstars-live",
			Name: "WinStars.live",
		},
	},
	"ROT": {
		{
			ID:   7164,
			Slug: "rotten",
			Name: "Rotten",
		},
	},
	"SVX": {
		{
			ID:   9109,
			Slug: "savix",
			Name: "Savix",
		},
	},
	"PZM": {
		{
			ID:   1681,
			Slug: "prizm",
			Name: "PRIZM",
		},
	},
	"DF": {
		{
			ID:   4758,
			Slug: "dforce",
			Name: "dForce",
		},
	},
	"SAITO": {
		{
			ID:   9194,
			Slug: "saito",
			Name: "Saito",
		},
	},
	"RUFF": {
		{
			ID:   2476,
			Slug: "ruff",
			Name: "Ruff",
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
	"DOGES": {
		{
			ID:   7377,
			Slug: "dogeswap",
			Name: "Dogeswap",
		},
	},
	"PRIA": {
		{
			ID:   7466,
			Slug: "pria",
			Name: "PRIA",
		},
	},
	"TBAKE": {
		{
			ID:   9796,
			Slug: "bakery-tools",
			Name: "Bakery Tools",
		},
	},
	"DISC": {
		{
			ID:   9582,
			Slug: "discas-vision",
			Name: "DisCas Vision",
		},
	},
	"APACHE": {
		{
			ID:   10094,
			Slug: "apache",
			Name: "Apache",
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
	"ETH3L": {
		{
			ID:   5739,
			Slug: "amun-ether-3x-daily-long",
			Name: "Amun Ether 3x Daily Long",
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
	"AGOV": {
		{
			ID:   7673,
			Slug: "agov-answer-governance",
			Name: "AGOV (ANSWER Governance)",
		},
	},
	"DOTX": {
		{
			ID:   6796,
			Slug: "deli-of-thrones",
			Name: "DeFi of Thrones",
		},
	},
	"BURNX20": {
		{
			ID:   10072,
			Slug: "burnx20",
			Name: "BurnX 2.0",
		},
	},
	"MRCR": {
		{
			ID:   10170,
			Slug: "mercor-finance",
			Name: "Mercor Finance",
		},
	},
	"MILK": {
		{
			ID:   9141,
			Slug: "milk-token",
			Name: "Milk Token",
		},
	},
	"HAPPY": {
		{
			ID:   9824,
			Slug: "thehappycoin",
			Name: "HappyCoin",
		},
	},
	"JIND": {
		{
			ID:   10127,
			Slug: "jindo-inu",
			Name: "JINDO INU",
		},
	},
	"FERA": {
		{
			ID:   6612,
			Slug: "fera",
			Name: "Fera",
		},
	},
	"SSP": {
		{
			ID:   2862,
			Slug: "smartshare",
			Name: "Smartshare",
		},
	},
	"NAME": {
		{
			ID:   9264,
			Slug: "polkadomain",
			Name: "PolkaDomain",
		},
	},
	"WAXE": {
		{
			ID:   8136,
			Slug: "waxe",
			Name: "WAXE",
		},
	},
	"ODIN": {
		{
			ID:   9546,
			Slug: "odin-protocol",
			Name: "ODIN PROTOCOL",
		},
	},
	"MOVE": {
		{
			ID:   7371,
			Slug: "mover",
			Name: "Mover",
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
	"SBF": {
		{
			ID:   9953,
			Slug: "steakbankfinance",
			Name: "SteakBankFinance",
		},
	},
	"AFEN": {
		{
			ID:   9752,
			Slug: "afen-blockchain",
			Name: "AFEN Blockchain",
		},
	},
	"MORK": {
		{
			ID:   6692,
			Slug: "mork",
			Name: "MORK",
		},
	},
	"B2X": {
		{
			ID:   10475,
			Slug: "b2x",
			Name: "B2X",
		},
	},
	"DTX": {
		{
			ID:   2913,
			Slug: "databroker",
			Name: "Databroker",
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
	"CARE": {
		{
			ID:   3266,
			Slug: "carebit",
			Name: "Carebit",
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
	"LIGHT": {
		{
			ID:   8801,
			Slug: "lightning",
			Name: "Lightning",
		},
	},
	"DXF": {
		{
			ID:   8079,
			Slug: "dexfin",
			Name: "Dexfin",
		},
	},
	"FFYI": {
		{
			ID:   6935,
			Slug: "fiscus-fyi",
			Name: "Fiscus.fyi",
		},
	},
	"DFIO": {
		{
			ID:   6672,
			Slug: "defi-omega",
			Name: "DeFi Omega",
		},
	},
	"ETC": {
		{
			ID:   1321,
			Slug: "ethereum-classic",
			Name: "Ethereum Classic",
		},
	},
	"NTX": {
		{
			ID:   8500,
			Slug: "nitroex",
			Name: "Nitroex",
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
	"ANI": {
		{
			ID:   8394,
			Slug: "anime-token",
			Name: "Anime Token",
		},
	},
	"MRS": {
		{
			ID:   8634,
			Slug: "marsan-exchange-token",
			Name: "Marsan Exchange token",
		},
	},
	"KSC": {
		{
			ID:   6493,
			Slug: "kstarcoin",
			Name: "KStarCoin",
		},
	},
	"FAIRLIFE": {
		{
			ID:   10546,
			Slug: "fairlife",
			Name: "FAIRLIFE",
		},
	},
	"HMR": {
		{
			ID:   5336,
			Slug: "homeros",
			Name: "Homeros",
		},
	},
	"FTC": {
		{
			ID:   8,
			Slug: "feathercoin",
			Name: "Feathercoin",
		},
	},
	"STPL": {
		{
			ID:   7322,
			Slug: "stream-protocol",
			Name: "Stream Protocol",
		},
	},
	"TRDG": {
		{
			ID:   10785,
			Slug: "tardigrades-finance-eth",
			Name: "Tardigrades.Finance (ETH)",
		},
	},
	"CRYPT": {
		{
			ID:   10135,
			Slug: "cryptonaught",
			Name: "Cryptonaught",
		},
	},
	"MDA": {
		{
			ID:   1954,
			Slug: "moeda-loyalty-points",
			Name: "Moeda Loyalty Points",
		},
	},
	"NOKU": {
		{
			ID:   3138,
			Slug: "noku",
			Name: "Noku",
		},
	},
	"PPBLZ": {
		{
			ID:   7435,
			Slug: "pepemon-pepeballs",
			Name: "Pepemon Pepeballs",
		},
	},
	"STARL": {
		{
			ID:   10821,
			Slug: "star-link",
			Name: "StarLink",
		},
	},
	"ETHHEDGE": {
		{
			ID:   5658,
			Slug: "1x-short-ethereum-token",
			Name: "1X Short Ethereum Token",
		},
	},
	"TAXI": {
		{
			ID:   8070,
			Slug: "taxi",
			Name: "Taxi",
		},
	},
	"ALCX": {
		{
			ID:   8613,
			Slug: "alchemix",
			Name: "Alchemix",
		},
	},
	"CHNG": {
		{
			ID:   9071,
			Slug: "chainge",
			Name: "Chainge",
		},
	},
	"WBB": {
		{
			ID:   831,
			Slug: "wild-beast-block",
			Name: "Wild Beast Block",
		},
	},
	"TLB": {
		{
			ID:   7544,
			Slug: "the-luxury-coin",
			Name: "The Luxury Coin",
		},
	},
	"TWI": {
		{
			ID:   7620,
			Slug: "trade-win",
			Name: "Trade.win",
		},
	},
	"SLS": {
		{
			ID:   1159,
			Slug: "salus",
			Name: "SaluS",
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
	"BIR": {
		{
			ID:   3285,
			Slug: "birake",
			Name: "Birake",
		},
	},
	"BDG": {
		{
			ID:   2374,
			Slug: "bitdegree",
			Name: "BitDegree",
		},
	},
	"SOCC": {
		{
			ID:   1774,
			Slug: "socialcoin-socc",
			Name: "SocialCoin",
		},
	},
	"HOP": {
		{
			ID:   10050,
			Slug: "hoppy",
			Name: "HOPPY",
		},
	},
	"GODL": {
		{
			ID:   10947,
			Slug: "godl",
			Name: "GODL",
		},
	},
	"SCU": {
		{
			ID:   8199,
			Slug: "securypto",
			Name: "Securypto",
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
	"DAY": {
		{
			ID:   1985,
			Slug: "chronologic",
			Name: "Chronologic",
		},
	},
	"BAS": {
		{
			ID:   7816,
			Slug: "basis-share",
			Name: "Basis Share",
		},
	},
	"DOG": {
		{
			ID:   8818,
			Slug: "underdog",
			Name: "UnderDog",
		},
	},
	"PM": {
		{
			ID:   10302,
			Slug: "pomskey",
			Name: "Pomskey",
		},
	},
	"ARPA": {
		{
			ID:   4039,
			Slug: "arpa-chain",
			Name: "ARPA Chain",
		},
	},
	"VIBE": {
		{
			ID:   1983,
			Slug: "vibe",
			Name: "VIBE",
		},
	},
	"KGO": {
		{
			ID:   8877,
			Slug: "kiwigo",
			Name: "KIWIGO",
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
	"KLP": {
		{
			ID:   6507,
			Slug: "kulupu",
			Name: "Kulupu",
		},
	},
	"CVAG": {
		{
			ID:   10918,
			Slug: "new-crypto-village-accelerator",
			Name: "Crypto Village Accelerator",
		},
	},
	"MERCI": {
		{
			ID:   5646,
			Slug: "merci",
			Name: "MERCI",
		},
	},
	"MIDBULL": {
		{
			ID:   6086,
			Slug: "3x-long-midcap-index-token",
			Name: "3X Long Midcap Index Token",
		},
	},
	"ILV": {
		{
			ID:   8719,
			Slug: "illuvium",
			Name: "Illuvium",
		},
	},
	"EOST": {
		{
			ID:   4124,
			Slug: "eos-trust",
			Name: "EOS TRUST",
		},
	},
	"EC2": {
		{
			ID:   9360,
			Slug: "employmentcoin",
			Name: "EmploymentCoin",
		},
	},
	"PAID": {
		{
			ID:   8329,
			Slug: "paid-network",
			Name: "PAID Network",
		},
	},
	"ERK": {
		{
			ID:   4966,
			Slug: "eureka-coin",
			Name: "Eureka Coin",
		},
	},
	"SCT": {
		{
			ID:   8429,
			Slug: "clash-token",
			Name: "Clash Token",
		},
	},
	"KEMA": {
		{
			ID:   5236,
			Slug: "kemacoin",
			Name: "Kemacoin",
		},
	},
	"LINKBULL": {
		{
			ID:   6037,
			Slug: "3x-long-chainlink-token",
			Name: "3X Long Chainlink Token",
		},
	},
	"INVEST": {
		{
			ID:   10468,
			Slug: "investdex",
			Name: "InvestDex",
		},
	},
	"CTXC": {
		{
			ID:   2638,
			Slug: "cortex",
			Name: "Cortex",
		},
	},
	"NYAN-2": {
		{
			ID:   7643,
			Slug: "nyan-v2",
			Name: "Nyan V2",
		},
	},
	"YMEN": {
		{
			ID:   7039,
			Slug: "ymen-finance",
			Name: "Ymen.Finance",
		},
	},
	"IHF": {
		{
			ID:   3301,
			Slug: "invictus-hyperion-fund",
			Name: "Invictus Hyperion Fund",
		},
	},
	"AXE": {
		{
			ID:   3898,
			Slug: "axe",
			Name: "Axe",
		},
	},
	"MOK": {
		{
			ID:   10819,
			Slug: "mocktailswap",
			Name: "MocktailSwap",
		},
	},
	"YFE": {
		{
			ID:   7250,
			Slug: "yfe-money",
			Name: "YFE Money",
		},
	},
	"PENDLE": {
		{
			ID:   9481,
			Slug: "pendle",
			Name: "Pendle",
		},
	},
	"ZORA": {
		{
			ID:   7826,
			Slug: "zoracles",
			Name: "Zoracles",
		},
	},
	"GGTK": {
		{
			ID:   8156,
			Slug: "ggdapp",
			Name: "GGDApp",
		},
	},
	"SAFEETH": {
		{
			ID:   10316,
			Slug: "safeeth",
			Name: "SafeETH",
		},
	},
	"LAYERX": {
		{
			ID:   9559,
			Slug: "unilayerx",
			Name: "UNILAYERX",
		},
	},
	"ESPRO": {
		{
			ID:   8760,
			Slug: "esportspro",
			Name: "EsportsPro",
		},
	},
	"BPS": {
		{
			ID:   5815,
			Slug: "bitcoinpos",
			Name: "BitcoinPoS",
		},
	},
	"LBY": {
		{
			ID:   8569,
			Slug: "libonomy",
			Name: "Libonomy",
		},
	},
	"BB": {
		{
			ID:   8345,
			Slug: "blackberry-tokenized-stock-ftx",
			Name: "BlackBerry tokenized stock FTX",
		},
	},
	"LDO": {
		{
			ID:   8000,
			Slug: "lido-dao",
			Name: "Lido DAO Token",
		},
	},
	"MTRG": {
		{
			ID:   5919,
			Slug: "meter-governance",
			Name: "Meter Governance",
		},
	},
	"NYANTE": {
		{
			ID:   8067,
			Slug: "nyantereum",
			Name: "Nyantereum International",
		},
	},
	"EDGT": {
		{
			ID:   9466,
			Slug: "edgecoin",
			Name: "Edgecoin",
		},
	},
	"STP": {
		{
			ID:   5785,
			Slug: "stpay",
			Name: "STPAY",
		},
	},
	"CHR": {
		{
			ID:   3978,
			Slug: "chromia",
			Name: "Chromia",
		},
	},
	"C20": {
		{
			ID:   2444,
			Slug: "c20",
			Name: "CRYPTO20",
		},
	},
	"IQN": {
		{
			ID:   3336,
			Slug: "iqeon",
			Name: "IQeon",
		},
	},
	"GMR": {
		{
			ID:   9629,
			Slug: "gmr-finance",
			Name: "GMR Finance",
		},
	},
	"INV": {
		{
			ID:   8720,
			Slug: "inverse-finance",
			Name: "Inverse Finance",
		},
	},
	"TRXBULL": {
		{
			ID:   5378,
			Slug: "3x-long-trx-token",
			Name: "3X Long TRX Token",
		},
	},
	"RSR": {
		{
			ID:   3964,
			Slug: "reserve-rights",
			Name: "Reserve Rights",
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
	"RSUN": {
		{
			ID:   11004,
			Slug: "risingsun",
			Name: "RisingSun",
		},
	},
	"YFP": {
		{
			ID:   6915,
			Slug: "yearn-finance-protocol",
			Name: "Yearn Finance Protocol",
		},
	},
	"CBANK": {
		{
			ID:   8113,
			Slug: "cryptobank",
			Name: "CryptoBank",
		},
	},
	"ARI": {
		{
			ID:   9565,
			Slug: "arise-finance",
			Name: "Arise Finance",
		},
	},
	"BOOMC": {
		{
			ID:   10575,
			Slug: "boomcoin",
			Name: "BoomCoin",
		},
	},
	"PRO": {
		{
			ID:   1974,
			Slug: "propy",
			Name: "Propy",
		},
	},
	"MRI": {
		{
			ID:   3254,
			Slug: "mirai",
			Name: "Mirai",
		},
	},
	"YFIIG": {
		{
			ID:   7193,
			Slug: "yfii-gold",
			Name: "YFII Gold",
		},
	},
	"SAFEARN": {
		{
			ID:   10988,
			Slug: "safe-earn",
			Name: "Safe Earn",
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
	"STOP": {
		{
			ID:   6766,
			Slug: "satopay-network",
			Name: "Satopay Network",
		},
	},
	"TOZ": {
		{
			ID:   8522,
			Slug: "tozex",
			Name: "TOZEX",
		},
	},
	"ROM": {
		{
			ID:   3670,
			Slug: "romtoken",
			Name: "ROMToken",
		},
	},
	"BAST": {
		{
			ID:   6821,
			Slug: "bast",
			Name: "Bast",
		},
	},
	"FARM": {
		{
			ID:   6859,
			Slug: "harvest-finance",
			Name: "Harvest Finance",
		},
	},
	"WEXPOLY": {
		{
			ID:   10725,
			Slug: "waultswap-polygon",
			Name: "WaultSwap Polygon",
		},
	},
	"ACAT": {
		{
			ID:   2525,
			Slug: "alphacat",
			Name: "Alphacat",
		},
	},
	"MDO": {
		{
			ID:   8486,
			Slug: "midas-dollar",
			Name: "Midas Dollar",
		},
	},
	"RES": {
		{
			ID:   5556,
			Slug: "resfinex-token",
			Name: "Resfinex Token",
		},
	},
	"PST": {
		{
			ID:   1930,
			Slug: "primas",
			Name: "Primas",
		},
	},
	"XANK": {
		{
			ID:   5659,
			Slug: "xank",
			Name: "Xank",
		},
	},
	"PBT": {
		{
			ID:   1841,
			Slug: "primalbase",
			Name: "Primalbase Token",
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
	"GNBU": {
		{
			ID:   10632,
			Slug: "nimbus-governance-token",
			Name: "Nimbus Governance Token",
		},
	},
	"NEBO": {
		{
			ID:   7597,
			Slug: "csp-dao",
			Name: "CSP DAO",
		},
	},
	"OCE": {
		{
			ID:   3880,
			Slug: "oceanex-token",
			Name: "OceanEx Token",
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
	"ADEL": {
		{
			ID:   6852,
			Slug: "akropolis-delphi",
			Name: "Akropolis Delphi",
		},
	},
	"QTCON": {
		{
			ID:   4746,
			Slug: "quiztok",
			Name: "Quiztok",
		},
	},
	"ABT": {
		{
			ID:   2545,
			Slug: "arcblock",
			Name: "Arcblock",
		},
	},
	"CEEK": {
		{
			ID:   2856,
			Slug: "ceek-vr",
			Name: "CEEK VR",
		},
	},
	"CREVA": {
		{
			ID:   986,
			Slug: "crevacoin",
			Name: "CrevaCoin",
		},
	},
	"AKITA": {
		{
			ID:   8378,
			Slug: "akita-inu",
			Name: "Akita Inu",
		},
	},
	"LPOOL": {
		{
			ID:   8545,
			Slug: "launchpool",
			Name: "Launchpool",
		},
	},
	"VRGX": {
		{
			ID:   7621,
			Slug: "vroomgo",
			Name: "VROOMGO",
		},
	},
	"MUNCH": {
		{
			ID:   9272,
			Slug: "munch-token",
			Name: "Munch Token",
		},
	},
	"ZBTC": {
		{
			ID:   10280,
			Slug: "zetta-bitcoin-hashrate-token",
			Name: "Zetta Bitcoin Hashrate Token",
		},
	},
	"FLA": {
		{
			ID:   7609,
			Slug: "fiola",
			Name: "Fiola",
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
	"BCI": {
		{
			ID:   2702,
			Slug: "bitcoin-interest",
			Name: "Bitcoin Interest",
		},
	},
	"MOONRISE": {
		{
			ID:   10681,
			Slug: "moonrise",
			Name: "MoonRise",
		},
	},
	"ALOHA": {
		{
			ID:   8548,
			Slug: "aloha",
			Name: "Aloha",
		},
	},
	"SWASS": {
		{
			ID:   10194,
			Slug: "swass-finance",
			Name: "SWASS.FINANCE",
		},
	},
	"SNOWGE": {
		{
			ID:   9772,
			Slug: "snowgecoin",
			Name: "SnowgeCoin",
		},
	},
	"CAKE": {
		{
			ID:   7186,
			Slug: "pancakeswap",
			Name: "PancakeSwap",
		},
	},
	"FSN": {
		{
			ID:   2530,
			Slug: "fusion",
			Name: "Fusion",
		},
	},
	"UDOKI": {
		{
			ID:   9474,
			Slug: "unicly-doki-doki-collection",
			Name: "Unicly Doki Doki Collection",
		},
	},
	"PALG": {
		{
			ID:   9819,
			Slug: "palgold",
			Name: "PalGold",
		},
	},
	"QRX": {
		{
			ID:   7510,
			Slug: "quiverx",
			Name: "QuiverX",
		},
	},
	"ZDX": {
		{
			ID:   4184,
			Slug: "zer-dex",
			Name: "Zer-Dex",
		},
	},
	"CSPC": {
		{
			ID:   6412,
			Slug: "cspc",
			Name: "CSPC",
		},
	},
	"EULER": {
		{
			ID:   9368,
			Slug: "euler-tools",
			Name: "Euler Tools",
		},
	},
	"GTON": {
		{
			ID:   10006,
			Slug: "graviton-one",
			Name: "Graviton",
		},
	},
	"RIGEL": {
		{
			ID:   8304,
			Slug: "rigel-finance",
			Name: "Rigel Finance",
		},
	},
	"GGT": {
		{
			ID:   7658,
			Slug: "gard-governance-token",
			Name: "GARD Governance Token",
		},
	},
	"SWAPS": {
		{
			ID:   9412,
			Slug: "nftswaps",
			Name: "NFTSwaps",
		},
	},
	"MEX": {
		{
			ID:   3286,
			Slug: "mex",
			Name: "MEX",
		},
	},
	"SURE": {
		{
			ID:   5113,
			Slug: "insure",
			Name: "inSure DeFi",
		},
	},
	"LAND": {
		{
			ID:   8752,
			Slug: "landbox",
			Name: "Landbox",
		},
	},
	"BIDCOM": {
		{
			ID:   10321,
			Slug: "bidcommerce",
			Name: "Bidcommerce",
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
	"XAMP": {
		{
			ID:   5998,
			Slug: "antiample",
			Name: "Antiample",
		},
	},
	"SYN": {
		{
			ID:   7520,
			Slug: "synlev",
			Name: "SynLev",
		},
	},
	"KOJI": {
		{
			ID:   10680,
			Slug: "koji",
			Name: "Koji",
		},
	},
	"LNCHX": {
		{
			ID:   9771,
			Slug: "launchx",
			Name: "LaunchX",
		},
	},
	"AVS": {
		{
			ID:   8832,
			Slug: "algovest",
			Name: "AlgoVest",
		},
	},
	"EGR": {
		{
			ID:   5075,
			Slug: "egoras",
			Name: "Egoras",
		},
	},
	"PSG": {
		{
			ID:   5226,
			Slug: "paris-saint-germain-fan-token",
			Name: "Paris Saint-Germain Fan Token",
		},
	},
	"POLK": {
		{
			ID:   8579,
			Slug: "polkamarkets",
			Name: "Polkamarkets",
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
	"DANK": {
		{
			ID:   9041,
			Slug: "mu-dank",
			Name: "MU DANK",
		},
	},
	"ICON": {
		{
			ID:   1528,
			Slug: "iconic",
			Name: "Iconic",
		},
	},
	"BETH": {
		{
			ID:   8353,
			Slug: "beacon-eth",
			Name: "Beacon ETH",
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
	"BKBT": {
		{
			ID:   2914,
			Slug: "beekan",
			Name: "BeeKan",
		},
	},
	"KODURO": {
		{
			ID:   9835,
			Slug: "koduro",
			Name: "Koduro",
		},
	},
	"LEX": {
		{
			ID:   7578,
			Slug: "elxis",
			Name: "Elxis",
		},
	},
	"BURNS": {
		{
			ID:   10734,
			Slug: "mr-burn-token",
			Name: "Mr Burn Token",
		},
	},
	"FNT": {
		{
			ID:   5871,
			Slug: "falcon-project",
			Name: "Falcon Project",
		},
	},
	"ESCE": {
		{
			ID:   3489,
			Slug: "escroco-emerald",
			Name: "Escroco Emerald",
		},
	},
	"1UP": {
		{
			ID:   4213,
			Slug: "uptrennd",
			Name: "Uptrennd",
		},
	},
	"SHB": {
		{
			ID:   3604,
			Slug: "skyhub-coin",
			Name: "SkyHub Coin",
		},
	},
	"$BABYDOGEINU": {
		{
			ID:   10914,
			Slug: "baby-doge-inu",
			Name: "BABY DOGE INU",
		},
	},
	"VX": {
		{
			ID:   5246,
			Slug: "vitex-coin",
			Name: "ViteX Coin",
		},
	},
	"ARX": {
		{
			ID:   5005,
			Slug: "arcs",
			Name: "ARCS",
		},
	},
	"SNB": {
		{
			ID:   5277,
			Slug: "synchrobitcoin",
			Name: "SynchroBitcoin",
		},
	},
	"CNB": {
		{
			ID:   5114,
			Slug: "coinsbit-token",
			Name: "Coinsbit Token",
		},
	},
	"PFR": {
		{
			ID:   2244,
			Slug: "payfair",
			Name: "Payfair",
		},
	},
	"NIA": {
		{
			ID:   8314,
			Slug: "nydronia",
			Name: "Nydronia",
		},
	},
	"MODEX": {
		{
			ID:   4916,
			Slug: "modex",
			Name: "Modex",
		},
	},
	"LABRA": {
		{
			ID:   9519,
			Slug: "labracoin",
			Name: "LabraCoin",
		},
	},
	"SEUR": {
		{
			ID:   10419,
			Slug: "seur",
			Name: "sEUR",
		},
	},
	"PAY": {
		{
			ID:   1758,
			Slug: "tenx",
			Name: "TenX",
		},
	},
	"VSO": {
		{
			ID:   9451,
			Slug: "verso-token",
			Name: "Verso Token",
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
	"HYBN": {
		{
			ID:   5934,
			Slug: "hey-bitcoin",
			Name: "Hey Bitcoin",
		},
	},
	"XMR": {
		{
			ID:   328,
			Slug: "monero",
			Name: "Monero",
		},
	},
	"MOF": {
		{
			ID:   2441,
			Slug: "molecular-future",
			Name: "Molecular Future",
		},
	},
	"IXT": {
		{
			ID:   1845,
			Slug: "ixledger",
			Name: "IXT",
		},
	},
	"BITTO": {
		{
			ID:   4534,
			Slug: "bitto",
			Name: "BITTO",
		},
	},
	"CLT": {
		{
			ID:   5401,
			Slug: "coinloan",
			Name: "CoinLoan",
		},
	},
	"WMATIC": {
		{
			ID:   8925,
			Slug: "wmatic",
			Name: "Wrapped Matic",
		},
	},
	"HANDY": {
		{
			ID:   7755,
			Slug: "handy",
			Name: "Handy",
		},
	},
	"GSMT": {
		{
			ID:   6851,
			Slug: "grafsound",
			Name: "GrafSound",
		},
	},
	"GRN": {
		{
			ID:   2746,
			Slug: "greenpower",
			Name: "GreenPower",
		},
	},
	"FOTO": {
		{
			ID:   8729,
			Slug: "unique-photo",
			Name: "Unique Photo",
		},
	},
	"ILG": {
		{
			ID:   9918,
			Slug: "ilgon",
			Name: "ILGON",
		},
	},
	"DAX": {
		{
			ID:   2696,
			Slug: "daex",
			Name: "DAEX",
		},
	},
	"LIBERTAS": {
		{
			ID:   5997,
			Slug: "libertas-token",
			Name: "Libertas Token",
		},
	},
	"JUI": {
		{
			ID:   6169,
			Slug: "juiice",
			Name: "JUIICE",
		},
	},
	"TEP": {
		{
			ID:   4677,
			Slug: "tepleton",
			Name: "Tepleton",
		},
	},
	"LELE": {
		{
			ID:   7781,
			Slug: "lelefoodchain",
			Name: "LeLeFoodChain",
		},
	},
	"AXS": {
		// {
		// 	ID:   1,
		// 	Slug: "1",
		// 	Name: "123",
		// },
		// {
		// 	ID:   1,
		// 	Slug: "1",
		// 	Name: "123",
		// },
		{
			ID:   6783,
			Slug: "axie-infinity",
			Name: "Axie Infinity",
		},
	},
	"NU": {
		{
			ID:   4761,
			Slug: "nucypher",
			Name: "NuCypher",
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
	"DEPAY": {
		{
			ID:   8181,
			Slug: "depay",
			Name: "DePay",
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
	"POE": {
		{
			ID:   1937,
			Slug: "poet",
			Name: "Po.et",
		},
	},
	"CUMINU": {
		{
			ID:   10061,
			Slug: "cuminu",
			Name: "CumInu",
		},
	},
	"POX": {
		{
			ID:   6682,
			Slug: "pollux-coin",
			Name: "Pollux Coin",
		},
	},
	"SNOGE": {
		{
			ID:   10007,
			Slug: "snoop-doge",
			Name: "Snoop Doge",
		},
	},
	"$HINA": {
		{
			ID:   10189,
			Slug: "hina-inu",
			Name: "Hina Inu",
		},
	},
	"AGIX": {
		{
			ID:   2424,
			Slug: "singularitynet",
			Name: "SingularityNET",
		},
	},
	"ELON": {
		{
			ID:   9436,
			Slug: "dogelon",
			Name: "Dogelon Mars",
		},
	},
	"BAX": {
		{
			ID:   2572,
			Slug: "babb",
			Name: "BABB",
		},
	},
	"CAPP": {
		{
			ID:   2248,
			Slug: "cappasity",
			Name: "Cappasity",
		},
	},
	"VISR": {
		{
			ID:   9170,
			Slug: "visor-finance",
			Name: "Visor.Finance",
		},
	},
	"PERI": {
		{
			ID:   9550,
			Slug: "peri-finance",
			Name: "PERI Finance",
		},
	},
	"AHT": {
		{
			ID:   6641,
			Slug: "ahatoken",
			Name: "AhaToken",
		},
	},
	"TKY": {
		{
			ID:   2507,
			Slug: "thekey",
			Name: "THEKEY",
		},
	},
	"NAVI": {
		{
			ID:   8340,
			Slug: "natus-vincere-fan-token",
			Name: "Natus Vincere Fan Token",
		},
	},
	"SCORGI": {
		{
			ID:   9716,
			Slug: "spacecorgi",
			Name: "SpaceCorgi",
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
	"MAPS": {
		{
			ID:   8166,
			Slug: "maps",
			Name: "MAPS",
		},
	},
	"TRA": {
		{
			ID:   7637,
			Slug: "trabzonspor-fan-token",
			Name: "Trabzonspor Fan Token",
		},
	},
	"ZASH": {
		{
			ID:   5610,
			Slug: "zimbocash",
			Name: "ZIMBOCASH",
		},
	},
	"AMON": {
		{
			ID:   4712,
			Slug: "amond",
			Name: "AmonD",
		},
	},
	"FYP": {
		{
			ID:   2126,
			Slug: "flypme",
			Name: "FlypMe",
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
	"POWER": {
		{
			ID:   5660,
			Slug: "unipower",
			Name: "UniPower",
		},
	},
	"SHIH": {
		{
			ID:   9712,
			Slug: "shih-tzu",
			Name: "Shih Tzu",
		},
	},
	"TMC": {
		{
			ID:   7235,
			Slug: "trading-membership-community",
			Name: "Trading Membership Community",
		},
	},
	"IOTX": {
		{
			ID:   2777,
			Slug: "iotex",
			Name: "IoTeX",
		},
	},
	"COTI": {
		{
			ID:   3992,
			Slug: "coti",
			Name: "COTI",
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
	"AKN": {
		{
			ID:   4618,
			Slug: "akoin",
			Name: "Akoin",
		},
	},
	"VLAD": {
		{
			ID:   9096,
			Slug: "vlad-finance",
			Name: "Vlad Finance",
		},
	},
	"YEED": {
		{
			ID:   3474,
			Slug: "yeed",
			Name: "YGGDRASH",
		},
	},
	"PLNC": {
		{
			ID:   263,
			Slug: "plncoin",
			Name: "PLNcoin",
		},
	},
	"F1C": {
		{
			ID:   3417,
			Slug: "future1coin",
			Name: "Future1coin",
		},
	},
	"CRT": {
		{
			ID:   6872,
			Slug: "carrot",
			Name: "Carrot",
		},
	},
	"MINTY": {
		{
			ID:   8694,
			Slug: "minty-art",
			Name: "Minty Art",
		},
	},
	"QUAM": {
		{
			ID:   8864,
			Slug: "quam-network",
			Name: "Quam Network",
		},
	},
	"WHEAT": {
		{
			ID:   9229,
			Slug: "wheat",
			Name: "WHEAT Token",
		},
	},
	"PMP": {
		{
			ID:   9090,
			Slug: "pumpy-farm",
			Name: "Pumpy farm",
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
	"LOWB": {
		{
			ID:   9673,
			Slug: "loser-coin",
			Name: "Loser Coin",
		},
	},
	"DML": {
		{
			ID:   2679,
			Slug: "decentralized-machine-learning",
			Name: "Decentralized Machine Learning",
		},
	},
	"DIAMONDS": {
		{
			ID:   10924,
			Slug: "black-diamond",
			Name: "Black Diamond",
		},
	},
	"BSY": {
		{
			ID:   5782,
			Slug: "bestay",
			Name: "Bestay",
		},
	},
	"WAULT": {
		{
			ID:   8588,
			Slug: "wault-finance",
			Name: "Wault Finance (OLD)",
		},
	},
	"SLME": {
		{
			ID:   8693,
			Slug: "slime-finance",
			Name: "Slime Finance",
		},
	},
	"ORCL5": {
		{
			ID:   8889,
			Slug: "oracle-top-5",
			Name: "Oracle Top 5 Tokens Index",
		},
	},
	"BTCDOM": {
		{
			ID:   7421,
			Slug: "bitfinex-bitcoin-dominance-perps",
			Name: "Bitfinex Bitcoin Dominance Perps",
		},
	},
	"BELT": {
		{
			ID:   8730,
			Slug: "belt",
			Name: "Belt Finance",
		},
	},
	"PWAR": {
		{
			ID:   10748,
			Slug: "polkawar",
			Name: "PolkaWar",
		},
	},
	"ZNT": {
		{
			ID:   3446,
			Slug: "zenswap-network-token",
			Name: "Zenswap Network Token",
		},
	},
	"ALA": {
		{
			ID:   6980,
			Slug: "aladiex",
			Name: "AladiEx",
		},
	},
	"STONK": {
		{
			ID:   5843,
			Slug: "stonk",
			Name: "STONK",
		},
	},
	"AR": {
		{
			ID:   5632,
			Slug: "arweave",
			Name: "Arweave",
		},
	},
	"DOCK": {
		{
			ID:   2675,
			Slug: "dock",
			Name: "Dock",
		},
	},
	"EQUAD": {
		{
			ID:   3625,
			Slug: "quadrantprotocol",
			Name: "QuadrantProtocol",
		},
	},
	"ITT": {
		{
			ID:   2103,
			Slug: "intelligent-trading-foundation",
			Name: "Intelligent Trading Foundation",
		},
	},
	"NOAHP": {
		{
			ID:   2599,
			Slug: "noah-coin",
			Name: "Noah Coin",
		},
	},
	"ETNX": {
		{
			ID:   4448,
			Slug: "electronero",
			Name: "Electronero",
		},
	},
	"XSR": {
		{
			ID:   4818,
			Slug: "xensor",
			Name: "Xensor",
		},
	},
	"UBEX": {
		{
			ID:   3140,
			Slug: "ubex",
			Name: "Ubex",
		},
	},
	"REW": {
		{
			ID:   5816,
			Slug: "rewardiqa",
			Name: "Rewardiqa",
		},
	},
	"KZC": {
		{
			ID:   2386,
			Slug: "kz-cash",
			Name: "KZ Cash",
		},
	},
	"COPE": {
		{
			ID:   9015,
			Slug: "cope",
			Name: "Cope",
		},
	},
	"XHT": {
		{
			ID:   10423,
			Slug: "hollaex-token",
			Name: "HollaEx Token",
		},
	},
	"BTNYX": {
		{
			ID:   7513,
			Slug: "bitonyx",
			Name: "BitOnyx",
		},
	},
	"SC": {
		{
			ID:   1042,
			Slug: "siacoin",
			Name: "Siacoin",
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
	"MUSD": {
		{
			ID:   5747,
			Slug: "mstable-usd",
			Name: "mStable USD",
		},
	},
	"ITC": {
		{
			ID:   2251,
			Slug: "iot-chain",
			Name: "IoT Chain",
		},
	},
	"1MT": {
		{
			ID:   4222,
			Slug: "1million-token",
			Name: "1Million Token",
		},
	},
	"NEWTON": {
		{
			ID:   6685,
			Slug: "newtonium",
			Name: "Newtonium",
		},
	},
	"XRGE": {
		{
			ID:   9106,
			Slug: "rougecoin",
			Name: "RougeCoin",
		},
	},
	"XEND": {
		{
			ID:   8519,
			Slug: "xend-finance",
			Name: "Xend Finance",
		},
	},
	"BASE": {
		{
			ID:   7838,
			Slug: "base-protocol",
			Name: "Base Protocol",
		},
	},
	"ENB": {
		{
			ID:   7716,
			Slug: "earnbase",
			Name: "Earnbase",
		},
	},
	"CCN": {
		{
			ID:   3770,
			Slug: "customcontractnetwork",
			Name: "CustomContractNetwork",
		},
	},
	"LEOS": {
		{
			ID:   10762,
			Slug: "leonicorn-swap",
			Name: "Leonicorn Swap",
		},
	},
	"BSCGIRL": {
		{
			ID:   10570,
			Slug: "binance-smart-chain-girl",
			Name: "Binance Smart Chain Girl",
		},
	},
	"MSD": {
		{
			ID:   2008,
			Slug: "msd",
			Name: "MSD",
		},
	},
	"$ANRX": {
		{
			ID:   8057,
			Slug: "anrkey-x",
			Name: "AnRKey X",
		},
	},
	"SENC": {
		{
			ID:   2624,
			Slug: "sentinel-chain",
			Name: "Sentinel Chain",
		},
	},
	"AAVEUP": {
		{
			ID:   7774,
			Slug: "aave-up",
			Name: "AAVEUP",
		},
	},
	"VDX": {
		{
			ID:   3962,
			Slug: "vodi-x",
			Name: "Vodi X",
		},
	},
	"BULK": {
		{
			ID:   10235,
			Slug: "bulk",
			Name: "Bulk",
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
	"XBI": {
		{
			ID:   3166,
			Slug: "bitcoin-incognito",
			Name: "Bitcoin Incognito",
		},
	},
	"BQT": {
		{
			ID:   3231,
			Slug: "blockchain-quotations-index-token",
			Name: "Blockchain Quotations Index Token",
		},
	},
	"WFX": {
		{
			ID:   3990,
			Slug: "webflix-token",
			Name: "Webflix Token",
		},
	},
	"OROS": {
		{
			ID:   10842,
			Slug: "oros-finance",
			Name: "OROS.finance",
		},
	},
	"GOGO": {
		{
			ID:   8375,
			Slug: "gogo-finance",
			Name: "GOGO.finance",
		},
	},
	"GAIA": {
		{
			ID:   9800,
			Slug: "gaiadao",
			Name: "GaiaDAO",
		},
	},
	"TROY": {
		{
			ID:   5007,
			Slug: "troy",
			Name: "TROY",
		},
	},
	"NFY": {
		{
			ID:   7389,
			Slug: "non-fungible-yearn",
			Name: "Non-Fungible Yearn",
		},
	},
	"TOAD": {
		{
			ID:   9983,
			Slug: "toad-network",
			Name: "toad.network",
		},
	},
	"IND": {
		{
			ID:   1967,
			Slug: "indorse-token",
			Name: "Indorse Token",
		},
	},
	"HOTCROSS": {
		{
			ID:   9867,
			Slug: "hot-cross",
			Name: "Hot Cross",
		},
	},
	"PUG": {
		{
			ID:   10460,
			Slug: "pug-cash",
			Name: "Pug Cash",
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
	"VIG": {
		{
			ID:   6562,
			Slug: "vig",
			Name: "VIG",
		},
	},
	"DBET": {
		{
			ID:   2175,
			Slug: "decent-bet",
			Name: "DecentBet",
		},
	},
	"BOOST": {
		{
			ID:   6862,
			Slug: "boosted-finance",
			Name: "Boosted Finance",
		},
	},
	"SWIFT": {
		{
			ID:   3933,
			Slug: "swiftcash",
			Name: "SwiftCash",
		},
	},
	"KINE": {
		{
			ID:   8790,
			Slug: "kine",
			Name: "KINE",
		},
	},
	"PASC": {
		{
			ID:   1473,
			Slug: "pascal",
			Name: "Pascal",
		},
	},
	"PRARE": {
		{
			ID:   9544,
			Slug: "polkarare",
			Name: "POLKARARE",
		},
	},
	"ROOM": {
		{
			ID:   8351,
			Slug: "optionroom",
			Name: "OptionRoom",
		},
	},
	"JFI": {
		{
			ID:   6898,
			Slug: "jackpool-finance",
			Name: "JackPool.finance",
		},
	},
	"MT": {
		{
			ID:   2712,
			Slug: "mytoken",
			Name: "MyToken",
		},
	},
	"DAO1": {
		{
			ID:   10384,
			Slug: "dao1",
			Name: "DAO1",
		},
	},
	"VOICE": {
		{
			ID:   7518,
			Slug: "nix-bridge-token",
			Name: "Voice Token",
		},
	},
	"TLW": {
		{
			ID:   5399,
			Slug: "tilwiki",
			Name: "TILWIKI",
		},
	},
	"FET": {
		{
			ID:   3773,
			Slug: "fetch",
			Name: "Fetch.ai",
		},
	},
	"NBOT": {
		{
			ID:   4047,
			Slug: "naka-bodhi-token",
			Name: "Naka Bodhi Token",
		},
	},
	"CCAKE": {
		{
			ID:   8747,
			Slug: "cheesecakeswap-token",
			Name: "CheesecakeSwap Token",
		},
	},
	"OCAT": {
		{
			ID:   10633,
			Slug: "orange-cat-token",
			Name: "Orange Cat Token",
		},
	},
	"pETH18C": {
		{
			ID:   9034,
			Slug: "peth18c",
			Name: "pETH18C",
		},
	},
	"ZEC": {
		{
			ID:   1437,
			Slug: "zcash",
			Name: "Zcash",
		},
	},
	"NCDT": {
		{
			ID:   6933,
			Slug: "nuco-cloud",
			Name: "Nuco.cloud",
		},
	},
	"BABYDOGE": {
		{
			ID:   10970,
			Slug: "babydoge-coin",
			Name: "BabyDoge ETH",
		},
	},
	"CASH": {
		{
			ID:   5038,
			Slug: "litecash",
			Name: "Litecash",
		},
	},
	"UPBNB": {
		{
			ID:   9671,
			Slug: "upbnb",
			Name: "upBNB",
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
	"APE": {
		{
			ID:   7257,
			Slug: "apecoin",
			Name: "APEcoin",
		},
	},
	"HUA": {
		{
			ID:   9941,
			Slug: "chihuahua",
			Name: "Chihuahua",
		},
	},
	"UOS": {
		{
			ID:   4189,
			Slug: "ultra",
			Name: "Ultra",
		},
	},
	"CNUS": {
		{
			ID:   3648,
			Slug: "coinus",
			Name: "CoinUs",
		},
	},
	"TANGO": {
		{
			ID:   8712,
			Slug: "keytango",
			Name: "keyTango",
		},
	},
	"XMY": {
		{
			ID:   182,
			Slug: "myriad",
			Name: "Myriad",
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
	"JEM": {
		{
			ID:   7723,
			Slug: "itchiro-games",
			Name: "Itchiro Games",
		},
	},
	"DOOS": {
		{
			ID:   6983,
			Slug: "doos-token",
			Name: "DOOS TOKEN",
		},
	},
	"BTAP": {
		{
			ID:   9060,
			Slug: "bta-protocol",
			Name: "BTA Protocol",
		},
	},
	"UBIN": {
		{
			ID:   7515,
			Slug: "ubiner",
			Name: "Ubiner",
		},
	},
	"L3P": {
		{
			ID:   8845,
			Slug: "lepricon",
			Name: "Lepricon",
		},
	},
	"CSPN": {
		{
			ID:   3894,
			Slug: "crypto-sports",
			Name: "Crypto Sports",
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
	"VEGA": {
		{
			ID:   10223,
			Slug: "vegaprotocol",
			Name: "Vega Protocol",
		},
	},
	"DACC": {
		{
			ID:   2986,
			Slug: "dacc",
			Name: "DACC",
		},
	},
	"LEV": {
		{
			ID:   9505,
			Slug: "lever-token",
			Name: "Lever Token",
		},
	},
	"SAFESPACE": {
		{
			ID:   10075,
			Slug: "safespace",
			Name: "SAFESPACE",
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
	"CATT": {
		{
			ID:   4045,
			Slug: "catex-token",
			Name: "Catex Token",
		},
	},
	"XPOSE": {
		{
			ID:   9704,
			Slug: "xpose",
			Name: "Xpose Protocol",
		},
	},
	"VNTW": {
		{
			ID:   9230,
			Slug: "value-network",
			Name: "Value Network",
		},
	},
	"FOXD": {
		{
			ID:   9434,
			Slug: "foxdcoin",
			Name: "FoxDcoin",
		},
	},
	"GNO": {
		{
			ID:   1659,
			Slug: "gnosis-gno",
			Name: "Gnosis",
		},
	},
	"CTI": {
		{
			ID:   7860,
			Slug: "clintex-cti",
			Name: "ClinTex CTi",
		},
	},
	"MBONK": {
		{
			ID:   6116,
			Slug: "bonk",
			Name: "megaBONK",
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
	"GSE": {
		{
			ID:   3123,
			Slug: "gsenetwork",
			Name: "GSENetwork",
		},
	},
	"PYX": {
		{
			ID:   9168,
			Slug: "pyxis-network",
			Name: "PYXIS Network",
		},
	},
	"THE": {
		{
			ID:   5078,
			Slug: "thenode",
			Name: "THENODE",
		},
	},
	"LEM": {
		{
			ID:   9142,
			Slug: "lemur-finance",
			Name: "Lemur Finance",
		},
	},
	"NXM": {
		{
			ID:   5830,
			Slug: "nxm",
			Name: "NXM",
		},
	},
	"LABS": {
		{
			ID:   8813,
			Slug: "labs-group",
			Name: "LABS Group",
		},
	},
	"ARTH": {
		{
			ID:   10865,
			Slug: "arth-new",
			Name: "ARTH [polygon]",
		},
	},
	"VENUS": {
		{
			ID:   10073,
			Slug: "venusia",
			Name: "Venusia",
		},
	},
	"CLO": {
		{
			ID:   2757,
			Slug: "callisto-network",
			Name: "Callisto Network",
		},
	},
	"USNBT": {
		{
			ID:   626,
			Slug: "nubits",
			Name: "NuBits",
		},
	},
	"EVIL": {
		{
			ID:   1165,
			Slug: "evil-coin",
			Name: "Evil Coin",
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
	"TBTC": {
		{
			ID:   5776,
			Slug: "tbtc",
			Name: "tBTC",
		},
	},
	"BNSD": {
		{
			ID:   6911,
			Slug: "bnsd-finance",
			Name: "BNSD Finance",
		},
	},
	"RNDR": {
		{
			ID:   5690,
			Slug: "render-token",
			Name: "Render Token",
		},
	},
	"CORE": {
		{
			ID:   7242,
			Slug: "cvault-finance",
			Name: "cVault.finance",
		},
	},
	"TRIAS": {
		{
			ID:   8936,
			Slug: "trias-token",
			Name: "Trias Token (new)",
		},
	},
	"PXG": {
		{
			ID:   3639,
			Slug: "playgame",
			Name: "PlayGame",
		},
	},
	"TCASH": {
		{
			ID:   3980,
			Slug: "tcash",
			Name: "TCASH",
		},
	},
	"XSH": {
		{
			ID:   2144,
			Slug: "shield-xsh",
			Name: "SHIELD",
		},
	},
	"UNL": {
		{
			ID:   8103,
			Slug: "unilock-network",
			Name: "unilock.network",
		},
	},
	"TST": {
		{
			ID:   5393,
			Slug: "touch-social",
			Name: "Touch Social",
		},
	},
	"BLK": {
		{
			ID:   170,
			Slug: "blackcoin",
			Name: "BlackCoin",
		},
	},
	"TAN": {
		{
			ID:   4272,
			Slug: "taklimakan-network",
			Name: "Taklimakan Network",
		},
	},
	"PLEX": {
		{
			ID:   7980,
			Slug: "mineplex",
			Name: "MinePlex",
		},
	},
	"AWS": {
		{
			ID:   9547,
			Slug: "aurussilver",
			Name: "AurusSILVER",
		},
	},
	"MPAD": {
		{
			ID:   10203,
			Slug: "moonpad",
			Name: "Moonpad",
		},
	},
	"MFT": {
		{
			ID:   2896,
			Slug: "mainframe",
			Name: "Hifi Finance",
		},
	},
	"FEI": {
		{
			ID:   8642,
			Slug: "fei-protocol",
			Name: "Fei Protocol",
		},
	},
	"eRSDL": {
		{
			ID:   7553,
			Slug: "unfederalreserve",
			Name: "unFederalReserve",
		},
	},
	"WCT": {
		{
			ID:   1527,
			Slug: "waves-community-token",
			Name: "Waves Community Token",
		},
	},
	"CLAIM": {
		{
			ID:   9349,
			Slug: "claim",
			Name: "CLAIM",
		},
	},
	"CARR": {
		{
			ID:   8806,
			Slug: "carnomaly",
			Name: "Carnomaly",
		},
	},
	"PINT": {
		{
			ID:   8571,
			Slug: "pub-finance",
			Name: "Pub Finance",
		},
	},
	"SHIFT": {
		{
			ID:   1050,
			Slug: "shift",
			Name: "Shift",
		},
	},
	"DPR": {
		{
			ID:   8894,
			Slug: "deeper-network",
			Name: "Deeper Network",
		},
	},
	"CLR": {
		{
			ID:   4681,
			Slug: "color-platform",
			Name: "Color Platform",
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
	"DSS": {
		{
			ID:   7145,
			Slug: "defi-shopping-stake",
			Name: "Defi Shopping Stake",
		},
	},
	"MYID": {
		{
			ID:   9187,
			Slug: "my-identity-coin",
			Name: "MY IDENTITY COIN",
		},
	},
	"CATS": {
		{
			ID:   10663,
			Slug: "catoshi-nakamoto",
			Name: "Catoshi Nakamoto",
		},
	},
	"AFC": {
		{
			ID:   7469,
			Slug: "apiary-fund-coin",
			Name: "Apiary Fund Coin",
		},
	},
	"TEL": {
		{
			ID:   2394,
			Slug: "telcoin",
			Name: "Telcoin",
		},
	},
	"HAI": {
		{
			ID:   5583,
			Slug: "hackenai",
			Name: "Hacken Token",
		},
	},
	"ROOBEE": {
		{
			ID:   4804,
			Slug: "roobee",
			Name: "ROOBEE",
		},
	},
	"L2P": {
		{
			ID:   5993,
			Slug: "lung-protocol",
			Name: "Lung Protocol",
		},
	},
	"PROT": {
		{
			ID:   9193,
			Slug: "prostarter",
			Name: "Prostarter",
		},
	},
	"NFTD": {
		{
			ID:   9249,
			Slug: "nftd-protocol",
			Name: "NFTD Protocol",
		},
	},
	"TOR": {
		{
			ID:   5518,
			Slug: "torex",
			Name: "Torex",
		},
	},
	"DHOLD": {
		{
			ID:   10230,
			Slug: "diamondhold",
			Name: "DiamondHold",
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
	"FCL": {
		{
			ID:   8544,
			Slug: "fractal",
			Name: "Fractal",
		},
	},
	"LCS": {
		{
			ID:   2970,
			Slug: "local-coin-swap",
			Name: "LocalCoinSwap",
		},
	},
	"C2": {
		{
			ID:   367,
			Slug: "coin2-1",
			Name: "Coin2.1",
		},
	},
	"YFBT": {
		{
			ID:   7237,
			Slug: "yearn-finance-bit",
			Name: "Yearn Finance Bit",
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
	"PRQBOOST": {
		{
			ID:   8488,
			Slug: "parsiq-boost",
			Name: "Parsiq Boost",
		},
	},
	"UENC": {
		{
			ID:   6525,
			Slug: "universalenergychain",
			Name: "UniversalEnergyChain",
		},
	},
	"ZLW": {
		{
			ID:   5614,
			Slug: "zelwin",
			Name: "Zelwin",
		},
	},
	"MTV": {
		{
			ID:   3853,
			Slug: "multivac",
			Name: "MultiVAC",
		},
	},
	"LNC": {
		{
			ID:   2677,
			Slug: "linker-coin",
			Name: "Linker Coin",
		},
	},
	"SWC": {
		{
			ID:   3760,
			Slug: "scanetchain",
			Name: "Scanetchain",
		},
	},
	"ETHPLO": {
		{
			ID:   4076,
			Slug: "ethplode",
			Name: "ETHplode",
		},
	},
	"CNZ": {
		{
			ID:   6397,
			Slug: "coinzo-token",
			Name: "Coinzo Token",
		},
	},
	"INO": {
		{
			ID:   3085,
			Slug: "ino-coin",
			Name: "INO COIN",
		},
	},
	"GERA": {
		{
			ID:   8270,
			Slug: "gera-coin",
			Name: "Gera Coin",
		},
	},
	"TAIYO": {
		{
			ID:   11030,
			Slug: "taiyo",
			Name: "TAIYO",
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
	"PHOON": {
		{
			ID:   8306,
			Slug: "typhoon-cash",
			Name: "Typhoon Cash",
		},
	},
	"GVT": {
		{
			ID:   2181,
			Slug: "genesis-vision",
			Name: "Genesis Vision",
		},
	},
	"BALPHA": {
		{
			ID:   8710,
			Slug: "balpha",
			Name: "bAlpha",
		},
	},
	"BVND": {
		{
			ID:   8058,
			Slug: "binance-vnd",
			Name: "Binance VND",
		},
	},
	"NOVA": {
		{
			ID:   4835,
			Slug: "nova",
			Name: "NOVA",
		},
	},
	"SCONEX": {
		{
			ID:   9827,
			Slug: "sportcash-one",
			Name: "Sportcash One",
		},
	},
	"GZRO": {
		{
			ID:   3488,
			Slug: "gravity",
			Name: "Gravity",
		},
	},
	"CALL": {
		{
			ID:   5019,
			Slug: "global-crypto-alliance",
			Name: "Global Crypto Alliance",
		},
	},
	"ARB": {
		{
			ID:   938,
			Slug: "arbit",
			Name: "ARbit",
		},
	},
	"OAK": {
		{
			ID:   9152,
			Slug: "octree-oak",
			Name: "Octree Finance",
		},
	},
	"WATCH": {
		{
			ID:   8621,
			Slug: "yieldwatch",
			Name: "yieldwatch",
		},
	},
	"REX": {
		{
			ID:   1961,
			Slug: "imbrex",
			Name: "imbrex",
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
	"TRAIL": {
		{
			ID:   10732,
			Slug: "polkatrail",
			Name: "PolkaTrail",
		},
	},
	"MDU": {
		{
			ID:   6028,
			Slug: "mdu",
			Name: "MDUKEY",
		},
	},
	"ICEBRK": {
		{
			ID:   10304,
			Slug: "icebreak-r",
			Name: "IceBreak-R",
		},
	},
	"OXT": {
		{
			ID:   5026,
			Slug: "orchid",
			Name: "Orchid",
		},
	},
	"MASS": {
		{
			ID:   5548,
			Slug: "massnet",
			Name: "Massnet",
		},
	},
	"MYST": {
		{
			ID:   1721,
			Slug: "mysterium",
			Name: "Mysterium",
		},
	},
	"INDEX": {
		{
			ID:   7336,
			Slug: "index-cooperative",
			Name: "Index Cooperative",
		},
	},
	"KEYFI": {
		{
			ID:   8561,
			Slug: "keyfi",
			Name: "KeyFi",
		},
	},
	"ASAP": {
		{
			ID:   9289,
			Slug: "chainswap",
			Name: "Chainswap",
		},
	},
	"WOOFY": {
		{
			ID:   9719,
			Slug: "woofy",
			Name: "Woofy",
		},
	},
	"MOONDAY": {
		{
			ID:   7481,
			Slug: "moonday-finance",
			Name: "Moonday Finance",
		},
	},
	"iOWN": {
		{
			ID:   5337,
			Slug: "iown-token",
			Name: "iOWN Token",
		},
	},
	"QI": {
		{
			ID:   8510,
			Slug: "qiswap",
			Name: "QiSwap",
		},
	},
	"PREMIA": {
		{
			ID:   8476,
			Slug: "premia",
			Name: "Premia",
		},
	},
	"SAFECITY": {
		{
			ID:   10473,
			Slug: "safecity",
			Name: "SafeCity",
		},
	},
	"DOGE": {
		{
			ID:   74,
			Slug: "dogecoin",
			Name: "Dogecoin",
		},
	},
	"DAOFI": {
		{
			ID:   7419,
			Slug: "daofi",
			Name: "DAOFi",
		},
	},
	"PLUT": {
		{
			ID:   7968,
			Slug: "pluto",
			Name: "Pluto",
		},
	},
	"XPD": {
		{
			ID:   260,
			Slug: "petrodollar",
			Name: "PetroDollar",
		},
	},
	"NEWB": {
		{
			ID:   10888,
			Slug: "newb-farm",
			Name: "NewB.Farm",
		},
	},
	"ASS": {
		{
			ID:   9456,
			Slug: "australian-safe-shepherd",
			Name: "Australian Safe Shepherd",
		},
	},
	"LSV": {
		{
			ID:   5577,
			Slug: "litecoin-sv",
			Name: "Litecoin SV",
		},
	},
	"C98": {
		{
			ID:   10903,
			Slug: "coin98",
			Name: "Coin98",
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
	"SLTH": {
		{
			ID:   10596,
			Slug: "slothi",
			Name: "SLOTHI",
		},
	},
	"LOTTO": {
		{
			ID:   8399,
			Slug: "lotto",
			Name: "Lotto",
		},
	},
	"XSUTER": {
		{
			ID:   9838,
			Slug: "xsuter",
			Name: "xSuter",
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
	"GZIL": {
		{
			ID:   8031,
			Slug: "governance-zil",
			Name: "governance ZIL",
		},
	},
	"VIKING": {
		{
			ID:   8555,
			Slug: "viking-swap",
			Name: "Viking Swap",
		},
	},
	"GM": {
		{
			ID:   5698,
			Slug: "gm-holding",
			Name: "GM Holding",
		},
	},
	"XNV": {
		{
			ID:   3596,
			Slug: "nerva",
			Name: "Nerva",
		},
	},
	"SKRT": {
		{
			ID:   10554,
			Slug: "sekuritance",
			Name: "Sekuritance",
		},
	},
	"ECC": {
		{
			ID:   212,
			Slug: "eccoin",
			Name: "ECC",
		},
	},
	"SCR": {
		{
			ID:   3094,
			Slug: "scorum-coins",
			Name: "Scorum Coins",
		},
	},
	"SFG": {
		{
			ID:   7187,
			Slug: "s-finance",
			Name: "S.Finance",
		},
	},
	"MSS": {
		{
			ID:   8465,
			Slug: "monster-slayer-share",
			Name: "Monster Slayer Share",
		},
	},
	"TOMO": {
		{
			ID:   2570,
			Slug: "tomochain",
			Name: "TomoChain",
		},
	},
	"CV": {
		{
			ID:   2450,
			Slug: "carvertical",
			Name: "carVertical",
		},
	},
	"DMS": {
		{
			ID:   5143,
			Slug: "documentchain",
			Name: "Documentchain",
		},
	},
	"BRT": {
		{
			ID:   10602,
			Slug: "base-reward-token",
			Name: "Base Reward Token",
		},
	},
	"SLT": {
		{
			ID:   2471,
			Slug: "smartlands-network",
			Name: "Smartlands Network",
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
	"GEAR": {
		{
			ID:   6593,
			Slug: "bitgear",
			Name: "Bitgear",
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
	"QDAO": {
		{
			ID:   4053,
			Slug: "q-dao-governance-token",
			Name: "Q DAO Governance token v1.0",
		},
	},
	"STND": {
		{
			ID:   9251,
			Slug: "standard-protocol",
			Name: "Standard Protocol",
		},
	},
	"IONC": {
		{
			ID:   3506,
			Slug: "ionchain",
			Name: "IONChain",
		},
	},
	"CSC": {
		{
			ID:   45,
			Slug: "casinocoin",
			Name: "CasinoCoin",
		},
	},
	"XMARK": {
		{
			ID:   8835,
			Slug: "xmark",
			Name: "xMARK",
		},
	},
	"JACK": {
		{
			ID:   5486,
			Slug: "jack-token",
			Name: "Jack Token",
		},
	},
	"KOIN": {
		{
			ID:   8282,
			Slug: "koinos",
			Name: "Koinos",
		},
	},
	"THX": {
		{
			ID:   3916,
			Slug: "thorenext",
			Name: "ThoreNext",
		},
	},
	"SPIZ": {
		{
			ID:   6626,
			Slug: "space-iz",
			Name: "SPACE-iZ",
		},
	},
	"OBIC": {
		{
			ID:   6982,
			Slug: "obic",
			Name: "OBIC",
		},
	},
	"WEWON": {
		{
			ID:   10645,
			Slug: "wewon-world",
			Name: "WeWon World",
		},
	},
	"CONI": {
		{
			ID:   2895,
			Slug: "coni",
			Name: "Coni",
		},
	},
	"ROBBIN": {
		{
			ID:   10883,
			Slug: "robbin-hood",
			Name: "ROBBIN HOOD",
		},
	},
	"ASA": {
		{
			ID:   3477,
			Slug: "asura-coin",
			Name: "Asura Coin",
		},
	},
	"SGO": {
		{
			ID:   10751,
			Slug: "sportemon-go",
			Name: "Sportemon-Go",
		},
	},
	"MOONRABBIT": {
		{
			ID:   10249,
			Slug: "moonrabbit-money",
			Name: "MoonRabbit",
		},
	},
	"IDX": {
		{
			ID:   5898,
			Slug: "index-chain",
			Name: "Index Chain",
		},
	},
	"KYTE": {
		{
			ID:   9192,
			Slug: "kambria-yield-tuning-engine",
			Name: "Kambria Yield Tuning Engine",
		},
	},
	"PTOY": {
		{
			ID:   1708,
			Slug: "patientory",
			Name: "Patientory",
		},
	},
	"CF": {
		{
			ID:   898,
			Slug: "californium",
			Name: "Californium",
		},
	},
	"UGAS": {
		{
			ID:   3863,
			Slug: "ugas",
			Name: "UGAS",
		},
	},
	"LHT": {
		{
			ID:   4230,
			Slug: "lighthouse-token",
			Name: "LHT",
		},
	},
	"imBTC": {
		{
			ID:   5292,
			Slug: "the-tokenized-bitcoin",
			Name: "The Tokenized Bitcoin",
		},
	},
	"DCD": {
		{
			ID:   7159,
			Slug: "digital-currency-daily",
			Name: "Digital Currency Daily",
		},
	},
	"KAI": {
		{
			ID:   5453,
			Slug: "kardiachain",
			Name: "KardiaChain",
		},
	},
	"SHIBERUS": {
		{
			ID:   10874,
			Slug: "shiberus-inu",
			Name: "Shiberus Inu",
		},
	},
	"DETS": {
		{
			ID:   6722,
			Slug: "dextrust",
			Name: "Dextrust",
		},
	},
	"AID": {
		{
			ID:   2462,
			Slug: "aidcoin",
			Name: "AidCoin",
		},
	},
	"FEX": {
		{
			ID:   3800,
			Slug: "fidex-token",
			Name: "FidexToken",
		},
	},
	"XGM": {
		{
			ID:   5847,
			Slug: "defis",
			Name: "Defis",
		},
	},
	"PAXEX": {
		{
			ID:   3338,
			Slug: "paxex",
			Name: "PAXEX",
		},
	},
	"B360": {
		{
			ID:   9128,
			Slug: "b360",
			Name: "B360",
		},
	},
	"VIAGRA": {
		{
			ID:   10037,
			Slug: "viagra-token",
			Name: "Viagra Token",
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
	"BYTS": {
		{
			ID:   6393,
			Slug: "bytus",
			Name: "Bytus",
		},
	},
	"INARI": {
		{
			ID:   10809,
			Slug: "inari",
			Name: "Inari",
		},
	},
	"BFF": {
		{
			ID:   5180,
			Slug: "bitcoffeen",
			Name: "Bitcoffeen",
		},
	},
	"MWG": {
		{
			ID:   8505,
			Slug: "metawhale-gold",
			Name: "Metawhale Gold",
		},
	},
	"CYC": {
		{
			ID:   8590,
			Slug: "cyclone-protocol",
			Name: "Cyclone Protocol",
		},
	},
	"PMA": {
		{
			ID:   3164,
			Slug: "pumapay",
			Name: "PumaPay",
		},
	},
	"ARGO": {
		{
			ID:   9663,
			Slug: "argoapp",
			Name: "ArGo",
		},
	},
	"BCX": {
		{
			ID:   2281,
			Slug: "bitcoinx",
			Name: "BitcoinX",
		},
	},
	"BARMY": {
		{
			ID:   10340,
			Slug: "bscarmy",
			Name: "BscArmy",
		},
	},
	"SYLO": {
		{
			ID:   5662,
			Slug: "sylo",
			Name: "Sylo",
		},
	},
	"CHONK": {
		{
			ID:   7487,
			Slug: "chonk",
			Name: "Chonk",
		},
	},
	"DGC": {
		{
			ID:   18,
			Slug: "digitalcoin",
			Name: "Digitalcoin",
		},
	},
	"VOYRME": {
		{
			ID:   10690,
			Slug: "voyr",
			Name: "VOYR",
		},
	},
	"FR": {
		{
			ID:   7712,
			Slug: "freedom-reserve",
			Name: "Freedom Reserve",
		},
	},
	"REM": {
		{
			ID:   2546,
			Slug: "remme",
			Name: "Remme",
		},
	},
	"MCHC": {
		{
			ID:   9686,
			Slug: "my-crypto-heroes",
			Name: "My Crypto Heroes",
		},
	},
	"ARCONA": {
		{
			ID:   6218,
			Slug: "arcona",
			Name: "Arcona",
		},
	},
	"AIR": {
		{
			ID:   10715,
			Slug: "air",
			Name: "AirCoin",
		},
	},
	"SME": {
		{
			ID:   10153,
			Slug: "safememe",
			Name: "SafeMeme",
		},
	},
	"GLQ": {
		{
			ID:   9029,
			Slug: "graphlinq-protocol",
			Name: "Graphlinq Protocol",
		},
	},
	"XMX": {
		{
			ID:   2859,
			Slug: "xmax",
			Name: "XMax",
		},
	},
	"KURT": {
		{
			ID:   1468,
			Slug: "kurrent",
			Name: "Kurrent",
		},
	},
	"ATH": {
		{
			ID:   3447,
			Slug: "atheios",
			Name: "Atheios",
		},
	},
	"BNBDOWN": {
		{
			ID:   7010,
			Slug: "bnbdown",
			Name: "BNBDOWN",
		},
	},
	"VRX": {
		{
			ID:   8278,
			Slug: "verox",
			Name: "VEROX",
		},
	},
	"PEARL": {
		{
			ID:   6829,
			Slug: "pearl",
			Name: "Pearl",
		},
	},
	"UNCL": {
		{
			ID:   7669,
			Slug: "uncl",
			Name: "UNCL",
		},
	},
	"ASTRA": {
		{
			ID:   10621,
			Slug: "astra-coin",
			Name: "Astra Coin",
		},
	},
	"PSL": {
		{
			ID:   8738,
			Slug: "pastel",
			Name: "Pastel",
		},
	},
	"MOZ": {
		{
			ID:   9802,
			Slug: "mozik",
			Name: "Mozik",
		},
	},
	"VJC": {
		{
			ID:   3923,
			Slug: "venjocoin",
			Name: "VENJOCOIN",
		},
	},
	"RNO": {
		{
			ID:   3679,
			Slug: "earneo",
			Name: "Earneo",
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
	"MNG": {
		{
			ID:   10456,
			Slug: "moon-nation-game",
			Name: "Moon Nation Game",
		},
	},
	"SHOPX": {
		{
			ID:   8863,
			Slug: "splyt",
			Name: "Splyt",
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
	"MLGC": {
		{
			ID:   4892,
			Slug: "marshal-lion-group-coin",
			Name: "Marshal Lion Group Coin",
		},
	},
	"ICX": {
		{
			ID:   2099,
			Slug: "icon",
			Name: "ICON",
		},
	},
	"BYN": {
		{
			ID:   9134,
			Slug: "beyond-finance",
			Name: "Beyond Finance",
		},
	},
	"SYNC": {
		{
			ID:   7812,
			Slug: "sync-network",
			Name: "SYNC Network",
		},
	},
	"NSFW": {
		{
			ID:   9840,
			Slug: "xxxnifty",
			Name: "xxxNifty",
		},
	},
	"$BITE": {
		{
			ID:   10314,
			Slug: "dragonbite",
			Name: "DragonBite",
		},
	},
	"XAT": {
		{
			ID:   7488,
			Slug: "shareat",
			Name: "ShareAt",
		},
	},
	"AMP": {
		{
			ID:   6945,
			Slug: "amp",
			Name: "Amp",
		},
	},
	"NLG": {
		{
			ID:   254,
			Slug: "gulden",
			Name: "Gulden",
		},
	},
	"BREE": {
		{
			ID:   6502,
			Slug: "cbdao",
			Name: "CBDAO",
		},
	},
	"NTR": {
		{
			ID:   3950,
			Slug: "netrum",
			Name: "Netrum",
		},
	},
	"G999": {
		{
			ID:   8431,
			Slug: "g999",
			Name: "G999",
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
	"CELC": {
		{
			ID:   5388,
			Slug: "celcoin",
			Name: "CelCoin",
		},
	},
	"DOT": {
		{
			ID:   6636,
			Slug: "polkadot-new",
			Name: "Polkadot",
		},
	},
	"CRM": {
		{
			ID:   1850,
			Slug: "cream",
			Name: "Cream",
		},
	},
	"YFIDOWN": {
		{
			ID:   7453,
			Slug: "yfidown",
			Name: "YFIDOWN",
		},
	},
	"VIDYX": {
		{
			ID:   8182,
			Slug: "vidyx",
			Name: "VidyX",
		},
	},
	"KHC": {
		{
			ID:   10246,
			Slug: "koho-chain",
			Name: "KoHo Chain",
		},
	},
	"PBTC": {
		{
			ID:   5434,
			Slug: "ptokens-btc",
			Name: "pTokens BTC",
		},
	},
	"LYFE": {
		{
			ID:   5786,
			Slug: "lyfe",
			Name: "LYFE",
		},
	},
	"DOGEDAO": {
		{
			ID:   9804,
			Slug: "dogedao-finance",
			Name: "DogeDao Finance",
		},
	},
	"SKC": {
		{
			ID:   5533,
			Slug: "skinchain",
			Name: "SKINCHAIN",
		},
	},
	"COLDKOALA": {
		{
			ID:   10383,
			Slug: "cold-koala",
			Name: "Cold Koala",
		},
	},
	"YIELD": {
		{
			ID:   7498,
			Slug: "yield-protocol",
			Name: "Yield Protocol",
		},
	},
	"XEC": {
		{
			ID:   10791,
			Slug: "ecash",
			Name: "eCash",
		},
	},
	"FND": {
		{
			ID:   9901,
			Slug: "fundum-capital",
			Name: "Fundum Capital",
		},
	},
	"BURGER": {
		{
			ID:   7158,
			Slug: "burger-swap",
			Name: "Burger Swap",
		},
	},
	"SMART": {
		{
			ID:   1828,
			Slug: "smartcash",
			Name: "SmartCash",
		},
	},
	"STEEL": {
		{
			ID:   10639,
			Slug: "steel",
			Name: "Steel",
		},
	},
	"NODOGE": {
		{
			ID:   10823,
			Slug: "no-doge",
			Name: "NO DOGE",
		},
	},
	"DPET": {
		{
			ID:   9665,
			Slug: "my-defi-pet",
			Name: "My DeFi Pet",
		},
	},
	"WEBD": {
		{
			ID:   3877,
			Slug: "webdollar",
			Name: "WebDollar",
		},
	},
	"FO": {
		{
			ID:   4058,
			Slug: "fibos",
			Name: "FIBOS",
		},
	},
	"FSBT": {
		{
			ID:   2884,
			Slug: "fsbt-api-token",
			Name: "FSBT API Token",
		},
	},
	"TREES": {
		{
			ID:   10512,
			Slug: "safetrees",
			Name: "SAFETREES",
		},
	},
	"IMM": {
		{
			ID:   10450,
			Slug: "imm",
			Name: "IMM",
		},
	},
	"XIOT": {
		{
			ID:   6667,
			Slug: "xiotri",
			Name: "Xiotri",
		},
	},
	"GROOT": {
		{
			ID:   8604,
			Slug: "growth-root-token",
			Name: "growth Root Token",
		},
	},
	"LYR": {
		{
			ID:   7762,
			Slug: "lyra",
			Name: "Lyra",
		},
	},
	"BHD": {
		{
			ID:   3966,
			Slug: "bitcoinhd",
			Name: "BitcoinHD",
		},
	},
	"STARK": {
		{
			ID:   7610,
			Slug: "stark-chain",
			Name: "STARK CHAIN",
		},
	},
	"UC": {
		{
			ID:   3259,
			Slug: "youlive-coin",
			Name: "YouLive Coin",
		},
	},
	"MATPAD": {
		{
			ID:   10696,
			Slug: "maticpad",
			Name: "MaticPad",
		},
	},
	"FXS": {
		{
			ID:   6953,
			Slug: "frax-share",
			Name: "Frax Share",
		},
	},
	"S": {
		{
			ID:   3423,
			Slug: "sharpay",
			Name: "Sharpay",
		},
	},
	"FLEX": {
		{
			ID:   5190,
			Slug: "flex",
			Name: "FLEX",
		},
	},
	"HIBS": {
		{
			ID:   6470,
			Slug: "hiblocks",
			Name: "Hiblocks",
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
	"DFR": {
		{
			ID:   8116,
			Slug: "diffract-finance",
			Name: "Diffract Finance",
		},
	},
	"YVS": {
		{
			ID:   8036,
			Slug: "yvs-finance",
			Name: "YVS.Finance",
		},
	},
	"GSHIBA": {
		{
			ID:   11043,
			Slug: "gambler-shiba",
			Name: "Gambler Shiba",
		},
	},
	"TTX": {
		{
			ID:   7392,
			Slug: "talent-token",
			Name: "Talent Token",
		},
	},
	"INTRATIO": {
		{
			ID:   6155,
			Slug: "intelligent-ratio-set",
			Name: "Intelligent Ratio Set",
		},
	},
	"SRSB": {
		{
			ID:   10025,
			Slug: "sirius-bond",
			Name: "Sirius Bond",
		},
	},
	"CASHDOG": {
		{
			ID:   10600,
			Slug: "cashdog",
			Name: "CashDog",
		},
	},
	"BSCS": {
		{
			ID:   9345,
			Slug: "bsc-station",
			Name: "BSC Station",
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
	"VINCI": {
		{
			ID:   4946,
			Slug: "vinci",
			Name: "Vinci",
		},
	},
	"sLINK": {
		{
			ID:   6190,
			Slug: "slink",
			Name: "sLINK",
		},
	},
	"YUSRA": {
		{
			ID:   6726,
			Slug: "yusra",
			Name: "YUSRA",
		},
	},
	"TCAP": {
		{
			ID:   9678,
			Slug: "total-crypto-market-cap-token",
			Name: "Total Crypto Market Cap Token",
		},
	},
	"BEPRO": {
		{
			ID:   5062,
			Slug: "bepro-network",
			Name: "BEPRO Network",
		},
	},
	"DIS": {
		{
			ID:   8310,
			Slug: "tosdis",
			Name: "TosDis",
		},
	},
	"B1P": {
		{
			ID:   5383,
			Slug: "b-one-payment",
			Name: "B ONE PAYMENT",
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
	"BIP": {
		{
			ID:   4957,
			Slug: "minter-network",
			Name: "Minter Network",
		},
	},
	"SWFTC": {
		{
			ID:   2341,
			Slug: "swftcoin",
			Name: "SwftCoin",
		},
	},
	"GDAO": {
		{
			ID:   7694,
			Slug: "governor-dao",
			Name: "Governor DAO",
		},
	},
	"HGOLD": {
		{
			ID:   8256,
			Slug: "hollygold",
			Name: "HollyGold",
		},
	},
	"SPHTX": {
		{
			ID:   2309,
			Slug: "sophiatx",
			Name: "SophiaTX",
		},
	},
	"RNT": {
		{
			ID:   2400,
			Slug: "oneroot-network",
			Name: "OneRoot Network",
		},
	},
	"IDT": {
		{
			ID:   2406,
			Slug: "investdigital",
			Name: "InvestDigital",
		},
	},
	"PANDA": {
		{
			ID:   10131,
			Slug: "hashpanda",
			Name: "HashPanda",
		},
	},
	"DIAMOND": {
		{
			ID:   9650,
			Slug: "diamondtoken",
			Name: "DiamondToken",
		},
	},
	"QBX": {
		{
			ID:   4100,
			Slug: "qiibee",
			Name: "qiibee",
		},
	},
	"VBETH": {
		{
			ID:   8370,
			Slug: "venus-beth",
			Name: "Venus BETH",
		},
	},
	"SLEEPY": {
		{
			ID:   9774,
			Slug: "sleepy-sloth-finance",
			Name: "Sleepy Sloth Finance",
		},
	},
	"NEO": {
		{
			ID:   1376,
			Slug: "neo",
			Name: "Neo",
		},
	},
	"WOWS": {
		{
			ID:   8496,
			Slug: "wolves-of-wall-street",
			Name: "Wolves of Wall Street",
		},
	},
	"OMN": {
		{
			ID:   10944,
			Slug: "omni-token",
			Name: "OMNI - People Driven",
		},
	},
	"NIIFI": {
		{
			ID:   9825,
			Slug: "niifi",
			Name: "NiiFi",
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
	"R3FI": {
		{
			ID:   8313,
			Slug: "r3fi-finance",
			Name: "Recharge Finance",
		},
	},
	"ROOT": {
		{
			ID:   8077,
			Slug: "rootkit-finance",
			Name: "Rootkit Finance",
		},
	},
	"ASK": {
		{
			ID:   7105,
			Slug: "permission-coin",
			Name: "Permission Coin",
		},
	},
	"XAS": {
		{
			ID:   1609,
			Slug: "asch",
			Name: "Asch",
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
	"MP4": {
		{
			ID:   8597,
			Slug: "mp4",
			Name: "MP4",
		},
	},
	"AUCTION": {
		{
			ID:   8602,
			Slug: "bounce-token",
			Name: "Bounce Token",
		},
	},
	"ZEBI": {
		{
			ID:   2685,
			Slug: "zebi-token",
			Name: "Zebi Token",
		},
	},
	"YAE": {
		{
			ID:   9149,
			Slug: "cryptonovae",
			Name: "Cryptonovae",
		},
	},
	"DVP": {
		{
			ID:   4520,
			Slug: "decentralized-vulnerability-platform",
			Name: "Decentralized Vulnerability Platform",
		},
	},
	"TOTM": {
		{
			ID:   8673,
			Slug: "totemfi",
			Name: "TotemFi",
		},
	},
	"NEKO": {
		{
			ID:   9963,
			Slug: "neko-network",
			Name: "Neko Network",
		},
	},
	"LEC": {
		{
			ID:   10779,
			Slug: "love-earth-coin",
			Name: "LOVE EARTH COIN",
		},
	},
	"HD": {
		{
			ID:   7541,
			Slug: "hubdao",
			Name: "HubDao",
		},
	},
	"HC": {
		{
			ID:   1903,
			Slug: "hypercash",
			Name: "HyperCash",
		},
	},
	"TSL": {
		{
			ID:   2215,
			Slug: "energo",
			Name: "Energo",
		},
	},
	"EOSDOWN": {
		{
			ID:   7000,
			Slug: "eosdown",
			Name: "EOSDOWN",
		},
	},
	"ISDT": {
		{
			ID:   7148,
			Slug: "istardust",
			Name: "ISTARDUST",
		},
	},
	"STRONG": {
		{
			ID:   6511,
			Slug: "strong",
			Name: "Strong",
		},
	},
	"BIOT": {
		{
			ID:   8034,
			Slug: "biopassport-token",
			Name: "BioPassport Token",
		},
	},
	"UCA": {
		{
			ID:   5479,
			Slug: "uca-coin",
			Name: "UCA Coin",
		},
	},
	"LDFI": {
		{
			ID:   9341,
			Slug: "lendefi",
			Name: "Lendefi",
		},
	},
	"FORESTPLUS": {
		{
			ID:   4848,
			Slug: "the-forbidden-forest",
			Name: "The Forbidden Forest",
		},
	},
	"yVault LP-yCurve(YYCRV)": {
		{
			ID:   6686,
			Slug: "yvault-lp-ycurve",
			Name: "yVault LP-yCurve",
		},
	},
	"VGW": {
		{
			ID:   3735,
			Slug: "vegawallet-token",
			Name: "VegaWallet Token",
		},
	},
	"DAM": {
		{
			ID:   2260,
			Slug: "datamine",
			Name: "Datamine",
		},
	},
	"BMH": {
		{
			ID:   2747,
			Slug: "blockmesh",
			Name: "BlockMesh",
		},
	},
	"ETHO": {
		{
			ID:   3452,
			Slug: "ether-1",
			Name: "Etho Protocol",
		},
	},
	"TRO": {
		{
			ID:   8636,
			Slug: "trodl",
			Name: "Trodl",
		},
	},
	"WEB": {
		{
			ID:   3027,
			Slug: "webcoin",
			Name: "Webcoin",
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
	"NZL": {
		{
			ID:   3624,
			Slug: "zealium",
			Name: "Zealium",
		},
	},
	"MIMATIC": {
		{
			ID:   10238,
			Slug: "mimatic",
			Name: "miMatic",
		},
	},
	"NABOX": {
		{
			ID:   9653,
			Slug: "nabox",
			Name: "Nabox",
		},
	},
	"EQO": {
		{
			ID:   9593,
			Slug: "equos-io",
			Name: "EQO",
		},
	},
	"ADL": {
		{
			ID:   1725,
			Slug: "adelphoi",
			Name: "Adelphoi",
		},
	},
	"ARAW": {
		{
			ID:   3792,
			Slug: "araw",
			Name: "ARAW",
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
	"AUSCM": {
		{
			ID:   7583,
			Slug: "auric-network",
			Name: "Auric Network",
		},
	},
	"BSP": {
		{
			ID:   8566,
			Slug: "ballswap",
			Name: "Ballswap",
		},
	},
	"COP": {
		{
			ID:   9763,
			Slug: "copiosa-coin",
			Name: "Copiosa Coin",
		},
	},
	"ECHO": {
		{
			ID:   9979,
			Slug: "echelon-dao",
			Name: "Echelon DAO",
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
	"MCB": {
		{
			ID:   5956,
			Slug: "mcdex",
			Name: "MCDEX",
		},
	},
	"MUSUBI": {
		{
			ID:   10708,
			Slug: "musubi",
			Name: "Musubi",
		},
	},
	"DGN": {
		{
			ID:   8667,
			Slug: "degen-protocol",
			Name: "Degen Protocol",
		},
	},
	"ZB": {
		{
			ID:   3351,
			Slug: "zb-token",
			Name: "ZB Token",
		},
	},
	"ΤBTC": {
		{
			ID:   9114,
			Slug: "t-bitcoin",
			Name: "τBitcoin",
		},
	},
	"XAUR": {
		{
			ID:   895,
			Slug: "xaurum",
			Name: "Xaurum",
		},
	},
	"YFID": {
		{
			ID:   7880,
			Slug: "yfidapp",
			Name: "YFIDapp",
		},
	},
	"VIP": {
		{
			ID:   719,
			Slug: "limitless-vip",
			Name: "Limitless VIP",
		},
	},
	"FOL": {
		{
			ID:   7687,
			Slug: "folder-protocol",
			Name: "Folder Protocol",
		},
	},
	"CSTL": {
		{
			ID:   3311,
			Slug: "castle",
			Name: "Castle",
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
	"KRL": {
		{
			ID:   2949,
			Slug: "kryll",
			Name: "Kryll",
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
	"VIDZ": {
		{
			ID:   1511,
			Slug: "purevidz",
			Name: "PureVidz",
		},
	},
	"KOMBAI": {
		{
			ID:   10319,
			Slug: "kombai-inu",
			Name: "Kombai Inu",
		},
	},
	"USDEX": {
		{
			ID:   8348,
			Slug: "usdex",
			Name: "USDEX",
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
	"mBABA": {
		{
			ID:   8006,
			Slug: "mirrored-alibaba",
			Name: "Mirrored Alibaba",
		},
	},
	"FOR": {
		{
			ID:   4118,
			Slug: "the-force-protocol",
			Name: "ForTube",
		},
	},
	"BUT": {
		{
			ID:   3258,
			Slug: "bitup-token",
			Name: "BitUP Token",
		},
	},
	"CHORD": {
		{
			ID:   9199,
			Slug: "chord-protocol",
			Name: "Chord Protocol",
		},
	},
	"YFI2": {
		{
			ID:   6973,
			Slug: "yearn2-finance",
			Name: "YEARN2.FINANCE",
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
	"DGX": {
		{
			ID:   2739,
			Slug: "digix-gold-token",
			Name: "Digix Gold Token",
		},
	},
	"UT": {
		{
			ID:   3233,
			Slug: "ulord",
			Name: "Ulord",
		},
	},
	"ICE": {
		{
			ID:   9073,
			Slug: "popsicle-finance",
			Name: "Popsicle Finance",
		},
	},
	"MANA": {
		{
			ID:   1966,
			Slug: "decentraland",
			Name: "Decentraland",
		},
	},
	"ORBS": {
		{
			ID:   3835,
			Slug: "orbs",
			Name: "Orbs",
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
	"SPHR": {
		{
			ID:   914,
			Slug: "sphere",
			Name: "Sphere",
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
	"BSYS": {
		{
			ID:   6391,
			Slug: "bsys",
			Name: "BSYS",
		},
	},
	"BR": {
		{
			ID:   10624,
			Slug: "bull-run-finance",
			Name: "Bull Run Finance",
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
	"SFUEL": {
		{
			ID:   8145,
			Slug: "sparkpoint-fuel",
			Name: "SparkPoint Fuel",
		},
	},
	"SST": {
		{
			ID:   5987,
			Slug: "simba-storage-token",
			Name: "SIMBA Storage Token",
		},
	},
	"VEO": {
		{
			ID:   3716,
			Slug: "amoveo",
			Name: "Amoveo",
		},
	},
	"SNT": {
		{
			ID:   1759,
			Slug: "status",
			Name: "Status",
		},
	},
	"CS": {
		{
			ID:   2556,
			Slug: "credits",
			Name: "Credits",
		},
	},
	"MPL": {
		{
			ID:   9417,
			Slug: "maple",
			Name: "Maple",
		},
	},
	"BON": {
		{
			ID:   2256,
			Slug: "bonpay",
			Name: "Bonpay",
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
	"AME": {
		{
			ID:   8162,
			Slug: "amepay",
			Name: "AMEPAY",
		},
	},
	"CUE": {
		{
			ID:   9121,
			Slug: "cue-protocol",
			Name: "CUE Protocol",
		},
	},
	"NCC": {
		{
			ID:   3016,
			Slug: "neurochain",
			Name: "NeuroChain",
		},
	},
	"SMGM": {
		{
			ID:   10867,
			Slug: "smegmars",
			Name: "SMEGMARS",
		},
	},
	"LRG": {
		{
			ID:   5329,
			Slug: "largo-coin",
			Name: "Largo Coin",
		},
	},
	"ALIA": {
		{
			ID:   9099,
			Slug: "xanalia",
			Name: "XANALIA",
		},
	},
	"DIVI": {
		{
			ID:   3441,
			Slug: "divi",
			Name: "Divi",
		},
	},
	"VEST": {
		{
			ID:   3607,
			Slug: "vestchain",
			Name: "VestChain",
		},
	},
	"ZUT": {
		{
			ID:   7318,
			Slug: "zero-utility-token",
			Name: "Zero Utility Token",
		},
	},
	"HDI": {
		{
			ID:   5797,
			Slug: "heidi",
			Name: "HEIDI",
		},
	},
	"GLOX": {
		{
			ID:   7886,
			Slug: "glox-finance",
			Name: "Glox Finance",
		},
	},
	"LAMB": {
		{
			ID:   3657,
			Slug: "lambda",
			Name: "Lambda",
		},
	},
	"UNIDOWN": {
		{
			ID:   7525,
			Slug: "unidown",
			Name: "UNIDOWN",
		},
	},
	"DEXI": {
		{
			ID:   9830,
			Slug: "dexioprotocol",
			Name: "Dexioprotocol",
		},
	},
	"BUL": {
		{
			ID:   3690,
			Slug: "bulleon",
			Name: "Bulleon",
		},
	},
	"PAPADOGE": {
		{
			ID:   11045,
			Slug: "papa-doge-coin",
			Name: "Papa Doge Coin",
		},
	},
	"WIKI": {
		{
			ID:   3090,
			Slug: "wiki-token",
			Name: "Wiki Token",
		},
	},
	"PIRATE": {
		{
			ID:   4460,
			Slug: "piratecash",
			Name: "PirateCash",
		},
	},
	"DPI": {
		{
			ID:   7055,
			Slug: "defi-pulse-index",
			Name: "DeFi Pulse Index",
		},
	},
	"STARLINKDOGE": {
		{
			ID:   10871,
			Slug: "baby-starlink-doge",
			Name: "Baby Starlink Doge",
		},
	},
	"SBEAR": {
		{
			ID:   9926,
			Slug: "ybearswap",
			Name: "yBEARSwap",
		},
	},
	"BCHC": {
		{
			ID:   5803,
			Slug: "bitcherry",
			Name: "BitCherry",
		},
	},
	"IFLT": {
		{
			ID:   1504,
			Slug: "inflationcoin",
			Name: "InflationCoin",
		},
	},
	"WRX": {
		{
			ID:   5161,
			Slug: "wazirx",
			Name: "WazirX",
		},
	},
	"FOXX": {
		{
			ID:   10387,
			Slug: "star-foxx",
			Name: "Star Foxx",
		},
	},
	"RENASCENT": {
		{
			ID:   9748,
			Slug: "renascent-finance",
			Name: "Renascent Finance",
		},
	},
	"xBLZD": {
		{
			ID:   8964,
			Slug: "blizzard-money",
			Name: "Blizzard.money",
		},
	},
	"ETHRSIAPY": {
		{
			ID:   6144,
			Slug: "eth-rsi-60-40-yield-set",
			Name: "ETH RSI 60/40 Yield Set",
		},
	},
}
