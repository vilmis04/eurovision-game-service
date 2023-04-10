import { Deprecated_CountryTypes } from "../deprecated_country.types";
import { VoteTypes } from "./votes.types";

export interface IGetVotesResponse {
	username?: string;
	[Deprecated_CountryTypes.ARMENIA]: VoteTypes | null;
	[Deprecated_CountryTypes.AUSTRALIA]: VoteTypes | null;
	[Deprecated_CountryTypes.AZERBAIJAN]: VoteTypes | null;
	[Deprecated_CountryTypes.BELGIUM]: VoteTypes | null;
	[Deprecated_CountryTypes.CZECH_REPUBLIC]: VoteTypes | null;
	[Deprecated_CountryTypes.ESTONIA]: VoteTypes | null;
	[Deprecated_CountryTypes.FINLAND]: VoteTypes | null;
	[Deprecated_CountryTypes.FRANCE]: VoteTypes | null;
	[Deprecated_CountryTypes.GERMANY]: VoteTypes | null;
	[Deprecated_CountryTypes.GREECE]: VoteTypes | null;
	[Deprecated_CountryTypes.ICELAND]: VoteTypes | null;
	[Deprecated_CountryTypes.ITALY]: VoteTypes | null;
	[Deprecated_CountryTypes.LITHUANIA]: VoteTypes | null;
	[Deprecated_CountryTypes.MOLDOVA]: VoteTypes | null;
	[Deprecated_CountryTypes.NETHERLANDS]: VoteTypes | null;
	[Deprecated_CountryTypes.NORWAY]: VoteTypes | null;
	[Deprecated_CountryTypes.POLAND]: VoteTypes | null;
	[Deprecated_CountryTypes.PORTUGAL]: VoteTypes | null;
	[Deprecated_CountryTypes.ROMANIA]: VoteTypes | null;
	[Deprecated_CountryTypes.SERBIA]: VoteTypes | null;
	[Deprecated_CountryTypes.SPAIN]: VoteTypes | null;
	[Deprecated_CountryTypes.SWEDEN]: VoteTypes | null;
	[Deprecated_CountryTypes.SWITZERLAND]: VoteTypes | null;
	[Deprecated_CountryTypes.UKRAINE]: VoteTypes | null;
	[Deprecated_CountryTypes.UNITED_KINGDOM]: VoteTypes | null;
}
