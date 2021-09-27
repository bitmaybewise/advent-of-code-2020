package main

import (
	"testing"
)

func TestPassportIsValid(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		p := passport("ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm")
		if !p.IsValid() {
			t.Errorf("Passport %s should be valid", p)
		}
	})
	t.Run("Missing hgt", func(t *testing.T) {
		p := passport("iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884 hcl:#cfa07d byr:1929")
		if p.IsValid() {
			t.Errorf("Passport %s should be invalid", p)
		}
	})
	t.Run("Optional cid", func(t *testing.T) {
		p := passport("hcl:#ae17e1 iyr:2013 eyr:2024 ecl:brn pid:760753108 byr:1931 hgt:179cm")
		if !p.IsValid() {
			t.Errorf("Passport %s should be valid", p)
		}
	})
	t.Run("Missing cid and byr", func(t *testing.T) {
		p := passport("hcl:#cfa07d eyr:2025 pid:166559648 iyr:2011 ecl:brn hgt:59in")
		if p.IsValid() {
			t.Errorf("Passport %s should be invalid", p)
		}
	})
}

func TestPassportIsReallyValid(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		p := passport("ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm")
		if !p.IsReallyValid() {
			t.Errorf("Passport %s should be valid", p)
		}
	})
	t.Run("byr", func(t *testing.T) {
		p := passport("ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1919 iyr:2017 cid:147 hgt:183cm")
		if p.IsReallyValid() {
			t.Errorf("Passport %s should be invalid", p)
		}
		p = passport("ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:2003 iyr:2017 cid:147 hgt:183cm")
		if p.IsReallyValid() {
			t.Errorf("Passport %s should be invalid", p)
		}
	})
	t.Run("iyr", func(t *testing.T) {
		p := passport("ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2009 cid:147 hgt:183cm")
		if p.IsReallyValid() {
			t.Errorf("Passport %s should be invalid", p)
		}
		p = passport("ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2021 cid:147 hgt:183cm")
		if p.IsReallyValid() {
			t.Errorf("Passport %s should be invalid", p)
		}
	})
	t.Run("eyr", func(t *testing.T) {
		p := passport("ecl:gry pid:860033327 eyr:2019 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm")
		if p.IsReallyValid() {
			t.Errorf("Passport %s should be invalid", p)
		}
		p = passport("ecl:gry pid:860033327 eyr:2031 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm")
		if p.IsReallyValid() {
			t.Errorf("Passport %s should be invalid", p)
		}
	})
	t.Run("hgt", func(t *testing.T) {
		t.Run("cm", func(t *testing.T) {
			p := passport("ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:149cm")
			if p.IsReallyValid() {
				t.Errorf("Passport %s should be invalid", p)
			}
			p = passport("ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:194cm")
			if p.IsReallyValid() {
				t.Errorf("Passport %s should be invalid", p)
			}
		})
		t.Run("in", func(t *testing.T) {
			p := passport("ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:58in")
			if p.IsReallyValid() {
				t.Errorf("Passport %s should be invalid", p)
			}
			p = passport("ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:77in")
			if p.IsReallyValid() {
				t.Errorf("Passport %s should be invalid", p)
			}
		})
		t.Run("No measure unit", func(t *testing.T) {
			p := passport("ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:58")
			if p.IsReallyValid() {
				t.Errorf("Passport %s should be invalid", p)
			}
		})
	})
	t.Run("hcl", func(t *testing.T) {
		p := passport("ecl:gry pid:860033327 eyr:2020 hcl:#fff byr:1937 iyr:2017 cid:147 hgt:183cm")
		if p.IsReallyValid() {
			t.Errorf("Passport %s should be invalid", p)
		}
		p = passport("ecl:gry pid:860033327 eyr:2020 hcl:#0011 byr:1937 iyr:2017 cid:147 hgt:183cm")
		if p.IsReallyValid() {
			t.Errorf("Passport %s should be invalid", p)
		}
		p = passport("ecl:gry pid:860033327 eyr:2020 hcl:#0011zz byr:1937 iyr:2017 cid:147 hgt:183cm")
		if p.IsReallyValid() {
			t.Errorf("Passport %s should be invalid", p)
		}
	})
	t.Run("ecl", func(t *testing.T) {
		p := passport("ecl:aaa pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm")
		if p.IsReallyValid() {
			t.Errorf("Passport %s should be invalid", p)
		}
	})
	t.Run("pid", func(t *testing.T) {
		p := passport("ecl:gry pid:0012 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm")
		if p.IsReallyValid() {
			t.Errorf("Passport %s should be invalid", p)
		}
		p = passport("ecl:gry pid:aaaaaaaaa eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm")
		if p.IsReallyValid() {
			t.Errorf("Passport %s should be invalid", p)
		}
	})
}

func TestAnswer1(t *testing.T) {
	p := []passport{
		"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm",
		"iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884 hcl:#cfa07d byr:1929",
		"hcl:#ae17e1 iyr:2013 eyr:2024 ecl:brn pid:760753108 byr:1931 hgt:179cm",
		"hcl:#cfa07d eyr:2025 pid:166559648 iyr:2011 ecl:brn hgt:59in",
	}
	if result := answer1(p); result != 2 {
		t.Errorf("Should have %d valid passports, got %d", 2, result)
	}
}

func TestAnswer2(t *testing.T) {
	p := []passport{
		// invalid
		"eyr:1972 cid:100 hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926",
		"iyr:2019 hcl:#602927 eyr:1967 hgt:170cm ecl:grn pid:012533040 byr:1946",
		"hcl:dab227 iyr:2012 ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277",
		"hgt:59cm ecl:zzz eyr:2038 hcl:74454a iyr:2023 pid:3556412378 byr:2007",
		// valid
		"pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980 hcl:#623a2f",
		"eyr:2029 ecl:blu cid:129 byr:1989 iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm",
		"hcl:#888785 hgt:164cm byr:2001 iyr:2015 cid:88 pid:545766238 ecl:hzl eyr:2022",
		"iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719",
	}
	if result := answer2(p); result != 4 {
		t.Errorf("Should have %d valid passports, got %d", 4, result)
	}
}
