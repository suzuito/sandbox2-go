package factoryerror

import "errors"

var ErrNoMatchedFetcherID = errors.New("NoMatchedFetcherID")
var ErrNoMatchedParserID = errors.New("NoMatchedParserID")
var ErrNoMatchedPublisherID = errors.New("NoMatchedPublisherID")
var ErrNoMatchedNotifierID = errors.New("NoMatchedNotifierID")
