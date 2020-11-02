package swid

import (
	"testing"
)

var (
	entitlementDataRequired = true

	testSoftwareMetaFull = SoftwareMeta{
		ActivationStatus:        "Licensed",
		ChannelType:             "Retail",
		ColloquialVersion:       "2020 Annus Horribilis edition",
		Description:             "The best of the best",
		Edition:                 "Enterprise",
		EntitlementDataRequired: &entitlementDataRequired,
		EntitlementKey:          "deadbeef",
		Generator:               "veraison-coswid-generator",
		PersistentID:            "356CCB1C-A873-42B0-8D17-FA33F56B9F3E",
		Product:                 "Alapage",
		ProductFamily:           "Email System",
		Revision:                "SP1",
		Summary:                 "quadragintesimal",
		UnspscCode:              "44121903",
		UnspscVersion:           "6.0315",
	}
)

func TestSoftwareMeta_RoundtripFull(t *testing.T) {
	tv := testSoftwareMetaFull
	/*
		af                                      # map(15)
		   18 2b                                # unsigned(43)
		   68                                   # text(8)
		      4c6963656e736564                  # "Licensed"
		   18 2c                                # unsigned(44)
		   66                                   # text(6)
		      52657461696c                      # "Retail"
		   18 2d                                # unsigned(45)
		   78 1d                                # text(29)
		      3230323020416e6e757320486f72726962696c69732065646974696f6e # "2020 Annus Horribilis edition"
		   18 2e                                # unsigned(46)
		   74                                   # text(20)
		      5468652062657374206f66207468652062657374 # "The best of the best"
		   18 2f                                # unsigned(47)
		   6a                                   # text(10)
		      456e7465727072697365              # "Enterprise"
		   18 30                                # unsigned(48)
		   f5                                   # primitive(21)
		   18 31                                # unsigned(49)
		   68                                   # text(8)
		      6465616462656566                  # "deadbeef"
		   18 32                                # unsigned(50)
		   78 19                                # text(25)
		      7665726169736f6e2d636f737769642d67656e657261746f72 # "veraison-coswid-generator"
		   18 33                                # unsigned(51)
		   78 24                                # text(36)
		      33353643434231432d413837332d343242302d384431372d464133334635364239463345 # "356CCB1C-A873-42B0-8D17-FA33F56B9F3E"
		   18 34                                # unsigned(52)
		   67                                   # text(7)
		      416c6170616765                    # "Alapage"
		   18 35                                # unsigned(53)
		   6c                                   # text(12)
		      456d61696c2053797374656d          # "Email System"
		   18 36                                # unsigned(54)
		   63                                   # text(3)
		      535031                            # "SP1"
		   18 37                                # unsigned(55)
		   70                                   # text(16)
		      71756164726167696e746573696d616c  # "quadragintesimal"
		   18 38                                # unsigned(56)
		   68                                   # text(8)
		      3434313231393033                  # "44121903"
		   18 39                                # unsigned(57)
		   66                                   # text(6)
			  362e30333135                      # "6.0315"
	*/
	expectedCBOR := []byte{
		0xaf, 0x18, 0x2b, 0x68, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65,
		0x64, 0x18, 0x2c, 0x66, 0x52, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18,
		0x2d, 0x78, 0x1d, 0x32, 0x30, 0x32, 0x30, 0x20, 0x41, 0x6e, 0x6e,
		0x75, 0x73, 0x20, 0x48, 0x6f, 0x72, 0x72, 0x69, 0x62, 0x69, 0x6c,
		0x69, 0x73, 0x20, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18,
		0x2e, 0x74, 0x54, 0x68, 0x65, 0x20, 0x62, 0x65, 0x73, 0x74, 0x20,
		0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x62, 0x65, 0x73, 0x74,
		0x18, 0x2f, 0x6a, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69,
		0x73, 0x65, 0x18, 0x30, 0xf5, 0x18, 0x31, 0x68, 0x64, 0x65, 0x61,
		0x64, 0x62, 0x65, 0x65, 0x66, 0x18, 0x32, 0x78, 0x19, 0x76, 0x65,
		0x72, 0x61, 0x69, 0x73, 0x6f, 0x6e, 0x2d, 0x63, 0x6f, 0x73, 0x77,
		0x69, 0x64, 0x2d, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f,
		0x72, 0x18, 0x33, 0x78, 0x24, 0x33, 0x35, 0x36, 0x43, 0x43, 0x42,
		0x31, 0x43, 0x2d, 0x41, 0x38, 0x37, 0x33, 0x2d, 0x34, 0x32, 0x42,
		0x30, 0x2d, 0x38, 0x44, 0x31, 0x37, 0x2d, 0x46, 0x41, 0x33, 0x33,
		0x46, 0x35, 0x36, 0x42, 0x39, 0x46, 0x33, 0x45, 0x18, 0x34, 0x67,
		0x41, 0x6c, 0x61, 0x70, 0x61, 0x67, 0x65, 0x18, 0x35, 0x6c, 0x45,
		0x6d, 0x61, 0x69, 0x6c, 0x20, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d,
		0x18, 0x36, 0x63, 0x53, 0x50, 0x31, 0x18, 0x37, 0x70, 0x71, 0x75,
		0x61, 0x64, 0x72, 0x61, 0x67, 0x69, 0x6e, 0x74, 0x65, 0x73, 0x69,
		0x6d, 0x61, 0x6c, 0x18, 0x38, 0x68, 0x34, 0x34, 0x31, 0x32, 0x31,
		0x39, 0x30, 0x33, 0x18, 0x39, 0x66, 0x36, 0x2e, 0x30, 0x33, 0x31,
		0x35,
	}

	roundTripper(t, tv, expectedCBOR)
}

func TestSoftwareMeta_RoundtripEmpty(t *testing.T) {
	tv := SoftwareMeta{}
	// encodes to the empty CBOR map
	expectedCBOR := []byte{0xa0}

	roundTripper(t, tv, expectedCBOR)
}
