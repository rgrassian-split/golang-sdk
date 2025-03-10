package openfeature

import (
	"context"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

// The `evaluation context` structure MUST define an optional `targeting key` field of type string,
// identifying the subject of the flag evaluation.
func TestRequirement_3_1_1(t *testing.T) {
	evalCtx := &EvaluationContext{}

	metaValue := reflect.ValueOf(evalCtx).Elem()
	fieldName := "TargetingKey"

	field := metaValue.FieldByName(fieldName)
	if field == (reflect.Value{}) {
		t.Errorf("field %s doesn't exist in the EvaluationContext struct", fieldName)
	}
}

// The evaluation context MUST support the inclusion of custom fields,
// having keys of type `string`, and values of type `boolean | string | number | datetime | structure`.
func TestRequirement_3_1_2(t *testing.T) {
	evalCtx := EvaluationContext{}

	tpe := reflect.TypeOf(evalCtx.Attributes)
	if tpe.Kind() != reflect.Map {
		t.Fatalf("expected EvaluationContext.Attributes kind to be map, got %s", tpe.Kind())
	}
	if tpe.Key().Kind() != reflect.String {
		t.Errorf("expected EvaluationContext.Attributes key to be string, got %s", tpe.Key().Kind())
	}
	if tpe.Elem().Kind() != reflect.Interface {
		t.Errorf("expected EvaluationContext.Attributes value to be interface{}, got %s", tpe.Elem().Kind())
	}
}

// The API, Client and invocation MUST have a method for supplying `evaluation context`.
func TestRequirement_3_2_1(t *testing.T) {
	defer t.Cleanup(initSingleton)

	t.Run("API MUST have a method for supplying `evaluation context`", func(t *testing.T) {
		SetEvaluationContext(EvaluationContext{})
	})

	t.Run("client MUST have a method for supplying `evaluation context`", func(t *testing.T) {
		client := NewClient("test")

		type requirement interface {
			SetEvaluationContext(evalCtx EvaluationContext)
		}

		var clientI interface{} = client
		if _, ok := clientI.(requirement); !ok {
			t.Error("client doesn't implement the required SetEvaluationContext func signature")
		}
	})

	t.Run("invocation MUST have a method for supplying `evaluation context`", func(t *testing.T) {
		client := NewClient("test")

		type requirement interface {
			BooleanValue(ctx context.Context, flag string, defaultValue bool, evalCtx EvaluationContext, options ...Option) (bool, error)
			StringValue(ctx context.Context, flag string, defaultValue string, evalCtx EvaluationContext, options ...Option) (string, error)
			FloatValue(ctx context.Context, flag string, defaultValue float64, evalCtx EvaluationContext, options ...Option) (float64, error)
			IntValue(ctx context.Context, flag string, defaultValue int64, evalCtx EvaluationContext, options ...Option) (int64, error)
			ObjectValue(ctx context.Context, flag string, defaultValue interface{}, evalCtx EvaluationContext, options ...Option) (interface{}, error)
			BooleanValueDetails(ctx context.Context, flag string, defaultValue bool, evalCtx EvaluationContext, options ...Option) (BooleanEvaluationDetails, error)
			StringValueDetails(ctx context.Context, flag string, defaultValue string, evalCtx EvaluationContext, options ...Option) (StringEvaluationDetails, error)
			FloatValueDetails(ctx context.Context, flag string, defaultValue float64, evalCtx EvaluationContext, options ...Option) (FloatEvaluationDetails, error)
			IntValueDetails(ctx context.Context, flag string, defaultValue int64, evalCtx EvaluationContext, options ...Option) (IntEvaluationDetails, error)
			ObjectValueDetails(ctx context.Context, flag string, defaultValue interface{}, evalCtx EvaluationContext, options ...Option) (InterfaceEvaluationDetails, error)
		}

		var clientI interface{} = client
		if _, ok := clientI.(requirement); !ok {
			t.Error("client doesn't implement the required func signatures containing EvaluationContext")
		}
	})
}

// Evaluation context MUST be merged in the order: API (global) - client - invocation,
// with duplicate values being overwritten.
func TestRequirement_3_2_2(t *testing.T) {
	defer t.Cleanup(initSingleton)
	ctrl := gomock.NewController(t)

	apiEvalCtx := EvaluationContext{
		TargetingKey: "API",
		Attributes: map[string]interface{}{
			"invocationEvalCtx": true,
			"foo":               2,
			"user":              2,
		},
	}
	SetEvaluationContext(apiEvalCtx)

	mockProvider := NewMockFeatureProvider(ctrl)
	mockProvider.EXPECT().Metadata().AnyTimes()
	SetProvider(mockProvider)

	client := NewClient("test")
	clientEvalCtx := EvaluationContext{
		TargetingKey: "Client",
		Attributes: map[string]interface{}{
			"clientEvalCtx": true,
			"foo":           1,
			"user":          1,
		},
	}
	client.SetEvaluationContext(clientEvalCtx)

	invocationEvalCtx := EvaluationContext{
		TargetingKey: "",
		Attributes: map[string]interface{}{
			"apiEvalCtx": true,
			"foo":        "bar",
		},
	}

	mockProvider.EXPECT().Hooks().AnyTimes()
	expectedMergedEvalCtx := EvaluationContext{
		TargetingKey: "Client",
		Attributes: map[string]interface{}{
			"apiEvalCtx":        true,
			"invocationEvalCtx": true,
			"clientEvalCtx":     true,
			"foo":               "bar",
			"user":              1,
		},
	}
	flatCtx := flattenContext(expectedMergedEvalCtx)
	mockProvider.EXPECT().StringEvaluation(gomock.Any(), gomock.Any(), gomock.Any(), flatCtx)

	_, err := client.StringValue(context.Background(), "foo", "bar", invocationEvalCtx)
	if err != nil {
		t.Error(err)
	}

}
