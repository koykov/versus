package inspector2

import "github.com/koykov/versus/inspector2/types"

var obj = types.T{
	L1: &types.L1{
		L2: &types.L2{
			L3: &types.L3{
				S: "foobar",
				I: 123456,
				F: 3.1415,
			},
		},
	},
}
