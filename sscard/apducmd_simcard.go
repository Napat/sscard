package sscard

//////////////////////////////////////////////////////////////////////////
// Master file

// APDUReadICCID ...
var APDUReadICCID = [][]byte{
	{0xA0, 0xA4, 0x00, 0x00, 0x02, 0x3F, 0x00}, // Select MF
	{0xA0, 0xA4, 0x00, 0x00, 0x02, 0x2F, 0xE2}, // Select ICCID EF
	{0xA0, 0xB0, 0x00, 0x00, 0x0A},             // Read ICCID // 10 bytes
}

// APDUReadICCIDclass00 ...
var APDUReadICCIDclass00 = [][]byte{
	{0x00, 0xA4, 0x00, 0x04, 0x02, 0x3F, 0x00}, // Select MF
	{0x00, 0xA4, 0x00, 0x04, 0x02, 0x2F, 0xE2}, // Select ICCID EF
	{0x00, 0xB0, 0x00, 0x00, 0x0A},             // Read ICCID // 10
}

//////////////////////////////////////////////////////////////////////////
// DF GSM

// APDUReadIMSI ...
var APDUReadIMSI = [][]byte{
	{0xA0, 0xA4, 0x00, 0x00, 0x02, 0x3F, 0x00}, // Select MF
	{0xA0, 0xA4, 0x00, 0x00, 0x02, 0x7F, 0x20}, // Select GSM DF
	{0xA0, 0xA4, 0x00, 0x00, 0x02, 0x6F, 0x07}, // Select IMSI EF
	{0xA0, 0xB0, 0x00, 0x00, 0x09},             // Read IMSI // 9 bytes
}

// APDUReadIMSIclass00 ...
var APDUReadIMSIclass00 = [][]byte{
	{0x00, 0xA4, 0x00, 0x04, 0x02, 0x3F, 0x00}, // Select MF
	{0x00, 0xA4, 0x00, 0x04, 0x02, 0x7F, 0x20}, // Select GSM DF
	{0x00, 0xA4, 0x00, 0x04, 0x02, 0x6F, 0x07}, // Select IMSI EF
	{0x00, 0xB0, 0x00, 0x00, 0x09},             // Read IMSI // 9 bytes
}

// APDUReadKC ...
var APDUReadKC = [][]byte{
	{0xA0, 0xA4, 0x00, 0x00, 0x02, 0x3F, 0x00}, // Select MF
	{0xA0, 0xA4, 0x00, 0x00, 0x02, 0x7F, 0x20}, // Select GSM DF
	{0xA0, 0xA4, 0x00, 0x00, 0x02, 0x6F, 0x20}, // Select KC EF
	{0xA0, 0xB0, 0x00, 0x00, 0x09},             // Read KC // 9 bytes
}

// APDUReadHPPLMN ...
var APDUReadHPPLMN = [][]byte{
	{0xA0, 0xA4, 0x00, 0x00, 0x02, 0x3F, 0x00}, // Select MF
	{0xA0, 0xA4, 0x00, 0x00, 0x02, 0x7F, 0x20}, // Select GSM DF
	{0xA0, 0xA4, 0x00, 0x00, 0x02, 0x6F, 0x31}, // Select HPPLMN EF
	{0xA0, 0xB0, 0x00, 0x00, 0x01},             // Read HPPLMN 1 bytes
}

// APDUReadBCCH ...
var APDUReadBCCH = [][]byte{
	{0xA0, 0xA4, 0x00, 0x00, 0x02, 0x3F, 0x00}, // Select MF
	{0xA0, 0xA4, 0x00, 0x00, 0x02, 0x7F, 0x20}, // Select GSM DF
	{0xA0, 0xA4, 0x00, 0x00, 0x02, 0x6F, 0x74}, // Select BCCH EF
	{0xA0, 0xB0, 0x00, 0x00, 0x10},             // Read BCCH 16 bytes
}

// APDUReadACC ...
var APDUReadACC = [][]byte{
	{0xA0, 0xA4, 0x00, 0x00, 0x02, 0x3F, 0x00}, // Select MF
	{0xA0, 0xA4, 0x00, 0x00, 0x02, 0x7F, 0x20}, // Select GSM DF
	{0xA0, 0xA4, 0x00, 0x00, 0x02, 0x6F, 0x78}, // Select ACC EF
	{0xA0, 0xB0, 0x00, 0x00, 0x02},             // Read ACC 2 bytes
}

// APDUReadFPLMN ...
var APDUReadFPLMN = [][]byte{
	{0xA0, 0xA4, 0x00, 0x00, 0x02, 0x3F, 0x00}, // Select MF
	{0xA0, 0xA4, 0x00, 0x00, 0x02, 0x7F, 0x20}, // Select GSM DF
	{0xA0, 0xA4, 0x00, 0x00, 0x02, 0x6F, 0x7B}, // Select FPLMN EF
	{0xA0, 0xB0, 0x00, 0x00, 0x0C},             // Read FPLMN 12 bytes
}

