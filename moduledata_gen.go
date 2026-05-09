// Copyright (C) 2019-2026 GoRE Authors.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package gore

import (
	"fmt"
	"sort"
)

type moduledata_1_5_32 struct {
	Pclntable, Pclntablelen, Pclntablecap uint32
	Ftab, Ftablen, Ftabcap                uint32
	Filetab, Filetablen, Filetabcap       uint32
	Findfunctab                           uint32
	Minpc                                 uint32
	Maxpc                                 uint32
	Text                                  uint32
	Etext                                 uint32
	Noptrdata                             uint32
	Enoptrdata                            uint32
	Data                                  uint32
	Edata                                 uint32
	Bss                                   uint32
	Ebss                                  uint32
	Noptrbss                              uint32
	Enoptrbss                             uint32
	End                                   uint32
	Gcdata                                uint32
	Gcbss                                 uint32
	Typelinks, Typelinkslen, Typelinkscap uint32
}

func (md moduledata_1_5_32) toModuledata() moduledata {
	return moduledata{
		TextAddr:      uint64(md.Text),
		TextLen:       uint64(md.Etext - md.Text),
		NoPtrDataAddr: uint64(md.Noptrdata),
		NoPtrDataLen:  uint64(md.Enoptrdata - md.Noptrdata),
		DataAddr:      uint64(md.Data),
		DataLen:       uint64(md.Edata - md.Data),
		BssAddr:       uint64(md.Bss),
		BssLen:        uint64(md.Ebss - md.Bss),
		NoPtrBssAddr:  uint64(md.Noptrbss),
		NoPtrBssLen:   uint64(md.Enoptrbss - md.Noptrbss),
		TypelinkAddr:  uint64(md.Typelinks),
		TypelinkLen:   uint64(md.Typelinkslen),
		FuncTabAddr:   uint64(md.Ftab),
		FuncTabLen:    uint64(md.Ftablen),
		PCLNTabAddr:   uint64(md.Pclntable),
		PCLNTabLen:    uint64(md.Pclntablelen),
	}
}

func (md moduledata_1_5_32) pointerOffsets() []int {
	return []int{
		0,
		12,
		24,
		36,
		40,
		44,
		48,
		52,
		56,
		60,
		64,
		68,
		72,
		76,
		80,
		84,
		88,
		92,
		96,
		100,
	}
}

type moduledata_1_5_64 struct {
	Pclntable, Pclntablelen, Pclntablecap uint64
	Ftab, Ftablen, Ftabcap                uint64
	Filetab, Filetablen, Filetabcap       uint64
	Findfunctab                           uint64
	Minpc                                 uint64
	Maxpc                                 uint64
	Text                                  uint64
	Etext                                 uint64
	Noptrdata                             uint64
	Enoptrdata                            uint64
	Data                                  uint64
	Edata                                 uint64
	Bss                                   uint64
	Ebss                                  uint64
	Noptrbss                              uint64
	Enoptrbss                             uint64
	End                                   uint64
	Gcdata                                uint64
	Gcbss                                 uint64
	Typelinks, Typelinkslen, Typelinkscap uint64
}

func (md moduledata_1_5_64) toModuledata() moduledata {
	return moduledata{
		TextAddr:      md.Text,
		TextLen:       md.Etext - md.Text,
		NoPtrDataAddr: md.Noptrdata,
		NoPtrDataLen:  md.Enoptrdata - md.Noptrdata,
		DataAddr:      md.Data,
		DataLen:       md.Edata - md.Data,
		BssAddr:       md.Bss,
		BssLen:        md.Ebss - md.Bss,
		NoPtrBssAddr:  md.Noptrbss,
		NoPtrBssLen:   md.Enoptrbss - md.Noptrbss,
		TypelinkAddr:  md.Typelinks,
		TypelinkLen:   md.Typelinkslen,
		FuncTabAddr:   md.Ftab,
		FuncTabLen:    md.Ftablen,
		PCLNTabAddr:   md.Pclntable,
		PCLNTabLen:    md.Pclntablelen,
	}
}

func (md moduledata_1_5_64) pointerOffsets() []int {
	return []int{
		0,
		24,
		48,
		72,
		80,
		88,
		96,
		104,
		112,
		120,
		128,
		136,
		144,
		152,
		160,
		168,
		176,
		184,
		192,
		200,
	}
}

type moduledata_1_7_32 struct {
	Pclntable, Pclntablelen, Pclntablecap uint32
	Ftab, Ftablen, Ftabcap                uint32
	Filetab, Filetablen, Filetabcap       uint32
	Findfunctab                           uint32
	Minpc                                 uint32
	Maxpc                                 uint32
	Text                                  uint32
	Etext                                 uint32
	Noptrdata                             uint32
	Enoptrdata                            uint32
	Data                                  uint32
	Edata                                 uint32
	Bss                                   uint32
	Ebss                                  uint32
	Noptrbss                              uint32
	Enoptrbss                             uint32
	End                                   uint32
	Gcdata                                uint32
	Gcbss                                 uint32
	Types                                 uint32
	Etypes                                uint32
	Typelinks, Typelinkslen, Typelinkscap uint32
	Itablinks, Itablinkslen, Itablinkscap uint32
}

func (md moduledata_1_7_32) toModuledata() moduledata {
	return moduledata{
		TextAddr:      uint64(md.Text),
		TextLen:       uint64(md.Etext - md.Text),
		NoPtrDataAddr: uint64(md.Noptrdata),
		NoPtrDataLen:  uint64(md.Enoptrdata - md.Noptrdata),
		DataAddr:      uint64(md.Data),
		DataLen:       uint64(md.Edata - md.Data),
		BssAddr:       uint64(md.Bss),
		BssLen:        uint64(md.Ebss - md.Bss),
		NoPtrBssAddr:  uint64(md.Noptrbss),
		NoPtrBssLen:   uint64(md.Enoptrbss - md.Noptrbss),
		TypesAddr:     uint64(md.Types),
		TypesLen:      uint64(md.Etypes - md.Types),
		TypelinkAddr:  uint64(md.Typelinks),
		TypelinkLen:   uint64(md.Typelinkslen),
		ITabLinkAddr:  uint64(md.Itablinks),
		ITabLinkLen:   uint64(md.Itablinkslen),
		FuncTabAddr:   uint64(md.Ftab),
		FuncTabLen:    uint64(md.Ftablen),
		PCLNTabAddr:   uint64(md.Pclntable),
		PCLNTabLen:    uint64(md.Pclntablelen),
	}
}

