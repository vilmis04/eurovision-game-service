import { Deprecated_CountryTypes as CountryTypes } from "@eurovision-game-monorepo/core";

type TCountryInfo = {
	artist: string;
	song: string;
};

type TCountryData = {
	[key in CountryTypes]: TCountryInfo;
};
// TODO: move to database? Maybe add data via admin portal for the future
export const deprecated_countryData: TCountryData = {
	[CountryTypes.UKRAINE]: {
		artist: "Kalush",
		song: "Stefania",
	},
	[CountryTypes.UNITED_KINGDOM]: {
		artist: "Sam Ryder",
		song: "Space Man",
	},
	[CountryTypes.SPAIN]: {
		artist: "Chanel",
		song: "SloMo",
	},
	[CountryTypes.SWEDEN]: {
		artist: "Cornelia Jakobs",
		song: "Hold Me Closer",
	},
	[CountryTypes.SERBIA]: {
		artist: "Konstranta",
		song: "In corpre sano",
	},
	[CountryTypes.NORWAY]: {
		artist: "Subwoolfer",
		song: "Give That Wolf a Banana",
	},
	[CountryTypes.BELGIUM]: {
		artist: "Jérémie Makiese",
		song: "Miss You",
	},
	[CountryTypes.CZECH_REPUBLIC]: {
		artist: "We Are Domi",
		song: "Lights Off",
	},
	[CountryTypes.ROMANIA]: {
		artist: "WRS",
		song: "Llámame",
	},
	[CountryTypes.PORTUGAL]: {
		artist: "MARO",
		song: "Saudade, Saudade",
	},
	[CountryTypes.FINLAND]: {
		artist: "The Rasmus",
		song: "Jezebel",
	},
	[CountryTypes.SWITZERLAND]: {
		artist: "Marius Bear",
		song: "Boys Do Cry",
	},
	[CountryTypes.FRANCE]: {
		artist: "Alvan & Ahez",
		song: "Fulenn",
	},
	[CountryTypes.ARMENIA]: {
		artist: "Rosa Linn",
		song: "Snap",
	},
	[CountryTypes.ITALY]: {
		artist: "Mahmood & Blanco",
		song: "Brividi",
	},
	[CountryTypes.NETHERLANDS]: {
		artist: "S10",
		song: "De Diepte",
	},
	[CountryTypes.GERMANY]: {
		artist: "Malik Harris",
		song: "Rockstars",
	},
	[CountryTypes.LITHUANIA]: {
		artist: "Monika Liu",
		song: "Sentimentai",
	},
	[CountryTypes.AZERBAIJAN]: {
		artist: "Nadir Rustamli",
		song: "Fade To Black",
	},
	[CountryTypes.GREECE]: {
		artist: "Amanda Georgiadi Tenfjord",
		song: "Die Together",
	},
	[CountryTypes.ICELAND]: {
		artist: "Systur",
		song: "Með Hækkandi Sól",
	},
	[CountryTypes.MOLDOVA]: {
		artist: "Zdob şi Zdub & Advahov Brothers",
		song: "Trenulețul",
	},
	[CountryTypes.AUSTRALIA]: {
		artist: "Sheldon Riley",
		song: "Not The Same",
	},
	[CountryTypes.POLAND]: {
		artist: "Ochman",
		song: "River",
	},
	[CountryTypes.ESTONIA]: {
		artist: "Stefan",
		song: "Hope",
	},
};
