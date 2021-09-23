package fee

import (
	"flag"
	"fmt"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
	"math/big"
	"math/rand"
	"os"
	"testing"
	"time"
)

// TestMain is used to setup/teardown a temporary CockroachDB instance
// and run all the unit tests in between.
func TestMain(m *testing.M) {
	flag.Parse()
	rand.Seed(time.Now().Unix())

	if err := setup(); err != nil {
		fmt.Println("cannot setup DB:", err)
		os.Exit(1)
	}

	exitCode := m.Run() // run all unittests
	os.Exit(exitCode)
}

var priceCache *TokenPriceCache

func setup() error {
	// Start polling
	priceCache = NewTokenPriceCache()
	return nil
}

func errIsNil(t *testing.T, err error) {
	if err != nil {
		t.Errorf("invalid error, it must be nil: %v", err)
	}
}

func TestFee(t *testing.T) {
	token := &types.Token{
		Symbol:  "DAI",
		Address: "",
		Decimal: 18,
	}
	chainToken := &types.Token{
		Symbol:  "HT",
		Address: "",
		Decimal: 18,
	}
	tokenUsdPrice, err := priceCache.GetTokenPrice(token, chainToken, new(big.Int).Mul(big.NewInt(2500), big.NewInt(19000000000)))
	errIsNil(t, err)
	t.Log("DAI eth prize:", tokenUsdPrice)
}
