package assert

import (
	"context"

	cetest "github.com/cloudevents/sdk-go/v2/test"

	"knative.dev/reconciler-test/pkg/eventshub"
	"knative.dev/reconciler-test/pkg/feature"
)

type MatchAssertionBuilder struct {
	storeName string
	matchers  []eventshub.EventInfoMatcher
}

// OnStore creates an assertion builder starting from the name of the store
func OnStore(name string) MatchAssertionBuilder {
	return MatchAssertionBuilder{
		storeName: name,
		matchers:  nil,
	}
}

// Match adds the provided matchers in this builder
func (m MatchAssertionBuilder) Match(matchers ...eventshub.EventInfoMatcher) MatchAssertionBuilder {
	m.matchers = append(m.matchers, matchers...)
	return m
}

// MatchReceivedEvent is a shortcut for Match(MatchKind(eventshub.EventReceived), MatchEvent(matchers...))
func (m MatchAssertionBuilder) MatchReceivedEvent(matchers ...cetest.EventMatcher) MatchAssertionBuilder {
	m.matchers = append(m.matchers, MatchKind(eventshub.EventReceived))
	m.matchers = append(m.matchers, MatchEvent(matchers...))
	return m
}

// MatchRejectedEvent is a shortcut for Match(MatchKind(eventshub.EventRejected), MatchEvent(matchers...))
func (m MatchAssertionBuilder) MatchRejectedEvent(matchers ...cetest.EventMatcher) MatchAssertionBuilder {
	m.matchers = append(m.matchers, MatchKind(eventshub.EventRejected))
	m.matchers = append(m.matchers, MatchEvent(matchers...))
	return m
}

// MatchSentEvent is a shortcut for Match(MatchKind(eventshub.EventSent), MatchEvent(matchers...))
func (m MatchAssertionBuilder) MatchSentEvent(matchers ...cetest.EventMatcher) MatchAssertionBuilder {
	m.matchers = append(m.matchers, MatchKind(eventshub.EventSent))
	m.matchers = append(m.matchers, MatchEvent(matchers...))
	return m
}

// MatchResponseEvent is a shortcut for Match(MatchKind(eventshub.EventResponse), MatchEvent(matchers...))
func (m MatchAssertionBuilder) MatchResponseEvent(matchers ...cetest.EventMatcher) MatchAssertionBuilder {
	m.matchers = append(m.matchers, MatchKind(eventshub.EventResponse))
	m.matchers = append(m.matchers, MatchEvent(matchers...))
	return m
}

// MatchEvent is a shortcut for Match(MatchEvent(), OneOf(MatchKind(eventshub.EventReceived), MatchKind(eventshub.EventSent)))
func (m MatchAssertionBuilder) MatchEvent(matchers ...cetest.EventMatcher) MatchAssertionBuilder {
	m.matchers = append(m.matchers, OneOf(
		MatchKind(eventshub.EventReceived),
		MatchKind(eventshub.EventSent),
	))
	m.matchers = append(m.matchers, MatchEvent(matchers...))
	return m
}

// AtLeast builds the assertion feature.StepFn
// OnStore(store).Match(matchers).AtLeast(min) is equivalent to StoreFromContext(ctx, store).AssertAtLeast(min, matchers)
func (m MatchAssertionBuilder) AtLeast(min int) feature.StepFn {
	return func(ctx context.Context, t feature.T) {
		eventshub.StoreFromContext(ctx, m.storeName).AssertAtLeast(t, min, m.matchers...)
	}
}

// InRange builds the assertion feature.StepFn
// OnStore(store).Match(matchers).InRange(min, max) is equivalent to StoreFromContext(ctx, store).AssertInRange(min, max, matchers)
func (m MatchAssertionBuilder) InRange(min int, max int) feature.StepFn {
	return func(ctx context.Context, t feature.T) {
		eventshub.StoreFromContext(ctx, m.storeName).AssertInRange(t, min, max, m.matchers...)
	}
}

// Exact builds the assertion feature.StepFn
// OnStore(store).Match(matchers).Exact(n) is equivalent to StoreFromContext(ctx, store).AssertExact(n, matchers)
func (m MatchAssertionBuilder) Exact(n int) feature.StepFn {
	return func(ctx context.Context, t feature.T) {
		eventshub.StoreFromContext(ctx, m.storeName).AssertExact(t, n, m.matchers...)
	}
}

// Not builds the assertion feature.StepFn
// OnStore(store).Match(matchers).Not() is equivalent to StoreFromContext(ctx, store).AssertNot(matchers)
func (m MatchAssertionBuilder) Not() feature.StepFn {
	return func(ctx context.Context, t feature.T) {
		eventshub.StoreFromContext(ctx, m.storeName).AssertNot(t, m.matchers...)
	}
}
