package flags

import (
	"fmt"

	buildv1alpha1 "github.com/shipwright-io/build/pkg/apis/build/v1alpha1"
)

// StrategyKindValue implements pflag.Value interface, to represent Shipwright's BuildStrategyKind as
// a string command-line in an cobra.Command instance.
type StrategyKindValue struct {
	kindPtr *buildv1alpha1.BuildStrategyKind
}

// String shows the value as string.
func (s *StrategyKindValue) String() string {
	if s.kindPtr == nil {
		return ""
	}
	return string(*s.kindPtr)
}

// Set set the informed string as BuildStrategyKind by casting.
func (s *StrategyKindValue) Set(str string) error {
	var strInterface interface{} = str
	var ok bool
	s.kindPtr, ok = strInterface.(*buildv1alpha1.BuildStrategyKind)
	if !ok {
		return fmt.Errorf("unable to cast '%s' as BuildStrategyKind value", str)
	}
	return nil
}

// Type analogous to the pflag "string".
func (s *StrategyKindValue) Type() string {
	return "string"
}

// NewStrategyKindValue creates a new instance of StrategyKindValue sharing an existing reference.
func NewStrategyKindValue(kindPtr *buildv1alpha1.BuildStrategyKind) *StrategyKindValue {
	return &StrategyKindValue{kindPtr: kindPtr}
}
