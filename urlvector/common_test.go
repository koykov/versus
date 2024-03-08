package urlvector

import "testing"

var (
	url0         = []byte("http://user:pass@[1080:0:0:0:8:800:200C:417A]:61616/foo?v=default&blockID=319385&page=https%3A%2F%2Fultra-software-base.ru%2Fsystem%2Fgoogle-chrome.html%3Fyclid%3D212247430717539672&domain=ultra-software-base.ru&uid=4f5d0edc-3a3e-48d0-9872-0b48a7998ac6&clientNotice=true&imgX=360&imgY=240&limit=1&subage_dt=2021-01-29&format=json&cur=RUB&ua=Mozilla%2F5.0+%28Windows+NT+6.1%3B+Win64%3B+x64%29+AppleWebKit%2F537.36+%28KHTML%2C+like+Gecko%29+Chrome%2F89.0.4389.105+YaBrowser%2F21.3.3.230+Yowser%2F2.5+Safari%2F537.36&ip=5.5.5.5&subage=102&language=ru#foobar")
	url0scheme   = "http"
	url0user     = "user"
	url0pass     = "pass"
	url0host     = "[1080:0:0:0:8:800:200C:417A]:61616"
	url0path     = "/foo"
	url0hash     = "#foobar"
	url0fragment = "foobar"
	url0qPage    = "https://ultra-software-base.ru/system/google-chrome.html?yclid=212247430717539672"
	url0qUa      = "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.105 YaBrowser/21.3.3.230 Yowser/2.5 Safari/537.36"
	url0qUid     = "4f5d0edc-3a3e-48d0-9872-0b48a7998ac6"
	url0qIp      = "5.5.5.5"
	url0qIp1     = "127.0.0.1"

	url0modHost  = "http://user:pass@x.com:3333/foo?v=default&blockID=319385&page=https%3A%2F%2Fultra-software-base.ru%2Fsystem%2Fgoogle-chrome.html%3Fyclid%3D212247430717539672&domain=ultra-software-base.ru&uid=4f5d0edc-3a3e-48d0-9872-0b48a7998ac6&clientNotice=true&imgX=360&imgY=240&limit=1&subage_dt=2021-01-29&format=json&cur=RUB&ua=Mozilla%2F5.0+%28Windows+NT+6.1%3B+Win64%3B+x64%29+AppleWebKit%2F537.36+%28KHTML%2C+like+Gecko%29+Chrome%2F89.0.4389.105+YaBrowser%2F21.3.3.230+Yowser%2F2.5+Safari%2F537.36&ip=5.5.5.5&subage=102&language=ru#foobar"
	url0newQuery = "?v=default&blockID=319385&page=https%3A%2F%2Fultra-software-base.ru%2Fsystem%2Fgoogle-chrome.html%3Fyclid%3D212247430717539672&domain=ultra-software-base.ru&uid=4f5d0edc-3a3e-48d0-9872-0b48a7998ac6&clientNotice=true&imgX=360&imgY=240&limit=1&subage_dt=2021-01-29&format=json&cur=RUB&ua=Mozilla%2F5.0+%28Windows+NT+6.1%3B+Win64%3B+x64%29+AppleWebKit%2F537.36+%28KHTML%2C+like+Gecko%29+Chrome%2F89.0.4389.105+YaBrowser%2F21.3.3.230+Yowser%2F2.5+Safari%2F537.36&ip=127.0.0.1&subage=102&language=ru"
	url0modQuery = "http://user:pass@[1080:0:0:0:8:800:200C:417A]:61616/foo?v=default&blockID=319385&page=https%3A%2F%2Fultra-software-base.ru%2Fsystem%2Fgoogle-chrome.html%3Fyclid%3D212247430717539672&domain=ultra-software-base.ru&uid=4f5d0edc-3a3e-48d0-9872-0b48a7998ac6&clientNotice=true&imgX=360&imgY=240&limit=1&subage_dt=2021-01-29&format=json&cur=RUB&ua=Mozilla%2F5.0+%28Windows+NT+6.1%3B+Win64%3B+x64%29+AppleWebKit%2F537.36+%28KHTML%2C+like+Gecko%29+Chrome%2F89.0.4389.105+YaBrowser%2F21.3.3.230+Yowser%2F2.5+Safari%2F537.36&ip=127.0.0.1&subage=102&language=ru#foobar"
)

