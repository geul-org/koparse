//ff:type feature=config type=model
//ff:what Split — 분리 판정 설정
package config

// Split configures the eojeol split judgement (§6.3).
type Split struct {
	// QStem is the Qualifier of the stem-support warrant of split candidates.
	QStem float64 `json:"q_stem"`
	// QWhole is the Qualifier of the whole-eojeol warrant.
	QWhole float64 `json:"q_whole"`
	// WholePrior is the whole warrant's verdict when the eojeol is not a known label.
	WholePrior float64 `json:"whole_prior"`
	// KnownDecay lowers a known stem's verdict by KnownDecay·ln(candidate count).
	KnownDecay float64 `json:"known_decay"`
	// KnownMin is the lowest verdict of a known stem.
	KnownMin float64 `json:"known_min"`
	// UnknownScale scales the suffix-chain weight of a stem that is not a known label.
	UnknownScale float64 `json:"unknown_scale"`
	// ForeignStem is the verdict of a stem that is exactly a non-Hangul chunk (GPU가).
	ForeignStem float64 `json:"foreign_stem"`
	// Attack forces of the rebuttals and the exception.
	ForceWholeKnown float64 `json:"force_whole_known"`
	ForceBatchim    float64 `json:"force_batchim"`
	ForceChain      float64 `json:"force_chain"`
	ForceShortStem  float64 `json:"force_short_stem"`
	ForceShorter    float64 `json:"force_shorter_known"`
	ForceQuoted     float64 `json:"force_quoted"`
	// Adopt is the verdict the first candidate must reach; Margin the gap to the second.
	Adopt  float64 `json:"adopt"`
	Margin float64 `json:"margin"`
}