// APDUReadLOCI ...
var APDUReadLOCI = [][]byte{
	{0xA0, 0xA4, 0x00, 0x00, 0x02, 0x3F, 0x00}, // Select MF
	{0xA0, 0xA4, 0x00, 0x00, 0x02, 0x7F, 0x20}, // Select GSM DF
	{0xA0, 0xA4, 0x00, 0x00, 0x02, 0x6F, 0x7E}, // Select LOCI EF
	{0xA0, 0xB0, 0x00, 0x00, 0x0B},             // Read LOCI 11 bytes
}

// APDUReadAD ...
var APDUReadAD = [][]byte{
	{0xA0, 0xA4, 0x00, 0x00, 0x02, 0x3F, 0x00}, // Select MF
	{0xA0, 0xA4, 0x00, 0x00, 0x02, 0x7F, 0x20}, // Select GSM DF
	{0xA0, 0xA4, 0x00, 0x00, 0x02, 0x6F, 0xAD}, // Select AD EF
	{0xA0, 0xB0, 0x00, 0x00, 0x03},             // Read AD 3+X byte
}

// APDUReadPHASE ...
var APDUReadPHASE = [][]byte{
	{0xA0, 0xA4, 0x00, 0x00, 0x02, 0x3F, 0x00}, // Select MF
	{0xA0, 0xA4, 0x00, 0x00, 0x02, 0x7F, 0x20}, // Select GSM DF
	{0xA0, 0xA4, 0x00, 0x00, 0x02, 0x6F, 0xAE}, // Select PHASE EF
	{0xA0, 0xB0, 0x00, 0x00, 0x01},             // Read PHASE 1 byte
}

//////////////////////////////////////////////////////////////////////////
// DF TELECOM

// APDUReadADN ...
var APDUReadADN = [][]byte{
	{0xA0, 0xA4, 0x00, 0x00, 0x02, 0x3F, 0x00}, // Select MF
	{0xA0, 0xA4, 0x00, 0x00, 0x02, 0x7F, 0x10}, // Select TELECOM DF
	{0xA0, 0xA4, 0x00, 0x00, 0x02, 0x6F, 0x3A}, // Select ADN EF
	{0xA0, 0xB0, 0x00, 0x00, 0x0D},             // Read ADN X+14 bytes
}

// APDUReadFDN ...
var APDUReadFDN = [][]byte{
	{0xA0, 0xA4, 0x00, 0x00, 0x02, 0x3F, 0x00}, // Select MF
	{0xA0, 0xA4, 0x00, 0x00, 0x02, 0x7F, 0x10}, // Select TELECOM DF
	{0xA0, 0xA4, 0x00, 0x00, 0x02, 0x6F, 0x0B}, // Select FDN EF
	{0xA0, 0xB0, 0x00, 0x00, 0x0D},             // Read FDN X+14 bytes
}

// APDUReadSMS ...
var APDUReadSMS = [][]byte{
	{0xA0, 0xA4, 0x00, 0x00, 0x02, 0x3F, 0x00}, // Select MF
	{0xA0, 0xA4, 0x00, 0x00, 0x02, 0x7F, 0x10}, // Select TELECOM DF
	{0xA0, 0xA4, 0x00, 0x00, 0x02, 0x6F, 0x3C}, // Select SMS EF
	{0xA0, 0xB2, 0x01, 0x04, 0xB0},             // Read SMS 176 bytes
}

// APDUReadMSISDN ...
var APDUReadMSISDN = [][]byte{
	{0xA0, 0xA4, 0x00, 0x00, 0x02, 0x3F, 0x00}, // Select MF
	{0xA0, 0xA4, 0x00, 0x00, 0x02, 0x7F, 0x10}, // Select TELECOM DF
	{0xA0, 0xA4, 0x00, 0x00, 0x02, 0x6F, 0x40}, // Select MSISDN EF
	{0xA0, 0xB0, 0x00, 0x00, 0x0D},             // Read MSISDN X+14 bytes
}

// APDUReadSMSP ...
var APDUReadSMSP = [][]byte{
	{0xA0, 0xA4, 0x00, 0x00, 0x02, 0x3F, 0x00}, // Select MF
	{0xA0, 0xA4, 0x00, 0x00, 0x02, 0x7F, 0x10}, // Select TELECOM DF
	{0xA0, 0xA4, 0x00, 0x00, 0x02, 0x6F, 0x42}, // Select SMSP EF
	{0xA0, 0xB0, 0x00, 0x00, 0x0D},             // Read SMSP X+14 bytes
}