func (md moduledata_1_7_32) pointerOffsets() []int {
	return []int{
		0,
		12,
		24,
		36,
		40,
		44,
		48,
		52,
		56,
		60,
		64,
		68,
		72,
		76,
		80,
		84,
		88,
		92,
		96,
		100,
		104,
		108,
		120,
	}
}

type moduledata_1_7_64 struct {
	Pclntable, Pclntablelen, Pclntablecap uint64
	Ftab, Ftablen, Ftabcap                uint64
	Filetab, Filetablen, Filetabcap       uint64
	Findfunctab                           uint64
	Minpc                                 uint64
	Maxpc                                 uint64
	Text                                  uint64
	Etext                                 uint64
	Noptrdata                             uint64
	Enoptrdata                            uint64
	Data                                  uint64
	Edata                                 uint64
	Bss                                   uint64
	Ebss                                  uint64
	Noptrbss                              uint64
	Enoptrbss                             uint64
	End                                   uint64
	Gcdata                                uint64
	Gcbss                                 uint64
	Types                                 uint64
	Etypes                                uint64
	Typelinks, Typelinkslen, Typelinkscap uint64
	Itablinks, Itablinkslen, Itablinkscap uint64
}

func (md moduledata_1_7_64) toModuledata() moduledata {
	return moduledata{
		TextAddr:      md.Text,
		TextLen:       md.Etext - md.Text,
		NoPtrDataAddr: md.Noptrdata,
		NoPtrDataLen:  md.Enoptrdata - md.Noptrdata,
		DataAddr:      md.Data,
		DataLen:       md.Edata - md.Data,
		BssAddr:       md.Bss,
		BssLen:        md.Ebss - md.Bss,
		NoPtrBssAddr:  md.Noptrbss,
		NoPtrBssLen:   md.Enoptrbss - md.Noptrbss,
		TypesAddr:     md.Types,
		TypesLen:      md.Etypes - md.Types,
		TypelinkAddr:  md.Typelinks,
		TypelinkLen:   md.Typelinkslen,
		ITabLinkAddr:  md.Itablinks,
		ITabLinkLen:   md.Itablinkslen,
		FuncTabAddr:   md.Ftab,
		FuncTabLen:    md.Ftablen,
		PCLNTabAddr:   md.Pclntable,
		PCLNTabLen:    md.Pclntablelen,
	}
}

func (md moduledata_1_7_64) pointerOffsets() []int {
	return []int{
		0,
		24,
		48,
		72,
		80,
		88,
		96,
		104,
		112,
		120,
		128,
		136,
		144,
		152,
		160,
		168,
		176,
		184,
		192,
		200,
		208,
		216,
		240,
	}
}

type moduledata_1_8_32 struct {
	Pclntable, Pclntablelen, Pclntablecap       uint32
	Ftab, Ftablen, Ftabcap                      uint32
	Filetab, Filetablen, Filetabcap             uint32
	Findfunctab                                 uint32
	Minpc                                       uint32
	Maxpc                                       uint32
	Text                                        uint32
	Etext                                       uint32
	Noptrdata                                   uint32
	Enoptrdata                                  uint32
	Data                                        uint32
	Edata                                       uint32
	Bss                                         uint32
	Ebss                                        uint32
	Noptrbss                                    uint32
	Enoptrbss                                   uint32
	End                                         uint32
	Gcdata                                      uint32
	Gcbss                                       uint32
	Types                                       uint32
	Etypes                                      uint32
	Textsectmap, Textsectmaplen, Textsectmapcap uint32
	Typelinks, Typelinkslen, Typelinkscap       uint32
	Itablinks, Itablinkslen, Itablinkscap       uint32
	Ptab, Ptablen, Ptabcap                      uint32
	Pluginpath, Pluginpathlen                   uint32
	Pkghashes, Pkghasheslen, Pkghashescap       uint32
}

func (md moduledata_1_8_32) toModuledata() moduledata {
	return moduledata{
		TextAddr:      uint64(md.Text),
		TextLen:       uint64(md.Etext - md.Text),
		NoPtrDataAddr: uint64(md.Noptrdata),
		NoPtrDataLen:  uint64(md.Enoptrdata - md.Noptrdata),
		DataAddr:      uint64(md.Data),
		DataLen:       uint64(md.Edata - md.Data),
		BssAddr:       uint64(md.Bss),
		BssLen:        uint64(md.Ebss - md.Bss),
		NoPtrBssAddr:  uint64(md.Noptrbss),
		NoPtrBssLen:   uint64(md.Enoptrbss - md.Noptrbss),
		TypesAddr:     uint64(md.Types),
		TypesLen:      uint64(md.Etypes - md.Types),
		TypelinkAddr:  uint64(md.Typelinks),
		TypelinkLen:   uint64(md.Typelinkslen),
		ITabLinkAddr:  uint64(md.Itablinks),
		ITabLinkLen:   uint64(md.Itablinkslen),
		FuncTabAddr:   uint64(md.Ftab),
		FuncTabLen:    uint64(md.Ftablen),
		PCLNTabAddr:   uint64(md.Pclntable),
		PCLNTabLen:    uint64(md.Pclntablelen),
	}
}

func (md moduledata_1_8_32) pointerOffsets() []int {
	return []int{
		0,
		12,
		24,
		36,
		40,
		44,
		48,
		52,
		56,
		60,
		64,
		68,
		72,
		76,
		80,
		84,
		88,
		92,
		96,
		100,
		104,
		108,
		120,
		132,
		144,
		156,
		164,
	}
}

