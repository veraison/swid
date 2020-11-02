package swid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLink_DecodeFullWithGlobalAttributes(t *testing.T) {
	/*
	   a9                               # map(9)
	      0f                            # unsigned(15)   -.
	      6b                            # text(11)        | global-attributes::lang
	         756e70726179657266756c     # "unprayerful"  -:
	      39 02d8                       # negative(728)   | global-attributes::custom(-729) <- ignored
	      19 059b                       # unsigned(1435) -'
	      18 25                         # unsigned(37)
	      6d                            # text(13)
	         726563656d656e746174696f6e # "recementation"
	      18 26                         # unsigned(38)
	      6a                            # text(10)
	         657863616c6174696f6e       # "excalation"
	      0a                            # unsigned(10)
	      6b                            # text(11)
	         65726f746f6d616e696163     # "erotomaniac"
	      18 27                         # unsigned(39)
	      03                            # unsigned(3)
	      18 28                         # unsigned(40)
	      01                            # unsigned(1)
	      18 29                         # unsigned(41)
	      6a                            # text(10)
	         4d61796f6c6f67697374       # "Mayologist"
	      18 2a                         # unsigned(42)
	      6c                            # text(12)
	   	  696d70657261746f7269616e      # "imperatorian"
	*/
	tv := []byte{
		0xa9, 0x0f, 0x6b, 0x75, 0x6e, 0x70, 0x72, 0x61, 0x79, 0x65, 0x72,
		0x66, 0x75, 0x6c, 0x39, 0x02, 0xd8, 0x19, 0x05, 0x9b, 0x18, 0x25,
		0x6d, 0x72, 0x65, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74,
		0x69, 0x6f, 0x6e, 0x18, 0x26, 0x6a, 0x65, 0x78, 0x63, 0x61, 0x6c,
		0x61, 0x74, 0x69, 0x6f, 0x6e, 0x0a, 0x6b, 0x65, 0x72, 0x6f, 0x74,
		0x6f, 0x6d, 0x61, 0x6e, 0x69, 0x61, 0x63, 0x18, 0x27, 0x03, 0x18,
		0x28, 0x01, 0x18, 0x29, 0x6a, 0x4d, 0x61, 0x79, 0x6f, 0x6c, 0x6f,
		0x67, 0x69, 0x73, 0x74, 0x18, 0x2a, 0x6c, 0x69, 0x6d, 0x70, 0x65,
		0x72, 0x61, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6e,
	}

	actual := Link{}

	err := dm.Unmarshal(tv, &actual)

	assert.Nil(t, err)

	expectedLang := "unprayerful"

	expected := Link{
		GlobalAttributes: GlobalAttributes{
			Lang: expectedLang,
		},
		Artifact:  "recementation",
		Href:      "excalation",
		Media:     "erotomaniac",
		Ownership: &Ownership{OwnershipAbandon},
		Rel:       Rel{RelAncestor},
		MediaType: "Mayologist",
		Use:       &Use{"imperatorian"},
	}

	assert.Equal(t, expected, actual)
	assert.Equal(t, "abandon", expected.GetOwnershipAsString())
	assert.Equal(t, "ancestor", expected.GetRelAsString())
}

func TestLink_EncodeMinset(t *testing.T) {
	tv := Link{
		Artifact:  "/bin/installer.sh",
		Href:      "swid:2df9de35-0aff-4a86-ace6-f7dddd1ade4c",
		Ownership: &Ownership{OwnershipShared},
		Rel:       Rel{RelInstallationMedia},
		Use:       &Use{UseRequired},
	}

	data, err := em.Marshal(tv)

	assert.Nil(t, err)
	assert.Equal(t, data,
		/*
			a5                                      # map(5)
			   18 25                                # unsigned(37)
			   71                                   # text(17)
				  # "/bin/installer.sh"
				  2f62696e2f696e7374616c6c65722e7368
			   18 26                                # unsigned(38)
			   78 29                                # text(41)
				  # "swid:2df9de35-0aff-4a86-ace6-f7dddd1ade4c"
				  737769643a32646639646533352d306166662d346138362d616365362d663764646464316164653463
			   18 27                                # unsigned(39)
			   01                                   # unsigned(1)
			   18 28                                # unsigned(40)
			   04                                   # unsigned(4)
			   18 2a                                # unsigned(42)
			   02                                   # unsigned(2)
		*/
		[]byte{
			0xa5, 0x18, 0x25, 0x71, 0x2f, 0x62, 0x69, 0x6e, 0x2f, 0x69, 0x6e,
			0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x73, 0x68, 0x18,
			0x26, 0x78, 0x29, 0x73, 0x77, 0x69, 0x64, 0x3a, 0x32, 0x64, 0x66,
			0x39, 0x64, 0x65, 0x33, 0x35, 0x2d, 0x30, 0x61, 0x66, 0x66, 0x2d,
			0x34, 0x61, 0x38, 0x36, 0x2d, 0x61, 0x63, 0x65, 0x36, 0x2d, 0x66,
			0x37, 0x64, 0x64, 0x64, 0x64, 0x31, 0x61, 0x64, 0x65, 0x34, 0x63,
			0x18, 0x27, 0x01, 0x18, 0x28, 0x04, 0x18, 0x2a, 0x02,
		},
	)
}
