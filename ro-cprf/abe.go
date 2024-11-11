package abe

// not sure if i did the package abe/rocprf stuff correctly
import {
	"ro-cprf/rocprf"
	// "ske"
}

// Master key for the CPRF
// length: length of the inner product
// modulus: inner product modulus
// z0: master key
type SecretKey struct {
	length  int
	modulus *big.Int
	z0      []*big.Int
}

type Plaintext struct {
	length int
	plaintext []byte
}

type Ciphertext struct {
	length int
	ciphertext []byte
}

// ABE.Setup(1λ) : It computes ppR← CPRF.Setup(1λ) and msk ← CPRF.KeyGen(pp), and outputs a public parameter pp and a master secret key msk.
func ABESetup(modulus *big.Int, length int) (msk *MasterKey) {
	msk, _ := KeyGen(modulus, length)
	return
}

// ABE.KeyGen(msk, f) : It computes skf ← CPRF.Constrain(msk, f), and outputs skf .
func ABEKeyGen(msk *MasterKey) (skf *SecretKey) {
	z := msk.z0
	skf, _ := msk.Constrain(z) // not sure how this works, Constrain has (msk *MasterKey)
	return
}

// ABE.Enc(msk, x, msg) : It computes k ← CPRF.Eval(msk, x) and ct ← SKE.Enc(k, msg), and outputs ct
func ABEEnc(msk *MasterKey, x []*big.Int, m *Plaintext) (ct *Ciphertext){
	k := msk.Eval(x) // Eval has (msk *MasterKey)
	// ct := SKEEnc(k, m)
	// return
}

// ABE.Dec(skf , ct) : It computes k ← CPRF.CEval(skf , x), and outputs SKE.Dec(k, ct).
func ABEDec(skf *SecretKey, x []*big.Int, ct *Ciphertext) (fm []byte){
	k = skf.CEval(x) // CEval has (csk *ConstrainedKey)
	// fm := SKEDec(k, ct)
	// return
}

