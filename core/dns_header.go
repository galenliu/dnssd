package core

//const KSizeBytes = 12
//
//const (
//	kIsResponseMask uint16 = 0x8000
//	kOpcodeMask     uint16 = 0x7000
//	kTruncationMask uint16 = 0x0400
//	kReturnCodeMask uint16 = 0x000F
//)
//
//const (
//	KMessageIdOffset       uint16 = 0
//	KFlagsOffset           uint16 = 2
//	KQueryCountOffset      uint16 = 4
//	KAnswerCountOffset     uint16 = 6
//	KAuthorityCountOffset  uint16 = 8
//	KAdditionalCountOffset uint16 = 10
//)
//
//// Flags BitPackedFlags 数据的标标志位 Flag
//
//type Flags struct {
//	QR     uint8
//	Opcode uint8
//	AA     uint8
//	TC     uint8
//	RD     uint8
//	RA     uint8
//	Z      uint8
//	AD     uint8
//	CD     uint8
//	RCODE  uint8
//}
//
//func isValidMdns(f uint16) bool {
//	return (f & (kOpcodeMask | kReturnCodeMask)) == 0
//}
//
///**
// * Allows operations on a DNS header. A DNS Header is defined in RFC 1035
// * and looks like this:
// *
// * | 0| 1 2 3 4| 5| 6| 7| 8| 9| 0| 1| 2 3 4 5 |
// * |               Message ID                 |
// * |QR| OPCODE |AA|TC|RD|RA| Z|AD|CD| RCODE   |
// * |       Items in QUESTION Section          |
// * |       Items in ANSWER Section            |
// * |       Items in AUTHORITY Section         |
// * |       Items in ADDITIONAL Section        |
// */
//
//type ConstHeaderRef struct {
//	ID      uint32
//	FLAGS   uint16
//	QDCOUNT uint16
//	ANCOUNT uint16
//	NSCOUNT uint16
//	ARCOUNT uint16
//}
//
//func (r ConstHeaderRef) GetMessageId() uint32 {
//	//return r.Get16At(kMessageIdOffset)
//	return r.ID
//}
//
//func (h *ConstHeaderRef) GetQueryCount() uint16 {
//	//return h.Get16At(kQueryCountOffset)
//	return h.QDCOUNT
//}
//
//func (h *ConstHeaderRef) GetAnswerCount() uint16 {
//	//return h.Get16At(kAnswerCountOffset)
//	return h.ANCOUNT
//}
//
//func (h *ConstHeaderRef) GetAuthorityCount() uint16 {
//	//return h.Get16At(kAuthorityCountOffset)
//	return h.NSCOUNT
//}
//
//func (h *ConstHeaderRef) GetAdditionalCount() uint16 {
//	//return h.Get16At(kAdditionalCountOffset)
//	return h.ARCOUNT
//}
//
//func (h *ConstHeaderRef) GetFlags() (flags Flags) {
//	flags.QR = uint8((h.FLAGS & 0b1000000000000000) >> 15)
//	flags.Opcode = uint8((h.FLAGS & 0b0111100000000000) >> 11)
//	flags.AA = uint8((h.FLAGS & 0b0000010000000000) >> 10)
//	flags.TC = uint8((h.FLAGS & 0b0000001000000000) >> 9)
//	flags.RD = uint8((h.FLAGS & 0b0000000100000000) >> 8)
//	flags.RA = uint8((h.FLAGS & 0b0000000010000000) >> 7)
//	flags.Z = uint8((h.FLAGS & 0b0000000001000000) >> 6)
//	flags.AD = uint8((h.FLAGS & 0b0000000000100000) >> 5)
//	flags.CD = uint8((h.FLAGS & 0b0000000000010000) >> 4)
//	flags.RCODE = uint8(h.FLAGS & 0b0000000000001111)
//	return
//}
//
//func (r *ConstHeaderRef) IsValidMdns() bool {
//	return isValidMdns(r.FLAGS)
//}
//
//type HeaderRef struct {
//	*ConstHeaderRef
//}
