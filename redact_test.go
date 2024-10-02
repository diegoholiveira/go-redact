package redact_test

import (
	"strings"
	"testing"

	"github.com/diegoholiveira/go-redact"
	"github.com/stretchr/testify/assert"
)

func TestRedact(t *testing.T) {
	sb := strings.Builder{}

	redact.Redact(strings.NewReader(input), &sb, sensitiveWords, "__REDACT__")

	assert.Equal(t, expected, sb.String())
}

var sensitiveWords = []string{
	"CompanyX",
	"ClientA",
	"ClientB",
	"ClientC",
	"RegionY",
	"RegionZ",
	"SystemOmega",
	"Person1",
	"CompanyY",
	"YearX",
	"PercentageX",
	"RegulatorX",
	"OrganizationY",
	"CountryY",
	"LawFirmX",
	"RegionW",
	"Person2",
	"Person3",
	"InfluencerA",
	"InfluencerB",
	"Person4",
	"UnionX",
}

var input = `In recent years, CompanyX has seen remarkable growth in several markets. Our expansion strategy involved partnerships with various high-profile clients, including ClientA, ClientB, and ClientC. These partnerships have led to significant revenue increases, particularly in the regions of RegionY and RegionZ.

One of the key factors behind our success is the implementation of a new proprietary system, SystemOmega. This system, developed by Person1 and his team, allows for real-time monitoring of transactions across multiple platforms. Additionally, we’ve optimized our cybersecurity measures by working closely with CompanyY, which specializes in encryption technologies.

Furthermore, internal reports suggest that by YearX, the company is on track to achieve a market share of PercentageX in the e-commerce sector. However, we must remain vigilant in addressing certain issues raised by RegulatorX and OrganizationY regarding compliance with local data privacy regulations, especially in CountryY.

Our legal department, led by Person2, is already in talks with LawFirmX to ensure that we meet all legal requirements before the upcoming expansion into RegionW. In addition, Person3, the head of marketing, has outlined a strategy targeting younger consumers, leveraging the influence of InfluencerA and InfluencerB.

Despite these advancements, we are still managing internal challenges, including the resignation of Person4 and potential conflicts with UnionX.`

var expected = `In recent years, __REDACT__ has seen remarkable growth in several markets. Our expansion strategy involved partnerships with various high-profile clients, including __REDACT__, __REDACT__, and __REDACT__. These partnerships have led to significant revenue increases, particularly in the regions of __REDACT__ and __REDACT__.

One of the key factors behind our success is the implementation of a new proprietary system, __REDACT__. This system, developed by __REDACT__ and his team, allows for real-time monitoring of transactions across multiple platforms. Additionally, we’ve optimized our cybersecurity measures by working closely with __REDACT__, which specializes in encryption technologies.

Furthermore, internal reports suggest that by __REDACT__, the company is on track to achieve a market share of __REDACT__ in the e-commerce sector. However, we must remain vigilant in addressing certain issues raised by __REDACT__ and __REDACT__ regarding compliance with local data privacy regulations, especially in __REDACT__.

Our legal department, led by __REDACT__, is already in talks with __REDACT__ to ensure that we meet all legal requirements before the upcoming expansion into __REDACT__. In addition, __REDACT__, the head of marketing, has outlined a strategy targeting younger consumers, leveraging the influence of __REDACT__ and __REDACT__.

Despite these advancements, we are still managing internal challenges, including the resignation of __REDACT__ and potential conflicts with __REDACT__.`