type moduledata_1_8_64 struct {
	Pclntable, Pclntablelen, Pclntablecap       uint64
	Ftab, Ftablen, Ftabcap                      uint64
	Filetab, Filetablen, Filetabcap             uint64
	Findfunctab                                 uint64
	Minpc                                       uint64
	Maxpc                                       uint64
	Text                                        uint64
	Etext                                       uint64
	Noptrdata                                   uint64
	Enoptrdata                                  uint64
	Data                                        uint64
	Edata                                       uint64
	Bss                                         uint64
	Ebss                                        uint64
	Noptrbss                                    uint64
	Enoptrbss                                   uint64
	End                                         uint64
	Gcdata                                      uint64
	Gcbss                                       uint64
	Types                                       uint64
	Etypes                                      uint64
	Textsectmap, Textsectmaplen, Textsectmapcap uint64
	Typelinks, Typelinkslen, Typelinkscap       uint64
	Itablinks, Itablinkslen, Itablinkscap       uint64
	Ptab, Ptablen, Ptabcap                      uint64
	Pluginpath, Pluginpathlen                   uint64
	Pkghashes, Pkghasheslen, Pkghashescap       uint64
}

func (md moduledata_1_8_64) toModuledata() moduledata {
	return moduledata{
		TextAddr:      md.Text,
		TextLen:       md.Etext - md.Text,
		NoPtrDataAddr: md.Noptrdata,
		NoPtrDataLen:  md.Enoptrdata - md.Noptrdata,
		DataAddr:      md.Data,
		DataLen:       md.Edata - md.Data,
		BssAddr:       md.Bss,
		BssLen:        md.Ebss - md.Bss,
		NoPtrBssAddr:  md.Noptrbss,
		NoPtrBssLen:   md.Enoptrbss - md.Noptrbss,
		TypesAddr:     md.Types,
		TypesLen:      md.Etypes - md.Types,
		TypelinkAddr:  md.Typelinks,
		TypelinkLen:   md.Typelinkslen,
		ITabLinkAddr:  md.Itablinks,
		ITabLinkLen:   md.Itablinkslen,
		FuncTabAddr:   md.Ftab,
		FuncTabLen:    md.Ftablen,
		PCLNTabAddr:   md.Pclntable,
		PCLNTabLen:    md.Pclntablelen,
	}
}

func (md moduledata_1_8_64) pointerOffsets() []int {
	return []int{
		0,
		24,
		48,
		72,
		80,
		88,
		96,
		104,
		112,
		120,
		128,
		136,
		144,
		152,
		160,
		168,
		176,
		184,
		192,
		200,
		208,
		216,
		240,
		264,
		288,
		312,
		328,
	}
}

type moduledata_1_16_32 struct {
	PcHeader                                    uint32
	Funcnametab, Funcnametablen, Funcnametabcap uint32
	Cutab, Cutablen, Cutabcap                   uint32
	Filetab, Filetablen, Filetabcap             uint32
	Pctab, Pctablen, Pctabcap                   uint32
	Pclntable, Pclntablelen, Pclntablecap       uint32
	Ftab, Ftablen, Ftabcap                      uint32
	Findfunctab                                 uint32
	Minpc                                       uint32
	Maxpc                                       uint32
	Text                                        uint32
	Etext                                       uint32
	Noptrdata                                   uint32
	Enoptrdata                                  uint32
	Data                                        uint32
	Edata                                       uint32
	Bss                                         uint32
	Ebss                                        uint32
	Noptrbss                                    uint32
	Enoptrbss                                   uint32
	End                                         uint32
	Gcdata                                      uint32
	Gcbss                                       uint32
	Types                                       uint32
	Etypes                                      uint32
	Textsectmap, Textsectmaplen, Textsectmapcap uint32
	Typelinks, Typelinkslen, Typelinkscap       uint32
	Itablinks, Itablinkslen, Itablinkscap       uint32
	Ptab, Ptablen, Ptabcap                      uint32
	Pluginpath, Pluginpathlen                   uint32
	Pkghashes, Pkghasheslen, Pkghashescap       uint32
}

func (md moduledata_1_16_32) toModuledata() moduledata {
	return moduledata{
		TextAddr:      uint64(md.Text),
		TextLen:       uint64(md.Etext - md.Text),
		NoPtrDataAddr: uint64(md.Noptrdata),
		NoPtrDataLen:  uint64(md.Enoptrdata - md.Noptrdata),
		DataAddr:      uint64(md.Data),
		DataLen:       uint64(md.Edata - md.Data),
		BssAddr:       uint64(md.Bss),
		BssLen:        uint64(md.Ebss - md.Bss),
		NoPtrBssAddr:  uint64(md.Noptrbss),
		NoPtrBssLen:   uint64(md.Enoptrbss - md.Noptrbss),
		TypesAddr:     uint64(md.Types),
		TypesLen:      uint64(md.Etypes - md.Types),
		TypelinkAddr:  uint64(md.Typelinks),
		TypelinkLen:   uint64(md.Typelinkslen),
		ITabLinkAddr:  uint64(md.Itablinks),
		ITabLinkLen:   uint64(md.Itablinkslen),
		FuncTabAddr:   uint64(md.Ftab),
		FuncTabLen:    uint64(md.Ftablen),
		PCLNTabAddr:   uint64(md.Pclntable),
		PCLNTabLen:    uint64(md.Pclntablelen),
	}
}

func (md moduledata_1_16_32) pointerOffsets() []int {
	return []int{
		0,
		4,
		16,
		28,
		40,
		52,
		64,
		76,
		80,
		84,
		88,
		92,
		96,
		100,
		104,
		108,
		112,
		116,
		120,
		124,
		128,
		132,
		136,
		140,
		144,
		148,
		160,
		172,
		184,
		196,
		204,
	}
}

