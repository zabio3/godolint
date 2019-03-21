package rules

//import (
//	"testing"
//)
//
//func TestValidateDL3012(t *testing.T) {
//	cases := []struct {
//		dockerfileStr string
//		file          string
//		expectedRst   []string
//		expectedErr   error
//	}{
//		{
//			dockerfileStr: `FROM busybox
//MAINTAINER zabio3
//`,
//			file:        "DL3012_Dockerfile",
//			expectedRst: nil,
//			expectedErr: nil,
//		},
//	}
//
//	for i, tc := range cases {
//		rst, err := parseDockerfile(tc.dockerfileStr)
//		if err != nil {
//			t.Errorf("#%d parse error %s", i, tc.dockerfileStr)
//		}
//
//		gotRst, gotErr := validateDL3012(rst.AST, tc.file)
//		if !sliceEq(gotRst, tc.expectedRst) {
//			t.Errorf("#%d results deep equal has returned: want %s, got %s", i, tc.expectedRst, gotRst)
//		}
//
//		if gotErr != tc.expectedErr {
//			t.Errorf("#%d error has returned: want %s, got %s", i, tc.expectedErr, gotErr)
//		}
//		cleanup(t)
//	}
//}
