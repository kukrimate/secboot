// -*- Mode: Go; indent-tabs-mode: t -*-

/*
 * Copyright (C) 2024 Canonical Ltd
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3 as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package efi

import "crypto/x509"

// SecureBootAuthorityIdentity corresponds to the identify of a secure boot
// authority.
type SecureBootAuthorityIdentity struct {
	Subject            []byte
	SubjectKeyId       []byte
	PublicKeyAlgorithm x509.PublicKeyAlgorithm

	Issuer             []byte
	AuthorityKeyId     []byte
	SignatureAlgorithm x509.SignatureAlgorithm
}

var (
	// MSUefiCA2011 corresponds to the 2011 Microsoft UEFI CA
	MSUefiCA2011 = &SecureBootAuthorityIdentity{
		// CN=Microsoft Corporation UEFI CA 2011,O=Microsoft Corporation,L=Redmond,ST=Washington,C=US
		Subject: []byte{
			0x30, 0x81, 0x81, 0x31, 0x0b, 0x30, 0x09, 0x06, 0x03, 0x55,
			0x04, 0x06, 0x13, 0x02, 0x55, 0x53, 0x31, 0x13, 0x30, 0x11,
			0x06, 0x03, 0x55, 0x04, 0x08, 0x13, 0x0a, 0x57, 0x61, 0x73,
			0x68, 0x69, 0x6e, 0x67, 0x74, 0x6f, 0x6e, 0x31, 0x10, 0x30,
			0x0e, 0x06, 0x03, 0x55, 0x04, 0x07, 0x13, 0x07, 0x52, 0x65,
			0x64, 0x6d, 0x6f, 0x6e, 0x64, 0x31, 0x1e, 0x30, 0x1c, 0x06,
			0x03, 0x55, 0x04, 0x0a, 0x13, 0x15, 0x4d, 0x69, 0x63, 0x72,
			0x6f, 0x73, 0x6f, 0x66, 0x74, 0x20, 0x43, 0x6f, 0x72, 0x70,
			0x6f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x31, 0x2b, 0x30,
			0x29, 0x06, 0x03, 0x55, 0x04, 0x03, 0x13, 0x22, 0x4d, 0x69,
			0x63, 0x72, 0x6f, 0x73, 0x6f, 0x66, 0x74, 0x20, 0x43, 0x6f,
			0x72, 0x70, 0x6f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20,
			0x55, 0x45, 0x46, 0x49, 0x20, 0x43, 0x41, 0x20, 0x32, 0x30,
			0x31, 0x31,
		},
		SubjectKeyId: []byte{
			0x13, 0xad, 0xbf, 0x43, 0x09, 0xbd, 0x82, 0x70, 0x9c, 0x8c,
			0xd5, 0x4f, 0x31, 0x6e, 0xd5, 0x22, 0x98, 0x8a, 0x1b, 0xd4,
		},
		PublicKeyAlgorithm: x509.RSA,
		Issuer: []byte{
			0x30, 0x81, 0x91, 0x31, 0x0b, 0x30, 0x09, 0x06, 0x03, 0x55,
			0x04, 0x06, 0x13, 0x02, 0x55, 0x53, 0x31, 0x13, 0x30, 0x11,
			0x06, 0x03, 0x55, 0x04, 0x08, 0x13, 0x0a, 0x57, 0x61, 0x73,
			0x68, 0x69, 0x6e, 0x67, 0x74, 0x6f, 0x6e, 0x31, 0x10, 0x30,
			0x0e, 0x06, 0x03, 0x55, 0x04, 0x07, 0x13, 0x07, 0x52, 0x65,
			0x64, 0x6d, 0x6f, 0x6e, 0x64, 0x31, 0x1e, 0x30, 0x1c, 0x06,
			0x03, 0x55, 0x04, 0x0a, 0x13, 0x15, 0x4d, 0x69, 0x63, 0x72,
			0x6f, 0x73, 0x6f, 0x66, 0x74, 0x20, 0x43, 0x6f, 0x72, 0x70,
			0x6f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x31, 0x3b, 0x30,
			0x39, 0x06, 0x03, 0x55, 0x04, 0x03, 0x13, 0x32, 0x4d, 0x69,
			0x63, 0x72, 0x6f, 0x73, 0x6f, 0x66, 0x74, 0x20, 0x43, 0x6f,
			0x72, 0x70, 0x6f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20,
			0x54, 0x68, 0x69, 0x72, 0x64, 0x20, 0x50, 0x61, 0x72, 0x74,
			0x79, 0x20, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c,
			0x61, 0x63, 0x65, 0x20, 0x52, 0x6f, 0x6f, 0x74,
		},
		AuthorityKeyId: []byte{
			0x45, 0x66, 0x52, 0x43, 0xe1, 0x7e, 0x58, 0x11, 0xbf, 0xd6,
			0x4e, 0x9e, 0x23, 0x55, 0x08, 0x3b, 0x3a, 0x22, 0x6a, 0xa8,
		},
		SignatureAlgorithm: x509.SHA256WithRSA,
	}

	// MSUefiCA2023 corresponds to the 2023 Microsoft UEFI CA, which will eventually
	// replace the 2011 CA.
	MSUefiCA2023 = &SecureBootAuthorityIdentity{
		// CN=Microsoft UEFI CA 2023,O=Microsoft Corporation,C=US
		Subject: []byte{
			0x30, 0x4e, 0x31, 0x0b, 0x30, 0x09, 0x06, 0x03, 0x55, 0x04,
			0x06, 0x13, 0x02, 0x55, 0x53, 0x31, 0x1e, 0x30, 0x1c, 0x06,
			0x03, 0x55, 0x04, 0x0a, 0x13, 0x15, 0x4d, 0x69, 0x63, 0x72,
			0x6f, 0x73, 0x6f, 0x66, 0x74, 0x20, 0x43, 0x6f, 0x72, 0x70,
			0x6f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x31, 0x1f, 0x30,
			0x1d, 0x06, 0x03, 0x55, 0x04, 0x03, 0x13, 0x16, 0x4d, 0x69,
			0x63, 0x72, 0x6f, 0x73, 0x6f, 0x66, 0x74, 0x20, 0x55, 0x45,
			0x46, 0x49, 0x20, 0x43, 0x41, 0x20, 0x32, 0x30, 0x32, 0x33,
		},
		SubjectKeyId: []byte{
			0x81, 0xaa, 0x6b, 0x32, 0x44, 0xc9, 0x35, 0xbc, 0xe0, 0xd6,
			0x62, 0x8a, 0xf3, 0x98, 0x27, 0x42, 0x1e, 0x32, 0x49, 0x7d,
		},
		PublicKeyAlgorithm: x509.RSA,
		Issuer: []byte{
			0x30, 0x5a, 0x31, 0x0b, 0x30, 0x09, 0x06, 0x03, 0x55, 0x04,
			0x06, 0x13, 0x02, 0x55, 0x53, 0x31, 0x1e, 0x30, 0x1c, 0x06,
			0x03, 0x55, 0x04, 0x0a, 0x13, 0x15, 0x4d, 0x69, 0x63, 0x72,
			0x6f, 0x73, 0x6f, 0x66, 0x74, 0x20, 0x43, 0x6f, 0x72, 0x70,
			0x6f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x31, 0x2b, 0x30,
			0x29, 0x06, 0x03, 0x55, 0x04, 0x03, 0x13, 0x22, 0x4d, 0x69,
			0x63, 0x72, 0x6f, 0x73, 0x6f, 0x66, 0x74, 0x20, 0x52, 0x53,
			0x41, 0x20, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x20,
			0x52, 0x6f, 0x6f, 0x74, 0x20, 0x43, 0x41, 0x20, 0x32, 0x30,
			0x32, 0x31,
		},
		AuthorityKeyId: []byte{
			0x84, 0x44, 0x86, 0x06, 0x00, 0x98, 0x3f, 0x2c, 0xaa, 0xb3,
			0xc5, 0x89, 0xf3, 0xac, 0x2e, 0xc9, 0xe6, 0x9d, 0x09, 0x03,
		},
		SignatureAlgorithm: x509.SHA256WithRSA,
	}
)
