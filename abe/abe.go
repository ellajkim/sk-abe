package abe

import (
	"math/big"
	"github.com/ellajkim/cprf/ro-cprf/rocprf"
	"github.com/ellajkim/cprf/ske"
)

// ABE.Setup(1λ) : It computes ppR← CPRF.Setup(1λ) and msk ← CPRF.KeyGen(pp), and outputs a public parameter pp and a master secret key msk.
func ABESetup(modulus *big.Int, length int) (*rocprf.MasterKey, error) {
    msk, err := rocprf.KeyGen(modulus, length)
    if err != nil {
        return nil, err
    }
    return msk, nil
}

// ABE.KeyGen(msk, f) : It computes skf ← CPRF.Constrain(msk, f), and outputs skf .
func ABEKeyGen(msk *rocprf.MasterKey, f []*big.Int) (*rocprf.ConstrainedKey, error) {
	skf, err := msk.Constrain(f)
	if err != nil {
		return nil, err
	}
	return skf, nil
}

// ABE.Enc(msk, x, msg) : It computes k ← CPRF.Eval(msk, x) and ct ← SKE.Enc(k, msg), and outputs ct
func ABEEnc(msk *rocprf.MasterKey, x []*big.Int, msg []byte) ([]byte, error){ // is it string?
	k := msk.Eval(x)
	ct, err := ske.Encrypt(k, msg)
	if err != nil {
		return nil, err
	}
	return ct, nil
}

// ABE.Dec(skf , ct) : It computes k ← CPRF.CEval(skf , x), and outputs SKE.Dec(k, ct).
func ABEDec(skf *rocprf.ConstrainedKey, x []*big.Int, ct []byte) ([]byte, error){
	k := skf.CEval(x)
	msg, err := ske.Decrypt(k, ct)
	if err != nil {
		return nil, err
	}
	return msg, nil
}