type moduledata_1_16_64 struct {
	PcHeader                                    uint64
	Funcnametab, Funcnametablen, Funcnametabcap uint64
	Cutab, Cutablen, Cutabcap                   uint64
	Filetab, Filetablen, Filetabcap             uint64
	Pctab, Pctablen, Pctabcap                   uint64
	Pclntable, Pclntablelen, Pclntablecap       uint64
	Ftab, Ftablen, Ftabcap                      uint64
	Findfunctab                                 uint64
	Minpc                                       uint64
	Maxpc                                       uint64
	Text                                        uint64
	Etext                                       uint64
	Noptrdata                                   uint64
	Enoptrdata                                  uint64
	Data                                        uint64
	Edata                                       uint64
	Bss                                         uint64
	Ebss                                        uint64
	Noptrbss                                    uint64
	Enoptrbss                                   uint64
	End                                         uint64
	Gcdata                                      uint64
	Gcbss                                       uint64
	Types                                       uint64
	Etypes                                      uint64
	Textsectmap, Textsectmaplen, Textsectmapcap uint64
	Typelinks, Typelinkslen, Typelinkscap       uint64
	Itablinks, Itablinkslen, Itablinkscap       uint64
	Ptab, Ptablen, Ptabcap                      uint64
	Pluginpath, Pluginpathlen                   uint64
	Pkghashes, Pkghasheslen, Pkghashescap       uint64
}

func (md moduledata_1_16_64) toModuledata() moduledata {
	return moduledata{
		TextAddr:      md.Text,
		TextLen:       md.Etext - md.Text,
		NoPtrDataAddr: md.Noptrdata,
		NoPtrDataLen:  md.Enoptrdata - md.Noptrdata,
		DataAddr:      md.Data,
		DataLen:       md.Edata - md.Data,
		BssAddr:       md.Bss,
		BssLen:        md.Ebss - md.Bss,
		NoPtrBssAddr:  md.Noptrbss,
		NoPtrBssLen:   md.Enoptrbss - md.Noptrbss,
		TypesAddr:     md.Types,
		TypesLen:      md.Etypes - md.Types,
		TypelinkAddr:  md.Typelinks,
		TypelinkLen:   md.Typelinkslen,
		ITabLinkAddr:  md.Itablinks,
		ITabLinkLen:   md.Itablinkslen,
		FuncTabAddr:   md.Ftab,
		FuncTabLen:    md.Ftablen,
		PCLNTabAddr:   md.Pclntable,
		PCLNTabLen:    md.Pclntablelen,
	}
}

func (md moduledata_1_16_64) pointerOffsets() []int {
	return []int{
		0,
		8,
		32,
		56,
		80,
		104,
		128,
		152,
		160,
		168,
		176,
		184,
		192,
		200,
		208,
		216,
		224,
		232,
		240,
		248,
		256,
		264,
		272,
		280,
		288,
		296,
		320,
		344,
		368,
		392,
		408,
	}
}

type moduledata_1_18_32 struct {
	PcHeader                                    uint32
	Funcnametab, Funcnametablen, Funcnametabcap uint32
	Cutab, Cutablen, Cutabcap                   uint32
	Filetab, Filetablen, Filetabcap             uint32
	Pctab, Pctablen, Pctabcap                   uint32
	Pclntable, Pclntablelen, Pclntablecap       uint32
	Ftab, Ftablen, Ftabcap                      uint32
	Findfunctab                                 uint32
	Minpc                                       uint32
	Maxpc                                       uint32
	Text                                        uint32
	Etext                                       uint32
	Noptrdata                                   uint32
	Enoptrdata                                  uint32
	Data                                        uint32
	Edata                                       uint32
	Bss                                         uint32
	Ebss                                        uint32
	Noptrbss                                    uint32
	Enoptrbss                                   uint32
	End                                         uint32
	Gcdata                                      uint32
	Gcbss                                       uint32
	Types                                       uint32
	Etypes                                      uint32
	Rodata                                      uint32
	Gofunc                                      uint32
	Textsectmap, Textsectmaplen, Textsectmapcap uint32
	Typelinks, Typelinkslen, Typelinkscap       uint32
	Itablinks, Itablinkslen, Itablinkscap       uint32
	Ptab, Ptablen, Ptabcap                      uint32
	Pluginpath, Pluginpathlen                   uint32
	Pkghashes, Pkghasheslen, Pkghashescap       uint32
}

func (md moduledata_1_18_32) toModuledata() moduledata {
	return moduledata{
		TextAddr:      uint64(md.Text),
		TextLen:       uint64(md.Etext - md.Text),
		NoPtrDataAddr: uint64(md.Noptrdata),
		NoPtrDataLen:  uint64(md.Enoptrdata - md.Noptrdata),
		DataAddr:      uint64(md.Data),
		DataLen:       uint64(md.Edata - md.Data),
		BssAddr:       uint64(md.Bss),
		BssLen:        uint64(md.Ebss - md.Bss),
		NoPtrBssAddr:  uint64(md.Noptrbss),
		NoPtrBssLen:   uint64(md.Enoptrbss - md.Noptrbss),
		TypesAddr:     uint64(md.Types),
		TypesLen:      uint64(md.Etypes - md.Types),
		TypelinkAddr:  uint64(md.Typelinks),
		TypelinkLen:   uint64(md.Typelinkslen),
		ITabLinkAddr:  uint64(md.Itablinks),
		ITabLinkLen:   uint64(md.Itablinkslen),
		FuncTabAddr:   uint64(md.Ftab),
		FuncTabLen:    uint64(md.Ftablen),
		PCLNTabAddr:   uint64(md.Pclntable),
		PCLNTabLen:    uint64(md.Pclntablelen),
		GoFuncVal:     uint64(md.Gofunc),
	}
}

func (md moduledata_1_18_32) pointerOffsets() []int {
	return []int{
		0,
		4,
		16,
		28,
		40,
		52,
		64,
		76,
		80,
		84,
		88,
		92,
		96,
		100,
		104,
		108,
		112,
		116,
		120,
		124,
		128,
		132,
		136,
		140,
		144,
		148,
		152,
		156,
		168,
		180,
		192,
		204,
		212,
	}
}

