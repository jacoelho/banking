package iban_test

import (
	"testing"

	"github.com/jacoelho/banking/iban"
)

func TestValidateIBAN(t *testing.T) {
	tests := []struct {
		iban    string
		wantErr bool
	}{
		{
			iban:    "AL47212110090000000235698741",
			wantErr: false,
		},
		{
			iban:    "AD1200012030200359100100",
			wantErr: false,
		},
		{
			iban:    "AT611904300234573201",
			wantErr: false,
		},
		{
			iban:    "AZ21NABZ00000000137010001944",
			wantErr: false,
		},
		{
			iban:    "BH67BMAG00001299123456",
			wantErr: false,
		},
		{
			iban:    "BE68539007547034",
			wantErr: false,
		},
		{
			iban:    "BA391290079401028494",
			wantErr: false,
		},
		{
			iban:    "BR9700360305000010009795493P1",
			wantErr: false,
		},
		{
			iban:    "BR1800000000141455123924100C2",
			wantErr: false,
		},
		{
			iban:    "BG80BNBG96611020345678",
			wantErr: false,
		},
		{
			iban:    "CR05015202001026284066",
			wantErr: false,
		},
		{
			iban:    "HR1210010051863000160",
			wantErr: false,
		},
		{
			iban:    "CY17002001280000001200527600",
			wantErr: false,
		},
		{
			iban:    "CZ6508000000192000145399",
			wantErr: false,
		},
		{
			iban:    "CZ9455000000001011038930",
			wantErr: false,
		},
		{
			iban:    "DK5000400440116243",
			wantErr: false,
		},
		{
			iban:    "FO6264600001631634",
			wantErr: false,
		},
		{
			iban:    "GL8964710001000206",
			wantErr: false,
		},
		{
			iban:    "DO28BAGR00000001212453611324",
			wantErr: false,
		},
		{
			iban:    "EE382200221020145685",
			wantErr: false,
		},
		{
			iban:    "FI2112345600000785",
			wantErr: false,
		},
		{
			iban:    "FI5542345670000081",
			wantErr: false,
		},
		{
			iban:    "FR1420041010050500013M02606",
			wantErr: false,
		},
		{
			iban:    "GE29NB0000000101904917",
			wantErr: false,
		},
		{
			iban:    "DE89370400440532013000",
			wantErr: false,
		},
		{
			iban:    "GI75NWBK000000007099453",
			wantErr: false,
		},
		{
			iban:    "GR1601101250000000012300695",
			wantErr: false,
		},
		{
			iban:    "GT82TRAJ01020000001210029690",
			wantErr: false,
		},
		{
			iban:    "HU42117730161111101800000000",
			wantErr: false,
		},
		{
			iban:    "IS140159260076545510730339",
			wantErr: false,
		},
		{
			iban:    "IE29AIBK93115212345678",
			wantErr: false,
		},
		{
			iban:    "IL620108000000099999999",
			wantErr: false,
		},
		{
			iban:    "IT60X0542811101000000123456",
			wantErr: false,
		},
		{
			iban:    "JO94CBJO0010000000000131000302",
			wantErr: false,
		},
		{
			iban:    "XK051212012345678906",
			wantErr: false,
		},
		{
			iban:    "KW81CBKU0000000000001234560101",
			wantErr: false,
		},
		{
			iban:    "LV80BANK0000435195001",
			wantErr: false,
		},
		{
			iban:    "LB62099900000001001901229114",
			wantErr: false,
		},
		{
			iban:    "LI21088100002324013AA",
			wantErr: false,
		},
		{
			iban:    "LT121000011101001000",
			wantErr: false,
		},
		{
			iban:    "LU280019400644750000",
			wantErr: false,
		},
		{
			iban:    "MK07250120000058984",
			wantErr: false,
		},
		{
			iban:    "MT84MALT011000012345MTLCAST001S",
			wantErr: false,
		},
		{
			iban:    "MR1300020001010000123456753",
			wantErr: false,
		},
		{
			iban:    "MU17BOMM0101101030300200000MUR",
			wantErr: false,
		},
		{
			iban:    "MD24AG000225100013104168",
			wantErr: false,
		},
		{
			iban:    "MC5811222000010123456789030",
			wantErr: false,
		},
		{
			iban:    "ME25505000012345678951",
			wantErr: false,
		},
		{
			iban:    "NL91ABNA0417164300",
			wantErr: false,
		},
		{
			iban:    "NO9386011117947",
			wantErr: false,
		},
		{
			iban:    "PK36SCBL0000001123456702",
			wantErr: false,
		},
		{
			iban:    "PS92PALS000000000400123456702",
			wantErr: false,
		},
		{
			iban:    "PL61109010140000071219812874",
			wantErr: false,
		},
		{
			iban:    "PT50000201231234567890154",
			wantErr: false,
		},
		{
			iban:    "QA58DOHB00001234567890ABCDEFG",
			wantErr: false,
		},
		{
			iban:    "RO49AAAA1B31007593840000",
			wantErr: false,
		},
		{
			iban:    "LC55HEMM000100010012001200023015",
			wantErr: false,
		},
		{
			iban:    "SM86U0322509800000000270100",
			wantErr: false,
		},
		{
			iban:    "ST68000100010051845310112",
			wantErr: false,
		},
		{
			iban:    "SA0380000000608010167519",
			wantErr: false,
		},
		{
			iban:    "RS35260005601001611379",
			wantErr: false,
		},
		{
			iban:    "SK3112000000198742637541",
			wantErr: false,
		},
		{
			iban:    "SI56191000000123438",
			wantErr: false,
		},
		{
			iban:    "ES9121000418450200051332",
			wantErr: false,
		},
		{
			iban:    "SE4550000000058398257466",
			wantErr: false,
		},
		{
			iban:    "CH9300762011623852957",
			wantErr: false,
		},
		{
			iban:    "TL380080012345678910157",
			wantErr: false,
		},
		{
			iban:    "TN5910006035183598478831",
			wantErr: false,
		},
		{
			iban:    "TR330006100519786457841326",
			wantErr: false,
		},
		{
			iban:    "AE070331234567890123456",
			wantErr: false,
		},
		{
			iban:    "GB29NWBK60161331926819",
			wantErr: false,
		},
		{
			iban:    "VG96VPVG0000012345678901",
			wantErr: false,
		},
		{
			iban:    "EG380019000500000000263180002",
			wantErr: false,
		},
		{
			iban:    "IQ98NBIQ850123456789012",
			wantErr: false,
		},
		{
			iban:    "SV62CENR00000000000000700025",
			wantErr: false,
		},
		{
			iban:    "VA59001123000012345678",
			wantErr: false,
		},
		{
			iban:    "SD2129010501234001",
			wantErr: false,
		},
		{
			iban:    "BI4210000100010000332045181",
			wantErr: false,
		},
		{
			iban:    "DJ2100010000000154000100186",
			wantErr: false,
		},
		{
			iban:    "RU0304452522540817810538091310419",
			wantErr: false,
		},
		{
			iban:    "SO211000001001000100141",
			wantErr: false,
		},
		{
			iban:    "MN121234123456789123",
			wantErr: false,
		},
		{
			iban:    "NI45BAPR00000013000003558124",
			wantErr: false,
		},
		{
			iban:    "FK88SC123456789012",
			wantErr: false,
		},
		{
			iban:    "OM810180000001299123456",
			wantErr: false,
		},
		{
			iban:    "YE15CBYE0001018861234567891234",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.iban, func(t *testing.T) {
			if err := iban.Validate(tt.iban); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
