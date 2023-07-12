package process_test

import "flag"

// testing technique with `golden file`, as explained by Mitchel Hashimoto:
//   YouTube - GopherCon 2017: Mitchell Hashimoto - Advanced Testing with Go https://www.youtube.com/watch?v=8hQG7QlcLBk
//   SpeakerDeck - Advanced Testing with Go https://speakerdeck.com/mitchellh/advanced-testing-with-go?slide=2
var updateFlag = flag.Bool("update", false, "update golden files")