var stages = []string{
	"https://google.com/?foo=bar",
	"https://www.mozilla.org/en-US/firefox/34.0/whatsnew/?oldversion=33.1",
	"https://www.msn.com/ru-ru/lifestyle/travel/на-фото-памятные-достопримечательности-из-разных-уголков-планеты/ss-AAGnFe0#image=4?ocid=ems.msn.dl.090919.TowerOfPisaItaly",
	"https://lamimyde.pro/bB3.VC0DPE2_lGjHPIXJB-zLJMmN9O0_PQUR5SETR-TVJWNXRYE_Va5bTcVdN-RfegEh5iq_SkTlBmNna-mpMqxrTsn_puVvMwUxp-FzbAFBVCK_RE0FpG5HY-jJJKGLaM1_lOtPRQnRV-aTQU1VFWr_SYkZRaKba-ldpeXfUgT_BiOjVk0l0-0nWoWppqC_as0t1uUvZ-GxxyPzRAG_dC3DTEkFR-JHMIVJpKt_WMmN1ONPR-ER5SpTTUX_pWRXeYEZ9-XbUcTdAem_cgnhJipjZ-Dl0mznYo2_Mq3rZsWtQ-zvMwjxUyx_NAGBYC2DZ-DFEG3HMIj_UK3LOMDNg-5PMQDRlSj_NUDVFWiXY-zZUa2bNcy_Zeyfcg3hJ-jjPkXlNmj_complqwrd-CtZu2vdwD_0yxzNAjBI-0DMEjFcG1_NIzJUK1LL-TNVOkPZQG_YS1TYUTVV-hXZYWZUa5_NczdgexfZ-WhRimjMkz_FmjnNoWpJ-krMsTtQu5_ZwWxIy5zY-TBlCjD?bd_vs=1.0.2&bd_t0=1624275754931&bd_a0=a0,a1,a2,a3,a4,a5,a6,a7,a8,a9,a10,a11,a12,a13,a14,a15,a16,a17,a18,a19,a20,a21,a22,a23,a24,a25,a26,a27,a28,a29,a30,a31,a32,a33,a34,a35,a36,a37,a38,a39,a40,a41,a42,a43,a44,a45,a46,a47,a48,a49,a50,a51,a52,a53,a54,a55,a56,a57,sm0&bd_a1=59&bd_n0=Mozilla/5.0 (Linux; Android 9; SAMSUNG SM-G950F/G950FXXUCDUD1) AppleWebKit/537.36 (KHTML, like Gecko) SamsungBrowser/14.0 Chrome/87.0.4280.141 Mobile Safari/537.36&bd_n1=[\\\"it-IT\\\",\\\"it\\\",\\\"en-US\\\",\\\"en\\\"]&bd_n2=8&bd_n3=20030107&bd_n4=Google Inc.&bd_o1=false&bd_o4=&bd_o5=false&sseq=3&dseq=3&ce=lnk",
	"https://hotels.andcastles.andhouseboats.andigloos.andteepees.andriversidecabins.andlakesidecabins.andpondsidecabins.andstreamadjacentcabins.andcabinsthatarentnearanybodiesofwaterwhatsoever.andlakehouses.andregularhousesandlodgesandskilodgesandallthings.ski/ChaletRelatedAndBoutiquesAnd5StarSuitesAndRetreatsAndBungalowsAndJungleBungalowsAndOtherKindaLessExcitingBungalowsAndCabanasAndOceansideCabanasAndSeaSideCabanasWhichSeemLikeTheSameThingAndBedAndBreakfastsAndJustBedsBecauseThoseAreKindOfARequirementInRoomsOfAnyKindInOurOpinionAndCottagesAndVacationHomesAndHostelsAndYouCanGetRewardedBasicallyEverywhereAndResortsAndGetawaysAndInnsAndHarborInnsAndVillasForGuysNamedJamieFromNorthEasternMiamiAndVillasForHumansNamedAnyOtherKindOfHumanNameAndTreehousesAndHousesNearTreesAndRanchesForPeopleWhoReallyLikeFarmAnimalsAndRanchesForPeopleWhoJustFeelKindOfSoSoAboutFarmAnimalsAndRanchesInGeneralAndCliffHousesAndRewardsAndYesWeAreSurprisedYouAreStillTypingAtThisPointButHeyWeStillHaveMoreAccomodationsSoWeWillJustKeepListingThemBroBroAndBeachGetawaysAndSnowyGetawaysAndThoseAreTheOnlyTwoKindsOfGetawaysAndMansionsAndLoftsAndFarmStaysAndFarmStayRetreatsSpecificallyInNewZealandBecauseNothingSaysLuxuryLikeAnExcessOfHayAndGoatNoisesAndIceHotelsAndIceCabinsAndIceLodgesAndHotelsWithBalconyCapabilitiesAndHotelsWithRooftopsBecauseEveryHotelNeedsARoofItIsKindOfJustLogisticalAndTimesharesAndCountryHomesAndPalazzinasAndYesWeKnowWhatPalazzinasAreDontActLikeWeDontWhoIsTheExpertHereAndPenthousesAndManorsAndStatelyManorsWhichFeelPrettySimilarAndTownhousesAndPalacesAndHousesAndCasasWhichAreHousesButInSpanishAndGuestHousesForWomenNamedTammyWhoLiveForAGreatDealAndRewardsForEveryoneElseWhoLikesAGreatDealAndAllTypesOfPlacesForPetsAndAllTypesOfPlacesForNotPetsAndThemedHotelsAndPirateHotelsAndFairyTaleHotelsAndHauntedHotelsAndHauntedMotelsOrToPutItAnotherWayMotelsAndYesWeAreStillGoingAndFamilyResortsAndPlacesNearGoatYogaAndPlacesWherePeopleHaveMusclesAndPlacesThatServeMelonBetweenTheHoursOf6And10AMAndHouseYachtsAndYachtClubsAndOtherThingsOnOrNearWaterAndDolphinMotelsInLowerSanDiegoAndAlsoYurts.com",
}

func reportErr(x testing.TB, key, need, got string) {
	x.Error(key, "mismatch", "need", need, "got", got)
}
