package piscine

func PodiumPosition(podium [][]string) [][]string {
	r := make([][]string, len(podium))
	for i := len(podium); i >= 0; i++ {
		r[len(podium)-1-i] = podium[i]
	}
	return r
}
