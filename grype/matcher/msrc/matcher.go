package msrc

import (
	"github.com/anchore/grype/grype/match"
	"github.com/anchore/grype/grype/matcher/common"
	"github.com/anchore/grype/grype/pkg"
	"github.com/anchore/grype/grype/vulnerability"
	"github.com/anchore/syft/syft/distro"
	syftPkg "github.com/anchore/syft/syft/pkg"
)

type Matcher struct {
}

func (m *Matcher) PackageTypes() []syftPkg.Type {
	// This looks like there is a special package, but in reality, this is just
	// a workaround. MSRC matching is done at the KB-patch level, and so this
	// treats KBs as "packages" but they aren't packages, they are patches
	return []syftPkg.Type{syftPkg.KbPkg}
}

func (m *Matcher) Type() match.MatcherType {
	return match.MsrcMatcher
}

func (m *Matcher) Match(store vulnerability.Provider, d *distro.Distro, p pkg.Package) ([]match.Match, error) {
	var matches []match.Match

	// find KB matches for the MSFT version given in the package and version.
	// The "distro" holds the information about the Windows version, and its
	// patch (KB)
	kbMatches, err := common.FindMatchesByPackageDistro(store, d, p, m.Type())
	if err != nil {
		return nil, err
	}

	matches = append(matches, kbMatches...)

	return matches, nil
}
