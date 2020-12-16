package diff

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/google/go-cmp/cmp"
	"go.uber.org/zap"
)

type Diffeable interface {
	Diff(right interface{}, options ...Option)
}

type Option interface {
	apply(o *options)
}

type optionFunc func(o *options)

func (f optionFunc) apply(opts *options) {
	f(opts)
}

func CmpOption(cmpOption cmp.Option) Option {
	return optionFunc(func(opts *options) { opts.cmpOptions = append(opts.cmpOptions, cmpOption) })
}

func OnEvent(callback func(Event)) Option {
	return optionFunc(func(opts *options) { opts.onEvent = callback })
}

type options struct {
	cmpOptions []cmp.Option
	onEvent    func(Event)
}

type Kind uint8

const (
	KindAdded Kind = iota
	KindChanged
	KindRemoved
)

func (k Kind) String() string {
	switch k {
	case KindAdded:
		return "added"
	case KindChanged:
		return "changed"
	case KindRemoved:
		return "removed"
	}

	return "unknown"
}

type Path cmp.Path

func (pa Path) String() string {
	if len(pa) == 1 {
		return ""
	}

	return strings.TrimPrefix(cmp.Path(pa[1:]).GoString(), ".")
}

type Event struct {
	Path Path
	Kind Kind
	Old  reflect.Value
	New  reflect.Value
}

// Match currently simply ensure that `pattern` parameter is the start of the path string
// which represents the direct access from top-level to struct.
func (p *Event) Match(pattern string) (match bool, matches []string) {
	regexRaw := regexp.QuoteMeta(pattern)
	regexRaw = strings.ReplaceAll("^"+regexRaw+"$", "#", `([0-9]+|.->[0-9]+|[0-9]+->.|[0-9]+->[0-9]+)`)

	regex := regexp.MustCompile(regexRaw)
	regexMatch := regex.FindAllStringSubmatch(p.Path.String(), 1)
	if len(regexMatch) != 1 {
		return false, nil
	}

	// For now we accept only array indices, will need to re-write logic if we ever need to check for keys also
	subMatches := regexMatch[0][1:]
	if len(subMatches) == 0 {
		return true, nil
	}

	return true, subMatches
}

func (p *Event) AddedKind() bool {
	return p.Kind == KindAdded
}

func (p *Event) ChangedKind() bool {
	return p.Kind == KindChanged
}

func (p *Event) RemovedKind() bool {
	return p.Kind == KindRemoved
}

// Element picks the element based on the Event's Kind, if it's removed, the element is the
// "old" value, if it's added or changed, the element is the "new" value.
func (p *Event) Element() reflect.Value {
	if p.Kind == KindRemoved {
		return p.Old
	}

	return p.New
}

func (p *Event) String() string {
	path := ""
	if len(p.Path) > 1 {
		path = " @ " + p.Path.String()
	}

	return fmt.Sprintf("%s => %s (%s%s)", reflectValueToString(p.Old), reflectValueToString(p.New), p.Kind, path)
}

func reflectValueToString(value reflect.Value) string {
	if !value.IsValid() {
		return "<nil>"
	}

	if value.CanInterface() {
		if reflectValueCanIsNil(value) && value.IsNil() {
			return fmt.Sprintf("<nil> (%s)", value.Type())
		}

		v := value.Interface()
		return fmt.Sprintf("%v (%T)", v, v)
	}

	return fmt.Sprintf("<type %T>", value.Type())
}

func reflectValueCanIsNil(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.UnsafePointer, reflect.Interface, reflect.Slice:
		return true
	default:
		return false
	}
}

func Diff(left interface{}, right interface{}, opts ...Option) {
	options := options{}
	for _, opt := range opts {
		opt.apply(&options)
	}

	if options.onEvent == nil {
		panic("the option diff.OnEvent(...) must always be defined")
	}

	reporter := &diffReporter{notify: options.onEvent}
	cmp.Equal(left, right, append(
		[]cmp.Option{cmp.Reporter(reporter)},
		options.cmpOptions...,
	)...)
}

type diffReporter struct {
	notify func(event Event)
	path   cmp.Path
	diffs  []string
}

func (r *diffReporter) PushStep(ps cmp.PathStep) {
	if traceEnabled {
		zlog.Debug("pushing path step", zap.Stringer("step", ps))
	}

	r.path = append(r.path, ps)
}

func (r *diffReporter) Report(rs cmp.Result) {
	if !rs.Equal() {
		lastStep := r.path.Last()
		vLeft, vRight := lastStep.Values()
		if !vLeft.IsValid() {
			if traceEnabled {
				zlog.Debug("added event", zap.Stringer("path", r.path))
			}

			// Left is not set but right is, we have added "right"
			r.notify(Event{Kind: KindAdded, Path: Path(r.path), New: vRight})
			return
		}

		if !vRight.IsValid() {
			if traceEnabled {
				zlog.Debug("removed event", zap.Stringer("path", r.path))
			}

			// Left is set but right is not, we have removed "left"
			r.notify(Event{Kind: KindRemoved, Path: Path(r.path), Old: vLeft})
			return
		}

		if isArrayPathStep(lastStep) {
			// We might want to do this only on certain circumstances?
			if traceEnabled {
				zlog.Debug("array changed event, splitting in removed, added", zap.Stringer("path", r.path))
			}

			r.notify(Event{Kind: KindRemoved, Path: Path(r.path), Old: vLeft})
			r.notify(Event{Kind: KindAdded, Path: Path(r.path), New: vRight})
			return
		}

		if traceEnabled {
			zlog.Debug("changed event", zap.Stringer("path", r.path))
		}

		r.notify(Event{Kind: KindChanged, Path: Path(r.path), Old: vLeft, New: vRight})
	}
}

func (r *diffReporter) PopStep() {
	if traceEnabled {
		zlog.Debug("popping path step", zap.Stringer("step", r.path[len(r.path)-1]))
	}

	r.path = r.path[:len(r.path)-1]
}

func isArrayPathStep(step cmp.PathStep) bool {
	_, ok := step.(cmp.SliceIndex)
	return ok
}

func copyPath(path cmp.Path) Path {
	if len(path) == 0 {
		return Path(path)
	}

	out := make([]cmp.PathStep, len(path))
	for i, step := range path {
		out[i] = step
	}

	return Path(cmp.Path(out))
}
