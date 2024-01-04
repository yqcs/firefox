// Code generated by 'mozilla extract crypto/x509'. DO NOT EDIT.

//go:build go1.20
// +build go1.20

package stdlib

import (
	"crypto/x509"
	"reflect"
)

func init() {
	Symbols["crypto/x509/x509"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"CANotAuthorizedForExtKeyUsage":             reflect.ValueOf(x509.CANotAuthorizedForExtKeyUsage),
		"CANotAuthorizedForThisName":                reflect.ValueOf(x509.CANotAuthorizedForThisName),
		"CreateCertificate":                         reflect.ValueOf(x509.CreateCertificate),
		"CreateCertificateRequest":                  reflect.ValueOf(x509.CreateCertificateRequest),
		"CreateRevocationList":                      reflect.ValueOf(x509.CreateRevocationList),
		"DSA":                                       reflect.ValueOf(x509.DSA),
		"DSAWithSHA1":                               reflect.ValueOf(x509.DSAWithSHA1),
		"DSAWithSHA256":                             reflect.ValueOf(x509.DSAWithSHA256),
		"DecryptPEMBlock":                           reflect.ValueOf(x509.DecryptPEMBlock),
		"ECDSA":                                     reflect.ValueOf(x509.ECDSA),
		"ECDSAWithSHA1":                             reflect.ValueOf(x509.ECDSAWithSHA1),
		"ECDSAWithSHA256":                           reflect.ValueOf(x509.ECDSAWithSHA256),
		"ECDSAWithSHA384":                           reflect.ValueOf(x509.ECDSAWithSHA384),
		"ECDSAWithSHA512":                           reflect.ValueOf(x509.ECDSAWithSHA512),
		"Ed25519":                                   reflect.ValueOf(x509.Ed25519),
		"EncryptPEMBlock":                           reflect.ValueOf(x509.EncryptPEMBlock),
		"ErrUnsupportedAlgorithm":                   reflect.ValueOf(&x509.ErrUnsupportedAlgorithm).Elem(),
		"Expired":                                   reflect.ValueOf(x509.Expired),
		"ExtKeyUsageAny":                            reflect.ValueOf(x509.ExtKeyUsageAny),
		"ExtKeyUsageClientAuth":                     reflect.ValueOf(x509.ExtKeyUsageClientAuth),
		"ExtKeyUsageCodeSigning":                    reflect.ValueOf(x509.ExtKeyUsageCodeSigning),
		"ExtKeyUsageEmailProtection":                reflect.ValueOf(x509.ExtKeyUsageEmailProtection),
		"ExtKeyUsageIPSECEndSystem":                 reflect.ValueOf(x509.ExtKeyUsageIPSECEndSystem),
		"ExtKeyUsageIPSECTunnel":                    reflect.ValueOf(x509.ExtKeyUsageIPSECTunnel),
		"ExtKeyUsageIPSECUser":                      reflect.ValueOf(x509.ExtKeyUsageIPSECUser),
		"ExtKeyUsageMicrosoftCommercialCodeSigning": reflect.ValueOf(x509.ExtKeyUsageMicrosoftCommercialCodeSigning),
		"ExtKeyUsageMicrosoftKernelCodeSigning":     reflect.ValueOf(x509.ExtKeyUsageMicrosoftKernelCodeSigning),
		"ExtKeyUsageMicrosoftServerGatedCrypto":     reflect.ValueOf(x509.ExtKeyUsageMicrosoftServerGatedCrypto),
		"ExtKeyUsageNetscapeServerGatedCrypto":      reflect.ValueOf(x509.ExtKeyUsageNetscapeServerGatedCrypto),
		"ExtKeyUsageOCSPSigning":                    reflect.ValueOf(x509.ExtKeyUsageOCSPSigning),
		"ExtKeyUsageServerAuth":                     reflect.ValueOf(x509.ExtKeyUsageServerAuth),
		"ExtKeyUsageTimeStamping":                   reflect.ValueOf(x509.ExtKeyUsageTimeStamping),
		"IncompatibleUsage":                         reflect.ValueOf(x509.IncompatibleUsage),
		"IncorrectPasswordError":                    reflect.ValueOf(&x509.IncorrectPasswordError).Elem(),
		"IsEncryptedPEMBlock":                       reflect.ValueOf(x509.IsEncryptedPEMBlock),
		"KeyUsageCRLSign":                           reflect.ValueOf(x509.KeyUsageCRLSign),
		"KeyUsageCertSign":                          reflect.ValueOf(x509.KeyUsageCertSign),
		"KeyUsageContentCommitment":                 reflect.ValueOf(x509.KeyUsageContentCommitment),
		"KeyUsageDataEncipherment":                  reflect.ValueOf(x509.KeyUsageDataEncipherment),
		"KeyUsageDecipherOnly":                      reflect.ValueOf(x509.KeyUsageDecipherOnly),
		"KeyUsageDigitalSignature":                  reflect.ValueOf(x509.KeyUsageDigitalSignature),
		"KeyUsageEncipherOnly":                      reflect.ValueOf(x509.KeyUsageEncipherOnly),
		"KeyUsageKeyAgreement":                      reflect.ValueOf(x509.KeyUsageKeyAgreement),
		"KeyUsageKeyEncipherment":                   reflect.ValueOf(x509.KeyUsageKeyEncipherment),
		"MD2WithRSA":                                reflect.ValueOf(x509.MD2WithRSA),
		"MD5WithRSA":                                reflect.ValueOf(x509.MD5WithRSA),
		"MarshalECPrivateKey":                       reflect.ValueOf(x509.MarshalECPrivateKey),
		"MarshalPKCS1PrivateKey":                    reflect.ValueOf(x509.MarshalPKCS1PrivateKey),
		"MarshalPKCS1PublicKey":                     reflect.ValueOf(x509.MarshalPKCS1PublicKey),
		"MarshalPKCS8PrivateKey":                    reflect.ValueOf(x509.MarshalPKCS8PrivateKey),
		"MarshalPKIXPublicKey":                      reflect.ValueOf(x509.MarshalPKIXPublicKey),
		"NameConstraintsWithoutSANs":                reflect.ValueOf(x509.NameConstraintsWithoutSANs),
		"NameMismatch":                              reflect.ValueOf(x509.NameMismatch),
		"NewCertPool":                               reflect.ValueOf(x509.NewCertPool),
		"NotAuthorizedToSign":                       reflect.ValueOf(x509.NotAuthorizedToSign),
		"PEMCipher3DES":                             reflect.ValueOf(x509.PEMCipher3DES),
		"PEMCipherAES128":                           reflect.ValueOf(x509.PEMCipherAES128),
		"PEMCipherAES192":                           reflect.ValueOf(x509.PEMCipherAES192),
		"PEMCipherAES256":                           reflect.ValueOf(x509.PEMCipherAES256),
		"PEMCipherDES":                              reflect.ValueOf(x509.PEMCipherDES),
		"ParseCRL":                                  reflect.ValueOf(x509.ParseCRL),
		"ParseCertificate":                          reflect.ValueOf(x509.ParseCertificate),
		"ParseCertificateRequest":                   reflect.ValueOf(x509.ParseCertificateRequest),
		"ParseCertificates":                         reflect.ValueOf(x509.ParseCertificates),
		"ParseDERCRL":                               reflect.ValueOf(x509.ParseDERCRL),
		"ParseECPrivateKey":                         reflect.ValueOf(x509.ParseECPrivateKey),
		"ParsePKCS1PrivateKey":                      reflect.ValueOf(x509.ParsePKCS1PrivateKey),
		"ParsePKCS1PublicKey":                       reflect.ValueOf(x509.ParsePKCS1PublicKey),
		"ParsePKCS8PrivateKey":                      reflect.ValueOf(x509.ParsePKCS8PrivateKey),
		"ParsePKIXPublicKey":                        reflect.ValueOf(x509.ParsePKIXPublicKey),
		"ParseRevocationList":                       reflect.ValueOf(x509.ParseRevocationList),
		"PureEd25519":                               reflect.ValueOf(x509.PureEd25519),
		"RSA":                                       reflect.ValueOf(x509.RSA),
		"SHA1WithRSA":                               reflect.ValueOf(x509.SHA1WithRSA),
		"SHA256WithRSA":                             reflect.ValueOf(x509.SHA256WithRSA),
		"SHA256WithRSAPSS":                          reflect.ValueOf(x509.SHA256WithRSAPSS),
		"SHA384WithRSA":                             reflect.ValueOf(x509.SHA384WithRSA),
		"SHA384WithRSAPSS":                          reflect.ValueOf(x509.SHA384WithRSAPSS),
		"SHA512WithRSA":                             reflect.ValueOf(x509.SHA512WithRSA),
		"SHA512WithRSAPSS":                          reflect.ValueOf(x509.SHA512WithRSAPSS),
		"SetFallbackRoots":                          reflect.ValueOf(x509.SetFallbackRoots),
		"SystemCertPool":                            reflect.ValueOf(x509.SystemCertPool),
		"TooManyConstraints":                        reflect.ValueOf(x509.TooManyConstraints),
		"TooManyIntermediates":                      reflect.ValueOf(x509.TooManyIntermediates),
		"UnconstrainedName":                         reflect.ValueOf(x509.UnconstrainedName),
		"UnknownPublicKeyAlgorithm":                 reflect.ValueOf(x509.UnknownPublicKeyAlgorithm),
		"UnknownSignatureAlgorithm":                 reflect.ValueOf(x509.UnknownSignatureAlgorithm),

		// type definitions
		"CertPool":                   reflect.ValueOf((*x509.CertPool)(nil)),
		"Certificate":                reflect.ValueOf((*x509.Certificate)(nil)),
		"CertificateInvalidError":    reflect.ValueOf((*x509.CertificateInvalidError)(nil)),
		"CertificateRequest":         reflect.ValueOf((*x509.CertificateRequest)(nil)),
		"ConstraintViolationError":   reflect.ValueOf((*x509.ConstraintViolationError)(nil)),
		"ExtKeyUsage":                reflect.ValueOf((*x509.ExtKeyUsage)(nil)),
		"HostnameError":              reflect.ValueOf((*x509.HostnameError)(nil)),
		"InsecureAlgorithmError":     reflect.ValueOf((*x509.InsecureAlgorithmError)(nil)),
		"InvalidReason":              reflect.ValueOf((*x509.InvalidReason)(nil)),
		"KeyUsage":                   reflect.ValueOf((*x509.KeyUsage)(nil)),
		"PEMCipher":                  reflect.ValueOf((*x509.PEMCipher)(nil)),
		"PublicKeyAlgorithm":         reflect.ValueOf((*x509.PublicKeyAlgorithm)(nil)),
		"RevocationList":             reflect.ValueOf((*x509.RevocationList)(nil)),
		"SignatureAlgorithm":         reflect.ValueOf((*x509.SignatureAlgorithm)(nil)),
		"SystemRootsError":           reflect.ValueOf((*x509.SystemRootsError)(nil)),
		"UnhandledCriticalExtension": reflect.ValueOf((*x509.UnhandledCriticalExtension)(nil)),
		"UnknownAuthorityError":      reflect.ValueOf((*x509.UnknownAuthorityError)(nil)),
		"VerifyOptions":              reflect.ValueOf((*x509.VerifyOptions)(nil)),
	}
}
