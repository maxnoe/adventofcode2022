package day13

import (
	"reflect"
	"testing"
)


var test_input = `
[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]
`

var test_pairs = []PacketPair {
	{
		[]interface{} {1.0, 1.0, 3.0, 1.0, 1.0,},
		[]interface{} {1.0, 1.0, 5.0, 1.0, 1.0,},
	},
	{
		[]interface{} {[]interface{}{1.0}, []interface{}{2.0, 3.0, 4.0}},
		[]interface{} {[]interface{}{1.0}, 4.0},
	},
	{
		[]interface{} {9.0},
		[]interface{} {[]interface{}{8.0, 7.0, 6.0}},
	},
	{
		[]interface{} {[]interface{} {4.0, 4.0}, 4.0, 4.0},
		[]interface{} {[]interface{} {4.0, 4.0}, 4.0, 4.0, 4.0},
	},
	{
		[]interface{} {7.0, 7.0, 7.0, 7.0},
		[]interface{} {7.0, 7.0, 7.0},
	},
	{
		[]interface{} {},
		[]interface{} {3.0},
	},
	{
		[]interface{} {[]interface{}{[] interface{}{}}},
		[]interface{} {[]interface{}{}},
	},
	{
		[]interface{} {1.0, []interface{}{2.0, []interface{}{3.0, []interface{}{4.0, []interface{} {5.0, 6.0, 7.0}}}}, 8.0, 9.0},
		[]interface{} {1.0, []interface{}{2.0, []interface{}{3.0, []interface{}{4.0, []interface{} {5.0, 6.0, 0.0}}}}, 8.0, 9.0},
	},
}


func TestParseInput(t *testing.T) {
	pairs := ParseInput(test_input)
	for i, pair := range test_pairs {
		if !reflect.DeepEqual(pair, pairs[i]) {
			t.Errorf("Parsing pair %d failed: \n%v\nexpected:\n%v", i, pairs[i], pair)
		}
	}
}