type moduledata_1_18_64 struct {
	PcHeader                                    uint64
	Funcnametab, Funcnametablen, Funcnametabcap uint64
	Cutab, Cutablen, Cutabcap                   uint64
	Filetab, Filetablen, Filetabcap             uint64
	Pctab, Pctablen, Pctabcap                   uint64
	Pclntable, Pclntablelen, Pclntablecap       uint64
	Ftab, Ftablen, Ftabcap                      uint64
	Findfunctab                                 uint64
	Minpc                                       uint64
	Maxpc                                       uint64
	Text                                        uint64
	Etext                                       uint64
	Noptrdata                                   uint64
	Enoptrdata                                  uint64
	Data                                        uint64
	Edata                                       uint64
	Bss                                         uint64
	Ebss                                        uint64
	Noptrbss                                    uint64
	Enoptrbss                                   uint64
	End                                         uint64
	Gcdata                                      uint64
	Gcbss                                       uint64
	Types                                       uint64
	Etypes                                      uint64
	Rodata                                      uint64
	Gofunc                                      uint64
	Textsectmap, Textsectmaplen, Textsectmapcap uint64
	Typelinks, Typelinkslen, Typelinkscap       uint64
	Itablinks, Itablinkslen, Itablinkscap       uint64
	Ptab, Ptablen, Ptabcap                      uint64
	Pluginpath, Pluginpathlen                   uint64
	Pkghashes, Pkghasheslen, Pkghashescap       uint64
}

func (md moduledata_1_18_64) toModuledata() moduledata {
	return moduledata{
		TextAddr:      md.Text,
		TextLen:       md.Etext - md.Text,
		NoPtrDataAddr: md.Noptrdata,
		NoPtrDataLen:  md.Enoptrdata - md.Noptrdata,
		DataAddr:      md.Data,
		DataLen:       md.Edata - md.Data,
		BssAddr:       md.Bss,
		BssLen:        md.Ebss - md.Bss,
		NoPtrBssAddr:  md.Noptrbss,
		NoPtrBssLen:   md.Enoptrbss - md.Noptrbss,
		TypesAddr:     md.Types,
		TypesLen:      md.Etypes - md.Types,
		TypelinkAddr:  md.Typelinks,
		TypelinkLen:   md.Typelinkslen,
		ITabLinkAddr:  md.Itablinks,
		ITabLinkLen:   md.Itablinkslen,
		FuncTabAddr:   md.Ftab,
		FuncTabLen:    md.Ftablen,
		PCLNTabAddr:   md.Pclntable,
		PCLNTabLen:    md.Pclntablelen,
		GoFuncVal:     md.Gofunc,
	}
}

func (md moduledata_1_18_64) pointerOffsets() []int {
	return []int{
		0,
		8,
		32,
		56,
		80,
		104,
		128,
		152,
		160,
		168,
		176,
		184,
		192,
		200,
		208,
		216,
		224,
		232,
		240,
		248,
		256,
		264,
		272,
		280,
		288,
		296,
		304,
		312,
		336,
		360,
		384,
		408,
		424,
	}
}

type moduledata_1_20_32 struct {
	PcHeader                                    uint32
	Funcnametab, Funcnametablen, Funcnametabcap uint32
	Cutab, Cutablen, Cutabcap                   uint32
	Filetab, Filetablen, Filetabcap             uint32
	Pctab, Pctablen, Pctabcap                   uint32
	Pclntable, Pclntablelen, Pclntablecap       uint32
	Ftab, Ftablen, Ftabcap                      uint32
	Findfunctab                                 uint32
	Minpc                                       uint32
	Maxpc                                       uint32
	Text                                        uint32
	Etext                                       uint32
	Noptrdata                                   uint32
	Enoptrdata                                  uint32
	Data                                        uint32
	Edata                                       uint32
	Bss                                         uint32
	Ebss                                        uint32
	Noptrbss                                    uint32
	Enoptrbss                                   uint32
	Covctrs                                     uint32
	Ecovctrs                                    uint32
	End                                         uint32
	Gcdata                                      uint32
	Gcbss                                       uint32
	Types                                       uint32
	Etypes                                      uint32
	Rodata                                      uint32
	Gofunc                                      uint32
	Textsectmap, Textsectmaplen, Textsectmapcap uint32
	Typelinks, Typelinkslen, Typelinkscap       uint32
	Itablinks, Itablinkslen, Itablinkscap       uint32
	Ptab, Ptablen, Ptabcap                      uint32
	Pluginpath, Pluginpathlen                   uint32
	Pkghashes, Pkghasheslen, Pkghashescap       uint32
}

func (md moduledata_1_20_32) toModuledata() moduledata {
	return moduledata{
		TextAddr:      uint64(md.Text),
		TextLen:       uint64(md.Etext - md.Text),
		NoPtrDataAddr: uint64(md.Noptrdata),
		NoPtrDataLen:  uint64(md.Enoptrdata - md.Noptrdata),
		DataAddr:      uint64(md.Data),
		DataLen:       uint64(md.Edata - md.Data),
		BssAddr:       uint64(md.Bss),
		BssLen:        uint64(md.Ebss - md.Bss),
		NoPtrBssAddr:  uint64(md.Noptrbss),
		NoPtrBssLen:   uint64(md.Enoptrbss - md.Noptrbss),
		TypesAddr:     uint64(md.Types),
		TypesLen:      uint64(md.Etypes - md.Types),
		TypelinkAddr:  uint64(md.Typelinks),
		TypelinkLen:   uint64(md.Typelinkslen),
		ITabLinkAddr:  uint64(md.Itablinks),
		ITabLinkLen:   uint64(md.Itablinkslen),
		FuncTabAddr:   uint64(md.Ftab),
		FuncTabLen:    uint64(md.Ftablen),
		PCLNTabAddr:   uint64(md.Pclntable),
		PCLNTabLen:    uint64(md.Pclntablelen),
		GoFuncVal:     uint64(md.Gofunc),
	}
}

func (md moduledata_1_20_32) pointerOffsets() []int {
	return []int{
		0,
		4,
		16,
		28,
		40,
		52,
		64,
		76,
		80,
		84,
		88,
		92,
		96,
		100,
		104,
		108,
		112,
		116,
		120,
		124,
		128,
		132,
		136,
		140,
		144,
		148,
		152,
		156,
		160,
		164,
		176,
		188,
		200,
		212,
		220,
	}
}

