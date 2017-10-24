// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package testdata

var PEMBytes = map[string][]byte{
	"dsa": []byte(`-----BEGIN DSA PRIVATE KEY-----
MIIBuwIBAAKBgQD6PDSEyXiI9jfNs97WuM46MSDCYlOqWw80ajN16AohtBncs1YB
lHk//dQOvCYOsYaE+gNix2jtoRjwXhDsc25/IqQbU1ahb7mB8/rsaILRGIbA5WH3
EgFtJmXFovDz3if6F6TzvhFpHgJRmLYVR8cqsezL3hEZOvvs2iH7MorkxwIVAJHD
nD82+lxh2fb4PMsIiaXudAsBAoGAQRf7Q/iaPRn43ZquUhd6WwvirqUj+tkIu6eV
2nZWYmXLlqFQKEy4Tejl7Wkyzr2OSYvbXLzo7TNxLKoWor6ips0phYPPMyXld14r
juhT24CrhOzuLMhDduMDi032wDIZG4Y+K7ElU8Oufn8Sj5Wge8r6ANmmVgmFfynr
FhdYCngCgYEA3ucGJ93/Mx4q4eKRDxcWD3QzWyqpbRVRRV1Vmih9Ha/qC994nJFz
DQIdjxDIT2Rk2AGzMqFEB68Zc3O+Wcsmz5eWWzEwFxaTwOGWTyDqsDRLm3fD+QYj
nOwuxb0Kce+gWI8voWcqC9cyRm09jGzu2Ab3Bhtpg8JJ8L7gS3MRZK4CFEx4UAfY
Fmsr0W6fHB9nhS4/UXM8
-----END DSA PRIVATE KEY-----
`),
	"ecdsa": []byte(`-----BEGIN EC PRIVATE KEY-----
MHcCAQEEINGWx0zo6fhJ/0EAfrPzVFyFC9s18lBt3cRoEDhS3ARooAoGCCqGSM49
AwEHoUQDQgAEi9Hdw6KvZcWxfg2IDhA7UkpDtzzt6ZqJXSsFdLd+Kx4S3Sx4cVO+
6/ZOXRnPmNAlLUqjShUsUBBngG0u2fqEqA==
-----END EC PRIVATE KEY-----
`),
	"rsa": []byte(`-----BEGIN RSA PRIVATE KEY-----
MIICXAIBAAKBgQC8A6FGHDiWCSREAXCq6yBfNVr0xCVG2CzvktFNRpue+RXrGs/2
a6ySEJQb3IYquw7HlJgu6fg3WIWhOmHCjfpG0PrL4CRwbqQ2LaPPXhJErWYejcD8
Di00cF3677+G10KMZk9RXbmHtuBFZT98wxg8j+ZsBMqGM1+7yrWUvynswQIDAQAB
AoGAJMCk5vqfSRzyXOTXLGIYCuR4Kj6pdsbNSeuuRGfYBeR1F2c/XdFAg7D/8s5R
38p/Ih52/Ty5S8BfJtwtvgVY9ecf/JlU/rl/QzhG8/8KC0NG7KsyXklbQ7gJT8UT
Ojmw5QpMk+rKv17ipDVkQQmPaj+gJXYNAHqImke5mm/K/h0CQQDciPmviQ+DOhOq
2ZBqUfH8oXHgFmp7/6pXw80DpMIxgV3CwkxxIVx6a8lVH9bT/AFySJ6vXq4zTuV9
6QmZcZzDAkEA2j/UXJPIs1fQ8z/6sONOkU/BjtoePFIWJlRxdN35cZjXnBraX5UR
fFHkePv4YwqmXNqrBOvSu+w2WdSDci+IKwJAcsPRc/jWmsrJW1q3Ha0hSf/WG/Bu
X7MPuXaKpP/DkzGoUmb8ks7yqj6XWnYkPNLjCc8izU5vRwIiyWBRf4mxMwJBAILa
NDvRS0rjwt6lJGv7zPZoqDc65VfrK2aNyHx2PgFyzwrEOtuF57bu7pnvEIxpLTeM
z26i6XVMeYXAWZMTloMCQBbpGgEERQpeUknLBqUHhg/wXF6+lFA+vEGnkY+Dwab2
KCXFGd+SQ5GdUcEMe9isUH6DYj/6/yCDoFrXXmpQb+M=
-----END RSA PRIVATE KEY-----
`),
	"ed25519": []byte(`-----BEGIN OPENSSH PRIVATE KEY-----
b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW
QyNTUxOQAAACA+3f7hS7g5UWwXOGVTrMfhmxyrjqz7Sxxbx7I1j8DvvwAAAJhAFfkOQBX5
DgAAAAtzc2gtZWQyNTUxOQAAACA+3f7hS7g5UWwXOGVTrMfhmxyrjqz7Sxxbx7I1j8Dvvw
AAAEAaYmXltfW6nhRo3iWGglRB48lYq0z0Q3I3KyrdutEr6j7d/uFLuDlRbBc4ZVOsx+Gb
HKuOrPtLHFvHsjWPwO+/AAAAE2dhcnRvbm1AZ2FydG9ubS14cHMBAg==
-----END OPENSSH PRIVATE KEY-----
`),
	"user": []byte(`-----BEGIN EC PRIVATE KEY-----
MHcCAQEEILYCAeq8f7V4vSSypRw7pxy8yz3V5W4qg8kSC3zJhqpQoAoGCCqGSM49
AwEHoUQDQgAEYcO2xNKiRUYOLEHM7VYAp57HNyKbOdYtHD83Z4hzNPVC4tM5mdGD
PLL8IEwvYu2wq+lpXfGQnNMbzYf9gspG0w==
-----END EC PRIVATE KEY-----
`),
}

