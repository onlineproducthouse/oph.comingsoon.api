package errorlocal_test

import (
	"errors"
	"testing"

	"oph.api.comingsoon/oph.api.comingsoon.testhelpers/utils/asserterrmsg"
	"oph.api.comingsoon/oph.api.comingsoon.util/errorlocal"
	"oph.api.comingsoon/oph.api.comingsoon.util/httpstatuslocal"
)

func TestValidationErr(t *testing.T) {
	op := "TestValidationErr"
	msg := "Testing Validation Error"
	innerMsg := "Inner Testing Validation Error"
	originalErr := errors.New(innerMsg)

	err := errorlocal.ValidationErr(msg, op, originalErr)

	if err.Kind() != errorlocal.ValidationErrStr {
		t.Errorf(asserterrmsg.BuildAssertErrorMessage("Error Kind", errorlocal.ValidationErrStr, err.Kind()))
	}

	code, _ := httpstatuslocal.BadRequest()
	if err.StatusCode() != code {
		t.Errorf(asserterrmsg.BuildAssertErrorMessage("Error StatusCode", code, err.StatusCode()))
	}

	if err.Error() != msg {
		t.Errorf(asserterrmsg.BuildAssertErrorMessage("Error", msg, err.Error()))
	}

	if err.OriginalErr().Error() != originalErr.Error() {
		t.Errorf(asserterrmsg.BuildAssertErrorMessage("OriginalErr", originalErr, err.OriginalErr()))
	}

	if err.Trace().InnerMessage != innerMsg {
		t.Errorf(asserterrmsg.BuildAssertErrorMessage("Trace InnerMessage", innerMsg, err.Trace().InnerMessage))
	}
}

func TestAuthErr(t *testing.T) {
	op := "TestAuthErr"
	msg := "Testing Auth Error"
	innerMsg := "Inner Testing Auth Error"
	originalErr := errors.New(innerMsg)

	err := errorlocal.AuthErr(msg, op, originalErr)

	if err.Kind() != errorlocal.AuthErrStr {
		t.Errorf(asserterrmsg.BuildAssertErrorMessage("Error Kind", errorlocal.AuthErrStr, err.Kind()))
	}
}

func TestForbiddenErr(t *testing.T) {
	op := "TestForbiddenErr"
	msg := "Testing ForbiddenErr Error"
	innerMsg := "Inner Testing ForbiddenErr Error"
	originalErr := errors.New(innerMsg)

	err := errorlocal.ForbiddenErr(msg, op, originalErr)

	if err.Kind() != errorlocal.ForbiddenErrStr {
		t.Errorf(asserterrmsg.BuildAssertErrorMessage("Error Kind", errorlocal.ForbiddenErrStr, err.Kind()))
	}
}

func TestNotFoundErr(t *testing.T) {
	op := "TestNotFoundErr"
	msg := "Testing NotFoundErr Error"
	innerMsg := "Inner Testing NotFoundErr Error"
	originalErr := errors.New(innerMsg)

	err := errorlocal.NotFoundErr(msg, op, originalErr)

	if err.Kind() != errorlocal.NotFoundErrStr {
		t.Errorf(asserterrmsg.BuildAssertErrorMessage("Error Kind", errorlocal.NotFoundErrStr, err.Kind()))
	}
}

func TestResorceLockedErr(t *testing.T) {
	op := "TestResorceLockedErr"
	msg := "Testing ResorceLockedErr Error"
	innerMsg := "Inner Testing ResorceLockedErr Error"
	originalErr := errors.New(innerMsg)

	err := errorlocal.ResorceLockedErr(msg, op, originalErr)

	if err.Kind() != errorlocal.ResorceLockedErrStr {
		t.Errorf(asserterrmsg.BuildAssertErrorMessage("Error Kind", errorlocal.ResorceLockedErrStr, err.Kind()))
	}
}

func TestUnknownErr(t *testing.T) {
	op := "TestUnknownErr"
	msg := "Testing UnknownErr Error"
	innerMsg := "Inner Testing UnknownErr Error"
	originalErr := errors.New(innerMsg)

	err := errorlocal.UnknownErr(msg, op, originalErr)

	if err.Kind() != errorlocal.UnknownErrStr {
		t.Errorf(asserterrmsg.BuildAssertErrorMessage("Error Kind", errorlocal.UnknownErrStr, err.Kind()))
	}
}

func TestNotImplementedErr(t *testing.T) {
	op := "TestNotImplementedErr"
	err := errorlocal.NotImplementedErr(op)

	if err.Kind() != errorlocal.NotImplementedErrStr {
		t.Errorf(asserterrmsg.BuildAssertErrorMessage("Error Kind", errorlocal.NotImplementedErrStr, err.Kind()))
	}
}

func TestDeprecated(t *testing.T) {
	op := "TestDeprecated"
	err := errorlocal.Deprecated(op)

	if err.Kind() != errorlocal.DeprecatedErrStr {
		t.Errorf(asserterrmsg.BuildAssertErrorMessage("Error Kind", errorlocal.DeprecatedErrStr, err.Kind()))
	}
}

func TestCatchErr(t *testing.T) {
	op := "TestCatchErr"
	msg := "Testing CatchErr Error"
	innerMsg := "Inner Testing CatchErr Error"
	originalErr := errors.New(innerMsg)

	err := errorlocal.CatchErr(errorlocal.UnknownErr(msg, op, originalErr), op)

	if err.Trace().InnerMessage != innerMsg {
		t.Errorf(asserterrmsg.BuildAssertErrorMessage("Error Trace Message", msg, err.Trace().InnerMessage))
	}
}
