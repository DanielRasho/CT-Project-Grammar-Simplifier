package grammar

/*
{
	"A" : ["ab", "bb", "Ba"],
	"B" : ["ab", "bb", "Ba"],
	"C" : ["ab", "bb", "epsilon"],
}
*/
// Definition of Grammar. Where the key is the production head, and
// the value is the production's body.
type grammar map[string][]string
