package main

import (
	"crypto/x509/pkix"
	"encoding/asn1"
	"fmt"
	"log"
	"math/big"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/ethereum/go-ethereum/crypto"
)

type PubKeyAsn struct {
	Algo   pkix.AlgorithmIdentifier
	PubKey asn1.BitString
}

type RS struct {
	R, S *big.Int
}

// test and verify aws kms apis
func main() {
	sess, err := getAwsSess("us-west-2")
	if err != nil {
		log.Fatal(err)
	}
	svc := kms.New(sess)
	keyAlias := aws.String("alias/jundatest")
	resp, _ := svc.GetPublicKey(&kms.GetPublicKeyInput{
		KeyId: keyAlias,
	})
	pub := new(PubKeyAsn)
	_, err = asn1.Unmarshal(resp.PublicKey, pub)
	// skip first byte as it's just an indicator whether compress, and use last 20 bytes of hash
	// see PubkeyToAddress https://github.com/ethereum/go-ethereum/blob/master/crypto/crypto.go#L276
	log.Printf("addr: %x", crypto.Keccak256(pub.PubKey.Bytes[1:])[12:])

	// log.Printf("addr: %x", crypto.Keccak256(resp.PublicKey)[:20])
	digest := GeneratePrefixedHash([]byte("123456"))
	sig, err := svc.Sign(&kms.SignInput{
		KeyId:            keyAlias,
		MessageType:      aws.String("DIGEST"),
		Message:          digest,
		SigningAlgorithm: aws.String("ECDSA_SHA_256"),
	})
	if err != nil {
		log.Fatal("sign err:", err)
	}
	rs := new(RS)
	_, err = asn1.Unmarshal(sig.Signature, rs)

	var ethSig []byte
	ethSig = append(ethSig, rs.R.Bytes()...)
	ethSig = append(ethSig, rs.S.Bytes()...)
	ethSig = append(ethSig, 0)
	log.Print(len(ethSig))

	pubKey, err := crypto.SigToPub(digest, ethSig)
	log.Print(crypto.PubkeyToAddress(*pubKey).String())
	ethSig[64] = 1
	pubKey, err = crypto.SigToPub(digest, ethSig)
	log.Print(crypto.PubkeyToAddress(*pubKey).String())
}

func GeneratePrefixedHash(data []byte) []byte {
	return crypto.Keccak256([]byte("\x19Ethereum Signed Message:\n32"), crypto.Keccak256(data))
}

func getAwsSess(region string) (*session.Session, error) {
	sess, err := session.NewSession(&aws.Config{
		// default will read credentials from env or dot file
		// Credentials: credentials.NewStaticCredentials(awsKey, awsSec, ""),
		Region: aws.String(region)},
	)
	if err != nil {
		return nil, fmt.Errorf("newsession fail: %w", err)
	}
	return sess, nil
}