type moduledata_1_20_64 struct {
	PcHeader                                    uint64
	Funcnametab, Funcnametablen, Funcnametabcap uint64
	Cutab, Cutablen, Cutabcap                   uint64
	Filetab, Filetablen, Filetabcap             uint64
	Pctab, Pctablen, Pctabcap                   uint64
	Pclntable, Pclntablelen, Pclntablecap       uint64
	Ftab, Ftablen, Ftabcap                      uint64
	Findfunctab                                 uint64
	Minpc                                       uint64
	Maxpc                                       uint64
	Text                                        uint64
	Etext                                       uint64
	Noptrdata                                   uint64
	Enoptrdata                                  uint64
	Data                                        uint64
	Edata                                       uint64
	Bss                                         uint64
	Ebss                                        uint64
	Noptrbss                                    uint64
	Enoptrbss                                   uint64
	Covctrs                                     uint64
	Ecovctrs                                    uint64
	End                                         uint64
	Gcdata                                      uint64
	Gcbss                                       uint64
	Types                                       uint64
	Etypes                                      uint64
	Rodata                                      uint64
	Gofunc                                      uint64
	Textsectmap, Textsectmaplen, Textsectmapcap uint64
	Typelinks, Typelinkslen, Typelinkscap       uint64
	Itablinks, Itablinkslen, Itablinkscap       uint64
	Ptab, Ptablen, Ptabcap                      uint64
	Pluginpath, Pluginpathlen                   uint64
	Pkghashes, Pkghasheslen, Pkghashescap       uint64
}

func (md moduledata_1_20_64) toModuledata() moduledata {
	return moduledata{
		TextAddr:      md.Text,
		TextLen:       md.Etext - md.Text,
		NoPtrDataAddr: md.Noptrdata,
		NoPtrDataLen:  md.Enoptrdata - md.Noptrdata,
		DataAddr:      md.Data,
		DataLen:       md.Edata - md.Data,
		BssAddr:       md.Bss,
		BssLen:        md.Ebss - md.Bss,
		NoPtrBssAddr:  md.Noptrbss,
		NoPtrBssLen:   md.Enoptrbss - md.Noptrbss,
		TypesAddr:     md.Types,
		TypesLen:      md.Etypes - md.Types,
		TypelinkAddr:  md.Typelinks,
		TypelinkLen:   md.Typelinkslen,
		ITabLinkAddr:  md.Itablinks,
		ITabLinkLen:   md.Itablinkslen,
		FuncTabAddr:   md.Ftab,
		FuncTabLen:    md.Ftablen,
		PCLNTabAddr:   md.Pclntable,
		PCLNTabLen:    md.Pclntablelen,
		GoFuncVal:     md.Gofunc,
	}
}

func (md moduledata_1_20_64) pointerOffsets() []int {
	return []int{
		0,
		8,
		32,
		56,
		80,
		104,
		128,
		152,
		160,
		168,
		176,
		184,
		192,
		200,
		208,
		216,
		224,
		232,
		240,
		248,
		256,
		264,
		272,
		280,
		288,
		296,
		304,
		312,
		320,
		328,
		352,
		376,
		400,
		424,
		440,
	}
}

type moduledata_1_21_32 struct {
	PcHeader                                    uint32
	Funcnametab, Funcnametablen, Funcnametabcap uint32
	Cutab, Cutablen, Cutabcap                   uint32
	Filetab, Filetablen, Filetabcap             uint32
	Pctab, Pctablen, Pctabcap                   uint32
	Pclntable, Pclntablelen, Pclntablecap       uint32
	Ftab, Ftablen, Ftabcap                      uint32
	Findfunctab                                 uint32
	Minpc                                       uint32
	Maxpc                                       uint32
	Text                                        uint32
	Etext                                       uint32
	Noptrdata                                   uint32
	Enoptrdata                                  uint32
	Data                                        uint32
	Edata                                       uint32
	Bss                                         uint32
	Ebss                                        uint32
	Noptrbss                                    uint32
	Enoptrbss                                   uint32
	Covctrs                                     uint32
	Ecovctrs                                    uint32
	End                                         uint32
	Gcdata                                      uint32
	Gcbss                                       uint32
	Types                                       uint32
	Etypes                                      uint32
	Rodata                                      uint32
	Gofunc                                      uint32
	Textsectmap, Textsectmaplen, Textsectmapcap uint32
	Typelinks, Typelinkslen, Typelinkscap       uint32
	Itablinks, Itablinkslen, Itablinkscap       uint32
	Ptab, Ptablen, Ptabcap                      uint32
	Pluginpath, Pluginpathlen                   uint32
	Pkghashes, Pkghasheslen, Pkghashescap       uint32
	Inittasks, Inittaskslen, Inittaskscap       uint32
}

func (md moduledata_1_21_32) toModuledata() moduledata {
	return moduledata{
		TextAddr:      uint64(md.Text),
		TextLen:       uint64(md.Etext - md.Text),
		NoPtrDataAddr: uint64(md.Noptrdata),
		NoPtrDataLen:  uint64(md.Enoptrdata - md.Noptrdata),
		DataAddr:      uint64(md.Data),
		DataLen:       uint64(md.Edata - md.Data),
		BssAddr:       uint64(md.Bss),
		BssLen:        uint64(md.Ebss - md.Bss),
		NoPtrBssAddr:  uint64(md.Noptrbss),
		NoPtrBssLen:   uint64(md.Enoptrbss - md.Noptrbss),
		TypesAddr:     uint64(md.Types),
		TypesLen:      uint64(md.Etypes - md.Types),
		TypelinkAddr:  uint64(md.Typelinks),
		TypelinkLen:   uint64(md.Typelinkslen),
		ITabLinkAddr:  uint64(md.Itablinks),
		ITabLinkLen:   uint64(md.Itablinkslen),
		FuncTabAddr:   uint64(md.Ftab),
		FuncTabLen:    uint64(md.Ftablen),
		PCLNTabAddr:   uint64(md.Pclntable),
		PCLNTabLen:    uint64(md.Pclntablelen),
		GoFuncVal:     uint64(md.Gofunc),
	}
}

