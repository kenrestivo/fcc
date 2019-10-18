package fcc

import (
	"strconv"
	"strings"
	"time"
)

const timeFormat = "01/02/2006 15:04:05"

type ErrZip struct {
	badValue string
}

func ReadRecord(record []byte) MinimalLicense {
	var ml MinimalLicense
	split := strings.Split(string(record), Sep)
	ml.Callsign = split[0]
	ml.Name = split[1]
	ml.Address = split[2]
	ml.City = split[3]
	ml.State = split[4]
	ml.ZIP = split[5]
	return ml
}

func (l *License) Minimal() MinimalLicense {
	return MinimalLicense{
		Callsign: l.Callsign,
		Name:     l.LicName,
		Address:  l.LicAddress,
		City:     l.LicCity,
		State:    l.LicState,
		ZIP:      l.LicZipCode,
	}
}

func (m *MinimalLicense) DiskFormat() []byte {
	join := strings.Join([]string{
		m.Callsign,
		m.Name,
		m.Address,
		m.City,
		m.State,
		m.ZIP}, Sep)
	return []byte(join)
}

func ParseLicense(line []string) (*License, error) {
	var err error
	var gd, ed, cd, lad time.Time
	var lid, frnI, fidI int
	var frnUi, fidUi uint32

	lic := License{}
	lid, err = strconv.Atoi(line[0])
	if err != nil {
		return &lic, err
	}
	lic.LicenseID = uint32(lid)

	lic.SourceSystem = line[1]

	lic.Callsign = line[2]
	if line[3] != "" {
		fidI, err = strconv.Atoi(line[3])
		if err != nil {
			return &lic, err
		}
		fidUi = uint32(fidI)
		lic.FacilityID = &fidUi
	}

	if line[4] != "" {
		frnI, err = strconv.Atoi(line[4])
		if err != nil {
			return &lic, err
		}
		frnUi = uint32(frnI)
		lic.FRN = &frnUi
	}

	lic.LicName = line[5]
	lic.CommonName = line[6]
	lic.RadioServiceCode = line[7]
	lic.RadioServiceDesc = line[8]
	lic.RollupCategoryCode = line[9]
	lic.RollupCategoryDesc = line[10]

	if line[11] != "" {
		gd, err = parseTime(line[11])
		if err != nil {
			return &lic, err
		}
		lic.GrantDate = &gd
	}

	if line[12] != "" {
		ed, err = parseTime(line[12])
		if err != nil {
			return &lic, err
		}
		lic.ExpiredDate = &ed
	}

	if line[13] != "" {
		cd, err = parseTime(line[13])
		if err != nil {
			return &lic, err
		}
		lic.CancellationDate = &cd
	}

	if line[14] != "" {
		lad, err = parseTime(line[14])
		if err != nil {
			return &lic, err
		}
		lic.LastActionDate = &lad
	}
	lic.LicStatusCode = line[15]
	lic.LicStatusDesc = line[16]
	lic.RollupStatusCode = line[17]
	lic.RollupStatusDesc = line[18]
	lic.EntityTypeCode = line[19]
	lic.EntityTypeDesc = line[20]
	lic.RollupEntityCode = line[21]
	lic.RollupEntityDesc = line[22]
	lic.LicAddress = line[23]
	lic.LicCity = line[24]
	lic.LicState = line[25]
	lic.LicZipCode = line[26]

	return &lic, nil
}

func parseTime(ts string) (time.Time, error) {
	return time.Parse(timeFormat, ts)
}
