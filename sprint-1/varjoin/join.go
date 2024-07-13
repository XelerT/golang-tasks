//go:build !solution

package varjoin

func Join(sep string, args ...string) string {
	if len(args) == 0 {
		return "" 
	}

	var res string = args[0]
	for _, str := range args[1:] {
		res += sep + str
	} 

	return res 
}