func (md moduledata_1_21_32) pointerOffsets() []int {
	return []int{
		0,
		4,
		16,
		28,
		40,
		52,
		64,
		76,
		80,
		84,
		88,
		92,
		96,
		100,
		104,
		108,
		112,
		116,
		120,
		124,
		128,
		132,
		136,
		140,
		144,
		148,
		152,
		156,
		160,
		164,
		176,
		188,
		200,
		212,
		220,
		232,
	}
}

type moduledata_1_21_64 struct {
	PcHeader                                    uint64
	Funcnametab, Funcnametablen, Funcnametabcap uint64
	Cutab, Cutablen, Cutabcap                   uint64
	Filetab, Filetablen, Filetabcap             uint64
	Pctab, Pctablen, Pctabcap                   uint64
	Pclntable, Pclntablelen, Pclntablecap       uint64
	Ftab, Ftablen, Ftabcap                      uint64
	Findfunctab                                 uint64
	Minpc                                       uint64
	Maxpc                                       uint64
	Text                                        uint64
	Etext                                       uint64
	Noptrdata                                   uint64
	Enoptrdata                                  uint64
	Data                                        uint64
	Edata                                       uint64
	Bss                                         uint64
	Ebss                                        uint64
	Noptrbss                                    uint64
	Enoptrbss                                   uint64
	Covctrs                                     uint64
	Ecovctrs                                    uint64
	End                                         uint64
	Gcdata                                      uint64
	Gcbss                                       uint64
	Types                                       uint64
	Etypes                                      uint64
	Rodata                                      uint64
	Gofunc                                      uint64
	Textsectmap, Textsectmaplen, Textsectmapcap uint64
	Typelinks, Typelinkslen, Typelinkscap       uint64
	Itablinks, Itablinkslen, Itablinkscap       uint64
	Ptab, Ptablen, Ptabcap                      uint64
	Pluginpath, Pluginpathlen                   uint64
	Pkghashes, Pkghasheslen, Pkghashescap       uint64
	Inittasks, Inittaskslen, Inittaskscap       uint64
}

func (md moduledata_1_21_64) toModuledata() moduledata {
	return moduledata{
		TextAddr:      md.Text,
		TextLen:       md.Etext - md.Text,
		NoPtrDataAddr: md.Noptrdata,
		NoPtrDataLen:  md.Enoptrdata - md.Noptrdata,
		DataAddr:      md.Data,
		DataLen:       md.Edata - md.Data,
		BssAddr:       md.Bss,
		BssLen:        md.Ebss - md.Bss,
		NoPtrBssAddr:  md.Noptrbss,
		NoPtrBssLen:   md.Enoptrbss - md.Noptrbss,
		TypesAddr:     md.Types,
		TypesLen:      md.Etypes - md.Types,
		TypelinkAddr:  md.Typelinks,
		TypelinkLen:   md.Typelinkslen,
		ITabLinkAddr:  md.Itablinks,
		ITabLinkLen:   md.Itablinkslen,
		FuncTabAddr:   md.Ftab,
		FuncTabLen:    md.Ftablen,
		PCLNTabAddr:   md.Pclntable,
		PCLNTabLen:    md.Pclntablelen,
		GoFuncVal:     md.Gofunc,
	}
}

func (md moduledata_1_21_64) pointerOffsets() []int {
	return []int{
		0,
		8,
		32,
		56,
		80,
		104,
		128,
		152,
		160,
		168,
		176,
		184,
		192,
		200,
		208,
		216,
		224,
		232,
		240,
		248,
		256,
		264,
		272,
		280,
		288,
		296,
		304,
		312,
		320,
		328,
		352,
		376,
		400,
		424,
		440,
		464,
	}
}

type moduledata_1_26_32 struct {
	PcHeader                                    uint32
	Funcnametab, Funcnametablen, Funcnametabcap uint32
	Cutab, Cutablen, Cutabcap                   uint32
	Filetab, Filetablen, Filetabcap             uint32
	Pctab, Pctablen, Pctabcap                   uint32
	Pclntable, Pclntablelen, Pclntablecap       uint32
	Ftab, Ftablen, Ftabcap                      uint32
	Findfunctab                                 uint32
	Minpc                                       uint32
	Maxpc                                       uint32
	Text                                        uint32
	Etext                                       uint32
	Noptrdata                                   uint32
	Enoptrdata                                  uint32
	Data                                        uint32
	Edata                                       uint32
	Bss                                         uint32
	Ebss                                        uint32
	Noptrbss                                    uint32
	Enoptrbss                                   uint32
	Covctrs                                     uint32
	Ecovctrs                                    uint32
	End                                         uint32
	Gcdata                                      uint32
	Gcbss                                       uint32
	Types                                       uint32
	Etypes                                      uint32
	Rodata                                      uint32
	Gofunc                                      uint32
	Epclntab                                    uint32
	Textsectmap, Textsectmaplen, Textsectmapcap uint32
	Typelinks, Typelinkslen, Typelinkscap       uint32
	Itablinks, Itablinkslen, Itablinkscap       uint32
	Ptab, Ptablen, Ptabcap                      uint32
	Pluginpath, Pluginpathlen                   uint32
	Pkghashes, Pkghasheslen, Pkghashescap       uint32
	Inittasks, Inittaskslen, Inittaskscap       uint32
}

func (md moduledata_1_26_32) toModuledata() moduledata {
	return moduledata{
		TextAddr:      uint64(md.Text),
		TextLen:       uint64(md.Etext - md.Text),
		NoPtrDataAddr: uint64(md.Noptrdata),
		NoPtrDataLen:  uint64(md.Enoptrdata - md.Noptrdata),
		DataAddr:      uint64(md.Data),
		DataLen:       uint64(md.Edata - md.Data),
		BssAddr:       uint64(md.Bss),
		BssLen:        uint64(md.Ebss - md.Bss),
		NoPtrBssAddr:  uint64(md.Noptrbss),
		NoPtrBssLen:   uint64(md.Enoptrbss - md.Noptrbss),
		TypesAddr:     uint64(md.Types),
		TypesLen:      uint64(md.Etypes - md.Types),
		TypelinkAddr:  uint64(md.Typelinks),
		TypelinkLen:   uint64(md.Typelinkslen),
		ITabLinkAddr:  uint64(md.Itablinks),
		ITabLinkLen:   uint64(md.Itablinkslen),
		FuncTabAddr:   uint64(md.Ftab),
		FuncTabLen:    uint64(md.Ftablen),
		PCLNTabAddr:   uint64(md.Pclntable),
		PCLNTabLen:    uint64(md.Pclntablelen),
		GoFuncVal:     uint64(md.Gofunc),
	}
}

