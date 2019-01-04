package leetcode

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// func isEqual(a,b []interface) bool {
// 	if a == nil && b == nil {
//         return true;
//     }

//     if a == nil || b == nil {
//         return false;
//     }

//     if len(a) != len(b) {
//         return false
//     }

//     for i := range a {
//         if a[i] != b[i] {
//             return false
//         }
//     }

//     return true
// }
