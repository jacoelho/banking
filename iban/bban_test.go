package iban_test

import (
	"reflect"
	"testing"

	"github.com/jacoelho/banking/iban"
)

func TestBBAN(t *testing.T) {
	tests := []struct {
		iban    string
		bban    iban.BBAN
		wantErr bool
	}{
		{
			iban: "AL47212110090000000235698741",
			bban: iban.BBAN{
				BBAN:          "212110090000000235698741",
				BankCode:      "212",
				BranchCode:    "11009",
				AccountNumber: "0000000235698741"},
		},
		{
			iban:    "AD1200012030200359100100",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "00012030200359100100",
				BankCode:      "0001",
				BranchCode:    "2030",
				AccountNumber: "200359100100"},
		},
		{
			iban:    "AT611904300234573201",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "1904300234573201",
				BankCode:      "19043",
				BranchCode:    "",
				AccountNumber: "00234573201"},
		},
		{
			iban:    "AZ21NABZ00000000137010001944",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "NABZ00000000137010001944",
				BankCode:      "NABZ",
				BranchCode:    "",
				AccountNumber: "00000000137010001944"},
		},
		{
			iban:    "BH67BMAG00001299123456",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "BMAG00001299123456",
				BankCode:      "BMAG",
				BranchCode:    "",
				AccountNumber: "00001299123456"},
		},
		{
			iban:    "BE68539007547034",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "539007547034",
				BankCode:      "539",
				BranchCode:    "",
				AccountNumber: "007547034"},
		},
		{
			iban:    "BA391290079401028494",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "1290079401028494",
				BankCode:      "129",
				BranchCode:    "007",
				AccountNumber: "9401028494"},
		},
		{
			iban:    "BR9700360305000010009795493P1",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "00360305000010009795493P1",
				BankCode:      "00360305",
				BranchCode:    "00001",
				AccountNumber: "0009795493P1"},
		},
		{
			iban:    "BR1800000000141455123924100C2",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "00000000141455123924100C2",
				BankCode:      "00000000",
				BranchCode:    "14145",
				AccountNumber: "5123924100C2"},
		},
		{
			iban:    "BG80BNBG96611020345678",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "BNBG96611020345678",
				BankCode:      "BNBG",
				BranchCode:    "9661",
				AccountNumber: "1020345678"},
		},
		{
			iban:    "CR05015202001026284066",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "015202001026284066",
				BankCode:      "0152",
				BranchCode:    "",
				AccountNumber: "02001026284066"},
		},
		{
			iban:    "HR1210010051863000160",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "10010051863000160",
				BankCode:      "1001005",
				BranchCode:    "",
				AccountNumber: "1863000160"},
		},
		{
			iban:    "CY17002001280000001200527600",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "002001280000001200527600",
				BankCode:      "002",
				BranchCode:    "00128",
				AccountNumber: "0000001200527600"},
		},
		{
			iban:    "CZ6508000000192000145399",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "08000000192000145399",
				BankCode:      "0800",
				BranchCode:    "",
				AccountNumber: "0000192000145399"},
		},
		{
			iban:    "CZ9455000000001011038930",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "55000000001011038930",
				BankCode:      "5500",
				BranchCode:    "",
				AccountNumber: "0000001011038930"},
		},
		{
			iban:    "DK5000400440116243",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "00400440116243",
				BankCode:      "0040",
				BranchCode:    "",
				AccountNumber: "0440116243"},
		},
		{
			iban:    "FO6264600001631634",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "64600001631634",
				BankCode:      "6460",
				BranchCode:    "",
				AccountNumber: "0001631634"},
		},
		{
			iban:    "GL8964710001000206",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "64710001000206",
				BankCode:      "6471",
				BranchCode:    "",
				AccountNumber: "0001000206"},
		},
		{
			iban:    "DO28BAGR00000001212453611324",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "BAGR00000001212453611324",
				BankCode:      "BAGR",
				BranchCode:    "",
				AccountNumber: "00000001212453611324"},
		},
		{
			iban:    "EE382200221020145685",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "2200221020145685",
				BankCode:      "22",
				BranchCode:    "",
				AccountNumber: "00221020145685"},
		},
		{
			iban:    "FI2112345600000785",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "12345600000785",
				BankCode:      "123",
				BranchCode:    "",
				AccountNumber: "45600000785"},
		},
		{
			iban:    "FI5542345670000081",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "42345670000081",
				BankCode:      "423",
				BranchCode:    "",
				AccountNumber: "45670000081"},
		},
		{
			iban:    "FR1420041010050500013M02606",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "20041010050500013M02606",
				BankCode:      "20041",
				BranchCode:    "",
				AccountNumber: "010050500013M02606"},
		},
		{
			iban:    "GE29NB0000000101904917",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "NB0000000101904917",
				BankCode:      "NB",
				BranchCode:    "",
				AccountNumber: "0000000101904917"},
		},
		{
			iban:    "DE89370400440532013000",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "370400440532013000",
				BankCode:      "37040044",
				BranchCode:    "",
				AccountNumber: "0532013000"},
		},
		{
			iban:    "GI75NWBK000000007099453",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "NWBK000000007099453",
				BankCode:      "NWBK",
				BranchCode:    "",
				AccountNumber: "000000007099453"},
		},
		{
			iban:    "GR1601101250000000012300695",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "01101250000000012300695",
				BankCode:      "011",
				BranchCode:    "0125",
				AccountNumber: "0000000012300695"},
		},
		{
			iban:    "GT82TRAJ01020000001210029690",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "TRAJ01020000001210029690",
				BankCode:      "TRAJ",
				BranchCode:    "",
				AccountNumber: "01020000001210029690"},
		},
		{
			iban:    "HU42117730161111101800000000",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "117730161111101800000000",
				BankCode:      "117",
				BranchCode:    "7301",
				AccountNumber: "61111101800000000"},
		},
		{
			iban:    "IS140159260076545510730339",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "0159260076545510730339",
				BankCode:      "01",
				BranchCode:    "59",
				AccountNumber: "260076545510730339"},
		},
		{
			iban:    "IE29AIBK93115212345678",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "AIBK93115212345678",
				BankCode:      "AIBK",
				BranchCode:    "931152",
				AccountNumber: "12345678"},
		},
		{
			iban:    "IL620108000000099999999",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "0108000000099999999",
				BankCode:      "010",
				BranchCode:    "800",
				AccountNumber: "0000099999999"},
		},
		{
			iban:    "IT60X0542811101000000123456",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "X0542811101000000123456",
				BankCode:      "05428",
				BranchCode:    "11101",
				AccountNumber: "000000123456"},
		},
		{
			iban:    "JO94CBJO0010000000000131000302",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "CBJO0010000000000131000302",
				BankCode:      "CBJO",
				BranchCode:    "0010",
				AccountNumber: "000000000131000302"},
		},
		{
			iban:    "XK051212012345678906",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "1212012345678906",
				BankCode:      "12",
				BranchCode:    "12",
				AccountNumber: "012345678906"},
		},
		{
			iban:    "KW81CBKU0000000000001234560101",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "CBKU0000000000001234560101",
				BankCode:      "CBKU",
				BranchCode:    "",
				AccountNumber: "0000000000001234560101"},
		},
		{
			iban:    "LV80BANK0000435195001",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "BANK0000435195001",
				BankCode:      "BANK",
				BranchCode:    "",
				AccountNumber: "0000435195001"},
		},
		{
			iban:    "LB62099900000001001901229114",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "099900000001001901229114",
				BankCode:      "0999",
				BranchCode:    "",
				AccountNumber: "00000001001901229114"},
		},
		{
			iban:    "LI21088100002324013AA",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "088100002324013AA",
				BankCode:      "08810",
				BranchCode:    "",
				AccountNumber: "0002324013AA"},
		},
		{
			iban:    "LT121000011101001000",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "1000011101001000",
				BankCode:      "10000",
				BranchCode:    "",
				AccountNumber: "11101001000"},
		},
		{
			iban:    "LU280019400644750000",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "0019400644750000",
				BankCode:      "001",
				BranchCode:    "",
				AccountNumber: "9400644750000"},
		},
		{
			iban:    "MK07250120000058984",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "250120000058984",
				BankCode:      "250",
				BranchCode:    "",
				AccountNumber: "120000058984"},
		},
		{
			iban:    "MT84MALT011000012345MTLCAST001S",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "MALT011000012345MTLCAST001S",
				BankCode:      "MALT",
				BranchCode:    "01100",
				AccountNumber: "0012345MTLCAST001S"},
		},
		{
			iban:    "MR1300020001010000123456753",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "00020001010000123456753",
				BankCode:      "00020",
				BranchCode:    "00101",
				AccountNumber: "0000123456753"},
		},
		{
			iban:    "MU17BOMM0101101030300200000MUR",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "BOMM0101101030300200000MUR",
				BankCode:      "BOMM01",
				BranchCode:    "01",
				AccountNumber: "101030300200000MUR"},
		},
		{
			iban:    "MD24AG000225100013104168",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "AG000225100013104168",
				BankCode:      "AG",
				BranchCode:    "",
				AccountNumber: "000225100013104168"},
		},
		{
			iban:    "MC5811222000010123456789030",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "11222000010123456789030",
				BankCode:      "11222",
				BranchCode:    "00001",
				AccountNumber: "0123456789030"},
		},
		{
			iban:    "ME25505000012345678951",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "505000012345678951",
				BankCode:      "505",
				BranchCode:    "",
				AccountNumber: "000012345678951"},
		},
		{
			iban:    "NL91ABNA0417164300",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "ABNA0417164300",
				BankCode:      "ABNA",
				BranchCode:    "",
				AccountNumber: "0417164300"},
		},
		{
			iban:    "NO9386011117947",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "86011117947",
				BankCode:      "8601",
				BranchCode:    "",
				AccountNumber: "1117947"},
		},
		{
			iban:    "PK36SCBL0000001123456702",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "SCBL0000001123456702",
				BankCode:      "SCBL",
				BranchCode:    "",
				AccountNumber: "0000001123456702"},
		},
		{
			iban:    "PS92PALS000000000400123456702",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "PALS000000000400123456702",
				BankCode:      "PALS",
				BranchCode:    "",
				AccountNumber: "000000000400123456702"},
		},
		{
			iban:    "PL61109010140000071219812874",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "109010140000071219812874",
				BankCode:      "10901014",
				BranchCode:    "",
				AccountNumber: "0000071219812874"},
		},
		{
			iban:    "PT50000201231234567890154",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "000201231234567890154",
				BankCode:      "0002",
				BranchCode:    "",
				AccountNumber: "01231234567890154"},
		},
		{
			iban:    "QA58DOHB00001234567890ABCDEFG",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "DOHB00001234567890ABCDEFG",
				BankCode:      "DOHB",
				BranchCode:    "",
				AccountNumber: "00001234567890ABCDEFG"},
		},
		{
			iban:    "RO49AAAA1B31007593840000",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "AAAA1B31007593840000",
				BankCode:      "AAAA",
				BranchCode:    "",
				AccountNumber: "1B31007593840000"},
		},
		{
			iban:    "LC55HEMM000100010012001200023015",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "HEMM000100010012001200023015",
				BankCode:      "HEMM",
				BranchCode:    "",
				AccountNumber: "000100010012001200023015"},
		},
		{
			iban:    "SM86U0322509800000000270100",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "U0322509800000000270100",
				BankCode:      "03225",
				BranchCode:    "09800",
				AccountNumber: "000000270100"},
		},
		{
			iban:    "ST23000100010051845310146",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "000100010051845310146",
				BankCode:      "0001",
				BranchCode:    "0001",
				AccountNumber: "0051845310146"},
		},
		{
			iban:    "SA0380000000608010167519",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "80000000608010167519",
				BankCode:      "80",
				BranchCode:    "",
				AccountNumber: "000000608010167519"},
		},
		{
			iban:    "RS35260005601001611379",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "260005601001611379",
				BankCode:      "260",
				BranchCode:    "",
				AccountNumber: "005601001611379"},
		},
		{
			iban:    "SK3112000000198742637541",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "12000000198742637541",
				BankCode:      "1200",
				BranchCode:    "",
				AccountNumber: "0000198742637541"},
		},
		{
			iban:    "SI56191000000123438",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "191000000123438",
				BankCode:      "19100",
				BranchCode:    "",
				AccountNumber: "0000123438"},
		},
		{
			iban:    "ES9121000418450200051332",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "21000418450200051332",
				BankCode:      "2100",
				BranchCode:    "0418",
				AccountNumber: "450200051332"},
		},
		{
			iban:    "SE4550000000058398257466",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "50000000058398257466",
				BankCode:      "500",
				BranchCode:    "",
				AccountNumber: "00000058398257466"},
		},
		{
			iban:    "CH9300762011623852957",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "00762011623852957",
				BankCode:      "00762",
				BranchCode:    "",
				AccountNumber: "011623852957"},
		},
		{
			iban:    "TL380080012345678910157",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "0080012345678910157",
				BankCode:      "008",
				BranchCode:    "",
				AccountNumber: "0012345678910157"},
		},
		{
			iban:    "TN5910006035183598478831",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "10006035183598478831",
				BankCode:      "10",
				BranchCode:    "006",
				AccountNumber: "035183598478831"},
		},
		{
			iban:    "TR330006100519786457841326",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "0006100519786457841326",
				BankCode:      "00061",
				BranchCode:    "",
				AccountNumber: "00519786457841326"},
		},
		{
			iban:    "AE070331234567890123456",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "0331234567890123456",
				BankCode:      "033",
				BranchCode:    "",
				AccountNumber: "1234567890123456"},
		},
		{
			iban:    "GB29NWBK60161331926819",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "NWBK60161331926819",
				BankCode:      "NWBK",
				BranchCode:    "601613",
				AccountNumber: "31926819"},
		},
		{
			iban:    "VG96VPVG0000012345678901",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "VPVG0000012345678901",
				BankCode:      "VPVG",
				BranchCode:    "",
				AccountNumber: "0000012345678901"},
		},
		{
			iban:    "EG380019000500000000263180002",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "0019000500000000263180002",
				BankCode:      "0019",
				BranchCode:    "0005",
				AccountNumber: "00000000263180002",
			},
		},
		{
			iban:    "IQ98NBIQ850123456789012",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "NBIQ850123456789012",
				BankCode:      "NBIQ",
				BranchCode:    "850",
				AccountNumber: "123456789012",
			},
		},
		{
			iban:    "SV62CENR00000000000000700025",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "CENR00000000000000700025",
				BankCode:      "CENR",
				BranchCode:    "",
				AccountNumber: "00000000000000700025",
			},
		},
		{
			iban:    "VA59001123000012345678",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "001123000012345678",
				BankCode:      "001",
				BranchCode:    "",
				AccountNumber: "123000012345678",
			},
		},
		{
			iban:    "BY13NBRB3600900000002Z00AB00",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "NBRB3600900000002Z00AB00",
				BankCode:      "NBRB",
				BranchCode:    "",
				AccountNumber: "3600900000002Z00AB00",
			},
		},
		{
			iban:    "SD2129010501234001",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "29010501234001",
				BankCode:      "29",
				BranchCode:    "",
				AccountNumber: "010501234001",
			},
		},
		{
			iban:    "LY83002048000020100120361",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "002048000020100120361",
				BankCode:      "002",
				BranchCode:    "048",
				AccountNumber: "000020100120361",
			},
		},
		{
			iban:    "BI4210000100010000332045181",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "10000100010000332045181",
				BankCode:      "10000",
				BranchCode:    "10001",
				AccountNumber: "0000332045181",
			},
		},
		{
			iban:    "DJ2100010000000154000100186",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "00010000000154000100186",
				BankCode:      "00010",
				BranchCode:    "00000",
				AccountNumber: "0154000100186",
			},
		},
		{
			iban:    "RU0304452522540817810538091310419",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "04452522540817810538091310419",
				BankCode:      "044525225",
				BranchCode:    "40817",
				AccountNumber: "810538091310419",
			},
		},
		{
			iban:    "SO211000001001000100141",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "1000001001000100141",
				BankCode:      "1000",
				BranchCode:    "001",
				AccountNumber: "001000100141",
			},
		},
		{
			iban:    "MN121234123456789123",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "1234123456789123",
				BankCode:      "1234",
				BranchCode:    "",
				AccountNumber: "123456789123",
			},
		},
		{
			iban:    "NI45BAPR00000013000003558124",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "BAPR00000013000003558124",
				BankCode:      "BAPR",
				BranchCode:    "",
				AccountNumber: "00000013000003558124",
			},
		},
		{
			iban:    "FK88SC123456789012",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "SC123456789012",
				BankCode:      "SC",
				BranchCode:    "",
				AccountNumber: "123456789012",
			},
		},
		{
			iban:    "OM810180000001299123456",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "0180000001299123456",
				BankCode:      "018",
				BranchCode:    "",
				AccountNumber: "0000001299123456",
			},
		},
		{
			iban:    "YE15CBYE0001018861234567891234",
			wantErr: false,
			bban: iban.BBAN{
				BBAN:          "CBYE0001018861234567891234",
				BankCode:      "CBYE",
				BranchCode:    "0001",
				AccountNumber: "018861234567891234",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.iban,
			func(t *testing.T) {
				got, err := iban.GetBBAN(tt.iban)
				if (err != nil) != tt.wantErr {
					t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
				}

				if !reflect.DeepEqual(tt.bban, got) {
					t.Errorf("GetBBAN() got = %v, want %v", got, tt.bban)
				}
			})
	}

}