func (md moduledata_1_26_32) pointerOffsets() []int {
	return []int{
		0,
		4,
		16,
		28,
		40,
		52,
		64,
		76,
		80,
		84,
		88,
		92,
		96,
		100,
		104,
		108,
		112,
		116,
		120,
		124,
		128,
		132,
		136,
		140,
		144,
		148,
		152,
		156,
		160,
		164,
		168,
		180,
		192,
		204,
		216,
		224,
		236,
	}
}

type moduledata_1_26_64 struct {
	PcHeader                                    uint64
	Funcnametab, Funcnametablen, Funcnametabcap uint64
	Cutab, Cutablen, Cutabcap                   uint64
	Filetab, Filetablen, Filetabcap             uint64
	Pctab, Pctablen, Pctabcap                   uint64
	Pclntable, Pclntablelen, Pclntablecap       uint64
	Ftab, Ftablen, Ftabcap                      uint64
	Findfunctab                                 uint64
	Minpc                                       uint64
	Maxpc                                       uint64
	Text                                        uint64
	Etext                                       uint64
	Noptrdata                                   uint64
	Enoptrdata                                  uint64
	Data                                        uint64
	Edata                                       uint64
	Bss                                         uint64
	Ebss                                        uint64
	Noptrbss                                    uint64
	Enoptrbss                                   uint64
	Covctrs                                     uint64
	Ecovctrs                                    uint64
	End                                         uint64
	Gcdata                                      uint64
	Gcbss                                       uint64
	Types                                       uint64
	Etypes                                      uint64
	Rodata                                      uint64
	Gofunc                                      uint64
	Epclntab                                    uint64
	Textsectmap, Textsectmaplen, Textsectmapcap uint64
	Typelinks, Typelinkslen, Typelinkscap       uint64
	Itablinks, Itablinkslen, Itablinkscap       uint64
	Ptab, Ptablen, Ptabcap                      uint64
	Pluginpath, Pluginpathlen                   uint64
	Pkghashes, Pkghasheslen, Pkghashescap       uint64
	Inittasks, Inittaskslen, Inittaskscap       uint64
}

func (md moduledata_1_26_64) toModuledata() moduledata {
	return moduledata{
		TextAddr:      md.Text,
		TextLen:       md.Etext - md.Text,
		NoPtrDataAddr: md.Noptrdata,
		NoPtrDataLen:  md.Enoptrdata - md.Noptrdata,
		DataAddr:      md.Data,
		DataLen:       md.Edata - md.Data,
		BssAddr:       md.Bss,
		BssLen:        md.Ebss - md.Bss,
		NoPtrBssAddr:  md.Noptrbss,
		NoPtrBssLen:   md.Enoptrbss - md.Noptrbss,
		TypesAddr:     md.Types,
		TypesLen:      md.Etypes - md.Types,
		TypelinkAddr:  md.Typelinks,
		TypelinkLen:   md.Typelinkslen,
		ITabLinkAddr:  md.Itablinks,
		ITabLinkLen:   md.Itablinkslen,
		FuncTabAddr:   md.Ftab,
		FuncTabLen:    md.Ftablen,
		PCLNTabAddr:   md.Pclntable,
		PCLNTabLen:    md.Pclntablelen,
		GoFuncVal:     md.Gofunc,
	}
}

func (md moduledata_1_26_64) pointerOffsets() []int {
	return []int{
		0,
		8,
		32,
		56,
		80,
		104,
		128,
		152,
		160,
		168,
		176,
		184,
		192,
		200,
		208,
		216,
		224,
		232,
		240,
		248,
		256,
		264,
		272,
		280,
		288,
		296,
		304,
		312,
		320,
		328,
		336,
		360,
		384,
		408,
		432,
		448,
		472,
	}
}

var moduledataVersions_32 = []struct {
	minVersion int
	factory    func() modulable
}{
	{5, func() modulable { return &moduledata_1_5_32{} }},
	{7, func() modulable { return &moduledata_1_7_32{} }},
	{8, func() modulable { return &moduledata_1_8_32{} }},
	{16, func() modulable { return &moduledata_1_16_32{} }},
	{18, func() modulable { return &moduledata_1_18_32{} }},
	{20, func() modulable { return &moduledata_1_20_32{} }},
	{21, func() modulable { return &moduledata_1_21_32{} }},
	{26, func() modulable { return &moduledata_1_26_32{} }},
}

var moduledataVersions_64 = []struct {
	minVersion int
	factory    func() modulable
}{
	{5, func() modulable { return &moduledata_1_5_64{} }},
	{7, func() modulable { return &moduledata_1_7_64{} }},
	{8, func() modulable { return &moduledata_1_8_64{} }},
	{16, func() modulable { return &moduledata_1_16_64{} }},
	{18, func() modulable { return &moduledata_1_18_64{} }},
	{20, func() modulable { return &moduledata_1_20_64{} }},
	{21, func() modulable { return &moduledata_1_21_64{} }},
	{26, func() modulable { return &moduledata_1_26_64{} }},
}

func selectModuleData(v int, bits int) (modulable, error) {
	var versions []struct {
		minVersion int
		factory    func() modulable
	}
	switch bits {
	case 32:
		versions = moduledataVersions_32
	case 64:
		versions = moduledataVersions_64
	default:
		return nil, fmt.Errorf("unsupported bits %d", bits)
	}
	idx := sort.Search(len(versions), func(i int) bool { return versions[i].minVersion > v }) - 1
	if idx < 0 {
		return nil, fmt.Errorf("unsupported version %d and bits %d", v, bits)
	}
	return versions[idx].factory(), nil
}