var PEMEncryptedKeys = []struct {
	Name          string
	EncryptionKey string
	PEMBytes      []byte
}{
	0: {
		Name:          "rsa-encrypted",
		EncryptionKey: "r54-G0pher_t3st$",
		PEMBytes: []byte(`-----BEGIN RSA PRIVATE KEY-----
Proc-Type: 4,ENCRYPTED
DEK-Info: AES-128-CBC,3E1714DE130BC5E81327F36564B05462

MqW88sud4fnWk/Jk3fkjh7ydu51ZkHLN5qlQgA4SkAXORPPMj2XvqZOv1v2LOgUV
dUevUn8PZK7a9zbZg4QShUSzwE5k6wdB7XKPyBgI39mJ79GBd2U4W3h6KT6jIdWA
goQpluxkrzr2/X602IaxLEre97FT9mpKC6zxKCLvyFWVIP9n3OSFS47cTTXyFr+l
7PdRhe60nn6jSBgUNk/Q1lAvEQ9fufdPwDYY93F1wyJ6lOr0F1+mzRrMbH67NyKs
rG8J1Fa7cIIre7ueKIAXTIne7OAWqpU9UDgQatDtZTbvA7ciqGsSFgiwwW13N+Rr
hN8MkODKs9cjtONxSKi05s206A3NDU6STtZ3KuPDjFE1gMJODotOuqSM+cxKfyFq
wxpk/CHYCDdMAVBSwxb/vraOHamylL4uCHpJdBHypzf2HABt+lS8Su23uAmL87DR
yvyCS/lmpuNTndef6qHPRkoW2EV3xqD3ovosGf7kgwGJUk2ZpCLVteqmYehKlZDK
r/Jy+J26ooI2jIg9bjvD1PZq+Mv+2dQ1RlDrPG3PB+rEixw6vBaL9x3jatCd4ej7
XG7lb3qO9xFpLsx89tkEcvpGR+broSpUJ6Mu5LBCVmrvqHjvnDhrZVz1brMiQtU9
iMZbgXqDLXHd6ERWygk7OTU03u+l1gs+KGMfmS0h0ZYw6KGVLgMnsoxqd6cFSKNB
8Ohk9ZTZGCiovlXBUepyu8wKat1k8YlHSfIHoRUJRhhcd7DrmojC+bcbMIZBU22T
Pl2ftVRGtcQY23lYd0NNKfebF7ncjuLWQGy+vZW+7cgfI6wPIbfYfP6g7QAutk6W
KQx0AoX5woZ6cNxtpIrymaVjSMRRBkKQrJKmRp3pC/lul5E5P2cueMs1fj4OHTbJ
lAUv88ywr+R+mRgYQlFW/XQ653f6DT4t6+njfO9oBcPrQDASZel3LjXLpjjYG/N5
+BWnVexuJX9ika8HJiFl55oqaKb+WknfNhk5cPY+x7SDV9ywQeMiDZpr0ffeYAEP
LlwwiWRDYpO+uwXHSFF3+JjWwjhs8m8g99iFb7U93yKgBB12dCEPPa2ZeH9wUHMJ
sreYhNuq6f4iWWSXpzN45inQqtTi8jrJhuNLTT543ErW7DtntBO2rWMhff3aiXbn
Uy3qzZM1nPbuCGuBmP9L2dJ3Z5ifDWB4JmOyWY4swTZGt9AVmUxMIKdZpRONx8vz
I9u9nbVPGZBcou50Pa0qTLbkWsSL94MNXrARBxzhHC9Zs6XNEtwN7mOuii7uMkVc
adrxgknBH1J1N+NX/eTKzUwJuPvDtA+Z5ILWNN9wpZT/7ed8zEnKHPNUexyeT5g3
uw9z9jH7ffGxFYlx87oiVPHGOrCXYZYW5uoZE31SCBkbtNuffNRJRKIFeipmpJ3P
7bpAG+kGHMelQH6b+5K1Qgsv4tpuSyKeTKpPFH9Av5nN4P1ZBm9N80tzbNWqjSJm
S7rYdHnuNEVnUGnRmEUMmVuYZnNBEVN/fP2m2SEwXcP3Uh7TiYlcWw10ygaGmOr7
MvMLGkYgQ4Utwnd98mtqa0jr0hK2TcOSFir3AqVvXN3XJj4cVULkrXe4Im1laWgp
-----END RSA PRIVATE KEY-----
`),
	},

	1: {
		Name:          "dsa-encrypted",
		EncryptionKey: "qG0pher-dsa_t3st$",
		PEMBytes: []byte(`-----BEGIN DSA PRIVATE KEY-----
Proc-Type: 4,ENCRYPTED
DEK-Info: AES-128-CBC,7CE7A6E4A647DC01AF860210B15ADE3E

hvnBpI99Hceq/55pYRdOzBLntIEis02JFNXuLEydWL+RJBFDn7tA+vXec0ERJd6J
G8JXlSOAhmC2H4uK3q2xR8/Y3yL95n6OIcjvCBiLsV+o3jj1MYJmErxP6zRtq4w3
JjIjGHWmaYFSxPKQ6e8fs74HEqaeMV9ONUoTtB+aISmgaBL15Fcoayg245dkBvVl
h5Kqspe7yvOBmzA3zjRuxmSCqKJmasXM7mqs3vIrMxZE3XPo1/fWKcPuExgpVQoT
HkJZEoIEIIPnPMwT2uYbFJSGgPJVMDT84xz7yvjCdhLmqrsXgs5Qw7Pw0i0c0BUJ
b7fDJ2UhdiwSckWGmIhTLlJZzr8K+JpjCDlP+REYBI5meB7kosBnlvCEHdw2EJkH
0QDc/2F4xlVrHOLbPRFyu1Oi2Gvbeoo9EsM/DThpd1hKAlb0sF5Y0y0d+owv0PnE
R/4X3HWfIdOHsDUvJ8xVWZ4BZk9Zk9qol045DcFCehpr/3hslCrKSZHakLt9GI58
vVQJ4L0aYp5nloLfzhViZtKJXRLkySMKdzYkIlNmW1oVGl7tce5UCNI8Nok4j6yn
IiHM7GBn+0nJoKTXsOGMIBe3ulKlKVxLjEuk9yivh/8=
-----END DSA PRIVATE KEY-----
`),
	},
}
